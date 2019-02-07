package procmaps

import (
	"fmt"
	"os"
)

func ReadSelf() ([]Mapping, error) {
	f, err := os.Open("/proc/self/maps")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ScanAll(f)
}

func Read(pid int) ([]Mapping, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/maps", pid))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ScanAll(f)
}
