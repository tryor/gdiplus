package gdiplus

import (
	"errors"
	. "github.com/trygo/winapi"
)

type Graphics struct {
	GdiplusBase
	nativeGraphics *GpGraphics
	//LastResult     Status
	//LastError      error
}

func FromHDC(hdc HDC, hdevice ...HANDLE) (*Graphics, error) {
	return NewGraphicsFromHDC(hdc, hdevice...)
}

func FromHWND(hwnd HANDLE, icm ...bool) (*Graphics, error) {
	return NewGraphicsFromHWND(hwnd, icm...)
}

func FromImage(image IImage) (*Graphics, error) {
	return NewGraphicsFromImage(image)
}

func NewGraphicsFromHDC(hdc HDC, hdevice ...HANDLE) (*Graphics, error) {
	var gpg *GpGraphics
	var lastResult Status
	var err error
	if len(hdevice) > 0 {
		lastResult, err = GdipCreateFromHDC2(hdc, hdevice[0], &gpg)
	} else {
		lastResult, err = GdipCreateFromHDC(hdc, &gpg)
	}

	return &Graphics{nativeGraphics: gpg, GdiplusBase: GdiplusBase{LastResult: Status(lastResult), LastError: err}}, err
}

func NewGraphicsFromHWND(hwnd HANDLE, icm ...bool) (*Graphics, error) {
	var gpg *GpGraphics
	var lastResult Status
	var err error
	if len(icm) > 0 && icm[0] {
		lastResult, err = GdipCreateFromHWNDICM(hwnd, &gpg)
	} else {
		lastResult, err = GdipCreateFromHWND(hwnd, &gpg)
	}
	return &Graphics{nativeGraphics: gpg, GdiplusBase: GdiplusBase{LastResult: Status(lastResult), LastError: err}}, err
}

func NewGraphicsFromImage(image IImage) (*Graphics, error) {
	var gpg *GpGraphics
	lastResult, err := GdipGetImageGraphicsContext(image.GetNativeImage(), &gpg)
	return &Graphics{nativeGraphics: gpg, GdiplusBase: GdiplusBase{LastResult: Status(lastResult), LastError: err}}, err
}

func (this *Graphics) Release() {
	if this.nativeGraphics != nil {
		this.setStatus(GdipDeleteGraphics(this.nativeGraphics))
		this.nativeGraphics = nil
	}
}

func (this *Graphics) Flush(intention ...FlushIntention) {
	var intn FlushIntention
	if len(intention) > 0 {
		intn = intention[0]
	} else {
		intn = FlushIntentionFlush
	}
	GdipFlush(this.nativeGraphics, GpFlushIntention(intn))
}

// Locks the graphics until ReleaseDC is called
func (this *Graphics) GetHDC() HDC {
	var hdc HDC
	this.setStatus(GdipGetDC(this.nativeGraphics, &hdc))
	return hdc
}

func (this *Graphics) ReleaseHDC(hdc HDC) {
	this.setStatus(GdipReleaseDC(this.nativeGraphics, hdc))
}

//------------------------------------------------------------------------
// Rendering modes
//------------------------------------------------------------------------

func (this *Graphics) SetRenderingOrigin(in_x, in_y INT) Status {
	return this.setStatus(GdipSetRenderingOrigin(this.nativeGraphics, in_x, in_y))
}

func (this *Graphics) GetRenderingOrigin(out_x, out_y *INT) Status {
	return this.setStatus(GdipGetRenderingOrigin(this.nativeGraphics, out_x, out_y))
}

func (this *Graphics) SetCompositingMode(compositingMode CompositingMode) Status {
	return this.setStatus(GdipSetCompositingMode(this.nativeGraphics, compositingMode))
}

func (this *Graphics) GetCompositingMode() CompositingMode {
	var mode CompositingMode
	this.setStatus(GdipGetCompositingMode(this.nativeGraphics, &mode))
	return mode
}

func (this *Graphics) SetCompositingQuality(compositingQuality CompositingQuality) Status {
	return this.setStatus(GdipSetCompositingQuality(this.nativeGraphics, compositingQuality))
}

func (this *Graphics) GetCompositingQuality() CompositingQuality {
	var quality CompositingQuality
	this.setStatus(GdipGetCompositingQuality(this.nativeGraphics, &quality))
	return quality
}

func (this *Graphics) SetTextRenderingHint(newMode TextRenderingHint) Status {
	if !DCR_USE_NEW_186764() { //#ifndef DCR_USE_NEW_186764
		newMode = (newMode | 0x0f000)
	}
	return this.setStatus(GdipSetTextRenderingHint(this.nativeGraphics, newMode))
}

func (this *Graphics) GetTextRenderingHint() TextRenderingHint {
	var hint TextRenderingHint
	this.setStatus(GdipGetTextRenderingHint(this.nativeGraphics, &hint))
	return hint
}

func (this *Graphics) SetTextContrast(contrast UINT) Status {
	return this.setStatus(GdipSetTextContrast(this.nativeGraphics, contrast))
}

func (this *Graphics) GetTextContrast() UINT {
	var contrast UINT
	this.setStatus(GdipGetTextContrast(this.nativeGraphics, &contrast))
	return contrast
}

func (this *Graphics) SetTextGammaValue(gammaValue UINT) Status {
	return this.setStatus(GdipSetTextGammaValue(this.nativeGraphics, gammaValue))
}

func (this *Graphics) GetTextGammaValue() UINT {
	var gammaValue UINT
	this.setStatus(GdipGetTextGammaValue(this.nativeGraphics, &gammaValue))
	return gammaValue
}

func (this *Graphics) SetInterpolationMode(mode InterpolationMode) Status {
	return this.setStatus(GdipSetInterpolationMode(this.nativeGraphics, mode))
}

func (this *Graphics) GetInterpolationMode() InterpolationMode {
	var mode InterpolationMode = InterpolationModeInvalid
	this.setStatus(GdipGetInterpolationMode(this.nativeGraphics, &mode))
	return mode
}

func (this *Graphics) SetSmoothingMode(mode SmoothingMode) Status {
	return this.setStatus(GdipSetSmoothingMode(this.nativeGraphics, mode))
}

func (this *Graphics) GetSmoothingMode() SmoothingMode {
	var mode SmoothingMode = SmoothingModeInvalid
	this.setStatus(GdipGetSmoothingMode(this.nativeGraphics, &mode))
	return mode
}

