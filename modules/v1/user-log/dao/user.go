package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateActivityLogUser struct {
	LogActivityUserName       string    `bson:"log_activity_user_name" json:"log_activity_user_name"`
	LogActivityUserAction     string    `bson:"log_activity_user_action" json:"log_activity_user_action"`
	LogActivityUserAuthority  string    `bson:"log_activity_user_authority" json:"log_activity_user_authority"`
	LogActivityUserEmail      string    `bson:"log_activity_user_email" json:"log_activity_user_email"`
	LogActivityUserURLAppName string    `bson:"log_activity_user_url_app" json:"log_activity_user_url_app"`
	CreatedAt                 time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt                 time.Time `bson:"updated_at" json:"updated_at"`
}

type ListUserActivityLog struct {
	LogActivityUserLogID     primitive.ObjectID `bson:"_id" json:"log_user_id"`
	LogActivityUserName      string             `bson:"log_activity_user_name" json:"log_activity_user_name"`
	LogActivityUserAction    string             `bson:"log_activity_user_action" json:"log_activity_user_action"`
	LogActivityUserAuthority string             `bson:"log_activity_user_authority" json:"log_activity_user_authority"`
}

type DetailUserActivityLog struct {
	LogActivityUserLogID      primitive.ObjectID `bson:"_id" json:"log_user_id"`
	LogActivityUserName       string             `bson:"log_activity_user_name" json:"log_activity_user_name"`
	LogActivityUserAction     string             `bson:"log_activity_user_action" json:"log_activity_user_action"`
	LogActivityUserAuthority  string             `bson:"log_activity_user_authority" json:"log_activity_user_authority"`
	LogActivityUserEmail      string             `bson:"log_activity_user_email" json:"log_activity_user_email"`
	LogActivityUserURLAppName string             `bson:"log_activity_user_url_app" json:"log_activity_user_url_app"`
	CreatedAt                 time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt                 time.Time          `bson:"updated_at" json:"updated_at"`
}
