package setting

import (
	"time"
)

//声明配置属性的结构体和读取区段配置的方法
type ServerSettingS struct {
	RunMode			string
	HttpPort		string
	ReadTimeout		time.Duration
	WriteTimeout	time.Duration
}

type AppSettingS struct {
	DefaultPageSize		int
	MaxPageSize			int
	LogSavePath			string
	LogFileName			string
	LogFileExt			string
}

type DataBaseSetting struct {
	DBType		string
	UserName	string
	Password	string
	Host		string
	DBName		string
	TablePrefix	string
	Charset		string
	ParseTime	bool
	MaxIdleConns int
	MaxOpenConns int
}



func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

