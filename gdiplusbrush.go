package gdiplus

import (
	. "github.com/trygo/winapi"
)

type IBrush interface {
	GetNativeBrush() *GpBrush
	SetNativeBrush(nativeBrush *GpBrush, status ...Status)
	GetType() BrushType
	Close()
}

type Brush struct {
	GdiplusBase
	nativeBrush *GpBrush
}

func newBrush() *Brush {
	return &Brush{GdiplusBase: GdiplusBase{LastResult: NotImplemented}}
}

func (this *Brush) Close() {
	if this.nativeBrush != nil {
		GdipDeleteBrush(this.nativeBrush)
	}
}

func (this *Brush) GetNativeBrush() *GpBrush {
	return this.nativeBrush
}

func (this *Brush) SetNativeBrush(nativeBrush *GpBrush, status ...Status) {
	if len(status) > 0 {
		this.LastResult = status[0]
	}
	this.nativeBrush = nativeBrush
}

func (this *Brush) GetType() BrushType {
	typ := GpBrushType(-1)
	this.setStatus(GdipGetBrushType(this.nativeBrush, &typ))
	return BrushType(typ)
}

func (this *Brush) Clone() *Brush {
	var brush *GpBrush
	this.setStatus(GdipCloneBrush(this.nativeBrush, &brush))
	if this.LastResult != Ok {
		return nil
	}
	return &Brush{nativeBrush: brush, GdiplusBase: GdiplusBase{LastResult: this.LastResult}}
}

//--------------------------------------------------------------------------
// Represent solid fill brush object
//--------------------------------------------------------------------------

type SolidBrush struct {
	Brush
}

func NewSolidBrush(color Color) (*SolidBrush, error) {
	sbrush := &SolidBrush{}
	var brush *GpSolidFill
	sbrush.setStatus(GdipCreateSolidFill(color.GetValue(), &brush))
	sbrush.SetNativeBrush((*GpBrush)(brush))
	return sbrush, sbrush.LastError
}

func (this *SolidBrush) GetColor() (Color, Status) {
	var argb ARGB
	this.setStatus(GdipGetSolidFillColor((*GpSolidFill)(this.nativeBrush), &argb))
	return Color{Argb: argb}, this.LastResult
}

func (this *SolidBrush) SetColor(color Color) Status {
	return this.setStatus(GdipSetSolidFillColor((*GpSolidFill)(this.nativeBrush), color.GetValue()))
}

//--------------------------------------------------------------------------
// Texture Brush Fill Object
//--------------------------------------------------------------------------

type TextureBrush struct {
	Brush
}

func NewTextureBrush(image IImage, wrapMode ...WrapMode) (*TextureBrush, error) {
	var wrapMode_ WrapMode
	if len(wrapMode) > 0 {
		wrapMode_ = wrapMode[0]
	} else {
		wrapMode_ = WrapModeTile
	}

	brush := &TextureBrush{}
	brush.setStatus(GdipCreateTexture(
		image.GetNativeImage(),
		wrapMode_, &brush.Brush.nativeBrush))
	return brush, brush.LastError
}

func NewTextureBrush2(image IImage, dstRect *RectF, wrapMode ...WrapMode) (*TextureBrush, error) {
	return NewTextureBrush3(image, dstRect.X, dstRect.Y, dstRect.Width, dstRect.Height, wrapMode...)
}

func NewTextureBrush3(image IImage, dstX, dstY, dstWidth, dstHeight REAL,
	wrapMode ...WrapMode) (*TextureBrush, error) {
	var wrapMode_ WrapMode
	if len(wrapMode) > 0 {
		wrapMode_ = wrapMode[0]
	} else {
		wrapMode_ = WrapModeTile
	}

	brush := &TextureBrush{}
	brush.setStatus(GdipCreateTexture2(
		image.GetNativeImage(),
		wrapMode_,
		dstX,
		dstY,
		dstWidth,
		dstHeight,
		&brush.Brush.nativeBrush))
	return brush, brush.LastError
}

func NewTextureBrushIA(image IImage, dstRect *RectF,
	imageAttributes *ImageAttributes) (*TextureBrush, error) {

	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	brush := &TextureBrush{}
	brush.setStatus(GdipCreateTextureIA(
		image.GetNativeImage(),
		nativeImageAttr,
		dstRect.X,
		dstRect.Y,
		dstRect.Width,
		dstRect.Height,
		&brush.Brush.nativeBrush))

	return brush, brush.LastError
}

