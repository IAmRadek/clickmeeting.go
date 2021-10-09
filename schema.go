package clickmeeting

import "time"

type NewRoom struct {
	//Name of the room that will be visible to attendees. This name will be part of your meeting room url.
	Name     string
	RoomType RoomType
	//PermanentRoom determines whether you create a one-time scheduled meeting or a permanent (endless) conference room.
	//	false	one-time scheduled meeting.
	//	true	permanent event.
	PermanentRoom bool
	AccessType    AccessType
}

type Room struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"name_url"`
	RoomURL      string `json:"room_url"`
	EmbedRoomURL string `json:"embed_room_url"`

	Status     RoomStatus `json:"status"`
	AccessType AccessType `json:"access_type"`
	RoomType   RoomType   `json:"room_type"`

	UpdatedAt time.Time `json:"updated_at"`

	LobbyEnabled        bool   `json:"lobby_enabled"`
	LobbyDescription    string `json:"lobby_description"`
	RegistrationEnabled int    `json:"registration_enabled"`
	PermanentRoom       bool   `json:"permanent_room"`

	RoomPin           int `json:"room_pin"`
	PhonePresenterPin int `json:"phone_presenter_pin"`
	PhoneListenerPin  int `json:"phone_listener_pin"`

	Timezone       string `json:"timezone"`
	TimezoneOffset int    `json:"timezone_offset"`

	AccessRoleHashes struct {
		Listener  string `json:"listener"`
		Presenter string `json:"presenter"`
		Host      string `json:"host"`
	} `json:"access_role_hashes"`
	CreatedAt    time.Time    `json:"created_at"`
	RecorderList []string     `json:"recorder_list"`
	WidgetsHash  string       `json:"widgets_hash"`
	Settings     RoomSettings `json:"settings"`
}

type AccessType int

func (a AccessType) String() string {
	switch a {
	case OpenType:
		return "OpenType"
	case PasswordProtected:
		return "PasswordProtected"
	case TokenProtected:
		return "TokenProtected"
	}
	return "UnknownType"
}

const (
	UnknownType AccessType = iota
	// OpenType room not protected by password or token.
	OpenType
	// PasswordProtected access to the conference room granted based on password provided in advance.
	PasswordProtected
	// TokenProtected means each invitee receives a unique token that grants access to the conference room.
	// Access token is single-use, so it can be used only once by one person.
	TokenProtected
)

type RoomType string

func (r RoomType) String() string {
	return string(r)
}

const (
	Meeting RoomType = "meeting"
	Webinar RoomType = "webinar"
)

type RoomStatus string

const (
	ActiveRoom   RoomStatus = "active"
	InactiveRoom RoomStatus = "active"
)

type RoomSettings struct {
	//ShowOnPersonalPage displays conference on personal page.
	ShowOnPersonalPage bool `json:"show_on_personal_page"`
	//ThankYouEmailsEnabled sends thank you email.
	ThankYouEmailsEnabled bool `json:"thank_you_emails_enabled"`
	//ConnectionTesterEnabled turns on connection tester.
	ConnectionTesterEnabled bool `json:"connection_tester_enabled"`
	//PhoneGatewayEnabled turns on phone gateway.
	PhoneGatewayEnabled bool `json:"phonegateway_enabled"`
	//RecorderAutostartEnabled turns on recorder autostart.
	RecorderAutostartEnabled bool `json:"recorder_autostart_enabled"`
	//RoomInviteButtonEnabled turns on invite option in conference room.
	RoomInviteButtonEnabled bool `json:"room_invite_button_enabled"`
	//SocialMediaSharingEnabled turns on social media sharing in conference room.
	SocialMediaSharingEnabled bool `json:"social_media_sharing_enabled"`
	//ConnectionStatusEnabled turns on connection status.
	ConnectionStatusEnabled bool `json:"connection_status_enabled"`
	//ThankYouPageUrl sets thank you page url.
	ThankYouPageUrl string `json:"thank_you_page_url"`
}

type AccessToken struct {
	Token        string     `json:"token"`
	SentToEmail  string     `json:"sent_to_email,omitempty"`
	FirstUseData *time.Time `json:"first_use_data,omitempty"`
}

type Participant struct {
	RegistrationDate      time.Time `json:"registration_date"`
	RegistrationConfirmed string    `json:"registration_confirmed"`
	Fields                struct {
		FirstName    string `json:"First Name"`
		LastName     string `json:"Last Name"`
		EmailAddress string `json:"Email Address"`
	} `json:"fields"`
	ID              int    `json:"id"`
	SessionID       int    `json:"session_id"`
	Email           string `json:"email"`
	VisitorNickname string `json:"visitor_nickname"`
}

type NewParticipant struct {
	FirstName    string
	LastName     string
	EmailAddress string
}

type SessionSummary struct {
	ID            int       `json:"id"`
	TotalVisitors int       `json:"total_visitors"`
	MaxVisitors   int       `json:"max_visitors"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

type Session struct {
	MaxVisitors   int                   `json:"max_visitors"`
	Attendees     []Attendee            `json:"attendees"`
	StartDate     time.Time             `json:"start_date"`
	PDF           map[string]PDFSummary `json:"pdf"`
	EndDate       time.Time             `json:"end_date"`
	TotalVisitors int                   `json:"total_visitors"`
}
type Attendee struct {
	Id        int       `json:"id"`
	StartDate time.Time `json:"start_date"`
	Email     string    `json:"email"`
	EndDate   time.Time `json:"end_date"`
	Login     string    `json:"login"`
}

type PDFSummary struct {
	URL      string `json:"generate_pdf_url"`
	Progress int    `json:"progress"`
}
