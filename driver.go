package MultipleDriver

import (
	"errors"
	"io"
	"strings"

	"github.com/goftp/server"
)

type MultipleDriver struct {
	drivers map[string]server.Driver
}

func (driver *MultipleDriver) ChangeDir(path string) error {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.ChangeDir(strings.TrimPrefix(path, prefix))
		}
	}
	return errors.New("Not a directory")
}

func (driver *MultipleDriver) Stat(path string) (server.FileInfo, error) {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.Stat(strings.TrimPrefix(path, prefix))
		}
	}
	return nil, errors.New("Not a file")
}

func (driver *MultipleDriver) DirContents(path string) ([]server.FileInfo, error) {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.DirContents(strings.TrimPrefix(path, prefix))
		}
	}
	return nil, errors.New("Not a directory")
}

func (driver *MultipleDriver) DeleteDir(path string) error {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.DeleteDir(strings.TrimPrefix(path, prefix))
		}
	}
	return errors.New("Not a directory")
}

func (driver *MultipleDriver) DeleteFile(path string) error {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.DeleteFile(strings.TrimPrefix(path, prefix))
		}
	}

	return errors.New("Not a file")
}

func (driver *MultipleDriver) Rename(fromPath string, toPath string) error {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(fromPath, prefix) {
			return driver.Rename(strings.TrimPrefix(fromPath, prefix), strings.TrimPrefix(toPath, prefix))
		}
	}

	return errors.New("Not a file")
}

func (driver *MultipleDriver) MakeDir(path string) error {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.MakeDir(strings.TrimPrefix(path, prefix))
		}
	}
	return errors.New("Not a directory")
}

func (driver *MultipleDriver) GetFile(path string, offset int64) (int64, io.ReadCloser, error) {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(path, prefix) {
			return driver.GetFile(strings.TrimPrefix(path, prefix), offset)
		}
	}

	return 0, nil, errors.New("Not a file")
}

func (driver *MultipleDriver) PutFile(destPath string, data io.Reader, appendData bool) (int64, error) {
	for prefix, driver := range driver.drivers {
		if strings.HasPrefix(destPath, prefix) {
			return driver.PutFile(strings.TrimPrefix(destPath, prefix), data, appendData)
		}
	}

	return 0, errors.New("Not a file")
}

type MultipleDriverFactory struct {
	drivers map[string]server.Driver
}

func (factory *MultipleDriverFactory) NewDriver() (server.Driver, error) {
	return &MultipleDriver{factory.drivers}, nil
}
