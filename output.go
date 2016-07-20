package main

import "fmt"

func printNormal(content ...interface{}){
    print(37, content...)
}

func printRed(content ...interface{}){
    print(31, content...)
}

func printGreen(content ...interface{}){
    print(32, content...)
}

func printYellow(content ...interface{}){
    print(33, content...)
}

func print(color int, content ...interface{}){
    content[0] = fmt.Sprintf("\x1b[0;%dm%+v", color, content[0])
    content = append(content, "\x1b[0m")
    fmt.Println(content...)
}