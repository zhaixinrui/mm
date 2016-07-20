package main

import "flag"
import "os"

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
    var fs = flag.NewFlagSet("md5", flag.ContinueOnError)
    fs.DurationVar(&timeout, "t", 0 , "command exec timeout per machine")
    fs.DurationVar(&sleep, "s", 0, "sleep time afer exec command")
    fs.IntVar(&concurrent, "c", 1, "concurrent when exec command")
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

    filename := cmdMd5.Flag.Arg(0)
    command := "md5 " + filename
    BatchExecTask(machines, command, concurrent, timeout, sleep)

    // for _,m := range machines {
    //     delete(machines, m.Host)
    // }
    // writeResult(lastResult)
    // for _,v := range lastResult{
    //     printNormal(fmt.Sprintf("%-20s%s", v.Ip, v.Host))
    // }
    // printYellow("There are", len(lastResult), "hosts")
    return 1
}







