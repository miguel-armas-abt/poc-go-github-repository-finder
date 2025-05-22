package response

type RepoResponseDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PushedAt    string `json:"pushedAt"`
	Url         string `json:"url"`
	Watchers    int    `json:"watchersCount"`
	ImageUrl    string `json:"imageUrl"`
}
