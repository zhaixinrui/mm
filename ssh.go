package main

import "flag"
import "os"

var cmdSsh = &Command{
    UsageLine: "ssh [-t=0s] [-s=0s] command",
    Short: "ssh all machines to exec command",
    Long: `
Ssh is used to login machines which find by command 'Find' and exec the command.
It must used after command 'Find' or 'List'

-s        sleep time afer exec command (default: 0s).
-t        command exec timeout per machine (default: 0s).
-c        concurrent when exec command (default: 1).
`,
}

func init() {
    var fs = flag.NewFlagSet("ssh", flag.ContinueOnError)
    fs.DurationVar(&timeout, "t", 0 , "command exec timeout per machine")
    fs.DurationVar(&sleep, "s", 0, "sleep time afer exec command")
    fs.IntVar(&concurrent, "c", 1, "concurrent when exec command")
    cmdSsh.Flag = *fs
    cmdSsh.Run = ssh
}

func ssh(cmd *Command, args []string) int {
    err := cmdSsh.Flag.Parse(args)
    if err != nil{
        return 1
    }
    machines, _ := readResult()
    if cmdMd5.Flag.NArg() <= 0 {
        tmpl(os.Stdout, helpTemplate, cmdMd5)
        return 1
    }

    if len(machines) == 0 {
        printRed("Use './mm find' to get hosts firstly")
        return 1
    }
    
    command := cmdSsh.Flag.Arg(0)
    // fmt.Println(machines, command, concurrent, timeout, sleep)
    BatchExecTask(machines, command, concurrent, timeout, sleep)

    return 1
}







