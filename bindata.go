// Code generated by go-bindata.
// sources:
// template/delete.tmpl
// template/delete_column.tmpl
// template/insert.tmpl
// template/insert_column.tmpl
// template/select.tmpl
// template/select_column.tmpl
// template/table.tmpl
// template/update.tmpl
// template/update_column.tmpl
// DO NOT EDIT!

package sqlla

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

var _templateDeleteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xc1\x4e\xfb\x30\x0c\xc6\xcf\xcd\x53\xf8\x6f\xfd\x0f\xab\x56\xe5\x01\x40\x3b\x6d\x45\x1c\x06\x88\x6d\x12\x07\xc4\xa1\xea\xdc\x51\x29\x4b\xd7\x34\x65\x9a\x4a\xde\x1d\x27\x29\x03\x84\xc6\xa5\x4a\xfd\x7d\xfe\xd9\xfe\x86\x01\xb6\x54\xd5\x9a\x00\x17\xa4\xc8\x12\x82\x73\x5c\xfc\x5f\x16\x7b\x52\xf7\xfc\x81\xab\x19\xc8\xf0\x78\x07\xdb\xcc\x7d\x39\xbc\x96\xcd\x91\x0c\x9b\x85\x3d\x1d\x08\x7e\xb6\x38\x17\x61\xeb\xc7\x25\x74\xd6\xf4\xa5\x85\x41\x24\xec\x91\xa3\xcc\x82\x70\x42\x54\xbd\x2e\x61\xd2\xfe\xea\x66\x39\x85\x88\x98\xa4\x7f\xb0\x19\x6a\xc8\xf6\x46\x5f\xf6\xb0\x25\x69\x33\x91\x38\x3f\x8f\x5d\xa6\xd0\x3b\x02\x39\x6f\x54\xbf\xd7\x5d\x3c\xd6\xd2\xfe\xa0\x0a\x7b\xce\x20\x8a\x08\x32\xca\xa4\xb7\xfe\xce\x4b\xcb\x9e\x47\xa5\xb0\x69\xd6\xad\xe2\x8d\x27\x7c\x74\xad\x77\x19\x3c\xbf\xd4\xda\x92\xa9\x8a\x92\x06\x97\x01\x19\xd3\x98\xd4\xaf\x7d\x7c\x25\x43\x5d\x06\x6f\x5d\xa8\xfa\x90\x5b\x19\x8a\x72\x84\x88\xa4\xae\x82\xf4\x6f\x06\xba\x56\xbe\xe9\xf3\x58\xc4\xcc\x97\x42\xa7\xbf\x4c\x24\x6d\x4f\xe6\xe4\x21\xb8\xc8\x97\xf9\x26\x87\x9b\xd5\xc3\x1d\x7c\x0b\x1c\x03\x2e\x4e\xf5\x44\xc4\x00\x8c\x7d\x53\xfe\x87\xa7\xdb\x7c\x95\x23\x4c\x47\x53\xe4\x8e\x03\x47\x1b\xe0\x35\xc6\x95\x79\x3a\xe7\xf9\x95\xcd\x47\x00\x00\x00\xff\xff\x87\xff\x05\xec\x47\x02\x00\x00")

func templateDeleteTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDeleteTmpl,
		"template/delete.tmpl",
	)
}

