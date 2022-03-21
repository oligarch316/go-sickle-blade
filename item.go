package blade

import "net/url"

type ItemType struct{ val string }

var (
	ItemTypeInvalid    = ItemType{""}
	ItemTypeCollection = ItemType{"collection"}
	ItemTypeMedia      = ItemType{"media"}
)

func (it ItemType) String() string {
	switch it {
	case ItemTypeCollection, ItemTypeMedia:
		return it.val
	}
	return "invalid"
}

type Item interface{ Fingerprint() ([]byte, error) }

type ClassifiedItem interface {
	Item
	URL() (*url.URL, error)
	Type() ItemType
}

type CollectionItem interface {
	Item
	TODOCollection()
}

type MediaItem interface {
	Item
	TODOMedia()
}
