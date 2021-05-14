package task

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Tree(dirPath string, depth int, lastElem bool, elemArr []bool) {
	if len(elemArr) < depth {
		elemArr = append(elemArr, true)
	}
	if lastElem {
		elemArr[depth-1] = false
	}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		for u := 0; u < depth; u++ {
			if elemArr[u] {
				fmt.Print("|    ")
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Print("|----", f.Name(), "\n")
		if f.IsDir() {
			if len(files)-1 == i {
				//fmt.Sprintf("%s/%s", dirPath, f.Name())
				Tree(dirPath+"/"+f.Name(), depth+1, true, elemArr)
			} else {
				Tree(dirPath+"/"+f.Name(), depth+1, false, elemArr)
			}
		}
	}
}
