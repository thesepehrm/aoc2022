package fs

type FileSystem struct {
	Root       *Directory
	CurrentDir *Directory
}

func NewFileSystem() *FileSystem {
	rootDir := RootDir()

	return &FileSystem{
		Root:       rootDir,
		CurrentDir: rootDir,
	}
}

func (f *FileSystem) ChangeDir(name string) {
	if name == ".." {
		f.Back()
		return
	}
	if name == "." {
		return
	}
	if name == "/" {
		f.CurrentDir = f.Root
		return
	}
	if dir, ok := f.CurrentDir.content[name].(*Directory); ok {

		f.CurrentDir = dir
	} else {
		println(name + "is not a directory")
	}

}
func (f *FileSystem) Back() {
	if f.CurrentDir.Parent() == nil {
		return
	}

	f.CurrentDir = f.CurrentDir.Parent()
}
