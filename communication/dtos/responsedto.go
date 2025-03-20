package communication

type ResponseDTO struct {
	Message    string
	StatusCode int
	Data       any
}
