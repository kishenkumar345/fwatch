package main

import (
	"fwatch/internal/collect"
	"fwatch/internal/hash"
	"fmt"
)

func main() {
	allFiles := collect.CollectFiles()
	fmt.Println(allFiles)
	hashedFiles := initHash.FileHash(allFiles)

	for {
		initHash.HashCheck(hashedFiles)
	}
}