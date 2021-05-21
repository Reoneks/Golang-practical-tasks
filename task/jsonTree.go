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
	Name      string   `json:"name"`
	Type      ElemType `json:"elemType"`
	Childrens []Elem   `json:"childrens"`
}

func ConvFileTypeToElem(fileInfo fs.FileInfo, arrayOfFilesInDir []Elem) (result Elem) {
	fileType := FileElemType
	if fileInfo.IsDir() {
		fileType = DirectoryElemType
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
		return
	}
	for _, f := range files {
		var directoryElements []Elem
		if f.IsDir() {
			directoryElements = jsonTree(dirPath + "/" + f.Name())
		}
		elements = append(elements, ConvFileTypeToElem(f, directoryElements))
	}
	return
}

func StartJsonTree(startDirPath string) {
	data, err := json.Marshal(jsonTree(startDirPath))
	if err != nil {
		log.Fatal(err)
		return
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(prettyJSON.String())
}
