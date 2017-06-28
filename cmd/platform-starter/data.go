// Code generated by go-bindata.
// sources:
// config/.csscomb.json
// config/.editorconfig
// hooks/pre-commit
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

var _configCsscombJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x41\xae\xdb\x30\x0c\x44\xf7\x39\x85\xa0\x75\x75\x81\x9c\xa5\x1b\x4a\x1e\x27\x42\x68\xd2\xa5\xe8\x1a\x49\xd1\xbb\x17\x4e\xea\x16\xf6\x77\xb2\xe5\x3c\xce\x88\xa4\x7e\x9d\x42\x08\x21\x1a\x06\xfd\x89\x84\x61\xf4\x7b\xb2\x89\xd1\xe0\x2d\x9e\x83\xdb\x84\x6f\x2f\x84\x78\xa6\x7b\x4b\x0d\x43\x2d\xca\x2a\x5b\x75\x29\x59\x2a\xd4\x10\xcf\x21\xb2\xce\xb0\xf8\x57\xca\xac\xe5\x96\xaa\x74\x10\x5f\xc4\x67\x71\xd3\xd6\xae\x6a\x7e\x25\xe9\xe2\x39\xf4\xc4\x6d\x35\x05\x63\x80\xf8\xb1\x2d\xb4\x4f\x82\x99\xab\x60\xfb\x14\x06\x75\x55\x2e\xe9\x01\xd3\x9d\xe1\x8f\x49\x1d\xcb\x5c\xb1\x55\xb9\x30\x56\xaf\xa6\xe6\x49\xad\x83\xa5\x9e\x98\x33\x95\xdb\x02\x51\x2e\xff\x88\x91\x0a\x52\x46\xaf\x86\xb4\xce\x1f\xb7\x2a\xf5\x0e\xfb\x2f\x86\x37\xbd\x43\xae\x42\xae\x76\xc0\xac\x0e\x1f\x90\x0c\x9f\x01\x49\x1d\x0a\x93\x91\x57\x95\xe7\x40\xdf\xe5\x30\x4e\x47\xc8\xb2\x8c\x6c\x54\xf0\x36\xf1\x0b\xb5\x77\x7b\x61\x0d\x8c\xe2\x6a\xa9\x03\xd7\xa1\x3a\xec\x43\xf2\x31\x7c\xbc\x13\xd6\xf6\x2e\xde\xad\x8e\xe9\x49\xef\xfe\xa3\x53\x4e\xad\x3e\x76\xc7\x9f\xa4\x3a\xa3\xb5\xf5\xfa\x8b\x74\xfa\x7d\xfa\x13\x00\x00\xff\xff\x77\x38\x8d\xf7\xe7\x02\x00\x00")

func configCsscombJsonBytes() ([]byte, error) {
	return bindataRead(
		_configCsscombJson,
		"config/.csscomb.json",
	)
}

