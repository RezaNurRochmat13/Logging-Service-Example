package repository

import (
	"context"
	"svc-logger-go/modules/v1/user-log/dao"
	"svc-logger-go/util"

	"go.mongodb.org/mongo-driver/bson"
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

func (ur *userLogActivityRepoImpl) FindAll() ([]dao.ListUserActivityLog, error) {
	var (
		userActivityDao    dao.ListUserActivityLog
		resultUserActivity []dao.ListUserActivityLog
	)

	queryFindAllUserActivity, errorHandlerQuery := ur.Connection.Collection("user_log").Find(cntx, bson.M{})
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when query all user acitivity", "Error", errorHandlerQuery.Error())
		return nil, errorHandlerQuery
	}

	for queryFindAllUserActivity.Next(cntx) {
		errorHandlerDecode := queryFindAllUserActivity.Decode(&userActivityDao)
		if errorHandlerDecode != nil {
			util.LoggerOutput("Error when decode data", "Error", errorHandlerDecode.Error())
			return nil, errorHandlerDecode
		}
		resultUserActivity = append(resultUserActivity, userActivityDao)
	}

	return resultUserActivity, nil
}

func (ur *userLogActivityRepoImpl) Count() (int64, error) {
	countAllUserActivityLog, errorHandlerQuery := ur.Connection.Collection("user_log").CountDocuments(cntx, bson.M{})
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when count all user activity log", "Error", errorHandlerQuery.Error())
		return 0, errorHandlerQuery
	}

	return countAllUserActivityLog, nil
}
