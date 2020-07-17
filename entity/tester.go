package entity

import "github.com/rocketlaunchr/dbq/v2"

type Tester struct {
	Id    int    `json:"id" dbq:"tester_id"`
	Key   string `json:"key" dbq:"tester_key"`
	Value string `json:"value" dbq:"tester_value"`
}

// func (u *Tester) TableName() string {
// 	return "m_tester"
// }

func TesterSingleOption() *dbq.Options {
	return &dbq.Options{
		ConcreteStruct: Tester{},
		SingleResult:   true,
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
}

func TesterMultiOption() *dbq.Options {
	return &dbq.Options{
		ConcreteStruct: Tester{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
}
