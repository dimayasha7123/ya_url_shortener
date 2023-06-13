package handlers

type redirectResp struct {
	URL string `json:"orig_url"`
}

type shortenBody struct {
	URL string `json:"url"`
}

type shortenResp struct {
	URL string `json:"shorten_url"`
}

type errResp struct {
	Error string `json:"error"`
}
