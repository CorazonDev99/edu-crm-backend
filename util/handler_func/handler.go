package handler_func

import (
	"EduCRM/config"
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	cfg     = config.Config()
	loggers = logrus_log.GetLogger()
)

const (
	userIdCtx    = "userId"
	userRoleCtx  = "userRole"
	paramInvalid = "%s param is invalid"
	queryInvalid = "%s query is invalid"
)

// GetStringQuery
//func GetOffsetParam(ctx *gin.Context) (offset int64,
//	err error) {
//	offsetStr := ctx.DefaultQuery("offset", cfg.DefaultOffset)
//	offset, err = strconv.ParseInt(offsetStr, 10, 64)
//	if err != nil {
//		return 0, response.ErrorNotANumberOffset
//	}
//	if offset < 0 {
//		return 0, response.ErrorOffsetNotAUnsignedInt
//	}
//	return offset, nil
//}

// GetStringQuery func GetLimitParam(ctx *gin.Context) (limit int64,
//
//		err error) {
//		limitStr := ctx.DefaultQuery("limit", cfg.DefaultLimit)
//		limit, err = strconv.ParseInt(limitStr, 10, 64)
//		if err != nil {
//			return 0, response.ErrorNotANumberLimit
//		}
//		if limit < 0 {
//			return 0, response.ErrorLimitNotAUnsignedInt
//		}
//		return limit, nil
//	}
func GetStringQuery(ctx *gin.Context,
	query string) (param string, err error) {
	param = ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(" %s param is empty", query)
		return "", errors.New(err)
	}
	return param, nil
}
func GetInt64Query(ctx *gin.Context,
	query string) (int64,
	error) {
	param := ctx.Query(query)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
func GetFloat64Query(ctx *gin.Context,
	query string) (float64,
	error) {
	param := ctx.Query(query)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
func GetArrayStringQuery(ctx *gin.Context,
	query string) ([]string, error) {
	param := ctx.Query(query)
	if param == "" {
		return []string{}, errors.New("param is empty")
	}
	chunks := strings.Split(param, ",")
	return chunks, nil
}
func JsonUnmarshal(pointData *interface{},
	data []byte) error {
	err := json.Unmarshal(data, pointData)
	if err != nil {
		return err
	}
	return nil
}
func GetBooleanQuery(ctx *gin.Context,
	query string) (bool,
	error) {
	param := ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(paramInvalid, query)
		return false, errors.New(err)
	}
	boolVal, err := strconv.ParseBool(param)
	if err != nil {
		err := fmt.Sprintf(paramInvalid, query)
		return false, errors.New(err)
	}
	return boolVal, nil
}
func GetUUIDQuery(ctx *gin.Context,
	query string) (string,
	error) {
	param := ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(paramInvalid, query)
		return "", errors.New(err)
	}
	paramUUID, err := uuid.Parse(param)
	if err != nil {
		logrus.Error(err)
		err := fmt.Sprintf(paramInvalid, query)
		return "", errors.New(err)
	}
	return paramUUID.String(), nil
}
func GetUserId(ctx *gin.Context) (string, error) {
	id, ok := ctx.Get(userIdCtx)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return "", errors.New(response.ErrorUserIDInvalid.Error())
	}
	userID, ok := id.(string)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return "", errors.New(response.ErrorUserIDInvalid.Error())
	}
	_, err := uuid.Parse(userID)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return userID, nil
}
func GetUserRole(ctx *gin.Context) (string, error) {
	id, ok := ctx.Get(userRoleCtx)
	if !ok {
		return "", errors.New("user role not found")
	}
	idInt, ok := id.(string)
	if !ok {
		return "", errors.New("user role not found")
	}
	return idInt, nil
}

func GetPageQuery(ctx *gin.Context) (offset int64,
	err error) {
	offsetStr := ctx.DefaultQuery("page", cfg.DefaultPage)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return 0, response.ErrorNotANumberPage
	}
	if offset < 0 {
		return 0, response.ErrorOffsetNotAUnsignedInt
	}
	return offset, nil
}
func GetPageSizeQuery(ctx *gin.Context) (limit int64,
	err error) {
	limitStr := ctx.DefaultQuery("pageSize", cfg.DefaultPageSize)
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, response.ErrorNotANumberPageSize
	}
	if limit < 0 {
		return 0, response.ErrorLimitNotAUnsignedInt
	}
	return limit, nil
}

func CalculationPagination(page, pageSize int64) (offset, limit int64) {
	if page < 0 {
		page = 1
	}
	offset = (page - 1) * pageSize
	limit = pageSize
	return offset, limit
}

func ListPagination(ctx *gin.Context) (pagination model.Pagination, err error) {
	page, err := GetPageQuery(ctx)
	if err != nil {
		logrus.Error(err)
		return pagination, err
	}
	pageSize, err := GetPageSizeQuery(ctx)
	if err != nil {
		logrus.Error(err)
		return pagination, err
	}
	offset, limit := CalculationPagination(page, pageSize)
	pagination.Limit = limit
	pagination.Offset = offset
	pagination.Page = page
	pagination.PageSize = pageSize
	return pagination, nil
}

func GetUUIDParam(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Param(query)
	if queryData == "" {
		err := fmt.Sprintf(queryInvalid, queryData)
		return uuid.Nil, errors.New(err)
	}
	queryUUID, err := uuid.Parse(queryData)
	if err != nil {
		err := fmt.Sprintf(queryInvalid, queryData)
		return uuid.Nil, errors.New(err)
	}
	return queryUUID, nil
}
