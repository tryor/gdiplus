package gdiplus

import (
	"unsafe"

	. "github.com/tryor/winapi"
	. "github.com/tryor/winapi/gdi"
)

type Font struct {
	GdiplusBase
	nativeFont *GpFont
}

func (this *Font) Release() {
	if this.nativeFont != nil {
		this.setStatus(GdipDeleteFont(this.nativeFont))
		this.nativeFont = nil
	}
}

func NewFont(family *FontFamily, emSize REAL, style FontStyle, unit Unit) (*Font, error) {
	var nativeFamily *GpFontFamily
	if family != nil {
		nativeFamily = family.nativeFamily
	}

	font := &Font{}
	var nativeFont *GpFont
	font.setStatus(GdipCreateFont(nativeFamily, emSize, style, unit, &nativeFont))
	font.nativeFont = nativeFont
	return font, font.LastError
}

func NewFont2(familyName string, emSize REAL, style FontStyle, unit Unit, fontCollection IFontCollection) (*Font, error) {

	font := &Font{}
	family, _ := NewFontFamily(familyName, fontCollection)
	defer family.Release()
	//nativeFamily := family.nativeFamily
	var nativeFamily *GpFontFamily
	if family != nil {
		nativeFamily = family.nativeFamily
		font.setStatus((family.GetLastStatus()), nil)
	}

	if font.LastResult != Ok {
		genericSansSerif := GenericSansSerif()
		nativeFamily = genericSansSerif.nativeFamily
		font.setStatus((genericSansSerif.LastResult), nil)
		if font.LastResult != Ok {
			return font, font.LastError
		}
	}

	font.setStatus(GdipCreateFont(nativeFamily,
		emSize,
		style,
		unit,
		&font.nativeFont))

	if font.LastResult != Ok {
		genericSansSerif := GenericSansSerif()
		nativeFamily = genericSansSerif.nativeFamily
		//lastResult = genericSansSerif.LastResult;
		if genericSansSerif.LastResult != Ok {
			return font, font.LastError
		}

		font.setStatus(GdipCreateFont(nativeFamily,
			emSize,
			style,
			unit,
			&font.nativeFont))
	}

	return font, font.LastError
}

//TODO 测试时报参数错误
func NewFont3(hdc HDC) (*Font, error) {
	font := &Font{}
	var nativeFont *GpFont
	font.setStatus(GdipCreateFontFromDC(hdc, &nativeFont))
	font.nativeFont = nativeFont
	return font, font.LastError
}

//TODO 测试时报参数错误
func NewFont4(hdc HDC, hfont HFONT) (*Font, error) {
	font := &Font{}
	if hfont != 0 {
		var lf LOGFONTA
		if GetObjectA(HANDLE(hfont), uintptr(LOGFONTA_SIZE), uintptr(unsafe.Pointer(&lf))) != 0 {
			font.setStatus(GdipCreateFontFromLogfontA(hdc, &lf, &font.nativeFont))
		} else {
			font.setStatus(GdipCreateFontFromDC(hdc, &font.nativeFont))
		}
	} else {
		font.setStatus(GdipCreateFontFromDC(hdc, &font.nativeFont))
	}
	return font, font.LastError
}

func NewFont5(hdc HDC, logfont *LOGFONTW) (*Font, error) {
	font := &Font{}
	if logfont != nil {
		font.setStatus(GdipCreateFontFromLogfontW(hdc, logfont, &font.nativeFont))
	} else {
		font.setStatus(GdipCreateFontFromDC(hdc, &font.nativeFont))
	}
	return font, font.LastError
}

func NewFont6(hdc HDC, logfont *LOGFONTA) (*Font, error) {
	font := &Font{}
	if logfont != nil {
		font.setStatus(GdipCreateFontFromLogfontA(hdc, logfont, &font.nativeFont))
	} else {
		font.setStatus(GdipCreateFontFromDC(hdc, &font.nativeFont))
	}
	return font, font.LastError
}

func (this *Font) GetLogFontA(g *Graphics) (logfontA *LOGFONTA, status Status) {
	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}
	logfontA = &LOGFONTA{}
	status = this.setStatus(GdipGetLogFontA(this.nativeFont, nativeGraphics, logfontA))
	return
}

func (this *Font) GetLogFontW(g *Graphics) (logfontW *LOGFONTW, status Status) {
	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}
	logfontW = &LOGFONTW{}
	status = this.setStatus(GdipGetLogFontW(this.nativeFont, nativeGraphics, logfontW))
	return
}

func (this *Font) Clone() (*Font, error) {
	cloneFont := &Font{}
	this.setStatus(GdipCloneFont(this.nativeFont, &cloneFont.nativeFont))
	return cloneFont, this.LastError
}

func (this *Font) IsAvailable() BOOL {
	return this.nativeFont != nil
}

func (this *Font) GetFamily() (family *FontFamily, status Status) {
	family = &FontFamily{}
	s, err := GdipGetFamily(this.nativeFont, &family.nativeFamily)
	this.setStatus(s, err)
	return family, family.setStatus(s, err)
}

func (this *Font) GetStyle() (style INT) {
	this.setStatus(GdipGetFontStyle(this.nativeFont, &style))
	return
}

func (this *Font) GetSize() (size REAL) {
	this.setStatus(GdipGetFontSize(this.nativeFont, &size))
	return
}

func (this *Font) GetUnit() (unit Unit) {
	this.setStatus(GdipGetFontUnit(this.nativeFont, &unit))
	return
}

func (this *Font) GetHeight(g *Graphics) (height REAL) {
	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}
	this.setStatus(GdipGetFontHeight(
		this.nativeFont,
		nativeGraphics,
		&height))
	return
}

func (this *Font) GetHeigh2(dpi REAL) (height REAL) {
	this.setStatus(GdipGetFontHeightGivenDPI(
		this.nativeFont,
		dpi,
		&height))
	return
}
