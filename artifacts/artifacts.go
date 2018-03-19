// Code generated by go-bindata.
// sources:
// artifacts/backup-crd.yaml
// artifacts/backup-cronjob.yaml
// artifacts/backup-pvc.yaml
// artifacts/mysql-configmap.yaml
// artifacts/mysql-crd.yaml
// artifacts/mysql-service-read.yaml
// artifacts/mysql-service.yaml
// artifacts/mysql-statefulset.yaml
// DO NOT EDIT!

package artifacts

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

var _artifactsBackupCrdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x4e\x04\x31\x0c\x44\xfb\x7c\x85\xbf\x20\x68\x3b\x94\x12\xe8\x38\x90\x00\x89\xde\x9b\x35\x8b\xb5\x49\x1c\x62\x67\xc5\xfd\x3d\xda\x9c\x28\x8e\x86\x72\x3c\x1e\xbd\xb1\xb1\xf2\x3b\x35\x65\x29\x01\xb0\x32\x7d\x1b\x95\x43\xa9\xdf\x6e\xd5\xb3\xdc\xec\xd3\x4c\x86\x93\xdb\xb8\x2c\x01\xee\xbb\x9a\xe4\x57\x52\xe9\x2d\xd2\x03\x7d\x70\x61\x63\x29\x2e\x93\xe1\x82\x86\xc1\x01\x14\xcc\x14\x20\x9f\xf5\x2b\xcd\x18\xb7\x5e\xd5\xc7\xe6\x87\x96\x4a\x0d\x4d\x9a\x5f\x9b\x25\xbf\xb2\x7d\xf6\xd9\x47\xc9\x4e\x2b\xc5\x23\xbb\x36\xe9\x35\xc0\xbf\xfb\x17\x8a\x1e\x11\x80\x4b\xb7\xa7\xf3\xdb\xcb\xe9\x6e\x00\xc7\x34\xb1\xda\xe3\x5f\xe7\xc4\x6a\xc3\x55\x2e\x6b\x4f\xd8\xae\x8a\x0e\xa7\xa6\xde\x30\x5d\x1f\xe0\x00\x34\x4a\xa5\x00\xcf\x07\xb6\x62\xa4\xc5\x01\xec\xbf\xaf\xdb\x27\xf7\x13\x00\x00\xff\xff\xe5\xf5\xf0\x5c\x4a\x01\x00\x00")

func artifactsBackupCrdYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupCrdYaml,
		"artifacts/backup-crd.yaml",
	)
}

