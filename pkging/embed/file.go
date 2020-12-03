package embed

import (
	"github.com/bingoohuang/pkger/here"
	"github.com/bingoohuang/pkger/pkging"
)

type File struct {
	Info   *pkging.FileInfo `json:"info"`
	Here   here.Info        `json:"her"`
	Path   here.Path        `json:"path"`
	Data   []byte           `json:"data"`
	Parent here.Path        `json:"parent"`
}
