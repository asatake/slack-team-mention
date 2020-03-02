package main

import (
	"./team"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"net/url"
	"os"
)

func GenMessage(team team.Team, txt string) {
	msg := ""

	for _, t := range team.Members {
		msg += "<@" + t.Name + "> "
	}
	return msg + txt
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	u, err := url.ParseQuery(request.Body)
	if err != nil {
		fmt.Println("url decode error: " + err.Error())
		return events.APIGatewayProxyResponse{Body: "decode error: " + err.Error(), StatusCode: 500}, nil
	}

	api := slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))

	var slashCommand slack.SlashCommand
	err := json.Unmarshal([]byte(u.Encode()), &slashCommand)

	var teams team.Team
	switch request.Path {
	case "/sample":
		teams = team.SampleTeam()
		msg := slack.Message{
			Msg: slack.Msg,
		}

		api.PostMessage(slashCommand.ChannelID)
	}
}
