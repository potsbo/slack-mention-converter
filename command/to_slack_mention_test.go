package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestToSlackMentionCommand_implement(t *testing.T) {
	var _ cli.Command = &ToSlackMentionCommand{}
}
