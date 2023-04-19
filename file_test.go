package utils_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/dmitrymomot/go-utils"
	"github.com/test-go/testify/assert"
)

func TestGetFileTypeByURI(t *testing.T) {
	tests := []struct {
		uri    string
		expect string
	}{
		{uri: "/path/to/file.png", expect: "image/png"},
		{uri: "/path/to/file.gif", expect: "image/gif"},
		{uri: "/path/to/file.jpg", expect: "image/jpeg"},
		{uri: "/path/to/file.mp4", expect: "video/mp4"},
		{uri: "/path/to/file.mov", expect: "video/quicktime"},
		{uri: "/path/to/file.mp3", expect: "audio/mpeg"},
		{uri: "/path/to/file.flac", expect: "audio/flac"},
		{uri: "/path/to/file.wav", expect: "audio/wav"},
		{uri: "/path/to/file.glb", expect: "model/gltf-binary"},
		{uri: "/path/to/file.gltf", expect: "model/gltf+json"},
		{uri: "/path/to/file.html", expect: "text/html"},
		{uri: "/path/to/file.js", expect: "application/javascript"},
		{uri: "/path/to/file.css", expect: "text/css"},
		{uri: "/path/to/file.json", expect: "application/json"},
		{uri: "/path/to/file.xml", expect: "application/xml"},
		{uri: "/path/to/file.svg", expect: "image/svg+xml"},
		{uri: "/path/to/file.ico", expect: "image/x-icon"},
		{uri: "/path/to/file.zip", expect: "application/zip"},
		{uri: "/path/to/file.pdf", expect: "application/pdf"},
		{uri: "/path/to/file.txt", expect: "text/plain"},
		{uri: "/path/to/file.md", expect: "text/markdown"},
		{uri: "/path/to/file.csv", expect: "text/csv"},
		{uri: "https://example.com/image.png", expect: "image/png"},
		{uri: "https://example.com/image?ext=png", expect: "image/png"},
		{uri: "https://example.com/image?format=png", expect: "image/png"},
		{uri: "https://example.com/file.mp4?ext=mkv", expect: "video/mp4"},
		{uri: "https://example.com/file.mp4?format=mkv", expect: "video/mp4"},
		{uri: "https://example.com/file.mp4?format=webm", expect: "video/mp4"},
		{uri: "https://example.com/file.unknown?format=xxx", expect: "application/octet-stream"},
	}

	for _, test := range tests {
		actual := utils.GetFileTypeByURI(test.uri)
		if actual != test.expect {
			t.Errorf("GetFileTypeByURI(%s): expected %s, but got %s", test.uri, test.expect, actual)
		}
	}
}

func TestGetFileByPath(t *testing.T) {
	// Test with empty path
	t.Run("empty path", func(t *testing.T) {
		_, err := utils.GetFileByPath("")
		if err == nil {
			t.Error("expected error but got none")
		}
	})

	// Test with unsupported URL schema
	t.Run("unsupported URL schema", func(t *testing.T) {
		_, err := utils.GetFileByPath("ftp://example.com/file.txt")
		if err == nil {
			t.Error("expected error but got none")
		}
	})

	// Test with non-existent local file
	t.Run("non-existent local file", func(t *testing.T) {
		_, err := utils.GetFileByPath("/path/to/non-existent-file.txt")
		if err == nil {
			t.Error("expected error but got none")
		}
	})

	// Test with existing local file
	t.Run("existing local file", func(t *testing.T) {
		data, err := utils.GetFileByPath("./testdata/testfile.txt")
		if err != nil {
			t.Errorf("failed to read local file: %v", err)
		}
		if string(data) != "test file content" {
			t.Errorf("got wrong data: %s", data)
		}
	})

	// Test with remote file
	t.Run("remote file", func(t *testing.T) {
		data, err := utils.GetFileByPath("https://raw.githubusercontent.com/golang/go/master/README.md")
		if err != nil {
			t.Errorf("failed to download remote file: %v", err)
		}
		if len(data) == 0 {
			t.Error("got empty data from remote file")
		}
	})
}

func TestDownloadFile(t *testing.T) {
	// Test with invalid URL
	t.Run("invalid URL", func(t *testing.T) {
		_, err := utils.DownloadFile("invalid-url")
		if err == nil {
			t.Error("expected error but got none")
		}
	})

	// Test with non-existent URL
	t.Run("non-existent URL", func(t *testing.T) {
		_, err := utils.DownloadFile("https://non-existent.com/file.txt")
		if err == nil {
			t.Error("expected error but got none")
		}
	})

	// Test with existing URL
	t.Run("existing URL", func(t *testing.T) {
		data, err := utils.DownloadFile("https://raw.githubusercontent.com/golang/go/master/README.md")
		if err != nil {
			t.Errorf("failed to download file: %v", err)
		}
		if len(data) == 0 {
			t.Error("got empty data from file")
		}
	})
}

