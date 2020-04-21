package gogit

import (
	"encoding/json"

	"github.com/NavenduDuari/gogit/utils"
)

func getCodeFrequency() []utils.CodeFreqStruct {
	repos := getRepos()
	var codeFreqs []utils.CodeFreqStruct
	go func() {
		for _, repo := range repos {
			codeFrequencyURL := getCodeFrequencyURL(repo.Name)
			go getInfo(codeFrequencyURL)

		}
	}()

	for i := 1; i <= len(repos); i++ {
		rawCodeFreq, ok := <-utils.RawInfo
		if !ok {
			continue
		}
		var codeFreq utils.CodeFreqStruct
		json.Unmarshal([]byte(rawCodeFreq), &codeFreq)
		codeFreqs = append(codeFreqs, codeFreq)
	}
	return codeFreqs
}

func GetLOC() int64 {
	var totalLOC int64
	codeFreqs := getCodeFrequency()
	for _, codeFreq := range codeFreqs {
		for _, weeklyArr := range codeFreq {
			if len(weeklyArr) == 0 {
				continue
			}
			totalLOC = totalLOC + weeklyArr[1] + weeklyArr[2]
		}
	}
	// utils.GetLOC <- totalLOC
	return totalLOC
}
