package blade

import (
	"context"
	"net/url"
)

type Consumer interface{ Namespace() string }

type ParsedConsumer interface {
	Consumer
	ConsumeParsed(context.Context, *url.URL) error
}

type ClassifiedConsumer interface {
	Consumer
	ConsumeClassified(context.Context, ClassifiedItem) error
}

type CollectionConsumer interface {
	Consumer
	ConsumeCollection(context.Context, CollectionItem) error
}

type MediaConsumer interface {
	Consumer
	ConsumeMedia(context.Context, MediaItem) error
}
