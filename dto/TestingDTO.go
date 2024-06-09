package dto

type SuccessResponse struct {
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	TotalCount    int64       `json:"totalCount,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}