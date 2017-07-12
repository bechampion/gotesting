package main
import (
    "fmt"
    "net/http"
    "reflect"
)



func main() {
    var urls []string
    for i := 1 ;  i < 11 ; i ++ {
        urls =  append(urls, "www.com.ar")
        //fmt.Println(i)
        //urls[i] = "www.com.ar"
    }

    fmt.Println(urls)
    resp, err := http.Get("http://example.com/")
    fmt.Println(reflect.TypeOf(resp))
    fmt.Println(reflect.TypeOf(resp.Status))
    fmt.Println(resp.Header)
    fmt.Println(reflect.TypeOf(resp.Header))
    for k, v := range resp.Header {
        fmt.Printf("k:[%s] -- v:[%s]\n",reflect.TypeOf(k),reflect.TypeOf(v))
        fmt.Printf("k:[%s] -- v:[%s]\n",k,v)
        fmt.Println(v)
    }
    fmt.Println(err)
}

