package main

import (
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
	"log"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
)

func main() {
	url := "https://github.com/marceloagmelo/cursodego.git"
	directory := "/tmp/teste"
	username := "marceloagmelo@gmail.com"
	password := "magm0101"

	// Clone the given repository to the given directory
	fmt.Printf("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:      url,
		Progress: os.Stdout,
	})	
	
	/*r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		log.Fatal(err)
	}*/

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commit)	
}