func (this *Graphics) SetPixelOffsetMode(mode PixelOffsetMode) Status {
	return this.setStatus(GdipSetPixelOffsetMode(this.nativeGraphics, mode))
}

func (this *Graphics) GetPixelOffsetMode() PixelOffsetMode {
	var mode PixelOffsetMode = PixelOffsetModeInvalid
	this.setStatus(GdipGetPixelOffsetMode(this.nativeGraphics, &mode))
	return mode
}

func (this *Graphics) SetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipSetWorldTransform(this.nativeGraphics, matrix.nativeMatrix))
}

func (this *Graphics) ResetTransform() Status {
	return this.setStatus(GdipResetWorldTransform(this.nativeGraphics))
}

func (this *Graphics) MultiplyTransform(matrix *Matrix, order ...MatrixOrder) Status {
	var order_ MatrixOrder
	if len(order) > 0 {
		order_ = order[0]
	} else {
		order_ = MatrixOrderPrepend
	}
	return this.setStatus(GdipMultiplyWorldTransform(this.nativeGraphics, matrix.nativeMatrix, order_))
}

func (this *Graphics) TranslateTransform(dx, dy REAL,
	order ...MatrixOrder) Status {
	var order_ MatrixOrder
	if len(order) > 0 {
		order_ = order[0]
	} else {
		order_ = MatrixOrderPrepend
	}

	return this.setStatus(GdipTranslateWorldTransform(this.nativeGraphics,
		dx, dy, order_))
}

func (this *Graphics) ScaleTransform(sx, sy REAL,
	order ...MatrixOrder) Status {
	var order_ MatrixOrder
	if len(order) > 0 {
		order_ = order[0]
	} else {
		order_ = MatrixOrderPrepend
	}

	return this.setStatus(GdipScaleWorldTransform(this.nativeGraphics,
		sx, sy, order_))
}

func (this *Graphics) RotateTransform(angle REAL,
	order ...MatrixOrder) Status {
	var order_ MatrixOrder
	if len(order) > 0 {
		order_ = order[0]
	} else {
		order_ = MatrixOrderPrepend
	}

	return this.setStatus(GdipRotateWorldTransform(this.nativeGraphics,
		angle, order_))
}

func (this *Graphics) GetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipGetWorldTransform(this.nativeGraphics,
		matrix.nativeMatrix))
}

func (this *Graphics) SetPageScale(scale REAL) Status {
	return this.setStatus(GdipSetPageScale(this.nativeGraphics, scale))
}

func (this *Graphics) GetPageUnit() Unit {
	var unit Unit
	this.setStatus(GdipGetPageUnit(this.nativeGraphics, (*GpUnit)(&unit)))
	return unit
}

func (this *Graphics) GetPageScale() REAL {
	var scale REAL
	this.setStatus(GdipGetPageScale(this.nativeGraphics, &scale))
	return scale
}

func (this *Graphics) GetDpiX() REAL {
	var dpi REAL
	this.setStatus(GdipGetDpiX(this.nativeGraphics, &dpi))
	return dpi
}

func (this *Graphics) GetDpiY() REAL {
	var dpi REAL
	this.setStatus(GdipGetDpiY(this.nativeGraphics, &dpi))
	return dpi
}

func (this *Graphics) TransformPoints(destSpace CoordinateSpace, srcSpace CoordinateSpace, pts []PointF) ([]PointF, Status) {
	this.setStatus(GdipTransformPoints(this.nativeGraphics, GpCoordinateSpace(destSpace), GpCoordinateSpace(srcSpace), (*GpPointF)(&pts[0]), INT(len(pts))))
	return pts, this.LastResult
}

func (this *Graphics) TransformPoints2(destSpace CoordinateSpace, srcSpace CoordinateSpace, pts []Point) ([]Point, Status) {
	this.setStatus(GdipTransformPointsI(this.nativeGraphics, GpCoordinateSpace(destSpace), GpCoordinateSpace(srcSpace), (*GpPoint)(&pts[0]), INT(len(pts))))
	return pts, this.LastResult
}

//------------------------------------------------------------------------
// GetNearestColor (for <= 8bpp surfaces).  Note: Alpha is ignored.
//------------------------------------------------------------------------

func (this *Graphics) GetNearestColor(color Color) (Color, Status) {
	argb := color.GetValue()
	this.setStatus(GdipGetNearestColor(this.nativeGraphics, &argb))
	color.SetValue(argb)
	return color, this.LastResult
}

func (this *Graphics) DrawLine(pen *Pen, x1, y1, x2, y2 REAL) Status {
	return this.setStatus(GdipDrawLine(this.nativeGraphics, pen.nativePen, x1, y1, x2, y2))
}

func (this *Graphics) DrawLine2(pen *Pen, pt1, pt2 *PointF) Status {
	return this.DrawLine(pen, pt1.X, pt1.Y, pt2.X, pt2.Y)
}

func (this *Graphics) DrawLines(pen *Pen, points []PointF) Status {
	return this.setStatus(GdipDrawLines(this.nativeGraphics, pen.nativePen, (*GpPointF)(&points[0]), INT(len(points))))
}

func (this *Graphics) DrawLineI(pen *Pen, x1, y1, x2, y2 INT) Status {
	return this.setStatus(GdipDrawLineI(this.nativeGraphics, pen.nativePen, x1, y1, x2, y2))
}

func (this *Graphics) DrawLineI2(pen *Pen, pt1, pt2 *Point) Status {
	return this.DrawLineI(pen, pt1.X, pt1.Y, pt2.X, pt2.Y)
}

func (this *Graphics) DrawLinesI(pen *Pen, points []Point) Status {
	return this.setStatus(GdipDrawLinesI(this.nativeGraphics, pen.nativePen, (*GpPoint)(&points[0]), INT(len(points))))
}

func (this *Graphics) DrawArc(pen *Pen, x, y, width, height, startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipDrawArc(this.nativeGraphics, pen.nativePen, x, y, width, height, startAngle, sweepAngle))
}

