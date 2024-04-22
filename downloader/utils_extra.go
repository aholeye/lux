package downloader

import (
	"github.com/iawia002/lux/extractors"
)

func GenSortedStreams(streams map[string]*extractors.Stream) []*extractors.Stream {
	return genSortedStreams(streams)
}
