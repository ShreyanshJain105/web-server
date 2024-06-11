package main

import (
    "fmt"
    "log" // Import log if you intend to use it
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful\n")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name = %s \n", name)
    fmt.Fprintf(w, "Address = %s \n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { // Corrected typo
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "method is not supported", http.StatusNotFound) // Improved error message
        return
    }
    fmt.Fprintf(w, "hello Buddy!")
}

func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler) // Corrected typo
    fmt.Println("starting server at port 1206\n")
    if err := http.ListenAndServe(":1206", nil); err != nil {
        log.Fatal(err) // Use log.Fatal if you imported the log package
    }
}

