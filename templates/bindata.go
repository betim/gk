// Code generated by go-bindata.
// sources:
// tmpl/file.tmpl
// tmpl/gk.json.tmpl
// tmpl/partials/constants.tmpl
// tmpl/partials/endpoint_func.tmpl
// tmpl/partials/func.tmpl
// tmpl/partials/func_parameters.tmpl
// tmpl/partials/func_results.tmpl
// tmpl/partials/func_return.tmpl
// tmpl/partials/imports.tmpl
// tmpl/partials/interface.tmpl
// tmpl/partials/interface_func.tmpl
// tmpl/partials/interface_stub.tmpl
// tmpl/partials/struct.tmpl
// tmpl/partials/struct_function.tmpl
// tmpl/partials/vars.tmpl
// tmpl/proto.pb.tmpl
// tmpl/proto_compile.bat.tmpl
// tmpl/proto_compile.sh.tmpl
// DO NOT EDIT!

package template

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _tmplFileTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xb1\x6a\x84\x40\x10\x86\x7b\x9f\x62\x58\x2c\x12\x08\x92\x3a\x90\x2a\xd5\x15\x09\x81\x40\xda\x30\xec\x8d\x46\xa2\xa3\x59\xe7\x6c\x06\xdf\x3d\xec\xed\x2a\x9b\x45\xe1\x3a\x99\xff\xff\xbe\x19\x77\x44\xfb\x83\x0d\x81\x6a\xf5\x1e\x3e\x97\x45\xb5\xad\xa1\x11\xb8\xeb\x88\xa1\x3a\xf5\xe3\xe0\x64\xba\x87\xc7\x65\x29\x54\x85\xfa\xb1\x43\x21\x30\x6d\x08\xcc\x56\xf1\x24\xf1\x39\x17\xbc\x0c\x3c\x09\xf2\x9e\xc2\xae\x91\x49\x6a\x07\x9a\x4f\x74\x3b\x86\x19\x9d\x87\x7d\x78\xc0\x9d\x58\xc8\xd5\x68\x69\xa3\x1d\x72\x43\x50\xb6\x0f\xe5\x0c\x4f\xcf\x69\xc3\xb3\xc9\x0f\xae\x73\x03\xe5\x7c\x25\xa3\x7e\x67\xcb\x87\xb8\x8b\x95\xc3\x15\x31\xfe\xef\x9f\xae\xc3\x1b\xe4\xaf\x24\xdf\xc3\x39\x93\x27\xf6\x98\x47\x8c\x7e\xa1\x9c\xe3\xc6\xea\x0d\x7b\x02\x63\xb2\x57\xab\x2f\x6c\xc3\x5e\x55\xea\x26\xca\xe2\x70\xd8\x97\x6f\x49\x3b\xf0\xd6\xf4\x97\x65\x87\x16\xc5\x5f\x00\x00\x00\xff\xff\x18\xaa\x49\x60\x42\x02\x00\x00"

func tmplFileTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplFileTmpl,
		"tmpl/file.tmpl",
	)
}