func artifactsBackupCrdYaml() (*asset, error) {
	bytes, err := artifactsBackupCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsBackupCronjobYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x41\x6f\xdb\x3a\x0c\xbe\xe7\x57\x10\x79\x2f\xe8\xe1\x41\x76\x82\x1e\x1e\x60\xa0\x87\x2d\xdd\xb0\x0d\x4b\xd7\xad\x5d\x81\xdd\x4a\xcb\x6c\xa2\xd6\x92\x5c\x8a\xce\x12\xa4\xfd\xef\x83\xe2\x74\x8b\xd5\x06\xe8\x78\xfc\xf8\x91\xfa\xf8\x51\xc4\xc6\x5c\x11\x07\xe3\x5d\x01\x25\x8a\x5e\xe4\xcb\x49\x49\x82\x93\xc1\x9d\x71\x55\x01\x53\xf6\xee\x93\x2f\x07\x96\x04\x2b\x14\x2c\x06\x00\x0e\x2d\x15\xb0\xd9\x40\x76\x86\x96\xe0\xf1\x51\xdd\xfa\x72\x00\xe0\x7f\x3a\xe2\x6f\x74\x43\x4c\x4e\x53\x88\x54\x00\x05\x5d\xa3\xd9\xfa\xe2\xeb\xe7\xb7\xa8\xef\xda\x66\x8b\x03\xec\x3f\xad\x39\xb3\xeb\x70\x5f\xfb\x86\x18\xc5\x73\x36\x67\xa9\xb3\xb9\x91\x45\x5b\x66\xda\xdb\x7c\x39\xd9\x55\x3d\x7b\x7c\x87\xb7\xa6\xea\xe0\xef\x1f\x4f\x23\x1a\x1a\xd2\x51\x41\xd0\x0b\xaa\xda\x9a\x0a\x18\xc6\xec\x45\x43\x3a\xbb\x34\xdb\xca\xe1\x00\xe0\xd6\x97\x97\x64\x9b\x1a\x85\x3a\xbd\x4f\x75\x31\xa4\x97\x49\xb3\x31\x98\x82\x20\xcb\xb9\xaf\x8d\x5e\x17\xf0\xc5\xbd\x47\x53\xb7\x4c\x7b\x94\xa5\xaf\x5b\xfb\xe4\x46\x17\x6a\x37\x45\xb9\x6f\x47\x17\x4d\x74\x24\x08\x39\xb9\xda\xd6\x4d\x6b\x34\xb6\xe8\x51\x00\x74\x04\xcf\x0e\x18\xb1\xff\x80\xf6\xee\xa6\x57\x4b\xb6\x91\xf5\xa9\xe1\x02\x36\x87\xf8\x66\xae\x2c\xf6\x45\x75\xf0\x0c\x9b\x54\x48\x57\xb5\xdd\xdc\xa0\x47\x17\x34\x8e\xf8\xb5\x43\x1b\x8b\x73\x2a\x20\xee\x3c\x5f\x09\x63\x47\x29\xa2\xf5\x41\x12\x25\xd6\xa2\xab\xfa\x3a\x14\x94\x18\x16\x09\x34\x54\x7a\x98\x40\x0f\x89\xfa\x40\x02\x8a\x56\x83\x04\xfe\x07\xa6\xb5\x77\x04\xf1\xbb\xc3\x0d\x7b\x0b\x16\x83\x10\x67\x09\xef\xf4\xcd\xe5\xbb\x93\xeb\x0a\x85\xe0\xe8\xbf\xd1\x0f\x35\xb2\x6a\x54\xa9\xd1\x07\x35\x9a\x1d\x5d\x27\x5c\x7b\x57\x19\x86\x7c\xeb\x54\xde\x8d\x97\xef\xad\x2e\xff\x77\x13\xbb\x3d\xa6\xf6\x6a\x14\x50\x8a\x49\x2f\x95\x77\xf5\x1a\x7e\xff\xe0\x69\xdd\x46\x49\xf1\xf6\xc6\xd9\x0b\x28\x1c\x1f\x8f\xff\x87\x07\x58\x95\x41\x98\xd0\x82\x5a\x81\x9a\xbe\x42\xc0\x33\x2f\xce\x99\x1a\x64\x02\x59\xd0\x6e\x75\xa9\x0f\x7f\x36\x06\x4a\x35\x3b\xba\x52\x82\x3c\x27\x51\x95\xe1\x93\xbf\x1a\xbb\xbb\x97\x99\x6f\x9d\x84\x74\xcf\x07\x7f\x10\x80\x8d\x05\xe7\x28\x8b\xe2\xf0\x94\x2f\x76\x7b\x76\x23\xfd\x5e\x24\x7a\xd7\x2f\x12\xb3\x2a\xa1\x32\x05\xdf\xb2\xa6\x90\x9e\x46\x4c\xdd\xb7\x14\xd2\x21\xba\xd0\x4d\x5b\xc0\x64\x3c\xb6\x2f\xe4\x2c\x59\xcf\xeb\x6d\x7a\x66\x7e\x05\x00\x00\xff\xff\x28\x08\xc5\x35\xa0\x05\x00\x00")

func artifactsBackupCronjobYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupCronjobYaml,
		"artifacts/backup-cronjob.yaml",
	)
}

func artifactsBackupCronjobYaml() (*asset, error) {
	bytes, err := artifactsBackupCronjobYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-cronjob.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsBackupPvcYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x31\x8a\xc3\x30\x14\x44\x7b\x9d\x62\x2e\xb0\x86\x6d\xd5\x6e\xed\x65\x59\x83\x53\x7f\xa4\x21\x88\x58\x92\xa3\xff\x1d\x08\xc6\x77\x0f\x22\x4e\xf9\xe6\x0d\x33\xb7\x54\xa2\xc7\x1f\x9b\x26\x35\x16\x9b\xeb\xb2\x65\xfe\x2c\x92\xb2\x93\x35\xcd\x5d\xd4\xe2\xf1\xf8\x76\x99\x26\x51\x4c\xbc\x03\x8a\x64\x7a\xec\x3b\x86\x5f\xc9\xc4\x71\x38\x5d\x19\xba\x91\x10\xa8\x3a\xd6\x48\xed\x08\x7c\xe1\x9f\x12\x2f\x2d\x19\x47\x29\x4f\x07\x34\x6a\xdd\x5a\xf8\x14\x1a\xef\x1b\xd5\x4e\x02\xd4\x6a\x93\xeb\x39\x3f\xad\x0c\xc3\xf4\x4e\xfa\xcd\x2b\x00\x00\xff\xff\x02\x00\x34\xe8\xb0\x00\x00\x00")

func artifactsBackupPvcYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupPvcYaml,
		"artifacts/backup-pvc.yaml",
	)
}

