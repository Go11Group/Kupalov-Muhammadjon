package logger

import (
	"path/filepath"

	"go.uber.org/zap"
)

func newDevelopmentConfig(file string) *zap.Config{
	configZap := zap.NewDevelopmentConfig()
	configZap.OutputPaths = []string{"stdout", file}
	configZap.ErrorOutputPaths = []string{"stderr"}

	return &configZap
}

func New(fileName string) (*zap.Logger, error){
	file := filepath.Join("./", fileName)

	configZap := newDevelopmentConfig(file)

	configZap.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	return configZap.Build()
}
