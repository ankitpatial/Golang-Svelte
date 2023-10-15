package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Channel string

const (
	ChannelPing            Channel = "PING"
	ChannelEstimate        Channel = "ESTIMATE"
	ChannelJob             Channel = "JOB"
	ChannelJobChat         Channel = "JOB_CHAT"
	ChannelJobNote         Channel = "JOB_NOTE"
	ChannelTask            Channel = "TASK"
	ChannelTrainingVideo   Channel = "TRAINING_VIDEO"
	ChannelPartner         Channel = "PARTNER"
	ChannelPartnerUser     Channel = "PARTNER_USER"
	ChannelInstallationJob Channel = "INSTALLATION_JOB"
)

var AllChannels = []Channel{
	ChannelPing,
	ChannelEstimate,
	ChannelJob,
	ChannelJobChat,
	ChannelTask,
	ChannelTrainingVideo,
	ChannelPartner,
	ChannelPartnerUser,
}

func (s Channel) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (Channel) Values() (kinds []string) {
	for _, s := range AllChannels {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Channel) IsValid() bool {
	for _, t := range AllChannels {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Channel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Channel(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s Channel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
