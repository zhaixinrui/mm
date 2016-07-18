package main 

import (
    "fmt"
    "flag"
    "os"
    "strings"
    "io"
    "html/template"
    // "reflect"
)


type Command struct {
    Run func (cmd *Command, args []string) int
    UsageLine string
    Short template.HTML
    Long template.HTML
    Flag flag.FlagSet
}

var commands = []*Command{
    cmdFind,
    cmdList,
    cmdSsh,
}

func (c *Command) Name() string {
    usageLine := strings.Split(c.UsageLine, " ")
    return usageLine[0]
}

var usageTemplate = `mm is a tool for managing machines

Usage:
    bee command [arguments]

The commands are:{{range .}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "bee help [command]" for more information about a command.
`

var helpTemplate = `usage: bee {{.UsageLine}}

{{.Long | trim}}
`
// 输出模板内容
func tmpl(w io.Writer, text string, data interface{}) {
    var trim = func (s template.HTML) template.HTML {
        return template.HTML(strings.TrimSpace(string(s)))
    }

    tpl := template.New("usage")
    tpl.Funcs(template.FuncMap {"trim": trim})
    template.Must(tpl.Parse(text))
    if err := tpl.Execute(w, data); err != nil {
        panic(err)
    }
}

func usage() {
    tmpl(os.Stdout, usageTemplate, commands)
    os.Exit(0)
}

func help(args []string) {
    if len(args) != 1 {
        usage()
        return
    }
    for _, cmd := range commands {
        if cmd.Name() == args[0] {
            tmpl(os.Stdout, helpTemplate, cmd)
            return
        }
    }

    fmt.Fprintf(os.Stdout, "Unknown help command %#q. Run 'mm help'.\n", args[0])
    os.Exit(2)
}

func main() {
    // flag.Usage = usage()
    flag.Parse()
    args := flag.Args()
    // fmt.Println(args)

    if len(args) < 1 {
        usage()
        return
    }

    if "help" == args[0] {
        help(args[1:])
        return
    }

    loadConfig()
    for _,cmd := range commands {
        if cmd.Name() == args[0] {
            cmd.Run(cmd, args[1:])
        }
    }
}




