package responses

// JSONResponse レスポンス
type JSONResponse struct {
	Data interface{} `json: "data"`
}
