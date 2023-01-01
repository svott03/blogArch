package responses

type ProfileResponse struct {
	UserID string
	Entries []string
}

type StatusResponse struct {
	Status string
}