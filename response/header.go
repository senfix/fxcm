package response

type Header struct {
	Response struct {
		Executed bool   `json:"executed"`
		Error    string `json:"error"`
	} `json:"response"`
}
