package conf

import "encoding/json"
import "os"
import "fmt"

// var Conf map[string] interface{}

var conf struct {
    HostFile string    `json:"host_file"`
    Concurrent int
    Timeout int
}

func LoadConfig(confFile string) (error) {
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
    fmt.Println(conf)
    return nil
}
