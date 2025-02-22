package srv

import srv_db "github.com/heyue7/BookHub-Server-Learn/service/srv-db"

func DB() *srv_db.Client {
	return srv_db.GetClient()
}
