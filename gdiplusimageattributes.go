package gdiplus

import (
	. "github.com/trygo/winapi"
)

type ImageAttributes struct {
	GdiplusBase
	nativeImageAttr *GpImageAttributes
}

func NewImageAttributes() (*ImageAttributes, error) {
	imageattr := &ImageAttributes{}
	imageattr.setStatus(GdipCreateImageAttributes(&imageattr.nativeImageAttr))
	return imageattr, imageattr.LastError
}

func (this *ImageAttributes) Release() {
	if this.nativeImageAttr != nil {
		GdipDisposeImageAttributes(this.nativeImageAttr)
	}
}

func (this *ImageAttributes) Clone() (clone *ImageAttributes) {
	clone = &ImageAttributes{}

	this.setStatus(GdipCloneImageAttributes(
		this.nativeImageAttr,
		&clone.nativeImageAttr))
	clone.setStatus(this.LastResult, this.LastError)
	return
}

//default typ= ColorAdjustTypeDefault
func (this *ImageAttributes) SetToIdentity(typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesToIdentity(
		this.nativeImageAttr, typ))
}

//default typ = ColorAdjustTypeDefault
func (this *ImageAttributes) Reset(typ ColorAdjustType) Status {
	return this.setStatus(GdipResetImageAttributes(
		this.nativeImageAttr, typ))
}

//default mode = ColorMatrixFlagsDefault,
//default typ = ColorAdjustTypeDefault
func (this *ImageAttributes) SetColorMatrix(
	colorMatrix *ColorMatrix,
	mode ColorMatrixFlags,
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesColorMatrix(
		this.nativeImageAttr,
		typ,
		TRUE,
		colorMatrix,
		nil,
		mode))
}

//default typ = ColorAdjustTypeDefault
func (this *ImageAttributes) ClearColorMatrix(typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesColorMatrix(
		this.nativeImageAttr,
		typ,
		FALSE,
		nil,
		nil,
		ColorMatrixFlagsDefault))
}

//default mode = ColorMatrixFlagsDefault,
//default typ = ColorAdjustTypeDefault
func (this *ImageAttributes) SetColorMatrices(
	colorMatrix *ColorMatrix,
	grayMatrix *ColorMatrix,
	mode ColorMatrixFlags,
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesColorMatrix(
		this.nativeImageAttr,
		typ,
		TRUE,
		colorMatrix,
		grayMatrix,
		mode))
}

func (this *ImageAttributes) ClearColorMatrices(typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesColorMatrix(
		this.nativeImageAttr,
		typ,
		FALSE,
		nil,
		nil,
		ColorMatrixFlagsDefault))
}

//default typ = ColorAdjustTypeDefault
func (this *ImageAttributes) SetThreshold(
	threshold REAL,
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesThreshold(
		this.nativeImageAttr,
		typ,
		TRUE,
		threshold))
}

func (this *ImageAttributes) SetThreshold2(typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesThreshold(
		this.nativeImageAttr,
		typ,
		TRUE,
		0.0))
}

func (this *ImageAttributes) SetGamma(
	gamma REAL, typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesGamma(
		this.nativeImageAttr,
		typ,
		TRUE,
		gamma))
}

func (this *ImageAttributes) ClearGamma(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesGamma(
		this.nativeImageAttr,
		typ,
		FALSE,
		0.0))
}

func (this *ImageAttributes) SetNoOp(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesNoOp(
		this.nativeImageAttr,
		typ,
		TRUE))
}

func (this *ImageAttributes) ClearNoOp(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesNoOp(
		this.nativeImageAttr,
		typ,
		FALSE))
}

func (this *ImageAttributes) SetColorKey(
	colorLow, colorHigh Color,
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesColorKeys(
		this.nativeImageAttr,
		typ,
		TRUE,
		colorLow.GetValue(),
		colorHigh.GetValue()))
}

func (this *ImageAttributes) ClearColorKey(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesColorKeys(
		this.nativeImageAttr,
		typ,
		FALSE,
		0,
		0))
}

func (this *ImageAttributes) SetOutputChannel(
	channelFlags ColorChannelFlags,
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesOutputChannel(
		this.nativeImageAttr,
		typ,
		TRUE,
		channelFlags))
}

func (this *ImageAttributes) ClearOutputChannel(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesOutputChannel(
		this.nativeImageAttr,
		typ,
		FALSE,
		ColorChannelFlagsLast))
}

func (this *ImageAttributes) SetOutputChannelColorProfile(
	colorProfileFilename string,
	typ ColorAdjustType) Status {

	return this.setStatus(GdipSetImageAttributesOutputChannelColorProfile(
		this.nativeImageAttr,
		typ,
		TRUE,
		&colorProfileFilename))
}

func (this *ImageAttributes) ClearOutputChannelColorProfile(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesOutputChannelColorProfile(
		this.nativeImageAttr,
		typ,
		FALSE,
		nil))
}

func (this *ImageAttributes) SetRemapTable(
	maps []ColorMap,
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesRemapTable(
		this.nativeImageAttr,
		typ,
		TRUE,
		maps))
}

func (this *ImageAttributes) ClearRemapTable(
	typ ColorAdjustType) Status {
	return this.setStatus(GdipSetImageAttributesRemapTable(
		this.nativeImageAttr,
		typ,
		FALSE,
		nil))
}

func (this *ImageAttributes) SetBrushRemapTable(maps []ColorMap) Status {
	return this.SetRemapTable(maps, ColorAdjustTypeBrush)
}

func (this *ImageAttributes) ClearBrushRemapTable() Status {
	return this.ClearRemapTable(ColorAdjustTypeBrush)
}

//default color = Color()
//default clamp = FALSE
func (this *ImageAttributes) SetWrapMode(wrap WrapMode,
	color Color,
	clamp BOOL) Status {
	return this.setStatus(GdipSetImageAttributesWrapMode(
		this.nativeImageAttr, wrap, color.GetValue(), clamp))
}

// The flags of the palette are ignored.
//IN OUT colorPalette *ColorPalette
/*
Usage:
	colorPaletteCount := UINT(4)
	paletteBuff := make([]byte, ColorPalette_SIZE*(colorPaletteCount-1))
	colorPalette := (*ColorPalette)(unsafe.Pointer(&paletteBuff[0]))
	colorPalette.Flags = 0
	colorPalette.Count = colorPaletteCount
	imageAttributes.GetAdjustedPalette(colorPalette, ColorAdjustTypeBitmap)
	entrie := &colorPalette.Entries[0]
	for j := UINT(0); j < colorPaletteCount; j++ {
		entrie = (*ARGB)(unsafe.Pointer(uintptr(unsafe.Pointer(entrie)) + uintptr(j*ARGB_SIZE)))
		fmt.Printf("Entries[%v]:%v\n", j, *entrie)
	}

*/
func (this *ImageAttributes) GetAdjustedPalette(colorPalette *ColorPalette,
	colorAdjustType ColorAdjustType) Status {
	return this.setStatus(GdipGetImageAttributesAdjustedPalette(
		this.nativeImageAttr, colorPalette, colorAdjustType))
}