func configCsscombJson() (*asset, error) {
	bytes, err := configCsscombJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/.csscomb.json", size: 743, mode: os.FileMode(436), modTime: time.Unix(1498571809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEditorconfig = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\x05\x81\x6c\x0f\x7e\x1e\x8a\x0e\x5d\x32\xb5\x19\x3a\xf4\x0b\x02\x43\x50\x2c\xca\x66\x22\x53\x86\xc8\x24\x4d\x83\xfc\x7b\x21\xd7\x6d\xd1\x14\xcd\x46\x5c\x93\xe7\x1e\x6b\x01\x2b\x4f\x9a\xf2\x63\xe2\x40\x1d\xf4\x18\x47\x01\x8f\x07\x8c\x69\xc4\x5c\xc6\x40\x8c\xe0\xd8\xc3\xe0\x88\xd5\x11\x43\x9b\x58\x48\x14\x59\xcd\x02\xda\xe4\x89\x3b\x10\x3d\x45\x14\xd8\xa0\x1e\x11\x19\x3c\x85\x80\x19\x59\x01\x27\xbc\x4c\x84\xe7\xa7\x95\x98\xc5\x1c\xb5\x53\x63\x9d\x72\x67\x4c\x4e\x49\x61\x09\x9a\xf7\x68\xcc\xfa\x5f\x63\x90\xbd\x4d\xc1\xc6\xd2\xbd\x84\x18\x4c\xdb\xbb\x2c\x58\x96\xf6\x1a\xfe\x3f\x18\xcd\x34\x58\xcd\x8e\x22\x71\x67\x8f\x3d\x29\xca\xe8\x5a\xfc\xa4\x10\x0b\x66\xb5\x81\xd8\x45\xcb\x78\x9c\x49\xf3\x37\x8f\xac\x76\x52\x86\x25\x4c\x77\x5f\x21\xbd\x95\xec\xbe\x68\xd4\xfd\x46\x9a\xbf\x48\xc1\x45\x99\x64\xeb\xf3\x56\x2a\x95\x6a\x2b\xaf\xd5\x69\x88\x97\xe6\x0a\x75\xf7\xb1\x54\x5e\xa4\x1a\xfc\xa5\xb9\xa5\xfe\x0d\xed\x52\x73\xed\xa9\x6e\x63\xcc\xfa\x3c\xba\x76\xe7\x3a\xac\xb7\x92\xb8\xaa\x35\xbb\x03\x49\xfd\xb3\xf8\xc6\x8f\x15\x9b\xf3\x8b\xdb\x61\xa0\x88\x15\x0c\xf3\xf4\xfb\xb8\xb4\xbd\x07\x00\x00\xff\xff\x74\xd1\xff\x15\x1d\x02\x00\x00")

func configEditorconfigBytes() ([]byte, error) {
	return bindataRead(
		_configEditorconfig,
		"config/.editorconfig",
	)
}

