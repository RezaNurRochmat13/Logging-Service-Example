package repository

import (
	"context"
	"svc-logger-go/modules/v1/user-log/dao"
	"svc-logger-go/util"

	"go.mongodb.org/mongo-driver/mongo"
)

var cntx context.Context

type userLogActivityRepoImpl struct {
	Connection *mongo.Database
}

func NewUserLogActivityRepository(ConnectionDB *mongo.Database) Repository {
	return &userLogActivityRepoImpl{Connection: ConnectionDB}
}

func (ur *userLogActivityRepoImpl) Save(payload *dao.CreateActivityLogUser) (*dao.CreateActivityLogUser, error) {
	_, errorHandlerQuery := ur.Connection.Collection("user_log").InsertOne(cntx, payload)
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when saving user log", "Error", errorHandlerQuery.Error())
		return nil, errorHandlerQuery
	}

	return payload, nil
}
