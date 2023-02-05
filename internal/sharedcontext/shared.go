package sharedcontext

import (
	"sync"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"go.uber.org/zap"
)

type SharedContext struct {
	Logger *zap.Logger
}

var once sync.Once
var (
	instance SharedContext
)

func GetInstance() SharedContext {

	once.Do(func() { // <-- atomic, does not allow repeating
		logutils.InitializeLogger()
		instance = SharedContext{
			Logger: logutils.Logger,
		}
	})

	return instance
}
