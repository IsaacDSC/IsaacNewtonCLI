package github

import (
	"context"
	"fmt"
)

// GetProject retorna informações detalhadas sobre um repositório específico do GitHub
func (c CLI) ViewProject(ctx context.Context, repoName string) (string, error) {
	if repoName == "" {
		return "", fmt.Errorf("nome do repositório não pode ser vazio")
	}

	output, err := c.Runner.RunCommand(ctx, "gh", "repo", "view", repoName)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
