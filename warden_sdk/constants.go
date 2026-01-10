package warden_sdk

// 命令常量 - Command Constants
const (
	// 设备属性类 - Device Properties (0x50-0x5D)
	CMD_SYNC_TIME    = 0x50 // 同步时间
	CMD_BATTERY      = 0x51 // 电量获取
	CMD_BRIGHTNESS   = 0x52 // 屏幕亮度
	CMD_LANGUAGE     = 0x53 // 设备语言
	CMD_DND          = 0x57 // 勿扰功能
	CMD_FIND_PHONE   = 0x59 // 寻找手机
	CMD_WEATHER_UNIT = 0x5A // 天气单位设置
	CMD_TIME_FORMAT  = 0x5B // 12h/24h时间制切换
	CMD_DEVICE_INFO  = 0x5C // 设备端信息
	CMD_APP_INFO     = 0x5D // 应用端信息

	// 状态控制类 - Status Control (0x80-0x91)
	CMD_SWITCH_CONTROL   byte = 0x80 // 开关设置
	CMD_DEVICE_BINDING   byte = 0x81 // 设备绑定
	CMD_BINDING_RESERVED byte = 0x82 // 绑定(预留)
	CMD_SWITCH_EXTEND    byte = 0x86 // 开关表扩展
	CMD_REMOTE_CAMERA    byte = 0x91 // 远程拍照

	// 消息推送类 - Message Push (0xA0-0xA1)
	CMD_MESSAGE_PUSH byte = 0xA0 // 消息推送
	CMD_WEATHER_INFO byte = 0xA1 // 天气信息
)

// 命令类型常量 - Command Type Constants
const (
	CMD_TYPE_REQUEST    = 0x01 // Request: 手机 -> 设备
	CMD_TYPE_RESPONSE   = 0x02 // Response: 设备 -> 手机
	CMD_TYPE_NOTIFY     = 0x03 // Notify: 主动上报，不需要对端回复
	CMD_TYPE_REQUEST_N  = 0x04 // Request_n: 设备 -> 手机
	CMD_TYPE_RESPONSE_N = 0x05 // Response_n: 手机 -> 设备
	CMD_TYPE_NOTIFY_N   = 0x06 // Notify_n: 主动上报，不需要对端回复
)

// Action命令常量 - Action Command Constants
const (
	ACTION_QUERY = 0x00 // 查询
	ACTION_SET   = 0x01 // 设置
)

// 返回值常量 - Response Code Constants (表7-1)
const (
	CMD_RESULT_OK          = 0x00 // 成功
	CMD_RESULT_FAIL        = 0x01 // 失败
	CMD_RESULT_LEN_INVALID = 0x02 // 包长度不合法
	CMD_RESULT_CMD_INVALID = 0x03 // 命令类型不合法
	CMD_RESULT_IDX_INVALID = 0x04 // 编号索引不合法
	CMD_RESULT_NOT_SUPPORT = 0x05 // 设备不支持该命令
	CMD_RESULT_SWI_ERROR   = 0x06 // 开关操作不合法
	CMD_RESULT_CHECK_ERROR = 0x07 // 数据校验错误
	CMD_RESULT_PACKET_LOSS = 0x08 // 上一包数据丢失
)

// 语言常量 - Language Constants (表7-2-1)
const (
	LANG_ENGLISH             = 0x00
	LANG_CHINESE_SIMPLIFIED  = 0x01
	LANG_ITALIAN             = 0x02
	LANG_SPANISH             = 0x03
	LANG_PORTUGUESE          = 0x04
	LANG_RUSSIAN             = 0x05
	LANG_JAPANESE            = 0x06
	LANG_CHINESE_TRADITIONAL = 0x07
	LANG_GERMAN              = 0x08
	LANG_KOREAN              = 0x09
	LANG_THAI                = 0x0A
	LANG_ARABIC              = 0x0B
	LANG_TURKISH             = 0x0C
	LANG_FRENCH              = 0x0D
	LANG_VIETNAMESE          = 0x0E
	LANG_POLISH              = 0x0F
	LANG_DUTCH               = 0x10
	LANG_HEBREW              = 0x11
	LANG_PERSIAN             = 0x12
	LANG_GREEK               = 0x13
	LANG_MALAYSIAN           = 0x14
	LANG_BURMESE             = 0x15
	LANG_DANISH              = 0x16
	LANG_UKRAINIAN           = 0x17
	LANG_INDONESIAN          = 0x18
	LANG_CZECH               = 0x19
	LANG_HINDI               = 0x20
)

