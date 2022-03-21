# itermcolors-to-vscode

Convert your [iTerm2](http://iterm2.com) color scheme into [VS Code](https://code.visualstudio.com) terminal color settings.

## WHY?

Having the same colors for iTerm2 and VS Code terminal is great!

## HOW?

Get your iTerm color preset in .itermcolors format

`iTerm settings > Profiles > Colors > Color Presets > Export`


```bash
go run main.go -f base16-eighties.itermcolors

"workbench.colorCustomizations": {
    "terminal.ansiBlack": "#2d2d2d",
    "terminal.ansiBlue": "#6699cc",
    "terminal.ansiBrightBlack": "#747369",
    "terminal.ansiBrightBlue": "#a09f93",
    "terminal.ansiBrightCyan": "#d27b53",
    "terminal.ansiBrightGreen": "#393939",
    "terminal.ansiBrightMagenta": "#e8e6df",
    "terminal.ansiBrightRed": "#f99157",
    "terminal.ansiBrightWhite": "#f2f0ec",
    "terminal.ansiBrightYellow": "#515151",
    "terminal.ansiCyan": "#66cccc",
    "terminal.ansiGreen": "#99cc99",
    "terminal.ansiMagenta": "#cc99cc",
    "terminal.ansiRed": "#f2777a",
    "terminal.ansiWhite": "#d3d0c8",
    "terminal.ansiYellow": "#ffcc66",
    "terminal.background": "#2d2d2d",
    "terminal.foreground": "#d3d0c8",
    "terminal.selectionBackground": "#515151",
    "terminalCursor.background": "#2d2d2d",
    "terminalCursor.foreground": "#d3d0c8"
}
```

Copy the output and add it to your `settings.json`
