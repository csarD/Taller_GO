package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3001", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\Users\\cesar\\GolandProjects\\Practica_Go\\basics\\web\\index.html")

}
