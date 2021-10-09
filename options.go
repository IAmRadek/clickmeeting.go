package clickmeeting

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type option func(values url.Values)

type CreateRoomOption option

//WithRoomSettings sets various settings of a conference.
func WithRoomSettings(s RoomSettings) CreateRoomOption {
	boolToString := func(b bool) string {
		if b {
			return "1"
		}
		return "0"
	}
	return func(v url.Values) {
		v.Add("settings[show_on_personal_page]", boolToString(s.ShowOnPersonalPage))
		v.Add("settings[thank_you_emails_enabled]", boolToString(s.ThankYouEmailsEnabled))
		v.Add("settings[connection_tester_enabled]", boolToString(s.ConnectionTesterEnabled))
		v.Add("settings[phonegateway_enabled]", boolToString(s.PhoneGatewayEnabled))
		v.Add("settings[recorder_autostart_enabled]", boolToString(s.RecorderAutostartEnabled))
		v.Add("settings[room_invite_button_enabled]", boolToString(s.RoomInviteButtonEnabled))
		v.Add("settings[social_media_sharing_enabled]", boolToString(s.SocialMediaSharingEnabled))
		v.Add("settings[connection_status_enabled]", boolToString(s.ConnectionStatusEnabled))
		v.Add("settings[thank_you_page_url]", s.ThankYouPageUrl)
	}
}

//WithLobby enabled lobby and sets description.
func WithLobby(enabled bool, description string) CreateRoomOption {
	return func(v url.Values) {
		if enabled {
			v.Add("lobby_enabled", "1")
		} else {
			v.Add("lobby_enabled", "0")
		}
		if description != "" {
			v.Add("lobby_description", description)
		}
	}
}

//WithRegistration enables registration.
func WithRegistration() CreateRoomOption {
	return func(v url.Values) {
		v.Add("registration[enabled]", "1")
	}
}

//WithRegistrationAndTemplate enables registration and sets meeting registration template.
// Valid template values: 1 - 3
func WithRegistrationAndTemplate(template int) CreateRoomOption {
	return func(v url.Values) {
		WithRegistration()(v)
		v.Add("registration[template]", strconv.Itoa(template))
	}
}

//WithDuration sets duration of a conference.
func WithDuration(d time.Duration) CreateRoomOption {
	return func(v url.Values) {
		v.Add("duration", fmt.Sprintf("%d:%d", int(d.Hours()), int(d.Minutes())%60))
	}
}

//WithPassword sets password of a conference.
func WithPassword(password string) CreateRoomOption {
	return func(v url.Values) {
		v.Add("password", password)
	}
}

type UpdateRoomOption option

func SetName(name string) UpdateRoomOption {
	return func(v url.Values) {
		v.Add("name", name)
	}
}
func SetRoomType(roomType RoomType) UpdateRoomOption {
	return func(v url.Values) {
		v.Add("room_type", roomType.String())
	}
}
func SetPermanence(p bool) UpdateRoomOption {
	return func(v url.Values) {
		if p {
			v.Add("permanent_room", "1")
		} else {
			v.Add("permanent_room", "0")
		}
	}
}
func SetAccessType(accessType AccessType) UpdateRoomOption {
	return func(v url.Values) {
		v.Add("access_type", strconv.Itoa(int(accessType)))
	}
}
func SetLobby(enabled bool, description string) UpdateRoomOption {
	return func(v url.Values) {
		WithLobby(enabled, description)(v)
	}
}
func SetDuration(d time.Duration) UpdateRoomOption {
	return func(v url.Values) {
		WithDuration(d)(v)
	}
}
func SetStartsAt(t time.Time) UpdateRoomOption {
	return func(v url.Values) {
		v.Add("starts_at", t.Format(time.RFC3339))
	}
}

func SetPassword(password string) UpdateRoomOption {
	return func(v url.Values) {
		SetAccessType(PasswordProtected)(v)
		WithPassword(password)(v)
	}
}
func SetStatus(status RoomStatus) UpdateRoomOption {
	return func(v url.Values) {
		v.Add("status", string(status))
	}
}
func SetRoomSettings(settings RoomSettings) UpdateRoomOption {
	return func(v url.Values) {
		WithRoomSettings(settings)(v)
	}
}

type SendInvitationOption option

type TemplateType string

const (
	AdvancedTemplate TemplateType = "advanced"
	BasicTemplate    TemplateType = "basic"
)

func SetTemplate(tp TemplateType) SendInvitationOption {
	return func(v url.Values) {
		v.Add("template", string(tp))
	}
}

type InviteeRole string

const (
	AsListener  InviteeRole = "listener"
	AsPresenter InviteeRole = "presenter"
)

func SetRole(role InviteeRole) SendInvitationOption {
	return func(v url.Values) {
		v.Add("role", string(role))
	}
}

type RegisterParticipantOption option

func WithEmailConfirmation(language string) RegisterParticipantOption {
	return func(values url.Values) {
		values.Add("confirmation_email[enabled]", "1")
		values.Add("confirmation_email[lang]", language)
	}
}
