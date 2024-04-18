# qss

`qss` is a utility for quickly capturing and automatically saving portions of the screen.

## Usage

Run the executable to see a brief help text. Select an area to capture by clicking and dragging, or capture the entire screen with a right click.

By default, screenshots are saved to `$HOME/Pictures`. You can configure a different path via the config file.

## Config file

Create a `.qss.conf` file in your default config directory (`$HOME/.config` or `C:\Users\<user>\AppData\Roaming`) to customize settings.

```
# Help text font size (default: 24)
font-size=24

# Toggle help text at startup (default: true)
show-help=true

# Custom screenshot save path (default: $HOME/Pictures)
save-path=/home/user/
```

# Roadmap

- [x] Show instructions (but hide during snapshotting)
- [x] Config file
  - [x] file path
  - [x] show help
  - [x] font size
- [x] Keys
  - [x] Right click: print whole screen
  - [x] Right click while dragging: cancel selection
  - [x] Esc: quit
  - [x] Esc/right click (during selection): cancel current selection, don't quit app
- [x] Show rect size while dragging
- [ ] Test on other OSs
- [x] Use log pkg instead of std.Fmt
- [ ] Run on multiple monitors
- [ ] Builds and installer/instructions
  - [ ] CD
  - [ ] Linux
    - [ ] Package managers?
  - [ ] Windows
  - [ ] Macos
