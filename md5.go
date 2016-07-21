package main

import "flag"
import "os"
// import "fmt"
import "strings"

var cmdMd5 = &Command{
    UsageLine: "md5 [-t=0s] [-s=0s] filename",
    Short: "calculate md5 of all machines",
    Long: `
Ssh is used to login machines which find by command 'Find' and exec the command.
It must used after command 'Find' or 'List'

-s        sleep time afer exec command (default: 0s).
-t        command exec timeout per machine (default: 0s).
-c        concurrent when exec command (default: 1).
`,
}

func init() {
    loadConfig()
    var fs = flag.NewFlagSet("md5", flag.ContinueOnError)
    fs.DurationVar(&timeout, "t", conf.Timeout, "command exec timeout per machine")
    fs.DurationVar(&sleep, "s", conf.Sleep, "sleep time afer exec command")
    fs.IntVar(&concurrent, "c", conf.Concurrent, "concurrent when exec command")
    cmdMd5.Flag = *fs
    cmdMd5.Run = md5
}

func md5(cmd *Command, args []string) int {
    err := cmdMd5.Flag.Parse(args)
    if err != nil{
        return 1
    }
    machines,_ := readResult()

    if cmdMd5.Flag.NArg() <= 0 {
        tmpl(os.Stdout, helpTemplate, cmdMd5)
        return 1
    }

    if len(machines) == 0 {
        printRed("Use './mm find' to get hosts firstly")
        return 1
    }

    filename   := cmdMd5.Flag.Arg(0)
    command    := "md5sum " + filename
    result     := BatchExecTask(machines, command, concurrent, timeout, sleep)
    successMap := make(map[string][]string)
    failMap    := make(map[string][]string)

    // fmt.Printf("%+v", result)
    for h,r := range result {
        if r.Error == nil {
            successMap[r.Stdout] = append(successMap[r.Stdout], h)
        }else{
            failMap[r.Stderr] = append(failMap[r.Stderr], h)
        }
    }

    // fmt.Printf("%+v", successMap)
    if len(failMap) == 0 && len(successMap) == 1 {
        printGreen("All files are the same")
        return 1
    }
    if len(successMap) > 1 {
        printRed("There are", len(successMap), "kinds of md5:")
        for md5, hosts := range successMap {
            printRed(md5, strings.Join(hosts, " "))
        }
    }
    if len(failMap) > 0 {
        printRed("These host exec fail:")
        for errMsg, hosts := range failMap {
            printRed(errMsg, strings.Join(hosts, " "))
        }
    }

    return 1
}







