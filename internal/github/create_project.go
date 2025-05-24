package github

import (
	"context"
	"fmt"
	"github.com/IsaacDSC/IsaacNewtonCLI/internal/entity"
)

const BaseUrlHttps = "https://github.com"

type Url string

func (u Url) String() string {
	return string(u)
}

// gh repo create example-test --private --gitignore Go --license=MIT --add-readme
func (c CLI) CreateRepo(ctx context.Context, repoName string, repoDescription string, platform entity.PlatformType) (Url, error) {
	_, err := c.Runner.RunCommand(
		ctx,
		"gh",
		"repo",
		"create",
		repoName,
		"--private",
		fmt.Sprintf("--gitignore=%s", platform),
		"--license=MIT",
		"--add-readme",
		"--description",
		repoDescription,
	)

	if err != nil {
		return "", err
	}

	//https://github.com/IsaacDSC/example-test1
	return Url(fmt.Sprintf("%s/%s/%s", BaseUrlHttps, Name, repoName)), nil
}
