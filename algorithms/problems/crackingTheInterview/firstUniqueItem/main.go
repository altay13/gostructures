package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/altay13/gostructures/hashtable"

	"github.com/altay13/gostructures/set"
)

type TestReader struct {
	rd io.Reader
	s  *set.Set
	h  *hashtable.Hash
	l  []int
}

func NewTestReader(r io.Reader) *TestReader {
	return &TestReader{
		rd: r,
		s:  set.NewSet(math.MaxInt32),
		h:  hashtable.NewHash(math.MaxInt32),
		l:  make([]int, 0),
	}
}

func (r *TestReader) Read(p []byte) (n int, err error) {
	return r.rd.Read(p)
}

func (r *TestReader) GetFirstUnique() int {
	for i := 0; i < len(r.l); i++ {
		if r.l[i] != -1 {
			return r.l[i]
		}
	}

	return -1
}

func main() {
	rdr := NewTestReader(strings.NewReader("TESTing"))

	// os.File

	var out []byte
	for {
		n, err := rdr.Read(out)
		if err != nil {
			if err == io.EOF {
				os.Exit(1)
			}
		}
		fmt.Println(string(out))
		if n == 0 {
			os.Exit(1)
		}
	}
}
