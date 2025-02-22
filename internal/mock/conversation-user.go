package mock

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func createUserMessage(createTime int64, input string) (conversation datatypes.ConversationHistory) {
	conversation.ID = define.GenerateUUID()
	conversation.Parent = ""
	conversation.Children = []string{}

	var message datatypes.ConversationMessage
	message.ID = conversation.ID
	message.Author = CreateMessageUser("user")
	message.CreateTime = float64(createTime)
	message.Content = CreateMessageContent("text", input)

	message.Metadata = createTimestampMessageMeta("absolute")
	// message.EndTurn = true
	message.Weight = 1.0
	message.Recipient = "all"

	conversation.Message = message
	return conversation
}
