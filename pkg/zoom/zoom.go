package zoom

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
	"github.com/zoom-lib-golang/zoom-lib-golang"
)

type Zoom struct {
	APIKey    string
	APISecret string
	UserID    string
}

func NewZoom(env *env.Env) *Zoom {
	return &Zoom{}
}

func (z *Zoom) CreateMeeting(ctx context.Context) (string, error) {
	createMeetingOpt := zoom.CreateMeetingOptions{
		Duration: 3600,
		HostID:   z.UserID,
		Topic:    "Consultation",
	}

	meeting, err := zoom.CreateMeeting(createMeetingOpt)
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to create zoom meeting").
			WithError(err)

	}

	return meeting.JoinURL, nil
}
