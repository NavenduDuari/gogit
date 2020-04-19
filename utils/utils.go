package utils

type RepoStruct struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type LanguageWithByteMap map[string]float64

type LanguageWithByteStruct struct {
	Language  string
	ByteCount float64
}

type LanguageWithPercentageStruct struct {
	Language   string
	Percentage float64
}

type CodeFreqStruct [][]int64

type CommitStruct []CommitBodystruct

type CommitBodystruct struct {
	SHA string `json:"sha"`
}
