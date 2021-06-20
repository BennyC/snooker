package storage

import (
	"github.com/fudge/snooker/internal/entity"
)

type Storage struct {
	Users entity.UserRepository
}
