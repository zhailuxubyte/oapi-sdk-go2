# 飞书开放接口SDK

旨在让开发者便捷的调用飞书开放API、处理订阅的消息事件、处理服务端推送的卡片行为。

# 目录


<!-- toc -->

- [安装](#安装)
- [API Client](#api-client)
    - [创建API Client](#创建api-client)
    - [配置API Client](#配置api-client)

- [API调用](#api调用)
    - [基本用法](#基本用法)
    - [配置请求选项](#配置请求选项)
    - [原生API调用方式](#原生api调用方式)

- [处理消息事件回调](#处理消息事件回调)
    - [基本用法](#基本用法-1)
    - [集成gin框架](#集成gin框架)
        - [安装集成包](#安装集成包)
        - [集成示例](#集成示例)

- [处理卡片行为回调](#处理卡片行为回调)
    - [基本用法](#基本用法-2)
    - [返回卡片消息](#返回卡片消息)
    - [返回自定义消息](#返回自定义消息)
    - [集成gin框架](#集成gin框架)
        - [安装集成包](#安装集成包)
        - [集成示例](#集成示例)

<!-- tocstop -->

# 安装

go get -u github.com/larksuite/oapi-sdk-go

# API Client

开发者在调用 API 前，需要先创建一个 API Client，然后才可以基于 API Client 发起 API 调用

## 创建API Client

- 对于自建应用，可使用下面代码来创建一个 API Client

```go
var client = lark.NewClient("appID", "appSecret") // 默认配置为自建应用
```

- 对于商店应用，需在创建 API Client 时，使用 lark.WithMarketplaceApp 方法指定 AppType 为商店应用

```go
var client = lark.NewClient("appID", "appSecret",lark.WithMarketplaceApp()) // 设置App为商店应用
```

## 配置API Client

创建 API Client 时，可对 API Client 进行一定的配置，比如我们可以在创建 API Client 时设置日志级别、设置 http 请求超时时间等等：

```go
var client = lark.NewClient("appID", "appSecret",
lark.WithLogLevel(larkcore.LogLevelDebug),
lark.WithReqTimeout(3*time.Second),
lark.WithEnableTokenCache(true),
lark.WithHelpdeskCredential("id", "token"),
lark.WithLogger(larkcore.NewEventLogger()),
lark.WithHttpClient(http.DefaultClient))
```

每个配置选项的具体含义，如下表格：

<table>
  <thead align=left>
    <tr>
      <th>
        配置选项
      </th>
      <th>
        配置方式
      </th>
       <th>
        描述
      </th>
    </tr>
  </thead>
  <tbody align=left valign=top>
    <tr>
      <th>
        <code>LogLevel</code>
      </th>
      <td>
        <code>lark.WithLogLevel(logLevel larkcore.LogLevel)</code>
      </td>
      <td>
设置 API Client 的日志输出级别(默认为 Info 级别)，枚举值如下：

- LogLevelDebug
- LogLevelInfo
- LogLevelWarn
- LogLevelError

</td>
</tr>

<tr>
      <th>
        <code>AppType</code>
      </th>
      <td>
        <code>lark.WithMarketplaceApp()</code>
      </td>
      <td>
设置 App 类型为商店应用，ISV 开发者必须要设置该选项，默认为自建应用

</td>
</tr>

<tr>
      <th>
        <code>ReqTimeout</code>
      </th>
      <td>
        <code>lark.WithReqTimeout(time time.Duration)</code>
      </td>
      <td>
设置 Http 整个调用过程的超时时间，单位为 time.Duration。
默认为0，表示永不超时

</td>
</tr>


<tr>
      <th>
        <code>BaseUrl</code>
      </th>
      <td>
        <code>lark.WithOpenBaseUrl(baseUrl string)</code>
      </td>
      <td>
设置飞书域名；默认为该值：

```go
var FeishuBaseUrl = "https://open.feishu.cn"
var LarkBaseUrl = "https://open.larksuite.com"
```

</td>
</tr>

<tr>
      <th>
        <code>EnableTokenCache</code>
      </th>
      <td>
        <code>lark.WithEnableTokenCache(enableTokenCache bool)</code>
      </td>
      <td>
是否开启 UserAccessToken,TenantAccessToken 的自动获取与缓存;
默认开启，如需要关闭可传递 false
</td>
</tr>

<tr>
      <th>
        <code>HelpDeskId、HelpDeskToken</code>
      </th>
      <td>
        <code>lark.WithHelpdeskCredential(helpdeskID, helpdeskToken string)</code>
      </td>
      <td>
仅在调用服务台业务的 API 时需要传递
</td>
</tr>


<tr>
      <th>
        <code>Logger</code>
      </th>
      <td>
        <code>lark.WithLogger(logger larkcore.Logger)</code>
      </td>
      <td>
设置自定义的日志器，开发者需要实现下面的 Logger 接口:

```go
type Logger interface {
Debug(context.Context, ...interface{})
Info(context.Context, ...interface{})
Warn(context.Context, ...interface{})
Error(context.Context, ...interface{})
}
```

默认为标准输出
</td>
</tr>

<tr>
      <th>
        <code>HttpClient</code>
      </th>
      <td>
        <code>lark.WithHttpClient(httpClient larkcore.HttpClient)</code>
      </td>
      <td>
设置自定义的 httpClient，开发者需要实现下面的 HttpClient 接口:

```go
type HttpClient interface {
   Do(*http.Request) (*http.Response, error)
}

```

</td>
</tr>

<tr>
      <th>
        <code>TokenCache</code>
      </th>
      <td>
        <code>lark.WithTokenCache(cache larkcore.Cache)</code>
      </td>
      <td>
设置自定义的 token 缓存，开发者需要实现下面的 Cache 接口:

```go
type Cache interface {
Set(ctx context.Context, key string, value string, expireTime time.Duration) error
Get(ctx context.Context, key string) (string, error)
}

```

</td>
</tr>


<tr>
      <th>
        <code>LogReqRespInfoAtDebugLevel</code>
      </th>
      <td>
        <code>lark.WithLogReqRespInfoAtDebugLevel(printReqRespLog bool)</code>
      </td>
      <td>
开启 Http 请求参数和响应参数的日志打印开关；开启后，在 debug 模式下会打印 http 请求的 headers,body 等信息

</td>
</tr>
  </tbody>
</table>

# API调用
创建完毕 API Client，我们可以使用 Client.业务域.资源.方法名称 来定位具体的 API 方法，然后对具体的 API 发起调用。

飞书开放平台开放的所有 API 列表，可点击[这里查看](https://open.feishu.cn/document/ukTMukTMukTM/uYTM5UjL2ETO14iNxkTN/server-api-list)

## 基本用法

如下代码我们通过 client 调用文档业务的 Create 方法，创建一个文档：

``` go
import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/larksuite/oapi-sdk-go"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/service/docx/v1"
)


func main() {
	// 创建 client
	client := lark.NewClient("appID", "appSecret")

	// 发起请求
	resp, err := client.Docx.Document.Create(context.Background(), larkdocx.NewCreateDocumentReqBuilder().
		Body(larkdocx.NewCreateDocumentReqBodyBuilder().
			FolderToken("token").
			Title("title").
			Build()).
		Build())

	//处理错误
	if err != nil {
           // 处理err
           return
	}

	// 服务端错误处理
	if !resp.Success() {
           fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	   return 
	}

	// 业务数据处理
	fmt.Println(larkcore.Prettify(resp.Data))
}
```

## 配置请求选项

开发者在每次发起 API 调用时，可以设置请求级别的一些参数，比如传递 UserAccessToken ,自定义 Headers 等：

```go
import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/larksuite/oapi-sdk-go"
	larkcore "github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/service/docx/v1"
)

func main() {
	// 创建client
	client := lark.NewClient("appID", "appSecret")

	// 自定义请求headers
	header := make(http.Header)
	header.Add("k1", "v1")
	header.Add("k2", "v2")

	// 发起请求
	resp, err := client.Docx.Document.Create(context.Background(), larkdocx.NewCreateDocumentReqBuilder().
		Body(larkdocx.NewCreateDocumentReqBodyBuilder().
			FolderToken("token").
			Title("title").
			Build(),
		).
		Build(),
		larkcore.WithHeaders(header), // 设置自定义headers
	)

	//处理错误
	if err != nil {
	   // 处理err
	   return
	}

	// 服务端错误处理
	if !resp.Success() {
	   fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	   return
	}

	// 业务数据处理
	fmt.Println(larkcore.Prettify(resp.Data))
}

```

如下表格，展示了所有请求级别可配置的选项：

<table>
  <thead align=left>
    <tr>
      <th>
        配置选项
      </th>
      <th>
        配置方式
      </th>
       <th>
        描述
      </th>
    </tr>
  </thead>
  <tbody align=left valign=top>
    <tr>
      <th>
        <code>Header</code>
      </th>
      <td>
        <code>larkcore.WithHeaders(header http.Header)</code>
      </td>
      <td>
设置自定义请求头

</td>
</tr>

<tr>
      <th>
        <code>UserAccessToken</code>
      </th>
      <td>
        <code>larkcore.WithUserAccessToken(userAccessToken string)</code>
      </td>
      <td>
设置用户token，当需要以用户身份发起调用时，需要配置该选项

</td>
</tr>

<tr>
      <th>
        <code>TenantAccessToken</code>
      </th>
      <td>
        <code>larkcore.WithTenantAccessToken(tenantAccessToken string)</code>
      </td>
      <td>
当开发者自己维护租户 toke n时，可以通过该选项传递组合 token

</td>
</tr>


<tr>
      <th>
        <code>RequestId</code>
      </th>
      <td>
        <code>larkCore.WithRequestId(requestId string)</code>
      </td>
      <td>
设置请求ID

</td>
</tr>

<tr>
      <th>
        <code>TenantKey</code>
      </th>
      <td>
        <code>larkcore.WithTenantKey(tenantKey string)</code>
      </td>
      <td>
设置租户 key, 商店应用必须设置该选项
</td>
</tr>

  </tbody>
</table>

## 原生API调用方式

有些老版本的开放接口，不能生成结构化的 API， 导致 SDK 内无法提供结构化的使用方式，这时可使用原生模式进行调用：

```go
import (
	"context"
	"fmt"
	"os"

	"github.com/larksuite/oapi-sdk-go"
	"github.com/larksuite/oapi-sdk-go/core"
)

func main() {
	// 创建 API Client
	var cli = lark.NewClient("appID", "appSecret")

	// 发起请求
	resp, err := cli.Post(context.Background(), "https://www.feishu.cn/approval/openapi/v2/approval/get", map[string]interface{}{
		"approval_code": "ou_c245b0a7dff2725cfa2fb104f8b48b9d",
	}, larkcore.AccessTokenTypeTenant)

	// 错误处理
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取请求ID
	fmt.Println(resp.RequestId())

	// 处理请求结果
	fmt.Println(resp.StatusCode) // http status code
	fmt.Println(resp.Header)     // http header
	fmt.Println(resp.RawBody)    // http body}
}
```

# 处理消息事件回调
关于消息订阅相关的知识，可以点击[这里查看](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)

飞书开放平台开放的所有事件列表，可点击[这里查看](https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-list)
## 基本用法

开发者订阅消息事件后，可以使用下面代码，对飞书开放平台推送的消息事件进行处理，如下代码基于 go-sdk 原生 http server 启动一个 httpServer：

```go
import (
	"context"
	"fmt"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/event"
	"github.com/larksuite/oapi-sdk-go/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/httpserverext"
	"github.com/larksuite/oapi-sdk-go/service/contact/v3"
	"github.com/larksuite/oapi-sdk-go/service/im/v1"
)

func main() {
    // 注册消息处理器
    handler := dispatcher.NewEventDispatcher("verificationToken", "eventEncryptKey").OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
        fmt.Println(larkcore.Prettify(event))
        fmt.Println(event.RequestId())
        return nil
    }).OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
        fmt.Println(larkcore.Prettify(event))
        fmt.Println(event.RequestId())
        return nil
    })
    
    // 注册 http 路由
    http.HandleFunc("/webhook/event", httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))
    
    // 启动 http 服务
    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        panic(err)
    }
}


```

其中 NewEventDispatcher 方法的参数用于签名验证和消息解密使用，默认可以传递为空串；但是如果开发者的应用在 [控制台](https://open.feishu.cn/app?lang=zh-CN) 的【事件订阅】里面开启了加密，则必须传递控制台上提供的值。
![Console](doc/console.jpeg)

## 集成Gin框架
如果开发者当前应用使用的是 Gin Web 框架，并且不想要使用 Go-Sdk 提供的原生的 Http Server，则可使用下面方式，把当前应用的 Gin 服务与 SDK进行集成。

要想把 SDK 集成已有 Gin 框架，开发者需要引入集成包 [oapi-sdk-gin](https://github.com/larksuite/oapi-sdk-gin)

### 安装集成包

```go
go get -u github.com/larksuite/oapi-sdk-gin
```

### 集成示例

```go
import (
	"context"
	"fmt"

	 "github.com/gin-gonic/gin"
	 "github.com/larksuite/oapi-sdk-gin"
	 "github.com/larksuite/oapi-sdk-go/card"
	 "github.com/larksuite/oapi-sdk-go/core"
	 "github.com/larksuite/oapi-sdk-go/event/dispatcher"
	 "github.com/larksuite/oapi-sdk-go/service/contact/v3"
	 "github.com/larksuite/oapi-sdk-go/service/im/v1"
)

func main() {
	// 注册消息处理器
	handler := dispatcher.NewEventDispatcher("verificationToken", "eventEncryptKey").OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	}).OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	}).OnP2UserCreatedV3(func(ctx context.Context, event *larkcontact.P2UserCreatedV3) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	})

	...

	// 在已有 Gin 实例上注册消息处理路由
	gin.POST("/webhook/event", sdkginext.NewEventHandlerFunc(handler))
}
```

# 处理卡片行为回调

关于卡片行为相关的知识，可点击[这里查看](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)
## 基本用法

开发者配置消息卡片回调地址后，可以使用下面代码，对飞书开放平台推送的卡片行为进行处理，如下代码基于go-sdk原生http server启动一个httpServer：

```go
import (
	"context"
	"fmt"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/card"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/httpserverext"
)

func main() {
	// 创建 card 处理器
	cardHandler := larkcard.NewCardActionHandler("v", "", func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		fmt.Println(larkcore.Prettify(cardAction))
	    fmt.Println(cardAction.RequestId())
		// 无返回值示例
		return nil, nil
	})

	// 注册处理器
	http.HandleFunc("/webhook/card", httpserverext.NewCardActionHandlerFunc(cardHandler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	// 启动 http 服务
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}

```

如上示例，如果不需要处理器内返回业务结果给飞书服务端，则直接返回 nil

## 返回卡片消息

如开发者需要卡片处理器内同步返回用于更新消息卡片的消息体，则可使用下面方法方式进行处理：

```go
import (
	"context"
	"fmt"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/card"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/httpserverext"
)

func main() {
	// 创建card处理器
	cardHandler := larkcard.NewCardActionHandler("v", "", func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		fmt.Println(larkcore.Prettify(cardAction))
	    fmt.Println(cardAction.RequestId())
		
		// 创建http body
		body := make(map[string]interface{})
		body["content"] = "hello"

		i18n := make(map[string]string)
		i18n["zh_cn"] = "你好"
		i18n["en_us"] = "hello"
		i18n["ja_jp"] = "こんにちは"
		body["i18n"] = i18n 
		
		// 创建自定义消息：http状态码，body内容
		resp := &larkcard.CustomResp{
			StatusCode: 400,
			Body:       body,
		}

		return resp, nil
	})

	// 注册处理器
	http.HandleFunc("/webhook/card", httpserverext.NewCardActionHandlerFunc(cardHandler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	// 启动http服务
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}

```

## 返回自定义消息

如开发者需卡片处理器内返回自定义内容，则可以使用下面方式进行处理：

```go 
import (
	"context"
	"fmt"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/card"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/httpserverext"
)

func main() {
	// 创建 card 处理器
	cardHandler := larkcard.NewCardActionHandler("v", "", func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		fmt.Println(larkcore.Prettify(cardAction))
	    fmt.Println(cardAction.RequestId())
		
		// 创建 http body
		body := make(map[string]interface{})
		body["content"] = "hello"

		i18n := make(map[string]string)
		i18n["zh_cn"] = "你好"
		i18n["en_us"] = "hello"
		i18n["ja_jp"] = "こんにちは"
		body["i18n"] = i18n 
		
		// 创建自定义消息：http状态码，body内容
		resp := &larkcard.CustomResp{
			StatusCode: 400,
			Body:       body,
		}

		return resp, nil
	})

	// 注册处理器
	http.HandleFunc("/webhook/card", httpserverext.NewCardActionHandlerFunc(cardHandler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	// 启动 http 服务
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}


```

## 集成gin框架

如果开发者当前应用使用的是 Gin Web 框架，并且不想要使用 Go-Sdk 提供的原生的 Http Server，则可使用下面方式，把当前应用的 Gin 服务与 SDK进行集成。

要想把 SDK 集成已有 Gin 框架，开发者需要引入集成包 [oapi-sdk-gin](https://github.com/larksuite/oapi-sdk-gin)

### 安装集成包

```go
go get -u github.com/larksuite/oapi-sdk-gin
```

### 集成示例

```go
import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/larksuite/oapi-sdk-gin"
	"github.com/larksuite/oapi-sdk-go/card"
	"github.com/larksuite/oapi-sdk-go/core"
)


func main() {
    // 创建 card 处理器
    cardHandler := larkcard.NewCardActionHandler("v", "", func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
        fmt.Println(larkcore.Prettify(cardAction))
	    fmt.Println(cardAction.RequestId())
	    
        return nil, nil
    })
    ...
    // 在已有的 Gin 实例上注册卡片处理路由
    gin.POST("/webhooXk/card", sdkginext.NewCardActionHandlerFunc(cardHandler))
    ...
}
```

# License
使用 MIT