func artifactsBackupPvcYaml() (*asset, error) {
	bytes, err := artifactsBackupPvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-pvc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xb1\x0a\xc2\x30\x10\xc6\xf1\x3d\x4f\xf1\x81\x73\x05\xd7\x6c\xe2\xec\xea\x22\x0e\xd7\xe6\xda\x06\x93\x4b\xcc\xc5\x42\xc1\x87\x97\x06\xba\x39\x38\x26\xf7\xbb\xff\x51\xf6\x37\x2e\xea\x93\x58\x2c\x27\xf3\xf4\xe2\x2c\x2e\x49\x46\x3f\x5d\x29\x9b\xc8\x95\x1c\x55\xb2\x06\x10\x8a\x6c\x11\x57\x7d\x05\x03\x04\xea\x39\xe8\xf6\x0f\x50\xce\xfb\x60\xc7\x91\xb4\x72\x39\x0e\x32\x5a\x7c\x1a\x3a\xe0\x9c\x73\x58\x51\x67\xaf\x18\xda\x05\x24\x09\x2b\x92\xa0\xce\xbc\x6f\x34\x7b\x6f\x31\xf7\x68\x8f\x90\xa6\xae\xf7\x62\x00\x0d\xb4\xf0\xbf\xcd\x86\xf5\x47\x4f\xdf\x99\x4b\x57\x98\x5c\xb7\x59\xf3\x0d\x00\x00\xff\xff\x1c\x35\x5e\xba\x03\x01\x00\x00")

func artifactsMysqlConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlConfigmapYaml,
		"artifacts/mysql-configmap.yaml",
	)
}

func artifactsMysqlConfigmapYaml() (*asset, error) {
	bytes, err := artifactsMysqlConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlCrdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x4e\x04\x31\x0c\x44\xfb\x7c\x85\xbf\x20\xe8\x3a\x94\xf6\xe8\x38\x90\x00\x89\xde\x97\x35\x8b\x75\x49\x1c\x6c\x67\xc5\xfd\x3d\xda\x5d\x51\x00\x05\xa5\x67\x3c\x7a\x63\x63\xe7\x57\x52\x63\x69\x09\xb0\x33\x7d\x3a\xb5\x75\xb2\x78\xb9\xb5\xc8\x72\xb3\x1c\xce\xe4\x78\x08\x17\x6e\x53\x82\xe3\x30\x97\xfa\x4c\x26\x43\x33\xdd\xd1\x1b\x37\x76\x96\x16\x2a\x39\x4e\xe8\x98\x02\x40\xc3\x4a\x09\xea\xd5\x3e\x4a\x2e\xc3\x9c\xd4\x62\xd6\xb8\x09\xd2\x49\xd1\x45\xe3\xac\x5e\xe2\xcc\xfe\x3e\xce\x31\x4b\x0d\xd6\x29\xaf\xe1\x59\x65\xf4\x04\xff\xee\xef\x18\x5b\x23\x00\x7b\xb9\x87\xeb\xcb\xd3\xe9\xb8\x13\x37\xb9\xb0\xf9\xfd\x1f\xeb\xc4\xe6\x9b\x6d\xdc\xe6\x51\x50\x7f\x76\xdd\xac\x5e\x86\x62\xf9\x75\x44\x00\xb0\x2c\x9d\x12\x3c\xae\xe8\x8e\x99\xa6\x00\xb0\x7c\xff\x6f\x39\x84\xaf\x00\x00\x00\xff\xff\x68\x53\xd4\x69\x4f\x01\x00\x00")

func artifactsMysqlCrdYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlCrdYaml,
		"artifacts/mysql-crd.yaml",
	)
}

