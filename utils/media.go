package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveMedia(fh *multipart.FileHeader, d string) error {
	var err error
	var f multipart.File
	if f, err = fh.Open(); err != nil {
		return err
	}
	defer f.Close()

	if err = os.MkdirAll(d, os.ModePerm); err != nil {
		return err
	}
	var out *os.File
	if out, err = os.Create(d); err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, f); err != nil {
		return err
	}
	return nil
}

func DeleteMedia(d string) error {
	var err error
	if err = os.Remove(d); err != nil {
		return err
	}
	return nil
}
