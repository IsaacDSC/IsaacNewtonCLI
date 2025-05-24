package github

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListProject(t *testing.T) {
	gh := NewGithubCLI()
	repos, err := gh.GetProjects(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, repos)
}
