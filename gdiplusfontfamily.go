package gdiplus

import (
	. "github.com/trygo/winapi"
	. "github.com/trygo/winapi/gdi"
	"syscall"
	"unsafe"
)

var GenericSansSerifFontFamily *FontFamily
var GenericSerifFontFamily *FontFamily
var GenericMonospaceFontFamily *FontFamily

var GenericSansSerifFontFamilyBuffer [unsafe.Sizeof(FontFamily{})]BYTE
var GenericSerifFontFamilyBuffer [unsafe.Sizeof(FontFamily{})]BYTE
var GenericMonospaceFontFamilyBuffer [unsafe.Sizeof(FontFamily{})]BYTE

type FontFamily struct {
	GdiplusBase
	nativeFamily *GpFontFamily
}

func NewFontFamily(name string, fontCollection IFontCollection) (*FontFamily, error) {
	var nativeFontCollection *GpFontCollection
	if fontCollection != nil {
		nativeFontCollection = fontCollection.GetNativeFontCollection()
	}
	family := &FontFamily{}
	var nativeFamily *GpFontFamily
	family.setStatus(GdipCreateFontFamilyFromName(name, nativeFontCollection, &nativeFamily))
	family.nativeFamily = nativeFamily
	return family, family.LastError
}

func (this *FontFamily) Close() {
	if this.nativeFamily != nil {
		this.setStatus(GdipDeleteFontFamily(this.nativeFamily))
	}
}

func (this *FontFamily) Clone() *FontFamily {
	family := &FontFamily{}
	this.setStatus(GdipCloneFontFamily(this.nativeFamily, &family.nativeFamily))
	family.setStatus(this.LastResult, this.LastError)
	return family
}

func (this *FontFamily) GetFamilyName(language LANGID) (name string, status Status) {
	namebuf := make([]uint16, LF_FACESIZE)
	status = this.setStatus(GdipGetFamilyName(this.nativeFamily,
		&namebuf[0],
		language))
	name = syscall.UTF16ToString(namebuf)
	return
}

func (this *FontFamily) IsStyleAvailable(style INT) BOOL {
	var styleAvailable BOOL
	this.setStatus(GdipIsStyleAvailable(this.nativeFamily, style, &styleAvailable))
	if this.LastResult != Ok {
		styleAvailable = FALSE
	}
	return styleAvailable
}

func (this *FontFamily) GetEmHeight(style INT) (emHeight UINT16) {
	this.setStatus(GdipGetEmHeight(this.nativeFamily, style, &emHeight))
	return
}

func (this *FontFamily) GetCellAscent(style INT) (cellAscent UINT16) {
	this.setStatus(GdipGetCellAscent(this.nativeFamily, style, &cellAscent))
	return
}

func (this *FontFamily) GetCellDescent(style INT) (cellDescent UINT16) {
	this.setStatus(GdipGetCellDescent(this.nativeFamily, style, &cellDescent))
	return
}

func (this *FontFamily) GetLineSpacing(style INT) (lineSpacing UINT16) {
	this.setStatus(GdipGetLineSpacing(this.nativeFamily, style, &lineSpacing))
	return
}

func GenericSansSerif() *FontFamily {
	if GenericSansSerifFontFamily != nil {
		return GenericSansSerifFontFamily
	}
	GenericSansSerifFontFamily =
		(*FontFamily)(unsafe.Pointer(&GenericSansSerifFontFamilyBuffer[0]))

	GenericSansSerifFontFamily.setStatus(GdipGetGenericFontFamilySansSerif(&GenericSansSerifFontFamily.nativeFamily))

	return GenericSansSerifFontFamily
}

func GenericSerif() *FontFamily {
	if GenericSerifFontFamily != nil {
		return GenericSerifFontFamily
	}
	GenericSerifFontFamily =
		(*FontFamily)(unsafe.Pointer(&GenericSerifFontFamilyBuffer[0]))

	GenericSerifFontFamily.setStatus(GdipGetGenericFontFamilySerif(
		&(GenericSerifFontFamily.nativeFamily)))

	return GenericSerifFontFamily
}

func GenericMonospace() *FontFamily {
	if GenericMonospaceFontFamily != nil {
		return GenericMonospaceFontFamily
	}

	GenericMonospaceFontFamily =
		(*FontFamily)(unsafe.Pointer(&GenericMonospaceFontFamilyBuffer[0]))

	GenericMonospaceFontFamily.setStatus(GdipGetGenericFontFamilyMonospace(
		&(GenericMonospaceFontFamily.nativeFamily)))

	return GenericMonospaceFontFamily
}
