package gdiplus

import (
	"errors"
)

type GdiplusBase struct {
	LastResult Status
	LastError  error
}

func (this *GdiplusBase) setStatus(status Status, err error) Status {
	this.LastResult = status
	if err == nil && this.LastResult != Ok {
		this.LastError = errors.New(StatusText[this.LastResult])
	} else {
		this.LastError = err
	}
	return this.LastResult
}

//func (this *GdiplusBase) SetStatus(status Status) Status {
//	if status != Ok {
//		this.LastResult = status
//		this.LastError = errors.New(StatusText[this.LastResult])
//		return this.LastResult
//	} else {
//		return status
//	}
//}

func (this *GdiplusBase) GetLastStatus() Status {
	lastStatus := this.LastResult
	this.LastResult = Ok
	return lastStatus
}
