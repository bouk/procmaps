package procmaps

import (
	"fmt"
	"io"
)

type Mapping struct {
	Start   uintptr
	End     uintptr
	Read    bool
	Write   bool
	Execute bool
	Private bool
	Offset  uintptr
	Device  string
	Inode   uintptr
	Path    string
}

func Scan(r io.Reader) (m Mapping, err error) {
	var n int
	var read, write, execute, private rune
	n, err = fmt.Fscanf(r, "%x-%x %c%c%c%c %x %s %d %s\n", &m.Start, &m.End, &read, &write, &execute, &private, &m.Offset, &m.Device, &m.Inode, &m.Path)
	if err != nil && n == 9 && (err.Error() == "unexpected newline" || err.Error() == "newline in input does not match format") {
		err = nil
	}
	m.Read = read == 'r'
	m.Write = write == 'w'
	m.Execute = execute == 'x'
	m.Private = private == 'p'
	return
}

func ScanAll(r io.Reader) (mappings []Mapping, err error) {
	for {
		m, err := Scan(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		mappings = append(mappings, m)
	}
	return
}
