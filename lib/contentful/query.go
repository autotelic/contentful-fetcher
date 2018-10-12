package contentful

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/autotelic/contentful-fetcher/lib/filesystem"
)

// Query is a query to be performed to Path and saved to Target.
type Query struct {
	Name   string            `json:"name"`
	Target string            `json:"target"`
	Path   string            `json:"path"`
	Params map[string]string `json:"params"`
}

// QueryExecutor is the interface that wraps the Execute method for executing a query.
type QueryExecutor interface {
	Execute(q *Query, apiURL string, accessToken string) error
}

type queryExecutor struct {
	fileOpener filesystem.Opener
	dirMaker   filesystem.DirMaker
}

// NewDefaultQueryExecutor constructs a QueryExecutor with the default dependencies.
func NewDefaultQueryExecutor() QueryExecutor {
	return NewQueryExecutor(
		filesystem.NewDefaultOpener(),
		filesystem.NewDefaultDirMaker(),
	)
}

// NewQueryExecutor constructs a struct that satisfies the QueryExecutor interface.
func NewQueryExecutor(
	fileOpener filesystem.Opener,
	dirMaker filesystem.DirMaker,
) QueryExecutor {
	return &queryExecutor{
		fileOpener: fileOpener,
		dirMaker:   dirMaker,
	}
}

// Execute executes the query, calling the apiURL with the query path and access token.
func (qe *queryExecutor) Execute(q *Query, apiURL string, accessToken string) error {
	u, err := url.Parse(apiURL)
	if err != nil {
		return err
	}

	u.Path = q.Path

	params := url.Values{}
	for k, v := range q.Params {
		params.Set(k, v)
	}

	params.Set("access_token", accessToken)

	u.RawQuery = params.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	targetPath := strings.Replace(q.Target, "/", string(filepath.Separator), -1)
	dir := filepath.Dir(targetPath)
	if err = qe.dirMaker.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	target, err := qe.fileOpener.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer target.Close()

	if _, err = io.Copy(target, res.Body); err != nil {
		return err
	}

	return nil
}
