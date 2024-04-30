package code

type ErrorCode int

const (
	Success                    ErrorCode = 10000 // 成功
	AccountNotFoundError       ErrorCode = 10001 // 账号不存在
	AccountAlreadyExistError   ErrorCode = 10002 // 账号已存在
	MailCodeError              ErrorCode = 10003 // 邮箱验证码错误
	PhoneCodeError             ErrorCode = 10004 // 手机验证码错误
	CreateAccountError         ErrorCode = 10005 // 创建账号失败
	CreateAccountDescribeError ErrorCode = 10006 // 创建账号描述失败
	PasswdError                ErrorCode = 10007 // 密码错误
	ErrTicketNotExist          ErrorCode = 10008 // 凭证不存在
	ResetPasswdError           ErrorCode = 10009 // 重置密码失败
	ErrTicketTypeError         ErrorCode = 10010 // 凭证类型错误
	ErrTicketNeedApply         ErrorCode = 10011 // 凭证需要申请
	VerifyCodeError            ErrorCode = 10012 // 验证码错误
	ErrSqlExecError            ErrorCode = 10013 // 数据库执行错误
	ErrSystem                  ErrorCode = 10014 // 系统错误
	ErrTicketTypeExist         ErrorCode = 10015 // 凭证类型已存在
	ErrCardIDNotMatch          ErrorCode = 10016 // 身份证号码不匹配
	CaptchaExpired             ErrorCode = 10017 // 验证码过期
	CaptchaFail                ErrorCode = 10018 // 验证码错误
	VerifyCardIDError          ErrorCode = 10019 // 身份证号码验证错误
	DeviceIDEmpty              ErrorCode = 10020 // 设备ID为空
	ErrTooMuchDevice           ErrorCode = 10021 // 设备数量过多
	ErrorVerifyNewDevice       ErrorCode = 10022 // 验证新设备
	ErrCodeNotMatch            ErrorCode = 10023 // 验证码不匹配
	ErrEMailNotExit            ErrorCode = 10024 // 邮箱不存在
	JWTTokenFail               ErrorCode = 10025 // JWT Token无效
	PasswordHashError          ErrorCode = 10026 // 密码加密错误
	ErrorNotAdult              ErrorCode = 10027 // 未成年
	SendVerificationCodeError  ErrorCode = 10028 // 发送验证码失败
)

const (
	RoleBegin    = iota + 2000
	RoleNotExist // 角色不存在
)

var codeName = map[ErrorCode]string{
	Success:                    " 成功",
	AccountNotFoundError:       " 账号不存在",
	AccountAlreadyExistError:   " 账号已存在",
	MailCodeError:              " 邮箱验证码错误",
	PhoneCodeError:             " 手机验证码错误",
	CreateAccountError:         " 创建账号失败",
	CreateAccountDescribeError: " 创建账号描述失败",
	PasswdError:                " 密码错误",
	ErrTicketNotExist:          " 凭证不存在",
	ResetPasswdError:           " 重置密码失败",
	ErrTicketTypeError:         " 凭证类型错误",
	ErrTicketNeedApply:         " 凭证需要申请",
	VerifyCodeError:            " 验证码错误",
	ErrSqlExecError:            " 数据库执行错误",
	ErrSystem:                  " 系统错误",
	ErrTicketTypeExist:         " 凭证类型已存在",
	ErrCardIDNotMatch:          " 身份证号码不匹配",
	CaptchaExpired:             " 验证码过期",
	CaptchaFail:                " 验证码错误",
	VerifyCardIDError:          " 身份证号码验证错误",
	DeviceIDEmpty:              " 设备ID为空",
	ErrTooMuchDevice:           " 设备数量过多",
	ErrorVerifyNewDevice:       " 验证新设备",
	ErrCodeNotMatch:            " 验证码不匹配",
	ErrEMailNotExit:            " 邮箱不存在",
	JWTTokenFail:               " JWT Token无效",
	PasswordHashError:          " 密码加密错误",
	ErrorNotAdult:              " 未成年",
	SendVerificationCodeError:  " 发送验证码失败	",
}

func (c *ErrorCode) String() string {
	if name, ok := codeName[*c]; ok {
		return name
	}
	return "未知错误"
}
