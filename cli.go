package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/tcnksm/go-gitconfig"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

var re = regexp.MustCompile("^-.*")


const (
	EnvGitHubtoken = "GITHUB_TOKEN"
	EnvGitHubAPI = "GITHUB_API"
	EnvDebug = "DEBUG"
)

const (
	ExitCodeOk    int = 0
	ExitCodeError int = 1 + iota
	//ExitCodeLoggingError
	ExitCodeBadArgs
)

const (
	defaultBaseURL = "https://api.github.com/"
	defaultCheckTimeout = 2 * time.Second
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {

	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			fmt.Fprintln(cli.errStream, helpText)
			return ExitCodeError
		case "--version":
			fmt.Fprintf(cli.errStream, "%s v%s\n", Name, Version)
			return ExitCodeError
		case "-p", "--path":
			args = args[1:]
		default:
			if re.Match([]byte(arg)) {
				fmt.Fprintf(cli.errStream, "gossh: %s : no such option\n", arg)
				return ExitCodeError
			}
		}
	}
	if len(args) == 0 {
		fmt.Fprintln(cli.errStream, "gossh: ", fmt.Errorf("too few arguments"))
		return ExitCodeBadArgs
	}

	key, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(cli.errStream,"gossh: %s error", err)
	}
	token, err := gitconfig.GithubToken()
	if err != nil {
		fmt.Println("To use gossh, you need a GitHub API Token")
	}


	baseURLStr := defaultBaseURL
	if urlStr := os.Getenv(EnvGitHubAPI); len(urlStr) != 0 {
		baseURLStr = urlStr
	}

	gitHubClient, err :=  NewGitHubClient(token,baseURLStr)

	k := new(github.Key)
	k.Key = github.String(string(key))
	_, err = gitHubClient.CreateKey(context.TODO(), k)
	if err != nil {
		fmt.Fprintln(cli.errStream, "gossh: ", err)
	}

	return ExitCodeOk
}

var helpText = `
Usage: gossh [path]
  gossh is  programmatically add ssh key to github.com user account.
`