func tmplFileTmpl() (*asset, error) {
	bytes, err := tmplFileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/file.tmpl", size: 578, mode: os.FileMode(420), modTime: time.Unix(1491773578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGkJsonTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\x3d\x6e\xc3\x30\x0c\x85\x77\x9f\x42\xd0\x5c\xf8\x00\x5e\x03\x74\xec\xe2\x74\x4e\x58\x8b\x8e\x85\xc8\x12\x41\xd1\x2d\x0a\x41\x77\x2f\x5c\xff\xd5\xa8\x87\x2c\x19\xa5\x47\x7e\xef\x03\x53\xa1\x94\x8e\xc8\x9f\xb6\x41\x5d\x8d\x2f\xa5\x34\x81\x74\xba\xd2\x29\x5d\x53\x92\x50\x7b\xb8\xe3\x09\x22\xaa\xb2\x9e\x06\xdf\xa0\xc7\x9c\xaf\x39\xa7\xd4\x5a\x87\x35\x12\x30\x48\xe0\x9c\xe9\x7e\xfb\xf7\xb7\xd0\x5f\x26\xf8\x98\x5e\x3c\xf4\xa8\xab\xa5\xb8\xbc\x85\x25\xb5\x5e\x90\x5b\x68\xd6\x91\x59\xe2\x9d\x08\xf9\xd5\x72\x94\x13\xf4\xe8\x8e\x75\xea\x7d\x53\x14\x1e\x1a\x59\xbb\x64\xf8\x98\x61\x0f\x20\x0a\xa5\xf2\x88\xd1\xbd\x35\xc6\xe1\x17\xf0\x76\x9e\x19\xb8\x25\xa3\xff\xba\x80\xde\x50\xb0\x5e\xe2\x93\xce\xb9\xf1\x0f\x0e\xba\x86\x3b\x25\x61\xf0\x91\x02\xcb\x93\x94\x7e\x31\xe5\x79\x69\x39\x7f\xd3\xb4\x7c\x64\xd8\x81\x37\x0e\x79\xe7\x67\xb0\x85\xc1\xc9\xe5\x8f\xa7\xee\x44\x48\x17\xf9\x27\x00\x00\xff\xff\x20\x06\xc3\x4e\xa0\x02\x00\x00"

func tmplGkJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplGkJsonTmpl,
		"tmpl/gk.json.tmpl",
	)
}

func tmplGkJsonTmpl() (*asset, error) {
	bytes, err := tmplGkJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/gk.json.tmpl", size: 672, mode: os.FileMode(420), modTime: time.Unix(1491829652, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsConstantsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x48\x2f\x51\xd0\xc8\x49\xcd\x53\x50\xd1\x54\x30\xa8\xad\x4d\xce\xcf\x2b\x2e\x51\xd0\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xe5\xaa\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x20\xec\xcc\x34\x05\x95\x32\x3d\x8f\xc4\xe2\xb0\xc4\x9c\xd2\xd4\xda\x5a\x5b\x88\x3c\x94\x57\x5d\x9d\x9a\x97\x02\xa7\xb8\x34\xa1\x0c\x40\x00\x00\x00\xff\xff\x13\xc6\xee\xf4\x7e\x00\x00\x00"

func tmplPartialsConstantsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsConstantsTmpl,
		"tmpl/partials/constants.tmpl",
	)
}

func tmplPartialsConstantsTmpl() (*asset, error) {
	bytes, err := tmplPartialsConstantsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/constants.tmpl", size: 126, mode: os.FileMode(420), modTime: time.Unix(1490947670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsEndpoint_funcTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x50\xb1\x6a\x2b\x31\x10\xec\xdf\x57\x6c\x71\xc5\x1d\x08\x7d\x80\xe1\x55\x86\x94\x29\x0c\x49\x2f\x2e\x7b\x46\x70\x59\x9d\x57\x7b\x87\x61\xd9\x7f\x0f\x92\xad\xe0\xf8\x8a\x14\x51\xb3\x5a\x98\x99\x9d\x99\x7f\x00\x00\x8c\xb2\x32\xc1\xb4\xd2\xd8\x8f\x72\x85\x31\x91\xe0\x55\xfc\xf1\x36\x1d\x30\x5e\x56\xcc\x02\x91\x04\x79\x0a\x23\xaa\x0d\xd0\x3f\x6c\x0e\x90\x39\xf1\x00\x5a\x05\xdb\x63\xbc\xc0\xe1\x7f\xa3\xfb\x5e\xd5\x9f\xee\xff\xd7\xf0\x89\x66\xc3\x0f\xb8\x2a\x07\x3a\x23\x74\xd1\x75\x5b\x21\xfa\x63\x98\xe7\x48\x67\x7f\xc2\xbc\xce\x92\xcd\x54\xbb\xed\xce\x55\x8d\x13\x50\x12\xe8\xe7\x90\x05\xba\x08\xdd\x33\x7e\x30\x73\xaa\x48\x1f\x05\x5d\x47\x51\xcd\xdb\xe8\x55\xbf\xb1\x37\xb5\x12\xdc\xed\x0d\x34\xbb\xef\x81\xb3\x59\x49\xe2\x7f\xb1\xf0\xc8\xd8\xdf\x1f\x9e\xfa\xa9\xc5\xd7\x5a\xf2\x92\x28\x63\x13\xde\x77\xb1\x0b\x57\x34\x25\xbd\x2d\x0b\xf2\x4b\xe4\x72\xbe\xd9\x3a\xfc\xb9\x25\x73\x40\x71\xae\x5e\xed\x2b\x00\x00\xff\xff\x23\x6f\x0b\xb8\x23\x02\x00\x00"

func tmplPartialsEndpoint_funcTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsEndpoint_funcTmpl,
		"tmpl/partials/endpoint_func.tmpl",
	)
}

