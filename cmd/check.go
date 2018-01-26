package main

import (
	"os"
	"gopkg.in/src-d/go-git.v4"
	"fmt"
	"encoding/json"
	"syscall"
)

func getRepo(path string) (*git.Repository, error) {
	var r *git.Repository
	if _, err := os.Stat(path); os.IsNotExist(err) {
		r, err = git.PlainClone(path, false, &git.CloneOptions{
			URL:          "https://github.com/ReadyTalk/avian.git",
			SingleBranch: true,
			Progress:     os.Stderr,
		})
		if err != nil {
			return r, err
		}
	} else {
		r, err = git.PlainOpen(path)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

type Response struct {
	Version  []Version      `json:"version"`
	Metadata []MetadataPair `json:"metadata"`
}

type Version struct {
	Ref string `json:"ref"`
}

type Versions []Version

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func onError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		syscall.Exit(-1)
	}
}

func main() {

	sourceRoot := "/tmp/git-resource-repo-cache"
	r, err := getRepo(sourceRoot)
	onError(err)

	r.Fetch(&git.FetchOptions{})

	w, err := r.Worktree()
	onError(err)

	w.Checkout(&git.CheckoutOptions{
		Branch: "FETCH_HEAD",
	})

	ref, err := r.Head()
	onError(err)

	c, err := r.CommitObject(ref.Hash())
	onError(err)

	json.NewEncoder(os.Stdout).Encode(Versions{
		Version{
			Ref: c.Hash.String(),
		},
	})
}
