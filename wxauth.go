package main

import (
    "github.com/chanxuehong/wechat/message/passive/request"
    "github.com/chanxuehong/wechat/message/passive/response"
    "github.com/chanxuehong/wechat/server"
    "log"
    "net/http"
    "time"
    "fmt"
)

// 实现 server.Agent
type CustomAgent struct {
    server.DefaultAgent // 可选, 不是必须!!! 提供了默认实现
}

// 文本消息处理函数, 覆盖默认的实现
func (this *CustomAgent) ServeTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64) {
    // TODO: 示例代码, 把用户发送过来的文本原样回复过去

    w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选

    // NOTE: 时间戳也可以用传入的参数 timestamp, 即微信服务器请求 URL 中的 timestamp
    resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.Content, time.Now().Unix())
    fmt.Println("FromUserName=",msg.FromUserName)
    fmt.Println("ToUserName=",msg.ToUserName)
    fmt.Println("COntent=",msg.Content)

    if err := this.WriteText(w, resp); err != nil {
        // TODO: 错误处理代码
    }
}

// 自定义错误请求处理函数
func CustomInvalidRequestHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
    // TODO: 这里只是简单的做下 log
    log.Println(err)
}

func init() {
    var agent CustomAgent
    agent.DefaultAgent.Id = "gh_b0c9f3b69b09"       // 参考 公众号设置-->帐号详情-->原始ID
    agent.DefaultAgent.Token = "a42d4607561aeda8a25352c32aae7feb" // 这里填上你自己的 Token

    // InvalidRequestHandler == nil 则会默认什么都不做
    agentFrontend := server.NewAgentFrontend(&agent, nil)

    // 注册这个 agentFrontend 到回调 URL 上
    // 比如你在公众平台后台注册的回调地址是 http://abc.xyz.com/weixin，那么可以这样注册
    http.Handle("/weixin", agentFrontend)
}

func main() {
    http.ListenAndServe("127.0.0.1:8015", nil)
}
