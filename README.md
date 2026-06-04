<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║                    🦊 EVILGINX3 — TELEGRAM EDITION                        ║
║         Man-in-the-Middle Attack Framework with 2FA Bypass                ║
║                Real-Time Telegram Alerts & Web Dashboard                  ║
║                                                                            ║
║                         by @officialmonsterz                               ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

<div align="center">

<a href="https://github.com/officialmonsterz/evilginx2">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/logo.png" alt="Evilginx3 Logo" width="220" style="filter: drop-shadow(0 4px 8px rgba(0,0,0,0.3));">
</a>

# 🦊 Evilginx3 — Telegram Edition

### *Man-in-the-Middle Attack Framework with 2FA Bypass & Real-Time Webdashboard & Telegram Alerts*

<br>

[![Telegram](https://img.shields.io/badge/Telegram-@officialmonsterz-26A5E4?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/officialmonsterz)
[![Email](https://img.shields.io/badge/Email-shapads@tutamail.com-D14836?style=for-the-badge&logo=mail.ru&logoColor=white)](mailto:shapads@tutamail.com)
[![GitHub](https://img.shields.io/badge/GitHub-officialmonsterz-181717?style=for-the-badge&logo=github)](https://github.com/officialmonsterz/evilginx2)
<br>
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue?style=flat-square)](https://github.com/officialmonsterz/evilginx2/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-3.3.0-brightgreen?style=flat-square)]()
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)]()

<br>
<sub><b>⚠️ FOR AUTHORIZED PENETRATION TESTING ONLY</b></sub>

</div>

---

<br>

# 📋 Table of Contents

<div style="column-count: 2; column-gap: 40px;">

1. [🧠 What Is Evilginx3?](#-what-is-evilginx3)
2. [⚡ Why This Fork?](#-why-this-fork)
3. [📊 Feature Comparison vs Original](#-feature-comparison-vs-original)
4. [🧬 Architecture Overview](#-architecture-overview)
5. [🆕 What's New — Complete File Inventory](#-whats-new--complete-file-inventory)
6. [💎 Core Features](#-core-features)
7. [📱 Telegram Integration Deep Dive](#-telegram-integration-deep-dive)
8. [📊 Web Dashboard](#-web-dashboard)
9. [🐳 Docker Support](#-docker-support)
10. [🚀 Quick Start](#-quick-start)
11. [🧪 Pro Tips & Tricks](#-pro-tips--tricks)
12. [📸 Screenshots](#-screenshots)
13. [⚖️ Disclaimer](#️-disclaimer)
14. [👏 Credits & Support](#-credits--support)

</div>

---

<br>

# 🧠 What Is Evilginx3?

**Evilginx3** is a **man-in-the-middle (MITM) attack framework** — a reverse proxy that sits between a victim and a real website (e.g., Office 365, Google, LinkedIn). It captures **login credentials** along with **session cookies**, effectively bypassing **2-factor authentication (2FA/MFA)** protection.

This is the successor to [Evilginx](https://github.com/kgretzky/evilginx) (2017), which used a custom nginx-based proxy. The present version is **fully written in Go** as a standalone application with its own HTTP and DNS servers — making deployment trivial.

---

┌─────────────────────────────────────────────────────────────────────────────────┐ │ │ │ VICTIM'S BROWSER │ │ (e.g., victim@company.com) │ │ │ │ │ │ https://login.yourdomain.com/abc123 │ │ ▼ │ │ ┌─────────────────────────────────────────────────────────────────────┐ │ │ │ EVILGINX3 PROXY ENGINE │ │ │ │ │ │ │ │ ┌──────────────┐ ┌──────────────┐ ┌──────────────────┐ │ │ │ │ │ HTTP Proxy │ │ URL Rewrite │ │ Content Filter │ │ │ │ │ │ (MITM) │───▶│ (Domain │───▶│ (Replace URLs │ │ │ │ │ │ │ │ Swapping) │ │ in HTML/JS/CSS)│ │ │ │ │ └──────────────┘ └──────────────┘ └──────────────────┘ │ │ │ │ │ │ │ │ ┌────────────────────────────────────────────────────────────┐ │ │ │ │ │ CAPTURE LAYER │ │ │ │ │ │ ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐ │ │ │ │ │ │ │ Username │ │ Password │ │ Session │ │ 2FA/OTP │ │ │ │ │ │ │ │ /Email │ │ │ │ Cookies │ │ Tokens │ │ │ │ │ │ │ └──────────┘ └──────────┘ └──────────┘ └──────────┘ │ │ │ │ │ └────────────────────────────────────────────────────────────┘ │ │ │ └─────────────────────────────────────────────────────────────────────┘ │ │ │ │ │ │ Proxied request (yourdomain.com → real website) │ │ ▼ │ │ ┌─────────────────────────────────────────────────────────────────────┐ │ │ │ REAL WEBSITE (e.g., login.microsoftonline.com) │ │ │ │ ┌──────────────────────────────────────────────────────────────┐ │ │ │ │ │ Login succeeds — victim sees NO ERROR — perfectly normal │ │ │ │ │ │ page loads, dashboard appears, everything looks legitimate │ │ │ │ │ └──────────────────────────────────────────────────────────────┘ │ │ │ └─────────────────────────────────────────────────────────────────────┘ │ │ │ │ ═══════════════════════════════════════════════════════════════════════ │ │ │ │ ▼ CAPTURED DATA DELIVERY CHANNELS ▼ │ │ │ │ ┌─────────────────────────────────┐ ┌─────────────────────────────────┐ │ │ │ 📱 TELEGRAM │ │ 📊 WEB DASHBOARD │ │ │ │ │ │ │ │ │ │ ✨ Session Information ✨ │ │ ┌────┬──────────┬──────────┐ │ │ │ │ 👤 user@corp.com │ │ │ ID │ Username │ Password │ │ │ │ │ 🔑 SuperSecret123! │ │ ├────┼──────────┼──────────┤ │ │ │ │ 📎 tokens.txt attached │ │ │ 1 │ user@... │ Pass123! │ │ │ │ │ 🔄 Auto-updating message │ │ └────┴──────────┴──────────┘ │ │ │ └─────────────────────────────────┘ └─────────────────────────────────┘ │ │ │ │ ┌─────────────────────────────────────────────────────────────────────┐ │ │ │ 💾 BUNTDB EMBEDDED DATABASE (Zero Config, No SQL Server Needed) │ │ │ │ └─────────────────────────────────────────────────────────────────┘ │ │ │ └────────────────────────────────────────────────────────────────────

<br>

## 🔁 How It Works — Visual Flow
