// Code generated by vfsgen; DO NOT EDIT.

package main

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

// Readme statically implements the virtual filesystem provided to vfsgen.
var Readme = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 2, 28, 7, 57, 28, 812468030, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2021, 2, 28, 7, 57, 28, 956468852, time.UTC),
			uncompressedSize: 8532,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\xdd\x72\xdb\x38\xb2\xbe\xf7\x53\xf4\x51\x9d\x4a\xe2\x1a\x93\xb4\x24\xdb\xf9\xb3\x35\xe3\xd8\x49\xc5\x75\x9c\xc4\x33\x56\x26\xc9\xb9\xd9\x82\xc8\x26\x09\x1b\x04\x18\x00\x94\x44\x5f\xed\x43\xec\x13\xee\x93\x6c\x35\x40\x52\xbf\xce\xda\x33\xb3\xb5\x17\x33\x16\xc1\x46\xa3\xbb\xd1\xfd\xf5\x0f\x73\xfc\x3f\xe7\x9f\xce\xc6\xdf\xae\xde\x42\x6e\x0b\x01\x57\x9f\xdf\x5c\x5e\x9c\x41\x2f\x88\xa2\x2f\xc3\xb3\x28\x3a\x1f\x9f\xc3\xd7\xf7\xe3\x0f\x97\xd0\x0f\xf7\x61\xac\x99\x34\xdc\x72\x25\x99\x88\xa2\xb7\x1f\x7b\xd0\xcb\xad\x2d\x5f\x45\xd1\x6c\x36\x0b\x67\xc3\x50\xe9\x2c\x1a\xff\x16\xcd\x89\x57\x9f\x36\x37\x3f\x03\xbb\xb4\x33\x4c\x6c\xd2\x1b\xed\x1c\xbb\x03\xe7\x85\x90\xe6\x64\x0b\x9b\xfe\xcb\x97\x2f\xfd\x6e\x47\x8b\x2c\x19\xed\x00\x1c\x17\x68\x19\x10\x75\x80\xdf\x2b\x3e\x3d\xe9\x9d\x29\x69\x51\xda\x60\x5c\x97\xd8\x83\xd8\x3f\x9d\xf4\x2c\xce\x6d\x44\xbb\x5f\x43\x9c\x33\x6d\xd0\x9e\x54\x36\x0d\x5e\xf4\x20\xfa\x31\x9f\x6b\x5b\x0b\xdc\xca\x2d\x36\x66\x65\xb7\x64\x05\x9e\xf4\x32\x94\xa8\x99\x55\x7a\x89\xbc\x64\x32\x51\x71\x4b\x6c\xb9\x15\x38\xea\x57\x61\x61\x8e\x23\xff\x40\xcb\x82\xcb\x5b\xd0\x28\x4e\x7a\x86\x8e\x34\x39\xa2\xed\x41\xae\x31\xf5\xe6\x30\xaf\xa2\x28\x4e\xe4\x8d\x09\x63\xa1\xaa\x24\x15\x4c\x63\x18\xab\x22\x62\x37\x6c\x1e\x09\x3e\x31\x51\xc6\x6d\x5e\x4d\x82\x82\xe9\xdb\x44\xcd\x64\x10\x1b\x13\x0d\xc3\xfd\xb0\xbf\xfe\x26\x2c\xb8\x0c\x9d\xfc\xb6\x2e\x71\x53\x1f\x13\x6b\x5e\x5a\x30\x3a\x7e\xf0\xd9\x39\xcf\x72\xc1\xb3\xdc\x86\x37\x26\x7a\x19\xf6\x0f\xc3\xfe\xfe\xd2\x22\x1d\x78\x63\x7a\xc0\xa5\xc5\x4c\x73\x5b\x9f\xf4\x4c\xce\x06\x87\x47\x41\xff\xae\xfa\x69\xf8\x46\x5e\x7e\xfb\xfd\xe5\x65\xc2\xbf\xbd\x38\xac\xbe\x7e\xb8\xe3\x7c\x38\x49\xf4\x2d\x8a\xba\x1c\x3e\xc7\xfd\xff\xaf\xc7\xa7\xbf\xe6\xfb\x27\x3d\x88\xb5\x32\x46\x69\x9e\x71\x79\xd2\x63\x52\xc9\xba\x50\x95\xe9\x8d\x8e\x23\x2f\xf3\x5f\x2a\xbe\x60\x32\xab\x58\x86\x26\x32\x39\x0a\xf1\x03\x25\xe4\xec\xd3\x87\xe1\xfc\x6d\x7c\x74\xf6\x2e\xd5\x57\x1f\xcf\x3f\xf6\xab\xf2\xeb\x4f\x87\xb5\xbc\xf9\xf2\x7f\xa7\x5f\x4d\xf6\xd3\xe4\xcb\xd1\xf0\xfa\xfa\xce\xe2\xc3\x94\xf8\x2b\xbc\x61\xab\x4a\x9e\x5b\x94\x60\xca\x2a\x61\x17\x7e\xb0\xa9\xd2\x5d\x5c\xc9\xef\xd7\xb2\x2f\x44\x76\x7a\xce\x2e\xae\xde\xd5\x77\xfa\xd7\x17\x9f\xaf\x2f\x6e\xbe\x0e\x7e\x2f\xab\xf9\xfb\xbb\xcf\xb3\x6f\xdc\x7c\x9a\xa9\x17\xf7\xab\xb4\xe2\x4f\xa3\x5c\xdc\x98\x90\x4b\x6e\xdf\xb7\xa2\x71\x99\x7d\x92\x97\x8a\x25\xcf\x76\x5f\x2f\x69\x1f\xf9\xf8\x3e\x9e\xa8\xa4\x86\x58\x30\x63\x4e\x7a\x9d\x57\xd3\xa2\xc3\x80\x7e\x1b\x44\x79\x9f\x1e\x07\xa3\x2f\x39\xb3\xc0\x0d\xd8\x9c\xd3\xea\x60\xb4\x73\x5c\x8e\xc6\x39\x37\xb4\xc8\xc0\x14\x4c\x08\x30\x68\x41\xa5\x70\x87\x5a\x05\xb1\x92\x29\xcf\x2a\xcd\x08\x89\xe0\xfc\xe3\x35\x54\x96\x0b\x6e\x39\x1a\x48\x95\x06\x66\x0c\x37\x24\x24\x70\x09\x09\x5a\x8c\x1d\x21\x93\x09\xe0\xbc\x14\x8a\x5b\xbf\x53\xa5\x70\x7d\xfd\xdb\xbb\x40\xa3\x60\x16\x13\x98\x56\x82\x50\x60\xd2\xf0\x0a\xe1\xc2\x42\xa9\xd5\x94\x27\x68\x00\x99\xa9\xc1\x2a\xa8\x0c\xba\x23\x35\x4e\xb8\x4c\xe8\x10\x7f\x78\xbd\x07\xcc\xc0\x0c\x85\xa0\xbf\x0c\x66\xcc\x91\x67\x68\x41\xa3\x51\x62\xca\x26\x02\xdd\xcf\x4a\xc7\xf4\x23\x56\x3a\x31\x30\xe3\x36\x07\x26\x6b\xc8\xf8\x14\x65\x0b\x3e\x26\x3c\x8e\xca\xc6\x0e\x08\x56\x29\x01\x89\x42\x03\x52\x59\xc0\xa2\x14\xaa\x76\x7b\xa4\x9a\xa2\x00\x8b\x71\x2e\xf9\xf7\x0a\x8d\x53\x91\x7b\xb2\xca\x2d\x91\x05\x88\xd2\xa0\x34\x18\xc2\xa9\x10\x90\x22\xb3\x95\x26\x62\x8d\x60\x35\x9f\x72\x26\x48\x52\x5e\x94\x02\x0b\x94\xd6\x71\x89\x99\x84\x09\x92\xd6\x5c\xd4\x90\xaa\x8a\x38\x4b\x50\x36\x47\x0d\x86\x17\x5c\x30\xed\x04\x5b\x15\xd5\xa0\x9e\xf2\x18\x49\x88\xb8\xd2\x1a\xa5\x15\x35\xe8\x4a\x02\xb3\x60\x73\x84\x63\xb6\x16\x0d\xce\x17\xa2\x5e\xeb\x13\x6c\x04\x89\x2a\x18\x97\xf0\xcc\xe9\x62\xcd\x53\x30\xd5\xc4\xaf\x99\xdd\xe6\xac\xfc\x60\x74\xed\xcd\x18\xab\x04\x8f\xa3\xfc\x60\x21\x80\x5f\x27\xcf\x99\x32\x2e\x9c\xd1\x95\xdc\x3c\xd7\x43\xab\x8b\x3d\x89\x38\xf7\x72\xf4\x46\x7e\x99\xe4\x08\xe1\xbd\x9a\xe1\x14\xf5\x9e\x13\x9c\x0e\x22\xae\x26\xe7\xd6\xd6\x74\xc5\x0e\x5b\x5a\xdd\x2f\x52\xa8\x55\x05\x33\x26\x2d\xd9\xd2\xa0\x48\x83\x5c\x19\x0b\xdc\xee\xf9\x37\x5c\x08\x90\x88\x09\xbd\x8e\x73\x26\x33\x04\xf2\xeb\x9c\xe9\x24\x20\xe6\x49\xab\x38\x25\x24\x13\xc2\xb5\x65\xda\x7a\xf7\x38\xa6\xf7\xa3\xc6\x63\xc2\x4c\x1d\x47\x6e\xc1\xdd\x93\x7f\x57\xd4\x89\x5c\x7a\xd3\x9a\x69\x30\xfa\x6c\x58\x86\x4d\x50\xe5\x07\xa3\xd3\xc0\x73\xe9\x4c\xb6\x26\x37\x79\x2b\x6b\x7c\x13\x2c\xc5\xa5\xf7\x5d\x34\xf4\x96\x49\xb8\xb8\xda\x73\x11\x40\x26\x49\x95\x10\x6a\x46\x11\xd0\xdd\xd0\xab\xd6\x1e\x8d\x58\xec\x16\x83\x27\xc2\xbe\xbe\xb8\x7a\x92\xd9\xd7\x81\xd6\x61\x73\xcf\xee\x75\x4b\xfc\x4e\x69\xc0\x39\x23\xff\xdb\x6b\xad\xb0\xc4\xa0\x1f\x0e\xc2\x61\x78\xb0\xbe\x7b\x45\x34\x4f\xde\x50\x36\x04\xad\x30\x1a\x5b\x30\x72\xf9\xa0\xd7\x08\xf7\xbf\xe0\x2e\x28\xb0\x70\x0a\x5b\xcf\xd9\xd9\xba\x0a\x39\x85\x77\x92\x68\x34\x06\x9a\x97\x3b\x0b\x85\x34\x3a\x95\xbe\xa9\xca\x85\x10\xd9\x2a\x61\x26\x47\x03\x5c\x1a\x8b\x2c\x21\xd4\x49\x94\x35\xe4\x42\x42\xc9\x8c\xfe\x92\x35\x2f\xae\xc8\xbd\xa6\x4c\xf0\xe4\xd1\x82\x07\x83\x60\x18\x6c\x08\xbe\xb6\xfa\x38\xc1\x4b\xc1\x62\x8a\xa6\x02\x5b\x24\x29\x35\xa6\x7c\x1e\x99\x2a\x4d\xf9\x1c\x26\x98\x2a\x8d\x4b\xd7\xd4\xde\x0a\xe1\x6f\x6a\x51\xb7\x5e\xab\xdb\x17\xcf\xbc\xd6\xda\xf9\xbd\x9a\x61\xb2\xfb\x08\x3d\x59\xe8\x8f\x0f\xd6\x54\x0b\x98\x4c\x42\x16\x78\xa1\x1a\xed\x1f\x41\xfb\x30\x9b\x7c\xa8\x84\xe5\xa5\x58\x80\x76\x03\x8e\x06\x4b\xa6\x5d\xe6\x98\xd4\x8d\xba\x74\x46\xf0\x47\xfd\xaf\x91\x94\x58\x1c\x06\x47\xc1\xf3\xe0\xc5\x3d\x57\xba\x8d\x62\xab\x2a\x8f\xdb\x75\x18\x1e\x85\xcf\xc3\x17\xeb\x06\xc8\x0f\x46\x2b\x09\x6f\x01\x1e\xb2\x81\x46\x49\x95\x68\x9b\x4e\x61\x52\x65\x66\x6f\x2d\x47\xba\x4c\x4e\xc9\x31\xb8\x95\x6a\x26\x17\x39\x0b\x2c\xd3\x19\xba\x64\x3d\xfe\x74\x36\xfe\xf4\xd9\xd5\xb6\xc4\x6b\x2d\x1b\x43\x52\x69\xa2\xba\xb8\x82\x89\x60\xf1\xad\x68\x52\xbc\xd2\x30\xcb\xb9\xc5\xe6\xd9\x25\x6d\x6e\xa0\x44\x9d\x2a\x5d\x60\x02\x95\x21\x32\xd6\x02\xcb\x2a\xaa\x51\x7a\x04\x81\x19\xb7\xc4\xb8\x39\xc2\x81\x1b\xd7\x86\xe8\x28\xad\x5a\x78\x16\xe7\x18\xdf\xee\x3a\xc4\xb5\xaa\x41\x3f\x3d\xe1\x49\x82\x12\x94\xc4\xe5\x9d\x06\x63\x25\x93\xc5\xd6\xca\xe0\xee\x22\x47\x12\xc6\xba\x56\x03\x17\x22\x11\xe2\x7b\xa4\xa7\x12\x08\x26\x98\xb3\x29\x57\x7a\x2b\xd2\xd6\xd2\xb2\xf9\xfd\x30\xdb\xf7\x38\xeb\xec\xde\x2c\x0d\x1e\x0c\xbd\x74\xd4\xb2\x48\xdb\x30\xd8\x73\xee\x1f\xbd\x0c\x07\x87\x07\xdd\xdf\x75\x68\x76\xd9\xae\x35\xa1\x33\xf5\x7d\x20\xed\x4d\x9a\xa3\x5c\x22\xf0\x4c\xc3\xe6\xef\x9f\x45\xf3\x1f\x4b\xbc\xf3\x18\xe2\xad\x31\xf6\x5f\x38\x79\xd5\x40\x9b\x80\x45\x85\x90\x50\x19\x8f\xc9\x97\xb8\x37\x70\x5b\xf7\xb9\x50\x34\x8d\x4b\x99\xc6\xaa\x95\x68\x8d\xca\x65\xac\x5d\x15\xc8\x5c\xc7\x2e\xf8\x88\xa7\xb4\x5d\x23\xcc\xe8\x7f\x52\xb5\x9e\x6d\x7c\x24\x70\xd3\xfa\x4c\x13\x03\x82\x19\x0b\x87\x4d\x18\x98\x3d\xe0\xf6\xa9\x69\xbd\x60\x11\x3d\xce\x37\x2e\xae\x5e\x1f\x47\x82\xfb\x73\x5c\x69\x39\xe3\x06\xef\xd9\xd2\xc4\x95\x92\x18\x36\x9b\xa2\x4a\x6c\x24\x5b\x0f\xfd\x68\x36\x13\x54\xb0\xec\x71\x4d\x16\x5b\xce\x54\xc1\x22\x55\x51\xa1\xab\x13\xd4\xae\xde\x27\x98\x62\x82\xdf\xe1\x46\x7c\x3c\xc3\x30\x0b\x9b\xdd\x2b\x29\xe7\xdf\xdd\xe9\x4a\x0e\x6a\x0e\xdd\x0d\xc1\xd5\xfa\xbc\x40\x55\x35\x89\xb2\x4d\x30\xae\xab\x41\x16\xe7\xcb\xc7\xaf\xd7\x9e\x6b\xc5\xa5\x13\xd6\xf7\x8b\x8b\xdb\x68\xd9\xff\x59\x68\x49\x95\xa6\x45\x6a\x3f\xf5\x94\x09\xf7\xe6\x01\x68\x33\x73\x7e\xe4\x39\xaf\xef\xef\x6c\x6f\x5c\xd5\x61\x73\xe7\x4b\xfc\xb6\xa5\xef\xef\x77\xe8\xf2\xac\xbf\xdf\x2a\xb4\x4b\xf8\xef\x09\x0e\x8b\xee\xfd\x21\x14\x5c\x56\x16\xcd\xee\x76\x23\x51\x0a\xe8\x75\x59\x03\x93\x1e\x81\xff\xb3\x59\xce\xe3\x9c\xce\xbf\xb8\xea\x3b\x17\xa8\x3a\x6c\x34\xbb\x94\xef\x35\xda\x4a\x4b\x4c\xa0\x68\x0b\x02\x32\x67\xe7\x6a\x5d\xb6\xfb\x0f\x58\xd7\x39\x2a\x2d\xca\xaa\xa0\x67\x77\xf0\xa3\x01\xfe\x5e\x40\x4f\x95\x1e\xee\x1b\x77\xc8\xc0\xb3\xee\x0f\x9e\x87\xfb\xe1\x7e\xd8\xdf\x8e\xed\x2d\xaa\x77\x95\xfa\x1a\xaa\xfb\xf8\xb6\x33\xe5\x4d\xb4\xb7\x80\xf9\x7b\xb6\xb7\xc7\x75\x0c\x94\x06\x49\x55\xc5\xb0\xbb\xeb\x45\xf3\x77\x7a\x7a\xba\xd1\xc9\x8c\x95\x43\x5f\xa8\xca\xfb\x72\xbd\x92\xa2\xee\xda\x98\xe9\x51\x0b\xa7\x8f\xbb\x2d\x5e\x06\xd3\xa3\x07\xf4\x34\x67\x4a\x28\x69\xa0\xa8\x8c\xf5\xae\xe3\x0a\xea\xc4\xa7\x79\x81\x76\x81\x3b\x71\xdb\xb4\xc1\xa9\x01\x26\x66\xac\x36\x7b\xa0\x99\x4c\x54\xd1\xa0\xd9\x32\x62\x35\xa5\x67\x65\xf0\x11\x6d\xc2\xe9\xe9\x29\x2c\x03\x94\x57\xa2\x1f\x0f\xe2\x78\xb8\x0e\x47\x3b\x0f\x25\x74\x29\x69\xd9\x90\xd0\x7f\x35\x78\xf5\x6a\xb8\xa5\x7c\x3c\xfb\x78\xfa\xe1\xad\xe9\xae\xea\x4d\xdd\x22\xd3\x1e\x54\xb2\x64\xda\xb8\x36\xbd\x61\xd3\x4c\x26\x62\x25\x0d\x4f\x50\x63\x42\xe9\xca\x73\x78\x98\xc6\x4e\xf2\xc6\xe7\xa9\xc1\x5f\xcb\xb8\xdb\xde\xb8\xa4\x28\x81\x09\xce\xfc\x08\x69\x89\x28\xdc\x09\xc3\x70\x4b\x8e\x55\x44\x18\x2f\x95\x71\x56\xd1\xcd\x30\x2f\xeb\x1e\xa9\xd3\xde\x30\xc1\x75\x97\x7e\xbc\x7d\x1f\x5b\xce\x78\x1e\xf7\x6b\x75\xff\xfb\x3f\xa0\xdb\x47\x65\xd1\x07\x8f\xb7\xbb\xf3\x5a\x55\x59\xe8\x5a\xb8\xd5\xdb\x59\x4c\x68\xbc\x7f\x57\xb2\x1d\xa9\x79\x6d\x4c\x35\x09\x4d\x1e\xde\x74\x4d\x62\x3b\xe3\x31\xb6\x2a\x79\xb2\x18\x3d\xf9\xb9\x5c\x82\x96\x71\xb1\xbb\x07\x1d\x22\x97\x5a\x4d\xd8\x44\xd4\x6e\xb2\x35\x23\xc1\xda\xa9\xc5\x63\xed\x28\x54\xcc\x04\xad\xad\xd9\x6f\x73\x7d\xc3\x6e\x1d\x49\xd8\xe9\xb3\x61\xbb\xfc\x60\xf4\xc9\x8d\xc8\xda\x19\x4a\x5d\xa2\x59\x9e\xb6\x10\xce\xf8\xdc\x36\x41\x3b\x43\x94\xf7\x55\x28\x9b\x45\x09\xb5\x50\x5d\xb0\x78\xfe\x54\x2a\xd1\x8b\x2e\x33\xb9\xf1\xa7\xac\xbb\x26\xaa\xa9\xd5\xc2\x87\x5a\x69\xfc\x75\xec\xed\x34\x11\x2c\xa7\xff\xd6\xcc\xb4\xb1\x0c\x09\xfa\xf9\x2f\xa7\x1a\x9f\x00\xfb\xc9\xf7\x4a\xd9\xd7\x2d\xa1\x7f\xda\x62\xa6\xf7\x38\x07\x94\xb1\x5a\xe9\x28\xdb\x4a\xce\xbd\xc0\x45\x77\x29\xad\x21\x6d\xba\xd1\x14\x97\x90\xa3\xc7\x45\x0a\x33\xd6\x58\x2b\xc7\xf9\x5a\x98\xad\x54\x78\xcb\x36\x7e\xac\xdf\x10\xeb\x61\x7f\x80\xc3\xc1\x00\x87\xc3\x01\x0e\xd7\x0b\xf9\x1f\x10\x3c\x64\xca\x90\x1f\xf8\xa8\x6b\x06\xda\xe3\xf1\xe5\xc2\x6d\xae\x55\xe1\x27\xab\xa8\x8d\x9b\x03\xcb\xa7\xb6\x1b\xd9\xe5\x4c\x26\x02\xdd\x60\x9c\x76\xb9\x4c\xe3\x86\xd7\xe7\x4d\xf1\x47\x8b\xdc\x40\xdf\xf9\x46\xaf\x19\xd0\xf6\xa0\x0d\x58\xb2\xe1\xbe\x7b\xe7\x0a\x70\x73\xdf\x48\xb3\x29\x2b\xc7\xe3\xcb\x65\x68\x33\x68\x83\xa6\x2c\x99\xa0\x76\x09\xd1\x5a\xb1\xf0\xe3\xda\x97\x7c\x5c\xae\x17\xcf\x8b\x64\x7e\xa9\x32\x98\x72\x9c\x2d\xfb\x41\xd3\xc0\xb8\x3b\x17\xa2\x99\x25\x34\x5d\x07\xc1\x41\x35\x11\x3c\x76\x05\xb3\x46\x3f\xc7\x5e\x49\xdf\x28\x93\x52\x71\x69\x1f\xd2\xde\xac\x4c\x86\xbb\x81\x34\x35\x31\xbd\xd1\xc6\x92\x1b\x50\xff\xf3\xef\xff\x70\xd3\x7a\xe3\x5b\x9d\xfe\xfe\x7e\x27\xdc\xa2\x9b\x59\xdd\xaa\x32\xb7\x8b\x49\xe0\x32\xe5\x92\x5b\x52\x8f\xb9\xc9\x48\xc9\xb2\xa6\xf9\x6f\x06\xe6\x4e\xf1\x66\xd2\xb4\xb7\x5c\x01\x13\x28\x42\x90\x76\xe5\xc2\x05\x05\x46\xe2\x26\x1c\xc4\x83\x60\xa0\x92\x92\x78\x36\x69\xa7\xd2\x02\xd6\xe5\x58\x6a\x71\x18\x58\xd4\x05\x97\x4c\x10\xd0\x0a\xba\x1e\xda\x5c\xfb\xba\xb7\x44\xcd\x8b\xe6\x53\x44\xab\xd4\x3d\xb6\x52\xd9\xcf\x99\xc6\xf2\x84\x1c\x41\x63\x86\xf3\x92\x1c\x61\xdd\x7c\x2d\xd5\xa2\xfc\x5f\xd0\x76\xc1\xd0\x98\xd7\x50\x7b\x45\x41\x33\x51\x53\xdc\x83\x49\x65\xc1\xe4\x6a\xe6\x4b\xb8\x82\xd9\xb8\xe9\x0e\x24\x9a\x95\x5e\x30\x1f\x8c\xce\x94\xb4\x2c\xb6\x06\x9e\xb0\xa2\x7c\x0d\xef\x4e\x7f\xed\xbe\x29\x35\x1e\x9d\xb3\x29\x3a\x90\x74\x77\xc6\xa9\x4c\x53\x1a\x4c\x95\x65\xed\x23\x97\xd4\x40\x24\x7b\x90\x22\x0a\x48\x35\xba\xf1\x45\xec\x39\x43\x81\x30\xe5\x6c\xf3\x93\x82\x0d\x0b\x74\x1f\x13\xb0\xe0\xa2\x37\xfa\xa5\xfd\xe9\xb4\x52\x12\xc6\x28\x30\xd3\xac\x70\x2d\xcb\xc6\xe6\x19\xa7\x9a\xd0\x7d\x90\xa0\x4d\x7f\x13\xa8\x25\xea\xde\xe8\x97\xa5\xa7\x8e\x53\x43\xdc\xc5\xd0\x85\xff\x7a\xe6\xbf\x14\x39\x79\xd7\x13\x81\xf3\x91\x9f\x17\x90\x52\xe9\xa5\x10\x7c\x53\x35\x59\x95\xa7\x70\xe1\xca\x61\x6e\xdb\x76\xde\x8d\x80\xb9\x10\x98\x31\x01\x2c\x99\xa2\x74\x53\x04\xa5\xe1\xfc\x5c\x5d\x13\xa1\xff\x84\x05\x79\x95\x21\xb0\x42\x55\xd2\x8d\x02\xad\x66\x69\xca\xe3\x9f\x97\x63\xda\xf8\x60\x65\xb3\xb4\x12\xc0\x13\x64\xa6\x49\xf1\x17\x90\x28\x02\x35\xf2\xf6\x25\xad\x9e\x4e\x29\xd1\xb1\x64\xf1\x05\x86\xd0\xca\xdd\x5f\xac\x64\x2c\xaa\xa4\xed\xe6\x6a\x55\x3d\x25\xce\x20\x95\x9a\xb4\x33\x40\x37\x35\xc8\xb9\xb5\x1c\x8d\xfb\x92\xe7\xac\x8f\x53\x67\xb9\x0e\x5b\xb5\xae\x9b\x23\xa3\x89\x4a\xea\xd1\xb1\xfb\x07\x07\xa3\x9d\x7f\x05\x00\x00\xff\xff\x90\x7c\x59\x2d\x54\x21\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
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
			gr:                        gr,
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
