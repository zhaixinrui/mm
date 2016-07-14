package main

import "fmt"
import "flag"

var cmdFind = &Command{
    UsageLine: "find",
    Short: "find machines and write the machine list into file",
    Long: `
Find is used to get machines by rules and write the result into a single file.
All other command are based on this machine list.

-a        find all machines (default: false).
-m        find machines by module name (exact match).
-M        find machines by module name (fuzzy match).
-s        find machines by host name (exact match).
-S        find machines by host name (fuzzy match).
`,
}

var (
    allMachine bool
    moduleName string
    moduleReg  string
    machineName string
    machineReg string
)

func init() {
    var fs = flag.NewFlagSet("find", flag.ContinueOnError)
    fs.BoolVar(&allMachine, "a", false, "is need get all machines, default false")
    cmdFind.Run = find
}

func find(cmd *Command, args []string) int {
    fmt.Println(args)
    return 1
}