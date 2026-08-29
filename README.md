# TermPlayer 0.3.2

Video player for terminal. Vibe coded, free, and open source.

Also displays images.

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
- Portrait mode
- Also landscape mode
- Also displays images (jpg, png, gif)
- Player auto-resizes as you resize the window
- Silent (no sound yet)

## Usage

To view a video at the terminal, type:

```
termplayer video.mp4
```

During playback, use "Q" to quit, and Spacebar to pause.

If you want to view an image instead, type:

```
termplayer image.jpg
```



## Compiling

1. Ensure Go is installed.

2. At the command line, enter:

```
go build -o termplayer.exe ./cmd/termplayer
```

## Notes

This version was focused on the auto-resizing capability. Surprisingly, we found that the app could already resize its output when the terminal window is resized. But on reflection, it was to be expected - we calculate the window size to render. The code for 0.3.2 is exactly the same as 0.3.1. But a milestone has been achieved.

## About this app

Copyright Kevin Koo Seng Kiat (C) 2026 with coding assistance from Microsoft Copilot and Gemini (v. 0.3.0 decoder).

## Licence

Licensed under MIT