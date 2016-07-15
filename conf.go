package main

import "encoding/json"
import "os"
import "fmt"

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

func loadConfig(confFile string) (error) {
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
    return fmt.Sprintf("%s/mm_%d_%d.txt", conf.ResultFilePath, os.Getuid(), os.Getppid())
}