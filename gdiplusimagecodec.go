package gdiplus

import (
	"unsafe"

	. "github.com/tryor/winapi"
)

//--------------------------------------------------------------------------
// Codec Management APIs
//--------------------------------------------------------------------------

func GetImageDecodersSize() (numDecoders UINT, size UINT, status Status) {
	status, _ = GdipGetImageDecodersSize(&numDecoders, &size)
	return
}

func GetImageDecoders(
	numDecoders UINT, size UINT) (decoders []*ImageCodecInfo, status Status) {

	decodersbuf := make([]byte, size)
	decoder := (*ImageCodecInfo)(unsafe.Pointer(&decodersbuf[0]))
	status, _ = GdipGetImageDecoders(numDecoders, size, decoder)
	if status == Ok {
		decoders = make([]*ImageCodecInfo, numDecoders)
		for i := UINT(0); i < numDecoders; i++ {
			decoder = (*ImageCodecInfo)(unsafe.Pointer((uintptr(unsafe.Pointer(decoder)) + uintptr(i*ImageCodecInfo_SIZE))))
			decoders[i] = decoder
		}
	}
	return
}

func GetImageEncodersSize() (numEncoders UINT, size UINT, status Status) {
	status, _ = GdipGetImageEncodersSize(&numEncoders, &size)
	return
}

func GetImageEncoders(
	numEncoders UINT, size UINT) (decoders []*ImageCodecInfo, status Status) {

	decodersbuf := make([]byte, size)
	decoder := (*ImageCodecInfo)(unsafe.Pointer(&decodersbuf[0]))
	status, _ = GdipGetImageEncoders(numEncoders, size, decoder)
	if status == Ok {
		decoders = make([]*ImageCodecInfo, numEncoders)
		for i := UINT(0); i < numEncoders; i++ {
			decoder = (*ImageCodecInfo)(unsafe.Pointer((uintptr(unsafe.Pointer(decoder)) + uintptr(i*ImageCodecInfo_SIZE))))
			decoders[i] = decoder
		}
	}
	return
}
