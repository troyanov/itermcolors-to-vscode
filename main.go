package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	mapping = map[string]string{
		"terminal.background":          "Background Color",
		"terminal.foreground":          "Foreground Color",
		"terminalCursor.background":    "Cursor Text Color",
		"terminalCursor.foreground":    "Cursor Color",
		"terminal.ansiBlack":           "Ansi 0 Color",
		"terminal.ansiBlue":            "Ansi 4 Color",
		"terminal.ansiBrightBlack":     "Ansi 8 Color",
		"terminal.ansiBrightBlue":      "Ansi 12 Color",
		"terminal.ansiBrightCyan":      "Ansi 14 Color",
		"terminal.ansiBrightGreen":     "Ansi 10 Color",
		"terminal.ansiBrightMagenta":   "Ansi 13 Color",
		"terminal.ansiBrightRed":       "Ansi 9 Color",
		"terminal.ansiBrightWhite":     "Ansi 15 Color",
		"terminal.ansiBrightYellow":    "Ansi 11 Color",
		"terminal.ansiCyan":            "Ansi 6 Color",
		"terminal.ansiGreen":           "Ansi 2 Color",
		"terminal.ansiMagenta":         "Ansi 5 Color",
		"terminal.ansiRed":             "Ansi 1 Color",
		"terminal.ansiWhite":           "Ansi 7 Color",
		"terminal.ansiYellow":          "Ansi 3 Color",
		"terminal.selectionBackground": "Selection Color",
	}
)

type itermcolors struct {
	XMLName xml.Name `xml:"plist"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Dict    struct {
		Text string   `xml:",chardata"`
		Key  []string `xml:"key"`
		Dict []struct {
			Text   string   `xml:",chardata"`
			Key    []string `xml:"key"`
			Real   []string `xml:"real"`
			String string   `xml:"string"`
		} `xml:"dict"`
	} `xml:"dict"`
}

func main() {
	path := flag.String("f", "", "path to .itermcolors file")
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("Failed to open %s: %s", *path, err)
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Failed to read %s: %s", *path, err)
	}

	colors := itermcolors{}
	xml.Unmarshal(data, &colors)

	result := make(map[string]string)
	for k, v := range mapping {
		result[k] = getColor(colors, v)
	}

	settings, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("\"workbench.colorCustomizations\": %s\n", string(settings))
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 32)
	return f
}

func getColor(colors itermcolors, key string) string {
	for i, color := range colors.Dict.Key {
		if color != key {
			continue
		}
		r := fmt.Sprintf("%02x",
			int(math.Round(parseFloat(colors.Dict.Dict[i].Real[3])*255)))
		g := fmt.Sprintf("%02x",
			int(math.Round(parseFloat(colors.Dict.Dict[i].Real[2])*255)))
		b := fmt.Sprintf("%02x",
			int(math.Round(parseFloat(colors.Dict.Dict[i].Real[1])*255)))
		return fmt.Sprintf("#%s%s%s", r, g, b)
	}
	return "#000000"
}
