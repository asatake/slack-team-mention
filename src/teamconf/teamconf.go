package teamconf

import (
	"errors"
	"github.com/nlopes/slack"
	"os"
)

type Team struct {
	TeamName string   `yaml:"team_name"`
	Members  []string `yaml:"member"`
}

func GetFullName(id string) (string, error) {
	api := slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	user, err := api.GetUserInfo(id)

	if err != nil {
		return "", errors.New("Unexpected user id")
	}

	return user.Name, nil
}

func SearchTeam(name string, teams []Team) (Team, error) {
	for _, team := range teams {
		if team.TeamName == name {
			return team, nil
		}
	}

	return Team{}, errors.New("Team not found.")
}
