![Gotion Banner](https://i.ibb.co.com/dwFDBKYx/gotion-banner-cmp.png)

# Gotion – Local-First CLI Note Manager

**Gotion** is a zero-dependency, cross-platform CLI that turns your terminal into a fast, offline note-taking workspace.  
Create, search, edit, tag, and archive **Markdown notes** stored directly on your machine—no servers, no cloud. All content lives as plain `.md` files in a folder you control, so you can sync, backup, or version-control them like any other project asset.

---

## ✨ Key Features

- **Blazing Fast** – sub-second cold start, millisecond queries
- **Offline & Private** – no cloud, no sign-ups, fully local storage
- **Portable Folder** – one directory of `.md` files; move, backup, or git-track it
- **Full-text Search** – fuzzy-find across titles, bodies, and tags

---

## 🚀 Installation & Run

| Target           | Command                      | Result                                                             |
| ---------------- | ---------------------------- | ------------------------------------------------------------------ |
| **Quick Try**    | `make build`                 | Produces `./gotion` (or `gotion.exe`) for immediate local testing. |
| **Run the App**  | `make run`                   | Builds (if needed) and starts Gotion in one step.                  |
| **Add Features** | `make dev feat=new-search`   | Scaffolds boilerplate for a new feature branch.                    |
| **Release**      | `git tag -a v1.2.0 -m "..."` | Creates annotated tag; triggers GoReleaser via GitHub Actions.     |
| **Install**      | `make install`               | Copies binary to `$GOPATH/bin` (or `%GOPATH%\bin` on Windows).     |
| **Run Anywhere** | `gotion`                     | Available system-wide after install; auto-detects `$HOME/Notes`.   |

---

## 🧪 Clone & Extend

1. **Install Go (if you haven’t already)**
   ```bash
   # macOS (Homebrew)
   brew install go

   # Ubuntu / Debian
   sudo apt update && sudo apt install golang-go

   # Arch Linux
   sudo pacman -S go

   # Windows (Scoop)
   scoop install go
   ```

2. **Clone the repo**
   ```bash
   git clone https://github.com/noyonalways/gotion.git
   cd gotion
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Hack on a new feature**
   ```bash
   make dev feat=my-awesome-idea
   # code, test, commit…
   ```

5. **Push & open a PR**
   ```bash
   git push origin feat/my-awesome-idea
   # open pull-request on GitHub
   ```

---

## 📦 Build & Release Pipeline

We follow a trunk-based workflow with automated releases:

1. **Stage changes**

   ```bash
   git add .
   git commit -m "feat: add encrypted note support"
   git push origin main
   ```

2. **Tag release**

   ```bash
   git tag -a v1.3.0 -m "release: encrypted notes, arm64 builds, zsh completion"
   git push origin v1.3.0
   ```

3. **CI takes over**
   - GoReleaser cross-compiles binaries for 10 targets
   - Checksums & SBOMs are generated and attached
   - Homebrew formula is auto-bumped
   - Docker images (`ghcr.io/your-org/gotion:v1.3.0`) are published
