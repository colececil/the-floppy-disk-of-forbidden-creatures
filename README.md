# The Floppy Disk of Forbidden Creatures

_**This began as an entry for Ludum Dare 55: "Summoning", by Cole Cecil. The initial version (v1.0.0) was created in under 72 hours.**_

What awaits you when you explore the contents of the Floppy Disk of Forbidden Creatures? A summoning, to be sure, but what will you summon? And what consequences will it bring?

_(**Disclaimer:** This game connects to the internet to enhance the gameplay experience. While no personal information is required or intentionally collected during the game, any information you enter may be transmitted over the internet and could be used as AI training data. Therefore, it is recommended that you do not enter any personal information.)_

## Instructions for Running the Game

This is a game you run and play on the command line. After [downloading](https://github.com/colececil/the-floppy-disk-of-forbidden-creatures/releases/latest) the zip file that matches your operating system and CPU architecture, unzip it. Then, on the command line, navigate to the unzipped directory and run the "summon" program.

Here are more specific command line instructions for each operating system:

- **Windows:**
  1. Open a terminal using either the "Command Prompt" app or "PowerShell" app.
  2. Copy the path to the directory the game was unzipped to. In the terminal, type `cd`, then a space, then paste the path to the directory and press enter. For example, if the game was unzipped to `C:\games\summon`, the command would be `cd C:\games\summon`.
  3. To run the game, type `.\summon.exe` and press enter.
- **macOS and Linux:**
  1. Open a terminal using the "Terminal" app.
  2. Copy the path to the directory the game was unzipped to. In the terminal, type `cd`, then a space, then paste the path to the directory and press enter. For example, if the game was unzipped to `/home/summon`, the command would be `cd /home/summon`.
  3. To run the game, type `./summon` and press enter. (If you get a permission error, make sure the `summon` file has executable permissions. You can add this by running the command `chmod +x summon`.)

## Instructions for Building the Game

### Prerequisites

- [Go](https://go.dev/) v1.22
- [Make](https://www.gnu.org/software/make/)

### Building the Game

To build the game for all supported operating systems and architectures, run the following command from the project root:

```
make build-all
```

To build the game for a single operating system and architecture, run the following command from the project root, where `<osName>` is one of `windows`, `darwin` or `linux`, and `<architectureName>` is one of `amd64` or `arm64`:

```
make build OS=<osName> ARCH=<architectureName>
```
  
## Attributions

This game was written in the [Go](https://go.dev/) programming language.

The following libraries were used:

- [Beep](https://github.com/gopxl/beep)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [go-openai](https://github.com/sashabaranov/go-openai)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [x/ansi](https://github.com/charmbracelet/x/tree/main/ansi)

The following audio files were used in the creation of sound effects:

- "[floppy drive on old pc](https://pixabay.com/sound-effects/floppy-drive-on-old-pc-52014/)", by Pixabay
- "[The Sound of dial-up Internet](https://freesound.org/s/546450/)", by wtermini