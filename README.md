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

## Scaling

You can choose to scale it between 25% to 125%. 

```
termplayer video.mp4 --scale 60
```
This scales a video playback to 60%.

You can also do it for a picture:

```
termplayer image.jpg --scale 50
```
The image is displayed at 50% scale.

If you put in an argument above 125, it's set at 125. If you put in a number 25, it's set at 25.


## Compiling

1. Ensure Go is installed.

2. At the command line, enter:

```
go build -o termplayer.exe ./cmd/termplayer
```

## Notes

This version was focused on the auto-resizing capability. Surprisingly, we found that the app could already resize its output when the terminal window is resized. But on reflection, it was to be expected - we calculate the window size to render. 

So we shipped another feature called scaling. Now you can playback videos or display images at a scale between 25% to 125%. But I don't recommend going beyong the 100% for horizontal videos and images. 

## About this app

Copyright Kevin Koo Seng Kiat (C) 2026 with coding assistance from Microsoft Copilot and Gemini (v. 0.3.0 decoder).

## Licence

Licensed under MIT