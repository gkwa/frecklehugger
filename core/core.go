package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-logr/logr"
)

type GitNotes interface {
	GetNotes(path string) ([]*Note, error)
}

type Note struct {
	CommitHash string
	Content    string
}

type GitNotesService struct {
	logger logr.Logger
}

func NewGitNotesService(logger logr.Logger) GitNotes {
	return &GitNotesService{
		logger: logger,
	}
}

func PrintNotes(ctx context.Context, path string, logger logr.Logger) {
	service := NewGitNotesService(logger)
	notes, err := service.GetNotes(path)
	if err != nil {
		logger.Error(err, "Failed to get notes")
		return
	}

	for _, note := range notes {
		fmt.Printf("Note for commit %s:\n%s\n", note.CommitHash, note.Content)
	}
}

func (g *GitNotesService) GetNotes(path string) ([]*Note, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return nil, err
	}

	repo, err := git.PlainOpen(absPath)
	if err != nil {
		return nil, err
	}

	notesRef, err := repo.Reference(plumbing.ReferenceName("refs/notes/commits"), true)
	if err != nil {
		return nil, err
	}

	notesCommit, err := repo.CommitObject(notesRef.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := notesCommit.Tree()
	if err != nil {
		return nil, err
	}

	var notes []*Note

	err = tree.Files().ForEach(func(f *object.File) error {
		commitHash := plumbing.NewHash(f.Name).String()

		content, err := f.Contents()
		if err != nil {
			return err
		}

		notes = append(notes, &Note{
			CommitHash: commitHash,
			Content:    content,
		})

		return nil
	})

	return notes, err
}

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")
}
