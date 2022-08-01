package urlvals

type PageParams struct {
	Limit  *int32 `page:"limit"`
	Offset *int32 `page:"offset"`
	Size   *int32 `page:"size,default=10"`
}
