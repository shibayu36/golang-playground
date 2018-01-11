package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestReleaseTagName(t *testing.T) {
	// NewServeMuxを使えば、特定のPATHでレスポンスを返すような
	// アプリケーションを簡単に書けるので、
	// shibayu36/sample-repoの最新リリースAPIを模倣する。
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/shibayu36/sample-repo/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"tag_name": "0.0.1"}`)
	})

	// テスト用に作ったアプリケーションでテストサーバを立てる
	apiGithubServer := httptest.NewServer(mux)
	defer apiGithubServer.Close()

	// APIアクセスを立てたテストサーバに向けてmyClientオブジェクトを作成
	c := &myClient{
		apiGithubURL: apiGithubServer.URL,
	}

	{
		// 成功パターンのテスト
		tagName, err := c.GetLatestReleaseTagName("shibayu36", "sample-repo")
		assert.NoError(t, err)
		assert.Equal(t, "0.0.1", tagName, "shibayu36/sample-repoのタグが取れている")
	}

	{
		// 失敗パターン(レポジトリがない)のテスト
		_, err := c.GetLatestReleaseTagName("shibayu36", "wrong-repo")
		assert.Error(t, err, "存在しないrepoなのでエラー")
	}
}
