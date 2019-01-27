# gossh
  gossh is  programmatically add ssh key to github.com user account.

## Usage

```
$ gossh [path]
```

specify a path to the public key file

### Github API Token

To use `gossh`, you need to get a GitHub token with an account 
To get token, first, visit GitHub account settings page, then go to Applications for the user.
set it in github.

token in gitconfig:

```
git config --global github.token "github token"

```

## Install

```
$ go get github.com/kumamo-n/gossh
```
