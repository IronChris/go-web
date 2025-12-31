# Neovim Cheatsheet

## Basic Navigation
| Command | Description |
|---------|-------------|
| `h`, `j`, `k`, `l` | Move left, down, up, right |
| `w` / `W` | Jump forward to start of word/WORD |
| `b` / `B` | Jump backward to start of word/WORD |
| `e` / `E` | Jump to end of word/WORD |
| `0` | Jump to start of line |
| `$` | Jump to end of line |
| `gg` | Go to first line |
| `G` | Go to last line |
| `:<number>` | Go to specific line |
| `Ctrl+o` | Jump back in history |
| `Ctrl+i` | Jump forward in history |

## File Management
| Command | Description |
|---------|-------------|
| `:w` | Save current file |
| `:w <filename>` | Save as new file |
| `:q` | Quit current window |
| `:q!` | Quit without saving |
| `:wq` or `:x` | Save and quit |
| `:e <file>` | Open file |
| `:e!` | Reload current file |
| `Ctrl+n` | Toggle file explorer (NvimTree) |

## Visual Mode
| Command | Description |
|---------|-------------|
| `v` | Enter visual (character) mode |
| `V` | Enter visual line mode |
| `Ctrl+v` | Enter visual block mode |
| `y` | Yank (copy) selection |
| `d` | Delete (cut) selection |
| `p` | Paste after cursor |
| `P` | Paste before cursor |
| `>` | Indent selection |
| `<` | Unindent selection |
| `=` | Auto-indent selection |

## Editing
| Command | Description |
|---------|-------------|
| `i` | Insert mode before cursor |
| `a` | Insert mode after cursor |
| `I` | Insert at beginning of line |
| `A` | Insert at end of line |
| `o` | Insert new line below |
| `O` | Insert new line above |
| `r` | Replace single character |
| `R` | Replace mode (overwrite) |
| `x` | Delete character under cursor |
| `X` | Delete character before cursor |
| `u` | Undo |
| `Ctrl+r` | Redo |
| `.` | Repeat last command |

## Text Manipulation
| Command | Description |
|---------|-------------|
| `dd` | Delete (cut) current line |
| `yy` | Yank (copy) current line |
| `p` | Paste below cursor |
| `P` | Paste above cursor |
| `J` | Join lines |
| `>>` | Indent line |
| `<<` | Unindent line |
| `~` | Toggle case of character |
| `gu` | Make selection lowercase |
| `gU` | Make selection uppercase |
| `g~` | Toggle case of selection |

## Search & Replace
| Command | Description |
|---------|-------------|
| `/pattern` | Search forward for pattern |
| `?pattern` | Search backward for pattern |
| `n` | Next search match |
| `N` | Previous search match |
| `*` | Search forward for word under cursor |
| `#` | Search backward for word under cursor |
| `:%s/old/new/g` | Replace all occurrences in file |
| `:%s/old/new/gc` | Replace with confirmation |
| `:noh` | Clear search highlighting |

## Window Management
| Command | Description |
|---------|-------------|
| `Ctrl+w s` | Split window horizontally |
| `Ctrl+w v` | Split window vertically |
| `Ctrl+w arrow` | Navigate between windows |
| `Ctrl+w =` | Equalize window sizes |
| `Ctrl+w _` | Maximize window height |
| `Ctrl+w \|` | Maximize window width |
| `Ctrl+w +` | Increase window height |
| `Ctrl+w -` | Decrease window height |
| `Ctrl+w >` | Increase window width |
| `Ctrl+w <` | Decrease window width |
| `Ctrl+w c` | Close current window |
| `Ctrl+w o` | Close other windows |

## Tabs
| Command | Description |
|---------|-------------|
| `:tabnew` | Create new tab |
| `:tabclose` | Close current tab |
| `:tabonly` | Close other tabs |
| `gt` | Next tab |
| `gT` | Previous tab |
| `:<number>gt` | Go to specific tab |

## Plugin-Specific Commands

### Telescope (File Finder)
| Command | Description |
|---------|-------------|
| `<leader>ff` | Find files |
| `<leader>fg` | Live grep (search in files) |
| `<leader>fb` | List open buffers |
| `<leader>fh` | Search help tags |
| `<Esc>` | Exit Telescope |

### NvimTree (File Explorer)
| Command | Description |
|---------|-------------|
| `Ctrl+n` | Toggle file explorer |
| `o` / `<CR>` | Open file/directory |
| `a` | Create file/directory |
| `d` | Delete file/directory |
| `r` | Rename file/directory |
| `y` | Copy file/directory |
| `p` | Paste file/directory |
| `-` | Go up to parent directory |
| `q` | Close NvimTree |

