package ulid

import (
	"sync"
	"time"

	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
)

type CustomULIDInterface interface {
	New() (ulid.ULID, error)
}

type CustomULIDStruct struct {
	mu      sync.Mutex
	entropy *rand.Rand
}

var ULID = getULID()

func getULID() CustomULIDInterface {
	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	return &CustomULIDStruct{
		entropy: entropy,
	}
}

func (u *CustomULIDStruct) New() (ulid.ULID, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, u.entropy)
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[ULID][New] Failed to generate ULID")

		return ulid.ULID{}, err
	}

	return id, nil
}
