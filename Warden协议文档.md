协议文档

VERSION：1.0.0

1 文档目标

本协议定义了APP端和蓝牙设备端之间的协议传输格式和规范， APP端和设备端根据该协议文档进行开

发对接。

2 全局说明

字节序: 文中的两字节和四字节字段，如无特殊说明，在传输中均采用小端格式进行传输, 多字节变长字

段(如字符串), 在传输中均采用大端格式进行传输。

开关：文中所有1bit大小的文段描述, 如无特殊说明，对应bit值为1表示打开，0表示关闭；

帧序：文中所有Frame Seq的文段描述, 只为0(default)时表示无需拼包;

3 指令格式

| Header | Payload |
| --- | --- |
| 5Byte | 0~N Byte |

Header格式如下：

| 字节序 | 说明 |
| --- | --- |
| 0 | Bit0 ~ Bit3: Seq Num，由0到15顺序递增，用于APP和设备各自做命令计数，检查是否有丢包 Bit4 ~ Bit6:暂时保留，填0 <br>Bit7:数据加密指示。0：不加密，1：加密 |
| 1 | Cmd，表示当前是哪条命令，如果该命令的类型为Request的，回复Response时的Cmd需和Request的Cmd保持一致 |
| 2 | Cmd Type，命令的类型，取值为： <br>//手机作主机 <br>1: Request, 手机(host) -> 设备(client) <br>2: Response, 设备(client) -> 手机(host) <br>3: Notify（主动上报数据，不需要对端回复） <br>//设备作主机 <br>4: Request_n, 设备(client) -> 手机(host) <br>5: Response_n, 手机(host) -> 设备(client) <br>6: Notify_n（主动上报数据，不需要对端回复） |
| 3 | Bit0 ~ Bit3: Frame Seq，帧序号，取值0~15，从0开始计数。 <br>Bit4 ~ Bit7:Total Frame，总的帧数，取值0~15，实际的总帧数等于（Bit4~Bit7）的值加1。 |

|  | 当数据长度太长需要拆包发送时，可根据该字节确定总的有多少帧数据以及当前是第几帧数据，若无需拆包，该字节为0 |
| --- | --- |
| 4 | Frame Length，表示当前帧的数据长度，即Payload<br>的长度 |

4 TLV格式说明

TLV是Type，Length和Value的缩写，一个基本的数据元包括这三个域。Type唯一标识该数据元，Length是Value域的长度，Value是数据本身。说明如下表：

| 字节序 | 命名 | 说明 |
| --- | --- | --- |
| 0 | Type | 占一个字节，表示类型 |
| 1 | Length | 占一个字节，表示数据的长度 |
| 2 ~ N | Value | 占N个字节，数据内容 |

当命令包含有多组功能集时，Payload通常使用TLV格式（注：并不是所有命令的Payload都需要采用TLV数据格式），TLV指令组可以组合进行使用（Type1+Length1+Value1+Type2+Length2+Value2+……）。如果手机端发起的命令数据格式为TLV格式，设备端也必须使用TLV格式回复，并且Type保持一致。

5 指令集

1 5.1 设备属性类

1.1  0x50：同步时间

大体流程:

(1) 绑定设备时, 手机端发起一次时间同步

设置:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x50 | 同步时间 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x06 | Payload的长度 |
| 5 | Action cmd | 1 | 0x01 | 0x01：设置时间 |
| 6 | Time zone | 1 | 0x00~0x | 0~11: 西十二区~西一区 |
| 6 | Time zone | 1 | 0x00~0x | 12:中时区/零时区 |

13~24:东一区~东十二区

| 7~10 | UTC | 4 | N | 时间戳 |
| --- | --- | --- | --- | --- |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x50 | 同步时间 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |

| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| --- | --- | --- | --- | --- |
| 5 | Action cmd | 1 | 0-0xff | 0x01：设置 |
| 6 | Response | 1 | 0x00/0x01 | 返回响应, 见表7-1返回值表 |

1.2  0x51：电量获取

大体流程:

(1) 绑定设备时, 设备端上报一次电量

(2) 设备端检测到电量变化, 上报电量

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x51 | 电量获取 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | Action cmd | 1 | 0-0xff | 0x00：查询电量值 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x51 | 电量获取 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0-0xff | 0x00：查询 |

电量值0~100

| 6 | Response | 1 | 0x00-0xff | bit7：充电状态：1：充电中；0：没有充电 |
| --- | --- | --- | --- | --- |
| 6 | Response | 1 | 0x00-0xff | bit6~bit0：电量值 |

