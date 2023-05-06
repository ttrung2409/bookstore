package repository

import "store/app/domain"

type Query[M domain.DataObject] interface {
	Select(fields ...string) Query[M]
	Include(ref string) Query[M]
	IncludeMany(ref string) Query[M]
	ThenInclude(ref string) Query[M]
	Where(field string) Where[M]
	Or(field string) Where[M]
	OrderBy(field string) Query[M]
	OrderByDesc(field string) Query[M]
	Find() ([]M, error)
	First() (M, error)
}

type Where[M domain.DataObject] interface {
	In(values any) Query[M]
	Eq(value any) Query[M]
	Contains(value string) Query[M]
}

type QueryFactory[M domain.DataObject] interface {
	New() Query[M]
}
