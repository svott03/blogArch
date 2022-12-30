package responses

type MainResponse struct {
	Output string
}

type ProfileResponse struct {
	UserID string
	Entries []string
}

type EntryResponse struct {
	Status string
}

type LoginResponse struct {
	User string
	Password string
}

type RegisterResponse struct {
	User string
	Password string
}