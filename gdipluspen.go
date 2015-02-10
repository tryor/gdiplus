package gdiplus

import (
	. "github.com/trygo/winapi"
)

type Pen struct {
	GdiplusBase
	nativePen *GpPen
}

func NewPen(color Color, width ...REAL) (*Pen, error) {
	var w REAL
	if len(width) > 0 {
		w = width[0]
	} else {
		w = 1.0
	}
	var unit Unit = UnitWorld
	var nativePen *GpPen
	lastResult, err := GdipCreatePen1(color.GetValue(), w, GpUnit(unit), &nativePen)

	return &Pen{nativePen: nativePen, GdiplusBase: GdiplusBase{LastResult: Status(lastResult), LastError: err}}, err
}

func NewPen2(brush IBrush, width ...REAL) (*Pen, error) {
	var w REAL
	if len(width) > 0 {
		w = width[0]
	} else {
		w = 1.0
	}
	pen := &Pen{}
	pen.setStatus(GdipCreatePen2(brush.GetNativeBrush(),
		w, UnitWorld, &pen.nativePen))

	return pen, pen.LastError
}

func (this *Pen) Close() {
	if this.nativePen != nil {
		GdipDeletePen(this.nativePen)
	}
}

func (this *Pen) Clone() *Pen {
	pen := &Pen{}
	pen.setStatus(GdipClonePen(this.nativePen, &pen.nativePen))
	return pen
}

func (this *Pen) SetWidth(width REAL) Status {
	return this.setStatus(GdipSetPenWidth(this.nativePen, width))
}

func (this *Pen) GetWidth() (width REAL) {
	this.setStatus(GdipGetPenWidth(this.nativePen, &width))
	return
}

func (this *Pen) SetLineCap(startCap, endCap LineCap,
	dashCap DashCap) Status {
	return this.setStatus(GdipSetPenLineCap197819(this.nativePen,
		startCap, endCap, dashCap))
}

func (this *Pen) SetStartCap(startCap LineCap) Status {
	return this.setStatus(GdipSetPenStartCap(this.nativePen, startCap))
}

func (this *Pen) SetEndCap(endCap LineCap) Status {
	return this.setStatus(GdipSetPenEndCap(this.nativePen, endCap))
}

func (this *Pen) SetDashCap(dashCap DashCap) Status {
	return this.setStatus(GdipSetPenDashCap197819(this.nativePen,
		dashCap))
}

func (this *Pen) GetStartCap() (startCap LineCap) {
	this.setStatus(GdipGetPenStartCap(this.nativePen, &startCap))
	return
}

func (this *Pen) GetEndCap() (endCap LineCap) {
	this.setStatus(GdipGetPenEndCap(this.nativePen, &endCap))
	return
}

func (this *Pen) GetDashCap() (dashCap DashCap) {
	this.setStatus(GdipGetPenDashCap197819(this.nativePen,
		&dashCap))
	return
}

func (this *Pen) SetLineJoin(lineJoin LineJoin) Status {
	return this.setStatus(GdipSetPenLineJoin(this.nativePen, lineJoin))
}

func (this *Pen) GetLineJoin() (lineJoin LineJoin) {
	this.setStatus(GdipGetPenLineJoin(this.nativePen, &lineJoin))
	return
}

func (this *Pen) SetCustomStartCap(customCap ICustomLineCap) Status {
	var nativeCap *GpCustomLineCap
	if customCap != nil {
		nativeCap = customCap.GetNativeCap()
	}

	return this.setStatus(GdipSetPenCustomStartCap(this.nativePen,
		nativeCap))
}

func (this *Pen) GetCustomStartCap() (customCap *CustomLineCap, status Status) {
	var nativeCap *GpCustomLineCap
	status = this.setStatus(GdipGetPenCustomStartCap(this.nativePen,
		&nativeCap))
	if status == Ok && nativeCap != nil {
		customCap = &CustomLineCap{nativeCap: nativeCap}
	}
	return
}

func (this *Pen) SetCustomEndCap(customCap ICustomLineCap) Status {
	var nativeCap *GpCustomLineCap
	if customCap != nil {
		nativeCap = customCap.GetNativeCap()
	}

	return this.setStatus(GdipSetPenCustomEndCap(this.nativePen,
		nativeCap))
}

func (this *Pen) GetCustomEndCap() (customCap *CustomLineCap, status Status) {
	var nativeCap *GpCustomLineCap
	status = this.setStatus(GdipGetPenCustomEndCap(this.nativePen, &nativeCap))
	if status == Ok && nativeCap != nil {
		customCap = &CustomLineCap{nativeCap: nativeCap}
	}
	return
}

func (this *Pen) SetMiterLimit(miterLimit REAL) Status {
	return this.setStatus(GdipSetPenMiterLimit(this.nativePen,
		miterLimit))
}

func (this *Pen) GetMiterLimit() (miterLimit REAL) {
	this.setStatus(GdipGetPenMiterLimit(this.nativePen, &miterLimit))
	return
}

func (this *Pen) SetAlignment(penAlignment PenAlignment) Status {
	return this.setStatus(GdipSetPenMode(this.nativePen, penAlignment))
}

