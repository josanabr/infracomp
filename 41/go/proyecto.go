package main

import (
  "fmt"
  "os"
  "net"
  "time"
  "net/http"
)

func procesarRaiz(w http.ResponseWriter, r *http.Request) {
  fmt.Println(time.Now().Format(time.RubyDate), " conexion desde ", r.RemoteAddr )
  fmt.Fprintf(w, "<h1> Bienvenido " + r.RemoteAddr + "</h1>" )
}

func main() {
   nombreServidor, _ := os.Hostname()
   ip, _ := net.LookupIP( nombreServidor )
   fmt.Println("Bienvenido a", nombreServidor, ip)
   http.HandleFunc("/", procesarRaiz)
   http.ListenAndServe(":8000", nil)
}
