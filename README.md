# TermPlayer 0.2.5

Video player for terminal.

Vibe coded, free, and open source.

![screenshot of Termplayer](screenshot.png)

![screenshot of Termplayer](screenshot2.png)


## Feature set

- RGB24 decoding
- Terminal truecolor output
- Half-block renderer
- Reusable image buffers
- Reusable frame buffers
- FPS counter
- Drop counter
- Single-key controls
- Pause support
- Hidden cursor
- Renderer optimization
- Published GitHub repository

## Usage

At the terminal, type:

```
termplayer video.mp4
```

During playback, use "Q" to quit, and Spacebar to pause.


## Compiling

1. Ensure Go is installed.

2. At the command line, enter:

```
go build -o termplayer.exe ./cmd/termplayer
```

## Notes

In the last version (0.2.4) we used strings.Builder which helped optimize the rendering and made the player smoother. In this version (0.2.5), we drop frames to enhance the user experience. 


## About this app

Copyright Kevin Koo Seng Kiat (C) 2026 with coding assistance from Microsoft Copilot.

## Licence

Licensed under MIT