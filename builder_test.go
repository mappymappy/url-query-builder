package url_query_builder

import (
	"testing"
	"time"
)

type TestTargetStruct struct {
	IntField    int       `url:"int_field"`
	Int64Field  int64     `url:"int64_field"`
	StringField string    `url:"string_field"`
	BoolField   bool      `url:"bool_field"`
	TimeField   time.Time `url:"time_field"`
}

func TestBuildURI(t *testing.T) {
	builder := NewBuilder()
	target := TestTargetStruct{
		1,
		int64(2),
		"hoge",
		true,
		time.Time{},
	}
	actual := builder.BuildURI(target)
	expected := "bool_field=true&int64_field=2&int_field=1&string_field=hoge&time_field=0001-01-01T00%3A00%3A00Z"
	if actual != expected {
		t.Errorf("failed test!:\n actual:%s \n expected: %s", actual, expected)
	}
}
