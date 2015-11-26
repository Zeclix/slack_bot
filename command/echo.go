package command

func EchoCommand(req Request) *Response {
	ret := new(Response)

	ret.Text = req.Text
	ret.ResponseType = deffered_in_channel

	return ret
}
