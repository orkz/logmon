// Code generated by go-bindata.
// sources:
// assets/config-template.json
// assets/favicon.ico
// assets/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsConfigTemplateJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xe6\x52\x50\x88\x56\x2a\xae\x2c\xce\xc9\x4f\x57\xd2\x51\x50\xd2\x2f\x4b\x2c\xd2\x07\xb2\xf5\xa1\x42\xb1\x5c\xb1\x80\x00\x00\x00\xff\xff\xb6\xbd\x4b\xbe\x23\x00\x00\x00")

func assetsConfigTemplateJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsConfigTemplateJson,
		"assets/config-template.json",
	)
}

func assetsConfigTemplateJson() (*asset, error) {
	bytes, err := assetsConfigTemplateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/config-template.json", size: 35, mode: os.FileMode(436), modTime: time.Unix(1462073477, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9d\x0b\x50\x54\xd7\xfd\xc7\x7f\x2b\x0a\xf8\xe0\x2f\xea\xdf\x47\x30\x0a\x5a\x33\x9a\x58\x22\xa0\x28\xc5\x17\xa0\x82\x46\x47\x51\xd6\x28\x89\x0a\x86\x87\x24\xd1\xa2\xd1\xa8\x51\xa9\xab\xa8\x91\x28\xf8\x1a\xd3\xc1\x32\x68\x62\x54\x06\x13\x12\x9b\x56\x46\x2b\x86\x40\x02\x51\x83\x89\x0f\xc4\xf8\x88\x45\x0a\x55\x0c\x08\x59\x97\x87\x50\xb6\xe7\xe4\x92\x89\xa0\xce\xd0\x5d\xd8\x73\xe1\x7c\x3f\xe3\x67\xd6\x7b\x87\xdd\xfb\x3b\xe7\x7b\xf7\x72\x76\x39\x7b\x96\x48\x43\xd6\x64\x6f\xcf\x6f\x9d\x28\xa2\x3d\x51\x38\x29\xf0\x6d\xea\x40\x74\x9e\xed\x73\x71\x51\xb6\x93\x9e\x63\xdb\x1d\x89\x9c\x9c\xea\xb7\xd9\xfd\x56\x0e\x23\xd2\xe9\x94\xed\x21\xb6\x1a\xea\xef\x45\xe4\xe5\x55\xbf\xed\x43\xe4\x3b\x59\x43\x43\xd8\xe3\xb1\x1f\xe5\x8f\xa8\x3c\x2e\x87\x3d\x6e\x8f\xce\x8a\xcd\x48\x3c\x33\xc9\x44\xdf\x67\x1a\xcd\xb0\xb6\x19\xee\x9f\xc8\xfc\xd8\x44\xf7\x35\x43\xff\xf9\x32\xb5\x26\x3a\xd1\xcc\xf6\xd7\x34\x43\xff\x95\x9b\x71\xff\x92\x66\xe8\x3f\x73\x29\x25\xa5\x1d\xa6\x78\x97\xc4\x9e\xbf\x3c\xbf\x49\x64\xfa\xf9\x33\xc9\x68\x7c\xbc\x43\x9e\xb6\xef\x49\xfb\x63\xfc\xa7\xfb\xd9\x75\x72\xe8\xc4\xfe\x6b\x37\x65\xf2\xc4\x59\xfc\x6a\xc3\xb5\x65\xd7\x1a\xaa\x58\x1b\xcc\x03\x6e\xf7\xda\x54\xbf\x89\x64\x7c\x77\xc7\xb7\x5f\xb0\xad\x8e\x4b\x27\x07\x2e\x27\xea\xdc\x83\xab\xa1\x7d\x09\x7d\xd8\x85\xc9\x67\xca\x44\x6f\xed\xea\x6b\x25\x37\xc7\x6f\xcd\x7b\xfb\xd9\xec\xf2\x7f\x57\xea\x33\x4b\x1f\x46\xcf\xba\x18\xec\x1f\xbf\xb0\xdb\xe5\x59\x03\x8a\xec\xad\x8b\x77\xce\xd5\x4e\xf5\xf5\x8e\xf7\x6f\xdf\xcd\xce\x7d\x64\x7b\xe7\xdc\x59\x01\x76\x61\xdd\x9c\x0f\xfb\x77\xfb\x2a\x30\x74\x66\xf7\x0b\x8e\xd7\xbf\x2d\xcd\xf8\xfd\xf2\x21\xf7\x17\xbb\xdc\xf7\xbc\x18\x94\xb9\xa3\xe0\xfc\xe1\x8d\xb5\x69\xf9\x0b\xbe\xcb\x5e\x79\xeb\x6e\x87\xfd\x83\x6d\xbe\x28\xd4\x6b\xa6\xec\x1e\xc4\xff\x6d\x0b\x5a\x95\x15\x74\xee\x44\xc2\xa1\xdc\x3d\x2b\x7a\x19\xe2\x17\x6e\x3c\x53\xdd\x55\xaf\x59\x5c\xe8\x7a\xa5\xf6\x77\xb6\xce\xbe\x51\xa7\xff\xf5\xce\x8f\x0e\xb5\xae\x1e\x65\x61\x3f\x85\x06\x18\x5e\xb4\xcb\x1d\x7d\xe3\x5c\x8c\x37\x1d\x5c\x77\x36\x36\x77\xed\x88\x15\xfb\x3f\x8b\xf9\xa1\x7f\xd9\xc0\xd4\x39\x91\x33\x37\xec\x7d\x50\x95\x14\xe2\x6e\x73\x67\x76\xe4\xcc\x98\xf1\x27\x3e\x3c\x96\x53\xb0\xc3\x3d\x23\x2a\x36\x36\x77\xa8\x36\x62\x42\xf6\xa0\xde\xb3\x3f\xf5\x8f\x77\x1c\x3d\xbd\xe6\xfa\xf8\xab\xc7\x8a\xca\xfd\xee\xbd\x62\x68\x97\xeb\xa6\x5b\x6a\x6d\x7f\x9e\x32\x9e\xf1\xf2\xb0\xb2\x5d\x44\xf9\x5d\x9c\x12\x34\x9b\x86\x9b\xbf\x73\xe7\x27\xb7\x3c\x02\xd3\xbc\x37\xd7\x45\x6c\x9f\x13\xe9\xff\x41\x74\xec\xb1\x62\xbb\x17\xfe\x76\x24\xa4\xe7\xad\x63\x1f\xad\x79\xf7\xa5\x84\xe0\xb0\x1d\x05\xdb\x97\x2e\x37\x74\xdf\xbb\x7a\x77\x5a\xbf\xef\x43\x52\x7b\x67\xd9\x5c\xfb\x72\xe4\x1d\x4d\xd4\x4a\xd7\xb3\xaf\xac\xef\x6e\x1d\x5f\xf8\x7c\xa2\x6e\x73\xa4\xf5\xfc\xaf\x33\x17\x07\xf7\x3c\x98\x7b\x9d\xf7\x4e\xf8\x46\xd7\x05\xb7\x97\x1a\xfe\xbf\xbe\xef\x1a\xff\x7b\x10\x99\x13\x64\xbf\xcc\xb5\xeb\x5f\xcf\x74\x1e\xd8\xeb\xc4\xb8\xe1\xdb\xf5\xe3\x46\x57\x45\x95\x2d\x7c\x51\xe7\xb7\xfa\x7d\xfd\xd1\xbc\xd7\x23\xe2\xf4\x43\x22\x33\xfe\x19\xe7\xb4\x3a\x6e\x91\xc7\x91\x35\x7b\xb6\xdc\x99\x57\x1a\xba\x25\x2a\xda\x48\x5d\xcb\xde\x30\xf4\xdf\x7b\x35\x63\x6d\xbf\x73\x1e\x5f\x6d\xce\x9d\xaa\xbd\x3d\xf8\xe5\x51\xa5\x95\x55\x21\xf3\xb3\xe6\xe7\x0c\x1d\x78\xf0\xe6\x95\x65\xdb\x0b\x59\x47\x16\x79\x65\x3f\x1f\x69\x6b\x57\xf1\xa6\xbe\xea\xa0\x61\x4c\x5d\xcb\x74\x1e\xdf\xe9\x5e\x97\x9c\x78\xc0\xe7\xf3\xd8\xa0\xe4\x51\x49\xdf\xef\xca\x3a\xcd\xba\xef\x41\x60\x81\x81\xdf\xf8\x86\x74\x5d\x7b\x74\x4e\xd2\x37\xd3\x2b\x56\xa5\x1f\x5a\xf7\xe5\x7b\xc5\xce\x0e\x27\x1d\xe7\x66\x6a\xad\xfe\x14\x3e\xd6\xe6\x68\x5e\xc4\x37\xa7\x17\x8e\x0d\x58\x3e\x6c\xcb\xb4\x9f\x3b\x5d\x9a\x97\x73\xb1\x78\xb6\xc3\x7b\x13\xc6\xe4\x4c\x88\x9e\x13\x7d\xb2\xc3\xfe\xd1\x29\xe5\x17\x92\x26\xa5\x7c\x37\xa2\xc7\xba\xa0\xb1\x13\x9e\xd2\x8f\x8e\xce\xa7\x1c\x83\xdd\xfb\xa6\x39\x06\xef\xba\x37\x3c\xd3\x23\xe4\xd9\xc4\x1a\x97\xde\xe9\x7b\x4f\xa5\x17\x15\xbf\xed\xb0\xfb\xd2\x4f\xb7\x97\xc5\x65\xcc\xa3\xd4\xb2\xe3\x79\x99\x9b\xaa\x63\x5f\xad\x7e\xc3\x3e\xfd\xef\xeb\xce\x77\xe0\x27\xe4\x9e\x1b\xe9\xf9\xbe\xd9\x6f\xf6\x5f\xff\x51\xe1\xa9\x2e\xfd\x5e\x5d\x9d\xba\x26\xe6\xec\xc7\x4b\x16\x7f\xf2\xe3\x01\x9f\x55\xa3\xfe\x98\xcc\x6e\xee\x79\x86\x07\xce\x4a\xf3\x09\x8b\x1b\xb9\x63\xfc\xb5\xad\x33\xb2\x9a\xab\xc7\x1e\xdf\x99\x72\xe1\x53\xd6\x87\x47\x1c\x02\x72\xfb\x6a\x23\x92\xcb\x97\x18\xac\x7b\x5e\xbb\x75\x75\x3e\x3b\x2d\x57\xae\x4a\xc8\xf6\xf8\xb3\xf6\xbe\x7f\x8c\xd5\xbc\x3c\x8f\x33\x2b\xa7\xcf\x19\x36\x73\x43\xe0\xf1\x65\xc9\x36\xd3\x96\x6e\xca\xa9\x73\xea\x32\x3c\x73\xec\x3a\x1a\x4b\x15\xfa\x01\x79\x9f\x7f\xc8\x4e\x94\x70\xc7\x55\x65\xae\x09\xd5\x11\x97\x3d\x47\x56\x45\x5c\x1e\x98\xb7\x68\x5f\xc7\xac\xaa\x93\x6b\xb6\x36\xea\xb9\x3a\xbf\xf5\x56\x07\xfa\x6c\xbf\xe9\x97\xba\x6d\x01\xbb\xc4\x3c\xb7\x62\xd2\xdc\x15\xa1\x0b\x57\x84\x79\x86\xbc\x15\xc6\x6e\xc8\xcd\xc5\xd5\xdd\xd9\x65\x84\xb3\x9b\x9b\xd6\xd5\xdd\xd3\xd5\xcd\xd3\xcd\xcd\xd9\xc5\xdd\xd3\xc5\xe5\xe1\x6b\x93\x93\x1b\xdc\x61\x49\x64\xe8\xeb\xe1\xef\x3c\xfd\x0e\xdb\xc6\x97\xf5\xe2\xd7\xbd\x29\x93\xa6\x4f\x4c\xf1\x09\xde\xc8\xc7\x59\x6c\xa8\x46\xc1\xf4\xc8\x38\x6b\x50\x8b\x8c\xb3\x00\x68\xcd\x6c\x25\xf3\xc7\x9b\x22\xe4\x35\x6f\x6c\xa5\xb5\xff\x6a\x85\x0a\x6a\x30\x47\x73\xc7\xc9\xa2\x45\xfd\xa8\x5f\xf6\xfa\x1f\xaa\xa0\x0e\x53\x35\x30\xa3\x99\x55\x64\xfa\x7b\x05\xa2\xe4\x35\x47\x11\x00\x00\x88\x03\xe3\x7f\x71\x62\xfc\x8f\xfa\x51\x7f\xeb\xb5\x2d\xd4\x8f\xf1\x3f\xc6\xff\x00\x00\x39\xd9\xcc\xac\x24\xf1\xd7\xc3\xff\x55\x5e\xb3\x8e\x30\xfe\x17\x69\x5b\x18\xff\x88\xae\x01\xf5\xb7\x5e\xdb\x42\xfd\xd5\x2a\xa8\xc3\x54\xf9\xf8\x7f\x03\xb5\xce\xdf\x5f\x18\xff\xab\x00\xe3\x2f\x3c\xfd\xd6\xd2\xc7\x93\xed\xf8\x7c\x9e\x96\x13\xd3\x8b\x1e\x99\xa7\x65\x8f\x79\x5a\x4f\xe0\x05\x32\x7d\x3e\xb1\x39\x06\x30\x07\x93\x79\xf3\xc9\xcd\xf5\x9e\xc0\x63\xff\xfa\x7b\x12\xc7\x17\x7b\xfc\x32\x81\xc7\x2f\x66\x0e\x25\x31\xcf\x3f\x2e\xbf\x4c\x02\xb1\xbc\x45\xa6\x7f\x9e\xca\x1c\x8f\x30\x97\x30\xeb\x48\xdc\xf9\x2f\xfa\x35\x82\x1a\xae\x3f\x32\x1f\x9f\xbf\xc7\xf6\xb5\xc0\xe3\xa7\x13\x90\x9d\x4c\x12\xf3\x3a\x9d\x9f\xfb\x69\x24\xfe\xf9\x27\xf3\xf5\x47\x0d\xc7\x17\x59\x03\xff\xfd\xcb\xc7\x00\x22\xc6\x1f\x49\xf5\xc7\x16\x8a\xd1\xd8\xb4\x77\x24\x8c\x8f\xd0\xd4\x9f\x6b\xca\xcf\xf2\x01\xb8\x8e\x6f\x68\x1e\x79\x9f\x82\xfd\x0f\xef\x53\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x20\x25\xff\xc7\xf4\x23\x71\x73\x06\x60\xcb\xcb\xe7\x46\xf1\x75\x1e\xbb\x50\x43\xc2\x98\x7a\x12\xfb\x37\x0b\x68\x39\xf9\xdc\xa4\x79\xa4\x30\x43\x05\xf5\x40\xcb\xcb\xe7\xe7\xf8\x30\xcf\xa8\xa0\x16\x28\xc6\x53\xd4\xba\x3f\x4f\x0d\xcd\xb3\x44\x05\x35\x40\x71\x3e\x50\x41\x0d\x10\xf9\x43\xe4\x0f\x91\x3f\x44\xfe\x10\xf9\x43\xe4\x0f\x91\x3f\x44\xfe\x10\xf9\x43\xe4\x0f\x91\x3f\x44\xfe\x10\xf9\x43\xe4\x0f\x91\x3f\x44\xfe\xd0\xbc\xfc\x5b\xfb\x9a\xde\xd0\x74\xef\x32\x4f\xab\xa0\x0e\x28\xc6\xcf\x98\xde\xcc\xff\xa8\xa0\x16\x68\x59\xf9\x1a\x5d\xee\xa4\x30\x97\x59\xaa\x82\x9a\xa0\x65\xe4\xeb\x13\xcf\xa0\x86\x74\x22\x65\x59\xeb\x69\xb0\x4d\x3b\x8e\x69\x4b\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xda\x22\xed\x98\x2f\x33\xff\x42\xe2\xbe\xb3\x0b\xb6\xbc\xfc\xbb\x59\xe3\x99\x33\xe9\x37\x6c\x98\xa9\x24\x7e\x4e\x0a\xb4\xac\x29\xcc\xf6\xcc\x0d\x2a\xa8\x05\x8a\x71\x39\x33\x5f\x05\x75\x40\x31\x5e\x22\xb1\xdf\xd3\x0d\xc5\xfa\xb3\x0a\x6a\x80\xe2\xc4\xe7\xbf\xe4\x16\xf9\xcb\x2d\xf2\x97\x5b\xe4\x2f\xb7\xc8\x5f\x6e\x91\xbf\xdc\x22\x7f\xb9\x45\xfe\x72\x8b\xfc\xe5\x16\xf9\xcb\x2d\xf2\x97\x5b\xe4\x2f\xb7\xc8\x5f\x6e\x91\xbf\xdc\x22\x7f\xb9\x45\xfe\x72\x8b\xfc\xe5\x56\x4f\xca\x5a\x70\xa2\xeb\x80\x62\xbc\xc9\x4c\x50\x41\x1d\x50\x8c\x71\xcc\x67\x98\x05\x2a\xa8\x05\x5a\x56\xfe\xdc\xef\x4e\x0a\x7d\x99\x1f\x30\xef\x90\x32\x26\x80\x6d\xd7\x22\x52\xae\xf9\xbd\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x6d\x89\x01\xa4\xac\x0d\x5a\xc2\xac\x85\x6d\xd6\x9a\xfa\x8c\x0f\x31\xfb\x91\x82\x13\xb3\x98\xc4\xcf\x47\x82\x96\x95\xcf\x03\x72\x60\x1e\x56\x41\x2d\x50\x8c\x7c\xbd\xef\xfb\x2a\xa8\x03\x8a\x11\xf3\x7e\xe5\x56\xaf\x82\x1a\xa0\x38\xf1\xf9\x2f\xb9\x45\xfe\x72\x8b\xfc\xe5\x16\xf9\xcb\x2d\xf2\x97\x5b\xe4\x2f\xb7\xc8\x5f\x6e\x91\xbf\xdc\x22\x7f\xb9\x45\xfe\x72\x8b\xfc\xe5\x16\xf9\xcb\x2d\xf2\x97\x5b\xe4\x2f\xb7\xc8\x5f\x6e\x91\xbf\xdc\xf2\xfc\x6b\x54\x50\x07\x14\x63\x19\xf3\x8a\x0a\xea\x80\x62\x3c\xcb\x8c\x54\x41\x1d\x50\x8c\xa1\xcc\x76\x84\x39\xe0\x32\x9a\xc8\xd4\xd0\x6f\xbc\xc4\xdc\xc5\x3c\x00\xdb\xb4\xbb\x99\x7e\x04\x00\x00\x00\x00\x00\xa4\xc4\xa8\x72\xd2\x1b\x96\x6b\xa5\x6b\xb8\xad\xb1\x54\x3f\x81\x27\xa3\x6b\xb8\xa9\x49\x6f\xb8\x6d\x25\xfa\xfc\xb1\x34\x8d\xfa\xc3\xaa\x05\xbb\x1e\x98\x80\xae\xe1\xa6\x46\xf4\xf9\x22\x9a\xc6\xfd\x61\xa9\x1c\x40\xd3\x48\x6f\xb8\xd9\xec\xd7\xd3\x21\xec\x41\xbd\x48\x79\x5e\x68\x7e\x59\x0a\x46\xd9\xd1\xa3\xb3\x22\x00\x00\x00\xea\xc3\xdc\xc9\x4c\x92\x48\xbe\x26\x5c\x2c\xb3\x27\x33\x8b\xc4\xff\xcd\x4a\x94\xff\x60\xd6\xa9\xa0\x0e\x51\x56\xa8\xa0\x06\x91\x56\xa9\xa0\x06\xb4\x1f\xed\x47\xfb\xd1\x7e\xb4\x1f\xed\x17\xd1\x7e\x99\xc7\x3f\x06\x52\xde\x8a\x11\x5d\x87\x28\x8f\x33\x7b\x93\x7c\xaf\x7f\xb8\xfc\xf5\x4f\x0f\x02\x00\x00\x00\xe4\x65\x0c\x53\x2b\x91\x01\xcc\x3f\xd4\xb7\x9d\x8f\x7d\x44\x8f\xc3\x44\xb9\x85\xe4\x7e\x0d\x24\xfb\x1a\xf9\x32\x67\x8f\xf6\xa3\xfd\x68\xbf\xf8\x1a\xd0\x7e\xb4\x1f\xed\x17\x63\x25\xc9\xfd\x5d\x31\x3f\x30\x27\x33\x0b\x49\xfc\x77\xb8\x59\xda\xdb\x4c\x6f\x02\x00\x00\x00\x80\x8c\x4c\x21\x65\x0c\x28\x7a\x3c\x62\xe9\xef\xcb\xcd\x67\x4e\x20\x65\xec\x27\x7a\x1c\x2a\xca\x1b\x2a\xa8\x41\xa4\x95\x2a\xa8\x41\xa4\xb2\xbf\xfe\x45\xfb\xc5\xd7\x80\xf6\xa3\xfd\x68\x3f\xda\x2f\xaa\xfd\x32\xaf\x95\x5c\xc2\xdc\xaa\x82\x3a\x44\xb9\x9e\x14\x46\x31\xfd\x25\x73\x24\x01\xe9\x69\xea\x7a\x02\x6a\xad\xab\x31\x3a\xe5\xee\x16\x5b\xef\xc2\xd4\x3a\x2d\xdd\xaf\x3a\xe5\xe6\xb1\x75\x51\xfe\x1b\x00\x00\xff\xff\x46\xb4\x2d\xb4\x6e\x8a\x01\x00")

func assetsFaviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_assetsFaviconIco,
		"assets/favicon.ico",
	)
}

func assetsFaviconIco() (*asset, error) {
	bytes, err := assetsFaviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/favicon.ico", size: 100974, mode: os.FileMode(436), modTime: time.Unix(1462600281, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x5a\x7b\x73\xdb\x36\x12\xff\xbf\x9f\x02\x65\x66\x2a\xb9\xb6\x28\x39\xee\x23\x91\x65\xdf\x24\x6e\x66\x92\x9b\x34\xe9\x9c\xed\xe9\xdc\xe4\x3c\x19\x88\x84\x24\x9c\x29\x82\x05\x40\xa9\xbe\xab\xbf\xfb\xed\x92\x20\xc5\x07\x40\xd2\xcd\x15\x9a\x89\x45\x02\xbb\x5c\xec\xe3\xb7\x3f\x50\x59\x7c\xfd\xd3\xc7\xab\x9b\x7f\xfe\xf2\x86\x6c\xf4\x36\xba\xfc\x6a\x51\xfc\x61\x34\xbc\xfc\x8a\xc0\x58\x68\xae\x23\x76\x19\x89\xf5\x56\xc4\x8b\x69\x7e\x95\xcf\x28\xfd\x50\x7c\xc7\xb1\x14\xe1\x03\xf9\x6f\x79\x99\xdd\xa2\xc1\xfd\x5a\x8a\x34\x0e\xe7\xe4\xd9\x6c\x36\x3b\x27\xd3\x6f\x9f\x9d\xbe\xc4\xcf\x39\xf9\x76\x5a\x5b\xbb\x12\xb1\x9e\xac\xe8\x96\x47\x0f\x73\xe2\xfd\x24\x05\x0f\xc9\x35\x8d\x95\x77\x42\x14\xfc\x99\x28\x26\xf9\xea\xbc\x14\x79\x2c\xbf\x3d\x53\x3c\x64\x4b\x2a\x1b\xcf\x4e\x84\xe2\x9a\x8b\x78\x4e\xe8\x52\x89\x28\xd5\xec\xbc\x36\xaf\x45\x32\x27\xb3\xfa\xbd\x88\xad\x74\xeb\xe6\x52\x68\x2d\xb6\xad\xdb\x7b\x1e\xea\xcd\x9c\x3c\xff\x7e\x96\xfc\x7e\xee\xde\xb5\xd9\x6d\xdd\x34\x1a\x86\x3c\x5e\xcf\xc9\x29\xc8\x92\x19\x7e\x3a\x37\x76\x49\x68\x63\x73\x81\x88\x84\x04\xed\xf4\x0c\x3f\x75\xed\x41\x2a\x15\x4e\x26\x82\xc7\x9a\xc9\xfa\x64\xc8\x55\x12\x51\xf0\xf0\x32\x12\xc1\xbd\xc3\xac\x1f\xc0\xaa\xd3\xef\x9b\xdb\xda\x52\xb9\xe6\xe0\x4e\xb4\xf7\x14\xed\xee\xd8\xf5\xf3\xd7\xf8\x39\x77\x04\x44\xb2\x88\x6a\xbe\x63\x7d\x9b\xf6\x69\x80\xcb\xba\x93\xea\x3b\xcc\xab\x2e\x45\x74\xbe\x11\x3b\xd6\x4c\x8f\xba\xb9\x2f\xe0\xf3\x72\xa0\x39\xfd\xea\x86\x18\x45\x54\x42\xe3\x27\xa6\xac\x89\xfa\x7e\xc3\x9b\x33\xb5\xc7\x9f\xbd\x78\xf9\xea\xec\x45\x7d\x81\xe4\xeb\x0d\xa4\xf6\x8f\xcd\xa0\x66\x55\xa7\xf8\x7f\x18\x24\xe3\xf3\xe6\x64\x56\x22\x3f\xb4\xd2\x5b\xc8\x90\xc9\x89\xa4\x21\x4f\x95\x65\xbe\x4c\x23\xd0\x47\xbe\xab\xce\x3a\xfd\x30\x67\xdb\x44\x37\xc1\xa3\x4c\xd5\x58\xc4\xf6\x54\x09\xc0\x78\x16\xeb\xff\x5b\xe1\x5b\x6a\xd9\xf8\x6d\x18\x24\xd4\x21\xec\x4a\xa4\x92\x43\xa2\x7c\x60\x7b\xc0\x30\x73\x75\x42\xbc\xf7\x69\xc0\x43\x9a\x61\x1b\xb9\x79\x48\xd8\x5e\x42\x38\xa5\x77\x98\xa9\xdd\x04\xd4\x15\xe0\xa2\x80\xb9\xe3\xe6\xf2\x71\xe1\x9e\x4b\x70\xe5\xee\xaf\x42\xc7\x27\x39\x08\x0b\x67\x15\x89\xfd\x04\xfc\x43\x53\x2d\x1c\xb3\xbf\xdb\x66\xcb\xb4\x6a\x01\x53\x56\x0f\x93\xcc\x49\x98\x2c\x7b\x49\x93\xfa\x82\x1d\x57\x7c\xc9\x23\xae\xe1\xb1\x1b\x1e\x86\x2c\xee\xf7\x97\x1d\x7b\xaa\x9a\xb2\xef\x51\x77\x66\xe6\xae\xb7\x05\xa0\x0b\x88\x0b\x74\xb7\x15\x72\xe9\x06\x07\x0c\xd7\x9c\x91\x48\x07\xc6\x6a\xaa\x53\xd5\xb0\x68\xc3\xf2\x50\x9e\x76\x36\x34\xc9\x42\x17\xac\x0f\xcf\x24\x77\xef\x6c\xe5\xd8\xc1\xea\xf9\x7c\xb2\x67\xcb\x7b\x0e\x79\x1f\x48\x11\x45\xed\x96\x6f\xd4\x9e\x9e\x35\xb5\x96\x5b\x3b\xb3\x97\x8a\x45\xf5\x44\x4b\xd8\x75\x67\xd3\x38\xc3\xcf\x60\x75\x9b\x74\xbb\xec\x52\xd7\x0c\x76\xb7\xba\x40\xc8\xb8\xa7\xa7\x59\xcd\x5b\x4c\x0d\x6b\xcb\xaf\x22\x1e\xdf\x63\x3b\xbe\xf0\x38\xa4\xac\x47\x36\x92\xad\x2e\xbc\x15\xdd\xe1\xa5\x0f\xff\x78\x40\x06\xa7\x39\x1b\x5c\x20\xc1\x33\xd4\x0f\xf3\x99\x87\x17\x9e\x81\x71\xef\x72\x31\x85\x5b\x8d\x49\x53\x04\xf6\xc9\x3c\x05\xeb\x73\xb0\x3d\x9e\xe8\x03\xa3\xdc\x41\x84\x6f\xdf\x91\x8b\xc6\x36\x23\xbe\xe5\xfa\xbd\x58\x43\xf3\x39\x83\x3e\x7b\xd2\x48\x37\x4d\xa3\xdb\x58\x82\xc9\x90\x47\xf5\x39\x1e\x73\x48\x82\x55\x1a\x07\x98\xb0\xe3\xa3\x86\xde\xe2\x91\x89\x90\x1a\x1e\xba\xe7\x71\x28\xf6\x7e\x28\x82\x74\x0b\xdb\xf0\xa1\x4e\x29\xca\xf9\x38\xdf\x12\x2c\xc6\xdf\xba\xe5\xc8\x9c\xbc\x80\xec\x6e\x89\xeb\x0d\x57\xbe\x02\x24\x60\xf8\xe8\x98\xed\xc9\xaf\x6c\x79\x9d\x5d\x8f\xbd\xbd\x9a\x4f\xa7\xa8\x26\xda\x08\xa5\xe7\x1e\x39\xce\x6d\x3c\x26\xde\x74\xaf\xbc\xa3\x6e\x7d\x3e\x74\x90\x84\xc5\xa0\x36\xbb\x29\xe2\x8f\x70\xe5\x2f\xc1\xca\x31\xde\x38\x3a\xef\x11\xde\x32\xa5\xe8\x9a\x1d\xe4\x7f\xce\x6f\x3c\x41\x45\x10\x09\x55\x51\x70\x85\x97\x35\x71\xbb\x7c\x28\xb6\x18\xfc\x47\x87\x7a\x98\xf6\x0b\x1e\x71\x41\x4a\x7f\xaf\x99\x7e\x13\x31\xfc\xfa\xfa\xe1\x5d\x38\x1e\x99\x25\x23\x97\x95\xa8\xa6\xc0\xeb\x0e\x35\x66\xc9\xc8\x66\x6d\x29\x24\xe2\x7b\xf6\x00\xd1\x47\x6f\x97\x79\xc6\x6c\x89\x86\x03\x93\x2d\x10\x21\x7a\x86\xf9\x20\x78\x85\xdf\xff\xf8\x03\x2e\x00\xc5\x83\x4d\xdb\x5e\x1c\x7c\x45\xc6\xb9\xd0\x05\x79\xfe\xa3\x4b\x75\xa1\x9e\x16\x5e\xaf\x38\xcb\xff\x2d\x65\xf2\xe1\x9a\x45\x2c\xd0\x42\x8e\xbd\x03\x19\x33\x4d\xcf\x3b\x3a\x71\x2a\xc5\xb1\x2f\x94\xe6\xa9\xae\x3e\x51\xf4\xd6\x2b\xad\x25\x5f\x02\xfe\x8f\x47\x21\xd5\x74\x12\xd3\x2d\x1b\x1d\xdd\x59\xdc\x55\xea\xf1\x79\x0c\x28\xf6\xf6\xe6\xe7\xf7\xa0\x71\x34\xb2\x6f\xf8\xb1\x75\xf7\xb1\x92\x3b\xb5\xc9\xc7\xba\xdd\x39\xc6\xcc\x2d\x2e\x0a\x91\x9c\xb8\x33\x26\x93\x1b\x59\xbc\xb0\x96\x8c\xc5\x3d\x10\x82\xe3\xe0\x72\x44\x5b\xff\x00\xcc\xb8\xcf\x4c\x89\x65\xb3\x8f\xed\xe7\x41\xb7\xfd\xd2\xa7\x81\x0a\xdb\xb3\xba\xfc\x96\x43\xc4\xbc\x27\x83\x31\x11\xb3\x27\x4b\x06\xb5\x11\x43\x32\xbd\xc3\xf3\xe6\x8e\x46\x2e\x3b\x83\x88\x51\x59\x2c\x72\xc9\xda\x93\x20\x84\x6c\xd5\x8c\xd8\x65\x6c\xdb\x73\x20\x52\x16\x59\x3f\x8b\xc0\xb8\xf1\xa8\x96\x13\x0c\xce\x35\xfd\x60\xed\x19\xff\x56\x02\x6b\xfe\xef\xd7\x1f\x3f\xf8\x09\x95\x8a\x8d\x99\x8f\x55\x60\x43\x0b\x74\x1c\xae\xf7\x57\x3c\x62\xaa\x3f\xa8\xa6\x6a\x87\xd4\x8a\x64\x3a\x95\x71\x2e\x8b\xdd\xee\x5a\x43\x23\xdc\xaa\xea\xe3\x7a\x52\x01\x07\x8b\x54\x93\xf8\x56\xf7\xda\x2a\x7f\xd4\xfe\x69\x76\x77\xe7\x46\x8d\xb0\x8a\xac\x01\xd8\xa4\x99\xa9\x38\x00\x0b\xbe\xb3\x82\x2a\x8e\x4c\xf3\xe9\x1d\x48\x9b\x6f\x9d\xb8\x04\x99\x01\x9c\x3a\x60\xe3\xe9\x37\xd3\x35\x9c\xa7\xbe\xa1\xdb\xe4\xdc\x6b\xef\xd8\x2e\xb3\xc8\x65\x22\x3d\x5c\xe4\x32\x17\x59\x3f\x41\xc4\xcb\x45\x7e\x4b\xc5\x13\x84\x46\xb9\xd0\xb3\xd9\xd9\xcb\x73\x6b\xbb\xc7\x11\xd6\x32\xc4\xf8\xcb\x9e\x26\x7b\x9f\x26\x50\xe1\xe1\xd5\x86\x47\xe1\x38\x74\x69\x9c\x4e\xc9\x75\xc6\x38\x81\x55\x99\x23\x9d\x43\x5d\x4e\x4c\x6f\x44\x82\xbc\xc9\x5c\xbd\xcd\x48\xb7\x5b\xf5\x15\xa0\x41\x4c\xd2\x84\x44\x40\xe5\xdc\x25\xf0\x19\x51\x23\xbe\x4d\xc6\x26\xcb\x5c\xc6\x62\x55\x7d\x0d\x1d\x45\xbd\xca\x7a\x58\x57\x63\xcc\xf4\x56\x88\xe2\xf1\xb1\xdd\x4d\xf9\xde\xd2\x01\x6b\x4c\x8d\xbe\xa5\x71\x08\x00\x5c\xef\xaf\x23\x7c\xc3\x31\x3a\xaa\x45\xa7\xd0\x3a\xb8\xe3\x39\xf0\xec\x73\x9a\x00\xca\xb0\x1b\x7c\x2b\xdb\x83\x67\xd5\xa5\x8d\x9e\x62\xd7\x5d\x71\x10\x58\x6c\x23\xad\x10\xc4\xb7\xb0\x71\x42\x21\x43\x0c\x16\xb4\xd6\xac\x84\x24\x63\x44\x0d\xe0\x37\x40\xc0\x6b\xc0\xe1\x8a\x51\xd9\x59\xcc\x3a\x7f\x43\xd5\xc7\x7d\xfc\x8b\x04\x1e\x2b\xf5\xc3\x18\x54\x1d\x3d\x25\xbe\xe4\xb8\x01\x58\xa0\xe0\xee\xcb\x23\x50\x9a\x59\xf3\x14\xb8\xca\x65\x5b\x89\x80\xd9\x5b\x74\x84\xf1\xfc\xb5\x7a\x6f\x7f\xc6\xd1\x01\xca\x2e\xbd\x64\x3c\x82\x33\x42\xdb\x1b\x64\x74\xf4\x54\x4a\x50\x54\x61\x25\x73\x90\xdf\xb9\x8e\x50\xc6\xd1\xcd\x46\x81\x22\x77\xf5\xf2\x78\x15\x45\x65\x0b\xb0\x39\xb8\x48\x81\x88\xc5\x6b\xbd\x21\x97\xb9\xc2\xf2\x08\xe8\x72\x75\x99\x76\x3c\xcb\x5d\xf8\xb3\xb8\x68\x29\x9b\x34\x95\xc1\xba\xe3\xe3\xae\xcc\x2a\x36\xc2\xef\xb0\xcb\x83\xcb\x3f\x00\x13\x07\x9c\xde\x8a\x1d\xcb\xc1\xf4\xb0\xc2\xc1\x64\x6c\xa9\xd5\xe1\x78\x73\x5a\xea\x23\x63\x55\x76\x03\x8c\xaf\x89\x05\xe5\x92\x92\x3b\xf5\xa0\x45\xb9\xae\x87\x7e\x3e\x9d\x04\xe6\x14\x65\x30\xc2\xb5\xf4\x42\x3c\x15\x2b\xaf\xc6\xc3\xc8\xb1\xe3\x64\x9d\xcd\xd1\x30\x94\xc0\xf4\x1c\xe1\xfa\xa2\x83\x74\x5b\xc1\x9f\x3a\x4c\xb7\xd5\xf4\x1e\xa8\x5b\xee\xad\xcc\x9e\x90\xd3\xd9\x6c\xd6\x1d\xff\x0a\x79\xac\x64\x80\x93\xb0\x56\xab\x7c\xe0\x91\xbd\x97\xcf\x36\x0f\xe7\x2d\x81\x76\xaf\x41\xf3\x7c\x28\xfd\x37\x34\xd8\x8c\x6b\x56\x9f\x10\x6e\xed\x74\x38\x32\x98\xc0\x1e\x87\x0b\x81\x61\x1c\x78\xd7\x27\xf2\x2f\xff\xee\x78\x7a\x42\x46\x13\x27\x45\x45\x26\x93\x11\x5a\x52\x1c\xa3\x37\x19\x0b\x70\x3f\x4a\xb3\x6d\x07\x1d\xa6\x36\x24\xcc\x42\x02\x72\x35\x1f\x18\x73\x11\xcf\x17\xc8\x31\x2e\x17\xd3\xec\x8f\xe3\x6c\x90\xcb\x67\x67\x42\xf5\x19\x1b\x03\x77\x34\xbf\x6c\xa1\x72\x1c\xe8\x4f\x8a\xc7\x76\xe5\x69\x35\xcc\x55\xa6\x89\x9a\xfb\xfd\x08\xcd\xcb\x00\xad\xd3\x87\xb0\xe4\xd7\x6c\x45\xff\xb9\xc2\xa6\xa2\x14\x37\xfe\xd8\x77\xfa\xe3\xb0\x3a\xad\xb0\xa1\x1e\xbd\x86\x86\x62\x98\x28\x74\xee\xbe\xf5\x35\x02\x09\x42\xe8\x29\x87\xa3\x5a\x75\x51\xf5\x70\xa9\xd1\xe5\xe6\x5a\x3f\x36\xa1\xc4\xf3\x55\x29\xe8\x0e\xcf\x6b\x98\x87\x33\x3c\x0f\xee\x09\xdb\xc1\x83\xdd\xc9\x83\x00\x85\xcb\x9a\x2f\xc1\x9c\x7d\x75\x08\x93\x2c\xc6\x9f\x61\x94\xa5\x75\x5f\xc8\x2c\x8b\xd1\x66\x93\x41\x44\x95\xfa\x00\x15\xd2\x71\x3a\xef\xd6\x30\x2c\x65\x8a\xd1\xa6\x11\x87\x99\x2e\x37\xff\x23\xe3\x2a\xc4\xfc\xba\x55\x40\x16\x46\xad\xfb\x3d\x62\x1c\x6c\x84\x54\xd5\x72\x6b\xd1\xb8\xca\x8b\x44\xcf\x51\x78\x38\x3e\xdd\x15\x10\xed\x07\x10\xef\xb1\x51\x7d\x72\xc8\x15\xda\x17\x02\x3a\xd8\xdd\x8f\xae\x3a\x30\xee\xc8\x5c\x8e\xa0\x83\xc4\xa0\x03\x74\x70\x30\x5f\x53\xb9\x86\xfe\x5b\x7b\x78\xee\xc9\x0e\x13\x4a\xb1\xfe\x63\xa1\xb5\xab\x15\xe3\x00\x16\x4f\x7b\xfe\x70\xf8\xaa\xaf\xae\xe4\xa3\x96\x29\xeb\xb0\xac\xf7\x28\x5a\x8c\xc7\xc1\x1c\xc5\x16\xb6\x6a\xb8\x56\x5c\x2a\xed\x0a\xd8\xc0\x14\xc5\x56\x9f\xe1\x54\x8b\x0b\x97\x57\x8f\x15\x33\xcc\xef\x3a\x22\x8e\x44\xe6\xc4\xdb\x77\xd9\x2b\xb6\xdc\xe8\xdb\x77\x46\x07\xf4\x60\xf3\x2b\xd6\x62\x9a\xff\x60\xb6\x98\x66\xff\xa9\xea\x7f\x01\x00\x00\xff\xff\x82\x16\xd7\x8f\x6b\x25\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 9579, mode: os.FileMode(436), modTime: time.Unix(1462711556, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/config-template.json": assetsConfigTemplateJson,
	"assets/favicon.ico":          assetsFaviconIco,
	"assets/index.html":           assetsIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"config-template.json": &bintree{assetsConfigTemplateJson, map[string]*bintree{}},
		"favicon.ico":          &bintree{assetsFaviconIco, map[string]*bintree{}},
		"index.html":           &bintree{assetsIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
