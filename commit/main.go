package main

import (
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"time"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/src-d/go-git.v4"
)

func main() {
	directory := "/tmp/teste"
	username := "marceloagmelo@gmail.com"
	password := "magm0101"

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// ... we need a file to commit so let's create a new file inside of the
	// worktree of the project using the go standard library.
	fmt.Println("echo \"hello world!\" > example-git-golang")
	filename := filepath.Join(directory, "example-git-golang")
	err = ioutil.WriteFile(filename, []byte("hello world alterado!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Adds the new file to the staging area.
	fmt.Println("git add example-git-golang")
	_, err = w.Add("example-git-golang")
	if err != nil {
		log.Fatal(err)
	}

	// We can verify the current status of the worktree using the method Status.
	fmt.Println("git status --porcelain")
	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	fmt.Println("git commit -m \"Teste pelo golang alterado\"")
	commit, err := w.Commit("Teste pelo golang alterado", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Marcelo Melo",
			Email: "marceloagmelo@gmail.com",
			When:  time.Now(),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Prints the current HEAD to verify that all worked well.
	fmt.Println("git show -s")
	obj, err := r.CommitObject(commit)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(obj)

	fmt.Println("git push")
	//err = r.Push(&git.PushOptions{})
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		}})
	if err != nil {
		log.Fatal(err)
	}
}