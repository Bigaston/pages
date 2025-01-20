package utils

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/gofiber/fiber/v2"
)

func UploadSite(siteName string, c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("file")

	if err != nil {
		return "", err
	}

	if _, err := os.Stat(TempDirectory); err != nil && os.IsNotExist(err) {
		os.Mkdir(TempDirectory, os.ModePerm)
	}

	if _, err := os.Stat(DataDirectory); err != nil && os.IsNotExist(err) {
		os.Mkdir(DataDirectory, os.ModePerm)
	}

	siteDirectory := fmt.Sprintf("%s/%s", DataDirectory, siteName)

	var repository *git.Repository

	if _, err := os.Stat(siteDirectory); err != nil && os.IsNotExist(err) {
		os.Mkdir(siteDirectory, os.ModePerm)
		repository, err = git.PlainInit(siteDirectory, false)

		if err != nil {
			return "", err
		}
	} else {
		repository, err = git.PlainOpen(siteDirectory)

		if err != nil {
			return "", err
		}
	}

	worktree, err := repository.Worktree()

	if err != nil {
		return "", err
	}

	tempFilePath := path.Join(TempDirectory, file.Filename)
	err = c.SaveFile(file, tempFilePath)

	if err != nil {
		return "", err
	}

	archive, err := zip.OpenReader(tempFilePath)

	if err != nil {
		return "", err
	}

	defer func() {
		err := os.Remove(tempFilePath)

		if err != nil {
			fmt.Println(err)
		}
	}()
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(siteDirectory, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(siteDirectory)+string(os.PathSeparator)) {
			return "", errors.New("invalid file path")
		}

		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return "", err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return "", err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return "", err
		}

		worktree.Add(f.Name)

		dstFile.Close()
		fileInArchive.Close()
	}

	commit, err := worktree.Commit(fmt.Sprintf("Change %s", time.Now()), &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Webmaster",
			Email: "webmaster@domain.com",
			When:  time.Now(),
		},
	})

	if err != nil {
		if errors.Is(err, git.ErrEmptyCommit) {
			return "OK, no change to commit", nil
		}
		return "", err
	}

	obj, err := repository.CommitObject(commit)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("OK: Commit Hash %s", obj.Hash), nil
}