func NewTextureBrushIAI(image IImage, dstRect *Rect,
	imageAttributes *ImageAttributes) (*TextureBrush, error) {

	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	brush := &TextureBrush{}
	brush.setStatus(GdipCreateTextureIAI(
		image.GetNativeImage(),
		nativeImageAttr,
		dstRect.X,
		dstRect.Y,
		dstRect.Width,
		dstRect.Height,
		&brush.Brush.nativeBrush))

	return brush, brush.LastError
}

func NewTextureBrush2I(image IImage, dstRect *Rect, wrapMode ...WrapMode) (*TextureBrush, error) {
	return NewTextureBrush3I(image, dstRect.X, dstRect.Y, dstRect.Width, dstRect.Height, wrapMode...)
}

func NewTextureBrush3I(image IImage, dstX, dstY, dstWidth, dstHeight INT, wrapMode ...WrapMode) (*TextureBrush, error) {
	var wrapMode_ WrapMode
	if len(wrapMode) > 0 {
		wrapMode_ = wrapMode[0]
	} else {
		wrapMode_ = WrapModeTile
	}

	brush := &TextureBrush{}
	brush.setStatus(GdipCreateTexture2I(
		image.GetNativeImage(),
		wrapMode_,
		dstX,
		dstY,
		dstWidth,
		dstHeight,
		&brush.Brush.nativeBrush))
	return brush, brush.LastError
}

func (this *TextureBrush) SetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipSetTextureTransform(this.nativeBrush,
		matrix.nativeMatrix))
}

//OUT matrix
func (this *TextureBrush) GetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipGetTextureTransform(this.nativeBrush,
		matrix.nativeMatrix))
}

func (this *TextureBrush) ResetTransform() Status {
	return this.setStatus(GdipResetTextureTransform(this.nativeBrush))
}

//order = MatrixOrderPrepend
func (this *TextureBrush) MultiplyTransform(matrix *Matrix,
	order MatrixOrder) Status {
	return this.setStatus(GdipMultiplyTextureTransform(this.nativeBrush,
		matrix.nativeMatrix, order))
}

//order = MatrixOrderPrepend
func (this *TextureBrush) TranslateTransform(dx, dy REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipTranslateTextureTransform(this.nativeBrush,
		dx, dy, order))
}

//order = MatrixOrderPrepend
func (this *TextureBrush) ScaleTransform(sx, sy REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipScaleTextureTransform(this.nativeBrush,
		sx, sy, order))
}

func (this *TextureBrush) RotateTransform(angle REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipRotateTextureTransform(this.nativeBrush,
		angle, order))
}

func (this *TextureBrush) SetWrapMode(wrapMode WrapMode) Status {
	return this.setStatus(GdipSetTextureWrapMode(this.nativeBrush,
		wrapMode))
}

func (this *TextureBrush) GetWrapMode() (wrapMode WrapMode) {
	this.setStatus(GdipGetTextureWrapMode(this.nativeBrush,
		&wrapMode))
	return
}

func (this *TextureBrush) GetImage() *Image {
	image := &Image{}
	this.setStatus(GdipGetTextureImage(this.nativeBrush,
		&image.nativeImage))
	return image
}

//--------------------------------------------------------------------------
// Linear Gradient Brush Object
//--------------------------------------------------------------------------

type LinearGradientBrush struct {
	Brush
}

func NewLinearGradientBrush(point1, point2 *PointF,
	color1, color2 Color) (*LinearGradientBrush, error) {
	brush := &LinearGradientBrush{}
	brush.setStatus(GdipCreateLineBrush(point1,
		point2,
		color1.GetValue(),
		color2.GetValue(),
		WrapModeTile,
		&brush.nativeBrush))
	return brush, brush.LastError
}

func NewLinearGradientBrushI(point1, point2 *Point,
	color1, color2 Color) (*LinearGradientBrush, error) {
	brush := &LinearGradientBrush{}
	brush.setStatus(GdipCreateLineBrushI(point1,
		point2,
		color1.GetValue(),
		color2.GetValue(),
		WrapModeTile,
		&brush.nativeBrush))
	return brush, brush.LastError
}

func NewLinearGradientBrushFromRect(rect *RectF,
	color1, color2 Color,
	mode LinearGradientMode) (*LinearGradientBrush, error) {
	brush := &LinearGradientBrush{}

	brush.setStatus(GdipCreateLineBrushFromRect(rect,
		color1.GetValue(),
		color2.GetValue(),
		mode,
		WrapModeTile,
		&brush.nativeBrush))

	return brush, brush.LastError
}

