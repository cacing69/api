package entity

import "github.com/rocketlaunchr/dbq/v2"

type T struct {
	Count int `json:"count" dbq:"count"`
}

func GenericSingleOption() *dbq.Options {
	return &dbq.Options{
		ConcreteStruct: T{},
		SingleResult:   true,
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
}
