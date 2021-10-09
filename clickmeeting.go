package clickmeeting

type Client interface {
	ListRooms(status RoomStatus) ([]Room, error)
	CreateRoom(room NewRoom, opts ...CreateRoomOption) (Room, error)
	UpdateRoom(roomID int, opts ...UpdateRoomOption) (Room, error)
	DeleteRoom(roomID int) error

	GetSessions(roomID int) ([]SessionSummary, error)
	GetSession(roomID int, sessionID int) (Session, error)

	GenerateAccessTokens(roomID int, howMany int) ([]AccessToken, error)
	GetAccessTokens(roomID int) ([]AccessToken, error)

	AutoLoginHash(roomID int) (string, error)

	SendInvitation(roomID int, language string, attendees []string, opts ...SendInvitationOption) error

	GetRegistrations(roomID int, status string) ([]Participant, error)
	RegisterParticipant(roomID int, participant NewParticipant, opts ...RegisterParticipantOption) (string, error)
	GetParticipants(roomID int, sessionID int) ([]Participant, error)
}
