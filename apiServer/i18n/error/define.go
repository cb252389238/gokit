package i18n_err

type ErrCode = int

var (
	SuccessCode ErrCode = 0 //成功
)

var (
	ErrorCodeParam                      ErrCode = 10000 //参数错误
	ErrorCodeToken                      ErrCode = 10001 //token错误
	ErrorCodeUserNotFound               ErrCode = 10002 //用户不存在
	ErrorCodeUserPassError              ErrCode = 10003 //密码有误
	ErrorCodeRepeatSubmit               ErrCode = 10004 //重复提交
	ErrorCodePassEasy                   ErrCode = 10005 //密码过于简单
	ErrorCodePassEq                     ErrCode = 10006 //新旧密码相同
	ErrorCodeNicknameRepeat             ErrCode = 10007 //用户昵称重复
	ErrorCodeUserFreezing               ErrCode = 10008 //用户已冻结
	ErrorCodeUserApplyInvalid           ErrCode = 10009 //用户已申请注销
	ErrorCodeUserInvalid                ErrCode = 10010 //用户已注销
	ErrorCodeIsPraiseTimeline           ErrCode = 10011 //已点赞过动态
	ErrorCodeNotPraiseTimeline          ErrCode = 10012 //没有点赞过动态
	ErrorCodeIsBlacklist                ErrCode = 10013 //用户被拉黑
	ErrorCodeCaptchaInvalid             ErrCode = 10014 //验证码错误
	ErrorCodeGetCaptcha                 ErrCode = 10015 //验证码获取失败
	ErrorCodeCaptchaExpre               ErrCode = 10016 //短信已经发送，请勿重复提交
	ErrorCodeSexEditNum                 ErrCode = 10017 //性别修改次数超过限制
	ErrorCodeNicknameEditNum            ErrCode = 10018 //本月昵称修改已达上限
	ErrorCodeVoiceEditNum               ErrCode = 10019 //本月声音修改已达上限
	ErrorCodeRepeatNickname             ErrCode = 10020
	ErrorCodeRequestAbnormal            ErrCode = 10021
	ErrorCodeSendGiftSuperAdmin         ErrCode = 10022 // 不能给超管或巡查送礼物
	ErrorCodeDiamondNotEnough           ErrCode = 10023 // 钻石余额不足
	ErrorCodeStarlightExchangeNotEnough ErrCode = 10024 // 星光兑换余额不足
	ErrorCodeSendGiftFromSuperAdmin     ErrCode = 10025
	ErrorCodeNewMobileOldMobileSame     ErrCode = 10026 //新旧手机号一致
	ErrorCodeMobileUsed                 ErrCode = 10027 //该手机号码已绑定其他账号

	ErrorCodeUserAlreadyFollow       ErrCode = 11000 //用户已关注
	ErrorCodeUserNotFollow           ErrCode = 11001 //用户未关注
	ErrorCodeGuildNotExist           ErrCode = 11002 //公会不存在
	ErrorCodeGuildAlreadyJoin        ErrCode = 11003 //已加入公会或已是其他公会成员
	ErrorCodeGuildAlreadyApplication ErrCode = 11004 //你已申请加入公会，请勿重复申请
	ErrorCodeGuildNotLeader          ErrCode = 11005 //不是会长
	ErrorCodeLvConfigNotExist        ErrCode = 11006 //等级配置不存在
	ErrorCodeLvMaxConfigNotExist     ErrCode = 11007 //最大等级配置不存在
	ErrorCodeLvNotExist              ErrCode = 11008 //用户等级不存在
	ErrorCodeRoomNotExist            ErrCode = 11009 //房间不存在
	ErrorCodeNotRoomOwner            ErrCode = 11010 //您不是房主，无权设置
	ErrorCodeUserIsAdmin             ErrCode = 11011 //该用户已经是管理员
	ErrorCodeUserNotInGuild          ErrCode = 11012 //用户不是该公会成员
	ErrorAccountIsBanned             ErrCode = 11013 //您的账号已被封禁,无法登陆！
	ErrorCodeUserBankNotExist        ErrCode = 11014 //请先绑定银行卡
	ErrorCodeUserAmountNotEnough     ErrCode = 11015 //不足100元不可提现
	ErrorCodeUserAmountMaxErr        ErrCode = 11016 //单笔提现金额上限为20000元人民币
	ErrorCodeUserAmountInvalid       ErrCode = 11017 //提现金额为100的倍数，如100、200、300
	ErrorCodeUserWithdrawDay         ErrCode = 11018 //今日不是结算日，无法提现
	ErrorCodeUserWithdrawDayMaxErr   ErrCode = 11019 //今日已提现，无法提现
	ErrorCodeUserAmountExceed        ErrCode = 11020 //输入金额超过可提现金额
	ErrorCodeBankCardExist           ErrCode = 11021 //该银行卡已存在
	ErrorCodeBankCardInfoErr         ErrCode = 11022 //银行卡信息获取错误
	ErrorCodeUserBankNotBind         ErrCode = 11023 //该用户未绑定该银行卡
	ErrorCodeIdCardNotExist          ErrCode = 11024 //	身份证号信息有误
	ErrorCodeAlreadyApplyRealName    ErrCode = 11025 //用户已申请实名认证
	ErrorCodeQuitGuildApplyExamine   ErrCode = 11026 //	已申请退出公会，请耐心等待！
	ErrorCodeForcedQuitGuildFee      ErrCode = 11027 //	您是从业者，请先缴纳违约金！
	ErrorCodeUserIsRoomOwner         ErrCode = 11028 //该成员有房主身份，请先取消对应身份后再进行操作！

	ErrorCodeUserPractitioner           ErrCode = 12000 //题库正在维护中，请稍后再试
	ErrorCodeUserPractitionerAwnser     ErrCode = 12001 //答案数据有误
	ErrorCodeIDCardAuth                 ErrCode = 12002 //用户未实名认证
	ErrorCodePractitionerResult         ErrCode = 12003 //获取从业者申请结果失败
	ErrorCodeIsTrue                     ErrCode = 12004 //记录正在审核或者已经有次身份了
	ErrorCodeCerdIsStatus               ErrCode = 12005 //已有当前身份,不需要考核
	ErrCodeCollect                      ErrCode = 12006 //收藏失败
	ErrCodeDelCollect                   ErrCode = 12007 //取消收藏失败
	ErrorCodeTooManysRequest            ErrCode = 12008 //请求太频繁
	ErrorCodePractitionerAnswer         ErrCode = 12009 //今日考试已达3次上限，请明天再来
	ErrorCodeRoomStatus                 ErrCode = 12010 //房间状态异常
	ErrCodeRoomPwdErr                   ErrCode = 12011 //房间密码错误
	ErrCodeRoomBlacklistErr             ErrCode = 12012 //没有权限进入
	ErrCodeRoomKickOutErr               ErrCode = 12013 //被房间踢出
	ErrCodeJoinRoomErr                  ErrCode = 12014 //加入房间失败
	ErrCodeLeaveRoomErr                 ErrCode = 12015 //退出房间失败
	ErrCodeOnlineUserErr                ErrCode = 12017 // 拉取在线用户列表失败
	ErrCodeDayUserErr                   ErrCode = 12018 // 拉取1000贡献榜列表失败
	ErrCodeRoomUpdateErr                ErrCode = 12019 //编辑房间失败
	ErrCodeKickOutErr                   ErrCode = 12020 //踢出用户失败
	ErrCodeReportingCenterErr           ErrCode = 12021 //举报失败
	ErrCodeAutoWelcomeErr               ErrCode = 12022 //设置自动欢迎语失败
	ErrCodeNotAnchor                    ErrCode = 12023 //您不是主播无法创建直播间
	ErrCodeBlackErr                     ErrCode = 12024 //您已被房间{{.roomName}}拉黑，不能进入房间
	ErrCodeBlackOutTimesErr             ErrCode = 12025 //被踢出了,还剩{{.times}}分钟能进入房间
	ErrCodeRoomAnchorErr                ErrCode = 12026 //房间未开播
	ErrCodeBlackListAddErr              ErrCode = 12027 //拉黑失败
	ErrCodeBlackListDelErr              ErrCode = 12028 //取消拉黑失败
	ErrCodeUserGoodsExpireErr           ErrCode = 12029 //商品已过期
	ErrCodeUserGoodsUseErr              ErrCode = 12030 //装扮使用失败
	ErrCodeUserGoodsDelErr              ErrCode = 12031 //删除装扮失败
	ErrCodeBlackOutImErr                ErrCode = 12032 //您被踢出{{.times}}分钟,请稍后再试
	ErrCodeHighGradeUsersErr            ErrCode = 12033 // 拉取高等级用户列表失败
	ErrCodeRoomMicErr                   ErrCode = 13001 //房间麦位异常
	ErrCodeRoomPermissionDenied         ErrCode = 13002 //房间操作无权限
	ErrCodeRoomSeatClosed               ErrCode = 13003 //房间麦位已关闭
	ErrCodeRoomSeatUsed                 ErrCode = 13004 //房间麦位使用中
	ErrCodeRoomCompereSeatEmpty         ErrCode = 13005 //主持不在麦
	ErrCodeRoomUserNotFound             ErrCode = 13006 //用户不在房间内
	ErrCodeRoomSeatUp                   ErrCode = 13007 //您已在麦位上
	ErrCodeRoomSeatFilled               ErrCode = 13008 //房间麦位已满
	ErrCodeAvTokenFailed                ErrCode = 13009 //获取声网token失败
	ErrCodeUserNotCompere               ErrCode = 13010 //该用户不是主持人
	ErrCodeFreeUpMicNotOpen             ErrCode = 13011 //自由上下麦没有开启
	ErrCodeUserNotOnSeat                ErrCode = 13012 //用户不在麦位上
	ErrCodeUserNotInRoom                ErrCode = 13013 //用户不在房间内
	ErrCodeUserMuteFail                 ErrCode = 13014 //禁言失败
	ErrCodeUserUnMuteFail               ErrCode = 13015 //取消禁言失败
	ErrCodeUserIsMute                   ErrCode = 13016 //用户已禁言
	ErrCodeUserMuteMsg                  ErrCode = 13017 //您被禁言？分钟
	ErrCodeHeMuteMsg                    ErrCode = 13019 //TA被禁言？分钟
	ErrCodeHiddenMicHasUser             ErrCode = 13020 //隐藏麦有人
	ErrCodeBeforeDownHiddenMic          ErrCode = 13021 //请先下隐藏麦
	ErrCodeUserCantNotOperate           ErrCode = 13022 //用户不可操作
	ErrCodeAgentSignErr                 ErrCode = 13023 //代理签名错误
	ErrCodeUserRealName                 ErrCode = 13024 //实名错误
	ErrCodeUserAccountExceed            ErrCode = 13025 //用户账号数量超出
	ErrCodeSwitchUserAccount            ErrCode = 13026 //切换账户失败
	ErrCodeSwitchUserAccountTokenExpire ErrCode = 13027 //切换账号授权已过期
	ErrCodePasswordUnlike               ErrCode = 13028 //两次密码不一致
	ErrCodePasswordRepeat               ErrCode = 13029 //和您的其他账号密码重复
	ErrCodeMobileOrRegionCode           ErrCode = 13030 //手机号或者区域码错误
	ErrCodeRoomRelateWheat              ErrCode = 13031 //连麦模板不支持
	ErrCodeOtherClientInRoom            ErrCode = 13032 //其他端在房间内
	ErrCodeUserPasswordNotSet           ErrCode = 13033 //用户密码未设置
	ErrCodeFreedSpeakClosed             ErrCode = 13034 //自由发言模式已关闭
	ErrCodeOnMicSeat                    ErrCode = 13035 //已在麦位请先下麦
	ErrCodeNotHaveRoomBg                ErrCode = 13036 //您还未获得该背景图，无法更换！
	ErrCodeRoomCompereSeatClosed        ErrCode = 13037 //主持人麦位已关闭，不能申请上麦
)

