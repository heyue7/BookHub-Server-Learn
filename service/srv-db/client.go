package srv_db

import (
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/log"
	"os"
)

type Client struct {
	*xorm.Engine
}

func newClient(conf DBConfig) (cli *Client, err error) {
	dsn := conf.Master

	engine, err := xorm.NewEngine(conf.Driver, dsn)
	if err != nil {
		return
	}

	f, err := os.OpenFile("sql.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	engine.SetLogger(log.NewSimpleLogger(f))

	cli = &Client{Engine: engine}
	return

}
