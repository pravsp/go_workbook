package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var s string
	//fmt.Scanln(&s)
	//fmt.Println(s)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	editor := bufio.NewReadWriter(reader, writer)
	editor.WriteString("Enter text:")
	editor.Flush()
	str, _ := editor.ReadString('\n')
	fmt.Println(str)
	fmt.Print("Enter Number:")
	str, _ = editor.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
