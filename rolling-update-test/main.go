package main
import (
    "fmt"
    "log"
    "net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is version 1.") //
}

func main() {
    http.HandleFunc("/", sayhello) //
    log.Println("This is version 1.")
    err := http.ListenAndServe(":9090", nil) //
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