func (this *Graphics) DrawArc2(pen *Pen, rect *RectF, startAngle, sweepAngle REAL) Status {
	return this.DrawArc(pen, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (this *Graphics) DrawArcI(pen *Pen, x, y, width, height INT, startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipDrawArcI(this.nativeGraphics, pen.nativePen, x, y, width, height, startAngle, sweepAngle))
}

func (this *Graphics) DrawArcI2(pen *Pen, rect *Rect, startAngle, sweepAngle REAL) Status {
	return this.DrawArcI(pen, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (this *Graphics) DrawBezier(pen *Pen, x1, y1, x2, y2, x3, y3, x4, y4 REAL) Status {
	return this.setStatus(GdipDrawBezier(this.nativeGraphics,
		pen.nativePen, x1, y1, x2, y2, x3, y3, x4, y4))
}

func (this *Graphics) DrawBezier2(pen *Pen, pt1, pt2, pt3, pt4 *PointF) Status {
	return this.DrawBezier(pen, pt1.X, pt1.Y, pt2.X, pt2.Y, pt3.X, pt3.Y, pt4.X, pt4.Y)
}

func (this *Graphics) DrawBeziers(pen *Pen, points []PointF) Status {
	return this.setStatus(GdipDrawBeziers(this.nativeGraphics,
		pen.nativePen,
		(*GpPointF)(&points[0]),
		INT(len(points))))
}

func (this *Graphics) DrawBezierI(pen *Pen, x1, y1, x2, y2, x3, y3, x4, y4 INT) Status {
	return this.setStatus(GdipDrawBezierI(this.nativeGraphics,
		pen.nativePen, x1, y1, x2, y2, x3, y3, x4, y4))
}

func (this *Graphics) DrawBezierI2(pen *Pen, pt1, pt2, pt3, pt4 *Point) Status {
	return this.DrawBezierI(pen, pt1.X, pt1.Y, pt2.X, pt2.Y, pt3.X, pt3.Y, pt4.X, pt4.Y)
}

func (this *Graphics) DrawBeziersI(pen *Pen, points []Point) Status {
	return this.setStatus(GdipDrawBeziersI(this.nativeGraphics,
		pen.nativePen,
		(*GpPoint)(&points[0]),
		INT(len(points))))
}

func (this *Graphics) DrawRectangle(pen *Pen, x, y, width, height REAL) Status {
	return this.setStatus(GdipDrawRectangle(this.nativeGraphics, pen.nativePen, x, y, width, height))
}

func (this *Graphics) DrawRectangle2(pen *Pen, rect *RectF) Status {
	return this.DrawRectangle(pen, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) DrawRectangles(pen *Pen, rects []RectF) Status {
	return this.setStatus(GdipDrawRectangles(this.nativeGraphics,
		pen.nativePen,
		(*GpRectF)(&rects[0]), INT(len(rects))))
}

func (this *Graphics) DrawRectangleI(pen *Pen, x, y, width, height INT) Status {
	return this.setStatus(GdipDrawRectangleI(this.nativeGraphics, pen.nativePen, x, y, width, height))
}

func (this *Graphics) DrawRectangleI2(pen *Pen, rect *Rect) Status {
	return this.DrawRectangleI(pen, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) DrawRectanglesI(pen *Pen, rects []Rect) Status {
	return this.setStatus(GdipDrawRectanglesI(this.nativeGraphics,
		pen.nativePen, (*GpRect)(&rects[0]), INT(len(rects))))
}

func (this *Graphics) DrawEllipse(pen *Pen, x, y, width, height REAL) Status {
	return this.setStatus(GdipDrawEllipse(this.nativeGraphics, pen.nativePen, x, y, width, height))
}

func (this *Graphics) DrawEllipse2(pen *Pen, rect *RectF) Status {
	return this.DrawEllipse(pen, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) DrawEllipseI(pen *Pen, x, y, width, height INT) Status {
	return this.setStatus(GdipDrawEllipseI(this.nativeGraphics, pen.nativePen, x, y, width, height))
}

func (this *Graphics) DrawEllipseI2(pen *Pen, rect *Rect) Status {
	return this.DrawEllipseI(pen, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) DrawPie(pen *Pen, x, y, width, height, startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipDrawPie(this.nativeGraphics,
		pen.nativePen, x, y, width, height, startAngle, sweepAngle))
}

func (this *Graphics) DrawPie2(pen *Pen, rect *RectF, startAngle, sweepAngle REAL) Status {
	return this.DrawPie(pen, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (this *Graphics) DrawPieI(pen *Pen, x, y, width, height INT, startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipDrawPieI(this.nativeGraphics,
		pen.nativePen, x, y, width, height, startAngle, sweepAngle))
}

func (this *Graphics) DrawPieI2(pen *Pen, rect *Rect, startAngle, sweepAngle REAL) Status {
	return this.DrawPieI(pen, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (this *Graphics) DrawPolygon(pen *Pen, points []PointF) Status {
	return this.setStatus(GdipDrawPolygon(this.nativeGraphics,
		pen.nativePen, (*GpPointF)(&points[0]), INT(len(points))))
}

func (this *Graphics) DrawPolygonI(pen *Pen, points []Point) Status {
	return this.setStatus(GdipDrawPolygonI(this.nativeGraphics,
		pen.nativePen, (*GpPoint)(&points[0]), INT(len(points))))
}

func (this *Graphics) DrawPath(pen *Pen, path *GraphicsPath) Status {
	var npen *GpPen
	if pen != nil {
		npen = pen.nativePen
	}
	var npath *GpPath
	if path != nil {
		npath = path.nativePath
	}
	return this.setStatus(GdipDrawPath(this.nativeGraphics, npen, npath))
}

func (this *Graphics) FillPath(brush IBrush, path *GraphicsPath) Status {
	return this.setStatus(GdipFillPath(this.nativeGraphics,
		brush.GetNativeBrush(),
		path.nativePath))
}

func (this *Graphics) SetPageUnit(unit Unit) Status {
	return this.setStatus(GdipSetPageUnit(this.nativeGraphics, unit))
}

func (this *Graphics) Clear(color Color) Status {
	return this.setStatus(GdipGraphicsClear(this.nativeGraphics, color.GetValue()))
}

func (this *Graphics) DrawCurve(pen *Pen, points []PointF, tension ...REAL) Status {
	if len(tension) > 0 {
		return this.setStatus(GdipDrawCurve2(this.nativeGraphics,
			pen.nativePen, (*GpPointF)(&points[0]), INT(len(points)), tension[0]))
	} else {
		return this.setStatus(GdipDrawCurve(this.nativeGraphics,
			pen.nativePen, (*GpPointF)(&points[0]), INT(len(points))))
	}
}

func (this *Graphics) DrawCurve2(pen *Pen, points []PointF, offset, numberOfSegments INT, tension ...REAL) Status {
	var tension_ REAL
	if len(tension) > 0 {
		tension_ = tension[0]
	} else {
		tension_ = 0.5
	}
	return this.setStatus(GdipDrawCurve3(this.nativeGraphics,
		pen.nativePen, (*GpPointF)(&points[0]), INT(len(points)), offset, numberOfSegments, tension_))
}

func (this *Graphics) DrawCurveI(pen *Pen, points []Point, tension ...REAL) Status {
	if len(tension) > 0 {
		return this.setStatus(GdipDrawCurve2I(this.nativeGraphics,
			pen.nativePen, (*GpPoint)(&points[0]), INT(len(points)), tension[0]))
	} else {
		return this.setStatus(GdipDrawCurveI(this.nativeGraphics,
			pen.nativePen, (*GpPoint)(&points[0]), INT(len(points))))
	}
}

func (this *Graphics) DrawCurveI2(pen *Pen, points []Point, offset, numberOfSegments INT, tension ...REAL) Status {
	var tension_ REAL
	if len(tension) > 0 {
		tension_ = tension[0]
	} else {
		tension_ = 0.5
	}
	return this.setStatus(GdipDrawCurve3I(this.nativeGraphics,
		pen.nativePen, (*GpPoint)(&points[0]), INT(len(points)), offset, numberOfSegments, tension_))
}

func (this *Graphics) DrawClosedCurve(pen *Pen, points []PointF, tension ...REAL) Status {
	if len(tension) > 0 {
		return this.setStatus(GdipDrawClosedCurve2(this.nativeGraphics,
			pen.nativePen, (*GpPointF)(&points[0]), INT(len(points)), tension[0]))
	} else {
		return this.setStatus(GdipDrawClosedCurve(this.nativeGraphics,
			pen.nativePen, (*GpPointF)(&points[0]), INT(len(points))))
	}
}

func (this *Graphics) DrawClosedCurveI(pen *Pen, points []Point, tension ...REAL) Status {
	if len(tension) > 0 {
		return this.setStatus(GdipDrawClosedCurve2I(this.nativeGraphics,
			pen.nativePen, (*GpPoint)(&points[0]), INT(len(points)), tension[0]))
	} else {
		return this.setStatus(GdipDrawClosedCurveI(this.nativeGraphics,
			pen.nativePen, (*GpPoint)(&points[0]), INT(len(points))))
	}
}

func (this *Graphics) FillRectangle(brush IBrush, x, y, width, height REAL) Status {
	return this.setStatus(GdipFillRectangle(this.nativeGraphics,
		brush.GetNativeBrush(), x, y, width, height))
}

func (this *Graphics) FillRectangle2(brush IBrush, rect *RectF) Status {
	return this.FillRectangle(brush, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) FillRectangles(brush IBrush, rects []RectF) Status {
	return this.setStatus(GdipFillRectangles(this.nativeGraphics,
		brush.GetNativeBrush(),
		(*GpRectF)(&rects[0]), INT(len(rects))))
}

func (this *Graphics) FillRectangleI(brush IBrush, x, y, width, height INT) Status {
	return this.setStatus(GdipFillRectangleI(this.nativeGraphics,
		brush.GetNativeBrush(), x, y, width, height))
}

func (this *Graphics) FillRectangleI2(brush IBrush, rect *Rect) Status {
	return this.FillRectangleI(brush, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) FillRectanglesI(brush IBrush, rects []Rect) Status {
	return this.setStatus(GdipFillRectanglesI(this.nativeGraphics,
		brush.GetNativeBrush(),
		(*GpRect)(&rects[0]), INT(len(rects))))
}

func (this *Graphics) FillPolygon(brush IBrush, points []PointF, fillMode ...FillMode) Status {
	var fillMode_ FillMode
	if len(fillMode) > 0 {
		fillMode_ = fillMode[0]
	} else {
		fillMode_ = FillModeAlternate
	}
	return this.setStatus(GdipFillPolygon(this.nativeGraphics,
		brush.GetNativeBrush(), (*GpPointF)(&points[0]), INT(len(points)), fillMode_))
}

func (this *Graphics) FillPolygonI(brush IBrush, points []Point, fillMode ...FillMode) Status {
	var fillMode_ FillMode
	if len(fillMode) > 0 {
		fillMode_ = fillMode[0]
	} else {
		fillMode_ = FillModeAlternate
	}
	return this.setStatus(GdipFillPolygonI(this.nativeGraphics,
		brush.GetNativeBrush(), (*GpPoint)(&points[0]), INT(len(points)), fillMode_))
}

func (this *Graphics) FillEllipse(brush IBrush, x, y, width, height REAL) Status {
	return this.setStatus(GdipFillEllipse(this.nativeGraphics, brush.GetNativeBrush(), x, y, width, height))
}

func (this *Graphics) FillEllipse2(brush IBrush, rect *RectF) Status {
	return this.FillEllipse(brush, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) FillEllipseI(brush IBrush, x, y, width, height INT) Status {
	return this.setStatus(GdipFillEllipseI(this.nativeGraphics, brush.GetNativeBrush(), x, y, width, height))
}

func (this *Graphics) FillEllipseI2(brush IBrush, rect *Rect) Status {
	return this.FillEllipseI(brush, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) FillPie(brush IBrush, x, y, width, height, startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipFillPie(this.nativeGraphics,
		brush.GetNativeBrush(), x, y, width, height, startAngle, sweepAngle))
}

func (this *Graphics) FillPie2(brush IBrush, rect *RectF, startAngle, sweepAngle REAL) Status {
	return this.FillPie(brush, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (this *Graphics) FillPieI(brush IBrush, x, y, width, height INT, startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipFillPieI(this.nativeGraphics,
		brush.GetNativeBrush(), x, y, width, height, startAngle, sweepAngle))
}

func (this *Graphics) FillPieI2(brush IBrush, rect *Rect, startAngle, sweepAngle REAL) Status {
	return this.FillPieI(brush, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (this *Graphics) FillClosedCurve(brush IBrush, points []PointF) Status {
	return this.setStatus(GdipFillClosedCurve(this.nativeGraphics,
		brush.GetNativeBrush(), (*GpPointF)(&points[0]), INT(len(points))))
}

func (this *Graphics) FillClosedCurve2(brush IBrush, points []PointF, fillMode FillMode, tension ...REAL) Status {
	var tension_ REAL
	if len(tension) > 0 {
		tension_ = tension[0]
	} else {
		tension_ = 0.5
	}
	return this.setStatus(GdipFillClosedCurve2(this.nativeGraphics,
		brush.GetNativeBrush(), (*GpPointF)(&points[0]), INT(len(points)), tension_, fillMode))
}

func (this *Graphics) FillClosedCurveI(brush IBrush, points []Point) Status {
	return this.setStatus(GdipFillClosedCurveI(this.nativeGraphics,
		brush.GetNativeBrush(), (*GpPoint)(&points[0]), INT(len(points))))
}

func (this *Graphics) FillClosedCurveI2(brush IBrush, points []Point, fillMode FillMode, tension ...REAL) Status {
	var tension_ REAL
	if len(tension) > 0 {
		tension_ = tension[0]
	} else {
		tension_ = 0.5
	}
	return this.setStatus(GdipFillClosedCurve2I(this.nativeGraphics,
		brush.GetNativeBrush(), (*GpPoint)(&points[0]), INT(len(points)), tension_, fillMode))
}

func (this *Graphics) FillRegion(brush IBrush, region *Region) Status {
	return this.setStatus(GdipFillRegion(this.nativeGraphics, brush.GetNativeBrush(), region.nativeRegion))
}

//注意: 调用此方法时一直返回GenericError，没找到原因。建议使用 GraphicsPath.AddString()代替
func (this *Graphics) DrawString(text string, font *Font, layoutRect *RectF, stringFormat *StringFormat, brush IBrush) Status {
	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}
	var nativeBrush *GpBrush
	if brush != nil {
		nativeBrush = brush.GetNativeBrush()
	}

	return this.setStatus(GdipDrawString(
		this.nativeGraphics,
		text,
		nativeFont,
		layoutRect,
		nativeFormat,
		nativeBrush))
}

func (this *Graphics) DrawDriverString(text string, font *Font, brush IBrush, positions *PointF, flags DriverStringOptions, matrix *Matrix) Status {
	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeMatrix *GpMatrix
	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}
	var nativeBrush *GpBrush
	if brush != nil {
		nativeBrush = brush.GetNativeBrush()
	}

	return this.setStatus(GdipDrawDriverString(
		this.nativeGraphics,
		text,
		nativeFont,
		nativeBrush,
		positions,
		flags,
		nativeMatrix))
}

func (this *Graphics) MeasureString(text string, font *Font, layoutRect *RectF, stringFormat *StringFormat) (boundingBox *RectF, codepointsFitted, linesFilled INT, status Status) {
	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}

	boundingBox = &RectF{}

	status = this.setStatus(GdipMeasureString(
		this.nativeGraphics,
		text,
		nativeFont,
		layoutRect,
		nativeFormat,
		boundingBox, &codepointsFitted, &linesFilled))

	if boundingBox.Width <= REAL_EPSILON {
		boundingBox.Width = 0.0
	}

	if boundingBox.Height <= REAL_EPSILON {
		boundingBox.Height = 0.0
	}

	return

}

func (this *Graphics) MeasureString2(text string, font *Font, layoutRectSize *SizeF, stringFormat *StringFormat) (size *SizeF, codepointsFitted INT, linesFilled INT, status Status) {

	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}

	layoutRect := NewRectF(0, 0, layoutRectSize.Width, layoutRectSize.Height)
	boundingBox := &RectF{}

	status = this.setStatus(GdipMeasureString(
		this.nativeGraphics,
		text,
		nativeFont,
		layoutRect,
		nativeFormat,
		boundingBox, &codepointsFitted, &linesFilled))

	size = &SizeF{}

	if status == Ok {
		if boundingBox.Width <= REAL_EPSILON {
			size.Width = 0.0
		} else {
			size.Width = boundingBox.Width
		}
		if boundingBox.Height <= REAL_EPSILON {
			size.Height = 0.0
		} else {
			size.Height = boundingBox.Height
		}
	}

	return
}

func (this *Graphics) MeasureString3(text string, font *Font, origin *PointF, stringFormat *StringFormat) (boundingBox *RectF, status Status) {
	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}

	rect := &RectF{origin.X, origin.Y, 0.0, 0.0}
	boundingBox = &RectF{}

	status = this.setStatus(GdipMeasureString(
		this.nativeGraphics,
		text,
		nativeFont,
		rect,
		nativeFormat,
		boundingBox,
		nil,
		nil))

	if boundingBox.Width <= REAL_EPSILON {
		boundingBox.Width = 0.0
	}

	if boundingBox.Height <= REAL_EPSILON {
		boundingBox.Height = 0.0
	}

	return
}

func (this *Graphics) MeasureString4(text string, font *Font, layoutRect *RectF) (boundingBox *RectF, status Status) {
	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	boundingBox = &RectF{}

	status = this.setStatus(GdipMeasureString(
		this.nativeGraphics,
		text,
		nativeFont,
		layoutRect,
		nil,
		boundingBox,
		nil,
		nil))

	if boundingBox.Width <= REAL_EPSILON {
		boundingBox.Width = 0.0
	}
	if boundingBox.Height <= REAL_EPSILON {
		boundingBox.Height = 0.0
	}
	return
}

func (this *Graphics) MeasureCharacterRanges(
	text string,
	font *Font,
	layoutRect *RectF,
	stringFormat *StringFormat,
	regions []*Region) (status Status) {

	if len(regions) <= 0 {
		this.LastResult = InvalidParameter
		this.LastError = errors.New(StatusText[InvalidParameter])
		return this.LastResult
	}

	regionCount := INT(len(regions))
	//regions = make([]Region, regionCount)

	//GpRegion **nativeRegions = new GpRegion* [regionCount];
	nativeRegions := make([]*GpRegion, regionCount)

	for i := INT(0); i < regionCount; i++ {
		nativeRegions[i] = regions[i].nativeRegion
	}

	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}
	status = this.setStatus(GdipMeasureCharacterRanges(
		this.nativeGraphics,
		text,
		nativeFont,
		layoutRect,
		nativeFormat,
		nativeRegions))

	return
}

func (this *Graphics) MeasureDriverString(
	text string,
	font *Font,
	positions *PointF,
	flags INT,
	matrix *Matrix) (boundingBox *RectF, status Status) {

	var nativeFont *GpFont
	if font != nil {
		nativeFont = font.nativeFont
	}
	var nativeMatrix *GpMatrix
	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}
	boundingBox = &RectF{}

	status = this.setStatus(GdipMeasureDriverString(
		this.nativeGraphics,
		text,
		nativeFont,
		positions,
		flags,
		nativeMatrix,
		boundingBox))

	return
}

func (this *Graphics) DrawCachedBitmap(cb *CachedBitmap, x, y INT) Status {
	return this.setStatus(GdipDrawCachedBitmap(
		this.nativeGraphics,
		cb.nativeCachedBitmap,
		x, y))
}

func (this *Graphics) DrawImage(image IImage, x, y REAL) Status {
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImage(this.nativeGraphics,
		nativeImage, x, y))
}

func (this *Graphics) DrawImage2(image IImage, point *PointF) Status {
	return this.DrawImage(image, point.X, point.Y)
}

func (this *Graphics) DrawImage3(image IImage, x, y, width, height REAL) Status {
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImageRect(this.nativeGraphics, nativeImage,
		x, y, width, height))
}

func (this *Graphics) DrawImage4(image IImage, rect *RectF) Status {
	return this.DrawImage3(image, rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) DrawImage5(image IImage, destPoints []PointF) Status {
	count := len(destPoints)
	if count != 3 && count != 4 {
		return this.setStatus(InvalidParameter, errors.New(StatusText[InvalidParameter]))
	}
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImagePoints(this.nativeGraphics,
		nativeImage, destPoints))
}

func (this *Graphics) DrawImage6(image IImage, x, y, srcx, srcy, srcwidth, srcheight REAL, srcUnit Unit) Status {
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImagePointRect(this.nativeGraphics, nativeImage,
		x, y, srcx, srcy, srcwidth, srcheight, GpUnit(srcUnit)))
}

func (this *Graphics) DrawImage7(image IImage,
	destRect *RectF,
	srcx,
	srcy,
	srcwidth,
	srcheight REAL,
	srcUnit Unit,
	imageAttributes *ImageAttributes,
	callback DrawImageAbort,
	callbackData DATA_PTR) Status {

	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipDrawImageRectRect(this.nativeGraphics,
		nativeImage,
		destRect.X,
		destRect.Y,
		destRect.Width,
		destRect.Height,
		srcx, srcy,
		srcwidth, srcheight,
		GpUnit(srcUnit),
		nativeImageAttr,
		callback,
		callbackData))
}

func (this *Graphics) DrawImage8(image IImage,
	destPoints []PointF,
	srcx,
	srcy,
	srcwidth,
	srcheight REAL,
	srcUnit Unit,
	imageAttributes *ImageAttributes,
	callback DrawImageAbort,
	callbackData DATA_PTR) Status {

	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipDrawImagePointsRect(this.nativeGraphics,
		nativeImage,
		destPoints,
		srcx, srcy,
		srcwidth, srcheight,
		GpUnit(srcUnit),
		nativeImageAttr,
		callback,
		callbackData))
}

func (this *Graphics) DrawImageI(image IImage, x, y INT) Status {
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImageI(this.nativeGraphics,
		nativeImage, x, y))
}

func (this *Graphics) DrawImageI2(image IImage, point *Point) Status {
	return this.DrawImageI(image, point.X, point.Y)
}

func (this *Graphics) DrawImageI3(image IImage, x, y, width, height INT) Status {
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImageRectI(this.nativeGraphics, nativeImage,
		x, y, width, height))
}

func (this *Graphics) DrawImageI4(image IImage, rect *Rect) Status {
	return this.DrawImageI3(image,
		rect.X,
		rect.Y,
		rect.Width,
		rect.Height)
}

func (this *Graphics) DrawImageI5(image IImage, destPoints []Point) Status {
	count := len(destPoints)
	if count != 3 && count != 4 {
		return this.setStatus(InvalidParameter, errors.New(StatusText[InvalidParameter]))
	}
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImagePointsI(this.nativeGraphics,
		nativeImage, destPoints))
}

