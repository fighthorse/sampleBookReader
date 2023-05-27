package httpserver

import "github.com/fighthorse/sampleBookReader/domain/component/httpclient"

var (
	Amap = &AmapClient{}
)

func Init() {
	Amap.name = "amap"
	Amap.Client = httpclient.New(Amap.name)
}
