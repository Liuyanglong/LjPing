package main

import (
    "testing"
    "fmt"
)

func Test_login( t *testing.T ) {
    method := "user.login"
    auth := ""
    params := make(map[string]string)
    params["username"] = "*******"
    params["password"] = "********"
    res, err := CurlSendCmdb(method,params,auth)
    fmt.Println("from Test_login:",res)
    if err != nil {
        t.Error("curl user.login error!")
    }
}

func Test_getauth( t *testing.T ) {
    auth,err := getAuth()   
    fmt.Println("from Test_getauth:",auth,err)
    if err != nil {
        t.Error("func getAuth error!")    
    }
}

func Test_getServerList( t *testing.T ) {
    res,err := getServerIdFromBaseTag()
    fmt.Println("from Test_getServerList:",len(res))
    if err != nil {
        t.Error("func Test_getServerList error!")    
    }
}

func Test_getServerIPList( t *testing.T ) {
    res,err := getServerIPList()
    fmt.Println("from getServerIPList:",len(res))    
    if err != nil {
        t.Error("func getServerIPList error!")    
    }
}
