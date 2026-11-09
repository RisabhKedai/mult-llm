package chat

type Model struct {
	base    string `json:"base"`
	version string `json:"version"`
}

type Msg struct {
	msg     string `json:"msg"`
	session string `json:"session"`
	model   Model  `json:"model"`
}
