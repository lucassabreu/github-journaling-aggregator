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
			modTime: mustUnmarshalTextTime("2017-09-26T00:37:31.905599487Z"),
		},
		"/.html.html.swp": &vfsgen۰CompressedFileInfo{
			name:             ".html.html.swp",
			modTime:          mustUnmarshalTextTime("2017-09-26T00:37:40.055630649Z"),
			uncompressedSize: 16384,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9a\x4b\x6c\x1c\x49\x19\xc7\xff\x9e\xbc\xfd\xe8\x01\x56\x48\x0b\x17\x8a\x01\xc4\x9a\xa5\xa7\xc7\x89\x40\x91\x69\x0f\x0a\x61\xa5\xb5\x48\x42\xe4\x38\x2b\x38\x59\xe5\xee\x9a\x99\x72\xaa\xbb\x3a\x55\xd5\x76\x66\x27\xc3\x4a\x5c\xe1\xc6\x09\x2e\x0b\x57\x56\x02\xad\xc4\x85\x03\xe2\xc4\x01\xb1\x20\x0e\x1c\x10\x5c\x40\x08\x09\xb4\x88\x03\x17\x40\x02\x81\xfa\x31\x4e\xcf\x4c\xdb\xf1\x7a\x97\x04\x44\xfd\x0e\x6e\x77\xd5\x57\xdf\xab\xbe\xae\x99\xaa\x9e\xdd\xce\x4b\x9b\x37\xc9\xd5\x76\x07\x00\xde\x05\x7c\xee\x5b\x3f\xfe\xe2\xef\x5e\xda\xc6\x9b\xcf\x02\x22\x0d\xa8\xc6\xe3\xb9\xf5\xf9\x6d\xf7\xc6\xdd\xeb\xd7\xee\x3c\x46\xee\x4b\xb9\x42\xef\x40\xaa\x7b\x3a\xa1\x01\xf3\xfa\x52\xd0\xb8\xef\x69\x15\x78\x7d\x6e\x06\xe9\x6e\x3b\x90\x91\x97\x4b\x69\xba\xab\x58\x5a\x36\xbb\x7b\x32\x55\x31\x15\x3c\xee\xbb\xb4\xdf\x57\xac\x4f\x8d\x54\x5e\x4f\xaa\x88\x1a\xc3\x94\x47\xb5\x66\x46\x7b\x03\x13\x89\x76\xf6\xe7\x04\x4e\x5b\x2c\xff\x4f\xa4\xa6\xe7\x5e\x5d\xc1\x95\xcb\x6b\xf9\xa3\xfe\xa1\xd6\x07\xc9\x33\xef\xb9\xfb\xb4\xbd\xb2\x58\x2c\x16\x8b\xc5\x62\xb1\x58\x2c\x4f\x10\x93\x34\xf0\x0a\x80\x46\x79\xff\xd9\xf2\xba\x30\x73\x3d\x53\x5e\x9d\xf2\x7a\x6d\xa6\xdf\x62\xb1\x58\x2c\x16\x8b\xc5\x62\xb1\x58\x2c\x16\xcb\x7f\x2f\x34\x04\xfe\x0c\xe0\xfd\x8d\xe2\xfd\xff\x64\xff\xff\xf7\x26\xf0\x97\x26\xf0\x8b\x26\x30\x6c\x02\xef\x6d\x02\x4b\x4d\xe0\x6f\x0e\xf0\x5b\x07\xf8\xb5\x03\xbc\xe1\x00\xaf\x3b\xc0\x6b\x0e\xf0\x6d\x07\xf8\xba\x03\x7c\xd9\x01\x1e\x3a\xc0\x7d\x07\xe0\x0e\xb0\xeb\x00\x8e\x03\x7c\x73\x05\x38\x58\x01\xba\x2b\xc0\x95\x15\xe0\xa3\x2b\x40\x63\x05\xf8\xd9\x32\xf0\xbd\x65\xe0\x95\x65\x40\x2e\x03\x6c\x19\x08\x97\x81\x73\xcb\xc0\x5f\x97\x80\xd7\x97\x80\xfb\x4b\xc0\xc7\x97\x80\x3f\x2e\x02\xbf\x5a\x04\xde\x58\x04\xbe\xb3\x08\xbc\xba\x08\x6c\x2f\x02\xff\xbc\x04\xbc\x76\x09\xd8\xbe\x04\xbc\xfb\x12\xf0\xaf\x8b\xc0\x0f\x2f\x02\xab\x17\x81\x37\x2f\x00\xfb\x17\x80\x4f\x5e\x00\xbe\x76\x1e\xb8\x73\x1e\x58\x3e\x0f\x7c\xe5\x1c\x10\x9f\x03\x2e\x9e\x03\xfe\x74\x16\xf8\xc6\x59\xe0\x13\x67\x81\xef\x9e\x01\xf6\xcf\x00\x9d\x33\xc0\x4f\x1a\xc0\x9d\x46\x91\x9b\xf7\x35\x80\x67\x1a\xc0\x3f\x16\x80\x3f\x2c\x00\x1f\x68\x00\xcf\x36\x80\x9f\x2e\x00\x3f\x5a\x00\x7e\xf0\x34\x0e\x42\xfc\xfc\x17\x0f\x5d\xf8\xde\xae\x0c\x87\x5d\xf8\x3a\x50\x3c\x31\x24\x64\x3d\xa6\x88\x56\xc1\x46\x6b\x60\x4c\xa2\xd7\x3d\x8f\xcc\x31\x1a\xd5\xb6\xb5\x6f\xcf\x37\x57\x7b\xe9\x50\x48\x1a\xb6\xaf\xcb\x28\x62\xb1\x69\x7f\x46\x86\x43\xf2\x90\x44\x54\xdd\x0b\xe5\x41\x4c\xc6\x63\xcc\x0f\xe4\x31\xe1\x5a\xa7\x8c\xf8\x94\x0c\x14\xeb\x6d\xb4\x46\x23\xf2\xe1\xfe\x20\x25\xe3\xb1\x97\xf7\x68\xaf\xaa\x7c\x33\x6b\x6a\xdf\x4a\xa3\x5d\xa6\xc8\x78\xdc\xea\xce\x77\x6e\x73\x23\x18\x19\x8f\x7d\x8f\x76\xc9\x01\x37\x83\xf5\x1a\xbb\xd5\x61\xd7\x02\xc3\x65\x4c\x1e\x12\x53\x8e\x9c\x72\x66\x2e\xae\x17\xb7\x6f\xde\xb8\xbb\x75\x23\xb7\x1e\x14\x6d\x99\xa9\x39\x23\xa3\x11\x61\x42\x33\xc2\x7b\x84\xdd\x27\xcf\x85\x4c\xb1\x9e\x36\x8a\xb4\xb7\x87\x09\x5b\x25\xad\xdc\xdb\x52\xe9\x0b\xfb\x2c\x36\xad\xfa\x14\x9d\x3a\xb7\x49\x2a\x04\x51\xec\x7e\xca\xb4\xa9\xcd\x6f\x26\x30\x95\xdd\xdb\xa9\x10\x5b\xc5\x80\xa3\x72\x5c\x15\xf9\x5f\xc9\x74\xc5\xe7\x2d\xb6\xcf\xd9\xc1\xc9\xb2\xce\xe2\xb0\xbe\x6f\x3a\xa8\x2d\xd6\xcb\xec\x1c\x15\xd5\x24\xd9\x45\x1a\xb7\x58\x22\xdb\xb7\x68\x74\x98\xb6\x40\x31\x6a\x58\x78\x84\x07\x59\x54\x47\xb8\x70\x2a\x27\x8c\x62\xcc\x9b\x19\x37\x37\xc5\x45\xdb\xe3\xbd\xe3\x3d\x12\xb3\x6a\xba\x67\x9c\x59\x25\x2d\xc5\x12\xa9\xb9\x91\x6a\x58\x9b\xe6\xc7\x4d\xdc\xf5\xdc\xfc\x71\xd3\xb4\xf9\x1f\x5a\x3c\xc8\x73\xc7\x0c\x5c\xcd\x73\x53\x53\xd8\xa7\x08\x31\x57\xad\x8f\x0b\xf1\x76\xaa\x07\x2c\x24\xd9\x03\xc0\xeb\x9f\xe3\xa2\x6b\x2a\xd4\x3b\x2f\x5e\x9b\x8b\xf0\x26\xd3\x9a\xf6\xd9\xf4\xba\x71\xca\x07\x4a\x0f\x8e\x77\x59\x08\xb2\xf5\xa4\x56\x9e\x13\x4f\xc4\x09\xd6\x87\xa3\xa2\xf2\x93\xe9\x2c\xf9\x26\x24\x81\xa0\x5a\x6f\xb4\xa2\x50\xb8\x21\x35\xd4\x35\x74\x57\xb0\x9d\x9d\x80\x09\xe1\xba\xb1\x8c\xdd\x38\x8d\x98\xe2\x01\x89\x8a\xc4\xb7\x4e\xaf\xa2\xd5\xf5\x93\x9a\xf5\x23\xe9\xfa\x9e\x09\xdf\x01\xb5\xf9\xda\xf1\x8e\x6a\xcc\x13\x79\xac\xa3\x93\x72\x58\xdf\x20\x89\xe2\xb1\x21\xf9\x37\x92\x75\xaf\xfa\x83\xcf\x56\x75\xb6\x0e\x63\x5f\x9d\x9e\x1f\xdf\xa8\x47\x9a\x47\x23\xa2\x68\xdc\x67\x64\x52\xee\xfa\x91\xb0\x6f\x8a\x2f\x43\xe5\x9d\x67\x06\x8c\x56\x9c\xf2\xcd\xe0\x2d\xc4\x59\xaa\xcf\xb4\x9c\x56\x45\x16\x10\xc9\x02\x7a\x3b\x4a\xb2\xa9\x7b\x3b\xe3\xf3\x89\x9a\xf3\xc2\xaf\xe4\xc6\xcf\x07\xd7\xab\x24\xd9\xed\x9e\x9e\x6d\xd1\x03\x1a\xca\x03\xd7\xbd\x1c\x26\xf9\x7d\x66\xf7\xf0\x1f\xd7\x5d\xbb\xec\x06\x52\xb4\xba\x28\xf4\x7b\xa9\x98\xd8\x9d\xff\xd8\xf5\x05\xaf\x9a\x16\x5c\x9b\x9d\x1d\x6e\x58\x54\xac\x0e\x2f\x28\x25\x55\x5e\x63\x82\x57\x94\x94\x35\x90\xf7\x1e\x56\x80\x9f\x8a\x59\x55\xc7\x7a\x47\x88\x1f\x51\x1e\x4f\x8d\xa1\x43\x99\x9a\x9d\x9d\x40\xc6\x26\x5b\x29\x72\x21\x2f\xcb\x15\x2b\x8b\xd0\xf7\x42\xbe\x7f\x98\x46\x9d\xd0\x9a\xf1\x6e\xfe\x21\x5d\x04\xf0\x68\x29\xcb\x64\x4b\x1d\x21\xdf\xaf\xb3\x5a\xd8\x71\x95\x3c\x28\x0c\x17\xf7\x47\x4b\x92\x9a\xc1\xae\x0e\x94\x9c\xc4\x2b\x85\x54\xae\xdb\x57\x6c\xe8\xae\x75\x3a\x95\x36\xc3\x1e\x98\xb2\xe3\x6a\xa7\xd3\xea\xa2\xde\xa5\x8a\x01\xd7\xed\xf1\x07\x2c\x74\x2b\x96\xf7\x74\x55\x6c\xc6\x58\xa6\xb3\xdc\x9b\x78\x93\x5a\xf3\x3d\x6d\x86\x82\x15\x49\x98\x94\x40\x4f\xc6\xc6\xd5\xfc\x65\xb6\x4e\xd6\xae\x24\x0f\x3e\x95\x37\x17\xa5\x66\x42\xf2\x31\x32\xc2\x94\xf8\xc1\x80\x1b\xe6\xe6\xbf\x22\x5f\x27\xb1\x54\x11\x15\x33\x43\x46\x53\xf2\x11\x7d\xe0\x1e\xf0\xd0\x0c\xd6\xc9\x5a\xa7\xf3\x91\x19\x59\x1e\xf5\x73\x79\xff\xd0\x31\x5f\xf0\xf8\x1e\x51\x4c\x6c\xb4\xf2\x36\x3d\x60\xcc\xb4\xca\x8f\xb8\xc9\x9e\x2a\x90\x21\x6b\xf7\x99\x89\x42\xd1\xe6\xd2\x5b\x6b\x5f\x69\x77\xbc\x88\x1a\xa6\x38\x15\xed\x90\xb1\x64\x27\x49\x55\x22\x98\xcb\xe3\x90\xf7\x65\x3b\xe2\x71\x3b\xd0\xba\x45\xbc\x13\xdb\xc8\xf2\xa2\xdb\x7d\x29\xfb\x82\xd1\x84\xeb\x7c\xc1\xe4\x81\x8c\x3f\xdd\xa3\x11\x17\xc3\x8d\x9b\xa5\xbd\xe7\x37\x03\x19\x1f\xaa\xce\x4b\x6f\xa6\xf2\x8a\x36\xf8\xc5\x3c\xf8\xc5\xd6\x31\xdb\xff\x67\xfb\xf5\x9f\x3b\xc5\xfe\x7f\xf2\x7e\xff\xf7\x4d\xe0\xfb\x4d\xe0\xd5\x26\xf0\xd5\x26\xf0\x72\x13\xd8\x6b\x02\x5f\x68\x02\x9b\x4d\xc0\x6f\x02\x97\x9b\xc0\xf3\x4d\xe0\x37\x0e\xf0\xcb\x89\x0e\x8b\xc5\x62\xb1\x58\x2c\x16\x8b\xc5\x62\xb1\x58\x9e\x00\x6f\xe1\x95\xe8\xf1\xdb\xf7\x6c\xaf\xbe\xa7\x5b\x5d\xdf\x2b\x34\x64\x2a\xcb\x33\x1f\xdf\x8b\x28\x9f\x1c\xe2\x78\xf9\x29\x42\xe5\xb8\xb1\x7a\xfa\x58\xff\x86\xc9\xf7\xaa\x67\x99\x65\x4b\x38\x7b\x40\xef\x7b\x49\xfd\x99\xfd\x51\x2f\xac\xb2\xdd\x7e\xed\xe9\xff\x91\x27\xff\xe3\x31\xfe\x1d\x00\x00\xff\xff\x90\x09\x83\xd2\x00\x40\x00\x00"),
		},
		"/html.html": &vfsgen۰CompressedFileInfo{
			name:             "html.html",
			modTime:          mustUnmarshalTextTime("2017-09-26T00:37:31.868932667Z"),
			uncompressedSize: 3862,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x38\x08\x1b\xd0\x6c\x90\x98\xb4\x2f\x83\xc7\x68\xe8\x8a\x02\x2d\xd0\x14\x45\x96\x3d\x1b\x8c\x78\xb6\xd8\xf2\x8f\x4a\x52\x71\xbc\xd4\xdf\x7d\x20\x29\x39\x92\xa3\xb8\xf9\xd7\x61\x7a\xa2\x8e\x77\x3f\xde\xef\xee\x78\x24\x69\xed\x95\x2c\x0f\x68\x8d\x8c\x97\x07\x00\xd4\x0b\x2f\xb1\xbc\xbe\x86\xe2\x3c\x8c\x60\xb3\xa1\x24\xc9\xc2\xac\x14\xfa\x0b\x58\x94\x27\x99\xf3\x6b\x89\xae\x46\xf4\x19\xd4\x16\x17\x27\x59\xed\x7d\xe3\x66\x84\x2c\x8c\xf6\xae\x58\x1a\xb3\x94\xc8\x1a\xe1\x8a\xca\x28\x22\x2a\xa3\xff\x58\x30\x25\xe4\xfa\xe4\x94\x79\xb4\x82\xc9\x5f\xdf\x57\x46\xbb\x0c\xc8\xbd\xa1\x2b\xc3\xb1\x58\xa2\x57\x5c\x16\xc2\x90\xe3\xe2\x55\x71\x44\x54\x07\x57\x70\xc4\x66\xde\xb4\xb6\x91\x98\x0b\xcd\xc5\xd2\x14\x4a\xe8\xa2\x72\xdb\x35\x22\x72\x18\x01\x78\x76\x21\x11\x3c\x07\xa1\x96\x70\x1d\x45\x00\x8a\x5d\xe5\x2b\xc1\x7d\x3d\x83\xe3\xa3\xa3\x9f\x7f\x8f\xe2\xcd\x58\xbf\xd7\x5d\xd5\xc2\x63\xee\x1a\x56\xe1\x0c\xb4\xb1\x8a\xc9\x5e\x7f\x6c\xf0\xcb\xd6\x24\x44\x26\x77\xe2\x1f\x9c\xc1\xf1\xab\xe6\xea\x06\x9e\x92\xce\x33\x4a\x52\x22\xe8\x85\xe1\xeb\xf2\x80\x72\x71\x09\x95\x64\xce\x9d\x64\x8a\xcb\x5c\xb2\xb5\x69\x3d\xdc\x0c\xf3\x7c\x21\xae\x90\xe7\xc1\x0a\x6d\x9c\xf8\xec\x86\x6a\x95\x91\xc6\xe6\xf9\xd2\xe2\x3a\x3f\x3e\x3a\xca\x62\x18\x3a\xed\x5b\xc0\xf3\xf9\x00\x67\x2c\xca\x73\x57\x59\x23\xe5\x14\xe8\x40\xe6\xf1\xca\x77\x13\xbf\x75\xab\x01\x4c\xb3\xd8\x22\x5b\xb3\xea\x14\x43\x86\x1a\xa6\x6f\xeb\xe6\xb1\x02\xb3\x9d\xb2\x0c\xba\xdd\x0a\x84\x8b\xcb\x48\x8d\x24\xd0\x38\x56\x4c\x4c\x60\xcd\xe7\x95\xd1\x1e\xb5\xef\xbd\x6b\xe5\x48\x49\xb8\x2e\x72\xd8\xb3\x45\x29\xf3\xfc\xf8\x65\xa0\xb8\x75\xf4\xfa\x1a\x2c\xd3\x4b\x84\xe2\xad\xb5\xc6\x3a\xd8\x6c\xba\x99\x58\xca\xbb\x88\xf3\xb9\xf0\xa8\x12\x81\x68\x10\x09\x48\x31\x80\x43\xcd\x7b\x10\x4a\x5a\x59\xa6\x2a\xa2\xa9\x8c\x06\x70\x9c\x79\x96\x27\x69\x97\xee\x1d\x89\xab\x19\x37\xab\x3c\x7f\xc9\x9b\x7b\x10\xa1\xbe\xdf\xfa\xdb\xff\xe9\xd5\xe6\xf3\x64\xaf\x8d\xce\x75\xab\xd0\x8a\x2a\x2b\xdf\x5e\xa2\xf6\xf0\x91\x29\xa4\xc4\xd7\x8f\x45\x39\x5f\x37\x4f\xb2\x3f\xc3\xc6\x3c\xd9\x89\x53\x74\x8e\x2d\x47\x10\x61\x3c\x08\x0e\xf5\x69\x57\xf6\xf8\x37\x35\xd0\xd9\x8e\xaa\x20\xe8\xdb\x72\xf0\x1b\x0d\x7e\x5a\xd6\x2d\xcc\x4e\xa0\xb1\x42\x7b\x88\x8d\x6d\x46\xc8\x52\xf8\xba\xbd\x88\x8d\x32\x83\x17\x1c\x2d\x2e\x9c\xb7\x50\x04\x62\x45\x20\x76\x38\x46\x0e\xd8\xfc\x01\xdc\x68\x93\x2a\x2f\x64\x2b\xc0\xc5\xea\x6b\x4a\x4a\x3c\x2f\x9f\x0e\x1b\xd2\xf7\xbc\x88\x5b\xde\xcf\x02\x0b\x2a\xa5\x27\x1b\x43\x00\xd0\x66\x57\x12\x73\x24\x16\x80\x5f\x87\x69\x08\x04\x0f\x21\xfb\xd4\x4a\x79\x86\x5f\x5b\x74\x3e\x46\x32\xdb\x4d\x4a\xfa\x82\x1a\x74\x7a\x40\x59\x77\x84\xf5\xb9\xdf\x6c\x48\xd3\x4a\x49\x02\xcd\x4f\x6c\x2d\x0d\xe3\xc5\x00\xb8\xf8\xd8\xaa\x0b\x0c\xed\x21\x35\x8b\x29\x95\x9b\x0e\xc8\x4a\x18\x2a\xbd\xae\xbc\x30\x7a\xca\xab\xd0\x5f\xa4\xc3\x3d\xd4\x5c\xbd\x9f\x93\xab\x91\x43\x65\x94\x12\xd3\xa4\xd2\xd4\x88\xd6\x5f\xef\x5e\xdf\xe2\xd1\xed\x14\xf8\x06\x8a\xd9\x2f\xdc\xac\x74\x47\xe4\xc1\x2e\xbf\x77\xae\x45\xb7\xcf\xe9\xa8\x31\xe9\xac\x88\xb6\x23\x67\xa3\xf2\x5d\xd1\x4f\x93\x7d\xdc\xe1\xc5\x1e\xc3\xc3\x67\xcc\xca\x1b\x8b\xcc\xe3\x3e\x8a\xa9\x5a\x35\x0e\x8d\xfb\x75\xcf\x70\xd1\xe1\x58\x6c\x8c\x13\xde\xd8\xf5\x1d\x30\x1d\xd4\xae\x29\x7c\x03\xdf\x73\x9e\x0a\xa3\xb7\x88\x64\xc7\xee\x56\xec\x92\x2c\x46\xa5\x8a\x7c\xf8\x34\x8f\x18\x89\x3b\xbc\x7b\x84\x6f\xd9\x44\x23\xf9\xae\x0b\x37\x47\xf0\x43\xd2\x34\xd8\x9b\x67\x78\x29\x70\xf5\xc6\x28\x85\x7a\x6f\x93\x98\x28\x90\x3b\x08\x6d\xd5\x3a\xd4\xe2\xdd\xf9\xe9\x87\xbf\xcf\x3e\xc4\x40\x57\x49\x36\xb9\x81\x00\x42\xa3\x01\xfb\x5f\x75\xa2\x95\xf0\xf5\xec\x3b\x54\x7b\x0e\x7f\x1a\xbe\x1e\xf7\x80\xc7\xed\xff\xff\x4b\xa4\x85\x06\xf1\x83\xba\xcd\x8f\x8c\xed\x9d\x01\x9b\x6c\xd3\x93\x30\x53\x5b\x26\x1c\xd9\xe3\xd3\x7a\xe7\xfc\xa6\x64\x78\x35\xda\x85\xa1\x64\x70\xd1\xa2\x24\x1e\xeb\xe9\x72\x1f\x6e\xf4\xe1\x91\x14\x6f\xfb\xd4\x55\x56\x34\x1e\x38\x2e\xd0\x82\xb3\xd5\x7d\xdf\x89\xe1\x51\xf8\xd9\x65\x25\x25\x09\x21\x20\x76\xef\x2d\x92\x9e\xc3\xff\x06\x00\x00\xff\xff\x31\xed\x50\x28\x16\x0f\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.html.html.swp"].(os.FileInfo),
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
