package github

import (
	"context"
	"github.com/IsaacDSC/IsaacNewtonCLI/internal/entity"
)

/*
# Add user with read permission
gh repo add-collaborator john-doe --permission read

# Add user with write permission
gh repo add-collaborator developer-jane --permission write

# Add user with admin permission
gh repo add-collaborator lead-dev --permission admin

# Add to specific repository (if not in current directory)
gh repo add-collaborator username --permission write --repo owner/repository

*/

func (c CLI) AddCollaborator(ctx context.Context, repoName string, collaborator entity.Collaborator, permission entity.CollaboratorPermissions) error {
	_, err := c.Runner.RunCommand(
		ctx,
		"gh",
		"repo",
		repoName,
		collaborator.Username,
		"--permission", permission.String())

	if err != nil {
		return err
	}

	return nil
}
