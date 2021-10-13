package utils

import "fmt"

var (
	ok          = NewBusinessException(0, "", "ok")
	NotFoundErr = NewBusinessException(404, "not found")
	ServeErr    = NewBusinessException(500, "serve error")
)

type BusinessException = ResultInfo

func NewBusinessException(code int, err string, msgs ...string) *BusinessException {
	msg := ""
	for _, m := range msgs {
		msg = msg + m
	}
	return &BusinessException{
		Code:    code,
		Err:     err,
		Message: msg,
	}
}

func (ex *BusinessException) Error() string {
	//return ex.Err
	return fmt.Sprintf("%s(%d)", ex.Err, ex.Code)
}

func OK(msgs ...string) *BusinessException {
	o := ok.Clone()

	msg := ""
	if len(msgs) > 0 {
		for _, v := range msgs {
			msg = msg + v
		}
		o.WithMessage(msg)
	}

	return o
}

func OK2(d interface{}, msgs ...string) *BusinessException {
	return OK(msgs...).WithData(d)
}

func IsOK(ex error) bool {
	e, c := ex.(*BusinessException)
	if !c {
		return false
	}

	return e.Err == ok.Err
}

func (ex *BusinessException) Clone() *BusinessException {
	return &BusinessException{
		Code:    ex.Code,
		Err:     ex.Err,
		Message: ex.Message,
		Data:    ex.Data,
	}
}
