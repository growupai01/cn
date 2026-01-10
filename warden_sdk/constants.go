package warden_sdk

// 命令常量 - Command Constants
const (
	// 设备属性类 - Device Properties (0x50-0x5D)
	CMD_SYNC_TIME    byte = 0x50 // 同步时间
	CMD_BATTERY      byte = 0x51 // 电量获取
	CMD_BRIGHTNESS   byte = 0x52 // 屏幕亮度
	CMD_LANGUAGE     byte = 0x53 // 设备语言
	CMD_DND          byte = 0x57 // 勿扰功能
	CMD_FIND_PHONE   byte = 0x59 // 寻找手机
	CMD_WEATHER_UNIT byte = 0x5A // 天气单位设置
	CMD_TIME_FORMAT  byte = 0x5B // 12h/24h时间制切换
	CMD_DEVICE_INFO  byte = 0x5C // 设备端信息
	CMD_APP_INFO     byte = 0x5D // 应用端信息

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
	CMD_TYPE_REQUEST    byte = 0x01 // Request: 手机 -> 设备
	CMD_TYPE_RESPONSE   byte = 0x02 // Response: 设备 -> 手机
	CMD_TYPE_NOTIFY     byte = 0x03 // Notify: 主动上报，不需要对端回复
	CMD_TYPE_REQUEST_N  byte = 0x04 // Request_n: 设备 -> 手机
	CMD_TYPE_RESPONSE_N byte = 0x05 // Response_n: 手机 -> 设备
	CMD_TYPE_NOTIFY_N   byte = 0x06 // Notify_n: 主动上报，不需要对端回复
)

// Action命令常量 - Action Command Constants
const (
	ACTION_QUERY byte = 0x00 // 查询
	ACTION_SET   byte = 0x01 // 设置
)

// 返回值常量 - Response Code Constants (表7-1)
const (
	CMD_RESULT_OK          byte = 0x00 // 成功
	CMD_RESULT_FAIL        byte = 0x01 // 失败
	CMD_RESULT_LEN_INVALID byte = 0x02 // 包长度不合法
	CMD_RESULT_CMD_INVALID byte = 0x03 // 命令类型不合法
	CMD_RESULT_IDX_INVALID byte = 0x04 // 编号索引不合法
	CMD_RESULT_NOT_SUPPORT byte = 0x05 // 设备不支持该命令
	CMD_RESULT_SWI_ERROR   byte = 0x06 // 开关操作不合法
	CMD_RESULT_CHECK_ERROR byte = 0x07 // 数据校验错误
	CMD_RESULT_PACKET_LOSS byte = 0x08 // 上一包数据丢失
)

// 语言常量 - Language Constants (表7-2-1)
const (
	LANG_ENGLISH             byte = 0x00
	LANG_CHINESE_SIMPLIFIED  byte = 0x01
	LANG_ITALIAN             byte = 0x02
	LANG_SPANISH             byte = 0x03
	LANG_PORTUGUESE          byte = 0x04
	LANG_RUSSIAN             byte = 0x05
	LANG_JAPANESE            byte = 0x06
	LANG_CHINESE_TRADITIONAL byte = 0x07
	LANG_GERMAN              byte = 0x08
	LANG_KOREAN              byte = 0x09
	LANG_THAI                byte = 0x0A
	LANG_ARABIC              byte = 0x0B
	LANG_TURKISH             byte = 0x0C
	LANG_FRENCH              byte = 0x0D
	LANG_VIETNAMESE          byte = 0x0E
	LANG_POLISH              byte = 0x0F
	LANG_DUTCH               byte = 0x10
	LANG_HEBREW              byte = 0x11
	LANG_PERSIAN             byte = 0x12
	LANG_GREEK               byte = 0x13
	LANG_MALAYSIAN           byte = 0x14
	LANG_BURMESE             byte = 0x15
	LANG_DANISH              byte = 0x16
	LANG_UKRAINIAN           byte = 0x17
	LANG_INDONESIAN          byte = 0x18
	LANG_CZECH               byte = 0x19
	LANG_HINDI               byte = 0x20
)

