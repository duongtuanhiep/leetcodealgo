package filesystem

// Interface
type metadata interface {
	getName() string
	isDirectory() bool
	getSize() int
}

// File
type File struct {
	name string
	size int
}

func (f *File) getName() string {
	return f.name
}

func (f *File) isDirectory() bool {
	return false
}

// Folder
type Folder struct {
	items []metadata
	name  string
	size  int
}

func (f *Folder) getName() string {
	return f.name
}

func (f *Folder) isDirectory() bool {
	return true
}

// Mock a file system
type FS struct {
}

func (fs *FS) goToDir(dir string) *Folder {
	return nil
}

type matcherFunc func(metadata) bool

func matchName(name string) func(metadata) bool {
	return func(m metadata) bool {
		return m.getName() == name
	}
}

func matchSize(size int) func(metadata) bool {
	return func(m metadata) bool {
		return m.getSize() == size
	}
}

// FindFile(dir, []matcherFunc{matchName, matchSize})
func (fs *FS) FindFile(dir string, matchers []matcherFunc) []metadata {
	item := fs.goToDir(dir)
	if item == nil {
		return nil
	}

	res := []metadata{}
	stack := []metadata{item}
	for len(stack) > 0 {
		items := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, item := range items {
			if item.isDirectory() {
				stack = append(stack, item)
			}

			valid = true
			for _, match := range matchers {
				if !match(item) {
					valid = false
				}
			}

			if valid {
				res = append(res, item)
			}
		}
	}

	return res
}