// 消息类型常量 - Message Type Constants (表7-3-1)
const (
	MSG_NULL               = 0x00
	MSG_INCOMING_CALL      = 0x01 // 来电
	MSG_MISSED_CALL        = 0x02 // 未接来电
	MSG_SMS                = 0x03 // 短信
	MSG_EMAIL              = 0x04 // 邮件
	MSG_SCHEDULE           = 0x05 // 日程
	MSG_FACETIME           = 0x06
	MSG_QQ                 = 0x07
	MSG_SKYPE              = 0x08
	MSG_WECHAT             = 0x09 // 微信
	MSG_WHATSAPP           = 0x0A
	MSG_GMAIL              = 0x0B
	MSG_HANGOUT            = 0x0C
	MSG_INBOX              = 0x0D
	MSG_LINE               = 0x0E
	MSG_TWITTER            = 0x0F // 推特
	MSG_FACEBOOK           = 0x10 // 脸书
	MSG_FACEBOOK_MESSENGER = 0x11
	MSG_INSTAGRAM          = 0x12
	MSG_WEIBO              = 0x13 // 微博
	MSG_KAKAOTALK          = 0x14
	MSG_FACEBOOK_MANAGER   = 0x15
	MSG_VIBER              = 0x16
	MSG_VKCLIENT           = 0x17
	MSG_TELEGRAM           = 0x18
	MSG_SNAPCHAT           = 0x1A
	MSG_DINGTALK           = 0x1B // 钉钉
	MSG_ALIPAY             = 0x1C // 支付宝
	MSG_TIKTOK             = 0x1D // 抖音
	MSG_LINKEDIN           = 0x1E // 领英
)

// 开关类型常量 - Switch Type Constants (表7-4-1)
const (
	SWITCH_ANTI_LOST     = 0x00 // 防丢开关
	SWITCH_RAISE_WRIST   = 0x01 // 抬手亮屏开关
	SWITCH_AUTO_SYNC     = 0x02 // 自动同步开关
	SWITCH_SLEEP_MONITOR = 0x04 // 睡眠监测开关
	SWITCH_MSG_TOTAL     = 0x05 // 消息提醒总开关
	SWITCH_HOURLY_SPORT  = 0x06 // 整点上传运动数据开关
	SWITCH_GOAL_ACHIEVED = 0x07 // 目标达成开关
	SWITCH_MSG_SCREEN    = 0x09 // 消息提醒亮屏开关
	SWITCH_SOUND         = 0x0A // 声音开关
	SWITCH_VIBRATE_TOTAL = 0x0B // 震动总开关
	SWITCH_HOURLY_HEALTH = 0x0C // 整点上传健康数据开关
	SWITCH_MSG_VIBRATE   = 0x0D // 消息提醒震动开关
)

// 天气类型常量 - Weather Type Constants
const (
	WEATHER_CLOUDY    = 0x00 // 多云
	WEATHER_SUNSHINE  = 0x01 // 晴天
	WEATHER_SNOW      = 0x02 // 雪天
	WEATHER_RAIN      = 0x03 // 雨天
	WEATHER_OVERCAST  = 0x04 // 阴天
	WEATHER_SAND_DUST = 0x05 // 沙尘天气
	WEATHER_WINDY     = 0x06 // 大风天气
	WEATHER_HAZE      = 0x07 // 阴霾天气
)

