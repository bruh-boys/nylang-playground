package setup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func unzipLib() {
	d, _ := os.UserHomeDir()
	unzip("lib.zip", filepath.Join(d, ".nylang"))
	unzip("codemirror.zip", "./")

}

//thanks to https://golangcode.com/unzip-files-in-go/

func unzip(src string, dest string) (filenames []string, err error) {

	r, err := zip.OpenReader(src)
	if err != nil {
		return
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		outFile, _ := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

		rc, _ := f.Open()

		defer outFile.Close()
		defer rc.Close()
		io.Copy(outFile, rc)

	}
	return
}
