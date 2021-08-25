package main

import "fmt"

type inode interface {
	print(indentation string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

func NewFile(name string) *file {
	return &file{name: name}
}

type folder struct {
	name string
	children []inode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.children {
		v.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode

	for _, v := range f.children {
		cloned := v.clone()
		tempChildren = append(tempChildren, cloned)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func NewFolder(name string) *folder {
	return &folder{name: name}
}

func main() {
	file1 := NewFile("File1")
	file2 := NewFile("File2")
	file3 := NewFile("File3")

	folder1 := NewFolder("Folder1")
	folder1.children = []inode{file1}

	folder2 := NewFolder("Folder2")
	folder2.children = []inode{folder1, file2, file3}

	folder2.print("  ")

	clonedFolder := folder2.clone()
	clonedFolder.print("  ")
}