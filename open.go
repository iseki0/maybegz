package maybegz

import (
	"compress/bzip2"
	"compress/gzip"
	"github.com/ulikunitz/xz"
	"io"
	"os"
	"strings"
)

type S struct {
	f *os.File
	r io.Reader
}

func (s *S) Read(p []byte) (n int, err error) {
	return s.r.Read(p)
}

func (s *S) Close() error {
	return s.f.Close()
}

func Open(path string) (io.ReadCloser, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	s := &S{f: f, r: f}
	if strings.HasSuffix(path, ".gz") {
		s.r, e = gzip.NewReader(s.f)
		if e != nil {
			_ = s.Close()
			return nil, e
		}
		return s, nil
	}
	if strings.HasSuffix(path, ".bz2") {
		s.r = bzip2.NewReader(s.f)
		return s, nil
	}
	if strings.HasSuffix(path, ".xz") {
		s.r, e = xz.NewReader(s.f)
		if e != nil {
			_ = s.Close()
			return nil, e
		}
		return s, nil
	}
	return s, nil
}
