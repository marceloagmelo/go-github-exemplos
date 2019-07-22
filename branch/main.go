package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {
	url := "https://github.com/marceloagmelo/cursodego.git"
	directory := "/tmp/teste"
	username := "marceloagmelo@gmail.com"
	password := "magm0101"
	nomeBranch := "branch-teste"

	// Clone the given repository to the given directory
	fmt.Printf("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new branch to the current HEAD
	fmt.Println("git branch branch-teste")

	headRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new plumbing.HashReference object with the name of the branch
	// and the hash from the HEAD. The reference name should be a full reference
	// name and not an abbreviated one, as is used on the git cli.
	//
	// For tags we should use `refs/tags/%s` instead of `refs/heads/%s` used
	// for branches.
	refName := plumbing.NewBranchReferenceName(nomeBranch)
	ref := plumbing.NewHashReference(refName, headRef.Hash())

	// The created reference is saved in the storage.
	err = r.Storer.SetReference(ref)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("git push")
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		}})
	if err != nil {
		log.Fatal(err)
	}

	// Or deleted from it.
	/*fmt.Println("git branch -D branch-teste")
	err = r.Storer.RemoveReference(ref.Name())
	if err != nil {
		log.Fatal(err)
	}*/
}
