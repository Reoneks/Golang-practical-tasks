package main

import (
	"fmt"
	"test/task"
)

func main() {
	fmt.Println(".")
	//task.Tree("E:\\My Programs", 0, false, []bool{true})
	array := []string{"Ubuntu", "Linux Mint", "OpenSUSE", "Debian", "Arch Linux", "Deepin OS", "Elementary OS", "Manjaro Linux", "Fedora", "Kali Linux"}
	toFind := "Linux Mint"
	fmt.Print("\n\nIndex: ", task.BinarySearch(array, toFind), "\n\n\n")
	//task.Test()
}
