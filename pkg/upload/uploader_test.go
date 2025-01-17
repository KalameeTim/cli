package upload

import (
	"bytes"
	"encoding/json"
	"github.com/debricked/cli/pkg/client"
	"github.com/debricked/cli/pkg/file"
	"github.com/debricked/cli/pkg/git"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestNewDebrickedUploader(t *testing.T) {
	uploader, err := NewUploader(nil)
	if err == nil {
		t.Error("failed to assert that error occurred")
	}
	if uploader != nil {
		t.Error("failed to assert that Uploader was nil")
	}
	var c client.IDebClient
	c = &debClientMock{}
	uploader, err = NewUploader(&c)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}

	if uploader == nil {
		t.Error("failed to assert that Uploader was not nil")
	}
}
func TestUpload(t *testing.T) {
	var c client.IDebClient
	c = &debClientMock{}
	uploader, _ := NewUploader(&c)
	metaObject, _ := git.NewMetaObject(
		"testdata/yarn",
		"testdata/yarn",
		"testdata/yarn-commit",
		"",
		"",
		"",
	)
	g := file.NewGroup("testdata/yarn/package.json", nil, []string{"testdata/yarn/yarn.lock"})
	groups := file.Groups{}
	groups.Add(*g)
	uploaderOptions := DebrickedOptions{FileGroups: groups, GitMetaObject: *metaObject, IntegrationsName: "CLI"}
	result, err := uploader.Upload(uploaderOptions)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if result == nil {
		t.Error("failed to assert that result was not nil")
	}
}

type debClientMock struct{}

func (mock *debClientMock) Post(uri string, _ string, _ *bytes.Buffer) (*http.Response, error) {
	res := &http.Response{
		Status:           "",
		StatusCode:       http.StatusOK,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          nil,
		TLS:              nil,
	}
	var resBodyBytes []byte
	if uri == "/api/1.0/open/uploads/dependencies/files" {
		f := uploadedFile{1, 0, 0, 0, "0", 0}
		resBodyBytes, _ = json.Marshal(f)

	} else if uri == "/api/1.0/open/finishes/dependencies/files/uploads" {
		res.StatusCode = http.StatusNoContent
	}

	res.Body = io.NopCloser(strings.NewReader(string(resBodyBytes)))

	return res, nil
}

var progress = 50

func (mock *debClientMock) Get(_ string, _ string) (*http.Response, error) {
	res := &http.Response{
		Status:           "",
		StatusCode:       http.StatusOK,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          nil,
		TLS:              nil,
	}

	f := &uploadStatus{progress, 0, 0, "", nil, ""}
	progress = progress + progress%100

	resBodyBytes, _ := json.Marshal(f)
	res.Body = io.NopCloser(strings.NewReader(string(resBodyBytes)))

	return res, nil
}
