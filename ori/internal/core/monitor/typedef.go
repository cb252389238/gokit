package monitor

var webHookMsgChan = make(chan webHookMsgTextData, 10)

// webhook消息详细内容
type webHookMsgTextData struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}
