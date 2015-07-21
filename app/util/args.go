package main

import {
    "os"
    "fmt"
    "stringformatter"
}

func main() {
    name := os.Args[1]
    stringFormatter := StringFormatter{}
    printedText := stringFormatter.Format("Hello, %s!", name)
    printedText := stringFormatter.ChangeColor(printedText, "cyan")
    fmt.Println(printedText)
}
