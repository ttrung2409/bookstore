package repository

type BookRepository interface {
	repository
	GetByGoogleBookId(googleBookId string) (interface{}, error)
}