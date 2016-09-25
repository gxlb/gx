//    CopyRight @Ally Dale 2016
//    Author  : Ally Dale(vipally@gmail.com)
//    Blog    : http://blog.csdn.net/vipally
//    Site    : https://github.com/vipally

//FilePath related operations

package fs

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileMode os.FileMode

func (this FileMode) os() os.FileMode { return os.FileMode(this) }

const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        = FileMode(os.ModeDir)        // d: is a directory
	ModeAppend     = FileMode(os.ModeAppend)     // a: append-only
	ModeExclusive  = FileMode(os.ModeExclusive)  // l: exclusive use
	ModeTemporary  = FileMode(os.ModeTemporary)  // T: temporary file (not backed up)
	ModeSymlink    = FileMode(os.ModeSymlink)    // L: symbolic link
	ModeDevice     = FileMode(os.ModeDevice)     // D: device file
	ModeNamedPipe  = FileMode(os.ModeNamedPipe)  // p: named pipe (FIFO)
	ModeSocket     = FileMode(os.ModeSocket)     // S: Unix domain socket
	ModeSetuid     = FileMode(os.ModeSetuid)     // u: setuid
	ModeSetgid     = FileMode(os.ModeSetgid)     // g: setgid
	ModeCharDevice = FileMode(os.ModeCharDevice) // c: Unix character device, when ModeDevice is set
	ModeSticky     = FileMode(os.ModeSticky)     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = FileMode(os.ModeType)
	ModePerm = FileMode(os.ModePerm) // Unix permission bits
)

//const (
//	ModeRead FileMode = 1 << iota
//	ModeWrite
//	ModeTruncate
//)

//file path object, to differentiate from common string
//file-path treat all paths separated with "/"
//StringSys return system-related string format
type filePath string

//to file-path object
func FilePath(str string) filePath {
	var f filePath
	f.Set(str)
	return f
}

//set value
func (this *filePath) Set(newPath string) {
	*this = filePath(filepath.ToSlash(filepath.Clean(newPath)))
}

//show as string
func (this filePath) String() string {
	return string(this)
}

//system-related string format
func (this filePath) StringSys() string {
	return filepath.FromSlash(string(this))
}

/////////////////////////////
//from std.filepath

func (this filePath) SplitList() []string {
	return filepath.SplitList(string(this))
}

func (this filePath) Split() (dir, file string) {
	return filepath.Split(string(this))
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
	s, _ := filepath.Rel(FilePath(root).String(), string(this))
	return FilePath(s)
}

func (this filePath) RelateGoPath() filePath {
	s, _ := filepath.Rel(GoPath(), string(this))
	return FilePath(s)
}

func (this filePath) RelateWorkPath() filePath {
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
	return filepath.HasPrefix(string(this), FilePath(prefix).String())
}

////////////////////////////////
//more
func (this filePath) SplitAll() []string {
	s := this.String()
	maxn := strings.Count(s, "/") + 1
	b := make([]string, maxn, maxn)
	i := maxn - 1
	for ; i >= 0; i-- {
		p, f := filepath.Split(s)
		s = strings.TrimSuffix(p, "/")
		if f != "" {
			b[i] = f
		} else {
			if p != "" {
				b[i] = p
			} else {
				i++
			}
			break
		}
	}
	return b[i:]
}

func Joins(elem ...string) string {
	s := filepath.Join(elem...)
	return FilePath(s).String()
}

func (this filePath) Joins(elem ...string) string {
	s := append([]string{this.String()}, elem...)
	return Joins(s...)
}

func (this filePath) Join(child string) string {
	return Joins(this.String(), child)
}

/////////////////////////////////////////////////////////////////
//OS operations

//////////////////////////
//from std.os

func (this filePath) Chmod(mod FileMode) error {
	return os.Chmod(string(this), mod.os())
}
func (this filePath) Chown(uid, gid int) error {
	return os.Chown(string(this), uid, gid)
}
func (this filePath) Chtimes(atime, mtime time.Time) error {
	return os.Chtimes(string(this), atime, mtime)
}
func (this filePath) Lchown(uid, gid int) error {
	return os.Lchown(string(this), uid, gid)
}
func (this filePath) Link(linkname string) (n filePath, err error) {
	n = this.destPath(linkname)
	err = os.Link(string(this), n.String())
	return
}
func (this filePath) Mkdir(perm FileMode) error {
	return os.Mkdir(string(this), perm.os())
}

//create full path dir
func (this filePath) MkdirAll(perm FileMode) error {
	return os.MkdirAll(string(this), perm.os())
}
func (this filePath) Readlink() (target string, err error) {
	return os.Readlink(string(this))
}

//remove this file or dir
func (this filePath) Remove() error {
	return os.Remove(string(this))
}

//remove this dir no mater whether it is empty
func (this filePath) RemoveAll() error {
	return os.RemoveAll(string(this))
}

//move to another path name,but never cross disk
func (this filePath) Rename(newname string) (newPath filePath, err error) {
	n := this.destPath(newname)
	return n, os.Rename(string(this), n.String())
}

func (this filePath) Symlink(newname string) (n filePath, err error) {
	n = this.destPath(newname)
	err = os.Symlink(string(this), n.String())
	return
}

//resize this file
func (this filePath) Truncate(size int64) error {
	return os.Truncate(string(this), size)
}

func (this filePath) Create() (*os.File, error) {
	return os.Create(string(this))
}

func (this filePath) NewFile(fd uintptr) *os.File {
	return os.NewFile(fd, string(this))
}

func (this filePath) Open() (*os.File, error) {
	return os.Create(string(this))
}

func (this filePath) OpenFile(flag int, perm FileMode) (*os.File, error) {
	return os.OpenFile(string(this), flag, perm.os())
}

func (this filePath) Lstat() (os.FileInfo, error) {
	return os.Lstat(string(this))
}

func (this filePath) Stat() (os.FileInfo, error) {
	return os.Stat(string(this))
}

//////////////////////////////////////////////////////
//more

//if _destPath is related path, then calculate from this.Dir
func (this filePath) destPath(_destPath string) filePath {
	n := FilePath(_destPath)
	if n.VolumeName() == "" { //related path, then calculate from this.Dir
		n.Set(Joins(this.Dir(), n.String()))
	}
	return n
}

//OS statistic infomation
func (this filePath) Statistic() (nDir, nFile int, size uint64, info string) {
	var buf bytes.Buffer
	this.Walk(func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if info.IsDir() {
				nDir++
			} else {
				nFile++
				size += uint64(info.Size())
			}
			//fmt.Println(path, info.Name(), info.Size(), info.IsDir())
		} else {
			buf.WriteString(err.Error())
			buf.WriteByte('\n')
			//fmt.Println(err)
		}
		return nil
	})
	info = buf.String()

	//fmt.Printf("%s\n[%s] %ddir(s) %dfile(s) %s\n", info, this.StringSys(), nDir, nFile, FileSize(size))
	return
}

//copy to destPath
func (this filePath) Copy(destPath string) (filePath, error) {
	//n := this.destPath(newname)
	return "", nil
}

//move to destPath
func (this filePath) Move(destPath string) (filePath, error) {
	n := this.destPath(destPath)
	if this.VolumeName() == n.VolumeName() { //not cross disk,use rename operation
		return this.Rename(n.String())
	}
	return "", nil
}

//calculate file hash
func (this filePath) Hash(method, salt string) string {
	return ""
}

//func (this filePath) Tree()  {
//	return this
//}
//func (this filePath) CollectSubs(opt FsOption) (subs []string, err error) {
//	return this
//}
