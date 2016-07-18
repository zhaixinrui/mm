package main

import "encoding/json"
import "os"
import "fmt"
import "os/user"

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
    Timeout int
}

func loadConfig() (error) {
    confFile := "conf/mm.conf"
    f, err := os.Open(confFile)
    if err != nil {
        return err
    }
    defer f.Close()

    d := json.NewDecoder(f)
    err = d.Decode(&conf)
    if err != nil {
        return err
    }

    // fmt.Println(Conf)
    return nil
}

func getResultFilePath() (string){
    user, _ := user.Current()
    return fmt.Sprintf("%s/mm_%s_%d.txt", conf.ResultFilePath, user.Username, os.Getppid())
}

func getAllMachines() map[string]machine{
    return nil
}