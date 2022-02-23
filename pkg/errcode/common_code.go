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

var (
	ErrorGetTagListFail = NewError(201001, "Get Tag List Failed")
	ErrorCreateTagFail  = NewError(201002, "Create Tag Failed")
	ErrorUpdateTagFail  = NewError(201003, "Update Tag Failed")
	ErrorDeleteTagFail  = NewError(201004, "Delete Tag Failed")
	ErrorCountTagFail   = NewError(201005, "Count Tag Failed")
)