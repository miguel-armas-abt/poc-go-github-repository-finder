package response

type RepoResponseWrapper struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PushedAt    string `json:"pushed_at"`
	Url         string `json:"clone_url"`
	Watchers    int    `json:"watchers_count"`
}
