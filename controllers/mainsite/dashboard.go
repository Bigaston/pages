package mainsite

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

type WebsiteData struct {
	Name       string
	LastChange time.Time
}

func Dashboard(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	entries, err := os.ReadDir(utils.DataDirectory)
	if err != nil {
		return err
	}

	websites := make([]WebsiteData, 0, 20)

	for _, e := range entries {
		repository, err := git.PlainOpen(path.Join(utils.DataDirectory, e.Name()))

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

		commit, err := cIter.Next()

		if err != nil {
			return err
		}

		websites = append(websites, WebsiteData{
			Name:       e.Name(),
			LastChange: commit.Author.When,
		})
	}

	fmt.Println(websites)

	return c.Render("dashboard", fiber.Map{
		"Websites": websites,
	})
}
