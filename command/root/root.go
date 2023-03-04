package root

import (
	"fmt"
	"os"

	"github.com/polygomic/polygomic-edge/command/backup"
	"github.com/polygomic/polygomic-edge/command/genesis"
	"github.com/polygomic/polygomic-edge/command/helper"
	"github.com/polygomic/polygomic-edge/command/ibft"
	"github.com/polygomic/polygomic-edge/command/license"
	"github.com/polygomic/polygomic-edge/command/monitor"
	"github.com/polygomic/polygomic-edge/command/peers"
	"github.com/polygomic/polygomic-edge/command/secrets"
	"github.com/polygomic/polygomic-edge/command/server"
	"github.com/polygomic/polygomic-edge/command/status"
	"github.com/polygomic/polygomic-edge/command/txpool"
	"github.com/polygomic/polygomic-edge/command/version"
	"github.com/polygomic/polygomic-edge/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
