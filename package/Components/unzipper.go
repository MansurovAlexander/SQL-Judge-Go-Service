package components

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func SaveZipFile(fileName, src, folder string) error {
	decodeData, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return err
	}
	var filePath strings.Builder
	filePath.WriteString(folder)
	filePath.WriteString(fileName)
	file, err := os.Create(filePath.String())
	if err != nil {
		return err
	}
	if _, err = file.Write(decodeData); err != nil {
		return err
	}
	if err = file.Sync(); err != nil {
		return err
	}
	file.Close()
	return nil
}

func UnzipPackage(zipName, src, sqlName, folder string) error {
	if err := SaveZipFile(zipName, src, folder); err != nil {
		return err
	}
	filePath := filepath.Join(folder, zipName)
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		return err
	}
	for _, f := range archive.File {

		// Create the destination file path
		filePath := filepath.Join(folder, sqlName)

		// Print the file path
		fmt.Println("extracting file ", filePath)

		// Check if the file is a directory
		if f.FileInfo().IsDir() {
			// Create the directory
			fmt.Println("creating directory...")
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		// Create an empty destination file
		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			return err
		}
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		// Close the files
		dstFile.Close()
		srcFile.Close()
	}
	archive.Close()
	//Remove zipName file
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
