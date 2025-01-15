package game

type Msg struct {
	Value   string
	IsError bool
	IsEmpty bool
}

func NewEmptyMsg() Msg {
	return Msg{
		Value:   "",
		IsError: false,
		IsEmpty: true,
	}
}

func NewErrorMsg(value string) Msg {
	return Msg{
		Value:   value,
		IsError: true,
		IsEmpty: false,
	}
}

func NewSuccessMsg(value string) Msg {
	return Msg{
		Value:   value,
		IsError: false,
		IsEmpty: false,
	}
}