func (this *Pen) GetAlignment() (penAlignment PenAlignment) {
	this.setStatus(GdipGetPenMode(this.nativePen, &penAlignment))
	return
}

func (this *Pen) SetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipSetPenTransform(this.nativePen,
		matrix.nativeMatrix))
}

//OUT Matrix* matrix
func (this *Pen) GetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipGetPenTransform(this.nativePen,
		matrix.nativeMatrix))
}

func (this *Pen) ResetTransform() Status {
	return this.setStatus(GdipResetPenTransform(this.nativePen))
}

//order = MatrixOrderPrepend
func (this *Pen) MultiplyTransform(matrix *Matrix,
	order MatrixOrder) Status {
	return this.setStatus(GdipMultiplyPenTransform(this.nativePen,
		matrix.nativeMatrix, order))
}

//order = MatrixOrderPrepend
func (this *Pen) TranslateTransform(dx, dy REAL, order MatrixOrder) Status {
	return this.setStatus(GdipTranslatePenTransform(
		this.nativePen,
		dx, dy, order))
}

//order = MatrixOrderPrepend
func (this *Pen) ScaleTransform(sx, sy REAL, order MatrixOrder) Status {
	return this.setStatus(GdipScalePenTransform(
		this.nativePen,
		sx, sy, order))
}

//order = MatrixOrderPrepend
func (this *Pen) RotateTransform(angle REAL, order MatrixOrder) Status {
	return this.setStatus(GdipRotatePenTransform(
		this.nativePen, angle, order))
}

func (this *Pen) GetPenType() (typ PenType) {
	this.setStatus(GdipGetPenFillType(this.nativePen, &typ))
	return
}

func (this *Pen) SetColor(color Color) Status {
	return this.setStatus(GdipSetPenColor(this.nativePen,
		color.GetValue()))
}

func (this *Pen) GetColor() (color Color, status Status) {
	typ := this.GetPenType()
	if typ != PenTypeSolidColor {
		status = WrongState
		return
	}
	status = this.setStatus(GdipGetPenColor(this.nativePen, &color.Argb))
	return
}

func (this *Pen) SetBrush(brush IBrush) Status {
	return this.setStatus(GdipSetPenBrushFill(this.nativePen,
		brush.GetNativeBrush()))
}

func (this *Pen) GetBrush() IBrush {
	typ := this.GetPenType()

	var brush IBrush
	switch typ {
	case PenTypeSolidColor:
		brush = &SolidBrush{}
	case PenTypeHatchFill:
		brush = &HatchBrush{}
	case PenTypeTextureFill:
		brush = &TextureBrush{}
	case PenTypePathGradient:
		brush = &Brush{}
	case PenTypeLinearGradient:
		brush = &LinearGradientBrush{}
	}

	if brush != nil {
		var nativeBrush *GpBrush
		this.setStatus(GdipGetPenBrushFill(this.nativePen,
			&nativeBrush))
		brush.SetNativeBrush(nativeBrush)
	}
	return brush
}

func (this *Pen) GetDashStyle() (dashStyle DashStyle) {
	this.setStatus(GdipGetPenDashStyle(this.nativePen, &dashStyle))
	return
}

func (this *Pen) SetDashStyle(dashStyle DashStyle) Status {
	return this.setStatus(GdipSetPenDashStyle(this.nativePen,
		dashStyle))
}

func (this *Pen) GetDashOffset() (dashOffset REAL) {
	this.setStatus(GdipGetPenDashOffset(this.nativePen, &dashOffset))
	return
}

func (this *Pen) SetDashOffset(dashOffset REAL) Status {
	return this.setStatus(GdipSetPenDashOffset(this.nativePen,
		dashOffset))
}

func (this *Pen) SetDashPattern(dashArray []REAL) Status {
	return this.setStatus(GdipSetPenDashArray(this.nativePen,
		dashArray))
}

func (this *Pen) GetDashPatternCount() (count INT) {
	this.setStatus(GdipGetPenDashCount(this.nativePen, &count))
	return
}

func (this *Pen) GetDashPattern(count INT) (dashArray []REAL, status Status) {
	if count <= 0 {
		status = this.setStatus(InvalidParameter, nil)
		return
	}
	dashArray = make([]REAL, count)
	status = this.setStatus(GdipGetPenDashArray(this.nativePen,
		dashArray))
	return
}

func (this *Pen) SetCompoundArray(compoundArray []REAL) Status {
	return this.setStatus(GdipSetPenCompoundArray(this.nativePen,
		compoundArray))
}

func (this *Pen) GetCompoundArrayCount() (count INT) {
	this.setStatus(GdipGetPenCompoundCount(this.nativePen, &count))
	return
}

func (this *Pen) GetCompoundArray(count INT) (compoundArray []REAL, status Status) {
	if count <= 0 {
		status = this.setStatus(InvalidParameter, nil)
		return
	}
	compoundArray = make([]REAL, count)
	status = this.setStatus(GdipGetPenCompoundArray(this.nativePen,
		compoundArray))
	return
}