设备端电量更新主动上报

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |

| 集0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| --- | --- | --- | --- | --- |
| 集0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x51 | 电量获取 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |

电量值0~100

| 5 | Response | 1 | 0x00-0xff | bit7：充电状态：1：充电中；0：没有充电 |
| --- | --- | --- | --- | --- |
| 5 | Response | 1 | 0x00-0xff | bit6~bit0：电量值 |

1.3  0x52：屏幕亮度

查询屏幕亮度

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x52 | 屏幕亮度 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询屏幕亮度值 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x52 | 屏幕亮度 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询屏幕亮度值 |
| 6 | brightness value | 1 | 0x00-0xff | 返回屏幕亮度值(0~100级) |

设置屏幕亮度

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x52 | 屏幕亮度 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置屏幕亮度值 |
| 6 | brightness value | 1 | 0x00-0xff | 屏幕亮度值(0~100级) |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x52 | 屏幕亮度 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置屏幕亮度值 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

设备端亮度更新主动上报

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x52 | 屏幕亮度 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | brightness value | 1 | 0x00-0xff | 返回屏幕亮度值 |

1.4  0x53：设备语言

查询设备语言

| Field | Size | Value | Description |
| --- | --- | --- | --- |
| Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| Seq Num & Enc | 1 | 0x00 | 加密 |
| Cmd | 1 | 0x53 | 设备语言 |
| Cmd Type | 1 | 0x01 | Request |
| Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| Frame Length | 1 | N | Payload的长度 |
| Action cmd | 1 | 0x00-0xff | 0x00：查询设备语言 |
| Field | Size | Value | Description |
| Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| Seq Num & Enc | 1 | 0x00 | 加密 |
| Cmd | 1 | 0x53 | 设备语言 |
| Cmd Type | 1 | 0x02 | Response |
| Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| Frame Length | 1 | N | Payload的长度 |
| Action cmd | 1 | 0x00-0xff | 0x00：查询设备语言 |
| language type | 1 | 0x00-0xff | 语言类型；（类型见表7-2-1 语言号 |
| language type | 1 | 0x00-0xff | 表） |

设置设备语言

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x53 | 设备语言 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置设备语言 |
| 6 | language type | 1 | 0x00-0xff | 设置语言类型 |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |

| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x53 | 设备语言 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置设备语言 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

设备端语言类型更新主动上报

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x53 | 设备语言 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | language type | 1 | 0x00-0xff | 语言类型；（类型见表7-2-1 语 |
| 5 | language type | 1 | 0x00-0xff | 言号表） |

1.8  0x57：勿扰功能

查询:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x57 | 勿扰功能 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x57 | 勿扰功能 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |
| 6 | switch | 1 | 0x00-0xff | [0]: 定时勿扰开关 |
| 7 | start hour | 1 | 0x00-0x18 | 开始时间：小时 |
| 8 | start minute | 1 | 0x00-0x3c | 开始时间：分钟 |
| 9 | end hour | 1 | 0x00-0x18 | 结束时间：小时 |
| 10 | end minute | 1 | 0x00-0x3c | 结束时间：分钟 |

设置:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x57 | 勿扰功能 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置 |
| 6 | switch | 1 | 0x00-0xff | [0]: 定时勿扰开关 |
| 7 | start hour | 1 | 0x00-0x18 | 开始时间：小时 |
| 8 | start minute | 1 | 0x00-0x3c | 开始时间：分钟 |
| 9 | end hour | 1 | 0x00-0x18 | 结束时间：小时 |
| 10 | end minute | 1 | 0x00-0x3c | 结束时间：分钟 |
| 7 | start hour | 1 | 0x00-0x18 | 开始时间：小时 |
| 8 | start minute | 1 | 0x00-0x3c | 开始时间：分钟 |
| 9 | end hour | 1 | 0x00-0x18 | 结束时间：小时 |
| 10 | end minute | 1 | 0x00-0x3c | 结束时间：分钟 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x57 | 勿扰功能 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置勿扰设置 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

1.10   0x59：寻找手机

| 0ffset | field | size | value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x59 | 寻找手机 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：寻找手机 <br>0x01：停止寻找 |

设置

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x59 | 寻找手机 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：停止寻找 |

设置响应

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x59 | 寻找手机 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：停止寻找 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

1.11   0x5A：天气单位设置

查询

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5a | 天气单位设置 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5a | 天气单位设置 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |
| 6 | Response | 1 | 0x00-0xff | 返回温度单位：0x00：摄氏度 0x01： |
| 6 | Response | 1 | 0x00-0xff | 华氏度 |

