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

## Development

Pre-requisites:

- Go
- [raylib-go](https://github.com/gen2brain/raylib-go?tab=readme-ov-file#raylib-go) requirements

```bash
make run # run the project
make build # builds the project
make build-w # builds the project for windows
```

Builds go to `/dist`.

## Roadmap

- [ ] Test on other OSs
  - [x] Linux
    - [x] x
    - [ ] Wayland
  - [x] Windows
- [ ] Work with multi-monitor
