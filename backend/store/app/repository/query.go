package repository

type Query interface {
	Select(fields ...string) Query
	Include(ref string) Query
	IncludeMany(ref string) Query
	ThenInclude(ref string) Query
	Where(field string) QueryCondition
	Or(field string) QueryCondition
	OrderBy(field string) Query
	OrderByDesc(field string) Query
	Find() ([]interface{}, error)
	First() (interface{}, error)
}

type QueryCondition interface {
	In(values interface{}) Query
	Eq(value interface{}) Query
	Contains(value string) Query
}

type QueryFactory interface {
	New(model interface{}) Query
}