func NewLinearGradientBrushFromRectI(rect *Rect,
	color1, color2 Color,
	mode LinearGradientMode) (*LinearGradientBrush, error) {
	brush := &LinearGradientBrush{}

	brush.setStatus(GdipCreateLineBrushFromRectI(rect,
		color1.GetValue(),
		color2.GetValue(),
		mode,
		WrapModeTile,
		&brush.nativeBrush))

	return brush, brush.LastError
}

//isAngleScalable = FALSE
func NewLinearGradientBrushFromRectWithAngle(rect *RectF,
	color1, color2 Color,
	angle REAL,
	isAngleScalable BOOL) (*LinearGradientBrush, error) {
	brush := &LinearGradientBrush{}

	brush.setStatus(GdipCreateLineBrushFromRectWithAngle(rect,
		color1.GetValue(),
		color2.GetValue(),
		angle,
		isAngleScalable,
		WrapModeTile,
		&brush.nativeBrush))

	return brush, brush.LastError
}

func NewLinearGradientBrushFromRectWithAngleI(rect *Rect,
	color1, color2 Color,
	angle REAL,
	isAngleScalable BOOL) (*LinearGradientBrush, error) {
	brush := &LinearGradientBrush{}

	brush.setStatus(GdipCreateLineBrushFromRectWithAngleI(rect,
		color1.GetValue(),
		color2.GetValue(),
		angle,
		isAngleScalable,
		WrapModeTile,
		&brush.nativeBrush))

	return brush, brush.LastError
}

func (this *LinearGradientBrush) SetLinearColors(color1, color2 Color) Status {
	return this.setStatus(GdipSetLineColors(this.nativeBrush,
		color1.GetValue(),
		color2.GetValue()))
}

func (this *LinearGradientBrush) GetLinearColors() (color1, color2 Color, status Status) {
	var argbs = make([]ARGB, 2) //[2]ARGB
	status = this.setStatus(GdipGetLineColors(this.nativeBrush, argbs))
	if status == Ok {
		// use bitwise copy operator for Color copy
		color1.SetValue(argbs[0])
		color2.SetValue(argbs[1])
	}
	return
}

func (this *LinearGradientBrush) GetRectangle() (rect *RectF, status Status) {
	rect = &RectF{}
	status = this.setStatus(GdipGetLineRect(this.nativeBrush, rect))
	return
}

func (this *LinearGradientBrush) GetRectangleI() (rect *Rect, status Status) {
	rect = &Rect{}
	status = this.setStatus(GdipGetLineRectI(this.nativeBrush, rect))
	return
}

func (this *LinearGradientBrush) SetGammaCorrection(useGammaCorrection BOOL) Status {
	return this.setStatus(GdipSetLineGammaCorrection(
		this.nativeBrush, useGammaCorrection))
}

func (this *LinearGradientBrush) GetGammaCorrection() (useGammaCorrection BOOL) {
	this.setStatus(GdipGetLineGammaCorrection(
		this.nativeBrush, &useGammaCorrection))
	return
}

func (this *LinearGradientBrush) GetBlendCount() (count INT) {
	this.setStatus(GdipGetLineBlendCount(
		this.nativeBrush, &count))
	return
}

func (this *LinearGradientBrush) SetBlend(blendFactors, blendPositions []REAL) Status {
	count := len(blendFactors)
	if count != len(blendPositions) || count <= 0 {
		return this.setStatus(InvalidParameter, nil)
	}
	return this.setStatus(GdipSetLineBlend(
		this.nativeBrush, blendFactors, blendPositions, INT(count)))
}

func (this *LinearGradientBrush) GetBlend(count INT) (blendFactors, blendPositions []REAL, status Status) {
	blendFactors = make([]REAL, count)
	blendPositions = make([]REAL, count)
	status = this.setStatus(GdipGetLineBlend(
		this.nativeBrush,
		blendFactors, blendPositions, count))
	return
}

func (this *LinearGradientBrush) GetInterpolationColorCount() (count INT) {
	this.setStatus(GdipGetLinePresetBlendCount(
		this.nativeBrush, &count))
	return
}

