# url-query-builder
provide simple and powerful url query builder by struct tag

# install

```
go get github.com/mappymappy/url-query-builder
```


# usage

please see `builder_test.go`

```
type TestTargetStruct struct {
	IntField    int       `url:"int_field"`
	Int64Field  int64     `url:"int64_field"`
	StringField string    `url:"string_field"`
	BoolField   bool      `url:"bool_field"`
	TimeField   time.Time `url:"time_field"`
}
target := TestTargetStruct{
	1,
	int64(2),
	"hoge",
	true,
	time.Time{},
}
builder := NewBuilder()
param := builder.BuildURI(target)
fmt.Printf("%s",param)
```



