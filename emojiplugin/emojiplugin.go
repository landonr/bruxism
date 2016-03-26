package emojiplugin

import (
	"fmt"
	"os"
	"strings"

	"github.com/iopred/bruxism"
)

func emojiFile(s string) string {
	found := ""
	filename := ""
	for _, r := range s {
		if filename != "" {
			filename = fmt.Sprintf("%s-%x", filename, r)
		} else {
			filename = fmt.Sprintf("%x", r)
		}

		if _, err := os.Stat(fmt.Sprintf("emoji/twitter/%s.png", filename)); err == nil {
			found = filename
		} else if found != "" {
			return found
		}
	}
	return found
}

func emojiLoadFunc(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if service.Name() != bruxism.DiscordServiceName {
		panic("Emoji Plugin only supports Discord.")
	}
	return nil
}

func emojiMessageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.Name() == bruxism.DiscordServiceName && !service.IsMe(message) {
		if bruxism.MatchesCommand(service, "emoji", message) {
			_, parts := bruxism.ParseCommand(service, message)
			if len(parts) == 1 {
				s := strings.TrimSpace(parts[0])
				for i := range s {
					filename := emojiFile(s[i:])
					if filename != "" {
						if f, err := os.Open(fmt.Sprintf("emoji/twitter/%s.png", filename)); err == nil {
							defer f.Close()
							service.SendFile(message.Channel(), "emoji.png", f)

							return
						}
					}
				}
			}
		}
	}
}

func emojiHelpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "emoji", "<emoji>", "Returns a big version of an emoji.")
}

// New creates a new emoji plugin.
func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("Emoji")
	p.LoadFunc = emojiLoadFunc
	p.MessageFunc = emojiMessageFunc
	p.HelpFunc = emojiHelpFunc
	return p
}
