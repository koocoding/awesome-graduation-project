package model

type UserCreateInput struct {
	Passport    string
	Password    string
	Nickname    string
	PhoneNumber string
	VerifyCode  string
}

type UserSignInInput struct {
	Passport    string
	PhoneNumber string
	Password    string
}
type AdminSignUp struct {
	Passport string
	Password string
	Nickname string
}
