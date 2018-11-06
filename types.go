package bot


type User struct {
	Id       int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Name     string `json:"first_name"`
	Lastname string `json:"last_name"`
	Username string `json:"username"`
	Lang     string `json:"language_code"`
}

type ChatPhoto struct {
	Small string `json:"small_file_id"`
	Big   string `json:"small_file_id"`
}

type PhotoSize struct {
	FileId string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"file_size"`
}

type Sticker struct {
	FileId string    `json:"file_id"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Size   int       `json:"file_size"`
	Thumb  PhotoSize `json:"thumb"`
}

type Document struct {
	FileId string    `json:"file_id"`
	Size   int       `json:"file_size"`
	Name   string    `json:"file_name"`
	Mime   string    `json:"mime_type"`
	Thumb  PhotoSize `json:"thumb"`
}

type File struct {
	FileId   string `json:"file_id,omitempty"`
	FilePath string `json:"file_path,omitempty"`
	Size     int    `json:"file_size,omitempty"`
}

type Video struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	Mime     string    `json:"mime_type"`
	Size     int       `json:"file_size"`
}

type Audio struct {
	FileId    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Title     string `json:"title"`
	Performer string `json:"performer"`
	Mime      string `json:"mime_type"`
	Size      int    `json:"file_size"`
}

type Voice struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	Mime     string `json:"mime_type"`
	Size     int    `json:"file_size"`
}

type Chat struct {
	Id       int       `json:"id"`
	ChatType string    `json:"type"`
	Title    string    `json:"title"`
	Username string    `json:"username"`
	Name     string    `json:"first_name"`
	Lastname string    `json:"last_name"`
	AMAA     bool      `json:"all_members_are_administrators"`
	Photo    ChatPhoto `json:"photo"`
	//some fields was omitted
}

type InlineKeyboardButton struct {
	Text string `json:"text,omtempty"`
	Url  string `json:"url,omitempty"`
	Data string `json:"callback_data,omitempty"`
}

type InlineKeyboard struct {
	Layout [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type UserMessage struct {
	ChatId      int            `json:"chat_id"`
	Photo       string         `json:"photo"` //Telegram id or url (doesn not support input file for now)
	Text        string         `json:"text,omitempty"`
	Caption     string         `json:"caption,omitempty"`
	Sticker     string	   `json:"sticker,omitempty"`
	ParseMode   string         `json:"parse_mode,omitempty"`
	Reply       int            `json:"reply_to_message_id,omitempty"`
	Type        string         `json:"-"`
	ReplyMarkup InlineKeyboard `json:"reply_markup,omitempty"`
}


type CallbackAnswer struct {
	Id        string `json:"callback_query_id"`
	Text      string `json:"text"`
	ShowAlert bool   `json:"show_alert"`
	url       string `json:"url"`
	CacheTime int    `json:"cache_time"`
}

type Message struct {
	Id                   int         `json:"message_id"`
	From                 User        `json:"from"`
	Date                 int         `json:"date"`
	Chat                 Chat        `json:"chat"`
	ForwardedFrom        User        `json:"forward_from"`
	ForwardedFromChat    Chat        `json:"forward_from_chat"`
	ForwardFromMessageId int         `json:"forward_from_message_id"`
	Reply                *Message    `json:"reply_to_message"`
	EditeDate            int         `json:"edit_date"`
	Text                 string      `json:"text"`
	Audio                Audio       `json:"audio"`
	Document             Document    `json:"document"`
	Photo                []PhotoSize `json:"photo"`
	Caption              string      `json:"caption"`
	Sticker              Sticker     `json:"sticker"`
	Video                Video       `json:"video"`
	Voice                Voice       `json:"voice"`
	//some fields was omitted
}

type Callback struct {
	Id           string  `json:"id"` //callback id, need for answer inline
	From         User    `json:"from"`
	Message      Message `json:"message"`
	InlineMesId  string  `json:"inline_message_id"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
}


type Response struct {
	Ok          bool    `json:"ok"`
	Result      Message `json:"result"`
	Error       int     `json:"error_code"`
	Description string  `json:"description"`
}
type GetFileResponse struct {
	Ok          bool    `json:"ok"`
	Error       int     `json:"error_code"`
	Description string  `json:"description"`
	Result      File    `json:"result"`

} 

type Update struct {
	UpdateId int        `json:"update_id"`
	Message  Message    `json:"message"`
	Callback Callback   `json:"callback_query"`
	ChannelPost Message `json:"channel_post"`
}
