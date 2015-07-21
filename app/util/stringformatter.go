package main

import (
    "strings"
    "fmt"
)

type StringFormatter struct {

}

// Format basically works like fmt.Printf, except that it only works for %f,
// and (and this is a big and), it returns strings as opposed to prints them.
// personally, I find this to be a lot less ugly than just concatenating them
// with plus signs or something.
func (s StringFormatter) Format(format string, specifiers...string) string {
    for i := 0; i < len(specifiers); i++ {
        format = strings.Replace(format, "%s", specifiers[i], 1)
    }
    return format
}

