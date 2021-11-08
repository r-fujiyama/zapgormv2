package zapgormv2_test

import (
	"github.com/r-fujiyama/zapgormv2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Example() {
	zapLogger, _ := zap.NewProduction()
	logger := zapgormv2.New(zapLogger)
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	db, _ := gorm.Open(nil, &gorm.Config{Logger: logger})

	// do stuff normally
	var _ = db // avoid "unused variable" warn
}
