package main

import (
	"fmt"
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	name        = flag.String("name", "teste-marcelo", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "teste-marcelo", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	//opt := &github.RepositoryListByOrgOptions{Type: "public"}
	//repos, _, err := client.Repositories.ListByOrg(ctx, "marceloagmelo", opt)
	repos, _, err := client.Repositories.List(ctx, "", nil)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Repos: %v\n", repos[0])
	for id, repo := range repos {
	//for _, repo := range repos {
		fmt.Printf("%d Repo: %s\n", id, *repo.Name)
		/*arquivoTXT, err := os.Create("/home/marcelo/github-repos.txt")
		if err != nil {
			fmt.Println("[main] Houve um erro ao criar o arquivo TXT. Erro: ", err.Error())
		}
		defer arquivoTXT.Close()
		escritor := bufio.NewWriter(arquivoTXT)
		for i := 0; i < len(repos); i++ {
			escritor.WriteString(repo[i]))
		}
		escritor.Flush()*/
	}
	//gerarArquivo(repos)
}

/*func gerarArquivo(repos []*github.Repository) {
	arquivoTXT, err := os.Create("/home/marcelo/github-repos.txt")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo TXT. Erro: ", err.Error())
		return
	}
	defer arquivoTXT.Close()
	escritor := bufio.NewWriter(arquivoTXT)
	for i := 0; i < len(repos); i++ {
		escritor.WriteString(repos[i])
	}
	escritor.Flush()
}*/
