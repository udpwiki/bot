package bot

import (
	"github.com/go-telegram/bot/models"
)

type SetWebhookParams struct {
	URL                string           `json:"url"`
	Certificate        models.InputFile `json:"certificate,omitempty"`
	IPAddress          string           `json:"ip_address,omitempty"`
	MaxConnections     int              `json:"max_connections,omitempty"`
	AllowedUpdates     []string         `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool             `json:"drop_pending_updates,omitempty"`
	SecretToken        string           `json:"secret_token,omitempty"`
}

type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

type SendMessageParams struct {
	ChatID                   any                    `json:"chat_id"`
	Text                     string                 `json:"text"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	Entities                 []models.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                   `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type ForwardMessageParams struct {
	ChatID              any    `json:"chat_id"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	MessageID           int    `json:"message_id"`
}

type CopyMessageParams struct {
	ChatID                   any                    `json:"chat_id"`
	FromChatID               string                 `json:"from_chat_id"`
	MessageID                int                    `json:"message_id"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendPhotoParams struct {
	ChatID                   any                    `json:"chat_id"`
	Photo                    models.InputFile       `json:"photo"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendAudioParams struct {
	ChatID                   any                    `json:"chat_id"`
	Audio                    models.InputFile       `json:"audio"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	Performer                string                 `json:"performer,omitempty"`
	Title                    string                 `json:"title,omitempty"`
	Thumb                    string                 `json:"thumb,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendDocumentParams struct {
	ChatID                      any                    `json:"chat_id"`
	Document                    string                 `json:"document"`
	Thumb                       string                 `json:"thumb,omitempty"`
	Caption                     string                 `json:"caption,omitempty"`
	ParseMode                   models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                   `json:"disable_notification,omitempty"`
	ProtectContent              bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID            int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendVideoParams struct {
	ChatID                   any                    `json:"chat_id"`
	Video                    models.InputFile       `json:"video"`
	Duration                 int                    `json:"duration,omitempty"`
	Width                    int                    `json:"width,omitempty"`
	Height                   int                    `json:"height,omitempty"`
	Thumb                    string                 `json:"thumb,omitempty"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	SupportsStreaming        bool                   `json:"supports_streaming,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendAnimationParams struct {
	ChatID                   any                    `json:"chat_id"`
	Animation                models.InputFile       `json:"animation"`
	Duration                 int                    `json:"duration,omitempty"`
	Width                    int                    `json:"width,omitempty"`
	Height                   int                    `json:"height,omitempty"`
	Thumb                    string                 `json:"thumb,omitempty"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendVoiceParams struct {
	ChatID                   any                    `json:"chat_id"`
	Voice                    models.InputFile       `json:"voice"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendVideoNoteParams struct {
	ChatID                   any                `json:"chat_id"`
	VideoNote                models.InputFile   `json:"video_note"`
	Duration                 int                `json:"duration,omitempty"`
	Length                   int                `json:"length,omitempty"`
	Thumb                    string             `json:"thumb,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendMediaGroupParams struct {
	ChatID                   any                 `json:"chat_id"`
	Media                    []models.InputMedia `json:"media"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageID         int                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
}

type SendLocationParams struct {
	ChatID                   any                `json:"chat_id"`
	Latitude                 float64            `json:"latitude"`
	Longitude                float64            `json:"longitude"`
	HorizontalAccuracy       float64            `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int                `json:"live_period,omitempty"`
	Heading                  int                `json:"heading,omitempty"`
	ProximityAlertRadius     int                `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageLiveLocationParams struct {
	ChatID               any                `json:"chat_id"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	Latitude             float64            `json:"latitude"`
	Longitude            float64            `json:"longitude"`
	HorizontalAccuracy   float64            `json:"horizontal_accuracy,omitempty"`
	Heading              int                `json:"heading,omitempty"`
	ProximityAlertRadius int                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type StopMessageLiveLocationParams struct {
	ChatID          string             `json:"chat_id,omitempty"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVenueParams struct {
	ChatID                   any                `json:"chat_id"`
	Latitude                 float64            `json:"latitude"`
	Longitude                float64            `json:"longitude"`
	Title                    string             `json:"title"`
	Address                  string             `json:"address"`
	FoursquareID             string             `json:"foursquare_id,omitempty"`
	FoursquareType           string             `json:"foursquare_type,omitempty"`
	GooglePlaceID            string             `json:"google_place_id,omitempty"`
	GooglePlaceType          string             `json:"google_place_type,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendContactParams struct {
	ChatID                   any                `json:"chat_id"`
	PhoneNumber              string             `json:"phone_number"`
	FirstName                string             `json:"first_name"`
	LastName                 string             `json:"last_name,omitempty"`
	VCard                    string             `json:"vcard,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendPollParams struct {
	ChatID                   any                    `json:"chat_id"`
	Question                 string                 `json:"question"`
	Options                  []string               `json:"options"`
	IsAnonymous              bool                   `json:"is_anonymous,omitempty"`
	Type                     string                 `json:"type,omitempty"`
	AllowsMultipleAnswers    bool                   `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID          int                    `json:"correct_option_id"`
	Explanation              string                 `json:"explanation,omitempty"`
	ExplanationParseMode     string                 `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []models.MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               int                    `json:"open_period,omitempty"`
	CloseDate                int                    `json:"close_date,omitempty"`
	IsClosed                 bool                   `json:"is_closed,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendDiceParams struct {
	ChatID                   any                `json:"chat_id"`
	Emoji                    string             `json:"emoji,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendChatActionParams struct {
	ChatID any    `json:"chat_id"`
	Action string `json:"action"`
}

type GetUserProfilePhotosParams struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type GetFileParams struct {
	FileID string `json:"file_id"`
}

type BanChatMemberParams struct {
	ChatID         any  `json:"chat_id"`
	UserID         int  `json:"user_id"`
	UntilDate      int  `json:"until_date,omitempty"`
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

type UnbanChatMemberParams struct {
	ChatID       any  `json:"chat_id"`
	UserID       int  `json:"user_id"`
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

type RestrictChatMemberParams struct {
	ChatID      any                     `json:"chat_id"`
	UserID      int                     `json:"user_id"`
	Permissions *models.ChatPermissions `json:"permissions,omitempty"`
	UntilDate   int                     `json:"until_date,omitempty"`
}

type PromoteChatMemberParams struct {
	ChatID             any  `json:"chat_id" rules:"required,chat_id"`
	UserID             int  `json:"user_id" rules:"required"`
	IsAnonymous        bool `json:"is_anonymous,omitempty"`
	CanManageChat      bool `json:"can_manage_chat,omitempty"`
	CanPostMessages    bool `json:"can_post_messages,omitempty"`
	CanEditMessages    bool `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool `json:"can_delete_messages,omitempty"`
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers  bool `json:"can_promote_members,omitempty"`
	CanChangeInfo      bool `json:"can_change_info,omitempty"`
	CanInviteUsers     bool `json:"can_invite_users,omitempty"`
	CanPinMessages     bool `json:"can_pin_messages,omitempty"`
}

type SetChatAdministratorCustomTitleParams struct {
	ChatID      any    `json:"chat_id"`
	UserID      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

type BanChatSenderChatParams struct {
	ChatID       any `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

type UnbanChatSenderChatParams struct {
	ChatID       any `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

type SetChatPermissionsParams struct {
	ChatID      any                    `json:"chat_id"`
	Permissions models.ChatPermissions `json:"permissions"`
}

type ExportChatInviteLinkParams struct {
	ChatID any `json:"chat_id"`
}

type CreateChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type EditChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type RevokeChatInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

type ApproveChatJoinRequestParams struct {
	ChatID any `json:"chat_id"`
	UserID int `json:"user_id"`
}

type DeclineChatJoinRequestParams struct {
	ChatID any `json:"chat_id"`
	UserID int `json:"user_id"`
}

type SetChatPhotoParams struct {
	ChatID any              `json:"chat_id"`
	Photo  models.InputFile `json:"photo"`
}

type DeleteChatPhotoParams struct {
	ChatID any `json:"chat_id"`
}

type SetChatTitleParams struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

type SetChatDescriptionParams struct {
	ChatID      any    `json:"chat_id"`
	Description string `json:"title"`
}

type PinChatMessageParams struct {
	ChatID              any  `json:"chat_id"`
	MessageID           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

type UnpinChatMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id,omitempty"`
}

type UnpinAllChatMessagesParams struct {
	ChatID any `json:"chat_id"`
}

type LeaveChatParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatAdministratorsParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatMemberCountParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatMemberParams struct {
	ChatID any `json:"chat_id"`
	UserID int `json:"user_id"`
}

type SetChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type DeleteChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

type SetMyCommandsParams struct {
	Commands     []models.BotCommand    `json:"commands"`
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type DeleteMyCommandsParams struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type GetMyCommandsParams struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type SetChatMenuButtonParams struct {
	ChatID     any                    `json:"chat_id"`
	MenuButton models.InputMenuButton `json:"menu_button"`
}

type GetChatMenuButtonParams struct {
	ChatID any `json:"chat_id"`
}

type SetMyDefaultAdministratorRightsParams struct {
	Rights      *models.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                            `json:"for_channels,omitempty"`
}

type GetMyDefaultAdministratorRightsParams struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

type EditMessageTextParams struct {
	ChatID                any                    `json:"chat_id"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Text                  string                 `json:"text"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	Entities              []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type EditMessageCaptionParams struct {
	ChatID                any                    `json:"chat_id"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Caption               string                 `json:"caption,omitempty"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type EditMessageMediaParams struct {
	ChatID          any                `json:"chat_id"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	Media           models.InputMedia  `json:"media"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkupParams struct {
	ChatID          any                `json:"chat_id"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type StopPollParams struct {
	ChatID      any                `json:"chat_id"`
	MessageID   int                `json:"message_id"`
	ReplyMarkup models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type DeleteMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id"`
}

type SendStickerParams struct {
	ChatID                   any                `json:"chat_id"`
	Sticker                  models.InputFile   `json:"sticker"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type GetStickerSetParams struct {
	Name string `json:"name"`
}

type UploadStickerFileParams struct {
	UserID     int              `json:"user_id"`
	PngSticker models.InputFile `json:"png_sticker"`
}

type CreateNewStickerSetParams struct {
	UserID        int                 `json:"user_id"`
	Name          string              `json:"name"`
	Title         string              `json:"title"`
	PngSticker    models.InputFile    `json:"png_sticker,omitempty"`
	TgsSticker    models.InputFile    `json:"tgs_sticker,omitempty"`
	WebmSticker   models.InputFile    `json:"webm_sticker,omitempty"`
	Emojis        string              `json:"emojis"`
	ContainsMasks bool                `json:"contains_masks,omitempty"`
	MaskPosition  models.MaskPosition `json:"mask_position,omitempty"`
}

type AddStickerToSetParams struct {
	UserID       int                 `json:"user_id"`
	Name         string              `json:"name"`
	PngSticker   models.InputFile    `json:"png_sticker,omitempty"`
	TgsSticker   models.InputFile    `json:"tgs_sticker,omitempty"`
	WebmSticker  models.InputFile    `json:"webm_sticker,omitempty"`
	Emojis       string              `json:"emojis"`
	MaskPosition models.MaskPosition `json:"mask_position,omitempty"`
}

type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"`
}

type SetStickerSetThumbParams struct {
	Name   string           `json:"name"`
	UserID int              `json:"user_id"`
	Thumb  models.InputFile `json:"thumb"`
}

type AnswerInlineQueryParams struct {
	InlineQueryID     string                     `json:"inline_query_id"`
	Results           []models.InlineQueryResult `json:"results"`
	CacheTime         int                        `json:"cache_time,omitempty"`
	IsPersonal        bool                       `json:"is_personal,omitempty"`
	NextOffset        string                     `json:"next_offset,omitempty"`
	SwitchPmText      string                     `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string                     `json:"switch_pm_parameter,omitempty"`
}

type AnswerWebAppQueryParams struct {
	WebAppQueryID string                   `json:"web_app_query_id"`
	Result        models.InlineQueryResult `json:"result"`
}

type SendInvoiceParams struct {
	ChatID                    any                   `json:"chat_id"`
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                `json:"start_parameter,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoURL                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int                   `json:"photo_size,omitempty"`
	PhotoWidth                int                   `json:"photo_width,omitempty"`
	PhotoHeight               int                   `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ProtectContent            bool                  `json:"protect_content,omitempty"`
	ReplyToMessageID          int                   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup               models.ReplyMarkup    `json:"reply_markup,omitempty"`
}

type CreateInvoiceLinkParams struct {
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoURL                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int                   `json:"photo_size,omitempty"`
	PhotoWidth                int                   `json:"photo_width,omitempty"`
	PhotoHeight               int                   `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
}

type AnswerShippingQueryParams struct {
	ShippingQueryID string                  `json:"shipping_query_id"`
	OK              bool                    `json:"ok"`
	ShippingOptions []models.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string                  `json:"error_message,omitempty"`
}

type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	OK                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

type SetPassportDataErrorsParams struct {
	UserID int                           `json:"user_id"`
	Errors []models.PassportElementError `json:"errors"`
}

type SendGameParams struct {
	ChatID                   any                `json:"chat_id"`
	GameShorName             string             `json:"game_short_name"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SetGameScoreParams struct {
	UserID             int  `json:"user_id"`
	Score              int  `json:"score"`
	Force              bool `json:"force,omitempty"`
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
	ChatID             int  `json:"chat_id,omitempty"`
	MessageID          int  `json:"message_id,omitempty"`
	InlineMessageID    int  `json:"inline_message_id,omitempty"`
}

type GetGameHighScoresParams struct {
	UserID          int `json:"user_id"`
	ChatID          int `json:"chat_id,omitempty"`
	MessageID       int `json:"message_id,omitempty"`
	InlineMessageID int `json:"inline_message_id,omitempty"`
}
