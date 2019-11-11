package response

type Header struct {
	Response Stats `json:"response"`
}

type Stats struct {
	Executed bool   `json:"executed"`
	Error    string `json:"error"`
}
