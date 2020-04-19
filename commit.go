package gogit

import (
	"encoding/json"

	"github.com/NavenduDuari/gogit/utils"
)

func getCommit() []utils.CommitStruct {
	var commitStructArr []utils.CommitStruct
	repos := getRepos()
	for _, repo := range repos {
		var commitStruct utils.CommitStruct
		commitURL := getCommitURL(repo.Name)
		rawCommitInfo := getInfo(commitURL)
		json.Unmarshal([]byte(rawCommitInfo), &commitStruct)
		commitStructArr = append(commitStructArr, commitStruct)
	}
	return commitStructArr
}

func GetCommitCount() int {
	var totalCommit int
	commitStructArr := getCommit()
	for _, commitStruct := range commitStructArr {
		totalCommit += len(commitStruct)
	}
	return totalCommit
}