设置

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5a | 天气单位设置 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置 |
| 6 | unit | 1 | 0x00-0x01 | 温度单位：0x00：摄氏度 0x01：华 |
| 6 | unit | 1 | 0x00-0x01 | 氏度 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5a | 天气单位设置 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |

1.12 	0x5B：12h/24h时间制切换

查询

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5b | 12h/24h时间制切换 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5b | 12h/24h时间制切换 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |
| 6 | Response | 1 | 0x00-0xff | 返回时间制：0x00:12h 0x01:24h |

设置

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5b | 12h/24h时间制切换 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |

| 4 | Frame Length | 1 | N | Payload的长度 |
| --- | --- | --- | --- | --- |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置 |
| 6 | unit | 1 | 0x00-0x01 | 0x00:12h 0x01:24h |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x5b | 12h/24h时间制切换 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

1.13 	0x5C: 设备端信息

描述: 获取设备信息,  app端绑定时查询;

查询:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; |
| 0 | Seq Num & Enc | 1 | 0x00 | [7]: 0：不加密，1：加密 |
| 1 | Cmd | 1 | 0x5C | 设备信息 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00: 查询 |
| 6 | Data type | 1 | 0x00-0xff | 0x00: 手表类型 <br>0x01: 支持的设备语言; 类型见下表<br>0x02: 序列号<br>0x03: 固件版本信息 |

查询响应:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; |
| 0 | Seq Num & Enc | 1 | 0x00 | [7]: 0：不加密，1：加密 |
| 1 | Cmd | 1 | 0x5C | 设备信息 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x00-0xff | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00: 查询 |
| 6 | Data type | 1 | 0x00-0xff | 0x00: 手表类型 <br>0x01: 支持的设备语言; 类型见下表<br>0x02: 序列号<br>0x03: 固件版本信息 |

当Data type域为0x00时, 接续

| 7 | Watch type | 1 | 0x00-0xff | [0,3]: 手表类型, 0: 方形手表, 1: 圆形 |
| --- | --- | --- | --- | --- |
| 7 | Watch type | 1 | 0x00-0xff | 手表, 2: 手环; |

当Data type域为0x01时, 接续

| 7~8 | Support | 4 | 0x00-0xffff | 支持的设备语言; （类型见下表） |
| --- | --- | --- | --- | --- |
| 7~8 | language | 4 | 0x00-0xffff | 支持的设备语言; （类型见下表） |

当Data type域为0x02时, 接续

| 7~12 | Serial number | 32 | N | 序列号(字节序: 大端) |
| --- | --- | --- | --- | --- |

当Data type域为0x03时, 接续

| 7 | Firmware | 1 | 0x00-0xff | 固件主版本号 |
| --- | --- | --- | --- | --- |
| 7 | Major Version | 1 | 0x00-0xff | 固件主版本号 |
| 8 | Firmware | 1 | 0x00-0xff | 固件从版本号 |
| 8 | Minor Version | 1 | 0x00-0xff | 固件从版本号 |

固件版本号格式： Major.Minor

0x03: 设备版本信息

上报:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; <br>[7]: 0：不加密，1：加密 |
| 1 | Cmd | 1 | 0x5C | 设备信息 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x15 | Payload的长度 |
| 5 | Watch type | 1 | 0x00-0xff | [0,3]: 手表类型, 0: 方形手表, 1: 圆 |
| 5 | Watch type | 1 | 0x00-0xff | 形手表, 2: 手环; |
| 6~9 | Support language | 4 | N | 设备支持的语言; （类型见7-2-1 |
| 6~9 | Support language | 4 | N | 表） |

[0]: 是否支持表盘市场

[1]: 是否支持消息提醒

[2]: 是否支持天气功能

[3]: 是否支持NFC

[4]: 是否支持通讯录

[5]: 是否支持钱包功能

| 10~13 | Function control | 4 | N | [6]: 是否支持名片显示 |
| --- | --- | --- | --- | --- |
| 10~13 | Function control | 4 | N | [7]: 是否支持固件升级 |
| 10~13 | flags | 4 | N | [7]: 是否支持固件升级 |
| 10~13 | flags | 4 | N | [8]: 是否支持勿扰功能 |

[9]: 是否支持亮屏设置

[10]: 是否支持拍一拍功能

[11]: 是否支持12/24小时制切换

[12]: 是否支持公制/英制单位切换

[13]: 是否支持静态卡路里

[14]:是否支持一键双连

[0]: 是否支持心率检测

