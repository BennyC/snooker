package playing

import "github.com/fudge/snooker/internal/entity"

type Service interface {
	NewMatch(entity.User, entity.User) entity.Match
}
