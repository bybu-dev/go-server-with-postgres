package utils

type _EmailSender struct {
	SendVerification func(code string) error
}

var EmailSender = _EmailSender{
	SendVerification: func(code string) error {
		return nil;
	},
}