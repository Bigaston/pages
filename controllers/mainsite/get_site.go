package mainsite

import (
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
	Changes    []siteChange
}

type siteChange struct {
	Hash        string
	FileChanges []fileChange
}

type fileChange struct {
	Name   string
	Action string
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
	var treeNext *object.Tree
	var commitNext *object.Commit

	siteChanges := make([]siteChange, 0, 10)

	err = cIter.ForEach(func(commit *object.Commit) error {
		if lastCommit == nil {
			lastCommit = commit
			commitNext = commit
			treeNext, err = commit.Tree()

			if err != nil {
				return err
			}

			return nil
		}

		myTree, err := commit.Tree()

		if err != nil {
			return err
		}

		changes, err := treeNext.Diff(myTree)

		if err != nil {
			return err
		}

		fileChanges := make([]fileChange, 0, 10)

		for _, change := range changes {
			action, _ := change.Action()

			name := change.From.Name

			if name == "" {
				name = change.To.Name
			}

			fileChanges = append(fileChanges, fileChange{
				Name:   name,
				Action: action.String(),
			})
		}

		siteChanges = append(siteChanges, siteChange{
			Hash:        commitNext.Hash.String(),
			FileChanges: fileChanges,
		})

		treeNext = myTree
		commitNext = commit

		return nil
	})

	if err != nil {
		return err
	}

	siteData := siteData{
		Name:       siteName,
		LastChange: lastCommit.Author.When,
		Changes:    siteChanges,
	}

	return c.Render("site_edit", siteData)
}
