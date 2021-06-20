package internal

type StringError string

func (s StringError) Error() string {
	return string(s)
}

const NotFound StringError = "Not found"
