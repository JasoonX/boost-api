package convert

func StringPtr(s string) *string {
	return &s
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func BoolPtr(b bool) *bool {
	return &b
}
