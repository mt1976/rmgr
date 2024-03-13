package errors

const (
	ErrInvalidRole               string = "invalid role"
	ErrCannotOpenFile            string = "cannot open file"
	ErrCannotBeSenderAndReceiver string = "sender and receiver are mutually exclusive"
	ErrInvalidContentType        string = "invalid Content Type: %s\n"
	ErrUnmarshalJSON             string = "ERROR: %s\n"
	ErrError                     string = "ERROR:"
)
