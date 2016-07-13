package main 

import "fmt"
import "time"

// import "conf"


func handleMachine(machine string, timeOut int) {
    time.Sleep(1000 * time.Millisecond)
    fmt.Println(machine, timeOut)
}

func execTask(taskIndex int, machines chan string, finishFlags chan int) (int){
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

var machineNum = 50
var taskNum    = 10

var machines    = make(chan string, machineNum)
var finishFlags = make(chan int, taskNum)

func zxr() error {
    err := LoadConfig("conf/mm.conf")
    return err
}
func main() {
    err := zxr()
    // fmt.Println(conf.conf["key"])
    // fmt.Println(conf.Conf["key2"])
    // for k,v := range(conf.Conf) {
    //     fmt.Println(k, v)
    // }
    fmt.Println(err)
    // // 添加机器列表
    // for i := 0; i < machineNum; i++ {
    //     machines <- fmt.Sprintf("host_%d", i)
    // }

    // // 启动任务协程
    // for i := 0; i < taskNum; i++ {
    //     go execTask(i, machines, finishFlags)
    // }

    // // 等待任务全部退出
    // for i := 0; i <    taskNum; i++ {
    //     <- finishFlags
    // }
}