package controllers

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const tempDirectory = "./temp"
const dataDirectory = "./data"

func Init(app *fiber.App) {
	app.Post("/~upload/:site", uploadSite)
}

func uploadSite(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	if _, err := os.Stat(tempDirectory); err != nil && os.IsNotExist(err) {
		os.Mkdir(tempDirectory, os.ModePerm)
	}

	if _, err := os.Stat(dataDirectory); err != nil && os.IsNotExist(err) {
		os.Mkdir(dataDirectory, os.ModePerm)
	}

	siteDirectory := fmt.Sprintf("%s/%s", dataDirectory, c.Params("site", "default"))

	if _, err := os.Stat(siteDirectory); err != nil && os.IsNotExist(err) {
		os.Mkdir(siteDirectory, os.ModePerm)
	}

	err = c.SaveFile(file, fmt.Sprintf("%s/%s", tempDirectory, file.Filename))

	if err != nil {
		return err
	}

	archive, err := zip.OpenReader(fmt.Sprintf("%s/%s", tempDirectory, file.Filename))

	if err != nil {
		return err
	}

	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(siteDirectory, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(siteDirectory)+string(os.PathSeparator)) {
			return errors.New("invalid file path")
		}

		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return c.SendString("OK")
}
