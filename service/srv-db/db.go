package srv_db

func Init(conf DBConfig) (err error) {
	__client, err = newClient(conf)
	if err != nil {
		return
	}
	return
}

var __client *Client

func GetClient() *Client {
	return __client
}
