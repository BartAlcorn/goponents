package todos

type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}
