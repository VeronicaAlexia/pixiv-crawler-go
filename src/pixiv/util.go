package pixiv

func NextUrl(next_url, path string, params map[string]string) (string, map[string]string) {
	if next_url == "" {
		next_url = API_BASE + path
	} else {
		params = nil
	}
	return next_url, params
}
