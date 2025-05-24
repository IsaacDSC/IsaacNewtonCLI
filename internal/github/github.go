package github

import (
	"context"
	"os/exec"
)

const Name = "IsaacDSC"

// CommandRunner define uma interface para executar comandos
type CommandRunner interface {
	RunCommand(ctx context.Context, name string, args ...string) ([]byte, error)
}

// DefaultCommandRunner implementa CommandRunner usando exec.Command
type DefaultCommandRunner struct{}

// RunCommand executa um comando e retorna sua saída
func (r DefaultCommandRunner) RunCommand(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.Output()
}

// CLI implementa operações usando GitHub CLI
type CLI struct {
	Runner CommandRunner
}

// NewGithubCLI cria uma nova instância de CLI com o runner padrão
func NewGithubCLI() *CLI {
	return &CLI{
		Runner: DefaultCommandRunner{},
	}
}