[1]: 是否支持血氧检测

| 14~17 | Health control | 4 | N | [2]: 是否支持血压检测 |
| --- | --- | --- | --- | --- |
| 14~17 | flags | 4 | N | [3]: 是否支持血糖检测 |

[4]: 是否支持睡眠检测

[5]: 是否支持计步/卡路里/距离检测

| 18~21 | Msg control | 4 | N | 支持的消息推送类型（见表7-3-2 |
| --- | --- | --- | --- | --- |
| 18~21 | Msg control | 4 | N | 消息位表） |
| 22~25 | Switch control | 4 | N | 支持的开关控制类型（见表7-4-2 |
| 22~25 | Switch control | 4 | N | 开关位表） |
| 26~29 | Sport control | 4 | N | 支持的运动类型（见7-5-2 运动位 |
| 26~29 | Sport control | 4 | N | 表） |
|  |  |  | N | 表） |

1.14   0x5D: 应用端信息

设置:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |

| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [7]: 0：不加密，1：加密 |
| 1 | Cmd | 1 | 0x5D | 应用端信息 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x03 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01: 设置 |
| 6 | Data type | 1 | 0x00-0xff | 0x00: 手机类型 |
| 7 | Phone type | 1 | 0x00-0xff | 0x00: 安卓手机 |
| 7 | Phone type | 1 | 0x00-0xff | 0x01: 苹果手机 |

设置响应:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1：加密 |
| 0 | Seq Num & Enc | 1 | 0x00 |  |
| 1 | Cmd | 1 | 0x5D | 应用端信息 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

3 5.3 状态控制类

3.1 0x80：开关设置

查询开关

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x80 | 开关设置 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x00：查询开关 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x80 | 开关设置 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x05 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询开关 |
| 6~9 | Response | 4 | 0x00- | 开关设置（见表7-4-2 开关位表） |
| 6~9 | Response | 4 | 0xffffffff | 开关设置（见表7-4-2 开关位表） |

设置开关

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x80 | 开关设置 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x05 | Payload的长度 |

| 5 | Action cmd | 1 | 0x00-0x01 | 0x01：设置开关 |
| --- | --- | --- | --- | --- |
| 6~9 | Response | 4 | 0x00-<br>0xffffffff | 返回开关设置（见表7-4-2 开关位表） |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x80 | 开关设置 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置开关 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

3.2 0x81：设备绑定

大体流程: 
(1) 手机端发起绑定, 首先发一条绑定开始, 然后同步数据给设备, 数据发完后发一条绑定数据结束; (2) 设备端直到初始化完成后发送绑定完成;

可能使用场景: 
(1) 首次绑定设备, 设备端需要清空设备运动数据, 重复绑定则无需清空; 
(2) 首次绑定设备, 设备端进入用户信息设置界面, 操作完成后, 设备端发送绑定完成; (3) IOS app断开后是没有断开ble的, 这个时候就要通过是否已绑定来判断是否推送消息;

设置:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x81 | 绑定开始 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x01: 设置 |
| 6 | Contrl | 1 | 0x00-0xff | 0x00: 绑定开始 <br>0x01: 绑定数据结束 <br>0x02: 断开绑定 <br>0x03: 二维码绑定开始(预留) 0x04: 二维码绑定结束(预留) |

设置响应:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x81 | 绑定开始 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x03 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x01: 设置 |
| 6 | Contrl | 1 | 0x00-0xff | 0x00: 绑定开始 <br>0x01: 绑定数据结束 <br>0x02: 断开绑定 <br>0x03: 二维码绑定开始(预留) 0x04: 二维码绑定结束(预留) |

当control域为0x00时, 接续

| 7 | Bind_first | 1 | 0x00-0xff | 0x00:未被绑定过 |
| --- | --- | --- | --- | --- |
| 7 | Bind_first | 1 | 0x00-0xff | 0x01:已被绑定过 |

当control域为0x01时, 接续

| 7 | Response | 1 | 0x00-0xff | [0]: 绑定是否完成 |
| --- | --- | --- | --- | --- |

当control域为0x02时, 接续

| 7 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |
| --- | --- | --- | --- | --- |

上报

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x82 | 绑定结束 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | State | 1 | 0x00 | 0x00: 绑定完成 |

3.3 0x82：绑定(预留)

查询:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x82 | 绑定 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询设备端初始化完成的标志 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x82 | 绑定 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询设备端初始化完成的标志 |
| 6 | Response | 1 | 0x00-0xff | 返回设备端初始化标志 |