func templateDeleteTmpl() (*asset, error) {
	bytes, err := templateDeleteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/delete.tmpl", size: 583, mode: os.FileMode(420), modTime: time.Unix(1469011271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDelete_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x4f\x4b\xc4\x30\x10\xc5\xcf\xc9\xa7\x78\x14\x0f\x2d\x94\xb0\xe7\x42\x4f\xeb\xde\x16\x17\x71\xf1\x22\x1e\xa2\x3b\x8b\x0b\xd3\x34\x4d\xff\xac\x52\xf3\xdd\x4d\x53\xdd\xb2\x28\x78\xea\x34\xef\xfd\x66\x26\x2f\xe3\x88\x03\x1d\x4f\x86\x90\xdc\x12\x53\x47\xeb\x9a\xfb\xca\x24\xf0\x3e\x48\x37\x6d\xa5\x99\xf7\xfa\x85\xe9\x4e\x57\x84\xa2\x84\x5a\xfe\x3e\xd1\xd5\xeb\x50\x70\xac\xb6\xf5\x99\x5c\xc0\xe4\xb1\x37\xaf\x48\x1b\xfc\xc1\x7b\x3f\x0f\x79\xb8\xdf\x66\x93\xae\x7e\xf5\xd9\x9f\x3a\x9e\x7c\xe9\x10\xf5\xfd\x87\xfd\x21\x73\xd0\xbb\x75\x2d\x94\x52\x6d\xc3\xac\xd5\xce\x92\xd3\x5d\xed\xb2\x7f\x26\x61\x94\x62\xd0\x0e\xb5\xc5\x35\x28\xc5\xe9\x08\x26\x93\xc6\xc6\x19\xca\x12\xab\xc9\x2c\x82\xb3\xbc\x78\x37\x4d\xaf\x59\x0a\x0f\xe2\x96\x16\x39\x32\x4f\xab\xe7\xa0\x48\x29\xce\x6f\xe4\x62\x3c\x33\xb5\x09\xe2\xd5\xfa\xcb\xc5\xc6\x47\xcd\x3d\x15\x18\x72\xec\x6c\x11\x96\xca\x31\x47\x5e\x20\xb9\x24\xe2\x7d\xe2\xa5\x68\xd4\xdc\xb6\x84\xb6\x96\xcc\x21\xfd\x3e\xc8\x11\x3f\x99\x14\x8e\xba\xde\x19\x34\x32\xec\x10\xe0\xe0\x99\x1e\xe0\x2b\x00\x00\xff\xff\x42\x59\x12\x68\xd4\x01\x00\x00")

func templateDelete_columnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDelete_columnTmpl,
		"template/delete_column.tmpl",
	)
}

func templateDelete_columnTmpl() (*asset, error) {
	bytes, err := templateDelete_columnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/delete_column.tmpl", size: 468, mode: os.FileMode(420), modTime: time.Unix(1469011271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateInsertTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\xc1\x6e\xdb\x30\x0c\x3d\xdb\x5f\xc1\x09\x1b\x60\xa3\xa9\x3e\xa0\x43\x2f\xeb\x0a\x2c\x40\x96\x75\x75\x6e\xc3\x0e\xae\x4d\x07\x46\x15\x39\x96\xe4\x75\x85\xa7\x7f\x1f\x25\x39\x8e\x96\xae\xc5\xb0\x5e\x02\x45\x24\x1f\x1f\xdf\x13\x3d\x8e\x50\x63\xd3\x4a\x04\xb6\x94\x1a\x95\x61\x60\x6d\x3a\x8e\xe7\xf0\xb6\x2a\x77\x28\xd6\xf4\x03\x17\x97\xc0\xfd\xe1\x17\x98\xee\xca\x5d\xfb\xd3\xaa\x7b\x40\x05\xe7\x73\x7e\x27\xb5\x51\x43\x65\x3a\xe5\x2a\xf6\xaa\x95\xa6\x01\xb6\xc6\x87\x77\xba\xf8\xba\x62\x90\x9d\x80\xe4\xc7\xda\x50\x38\x37\x2b\x8e\x7f\x5d\x8a\x79\xdc\x23\x10\xd3\x88\x92\xb5\x81\x2e\x01\x43\x28\x86\x31\x4d\x28\x87\x4f\x61\x0a\xa4\x89\x46\xf3\xb9\xdc\x27\xba\x17\xa2\xe4\x85\xff\x93\x26\x57\x9d\x18\x76\x52\xc3\xb7\xef\x54\xd8\xca\x6d\x6a\xd3\xb4\x19\x64\x05\x59\xff\xa4\x09\xa1\xe4\x10\x3a\x65\xf9\x0b\x14\xa8\xb7\x42\x33\x28\xf9\x7c\x0e\xa5\x24\x7f\x81\xbf\x80\x7e\x41\x91\xc0\xf4\x02\x62\xaa\xa3\xa5\x88\x75\xf4\xa8\x4e\x95\x72\x8b\xc0\x0f\xe4\xad\xa5\x3b\x83\xbb\xbd\x28\xcd\xec\x5d\x08\x32\xe0\x21\x8c\xb2\x76\x66\x3e\x37\xdb\xcc\x2c\x87\x4d\x57\xf4\x82\x06\xcc\x82\x22\x0b\xd2\x86\xcc\x43\xd5\x94\x15\x12\x0b\x40\xa5\x3a\x95\xbb\x29\x7b\xbd\x80\x1f\xda\xdf\x38\xa7\x7a\x1e\x88\xf3\x4d\x37\xc1\x39\x9c\x34\x69\x1b\x9f\xf1\xe6\x12\x64\x2b\x5c\xdd\x41\x1e\xc6\x4e\xc0\x27\x78\x37\x28\xa1\x0f\xa8\x1e\x1d\x2e\x5b\xae\x8b\xeb\xdb\x0d\x2c\xd7\x9b\x2f\x10\xb9\x0a\x0c\xce\xa0\xd7\xe9\x2c\x77\xa8\x38\x03\xf6\x9e\x05\x62\xd4\x6f\x52\x8c\x38\xf0\x4f\xa5\xbe\xb9\xf7\x6f\xe8\x1f\x54\xb8\xfe\x89\x55\x56\xdf\x4d\x1e\x7c\xfc\x40\x7a\xb8\xec\xe8\x6d\xda\x58\x0b\xf7\x70\x51\x68\x7c\x05\x7c\x7d\xc7\x6f\x51\x0f\xc2\x9c\xc2\x06\xe7\x82\x1e\x0b\x28\xd5\x36\x96\x7c\x72\xeb\x05\x95\x9f\xd0\x8e\x54\xa6\x9c\xb9\xa3\xc3\x23\x0e\x9e\x59\xd4\x8b\x73\x9e\xfb\x65\xfa\x53\xc2\xff\xec\xd7\xd6\x73\xaf\xd0\x9a\xaf\x4a\x6d\x82\x30\xcb\xfa\x35\x73\x1c\x17\x2e\xfa\xf4\x58\x9b\xe5\xb4\x40\x02\x2b\x5a\x59\x7e\x73\x1f\x96\x22\x6b\x6b\xba\xa5\xb7\x2d\x90\x54\xf7\xd3\x45\xee\x1d\xb0\x22\x69\xa6\x8c\x60\x84\xfb\x4a\x1d\xce\xbf\x03\x00\x00\xff\xff\x43\x27\x46\x98\x32\x05\x00\x00")

func templateInsertTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateInsertTmpl,
		"template/insert.tmpl",
	)
}

func templateInsertTmpl() (*asset, error) {
	bytes, err := templateInsertTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/insert.tmpl", size: 1330, mode: os.FileMode(420), modTime: time.Unix(1474539285, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateInsert_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8e\xb1\x0a\xc2\x30\x10\x86\xe7\xe6\x29\x8e\xe0\xd0\x2e\x79\x00\xa1\x53\x27\xa1\x0a\x62\x71\x11\x87\x68\xaf\x50\xb8\xa6\x6d\x9a\x54\xa4\xde\xbb\x9b\x2a\xea\xa0\xe0\xf6\x87\xff\xff\xbe\xdc\x34\x41\x89\x55\x6d\x10\xe4\xca\x0c\x68\x5d\xd6\x92\x6f\x8c\x04\xe6\x50\x2d\x86\x46\x13\x15\xfa\x44\xb8\xd1\x0d\xc2\x32\x05\xf5\x79\xdd\xc0\xb5\x59\x08\xf4\x48\x79\x7b\x41\x1b\x30\x51\x79\x73\x86\xb8\x87\x1f\x3c\xf3\xf3\x93\xdd\x36\x4f\x60\xaf\xc9\x63\x18\xa9\x2f\x59\x51\x3b\x9a\xc7\xf1\x38\x4b\x54\x71\xed\x5e\x78\xf2\xc7\x0a\x93\x88\x7a\x35\xa0\x5b\xeb\xee\x20\xdf\x72\x66\x79\x84\x14\x46\x11\x59\x74\xde\x1a\xe8\x05\x0b\x11\x6a\x34\xe5\x7c\xf2\x3d\x00\x00\xff\xff\x69\xeb\x49\xef\x06\x01\x00\x00")

func templateInsert_columnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateInsert_columnTmpl,
		"template/insert_column.tmpl",
	)
}