// 天气命令类型 - Weather Command Type
const (
	WEATHER_CMD_TYPE         = 0x00 // 天气类型
	WEATHER_CMD_CURRENT_TEMP = 0x01 // 当前温度
	WEATHER_CMD_MIN_TEMP     = 0x02 // 最低温度
	WEATHER_CMD_MAX_TEMP     = 0x03 // 最高温度
	WEATHER_CMD_LOCATION     = 0x04 // 地点位置
	WEATHER_CMD_HUMIDITY     = 0x05 // 湿度
	WEATHER_CMD_WIND_SPEED   = 0x06 // 风速
	WEATHER_CMD_UV_INDEX     = 0x07 // 紫外线指数
	WEATHER_CMD_PRESSURE     = 0x08 // 气压
	WEATHER_CMD_SUNRISE      = 0x09 // 日出时间
	WEATHER_CMD_SUNSET       = 0x0A // 日落时间
	WEATHER_CMD_RAIN_PROB    = 0x0B // 降雨概率
	WEATHER_CMD_RAINFALL     = 0x0C // 降雨量
	WEATHER_CMD_VISIBILITY   = 0x0D // 能见度
	WEATHER_CMD_AIR_QUALITY  = 0x0E // 空气质量指数
)

// 设备类型 - Device Type
const (
	DEVICE_TYPE_SQUARE   = 0x00 // 方形手表
	DEVICE_TYPE_ROUND    = 0x01 // 圆形手表
	DEVICE_TYPE_BRACELET = 0x02 // 手环
)

// 手机类型 - Phone Type
const (
	PHONE_TYPE_ANDROID = 0x00 // 安卓手机
	PHONE_TYPE_APPLE   = 0x01 // 苹果手机
)

// 时区常量 - Time Zone Constants
const (
	TIMEZONE_WEST_12 byte = 0x00 // 西十二区
	TIMEZONE_WEST_11 byte = 0x01
	TIMEZONE_WEST_10 byte = 0x02
	TIMEZONE_WEST_9  byte = 0x03
	TIMEZONE_WEST_8  byte = 0x04
	TIMEZONE_WEST_7  byte = 0x05
	TIMEZONE_WEST_6  byte = 0x06
	TIMEZONE_WEST_5  byte = 0x07
	TIMEZONE_WEST_4  byte = 0x08
	TIMEZONE_WEST_3  byte = 0x09
	TIMEZONE_WEST_2  byte = 0x0A
	TIMEZONE_WEST_1  byte = 0x0B
	TIMEZONE_ZERO    byte = 0x0C // 中时区/零时区
	TIMEZONE_EAST_1  byte = 0x0D
	TIMEZONE_EAST_2  byte = 0x0E
	TIMEZONE_EAST_3  byte = 0x0F
	TIMEZONE_EAST_4  byte = 0x10
	TIMEZONE_EAST_5  byte = 0x11
	TIMEZONE_EAST_6  byte = 0x12
	TIMEZONE_EAST_7  byte = 0x13
	TIMEZONE_EAST_8  byte = 0x14
	TIMEZONE_EAST_9  byte = 0x15
	TIMEZONE_EAST_10 byte = 0x16
	TIMEZONE_EAST_11 byte = 0x17
	TIMEZONE_EAST_12 byte = 0x18
)

// 温度单位 - Temperature Unit
const (
	TEMP_UNIT_CELSIUS    = 0x00 // 摄氏度
	TEMP_UNIT_FAHRENHEIT = 0x01 // 华氏度
)

// 时间制 - Time Format
const (
	TIME_FORMAT_12H = 0x00 // 12小时制
	TIME_FORMAT_24H = 0x01 // 24小时制
)

// BLE广播UUID
const (
	SERVICE_UUID          = "6e400001-b5a3-f393-e0a9-e50e24dcca9d"
	CHARACTERISTIC_WRITE  = "6e400002-b5a3-f393-e0a9-e50e24dcca9d" // APP写数据到设备
	CHARACTERISTIC_NOTIFY = "6e400003-b5a3-f393-e0a9-e50e24dcca9d" // 设备Notify数据给APP
)

// 协议常量
const (
	HEADER_SIZE      = 5   // 协议头大小
	DEFAULT_MTU      = 244 // 默认MTU值
	MAX_SEQ_NUM      = 15  // 最大序列号
	MAX_FRAME_SEQ    = 15  // 最大帧序号
	MAX_TOTAL_FRAMES = 16  // 最大总帧数
)
