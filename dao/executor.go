//Package dao db处理类
package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

//DBParams 初始化参数
type DBParams struct {
	Host string `schema:"host"`
	Port int    `schema:"port"`
	User string `schema:"user"`
	Pwd  string `schema:"pwd"`
	DB   string `schema:"db"`
}

//Init 初始化
func Init(params *DBParams) error {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		params.User, params.Pwd, params.Host, params.Port, params.DB)
	eng, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		return err
	}
	engine = eng
	return nil
}

//GetEngine 获取引擎
func GetEngine() *xorm.Engine {
	return engine
}
