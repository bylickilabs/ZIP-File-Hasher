package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ExtractZipToTemp(zipPath string) ([]string, error) {
	var extracted []string

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	tempDir := zipPath + "_extracted"
	os.MkdirAll(tempDir, 0755)

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}

		outPath := filepath.Join(tempDir, f.Name)
		os.MkdirAll(filepath.Dir(outPath), 0755)

		src, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		dest, err := os.Create(outPath)
		if err != nil {
			return nil, err
		}
		defer dest.Close()

		_, err = io.Copy(dest, src)
		if err != nil {
			return nil, err
		}

		extracted = append(extracted, outPath)
	}

	return extracted, nil
}