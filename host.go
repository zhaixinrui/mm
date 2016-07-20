package main

import "os"
import "fmt"
import "bufio"
import "io"
import "strings"

func writeResult(machines map[string]machine) (error) {
    resultFile := getResultFilePath()
    f, err := os.Create(resultFile)
    if err != nil {
        return err
    }
    defer f.Close()
    for _,machine := range machines {
        s := fmt.Sprintf("%-20s%s\n", machine.Ip, machine.Host)
        f.WriteString(s)
    }
    f.Sync()
    return nil
}

func readResult() (machines map[string]machine, err error) {
    machines = make(map[string]machine)
    resultFile := getResultFilePath()
    f,err := os.Open(resultFile)
    if err != nil {
        // printRed(err)
        return
    }
    defer f.Close()

    reader := bufio.NewReader(f)
    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            if err != io.EOF {
                printRed("read last query result error, file:", resultFile)
            }
            break
        }

        str := string(line)
        if "" == str {
            continue
        }

        l := strings.Fields(string(line))
        m := machine{l[0], l[1], make([]string, 0)}
        machines[m.Host] = m
        // fmt.Println(resultFile, f, machines, len(l))
    }
    // fmt.Println(resultFile)
    return
}
