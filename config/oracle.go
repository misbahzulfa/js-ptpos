package config

import (
	"time"

	_ "github.com/mattn/go-oci8"
	"xorm.io/xorm"
)

var (
	EngineEcha *xorm.Engine
	EngineOltp *xorm.Engine
)

func InitEcha() (err error) {

	configuration := New()

	EngineEcha, err = xorm.NewEngine(configuration.Get("ORACLE_DRIVER"), configuration.Get("ORACLE_ECHA_CONNECTION_STRING"))
	if err != nil {
		return err
	}

	EngineEcha.SetConnMaxLifetime(1 * time.Minute)
	// EngineEcha.ShowSQL(true)

	err = EngineEcha.Ping()
	if err != nil {
		return err
	}
	return
}

func InitOltp() (err error) {

	configuration := New()

	EngineOltp, err = xorm.NewEngine(configuration.Get("ORACLE_DRIVER"), configuration.Get("ORACLE_OLTP_CONNECTION_STRING"))
	if err != nil {
		return err
	}

	EngineOltp.SetConnMaxLifetime(1 * time.Minute)
	// EngineOltp.ShowSQL(true)

	err = EngineOltp.Ping()
	if err != nil {
		return err
	}
	return
}
