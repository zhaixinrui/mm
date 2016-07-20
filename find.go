package main

import "fmt"
import "flag"
import "strings"

var cmdFind = &Command{
    UsageLine: "find",
    Short: "find machines and write the machine list into file",
    Long: `
Find is used to get machines by rules and write the result into a single file.
All other command are based on this machine list.

-m        find machines by module name (fuzzy match).
-s        find machines by host name (fuzzy match).
`,
}

func init() {
    var fs = flag.NewFlagSet("find", flag.ContinueOnError)
    fs.StringVar(&moduleName, "m", "", "module fuzzy name")
    fs.StringVar(&machineName, "s", "", "host fuzzy name")
    cmdFind.Flag = *fs
    cmdFind.Run = find
}

func find(cmd *Command, args []string) int {
    cmdFind.Flag.Parse(args)
    machines := filter(moduleName, machineName)
    writeResult(machines)
    for _,v := range machines{
        printNormal(fmt.Sprintf("%-20s%s", v.Ip, v.Host))
    }
    printYellow("There are", len(machines), "hosts")
    return 1
}

func needDelete(moduleName string, machineName string, machine machine) bool {
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

func filter(moduleName string, machineName string) (machines map[string]machine) {
    machines = make(map[string]machine)
    // allMachine是全量的机器列表，然后根据规则进行过滤，最终得到待操作的机器列表
    for _,m := range conf.HostList{
        if !needDelete(moduleName, machineName, m){
            machines[m.Ip] = m
        }
    }

    return machines
}







