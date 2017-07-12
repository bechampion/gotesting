package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "fmt"
    "encoding/json"
)

type People struct {
    Id      string      `json:"id"`
    Name    string      `json:"name"`
}

/// demo people handler
func PeopleHandler( w  http.ResponseWriter , r *http.Request) { 
    params := mux.Vars(r)
    id := params["id"]
    var pp = &People {id, "tusia"}
    b, _ := json.Marshal(pp)
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


    log.Fatal(http.ListenAndServe(":8000", r))
}