绑定完成后发送下面协议，设备端则从初始化界面进入主界面，清零运动显示

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x82 | 绑定 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：绑定完成 |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x82 | 绑定 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：绑定完成 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x82 | 绑定 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：绑定完成 |

3.7 0x86：开关表扩展

查询开关表

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x86 | 开关表设置 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x00：查询开关 |
| 6 | Switch type | 1 | 0x00 | 0x00：社交开关类型 |
| 6 | Switch type | 1 | 0x00 | 可扩展 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x86 | 加密 |
| 1 | Cmd | 1 | 0x86 | 开关表设置 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x06 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x00：查询开关 |
| 6 | Switch type | 1 | 0x00 | 0x00：社交开关类型 |
| 6 | Switch type | 1 | 0x00 | 可扩展 |
| 7~10 | Response | 4 | 0x00- | 返回开关设置；（见表7-3-2 消息位 |
| 7~10 | Response | 4 | 0xffffffff | 表） |

设置:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x86 | 开关表设置 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x06 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x01：设置开关 |
| 6 | switch type | 1 | 0x00 | 0x00：社交开关类型 |
| 6 | switch type | 1 | 0x00 | 可扩展 |
| 7~10 | Response | 4 | 0x00- | 返回开关设置（见表7-3-2 消息位 |
| 7~10 | Response | 4 | 0xffffffff | 表） |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x86 | 开关表设置 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置开关 |
| 6 | Response | 1 | 0x00-0x01 | 返回响应, 见表7-1返回值表 |

3.9 0x91：远程拍照

大体流程:

(1) 手机端进入拍一拍后, 需设置打开拍照开关, app退出拍一拍时, 需关闭拍照开关;

(2) 设备端触发拍照

1) 若手机端未打开系统相机界面, 则呼出手机端系统相机

2) 若手机端已打开系统相机界面, 则触发拍照

(3)

查询:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x91 | 远程拍照 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1：加密 |

| 1 | Cmd | 1 | 0x91 | 远程拍照 |
| --- | --- | --- | --- | --- |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：查询 |
| 6 | Response | 1 | 0x00-0xff | 返回远程拍照开关状态：0x00：关闭 |
| 6 | Response | 1 | 0x00-0xff | 0x01：打开 |

设置

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x91 | 远程拍照 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：设置 |
| 6 | Camera ctrl | 1 | 0x00-0x01 | 设置拍照开关状态：0x00：关闭 ;  0x01:打开 |

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0x91 | 远程拍照 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：查询 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

上报

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密 |
| 0 | Seq Num & Enc | 1 | 0x00 | 1:加密 |
| 1 | Cmd | 1 | 0x91 | 远程拍照 |
| 2 | Cmd Type | 1 | 0x03 | Notify |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x01 | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：远程拍照 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：退出APP相机     0x02：呼出APP相机 |

4 5.4 消息推送类

4.1 0xA0：消息推送

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0xa0 | 消息推送 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：设置 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：推送内容 |

当Action cmd域为0x00时,

| 6 | Control | 1 | 0x00-0xff | [0]: 消息提醒总开关 
[1]: 消息提醒亮屏开关 
[2]: 消息提醒震动开关 
(与命令0x80中设置的效果相同) |
| --- | --- | --- | --- | --- |

当Action cmd为0x01时,

| 6 | Type | 1 | 0x00-0xff | 消息类型（见表7-3-1 消息号表） |
| --- | --- | --- | --- | --- |
| 7~N(小于<br>MTU值) | Value | MTU - 7 | N | 内容 |

当需要拼包时, 第二帧开始接续:

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0xa0 | 加密 |
| 1 | Cmd | 1 | 0xa0 | 消息推送 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5~N(小于 | Value | MTU - 7 | N | 内容 |
| MTU值) | Value | MTU - 7 | N | 内容 |

响应

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0xa0 | 消息推送 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | N | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x00：设置 |
| 5 | Action cmd | 1 | 0x00-0xff | 0x01：推送 |
| 6 | Response | 1 | 0x00-0xff | 返回响应, 见表7-1返回值表 |

当响应包的Response域为CMD_RESULT_PACKET_LOSS时, 意为上一包数据丢失, 本次数据传输需要重传;

4.2 0xA1：天气信息

设置天气

| Offset | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0xa1 | 天气信息 |
| 2 | Cmd Type | 1 | 0x01 | Request |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x05-0xff | Payload的长度 |
| 5 | Action cmd | 1 | 0x00-0x01 | 0x01：设置 |
| 6 | Date type | 1 | 0x00-0x02 | 0x00:当天天气 |
| 6 | Date type | 1 | 0x00-0x02 | 0x01:明天天气 |

