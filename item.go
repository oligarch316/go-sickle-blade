package blade

import (
	"context"
	"io"
	"net/url"
)

type ItemClass struct{ val string }

var (
	ItemClassInvalid    = ItemClass{""}
	ItemClassCollection = ItemClass{"collection"}
	ItemClassMedia      = ItemClass{"media"}
)

func (ic ItemClass) String() string {
	switch ic {
	case ItemClassCollection, ItemClassMedia:
		return ic.val
	}
	return "invalid"
}

type Item interface{ Fingerprint() ([]byte, error) }

type ClassifiedItem interface {
	Item
	Class() ItemClass
	NormalURL() (*url.URL, error)
}

type CollectionItem interface {
	Item
	ChildURLs(context.Context) ([]*url.URL, error)
}

type MediaItem interface {
	Item
	Data(context.Context) (io.ReadCloser, error)
}
