package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Topic string

const (
	TopicBroadcast    Topic = "BROADCAST"
	TopicChange       Topic = "CHANGE"
	TopicStatusChange Topic = "STATUS_CHANGE"
	TopicProgress     Topic = "PROGRESS"
	TopicFileUpload   Topic = "FILE_UPLOAD"
	TopicAssigned     Topic = "ASSIGNED"
	TopicCreated      Topic = "CREATED"
	TopicUpdated      Topic = "UPDATED"
	TopicNewMessage   Topic = "NEW_MESSAGE"
)

var AllChannelTopics = []Topic{
	TopicBroadcast,
	TopicChange,
	TopicStatusChange,
	TopicProgress,
	TopicFileUpload,
	TopicAssigned,
	TopicCreated,
	TopicUpdated,
	TopicNewMessage,
}

func (s Topic) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (Topic) Values() (kinds []string) {
	for _, s := range AllChannelTopics {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Topic) IsValid() bool {
	for _, t := range AllChannelTopics {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Topic) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Topic(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s Topic) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
