package github

import (
	"context"
	"fmt"
)

func (c CLI) GetRepo(ctx context.Context, repoName string) error {
	if repoName == "" {
		return fmt.Errorf("nome do repositório não pode ser vazio")
	}

	_, err := c.Runner.RunCommand(
		ctx,
		"git",
		"clone",
		fmt.Sprintf("git@github.com:%s/%s.git", Name, repoName),
	)

	if err != nil {
		return err
	}

	return nil
}
