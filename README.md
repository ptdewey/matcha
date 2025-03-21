# Matcha

Matcha is a beautiful and featureful note taking helper program built for the terminal.
Matcha lets you quickly search and edit your notes, create new ones from templates, and more.

Matcha is built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

Create Mode:
![Create mode gif](https://vhs.charm.sh/vhs-4tSj59aPfirr75FPTkCEDk.gif)

Search Mode:
![Search mode gif](https://vhs.charm.sh/vhs-4kSCJXt5B2VcZC14XOeFdu.gif)


## Installation

With Go install:
```bash
go install github.com/ptdewey/matcha@latest
```

(Ensure your Go binary installation location is in your PATH)

From source:
```bash
git clone https://github.com/ptdewey/matcha.git
cd matcha
go mod tidy
go build
# add to path or run with ./matcha
```

## Usage

```bash
# run matcha
matcha
```

TODO: explanation and usage of modes

## Configuration

Matcha looks for one of the following configuration files in your home directory:
- `matcha.toml`
- `.matcha.toml`
- `.matcharc`

Currently, matcha provides the following configuration options:
| Option | Type | Description |
|--------|------|-------------|
| noteSources | string array | Note directories to be used by Matcha |
| defaultExt | string | Default file extension for new notes used when none is specified |
| useTemplate | boolean | Whether or not to attempt to create new notes with template by default (experimental/WIP) |
| templateDir | string | Path to template directories for creating new notes (experimental/WIP) |
| noteExts | string array | Note file extensions to be used by Matcha |

Example configuration:
```toml
noteSources = [
    "~/notes",
]
defaultExt = ".md"
useTemplate = false
templateDir = ""
noteExts = [
    ".md",
    ".txt",
    ".tex",
    ".typ",
]
```

