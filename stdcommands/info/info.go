package info

import (
	"fmt"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

var Command = &commands.YAGCommand{
	CmdCategory:         commands.CategoryGeneral,
	Name:                "Info",
	Description:         "Responds with bot information",
	SlashCommandEnabled: true,
	DefaultEnabled:      true,
	RunInDM:             true,
	RunFunc: func(data *dcmd.Data) (any, error) {
		info := fmt.Sprintf(`**StonedBot - lets get high Bot**
This bot focuses on being configurable and therefore is one of the most advanced bots.
let's get stoned together!
				`)

		return info, nil
	},
}
