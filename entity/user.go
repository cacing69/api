package entity

import "github.com/rocketlaunchr/dbq/v2"

// type User struct {
// 	Id   int    `json:"id" dbq:"user_id"`
// 	Name string `json:"name" dbq:"user_name"`
// }

type User struct {
	Id   int    `json:"id" dbq:"user_id"`
	Name string `json:"name" dbq:"user_name"`
}

// func (u *User) TableName() string {
// 	return "m_user"
// }

func UserSingleOption() *dbq.Options {
	return &dbq.Options{
		ConcreteStruct: User{},
		SingleResult:   true,
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
}

func UserMultiOption() *dbq.Options {
	return &dbq.Options{
		ConcreteStruct: User{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
}
