//    CopyRight @Ally Dale 2016
//    Author  : Ally Dale(vipally@gmail.com)
//    Blog    : http://blog.csdn.net/vipally
//    Site    : https://github.com/vipally

//FilePath related operations

package fs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	xhash "github.com/vipally/gx/hash"
)

type FileMode os.FileMode

func (me FileMode) os() os.FileMode { return os.FileMode(me) }

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

const ( //re-export consts for hash
	NONE        = xhash.NONE
	MD5         = xhash.MD5
	SHA1        = xhash.SHA1
	SHA256      = xhash.SHA256
	SHA512      = xhash.SHA512
	CRC32       = xhash.CRC32
	CRC64       = xhash.CRC64
	FINGERPRINT = xhash.FINGERPRINT
)

//const (
//	ModeRead FileMode = 1 << iota
//	ModeWrite
//	ModeTruncate
//)

//
// !!! NEVER USE THIS TYPE TO CAST STRINGS DIRECTLY !!!
//  File-path object, to deal-with filepath or Urls as string.
//  file-path object treat all paths separated with "/".
//  StringSys return system-related string format.
//  Function FilePath() is used to create this object from raw-string.
type FilePathObject string

//  create file-path object from raw-string.
//  !!!NEVER USE FilePathObject TO CAST STRINGS DIRECTLY!!!
func FilePath(strPath string) FilePathObject {
	var f FilePathObject
	f.Set(strPath)
	return f
}

//set value
func (this *FilePathObject) Set(newPath string) {
	*this = FilePathObject(filepath.ToSlash(filepath.Clean(newPath)))
}

//show as string
func (me FilePathObject) String() string {
	return string(me)
}

//short for String()
func (me FilePathObject) S() string {
	return string(me)
}

//system-related string format
func (me FilePathObject) StringSys() string {
	return filepath.FromSlash(string(me))
}

/////////////////////////////
//from std.filepath

//adapt for filepath.SplitList
func (me FilePathObject) SplitList() []string {
	return filepath.SplitList(string(me))
}

//adapt for filepath.Split
func (me FilePathObject) Split() (dir, file string) {
	return filepath.Split(string(me))
}

//adapt for filepath.Ext
func (me FilePathObject) Ext() string {
	return filepath.Ext(string(me))
}

//adapt for filepath.EvalSymlinks
func (me FilePathObject) EvalSymlinks() (string, error) {
	return filepath.EvalSymlinks(string(me))
}

//adapt for filepath.Abs
func (me FilePathObject) Abs() (string, error) {
	return filepath.Abs(string(me))
}

//adapt for filepath.Rel
func (me FilePathObject) Relate(root string) FilePathObject {
	s, _ := filepath.Rel(FilePath(root).String(), string(me))
	return FilePath(s)
}

//related from GoPath
func (me FilePathObject) RelateGoPath() FilePathObject {
	s, _ := filepath.Rel(GoPath(), string(me))
	return FilePath(s)
}

//related from current working path
func (me FilePathObject) RelateWorkPath() FilePathObject {
	s, _ := filepath.Rel(WorkPath(), string(me))
	return FilePath(s)
}

//adapt for os.FileInfo
type FileInfo os.FileInfo

//walk func, adapt for fitlepath.WalkFunc
type WalkFunc func(path string, info FileInfo, err error) error

//adapt for filepath.Walk
func (me FilePathObject) Walk(walkFn WalkFunc) error {
	return filepath.Walk(
		string(me), func(path string, info os.FileInfo, err error) error {
			return walkFn(path, FileInfo(info), err)
		})
}

//adapt for filepath.Base
func (me FilePathObject) Base() string {
	return filepath.Base(string(me))
}

//adapt for filepath.Dir
func (me FilePathObject) Dir() string {
	return filepath.Dir(string(me))
}

//adapt for filepath.VolumeName
func (me FilePathObject) VolumeName() string {
	return filepath.VolumeName(string(me))
}

//adapt for filepath.Match
func (me FilePathObject) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, string(me))
}

//adapt for filepath.HasPrefix
func (me FilePathObject) HasPrefix(prefix string) bool {
	return filepath.HasPrefix(string(me), FilePath(prefix).String())
}

