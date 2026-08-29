# TermPlayer 0.2.4

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

In this update, we optimize the rendering engine. Rather than 1 call / write for every character, it's formed entirely as a string in the app, then written.

## About this app

Copyright Kevin Koo Seng Kiat (C) 2026 with coding assistance from Microsoft Copilot.

## Licence

Licensed under MIT