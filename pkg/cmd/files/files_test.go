package files

import (
	"fmt"
	"github.com/debricked/cli/pkg/client"
	"testing"
)

func TestNewFilesCmd(t *testing.T) {
	var debClient client.IDebClient = client.NewDebClient(nil)
	cmd := NewFilesCmd(&debClient)
	commands := cmd.Commands()
	nbrOfCommands := 1
	if len(commands) != nbrOfCommands {
		t.Error(fmt.Sprintf("failed to assert that there were %d sub commands connected", nbrOfCommands))
	}
}
