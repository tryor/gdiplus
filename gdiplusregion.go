package gdiplus

import (
	. "github.com/trygo/winapi"
)

type Region struct {
	GdiplusBase
	nativeRegion *GpRegion
}

func NewRegion() (*Region, error) {
	region := &Region{}
	var nativeRegion *GpRegion
	region.setStatus(GdipCreateRegion(&nativeRegion))
	region.nativeRegion = nativeRegion
	return region, region.LastError
}

func (this *Region) Close() {
	if this.nativeRegion != nil {
		GdipDeleteRegion(this.nativeRegion)
	}
}

func NewRegionRect(rect *RectF) (*Region, error) {
	region := &Region{}
	var nativeRegion *GpRegion
	region.setStatus(GdipCreateRegionRect((*GpRectF)(rect), &nativeRegion))
	region.nativeRegion = nativeRegion
	return region, region.LastError
}

func NewRegionRectI(rect *Rect) (*Region, error) {
	region := &Region{}
	var nativeRegion *GpRegion
	region.setStatus(GdipCreateRegionRectI((*GpRect)(rect), &nativeRegion))
	region.nativeRegion = nativeRegion
	return region, region.LastError
}

func NewRegionPath(path *GraphicsPath) (*Region, error) {
	region := &Region{}
	var nativeRegion *GpRegion
	region.setStatus(GdipCreateRegionPath(path.nativePath, &nativeRegion))
	region.nativeRegion = nativeRegion
	return region, region.LastError
}

func NewRegionRgnData(regionData []BYTE) (*Region, error) {
	region := &Region{}
	var nativeRegion *GpRegion
	region.setStatus(GdipCreateRegionRgnData(&regionData[0], INT(len(regionData)), &nativeRegion))
	region.nativeRegion = nativeRegion
	return region, region.LastError
}

func NewRegionHRGN(hRgn HRGN) (*Region, error) {
	region := &Region{}
	var nativeRegion *GpRegion
	region.setStatus(GdipCreateRegionHrgn(hRgn, &nativeRegion))
	region.nativeRegion = nativeRegion
	return region, region.LastError
}

func (this *Region) GetData(bufferSize UINT) (buffer []BYTE, s Status) {
	buffer = make([]BYTE, bufferSize)
	var sizeFilled UINT
	this.setStatus(GdipGetRegionData(this.nativeRegion, &buffer[0], bufferSize, &sizeFilled))
	buffer = buffer[0:sizeFilled]
	s = this.LastResult
	return
}

func (this *Region) GetHRGN(g *Graphics) HRGN {
	var hrgn HRGN
	this.setStatus(GdipGetRegionHRgn(this.nativeRegion, g.nativeGraphics, &hrgn))
	return hrgn
}

func (this *Region) Clone() *Region {
	var region *GpRegion
	this.setStatus(GdipCloneRegion(this.nativeRegion, &region))
	return &Region{nativeRegion: region}
}

func (this *Region) MakeInfinite() Status {
	return this.setStatus(GdipSetInfinite(this.nativeRegion))
}

func (this *Region) MakeEmpty() Status {
	return this.setStatus(GdipSetEmpty(this.nativeRegion))
}

func (this *Region) Intersect(rect *RectF) Status {
	return this.setStatus(GdipCombineRegionRect(this.nativeRegion, rect, CombineModeIntersect))
}

func (this *Region) IntersectI(rect *Rect) Status {
	return this.setStatus(GdipCombineRegionRectI(this.nativeRegion, rect, CombineModeIntersect))
}

func (this *Region) IntersectPath(path *GraphicsPath) Status {
	return this.setStatus(GdipCombineRegionPath(this.nativeRegion,
		path.nativePath, CombineModeIntersect))
}

func (this *Region) IntersectRegion(region *Region) Status {
	return this.setStatus(GdipCombineRegionRegion(this.nativeRegion,
		region.nativeRegion, CombineModeIntersect))
}

func (this *Region) Union(rect *RectF) Status {
	return this.setStatus(GdipCombineRegionRect(this.nativeRegion, rect,
		CombineModeUnion))
}

func (this *Region) UnionI(rect *Rect) Status {
	return this.setStatus(GdipCombineRegionRectI(this.nativeRegion, rect,
		CombineModeUnion))
}

func (this *Region) UnionPath(path *GraphicsPath) Status {
	return this.setStatus(GdipCombineRegionPath(this.nativeRegion,
		path.nativePath, CombineModeUnion))
}

func (this *Region) UnionRegion(region *Region) Status {
	return this.setStatus(GdipCombineRegionRegion(this.nativeRegion,
		region.nativeRegion, CombineModeUnion))
}

func (this *Region) Xor(rect *RectF) Status {
	return this.setStatus(GdipCombineRegionRect(this.nativeRegion, rect,
		CombineModeXor))
}

func (this *Region) XorI(rect *Rect) Status {
	return this.setStatus(GdipCombineRegionRectI(this.nativeRegion, rect,
		CombineModeXor))
}

func (this *Region) XorPath(path *GraphicsPath) Status {
	return this.setStatus(GdipCombineRegionPath(this.nativeRegion,
		path.nativePath, CombineModeXor))
}

func (this *Region) XorRegion(region *Region) Status {
	return this.setStatus(GdipCombineRegionRegion(this.nativeRegion,
		region.nativeRegion, CombineModeXor))
}

