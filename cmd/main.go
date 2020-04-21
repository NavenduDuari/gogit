package main

import (
	"fmt"
	"html"
	"os"
	"strconv"
	"time"

	"github.com/NavenduDuari/gogit"
)

func main() {
	userName := os.Args[1:][0]
	fmt.Println(time.Now())

	emojiFingure := html.UnescapeString("&#" + strconv.Itoa(128073) + ";")
	emojiAvatar := html.UnescapeString("&#" + strconv.Itoa(128100) + ";")
	var langResp = emojiFingure + ` Languages(%) => `
	langauges := gogit.GetLanguagePercentage(userName)
	for _, lang := range langauges {
		langResp += fmt.Sprintf("\n\t\t\t\t%s -- *%.2f*", lang.Language, lang.Percentage)
	}
	locResp := fmt.Sprintf("%s Total LOC => *%d* ", emojiFingure, gogit.GetLOC(userName))
	commitResp := fmt.Sprintf("%s Total Commits => *%d* ", emojiFingure, gogit.GetCommitCount(userName))
	response := `
	` + emojiAvatar + ` *` + userName + `*
	-----------------------------------------------
	` + locResp + `
	` + commitResp + `
	` + langResp
	fmt.Println(response)
	fmt.Println(time.Now())
}

// func main() {
// 	fmt.Println(time.Now())
// 	go gogit.GetLanguagePercentage()
// 	go gogit.GetCommitCount()
// 	go gogit.GetLOC()

// 	var loc int64
// 	var lang []utils.LanguageWithPercentageStruct
// 	var commit int
// 	var ok bool
// 	go func() {
// 		for {
// 			loc, ok = <-utils.GetLOC
// 			if !ok {
// 				continue
// 			}
// 			break
// 		}
// 		utils.LocDone <- true
// 		fmt.Println("Total LOC => ", loc)
// 	}()
// 	go func() {
// 		for {
// 			lang, ok = <-utils.GetLang
// 			if !ok {
// 				continue
// 			}
// 			break
// 		}
// 		utils.LangDone <- true
// 		fmt.Println("Languages => ", lang)
// 	}()

// 	go func() {
// 		for {
// 			commit, ok = <-utils.GetCommit
// 			if !ok {
// 				continue
// 			}
// 			break
// 		}
// 		utils.CommitDone <- true
// 		fmt.Println("Commit => ", commit)
// 	}()
// 	<-utils.LocDone
// 	<-utils.LangDone
// 	<-utils.CommitDone
// 	fmt.Println(time.Now())

// }
