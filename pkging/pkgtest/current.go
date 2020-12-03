package pkgtest

import (
	"testing"

	"github.com/bingoohuang/pkger/pkging"
)

func CurrentTest(t *testing.T, ref *Ref, pkg pkging.Pkger) {
	cur, err := pkg.Current()
	if err != nil {
		t.Fatal(err)
	}

	cmpHereInfo(t, ref.Info, cur)
}
