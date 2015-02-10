package gdiplus

import (
	. "github.com/trygo/winapi"
)

type ICustomLineCap interface {
	GetNativeCap() *GpCustomLineCap
}

type CustomLineCap struct {
	GdiplusBase
	nativeCap *GpCustomLineCap
}

func NewCustomLineCap(
	fillPath *GraphicsPath,
	strokePath *GraphicsPath,
	baseCap LineCap,
	baseInset REAL) (*CustomLineCap, error) {

	linecap := &CustomLineCap{}

	var nativeFillPath *GpPath
	var nativeStrokePath *GpPath

	if fillPath != nil {
		nativeFillPath = fillPath.nativePath
	}
	if strokePath != nil {
		nativeStrokePath = strokePath.nativePath
	}

	linecap.setStatus(GdipCreateCustomLineCap(
		nativeFillPath, nativeStrokePath,
		baseCap, baseInset, &linecap.nativeCap))

	return linecap, linecap.LastError
}

func (this *CustomLineCap) GetNativeCap() *GpCustomLineCap {
	return this.nativeCap
}

func (this *CustomLineCap) Close() {
	if this.nativeCap != nil {
		GdipDeleteCustomLineCap(this.nativeCap)
	}
}

func (this *CustomLineCap) SetStrokeCaps(
	startCap, endCap LineCap) Status {
	return this.setStatus(GdipSetCustomLineCapStrokeCaps(this.nativeCap,
		startCap, endCap))
}

func (this *CustomLineCap) GetStrokeCaps() (startCap, endCap LineCap, status Status) {
	status = this.setStatus(GdipGetCustomLineCapStrokeCaps(this.nativeCap,
		&startCap, &endCap))
	return
}

func (this *CustomLineCap) SetStrokeJoin(lineJoin LineJoin) Status {
	return this.setStatus(GdipSetCustomLineCapStrokeJoin(this.nativeCap,
		lineJoin))
}

func (this *CustomLineCap) GetStrokeJoin() (lineJoin LineJoin) {
	this.setStatus(GdipGetCustomLineCapStrokeJoin(this.nativeCap,
		&lineJoin))
	return
}

func (this *CustomLineCap) SetBaseCap(baseCap LineCap) Status {
	return this.setStatus(GdipSetCustomLineCapBaseCap(this.nativeCap,
		baseCap))
}

func (this *CustomLineCap) GetBaseCap() (baseCap LineCap) {
	this.setStatus(GdipGetCustomLineCapBaseCap(this.nativeCap, &baseCap))
	return
}

func (this *CustomLineCap) SetBaseInset(inset REAL) Status {
	return this.setStatus(GdipSetCustomLineCapBaseInset(this.nativeCap,
		inset))
}

func (this *CustomLineCap) GetBaseInset() (inset REAL) {
	this.setStatus(GdipGetCustomLineCapBaseInset(this.nativeCap, &inset))
	return
}

func (this *CustomLineCap) SetWidthScale(widthScale REAL) Status {
	return this.setStatus(GdipSetCustomLineCapWidthScale(this.nativeCap,
		widthScale))
}

func (this *CustomLineCap) GetWidthScale() (widthScale REAL) {
	this.setStatus(GdipGetCustomLineCapWidthScale(this.nativeCap, &widthScale))
	return
}

func (this *CustomLineCap) Clone() *CustomLineCap {
	linecap := &CustomLineCap{}

	this.setStatus(GdipCloneCustomLineCap(this.nativeCap,
		&linecap.nativeCap))

	return linecap
}

type AdjustableArrowCap struct {
	CustomLineCap
}

//isFilled = TRUE
func NewAdjustableArrowCap(
	height, width REAL, isFilled BOOL) (*AdjustableArrowCap, error) {
	arrowCap := &AdjustableArrowCap{}
	arrowCap.setStatus(GdipCreateAdjustableArrowCap(
		height, width, isFilled, &arrowCap.nativeCap))
	return arrowCap, arrowCap.LastError
}

func (this *AdjustableArrowCap) SetHeight(height REAL) Status {
	return this.setStatus(GdipSetAdjustableArrowCapHeight(
		this.nativeCap, height))
}

func (this *AdjustableArrowCap) GetHeight() (height REAL) {
	this.setStatus(GdipGetAdjustableArrowCapHeight(
		this.nativeCap, &height))
	return
}

func (this *AdjustableArrowCap) SetWidth(width REAL) Status {
	return this.setStatus(GdipSetAdjustableArrowCapWidth(
		this.nativeCap, width))
}

func (this *AdjustableArrowCap) GetWidth() (width REAL) {
	this.setStatus(GdipGetAdjustableArrowCapWidth(
		this.nativeCap, &width))
	return
}

func (this *AdjustableArrowCap) SetMiddleInset(middleInset REAL) Status {
	return this.setStatus(GdipSetAdjustableArrowCapMiddleInset(
		this.nativeCap, middleInset))
}

func (this *AdjustableArrowCap) GetMiddleInset() (middleInset REAL) {
	this.setStatus(GdipGetAdjustableArrowCapMiddleInset(
		this.nativeCap, &middleInset))
	return
}

func (this *AdjustableArrowCap) SetFillState(isFilled BOOL) Status {
	return this.setStatus(GdipSetAdjustableArrowCapFillState(
		this.nativeCap, isFilled))
}

func (this *AdjustableArrowCap) IsFilled() (isFilled BOOL) {
	this.setStatus(GdipGetAdjustableArrowCapFillState(
		this.nativeCap, &isFilled))
	return
}
