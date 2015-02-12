package gdiplus

import (
	"errors"
	. "github.com/trygo/winapi"
)

type GraphicsPath struct {
	GdiplusBase
	nativePath *GpPath
}

func NewGraphicsPath(fillMode ...FillMode) (*GraphicsPath, error) {
	var fillMode_ FillMode
	if len(fillMode) > 0 {
		fillMode_ = fillMode[0]
	} else {
		fillMode_ = FillModeAlternate
	}
	var nativePath *GpPath
	lastResult, err := GdipCreatePath(fillMode_, &nativePath)
	return &GraphicsPath{nativePath: nativePath, GdiplusBase: GdiplusBase{LastResult: Status(lastResult), LastError: err}}, err
}

func NewGraphicsPath2(points []PointF, types []byte,
	fillMode ...FillMode) (*GraphicsPath, error) {
	path := &GraphicsPath{}
	if len(points) != len(types) || len(types) == 0 {
		path.setStatus(InvalidParameter, errors.New(StatusText[InvalidParameter]))
		return path, path.LastError
	}

	var fillMode_ FillMode
	if len(fillMode) > 0 {
		fillMode_ = fillMode[0]
	} else {
		fillMode_ = FillModeAlternate
	}

	path.setStatus(GdipCreatePath2(points,
		types,
		fillMode_,
		&path.nativePath))

	return path, path.LastError
}

func NewGraphicsPath2I(points []Point, types []byte,
	fillMode ...FillMode) (*GraphicsPath, error) {
	path := &GraphicsPath{}
	if len(points) != len(types) || len(types) == 0 {
		path.setStatus(InvalidParameter, errors.New(StatusText[InvalidParameter]))
		return path, path.LastError
	}

	var fillMode_ FillMode
	if len(fillMode) > 0 {
		fillMode_ = fillMode[0]
	} else {
		fillMode_ = FillModeAlternate
	}

	path.setStatus(GdipCreatePath2I(points,
		types,
		fillMode_,
		&path.nativePath))

	return path, path.LastError
}

func (this *GraphicsPath) Release() {
	if this.nativePath != nil {
		GdipDeletePath(this.nativePath)
	}
}

func (this *GraphicsPath) Clone() *GraphicsPath {
	path := &GraphicsPath{}
	this.setStatus(GdipClonePath(this.nativePath, &path.nativePath))
	return path
}

func newGraphicsPath(path *GraphicsPath) (*GraphicsPath, error) {
	newpath := &GraphicsPath{}
	newpath.setStatus(GdipClonePath(path.nativePath, &newpath.nativePath))
	return newpath, newpath.LastError
}

func newGraphicsPath2(nativePath *GpPath) *GraphicsPath {
	return &GraphicsPath{GdiplusBase: GdiplusBase{LastResult: Ok}, nativePath: nativePath}
}

// Reset the path object to empty (and fill mode to FillModeAlternate)
func (this *GraphicsPath) Reset() Status {
	return this.setStatus(GdipResetPath(this.nativePath))
}

func (this *GraphicsPath) GetFillMode() FillMode {
	fillmode := FillModeAlternate
	this.setStatus(GdipGetPathFillMode(this.nativePath, (*GpFillMode)(&fillmode)))
	return fillmode
}

func (this *GraphicsPath) SetFillMode(fillmode FillMode) Status {
	return this.setStatus(GdipSetPathFillMode(this.nativePath,
		GpFillMode(fillmode)))
}

func (this *GraphicsPath) GetPointCount() (count INT) {
	this.setStatus(GdipGetPointCount(this.nativePath, &count))
	return
}

func (this *GraphicsPath) GetPathData() (pathData *PathData, status Status) {
	count := this.GetPointCount()
	if count <= 0 {
		pathData = NewPathData(0)
		status = this.setStatus(Ok, nil)
		return
	}
	pathData = NewPathData(count)
	status = this.setStatus(GdipGetPathData(this.nativePath, pathData))
	return
}

func (this *GraphicsPath) StartFigure() Status {
	return this.setStatus(GdipStartPathFigure(this.nativePath))
}

func (this *GraphicsPath) CloseFigure() Status {
	return this.setStatus(GdipClosePathFigure(this.nativePath))
}

func (this *GraphicsPath) CloseAllFigures() Status {
	return this.setStatus(GdipClosePathFigures(this.nativePath))
}

func (this *GraphicsPath) SetMarker() Status {
	return this.setStatus(GdipSetPathMarker(this.nativePath))
}

