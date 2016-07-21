package main

import "encoding/json"
import "os"
import "fmt"
import "time"

// var Conf map[string] interface{}
type machine struct {
    Ip string
    Host string
    Module []string
}

var conf struct {
    ResultFilePath string    `json:"result_file_path"`
    HostList []machine       `json:"host_list"`
    Concurrent int
    Timeout time.Duration
    Sleep  time.Duration
}

func init() {
    confFile := "conf/mm.conf"
    f, err := os.Open(confFile)
    if err != nil {
        printRed("load conf:", confFile, "fail:", err)
        return
    }
    defer f.Close()

    d := json.NewDecoder(f)
    err = d.Decode(&conf)
    if err != nil {
        printRed("parse conf:", confFile, "fail:", err)
        return
    }
}

func getResultFilePath() (string){
    return fmt.Sprintf("%s/mm_%s_%d.txt", conf.ResultFilePath, os.Getenv("USER"), os.Getppid())
}

func getAllMachines() map[string]machine{
    return nil
}