func (this *LinearGradientBrush) SetInterpolationColors(presetColors []Color,
	blendPositions []REAL) Status {
	count := len(presetColors)
	if count != len(blendPositions) || count <= 0 {
		return this.setStatus(InvalidParameter, nil)
	}

	argbs := make([]ARGB, count)
	for i := int(0); i < count; i++ {
		argbs[i] = presetColors[i].GetValue()
	}
	return this.setStatus(GdipSetLinePresetBlend(
		this.nativeBrush,
		argbs,
		blendPositions,
		INT(count)))
}

func (this *LinearGradientBrush) GetInterpolationColors(count INT) (presetColors []Color, blendPositions []REAL, status Status) {
	if count <= 0 {
		status = this.setStatus(InvalidParameter, nil)
		return
	}

	argbs := make([]ARGB, count)
	blendPositions = make([]REAL, count)
	presetColors = make([]Color, count)
	status = this.setStatus(GdipGetLinePresetBlend(
		this.nativeBrush,
		argbs,
		blendPositions,
		count))

	for i := INT(0); i < count; i++ {
		presetColors[i].Argb = argbs[i]
	}
	return
}

//scale = 1.0
func (this *LinearGradientBrush) SetBlendBellShape(focus REAL, scale REAL) Status {
	return this.setStatus(GdipSetLineSigmaBlend(
		this.nativeBrush, focus, scale))
}

//scale = 1.0
func (this *LinearGradientBrush) SetBlendTriangularShape(focus REAL, scale REAL) Status {
	return this.setStatus(GdipSetLineLinearBlend(
		this.nativeBrush, focus, scale))
}

func (this *LinearGradientBrush) GetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipGetLineTransform(
		this.nativeBrush,
		matrix.nativeMatrix))
}

func (this *LinearGradientBrush) SetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipSetLineTransform(
		this.nativeBrush,
		matrix.nativeMatrix))
}

func (this *LinearGradientBrush) ResetTransform() Status {
	return this.setStatus(GdipResetLineTransform(
		this.nativeBrush))
}

//order = MatrixOrderPrepend
func (this *LinearGradientBrush) MultiplyTransform(matrix *Matrix, order MatrixOrder) Status {
	return this.setStatus(GdipMultiplyLineTransform(
		this.nativeBrush,
		matrix.nativeMatrix,
		order))
}

//order = MatrixOrderPrepend
func (this *LinearGradientBrush) TranslateTransform(dx, dy REAL, order MatrixOrder) Status {
	return this.setStatus(GdipTranslateLineTransform(
		this.nativeBrush,
		dx, dy, order))
}

//order = MatrixOrderPrepend
func (this *LinearGradientBrush) ScaleTransform(sx, sy REAL, order MatrixOrder) Status {
	return this.setStatus(GdipScaleLineTransform(
		this.nativeBrush,
		sx, sy, order))
}

//order = MatrixOrderPrepend
func (this *LinearGradientBrush) RotateTransform(angle REAL, order MatrixOrder) Status {
	return this.setStatus(GdipRotateLineTransform(
		this.nativeBrush, angle, order))
}

func (this *LinearGradientBrush) GetWrapMode() (wrapMode WrapMode) {
	this.setStatus(GdipGetLineWrapMode(
		this.nativeBrush, &wrapMode))
	return
}

func (this *LinearGradientBrush) SetWrapMode(wrapMode WrapMode) Status {
	return this.setStatus(GdipSetLineWrapMode(
		this.nativeBrush, wrapMode))
}

//--------------------------------------------------------------------------
// Hatch Brush Object
//--------------------------------------------------------------------------

type HatchBrush struct {
	Brush
}

func NewHatchBrush(hatchStyle HatchStyle,
	foreColor Color, backColor Color) (*HatchBrush, error) {
	brush := &HatchBrush{}

	brush.setStatus(GdipCreateHatchBrush(hatchStyle,
		foreColor.GetValue(),
		backColor.GetValue(),
		&brush.nativeBrush))
	return brush, brush.LastError
}

func (this *HatchBrush) GetHatchStyle() (hatchStyle HatchStyle) {
	this.setStatus(GdipGetHatchStyle(this.nativeBrush,
		&hatchStyle))
	return
}

func (this *HatchBrush) GetForegroundColor() (color Color) {
	this.setStatus(GdipGetHatchForegroundColor(
		this.nativeBrush,
		&color.Argb))
	return
}

func (this *HatchBrush) GetBackgroundColor() (color Color) {
	this.setStatus(GdipGetHatchBackgroundColor(
		this.nativeBrush,
		&color.Argb))
	return
}