0x02:后天天气

Cmd type 	指令类型：（以下括号内数字为数据

建议字节长度, 温度为int型，区分温

度为负数情况）

0x00：天气类型（id见下表）

0x01：当前温度（℃，1）

0x02：最低温度（℃，1）

0x03：最高温度（℃，1）

0x04：地点位置(utf8编码,  N)

//以下参数暂未支持

0x05：湿度（%，1）

0x06：风速（km/h，1）

0x07：紫外线指数（1）

0x08：气压（pa，4）

| 7 | 1 | 0x00-0x0f | 0x09：日出时间（时分，2） |
| --- | --- | --- | --- |
| 7 | 1 | 0x00-0x0f | 0x0a：日落时间（时分，2） |

0x0b：降雨概率（%，1）

0x0c：降雨量（0.1mm，2）

0x0d：能见度（m，4）

0x0e：空气质量指数（等级由指数决

定）（2）

美国标准：

0-50 Good

51-100Moderate

101-150 Unhealthy for sensitive

groups

151-200 Unhealthy

200-300 Very unhealthy

301-500 Hazardous

后续可扩展

| 8 | Length | 1 | 0x00-0xff | 数值长度 |
| --- | --- | --- | --- | --- |
| 9~N | value | N | N | 数值 |
| N | TLV格式 | N | N | 指令类型+数值长度 +数值  <br>(多种指令拼接发送) |

示例: 00a101000e0100000102010119020105030119, 设置当天天气, 雪天, 当前气温25℃, 最低气温5℃最	高25℃

天气类型id表

| 0 | Field | Size | Cloudy | 多云 |
| --- | --- | --- | --- | --- |
| 1 | Field | Size | Sunshine | 晴天 |
| 2 | Field | Size | Snow | 雪天 |
| 3 | Field | Size | Rain | 雨天 |
| 4 | Field | Size | Overcast | 阴天 |
| 5 | Field | Size | Sand&Dust | 沙尘天气 |
| 6 | Field | Size | Windy | 大风天气 |
| 7 | Field | Size | Haze | 阴霾天气 |
| Offset | Field | Size | Value | Description |
| 0 | Seq Num & Enc | 1 | 0x00 | [0,3]: Seq Num ; [7]: 0：不加密，1： |
| 0 | Seq Num & Enc | 1 | 0x00 | 加密 |
| 1 | Cmd | 1 | 0xa1 | 天气信息 |
| 2 | Cmd Type | 1 | 0x02 | Response |
| 3 | Frame Seq | 1 | 0x00 | [0,3]: 帧序号; [4,7]:总的帧数; |
| 4 | Frame Length | 1 | 0x02 | Payload的长度 |
| 5 | Action cmd | 1 | 0x01 | 0x01：设置 |
| 6 | Response | 1 | 0x00-0x01 | 返回响应, 见表7-1返回值表 |

6 规范

1 6.1 ble广播规范

广播包必须包含规定的厂商自定义格式，如下：

| 字节序 | 名称 | 取值 | 说明 |
| --- | --- | --- | --- |
| 0 | Length | 0x10 | 广播包长度 |
| 1 | Type | 0xFF | 厂商自定义格式 |
| 2~3 | CID | 0x0642 | 公司编码 |
| 4 | VID | 0x01 | Bit0 ~ Bit3： 协议版本号，当前版本号为 1 |
| 4 | VID | 0x01 | Bit4 ~ Bit7：保留 |

产品 ID

| 5~6 | PID | 0x01 ~0x02 | 0x01: 耳机 |
| --- | --- | --- | --- |

0x02: 手表

| 7~12 | MAC | 经典蓝牙 MAC 地址的每个字节异或上 0xAD 组 |
| --- | --- | --- |
| 7~12 | MAC | 成; bt地址则为该地址异或0x55组成; |

Bit0：是否需要进行安全认证

0：不需要安全认证

1：需要安全认证

Bit1：保留

| 13 | FMASK | 0x00 | Bit2 ~ Bit3：设备状态 |
| --- | --- | --- | --- |

00：经典蓝牙未连接

01：经典蓝牙已连接

1x：保留

Bit4 ~ Bit7：保留

品牌 ID，可用于区分不同代理商和客户

| 14~16 | BID | 0 ~ 0xffffff | 0x00 : 厂家 |
| --- | --- | --- | --- |

其他: 代理商和客户

例如广播出来的自定义数据为：0x4206020200ECEF0E6566E500000000

