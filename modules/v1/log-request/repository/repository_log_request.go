package repository

import (
	"context"
	"svc-logger-go/modules/v1/log-request/dao"
	"svc-logger-go/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cntx context.Context

type logRequestRepositoryImpl struct {
	Connection *mongo.Database
}

func NewLogRequestRepository(ConnectionDatabase *mongo.Database) Repository {
	return &logRequestRepositoryImpl{Connection: ConnectionDatabase}
}

func (lr *logRequestRepositoryImpl) FindAll() ([]dao.ListLogRequest, error) {
	var (
		logRequestDao    dao.ListLogRequest
		resultLogRequest []dao.ListLogRequest
	)

	queryFindAll, errorHandlerQuery := lr.Connection.Collection("http_request_log").Find(cntx, bson.M{})
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error query", "Error", errorHandlerQuery.Error())
		return nil, errorHandlerQuery
	}

	for queryFindAll.Next(cntx) {
		erorHandlerScanValue := queryFindAll.Decode(&logRequestDao)
		if erorHandlerScanValue != nil {
			util.LoggerOutput("Error scan value", "Error", erorHandlerScanValue.Error())
			return nil, erorHandlerScanValue
		}

		resultLogRequest = append(resultLogRequest, logRequestDao)
	}

	return resultLogRequest, nil
}

func (lr *logRequestRepositoryImpl) Save(payload *dao.CreateLogRequest) (*dao.CreateLogRequest, error) {
	_, errorHandlerQuery := lr.Connection.Collection("http_request_log").InsertOne(cntx, payload)
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error create new log", "Error", errorHandlerQuery.Error())
	}

	return payload, nil
}

func (lr *logRequestRepositoryImpl) FindById(id string) (dao.DetailLogRequest, error) {
	var (
		detailLogRequest dao.DetailLogRequest
		primitiveId, _   = primitive.ObjectIDFromHex(id)
		filter           = bson.M{"_id": primitiveId}
	)

	errorHandlerQuery := lr.Connection.Collection("http_request_log").FindOne(cntx, filter).Decode(&detailLogRequest)
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when decode data detail log", "Error", errorHandlerQuery.Error())
		return dao.DetailLogRequest{}, errorHandlerQuery
	}

	return detailLogRequest, nil
}

func (lr *logRequestRepositoryImpl) Update(id string, payload *dao.UpdateLogRequest) (*dao.UpdateLogRequest, error) {
	primitiveId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": primitiveId}
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	updateFieldLogRequest := bson.M{
		"$set": bson.M{
			"request_name": payload.LogHttpRequestName,
			"status":       payload.LogHttpRequestStatus,
			"url":          payload.LogHttpRequestUrl,
			"created_at":   payload.CreatedAt,
			"updated_at":   payload.UpdatedAt,
		},
	}

	_, errorHandlerQuery := lr.Connection.Collection("http_request_log").UpdateOne(cntx, filter, updateFieldLogRequest)
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when update log request", "Error", errorHandlerQuery.Error())
		return nil, errorHandlerQuery
	}

	return payload, nil
}

func (lr *logRequestRepositoryImpl) Delete(id string) error {
	primitiveId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": primitiveId}

	_, errorHandlerQuery := lr.Connection.Collection("http_request_log").DeleteOne(cntx, filter)
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when delete log request ", "Error", errorHandlerQuery.Error())
		return errorHandlerQuery
	}

	return nil
}
