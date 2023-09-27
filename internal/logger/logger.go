package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func MustLogger(cfg LoggerConfig) {
	congig := zap.NewProductionConfig()
	congig.Encoding = cfg.Encoding
	err := congig.Level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		panic(err.Error())
	}
	congig.OutputPaths = cfg.OutputPaths
	congig.ErrorOutputPaths = cfg.ErrorOutputPaths
	congig.EncoderConfig=zap.NewProductionEncoderConfig()
	// congig.EncoderConfig = cfg.EncoderConfig
	Logger, err = congig.Build()
	if err != nil {
		panic(err.Error())
	}
}
