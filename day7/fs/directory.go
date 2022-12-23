package fs

import (
	"strings"

	"golang.org/x/exp/maps"
)

type Directory struct {
	name    string
	content map[string]FSData
	parent  *Directory

	calculatedSize int
}

func RootDir() *Directory {
	return &Directory{
		name:    "/",
		content: make(map[string]FSData),
		parent:  nil,
	}
}

func NewDir(name string, parent *Directory) *Directory {
	dir := &Directory{
		name:           name,
		content:        make(map[string]FSData),
		parent:         parent,
		calculatedSize: 0,
	}

	parent.content[dir.name] = dir

	return dir
}

func NewDirFromOutput(o string, parent *Directory) *Directory {
	params := strings.Split(o, " ")
	name := params[1]
	return NewDir(name, parent)
}

func (d *Directory) AddData(fd FSData) {
	d.content[fd.Name()] = fd
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	if d.calculatedSize == 0 {
		d.calculate()
	}
	return d.calculatedSize
}

func (d *Directory) calculate() {
	d.calculatedSize = 0
	for _, fd := range d.content {
		d.calculatedSize += fd.Size()
	}
}

func (d *Directory) Content() []FSData {
	return maps.Values(d.content)
}

func (d *Directory) Parent() *Directory {
	return d.parent
}

func (d *Directory) Print(level int) string {
	output := d.name
	for _, v := range d.content {
		output += "\n" + v.Print(level+1)
	}
	return generateTab(level) + output
}
