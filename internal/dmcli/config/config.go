package config

import (
	"time"
)

type Config struct {
	DatabaseList []*Database
	RedisList    []*Redis
}

type Database struct {
	Driver          string
	Source          string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type Redis struct {
	Addr         string
	Db           int32
	DialTimeout  int32
	ReadTimeout  int32
	WriteTimeout int32
}
