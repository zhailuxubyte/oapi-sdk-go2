package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var appID, appSecret = os.Getenv("APP_ID"), os.Getenv("APP_SECRET")
	larkApp := lark.NewApp(lark.DomainFeiShu, lark.WithAppCredential(appID, appSecret),
		lark.WithLogger(lark.NewDefaultLogger(), lark.LogLevelDebug))

	ctx := context.Background()
	sendMessage(ctx, larkApp)
	uploadImage(ctx, larkApp)
	downloadImage(ctx, larkApp)
}

func sendMessage(ctx context.Context, larkApp *lark.App) {
	resp, err := larkApp.SendRequest(ctx, "POST", "/open-apis/message/v4/send", map[string]interface{}{
		"user_id":  "77bbc392",
		"msg_type": "text",
		"content":  &lark.MessageText{Text: "test"},
	}, lark.AccessTokenTypeTenant, lark.WithHTTPHeader(map[string][]string{"X-Request-Id": []string{"3a87a759a4f03391607ed1c7"}}))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.RequestId())
	fmt.Println(resp)
	fmt.Println()
	fmt.Println()
}

func uploadImage(ctx context.Context, larkApp *lark.App) {
	img, err := os.Open("test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer img.Close()
	resp, err := larkApp.SendRequest(ctx, "POST", "/open-apis/image/v4/put",
		lark.NewFormdata().AddField("image_type", "message").AddFile("image", img), lark.AccessTokenTypeTenant)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.RequestId())
	fmt.Println(resp)
	fmt.Println()
	fmt.Println()
}

func downloadImage(ctx context.Context, larkApp *lark.App) {
	resp, err := larkApp.SendRequest(ctx, "GET", "/open-apis/image/v4/get?image_key=img_v2_a0cea146-64d2-4dcb-94c7-636586fea98g",
		nil, lark.AccessTokenTypeTenant)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		return
	}
	err = ioutil.WriteFile("test_download_v2.png", resp.RawBody, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.RequestId())
	fmt.Println(resp)
}