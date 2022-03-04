package lib

import (
	"context"

	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client  *github.Client
	repos   []*github.Repository
	ctx     context.Context
	orgName string
}

func NewGitHubClient(token string, org string) *GitHubClient {
	log.Println("Initializing GitHub Client")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	ctx := context.Background()
	return &GitHubClient{
		ctx:     ctx,
		client:  github.NewClient(oauth2.NewClient(ctx, ts)),
		orgName: org,
	}
}

func (gh *GitHubClient) RetrieveAllRepositories() {
	log.Println("Retrieving all repositories...")
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var repos []*github.Repository
	for {
		result, resp, err := gh.client.Repositories.ListByOrg(gh.ctx, gh.orgName, opt)
		if err != nil {
			log.Fatalf("Error listing Organization repositories: %s", err.Error())
		}
		repos = append(repos, result...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	gh.repos = repos
	log.Printf("%d repositories retieved", len(repos))
}

func (gh *GitHubClient) ListPrivateRepositories() {
	log.Println("Listing Private repositories...")
	if len(gh.repos) == 0 {
		log.Println("No repositories have been retrieved yet")
		return
	}
	var private []string
	for _, r := range gh.repos {
		if *r.Private {
			println(*r.Name)
			private = append(private, *r.Name)
		}
	}
	log.Printf("%d Private repositories", len(private))
}

func (gh *GitHubClient) ListPublicRepositories() {
	log.Println("Listing Public repositories...")
	if len(gh.repos) == 0 {
		log.Println("No repositories have been retrieved yet")
		return
	}
	var public []string
	for _, r := range gh.repos {
		if !*r.Private {
			println(*r.Name)
			public = append(public, *r.Name)
		}
	}
	log.Printf("%d Public repositories", len(public))
}

func (gh *GitHubClient) ListActiveRepositories() {
	log.Println("Listing Active repositories...")
	if len(gh.repos) == 0 {
		log.Println("No repositories have been retrieved yet")
		return
	}
	var active []string
	for _, r := range gh.repos {
		if !*r.Archived {
			println(*r.Name)
			active = append(active, *r.Name)
		}
	}
	log.Printf("%d Active repositories", len(active))
}

func (gh *GitHubClient) ListArchivedRepositories() {
	log.Println("Listing Archived repositories...")
	if len(gh.repos) == 0 {
		log.Println("No repositories have been retrieved yet")
		return
	}
	var archived []string
	for _, r := range gh.repos {
		if *r.Archived {
			println(*r.Name)
			archived = append(archived, *r.Name)
		}
	}
	log.Printf("%d Archived repositories", len(archived))
}
