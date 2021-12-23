package apicode

// 错误码规则:
// (1) 错误码需为 > 0 的数;
// (2) 错误码为 5 位数:
// -------------------------------------------
//     第1位          2、3位          4、5位
// -------------------------------------------
//   服务级错误码     模块级错误码	   具体错误码
// -------------------------------------------

var (
	// OK
	OK  = response(200, "ok") // 通用成功
	Err = response(500, "")   // 通用错误

	// 服务级错误码
	ErrParam               = response(10001, "参数有误")
	ErrSign                = response(10002, "签名参数有误")
	ErrParamNil            = response(10003, "参数不能为空")
	ErrParamBind           = response(10004, "参数格式错误,数据绑定失败")
	ErrNotFound            = response(10005, "请求的资源(接口)不存在")
	ErrMediaTypeNotSupport = response(10006, "文件类型不支持")

	// 模块级错误码 - 用户模块
	ErrUserService         = response(20100, "用户服务异常")
	ErrUserPhone           = response(20101, "用户手机号不合法")
	ErrUserCaptcha         = response(20102, "验证码有误")
	ErrUserAccount         = response(20103, "帐号错误")
	ErrUserPassword        = response(20104, "密码错误")
	ErrUserAccountNotExist = response(20105, "帐号不存在")
	ErrUserAccountExist    = response(20106, "帐号已存在")
	ErrUserPhoneExist      = response(20107, "手机号已存在")
	ErrUserEmailExist      = response(20108, "邮箱已存在")

	// 库存模块
	ErrOrderService = response(20200, "订单服务异常")
	ErrOrderOutTime = response(20201, "订单超时")

	// ......
)
