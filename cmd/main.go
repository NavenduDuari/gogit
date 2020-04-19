package main

import (
	"fmt"

	"github.com/NavenduDuari/gogit"
)

func main() {
	fmt.Println("Total LOC => ", gogit.GetLOC())
	fmt.Println("Languages => ", gogit.GetLanguagePercentage())
	fmt.Println("Commit => ", gogit.GetCommitCount())
}
