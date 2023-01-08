package errorx

type Error struct {
	Message    string      `json:"message"`
	Details    interface{} `json:"details"`
	StatusCode int         `json:"statusCode"`
	StatusText []byte      `json:"statusText"`
}
