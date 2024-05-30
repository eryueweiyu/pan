package file

import (
	"github.com/jsyzchen/pan/conf"
	"testing"
)

func TestUpload(t *testing.T) {
	fileUploader := NewUploader(conf.TestData.AccessToken, conf.TestData.Path, conf.TestData.LocalFilePath, 0)
	res, err := fileUploader.Upload()
	if err != nil {
		t.Fail()
	} else {
		t.Logf("TestUpload Success res: %+v", res)
	}
}
