package data

type Query interface {
	Select(columns ...string)
	Include(relation string) Query
	ThenInclude(relation string) Query
	Where(condition string, args ...interface{}) Query
	OrderBy(column string) Query
	OrderByDesc(column string) Query
	Find() ([]interface{}, error)
	First() (interface{}, error)
}
