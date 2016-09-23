package fs

import "fmt"

const (
	kBits = 10
	kb    = 1 << kBits
)

var bytesNames = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"} //max16EB ignore the next: "ZB", "YB", "BB"

//file-size object
type FileSize uint64

//show file-size
func (me FileSize) String() string {
	m64 := uint64(me)
	i, b := 0, uint64(0)
	for ; i < len(bytesNames); i, b = i+1, b+kBits {
		if (m64 >> b) < kb {
			break
		}
	}
	m := 1 << b
	d := float64(me) / float64(m)
	return fmt.Sprintf("%.3f%s", d, bytesNames[i])
}

//to uint64
func (me FileSize) N() uint64 {
	return uint64(me)
}

//from uint64
func (this *FileSize) Set(val uint64) {
	*this = FileSize(val)
}