func (this *Region) Exclude(rect *RectF) Status {
	return this.setStatus(GdipCombineRegionRect(this.nativeRegion, rect,
		CombineModeExclude))
}

func (this *Region) ExcludeI(rect *Rect) Status {
	return this.setStatus(GdipCombineRegionRectI(this.nativeRegion, rect,
		CombineModeExclude))
}

func (this *Region) ExcludePath(path *GraphicsPath) Status {
	return this.setStatus(GdipCombineRegionPath(this.nativeRegion,
		path.nativePath, CombineModeExclude))
}

func (this *Region) ExcludeRegion(region *Region) Status {
	return this.setStatus(GdipCombineRegionRegion(this.nativeRegion,
		region.nativeRegion, CombineModeExclude))
}

func (this *Region) Complement(rect *RectF) Status {
	return this.setStatus(GdipCombineRegionRect(this.nativeRegion, rect,
		CombineModeComplement))
}

func (this *Region) ComplementI(rect *Rect) Status {
	return this.setStatus(GdipCombineRegionRectI(this.nativeRegion, rect,
		CombineModeComplement))
}

func (this *Region) ComplementPath(path *GraphicsPath) Status {
	return this.setStatus(GdipCombineRegionPath(this.nativeRegion,
		path.nativePath, CombineModeComplement))
}

func (this *Region) ComplementRegion(region *Region) Status {
	return this.setStatus(GdipCombineRegionRegion(this.nativeRegion,
		region.nativeRegion, CombineModeComplement))
}

func (this *Region) Translate(dx, dy REAL) Status {
	return this.setStatus(GdipTranslateRegion(this.nativeRegion, dx, dy))
}

func (this *Region) TranslateI(dx, dy INT) Status {
	return this.setStatus(GdipTranslateRegionI(this.nativeRegion, dx, dy))
}

func (this *Region) Transform(matrix *Matrix) Status {
	return this.setStatus(GdipTransformRegion(this.nativeRegion,
		matrix.nativeMatrix))
}

func (this *Region) GetBounds(g *Graphics) (rect *RectF, status Status) {
	rect = &RectF{}
	status = this.setStatus(GdipGetRegionBounds(this.nativeRegion,
		g.nativeGraphics,
		rect))
	return
}

func (this *Region) GetBoundsI(g *Graphics) (rect *Rect, status Status) {
	rect = &Rect{}
	status = this.setStatus(GdipGetRegionBoundsI(this.nativeRegion,
		g.nativeGraphics, rect))
	return
}

func (this *Region) IsEmpty(g *Graphics) (booln BOOL) {
	this.setStatus(GdipIsEmptyRegion(this.nativeRegion,
		g.nativeGraphics, &booln))
	return booln
}

func (this *Region) IsInfinite(g *Graphics) (booln BOOL) {
	this.setStatus(GdipIsInfiniteRegion(this.nativeRegion,
		g.nativeGraphics, &booln))
	return booln
}

func (this *Region) Equals(region *Region, g *Graphics) (booln BOOL) {
	this.setStatus(GdipIsEqualRegion(this.nativeRegion,
		region.nativeRegion,
		g.nativeGraphics,
		&booln))
	return
}

func (this *Region) GetDataSize() (bufferSize UINT) {
	this.setStatus(GdipGetRegionDataSize(this.nativeRegion, &bufferSize))
	return
}

func (this *Region) IsVisible(point *PointF,
	g *Graphics) (booln BOOL) {

	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	this.setStatus(GdipIsVisibleRegionPoint(this.nativeRegion,
		point.X, point.Y, nativeGraphics, &booln))
	return
}

func (this *Region) IsVisible2(rect *RectF,
	g *Graphics) (booln BOOL) {

	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	this.setStatus(GdipIsVisibleRegionRect(this.nativeRegion, rect.X,
		rect.Y, rect.Width,
		rect.Height,
		nativeGraphics,
		&booln))
	return
}

func (this *Region) IsVisibleI(point *Point,
	g *Graphics) (booln BOOL) {

	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	this.setStatus(GdipIsVisibleRegionPointI(this.nativeRegion,
		point.X, point.Y, nativeGraphics, &booln))
	return
}

func (this *Region) IsVisible2I(rect *Rect,
	g *Graphics) (booln BOOL) {

	var nativeGraphics *GpGraphics
	if g != nil {
		nativeGraphics = g.nativeGraphics
	}

	this.setStatus(GdipIsVisibleRegionRectI(this.nativeRegion, rect.X,
		rect.Y, rect.Width,
		rect.Height,
		nativeGraphics,
		&booln))
	return
}

func (this *Region) GetRegionScansCount(matrix *Matrix) (count UINT) {
	this.setStatus(GdipGetRegionScansCount(this.nativeRegion,
		&count,
		matrix.nativeMatrix))
	return
}

func (this *Region) GetRegionScans(
	matrix *Matrix, count INT) (rects []RectF, status Status) {
	rects = make([]RectF, count)
	status = this.setStatus(GdipGetRegionScans(this.nativeRegion,
		rects,
		&count,
		matrix.nativeMatrix))

	rects = rects[0:count]
	return
}

func (this *Region) GetRegionScansI(
	matrix *Matrix, count INT) (rects []Rect, status Status) {
	rects = make([]Rect, count)
	status = this.setStatus(GdipGetRegionScansI(this.nativeRegion,
		rects,
		&count,
		matrix.nativeMatrix))

	rects = rects[0:count]
	return
}
