package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-github/github"
)

type myClient struct {
	apiGithubURL string
}

// ownerとrepoから最新のリリースタグ名を取得するメソッド
func (c *myClient) GetLatestReleaseTagName(owner, repo string) (string, error) {
	ctx := context.Background()
	client := github.NewClient(nil)

	// Github APIのベースURLを切り替えられるようにしておく
	client.BaseURL = c.getAPIGithubURL()

	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return "", err
	}
	return release.GetTagName(), nil
}

// go-githubのクライアントのClient.BaseURLに渡せるURLオブジェクト作成
// myClientインスタンス作成時にapiGithubURLを明示的に指定したら、
// それを利用するように。そうでなければデフォルトのURLを使う。
func (c *myClient) getAPIGithubURL() *url.URL {
	u := "https://api.github.com"
	if c.apiGithubURL != "" {
		u = c.apiGithubURL
	}
	apiURL, _ := url.Parse(u + "/")
	return apiURL
}

func main() {
	tagName, err := (&myClient{}).GetLatestReleaseTagName("shibayu36", "shibayu36")
	fmt.Println(tagName, err)
}
