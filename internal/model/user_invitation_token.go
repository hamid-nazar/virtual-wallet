package model

import (
	"time"
)

type UserInvitationToken struct {
	id           string
	token        string
	owner        User
	createdAt    time.Time
	expiry       time.Time
	invitedEmail string
	used         bool
}

func NewUserInvitationToken(id, token string, owner User, invitedEmail string) *UserInvitationToken {
	return &UserInvitationToken{
		id:           id,
		token:        token,
		owner:        owner,
		createdAt:    time.Now(),
		expiry:       time.Now().Add(5 * time.Minute),
		invitedEmail: invitedEmail,
		used:         false,
	}
}
func (t *UserInvitationToken) GetId() string {
	return t.id
}
func (t *UserInvitationToken) GetToken() string {
	return t.token
}
func (t *UserInvitationToken) GetOwner() User {
	return t.owner
}
func (t *UserInvitationToken) GetCreatedAt() time.Time {
	return t.createdAt
}
func (t *UserInvitationToken) GetExpiry() time.Time {
	return t.expiry
}
func (t *UserInvitationToken) GetInvitedEmail() string {
	return t.invitedEmail
}
func (t *UserInvitationToken) GetUsed() bool {
	return t.used
}
func (t *UserInvitationToken) SetId(id string) {
	t.id = id
}
func (t *UserInvitationToken) SetToken(token string) {
	t.token = token
}
func (t *UserInvitationToken) SetOwner(owner User) {
	t.owner = owner
}
func (t *UserInvitationToken) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}
func (t *UserInvitationToken) SetExpiry(expiry time.Time) {
	t.expiry = expiry
}
func (t *UserInvitationToken) SetInvitedEmail(invitedEmail string) {
	t.invitedEmail = invitedEmail
}
func (t *UserInvitationToken) SetUsed(used bool) {
	t.used = used
}
