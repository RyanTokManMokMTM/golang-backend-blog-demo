package errcode

//Define common ErrorCode/Error -> Error
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
	ErrorGetTagListFail = NewError(2001001, "Get Tag List Failed")
	ErrorCreateTagFail  = NewError(2001002, "Create Tag Failed")
	ErrorUpdateTagFail  = NewError(2001003, "Update Tag Failed")
	ErrorDeleteTagFail  = NewError(2001004, "Delete Tag Failed")
	ErrorCountTagFail   = NewError(2001005, "Count Tag Failed")
)

var (
	ErrorGetArticleFailed     = NewError(2002001, "Get Article Failed")
	ErrorCreateArticleFailed  = NewError(2002002, "Create Article Failed")
	ErrorUpdateArticleFailed  = NewError(2002003, "Update Article Failed")
	ErrorDeleteArticleFailed  = NewError(2002004, "Delete Article Failed")
	ErrorGetArticleListFailed = NewError(2002005, "Get Article List Failed")
	//ErrorCountArticleFailed   = NewError(2002006, "Count Article Failed")
)

var (
	ErrorUploadFileFailed = NewError(2003001, "Upload File Failed")
)
