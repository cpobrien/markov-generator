package main

import (
    "strings"
    "fmt"
)

type str struct {
}

func (e str) Format(format string, specifiers...string) string {
    for i := 0; i < len(specifiers); i++ {
        format = strings.Replace(format, "%s", specifiers[i], 1)
    }
    return format;
}

func main() {
    str := &str()
    fmt.Println(str.Format("%s, world!", "hello"))
}