func (this *GraphicsPath) ClearMarkers() Status {
	return this.setStatus(GdipClearPathMarkers(this.nativePath))
}

func (this *GraphicsPath) Reverse() Status {
	return this.setStatus(GdipReversePath(this.nativePath))
}

func (this *GraphicsPath) GetLastPoint() (lastPoint *PointF, status Status) {
	lastPoint = &PointF{}
	this.setStatus(GdipGetPathLastPoint(this.nativePath, lastPoint))
	status = this.LastResult
	return
}

func (this *GraphicsPath) AddLine(x1, y1, x2, y2 REAL) Status {
	return this.setStatus(GdipAddPathLine(this.nativePath, x1, y1, x2, y2))
}

func (this *GraphicsPath) AddLine2(pt1, pt2 *PointF) Status {
	return this.AddLine(pt1.X, pt1.Y, pt2.X, pt2.Y)
}

func (this *GraphicsPath) AddLines(points []PointF) Status {
	return this.setStatus(GdipAddPathLine2(this.nativePath, points))
}

func (this *GraphicsPath) AddLineI(x1, y1, x2, y2 INT) Status {
	return this.setStatus(GdipAddPathLineI(this.nativePath, x1, y1, x2, y2))
}

func (this *GraphicsPath) AddLineI2(pt1, pt2 *Point) Status {
	return this.AddLineI(pt1.X, pt1.Y, pt2.X, pt2.Y)
}

func (this *GraphicsPath) AddLinesI(points []Point) Status {
	return this.setStatus(GdipAddPathLine2I(this.nativePath, points))
}

func (this *GraphicsPath) AddArc(x, y, width, height, startAngle,
	sweepAngle REAL) Status {
	return this.setStatus(GdipAddPathArc(this.nativePath, x, y, width,
		height, startAngle, sweepAngle))
}

func (this *GraphicsPath) AddArc2(rect *RectF,
	startAngle, sweepAngle REAL) Status {
	return this.AddArc(rect.X, rect.Y, rect.Width, rect.Height,
		startAngle, sweepAngle)
}

func (this *GraphicsPath) AddArcI(x, y, width, height INT, startAngle,
	sweepAngle REAL) Status {
	return this.setStatus(GdipAddPathArcI(this.nativePath, x, y, width,
		height, startAngle, sweepAngle))
}

func (this *GraphicsPath) AddArcI2(rect *Rect,
	startAngle, sweepAngle REAL) Status {
	return this.AddArcI(rect.X, rect.Y, rect.Width, rect.Height,
		startAngle, sweepAngle)
}

func (this *GraphicsPath) AddBezier(x1, y1, x2, y2, x3, y3, x4, y4 REAL) Status {
	return this.setStatus(GdipAddPathBezier(this.nativePath, x1, y1, x2, y2, x3, y3, x4, y4))
}

func (this *GraphicsPath) AddBezier2(pt1, pt2, pt3, pt4 *PointF) Status {
	return this.AddBezier(pt1.X, pt1.Y, pt2.X, pt2.Y, pt3.X, pt3.Y, pt4.X, pt4.Y)
}

func (this *GraphicsPath) AddBeziers(points []PointF) Status {
	return this.setStatus(GdipAddPathBeziers(this.nativePath, points))
}

func (this *GraphicsPath) AddBezierI(x1, y1, x2, y2, x3, y3, x4, y4 INT) Status {
	return this.setStatus(GdipAddPathBezierI(this.nativePath, x1, y1, x2, y2, x3, y3, x4, y4))
}

func (this *GraphicsPath) AddBezierI2(pt1, pt2, pt3, pt4 *Point) Status {
	return this.AddBezierI(pt1.X, pt1.Y, pt2.X, pt2.Y, pt3.X, pt3.Y, pt4.X, pt4.Y)
}

func (this *GraphicsPath) AddBeziersI(points []Point) Status {
	return this.setStatus(GdipAddPathBeziersI(this.nativePath, points))
}

func (this *GraphicsPath) AddCurve(points []PointF) Status {
	return this.setStatus(GdipAddPathCurve(this.nativePath, points))
}

func (this *GraphicsPath) AddCurve2(points []PointF, tension REAL) Status {
	return this.setStatus(GdipAddPathCurve2(this.nativePath, points, tension))
}

func (this *GraphicsPath) AddCurve3(points []PointF,
	offset INT, numberOfSegments INT, tension REAL) Status {
	return this.setStatus(GdipAddPathCurve3(this.nativePath,
		points, offset, numberOfSegments, tension))
}

