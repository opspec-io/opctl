// Code generated by "esc -pkg=opfile -o validator_schema.go -private ../../../../opspec/opfile/jsonschema.json"; DO NOT EDIT.

package opfile

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/opspec/opfile/jsonschema.json": {
		name:    "jsonschema.json",
		local:   "../../../../opspec/opfile/jsonschema.json",
		size:    42783,
		modtime: 1587241275,
		compressed: `
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
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}