func artifactsMysqlCrdYaml() (*asset, error) {
	bytes, err := artifactsMysqlCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlServiceReadYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcd\x4a\x03\x31\x14\x85\xf7\x79\x8a\xfb\x02\x13\x2d\x05\x17\xd9\xea\x46\x50\x41\x8b\xee\x6f\x33\xc7\x1a\xcc\x9f\x37\x99\x91\x52\xfa\xee\x92\x99\x80\x88\x74\x7b\x7e\xee\xf9\x2e\x67\xf7\x06\x29\x2e\x45\x43\xf3\x46\x7d\xba\x38\x1a\xda\x41\x66\x67\xa1\x02\x2a\x8f\x5c\xd9\x28\xa2\xc8\x01\x86\x4e\x27\xd2\x4f\x1c\x40\xe7\xf3\x20\xe0\x51\x11\x79\xde\xc3\x97\x16\x21\xe2\x9c\xff\x64\x14\x51\xfa\x8e\x90\x17\xbc\x43\x10\x2d\x7a\x6e\xa0\x75\xe8\xf1\xb8\x7b\x7e\xb8\xf5\x53\xa9\x90\xc5\x68\x27\x7e\x81\xac\xe8\x70\x2c\x5f\x3e\x65\x08\xd7\x24\xfa\x20\xd5\xeb\x83\xab\x1f\xd3\x5e\xdb\x14\xae\xe6\x4d\x6f\xfd\xa3\xeb\xfa\xe4\xc6\x55\x7e\xbd\xbf\x6b\x6a\xc9\xb0\x0d\x21\x27\xa9\x0b\xcb\xd0\xab\xcb\xce\x52\x6a\x96\xa1\xed\xf6\xfa\x46\x11\x15\x78\xd8\x9a\xe4\xc2\x7b\x3f\x01\x00\x00\xff\xff\xbd\x08\x38\x7a\x3e\x01\x00\x00")

func artifactsMysqlServiceReadYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlServiceReadYaml,
		"artifacts/mysql-service-read.yaml",
	)
}

func artifactsMysqlServiceReadYaml() (*asset, error) {
	bytes, err := artifactsMysqlServiceReadYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-service-read.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcf\x4a\x03\x31\x1c\x84\xef\x79\x8a\x79\x81\xae\x96\x82\x87\x5c\xf5\x52\xd0\xa2\x16\xbd\xa7\xd9\xb1\x06\xf3\xcf\x5f\xb2\x2b\xa5\xf4\xdd\x25\xbb\x0b\x82\x68\x8e\xf3\x65\x32\x5f\x4c\x76\xaf\x94\xe2\x52\xd4\x18\xd7\xea\xc3\xc5\x5e\x63\x4f\x19\x9d\xa5\x0a\xac\xa6\x37\xd5\x68\x05\x44\x13\xa8\x71\x3e\xa3\xdb\x99\x40\x5c\x2e\x0a\xf0\xe6\x40\x5f\x1a\x05\x4c\xce\xbf\x71\xfa\x8a\x94\x67\xbe\x51\x18\x2d\x97\x7b\x2b\xcc\x1b\x0f\xa7\xfd\xd3\xfd\xad\x1f\x4a\xa5\x4c\xa0\x3d\xf1\xe3\x62\xa5\x0b\xa7\xf2\xe9\x53\xa6\x98\x9a\xa4\x3b\x4a\xf5\xdd\xd1\xd5\xf7\xe1\xd0\xd9\x14\xae\xc6\xf5\xd2\xfa\x4b\xac\x9d\xc1\xf5\x73\xfc\xb2\xbd\x6b\x69\xc9\xb4\x4d\x21\x27\xa9\x93\xcb\x6a\xa9\x4e\x3b\x53\xa9\x21\x8d\xcd\xe6\xfa\x46\x01\x76\x76\xdb\x3e\x6a\xec\x52\xa4\x02\x0a\x3d\x6d\x4d\xf2\xcf\x87\xbf\x03\x00\x00\xff\xff\xe1\xa3\x4b\x10\x4b\x01\x00\x00")

func artifactsMysqlServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlServiceYaml,
		"artifacts/mysql-service.yaml",
	)
}

