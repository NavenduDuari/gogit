package gogit

import (
	"encoding/json"

	"github.com/NavenduDuari/gogit/utils"
)

var (
	totalByteCount            float64
	languageWithByteStructArr []*utils.LanguageWithByteStruct
)

func getLanguages() []utils.LanguageWithByteMap {
	var languages []utils.LanguageWithByteMap
	repos := getRepos()
	go func() {
		for _, repo := range repos {
			languagesURL := getLanguagesURL(repo.Name)
			go getInfo(languagesURL)
		}
	}()

	for i := 1; i <= len(repos); i++ {
		rawLang, ok := <-utils.RawInfo
		if !ok {
			continue
		}
		var lang utils.LanguageWithByteMap
		json.Unmarshal([]byte(rawLang), &lang)
		languages = append(languages, lang)
	}

	return languages
}

func calculateLanguageByte() []*utils.LanguageWithByteStruct {
	languageWithByteMapArr := getLanguages()
	for _, languageWithByteMap := range languageWithByteMapArr {
		for lang, byteCount := range languageWithByteMap {
			isUpdated := false
			if len(languageWithByteStructArr) == 0 {
				languageWithByteStructArr = append(languageWithByteStructArr,
					&utils.LanguageWithByteStruct{
						Language:  lang,
						ByteCount: byteCount,
					})
				totalByteCount += byteCount
				isUpdated = true
			} else {
				for _, languageWithByteStruct := range languageWithByteStructArr {
					if lang == languageWithByteStruct.Language {
						languageWithByteStruct.ByteCount += byteCount
						totalByteCount += byteCount
						isUpdated = true
					}
				}
				if !isUpdated {
					languageWithByteStructArr = append(languageWithByteStructArr,
						&utils.LanguageWithByteStruct{
							Language:  lang,
							ByteCount: byteCount,
						})
					totalByteCount += byteCount
				}
			}
		}
	}

	return languageWithByteStructArr
}

func GetLanguagePercentage() []utils.LanguageWithPercentageStruct {
	var languageWithPercentageStructArr []utils.LanguageWithPercentageStruct

	languageWithByteStructArr := calculateLanguageByte()
	for _, languageWithByteStruct := range languageWithByteStructArr {
		languageWithPercentageStructArr = append(languageWithPercentageStructArr,
			utils.LanguageWithPercentageStruct{
				Language:   languageWithByteStruct.Language,
				Percentage: (languageWithByteStruct.ByteCount / totalByteCount * 100),
			})
	}
	// utils.GetLang <- languageWithPercentageStructArr
	return languageWithPercentageStructArr
}