func tmplPartialsEndpoint_funcTmpl() (*asset, error) {
	bytes, err := tmplPartialsEndpoint_funcTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/endpoint_func.tmpl", size: 547, mode: os.FileMode(420), modTime: time.Unix(1491900788, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFuncTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x31\x4e\xc5\x30\x10\x44\x7b\x9f\x62\xe4\xea\xff\x26\x9f\x33\xd0\xd1\x72\x81\xaf\x90\x4c\x82\x25\x67\x1d\xec\x75\x11\xad\xf6\xee\x28\x01\x21\xd1\xbd\xdd\x19\x8d\xde\xd2\x65\x82\x99\x72\xdb\xf3\xa8\x44\x4c\xa2\xac\xcb\x38\xf1\x79\x46\x11\x83\x3b\x0c\x66\x69\x81\x14\xc5\x8d\x5f\x18\x5e\xcb\x7c\x20\xc6\xbb\xbb\xd9\x75\x9c\xc0\xdc\xe8\x1e\x1e\x0f\xbc\x6d\x7b\xe6\x46\x51\x1c\xa5\x57\x7c\xf4\x96\x84\xad\x21\x97\x35\x4d\xf8\x64\x65\xb8\xf6\x56\xc5\x2d\x53\x30\xbc\xb3\xf5\xac\xed\x8e\x17\xf7\x4a\xed\x55\xfe\x29\x9d\x22\xcf\x9f\x7f\xfc\x2b\xbb\x07\x33\xca\xec\x0e\xfc\x42\xf0\xef\x00\x00\x00\xff\xff\x19\xbb\xd8\xc9\xce\x00\x00\x00"

func tmplPartialsFuncTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFuncTmpl,
		"tmpl/partials/func.tmpl",
	)
}

func tmplPartialsFuncTmpl() (*asset, error) {
	bytes, err := tmplPartialsFuncTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func.tmpl", size: 206, mode: os.FileMode(420), modTime: time.Unix(1491774296, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFunc_parametersTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xad\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x52\x41\xc2\x99\x69\x0a\x79\xf9\x25\x0a\x1a\x39\x89\xc5\x25\x0a\x2a\x99\x0a\x2a\x9a\xb5\xb5\x3a\xd5\xd5\xa9\x79\x29\x20\x59\x30\x05\x08\x00\x00\xff\xff\x15\x5b\x78\xe8\x50\x00\x00\x00"

func tmplPartialsFunc_parametersTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFunc_parametersTmpl,
		"tmpl/partials/func_parameters.tmpl",
	)
}

func tmplPartialsFunc_parametersTmpl() (*asset, error) {
	bytes, err := tmplPartialsFunc_parametersTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func_parameters.tmpl", size: 80, mode: os.FileMode(420), modTime: time.Unix(1490904745, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFunc_resultsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x48\x2f\x51\xd0\xc8\x49\xcd\x53\x50\xd1\x54\x30\xa8\xad\xd5\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xad\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x52\x41\xc2\x99\x69\x0a\x79\xf9\x20\xfd\x89\xc5\x25\x0a\x2a\x99\x0a\x2a\x9a\xb5\xb5\x3a\xd5\xd5\xa9\x79\x29\x20\x59\x30\xa5\x09\xa5\x01\x01\x00\x00\xff\xff\xcc\x6e\x26\x22\x6c\x00\x00\x00"

func tmplPartialsFunc_resultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFunc_resultsTmpl,
		"tmpl/partials/func_results.tmpl",
	)
}