func (this *Graphics) DrawImageI6(image IImage, x, y, srcx, srcy, srcwidth, srcheight INT, srcUnit Unit) Status {
	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	return this.setStatus(GdipDrawImagePointRectI(this.nativeGraphics, nativeImage,
		x, y, srcx, srcy, srcwidth, srcheight, GpUnit(srcUnit)))
}

func (this *Graphics) DrawImageI7(image IImage,
	destRect *Rect,
	srcx,
	srcy,
	srcwidth,
	srcheight INT,
	srcUnit Unit,
	imageAttributes *ImageAttributes,
	callback DrawImageAbort,
	callbackData DATA_PTR) Status {

	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipDrawImageRectRectI(this.nativeGraphics,
		nativeImage,
		destRect.X,
		destRect.Y,
		destRect.Width,
		destRect.Height,
		srcx, srcy,
		srcwidth, srcheight,
		GpUnit(srcUnit),
		nativeImageAttr,
		callback,
		callbackData))
}

func (this *Graphics) DrawImageI8(image IImage,
	destPoints []Point,
	srcx,
	srcy,
	srcwidth,
	srcheight INT,
	srcUnit Unit,
	imageAttributes *ImageAttributes,
	callback DrawImageAbort,
	callbackData DATA_PTR) Status {

	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipDrawImagePointsRectI(this.nativeGraphics,
		nativeImage,
		destPoints,
		srcx, srcy,
		srcwidth, srcheight,
		GpUnit(srcUnit),
		nativeImageAttr,
		callback,
		callbackData))
}

