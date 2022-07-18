package licensedb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-enry/go-license-detector/v4/licensedb/filer"
)

func BenchmarkDetect(b *testing.B) {
	f := pwdFiler()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Detect(f)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkDetectWithPreload(b *testing.B) {
	f := pwdFiler()
	Preload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Detect(f)
		if err != nil {
			panic(err)
		}
	}
}

func pwdFiler() filer.Filer {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root := filepath.Dir(pwd)
	f, err := filer.FromDirectory(root)
	if err != nil {
		panic(err)
	}
	return f
}

func TestLicenseURLs(t *testing.T) {
	t.Run("existing license", func(t *testing.T) {
		res, err := LicenseURLs("ODbL-1.0")
		require.NoError(t, err)
		assert.Equal(t, []string{"http://www.opendatacommons.org/licenses/odbl/1.0/", "https://opendatacommons.org/licenses/odbl/1-0/"}, res)
	})

	t.Run("not existing license", func(t *testing.T) {
		_, err := LicenseURLs("bad-license-key")
		require.Equal(t, ErrUnknownLicenseID, err)
	})
}

func TestLicenseName(t *testing.T) {
	t.Run("existing license", func(t *testing.T) {
		res, err := LicenseName("ODbL-1.0")
		require.NoError(t, err)
		assert.Equal(t, "Open Data Commons Open Database License v1.0", res)
	})

	t.Run("not existing license", func(t *testing.T) {
		_, err := LicenseName("bad-license-key")
		require.Equal(t, ErrUnknownLicenseID, err)
	})
}
