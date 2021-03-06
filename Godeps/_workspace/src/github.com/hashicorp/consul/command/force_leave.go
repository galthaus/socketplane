package command

import (
	"flag"
	"fmt"
	"github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/mitchellh/cli"
	"strings"
)

// ForceLeaveCommand is a Command implementation that tells a running Consul
// to force a member to enter the "left" state.
type ForceLeaveCommand struct {
	Ui cli.Ui
}

func (c *ForceLeaveCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("join", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	rpcAddr := RPCAddrFlag(cmdFlags)
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	nodes := cmdFlags.Args()
	if len(nodes) != 1 {
		c.Ui.Error("A node name must be specified to force leave.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	client, err := RPCClient(*rpcAddr)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Consul agent: %s", err))
		return 1
	}
	defer client.Close()

	err = client.ForceLeave(nodes[0])
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error force leaving: %s", err))
		return 1
	}

	return 0
}

func (c *ForceLeaveCommand) Synopsis() string {
	return "Forces a member of the cluster to enter the \"left\" state"
}

func (c *ForceLeaveCommand) Help() string {
	helpText := `
Usage: consul force-leave [options] name

  Forces a member of a Consul cluster to enter the "left" state. Note
  that if the member is still actually alive, it will eventually rejoin
  the cluster. This command is most useful for cleaning out "failed" nodes
  that are never coming back. If you do not force leave a failed node,
  Consul will attempt to reconnect to those failed nodes for some period of
  time before eventually reaping them.

Options:

  -rpc-addr=127.0.0.1:8400 RPC address of the Consul agent.

`
	return strings.TrimSpace(helpText)
}
