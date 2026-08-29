# TermPlayer 0.2.2

Video player for terminal.

Vibe coded, free, and open source.

![screenshot of Termplayer](screenshot.png)

![screenshot of Termplayer](screenshot2.png)


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

FFMPEG continues to process the video in the backend, so the video shouldn't be paused for too long. (Even though it's possible to pause, the buffer gets clogged up eventually.)

## About this app

Copyright Kevin Koo Seng Kiat (C) 2026 with coding assistance from Microsoft Copilot.

## Licence

Licensed under MIT