func tmplPartialsFunc_resultsTmpl() (*asset, error) {
	bytes, err := tmplPartialsFunc_resultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func_results.tmpl", size: 108, mode: os.FileMode(420), modTime: time.Unix(1490905608, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFunc_returnTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xad\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\x05\x31\x33\xd3\x14\xf2\xf2\x4b\x14\x34\x72\x12\x8b\x4b\x14\x54\x32\x15\x54\x34\x6b\x6b\x75\xaa\xab\x53\xf3\x52\x40\xb2\x60\x0a\x10\x00\x00\xff\xff\xbc\x66\xba\xd5\x44\x00\x00\x00"

func tmplPartialsFunc_returnTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFunc_returnTmpl,
		"tmpl/partials/func_return.tmpl",
	)
}

func tmplPartialsFunc_returnTmpl() (*asset, error) {
	bytes, err := tmplPartialsFunc_returnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func_return.tmpl", size: 68, mode: os.FileMode(420), modTime: time.Unix(1490905876, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsImportsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x50\xa9\xad\xcd\xcc\x2d\xc8\x2f\x2a\xd1\xe0\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x05\xc9\x81\xd5\xe4\xe5\x97\x28\x68\xa4\x16\x2a\xa8\x94\xe9\xf9\x25\xe6\xa6\x2a\x28\x29\x69\x82\x64\xa0\xdc\xda\x5a\x85\xea\xea\xd4\xbc\x14\xa8\x50\x48\x65\x41\x6a\x6d\x2d\x17\x54\x88\x4b\x13\xca\x00\x04\x00\x00\xff\xff\x34\xce\x15\xc7\x6d\x00\x00\x00"

func tmplPartialsImportsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsImportsTmpl,
		"tmpl/partials/imports.tmpl",
	)
}

func tmplPartialsImportsTmpl() (*asset, error) {
	bytes, err := tmplPartialsImportsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/imports.tmpl", size: 109, mode: os.FileMode(420), modTime: time.Unix(1491090111, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsInterfaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\xc1\x0a\x82\x40\x14\x85\xe1\xbd\x4f\x71\x10\x17\x05\xa1\xb4\x0d\x7a\x84\xda\xb6\x8c\x41\xcf\xd4\xc0\xcc\x28\xd3\x55\x90\xcb\x7d\xf7\x88\x42\xda\xff\x1f\xbf\xac\x13\xa1\xda\x5e\x5d\xa2\x19\x42\x16\x16\xef\x7a\x42\x2b\xd5\xe2\xf2\x83\x68\xc2\xa1\x59\x70\x3a\xa3\xbd\x50\x9e\xe3\xf0\x32\x53\x15\xa6\x29\x3a\x21\xea\x8d\xdc\xfd\x9c\xfb\x1a\xcd\x62\x56\xa9\x32\x0f\x9f\x2e\x78\x44\xc1\x2e\x32\x6f\x7c\x8f\xa3\x59\xd7\xe1\x56\x82\x10\xeb\x38\x97\xbf\x6d\xfa\x36\x3f\x5f\xd9\x3b\x00\x00\xff\xff\x7d\xe1\x7d\xe3\xa0\x00\x00\x00"

func tmplPartialsInterfaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsInterfaceTmpl,
		"tmpl/partials/interface.tmpl",
	)
}

