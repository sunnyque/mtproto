package mtproto

const (
	PEER_TYPE_USER    = "USER"
	PEER_TYPE_CHAT    = "CHAT"
	PEER_TYPE_CHANNEL = "CHANNEL"
)

const (
	DIALOG_TYPE_CHAT    = "CHAT"
	DIALOG_TYPE_USER    = "USER"
	DIALOG_TYPE_CHANNEL = "CHANNEL"
)

const (
	USER_STATUS_OFFLINE    = "OFFLINE"
	USER_STATUS_ONLINE     = "ONLINE"
	USER_STATUS_RECENTLY   = "RECENTLY"
	USER_STATUS_LAST_WEEK  = "LAST_WEEK"
	USER_STATUS_LAST_MONTH = "LAST_MONTH"
)

const (
	// Message Types
	MESSAGE_TYPE_EMPTY     = "EMPTY"
	MESSAGE_TYPE_NORMAL    = "NORMAL"
	MESSAGE_TYPE_SERVICE   = "SERVICE"
	MESSAGE_TYPE_FORWARDED = "FORWARDED"

	// Message Actions
	MESSAGE_ACTION_CHAT_CREATED         = "CHAT_CREATED"
	MESSAGE_ACTION_CHAT_EDIT_TITLE      = "CHAT_EDIT_TITLE"
	MESSAGE_ACTION_CHAT_EDIT_PHOTO      = "CHAT_EDIT_PHOTO"
	MESSAGE_ACTION_CHAT_DELETE_PHOTO    = "CHAT_DELETE_PHOTO"
	MESSAGE_ACTION_CHAT_ADD_USER        = "CHAT_ADD_USER"
	MESSAGE_ACTION_CHAT_DELETE_USER     = "CHAT_DELETE_USER"
	MESSAGE_ACTION_CHAT_JOINED_BY_LINK  = "CHAT_JOINED_BY_LINK"
	MESSAGE_ACTION_CHAT_MIGRATE_TO      = "CHAT_MIGRATE_TO"
	MESSAGE_ACTION_CHANNEL_CREATED      = "CHANNEL_CREATED"
	MESSAGE_ACTION_CHANNEL_MIGRATE_FROM = "CHANNEL_MIGRATE"
	MESSAGE_ACTION_GAME_SCORE           = "GAME_SCORE"
	MESSAGE_ACTION_HISTORY_CLEAN        = "HISTORY_CLEAN"

	// Message Entities
	MESSAGE_ENTITY_UNKNOWN      = "UNKNOWN"
	MESSAGE_ENTITY_HASHTAG      = "HASHTAG"
	MESSAGE_ENTITY_URL          = "URL"
	MESSAGE_ENTITY_MENTION      = "MENTION"
	MESSAGE_ENTITY_MENTION_NAME = "MENTION_NAME"
	MESSAGE_ENTITY_EMAIL        = "EMAIL"
	MESSAGE_ENTITY_BOLD         = "BOLD"
	MESSAGE_ENTITY_ITALIC       = "ITALIC"
	MESSAGE_ENTITY_CODE         = "CODE"
	MESSAGE_ENTITY_BOT_COMMAND  = "BOT_COMMAND"
	MESSAGE_ENTITY_PRE          = "PRE"
	MESSAGE_ENTITY_TEXT_URL     = "TEXT_URL"

	// Message Media Types
	MESSAGE_MEDIA_TYPE_EMPTY    = "EMPTY"
	MESSAGE_MEDIA_TYPE_PHOTO    = "PHOTO"
	MESSAGE_MEDIA_TYPE_VIDEO    = "VIDEO"
	MESSAGE_MEDIA_TYPE_GEO      = "GEO"
	MESSAGE_MEDIA_TYPE_CONTACT  = "CONTACT"
	MESSAGE_MEDIA_TYPE_DOCUMENT = "DOCUMENT"
	MESSAGE_MEDIA_TYPE_AUDIO    = "AUDIO"
)

const ()

const ()