func templateInsert_columnTmpl() (*asset, error) {
	bytes, err := templateInsert_columnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/insert_column.tmpl", size: 262, mode: os.FileMode(420), modTime: time.Unix(1469011271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSelectTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\xd1\x4f\xdb\x3e\x10\x7e\x4e\xfe\x8a\xfb\x59\x3f\xa1\x04\x82\xc5\xa4\x69\x0f\x9b\xfa\xb0\x75\x45\x63\x2a\x4c\x10\xa6\x3d\x20\x1e\x42\xea\xb2\x68\xae\xdd\x38\x29\x05\x75\xf9\xdf\x77\x67\xc7\x4d\x58\x3b\xda\x49\x93\x50\x88\xed\xbb\xef\xfb\xee\xbb\xab\xb3\x5a\xc1\x44\x4c\x0b\x25\x80\xa5\x42\x8a\xbc\x66\xd0\x34\xe1\x6a\x75\x0c\xff\xe7\xd9\x4c\xc8\x0b\x7c\xc0\xdb\x01\x70\xfb\xf2\x13\x6a\x3d\xa4\x6d\xfb\x36\xd6\x4b\x61\xe0\xd8\xc7\x57\xb5\x59\xe4\xf5\x3a\x21\xed\x96\x14\xf2\x90\x19\xc8\xa4\x1c\x6a\xb9\x98\xa9\x0a\x06\x70\x73\x8b\x09\x85\xba\x5f\x85\x01\x8a\x30\x99\xba\x17\xc0\xfd\x71\xd3\x30\xdc\x74\xa4\xf8\x9e\xe0\x42\xa8\x09\x49\x6b\xc2\xb0\x7e\x9a\x0b\xc0\x9d\x9e\xc2\xa6\x71\xea\xd3\xcb\x31\x38\x1d\xe0\x70\x3d\x04\x1e\x84\x81\x47\xf7\xd4\x61\xa0\xcd\x04\x4b\x00\xf0\x6b\x59\xcc\x8a\x1a\xd7\x87\x8b\x42\xd5\x6f\x5e\x13\xdb\x74\xa1\x72\x88\xca\x4d\xc2\xcb\x71\x0c\x8e\x35\x8a\x5f\x90\x83\x3a\x8c\xa8\x17\x46\xfd\x39\x06\x43\x82\x32\xc1\x47\x67\x10\xad\x18\xa3\xa7\x2a\x24\xfe\x6b\x5e\x94\xe2\x81\x62\x18\x53\x01\x91\x04\xa7\x7f\x87\xae\x92\xbb\x7a\x07\x70\x20\xd7\x2a\x4b\x62\xda\xd6\x11\xdc\xab\xc5\x6c\x2e\xb3\x7a\x3d\x2c\xee\x90\x01\x77\xc7\x6d\x8b\xf6\x90\x79\xad\xd3\x52\xa2\x6d\x91\x33\x3e\xc1\x96\xa0\x60\x61\xa6\x59\x2e\x56\x4d\x02\xc2\x18\x6d\x62\xd2\x98\xb7\xfc\x38\x52\x2e\xb6\xe2\x9f\x75\xa1\xa2\xd2\x2b\x4b\x80\xe1\x5f\x1c\x06\xcb\xef\xc2\x08\x5c\x3e\x54\x36\x9f\x32\x4a\x6e\x37\x79\x4b\x17\x06\xc5\xd4\x1e\xfd\x37\x00\xb4\x95\xe0\x7d\xd5\xe8\x35\x6d\xd9\x4c\x72\x1b\xcd\x59\x08\xf3\x44\x20\x2c\x1d\x8d\x47\xc3\x6b\x60\x70\x04\x5e\xcd\x11\xae\x4e\xaf\xbe\x9c\x43\x7f\x4a\x2d\xbc\x53\x41\x0c\x8c\x59\x02\x87\x73\x84\x6b\xf8\xf6\x69\x74\x35\x22\x18\x17\x44\x3c\xdd\x71\xc9\xed\x30\x5a\x10\xdf\x98\x9e\xce\x1e\xcc\xf8\xec\xfc\xcc\xc9\x41\x47\x72\xad\x1e\xf8\xa9\x36\xb3\xac\xfe\x8a\x16\x46\x87\x6d\x6e\x02\xaf\x4e\x62\x57\x89\x6f\xac\x43\x00\xf6\x8e\x39\x93\x10\x7a\xcf\xa9\x4a\xd1\x77\x29\xa2\xc9\x1d\x54\xa5\x94\x19\xff\xf8\x01\x5b\x47\xf1\xbd\xdf\x7c\xf3\xac\x6d\xeb\xf6\xe0\x70\x75\x53\xdd\x96\x9b\x40\x66\xee\xfb\x5d\xda\xdd\x9f\x0d\xb2\x76\x4a\xda\x0a\xf5\x92\x80\x26\x77\xfc\x92\xf0\xaf\xf4\x32\xea\x11\x71\xce\xe3\xce\x05\x9e\xe6\x99\x8a\x30\x23\xde\xb3\xf8\xf7\x52\xfe\x56\xf9\xcd\xed\x8b\xb5\x1b\x3b\xad\xb3\xec\x87\xd8\x1e\x79\xd2\xf6\xe6\x5f\x7b\xf4\x7c\x7e\xb1\xc2\x2e\xdd\x3b\xb3\x69\xcb\x5e\x68\x01\x7e\x23\xf0\x9e\x24\x48\x3e\x94\xba\x12\xa4\x63\xaa\xdb\x9d\x0b\xf1\x68\xaf\x40\xca\xed\x09\xf6\x3e\x57\x18\xbb\x85\x66\x93\x87\x88\xc8\x3c\x74\x63\x3e\xc7\xbb\x24\x32\x58\x80\xb1\x43\xec\x83\xcd\xdf\x8d\x2d\x29\xa8\xda\xce\xd1\x42\x09\xb3\x6b\x70\xe9\x3b\x45\xe3\xb4\x11\x14\x06\x6d\x65\x95\xab\x0c\xb5\x6e\xbb\x25\x0f\x30\x99\xe3\x41\x54\xa8\x89\x78\x74\x97\x43\x05\x27\x31\x7d\x0f\xf1\x57\x44\x6c\x2e\xb3\xbd\x2c\x83\x6e\x34\x31\xd3\x39\x41\x1f\x54\x7f\xfe\x2b\x00\x00\xff\xff\xc1\xb2\xca\x2c\xa1\x07\x00\x00")

func templateSelectTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateSelectTmpl,
		"template/select.tmpl",
	)
}

func templateSelectTmpl() (*asset, error) {
	bytes, err := templateSelectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/select.tmpl", size: 1953, mode: os.FileMode(420), modTime: time.Unix(1474534793, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSelect_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x52\xc1\x8e\x9b\x30\x10\x3d\xe3\xaf\x18\xa1\x1e\x40\x65\xad\x3d\x54\x3d\x20\x71\xd8\xcd\x72\xa8\xb4\x2a\xdb\x4d\x54\xa9\xaa\x7a\x70\x93\x89\x8a\xe2\x80\x31\x90\x34\xa2\xfe\xf7\x8e\x0d\x24\xa1\x8d\x94\x54\x3d\xc5\xe1\xcd\x9b\x37\xf3\xde\x74\x1d\xac\x70\x9d\x17\x08\xfe\x1c\x25\x2e\x9b\x59\x29\xdb\x6d\xe1\x83\x31\x04\xbd\xa9\xb7\x42\xca\x85\xf8\x2e\xf1\xa3\xd8\x22\xc4\x09\xf0\xd3\xbf\x5f\xd0\x94\x33\x7a\x48\xf7\x7a\x2e\xf7\xa8\x89\xc6\xd6\x6d\xb1\x84\xa0\x82\x0b\x7c\x63\x7a\x91\xf9\xa7\xe7\xd0\xe2\xfc\xaf\x3e\x8b\xbc\x91\xb6\x2e\xd8\x39\x7c\x71\x50\x23\x33\x02\xfc\xa9\x74\x0d\x9c\xf3\xba\x92\x52\xf0\x4c\xa1\x16\x4d\xa9\xc3\x2b\x4a\xd0\x31\x6f\x27\x34\x94\x0a\xa6\x44\xe6\xe5\x6b\x90\x58\x04\xae\x71\x08\x49\x02\xf7\xb6\xd8\xa3\xca\xe4\x58\x9b\x56\xad\x90\xcc\x33\x80\xb2\xc6\x13\xec\x38\x5f\xef\xbf\x11\xc2\x98\xb7\xff\x81\xda\xd9\xd3\xb3\x52\x02\x27\xe3\x9f\x16\xeb\x3e\x0b\xd9\x62\x0c\xbb\x08\x32\x15\xd3\x50\x11\xf4\x96\xc7\xe0\x1f\x1d\x31\xc6\x37\xcc\xab\x78\xdf\x36\x01\xa1\x14\x16\xab\x60\xf8\x10\x81\xfb\x09\x99\xa7\xb1\x69\x75\x01\x15\xa3\x19\x88\x4c\xeb\xf0\x0f\xf5\xcb\x06\xee\xfe\x21\x86\x97\x4d\xaf\x1f\xa8\x0d\xe4\x45\xf3\xfe\xdd\xff\x19\x6d\x4d\xf8\x23\x39\xea\x7c\x36\x2b\xbf\x92\xfb\x20\x4f\xea\x21\xad\xd5\x75\x77\x40\xab\xdb\xbb\xba\x79\xa3\x4c\xaf\x50\x3f\x1e\x2e\xe8\x50\xff\xd2\x82\x63\xb8\xf6\x7d\xcb\x56\x15\xef\x69\x09\xf8\x90\xbd\x3e\xa5\xaf\xf0\xf8\x05\xce\xd3\x72\xa7\x34\xd4\x8c\x47\xf0\x50\x2f\xdd\xb9\x8c\xe4\xb7\x96\xfd\x30\x9f\xf9\x93\x63\x9a\xa0\x4f\x69\x0f\xb3\x49\xb4\x24\x34\x58\xf0\x3b\x00\x00\xff\xff\x30\xfb\x53\xdf\xae\x03\x00\x00")