////////////////////////////////
//more

//split all path elments
func (me FilePathObject) SplitAll() []string {
	s := me.String()
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

//adapt for filepath.Join
func Joins(elem ...string) string {
	s := filepath.Join(elem...)
	return FilePath(s).String()
}

//join elements after me
func (me FilePathObject) Joins(elem ...string) string {
	s := append([]string{me.String()}, elem...)
	return Joins(s...)
}

//join child after me
func (me FilePathObject) Join(child string) string {
	return Joins(me.String(), child)
}

/////////////////////////////////////////////////////////////////
//OS operations

//////////////////////////
//from std.os

//adapt for os.Chmod
func (me FilePathObject) Chmod(mod FileMode) error {
	return os.Chmod(string(me), mod.os())
}

//adapt for os.Chown
func (me FilePathObject) Chown(uid, gid int) error {
	return os.Chown(string(me), uid, gid)
}

//adapt for os.Chtimes
func (me FilePathObject) Chtimes(atime, mtime time.Time) error {
	return os.Chtimes(string(me), atime, mtime)
}

//adapt for os.Lchown
func (me FilePathObject) Lchown(uid, gid int) error {
	return os.Lchown(string(me), uid, gid)
}

//adapt for os.Link, if linkname is not absolute, it will relate from me.Dir
func (me FilePathObject) Link(linkname string) (n FilePathObject, err error) {
	n = me.relateThisPath(linkname)
	err = os.Link(string(me), n.String())
	return
}

//adapt for os.Mkdir
func (me FilePathObject) Mkdir(perm FileMode) error {
	return os.Mkdir(string(me), perm.os())
}

//adapt for os.MkdirAll
func (me FilePathObject) MkdirAll(perm FileMode) error {
	return os.MkdirAll(string(me), perm.os())
}

//adapt for os.Readlink
func (me FilePathObject) Readlink() (target string, err error) {
	return os.Readlink(string(me))
}

//adapt for os.Remove
func (me FilePathObject) Remove() error {
	return os.Remove(string(me))
}

//adapt for os.RemoveAll
//remove this dir no mater whether it is empty
func (me FilePathObject) RemoveAll() error {
	return os.RemoveAll(string(me))
}

//adapt for os.Rename, if newname is not absolute, it will relate from me.Dir
//move to another path name,but never cross disk
func (me FilePathObject) Rename(newname string) (newPath FilePathObject, err error) {
	n := me.relateThisPath(newname)
	return n, os.Rename(string(me), n.String())
}

//adapt for os.Symlink, if linkname is not absolute, it will relate from me.Dir
func (me FilePathObject) Symlink(newname string) (n FilePathObject, err error) {
	n = me.relateThisPath(newname)
	err = os.Symlink(string(me), n.String())
	return
}

//adapt for os.Truncate
func (me FilePathObject) Truncate(size int64) error {
	return os.Truncate(string(me), size)
}

//adapt for os.Create
func (me FilePathObject) Create() (*os.File, error) {
	return os.Create(string(me))
}

//adapt for os.NewFile
func (me FilePathObject) NewFile(fd uintptr) *os.File {
	return os.NewFile(fd, string(me))
}

//adapt for os.Open
func (me FilePathObject) Open() (*os.File, error) {
	return os.Open(string(me))
}

//adapt for os.OpenFile
func (me FilePathObject) OpenFile(flag int, perm FileMode) (*os.File, error) {
	return os.OpenFile(string(me), flag, perm.os())
}

//adapt for os.Lstat
func (me FilePathObject) Lstat() (os.FileInfo, error) {
	return os.Lstat(string(me))
}

//adapt for os.Stat
func (me FilePathObject) Stat() (os.FileInfo, error) {
	return os.Stat(string(me))
}

//////////////////////////////////////////////////////
//more

//if _destPath is related path, then calculate from me.Dir
func (me FilePathObject) relateThisPath(_destPath string) FilePathObject {
	n := FilePath(_destPath)
	if n.VolumeName() == "" { //related path, then calculate from me.Dir
		n.Set(Joins(me.Dir(), n.String()))
	}
	return n
}

//OS statistic infomation
func (me FilePathObject) Statistic() (nDir, nFile int, size uint64, info string) {
	var buf bytes.Buffer
	me.Walk(func(path string, info FileInfo, err error) error {
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

	//fmt.Printf("%s\n[%s] %ddir(s) %dfile(s) %s\n", info, me.StringSys(), nDir, nFile, FileSize(size))
	return
}

//file hash value
type FileHash struct {
	Path, Hash string
	Size       FileSize
	IsDir      bool
}

//sort by Size, Hash  decend,the larger one first
func (this FileHash) Less(o FileHash) bool {
	if this.Size < o.Size {
		return true
	}
	if this.Size == o.Size && this.Hash < o.Hash {
		return true
	}

	return false
}

//Fingerprint get fingerprint of a file or dir
func (me FilePathObject) Fingerprint() (root *FileHashTreeNode, err error) {
	var f *os.File
	if f, err = me.Open(); err == nil {
		defer f.Close()
		var fi os.FileInfo
		root = &FileHashTreeNode{}
		if fi, err = f.Stat(); err == nil {
			h := xhash.NewFingerprint()
			if fi.IsDir() {
				var dirs []os.FileInfo
				if dirs, err = f.Readdir(0); err == nil {
					for _, v := range dirs {
						_sub := v.Name()
						if _sub != "." && _sub != ".." {
							_subfullname := me.Join(_sub)
							var sub *FileHashTreeNode
							if sub, err = FilePath(_subfullname).Fingerprint(); err == nil {
								root.AddChildNode(sub, -1)
							} else {
								return
							}
						}
					}
					root.SortChildren()
					for _, v := range root.Children() {
						r := strings.NewReader(v.Hash)
						io.Copy(h, r)
						root.Size += v.Size
					}
					root.IsDir = true
				}
			} else {
				r := io.Reader(f)
				io.Copy(h, r)
				root.Size = FileSize(fi.Size())
			}
			h.Resetsize(int64(root.Size))
			p := h.String()
			root.Path = me.String()
			root.Hash = p
		}
	}
	return
}

//Fingerprint get hash of a file or dir by method
func (me FilePathObject) Hash(method xhash.HashMethod, salt string) (root *FileHashTreeNode, err error) {
	var f *os.File
	if f, err = me.Open(); err == nil {
		defer f.Close()
		var fi os.FileInfo
		root = &FileHashTreeNode{}
		if fi, err = f.Stat(); err == nil {
			h := method.New()
			if salt != "" {
				h.Write([]byte(salt))
			}
			if fi.IsDir() {
				var dirs []os.FileInfo
				if dirs, err = f.Readdir(0); err == nil {
					for _, v := range dirs {
						_sub := v.Name()
						if _sub != "." && _sub != ".." {
							_subfullname := me.Join(_sub)
							var sub *FileHashTreeNode
							if sub, err = FilePath(_subfullname).Hash(method, salt); err == nil {
								root.AddChildNode(sub, -1)
							} else {
								return
							}
						}
					}
					root.SortChildren()
					for _, v := range root.Children() {
						r := strings.NewReader(v.Hash)
						io.Copy(h, r)
						root.Size += v.Size
					}
					root.IsDir = true
				}
			} else {
				r := io.Reader(f)
				io.Copy(h, r)
				root.Size = FileSize(fi.Size())
			}
			p := fmt.Sprintf("%x", h.Sum(nil))
			root.Path = me.String()
			root.Hash = p
		}
	}
	return
}

//copy to destPath
func (me FilePathObject) Copy(destPath string) (FilePathObject, error) {
	//n := me.destPath(newname)
	return "", nil
}

//move to destPath
func (me FilePathObject) Move(destPath string) (FilePathObject, error) {
	n := me.relateThisPath(destPath)
	if me.VolumeName() == n.VolumeName() { //not cross disk,use rename operation
		return me.Rename(n.String())
	}
	return "", nil
}

//func (me filePath) Tree()  {
//	return me
//}
//func (me filePath) CollectSubs(opt FsOption) (subs []string, err error) {
//	return me
//}
