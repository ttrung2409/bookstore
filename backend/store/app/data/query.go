package data

type Query interface {
	Select(fields ...string)
	Include(ref string) Query
	IncludeMany(ref string) Query
	ThenInclude(ref string) Query
	Where(condition string, args ...interface{}) Query
	OrderBy(field string) Query
	OrderByDesc(field string) Query
	Find() ([]interface{}, error)
	First() (interface{}, error)
}
