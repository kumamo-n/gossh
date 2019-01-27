package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/url"
)

type GitHub interface {
	CreateKey(ctx context.Context, key *github.Key) (int, error)
}

type GitHubClient struct {
	key string
	*github.Client
}


func NewGitHubClient(token, urlStr string) (GitHub, error) {
	if len(token) == 0 {
		fmt.Errorf("missing GitHub token")
	}
	if len(urlStr)  == 0{
		fmt.Errorf("missing GitHub urlStr")
	}

	baseURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse GitHub API URL")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(context.TODO(), ts)

	client := github.NewClient(tc)
	client.BaseURL = baseURL

	return &GitHubClient{
		Client: client,
	}, nil
}

func (c *GitHubClient) CreateKey(ctx context.Context, key *github.Key)  (int, error){
	_, _, err := c.Users.CreateKey(context.TODO(), key)
	if err != nil {
		return ExitCodeError, errors.Errorf("%s", err)
	}
	fmt.Println("A new public key was added to your account")
	return ExitCodeOk, err
}
