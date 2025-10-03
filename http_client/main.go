package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getNameAndRepos(resp *http.Response) (name string, repos int, err error) {
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %v", resp.Status)
		return
	}
	type data struct {
		Login       string `json:"login"`
		PublicRepos int    `json:"public_repos"`
	}
	var d data
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		err = fmt.Errorf("error: %v", err)
		return
	}
	return d.Login, d.PublicRepos, nil
}

func main() {
	resp, err := http.Get("https://api.github.com/users/inbharajmani")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()
	name, repos, err := getNameAndRepos(resp)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Name: %v\nRepos: %v\n", name, repos)
}
