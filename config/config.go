package config

import (
	"fmt"
	"os"
)

const (
	Port = 8080
)

// 1-3. Client ID、Client Secretを定義
// 誤ってgithubにアップロードしないように環境変数に変更しておく
var ClientID = os.Getenv("ClientID")
var ClientSecret= os.Getenv("ClientSecret")

// 1-4. リダイレクトURIを定義
var RedirectURI = fmt.Sprintf("http://localhost:%d/callback", Port)

// 1-5. OpenID ConnectのURLを定義
const (
	OIDCURL = "https://auth.login.yahoo.co.jp"
)
