package params

type PageParams struct {
	Limit  *int32 `page:"limit,default=10"`
	Offset *int32 `page:"offset"`
	Total  *int32
}
