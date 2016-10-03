//package io extends std.io
package io

import (
	"io"
)

//copy src to multi dsts
func Copys(src io.Reader, dsts ...io.Writer) (written int64, err error) {
	return CopysBuffer(src, nil, dsts...)
}

//copy src to multi dsts
func CopysBuffer(src io.Reader, buf []byte, dsts ...io.Writer) (written int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		for _, dst := range dsts {
			if written, err = wt.WriteTo(dst); err != nil {
				break
			}
		}
		return
	}

	if buf == nil {
		buf = make([]byte, 32*1024)
	}
ReadAll:
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := int(0), error(nil)
			for _, dst := range dsts {
				if nw, ew = dst.Write(buf[0:nr]); ew != nil {
					err = ew
					break ReadAll
				}
			}
			if nw > 0 {
				written += int64(nw)
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}

		if er != nil {
			if er != io.EOF { //ignore EOF
				err = er
			}
			break
		}
	}
	return
}
