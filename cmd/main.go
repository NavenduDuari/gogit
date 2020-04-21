package main

import (
	"fmt"
	"time"

	"github.com/NavenduDuari/gogit"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println("LOC => ", gogit.GetLOC())
	fmt.Println("Commit => ", gogit.GetCommitCount())
	fmt.Println("Languages => ", gogit.GetLanguagePercentage())
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
