package blade

import (
	"context"
	"net/url"
)

// ----- Data
type PatternData struct {
	Parameters map[string]string
	Tail       []string
}

type HostData struct {
	URL  *url.URL
	Host PatternData
}

type PathData struct {
	URL  *url.URL
	Host PatternData
	Path PatternData
}

// ----- Targets
type HostTarget interface{ HostPattern() string }
type PathTarget interface{ PathPattern() string }

// ----- General Handlers
type HostHandler interface {
	Transformer
	HostTarget
}

type PathHandler interface {
	Transformer
	HostTarget
	Paths() []PathTarget
}

// ----- Specific Handlers
type HostClassifier interface {
	HostHandler
	Classify(context.Context, HostData) (ClassifiedItem, error)
}

type HostCollector interface {
	HostHandler
	Collect(context.Context, HostData) (CollectionItem, error)
}

type HostDownloader interface {
	HostHandler
	Download(context.Context, HostData) (MediaItem, error)
}

type PathClassifier interface {
	PathTarget
	Classify(context.Context, PathData) (ClassifiedItem, error)
}

type PathCollector interface {
	PathTarget
	Collect(context.Context, PathData) (CollectionItem, error)
}

type PathDownloader interface {
	PathTarget
	Download(context.Context, PathData) (MediaItem, error)
}
