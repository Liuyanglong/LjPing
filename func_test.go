package main

import (
    "testing"
)

func Test_ping( t *testing.T ) {
    ip1 := "172.16.3.124"
    pr := Ping(ip1,10)
    if ! pr {
        t.Error("ping dns error!")    
    }
}

func Test_baidu( t *testing.T ) {
    l := "baidu.com"
    pr := Ping(l, 10)
    if ! pr {
        t.Error("ping baidu error!")    
    }    
}

func Test_noexistip( t *testing.T ) {
    l := "172.25.25.25"    
    pr := Ping(l,1)
    if pr {
        t.Error("ping none exist ip error!")    
    }
}
