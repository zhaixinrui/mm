package main

import "fmt"
import "flag"

var cmdList = &Command{
    UsageLine: "list",
    Short: "list or modify last result",
    Long: `
List is used to get machines by rules and write the result into a single file.
It will modify the result file of comman 'Find'

-a        append machines to last result (default: false).
-d        delete machines from last result (default: false).
-m        find machines by module name (fuzzy match).
-s        find machines by host name (fuzzy match).
`,
}

// var (
//     isAppend bool
//     isDelete bool
//     moduleName string
//     machineName string
// )

func init() {
    var fs = flag.NewFlagSet("list", flag.ContinueOnError)
    fs.BoolVar(&isAppend, "a", false, "append machines to last result")
    fs.BoolVar(&isDelete, "d", false, "delete machines from last result")
    fs.StringVar(&moduleName, "m", "", "module fuzzy name")
    fs.StringVar(&machineName, "s", "", "host fuzzy name")
    cmdList.Flag = *fs
    cmdList.Run = list
}

func list(cmd *Command, args []string) int {
    cmdList.Flag.Parse(args)
    lastResult,_ := readResult()
    
    // 仅查看
    if len(args) == 0 {
        // fmt.Println(lastResult, len(lastResult))
        if len(lastResult) > 0{
            for _,v := range lastResult{
                printNormal(fmt.Sprintf("%-20s%s", v.Ip, v.Host))
            }
            printYellow("There are", len(lastResult), "hosts")
        } else {
            printRed("Use './mm find' to get hosts firstly")
        }
        return 1
    }

    // 修改
    machines := filter(moduleName, machineName)
    if isAppend {
        for _,m := range machines {
            lastResult[m.Host] = m
        }
    }
    if isDelete {
        for _,m := range machines {
            delete(lastResult, m.Host)
        }
    }
    writeResult(lastResult)
    for _,v := range lastResult{
        printNormal(fmt.Sprintf("%-20s%s", v.Ip, v.Host))
    }
    printYellow("There are", len(lastResult), "hosts")
    return 1
}







