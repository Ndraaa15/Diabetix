package errx

type Errx struct {
	Code    int
	Message string
	Err     error
}

func New() *Errx {
	return &Errx{}
}

func (e *Errx) WithCode(code int) *Errx {
	e.Code = code
	return e
}

func (e *Errx) WithMessage(message string) *Errx {
	e.Message = message
	return e
}

func (e *Errx) WithError(err error) *Errx {
	e.Err = err
	return e
}

func (e *Errx) Error() string {
	return e.Message
}
