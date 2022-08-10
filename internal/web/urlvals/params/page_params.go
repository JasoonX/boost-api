package params

type PageParams struct {
	Limit  *int64 `page:"limit,default=10"`
	Offset *int64 `page:"offset"`
	Total  *int64
}