0x0642:为CID, 0x01是VID，0x0002是PID，0xECEF0E6566E5为MAC地址, 0x00为FMASK,

0x000000为BID;

蓝牙工具显示如下图:

连接UUID

| UUID | Property | Description |
| --- | --- | --- |
| 6e400001-b5a3-f393-e0a9- | Write No Reasponse | Service UUID |
| e50e24dcca9d | Write No Reasponse | Service UUID |
| 6e400002-b5a3-f393-e0a9- | Write No Reasponse | Characteristics UUID 用于APP写数据到设备 |
| e50e24dcca9d | Write No Reasponse | Characteristics UUID 用于APP写数据到设备 |
| 6e400003-b5a3-f393-e0a9- | Read Notify | Characteristics UUID 用于设备Notify数据给APP |
| e50e24dcca9d | Read Notify | Characteristics UUID 用于设备Notify数据给APP |

2 6.2 序列号分配规则

| Byte | Field | Size | Value | Description |
| --- | --- | --- | --- | --- |
| 0 | CID | 2 | 0x00~0xffff | 公司识别码(最大6万5535) |
| 0 | CID | 2 | 0x00~0xffff | 0x00: 厂家 |
| 0 | CID | 2 | 0x00~0xffff | 0x01: 华胜杰 |

0xff: 测试设备

| 1~2 | PID | 1 | 0x00-0xff | 产品代码(最大255) |
| --- | --- | --- | --- | --- |
| 3~5 | ID | 3 | 0x00- | 设备唯一识别号 |
| 3~5 | ID | 3 | 0xffffff | (最大1677万7215) |

7 附录

1 7.1 返回值表

Cmd Type 为 Response的包中,  response域统一使用如下含义:

| Value | Desc | 描述 |
| --- | --- | --- |
| 0 | CMD_RESULT_OK | 成功 |
| 1 | CMD_RESULT_FAILE | 失败 |
| 2 | CMD_RESULT_LEN_INVALID | 包长度不合法 |
| 3 | CMD_RESULT_CMD_INVALID | 命令类型不合法 |
| 4 | CMD_RESULT_IDX_INVALID | 编号索引不合法 |
| 5 | CMD_RESULT_NOT_SUPPORT | 设备不支持该命令 |
| 6 | CMD_RESULT_SWI_ERROR | 开关操作不合法 |
| 7 | CMD_RESULT_CHECK_ERROR | 数据校验错误 |
| 8 | CMD_RESULT_PACKET_LOSS | 上一包数据丢失 |

表7-1 返回值表

2 7.2 语言表

| Language ID | Type | Language ID | Type |
| --- | --- | --- | --- |
| 0x00 | 英文 | 0x01 | 简体中文 |
| 0x02 | 意大利语 | 0x03 | 西班牙语 |
| 0x04 | 葡萄牙语 | 0x05 | 俄语 |
| 0x06 | 日语 | 0x07 | 繁体中文 |
| 0x08 | 德语 | 0x09 | 韩语 |
| 0x0a | 泰语 | 0x0b | 阿拉伯语 |
| 0x0c | 土耳其语 | 0x0d | 法语 |
| 0x0e | 越南语 | 0x0f | 波兰语 |
| 0x10 | 荷兰语 | 0x11 | 希伯来语 |
| 0x12 | 波斯语 | 0x13 | 希腊语 |
| 0x14 | 马来西亚语 | 0x15 | 缅甸语 |
| 0x16 | 丹麦语 | 0x17 | 乌克兰语 |

| 0x18 | 印度尼西亚语 | 0x19 | 捷克语 |
| --- | --- | --- | --- |
| 0x20 | 印地语 | 0x21 |  |

表7-2-1 语言号表

| Language ID | Type | Language ID | Type |
| --- | --- | --- | --- |
| BIT(0) | 英文 | BIT(1) | 简体中文 |
| BIT(2) | 意大利语 | BIT(3) | 西班牙语 |
| BIT(4) | 葡萄牙语 | BIT(5) | 俄语 |
| BIT(6) | 日语 | BIT(7) | 繁体中文 |
| BIT(8) | 德语 | BIT(9) | 韩语 |
| BIT(10) | 泰语 | BIT(11) | 阿拉伯语 |
| BIT(12) | 土耳其语 | BIT(13) | 法语 |
| BIT(14) | 越南语 | BIT(15) | 波兰语 |
| BIT(16) | 荷兰语 | BIT(17) | 希伯来语 |
| BIT(18) | 波斯语 | BIT(19) | 希腊语 |
| BIT(20) | 马来西亚语 | BIT(21) | 缅甸语 |
| BIT(22) | 丹麦语 | BIT(23) | 乌克兰语 |
| BIT(24) | 印度尼西亚语 | BIT(25) | 捷克语 |
| BIT(26) | 印地语 | BIT(25) | 捷克语 |

