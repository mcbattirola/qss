# Printit

TODO

- [ ] Show instructions (but hide during snapshotting)
- [ ] Config file
- [ ] Keys
  - [ ] Right click: print whole screen
  - [ ] Right click while dragging: cancel selection
  - [x] Esc: quit
  - [ ] Esc/right click (during selection): cancel current selection, don't quit app
- [ ] Show rect size while dragging
- [ ] New name
- [ ] File name
- [ ] Test on other OSs
- [ ] Use slog instead of std.Fmt
- [ ] CLI interface
- [ ] Run on multiple monitors

## CLI

```bash
# simply invoke the binary to manually select the print area
qss

# -c or --config to set a specific config file
qss -c ./my-config-file.conf

# use -h or --help for detailed help
qss -h
```