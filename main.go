package main

import (
	"os"

	"github.com/Silby17/github-repositories-lister/lib"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Println("Starting...")
	token, org := getEnv()

	client := lib.NewGitHubClient(token, org)

	client.RetrieveAllRepositories()
	client.ListActiveRepositories()
	client.ListArchivedRepositories()
	client.ListPublicRepositories()
	client.ListPrivateRepositories()

	log.Println("Done")
}

// getEnv will return the token and org environment variables if they exist
func getEnv() (string, string) {
	token := os.Getenv("TOKEN")
	org := os.Getenv("ORG")

	if token == "" || org == "" {
		log.Fatal("Missing required environment variable: (TOKEN/ORG)")
	}
	return token, org
}
