# Home Deck Alpha

Home Deck Alpha is a private project on alpha develpoment to provide control over designated IoT devices.
It manages connections with devices using a pool of websockets, wich each device has a dedicated communication channel.

---

## 🔧 Quick Start Guide

This project uses `make` commands to simplify building, installing, and managing the `lock-server` as a systemd service.

### 📦 Prerequisites

- Go installed (1.22+ recommended)
- `make` installed
- `systemd` for daemon management (Linux only)
- Sudo privileges for install and service setup

---

### 🛠️ Build the Project

Compiles the Go project and outputs the binary to the `build/` directory.

```bash
make build
```

> Output: `build/server`

---

### 🧹 Clean Build Artifacts

Removes the `build/` directory and compiled binary.

```bash
make clean
```

---

### 📥 Install the Binary

Copies the compiled binary to `/usr/local/bin` with the name `lock-server`.

```bash
make install
```

> This step requires `sudo` since it writes to a system directory.

---

### ⚙️ Install as a systemd Service

Cleans, builds, installs, and sets up the `lock-server` as a systemd service.

```bash
make daemon
```

This does the following:

1. Cleans and rebuilds the project
2. Installs the binary
3. Copies `lock-server.service` to `/etc/systemd/system/`
4. Reloads the systemd daemon
5. Enables the service on boot

> After this, you can manage the service using systemd:

```bash
sudo systemctl start lock-server
sudo systemctl status lock-server
```

---

### 🧽 Remove the systemd Service

Stops and removes the systemd unit file for the `lock-server`.

```bash
make daemon-rm
```

This will:

1. Remove the `lock-server.service` file from `/etc/systemd/system/`
2. Reload the systemd daemon
