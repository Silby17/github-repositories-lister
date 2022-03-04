package main

import (
	"os"

	"github.com/Silby17/github-repositories-lister/lib"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(false)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
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

	log.Println("\nDone")
}

// getEnv will return the token and org environment variables if they exist
func getEnv() (string, string) {
	token := os.Getenv("token")
	org := os.Getenv("org")

	if token == "" || org == "" {
		log.Fatal("Missing required environment variable: (token or org)")
	}
	return token, org
}
