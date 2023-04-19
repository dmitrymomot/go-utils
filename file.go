package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

// default mime type
var defaultMimeType = "application/octet-stream"

// PRedefined mime types map
var mimeTypesMap = map[string]string{
	"png":  "image/png",
	"gif":  "image/gif",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"mp4":  "video/mp4",
	"webm": "video/mp4",
	"mkv":  "video/mp4",
	"mov":  "video/quicktime",
	"mp3":  "audio/mpeg",
	"flac": "audio/flac",
	"wav":  "audio/wav",
	"glb":  "model/gltf-binary",
	"gltf": "model/gltf+json",
	"html": "text/html",
	"js":   "application/javascript",
	"css":  "text/css",
	"json": "application/json",
	"xml":  "application/xml",
	"svg":  "image/svg+xml",
	"ico":  "image/x-icon",
	"zip":  "application/zip",
	"pdf":  "application/pdf",
	"txt":  "text/plain",
	"md":   "text/markdown",
	"csv":  "text/csv",
	"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

// GetFileTypeByURI returns the file type of the given URI
func GetFileTypeByURI(uri string) string {
	// Parse the uri.
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return defaultMimeType
	}

	// Get the file extension.
	ext := filepath.Ext(parsedUri.Path)
	if ext == "" {
		ext = parsedUri.Query().Get("ext")
		if ext == "" {
			ext = parsedUri.Query().Get("format")
		}
	}

	// Remove the query string.
	if strings.Contains(ext, "?") {
		ext = strings.Split(ext, "?")[0]
	}

	// get the mime type from the map.
	mime, ok := mimeTypesMap[strings.Trim(ext, ".")]
	if !ok {
		return defaultMimeType
	}

	return mime
}

// GetFileByPath returns the file bytes from the given path.
// If the path is a URL, it will download the file and return the bytes.
// If the path is a local file, it will read the file and return the bytes.
func GetFileByPath(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("path cannot be empty")
	}

	if strings.Contains(path, "://") {
		if !strings.HasPrefix(path, "http") {
			return nil, fmt.Errorf("this url schema is not supported: %s", path)
		}

		b, err := DownloadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to download file from url: %w", err)
		}

		return b, nil
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file from local disk: %w", err)
	}

	return b, nil
}

// DownloadFile downloads the file from the given URL and returns the bytes.
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return b, nil
}

// GetFileContentType returns the content type of a file.
func GetFileContentType(input io.Reader) (string, error) {
	if input == nil {
		return "", fmt.Errorf("invalid reader: %w", ErrInvalidReader)
	}

	mtype, err := mimetype.DetectReader(input)
	if err != nil {
		return "", fmt.Errorf("failed to detect file content type: %w", err)
	}

	parts := strings.Split(mtype.String(), ";")
	return parts[0], nil
}

// GetFileContentTypeByBytes returns the content type of a file.
func GetFileContentTypeByBytes(input []byte) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("invalid reader: %w", ErrEmptyInput)
	}

	mtype := mimetype.Detect(input)
	parts := strings.Split(mtype.String(), ";")
	return parts[0], nil
}

// Get max file parts can be if the file is split into parts with the given part size.
// The max file parts is 10000.
// file io.ReadSeeker: the file to be uploaded.
// partSize int64: the part size.
func GetMaxFileParts(file io.ReadSeeker, partSize int64) (int64, error) {
	if file == nil {
		return 0, fmt.Errorf("invalid file: %w", ErrInvalidReader)
	}
	if partSize <= 0 {
		return 0, ErrInvalidPartSize
	}

	// Get the file size.
	fileSize, err := GetFileSize(file)
	if err != nil {
		return 0, fmt.Errorf("failed to get file size: %w", err)
	}

	// Get the max file parts.
	maxFileParts := int64(fileSize / partSize)
	if fileSize%partSize != 0 {
		maxFileParts++
	}

	// Limit the number of parts to 100000.
	if maxFileParts > 100000 {
		maxFileParts = 100000
	}

	return maxFileParts, nil
}

// Get the file size.
// file io.ReadSeeker: the file to be uploaded.
func GetFileSize(file io.ReadSeeker) (int64, error) {
	if file == nil || reflect.ValueOf(file).IsNil() {
		return 0, fmt.Errorf("invalid file: %w", ErrInvalidReader)
	}

	// Get the file size.
	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, fmt.Errorf("failed to get file size: %w", err)
	}

	// Reset the read position to the beginning.
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return 0, fmt.Errorf("failed to reset file read position: %w", err)
	}

	return fileSize, nil
}

// Get the file name without extension.
// fileName string: the file name.
func GetFileNameWithoutExtension(fileName string) string {
	if fileName == "" {
		return ""
	}
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
