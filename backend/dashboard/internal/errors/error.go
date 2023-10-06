package app_errors

type AppError struct {
	Message   string         `json:"message"`
	UIMessage string         `json:"ui_message"`
	Options   map[string]any `json:"options"`
}

func New(err error, uiMessage string, options map[string]any) AppError {
	a := AppError{
		Message:   err.Error(),
		UIMessage: uiMessage,
		Options:   options,
	}
	return a
}

func (a *AppError) Error() string {
	return a.Message
}

const (
	ErrInternal      = "внутренняя ошибка сервер"
	ErrNotFound      = "не найдено"
	ErrBadRequest    = "неверный запрос"
	ErrWrongIdFormat = "неверный id"
	ErrNoId          = "нет id"
)
