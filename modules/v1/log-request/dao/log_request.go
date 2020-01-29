package dao

import "time"

type ListLogRequest struct {
	IdLogHttpRequest     string `bson:"_id" json:"id_log_http_request"`
	LogHttpRequestName   string `bson:"request_name" json:"log_http_name"`
	LogHttpRequestStatus int    `bson:"status" json:"log_http_status"`
}

type DetailLogRequest struct {
	IdLogHttpRequest     string `bson:"_id" json:"id_log_http_request"`
	LogHttpRequestName   string `bson:"request_name" json:"log_http_name"`
	LogHttpRequestStatus int    `bson:"status" json:"log_http_status"`
	LogHttpRequestUrl    string `bson:"url" json:"log_http_url"`
}

type CreateLogRequest struct {
	LogHttpRequestName   string    `bson:"request_name" json:"log_http_name"`
	LogHttpRequestStatus int       `bson:"status" json:"log_http_status"`
	LogHttpRequestUrl    string    `bson:"url" json:"log_http_url"`
	CreatedAt            time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time `bson:"updated_at" json:"updated_at"`
}
