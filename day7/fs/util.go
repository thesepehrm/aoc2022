package fs

func generateTab(count int) string {
	o := ""
	for i := 0; i < count; i++ {
		o += "  "
	}
	return o
}