func (this *GraphicsPath) AddCurveI(points []Point) Status {
	return this.setStatus(GdipAddPathCurveI(this.nativePath, points))
}

func (this *GraphicsPath) AddCurveI2(points []Point, tension REAL) Status {
	return this.setStatus(GdipAddPathCurve2I(this.nativePath, points, tension))
}

func (this *GraphicsPath) AddCurveI3(points []Point,
	offset INT, numberOfSegments INT, tension REAL) Status {
	return this.setStatus(GdipAddPathCurve3I(this.nativePath,
		points, offset, numberOfSegments, tension))
}

func (this *GraphicsPath) AddClosedCurve(points []PointF) Status {
	return this.setStatus(GdipAddPathClosedCurve(this.nativePath, points))
}

func (this *GraphicsPath) AddClosedCurve2(points []PointF, tension REAL) Status {
	return this.setStatus(GdipAddPathClosedCurve2(this.nativePath, points, tension))
}

func (this *GraphicsPath) AddClosedCurveI(points []Point) Status {
	return this.setStatus(GdipAddPathClosedCurveI(this.nativePath, points))
}

func (this *GraphicsPath) AddClosedCurveI2(points []Point, tension REAL) Status {
	return this.setStatus(GdipAddPathClosedCurve2I(this.nativePath, points, tension))
}

func (this *GraphicsPath) AddRectangle(rect *RectF) Status {
	return this.setStatus(GdipAddPathRectangle(this.nativePath,
		rect.X,
		rect.Y,
		rect.Width,
		rect.Height))
}

func (this *GraphicsPath) AddRectangles(rects []RectF) Status {
	return this.setStatus(GdipAddPathRectangles(this.nativePath,
		rects))
}

func (this *GraphicsPath) AddRectangleI(rect *Rect) Status {
	return this.setStatus(GdipAddPathRectangleI(this.nativePath,
		rect.X,
		rect.Y,
		rect.Width,
		rect.Height))
}

func (this *GraphicsPath) AddRectanglesI(rects []Rect) Status {
	return this.setStatus(GdipAddPathRectanglesI(this.nativePath,
		rects))
}

func (this *GraphicsPath) AddEllipse(x, y, width, height REAL) Status {
	return this.setStatus(GdipAddPathEllipse(this.nativePath,
		x, y, width, height))
}

