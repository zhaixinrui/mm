package main 

import "fmt"
import "time"
import "bytes"
import "os"
import "io/ioutil"
// import "math"

import gossh "golang.org/x/crypto/ssh"

type Task struct {
    Host string
    Command string
    Timeout time.Duration
    Sleep time.Duration
    Stdout string
    Stderr string
    Error error
}

func handleMachine(task Task) (Task){
    // 通过信任关系与目标机器建立连接
    privateKeyFile := os.Getenv("HOME") + "/.ssh/id_rsa"
    privateBytes, err := ioutil.ReadFile(privateKeyFile)
    if err != nil {
        panic("Failed to load private key from file : " + privateKeyFile + " error: " + err.Error())
    }
    signer, _ := gossh.ParsePrivateKey(privateBytes)
    config := &gossh.ClientConfig{
        User: os.Getenv("USER"),
        Auth: []gossh.AuthMethod{
            gossh.PublicKeys(signer),
        },
        Timeout: task.Timeout,
    }
    client, err := gossh.Dial("tcp", task.Host + ":22", config)
    if err != nil {
        fmt.Println("============", os.Getenv("USER"), "@", task.Host, "============")
        printRed("Failed to connect: " + err.Error())
        task.Error = err
        return task
    }
    session, err := client.NewSession()
    if err != nil {
        fmt.Println("============", os.Getenv("USER"), "@", task.Host, "============")
        printRed("Failed to create session: " + err.Error())
        task.Error = err
        return task
    }
    defer session.Close()

    // 建立连接后执行命令
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    session.Stdout = &stdout
    session.Stderr = &stderr
    err = session.Run(task.Command)
    fmt.Println("============", os.Getenv("USER"), "@", task.Host, "============")
    if err != nil {
        printRed(stderr.String())
    }else{
        printNormal(stdout.String())
    }
    task.Stdout = stdout.String()
    task.Stderr = stderr.String()
    task.Error = err
    return task
}

func execTask(taskIndex int, tasks chan Task, outputs chan Task){
    for{
        select {
            case task := <- tasks:
                // fmt.Println(taskIndex, "handle machine begin", task.Host)
                outputs <- handleMachine(task)
                time.Sleep(task.Sleep)
                // fmt.Println(taskIndex, "handle machine end", task.Host)
            default:
                // 没有拿到机器说明任务全部执行完成，退出
                // fmt.Println("task", taskIndex, "exit")
                return
        }
    }
}

func BatchExecTask(machines map[string]machine, command string, concurrent int, timeout time.Duration, sleep time.Duration) (map[string]Task) {
    var tasks   = make(chan Task, len(machines))
    var outputs = make(chan Task, len(machines))
    // 添加机器列表
    for _, m := range machines {
        tasks <- Task{
            Host: m.Host,
            Command: command,
            Timeout: timeout, 
            Sleep: sleep,
        }
    }
    // 启动任务协程
    // concurrent = math.Max(concurrent, 1)
    if(concurrent < 1){
        concurrent = 1
    }
    // concurrent = math.Min(concurrent, len(machines))
    for i := 1; i <= concurrent; i++ {
        go execTask(i, tasks, outputs)
    }
    result := make(map[string]Task)
    // 等待任务全部执行完成
    for i := 1; i <= len(machines); i++ {
        task := <- outputs
        result[task.Host] = task
        // printRed("output", output)
    }
    return result
}