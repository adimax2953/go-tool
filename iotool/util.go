package iotool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	PathError = errors.New("invalid path")
)

// EnsureDir - 確保目錄已創建
func EnsureDir(path string, perm os.FileMode) error {
	if 0 >= len(path) {
		return PathError
	}

	fi, err := os.Stat(path)
	if nil != err && !os.IsNotExist(err) {
		return err
	}

	if nil != fi && !fi.IsDir() {
		err = os.Remove(path)
		if nil != err {
			return err
		}
		fi = nil
	}
	if nil == fi {
		return os.Mkdir(path, perm)
	}

	return nil
}

// FileSize - 文件大小
func FileSize(path string) (size int64, err error) {
	if 0 >= len(path) {
		err = PathError
		return
	}

	var fi os.FileInfo
	fi, err = os.Stat(path)
	if nil == err {
		size = fi.Size()
	}

	return
}

// MakeOrDeleteDir - 圖片文件
func MakeOrDeleteDir(path string, perm os.FileMode) error {
	if 0 >= len(path) {
		return PathError
	}

	static, err := os.Stat(path)
	if err != nil {
		if mkErr := os.Mkdir(path, perm); mkErr == nil {
			return nil
		} else {
			return mkErr
		}
	}

	if static != nil && static.IsDir() {
		if deleteErr := os.RemoveAll(path); deleteErr == nil {
			if makeErr := os.Mkdir(path, perm); makeErr == nil {
				return nil
			} else {
				return makeErr
			}
		} else {
			return deleteErr
		}
	}

	return nil
}

// DirSize - 文件夾大小
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// CleanEmptyLogFile -清理空日誌文件
func CleanEmptyLogFile(logDir string) {
	arr, err := ioutil.ReadDir(logDir)
	if nil != err {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "ioutil.ReadDir error: %s\n", err)
		}
		return
	}
	for _, fi := range arr {
		if 0 == fi.Size() {
			absPath, err := filepath.Abs(filepath.Join(logDir, fi.Name()))
			if nil != err {
				fmt.Fprintf(os.Stderr, "filepath.Abs error: %s\n", err)
				continue
			}
			if err = os.Remove(absPath); nil != err {
				fmt.Fprintf(os.Stderr, "os.Remove %s error: %s\n", absPath, err)
				continue
			}
		}
	}
}
