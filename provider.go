package blade

import "go.uber.org/zap"

type Logger interface {
	Debug(string, ...zap.Field)
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)

	Named(string) Logger
	With(...zap.Field) Logger
}

type ProviderData struct {
	Logger Logger
}

type ProviderInfo struct {
	Name  string
	Short string
	Long  string
}

type Provider interface{ Info() ProviderInfo }

type ConsumerProvider interface {
	Provider
	Consumers(ProviderData) ([]Consumer, error)
}

type TransformerProvider interface {
	Provider
	Transformers(ProviderData) ([]Transformer, error)
}
