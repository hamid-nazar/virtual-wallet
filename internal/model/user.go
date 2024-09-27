package model

import "time"

type User struct {
	id                    string
	username              string
	name                  string
	phone                 string
	email                 string
	password              string
	firstName             string
	lastName              string
	joinedDate            time.Time
	enabled               bool
	blocked               bool
	confirmedRegistration bool
	defaultWallet         Wallet
	invitedUser           int
}

func NewUser(id, username, name, phone, email, password, firstName, lastName string, joinedDate time.Time, enabled, blocked, confirmedRegistration bool, defaultWallet Wallet, invitedUser int) *User {
	return &User{
		id:                    id,
		username:              username,
		name:                  name,
		phone:                 phone,
		email:                 email,
		password:              password,
		firstName:             firstName,
		lastName:              lastName,
		joinedDate:            joinedDate,
		enabled:               enabled,
		blocked:               blocked,
		confirmedRegistration: confirmedRegistration,
		defaultWallet:         defaultWallet,
		invitedUser:           invitedUser,
	}
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetPhone() string {
	return u.phone
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetJoinedDate() time.Time {
	return u.joinedDate
}

func (u *User) GetEnabled() bool {
	return u.enabled
}

func (u *User) GetBlocked() bool {
	return u.blocked
}

func (u *User) GetConfirmedRegistration() bool {
	return u.confirmedRegistration
}

func (u *User) GetDefaultWallet() Wallet {
	return u.defaultWallet
}

func (u *User) GetInvitedUser() int {
	return u.invitedUser
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) SetJoinedDate(joinedDate time.Time) {
	u.joinedDate = joinedDate
}

func (u *User) SetEnabled(enabled bool) {
	u.enabled = enabled
}

func (u *User) SetBlocked(blocked bool) {
	u.blocked = blocked
}

func (u *User) SetConfirmedRegistration(confirmedRegistration bool) {
	u.confirmedRegistration = confirmedRegistration
}

func (u *User) SetDefaultWallet(defaultWallet Wallet) {
	u.defaultWallet = defaultWallet
}

func (u *User) SetInvitedUser(invitedUser int) {
	u.invitedUser = invitedUser
}