func (this *GraphicsPath) AddEllipse2(rect *RectF) Status {
	return this.AddEllipse(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *GraphicsPath) AddEllipseI(x, y, width, height INT) Status {
	return this.setStatus(GdipAddPathEllipseI(this.nativePath,
		x, y, width, height))
}

func (this *GraphicsPath) AddEllipseI2(rect *Rect) Status {
	return this.AddEllipseI(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *GraphicsPath) AddPie(x, y, width, height,
	startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipAddPathPie(this.nativePath, x, y, width,
		height, startAngle, sweepAngle))
}

func (this *GraphicsPath) AddPie2(rect *RectF,
	startAngle, sweepAngle REAL) Status {
	return this.AddPie(rect.X, rect.Y, rect.Width, rect.Height, startAngle,
		sweepAngle)
}

func (this *GraphicsPath) AddPieI(x, y, width, height INT,
	startAngle, sweepAngle REAL) Status {
	return this.setStatus(GdipAddPathPieI(this.nativePath, x, y, width,
		height, startAngle, sweepAngle))
}

func (this *GraphicsPath) AddPieI2(rect *Rect,
	startAngle, sweepAngle REAL) Status {
	return this.AddPieI(rect.X, rect.Y, rect.Width, rect.Height, startAngle,
		sweepAngle)
}

func (this *GraphicsPath) AddPolygon(points []PointF) Status {
	return this.setStatus(GdipAddPathPolygon(this.nativePath, points))
}

func (this *GraphicsPath) AddPolygonI(points []Point) Status {
	return this.setStatus(GdipAddPathPolygonI(this.nativePath, points))
}

func (this *GraphicsPath) AddPath(addingPath *GraphicsPath, connect BOOL) Status {
	var nativePath2 *GpPath
	if addingPath != nil {
		nativePath2 = addingPath.nativePath
	}
	return this.setStatus(GdipAddPathPath(this.nativePath, nativePath2, connect))
}

func (this *GraphicsPath) AddString(text string, family *FontFamily, style FontStyle, emSize REAL, layoutRect *RectF, stringFormat *StringFormat) Status {
	var nativeFamily *GpFontFamily
	if family != nil {
		nativeFamily = family.nativeFamily
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}

	return this.setStatus(GdipAddPathString(
		this.nativePath,
		text,
		nativeFamily,
		style,
		emSize,
		layoutRect, nativeFormat))
}

func (this *GraphicsPath) AddString2(text string, family *FontFamily, style FontStyle,
	emSize REAL, origin *PointF, stringFormat *StringFormat) Status {
	rect := &RectF{origin.X, origin.Y, 0.0, 0.0}
	return this.AddString(text, family, style, emSize, rect, stringFormat)
}

func (this *GraphicsPath) AddStringI(text string, family *FontFamily, style FontStyle, emSize REAL, layoutRect *Rect, stringFormat *StringFormat) Status {
	var nativeFamily *GpFontFamily
	if family != nil {
		nativeFamily = family.nativeFamily
	}
	var nativeFormat *GpStringFormat
	if stringFormat != nil {
		nativeFormat = stringFormat.nativeFormat
	}

	return this.setStatus(GdipAddPathStringI(
		this.nativePath,
		text,
		nativeFamily,
		style,
		emSize,
		layoutRect, nativeFormat))
}

func (this *GraphicsPath) AddStringI2(text string, family *FontFamily, style FontStyle,
	emSize REAL, origin *Point, stringFormat *StringFormat) Status {
	rect := &Rect{origin.X, origin.Y, 0.0, 0.0}
	return this.AddStringI(text, family, style, emSize, rect, stringFormat)
}

func (this *GraphicsPath) Transform(matrix *Matrix) Status {
	if matrix != nil {
		return this.setStatus(GdipTransformPath(this.nativePath,
			matrix.nativeMatrix))
	} else {
		return Ok
	}
}

// The GetBounds rectangle may not be the tightest bounds.
func (this *GraphicsPath) GetBounds(
	matrix *Matrix, pen *Pen) (bounds *RectF, status Status) {
	var nativeMatrix *GpMatrix
	var nativePen *GpPen

	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}

	if pen != nil {
		nativePen = pen.nativePen
	}
	bounds = &RectF{}
	this.setStatus(GdipGetPathWorldBounds(this.nativePath, bounds,
		nativeMatrix, nativePen))
	return bounds, this.LastResult
}

func (this *GraphicsPath) GetBoundsI(
	matrix *Matrix, pen *Pen) (bounds *Rect, status Status) {
	var nativeMatrix *GpMatrix
	var nativePen *GpPen

	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}

	if pen != nil {
		nativePen = pen.nativePen
	}
	bounds = &Rect{}
	this.setStatus(GdipGetPathWorldBoundsI(this.nativePath, bounds,
		nativeMatrix, nativePen))
	return bounds, this.LastResult
}

// Once flattened, the resultant path is made of line segments and
// the original path information is lost.  When matrix is NULL the
// identity matrix is assumed.
func (this *GraphicsPath) Flatten(matrix *Matrix, flatness ...REAL) Status {
	var nativeMatrix *GpMatrix
	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}
	var flatness_ REAL
	if len(flatness) > 0 {
		flatness_ = flatness[0]
	} else {
		flatness_ = FlatnessDefault
	}
	return this.setStatus(GdipFlattenPath(
		this.nativePath,
		nativeMatrix,
		flatness_))
}

func (this *GraphicsPath) Widen(pen *Pen, matrix *Matrix, flatness ...REAL) Status {
	var nativeMatrix *GpMatrix
	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}
	var flatness_ REAL
	if len(flatness) > 0 {
		flatness_ = flatness[0]
	} else {
		flatness_ = FlatnessDefault
	}

	return this.setStatus(GdipWidenPath(
		this.nativePath,
		pen.nativePen,
		nativeMatrix,
		flatness_))
}

func (this *GraphicsPath) Outline(matrix *Matrix, flatness ...REAL) Status {
	var nativeMatrix *GpMatrix
	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}
	var flatness_ REAL
	if len(flatness) > 0 {
		flatness_ = flatness[0]
	} else {
		flatness_ = FlatnessDefault
	}

	return this.setStatus(GdipWindingModeOutline(
		this.nativePath, nativeMatrix, flatness_))
}

