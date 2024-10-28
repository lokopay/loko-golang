package payloads

import (
	"net/url"
	"reflect"
	"strconv"
)

type QueryParam struct {
	Limit         int64  `json:"limit,omitempty"`
	StartingAfter string `json:"starting_after,omitempty"`
	EndingBefore  string `json:"ending_before,omitempty"`
	CreatedFrom   int64  `json:"created_from,omitempty"`
	CreatedTo     int64  `json:"created_to,omitempty"`
	CompletedFrom int64  `json:"completed_from,omitempty"`
	CompletedTo   int64  `json:"completed_to,omitempty"`
	Status        string `json:"status,omitempty"`
}

func NewQueryParam() *QueryParam {
	return &QueryParam{}
}

func (q *QueryParam) SetLimit(limit int64) *QueryParam {
	q.Limit = limit
	return q
}

func (q *QueryParam) SetStartingAfter(startingAfter string) *QueryParam {
	q.StartingAfter = startingAfter
	return q
}

func (q *QueryParam) SetEndingBefore(endingBefore string) *QueryParam {
	q.EndingBefore = endingBefore
	return q
}

func (q *QueryParam) SetCreatedFrom(createdFrom int64) *QueryParam {
	q.CreatedFrom = createdFrom
	return q
}

func (q *QueryParam) SetCreatedTo(createdTo int64) *QueryParam {
	q.CreatedTo = createdTo
	return q
}

func (q *QueryParam) SetCompletedFrom(completedFrom int64) *QueryParam {
	q.CompletedFrom = completedFrom
	return q
}

func (q *QueryParam) SetCompletedTo(completedTo int64) *QueryParam {
	q.CompletedTo = completedTo
	return q
}

func (q *QueryParam) SetStatus(status string) *QueryParam {
	q.Status = status
	return q
}

func (q *QueryParam) String() string {
	values := url.Values{}
	v := reflect.ValueOf(q)

	// Ensure v is not a pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Tag.Get("json")

		// Skip fields without json tag or with omitempty suffix when empty
		if fieldName == "" || (field.Kind() == reflect.String && field.String() == "") || (field.Kind() == reflect.Int64 && field.Int() == 0) {
			continue
		}

		// Remove ",omitempty" from the field name
		fieldName = fieldName[:len(fieldName)-len(",omitempty")]

		switch field.Kind() {
		case reflect.String:
			values.Add(fieldName, field.String())
		case reflect.Int64:
			values.Add(fieldName, strconv.FormatInt(field.Int(), 10))
		}
	}

	return values.Encode()
}
