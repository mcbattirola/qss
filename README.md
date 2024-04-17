# qss

`qss` is a utility for quickly capturing and automatically saving portions of the screen.

## Config file

```
# font-size sets the help font size, defaults to 24
font-size=12
# show-help sets if the help text show up when the app starts, defaults to true
show-help=true
# save-path is the path where screenshots are saved, defaults to $HOME/Pictures
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
- [ ] Use slog instead of std.Fmt
- [ ] Run on multiple monitors
- [ ] Builds and installer/instructions
  - [ ] CD
  - [ ] Linux
    - [ ] Package managers?
  - [ ] Windows
  - [ ] Macos
