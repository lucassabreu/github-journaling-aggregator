// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package formatter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	mustUnmarshalTextTime := func(text string) time.Time {
		var t time.Time
		err := t.UnmarshalText([]byte(text))
		if err != nil {
			panic(err)
		}
		return t
	}

	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: mustUnmarshalTextTime("2017-09-24T12:12:26.96547426Z"),
		},
		"/html.html": &vfsgen۰CompressedFileInfo{
			name:             "html.html",
			modTime:          mustUnmarshalTextTime("2017-09-24T12:37:37.763949483Z"),
			uncompressedSize: 1631,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x3d\x53\x1b\x3d\x10\xc7\x7b\x7f\x0a\xcd\xb5\xcf\x48\x02\x9e\x26\x43\xe4\x4b\x45\x91\x82\x14\x99\xf4\x37\x42\x5a\xdf\x09\xf4\x72\x23\xad\x01\x8f\xc7\xdf\x3d\xa3\x17\xc3\x19\x5c\x90\x90\x6e\x6f\x77\xf5\x93\xfe\x2b\xed\xde\x4a\x4c\xe8\x6c\xbf\x12\x13\x48\xdd\xaf\x08\x11\x68\xd0\x42\xbf\xdf\x13\xf6\x2b\x5b\xe4\x70\x10\xbc\xfa\x72\xd4\x1a\xff\x40\x22\xd8\x75\x97\x70\x67\x21\x4d\x00\xd8\x91\x29\xc2\x66\xdd\x4d\x88\x73\xba\xe6\x7c\x13\x3c\x26\x36\x86\x30\x5a\x90\xb3\x49\x4c\x05\xc7\x8d\x0a\xfe\xdb\x46\x3a\x63\x77\xeb\x5b\x89\x10\x8d\xb4\xff\x7d\x57\xc1\xa7\xee\xc3\x60\x15\x34\xb0\x11\xd0\x69\xcb\x4c\xe0\x97\xec\x7f\x76\xc1\x5d\x83\x31\xe3\xb5\x19\x03\x9d\x8d\x7f\x60\xce\x78\xa6\x52\x43\x17\x60\xb6\x08\x41\x79\x67\x81\xa0\x66\x0e\x52\x92\x23\x90\x7d\x71\x13\xf2\x34\x19\x04\x9a\x66\xa9\xe0\x9a\xf8\x10\x9d\xb4\x5f\x4b\xe8\x90\x09\xbc\x21\x04\xaf\x65\x12\x77\x41\xef\xfa\x95\xd0\xe6\x91\x28\x2b\x53\x5a\x77\x4e\x5b\x6a\xe5\x2e\x6c\x91\xbc\x9a\x94\x6e\xcc\x33\x68\x9a\x57\x41\x2c\x81\xfb\xb4\x4c\x53\xc1\x86\x48\xe9\x18\x61\x47\x2f\x2f\x2e\xea\x79\x5b\xf6\x3b\xf0\x30\x2c\x38\xa7\x2e\x4a\x93\x8a\xc1\xda\x73\xd0\x85\x0f\xe1\x19\x5b\xe0\x4b\xdb\x8d\x90\xf3\x2a\x5e\xc8\x31\x3c\xb5\xc4\x5c\xca\x59\xfa\xf7\xb9\xb4\xbc\x8f\xee\xcd\xa3\xc9\xb9\x6d\x07\xae\xcd\x63\x91\xc6\x2b\xb4\xd8\x4e\x9a\x33\xac\x61\x50\xc1\x23\x78\x3c\x9e\x6e\x6b\x4f\x92\x4c\x6a\x95\x83\xa3\x5a\xb0\x96\xd2\xcb\xab\x2c\xf1\xe5\xa0\xfb\x3d\x89\xd2\x8f\x40\xd8\x4d\x8c\x21\x26\x72\x38\xb4\x48\x79\x6a\x6f\x89\xc3\x60\x10\x5c\x15\x50\x16\x14\x01\xd6\x2c\x70\xe0\xf5\x11\x22\xf8\xd6\xf6\xab\x6a\xd6\x07\xb5\xc0\x69\x89\x92\x56\x6f\xbb\xee\x37\x9e\x34\x49\x1d\x9e\x28\xbd\xd2\xf3\x07\x84\x08\x3c\x36\xe6\xcb\xf7\xf9\xdd\x86\xa1\xae\xf7\xc1\x53\xbf\x75\x10\x8d\xea\xfa\x9b\x47\xf0\x48\x7e\x48\x07\x82\xe3\xf4\xb7\x94\x9f\x30\x87\x4f\x43\x6e\x6b\xc3\x2d\x11\xd9\x5e\x88\x13\x58\xbb\xea\xc8\x7f\xbd\xc3\xb6\xf6\xe4\x16\x73\x7e\xec\x17\x9f\xd9\xa1\xff\xe0\x40\xe5\xae\x73\x7d\xb2\xb2\x3a\xe5\xf4\x67\x79\xb9\x52\xec\x9f\xf0\x48\x1b\x50\x95\xdb\x0a\x70\x86\x2a\xf8\xb2\x0a\xa7\xef\xb4\x44\x5f\x6b\x2a\x78\xd9\xac\xf6\x61\x6e\xbe\x3c\xcf\x4a\x63\x8a\xa4\xa2\x99\x91\x68\xd8\x40\x24\x29\xaa\x8f\x8e\xdc\x3c\x66\xef\x53\xd7\x0b\x5e\x09\x99\xd8\x46\x23\x2f\xff\x95\xdf\x01\x00\x00\xff\xff\x2a\xaa\x8e\x70\x5f\x06\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/html.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}