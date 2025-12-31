# 🚀 Development Cheat Sheet

## 🛠 Project Commands
- **Run Server:** `go run main.go`
- **Revert Uncommitted Changes:** `git restore .`
- **Clean New Files:** `git clean -fd`
- **Check Formatter Status:** `:ConformInfo`

## ⌨️ Neovim Navigation (LSP)
- `gd` : **Go to Definition** (Jump to where a function/variable is defined)
- `gr` : **Go to References** (See everywhere a function is being used)
- `K`  : **Hover Documentation** (See function signature/types)
- `<leader>rn` : **Rename** (Rename variable across the whole project)
- `ca` : **Code Action** (Quick fixes for errors)

## 🔍 Search & Completion
- `CTRL + n` : **Next** (Move down in the autocomplete/suggestion list)
- `CTRL + p` : **Previous** (Move up in the autocomplete/suggestion list)
- `Enter`    : **Confirm** selection
- `<leader>ff` : **Find Files** (Telescope)
- `<leader>fg` : **Live Grey** (Search text inside files)

## 📂 File Management
- `CTRL + h/j/k/l` : Move between split windows
- `:w` : Save file (Triggers Prettier formatter)
- `:q` : Quit
