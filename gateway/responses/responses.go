package responses

type MainResponse struct {
	Output string
}

type ProfileResponse struct {
	UserID string
	Entries []string
}

type StatusResponse struct {
	Status string
}