### LSP (Language Server Protocol)
| Command | Description |
|---------|-------------|
| `K` | Show hover information |
| `gd` | Go to definition |
| `<leader>ca` | Show code actions |
| `:Format` | Format current buffer |
| `gr` | Go to references |
| `gi` | Go to implementation |
| `<leader>rn` | Rename symbol |

### Autocompletion (nvim-cmp)
| Command | Description |
|---------|-------------|
| `Ctrl+Space` | Trigger completion |
| `Enter` | Confirm selection |
| `Ctrl+b` | Scroll docs up |
| `Ctrl+f` | Scroll docs down |
| `Tab` / `S-Tab` | Navigate completion items |

## Buffer Management
| Command | Description |
|---------|-------------|
| `:bnext` or `:bn` | Next buffer |
| `:bprev` or `:bp` | Previous buffer |
| `:bd` | Delete current buffer |
| `:ls` | List all buffers |
| `:b <number>` | Switch to buffer by number |
| `:b <name>` | Switch to buffer by name |

## Marks
| Command | Description |
|---------|-------------|
| `m{a-z}` | Set local mark |
| `m{A-Z}` | Set global mark |
| `` `{mark}`` | Jump to exact mark position |
| `'{mark}` | Jump to line of mark |
| `` `.`` | Jump to last change |
| `` `"`` | Jump to last position when file closed |
| `:marks` | List all marks |

## Registers
| Command | Description |
|---------|-------------|
| `"{register}y` | Yank to specific register |
| `"{register}p` | Paste from specific register |
| `:reg` | List all registers |
| `"0` | Last yank register |
| `"+` | System clipboard |
| `"*` | System selection (X11) |

## Macros
| Command | Description |
|---------|-------------|
| `q{register}` | Start recording macro to register |
| `q` | Stop recording |
| `@{register}` | Execute macro |
| `@@` | Repeat last macro |
| `:{number}@{register}` | Execute macro {number} times |

## Useful Vim Commands
| Command | Description |
|---------|-------------|
| `:set paste` | Enable paste mode (prevents auto-indent) |
| `:set nopaste` | Disable paste mode |
| `:set spell` | Enable spell checking |
| `:set nospell` | Disable spell checking |
| `zg` | Add word to spell dictionary |
| `zw` | Mark word as wrong |
| `]s` | Next spelling error |
| `[s` | Previous spelling error |
| `z=` | Show spelling suggestions |

## Formatting
| Command | Description |
|---------|-------------|
| `==` | Auto-indent current line |
| `gg=G` | Auto-indent entire file |
| `:Format` | Format current buffer with conform.nvim |
| `gq` | Format selected text |

## Folding
| Command | Description |
|---------|-------------|
| `zc` | Close fold |
| `zo` | Open fold |
| `zC` | Close all folds |
| `zO` | Open all folds |
| `zm` | Fold more (increase foldlevel) |
| `zr` | Fold less (decrease foldlevel) |
| `za` | Toggle fold |
| `zR` | Open all folds |
| `zM` | Close all folds |

## Custom Keybindings in Your Config
| Command | Description |
|---------|-------------|
| `<leader>ff` | Telescope find files |
| `<leader>fg` | Telescope live grep |
| `<leader>fb` | Telescope buffers |
| `<leader>fh` | Telescope help tags |
| `<leader>ca` | LSP code actions |
| `Ctrl+n` | Toggle NvimTree |

## Tips & Tricks
1. **Leader Key**: Your leader key is set to Space (` `)
2. **Relative Numbers**: Enabled for easier navigation (`10j` moves down 10 lines)
3. **Auto-formatting**: Go files auto-format on save
4. **Clipboard**: System clipboard integration for WSL
5. **File Explorer**: Use `Ctrl+n` to toggle file tree
6. **Fuzzy Finder**: Space+ff for file search, Space+fg for text search

## Common Workflows
1. **Open a file**: `:e filename` or use Telescope (`<leader>ff`)
2. **Search text**: `<leader>fg` then type search term
3. **Navigate code**: `gd` to go to definition, `K` for docs
4. **Format code**: `:Format` or save Go files (auto-format)
5. **Manage files**: `Ctrl+n` for tree, use `a`, `d`, `r` commands

---

**Note**: This config uses Dracula theme and is optimized for WSL. All plugins are managed by lazy.nvim.
