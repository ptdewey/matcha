# Matcha

Matcha is a beautiful and featureful note taking helper program built for the terminal.
Matcha lets you quickly search and edit your notes, create new ones from templates, and more.

Matcha is built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

<!-- TODO: screenshot/gif -->

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
go build
# add to path or run with ./matcha
```

## Configuration

Matcha looks for one of the following configuration files in your home directory:
- `matcha.json`
- `.matcha.json`
- `.matcharc`

TODO: explanation

Example configuration:
```json
{
    "noteSources": [
        "~/notes"
    ],
    "defaultExt": ".md",
    "useTemplate": true,
    "templateDir": "~/notes/.templates"
}
```
