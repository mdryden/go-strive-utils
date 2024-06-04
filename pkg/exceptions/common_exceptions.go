package exceptions

func ServerError(err error) Exception {
	return Exception{
		FullError: err,
		Code:      500,
		Message:   "Internal Server Error",
		Details:   "An unexpected error occurred",
	}
}

