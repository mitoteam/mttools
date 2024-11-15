package mttools

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Returns true if file exists and access is not denied.
func IsFileExists(path string) bool {
	file_info, err := os.Stat(path)

	if os.IsNotExist(err) || os.IsPermission(err) {
		return false
	}

	if file_info.IsDir() {
		return false
	}

	return true
}

// Returns true if directory exists and access is not denied.
func IsDirExists(path string) bool {
	file_info, err := os.Stat(path)

	if os.IsNotExist(err) || os.IsPermission(err) {
		return false
	}

	if !file_info.IsDir() {
		return false
	}

	return true
}

// Returns absolute path for directory. If `path` is "" current working directory is used.
func GetDirAbsolutePath(path string) (abs_path string, err error) {
	abs_path = path

	if abs_path == "" {
		abs_path = "." //current directory
	}

	if !filepath.IsAbs(abs_path) {
		abs_path, err = filepath.Abs(path)
		if err != nil {
			return
		}
	}

	if !IsDirExists(abs_path) {
		return abs_path, errors.New("\"" + path + "\" directory does not exists")
	}

	return
}

// Convert fille size to human-readable string
func FormatFileSize(size int64) string {
	if size > 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fTb", float64(size)/(1024*1024*1024*1024))
	} else if size > 1024*1024*1024 {
		return fmt.Sprintf("%2.2fGb", float64(size)/(1024*1024*1024))
	} else if size > 1024*1024 {
		return fmt.Sprintf("%2.2fMb", float64(size)/(1024*1024))
	} else if size > 1024 {
		return fmt.Sprintf("%2.2fKb", float64(size)/1024)
	}

	return fmt.Sprintf("%db", size)
}

func FileSha256(path string) (string, error) {
	if !IsFileExists(path) {
		return "", errors.New(fmt.Sprintf("Inaccessible file: %s", path))
	}

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
