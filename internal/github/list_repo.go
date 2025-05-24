package github

import (
	"bufio"
	"context"
	"strings"
)

// GetProjects retorna uma lista de repositórios do GitHub usando gh CLI
func (c CLI) GetProjects(ctx context.Context) ([]string, error) {
	output, err := c.Runner.RunCommand(ctx, "gh", "repo", "list")
	if err != nil {
		return nil, err
	}

	var projects []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))

	// Process each line and extract project names
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Remove trailing whitespace and add to projects list
		projectName := strings.TrimSpace(line)
		if projectName != "" {
			projects = append(projects, projectName)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