//IN const Matrix* matrix = NULL,
//IN WarpMode warpMode = WarpModePerspective,
//IN REAL flatness = FlatnessDefault
func (this *GraphicsPath) Warp(destPoints []PointF,
	srcRect *RectF,
	matrix *Matrix,
	warpMode WarpMode,
	flatness ...REAL) Status {

	var nativeMatrix *GpMatrix
	if matrix != nil {
		nativeMatrix = matrix.nativeMatrix
	}
	var flatness_ REAL
	if len(flatness) > 0 {
		flatness_ = flatness[0]
	} else {
		flatness_ = FlatnessDefault
	}

	return this.setStatus(GdipWarpPath(
		this.nativePath,
		nativeMatrix,
		destPoints,
		srcRect.X,
		srcRect.Y,
		srcRect.Width,
		srcRect.Height,
		warpMode,
		flatness_))
}

func (this *GraphicsPath) GetPathTypes(types []BYTE) Status {
	return this.setStatus(GdipGetPathTypes(this.nativePath, types))
}

func (this *GraphicsPath) GetPathPoints(points []PointF) Status {
	return this.setStatus(GdipGetPathPoints(this.nativePath, points))
}

func (this *GraphicsPath) GetPathPointsI(points []Point) Status {
	return this.setStatus(GdipGetPathPointsI(this.nativePath, points))
}

func (this *GraphicsPath) IsVisible(x, y REAL, g *Graphics) (booln BOOL) {
	var nativeGraphics *GpGraphics

	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	this.setStatus(GdipIsVisiblePathPoint(this.nativePath,
		x, y, nativeGraphics, &booln))

	return
}

func (this *GraphicsPath) IsVisible2(point *PointF, g *Graphics) BOOL {
	return this.IsVisible(point.X, point.Y, g)
}

func (this *GraphicsPath) IsVisibleI(x, y INT, g *Graphics) (booln BOOL) {
	var nativeGraphics *GpGraphics

	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	this.setStatus(GdipIsVisiblePathPointI(this.nativePath,
		x, y, nativeGraphics, &booln))

	return
}

func (this *GraphicsPath) IsVisibleI2(point *Point, g *Graphics) BOOL {
	return this.IsVisibleI(point.X, point.Y, g)
}

func (this *GraphicsPath) IsOutlineVisible(
	x, y REAL, pen *Pen, g *Graphics) (booln BOOL) {

	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	var nativePen *GpPen
	if pen != nil {
		nativePen = pen.nativePen
	}

	this.setStatus(GdipIsOutlineVisiblePathPoint(this.nativePath,
		x, y, nativePen, nativeGraphics, &booln))

	return
}

func (this *GraphicsPath) IsOutlineVisible2(point *PointF, pen *Pen, g *Graphics) BOOL {
	return this.IsOutlineVisible(point.X, point.Y, pen, g)
}

func (this *GraphicsPath) IsOutlineVisibleI(
	x, y INT, pen *Pen, g *Graphics) (booln BOOL) {

	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}
	var nativePen *GpPen
	if pen != nil {
		nativePen = pen.nativePen
	}

	this.setStatus(GdipIsOutlineVisiblePathPointI(this.nativePath,
		x, y, nativePen, nativeGraphics, &booln))

	return
}

func (this *GraphicsPath) IsOutlineVisibleI2(point *Point, pen *Pen, g *Graphics) BOOL {
	return this.IsOutlineVisibleI(point.X, point.Y, pen, g)
}

func (this *GraphicsPath) SetNativePath(nativePath *GpPath) {
	this.nativePath = nativePath
}

type GraphicsPathIterator struct {
	GdiplusBase
	nativeIterator *GpPathIterator
}

func NewGraphicsPathIterator(path *GraphicsPath) (*GraphicsPathIterator, error) {
	var nativePath *GpPath
	if path != nil {
		nativePath = path.nativePath
	}
	iter := &GraphicsPathIterator{}
	iter.setStatus(GdipCreatePathIter(&iter.nativeIterator, nativePath))
	return iter, iter.LastError
}

func (this *GraphicsPathIterator) Release() {
	if this.nativeIterator != nil {
		GdipDeletePathIter(this.nativeIterator)
	}
}

func (this *GraphicsPathIterator) NextSubpath() (startIndex, endIndex INT, isClosed BOOL, resultCount INT) {
	this.setStatus(GdipPathIterNextSubpath(this.nativeIterator,
		&resultCount, &startIndex, &endIndex, &isClosed))
	return
}

func (this *GraphicsPathIterator) NextSubpath2(outpath *GraphicsPath) (isClosed BOOL, resultCount INT) {
	var nativePath *GpPath
	if outpath != nil {
		nativePath = outpath.nativePath
	}
	this.setStatus(GdipPathIterNextSubpathPath(this.nativeIterator,
		&resultCount, nativePath, &isClosed))

	return
}

