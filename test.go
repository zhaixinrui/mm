package main

import "fmt"

var cmdTest = &Command{
    UsageLine: "test [appname]",
}

func init() {
    cmdTest.Run = test
}

func test(cmd *Command, args []string) int {
    fmt.Println(args)
    return 1
}