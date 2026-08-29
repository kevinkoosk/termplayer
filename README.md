# TermPlayer 0.3.0

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
- Portrait mode!!
- Silent (no sound yet)

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

Managed to get Portrait Mode working. It involved changes in the decoder, renderer, and player.

## About this app

Copyright Kevin Koo Seng Kiat (C) 2026 with coding assistance from Microsoft Copilot and Gemini (v. 0.3.0 decoder).

## Licence

Licensed under MIT