func tmplPartialsInterfaceTmpl() (*asset, error) {
	bytes, err := tmplPartialsInterfaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/interface.tmpl", size: 160, mode: os.FileMode(420), modTime: time.Unix(1491592614, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsInterface_funcTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xd5\xa8\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x2b\xcd\x4b\x8e\x2f\x48\x2c\x4a\xcc\x4d\x2d\x49\x2d\x2a\x56\x52\xd0\x0b\x80\x73\x6a\x6b\x35\x31\x94\x16\xa5\x16\x97\xe6\x94\x80\xd4\x05\x41\x58\xb5\xb5\x80\x00\x00\x00\xff\xff\xc1\x2e\xd4\xb5\x59\x00\x00\x00"

func tmplPartialsInterface_funcTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsInterface_funcTmpl,
		"tmpl/partials/interface_func.tmpl",
	)
}

func tmplPartialsInterface_funcTmpl() (*asset, error) {
	bytes, err := tmplPartialsInterface_funcTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/interface_func.tmpl", size: 89, mode: os.FileMode(420), modTime: time.Unix(1490906014, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsInterface_stubTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x2a\x2e\x29\x2a\x4d\x2e\x51\x52\xd0\xab\xad\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\xd0\xf3\xd4\xf3\x4d\x2d\xc9\xc8\x4f\x29\xae\xad\x45\xd5\x54\x9a\x14\x9f\x56\x9a\x97\x5c\x92\x99\x9f\xa7\xa4\xa0\x52\x56\x5b\xcb\xc5\x55\x5d\x9d\x9a\x97\x52\x5b\x0b\x08\x00\x00\xff\xff\x39\x10\x5c\x96\x5d\x00\x00\x00"

func tmplPartialsInterface_stubTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsInterface_stubTmpl,
		"tmpl/partials/interface_stub.tmpl",
	)
}

func tmplPartialsInterface_stubTmpl() (*asset, error) {
	bytes, err := tmplPartialsInterface_stubTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/interface_stub.tmpl", size: 93, mode: os.FileMode(420), modTime: time.Unix(1490908658, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsStructTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\x28\x2e\x29\x2a\x4d\x2e\x51\xa8\x56\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\xd0\x0b\x4b\x2c\x2a\xae\xad\xe5\xaa\xae\x56\x29\x83\xa9\x06\xb3\x43\x2a\x0b\x52\x6b\x6b\xab\xab\x53\xf3\x52\x6a\x6b\x15\x6a\x01\x01\x00\x00\xff\xff\x65\x47\x74\x51\x51\x00\x00\x00"

func tmplPartialsStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsStructTmpl,
		"tmpl/partials/struct.tmpl",
	)
}

func tmplPartialsStructTmpl() (*asset, error) {
	bytes, err := tmplPartialsStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/struct.tmpl", size: 81, mode: os.FileMode(420), modTime: time.Unix(1491062995, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsStruct_functionTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x31\x4e\xc4\x40\x0c\x45\xfb\x9c\xe2\x6b\xaa\xa4\x99\xe5\x0c\x74\x34\x14\x40\xbf\x0a\xd9\x9f\x25\xd2\xc4\x09\x33\x9e\x22\xb2\x7c\x77\x34\x41\x02\x6d\x65\xfb\xc9\xf6\x7f\x73\x95\x09\xbd\x59\x7c\xd7\x5c\x27\x8d\xaf\xe3\x4a\x77\xfc\x83\x8f\x63\xa7\xfb\x00\x33\xe5\xba\xa7\x51\x89\xb0\x88\x32\xcf\xe3\xc4\x6b\x3b\x0f\x88\xee\x06\xb3\x65\x86\x6c\x8a\x9e\xdf\x88\xcf\xdb\xed\x40\x08\x83\xbb\xd9\x39\xb4\x86\xa9\xd0\xbd\xbb\x5c\xf0\xb2\xee\x89\x2b\x45\x71\x6c\x35\xe3\xb3\x96\x45\x58\x0a\xd2\x76\x5f\x26\x7c\x31\xb3\x3b\xff\xdd\x15\x7d\xa2\x20\xbe\xb1\xd4\xa4\x65\xc0\x93\x7b\xa6\xd6\x2c\x0f\x46\xcd\xe3\xfa\xcb\xc3\xdf\xb2\x7b\x67\x46\xb9\x9d\xd1\xad\xc0\x7f\x02\x00\x00\xff\xff\x00\x38\xee\x95\xef\x00\x00\x00"

func tmplPartialsStruct_functionTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsStruct_functionTmpl,
		"tmpl/partials/struct_function.tmpl",
	)
}

