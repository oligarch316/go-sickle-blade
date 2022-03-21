package blade

import (
	"context"
	"net/url"
)

type Transformer interface{ Namespace() string }

type Classifier interface {
	Transformer
	Classify(context.Context, *url.URL) (ClassifiedItem, error)
}

type Collector interface {
	Transformer
	Collect(context.Context, *url.URL) (CollectionItem, error)
}

type Download interface {
	Transformer
	Download(context.Context, *url.URL) (MediaItem, error)
}
