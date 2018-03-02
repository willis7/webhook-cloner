package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

var refs string

func main() {
	if len(os.Args) < 2 {
		log.Println("Missing refs, assuming master")
		refs = "refs/heads/master"
	} else {
		refs = os.Args[1]
	}

	_, lookErr := exec.LookPath("git")
	if lookErr != nil {
		panic(lookErr)
	}

	http.HandleFunc("/gitlab", gitlab(refs))
	http.HandleFunc("/github", github(refs))
	http.ListenAndServe(":4567", nil)
}
