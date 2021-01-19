package exceptions

// APIException API例外
type APIException struct {
	Details interface{} `json: details`
}