func (this *GraphicsPathIterator) NextPathType() (pathType BYTE, startIndex, endIndex, resultCount INT) {
	this.setStatus(GdipPathIterNextPathType(this.nativeIterator,
		&resultCount, &pathType, &startIndex, &endIndex))
	return
}

func (this *GraphicsPathIterator) NextMarker() (startIndex, endIndex, resultCount INT) {
	this.setStatus(GdipPathIterNextMarker(this.nativeIterator,
		&resultCount, &startIndex, &endIndex))
	return
}

func (this *GraphicsPathIterator) NextMarker2(path *GraphicsPath) (resultCount INT) {
	var nativePath *GpPath
	if path != nil {
		nativePath = path.nativePath
	}
	this.setStatus(GdipPathIterNextMarkerPath(this.nativeIterator,
		&resultCount, nativePath))
	return
}

func (this *GraphicsPathIterator) GetCount() (resultCount INT) {
	this.setStatus(GdipPathIterGetCount(this.nativeIterator,
		&resultCount))
	return
}

func (this *GraphicsPathIterator) GetSubpathCount() (resultCount INT) {
	this.setStatus(GdipPathIterGetSubpathCount(this.nativeIterator,
		&resultCount))
	return
}

func (this *GraphicsPathIterator) HasCurve() (hasCurve BOOL) {
	this.setStatus(GdipPathIterHasCurve(this.nativeIterator, &hasCurve))
	return
}

func (this *GraphicsPathIterator) Rewind() {
	this.setStatus(GdipPathIterRewind(this.nativeIterator))
}

func (this *GraphicsPathIterator) Enumerate(count INT) (points []PointF, types []BYTE, resultCount INT) {
	points = make([]PointF, count)
	types = make([]BYTE, count)
	this.setStatus(GdipPathIterEnumerate(this.nativeIterator,
		&resultCount, points, types, count))
	return
}

func (this *GraphicsPathIterator) CopyData(startIndex, endIndex INT) (points []PointF, types []BYTE, resultCount INT) {
	count := endIndex - startIndex + 1
	points = make([]PointF, count)
	types = make([]BYTE, count)
	this.setStatus(GdipPathIterCopyData(this.nativeIterator,
		&resultCount, points, types, startIndex, endIndex))
	return
}

type PathGradientBrush struct {
	Brush
}

func NewPathGradientBrush(points []PointF, wrapMode ...WrapMode) (*PathGradientBrush, error) {
	var wrapMode_ WrapMode
	if len(wrapMode) > 0 {
		wrapMode_ = wrapMode[0]
	} else {
		wrapMode_ = WrapModeClamp
	}

	brush := &PathGradientBrush{}
	brush.setStatus(GdipCreatePathGradient(
		points, GpWrapMode(wrapMode_), &brush.nativeBrush))

	return brush, brush.LastError
}

func NewPathGradientBrushI(points []Point, wrapMode ...WrapMode) (*PathGradientBrush, error) {
	var wrapMode_ WrapMode
	if len(wrapMode) > 0 {
		wrapMode_ = wrapMode[0]
	} else {
		wrapMode_ = WrapModeClamp
	}

	brush := &PathGradientBrush{}
	brush.setStatus(GdipCreatePathGradientI(
		points, GpWrapMode(wrapMode_), &brush.nativeBrush))

	return brush, brush.LastError
}

func NewPathGradientBrush2(path *GraphicsPath) (*PathGradientBrush, error) {

	brush := &PathGradientBrush{}
	brush.setStatus(GdipCreatePathGradientFromPath(
		path.nativePath, &brush.nativeBrush))

	return brush, brush.LastError
}

func (this *PathGradientBrush) GetCenterColor() (color Color, status Status) {
	status = this.setStatus(GdipGetPathGradientCenterColor(
		this.nativeBrush, &color.Argb))
	return
}

func (this *PathGradientBrush) SetCenterColor(color Color) Status {
	return this.setStatus(GdipSetPathGradientCenterColor(
		this.nativeBrush,
		color.GetValue()))
}

func (this *PathGradientBrush) GetPointCount() (count INT) {
	this.setStatus(GdipGetPathGradientPointCount(
		this.nativeBrush, &count))
	return
}

func (this *PathGradientBrush) GetSurroundColorCount() (count INT) {
	this.setStatus(GdipGetPathGradientSurroundColorCount(
		this.nativeBrush, &count))

	return
}

