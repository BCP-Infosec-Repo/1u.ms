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
			modTime: time.Date(2021, 2, 28, 20, 39, 10, 596664267, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2021, 2, 28, 20, 39, 10, 740665385, time.UTC),
			uncompressedSize: 8488,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x5a\xd1\x72\xdb\x38\xb2\x7d\xf7\x57\xf4\x55\xdd\x4a\xe2\x1a\x8b\xb4\x24\xdb\x49\x1c\x5b\x33\x8e\x9d\xdc\xb8\xae\x93\x78\xc6\xca\x24\xd9\x97\x2d\x88\x6c\x92\x88\x41\x40\x01\x40\xc9\xf4\xd3\x7e\xc4\x7e\xe1\x7e\xc9\x56\x03\x20\x45\xc9\x72\xd6\xde\x99\xad\x7d\x98\xb1\x08\x36\x80\xee\x46\xf7\xe9\xd3\x60\x8e\xfe\xe7\xec\xe3\xe9\xe4\xeb\xe5\x1b\x28\x6c\x29\xe0\xf2\xd3\xeb\x8b\xf3\x53\xe8\xf5\xe3\xf8\xf3\xe8\x34\x8e\xcf\x26\x67\xf0\xe5\xdd\xe4\xfd\x05\x0c\xa2\x5d\x98\x68\x26\x0d\xb7\x5c\x49\x26\xe2\xf8\xcd\x87\x1e\xf4\x0a\x6b\x67\x87\x71\xbc\x58\x2c\xa2\xc5\x28\x52\x3a\x8f\x27\xbf\xc5\x37\xb4\xd6\x80\x26\x87\x9f\x7d\xdb\x99\x19\xa5\x36\xed\x8d\xb7\x8e\xdc\x86\x37\xa5\x90\xe6\x78\xc3\x32\x83\x97\x2f\x5f\xfa\xd9\x4e\x16\x59\x3a\xde\x02\x38\x2a\xd1\x32\x20\xe9\x3e\x7e\xaf\xf8\xfc\xb8\x77\xaa\xa4\x45\x69\xfb\x93\x7a\x86\x3d\x48\xfc\xd3\x71\xcf\xe2\x8d\x8d\x69\xf6\x2b\x48\x0a\xa6\x0d\xda\xe3\xca\x66\xfd\x17\x3d\x88\x7f\xbc\xce\x95\xad\x05\x6e\x5c\x2d\x31\x66\x65\xb6\x64\x25\x1e\xf7\x72\x94\xa8\x99\x55\xba\x23\x3e\x63\x32\x55\x49\x23\x6c\xb9\x15\x38\x1e\x54\x51\x69\x8e\x62\xff\x40\xc3\x82\xcb\x6b\xd0\x28\x8e\x7b\x86\xb6\x34\x05\xa2\xed\x41\xa1\x31\xf3\xee\x30\x87\x71\x9c\xa4\xf2\x9b\x89\x12\xa1\xaa\x34\x13\x4c\x63\x94\xa8\x32\x66\xdf\xd8\x4d\x2c\xf8\xd4\xc4\x39\xb7\x45\x35\xed\x97\x4c\x5f\xa7\x6a\x21\xfb\x89\x31\xf1\x28\xda\x8d\x06\xeb\x6f\xa2\x92\xcb\xc8\xe9\x6f\xeb\x19\xde\xb5\xc7\x24\x9a\xcf\x2c\x18\x9d\x3c\x78\xef\x82\xe7\x85\xe0\x79\x61\xa3\x6f\x26\x7e\x19\x0d\xf6\xa3\xc1\x6e\x67\x90\x36\xfc\x66\x7a\xc0\xa5\xc5\x5c\x73\x5b\x1f\xf7\x4c\xc1\x86\xfb\x07\xfd\xc1\x6d\xf5\xd3\xe8\xb5\xbc\xf8\xfa\xfb\xcb\x8b\x94\x7f\x7d\xb1\x5f\x7d\x79\x7f\xcb\xf9\x68\x9a\xea\x6b\x14\xf5\x6c\xf4\x1c\x77\xff\x52\x4f\x4e\x7e\x2d\x76\x8f\x7b\x90\x68\x65\x8c\xd2\x3c\xe7\xf2\xb8\xc7\xa4\x92\x75\xa9\x2a\xd3\x1b\x1f\xc5\x5e\xe7\x3f\x55\x7d\xc1\x64\x5e\xb1\x1c\x4d\x6c\x0a\x14\xe2\x07\x46\xc8\xc5\xc7\xf7\xa3\x9b\x37\xc9\xc1\xe9\xdb\x4c\x5f\x7e\x38\xfb\x30\xa8\x66\x5f\x7e\xda\xaf\xe5\xb7\xcf\xff\x7f\xf2\xc5\xe4\x3f\x4d\x3f\x1f\x8c\xae\xae\x6e\x2d\x3e\xcc\x88\x3f\x23\x1a\x36\x9a\xe4\x57\x8b\x53\xcc\x58\x25\xec\x32\x0e\xee\x9a\x74\x9b\x54\xf2\xfb\x95\x1c\x08\x91\x9f\x9c\xb1\xf3\xcb\xb7\xf5\xad\xfe\xf5\xc5\xa7\xab\xf3\x6f\x5f\x86\xbf\xcf\xaa\x9b\x77\xb7\x9f\x16\x5f\xb9\xf9\xb8\x50\x2f\xee\x37\x69\x25\x9e\xc6\x85\xf8\x66\x22\x2e\xb9\x7d\xd7\xa8\xc6\x65\xfe\x51\x5e\x28\x96\x3e\xdb\x7e\xd5\xb1\x3e\xf6\xf9\x7d\x34\x55\x69\x0d\x89\x60\xc6\x1c\xf7\xda\xa8\xa6\x41\x87\x01\x83\x26\x89\x8a\x01\x3d\x0e\xc7\x9f\x0b\x66\x81\x1b\xb0\x05\xa7\xd1\xe1\x78\xeb\x68\x36\x9e\x14\xdc\xd0\x20\x03\x53\x32\x21\xc0\xa0\x05\x95\xc1\x2d\x6a\xd5\x4f\x94\xcc\x78\x5e\x69\x46\x48\x04\x67\x1f\xae\xa0\xb2\x5c\x70\xcb\xd1\x40\xa6\x34\x30\x63\xb8\x21\x25\x81\x4b\x48\xd1\x62\xe2\x04\x99\x4c\x01\x6f\x66\x42\x71\xeb\x67\xaa\x0c\xae\xae\x7e\x7b\xdb\xd7\x28\x98\xc5\x14\xe6\x95\x20\x14\x98\x86\xb5\x22\x38\xb7\x30\xd3\x6a\xce\x53\x34\x80\xcc\xd4\x60\x15\x54\x06\xdd\x96\x1a\xa7\x5c\xa6\xb4\x89\xdf\xbc\xde\x01\x66\x60\x81\x42\xd0\x5f\x06\x0b\xe6\xc4\x73\xb4\xa0\xd1\x28\x31\x67\x53\x81\xee\x67\xa5\x13\xfa\x91\x28\x9d\x1a\x58\x70\x5b\x00\x93\x35\xe4\x7c\x8e\xb2\x01\x1f\x13\x1d\xc5\xb3\xe0\x07\x04\xab\x94\x80\x54\xa1\x01\xa9\x2c\x60\x39\x13\xaa\x76\x73\xa4\x9a\xa3\x00\x8b\x49\x21\xf9\xf7\x0a\x8d\x33\x91\x7b\xb1\xca\x0d\x91\x07\x48\xd2\xa0\x34\x18\xc1\x89\x10\x90\x21\xb3\x95\x26\x61\x8d\x60\x35\x9f\x73\x26\x48\x53\x5e\xce\x04\x96\x28\xad\x5b\x25\x61\x12\xa6\x48\x56\x73\x51\x43\xa6\x2a\x5a\x59\x82\xb2\x05\x6a\x30\xbc\xe4\x82\x69\xa7\xd8\xaa\xaa\x06\xf5\x9c\x27\x48\x4a\x24\x95\xd6\x28\xad\xa8\x41\x57\x12\x98\x05\x5b\x20\x1c\xb1\x4e\x36\x1c\xc6\xb1\x0b\x85\xb8\xd7\x84\x04\x1b\x43\xaa\x4a\xc6\x25\x3c\x73\xa6\x58\xf3\x14\x4c\x35\xf5\x63\x66\x3b\x6c\x55\xec\x8d\xaf\xbc\x17\x13\x95\x22\x3c\x61\xe5\xec\x15\x18\x14\x59\xbf\x50\xc6\x1e\xc5\xc5\xde\x52\x1f\x2f\x47\x81\x34\x67\x5c\xb8\x33\x50\x72\x55\x0d\x4a\x4a\x8f\xb4\x2e\x15\x25\xe2\x8d\xd7\xab\x37\xf6\xc3\xa4\x57\x04\xef\xd4\x02\xe7\xa8\x77\x9c\x1d\x6e\x63\x6e\xc0\x14\xdc\xda\x9a\x4e\xdc\x41\x4d\xe3\x8a\xf3\x0c\x6a\x55\xc1\x82\x49\x4b\xae\x6d\x75\x03\x6e\x77\x20\x53\x42\xa8\x85\x0b\x77\x30\x16\x67\xe6\xd0\x4f\x53\xa2\x49\x1a\x2e\x13\xed\x8e\x82\xb9\xb2\x29\xf8\xf8\xff\xd0\x02\x6b\x7c\x43\x25\x8b\x82\xb7\x56\x95\x86\xa4\x50\xe4\x70\x72\xd7\xac\xb2\xe0\x22\xd3\x87\x96\x51\x60\x29\xb5\x9c\x18\x1d\x0c\x86\x3f\xc6\x8f\xfb\xd5\xa2\xa3\x58\x70\xbf\xc9\x25\xea\x4c\xe9\x12\x8e\xc8\xba\x71\xee\x83\x77\xa3\x67\x8e\x62\x27\xd2\x99\x7b\xa6\x16\x52\x28\x96\x3a\x45\x4a\x95\xf2\xac\x7e\xa0\x97\xe3\xa9\x50\xd3\xb8\x64\xc6\xa2\xf6\x23\x51\xcd\x88\x2f\x2c\x7f\xbb\x03\x58\xee\xf5\x5b\x25\x41\xf0\x6b\x74\x2e\x3c\x0c\xea\x3a\x69\xe8\xce\x59\x51\x31\x56\xc2\xa3\xcd\x27\xc3\x72\x0c\x30\x53\xec\x8d\x4f\xfa\xde\x5d\x6d\xd4\xac\x1d\x5d\xee\x1c\xef\x65\xbc\xdb\x7c\x36\x93\x13\x15\x30\x09\xe7\x97\x3b\x0e\x13\x28\x2a\xfc\xc9\x12\x26\xb4\x41\x7b\xd8\x84\x84\xd7\xb2\x64\xd7\xd8\x7f\x22\xec\xab\xf3\xcb\x27\xb9\x7d\xd5\xd7\x3a\xea\x3a\xb4\x11\x7e\xab\x34\xe0\x0d\xa3\x8c\xdc\x69\x4e\xbd\xb3\xc0\x20\x1a\x46\xa3\x68\x6f\x7d\xf6\x8a\x6a\xc1\x2b\x5e\x32\x08\x34\xca\x68\x6c\x22\xcd\x55\xc8\x5e\x50\xee\x7f\xc1\xc5\x68\xdf\xc2\x09\x6c\xdc\x67\x6b\xe3\x28\x14\x04\x78\x69\xaa\xd1\x18\x08\x2f\xb7\x96\x06\x69\x74\x26\x7d\x55\x95\x03\x15\xf2\x55\xca\x4c\x81\x06\xb8\x34\x16\x59\x4a\xa1\x9c\x2a\x6b\x28\x8b\x84\x92\x39\xfd\x25\x6f\x9e\x5f\x52\x86\xcd\x99\xe0\xe9\xa3\x15\xef\x0f\xfb\xa3\xfe\x1d\xc5\xd7\x46\x1f\xa7\xf8\x4c\xb0\x84\x00\xa5\xc4\x06\x5b\x67\x1a\x33\x7e\x13\x9b\x2a\xcb\xf8\x0d\x4c\x31\x53\x1a\x3b\xc7\xd4\x9c\x0a\x55\xa4\xcc\xa2\x0e\xaf\xb4\x6e\x5e\x3c\xf3\x56\x6b\x04\x46\x71\x83\xe9\xf6\x23\xec\x64\x91\xdf\xbe\xbf\x66\x5a\x9f\xc9\x34\x62\x7d\xaf\x54\xb0\xfe\x11\xb2\x0f\xf3\xc9\xfb\x4a\x58\x3e\x13\xcb\x32\x16\xca\x85\xc1\x19\xd3\xae\x96\x4e\xeb\x60\x2e\xed\xd1\xff\x77\xe3\x2f\x68\x4a\x4b\xec\xf7\x0f\xfa\xcf\xfb\x2f\xee\x39\xd2\x4d\x12\x1b\x4d\x79\xdc\xac\xfd\xe8\x20\x7a\x1e\xbd\x58\x77\x40\xb1\x37\x5e\xa1\x00\x4b\xf0\x90\xa1\x3a\x48\xe2\xe6\x0d\xc1\x80\x69\x95\x9b\x9d\x35\xd6\xe0\xb8\x0d\xd1\x85\xfe\xb5\x54\x0b\xb9\xac\xe2\x60\x99\xce\xd1\xd1\x97\xc9\xc7\xd3\xc9\xc7\x4f\x8e\xed\xd3\x5a\x6b\xfc\x04\xd2\x4a\x93\xd4\xf9\x25\x4c\x05\x4b\xae\x45\x20\x3d\x4a\xc3\xa2\xe0\x16\xc3\xb3\xa3\x31\xdc\xc0\xcc\xa3\x3b\xa6\x50\x19\x12\x6b\xcb\xc9\x2a\xaa\x11\x61\x00\x81\x39\xb7\xb4\x70\xd8\xc2\x81\x1b\xd7\x86\xe4\x88\x68\x58\x78\x96\x14\x98\x5c\x6f\x3b\xb8\xb7\x2a\xa0\x9f\x9e\xf2\x34\x45\x09\x4a\x62\x77\xa6\xc1\x44\xc9\x74\x39\xb5\x32\xb8\xbd\x64\x0d\x84\xb1\xae\xf9\xc2\xb5\x0a\xe7\xa8\x91\xab\x92\x53\x2c\xd8\x9c\x2b\xbd\x11\x69\x6b\x69\xd9\xcd\xfd\x30\x3b\xf0\x38\xeb\xfc\x1e\x86\x86\x0f\x86\x5e\xda\xaa\xab\xd2\x26\x0c\xf6\x2b\x0f\x0e\x5e\x46\xc3\xfd\xbd\xf6\xef\x3a\x34\x2f\x38\x51\xaf\xe0\x42\xe7\xea\xfb\x40\xda\xbb\xb4\x40\xd9\x11\xf0\x8b\x46\xe1\xef\x1f\x45\xf3\x1f\x6b\xbc\xf5\x18\xe1\x8d\x39\xf6\x5f\xd8\x79\xd5\x41\x77\x01\x8b\xb8\xa0\x50\x39\x4f\x28\x96\xb8\x77\x70\xc3\x84\x5d\x2a\x9a\x10\x52\x0d\x19\xab\x7e\x40\xc6\x78\x46\xd3\x35\xc2\x82\xfe\x27\x55\x13\xd9\xc6\x67\x02\x37\x4d\xcc\x84\x1c\x10\xcc\x58\xd8\x0f\x69\x60\x76\x80\xdb\xa7\xa6\x89\x82\x65\xf6\xb8\xd8\x38\xbf\x7c\xb5\xe4\x39\x8e\x6c\x2f\xb8\xc1\x7b\xa6\x84\xbc\x52\x12\x5b\x96\x53\x89\x3b\xc5\xd6\x43\x3f\x9a\xbb\x05\xaa\xdf\x8d\xb8\x50\xc5\xba\x95\xaa\xbf\x2c\x55\x44\xfd\x75\x8a\xda\x75\x40\x04\x53\x4c\xf0\x5b\xbc\x93\x1f\xcf\x30\xca\xa3\x30\x7b\xa5\xe4\xfc\xab\x33\x5d\xa9\x41\x61\xd3\xed\x08\x5c\xf7\xc3\x4b\x54\x55\x28\x94\x4d\x81\x71\x7d\x1e\xb2\xa4\xe8\x6e\xbf\x4e\xbf\x25\x7a\x5f\x25\x05\x93\x79\x50\xd6\x77\xd0\xcb\xd3\x68\x96\xff\xa3\xd0\x92\x29\x4d\x83\xd4\x90\xeb\x39\x13\xee\xcd\x03\xd0\x66\xe1\xe2\xc8\xaf\xbc\x3e\xbf\xf5\xbd\x71\xac\xc3\x16\x2e\x96\x88\xf7\x06\x54\xd8\x6d\xd1\xe5\xd9\x60\xb7\x31\x68\x9b\xf0\xdf\x0b\xec\x97\xed\xfb\x7d\x28\xb9\xac\x2c\x9a\xed\xcd\x4e\xa2\x12\xd0\x6b\xab\x06\xa6\x3d\x02\xff\x67\x8b\x82\x27\x05\xed\x7f\x7e\x39\x70\x21\x50\xb5\xd8\x68\xb6\xa9\xde\x6b\xb4\x95\x96\x98\x42\xd9\x10\x02\x72\x67\x1b\x6a\x6d\xb5\xfb\x0f\x78\xd7\x05\x2a\x0d\xca\xaa\xa4\x67\xb7\xf1\xa3\x01\xfe\x5e\x40\xcf\x94\x1e\xed\x1a\xb7\xc9\xd0\x2f\x3d\x18\x3e\x8f\x76\xa3\xdd\x68\xb0\x19\xdb\x1b\x54\x6f\x99\xfa\x1a\xaa\xfb\xfc\xb6\x0b\xe5\x5d\xb4\xb3\x84\xf9\x7b\xa6\x37\xdb\xb5\x0b\x28\x0d\x92\x58\xc5\xa8\x3d\xeb\x65\x3f\x7c\x72\x72\x72\xa7\x93\x99\x28\x87\xbe\x50\xcd\xee\xab\xf5\x4a\x8a\xba\x6d\x63\xe6\x07\x0d\x9c\x3e\xee\xb4\xf8\xac\x3f\x3f\x78\x40\x4f\x73\xaa\x84\x92\x06\xca\xca\x58\x1f\x3a\x8e\x50\xa7\xbe\xcc\x0b\xb4\x4b\xdc\x49\x9a\xde\x0d\x4e\x0c\x30\xb1\x60\xb5\xd9\x01\xcd\x64\xaa\xca\x80\x66\x5d\xc4\x0a\xd4\xb3\x32\xf8\x88\x36\xe1\xe4\xe4\x04\xba\x00\xe5\x8d\x18\x24\xc3\x24\x19\xad\xc3\xd1\xd6\x43\x05\x5d\x49\xea\x3a\x12\x06\x87\xc3\xc3\xc3\xd1\x06\xfa\x78\xfa\xe1\xe4\xfd\x1b\xd3\x1e\xd5\xeb\xba\x41\xa6\x1d\xa8\xe4\x8c\x69\xe3\x6e\x2a\xc2\x32\xe1\xae\x26\x51\xd2\xf0\x14\x35\xa6\x54\xae\xfc\x0a\x0f\xb3\xd8\x69\x1e\x62\x9e\xba\xef\xb5\x8a\xbb\xe9\x8d\x2b\x8a\x12\x98\xe0\xcc\x5f\xaa\x75\x84\xa2\xad\x28\x8a\x36\xd4\x58\x45\x82\x49\x87\xc6\x59\x45\x27\xc3\xbc\xae\x3b\x64\x4e\x73\xc2\x04\xd7\x6d\xf9\xf1\xfe\x7d\x2c\x9d\xf1\x6b\xdc\x6f\xd5\xfd\xef\x1f\x6d\x5b\xb1\x37\xfe\xe8\x2e\xbe\x9a\x7b\x80\x7a\x86\xa6\x7b\x63\x40\xb9\xe2\xf1\x79\x8a\x76\x81\x28\xef\xab\xb2\x77\x0b\x2b\xb5\x01\xed\x81\xfb\xf5\xa9\xdc\xd3\x8b\x16\x5d\xdd\xa5\xa6\xac\xdb\x46\x20\xf0\x8d\xe8\xa1\x1e\x9b\x7c\x99\x78\x9f\x4d\x05\x2b\xe8\xbf\x35\x57\xdd\x19\x86\x14\xfd\xad\x2e\x27\x9e\x4a\xa0\xf3\xe4\x7b\xa5\xec\xab\x46\xd0\x3f\x6d\x70\xd3\x3b\xbc\x01\x94\x89\x5a\xe9\x8a\x1a\x36\xe2\x5e\xe0\xb2\x43\x92\xd6\x90\x35\xed\xf5\x0a\x97\x50\xa0\xcf\x6d\x0a\x15\x16\xbc\x55\xe0\xcd\x5a\xa8\xac\xb0\x94\xae\x8f\x1f\x1b\x43\xb4\xf4\x68\x30\xc4\xd1\x70\x88\xa3\xd1\x10\x47\xeb\x64\xf4\x07\x02\x0f\xe9\x94\x8b\xbd\xf1\x07\x65\xdd\x95\x23\x75\x7f\x93\xc9\xc5\x32\x6c\xae\x54\x89\xe1\x5a\x8e\x08\xa3\x7c\x6a\xdb\x4b\xa7\x82\xc9\x54\xa0\xbb\xec\xa6\x39\x0e\x2b\xdd\x85\xf4\x59\xa0\x2f\x34\xc8\x0d\x0c\x5c\x64\xf4\xc2\xa5\x6b\x2f\xa4\x9d\xbf\x06\xde\x75\xef\x1c\x85\x34\xf7\xdd\x4b\x06\x62\x34\x99\x5c\x74\x93\xd3\xa0\xed\x87\xc2\x3a\x45\xed\x20\xdd\x5a\xb1\x8c\xe2\xda\x93\x16\x2e\xd7\xe9\xdf\xb2\x1c\x5d\xa8\x1c\xe6\x1c\x17\xdd\x28\x08\x14\xdc\x9d\xb8\x10\xa1\x1b\x0e\xbc\x99\xda\xd3\x6a\x2a\x78\xe2\x28\x9f\x46\x7f\x37\xbd\x52\x80\x50\xa6\x33\xc5\xa5\x7d\x08\x41\xdf\x7c\xcb\x4c\x34\xbc\x37\xbe\x33\xe4\x6e\x9d\xff\xf1\xb7\xbf\xbb\x1b\x78\xe3\xc9\xfa\x60\x77\xb7\x55\x6e\xc9\xc7\x57\xa7\xaa\xdc\xcd\x62\x12\xb8\xcc\xb8\xe4\x96\xcc\x63\xae\xb7\x9f\xb1\x3c\xb4\xaf\xe1\x12\xdc\x19\x1e\xee\x4a\x76\xba\x1c\xce\x32\x2e\xa0\x9f\xb5\x05\xef\x9c\xd2\x22\x75\x3d\x3a\xad\x41\x20\x50\x49\x49\x6b\x06\xe0\xac\xb4\x80\x75\x3d\x3a\x24\x9d\x81\x45\x5d\x72\xc9\x04\x2c\x0a\x2e\xe8\x78\x68\x72\xed\x99\xdb\x0c\x35\x2f\xc3\xe7\x85\xc6\xa8\x7b\x7c\xa5\xf2\x9f\x73\x8d\xb3\x63\x0a\x04\x8d\x39\xde\xcc\x28\x10\x7a\xe3\xcd\x52\x4b\x02\xbb\x94\x6d\x53\x21\xb8\xd7\x50\x83\x40\x29\x33\x55\x73\xdc\x81\x69\x65\xc1\x14\x6a\xe1\x49\x48\xc9\x6c\x12\xf8\xad\x44\xb3\xd2\xcd\x14\xc3\xf1\xa9\x92\x96\x25\xd6\x84\xab\xfe\xb7\x27\xbf\xb6\xdf\x89\x42\x44\x17\x6c\x8e\x0e\x22\xdd\x99\x71\x22\x1a\x4a\x83\xa9\xf2\xbc\x79\xe4\x92\x28\x70\xba\x03\x19\xa2\x80\x4c\xa3\x6b\xc0\x13\xbf\x32\x94\x08\x73\xce\xee\xde\x58\xdb\xa8\x44\x77\x57\x8d\x25\x17\xbd\xf1\x2f\xcd\x4f\x67\x95\x92\x30\x41\x81\xb9\x66\xa5\x23\xdd\x77\x26\x2f\x38\xb1\x1a\x77\xdf\x4d\x93\xfe\x2a\x50\x4b\xd4\xbd\xf1\x2f\x9d\xa7\x76\xa5\x20\xdc\xe6\xd0\xb9\xff\x22\xe6\xbf\xfe\x38\x7d\xd7\xcb\x80\x8b\x91\x9f\x97\x80\x52\xe9\x4e\x0a\xbe\xae\x2c\x2c\xdc\x97\xb5\x0c\xce\x1d\xa1\xe3\xb6\x69\x48\xdd\x25\x26\x17\x02\x73\x26\x80\xa5\x73\x94\xae\x0f\x56\x1a\xce\xce\xd4\x15\x09\xfa\xcf\x52\x50\x54\x39\x02\x2b\x55\x25\xdd\x65\x96\xd5\x2c\xcb\x78\xf2\x73\x37\xa7\x8d\x4f\x56\xb6\xc8\x2a\x01\x3c\x45\x66\xc0\xb7\x0d\xe7\x01\xd4\x28\xda\x3b\x56\x3d\x9d\x53\x99\x63\xe9\xf2\x33\x0a\xa1\x95\x3b\xbf\x44\xc9\x44\x54\x69\xd3\x8f\xd4\xaa\x7a\x4a\x2b\x83\x54\x6a\xda\xdc\x62\xb9\xbe\xb7\xe0\xd6\x72\x34\xee\xeb\x9c\xf3\x3e\xce\x9d\xe7\x5a\x64\xd5\xba\x0e\x5b\xc6\x53\x95\xd6\xe3\x23\xf7\x8f\x08\xc6\x5b\xff\x0c\x00\x00\xff\xff\xad\x80\x7b\xae\x28\x21\x00\x00"),
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
