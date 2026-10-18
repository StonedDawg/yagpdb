package invite

import (
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

var Command = &commands.YAGCommand{
	CmdCategory:         commands.CategoryGeneral,
	Name:                "Invite",
	Description:         "Responds with bot invite link",
	SlashCommandEnabled: true,
	DefaultEnabled:      true,
	RunInDM:             true,

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		return "Please don't think about it too much", nil
	},
}