func artifactsMysqlServiceYaml() (*asset, error) {
	bytes, err := artifactsMysqlServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlStatefulsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x6d\x73\x1a\x39\x12\xfe\xee\x5f\xd1\x45\x5c\x6b\x7c\xb1\x30\x24\x95\x4b\x1d\x09\xb9\xf2\x12\xf2\x52\x67\x1b\x2f\xb0\xc9\x6d\x39\x1c\x2b\x34\x0d\xa8\x3c\x23\x8d\x25\x0d\x36\xe7\xf8\x7e\xfb\x95\x46\xc3\xa0\x19\x06\xc7\x49\xd5\x5e\xd5\x1e\x1f\xec\x41\x2f\xdd\x4f\x3f\xea\x6e\x75\x0f\x34\xe6\x9f\x50\x69\x2e\x45\x1b\x68\x1c\xeb\xe3\x65\x6b\xef\x8a\x8b\xa0\x0d\x43\x43\x0d\xce\x92\x70\x88\x66\x2f\x42\x43\x03\x6a\x68\x7b\x0f\x40\xd0\x08\xdb\x70\x77\x07\x8d\x73\x1a\x21\xdc\xdf\xef\x01\xc8\x1b\x81\x6a\x80\x33\x54\x28\x18\x6a\xbb\x0c\x80\x80\x13\x74\xb6\x1a\xfe\x72\xda\x0d\x13\x6d\x50\xa5\x13\x00\xbe\x56\xa6\x1a\xd1\x4a\x5f\x87\x32\x46\x45\x8d\x54\x8d\xb9\x32\x61\x63\xce\xcd\x22\x99\x36\x98\x8c\x2c\x22\xb7\xab\x4a\xb3\xfd\x24\x3c\x70\xc3\xbf\x7e\x7c\x6b\x47\x75\x8c\xcc\x42\xd0\x18\x22\x33\x52\x39\x38\x11\x35\x6c\x71\x4a\xa7\x18\x66\xf8\x2c\x8c\xb8\x2c\x4f\xa3\x5a\x72\x86\xe7\x15\xaa\x14\xc6\x21\x67\x54\xbb\xf1\x61\x8c\xac\x31\xc8\x86\xdc\x02\x83\x51\x1c\x52\x83\x99\x3e\x8f\x33\xfb\x09\x0b\xaa\x2b\x95\x03\xac\xa1\xdb\x0f\x17\xdc\x74\xa5\x30\x94\x0b\x54\xf9\x46\x92\xd1\x60\x67\x49\x4a\x5c\x2e\x91\x47\x74\x8e\x6d\x48\x07\xdb\x16\x88\x36\xf9\x1c\x93\x51\x44\x45\xb0\x51\x4f\x60\x4a\xf5\xc2\xfb\x5a\x23\xac\xe6\x7d\xfd\x9a\x3f\x5b\x52\x0c\x10\xbc\xf5\x46\x9e\xc0\x7b\x14\xf6\xbc\xd0\xa9\x4b\x79\x43\x45\x78\x00\x33\x25\x23\x88\x65\x00\x52\x05\x5c\xd0\x10\xb8\x08\xf0\xb6\xe1\x6d\xbe\xbc\x84\xdf\x17\x52\x1b\x6b\xc8\xef\xd0\xf9\x0f\x90\xfa\x65\x93\xfc\x6d\xfc\xf4\x70\x1f\xc6\x63\xf8\xfa\x15\xf0\x96\x1b\x68\x79\x5b\x32\x59\x9d\xfd\xbb\x9f\x4f\x86\x1f\x26\x83\xde\xd9\xc9\xa8\xfb\xe1\xb2\x35\xbe\xf7\x16\x21\x5b\x48\xb8\x4c\xf1\x04\x63\x78\x03\xc7\x91\x30\xc7\x4c\x8a\x59\x23\x38\xce\xf1\x35\x98\x98\x15\x0c\x39\x09\x02\xa0\x02\xe4\x6c\x66\xcd\x34\x12\xe8\x52\xf2\x00\x14\xa6\x5b\x82\x8d\x65\x9d\x26\x2c\x69\x98\x60\xa3\xac\x72\xb3\x62\xbf\x5e\x6f\x35\x9b\xf0\x14\xf6\x33\xc0\x87\x87\xf0\xe6\xb1\x40\xba\x32\x5e\x59\xaf\x50\x32\x56\xdc\x32\xeb\x76\xc0\x8c\x87\xa8\x1d\xad\x76\x84\xcf\x49\x44\x63\x0b\x14\xa3\xd8\xac\xde\x72\xe5\x03\xe2\x33\x4b\xef\x5a\x3d\x10\xbc\x86\x26\x8c\xc7\xaf\xc0\x2c\x50\x78\xeb\x00\x58\xbc\xc1\xe5\x64\x1e\x47\xd4\xc6\xa8\x05\x56\x80\xec\xdb\x1b\x6a\xfc\x96\x14\x1d\xd2\x25\x3e\x24\x64\xc6\xf3\x2f\x4b\x19\x26\x11\x9e\xc9\x44\x18\xed\xfb\xa6\xf3\x71\xbb\xd5\xdb\x17\xd9\x65\x17\xd4\x2c\xda\xbe\xe4\xca\x5d\x0e\xca\x37\xf6\x16\x16\xe5\xbb\x43\x29\xb0\x3a\xb0\x6c\x62\x3a\xbe\x35\x8a\x4e\x29\xbb\x4a\xe2\x3f\x3c\xc4\x86\x57\x3c\xb6\xe7\xe6\x30\xd9\xa3\xb5\xf9\x04\x68\xa8\x90\x06\x2b\x1b\x24\xda\xe8\x52\x60\x91\x00\x8e\x97\x54\x1d\x87\x7c\x7a\x9c\x1a\xe1\xfe\xda\xb8\xfa\xe9\x27\x17\x57\xcd\xff\x41\x28\x3e\x81\x01\x6a\x23\x15\x3a\xc8\xa9\xf3\x3a\xd6\x40\x0a\xc8\xfc\xec\x87\xdc\xf6\xee\xce\x2e\x16\x98\x65\xe0\x77\x4a\x46\x3f\xa7\x82\x1b\xee\x5f\x9a\x4d\x6b\xb5\xcd\xed\xf0\x63\x80\x00\x36\x07\x0d\x84\xc4\x0a\x63\xaa\x10\x08\x31\x54\xcd\xd1\x90\x80\xab\x4e\x46\xb0\x5b\x74\xbc\x7b\x33\x93\xf1\x8a\xd8\x6f\x8f\xde\x7e\x77\x07\x28\x82\x6d\x1b\x4a\x1e\x91\x23\x87\x7a\x21\xd9\x42\xf3\xb0\x68\xcb\xd6\xb9\x7b\x21\x98\xa6\x9e\x54\xdc\x86\x99\x58\xe1\x92\xcb\x44\x43\x8c\x45\x5a\x04\xa3\x06\x08\x51\xc8\x96\x44\x8a\x70\xe5\x5f\x60\x64\xbf\x5e\x5f\x9f\x20\x69\x1d\x1e\x36\xbc\x39\x78\xfe\xbc\xf9\x12\xbe\xc2\xed\x54\x1b\x85\x34\x02\x72\x0b\xa4\x5b\x72\xd5\x02\xa4\x8b\x8c\x71\x6b\xad\xa3\xc8\xc7\xf1\x88\xc3\xa9\x16\xfd\x70\xc2\xb1\x0c\xec\x48\x1a\xbb\x90\xea\x64\xea\x56\x14\x27\x1e\x91\xc2\xd0\xb0\xcc\x05\x4a\x89\xec\x07\x5c\x7c\xad\x4e\x39\x27\xdf\x95\xf8\x3c\x87\xab\x32\x21\xaf\x6b\x3c\x95\x1f\x85\x36\x54\x30\xf4\xb5\x95\xdd\x93\xed\xac\x52\x1e\x5f\xa0\xa0\x58\x6e\x9f\xc7\xd9\x6f\xc3\x5f\x4e\x27\x83\x7e\x7f\x34\xb9\x38\x19\x0e\x3f\xf7\x07\x6f\x3d\xdc\xe9\x6d\x6c\xa1\xb6\x0b\xce\xae\x91\x29\x34\xff\xc0\xd5\x00\x67\xc5\x19\xbf\x84\x4c\xed\xbc\xa0\x5a\xdf\x48\xb5\x15\x69\x00\x57\xb8\x6a\x43\x9c\x4d\xe7\x73\xb1\x54\x55\x6e\x53\xf6\x89\x9c\x8f\x0b\xa9\x4c\xdb\xfa\xfe\x5f\xff\x54\x0e\xa8\x50\xcb\x44\xe5\x85\xfc\x7a\xf0\x3a\x41\x6d\x74\xbb\x74\xff\x27\x6d\x78\xd1\x6c\x46\x85\xd1\x08\x23\xa9\x56\x6d\x68\xbd\xdf\xe4\x99\x90\x2f\x51\xa0\xd6\x17\x4a\x4e\xd1\x17\x82\xb7\x9b\x9a\x77\xcd\x9f\xbb\x49\xe1\xb2\x66\x6f\xd0\xda\x51\x7a\x73\x1e\x41\x2d\x45\x4b\x83\x88\x0b\x88\xb9\x98\x03\x89\xbf\xd4\xf6\xef\x2a\xbc\xe4\xfe\x4b\xad\x36\xf6\xef\x17\xc1\x0d\xa7\xe1\x5b\x0c\xe9\x6a\x88\x4c\x8a\x40\xb7\xe1\xb9\x9f\x11\x63\x54\x5c\x06\xf9\x5c\xcb\x9f\x33\x3c\x42\x99\x98\x7c\xf2\x85\xc7\x14\x0d\xf8\x23\xad\x7a\x02\xdd\x05\xb2\x2b\xb8\x41\x60\x54\xa4\x0b\x12\x83\x70\x9d\xa0\xe2\xa8\x41\x2e\x51\xc1\xa8\x7b\x01\x75\x7d\xc5\x63\x22\xd0\xdc\x48\x75\x65\xad\xe4\xda\x56\xa7\xa5\x94\xfe\x30\x47\x40\x16\xd0\x7a\xf6\xb2\xd1\x6c\x34\x1b\xad\x07\x69\x02\x82\x70\x30\xec\x9d\xf6\xba\x23\x68\x1d\x7c\x9b\xb4\x17\xbb\x39\x7b\xf6\x00\x65\xad\x52\x62\xd8\x24\xf0\xef\xac\xb2\xfe\xcc\x79\xa2\xc2\xe8\xaa\x64\xf1\xf2\x8f\x2a\x29\x59\xb9\x38\xdc\x2b\x5c\xb9\x6f\xd1\xa0\x8a\xb8\x40\x98\x72\x11\xca\x39\xc4\x52\x73\xc3\xa5\x6d\x8e\x5c\xc5\x11\xa4\x09\xea\xc8\x5e\x4f\x54\xac\xb6\x0b\x38\x32\xf3\x6c\x9c\xa4\x8d\xc0\x84\x8b\x99\xac\x2e\xe5\x9e\xc0\x3f\x8d\xa2\xee\x9e\xc9\xab\xda\x79\xd6\x57\x06\x40\x21\xa6\xca\x3a\x20\xd4\xba\x1f\x4e\xce\xdf\xf7\xe0\xec\x64\x38\xea\x0d\x60\xd4\xaf\xa5\x51\xb3\x2a\x49\x9b\x22\xa3\x89\x46\xb8\xc1\x03\xe5\x4a\x24\x1b\x3d\x69\x41\x93\xc6\x1b\xd7\xc6\x0e\xb8\xfe\xa4\x98\xaf\x96\x3b\x70\xb3\x05\x15\x73\x9c\xb8\x2a\x6b\x62\x64\x43\x5f\x87\x0d\x5e\x36\xe3\xe3\x5c\xd8\xca\xd2\x13\xe1\x08\x74\x32\xb8\x00\xb3\xe0\x1a\x18\xd5\x08\x75\x6e\x0e\x34\x24\x1a\x43\xd4\xba\x14\xd2\x2a\x2a\x11\xe8\x49\x29\xf4\x60\x55\x64\xfb\x1a\x77\xb0\xfd\xb9\xc0\x4b\xc0\x15\x32\x13\xae\x1c\x41\x59\x05\x0c\x17\x54\xe9\xad\xf3\x2f\xc2\xb4\x3d\x83\x2d\x03\xab\x95\xa7\x5d\xc4\xbf\xea\x8d\xbf\xfc\xfd\xf0\xf2\xb2\xad\x63\xca\xb0\x3d\x1e\x3f\x4d\x07\x76\xf6\x14\xa9\xed\xdf\x34\x3c\x6b\xb7\xb7\xbc\x21\x7b\x9a\x9c\xf6\xdf\x4f\xde\x7d\x3c\xed\x75\x0e\xb6\x7b\x93\x83\xa3\x2f\xa5\xc8\xb5\x1f\x6f\xe3\x45\x7f\x58\x6e\x69\x9e\x8d\xef\x6b\xf0\xe6\x11\x3e\x30\xe3\xc5\x40\x72\x79\x9e\xcf\x6c\xaa\x17\x88\x81\x6d\xd6\x99\x8c\xe2\x10\x0d\x02\xcd\x8a\xf7\xe9\x0a\xb4\xb1\x0e\x2e\xe6\xeb\x57\x4b\x65\xae\xf3\x73\xde\x01\xa1\xfa\xa0\x1d\x49\x9f\x29\x4f\x45\xcf\xa4\x72\x95\x41\x8a\x62\x8a\xe0\xc2\xac\x4e\x19\xc3\x38\x5d\xc1\xa4\x10\xc8\xac\x6e\x7d\x58\x2b\x48\x4a\x84\xe1\x21\x54\x5e\x28\x3b\xee\x93\xf4\x3a\xa9\xad\xaf\x93\xda\x2b\x08\x24\xe8\x10\x31\x86\x96\x7d\x16\xb8\x57\x01\xf5\xa3\xbb\x67\xf8\xbf\x4b\x54\x64\xef\x3c\x52\xb6\xd6\xae\x58\x2b\x87\x9e\x70\x71\x65\x33\xd4\x3a\x91\xa6\x55\x30\x55\xe6\x08\xa8\x31\x18\xc5\xc6\xc5\x1f\x35\x24\x92\xda\x10\x29\xd8\x76\xf8\xef\x62\xb8\x72\x5c\x2a\x3e\x2f\x0a\xf8\x2e\x8a\x5e\xbf\xee\xf5\xdf\x79\xfb\xf7\xeb\xaf\x77\xaa\x39\x3c\xda\xab\x70\xd8\x0f\xfd\xe1\xa8\x73\xe0\xf7\x5e\x4d\xbf\xdb\x3a\xa8\xdc\xf4\xeb\xb0\x37\xe8\x1c\x28\x29\x4d\xf5\xfc\x1a\xa1\x0d\x9f\x2a\xe0\xd5\xbb\xba\xfd\xf3\xf3\x5e\x77\x34\x19\xf4\x46\x83\xdf\x3a\xad\xe6\x2b\x6f\xd1\x70\x74\x32\x18\xc1\xf0\xf4\xe4\x53\xcf\x1f\x2e\x5a\x5f\x0e\x9e\xa1\x3d\x3a\xa0\xd9\x5b\x35\xeb\xb4\xda\x36\x1b\x2e\x33\x68\xb8\x59\xa0\x58\xd7\xa2\x18\xd8\x28\xb2\x3d\x6a\xe1\x35\x88\xad\xac\xd6\x9d\x6a\xc8\xb5\x41\x01\x84\x5c\x21\xc6\x44\xc6\xe9\xb3\x15\xe8\xda\x57\x42\x22\x7a\x4b\x6c\x04\xe8\x4e\xcb\x35\xa9\x84\x41\x31\x55\xd4\x0a\xcd\x66\xfe\x90\xde\x12\x24\xcd\xb7\x84\xb8\xae\xb6\xb3\x69\x6f\xc9\x42\x6a\xd3\xf1\xdc\x81\x24\x1a\x55\xc7\xb2\x6f\x3b\xd6\xac\x60\xe8\x3c\x50\xbe\xfe\x5f\x77\x0c\xad\x9d\x1d\x43\xb3\x79\xb6\xee\x19\x9c\xe1\x5b\x3d\x65\x01\xdf\xfa\x1d\x68\x1b\xee\xee\x2b\xd6\x95\xde\x03\xba\xa1\x33\x1a\xfb\x90\xb6\x7b\xb8\xef\x6e\xc0\x77\xb5\xdf\x31\x2a\x9d\x3a\xa0\xf9\x94\xda\xd2\x0d\x29\x2f\x54\xa1\xcc\x0e\x9c\x17\x8b\xcd\x6a\x75\xb9\x2e\xbf\xf9\x5e\x6e\xa4\x8e\xb2\x1f\x1f\x52\xba\xc8\xd6\xaf\x0f\x25\x67\xf1\x7f\x67\xb0\x37\x81\xd6\x67\x32\x40\x6d\xbb\x89\x01\xd2\xe0\xb3\xe2\x06\xfb\x82\x61\xde\x0f\x54\x1c\x70\xd5\xf1\x5a\xf3\xd3\x2a\x3e\xb7\x66\xe8\x46\x2c\xde\xff\x06\x00\x00\xff\xff\x04\x94\x6c\x8c\x6c\x1a\x00\x00")

func artifactsMysqlStatefulsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlStatefulsetYaml,
		"artifacts/mysql-statefulset.yaml",
	)
}

