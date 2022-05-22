package response

const (
	MsgError = "error"
	MsgOK    = "OK"
)

/* type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
*/

type response struct {
	MessageOK    message     `json:"message_ok"`
	MessageError message     `json:"message_error"`
	Data         interface{} `json:"data"`
}

type message struct {
	MessageCode uint16 `json:"code"`
	Message     string `json:"message"`
}

func New(msgType string, msgContent string, msgCode uint16, data interface{}) response {
	newMessage := message{
		MessageCode: msgCode,
		Message:     msgContent,
	}
	if msgType == MsgOK {
		return response{
			MessageOK:    newMessage,
			MessageError: message{},
			Data:         data,
		}
	} else {
		return response{
			MessageOK:    message{},
			MessageError: newMessage,
			Data:         data,
		}
	}
}

// func JSONResponse(msgType string, msgContent string, msgCode uint16, data interface{}, c echo.Context) error {
// 	response := NewResponse(msgType, msgContent, msgCode, data)
// 	return c.JSON(int(msgCode), response)
// }