//#if (GDIPVER >= 0x0110)
func (this *Graphics) DrawImage9(
	image IImage,
	destRect *RectF,
	sourceRect *RectF,
	srcUnit Unit,
	imageAttributes *ImageAttributes) Status {

	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipDrawImageRectRect(
		this.nativeGraphics,
		nativeImage,
		destRect.X,
		destRect.Y,
		destRect.Width,
		destRect.Height,
		sourceRect.X,
		sourceRect.Y,
		sourceRect.Width,
		sourceRect.Height,
		GpUnit(srcUnit),
		nativeImageAttr,
		nil,
		DATA_PTR(0)))
}

func (this *Graphics) DrawImage10(
	image IImage,
	sourceRect *RectF,
	xForm *Matrix,
	effect IEffect,
	imageAttributes *ImageAttributes,
	srcUnit Unit) Status {

	var nativeImage *GpImage
	if image != nil {
		nativeImage = image.GetNativeImage()
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}
	var nativeMatrix *GpMatrix
	if xForm != nil {
		nativeMatrix = xForm.nativeMatrix
	}
	var nativeEffect *CGpEffect
	if effect != nil {
		nativeEffect = effect.GetNativeEffect()
	}

	return this.setStatus(GdipDrawImageFX(
		this.nativeGraphics,
		nativeImage,
		(*GpRectF)(sourceRect),
		nativeMatrix,
		nativeEffect,
		nativeImageAttr,
		GpUnit(srcUnit)))
}

