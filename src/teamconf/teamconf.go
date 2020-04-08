package teamconf

import (
	"errors"
)

type Team struct {
	TeamName string   `yaml:"team_name"`
	Members  []string `yaml:"member"`
}

func SearchTeam(name string, teams []Team) (Team, error) {
	for _, team := range teams {
		if team.TeamName == name {
			return team, nil
		}
	}

	return Team{}, errors.New("Team not found.")
}
