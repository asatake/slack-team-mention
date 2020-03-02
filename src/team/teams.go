package team

import (
	"github.com/nlopes/slack"
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
	var user slack.User = api.GetUserInfo(id)
	return user.Name
}

func SampleTeam() Team {
	return team.Team{
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
