package clickmeeting_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/IAmRadek/clickmeeting.go"
	"github.com/matryer/is"
)

func Test_RoomAPI(t *testing.T) {
	t.SkipNow()

	api := clickmeeting.NewAPI("eu97744d4ec7fd54ba14276f5cc99c59a31097473b")

	var roomID int
	t.Run("CreateRoom", func(t *testing.T) {
		is := is.New(t)
		room, err := api.CreateRoom(
			clickmeeting.NewRoom{
				Name:          "My new sweet room",
				RoomType:      clickmeeting.Webinar,
				PermanentRoom: true,
				AccessType:    clickmeeting.PasswordProtected,
			},
			clickmeeting.WithPassword("test"),
			clickmeeting.WithDuration(3*time.Hour),
			clickmeeting.WithLobby(true, "Testing lobby message"),
			clickmeeting.WithRoomSettings(clickmeeting.RoomSettings{
				ShowOnPersonalPage:        true,
				ThankYouEmailsEnabled:     true,
				ConnectionTesterEnabled:   true,
				PhoneGatewayEnabled:       true,
				RecorderAutostartEnabled:  true,
				RoomInviteButtonEnabled:   true,
				SocialMediaSharingEnabled: true,
				ConnectionStatusEnabled:   true,
				ThankYouPageUrl:           "https://adasd.pl/asdasd",
			}),
		)
		is.NoErr(err)

		is.Equal(room.Name, "My new sweet room")
		is.Equal(room.RoomType, clickmeeting.Webinar)
		is.Equal(room.PermanentRoom, true)
		is.Equal(room.AccessType, clickmeeting.PasswordProtected)
		is.Equal(room.LobbyEnabled, true)
		is.Equal(room.LobbyDescription, "Testing lobby message")
		is.Equal(room.Settings.ThankYouPageUrl, "https://adasd.pl/asdasd")

		roomID = room.ID
	})

	t.Run("UpdateRoom", func(t *testing.T) {
		is := is.New(t)
		room, err := api.UpdateRoom(roomID,
			clickmeeting.SetName("Webinarium"),
			clickmeeting.SetLobby(false, ""),
			clickmeeting.SetDuration(4*time.Hour),
			clickmeeting.SetPermanence(false),
			clickmeeting.SetStartsAt(time.Now().Add(24*time.Hour)),
			clickmeeting.SetRoomType(clickmeeting.Webinar),
			clickmeeting.SetPassword("qwesdwdrty"),
			//clickmeeting.SetAccessType(clickmeeting.TokenProtected)
			clickmeeting.SetStatus(clickmeeting.InactiveRoom),
		)
		is.NoErr(err)
		is.Equal(room.Name, "Webinarium")
		is.Equal(room.LobbyEnabled, false)
	})

	t.Run("DeleteRoom", func(t *testing.T) {
		is := is.New(t)
		err := api.DeleteRoom(roomID)
		is.NoErr(err)
	})

}
func Test_Participants(t *testing.T) {
	t.SkipNow()

	api := clickmeeting.NewAPI("eu97744d4ec7fd54ba14276f5cc99c59a31097473b")

	var roomID int

	t.Run("CreateRoom", func(t *testing.T) {
		is := is.New(t)

		room, err := api.CreateRoom(clickmeeting.NewRoom{
			Name:          "Testing",
			RoomType:      clickmeeting.Webinar,
			PermanentRoom: false,
			AccessType:    clickmeeting.TokenProtected,
		}, clickmeeting.WithRegistration())
		is.NoErr(err)
		roomID = room.ID
	})
	t.Cleanup(func() {
		if roomID != 0 {
			api.DeleteRoom(roomID)
		}
	})

	t.Run("RegisterParticipant", func(t *testing.T) {
		is := is.New(t)

		attendURL, err := api.RegisterParticipant(roomID, clickmeeting.NewParticipant{
			FirstName:    "Jon",
			LastName:     "Doe",
			EmailAddress: "jon@doe.com",
		}, clickmeeting.WithEmailConfirmation("pl"))
		is.NoErr(err)
		is.True(attendURL != "")
	})

	t.Run("ListParticipants", func(t *testing.T) {
		is := is.New(t)

		people, err := api.GetRegistrations(roomID, "all")
		is.NoErr(err)
		is.Equal(len(people), 1)
		is.Equal(people[0].Email, "jon@doe.com")
	})
}

func Test_Invitations(t *testing.T) {
	t.SkipNow()
	api := clickmeeting.NewAPI("eu97744d4ec7fd54ba14276f5cc99c59a31097473b")

	var roomID int

	t.Run("CreateRoom", func(t *testing.T) {
		is := is.New(t)

		room, err := api.CreateRoom(clickmeeting.NewRoom{
			Name:          "Testing",
			RoomType:      clickmeeting.Webinar,
			PermanentRoom: false,
			AccessType:    clickmeeting.TokenProtected,
		}, clickmeeting.WithRegistration())
		is.NoErr(err)
		roomID = room.ID
	})
	t.Cleanup(func() {
		if roomID != 0 {
			api.DeleteRoom(roomID)
		}
	})

	t.Run("SendInvitation", func(t *testing.T) {
		is := is.New(t)

		err := api.SendInvitation(roomID, "pl", []string{"jon@doe.com"}, clickmeeting.SetRole(clickmeeting.AsListener))
		is.NoErr(err)
	})

	t.Run("ListParticipants", func(t *testing.T) {
		is := is.New(t)

		people, err := api.GetRegistrations(roomID, "all")
		is.NoErr(err)
		is.Equal(len(people), 1)
		is.Equal(people[0].Email, "jon@doe.com")
	})
}

func pp(d interface{}) {
	r, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(r))
}
