package design_pattern

import (
	"fmt"
)

type component interface {
	Search(string)
}

type Folder struct {
	components []component
	Name       string
}

func (f *Folder) AddComponent(c component) {
	f.components = append(f.components, c)
}

func (f *Folder) Search(str string) {
	if len(f.components) == 0 {
		fmt.Printf("search %s in folder %s\n", str, f.Name)
	}
	for _, component := range f.components {
		component.Search(str)
	}
}

type File struct {
	Name string
}

func (f *File) Search(str string) {
	fmt.Printf("search %s in file %s\n", str, f.Name)
}
