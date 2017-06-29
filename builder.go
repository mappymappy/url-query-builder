package url_query_builder

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const ConvertTargetTag = "url"

type QueryBuilder struct {
	timeFormat string
}

func NewBuilder() QueryBuilder {
	return QueryBuilder{
		timeFormat: time.RFC3339,
	}
}

func (q *QueryBuilder) BuildURI(target interface{}) string {
	s := reflect.TypeOf(target)
	v := reflect.ValueOf(target)
	urlParam := url.Values{}
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		q.convertUrlParam(field, v.FieldByName(field.Name).Interface(), &urlParam)
	}
	return urlParam.Encode()
}

func (q *QueryBuilder) convertUrlParam(field reflect.StructField, fieldValue interface{}, urlParam *url.Values) {
	tag := field.Tag.Get(ConvertTargetTag)
	switch fieldValue.(type) {
	case string:
		urlParam.Add(tag, fieldValue.(string))
	case int:
		urlParam.Add(tag, fmt.Sprintf("%d", fieldValue.(int)))
	case int64:
		urlParam.Add(tag, fmt.Sprintf("%d", fieldValue.(int64)))
	case bool:
		urlParam.Add(tag, strconv.FormatBool(fieldValue.(bool)))
	case time.Time:
		formattedTimeString := fieldValue.(time.Time).Format(q.timeFormat)
		urlParam.Add(tag, formattedTimeString)
	default:
	}

}

func (q *QueryBuilder) setTimeFormat(format string) {
	q.timeFormat = format
}