表7-2-2 语言位表

3 7.3 消息推送表

| Message ID | Type | Message ID | Type |
| --- | --- | --- | --- |
| 0x00 | NULL | 0x01 | Incoming call(来电) |
| 0x02 | Missed call(未接来电) | 0x03 | Messages(短信) |
| 0x04 | Email(邮件) | 0x05 | Schedule (日程) |
| 0x06 | Facetime | 0x07 | QQ |
| 0x08 | Skype | 0x09 | Wechat(微信) |
| 0x0a | Whatsapp | 0x0b | Gmail |
| 0x0c | Hangout | 0x0d | Inbox |
| 0x0e | Line | 0x0f | Twitter(推特) |
| 0x10 | Facebook(脸书) | 0x11 | Facebook messenger |
| 0x12 | Instagram | 0x13 | Weibo(微博) |
| 0x14 | Kakaotalk | 0x15 | Facebook page manager |
| 0x16 | Viber | 0x17 | Vkclient |
| 0x18 | Telegram | 0x19 | Resv(保留) |
| 0x1a | Snapchat | 0x1b | DingTalk(钉钉) |
| 0x1c | Alipay(支付宝) | 0x1d | Tiktok(抖音) |
| 0x1e | Linkedln(领英) | 0x1f | Resv(保留) |

表7-3-1 消息号表

| Language ID | Type | Language ID | Type |
| --- | --- | --- | --- |
| BIT(0) | NULL | BIT(1) | Incoming call(来电) |
| BIT(2) | Missed call(未接来电) | BIT(3) | Messages(短信) |
| BIT(4) | Email(邮件) | BIT(5) | Schedule(日程) |
| BIT(6) | Facetime | BIT(7) | QQ |
| BIT(8) | Skype | BIT(9) | Wechat(微信) |
| BIT(10) | Whatsapp | BIT(11) | Gmail |
| BIT(12) | Hangout | BIT(13) | Inbox |
| BIT(14) | Line | BIT(15) | Twitter(推特) |
| BIT(16) | Facebook(脸书) | BIT(17) | Facebook messenger |
| BIT(18) | Instagram | BIT(19) | Weibo(微博) |
| BIT(20) | Kakaotalk | BIT(21) | Facebook page manager |
| BIT(22) | Viber | BIT(23) | Vkclient |
| BIT(24) | Telegram | BIT(25) | Resv(保留) |
| BIT(26) | Snapchat | BIT(27) | DingTalk(钉钉) |
| BIT(28) | Alipay(支付宝) | BIT(29) | Tiktok(抖音) |

| BIT(30) | Linkedln(领英) | BIT(31) | Resv(保留) |
| --- | --- | --- | --- |

表7-3-2 消息位表

4 7.4 开关表

| Message ID | Type | Message ID | Type |
| --- | --- | --- | --- |
| 0x00 | 防丢开关 | 0x01 | 抬手亮屏开关 |
| 0x02 | 自动同步开关 | 0x03 | Reserve |
| 0x04 | 睡眠监测开关 | 0x05 | 消息提醒总开关 |
| 0x06 | 整点上传运动数据开关 | 0x07 | 目标达成开关 |
| 0x08 | Reserve | 0x09 | 消息提醒亮屏开关 |
| 0x0a | 声音开关 | 0x0b | 震动总开关 |
| 0x0c | 整点上传健康数据开关 | 0x0d | 消息提醒震动开关 |

表7-4-1 开关号表

| Language ID | Type | Language ID | Type |
| --- | --- | --- | --- |
| BIT(0) | 防丢开关 | BIT(1) | 抬手亮屏开关 |
| BIT(2) | 自动同步开关 | BIT(3) | Reserve |
| BIT(4) | 睡眠监测开关 | BIT(5) | 消息提醒总开关 |
| BIT(6) | 整点上传运动数据开关 | BIT(7) | 目标达成开关 |
| BIT(8) | Reserve | BIT(9) | 消息提醒亮屏开关 |
| BIT(10) | 声音开关 | BIT(11) | 震动总开关 |
| BIT(12) | 整点上传健康数据开关 | BIT(13) | 消息提醒震动开关 |

表7-4-2 开关

