package utils

import (
	"context"
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/ent/message"
	"discord-metrics-server/v2/ent/user"
	"fmt"
	"regexp"
)

// GetMentionedUserIDs Returns a list of discord User IDs that are
// mentioned in the message
func GetMentionedUserIDs(messageContents string) []string {
	r, _ := regexp.Compile("<@(\\d+)>")
	var IDs []string
	results := r.FindAllStringSubmatch(messageContents, -1)
	for _, ID := range results {
		IDs = append(IDs, ID[1])
	}
	return IDs
}

// GetUserMentions Adds mentions of users to database
func GetUserMentions(messageID string) {
	client := db.GetClient()

	messageObject, err := client.
		Message.
		Query().
		Where(message.MessageID(messageID)).
		WithMentions().
		Only(context.Background())

	if err != nil {
		fmt.Println("Unable to get message by ID", messageID)
		return
	}

	// Delete existing mentions, to prevent duplicates
	for _, userObject := range messageObject.Edges.Mentions {
		messageObject.Update().
			RemoveMentions(userObject).
			SaveX(context.Background())
	}

	// Get all mentioned user IDs in the message
	mentionedIDs := GetMentionedUserIDs(messageObject.Contents)

	for _, ID := range mentionedIDs {
		userObject, err := client.User.
			Query().
			Where(user.UserID(ID)).
			Only(context.Background())

		if err != nil {
			fmt.Println("Error getting user")
			fmt.Println(err.Error())
		}

		_, err = userObject.Update().
			AddMentionedIn(messageObject).
			Save(context.Background())
		if err != nil {
			fmt.Println("Error saving user message link")
		}
	}

}
