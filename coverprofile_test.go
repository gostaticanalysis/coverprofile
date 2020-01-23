package coverprofile_test

import (
	"testing"

	"github.com/gostaticanalysis/coverprofile"
	"golang.org/x/tools/cover"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	cases := []struct {
		pkg             string
		hasCoverProfile bool
	}{
		{"a", true},
		{"b", false},
	}

	testdata := analysistest.TestData()
	for _, tt := range cases {
		tt := tt
		t.Run(tt.pkg, func(t *testing.T) {
			rs := analysistest.Run(t, testdata, coverprofile.Analyzer, tt.pkg)
			
			// ignore xxx_test and xxx.test
			var r *analysistest.Result
			for i := range rs {
				if rs[i].Pass.Pkg.Path() == tt.pkg {
					r = rs[i]
					break
				}
			}

			if r == nil {
				t.Fatal("unexpected failure of analysistest.Run")
			}

			switch {
			case tt.hasCoverProfile && rs[0].Result == nil:
				t.Error("coverprofile cannot parse")
			case !tt.hasCoverProfile && rs[0].Result.([]*cover.Profile) != nil: // memo: typed nil
				t.Error("an unexpected coverprofile has parsed")
			}
		})
	}

}
