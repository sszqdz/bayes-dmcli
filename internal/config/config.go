// Copyright 2024 Moran. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"time"
)

type Config struct {
	DatabaseList []*Database
	RedisList    []*Redis
}

type Database struct {
	Name            string
	Desc            string
	Driver          string
	Source          string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type Redis struct {
	Name         string
	Desc         string
	Addr         string
	Db           int32
	DialTimeout  int32
	ReadTimeout  int32
	WriteTimeout int32
}
