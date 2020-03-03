package team

import (
	"errors"
	"github.com/nlopes/slack"
	"os"
)

type Member struct {
	Name string
	ID   string
}

type Team struct {
	Command string
	Members []Member
}

func GetFullName(id string) string {
	api := slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	user, err := api.GetUserInfo(id)

	if err != nil {
		errors.New("Unexpected user id")
	}

	return user.Name
}

func TeamsFactory(name string) (Team, error) {
	switch name {
	case "sample":
		return SampleTeam(), nil
	default:
		return Team{}, errors.New("Invalid team name given: '" + name + "'.")
	}
}

func SampleTeam() Team {
	return Team{
		Command: "sample",
		Members: []Member{
			{
				ID:   "YAMADA",
				Name: "taro",
			},
			{
				ID:   "TANAKA",
				Name: "hanako",
			},
		},
	}
}