func configEditorconfig() (*asset, error) {
	bytes, err := configEditorconfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/.editorconfig", size: 541, mode: os.FileMode(436), modTime: time.Unix(1498638997, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hooksPreCommit = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x96\x6f\x6f\xdc\x36\x12\xc6\x5f\x9b\x9f\xe2\xa9\x76\x91\x8d\x03\x4b\xeb\x18\x6e\x90\x3f\x75\x0e\x3e\xc7\xb9\x5b\xc0\x89\x0f\xb5\x81\xbe\x48\x83\x80\x12\x47\x2b\xc6\x12\xa9\x92\xd4\xae\xb7\x08\xfa\xd9\x0f\x43\x49\x6b\x6f\xed\x43\x73\xf5\x2b\x4b\x24\x87\xcf\x3c\xf3\x9b\xd1\x4e\x7e\x98\xe7\xda\xcc\x7d\x25\x26\x62\x82\x53\x03\xba\x95\x4d\x5b\x13\x2a\x6b\x6f\xe0\x0b\xa7\xdb\x80\x60\xb1\x22\xa7\xcb\x0d\xd6\x95\x0c\xd0\x1e\x32\xb7\x5d\x7c\x9f\x13\x0a\xdb\x34\x3a\x04\x52\x99\x98\xe0\x4c\xd6\x35\x29\xe4\x1b\x24\x4b\x1d\x86\xb5\x04\x6b\x1d\x2a\x18\x0b\xe9\x96\x5d\x43\x26\xf8\x0c\xb8\xae\xc6\x5b\x2a\xdb\xd5\x4a\x4c\x40\xb7\x3a\x8c\x5b\x4d\xfa\x3b\x39\x0b\x1f\x64\xe8\x3c\x64\x19\xc8\x41\x7b\xdf\x69\xb3\x84\x34\x90\x6d\xeb\x6c\xeb\xb4\x0c\x84\x86\xbc\x97\x4b\x82\x2e\xc5\x04\x1c\x41\x9a\xe0\x59\x9c\x0f\xb6\x45\xa8\x46\x89\x59\xcc\xf1\xda\x82\x8c\xcc\x6b\x42\xa8\xb4\x8f\x0a\x0e\xe0\xc8\xc8\x66\x78\x53\x6a\x5e\xb3\x48\x5a\x47\xe9\x90\x40\x26\x84\x2e\xc1\x19\x39\x5a\xa5\xad\x74\x9e\x90\xa6\x83\x29\xff\x3e\x3f\x7d\x87\xb7\x73\x45\xab\xb9\xe9\xea\x1a\x47\x6f\x9f\x3c\x17\xa1\x22\x23\xf6\xe4\x52\x6a\xe3\xc3\x09\x6f\x11\x54\x7b\x12\x7b\x13\x2c\x8c\x0e\x5a\xd6\x83\xaa\xd7\x50\xba\x2c\x31\xec\xe4\xdc\xa8\x69\xc3\x06\xc1\x11\xc1\xe6\x5f\xa9\x08\x77\x71\x8e\xf3\x97\x47\x3f\xaa\xe2\xc5\xf1\x51\x91\xbf\xa0\xfc\x95\x3c\x7c\x71\x48\x3f\x1e\xe7\xe5\x4b\xf5\xe2\xd5\xd1\xcb\x97\x65\x4e\x74\xfc\xea\xf0\x58\x94\x5a\x88\x09\x16\x25\x36\xb6\x8b\x86\x70\x46\xb2\xae\xed\x3a\x7a\x7b\x7a\x75\xb6\x58\xc4\x4c\x39\x6f\x0f\x4f\xa1\x4f\x7e\x25\x9d\xee\xcd\xb1\x08\xae\xa3\x4c\xc4\x43\xc6\x1a\xe9\x0b\xad\x4f\xa6\x4f\xfb\xb2\x9a\x52\x2f\x91\xa6\xb9\xb5\x75\xb4\xd0\x67\x3b\xfb\xf6\xf9\xf6\x33\x67\xbd\x47\x5b\xcb\x50\x5a\xd7\xa0\x75\x96\x93\xf1\x08\x64\x54\x94\xb3\xb2\x5a\x3d\x26\xe7\x0d\x5a\x47\x2b\x32\x41\x4c\xb8\x7c\x0d\x4a\x67\x1b\xe4\x14\x6b\xaf\x14\xc5\xd3\x5c\x57\x47\xad\xf5\x3a\x58\xb7\xc9\xf0\x0b\x81\x6e\xdb\xda\xea\x10\x97\x4a\x59\xf0\x3f\x32\x3e\x89\x09\x5a\xa7\x4d\x88\x99\x39\x69\x96\xc4\x60\xb9\xe0\xd1\xaf\xc3\xb7\xb2\x20\x14\x95\x74\xb2\x60\xd2\xa4\x51\x20\xa3\x7c\x4f\x63\xd0\xb5\xa2\x8c\x09\xf8\x84\x64\xba\x93\x68\x82\x1f\x4e\x90\xb0\x53\x09\x3e\xe3\xc9\x13\xae\xef\x47\x1b\x68\x7b\x35\x3a\x4f\xb0\x25\x72\x27\x8b\x1b\xe2\x1b\x9d\xed\x8c\x82\x44\x70\x83\x14\xed\x61\x6f\x50\x91\xa3\x03\x3c\xd5\x61\xe6\x39\x08\xe7\x0f\x47\xbf\x75\xda\x91\x3a\x40\x69\x1d\x5a\xeb\x82\xcc\x75\xad\x99\x0e\x8b\x2b\x5b\x4b\xa7\x3d\x9e\x1f\xce\x3c\xe6\x9d\x77\xb1\x93\x83\xdb\x3f\x80\xd7\xa6\x88\xa8\xc5\xdc\x7e\xeb\xa4\xa3\x51\x00\xf2\x4d\x20\x8f\x4a\xb6\x2d\x19\x0e\x53\xca\xba\x86\x36\x71\xab\x22\xaf\x97\x46\x06\x52\xbd\xb4\x4c\xec\x05\xf2\x01\x7d\xd5\x23\xa7\x69\x5a\xc8\xa2\x22\x85\x34\xe5\x5a\xa5\xd6\xd4\x1b\xa4\x29\xaf\xa5\xa5\xae\x03\xb9\x93\x53\xa4\xbf\x63\x3a\x02\xfd\x4d\xec\x01\x17\x67\x5f\x4e\x2f\x2e\x4e\xce\x38\xe9\x54\x61\xf6\x09\xe9\x1f\x9f\x7f\x3d\x9c\xe1\x1b\xd6\x05\xd2\x62\x9f\x6d\x3c\x1c\x9a\xa6\x90\x01\x3f\xfd\xf4\xeb\xf9\xe5\x7b\x71\xee\x9c\x75\xaf\x71\x1a\x02\x37\x45\x84\x46\xb1\x77\xbb\xd0\x80\x95\x64\x42\x5c\x33\xc1\x85\x34\x28\x24\xbb\xde\x3a\x9b\xd7\xd4\x78\xe8\xdd\x36\x58\x5b\x77\xd3\x17\xb6\x25\xcb\xb3\xce\x1a\xd8\x50\x91\xdb\xb2\xea\x39\x58\x1c\x6e\xbd\xe7\x35\xf1\x5c\xe1\xb9\xa7\x56\xda\x8f\xfd\xb1\x1d\x1a\x14\x55\x64\x42\x0c\xfd\x76\x63\xec\xba\x1f\x95\xfc\xc4\xe6\x2b\xcb\xec\xf2\x13\xcb\x53\x63\x8c\xa8\xb7\xa2\xe2\x06\x9d\xd7\x66\xf9\x5a\x08\xe0\x5e\x7f\x3d\xd2\x58\xb1\x29\x05\x5b\xb3\x17\xa7\xe5\xf3\xbb\x4e\xe7\x0c\x28\xde\xb6\xae\x74\xa0\x1e\x69\x62\x03\xfd\x41\x8f\x7f\x94\x6a\xcb\x92\x8c\x62\x39\x5b\xeb\x7c\xe4\xbd\x94\xba\xce\x04\xdd\x52\x81\xb1\xda\xa9\x36\x8a\x6e\xb9\xe6\x51\xe4\xb6\xf6\xdb\xe2\xa6\xa9\x10\x1f\x2e\xdf\x2d\xde\x2f\xce\xdf\x7d\x79\xbf\xb8\x38\xbf\x3a\xd9\x61\xe5\x0e\x11\x9e\x80\x7f\xe0\x1b\x96\x8e\x5a\xa4\x2b\xcc\x32\xbf\x5a\xce\x76\x5e\xb4\xe6\x4f\x2f\xbe\xb6\x7f\x7a\xe1\x8d\x6c\x67\xfd\x6c\x89\x82\x0a\xef\x0b\xdb\xe4\x42\x9c\x5d\x5d\x9d\x5d\x7e\xf8\xe7\xc9\xf4\xe9\xba\xd2\x45\x35\x2e\xe0\xe8\x2d\xb6\x73\x79\x3f\x36\xf0\x27\x86\x33\x99\x0e\x07\x12\x7c\xfe\xfc\x06\x11\x3b\x0c\x7f\x0f\xe9\x1b\x83\x69\x0f\x63\x03\x38\xf1\xf8\x95\xcb\x84\xf8\x85\x29\xda\x48\x67\x62\xe9\xfa\x3f\x7e\xc4\xb2\xb6\xb9\xac\x23\xad\x5b\x91\x71\xf3\xc7\xff\x7c\x78\xbd\xdd\x6a\xda\x66\x0c\x87\x74\x79\xb7\xf3\x34\xb7\x2e\x70\x89\xc6\x0f\x17\xab\x19\x0f\xed\xd4\xbd\x2f\x7a\xa3\x97\x55\x60\x5c\x9b\xae\x0e\x9a\x91\xf6\x5d\xae\xb4\xa3\x22\x58\xa7\x69\x18\x63\xdb\xc5\x31\xa1\x1e\xb3\xce\xc9\xa0\xad\xf1\x82\x67\xcc\xb0\xf4\x65\x20\x50\x1b\x4c\x9f\x96\xda\x28\x64\x88\xc5\xc4\x2c\x1b\xb6\x64\x5f\xbd\x35\xb3\xfd\x37\x50\xf6\xce\xbb\x21\x70\x9a\x0e\xe7\x93\xe9\x6e\xc0\x04\x69\x5a\x33\x89\xc9\x74\x17\x9b\x64\x1b\xa3\x2f\xd2\xf4\x1f\x48\x0d\xe1\xf0\x61\x7d\xb6\x36\x14\x95\x45\x12\x4b\xc4\xd3\xb1\x91\x21\x3a\x16\x1b\x69\x14\x72\xc0\x3f\x54\x76\x9c\x4c\x1e\xc6\xe9\xed\x1c\x1f\x4b\x2d\x94\x35\x74\xc7\x18\x29\xfe\xc0\xf4\xf2\x85\x38\x7f\xb7\xb8\xbe\xfc\xf9\xec\xf2\xe3\xfb\xc5\xbf\xbe\x5c\x5f\x5e\x5e\x5c\x6d\x99\xbb\xbf\x31\x0d\xd6\xd6\xfe\x7f\xe3\xf7\x30\xcc\x77\x91\xf8\xc8\x15\x7f\x13\xca\x87\x91\xbe\x83\xcf\xc7\x0e\x7d\x2f\xaa\x0c\x57\x1c\x38\xda\x3c\x2c\xfe\x0e\x44\x8f\x24\xd9\x8f\x9f\x64\xca\x01\x12\xf1\x7f\xa3\x72\xcf\xc8\xeb\xed\x90\x6c\xb4\x1f\xb0\x21\x15\x95\xf1\x94\xac\x49\x7a\x9e\xe6\xb7\xf1\x77\x47\x86\x9f\xa9\xa1\x26\x27\x77\x37\xba\xad\x88\x53\x3b\xdf\xe0\x6b\xe7\xc3\x80\x5b\xfc\x04\x58\x9e\xd3\xa3\x13\xd2\xa8\x7b\xee\x3f\x92\x12\xdf\xd1\xca\x50\xcd\x83\x9d\x37\x9b\x79\x14\x30\x7f\xf6\x6c\xfe\xec\x2f\x2c\xfd\x0b\x6c\xff\x1b\x00\x00\xff\xff\x1a\x20\x17\xe7\xd0\x0b\x00\x00")

func hooksPreCommitBytes() ([]byte, error) {
	return bindataRead(
		_hooksPreCommit,
		"hooks/pre-commit",
	)
}

func hooksPreCommit() (*asset, error) {
	bytes, err := hooksPreCommitBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hooks/pre-commit", size: 3024, mode: os.FileMode(509), modTime: time.Unix(1498639133, 0)}
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
	"config/.csscomb.json": configCsscombJson,
	"config/.editorconfig": configEditorconfig,
	"hooks/pre-commit": hooksPreCommit,
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
	"config": &bintree{nil, map[string]*bintree{
		".csscomb.json": &bintree{configCsscombJson, map[string]*bintree{}},
		".editorconfig": &bintree{configEditorconfig, map[string]*bintree{}},
	}},
	"hooks": &bintree{nil, map[string]*bintree{
		"pre-commit": &bintree{hooksPreCommit, map[string]*bintree{}},
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

