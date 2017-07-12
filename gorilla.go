package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "fmt"
    "encoding/json"
    "time"
    "os/exec"
    //"reflect"
)

type People struct {
    Id      string      `json:"id"`
    Name    string      `json:"name"`
}

var Messages = make(chan string)


/// demo people handler
func PeopleHandler( w  http.ResponseWriter , r *http.Request  ) { 
    params := mux.Vars(r)
    id := params["id"]
    pp := &People {id, "tusia"}
    b, _ := json.Marshal(pp)
    go func(){ 
        time.Sleep(2000 * time.Millisecond)
        cmd , _ := exec.Command("date","+%s").Output()
        Messages <- string(cmd)


    }()
    fmt.Fprintf(w,"%s" , b)
    b = nil
}
func main() {
    r := mux.NewRouter()
    //demo route
    r.HandleFunc("/", func( w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w,"Nothing here\n") }).Methods("GET")
    //some other router
    r.HandleFunc("/people/{id:[0-9]+}", PeopleHandler).Methods("GET")
    //not found handler
    r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter ,  r *http.Request) {
        w.WriteHeader(http.StatusNotImplemented)
        fmt.Fprintf(w , "NO")
    } )
    fmt.Println(Messages)
    go func() { for elem := range Messages{
        fmt.Println(elem)
    }}()
    log.Fatal(http.ListenAndServe(":8000", r))

}
