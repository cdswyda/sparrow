package StreamResponser

import (
	"time"

	eb "github.com/soulteary/sparrow/components/event-broker"
	OpenaiAPI "github.com/soulteary/sparrow/connectors/openai-api"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

type StreamMessageMode int

var (
	MSG_STATUS_AUTO_MODE StreamMessageMode = 0
	MSG_STATUS_CONTINUE  StreamMessageMode = 1
	MSG_STATUS_DONE      StreamMessageMode = 2
)

func StreamBuilder(parentMessageID string, conversationID string, modelSlug string, broker *eb.Broker, input string, mode StreamMessageMode) bool {
	messageID, modelSlug := GetBuilderParams(modelSlug)
	var sequences []string

	switch modelSlug {
	case datatypes.MODEL_OPENAI_API_3_5.Slug:
	case datatypes.MODEL_TEXT_DAVINCI_002_PLUGINS.Slug:
	case datatypes.MODEL_TEXT_DAVINCI_002_RENDER_PAID.Slug:
	case datatypes.MODEL_TEXT_DAVINCI_002_RENDER_SHA.Slug:
	case datatypes.MODEL_GPT4.Slug:
		if define.ENABLE_OPENAI_API {
			sequences = MakeStreamingMessage(OpenaiAPI.Get(input), modelSlug, conversationID, messageID, mode)
		}
	case datatypes.MODEL_MIDJOURNEY.Slug:
		if define.ENABLE_MIDJOURNEY {
			sequences = MakeStreamingMessage(input, modelSlug, conversationID, messageID, mode)
		}
		// case datatypes.MODEL_NO_MODELS.Slug:
		// default:
	}

	if len(sequences) > 0 {
		return MakeStreamingResponse(parentMessageID, broker, sequences)
	}

	sequences = MakeStreamingMessage("The administrator has disabled the export capability of this model.\nProject: [soulteary/sparrow](https://github.com/soulteary/sparrow).\nTalk is Cheap, Let's coding together.", modelSlug, conversationID, messageID, mode)
	return MakeStreamingResponse(parentMessageID, broker, sequences)
}

func GetBuilderParams(modelSlug string) (string, string) {
	messageID := define.GenerateUUID()
	if modelSlug == "" {
		modelSlug = "text-davinci-002-render-sha"
	}
	return messageID, modelSlug
}

func MakeStreamingResponse(parentMessageID string, broker *eb.Broker, sequences []string) bool {
	count := len(sequences)
	if count == 0 {
		return false
	}

	simulateDelay := 800 / define.RESPONSE_SPEED
	if define.DEV_MODE {
		simulateDelay = 10
	}

	go func() {
		lastThreeBefore := count - 3
		for id, sequence := range sequences {
			if id <= 2 {
				time.Sleep(time.Millisecond * time.Duration(simulateDelay))
			}

			broker.Event <- eb.Event{
				Name:    parentMessageID,
				Payload: sequence,
			}

			if id < lastThreeBefore {
				time.Sleep(time.Millisecond * time.Duration(RandomResponseTime(40, 120)))
			} else {
				// Acceleration end output
				time.Sleep(time.Millisecond * time.Duration(50))
			}
		}
	}()
	return true
}
