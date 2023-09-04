package zap

import (
	"hyssa/hyssa_go_billing_service/pkg/logger"
	"hyssa/hyssa_go_billing_service/pkg/logger/options"
)

// Factory is the receiver for zap factory
type Factory struct{}

// Build zap logger
func (_ *Factory) Build(cfg *options.Logging) (logger.Logger, error) {
	l, err := RegisterLog(cfg)
	if err != nil {
		return nil, err
	}

	return l, nil
}