//#endif

// The following methods are for playing an EMF+ to a graphics
// via the enumeration interface.  Each record of the EMF+ is
// sent to the callback (along with the callbackData).  Then
// the callback can invoke the Metafile::PlayRecord method
// to play the particular record.

func (this *Graphics) EnumerateMetafile(
	metafile *Metafile,
	destPoint *PointF,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileDestPoint(
		this.nativeGraphics,
		gpMetafile,
		destPoint,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafile2(
	metafile *Metafile,
	destRect *RectF,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileDestRect(
		this.nativeGraphics,
		gpMetafile,
		destRect,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafile3(
	metafile *Metafile,
	destPoints []PointF,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileDestPoints(
		this.nativeGraphics,
		gpMetafile,
		destPoints,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafileI(
	metafile *Metafile,
	destPoint *Point,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileDestPointI(
		this.nativeGraphics,
		gpMetafile,
		destPoint,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafileI2(
	metafile *Metafile,
	destRect *Rect,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileDestRectI(
		this.nativeGraphics,
		gpMetafile,
		destRect,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafileI3(
	metafile *Metafile,
	destPoints []Point,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileDestPointsI(
		this.nativeGraphics,
		gpMetafile,
		destPoints,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafile4(
	metafile *Metafile,
	destPoint *PointF,
	srcRect *RectF,
	srcUnit Unit,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileSrcRectDestPoint(
		this.nativeGraphics,
		gpMetafile,
		destPoint,
		srcRect,
		srcUnit,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafileI4(
	metafile *Metafile,
	destPoint *Point,
	srcRect *Rect,
	srcUnit Unit,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileSrcRectDestPointI(
		this.nativeGraphics,
		gpMetafile,
		destPoint,
		srcRect,
		srcUnit,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafile5(
	metafile *Metafile,
	destRect *RectF,
	srcRect *RectF,
	srcUnit Unit,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileSrcRectDestRect(
		this.nativeGraphics,
		gpMetafile,
		destRect,
		srcRect,
		srcUnit,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafileI5(
	metafile *Metafile,
	destRect *Rect,
	srcRect *Rect,
	srcUnit Unit,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileSrcRectDestRectI(
		this.nativeGraphics,
		gpMetafile,
		destRect,
		srcRect,
		srcUnit,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafile6(
	metafile *Metafile,
	destPoints []PointF,
	srcRect *RectF,
	srcUnit Unit,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileSrcRectDestPoints(
		this.nativeGraphics,
		gpMetafile,
		destPoints,
		srcRect,
		srcUnit,
		callback,
		callbackData,
		nativeImageAttr))
}

func (this *Graphics) EnumerateMetafileI6(
	metafile *Metafile,
	destPoints []Point,
	srcRect *Rect,
	srcUnit Unit,
	callback EnumerateMetafileProc,
	callbackData DATA_PTR,
	imageAttributes *ImageAttributes) Status {
	var gpMetafile *GpMetafile
	if metafile != nil {
		gpMetafile = (*GpMetafile)(metafile.nativeImage)
	}
	var nativeImageAttr *GpImageAttributes
	if imageAttributes != nil {
		nativeImageAttr = imageAttributes.nativeImageAttr
	}

	return this.setStatus(GdipEnumerateMetafileSrcRectDestPointsI(
		this.nativeGraphics,
		gpMetafile,
		destPoints,
		srcRect,
		srcUnit,
		callback,
		callbackData,
		nativeImageAttr))
}

//IN CombineMode combineMode = CombineModeReplace
func (this *Graphics) SetClipGraphics(g *Graphics, combineMode ...CombineMode) Status {
	var combineMode_ CombineMode
	if len(combineMode) > 0 {
		combineMode_ = combineMode[0]
	} else {
		combineMode_ = CombineModeReplace
	}

	return this.setStatus(GdipSetClipGraphics(this.nativeGraphics,
		g.nativeGraphics, combineMode_))
}

func (this *Graphics) SetClipRect(rect *RectF, combineMode ...CombineMode) Status {
	var combineMode_ CombineMode
	if len(combineMode) > 0 {
		combineMode_ = combineMode[0]
	} else {
		combineMode_ = CombineModeReplace
	}

	return this.setStatus(GdipSetClipRect(this.nativeGraphics,
		rect.X, rect.Y, rect.Width, rect.Height, combineMode_))
}

func (this *Graphics) SetClipRectI(rect *Rect, combineMode ...CombineMode) Status {
	var combineMode_ CombineMode
	if len(combineMode) > 0 {
		combineMode_ = combineMode[0]
	} else {
		combineMode_ = CombineModeReplace
	}

	return this.setStatus(GdipSetClipRectI(this.nativeGraphics,
		rect.X, rect.Y, rect.Width, rect.Height, combineMode_))
}

func (this *Graphics) SetClipPath(path *GraphicsPath, combineMode ...CombineMode) Status {
	var combineMode_ CombineMode
	if len(combineMode) > 0 {
		combineMode_ = combineMode[0]
	} else {
		combineMode_ = CombineModeReplace
	}

	return this.setStatus(GdipSetClipPath(this.nativeGraphics,
		path.nativePath, combineMode_))
}

func (this *Graphics) SetClipRegion(region *Region, combineMode ...CombineMode) Status {
	var combineMode_ CombineMode
	if len(combineMode) > 0 {
		combineMode_ = combineMode[0]
	} else {
		combineMode_ = CombineModeReplace
	}

	return this.setStatus(GdipSetClipRegion(this.nativeGraphics,
		region.nativeRegion, combineMode_))
}

// This is different than the other SetClip methods because it assumes
// that the HRGN is already in device units, so it doesn't transform
// the coordinates in the HRGN.

func (this *Graphics) SetClipHrgn(hRgn HRGN, combineMode ...CombineMode) Status {
	var combineMode_ CombineMode
	if len(combineMode) > 0 {
		combineMode_ = combineMode[0]
	} else {
		combineMode_ = CombineModeReplace
	}

	return this.setStatus(GdipSetClipHrgn(this.nativeGraphics,
		hRgn, combineMode_))
}

func (this *Graphics) IntersectClipRect(rect *RectF) Status {
	return this.setStatus(GdipSetClipRect(this.nativeGraphics,
		rect.X, rect.Y, rect.Width, rect.Height, CombineModeIntersect))
}

func (this *Graphics) IntersectClipRectI(rect *Rect) Status {
	return this.setStatus(GdipSetClipRectI(this.nativeGraphics,
		rect.X, rect.Y, rect.Width, rect.Height, CombineModeIntersect))
}

func (this *Graphics) IntersectClipRegion(region *Region) Status {
	return this.setStatus(GdipSetClipRegion(this.nativeGraphics,
		region.nativeRegion, CombineModeIntersect))
}

func (this *Graphics) ExcludeClipRect(rect *RectF) Status {
	return this.setStatus(GdipSetClipRect(this.nativeGraphics,
		rect.X, rect.Y, rect.Width, rect.Height, CombineModeExclude))
}

func (this *Graphics) ExcludeClipRectI(rect *Rect) Status {
	return this.setStatus(GdipSetClipRectI(this.nativeGraphics,
		rect.X, rect.Y, rect.Width, rect.Height, CombineModeExclude))
}

func (this *Graphics) ExcludeClipRegion(region *Region) Status {
	return this.setStatus(GdipSetClipRegion(this.nativeGraphics,
		region.nativeRegion, CombineModeExclude))
}

func (this *Graphics) ResetClip() Status {
	return this.setStatus(GdipResetClip(this.nativeGraphics))
}

func (this *Graphics) TranslateClip(dx, dy REAL) Status {
	return this.setStatus(GdipTranslateClip(this.nativeGraphics, dx, dy))
}

func (this *Graphics) TranslateClipI(dx, dy INT) Status {
	return this.setStatus(GdipTranslateClipI(this.nativeGraphics, dx, dy))
}

func (this *Graphics) GetClip(region *Region) (status Status) {
	return this.setStatus(GdipGetClip(this.nativeGraphics, region.nativeRegion))
}

func (this *Graphics) GetClipBounds() (rect *RectF, status Status) {
	rect = &RectF{}
	this.setStatus(GdipGetClipBounds(this.nativeGraphics, rect))
	status = this.LastResult
	return
}

func (this *Graphics) GetClipBoundsI() (rect *Rect, status Status) {
	rect = &Rect{}
	this.setStatus(GdipGetClipBoundsI(this.nativeGraphics, rect))
	status = this.LastResult
	return
}

func (this *Graphics) IsClipEmpty() (booln BOOL, status Status) {
	this.setStatus(GdipIsClipEmpty(this.nativeGraphics, &booln))
	return booln, this.LastResult
}

func (this *Graphics) GetVisibleClipBounds() (rect *RectF, status Status) {
	rect = &RectF{}
	this.setStatus(GdipGetVisibleClipBounds(this.nativeGraphics, rect))
	status = this.LastResult
	return
}

func (this *Graphics) GetVisibleClipBoundsI() (rect *Rect, status Status) {
	rect = &Rect{}
	this.setStatus(GdipGetVisibleClipBoundsI(this.nativeGraphics, rect))
	status = this.LastResult
	return
}

func (this *Graphics) IsVisibleClipEmpty() (booln BOOL) {
	this.setStatus(GdipIsVisibleClipEmpty(this.nativeGraphics, &booln))
	return
}

func (this *Graphics) IsVisiblePointI(x, y INT) (booln BOOL) {
	this.setStatus(GdipIsVisiblePointI(this.nativeGraphics,
		x, y, &booln))
	return
}

func (this *Graphics) IsVisiblePointI2(point *Point) (booln BOOL) {
	return this.IsVisiblePointI(point.X, point.Y)
}

func (this *Graphics) IsVisibleRectI(x, y, width,
	height INT) (booln BOOL) {
	this.setStatus(GdipIsVisibleRectI(this.nativeGraphics,
		x, y, width, height, &booln))
	return booln
}

func (this *Graphics) IsVisibleRectI2(rect *Rect) (booln BOOL) {
	return this.IsVisibleRectI(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) IsVisiblePoint(x, y REAL) (booln BOOL) {
	this.setStatus(GdipIsVisiblePoint(this.nativeGraphics,
		x, y, &booln))
	return
}

func (this *Graphics) IsVisiblePoint2(point *PointF) (booln BOOL) {
	return this.IsVisiblePoint(point.X, point.Y)
}

func (this *Graphics) IsVisibleRect(x, y, width,
	height REAL) (booln BOOL) {
	this.setStatus(GdipIsVisibleRect(this.nativeGraphics,
		x, y, width, height, &booln))
	return
}

func (this *Graphics) IsVisibleRect2(rect *RectF) (booln BOOL) {
	return this.IsVisibleRect(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *Graphics) Save() (gstate GraphicsState) {
	this.setStatus(GdipSaveGraphics(this.nativeGraphics, &gstate))
	return
}

func (this *Graphics) Restore(gstate GraphicsState) Status {
	return this.setStatus(GdipRestoreGraphics(this.nativeGraphics, gstate))
}

func (this *Graphics) BeginContainer2(dstrect, srcrect *RectF,
	unit Unit) (state GraphicsContainer) {
	this.setStatus(GdipBeginContainer(this.nativeGraphics, dstrect,
		srcrect, unit, &state))
	return
}

func (this *Graphics) BeginContainer2I(dstrect, srcrect *Rect,
	unit Unit) (state GraphicsContainer) {
	this.setStatus(GdipBeginContainerI(this.nativeGraphics, dstrect,
		srcrect, unit, &state))
	return
}

func (this *Graphics) BeginContainer() (state GraphicsContainer) {
	this.setStatus(GdipBeginContainer2(this.nativeGraphics, &state))
	return
}

func (this *Graphics) EndContainer(state GraphicsContainer) Status {
	return this.setStatus(GdipEndContainer(this.nativeGraphics, state))
}

// Only valid when recording metafiles.
func (this *Graphics) AddMetafileComment(data []byte) Status {
	return this.setStatus(GdipComment(this.nativeGraphics, data))
}

func GetHalftonePalette() (HPALETTE, error) {
	return GdipCreateHalftonePalette()
}
