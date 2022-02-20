package errcode

//Define common code/Error -> Error
var (
	Success = NewError(0, "succeed")

	ServerError              = NewError(100, "server internal error")
	InvalidParams            = NewError(101, "parameters invalid")
	NotFound                 = NewError(102, "not found")
	UnauthorizedAuthNotExist = NewError(103, "authorized failed, required key is not exist")

	UnauthorizedTokenError         = NewError(104, "authorized failed, token error")
	UnauthorizedTokenTimeOut       = NewError(105, "authorized failed, token expired")
	UnauthorizedTokenGenerateError = NewError(106, "authorized failed, token generation failed")
	TooManyRequest                 = NewError(107, "request is overloaded")
)
