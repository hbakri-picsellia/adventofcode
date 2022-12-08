package models

import (
	"math"
	"strings"
)

type Folder struct {
	Name    string
	Files   []File
	Folders []*Folder
	Parent  *Folder
	Scanned bool
}

func (folder *Folder) DecodeChildren(input string) {
	if !folder.Scanned {
		for _, value := range strings.Split(input, "\n") {
			parts := strings.SplitN(value, " ", 2)
			if parts[0] == "dir" {
				newFolder := Folder{Name: parts[1], Parent: folder}
				folder.Folders = append(folder.Folders, &newFolder)
			} else {
				newFile := File{}
				newFile.Decode(value)
				folder.Files = append(folder.Files, newFile)
			}
		}
	}
	folder.Scanned = true
}

func (folder *Folder) GetChild(name string) *Folder {
	for _, value := range folder.Folders {
		if value.Name == name {
			return value
		}
	}
	return &Folder{}
}

func (folder *Folder) GetRoot() (root *Folder) {
	root = folder
	for {
		if root.Parent == nil {
			return
		}
		root = root.Parent
	}
}

func (folder *Folder) GetSize() (size int) {
	for _, file := range folder.Files {
		size += file.Size
	}
	for _, child := range folder.Folders {
		size += child.GetSize()
	}
	return size
}

func (folder *Folder) Step1() (result int) {
	size := folder.GetSize()
	if size < 100000 {
		result = size
	}
	for _, child := range folder.Folders {
		result += child.Step1()
	}
	return result
}

func (folder *Folder) Step2(targetSize int) (result int) {
	size := folder.GetSize()
	if size > targetSize {
		result = size
		for _, child := range folder.Folders {
			result = int(math.Min(float64(result), float64(child.Step2(targetSize))))
		}
		return result
	} else {
		return math.MaxInt
	}
}
