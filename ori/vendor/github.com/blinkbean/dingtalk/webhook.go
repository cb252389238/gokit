package dingtalk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func InitDingTalk(tokens []string, key string) *DingTalk {
	if len(tokens) == 0 {
		panic("no token")
	}
	return &DingTalk{
		robotToken: tokens,
		keyWord: key,
	}
}

func InitDingTalkWithSecret(tokens string, secret string) *DingTalk {
	if len(tokens) == 0 || secret==""{
		panic("no token")
	}
	return &DingTalk{
		robotToken: []string{tokens},
		secret: secret,
	}
}

func (d *DingTalk) sendMessage(msg iDingMsg) error {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		uri    string
		resp   *http.Response
		err    error
	)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	value := url.Values{}
	value.Set("access_token", d.robotToken[rand.Intn(len(d.robotToken))])
	if d.secret!=""{
		t := time.Now().UnixNano() / 1e6
		value.Set("timestamp", fmt.Sprintf("%d", t))
		value.Set("sign", d.sign(t, d.secret))

	}
	uri = dingTalkURL + value.Encode()
	header := map[string]string{
		"Content-type": "application/json",
	}
	resp, err = doRequest(ctx, "POST", uri, header, msg.Marshaler())

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("send msg err: %s, token: %s, msg: %s", string(body), d.robotToken, msg.Marshaler())
	}
	return nil
}

func (d *DingTalk) sign(t int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func (d *DingTalk) OutGoing(r io.Reader) (outGoingMsg outGoingModel, err error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &outGoingMsg)
	return
}

func (d *DingTalk) SendTextMessage(content string, opt ...atOption) error {
	content = content + d.keyWord
	return d.sendMessage(NewTextMsg(content, opt...))
}

func (d *DingTalk) SendMarkDownMessage(title, text string, opts ...atOption) error {
	title = title + d.keyWord
	return d.sendMessage(NewMarkDownMsg(title, text, opts...))
}

// 利用dtmd发送点击消息
func (d *DingTalk) SendDTMDMessage(title string, dtmdMap *dingMap, opt ...atOption) error {
	title = title + d.keyWord
	return d.sendMessage(NewDTMDMsg(title, dtmdMap, opt...))
}

func (d DingTalk) SendMarkDownMessageBySlice(title string, textList []string, opts ...atOption) error {
	title = title + d.keyWord
	text := ""
	for _, t := range textList {
		text = text + "\n" + t
	}
	return d.sendMessage(NewMarkDownMsg(title, text, opts...))
}

func (d *DingTalk) SendLinkMessage(title, text, picUrl, msgUrl string) error {
	title = title + d.keyWord
	return d.sendMessage(NewLinkMsg(title, text, picUrl, msgUrl))
}

func (d *DingTalk) SendActionCardMessage(title, text string, opts ...actionCardOption) error {
	title = title + d.keyWord
	return d.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *DingTalk) SendActionCardMessageBySlice(title string, textList []string, opts ...actionCardOption) error {
	title = title + d.keyWord
	text := ""
	for _, t := range textList {
		text = text + "\n" + t
	}
	return d.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *DingTalk) SendFeedCardMessage(feedCard []FeedCardLinkModel) error {
	if len(feedCard) > 0 {
		feedCard[0].Title = feedCard[0].Title + d.keyWord
	}
	return d.sendMessage(NewFeedCardMsg(feedCard))
}
