package fs

type FSData interface {
	Name() string
	Size() int
	Parent() *Directory
	Print(level int) string
}
