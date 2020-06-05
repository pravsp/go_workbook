package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string
	var relPath []string
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]
	testFile := "/Users/praveen/Desktop/my_data/Learning/Go_Programming/go_workbook/.git/index"
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	fmt.Println(strings.TrimPrefix(testFile, arg))
	fmt.Println("~~~~~~~~~~~~~~~~")
	root := arg
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileName := filepath.Base(path)
			//relativeFilePath := strings.TrimPrefix(path, root)
			relativeFilePath, _ := filepath.Rel(root, path)
			if !strings.HasPrefix(relativeFilePath, ".") {
				files = append(files, fileName)
				if len(relativeFilePath) == len(fileName) {
					relPath = append(relPath, "./")
				} else {
					relPath = append(relPath, relativeFilePath[:len(relativeFilePath)-len(fileName)])
				}
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	outputFile := os.Args[2]
	fo, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for id, file := range files {
		fmt.Println(file, relPath[id])
		l, err := fo.WriteString(file + "," + relPath[id] + "\n")
		if err != nil {
			fmt.Println(err)
			panic(err)
		} else {
			fmt.Println(l, "bytes written successfully")
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
