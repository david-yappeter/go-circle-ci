package main

import(
  "github.com/gorilla/mux"
  "log"
  "os"
  "fmt"
  "net/http"
)

var temp = 0

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  router := mux.NewRouter()
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World!")
  })

  fmt.Println(an err)

  router.HandleFunc("/hit", func(w http.ResponseWriter, r *http.Request) {
    temp++
    fmt.Fprintf(w, "I have been hit %d time(s)", temp)
  })

  log.Println("Listen and Serve at http://localhost:"+port)
  log.Println(http.ListenAndServe(":"+port, router))
}