func templateSelect_columnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateSelect_columnTmpl,
		"template/select_column.tmpl",
	)
}

func templateSelect_columnTmpl() (*asset, error) {
	bytes, err := templateSelect_columnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/select_column.tmpl", size: 942, mode: os.FileMode(420), modTime: time.Unix(1474539356, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateTableTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x4d\x4b\xf4\x30\x14\x85\xd7\xcd\xaf\xb8\x84\x77\x31\xb3\x78\x33\x7b\xc1\x95\x6e\x84\x61\x50\x46\x71\x1d\xe3\xb5\x53\xcc\x57\xd3\x1b\x8b\xd4\xfc\x77\x6f\x32\x22\x8a\x5d\xf5\xa1\xb7\xe7\x39\xa7\x51\x9b\x57\xdd\x23\x2c\x0b\xa8\xdb\x33\x1f\xb4\x43\x28\x45\x88\xc1\xc5\x90\x08\x36\xa2\x93\x13\xa5\xc1\xf7\x93\x3c\xa3\x09\xfe\x4d\x0a\xe6\x7e\xa0\x53\x7e\x52\x26\xb8\x9d\xe3\x30\xe2\xae\x0f\xff\xa7\xd1\x5a\x2d\xc5\x56\xb0\xf3\x9f\x61\x99\x6d\xc6\x8b\x4b\x50\x0d\x3e\x80\xc2\x55\x7d\xdd\x68\x1f\x66\x4c\xb5\x8e\xde\x63\x9b\xf1\x23\x52\xca\xf1\x6e\x0f\x5c\x98\x0d\xc1\x22\xba\xf9\x84\x09\xa1\xf9\xd5\x63\x65\xc1\x2b\x5f\xb2\x37\x70\xc0\xb9\xfe\xc1\x1f\xff\xfd\x40\xf6\xcb\xb3\xd9\x42\x9e\x30\x55\x23\xab\xc6\xba\x67\xa5\x6d\x29\xa2\x4b\x48\x39\x79\x18\xab\x9d\x3f\x21\x74\xd1\x6a\x42\x90\x47\xb4\x68\x48\x82\xaa\x7b\x7f\x5d\x1e\xe2\x33\x3f\xd7\x2e\x37\x9e\x4b\x57\x33\xd7\x6c\xfb\xce\x7c\x06\x00\x00\xff\xff\xae\xa3\xe9\x6c\x88\x01\x00\x00")

func templateTableTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateTableTmpl,
		"template/table.tmpl",
	)
}

func templateTableTmpl() (*asset, error) {
	bytes, err := templateTableTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/table.tmpl", size: 392, mode: os.FileMode(420), modTime: time.Unix(1469011271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\xc1\x6e\xa3\x30\x10\x86\xcf\xf6\x53\xcc\x5a\x7b\x00\x05\xf9\x01\xb2\xca\x61\x95\x45\xda\x43\x76\xd5\x16\xa2\x1e\xaa\x1e\x10\x19\x52\x24\x62\xc0\x98\x46\x11\xe5\xdd\x3b\xb6\x49\x93\xb4\x4d\x0f\xbd\x20\xdb\x33\xf3\xff\xdf\xcc\x30\x0c\xb0\xc1\xa2\x54\x08\x62\xdd\x6c\x32\x83\x02\xc6\x91\x1e\x7f\xe6\xd9\x0e\xab\xff\xf4\x81\xf9\x02\xa4\x3b\xbc\x80\xa9\x97\xf6\xd9\x9d\x56\xf5\x1e\x35\x25\x73\x73\x68\x10\x2e\x4b\xc6\xd1\x8b\x25\xb7\x2b\xe8\x8c\xee\x73\x03\x03\x67\x94\x23\xa7\x30\x05\x38\xeb\xd0\xfc\xcb\x1a\xd6\xb5\x55\x95\xc9\xc4\x5d\x38\x5b\xd6\x55\xbf\x53\x1d\x3c\x3c\x52\x61\xa9\xb6\x7c\xe4\xbc\xe8\x55\x0e\x41\xfb\xc1\x84\x54\x42\xf0\x4e\x41\xf8\x05\x02\x79\x6b\x34\xbd\x56\xd7\x73\x28\x85\x7d\x22\x3f\x87\x36\xa2\x88\x27\x9d\xc3\x39\xea\x30\x52\x64\xb4\x78\x54\xa7\x33\xb5\x45\x90\x47\x78\x37\x42\x83\xbb\xa6\x22\xf5\xe3\x64\x7d\x50\x80\xf4\x61\x54\x1b\x3b\xbd\x6b\xbd\xbd\x91\x85\x90\xd6\x49\x5b\x51\x83\x81\x9f\x48\x44\xb3\x29\x95\x41\x5d\x64\x39\x12\x05\xa0\xd6\xb5\x0e\x6d\x97\xc4\x39\x21\x44\xd0\x3d\x77\x2e\x64\xf7\xd7\x4a\xdf\x81\x4c\xeb\x49\xd7\x0a\x72\x56\x16\x2e\xe3\xc7\x02\x54\x59\x59\x81\xe3\x9c\x84\x78\xe7\x32\xf9\xd8\x8e\xd9\xfe\x09\x35\x92\xf8\xfe\xc2\xc1\xbd\xca\x89\xf5\x9b\xd2\x9c\xb5\x3d\xea\x83\x15\x14\xeb\x9b\x3f\xbf\xd3\x18\xce\x7e\x1a\x48\xe2\x54\xc0\x0c\x4e\x5d\x3a\x1b\x8f\x63\x9d\x84\x70\x46\x5e\x63\x46\x77\xb8\xff\x1b\xdf\xc5\xb6\xc6\x27\x79\x8f\x09\x64\x4a\x03\xf1\x8b\x88\xb2\xa6\xa1\x85\x04\x6e\x68\xd4\x97\x94\x32\x8c\x2c\x39\xed\xf7\xb4\xab\xd7\x00\x00\x00\xff\xff\xf2\xc8\x0d\xb3\x2d\x03\x00\x00")

func templateUpdateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateUpdateTmpl,
		"template/update.tmpl",
	)
}

