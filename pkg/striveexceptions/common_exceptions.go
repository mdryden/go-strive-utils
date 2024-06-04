package striveexceptions

const (
    ServerErrorCode    = 500
    ServerErrorMsg     = "Internal Server Error"
    ServerErrorDetails = "An unexpected error occurred"
)



func ServerError(err error, message *string, details *string) Exception {
    if message == nil {
        msg := ServerErrorMsg
        message = &msg
    }
    if details == nil {
        msg := ServerErrorDetails
        details = &msg
    }

    return Exception{
        FullError: err,
        Code:      ServerErrorCode,
        Message:   *message,
        Details:   *details,
    }
}