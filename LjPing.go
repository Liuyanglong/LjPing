package main

import(
    "fmt"
    "os"
    "flag"
)

func main() {
    var concurrent int
    var timeout int

    flag.IntVar(&concurrent, "con", 100, "并发数")
    flag.IntVar(&timeout, "timeout", 30, "ping 超时时间")
    flag.Parse()

    iplist,err := getServerIPList()
    if err != nil {
       fmt.Println("Error! getServerIPList error,",err) 
       os.Exit(1)
    }

    result := make(map[string]bool)

    sem := make(chan bool, concurrent)
    for _, ip := range iplist {
        sem <-true
        go func(ip string,timeout int){
            defer func(){<-sem }()

            pres := Ping(ip, timeout)
            result[ip] = pres

        }(ip,timeout)
    }
    for i :=0; i < cap(sem); i++{
        sem <-true
    }

    fmt.Println("Total Count: ",len(result))
    succ := 0
    fail := 0
    failList := make([]string,0)
    for ip,b := range result {
        if !b {
            fail += 1
            failList = append(failList,ip)
        } else {
            succ += 1    
        }
    }
    fmt.Println("Success Ping Count: ",succ)
    fmt.Println("Fail Ping Count: ",fail)
    fmt.Println("Fail Ping list: ")
    for _, fip := range failList {
        fmt.Println(fip)    
    }
}


