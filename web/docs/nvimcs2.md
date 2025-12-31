# Neovim Keybindings Cheatsheet

## Leader Key
- **Leader**: `Space`
- **Local Leader**: `Space`

## LSP (Language Server Protocol)
| Key | Action | Description |
|-----|--------|-------------|
| `K` | Hover Documentation | Show documentation for item under cursor |
| `gd` | Go to Definition | Jump to definition of symbol under cursor |
| `<leader>ca` | Code Action | Show available code actions |

## File Explorer (nvim-tree)
| Key | Action | Description |
|-----|--------|-------------|
| `Ctrl+n` | Toggle File Tree | Open/close the file explorer |

## Telescope (Fuzzy Finder)
| Key | Action | Description |
|-----|--------|-------------|
| `<leader>ff` | Find Files | Search for files by name |
| `<leader>fg` | Live Grep | Search for text inside files |
| `<leader>fb` | Buffers | Browse open buffers |
| `<leader>fh` | Help Tags | Search through help documentation |

## Autocompletion (nvim-cmp)
| Key | Action | Description |
|-----|--------|-------------|
| `Ctrl+b` | Scroll Docs Up | Scroll documentation window up |
| `Ctrl+f` | Scroll Docs Down | Scroll documentation window down |
| `Ctrl+Space` | Trigger Completion | Manually trigger completion menu |
| `Enter` | Confirm | Accept the selected completion |

## Custom Commands
| Command | Description |
|---------|-------------|
| `:Format` | Format current buffer or selection using Conform |

## Notes
- Format on save is enabled for JavaScript, CSS, HTML, and Go files
- Go files auto-format on save using LSP
- Mouse support is enabled
- System clipboard integration is configured for WSL