func tmplPartialsStruct_functionTmpl() (*asset, error) {
	bytes, err := tmplPartialsStruct_functionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/struct_function.tmpl", size: 239, mode: os.FileMode(420), modTime: time.Unix(1491400930, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsVarsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x48\x2f\x51\xd0\xc8\x49\xcd\x53\x50\xd1\x54\x30\xa8\xad\x2d\x4b\x2c\x52\xd0\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xe5\xaa\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x20\xec\xcc\x34\x05\x95\x32\x3d\x8f\xc4\xe2\xb0\xc4\x9c\xd2\xd4\xda\x5a\x5b\x88\x3c\x94\x57\x5d\x9d\x9a\x97\x02\xa7\xb8\x34\xa1\x0c\x40\x00\x00\x00\xff\xff\xd8\x54\xdd\xc2\x7c\x00\x00\x00"

func tmplPartialsVarsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsVarsTmpl,
		"tmpl/partials/vars.tmpl",
	)
}

func tmplPartialsVarsTmpl() (*asset, error) {
	bytes, err := tmplPartialsVarsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/vars.tmpl", size: 124, mode: os.FileMode(420), modTime: time.Unix(1490946107, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplProtoPbTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xc1\x0a\x82\x40\x10\x86\xef\xfb\x14\x83\x78\x28\x08\x3d\x74\x4b\x7c\x84\x3a\x44\x2f\xb0\xe9\xa4\x4b\xb9\xbb\xed\xac\x92\x0c\xf3\xee\xa1\x64\x51\x87\x8e\xf3\x7f\x1f\x1f\x43\xa3\x8d\xfa\x01\x25\x24\x3e\xb8\xe8\xb6\x49\xa1\xbc\xae\xae\xba\x41\xf0\xe7\x42\xe5\x39\x9c\x5a\x04\xe6\xec\xa0\x3b\x14\x01\xc2\x30\x98\x0a\xa1\xc6\x8b\xb1\x26\x1a\x67\x33\xb5\x6c\x1f\x8b\x15\x73\xd0\xb6\x41\x48\xcd\x26\x1d\x60\x57\x42\xb6\xc7\xd8\xba\x9a\x44\x20\xf8\x0a\x98\xd3\x61\xb1\x57\xf3\x71\xc4\x7b\x8f\x14\x45\xd6\x10\x30\xf6\xc1\xd2\x1b\xf8\xdb\x38\xcd\x2c\x8a\x19\x6d\x2d\x22\xff\xfa\x1d\x12\x4d\xff\x7f\x57\x81\x95\xa8\x1f\x34\x77\x67\xf0\xea\x3e\x03\x00\x00\xff\xff\x30\xc0\xb8\x3f\x0e\x01\x00\x00"

func tmplProtoPbTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplProtoPbTmpl,
		"tmpl/proto.pb.tmpl",
	)
}

