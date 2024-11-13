package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitNotes(t *testing.T) {
	// Create temp directory
	tempDir, err := os.MkdirTemp("", "git-notes-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Initialize repo
	repo, err := git.PlainInit(tempDir, false)
	require.NoError(t, err)

	// Configure git
	cmd := exec.Command("git", "config", "--global", "user.email", "test@example.com")
	err = cmd.Run()
	require.NoError(t, err)

	cmd = exec.Command("git", "config", "--global", "user.name", "test")
	err = cmd.Run()
	require.NoError(t, err)

	// Create test file and commit
	testFile := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test content"), 0o644)
	require.NoError(t, err)

	w, err := repo.Worktree()
	require.NoError(t, err)

	_, err = w.Add("test.txt")
	require.NoError(t, err)

	commit, err := w.Commit("test commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "test",
			Email: "test@example.com",
		},
	})
	require.NoError(t, err)

	// Add notes using git command
	cmd = exec.Command("git", "-C", tempDir, "notes", "add", "-m", "test note 1", commit.String())
	err = cmd.Run()
	require.NoError(t, err)

	cmd = exec.Command("git", "-C", tempDir, "notes", "add", "-f", "-m", "test note 2", commit.String())
	err = cmd.Run()
	require.NoError(t, err)

	// Test notes retrieval
	service := NewGitNotesService(logr.Discard())
	notes, err := service.GetNotes(tempDir)
	require.NoError(t, err)

	assert.Len(t, notes, 1)
	assert.Equal(t, commit.String(), notes[0].CommitHash)
	fmt.Printf("Note content: %s\n", notes[0].Content)
	assert.Equal(t, "test note 2", strings.TrimSpace(notes[0].Content))
}