func TestGetFileContentType(t *testing.T) {
	testCases := []struct {
		name   string
		input  io.Reader
		output string
		err    error
	}{
		{
			name:   "empty reader",
			input:  nil,
			output: "",
			err:    utils.ErrInvalidReader,
		},
		{
			name:   "text file",
			input:  bytes.NewBufferString("hello, world!"),
			output: "text/plain",
			err:    nil,
		},
		{
			name:   "unknown file type",
			input:  bytes.NewBuffer([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}),
			output: "application/octet-stream",
			err:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contentType, err := utils.GetFileContentType(tc.input)

			if contentType != tc.output {
				t.Errorf("expected content type %q but got %q", tc.output, contentType)
			}

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v but got %v", tc.err, err)
			}
		})
	}
}

func TestGetFileContentTypeByBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
		err      error
	}{
		{
			name:     "valid image/jpeg file",
			input:    []byte{0xff, 0xd8, 0xff},
			expected: "image/jpeg",
			err:      nil,
		},
		{
			name:     "empty input",
			input:    []byte{},
			expected: "",
			err:      utils.ErrEmptyInput,
		},
		{
			name:     "invalid input",
			input:    []byte{0x01, 0x02, 0x03},
			expected: "application/octet-stream",
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := utils.GetFileContentTypeByBytes(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q but got %q", tt.expected, actual)
			}
			if tt.err != nil && !errors.Is(err, tt.err) {
				t.Errorf("expected error %v but got %v", tt.err, err)
			}
		})
	}
}

func TestGetMaxFileParts(t *testing.T) {
	testCases := []struct {
		name           string
		file           io.ReadSeeker
		partSize       int64
		expectedResult int64
		expectedError  error
	}{
		{
			name:           "Valid input",
			file:           strings.NewReader("This is a test file"),
			partSize:       5,
			expectedResult: 4,
			expectedError:  nil,
		},
		{
			name:           "Invalid file",
			file:           nil,
			partSize:       10,
			expectedResult: 0,
			expectedError:  utils.ErrInvalidReader,
		},
		{
			name:           "Invalid part size",
			file:           strings.NewReader("This is another test file"),
			partSize:       -1,
			expectedResult: 0,
			expectedError:  utils.ErrInvalidPartSize,
		},
		{
			name:           "File size less than part size",
			file:           strings.NewReader("Test"),
			partSize:       10,
			expectedResult: 1,
			expectedError:  nil,
		},
		{
			name:           "File size equals to part size",
			file:           strings.NewReader("Testing"),
			partSize:       7,
			expectedResult: 1,
			expectedError:  nil,
		},
		{
			name:           "File size larger than part size",
			file:           strings.NewReader("Hello world!"),
			partSize:       5,
			expectedResult: 3,
			expectedError:  nil,
		},
		{
			name:           "Large file, max number of parts is limited to 100000",
			file:           strings.NewReader(strings.Repeat("a", 100000000)),
			partSize:       1000,
			expectedResult: 100000,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := utils.GetMaxFileParts(tc.file, tc.partSize)

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError != nil {
				if !errors.Is(err, tc.expectedError) {
					t.Errorf("expected error %v but got %v", tc.expectedError, err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetFileSize(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		reader := strings.NewReader("hello, world")
		size, err := utils.GetFileSize(reader)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if size != int64(len("hello, world")) {
			t.Errorf("expected size %d, got %d", len("hello, world"), size)
		}
	})

	t.Run("invalid file", func(t *testing.T) {
		var reader *strings.Reader
		_, err := utils.GetFileSize(reader)
		if !errors.Is(err, utils.ErrInvalidReader) {
			t.Errorf("expected error %v, got %v", utils.ErrInvalidReader, err)
		}
	})

	t.Run("failed to get file size", func(t *testing.T) {
		reader := &errReader{}
		_, err := utils.GetFileSize(reader)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

// mock ReadSeeker that always returns an error
type errReader struct{}

// implements io.Seeker
func (r *errReader) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("failed to seek")
}

// implements io.Reader
func (r *errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failed to read")
}