// 消息类型常量 - Message Type Constants (表7-3-1)
const (
	MSG_NULL               byte = 0x00
	MSG_INCOMING_CALL      byte = 0x01 // 来电
	MSG_MISSED_CALL        byte = 0x02 // 未接来电
	MSG_SMS                byte = 0x03 // 短信
	MSG_EMAIL              byte = 0x04 // 邮件
	MSG_SCHEDULE           byte = 0x05 // 日程
	MSG_FACETIME           byte = 0x06
	MSG_QQ                 byte = 0x07
	MSG_SKYPE              byte = 0x08
	MSG_WECHAT             byte = 0x09 // 微信
	MSG_WHATSAPP           byte = 0x0A
	MSG_GMAIL              byte = 0x0B
	MSG_HANGOUT            byte = 0x0C
	MSG_INBOX              byte = 0x0D
	MSG_LINE               byte = 0x0E
	MSG_TWITTER            byte = 0x0F // 推特
	MSG_FACEBOOK           byte = 0x10 // 脸书
	MSG_FACEBOOK_MESSENGER byte = 0x11
	MSG_INSTAGRAM          byte = 0x12
	MSG_WEIBO              byte = 0x13 // 微博
	MSG_KAKAOTALK          byte = 0x14
	MSG_FACEBOOK_MANAGER   byte = 0x15
	MSG_VIBER              byte = 0x16
	MSG_VKCLIENT           byte = 0x17
	MSG_TELEGRAM           byte = 0x18
	MSG_SNAPCHAT           byte = 0x1A
	MSG_DINGTALK           byte = 0x1B // 钉钉
	MSG_ALIPAY             byte = 0x1C // 支付宝
	MSG_TIKTOK             byte = 0x1D // 抖音
	MSG_LINKEDIN           byte = 0x1E // 领英
)

// 开关类型常量 - Switch Type Constants (表7-4-1)
const (
	SWITCH_ANTI_LOST     byte = 0x00 // 防丢开关
	SWITCH_RAISE_WRIST   byte = 0x01 // 抬手亮屏开关
	SWITCH_AUTO_SYNC     byte = 0x02 // 自动同步开关
	SWITCH_SLEEP_MONITOR byte = 0x04 // 睡眠监测开关
	SWITCH_MSG_TOTAL     byte = 0x05 // 消息提醒总开关
	SWITCH_HOURLY_SPORT  byte = 0x06 // 整点上传运动数据开关
	SWITCH_GOAL_ACHIEVED byte = 0x07 // 目标达成开关
	SWITCH_MSG_SCREEN    byte = 0x09 // 消息提醒亮屏开关
	SWITCH_SOUND         byte = 0x0A // 声音开关
	SWITCH_VIBRATE_TOTAL byte = 0x0B // 震动总开关
	SWITCH_HOURLY_HEALTH byte = 0x0C // 整点上传健康数据开关
	SWITCH_MSG_VIBRATE   byte = 0x0D // 消息提醒震动开关
)

// 天气类型常量 - Weather Type Constants
const (
	WEATHER_CLOUDY    byte = 0x00 // 多云
	WEATHER_SUNSHINE  byte = 0x01 // 晴天
	WEATHER_SNOW      byte = 0x02 // 雪天
	WEATHER_RAIN      byte = 0x03 // 雨天
	WEATHER_OVERCAST  byte = 0x04 // 阴天
	WEATHER_SAND_DUST byte = 0x05 // 沙尘天气
	WEATHER_WINDY     byte = 0x06 // 大风天气
	WEATHER_HAZE      byte = 0x07 // 阴霾天气
)

// 天气命令类型 - Weather Command Type
const (
	WEATHER_CMD_TYPE         byte = 0x00 // 天气类型
	WEATHER_CMD_CURRENT_TEMP byte = 0x01 // 当前温度
	WEATHER_CMD_MIN_TEMP     byte = 0x02 // 最低温度
	WEATHER_CMD_MAX_TEMP     byte = 0x03 // 最高温度
	WEATHER_CMD_LOCATION     byte = 0x04 // 地点位置
	WEATHER_CMD_HUMIDITY     byte = 0x05 // 湿度
	WEATHER_CMD_WIND_SPEED   byte = 0x06 // 风速
	WEATHER_CMD_UV_INDEX     byte = 0x07 // 紫外线指数
	WEATHER_CMD_PRESSURE     byte = 0x08 // 气压
	WEATHER_CMD_SUNRISE      byte = 0x09 // 日出时间
	WEATHER_CMD_SUNSET       byte = 0x0A // 日落时间
	WEATHER_CMD_RAIN_PROB    byte = 0x0B // 降雨概率
	WEATHER_CMD_RAINFALL     byte = 0x0C // 降雨量
	WEATHER_CMD_VISIBILITY   byte = 0x0D // 能见度
	WEATHER_CMD_AIR_QUALITY  byte = 0x0E // 空气质量指数
)

// 设备类型 - Device Type
const (
	DEVICE_TYPE_SQUARE   byte = 0x00 // 方形手表
	DEVICE_TYPE_ROUND    byte = 0x01 // 圆形手表
	DEVICE_TYPE_BRACELET byte = 0x02 // 手环
)

// 手机类型 - Phone Type
const (
	PHONE_TYPE_ANDROID byte = 0x00 // 安卓手机
	PHONE_TYPE_APPLE   byte = 0x01 // 苹果手机
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
	TEMP_UNIT_CELSIUS    byte = 0x00 // 摄氏度
	TEMP_UNIT_FAHRENHEIT byte = 0x01 // 华氏度
)

// 时间制 - Time Format
const (
	TIME_FORMAT_12H byte = 0x00 // 12小时制
	TIME_FORMAT_24H byte = 0x01 // 24小时制
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
