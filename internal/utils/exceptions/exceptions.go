package exceptions

type Exception struct {
	Description string
	StatusCode  int
}

func New(description string, status int) *Exception {
	return &Exception{
		Description: description,
		StatusCode:  status,
	}
}

func (e *Exception) Error() string {
	return e.Description
}
