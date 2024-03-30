package history

type AuthoredMessage struct {
	Author  Author
	Content string
}

type Author string

const AuthorClaude Author = "claude"
const AuthorOpenAI Author = "openai"

type MessageHistory []AuthoredMessage

func NewMessageHistory() MessageHistory {
	return make([]AuthoredMessage, 0)
}

func (mh *MessageHistory) Add(a Author, c string) {
	*mh = append(*mh, AuthoredMessage{Author: a, Content: c})
}
