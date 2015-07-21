package main

import (
    "strings"
    "fmt"
    "bytes"
)
var colors = map[string]string{
    "red": "31",
    "green": "32",
    "yellow": "33",
    "blue": "34",
    "magenta": "35",
    "cyan": "36",
    "white": "37",
}

type StringFormatter struct {
}

// ChangeColor takes a string and a color available in the ANSI escape
// colors, and returns a string with that escape code attached. This way,
// the newly formatted string will print in that color.
func (s StringFormatter) ChangeColor(text string, color string) string {
    var buffer bytes.Buffer
    buffer.WriteString("\033[1;%sm")
    buffer.WriteString(text)
    buffer.WriteString("\033[0;0m")
    ansiColor := colors[color]
    return s.Format(buffer.String(), ansiColor)
}

// Format basically works like fmt.Printf, except that it only works for %f,
// and (and this is a big and), it returns strings as opposed to prints them.
// personally, I find this to be a lot less ugly than just concatenating them
// with plus signs or something.
func (s StringFormatter) Format(text string, specifiers...string) string {
    formattedString := text
    for i := 0; i < len(specifiers); i++ {
        formattedString = strings.Replace(formattedString, "%s", specifiers[i], 1)
    }
    return formattedString
}

func main() {
    stringFormatter := StringFormatter{}

    colors := []string{"red", "green", "yellow", "blue", "magenta", "cyan", "white"}
    rawText := "Hello, %s!"
    textToPrint := stringFormatter.Format(rawText, "me")

    for i := 0; i < len(colors); i++ {
        currentColor := colors[i]
        coloredText := stringFormatter.ChangeColor(textToPrint, currentColor)
        fmt.Println(coloredText)
    }
}