var (
	ErrorCodeReadDB       ErrCode = 20000 // 数据库读取失败
	ErrorCodeUpdateDB     ErrCode = 20001 //更新数据库失败
	ErrorCodeDataNotFound ErrCode = 20002 //数据不存在

)

var (
	ErrorCodeCreateToken            ErrCode = 30000 //token生成错误
	ErrorCodeRedisAction            ErrCode = 30001 //redis错误
	ErrorCodeOperationFail          ErrCode = 30002 //操作失败
	ErrorCodeNicknameCheckReject    ErrCode = 30003 //昵称检测违规
	ErrorCodeAvatarCheckReject      ErrCode = 30004 //头像检测违规
	ErrorCodeSignCheckReject        ErrCode = 30005 //签名违规
	ErrorCodeTextCheckReject        ErrCode = 30006 //消息内容违规
	ErrorCodeImageCheckReject       ErrCode = 30007 //图片违规
	ErrorCodePrivateChatIsBlacklist ErrCode = 30008 //您的消息已被对方拒收
)

var (
	ErrorCodeAppleIapError   ErrCode = 40000 //苹果内购支付失败
	ErrorCodeWxPayError      ErrCode = 40001 //微信支付错误
	ErrorCodeAliPayError     ErrCode = 40002 //支付宝支付错误
	ErrorCodePayChannelError ErrCode = 40003 //获取支付渠道失败
	ErrorCodeProductNotFound ErrCode = 40004 //商品不存在
)

