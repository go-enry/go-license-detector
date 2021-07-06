package licensedb

import (
	"os"
	"path/filepath"
	"testing"

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
