package fs

import (
	"os"
	"path"
	"strings"
)

const (
	FsDir  FsOption = 1 << iota //dir
	FsFile                      //file

	FsFullPath //full path
	FsDeep     //deep collect
	FsLeaf     //leaf items

	FsNone = FsOption(0)
	FsAll  = FsDir | FsFile
)

type FsOption uint

func (me FsOption) has(opt FsOption) bool {
	return opt != 0 && me&opt == opt
}
func (me FsOption) has_any(opt FsOption) bool {
	return opt != 0 && me&opt != 0
}
func (me FsOption) choose_name(full_name, name string) string {
	if me.has(FsFullPath) {
		return full_name
	} else {
		return name
	}
}

//to format "\\" with "/"
func FormatPath(filepath string) string {
	return strings.Replace(filepath, "\\", "/", -1)
}

func GetLeafDirs(dir string) (subs []string, err error) {
	return CollectSubs(dir, FsDir|FsLeaf|FsFullPath)
}

func GetSubDirs(dir string) (subs []string, err error) {
	return CollectSubs(dir, FsDir)
}

func GetSubFiles(dir string) (subs []string, err error) {
	return CollectSubs(dir, FsFile)
}

func GetSubs(dir string) (subs []string, err error) {
	return CollectSubs(dir, FsAll)
}

func GetDeepFiles(dir string) (subs []string, err error) {
	return CollectSubs(dir, FsFile|FsDeep)
}

func CollectSubs(dir string, opt FsOption) (subs []string, err error) {
	if _f, err := os.Open(dir); err == nil {
		defer _f.Close()
		if _subs, err := _f.Readdir(0); err == nil {
			_has_sub_dir := false
			for _, _v := range _subs {
				_sub_name := _v.Name()
				_fullsubname := Join(dir, _sub_name)

				if _v.IsDir() {
					if _sub_name == "." || _sub_name == ".." {
						continue
					}
					_has_sub_dir = true

					if opt.has(FsDir) && !opt.has(FsLeaf) { //if contain a sub dir, it's not leaf dir
						subs = append(subs, opt.choose_name(_fullsubname, _sub_name))
					}
					if opt.has_any(FsDeep | FsLeaf) {
						if _r, err := CollectSubs(_fullsubname, opt); err == nil {
							subs = append(subs, _r...)
						} else {
							return nil, err
						}
					}
				} else {
					if opt.has(FsFile) {
						subs = append(subs, opt.choose_name(_fullsubname, _sub_name))
					}
				}
			}
			if !_has_sub_dir && opt.has(FsLeaf|FsDir) {
				subs = append(subs, opt.choose_name(dir, FileName(dir)))
			}
		}
	}
	return
}

func Create(_dir string) error {
	_parent := _dir //Dir(_dir)
	if f, e := os.Open(_dir); e == nil {
		if e2 := Create(_parent); e2 == nil {
			_, e2 = os.Create(_parent)
			return e2
		} else {
			return e2
		}
	} else {
		f.Close()
	}
	return nil
}

func Split(file_path string) (_path, _file string) {
	return path.Split(file_path)
}

func Path(file_path string) (path string) {
	path, _ = Split(file_path)
	return
}

func FileName(file_path string) (file string) {
	_, file = Split(file_path)
	return
}
func FileBase(file_path string) (file string) {
	full := FileName(file_path)
	ext := FileExt(file_path)
	file = strings.TrimSuffix(full, ext)
	return
}

func FileExt(file_path string) (ext string) {
	return path.Ext(file_path)
}

func Dir(file_path string) (dir string) {
	return path.Dir(file_path)
}

func IsAbs(file_path string) (b bool) {
	if strings.HasPrefix(file_path, "/") || strings.Contains(file_path, ":/") { //windows D:/... linux /root/...
		return true
	}
	return path.IsAbs(file_path)
}

func Join(elem ...string) (file_path string) {
	return path.Join(elem...)
}

func Match(pattern, name string) (matched bool, err error) {
	return path.Match(pattern, name)
}

func WorkDir() (dir string) {
	dir, _ = os.Getwd()
	return
}