func tmplProtoPbTmpl() (*asset, error) {
	bytes, err := tmplProtoPbTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/proto.pb.tmpl", size: 270, mode: os.FileMode(420), modTime: time.Unix(1492288367, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplProto_compileBatTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\xb1\x4a\x04\x31\x10\xc6\xf1\x7e\x9f\x62\x1e\xc0\x24\x85\x5d\xe0\x6a\xb1\xb1\x11\x6b\x99\xdd\x1b\xe7\x02\xd9\x4c\xc8\x4c\x44\x58\xf6\xdd\x25\x39\x41\x04\x9b\x10\x86\x3f\x3f\xbe\x18\xe1\xb9\xa8\x61\xce\x50\x9b\x98\x3c\xfa\x25\x46\xb8\x99\x55\x8d\x21\x70\xb2\x5b\x5f\xfd\x26\x7b\x60\x11\xce\x14\x66\xb4\xf6\x8f\xd0\x28\x13\x2a\xe9\xc8\xdf\xea\x15\x8d\xee\xc0\x06\x4f\x02\x6b\x2a\xd7\x54\x58\xe1\x33\xe1\x08\x80\x05\x98\x0c\x5c\x87\x3f\x64\xc6\xc2\xbf\xe4\x31\x7f\x0f\x77\xc6\x31\x15\xc7\x72\x2e\x31\x0e\xe1\x95\x08\x30\xab\x4c\xed\xbf\x79\xad\x6e\xf3\x71\x2c\xc1\x1a\x51\xd8\x51\x8d\x5a\xa0\x2f\xdc\x6b\x26\x5d\x96\x9f\x79\xc7\xe1\x5f\x70\xa7\xf3\xf4\xf3\x00\xce\xb1\xbc\x4b\xb7\x4b\xcd\x9d\x53\xd1\xcb\x50\xa2\xff\x0e\x00\x00\xff\xff\x01\x54\xb0\xf4\x19\x01\x00\x00"

func tmplProto_compileBatTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplProto_compileBatTmpl,
		"tmpl/proto_compile.bat.tmpl",
	)
}

func tmplProto_compileBatTmpl() (*asset, error) {
	bytes, err := tmplProto_compileBatTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/proto_compile.bat.tmpl", size: 281, mode: os.FileMode(420), modTime: time.Unix(1492288812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplProto_compileShTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xbd\x6a\x2b\x31\x14\x84\x7b\x3d\xc5\x5c\xb6\xbd\x2b\x15\xe9\x6c\x5c\x87\x34\x49\x61\x52\x07\xad\x7c\x7c\x56\x44\xd2\x59\xf4\xe3\x24\x18\xbf\x7b\x58\x79\x09\x04\xd2\x48\x83\x18\xbe\x6f\x34\xfc\x33\xad\x64\x33\xf9\x64\x28\x5d\x50\x66\xa5\x06\x3c\xa5\x52\x6d\x08\x58\xb2\x54\x79\xc0\x39\x4b\x44\x91\x96\x1d\x21\x5a\xf7\x72\x84\xa4\xf0\xa5\xd5\x00\x4c\x99\x3e\xe0\xb7\xba\x6d\x55\x9c\xa4\x73\x0f\xd1\xbe\x13\x82\x9f\xaa\x48\x58\x9b\xec\x2b\x5c\x90\x44\x98\x6b\x5d\xca\xce\x18\xf6\x75\x6e\x93\x76\x12\x0d\x8b\x70\x20\xd3\x7d\x53\x3b\xaf\x7d\x6d\x56\x0a\x53\xd2\x65\xc6\x1e\xda\xac\x64\xcf\x2d\x13\xf6\xe8\xf0\xed\xda\xec\x6a\x50\x03\x5e\x97\x93\xad\x74\xdf\xed\xf0\x28\x98\x7c\x3a\xf9\xc4\x05\x17\x6f\xfb\x0a\x01\x53\xc5\xd8\xf0\xcb\x1e\x6c\xe2\x1f\xbb\xb9\xf6\xf4\xff\x4e\x19\x99\xd2\xc8\x72\xeb\xfc\x23\x11\x6c\x28\xb2\xa2\xfe\xfa\x46\x5e\x5c\x3f\x46\x16\x53\x33\x91\x89\xb6\x54\xca\x86\x3e\x6d\x5c\x02\x15\xa5\xb6\x69\xd7\xab\x7e\xb6\x91\x6e\x37\xdd\x1f\x30\x8e\x2c\x6f\xd2\xea\x61\x09\x8d\x7d\x2a\x87\x95\xb2\xd3\xdf\x01\x00\x00\xff\xff\xf5\x6a\x2e\xbe\x9e\x01\x00\x00"

func tmplProto_compileShTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplProto_compileShTmpl,
		"tmpl/proto_compile.sh.tmpl",
	)
}

