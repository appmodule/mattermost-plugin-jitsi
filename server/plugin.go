package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

const (
	commandTrigger = "jitsi"
)

// JitsiPlugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type JitsiPlugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// OnActivate register the plugin command
func (p *JitsiPlugin) OnActivate() error {

	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          commandTrigger,
		AutoComplete:     true,
		AutoCompleteHint: "[roomname]",
		AutoCompleteDesc: "Create Jitsi Meeting.",
	}); err != nil {
		return errors.Wrapf(err, "failed to register %s command", commandTrigger)
	}

	return nil
}

// ExecuteCommand executes a command that has been previously registered via the RegisterCommand
// API.
//
// This demo implementation responds to a /demo_plugin command, allowing the user to enable
// or disable the demo plugin's hooks functionality (but leave the command and webapp enabled).
func (p *JitsiPlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	trigger := strings.TrimPrefix(strings.Fields(args.Command)[0], "/")
	switch trigger {
	case commandTrigger:
		return p.executeCommand(args), nil

	default:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Unknown command: " + args.Command),
		}, nil
	}
}

func (p *JitsiPlugin) executeCommand(args *model.CommandArgs) *model.CommandResponse {
	channel, _ := p.API.GetChannel(args.ChannelId)
	team, _ := p.API.GetTeam(args.TeamId)
	user, _ := p.API.GetUser(args.UserId)
	command := strings.Fields(args.Command)
	room := fmt.Sprintf("%s_%s", team.Name, channel.Name)

	if len(command) > 1 {
		room = command[1]
	}

	config := p.getConfiguration()
	jitsiURL := strings.TrimSpace(config.JitsiURL)
	if len(jitsiURL) == 0 {
		jitsiURL = "https://meet.jit.si"
	}

	titleLink := fmt.Sprintf("%s/%s", jitsiURL, room)
	text := fmt.Sprintf("Meeting room created by %s", user.Username)

	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
		Props: model.StringInterface{
			"attachments": []*model.SlackAttachment{{
				AuthorName: "jitsi",
				AuthorIcon: "http://is3.mzstatic.com/image/thumb/Purple128/v4/33/0f/99/330f99b7-4e02-4990-ab79-d3440c4237be/source/512x512bb.jpg",
				Title:      fmt.Sprintf("Click here to join the meeting: %s.", room),
				TitleLink:  titleLink,
				Text:       text,
				Color:      "#ff0000",
			}},
		},
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
