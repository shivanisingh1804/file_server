// package main
// import (
//
//	"net/http"
//	"os"
//
// )
// func main(
//
//	dir,_ := os.Getwd()
//	http.ListenAndServe(":3000",FileServer(http.Dir(dir)))
//
// )
package main

import (
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}