func artifactsMysqlStatefulsetYaml() (*asset, error) {
	bytes, err := artifactsMysqlStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"artifacts/backup-crd.yaml":         artifactsBackupCrdYaml,
	"artifacts/backup-cronjob.yaml":     artifactsBackupCronjobYaml,
	"artifacts/backup-pvc.yaml":         artifactsBackupPvcYaml,
	"artifacts/mysql-configmap.yaml":    artifactsMysqlConfigmapYaml,
	"artifacts/mysql-crd.yaml":          artifactsMysqlCrdYaml,
	"artifacts/mysql-service-read.yaml": artifactsMysqlServiceReadYaml,
	"artifacts/mysql-service.yaml":      artifactsMysqlServiceYaml,
	"artifacts/mysql-statefulset.yaml":  artifactsMysqlStatefulsetYaml,
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
	"artifacts": {nil, map[string]*bintree{
		"backup-crd.yaml":         {artifactsBackupCrdYaml, map[string]*bintree{}},
		"backup-cronjob.yaml":     {artifactsBackupCronjobYaml, map[string]*bintree{}},
		"backup-pvc.yaml":         {artifactsBackupPvcYaml, map[string]*bintree{}},
		"mysql-configmap.yaml":    {artifactsMysqlConfigmapYaml, map[string]*bintree{}},
		"mysql-crd.yaml":          {artifactsMysqlCrdYaml, map[string]*bintree{}},
		"mysql-service-read.yaml": {artifactsMysqlServiceReadYaml, map[string]*bintree{}},
		"mysql-service.yaml":      {artifactsMysqlServiceYaml, map[string]*bintree{}},
		"mysql-statefulset.yaml":  {artifactsMysqlStatefulsetYaml, map[string]*bintree{}},
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
