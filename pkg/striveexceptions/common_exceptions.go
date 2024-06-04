package striveexceptions

const (
    ServerErrorCode    = 500
    ServerErrorMsg     = "Internal Server Error"
    ServerErrorDetails = "An unexpected error occurred"
)



func ServerError(err error, message string, details string) Exception {
    if message == "" {
        message = ServerErrorMsg
    }
    if details == "" {
        details = ServerErrorDetails
    }
    return Exception{
        FullError: err,
        Code:      ServerErrorCode,
        Message:   message,
        Details:   details,
    }
}