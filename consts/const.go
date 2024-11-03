package consts

const (
	// 成功
	SuccessCode   = 200 // 請求成功
	CreatedCode   = 201 // 資源創建成功
	NoContentCode = 204 // 請求成功但無內容

	// 客戶端錯誤
	BadRequestCode       = 400 // 無效請求
	UnauthorizedCode     = 401 // 未授權
	ForbiddenCode        = 403 // 禁止訪問
	NotFoundCode         = 404 // 資源不存在
	MethodNotAllowedCode = 405 // 方法不允許

	// 伺服器錯誤
	InternalServerErrorCode = 500 // 伺服器錯誤
	ServiceUnavailableCode  = 503 // 服務不可用

	// 自定義錯誤碼 (業務邏輯相關)
	InvalidParameterCode     = 1001 // 無效的參數
	DatabaseErrorCode        = 1002 // 資料庫錯誤
	AuthenticationFailedCode = 1003 // 驗證失敗
	ResourceConflictCode     = 1004 // 資源衝突
	PermissionDeniedCode     = 1005 // 權限不足
	BalanceNotEnough        = 1006 // 餘額不足
)