var (
	ErrorCodeUserLoginCheckMoreAccount ErrCode = 99993 //用户拥有多账号
	ErrorCodeUserLoginOtherClient      ErrCode = 99994 //用户在其他端登录
	ErrorCodeUserImServer              ErrCode = 99995 //im服务分配错误
	ErrorCodeUserNoPermissions         ErrCode = 99996 //权限不足
	ErrorCodeOperationFrequent         ErrCode = 99997 //操作频繁
	ErrorCodeSystemBusy                ErrCode = 99998 //系统繁忙
	ErrorCodeUnknown                   ErrCode = 99999 //未知错误
)

// 公会后台&房主后台
var (
	ErrorCodeTokenInvalid             ErrCode = 50000 // token失效
	ErrorCodeNotExistGuild            ErrCode = 50001 // 公会不存在
	ErrorCodeCheckMobile              ErrCode = 50002 // 手机号错误，请检查手机号！
	ErrorCodeLoginFreezing            ErrCode = 50003 // 账号已被封禁，请联系平台客服！
	ErrorCodeNoRoomID                 ErrCode = 50004 // 房间ID不存在
	ErrorCodeGuildRoomMaxNum          ErrCode = 50005 //公会房间上限
	ErrorCodeDaySettleInvalidRole     ErrCode = 50006 // 日结算人不是房主/会长
	ErrorCodeMonthSettleInvalidRole   ErrCode = 50007 // 月结算人不是房主/会长
	ErrorCodeRoomTypeNotExist         ErrCode = 50008 // 房间类型不存在
	ErrorCodeRoomTemplateNotExist     ErrCode = 50009 // 房间模板不存在
	ErrorCodeMemberApplyNotExist      ErrCode = 50010 // 成员申请信息不存在
	ErrorCodeRoomOwner                ErrCode = 50011 //当前用户是房主,不能踢出
	ErrorCodeRoomGuildOwner           ErrCode = 50012 //当前用户是会长,不能踢出
	ErrorCodeRoomApplyExist           ErrCode = 50013 // 你已申请房间,不能重复申请
	ErrorCodeWithdrawConfigNotExist   ErrCode = 50014 //提现配置信息不存在
	ErrorCodeSubsidyInfoNotExist      ErrCode = 50015 //补贴星光信息不存在
	ErrorCodeNotPractitionerCredExist ErrCode = 50016 //该用户无从业者资质，无法添加！
	ErrorCodeNotGuildMember           ErrCode = 50017 //该用户不在本公会，无法添加！
	ErrorCodeCheckUserID              ErrCode = 50018 //请输入正确的用户ID
	ErrorCodeIsPractitionerExist      ErrCode = 50019 //从业者已有该身份，无法添加！
	ErrorCodePractitionerExamine      ErrCode = 50020 //该从业者正在审核中，请耐心等待！
	ErrorCodeNotRoomPractitioner      ErrCode = 50021 //不是本房间从业者
	ErrorCodeRoomFreezing             ErrCode = 50022 //房间被平台封禁无法开启，请联系平台客服处理！
	ErrorCodeGuildApplyStatus         ErrCode = 50023 //非待审核状态，请从新载入！
	ErrorCodeGuildApplyNotRefused     ErrCode = 50024 //该申请不能拒绝！
	ErrorCodeCheckCaptcha             ErrCode = 50025 // 请输入正确的验证码！
)

var (
	ErrorCodeGroupIsMute ErrCode = 60000 //禁止发言
)
