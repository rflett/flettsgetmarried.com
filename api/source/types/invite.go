package types

// Invite represents the mapping of a Guest to a single Invite
type Invite struct {
	PK      string `json:"-" dynamodbav:"PK"`
	SK      string `json:"-" dynamodbav:"SK"`
	ID      string `json:"id"`
	GuestID string `json:"guestId"`
}
