package factory

import (
	"fmt"
	"hyssa/hyssa_go_billing_service/pkg/logger"
	"hyssa/hyssa_go_billing_service/pkg/logger/logrus"
	"hyssa/hyssa_go_billing_service/pkg/logger/options"
	"hyssa/hyssa_go_billing_service/pkg/logger/zap"
)

// logger map to map logger code to logger builder
var logFactoryBuilderMap = map[string]loggerBuilder{
	options.LOGRUS: &logrus.Factory{},
	options.ZAP:    &zap.Factory{},
}

// interface for logger factory
type loggerBuilder interface {
	Build(cfg *options.Logging) (logger.Logger, error)
}

// accessors for factoryBuilderMap
func getLogFactoryBuilder(key string) (loggerBuilder, error) {
	logFactoryBuilder, ok := logFactoryBuilderMap[key]
	if !ok {
		return nil, fmt.Errorf("not supported logger: %s", key)
	}

	return logFactoryBuilder, nil
}

// Build logger using appropriate log factory
func Build(cfg *options.Logging) (logger.Logger, error) {
	logFactoryBuilder, err := getLogFactoryBuilder(cfg.Code)
	if err != nil {
		return nil, err
	}

	log, err := logFactoryBuilder.Build(cfg)
	if err != nil {
		return nil, err
	}

	return log, nil
}