func (this *PathGradientBrush) GetSurroundColors() (colors []Color, status Status) {
	var count INT
	this.setStatus(GdipGetPathGradientSurroundColorCount(
		this.nativeBrush, &count))

	if this.LastResult != Ok {
		colors = make([]Color, 0)
		status = this.LastResult
		return
	}

	if count <= 0 {
		colors = make([]Color, 0)
		status = this.setStatus((InsufficientBuffer), nil)
		return
	}

	argbs := make([]ARGB, count)
	status = this.setStatus(GdipGetPathGradientSurroundColorsWithCount(this.nativeBrush, argbs, &count))

	if this.LastResult == Ok {
		colors = make([]Color, count)
		for i := INT(0); i < count; i++ {
			colors[i].SetValue(argbs[i])
		}
	} else {
		colors = make([]Color, 0)
	}
	return
}

func (this *PathGradientBrush) SetSurroundColors(colors []Color) (count INT, status Status) {
	if colors == nil {
		status = this.setStatus(InvalidParameter, nil)
		return
	}

	incount := INT(len(colors))
	count1 := this.GetPointCount()

	if (incount > count1) || (count1 <= 0) {
		status = this.setStatus(InvalidParameter, nil)
		return
	}
	count1 = incount
	argbs := make([]ARGB, count1)
	for i := INT(0); i < count1; i++ {
		argbs[i] = colors[i].GetValue()
	}
	status = this.setStatus(GdipSetPathGradientSurroundColorsWithCount(this.nativeBrush, argbs, &count1))

	if this.LastResult == Ok {
		count = count1
	}
	return
}

func (this *PathGradientBrush) GetGraphicsPath(path *GraphicsPath) Status {
	if path == nil {
		return this.setStatus(InvalidParameter, nil)
	}
	return this.setStatus(GdipGetPathGradientPath(
		this.nativeBrush, path.nativePath))
}

func (this *PathGradientBrush) SetGraphicsPath(path *GraphicsPath) Status {
	if path == nil {
		return this.setStatus(InvalidParameter, nil)
	}
	return this.setStatus(GdipSetPathGradientPath(
		this.nativeBrush, path.nativePath))
}

func (this *PathGradientBrush) GetCenterPoint() (point *PointF, status Status) {
	point = &PointF{}
	status = this.setStatus(GdipGetPathGradientCenterPoint(
		this.nativeBrush, point))
	return
}

func (this *PathGradientBrush) GetCenterPointI() (point *Point, status Status) {
	point = &Point{}
	status = this.setStatus(GdipGetPathGradientCenterPointI(
		this.nativeBrush, point))
	return
}

func (this *PathGradientBrush) SetCenterPoint(point *PointF) Status {
	return this.setStatus(GdipSetPathGradientCenterPoint(
		this.nativeBrush, point))
}

func (this *PathGradientBrush) SetCenterPointI(point *Point) Status {
	return this.setStatus(GdipSetPathGradientCenterPointI(
		this.nativeBrush, point))
}

func (this *PathGradientBrush) GetRectangle() (rect *RectF, status Status) {
	rect = &RectF{}
	status = this.setStatus(GdipGetPathGradientRect(
		this.nativeBrush, rect))
	return
}

func (this *PathGradientBrush) GetRectangleI() (rect *Rect, status Status) {
	rect = &Rect{}
	status = this.setStatus(GdipGetPathGradientRectI(
		this.nativeBrush, rect))
	return
}

func (this *PathGradientBrush) SetGammaCorrection(useGammaCorrection BOOL) Status {
	return this.setStatus(GdipSetPathGradientGammaCorrection(
		this.nativeBrush, useGammaCorrection))
}

func (this *PathGradientBrush) GetGammaCorrection() (useGammaCorrection BOOL) {
	this.setStatus(GdipGetPathGradientGammaCorrection(
		this.nativeBrush, &useGammaCorrection))
	return
}

func (this *PathGradientBrush) GetBlendCount() (count INT) {
	this.setStatus(GdipGetPathGradientBlendCount(
		this.nativeBrush, &count))
	return
}

func (this *PathGradientBrush) GetBlend(count INT) (blendFactors, blendPositions []REAL, status Status) {
	blendFactors = make([]REAL, count)
	blendPositions = make([]REAL, count)
	status = this.setStatus(GdipGetPathGradientBlend(
		this.nativeBrush,
		blendFactors, blendPositions, count))
	return
}

