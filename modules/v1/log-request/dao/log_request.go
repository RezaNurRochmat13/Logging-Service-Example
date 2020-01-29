package dao

import "time"

type ListLogRequest struct {
	IdLogHttpRequest     string `bson:"_id" json:"id_log_http_request"`
	LogHttpRequestName   string `bson:"request_name" json:"log_http_name"`
	LogHttpRequestStatus string `bson:"status" json:"log_http_status"`
}

type CreateLogRequest struct {
	LogHttpRequestName   string    `bson:"request_name" json:"log_http_name"`
	LogHttpRequestStatus string    `bson:"status" json:"log_http_status"`
	LogHttpRequestUrl    string    `bson:"status" json:"log_http_url"`
	CreatedAt            time.Time `bson:"created_at"`
	UpdatedAt            time.Time `bson:"updated_at"`
}
