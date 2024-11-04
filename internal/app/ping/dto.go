package ping

import "gorm.io/gorm"

type ExamplePost struct {
	Msg string `json:"msg"`
}

type TestModel struct {
	gorm.Model

	Name string `json:"name"`
	Age  int    `json:"age"`
}
