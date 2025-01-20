package mainsite

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

type siteData struct {
	Name       string
	LastChange time.Time
}

func GetSite(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	siteName := c.Params("site")
	siteDirectory := path.Join(utils.DataDirectory, siteName)

	if _, err := os.Stat(siteDirectory); err != nil && os.IsNotExist(err) {
		return c.SendStatus(http.StatusNotFound)
	}

	repository, err := git.PlainOpen(siteDirectory)

	if err != nil {
		return err
	}

	head, err := repository.Head()

	if err != nil {
		return err
	}

	cIter, err := repository.Log(&git.LogOptions{
		From: head.Hash(),
	})

	if err != nil {
		return err
	}

	var lastCommit *object.Commit

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		if lastCommit == nil {
			lastCommit = c
		}

		return nil
	})

	if err != nil {
		return err
	}

	siteData := siteData{
		Name:       siteName,
		LastChange: lastCommit.Author.When,
	}

	return c.Render("site_edit", siteData)
}
