// https://go-tour-jp.appspot.com/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i, v := range b {
		switch {
		case 'A' <= v && v <= 'Z'-13, 'a' <= v && v <= 'z'-13:
			b[i] += 13
		case 'z'-13 < v && v <= 'Z', 'z'-13 < v && v <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// >result
// You cracked the code!
