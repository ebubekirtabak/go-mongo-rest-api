package common

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func Init() {
	setBasePath()
}

func setBasePath() {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	dir, err := os.Open(path.Join(basepath, "../"))
	if err != nil {
		panic(err)
	}

	os.Setenv("BASE_PATH", dir.Name() + "/")
}
