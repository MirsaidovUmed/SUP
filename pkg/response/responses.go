package response

var BadRequest = Response{
	Code:    400,
	Message: "Bad Request",
}

var Success = Response{
	Code:    200,
	Message: "Success",
}

var InternalServer = Response{
	Code:    500,
	Message: "Internal Error",
}

var NotFound = Response{
	Code:    404,
	Message: "Not Found",
}

var Unauthorized = Response{
	Code:    401,
	Message: "Unauthorized",
}
