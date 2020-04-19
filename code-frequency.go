package gogit

import (
	"encoding/json"

	"github.com/NavenduDuari/gogit/utils"
)

func getCodeFrequency() []utils.CodeFreqStruct {
	repos := getRepos()
	var codeFreqs []utils.CodeFreqStruct
	for _, repo := range repos {
		var codeFreq utils.CodeFreqStruct
		codeFrequencyURL := getCodeFrequencyURL(repo.Name)
		rawCodeFreq := getInfo(codeFrequencyURL)
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
			totalLOC = totalLOC + weeklyArr[1] + weeklyArr[2]
		}
	}
	return totalLOC
}
