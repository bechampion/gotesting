package main

import (
    "fmt"
    "crypto/sha1"
    "encoding/hex"
)

func main() {

    s := "sha1 this stringdsadsada"
    h := sha1.New()
    h.Write([]byte(s))
    sha1_hash := hex.EncodeToString(h.Sum(nil))
    fmt.Println(s, sha1_hash)
}
