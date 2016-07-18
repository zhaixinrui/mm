package main 

import "fmt"
import "time"
import "bytes"
import "os/user"
// import 
import gossh "golang.org/x/crypto/ssh"


type execResult struct {
    returnCode int

}

func handleMachine(machine string, timeout int) {
    user, _ := user.Current()
    config := &gossh.ClientConfig{
        User: user.Username,
        Auth: []gossh.AuthMethod{
            gossh.Password("654321"),
        },
        Timeout: 1 * time.Second,
    }
    client, err := gossh.Dial("tcp", "11.22.33.44:22", config)
    if err != nil {
        panic("Failed to dial: " + err.Error())
    }

    // Each ClientConn can support multiple interactive sessions,
    // represented by a Session.
    session, err := client.NewSession()
    if err != nil {
        panic("Failed to create session: " + err.Error())
    }
    defer session.Close()

    // Once a Session is created, you can execute a single command on
    // the remote side using the Run method.
    var b bytes.Buffer
    session.Stdout = &b
    if err := session.Run("/usr/bin/whoami"); err != nil {
        panic("Failed to run: " + err.Error())
    }
    fmt.Println(b.String())
}

func execCmd(taskIndex int, machines chan string, command string, finishFlags chan int, timeout int, sleep int) (int){
    for{
        select {
            case machine := <- machines:
                fmt.Println(taskIndex, "handle machine begin", machine)
                handleMachine(machine, 1)
                fmt.Println(taskIndex, "handle machine end", machine)
            default:
                // 没有拿到机器说明任务全部执行完成，退出
                fmt.Println("task", taskIndex, "exit")
                finishFlags <- taskIndex
                return 0
        }
    }
    return 0
}

func batchExec(machines map[string]machine, command string, concurrent int, timeout int, sleep int) {
    var cMachines    = make(chan string, len(machines))
    var cFinishFlags = make(chan int, concurrent)
    // var 
    // 添加机器列表
    for _, m := range machines {
        cMachines <- m.Host
    }
    // 启动任务协程
    for i := 0; i < concurrent; i++ {
        go execCmd(i, cMachines, command, cFinishFlags, timeout, sleep)
    }

    // 等待任务全部退出
    for i := 0; i < concurrent; i++ {
        <- cFinishFlags
    }
}