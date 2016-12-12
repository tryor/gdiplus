package gdiplus

import (
	"unsafe"

	. "github.com/tryor/winapi"
)

//#if (GDIPVER >= 0x0110)
//----------------------------------------------------------------------------
// Color channel look up table (LUT)
//----------------------------------------------------------------------------

//typedef BYTE ColorChannelLUT[256];
type ColorChannelLUT [256]BYTE

//----------------------------------------------------------------------------
// Per-channel Histogram for 8bpp images.
//----------------------------------------------------------------------------

type HistogramFormat INT

const (
	HistogramFormatARGB HistogramFormat = iota
	HistogramFormatPARGB
	HistogramFormatRGB
	HistogramFormatGray
	HistogramFormatB
	HistogramFormatG
	HistogramFormatR
	HistogramFormatA
)

//#endif //(GDIPVER >= 0x0110)

//----------------------------------------------------------------------------
// Color matrix
//----------------------------------------------------------------------------

type ColorMatrix struct {
	m [5][5]REAL //REAL m[5][5];
}

var ColorMatrix_SIZE = UINT(unsafe.Sizeof(ColorMatrix{}))

//----------------------------------------------------------------------------
// Color Matrix flags
//----------------------------------------------------------------------------

type ColorMatrixFlags INT

const (
	ColorMatrixFlagsDefault   ColorMatrixFlags = 0
	ColorMatrixFlagsSkipGrays ColorMatrixFlags = 1
	ColorMatrixFlagsAltGray   ColorMatrixFlags = 2
)

//----------------------------------------------------------------------------
// Color Adjust Type
//----------------------------------------------------------------------------

type ColorAdjustType INT

const (
	ColorAdjustTypeDefault ColorAdjustType = iota
	ColorAdjustTypeBitmap
	ColorAdjustTypeBrush
	ColorAdjustTypePen
	ColorAdjustTypeText
	ColorAdjustTypeCount
	ColorAdjustTypeAny // Reserved
)

//----------------------------------------------------------------------------
// Color Map
//----------------------------------------------------------------------------

type ColorMap struct {
	OldColor Color
	NewColor Color
}
