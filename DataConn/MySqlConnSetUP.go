package DataConn

import (
	"database/sql"
	"fmt"
)

func MysqlConn()  {
	fmt.Println("mysql", DataSourceCfg.MySqlUser+":"+DataSourceCfg.MySqlPW+"@tcp("+DataSourceCfg.MySqlIP+")/"+DataSourceCfg.MySqlDB+"?charset=utf8")
	Db, err := sql.Open("mysql", DataSourceCfg.MySqlUser+":"+DataSourceCfg.MySqlPW+"@tcp("+sDataSourceCfg.MySqlIP+")/"+DataSourceCfg.MySqlDB+"?charset=utf8")
	CheckError(err)
	Db.SetMaxOpenConns(2000)
	Db.SetMaxIdleConns(1000)
	err = Db.Ping()
	CheckError(err)
}