func (this *PathGradientBrush) SetBlend(blendFactors, blendPositions []REAL) Status {
	count := len(blendFactors)
	if count != len(blendPositions) || count <= 0 {
		return this.setStatus(InvalidParameter, nil)
	}
	return this.setStatus(GdipSetPathGradientBlend(
		this.nativeBrush, blendFactors, blendPositions, INT(count)))
}

func (this *PathGradientBrush) GetInterpolationColorCount() (count INT) {
	this.setStatus(GdipGetPathGradientPresetBlendCount(
		this.nativeBrush, &count))
	return
}

func (this *PathGradientBrush) SetInterpolationColors(presetColors []Color,
	blendPositions []REAL) Status {
	count := len(presetColors)
	if count != len(blendPositions) || count <= 0 {
		return this.setStatus(InvalidParameter, nil)
	}

	argbs := make([]ARGB, count)
	for i := int(0); i < count; i++ {
		argbs[i] = presetColors[i].GetValue()
	}
	return this.setStatus(GdipSetPathGradientPresetBlend(
		this.nativeBrush,
		argbs,
		blendPositions,
		INT(count)))
}

func (this *PathGradientBrush) GetInterpolationColors(count INT) (presetColors []Color, blendPositions []REAL, status Status) {
	if count <= 0 {
		status = this.setStatus(InvalidParameter, nil)
		return
	}

	argbs := make([]ARGB, count)
	blendPositions = make([]REAL, count)
	presetColors = make([]Color, count)
	status = this.setStatus(GdipGetPathGradientPresetBlend(
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
func (this *PathGradientBrush) SetBlendBellShape(focus REAL, scale REAL) Status {
	return this.setStatus(GdipSetPathGradientSigmaBlend(
		this.nativeBrush, focus, scale))
}

//scale = 1.0
func (this *PathGradientBrush) SetBlendTriangularShape(focus REAL, scale REAL) Status {
	return this.setStatus(GdipSetPathGradientLinearBlend(
		this.nativeBrush, focus, scale))
}

func (this *PathGradientBrush) GetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipGetPathGradientTransform(
		this.nativeBrush,
		matrix.nativeMatrix))
}

func (this *PathGradientBrush) SetTransform(matrix *Matrix) Status {
	return this.setStatus(GdipSetPathGradientTransform(
		this.nativeBrush,
		matrix.nativeMatrix))
}

func (this *PathGradientBrush) ResetTransform() Status {
	return this.setStatus(GdipResetPathGradientTransform(
		this.nativeBrush))
}

//order = MatrixOrderPrepend
func (this *PathGradientBrush) MultiplyTransform(matrix *Matrix, order MatrixOrder) Status {
	return this.setStatus(GdipMultiplyPathGradientTransform(
		this.nativeBrush,
		matrix.nativeMatrix,
		order))
}

//order = MatrixOrderPrepend
func (this *PathGradientBrush) TranslateTransform(dx, dy REAL, order MatrixOrder) Status {
	return this.setStatus(GdipTranslatePathGradientTransform(
		this.nativeBrush,
		dx, dy, order))
}

//order = MatrixOrderPrepend
func (this *PathGradientBrush) ScaleTransform(sx, sy REAL, order MatrixOrder) Status {
	return this.setStatus(GdipScalePathGradientTransform(
		this.nativeBrush,
		sx, sy, order))
}

//order = MatrixOrderPrepend
func (this *PathGradientBrush) RotateTransform(angle REAL, order MatrixOrder) Status {
	return this.setStatus(GdipRotatePathGradientTransform(
		this.nativeBrush, angle, order))
}

func (this *PathGradientBrush) GetFocusScales() (xScale, yScale REAL, status Status) {
	status = this.setStatus(GdipGetPathGradientFocusScales(
		this.nativeBrush, &xScale, &yScale))
	return
}

func (this *PathGradientBrush) SetFocusScales(xScale, yScale REAL) Status {
	return this.setStatus(GdipSetPathGradientFocusScales(
		this.nativeBrush, xScale, yScale))
}

func (this *PathGradientBrush) GetWrapMode() (wrapMode WrapMode) {
	this.setStatus(GdipGetPathGradientWrapMode(
		this.nativeBrush, &wrapMode))
	return
}

func (this *PathGradientBrush) SetWrapMode(wrapMode WrapMode) Status {
	return this.setStatus(GdipSetPathGradientWrapMode(
		this.nativeBrush, wrapMode))
}
