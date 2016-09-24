package fs

import (
	"os"
	"path"
	"path/filepath"
)

type FileMode uint

const (
	ModeRead FileMode = 1 << iota
	ModeWrite
	ModeTruncate
)

//file path object, to differentiate from common string
//file-path treat all paths separated with "/"
//StringSys return system-related string format
type filePath string

//file-path object
func FilePath(str string) filePath {
	var f filePath
	f.Set(str)
	return f
}

//set value
func (this *filePath) Set(str string) {
	*this = filePath(path.Clean(filepath.ToSlash(str)))
}

//show string
func (this filePath) String() string {
	return string(this)
}

//system-related string format
func (this filePath) StringSys() string {
	return filepath.FromSlash(string(this))
}

func (this filePath) SplitList() []string {
	return filepath.SplitList(string(this))
}

func (this filePath) Split() (dir, file string) {
	return filepath.Split(string(this))
}

func Join(elem ...string) string {
	s := filepath.Join(elem...)
	return FormatPath(s)
}

func (this filePath) Join(elem ...string) string {
	return Join(elem...)
}

func (this filePath) Ext() string {
	return filepath.Ext(string(this))
}

func (this filePath) EvalSymlinks() (string, error) {
	return filepath.EvalSymlinks(string(this))
}

func (this filePath) Abs() (string, error) {
	return filepath.Abs(string(this))
}

func (this filePath) Relate(root string) filePath {
	s, _ := filepath.Rel(FormatPath(root), string(this))
	return FilePath(s)
}

func (this filePath) RelateGoPath() filePath {
	s, _ := filepath.Rel(GoPath(), string(this))
	return FilePath(s)
}

func (this filePath) RelateWorkPath(root string) filePath {
	s, _ := filepath.Rel(WorkPath(), string(this))
	return FilePath(s)
}

func (this filePath) Walk(walkFn filepath.WalkFunc) error {
	return filepath.Walk(string(this), walkFn)
}

func (this filePath) Base() string {
	return filepath.Base(string(this))
}

func (this filePath) Dir() string {
	return filepath.Dir(string(this))
}

func (this filePath) VolumeName() string {
	return filepath.VolumeName(string(this))
}

func (this filePath) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, string(this))
}

func (this filePath) HasPrefix(prefix string) bool {
	return filepath.HasPrefix(string(this), FormatPath(prefix))
}

/////////////////////////////////////////////////////////////////

//func (this filePath) Tree()  {
//	return this
//}
//func (this filePath) CollectSubs(opt FsOption) (subs []string, err error) {
//	return this
//}

//OS operation
func (this filePath) Open() {
}
func (this filePath) Create() {
}
func (this filePath) Delete() {
}
func (this filePath) Rename() {
}

func (this filePath) CreateAll(perm os.FileMode) error {
	return os.MkdirAll(string(this), perm)
}
