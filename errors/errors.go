package errors

const (
	ErrInvalidRole               string = "invalid role"
	ErrCannotOpenFile            string = "cannot open file"
	ErrCannotBeSenderAndReceiver string = "sender and receiver are mutually exclusive\n"
	ErrInvalidContentType        string = "invalid Content Type: %s\n"
	ErrUnmarshalJSON             string = "error: %s\n"
	ErrError                     string = "error:"
)
