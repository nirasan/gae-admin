// Code generated by go-bindata.
// sources:
// templates/index.html
// templates/toml.html
// DO NOT EDIT!

package bindata

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x91\xc1\x4a\xc4\x40\x0c\x86\xef\x3e\x45\xec\x49\x41\xcc\x61\xaf\x71\x40\xd4\x83\x20\x28\xb8\x8b\x78\xec\x6e\xb3\x6d\x20\x33\x59\xa6\x53\x50\x4a\xdf\x5d\xda\xa1\xdb\xba\xa7\x99\xfc\x21\x1f\x5f\x66\xe8\xfa\xf9\xfd\x69\xfb\xfd\xf1\x02\x4d\xf2\xea\xae\x28\x1f\x00\x00\xd4\x70\x59\xe5\xeb\x54\x26\x49\xca\xee\xb1\xf2\x12\x40\x42\xc5\x3f\x70\x2a\x6b\x26\xcc\x79\x1e\xc1\x65\x86\xf6\x56\xfd\x2e\xe3\x7d\x0f\x72\x84\xfb\x5d\xcb\x11\x86\xe1\x1c\x7f\xb1\x1e\xcc\xf3\xdd\xd8\x9f\x9b\x70\x43\x25\x34\x91\x8f\x0f\xc5\x98\xbe\x59\x6d\x5d\xda\x45\x85\x61\x28\x5c\x2b\x75\x00\xeb\x12\x61\xe9\x6e\x17\xb9\x66\xe3\x0e\x16\x12\x87\xd4\x12\x36\x9b\x95\x76\xa7\x4b\x31\x05\x2a\xee\xcc\xc7\x72\x5c\x07\x93\x79\x2d\xdc\xd6\xbc\xc2\xab\x3f\x59\x4c\x1c\x47\x3e\xa1\xca\x8a\x84\x6b\x54\xdf\x03\x6b\xcb\xeb\x5d\x2e\xad\x25\xcc\xd2\x9f\xa3\xb4\x04\xb0\x08\x91\x6b\x69\x67\xfe\x3e\xfe\x07\x86\x6a\xe6\x11\xe6\xe7\x23\x9c\x3e\xe4\x2f\x00\x00\xff\xff\x76\x1a\xfd\x74\xa7\x01\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 423, mode: os.FileMode(420), modTime: time.Unix(1479784498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tomlHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xc1\x4e\x44\x21\x0c\x45\xf7\x7e\x45\xed\xde\xf0\x03\x30\x89\x51\xd7\xba\x70\xe3\xb2\x33\x74\xf2\x9a\x50\x4a\x98\x62\x32\x7f\x6f\x1e\x44\x47\x37\xae\x80\x92\x73\x2e\x97\x78\xff\xfc\xfa\xf4\xfe\xf1\xf6\x02\x9b\x6b\x39\xdc\xc5\xb5\x00\x00\xc4\x8d\x29\xaf\xed\x3c\xba\x78\xe1\xc3\x63\x56\xa9\xe0\xa6\x05\x44\x9b\x75\xe7\x1e\xc3\xba\x5a\x54\xb8\x61\xf1\x68\xf9\xfa\xcb\x70\xb6\xae\xc0\xf5\xe4\xd7\xc6\x09\x75\x14\x97\x46\xdd\xc3\x3e\x7f\xc8\xe4\x84\x40\x27\x17\xab\x09\x03\xed\x39\x61\xcf\x41\x50\xf6\xcd\x72\xc2\x66\x17\xc7\x9b\x6f\x3a\xa5\xb6\xe1\xb0\x8c\x67\x29\x8c\x50\x49\x39\xe1\x68\xc5\x28\xaf\x49\xf8\x87\xb9\x8c\xa3\x8a\x23\x7c\x52\x19\x3f\xd8\x1f\x24\xce\xf7\x7d\xb7\x5b\x95\x62\x98\xff\xf4\x15\x00\x00\xff\xff\x3f\xae\x56\x02\x3e\x01\x00\x00")

func tomlHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tomlHtml,
		"toml.html",
	)
}

func tomlHtml() (*asset, error) {
	bytes, err := tomlHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "toml.html", size: 318, mode: os.FileMode(420), modTime: time.Unix(1479784189, 0)}
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
	"index.html": indexHtml,
	"toml.html": tomlHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"toml.html": &bintree{tomlHtml, map[string]*bintree{}},
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