func tmplProto_compileShTmpl() (*asset, error) {
	bytes, err := tmplProto_compileShTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/proto_compile.sh.tmpl", size: 414, mode: os.FileMode(420), modTime: time.Unix(1492288693, 0)}
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
	"tmpl/file.tmpl":                     tmplFileTmpl,
	"tmpl/gk.json.tmpl":                  tmplGkJsonTmpl,
	"tmpl/partials/constants.tmpl":       tmplPartialsConstantsTmpl,
	"tmpl/partials/endpoint_func.tmpl":   tmplPartialsEndpoint_funcTmpl,
	"tmpl/partials/func.tmpl":            tmplPartialsFuncTmpl,
	"tmpl/partials/func_parameters.tmpl": tmplPartialsFunc_parametersTmpl,
	"tmpl/partials/func_results.tmpl":    tmplPartialsFunc_resultsTmpl,
	"tmpl/partials/func_return.tmpl":     tmplPartialsFunc_returnTmpl,
	"tmpl/partials/imports.tmpl":         tmplPartialsImportsTmpl,
	"tmpl/partials/interface.tmpl":       tmplPartialsInterfaceTmpl,
	"tmpl/partials/interface_func.tmpl":  tmplPartialsInterface_funcTmpl,
	"tmpl/partials/interface_stub.tmpl":  tmplPartialsInterface_stubTmpl,
	"tmpl/partials/struct.tmpl":          tmplPartialsStructTmpl,
	"tmpl/partials/struct_function.tmpl": tmplPartialsStruct_functionTmpl,
	"tmpl/partials/vars.tmpl":            tmplPartialsVarsTmpl,
	"tmpl/proto.pb.tmpl":                 tmplProtoPbTmpl,
	"tmpl/proto_compile.bat.tmpl":        tmplProto_compileBatTmpl,
	"tmpl/proto_compile.sh.tmpl":         tmplProto_compileShTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"file.tmpl":    &bintree{tmplFileTmpl, map[string]*bintree{}},
		"gk.json.tmpl": &bintree{tmplGkJsonTmpl, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"constants.tmpl":       &bintree{tmplPartialsConstantsTmpl, map[string]*bintree{}},
			"endpoint_func.tmpl":   &bintree{tmplPartialsEndpoint_funcTmpl, map[string]*bintree{}},
			"func.tmpl":            &bintree{tmplPartialsFuncTmpl, map[string]*bintree{}},
			"func_parameters.tmpl": &bintree{tmplPartialsFunc_parametersTmpl, map[string]*bintree{}},
			"func_results.tmpl":    &bintree{tmplPartialsFunc_resultsTmpl, map[string]*bintree{}},
			"func_return.tmpl":     &bintree{tmplPartialsFunc_returnTmpl, map[string]*bintree{}},
			"imports.tmpl":         &bintree{tmplPartialsImportsTmpl, map[string]*bintree{}},
			"interface.tmpl":       &bintree{tmplPartialsInterfaceTmpl, map[string]*bintree{}},
			"interface_func.tmpl":  &bintree{tmplPartialsInterface_funcTmpl, map[string]*bintree{}},
			"interface_stub.tmpl":  &bintree{tmplPartialsInterface_stubTmpl, map[string]*bintree{}},
			"struct.tmpl":          &bintree{tmplPartialsStructTmpl, map[string]*bintree{}},
			"struct_function.tmpl": &bintree{tmplPartialsStruct_functionTmpl, map[string]*bintree{}},
			"vars.tmpl":            &bintree{tmplPartialsVarsTmpl, map[string]*bintree{}},
		}},
		"proto.pb.tmpl":          &bintree{tmplProtoPbTmpl, map[string]*bintree{}},
		"proto_compile.bat.tmpl": &bintree{tmplProto_compileBatTmpl, map[string]*bintree{}},
		"proto_compile.sh.tmpl":  &bintree{tmplProto_compileShTmpl, map[string]*bintree{}},
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
