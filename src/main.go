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

func GenMessage(team team.Team, txt string) string {
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
	if unserializeErr := json.Unmarshal([]byte(u.Encode()), &slashCommand); unserializeErr != nil {
		fmt.Println(unserializeErr.Error())
		return events.APIGatewayProxyResponse{Body: "decode error: " + err.Error(), StatusCode: 500}, nil
	}

	var teams team.Team
	team.TeamsFactory(slashCommand.Command)

	postMessageParams := slack.PostMessageParameters{
		AsUser:    true,
		Username:  slashCommand.UserName,
		LinkNames: 1,
		Markdown:  true,
	}

	msg := GenMessage(teams, slashCommand.Text)

	option1 := slack.MsgOptionPostMessageParameters(postMessageParams)
	option2 := slack.MsgOptionText(msg, false)

	if _, _, err := api.PostMessage(slashCommand.ChannelID, option1, option2); err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "decode error: " + err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
