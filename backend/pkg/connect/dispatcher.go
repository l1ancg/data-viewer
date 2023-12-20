package connect

import "github.com/l1ancg/data-viewer/backend/pkg"

type ConnectDispatcher struct {
	clients map[string]pkg.Client
}

func (c *ConnectDispatcher) Query(typ string, id int, uri string, ql string) ([]map[string]interface{}, error) {
	return c.clients[typ].Query(typ, id, uri, ql)
}

func ConnectProvider() *ConnectDispatcher {
	return &ConnectDispatcher{
		clients: map[string]pkg.Client{
			"MySQL": NewMySQLClient(),
		},
	}
}
