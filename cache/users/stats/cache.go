package stats

import (
	"errors"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/utils"
	"gorm.io/gorm"
	"sync"
)

type Identifier struct {
	User  int
	Mode  constants.Mode
	Relax bool
}

func (i *Identifier) TableName() string {
	return "stats_" + utils.ModeToDbSuffix(i.Mode, i.Relax)
}

var values = map[Identifier]*layouts.Stats{}
var mutex sync.RWMutex

func AutoMigrate(db *gorm.DB) {
	for _, v := range utils.Modes {
		db.Table("stats_" + v).AutoMigrate(&layouts.Stats{})
	}
}

func Get(identifier Identifier) *layouts.Stats {
	mutex.RLock()
	v, ok := values[identifier]
	mutex.RUnlock()

	if ok {
		return v
	} else {
		return Fetch(identifier)
	}
}

func Fetch(identifier Identifier) *layouts.Stats {
	stats := &layouts.Stats{User: identifier.User}
	err := database.DB.
		Table(identifier.TableName()).
		Where(stats).
		First(stats).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
		}
		return nil
	}

	mutex.Lock()
	values[identifier] = stats
	mutex.Unlock()

	return stats
}