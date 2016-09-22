package fs

//file path object, to differentiate from common string
type FilePath string

func ToFilePath(str string) FilePath {
	var f FilePath
	f.Set(str)
	return f
}

func (this *FilePath) Set(str string) {
	*this = FilePath(FormatPath(str))
}

func (this FilePath) String() string {
	return string(this)
}

func (this FilePath) Dir() FilePath {
	return this
}

func (this FilePath) FileName() FilePath {
	return this
}
