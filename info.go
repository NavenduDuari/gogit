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

func getInfo(url string) []byte {
	var bearer = "Bearer " + utils.GithubToken
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return []byte(body)
}

func getRepos() []utils.RepoStruct {
	var repos []utils.RepoStruct
	rawRepos := getInfo(userURL)
	json.Unmarshal(rawRepos, &repos)
	return repos
}
