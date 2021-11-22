package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
)

func SaveMedia(fh *multipart.FileHeader, path string) error {
	var err error
	var f multipart.File
	if f, err = fh.Open(); err != nil {
		return err
	}
	defer f.Close()

	dir := filepath.Join(configs.AppConfig.StaticBaseDir, path)
	if err = os.MkdirAll(filepath.Dir(dir), os.ModePerm); err != nil {
		return err
	}
	var out *os.File
	if out, err = os.Create(dir); err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, f); err != nil {
		return err
	}
	return nil
}

func DeleteMedia(path string) error {
	var err error
	dir := filepath.Join(configs.AppConfig.StaticBaseDir, path)
	if err = os.Remove(dir); err != nil {
		return err
	}
	return nil
}
