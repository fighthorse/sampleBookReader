package login

import (
	"github.com/fighthorse/sampleBookReader/domain/component/db"
	"github.com/fighthorse/sampleBookReader/domain/component/gocache"
	"time"
)

type Service struct {
	LocalCache *gocache.Cache
	Dao        *db.StoreDbClient
}

func New() (*Service, error) {

	return &Service{
		LocalCache: gocache.New(5*time.Minute, 10*time.Minute),
		Dao:        db.LoadDBByName("base"),
	}, nil
}
