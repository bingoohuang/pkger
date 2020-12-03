package mem_test

import (
	"testing"

	"github.com/bingoohuang/pkger/pkging"
	"github.com/bingoohuang/pkger/pkging/mem"
	"github.com/bingoohuang/pkger/pkging/pkgtest"
)

func Test_Pkger(t *testing.T) {
	pkgtest.All(t, func(ref *pkgtest.Ref) (pkging.Pkger, error) {
		return mem.New(ref.Info)
	})
}
