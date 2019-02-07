package procmaps

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanSimple(t *testing.T) {
	r := strings.NewReader(`561534781000-561534783000 r--p 00000000 fe:00 18088321                   /usr/bin/cat
`)
	m, err := Scan(r)

	require.NoError(t, err)
	require.Equal(t, Mapping{
		Start:   0x561534781000,
		End:     0x561534783000,
		Read:    true,
		Private: true,
		Offset:  0x00000000,
		Device:  "fe:00",
		Inode:   18088321,
		Path:    "/usr/bin/cat",
	}, m)
}

func TestScanMalformed(t *testing.T) {
	r := strings.NewReader(`561534781000-561534783000 r--p 00000000 fe:00`)
	_, err := Scan(r)

	require.Error(t, err)
}

func TestScanNoPath(t *testing.T) {
	r := strings.NewReader(`561534781000-561534783000 r--p 00000000 fe:00 18088321
`)
	m, err := Scan(r)

	require.NoError(t, err)
	require.Equal(t, Mapping{
		Start:   0x561534781000,
		End:     0x561534783000,
		Read:    true,
		Private: true,
		Offset:  0x00000000,
		Device:  "fe:00",
		Inode:   18088321,
	}, m)
}

func TestScanAll(t *testing.T) {
	r := strings.NewReader(`561534781000-561534783000 r--p 00000000 fe:00 18088321                   /usr/bin/cat
7f9ce6308000-7f9ce630a000 rw-p 001bd000 fe:00 18096430                   /usr/lib/libc-2.28.so
7f9ce630a000-7f9ce6310000 rw-p 00000000 00:00 0 
ffffffffff600000-ffffffffff601000 r-xp 00000000 00:00 0                  [vsyscall]
`)
	m, err := ScanAll(r)

	require.NoError(t, err)
	require.Equal(t, []Mapping{
		{
			Start:   0x561534781000,
			End:     0x561534783000,
			Read:    true,
			Private: true,
			Offset:  0x00000000,
			Device:  "fe:00",
			Inode:   18088321,
			Path:    "/usr/bin/cat",
		},
		{
			Start:   0x7f9ce6308000,
			End:     0x7f9ce630a000,
			Read:    true,
			Write:   true,
			Private: true,
			Offset:  0x001bd000,
			Device:  "fe:00",
			Inode:   18096430,
			Path:    "/usr/lib/libc-2.28.so",
		},
		{
			Start:   0x7f9ce630a000,
			End:     0x7f9ce6310000,
			Read:    true,
			Write:   true,
			Private: true,
			Offset:  0x00000000,
			Device:  "00:00",
			Inode:   0,
		},
		{
			Start:   0xffffffffff600000,
			End:     0xffffffffff601000,
			Read:    true,
			Execute: true,
			Private: true,
			Offset:  0x00000000,
			Device:  "00:00",
			Inode:   0,
			Path:    "[vsyscall]",
		},
	}, m)
}
