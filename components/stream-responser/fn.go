package StreamResponser

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"regexp"

	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func ParseConversationBody(r io.Reader) (originBody []byte, result datatypes.Conversation, err error) {
	originBody, err = io.ReadAll(r)
	if err != nil {
		return nil, result, err
	}
	result, err = PromptSerialization(originBody)
	return originBody, result, err
}

func PromptSerialization(buf []byte) (datatypes.Conversation, error) {
	var conversation datatypes.Conversation
	err := json.Unmarshal(buf, &conversation)
	if err != nil {
		return conversation, err
	}
	return conversation, nil
}

func RandomResponseTime(min, max int) int {
	const limit = 10
	if define.DEV_MODE {
		return limit
	}

	const defaultMin = 40
	const defaultMax = 120
	if min > max || min <= 0 || max <= 0 {
		min = defaultMin
		max = defaultMax
	}

	var i int
	i = rand.Intn(max-min+1) + min
	if i >= (int(float64(max) * 0.9)) {
		// 10% chance to get a longer delay, If it was originally a long delay
		i = rand.Intn(500-300+1) + 300
	}

	if define.RESPONSE_SPEED <= 0 {
		return i
	}
	// set speed
	delay := i / define.RESPONSE_SPEED
	if delay < limit {
		return limit
	}

	return delay
}

func ContainMarkdownImage(str string) bool {
	var re = regexp.MustCompile(`!\[.*?\]\(.*?\)`)
	return len(re.FindAllString(str, -1)) > 0
}

func ContainMarkdownLink(str string) bool {
	var re = regexp.MustCompile(`\[.*?\]\(.*?\)`)
	return len(re.FindAllString(str, -1)) > 0
}

func MakeJSON(data any) (string, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(ret), nil
}
