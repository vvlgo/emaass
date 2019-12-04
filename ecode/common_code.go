package ecode

//公共ecode
var (
	OK = New(0, "OK")

	UploadDataParsingException  = New(1000, "上传数据解析异常")
	UUIDGetException            = New(1001, "UUID获取异常")
	InsertDBException           = New(1002, "数据插入异常")
	DataAlreadyExists           = New(1003, "数据已存在")
	UpdateDBException           = New(1004, "更新数据异常")
	JSONStringToStructException = New(1005, "json字符串转结构体异常")
	StructToJSONStringException = New(1006, "结构体转json字符串异常")
	DataInUse                   = New(1007, "数据正在被使用")
	DataLockException           = New(1008, "数据锁定异常")
	DataUnLockException         = New(1009, "数据解除锁定异常")
	UploadDataIsNull            = New(1010, "上传数据为空")
	CloseTheWindowTryAgain      = New(1011, "关闭窗口重新尝试")
)
