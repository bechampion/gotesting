package main
import fmt "fmt";


type dada struct {
    pepe int
    dada string
}


func main() {
    fmt.Println(dada{20,"pepe"})
    s := dada{ pepe : 1 , dada : "caca" }
    fmt.Printf("dada");
    fmt.Println(s.pepe);


}
