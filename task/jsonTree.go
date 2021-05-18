package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
)

type ElemType string

const (
	FileElemType      ElemType = "file"
	DirectoryElemType ElemType = "directory"
)

type Elem struct {
	Name      string
	Type      ElemType
	Childrens []Elem
}

func convFileTypeToElem(fileInfo fs.FileInfo, arrayOfFilesInDir []Elem) (result Elem) {
	var fileType ElemType
	if fileInfo.IsDir() {
		fileType = DirectoryElemType
	} else {
		fileType = FileElemType
	}
	result = Elem{
		Name:      fileInfo.Name(),
		Type:      fileType,
		Childrens: arrayOfFilesInDir,
	}
	return
}

func jsonTree(dirPath string) (elements []Elem) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		var DirectoryElements []Elem
		if f.IsDir() {
			DirectoryElements = jsonTree(dirPath + "/" + f.Name())
		}
		elements = append(elements, convFileTypeToElem(f, DirectoryElements))
	}
	return
}

func StartJsonTree(startDirPath string) {
	data, err := json.Marshal(jsonTree(startDirPath))
	if err != nil {
		log.Fatal(err)
	} else {
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, data, "", "  ")
		fmt.Println(prettyJSON.String())
	}
}
