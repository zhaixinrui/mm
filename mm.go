package main 

import "fmt"
// import "time"
import "flag"
import "os"


type Command struct {
    Run func (cmd *Command, args []string) int
    UsageLine string
}

var comands = []*Command{
    // cmdTest,
}

func (c *Command) Name() string {
    return "" 
}

func main() {
    LoadConfig("conf/mm.conf")
    // flag.Usage = usage()
    flag.Parse()
    args := flag.Args()
    fmt.Println(args)

    if len(args) < 1 {
        usage()
        return
    }

    if "help" == args[0] {
        help(args[0:])
        return
    }

    for _,cmd := range comands {
        if cmd.Name() == args[0] {
            cmd.Run(cmd, args)
        }
    }
}

func usage() {
    fmt.Println("usage")
    os.Exit(0)
}

func help(cmd []string) error {
    // fmt.Println(cmd)
    return nil
}