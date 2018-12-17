package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		return b + 'N' - 'A'
	} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		return b - 'N' + 'A'
	} else {
		return b
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, error := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
