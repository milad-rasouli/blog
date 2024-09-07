# tmux Quick Shortcut Guide

## General Commands

| Shortcut                 | Description                                      |
|--------------------------|--------------------------------------------------|
| `Ctrl + b`               | Default prefix key for tmux commands             |
| `Ctrl + b + ?`           | List all tmux shortcuts                          |
| `Ctrl + b + d`           | Detach the current session                       |
| `Ctrl + b + s`           | List sessions and windows                        |
| `Ctrl + b + :`           | Enter tmux command prompt                        |
| `Ctrl + b + : then setw synchronize-panes on` | Enable pane synchronization (type in all panes at once) |
| `Ctrl + b + : then setw synchronize-panes off` | Disable pane synchronization |
| `tmux ls`                | List all tmux sessions                           |
| `tmux attach -t <session>` | Attach to a session by name                     |
| `tmux new -s <name>`     | Create a new session with a name                 |
| `tmux kill-session -t <session>` | Kill a specific session by ID            |

## Window Management

| Shortcut                 | Description                                      |
|--------------------------|--------------------------------------------------|
| `Ctrl + b + c`           | Create a new window                              |
| `Ctrl + b + w`           | List all windows                                 |
| `Ctrl + b + ,`           | Rename the current window                        |
| `Ctrl + b + &`           | Kill the current window                          |
| `Ctrl + b + p`           | Switch to the previous window                    |
| `Ctrl + b + n`           | Switch to the next window                        |
| `Ctrl + b + <number>`    | Switch to window by number                       |

## Pane Management

| Shortcut                 | Description                                      |
|--------------------------|--------------------------------------------------|
| `Ctrl + b + "`           | Split the current window horizontally            |
| `Ctrl + b + %`           | Split the current window vertically              |
| `Ctrl + b + o`           | Switch between panes                             |
| `Ctrl + b + x`           | Close the current pane                           |
| `Ctrl + b + z`           | Toggle zoom for the current pane                 |
| `Ctrl + b + q`           | Show pane numbers for quick pane switching       |
| `Ctrl + b + ;`           | Toggle between last two active panes             |

## Pane Resizing

| Shortcut                 | Description                                      |
|--------------------------|--------------------------------------------------|
| `Ctrl + b + <arrow key>`  | Resize pane in the direction of the arrow (left, right, up, down) |

## Copy Mode & Scrolling

| Shortcut                 | Description                                      |
|--------------------------|--------------------------------------------------|
| `Ctrl + b + [`           | Enter copy mode                                  |
| `q`                      | Exit copy mode                                   |
| `Space`                  | Start text selection in copy mode                |
| `Enter`                  | Copy selected text                               |
| `Ctrl + b + ]`           | Paste copied text                                |
| `Ctrl + b + PgUp`        | Scroll up in a window                            |

## Miscellaneous

| Shortcut                 | Description                                      |
|--------------------------|--------------------------------------------------|
| `Ctrl + b + t`           | Show the current time                            |
| `Ctrl + b + m`           | Mark the current pane                            |

---

**Tip**: Press `Ctrl + b` and release it, then press the next key combination to trigger a tmux command.

