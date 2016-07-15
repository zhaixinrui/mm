package main

import "fmt"
import "flag"
// import "os"
import "strings"
// import "sort"

var cmdFind = &Command{
    UsageLine: "find",
    Short: "find machines and write the machine list into file",
    Long: `
Find is used to get machines by rules and write the result into a single file.
All other command are based on this machine list.

-a        find all machines (default: false).
-m        find machines by module name (fuzzy match).
-s        find machines by host name (fuzzy match).
`,
}

var (
    allMachine bool
    moduleName string
    machineName string
)

func init() {
    var fs = flag.NewFlagSet("find", flag.ContinueOnError)
    fs.BoolVar(&allMachine, "a", false, "is need get all machines, default false")
    fs.StringVar(&moduleName, "m", "", "module fuzzy name")
    fs.StringVar(&machineName, "s", "", "host fuzzy name")
    cmdFind.Flag = *fs
    cmdFind.Run = find
}

func find(cmd *Command, args []string) int {
    cmdFind.Flag.Parse(args)
    machines := filter(allMachine, moduleName, machineName)
    fmt.Println(getResultFilePath())
    for _,v := range machines{
        printNormal(fmt.Sprintf("%-20s%s", v.Ip, v.Host))
    }
    printYellow("There are", len(machines), "hosts")
    return 1
}

func needDelete(moduleName string, machineName string, machine machine) bool {
    // fmt.Println(machine.Module, moduleName, sort.SearchStrings(machine.Module, moduleName))
    if "" != moduleName {
        find := false
        for _,v := range machine.Module {
            if v == moduleName {
                find = true
            }
        }
        if(false == find){
            return true
        }
    }
    if "" != machineName && !strings.Contains(machine.Host, machineName) {
        return true
    }
    return false
}

func filter(allMachine bool, moduleName string, machineName string) (machines []machine) {
    loadConfig("conf/mm.conf")
    // if(allMachine){
    //     return conf.HostList
    // }
    machines = conf.HostList

    for index := 0; index < len(machines); {
        // fmt.Println(index, machine, len(machines), machines[:index], machines[index+1:])
        if needDelete(moduleName, machineName, machines[index]) {
            machines = append(machines[:index], machines[index+1:]...)
        }else{
            // 没有找到的话，游标前移一位，继续检查下一个
            index++
        }
    }

    return machines
}







