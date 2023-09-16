package repository

type Query[TParams any, TResult any] interface {
	Execute(args TParams) (TResult, error)
}
