package main

import "fmt"
import "flag"
import "os"
import "time"


var cmdSsh = &Command{
    UsageLine: "ssh [-t=1] [-s=10] command",
    Short: "ssh all machines to exec command",
    Long: `
Ssh is used to login machines which find by command 'Find' and exec the command.
It must used after command 'Find' or 'List'

-s        sleep time afer exec command (default: 0s).
-t        command exec timeout per machine (default: 0s).
-c        concurrent when exec command (default: 1).
`,
}

var (
    timeout time.Duration
    sleep time.Duration
    concurrent int
)

func init() {
    var fs = flag.NewFlagSet("ssh", flag.ContinueOnError)
    fs.DurationVar(&timeout, "t", 0 , "command exec timeout per machine")
    fs.DurationVar(&sleep, "s", 0, "sleep time afer exec command")
    fs.IntVar(&concurrent, "c", 1, "concurrent when exec command")
    cmdSsh.Flag = *fs
    cmdSsh.Run = ssh
}

func ssh(cmd *Command, args []string) int {
    cmdSsh.Flag.Parse(args)
    machines,_ := readResult()
    
    if cmdSsh.Flag.NArg() > 0 {
        cmd := cmdSsh.Flag.Arg(0)
        // fmt.Println(machines, cmd, concurrent, timeout, sleep)
        BatchExecTask(machines, cmd, concurrent, timeout, sleep)
    } else {
        tmpl(os.Stdout, helpTemplate, cmdSsh)
    }

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







