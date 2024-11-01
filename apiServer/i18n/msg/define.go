package i18n_msg

type MsgKey = string

var (
	TextMsg MsgKey = "text_msg"

	JoinRoomMsgKey           = "join_room"                   //加入房间
	LeaveRoomMsgKey          = "leave_room"                  //离开房间
	FollowJoinRoomMsgKey     = "follow_join_room"            //踩着%v的小尾巴进入房间~
	CompereMicMsgKey         = "compere_mic"                 //主持位
	GuestMicMsgKey           = "guest_mic"                   //嘉宾位
	MicSeatMsgKey            = "mic_seat"                    //几号麦
	UpMsgKey                 = "up"                          //上
	DownMsgKey               = "down"                        //下
	YouHaveBeenMutedMsgKey   = "you_have_been_mute"          //您被禁言?分钟
	LikedYourPostMsgKey      = "liked_your_post"             //点赞了你的动态
	RepliedYourCommentMsgKey = "replied_your_comment"        //回复了你的评论
	CommentedYourPostMsgKey  = "commented_your_post"         //评论的你的动态
	JustNowMsgKey            = "just_now"                    //刚刚
	MinutesAgoMsgKey         = "minutes_ago"                 //?分钟前
	VisitMinuteMeMsgKey      = "visit_minute_me"             //我在*分钟前访问
	VisitMinuteHeMsgKey      = "visit_minute_he"             //ta在*分钟前访问
	VisitHourMeMsgKey        = "visit_hour_me"               //我在*小时前访问
	VisitHourHeMsgKey        = "visit_hour_he"               //ta在*小时前访问
	VisitDayMeMsgKey         = "visit_day_me"                //我在*天前访问
	VisitDayHeMsgKey         = "visit_day_he"                //ta在*天前访问
	VisitMonthMeMsyKey       = "visit_month_me"              // 我在*月*日访问
	VisitMonthHeMsyKey       = "visit_month_he"              // ta在*月*日访问
	VisitYearMeMsyKey        = "visit_year_me"               // 我在*年*月访问
	VisitYearHeMsyKey        = "visit_year_he"               // ta在*年*月访问
	VisitExtraMeMsgKey       = "visit_extra_me"              //你也看过ta
	VisitExtraHeMsgKey       = "visit_extra_he"              //ta也看过你
	ChatRoomKey              = "chat_name"                   //聊天室
	CastRoomKey              = "cast_room"                   //直播间
	PersonalRoomKey          = "personal_room"               //个人房
	JoinRoomFromOtherClient  = "join_room_from_other_client" //在其他房间进入聊天室
	OtherClientInRoom        = "other_client_in_room"        //其他端用户已在房间
	Diamond                  = "diamond"                     //钻石
	RechargeDiamond          = "recharge_diamond"            //充值钻石
	CollectKey               = "collect"                     //收藏
	RecommendedKey           = "recommended"                 //推荐
	EmotionKey               = "emotion"                     //情感
	LiveKey                  = "live"                        //直播
	SingKey                  = "sing"                        //唱歌
	EmotionManKey            = "emotion_man"                 //情感男
	EmotionWomanKey          = "emotion_woman"               //情感女
	LiveVideoKey             = "live_video"                  //视频直播
	LiveVoiceKey             = "live_voice"                  //语音直播
	MondayKey                = "monday"                      //周一
	TuesdayKey               = "tuesday"                     //周二
	WednesdayKey             = "wednesday"                   //周三
	ThursdayKey              = "thursday"                    //周四
	FridayKey                = "friday"                      //周五
	SaturdayKey              = "saturday"                    //周六
	SundayKey                = "sunday"                      //周日
	UnknownKey               = "unknown"                     //未知
	NewPostPublishedKey      = "new_post_published"          //发布了一条新动态，快去看看吧
	LiveStreamCurrentlyKey   = "live_stream_currently"       //正在直播中，快去围观
	LockRoomKey              = "lock_room"                   //锁定房间
	UnLockRoomKey            = "unlock_room"                 //解锁房间
	OpenRoomKey              = "open_room"                   //开启房间
	CloseRoomKe              = "close_room"                  //关闭房间
	NicknameRejectKey        = "nickname_reject"             //昵称违规
	NicknameRejectContextKey = "nickname_reject_context"     //昵称违规内容
	AvatarRejectKey          = "avatar_reject"               //头像违规
	AvatarRejectContextKey   = "avatar_reject_context"       //头像违规内容
)
