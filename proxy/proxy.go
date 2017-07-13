package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
)

func main() {
	tr := &http.Transport{
		DisableCompression: true,
		DisableKeepAlives:  false,
	}
	client := &http.Client{Transport: tr}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sending Request --> ", r.URL.Host+r.URL.Path)
		response := send(r, client)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Pepe", "pepe")
		fmt.Fprintf(w, response)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func send(r *http.Request, client *http.Client) string {
	req, _ := http.NewRequest("GET", r.URL.Scheme+"://"+r.URL.Host+r.URL.Path, nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}
