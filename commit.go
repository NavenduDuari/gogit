package gogit

import (
	"encoding/json"

	"github.com/NavenduDuari/gogit/utils"
)

func getCommit() []utils.CommitStruct {
	var commitStructArr []utils.CommitStruct
	repos := getRepos()
	go func() {
		for _, repo := range repos {
			commitURL := getCommitURL(repo.Name)
			go getInfo(commitURL)
		}
	}()

	for i := 1; i <= len(repos); i++ {
		rawCommitInfo, ok := <-utils.RawInfo
		if !ok {
			continue
		}
		var commitStruct utils.CommitStruct
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
	// utils.GetCommit <- totalCommit
	return totalCommit
}
