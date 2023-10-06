package app_errors

type AppError string

const (
	ErrBadRequest    = AppError("Некорректный запрос")
	ErrNotFound      = AppError("Не найдено")
	ErrInternal      = AppError("Внутренняя ошибка сервера")
	ErrNoId          = AppError("Не указан id")
	ErrWrongIdFormat = AppError("Некорректный id")
)