func templateUpdateTmpl() (*asset, error) {
	bytes, err := templateUpdateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/update.tmpl", size: 813, mode: os.FileMode(420), modTime: time.Unix(1469011271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUpdate_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\xcf\x4b\xc3\x30\x14\x3e\x37\x7f\xc5\x47\xf1\xd0\x42\x09\x3b\x17\x7a\x1a\xbb\x4d\x87\x6c\xea\x61\xec\x10\xdd\x1b\x0e\xb2\x34\x4d\xd3\x4e\xa9\xf9\xdf\x4d\x5a\x5d\x29\x0a\xce\x53\x5f\xf3\xfd\xc8\xf7\x5e\x5e\xd7\x61\x4f\x87\xa3\x22\xc4\x0f\x7a\x2f\x2c\xcd\x4b\xd9\x9c\x54\x0c\xe7\x3c\x74\x53\x9f\x84\x94\x1b\xf1\x2c\xe9\x4e\x9c\x08\x79\x01\x3e\xfe\x7d\xc0\x96\x73\x5f\xc8\xbe\x5a\x96\x67\x32\x5e\xc6\x0e\x8d\x7a\x41\x52\xe1\x17\xbd\x73\xc3\x25\xeb\xfb\x65\x8a\x35\x59\x4f\xe1\x3f\xac\x36\x47\x2b\x03\x35\x69\x83\x05\xdf\xbc\xeb\x6f\x71\xfa\x87\x27\x3a\x16\x55\xbc\x26\x7b\x2b\xf4\x36\xbe\x98\x3b\x17\xef\x50\xa0\x65\x91\x21\xdb\x18\x85\x8a\x39\x76\x75\xcc\xa7\x57\x32\xf4\xcf\xa0\x19\xe8\x4d\x9b\x1a\x9c\xf3\xba\x92\x52\xf0\x95\x26\x23\x6c\x69\xae\xe9\xa0\x15\x06\xa5\xc6\x54\xc8\xa2\xe3\x01\x92\x54\xd2\x1b\xa7\x28\x0a\xcc\x02\x39\xf2\xcc\xe2\xc2\x5d\x54\x8d\x90\x2c\x72\x20\x59\xd3\x08\xf7\x9a\xed\x6c\xe7\x11\xc6\xa2\x73\x68\x28\x3c\xe5\xa0\x5a\x78\x70\x12\x7f\x6c\xac\x7b\x14\xb2\xa1\x1c\x6d\x86\x95\xce\x7d\xa8\x0c\xc3\x7a\xe4\x98\x4c\xd7\x85\xb1\x0f\xb6\x05\x84\xd6\xa4\xf6\xc9\xd7\x41\x86\xfe\x93\x4e\x67\xef\xc5\x9e\x13\x96\xe5\x33\x00\x00\xff\xff\x12\x0e\xef\xee\x80\x02\x00\x00")

func templateUpdate_columnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateUpdate_columnTmpl,
		"template/update_column.tmpl",
	)
}

func templateUpdate_columnTmpl() (*asset, error) {
	bytes, err := templateUpdate_columnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/update_column.tmpl", size: 640, mode: os.FileMode(420), modTime: time.Unix(1469011271, 0)}
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
	"template/delete.tmpl": templateDeleteTmpl,
	"template/delete_column.tmpl": templateDelete_columnTmpl,
	"template/insert.tmpl": templateInsertTmpl,
	"template/insert_column.tmpl": templateInsert_columnTmpl,
	"template/select.tmpl": templateSelectTmpl,
	"template/select_column.tmpl": templateSelect_columnTmpl,
	"template/table.tmpl": templateTableTmpl,
	"template/update.tmpl": templateUpdateTmpl,
	"template/update_column.tmpl": templateUpdate_columnTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"delete.tmpl": &bintree{templateDeleteTmpl, map[string]*bintree{}},
		"delete_column.tmpl": &bintree{templateDelete_columnTmpl, map[string]*bintree{}},
		"insert.tmpl": &bintree{templateInsertTmpl, map[string]*bintree{}},
		"insert_column.tmpl": &bintree{templateInsert_columnTmpl, map[string]*bintree{}},
		"select.tmpl": &bintree{templateSelectTmpl, map[string]*bintree{}},
		"select_column.tmpl": &bintree{templateSelect_columnTmpl, map[string]*bintree{}},
		"table.tmpl": &bintree{templateTableTmpl, map[string]*bintree{}},
		"update.tmpl": &bintree{templateUpdateTmpl, map[string]*bintree{}},
		"update_column.tmpl": &bintree{templateUpdate_columnTmpl, map[string]*bintree{}},
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

