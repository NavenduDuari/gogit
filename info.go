package gogit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NavenduDuari/gogit/utils"
)

var (
	reposURL = "https://api.github.com/repos/NavenduDuari/"
	userURL  = "https://api.github.com/users/NavenduDuari/repos"
)

func getCodeFrequencyURL(repoName string) string {
	return reposURL + repoName + "/stats/code_frequency"
}

func getLanguagesURL(repoName string) string {
	return reposURL + repoName + "/languages"
}

func getCommitURL(repoName string) string {
	return reposURL + repoName + "/commits"
}

func getInfo(url string) {
	var bearer = "Bearer " + utils.GithubToken
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	utils.RawInfo <- body
	// return body
}

func getRepos() []utils.RepoStruct {
	ok := false
	var repos []utils.RepoStruct
	var rawRepos []byte
	go getInfo(userURL)
	for {
		rawRepos, ok = <-utils.RawInfo
		if !ok {
			continue
		}
		break
	}
	json.Unmarshal(rawRepos, &repos)
	return repos
}
