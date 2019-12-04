package ecode

var (
	AuditStatusWait   = 0  //审核状态待审核
	AuditStatusOK     = 1  //审核状态通过
	AuditStatusNO     = -1 //审核状态不通过
	AuditStatusCancel = 2  //审核状态撤销

	AuditTemplateStatusNormal   = 0  //审核模版正常
	AuditTemplateStatusAbnormal = -1 //审核模版异常

	NodeAuditStatusTransfer = 3  //节点审核移交
	NotTransfer             = 0  //没有移交
	Transfer                = 1  //移交
	Origina                 = 0  //原始
	NotOrigina              = -1 //非原始
)
