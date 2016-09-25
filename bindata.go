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

var _templateDeleteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x93\xd1\x6e\xd3\x30\x14\x86\xaf\xe3\xa7\x38\x58\x20\x25\x5a\xea\x07\x00\xf5\x86\x35\xd3\x2e\xc2\x58\x9b\x22\x2e\x10\x17\x59\x72\x52\xac\xb9\x4e\x63\x3b\x8c\xc9\xf3\xbb\x23\xdb\x69\x1b\x98\x8a\xb8\x8b\x8e\xcf\xf9\xfe\xdf\xbf\x4f\xac\x85\x16\x3b\x2e\x11\xe8\x0a\x05\x1a\xa4\xe0\x1c\xb1\x76\x01\x6f\x9b\x7a\x8f\xe2\xae\xde\x23\xbc\x5f\x02\x0b\x1f\x2f\x60\xfa\x6b\x5f\x86\x17\xf8\x22\x0d\x37\x02\x61\x71\xea\xef\xa5\x36\x6a\x6c\x4c\xaf\xfc\xc4\x41\x71\x69\x3a\xa0\x77\xf8\xf4\x4e\x57\xeb\x92\x42\xfa\x1a\xb2\xf5\x88\x2c\x30\xcc\xf3\x01\xc1\xda\xb9\xae\x73\xd1\x53\xb5\x2e\x21\xa2\xc1\x92\xe4\xef\x9e\x6a\x5d\x12\x47\x48\x37\xca\x06\xd2\xe1\x15\xa2\x5a\x97\x19\x44\x4e\x9a\xfd\x43\xc0\x92\x44\xa1\x19\x95\xbc\xdc\x63\x49\x92\x0c\x39\x49\x9c\xd7\xb3\x16\x54\x2d\x77\x08\xec\xba\x17\xe3\x5e\x6a\x70\xce\x5a\x30\xb8\x3f\x88\xda\x9c\xf2\x8c\x87\x14\x58\x3c\x46\xd9\xfa\x80\x2f\x99\x3d\x49\x65\xb0\xed\xab\x41\xa4\x19\xa4\xda\x28\x2e\x77\x39\x7c\xfb\xce\xa5\x41\xd5\xd5\x0d\x5a\x97\x03\x2a\xd5\xab\xcc\xdb\x7e\xfa\x81\x0a\x75\x0e\x3f\x75\xa8\xfa\xf0\x07\x16\x8a\x6c\x82\x90\x84\x77\xe1\xe8\xcd\x12\x24\x17\x7e\xe8\x78\x59\x4a\x73\x5f\x0a\x93\xfe\x66\x24\x19\x46\x54\xcf\x1e\x42\x57\x45\x59\x6c\x0b\xb8\xd9\x7c\xfe\xe4\x9d\xb2\xc9\x24\x0d\xb8\xa8\xea\x89\x94\x06\x60\x9c\xbb\x5a\x02\x85\xaf\xb7\xc5\xa6\xa0\x70\x35\x35\x45\xee\x24\x38\xb5\x01\xfd\x40\xa3\x65\xc9\x45\xcc\x73\x01\xbc\x03\x76\x5b\xeb\xfb\xc7\x73\x46\x3a\x28\x57\xe1\xf5\x27\xfd\xd3\x73\xb6\x0f\xa0\x07\x21\x6a\xb6\xfa\xe8\x73\x1a\x04\xdb\xa0\x1e\x85\x99\x87\x13\xe4\x72\xa8\xd5\xee\x9c\x4e\x48\x7d\xb6\xad\xce\xa5\x19\x3b\xae\x08\xf3\x7a\xf7\x8f\xf1\xdd\x2e\x6c\xac\x9f\xd0\x7f\x36\xde\x70\x14\xed\xd1\xe0\x7f\xe4\x3e\xcf\xfc\x58\x6b\x1f\x58\xf1\x0b\x9b\x74\xe6\x99\x31\x96\x91\xf8\x83\x4d\xab\x73\x5e\xa2\xdf\x01\x00\x00\xff\xff\xf6\x40\x0d\x3d\xbc\x03\x00\x00")

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

	info := bindataFileInfo{name: "template/delete.tmpl", size: 956, mode: os.FileMode(420), modTime: time.Unix(1474801346, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDelete_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\xc1\x4b\xc3\x30\x18\xc5\xcf\xc9\x5f\xf1\x28\x1e\x5a\x28\x61\xe7\x42\x4f\x73\x37\x71\x88\xd3\x8b\x78\x88\xee\x1b\x16\xbe\xa6\x69\x9a\xce\x69\x9a\xff\x5d\xda\xce\xcd\xa1\xe0\xa9\xf4\x7b\xdf\xef\xbd\xe4\x25\x04\x6c\x69\x57\x19\x42\x72\x4d\x4c\x9e\x96\x0d\xf7\xb5\x49\x10\x63\x08\xb8\xea\x6a\xcd\xbc\xd1\x2f\x4c\xb7\xba\x26\x14\x25\xd4\xf9\x6f\x80\x6f\x96\xba\x26\xc6\x80\x07\xe3\x2b\xcf\x84\x18\xe5\xae\x37\xaf\x48\x5b\xfc\xc1\xc7\x38\x87\xdc\xdf\xdd\x64\xa3\xae\x7e\xf9\x6c\x8e\x2e\xe9\x7e\xd2\x37\x1f\xf6\x9b\xcc\x41\x07\xeb\x3a\x28\xa5\xba\x96\x59\xab\xb5\x25\xa7\x7d\xe3\xb2\x7f\x92\x10\xa4\xd8\x6b\x87\xc6\xe2\x12\x94\xa2\xda\x81\xc9\xa4\x93\x71\x86\xb2\xc4\x62\x5c\x16\x8d\x45\x79\xda\x5d\xb5\xbd\x66\x29\x22\x88\x3b\x3a\xcb\x13\xf3\xb4\x78\x96\x22\x4a\x29\xde\xdf\xc8\x4d\xf5\xcc\xd4\xea\x60\xdd\xc5\xf1\x07\x8c\xa3\xea\x93\x7e\x5c\x31\x3c\x6a\xee\xa9\xc0\x3e\xc7\xda\x16\x68\x6c\x8e\xb9\xfc\x02\xc9\xa9\x9b\x18\x93\x28\x45\xab\xe6\x80\x12\xda\x5a\x32\xdb\xf4\x38\xc8\x31\x7d\x32\x29\x1c\xf9\xde\x19\xb4\x32\x4a\x19\x02\xc8\x6c\xc7\xa7\xf8\x0a\x00\x00\xff\xff\x5f\x4d\xfe\x90\xde\x01\x00\x00")

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

	info := bindataFileInfo{name: "template/delete_column.tmpl", size: 478, mode: os.FileMode(420), modTime: time.Unix(1474800935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateInsertTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xd1\x6e\xe2\x3a\x10\x7d\xb6\xbf\x62\x6e\x74\xaf\x64\xab\xa9\x3f\x80\x8a\x97\xdb\x56\x2a\x12\xcb\x96\x86\x7d\xaa\xfa\x10\xc8\x04\x45\x18\x87\xd8\x4e\x77\x2b\xd7\xff\xbe\x72\x1c\x68\x16\x68\x77\xb5\x7d\xc3\xf1\x78\xce\x99\x73\x66\x06\xe7\xa0\xc0\xb2\x52\x08\xc9\x44\x19\xd4\x36\x01\xef\xa9\x73\x97\xf0\xef\x2a\xdf\xa2\x9c\xe5\x5b\x84\xd1\x18\x44\xf7\xe3\x15\x6c\x7d\x1d\x3e\xc3\x2b\x7c\x53\xb6\xb2\x12\xe1\xf2\x10\x5f\x2b\x63\x75\xbb\xb2\xb5\x0e\x2f\x76\xba\x52\xb6\x84\x64\x86\xdf\xff\x33\xd9\x7c\x9a\x00\x3b\x4a\xc2\xbb\xb7\xf6\x65\x87\xe0\xdc\x10\xcf\xfb\xc8\x25\x9b\x4f\x21\xa6\x04\x47\xc9\x71\x4c\x36\x9f\x52\x62\xd0\x7e\xc9\x77\xc4\x34\x52\xe6\x22\xeb\x0e\x94\x5c\xd7\xb2\xdd\x2a\x03\x8f\x4f\xc6\xea\x4a\xad\xa9\xa7\xb4\x6c\xd5\x0a\x58\x73\x82\x94\xcd\xa7\x1c\x22\x1c\xe3\x1f\xf0\x70\x94\x68\xb4\xad\x56\xef\xc7\x38\x4a\xce\x91\x1c\x41\x93\x52\xd2\x33\x1d\xc1\x90\xaa\xf3\x29\x25\x3e\xd0\x73\x0e\x74\xae\xd6\x08\x62\x4f\xde\x7b\xe7\xc0\xe2\x76\x27\x73\x7b\x70\x27\x5e\x26\x20\xe2\x35\xaa\x22\xd8\xf5\x5e\x6d\x07\x66\x1c\x16\x75\xd6\x48\xc6\x81\x45\x45\x52\x78\x7c\xaa\x94\x45\x5d\xe6\x2b\x74\x3e\x05\xd4\xba\xd6\x3c\x54\xf9\x9c\xeb\x70\x8a\x5f\xe2\xd1\xc0\x20\x16\xc6\x01\x47\x64\x9d\x2f\x3d\x90\xf3\x94\x54\x25\xd8\x14\xea\x4d\xf0\xde\x08\x76\xcc\xe5\x06\xcb\xbc\x95\x36\x52\xba\xab\xeb\x0d\x6a\x7e\x15\xc2\x83\x6a\x4d\x47\x00\xc6\x60\xc5\x49\x1c\x6b\x38\x25\x21\x7b\x88\xf8\x67\x0c\xaa\x92\xdd\x9b\xbd\x1d\x49\x72\x54\x4c\x5f\x0e\x25\xc4\x07\x71\x49\x63\x52\x78\x36\x11\x61\x34\x86\x46\x44\x27\xc4\xa2\xee\xf5\x09\xc2\xd0\x33\x08\xbf\x05\xf0\x94\x92\xa6\x45\xfd\x12\xf2\x26\x93\x59\x76\xfb\xb0\x80\xc9\x6c\xf1\xb5\x53\xa8\x2f\x1c\x12\xb8\x80\xc6\xd0\x43\xff\xc4\x17\x17\x90\x5c\x25\x91\x98\xaa\x64\xdf\x02\x55\x09\xe2\x2e\x37\xf7\x9b\x6e\x32\xfe\xc0\xd6\xdb\x1f\xb8\x62\xc5\xb2\x6f\xaa\x9b\xff\x39\xb0\x13\x73\x06\xe6\x86\x39\x45\x69\xf0\x13\xe9\x8b\xa5\x78\x40\xd3\x4a\x7b\x9c\x36\xb6\x62\xd4\x23\x85\x5c\xaf\x87\x92\xf7\xed\xf7\x81\xca\x67\x7a\xea\xa0\x32\xd1\x6f\x88\x21\x5f\xb1\x14\x1d\xb3\x01\x96\x10\x82\x77\x2b\xe2\x57\x09\xff\x12\xaf\x2a\x0e\x58\x11\x5a\x4c\x73\xd3\xf7\xe4\xa4\xf8\x4c\x1d\x6f\x1b\x64\xb0\x2d\xbd\x67\x5c\x64\x28\x71\x65\x19\x17\xf7\x9b\x38\xe5\xac\x2a\xb8\xc8\x2a\xb5\x96\xc8\x8a\x65\x57\xdd\xc0\xbd\x7d\xae\x81\x34\x7d\x44\x34\xc2\xd3\xf3\x9b\xf5\xcc\x1c\xbe\x4d\x77\x28\xe4\x74\x00\x3f\x68\x90\xf7\xef\xf6\xed\x41\xc3\xbf\xc3\x9e\xd4\xcf\x00\x00\x00\xff\xff\x73\x47\xef\x69\x6e\x06\x00\x00")

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

	info := bindataFileInfo{name: "template/insert.tmpl", size: 1646, mode: os.FileMode(420), modTime: time.Unix(1474801356, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateInsert_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xce\xbb\x6a\xc4\x30\x10\x85\xe1\xda\x7a\x8a\x83\x48\x61\x37\x7a\x80\x80\x2b\x57\x81\x24\x10\xe2\xa4\x09\x29\xb4\xeb\x31\x18\x24\xf9\x26\x19\x96\xf1\xbc\xfb\x62\xf6\x56\xec\xc2\x96\x53\xfc\xdf\x1c\x66\x34\xd4\x76\x81\xa0\xdf\xc2\x4c\x53\xac\x7a\x97\x7c\xd0\x10\x61\xc6\xcb\xec\xad\x73\xb5\xdd\x39\xfa\xb4\x9e\xf0\x5a\xc2\xdc\xae\x15\xb1\xaf\xac\x27\x87\x15\x3f\x21\x76\xd1\x11\x44\x54\x9b\xc2\x1e\xf9\x88\x07\xbd\xc8\xe9\xc9\xf7\xd7\x7b\x81\x5f\xeb\x12\x31\xc3\xdc\x61\xf5\x99\xca\x97\x0d\x31\xf5\x61\xb8\xe4\xc5\x13\x15\xac\xb2\xd1\xcc\x14\x3f\xec\xf0\xa7\xaf\xb8\x88\xfe\x47\x89\x45\x65\x13\xc5\x34\x05\x8c\x4a\x94\x62\x06\x85\x66\x9b\x7c\x0c\x00\x00\xff\xff\x30\xe2\x3a\x61\x06\x01\x00\x00")

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

	info := bindataFileInfo{name: "template/insert_column.tmpl", size: 262, mode: os.FileMode(420), modTime: time.Unix(1474800935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSelectTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x55\xcf\x6f\xdb\x36\x14\x3e\x8b\x7f\xc5\x1b\xb1\x15\x52\xa3\x10\x19\x30\xec\xd0\xc1\x87\x2e\x4b\xd0\x0e\x6e\x56\xc7\x29\x76\x08\x72\x50\x25\xca\x23\x42\x93\x16\x49\x47\x2d\x58\xfd\xef\xc3\x23\x25\x5b\xa9\x12\xc7\x97\x5e\x0c\x8b\x7c\xef\xfb\xbe\xf7\x93\xde\x43\xc5\x6b\xa1\x38\xd0\x25\x97\xbc\x74\x14\xba\x8e\x78\x7f\x0a\x3f\x97\xc5\x9a\xcb\xab\x62\xcd\xe1\xcd\x0c\x58\xf8\xf3\x0d\x9c\x3e\xc7\x63\xf8\x06\x9f\x94\x13\x4e\x72\x38\xdd\xd9\x6b\x65\x9d\xd9\x96\x4e\x1b\xf4\xd8\x18\xa1\x5c\x0d\xf4\x8a\xb7\xbf\xd8\xe5\x62\x4e\x21\x9d\x82\xdc\x20\x44\x16\x30\x1e\x0a\x03\xde\x8f\x69\xbb\xee\xad\x94\xe7\x5a\x6e\xd7\xca\xc2\x0c\x6e\xef\xac\x33\x42\xad\x3c\x49\xbc\x07\x53\xa8\x15\x07\x36\x5c\x77\x1d\xf5\xbe\x57\xd9\x75\x34\xf7\x1e\xb8\xaa\x30\x96\x8e\x10\xf7\x75\xc3\x27\xd8\x31\xdc\xe5\x62\x0e\x51\x35\x44\xdc\xc7\x36\x8b\x39\x49\x06\x8a\x81\x9f\x24\xda\x54\xdc\x00\xc0\xf0\x2d\xc5\x5a\x38\x00\x78\xbd\x15\xca\xfd\xfe\x1b\x52\xd6\x5b\x55\x42\xda\x4c\x59\x17\xf3\x0c\x22\x75\x9a\x1d\xd0\xe4\x49\x62\xb8\xdb\x1a\xf5\xbc\x8d\x27\x49\xd2\xe4\x24\x99\xa8\xde\x67\x0d\x6f\x29\xc5\x5f\x25\x64\x4e\x92\xee\xa0\xb4\x01\x38\x83\x39\x06\x94\x4a\x88\xf1\xbc\xa0\xb3\x61\x31\xfe\x19\xbc\x92\x3b\xd5\x0d\x32\x3d\x55\x26\xef\xc1\xf1\xf5\x46\x16\x6e\xd7\x72\xf1\x92\x02\x8b\xd7\x7d\xdd\x8e\x90\x79\xa3\x97\x8d\x4c\x33\x48\x63\x21\x72\xb8\xbd\x13\xca\x71\x53\x17\x25\xf7\x5d\x0e\xdc\x18\x6d\x32\xd4\x58\xf6\xfc\x6f\x66\x7d\xd1\x2c\xfb\x5b\x0b\x95\x36\x83\xb2\x1c\x68\x0e\x34\x23\x49\xfb\x1f\x37\xdc\xe6\xf0\x60\x83\x3f\x7a\x34\x2c\x1c\xb2\x9e\x8e\x24\xa2\x0e\x57\x3f\xcd\x40\x09\x89\xf0\x43\xd4\x94\xe6\x78\x14\x3c\x31\xdb\x24\x69\xb6\xdc\x7c\x45\x10\xba\xbc\x98\x5f\x9c\xdf\x00\x85\x13\x18\xd4\x9c\x00\x85\xcb\xeb\x7f\x3e\xc0\xb8\x75\x03\x7c\x54\x81\x0c\x94\x06\x82\x88\x73\x32\x03\x0a\xff\xbe\xbb\xb8\xbe\x40\x98\x68\x84\x3c\xfb\xeb\x86\x85\xe6\x0c\x20\x43\x61\x46\x3a\x47\x30\xf3\xf7\x1f\xde\x47\x39\xd6\x99\x52\xab\x07\x76\xa9\xcd\xba\x70\x9f\x84\x72\xe9\xeb\xde\x37\x87\x5f\xcf\xb2\x18\xc9\x50\xd8\x88\x00\xf4\x0f\x1a\x93\xa4\x84\x8c\xb5\x3e\x05\x51\x03\x7b\x57\xd8\x8f\xf7\xfb\xfa\xd9\x10\xdb\x32\x4c\x58\x1f\xe1\xa8\xff\xd3\x03\xc5\xfd\x6e\x02\x46\xcb\xa5\xeb\xd2\x8c\x0d\x18\x0c\xf1\x3f\xde\xc7\x32\x3e\xb3\x60\xd0\xc3\x3e\x36\xbc\x14\x5c\x56\x83\x20\x82\x1b\x2c\xf4\x1d\x2e\xa2\x63\x3a\x6f\x29\xd4\x4a\xf2\xb4\xfa\x0c\xb6\x91\xb2\x60\x7f\xfd\x19\x83\x79\x1c\xe9\xb8\x03\x77\x9d\x06\xb3\x03\x6b\xae\xaf\x64\x0e\x85\x59\x8d\x1b\xf0\xe5\xd6\x9b\x90\xf7\x03\xd0\x17\x4f\xb7\x08\x54\x7d\x66\x0b\xc4\xbf\xd6\x6d\x3a\x22\x62\x8c\x65\xfb\x02\xb3\x65\x59\xa8\xd4\xe8\x36\x3b\x72\x5b\xbc\x95\xf2\xbb\x4c\xdc\xde\x1d\xcc\x85\x09\x83\xb8\x2e\xee\xf9\xd3\x96\x67\x7d\xdb\xfd\xe8\x9c\x3d\x1e\x55\xa3\xdb\xbd\xfb\x90\xa9\x69\x9a\x8e\x42\x4b\x2a\x5e\x73\x03\x08\xc9\xce\xa5\xb6\x1c\x75\xd4\xba\x3f\xb9\xe2\x5f\xc2\xf6\x47\xdf\x91\xe0\x21\xef\x36\x23\xc9\x13\x34\x53\x1e\x24\xc2\x64\xce\xa0\xd8\x6c\xb8\xaa\x52\x63\x73\x30\x61\x5e\x07\x63\xb3\x9f\xd0\x63\xda\x1a\x15\xd8\xbe\x92\xf8\xa1\xb8\x79\xa9\xb1\xf1\xdd\xc6\xf6\x9a\x18\x91\xa4\x8f\xcc\xc6\xc8\xe2\x4b\x35\x79\x10\x5e\x19\xdd\xe2\x6c\xa6\x42\x55\xfc\x4b\xdc\x83\x16\xce\x32\x04\x13\x6a\x85\x6c\xd1\xb3\x7f\x17\x92\x7d\xab\x1a\xdd\xc6\x4c\xec\xe6\xb7\xeb\xc8\xff\x01\x00\x00\xff\xff\x20\xea\x51\x1c\xd2\x08\x00\x00")

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

	info := bindataFileInfo{name: "template/select.tmpl", size: 2258, mode: os.FileMode(420), modTime: time.Unix(1474801322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSelect_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x93\x5f\x6f\xda\x30\x14\xc5\x9f\xe3\x4f\x71\x14\xed\x21\x51\xa9\xd5\x49\xd3\x1e\x90\xf2\xd0\x52\x1e\x90\xd6\xd1\x15\x36\x69\xaa\xfa\xe0\xc1\x45\x8b\x30\x8e\xe3\xfc\x59\x3b\xe3\xef\x3e\x39\x21\x2d\x61\x48\x65\x5a\x9f\x20\xbe\x3e\x3e\xbe\xbf\xe3\x6b\x2d\x96\xb4\x4a\x15\x21\x9c\x91\xa4\x45\x39\xca\x64\xb5\x51\x21\x9c\xb3\x16\xef\x8a\x8d\x90\x72\x2e\x7e\x48\xfa\x2c\x36\x84\x61\x02\xfe\xf2\xb5\x45\x99\x8d\xc4\x86\x24\xb6\xf8\xaa\xca\xb4\x94\x04\xe7\xd8\xaa\x52\x0b\x44\x39\x8e\xe8\x9d\x6b\x4d\x66\x5f\x3e\xc5\xbe\xce\xff\x3a\x67\xbe\x3b\x25\xaa\x9b\xfa\xfc\x49\x77\xca\x01\xe8\x51\x9b\x02\x9c\xf3\x22\x97\x52\xf0\xa9\x26\x23\xca\xcc\xc4\xaf\x38\xc1\xb2\xa0\x16\x06\x99\x46\x5f\xc8\x82\x74\x05\x49\x2a\x6a\x0e\x8e\x91\x24\xb8\xf0\x9b\x83\x4c\x23\x79\xde\x3b\xce\x2b\x21\x59\xe0\x40\xb2\xa0\x97\x72\xa3\xb9\xbf\x78\x60\x81\x63\x2c\xf8\xf5\x93\x4c\x83\xa7\x55\x8d\x1f\xb5\xe9\x5d\x7f\x0b\xbf\x94\xfe\xa6\xbd\x16\xed\x37\x21\x2b\x1a\xa2\x1e\x60\xaa\x87\xc8\xf4\x00\x2d\xfc\x21\xc2\x67\x36\xce\x85\x8e\x05\x39\x6f\x0d\x12\x08\xad\x49\x2d\xa3\xdd\xc2\x00\xcd\x4f\xcc\x02\x43\x65\x65\x14\x72\xe6\xd8\xdb\x04\x30\x51\xc7\x22\xa8\x1b\xfe\x07\xcb\xa7\x04\x70\x04\xd0\x4d\x25\xcb\xf4\x44\x4a\xc5\xb0\xeb\xfc\xfe\xe1\xc0\xdd\xd6\xcd\xbd\x38\xe7\x71\x0b\xb2\x75\xb8\x11\x6b\x9a\xa8\x2e\xea\xc8\xc7\x5c\x17\x31\xce\xf0\x3e\x7e\x2b\xce\xd6\x22\x5d\x81\x4f\x8a\xdb\x35\xce\xff\xe1\xe1\xdf\xae\x5b\xff\x48\xaf\x91\xaa\xf2\xe3\x87\xff\x7b\xda\x9e\xea\x01\x93\x48\xaf\xf7\xee\xca\x5f\x99\xb4\x9d\xbd\x27\xc8\x1c\xb3\xf6\x1c\xa4\x96\x7e\x92\x4f\xee\x68\x6a\x96\x64\xae\x9e\x8e\xf8\x38\x17\x65\xbe\xd8\x8d\x93\xff\x7f\x4a\x57\x39\x6f\x65\x09\x42\x4c\xef\xae\xc7\x77\xb8\xfa\x8e\xfd\xb4\x9a\xe1\xdd\xed\xe9\x5e\xd5\x65\xb1\x68\x06\xb4\x13\x9f\x79\xf5\xe5\x6c\x14\xf6\xc6\xb7\x57\xbd\x1e\xb7\x65\xd6\x8b\xd6\xda\x0e\xc1\x9f\x00\x00\x00\xff\xff\xea\x45\xa9\x39\x20\x05\x00\x00")

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

	info := bindataFileInfo{name: "template/select_column.tmpl", size: 1312, mode: os.FileMode(420), modTime: time.Unix(1474800935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateTableTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\xcd\x6a\xeb\x30\x10\x85\xd7\x9e\xa7\x18\xc4\x5d\x24\x8b\x38\xfb\x0b\x5d\x94\x76\x53\x08\xa1\x21\x0d\x5d\xab\xf2\xd4\x11\xd1\x8f\x2d\x8d\x1b\x8a\xa2\x77\x2f\x52\x42\x69\xa9\x77\x87\x39\x3e\x9f\x3f\x34\x48\x75\x92\x3d\x61\x4a\xd8\x3e\x5f\xf3\x56\x5a\xc2\x9c\x01\xb4\x1d\x7c\x60\x5c\x40\x23\x22\x07\xed\xfa\x28\xae\x51\x79\xf7\x21\x00\x9a\x94\x30\x48\xd7\x13\xb6\xf7\x5d\xa7\x59\x7b\x27\xcd\x0d\x12\x71\x95\x33\x34\xa2\x70\x31\x67\x51\x3f\x26\xd7\x61\xbd\xf6\x9a\x8f\xd3\x5b\xab\xbc\x5d\x5b\xa9\x4e\x44\xeb\xde\xaf\xe2\x68\x8c\x14\xb0\x84\x94\xf0\x9f\x92\x96\x4c\x35\xf9\x7f\x87\x6d\x0d\x17\x64\xff\x50\xce\x78\xc1\x83\x63\xcd\xa6\x6a\xf2\xe7\x50\xf5\x7f\x4c\x72\xde\xef\x36\x18\x39\x4c\x8a\x31\x41\x73\x3e\x52\x20\xac\xfc\xf6\xb5\x64\xc8\x00\xef\x93\x53\xb8\xa5\x73\x31\xfc\xc3\x7f\xb9\xd1\xf7\xbb\xcd\x62\x39\x4b\x4f\xd0\x8c\xc5\x6d\xa6\x4b\x19\x9a\x40\x3c\x05\x87\x63\xf9\x53\x4a\xc8\x64\x07\x23\x99\x50\xec\xc9\x90\x62\x51\x5f\xe5\x77\x73\x18\x3a\xc9\x34\xd7\x3c\xb9\x48\x61\x76\xf3\x48\x86\xbe\x37\x5f\x01\x00\x00\xff\xff\xa4\x2d\xbc\xde\xcc\x01\x00\x00")

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

	info := bindataFileInfo{name: "template/table.tmpl", size: 460, mode: os.FileMode(420), modTime: time.Unix(1474800645, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\xc1\x4e\xe3\x30\x10\x3d\xdb\x5f\x31\x6b\xb1\x52\x22\x82\x3f\xa0\xa8\x07\x16\xba\xe2\xd0\x45\x94\x14\x71\x40\x68\x15\x9a\x29\x1b\xd5\x75\x1a\xdb\xa5\x8b\x8c\xff\x7d\x65\x3b\xa5\x81\xb6\x80\xd0\xde\x1a\x7b\x3c\xef\xcd\x7b\x33\x53\x6b\xa1\xc4\x69\x25\x11\xd8\xf5\xa2\x2c\x0c\x32\x70\xce\x5a\x38\x98\x14\x73\x14\x17\xc5\x1c\xa1\xd7\x07\x1e\x7e\x3c\x83\xa9\x4f\xfd\x31\x3c\xc3\xb5\x34\x95\x11\x08\xce\x51\x6b\x8f\xe0\x60\x52\x4b\x6d\xd4\x72\x62\x6a\xe5\x1f\x2c\x54\x25\xcd\x14\xd8\x05\xae\xbe\xeb\x7c\x34\x64\x90\x6c\xe7\x18\xfb\x0c\x29\x1c\x39\x47\xcd\xd3\x02\xe1\x35\xac\x73\x91\x50\x3e\x1a\x42\x4c\x0d\x96\x92\xb7\x31\xf9\x68\x48\x89\x46\xf3\xab\x58\x10\xdd\x08\x51\xf0\x3c\x7c\x50\x72\x5a\x8b\xe5\x5c\x6a\xb8\xbd\xd3\x46\x55\xf2\x81\x3a\x4a\xa7\x4b\x39\x81\xa4\xd9\x42\xca\x47\xc3\x14\x22\x5c\x92\xbe\xc3\xc3\x52\xa2\xd0\x2c\x95\xdc\x1f\x63\x29\xd9\x45\xb2\x07\x4d\x46\x49\xcb\xb4\x07\x5d\xaa\xd6\x65\x94\x38\x4f\xcf\x5a\x50\x85\x7c\x40\xe0\x6b\xf2\xc1\x0b\x83\xf3\x85\x28\xcc\x8b\x45\xf1\x92\x01\x8f\xd7\x28\x4b\x6f\xc3\xbe\xda\x5e\x98\xa5\x30\xae\xf3\x46\x24\x29\x24\x51\x91\x0c\x6e\xef\x2a\x69\x50\x4d\x8b\x09\x5a\x97\x01\x2a\x55\xab\xd4\x57\xf9\x58\x28\xff\x15\x4f\xe2\xa7\x86\x4e\x2c\xf4\x3d\x0e\xcf\x83\x2f\x2d\x90\x75\x94\x54\x53\x30\x19\xd4\x33\xdf\x03\x9a\x27\x6f\xb9\x9c\xe1\xb4\x58\x0a\x13\x29\x9d\xd7\xf5\x0c\x55\x7a\xec\xc3\xbd\x6a\x4d\x20\x00\x7d\x30\x7c\x2b\x2e\x69\x52\x4a\x7c\x76\x1f\xf1\xad\x0f\xb2\x12\xe1\xcd\xda\x0e\xc6\xde\x14\xd3\x96\x43\x09\x71\x5e\x5c\x2f\x7c\xab\x69\x06\xfa\x51\x47\xa8\x5e\x1f\x1a\x1e\x2d\xe1\xe3\xba\x15\xca\x2b\x44\x77\x40\x7d\x88\xe4\x28\x59\xfd\x41\x85\x3a\x83\xd5\x2b\x84\x70\xca\x5b\xf1\xbf\x98\x9a\x92\x66\x89\xea\xc9\x27\x64\xd7\x97\x67\x27\xe3\x41\x30\xa0\xd5\x15\xf2\xc1\x98\xc1\x21\x6c\xaa\x0c\x30\x91\x8e\x47\x62\x2c\x4a\x1c\x72\x1c\xf6\x81\xc1\xcd\xf9\xe0\x6a\xe0\xdf\xc4\xa0\x88\xd1\x12\x69\xc3\x80\x1d\xb3\x0c\x8a\xc5\x02\x65\x99\x04\xd1\x56\x8f\x9a\x73\x9e\x66\x9e\x79\x6c\xd8\x23\xa8\xa6\xc0\xcf\x0b\x7d\x39\xdb\x34\xa1\xde\x6e\x8e\xaf\x8c\x57\x67\xa5\x38\x97\xa4\x7c\x9d\x81\xdf\x78\xca\x1e\xe2\x72\x16\xcb\xdd\xb3\x5b\xfc\x33\xcd\x5f\x05\xfe\xac\x50\x94\x6b\x4e\xef\xad\x84\xce\xd8\x0c\xfe\xe2\x24\x29\xef\xdb\xa1\x3d\xfb\x91\x42\x72\x7b\xb7\x55\x61\x77\x7c\x82\x80\x19\x14\xea\xa1\xdb\x07\x1f\x77\x80\xac\xc4\xa6\x99\x7e\xaf\xe7\xa1\xbc\xe7\x81\x42\x27\xab\x77\xe1\xb3\x79\x9a\x26\xc2\xef\x5a\x9d\x1b\xcb\x1b\x9e\xa3\xc0\x89\x49\x52\x7e\x22\x44\x52\xde\x7b\x75\xbc\xbf\x28\x34\x7e\x72\xbf\xec\x10\x4a\x37\x82\x5f\xa1\x5e\x0a\xf3\xdf\xf5\x69\xcf\xf6\xa9\xd3\xb2\x8f\xcb\x71\xf7\x1f\xcc\x8e\x75\xb4\x59\x72\x1e\x77\x7b\x0f\xbd\x53\xfd\xfe\xbb\x75\xe9\x81\xd3\x9a\xd2\xbf\x00\x00\x00\xff\xff\xeb\x93\xdb\xc2\x7a\x07\x00\x00")

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

	info := bindataFileInfo{name: "template/update.tmpl", size: 1914, mode: os.FileMode(420), modTime: time.Unix(1474801336, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUpdate_columnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\x41\x6b\xe3\x30\x10\x85\xcf\xd6\xaf\x78\x98\x3d\xd8\x60\x44\xce\x06\x9f\x42\x6e\xbb\x1b\x4a\x92\xf6\x10\x72\x50\xeb\x09\x35\xc8\xb2\x2c\xcb\x6e\x5a\x45\xff\xbd\x28\x4e\x93\x9a\x16\x9a\x9e\x84\x46\xf3\xbe\x99\xd1\x1b\xe7\x50\xd2\xbe\x52\x84\x78\xa3\x4b\x61\x69\xde\xc8\xbe\x56\x31\xbc\x77\x0e\x7f\xba\x5a\x48\xb9\x16\x8f\x92\xfe\x8b\x9a\x90\x17\xe0\xd7\xdb\x11\xb6\x99\x8b\x9a\x24\x8e\xd8\x28\x5b\x59\x49\xf0\x9e\xed\x7b\xf5\x84\xa4\xc5\x37\x7a\xef\xc7\x22\xab\xbb\xbf\x29\x56\x64\x9d\x03\xff\x82\x5a\x9f\x41\xc9\x10\x10\x7c\xfd\xaa\x3f\xc4\xe9\x0f\x4c\x38\x16\xb5\xbc\x23\xfb\x4f\xe8\x6d\x7c\x81\x7b\x1f\xef\x50\x60\x60\x91\x21\xdb\x1b\x85\x96\x79\x76\x73\x9b\x0f\xcf\x64\xe8\x97\x8d\x66\xa0\x83\x36\x1d\x38\xe7\x5d\x2b\xa5\xe0\x4b\x4d\x46\xd8\xc6\xdc\x32\xc1\x20\x0c\x1a\x8d\xa9\x90\x45\xd5\x1e\x92\x54\x72\x02\xa7\x28\x0a\xcc\x42\x72\xd4\x68\x14\x97\xdc\x45\xdb\x0b\xc9\x22\x0f\x92\x1d\x5d\x9f\x4f\x9a\xed\x6c\xc7\x22\xcf\x58\xf4\x12\x06\x0a\x56\x8e\xaa\xc5\x41\x9b\x49\xfb\x47\x84\x50\xf5\x46\x9f\x46\x74\xf7\x42\xf6\x94\x63\xc8\xb0\xd4\x39\x1a\x9d\x61\x5c\x94\x1c\x93\x7f\xf6\xc1\x80\xb1\x40\x01\xa1\x35\xa9\x32\x39\x07\x32\x9c\x8e\x74\xea\x82\x73\x20\x55\x86\xb5\x79\x0f\x00\x00\xff\xff\xf5\x51\xfc\xe7\x8a\x02\x00\x00")

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

	info := bindataFileInfo{name: "template/update_column.tmpl", size: 650, mode: os.FileMode(420), modTime: time.Unix(1474800935, 0)}
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

