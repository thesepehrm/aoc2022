package fs

import (
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	size   int
	name   string
	parent *Directory
}

func NewFileFromOutput(output string, parent *Directory) *File {
	f := parseFileFromOutput(output)
	f.parent = parent
	f.parent.content[f.name] = f
	return f
}

func parseFileFromOutput(output string) *File {
	params := strings.Split(output, " ")
	name := params[1]
	size, _ := strconv.Atoi(params[0])
	return &File{
		name:   name,
		size:   size,
		parent: nil,
	}
}

func (f File) Size() int {
	return f.size
}

func (f File) Name() string {
	return f.name
}

func (f File) Parent() *Directory {
	return f.parent
}

func (f File) Print(level int) string {
	return generateTab(level) + fmt.Sprintf("%d %s", f.size, f.name)
}
