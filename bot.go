package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	_ "os"
)

const (
	//main API url
	URL = "https://api.telegram.org/bot"
	//Send Methods
	SendTxtMethod        = "sendMessage"
	SendPicMethod        = "sendPhoto"
	SendStickerMethod    = "sendSticker"
	AnswerCallbackMethod = "answerCallbackQuery"
	// Update methods
	EditMessageTextMethod        = "editMessageText"
	EditMessageCaptionMethod     = "editMessageCaption"
	EditMessageReplyMarkupMethod = "editMessageReplyMarkup"
	//Other
	PayloadType = "application/json"
)

func init() {

}

func SendTextMessage(key string, chat int, txt string, reply int, kbs ...InlineKeyboardButton) (int, error) {
	msg := new(UserMessage)

	msg.Text = txt
	msg.Reply = reply
	msg.ChatId = chat
	msg.ReplyMarkup = genInlineKeyboard(3, kbs...)

	pl, _ := json.Marshal(msg)
	plr := bytes.NewReader(pl)

	txtURL := genUrl(key, SendTxtMethod)

	resp, _ := http.Post(txtURL, PayloadType, plr)
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}

	if r.Ok {
		return r.Result.Id, nil
	}

	return r.Error, fmt.Errorf("%s", r.Description)
}

//SendPhotoByUrl returns fileid to reuse (original size), msg id, and err desc
func SendPhoto(key string, chat int, pic, cap string, reply int, kbs ...InlineKeyboardButton) (string, int, error) {
	msg := new(UserMessage)

	msg.Photo = pic
	msg.Caption = cap
	msg.ChatId = chat
	msg.Reply = reply
	msg.ReplyMarkup = genInlineKeyboard(3, kbs...)

	pl, _ := json.Marshal(msg)
	plr := bytes.NewReader(pl)

	picURL := genUrl(key, SendPicMethod)

	resp, _ := http.Post(picURL, PayloadType, plr)
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", 0, err
	}

	if r.Ok {
		lastItem := len(r.Result.Photo) - 1 // best quality
		return r.Result.Photo[lastItem].FileId, r.Result.Id, nil
	}

	return "", r.Error, fmt.Errorf("%s", r.Description)
}

func SendSticker(key string, chat int, st string, reply int) (int, error) {
	msg := new(UserMessage)

	msg.Sticker = st
	msg.ChatId = chat
	msg.Reply = reply

	pl, _ := json.Marshal(msg)
	plr := bytes.NewReader(pl)

	stickerURL := genUrl(key, SendStickerMethod)

	resp, _ := http.Post(stickerURL, PayloadType, plr)
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}
	//r.Result.Sticker.FileId
	if r.Ok {
		return r.Result.Id, nil
	}

	return r.Error, fmt.Errorf("%s", r.Description)
}

// response to button ket press with pop up msg in client
func AnswerCallbackQuery(key, id, txt string) error {
	ac := new(CallbackAnswer)

	ac.Id = id
	ac.Text = txt
	ac.ShowAlert = true

	pl, _ := json.Marshal(ac)
	plr := bytes.NewReader(pl)

	cbURL := genUrl(key, AnswerCallbackMethod)

	resp, _ := http.Post(cbURL, PayloadType, plr)
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}

	if !r.Ok {
		return fmt.Errorf("%s", r.Description)
	}

	return nil
}

func genUrl(k, m string) string {
	var b bytes.Buffer
	b.Grow(128)

	b.WriteString(URL)
	b.WriteString(k)
	b.WriteString("/")
	b.WriteString(m)

	return b.String()
}

func genInlineKeyboard(fc int, kbs ...InlineKeyboardButton) InlineKeyboard {
	//fc = field count (num of buttons in one line). 3 is most common and pretty
	ln := len(kbs) / fc + 1
	// lookup if we have rounded num of args
	if l := len(kbs) % fc; l == 0 {ln--}
	
	ikb := make([][]InlineKeyboardButton, ln)
	
	//gen all lines
	for i := 0; i < ln; i++ {
		ikb[i] = make([]InlineKeyboardButton, fc)
	}	
	// fill all lines
	for i, kb := range kbs {
		fn := i / fc
		pn := i % fc
		
		ikb[fn][pn] = kb
	}
	
	return  InlineKeyboard{Layout: ikb}
}

