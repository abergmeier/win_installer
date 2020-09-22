package unarchive

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Run(config map[string]interface{}) error {

	srcConfig, ok := config["src"]
	if !ok {
		return errors.New("Missing src config")
	}
	if srcConfig == nil {
		return errors.New("Missing value in src config")
	}
	destConfig, ok := config["dest"]
	if !ok {
		return errors.New("Missing dest config")
	}
	if destConfig == nil {
		return errors.New("Missing value in dest config")
	}

	src := srcConfig.(string)
	dest := destConfig.(string)

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	return writeFilesToDir(r.File, dest)
}

func writeFilesToDir(files []*zip.File, dest string) error {
	for _, f := range files {

		fp := filepath.Join(dest, f.Name)

		// Prevent ZipSlip
		if !strings.HasPrefix(fp, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("Illegal file path: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fp, os.ModePerm)
			continue
		}

		err := writeZippedFile(f, fp)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeZippedFile(f *zip.File, fp string) error {
	r, err := f.Open()
	if err != nil {
		return err
	}
	defer r.Close()
	return copyFileContent(r, fp, f.Mode())
}

func copyFileContent(r io.Reader, fp string, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	if err != nil {
		return err
	}
	outFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, r)
	if err != nil {
		return err
	}
	return outFile.Sync()
}
