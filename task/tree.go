package task

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Tree(DirPath string, Depth int, LastElem bool, ElemArr []bool) {
	if len(ElemArr) < Depth {
		ElemArr = append(ElemArr, true)
	}
	if LastElem {
		ElemArr[Depth-1] = false
	}
	files, err := ioutil.ReadDir(DirPath)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		for u := 0; u < Depth; u++ {
			if ElemArr[u] {
				fmt.Print("|    ")
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Print("|----", f.Name(), "\n")
		if f.IsDir() {
			if len(files)-1 == i {
				Tree(DirPath+"/"+f.Name(), Depth+1, true, ElemArr)
			} else {
				Tree(DirPath+"/"+f.Name(), Depth+1, false, ElemArr)
			}
		}
	}
}
