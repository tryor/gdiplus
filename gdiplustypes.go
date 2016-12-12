package gdiplus

import (
	"math"
	"unsafe"
)
import (
	. "github.com/tryor/winapi"
)

//通过syscall.NewCallback() 返回
type CallbackHandle uintptr

//--------------------------------------------------------------------------
// Callback functions
//--------------------------------------------------------------------------

//extern "C" {
//typedef BOOL (CALLBACK * ImageAbort)(VOID *);
//typedef ImageAbort DrawImageAbort;
//typedef ImageAbort GetThumbnailImageAbort;
//}
type ImageAbort func(data DATA_PTR) uintptr
type DrawImageAbort ImageAbort
type GetThumbnailImageAbort ImageAbort

// Callback for EnumerateMetafile methods.  The parameters are:

//      recordType      WMF, EMF, or EMF+ record type
//      flags           (always 0 for WMF/EMF records)
//      dataSize        size of the record data (in bytes), or 0 if no data
//      data            pointer to the record data, or NULL if no data
//      callbackData    pointer to callbackData, if any

// This method can then call Metafile::PlayRecord to play the
// record that was just enumerated.  If this method  returns
// FALSE, the enumeration process is aborted.  Otherwise, it continues.

//extern "C" {
//typedef BOOL (CALLBACK * EnumerateMetafileProc)(EmfPlusRecordType,UINT,UINT,const BYTE*,VOID*);
//}

type EnumerateMetafileProc func(recordType EmfPlusRecordType, flags UINT, dataSize UINT, pStr *BYTE, callbackData ULONG_PTR) uintptr

//#if (GDIPVER >= 0x0110)
// This is the main GDI+ Abort interface

type GdiplusAbort struct {
	//virtual HRESULT __stdcall Abort(void) = 0;
	Abort CallbackHandle
}

//#endif //(GDIPVER >= 0x0110)

type DATA_PTR uintptr
type REAL float32

var REAL_SIZE = UINT(unsafe.Sizeof(REAL(0.0)))

const DATA_PTR_NULL = DATA_PTR(0)

//if REAL is float32
var REALbits = func(f REAL) uintptr { return uintptr(math.Float32bits(float32(f))) }

//if REAL is float64
//var REALbits = func(f REAL) uintptr { return uintptr(math.Float64bits(float64(f))) }

//type REAL64 float64

//var REAL64bits = func(f REAL64) uintptr { return uintptr(math.Float64bits(float64(f))) }

const REAL_MAX = 3.402823466e+38 //FLT_MAX
const REAL_MIN = 1.175494351e-38 //FLT_MIN
const REAL_TOLERANCE = (REAL_MIN * 100)
const REAL_EPSILON = 1.192092896e-07 //1.192092896e-07F        /* FLT_EPSILON */

type Status uint

var StatusText = func() map[Status]string {
	stmap := make(map[Status]string)
	stmap[Ok] = "Ok"
	stmap[GenericError] = "GenericError"
	stmap[InvalidParameter] = "InvalidParameter"
	stmap[OutOfMemory] = "OutOfMemory"
	stmap[ObjectBusy] = "ObjectBusy"
	stmap[InsufficientBuffer] = "InsufficientBuffer"
	stmap[NotImplemented] = "NotImplemented"
	stmap[Win32Error] = "Win32Error"
	stmap[WrongState] = "WrongState"
	stmap[Aborted] = "Aborted"
	stmap[FileNotFound] = "FileNotFound"
	stmap[ValueOverflow] = "ValueOverflow"
	stmap[AccessDenied] = "AccessDenied"
	stmap[UnknownImageFormat] = "UnknownImageFormat"
	stmap[FontFamilyNotFound] = "FontFamilyNotFound"
	stmap[FontStyleNotFound] = "FontStyleNotFound"
	stmap[NotTrueTypeFont] = "NotTrueTypeFont"
	stmap[NotFound] = "NotFound"
	stmap[UnsupportedGdiplusVersion] = "UnsupportedGdiplusVersion"
	stmap[GdiplusNotInitialized] = "GdiplusNotInitialized"

	return stmap
}()

const (
	Ok                 Status = 0
	GenericError       Status = 1
	InvalidParameter   Status = 2
	OutOfMemory        Status = 3
	ObjectBusy         Status = 4
	InsufficientBuffer Status = 5
	NotImplemented     Status = 6
	Win32Error         Status = 7
	WrongState         Status = 8
	Aborted            Status = 9
	//#ifdef DCR_USE_NEW_135429
	FileNotFound       Status = 10
	ValueOverflow      Status = 11
	AccessDenied       Status = 12
	UnknownImageFormat Status = 13
	FontFamilyNotFound Status = 14
	FontStyleNotFound  Status = 15
	NotTrueTypeFont    Status = 16
	//#else
	NotFound Status = 10
	//ValueOverflow Status = 11
	//#endif
	UnsupportedGdiplusVersion Status = 17
	GdiplusNotInitialized     Status = 18
)

//--------------------------------------------------------------------------
// Represents a dimension in a 2D coordinate system
//  (floating-point coordinates)
//--------------------------------------------------------------------------

type SizeF struct {
	Width  REAL
	Height REAL
}

//// Default constructor
// SizeF()
// {
//     Width = Height = 0.0f;
// }

func NewSizeF(width, height REAL) *SizeF {
	return &SizeF{Width: width, Height: height}
}

func NewSizeF2(size *SizeF) *SizeF {
	return &SizeF{Width: size.Width, Height: size.Height}
}

// SizeF operator+(IN const SizeF& sz) const
// {
//     return SizeF(Width + sz.Width,
//                  Height + sz.Height);
// }
func (this *SizeF) AddSizeF(sz *SizeF) *SizeF {
	return &SizeF{Width: this.Width + sz.Width, Height: this.Height + sz.Height}
}

// SizeF operator-(IN const SizeF& sz) const
// {
//     return SizeF(Width - sz.Width,
//                  Height - sz.Height);
// }
func (this *SizeF) SubSizeF(sz *SizeF) *SizeF {
	return &SizeF{Width: this.Width - sz.Width, Height: this.Height - sz.Height}
}

func (this *SizeF) Equals(sz *SizeF) bool {
	return (this.Width == sz.Width) && (this.Height == sz.Height)
}

func (this *SizeF) Empty() bool {
	return (this.Width == 0.0 && this.Height == 0.0)
}

//--------------------------------------------------------------------------
// Represents a location in a 2D coordinate system
//  (integer coordinates)
//--------------------------------------------------------------------------

type Size struct {
	Width  INT
	Height INT
}

//// Default constructor
// Size()
// {
//     Width = Height = 0.0f;
// }

func NewSize(width, height INT) *Size {
	return &Size{Width: width, Height: height}
}

func NewSize2(size *Size) *Size {
	return &Size{Width: size.Width, Height: size.Height}
}

// Size operator+(IN const Size& sz) const
// {
//     return Size(Width + sz.Width,
//                  Height + sz.Height);
// }
func (this *Size) AddSize(sz *Size) *Size {
	return &Size{Width: this.Width + sz.Width, Height: this.Height + sz.Height}
}

// Size operator-(IN const Size& sz) const
// {
//     return Size(Width - sz.Width,
//                  Height - sz.Height);
// }
func (this *Size) SubSize(sz *Size) *Size {
	return &Size{Width: this.Width - sz.Width, Height: this.Height - sz.Height}
}

func (this *Size) Equals(sz *Size) bool {
	return (this.Width == sz.Width) && (this.Height == sz.Height)
}

func (this *Size) Empty() bool {
	return (this.Width == 0.0 && this.Height == 0.0)
}

//--------------------------------------------------------------------------
// Represents a location in a 2D coordinate system
//  (floating-point coordinates)
//--------------------------------------------------------------------------

var PointF_SIZE = UINT(unsafe.Sizeof(PointF{}))

type PointF struct {
	X REAL
	Y REAL
}

func NewPointF(x, y REAL) *PointF {
	return &PointF{X: x, Y: y}
}

func NewPointF2(point *PointF) *PointF {
	return &PointF{X: point.X, Y: point.Y}
}

func NewPointF3(size *SizeF) *PointF {
	return &PointF{X: size.Width, Y: size.Height}
}

//PointF operator+(IN const PointF& point) const
//{
//    return PointF(X + point.X,
//                  Y + point.Y);
//}

func (this *PointF) AddPointF(point *PointF) *PointF {
	return &PointF{X: this.X + point.X, Y: this.Y + point.Y}
}

//PointF operator-(IN const PointF& point) const
//{
//    return PointF(X - point.X,
//                  Y - point.Y);
//}
func (this *PointF) SubPointF(point *PointF) *PointF {
	return &PointF{X: this.X - point.X, Y: this.Y - point.Y}
}

func (this *PointF) Equals(point *PointF) bool {
	return (this.X == point.X) && (this.Y == point.Y)
}

type Point struct {
	X INT
	Y INT
}

func NewPoint(x, y INT) *Point {
	return &Point{X: x, Y: y}
}

func NewPoint2(point *Point) *Point {
	return &Point{X: point.X, Y: point.Y}
}

func NewPoint3(size *Size) *Point {
	return &Point{X: size.Width, Y: size.Height}
}

//Point operator+(IN const Point& point) const
//{
//    return Point(X + point.X,
//                  Y + point.Y);
//}

func (this *Point) AddPoint(point *Point) *Point {
	return &Point{X: this.X + point.X, Y: this.Y + point.Y}
}

//Point operator-(IN const Point& point) const
//{
//    return Point(X - point.X,
//                  Y - point.Y);
//}
func (this *Point) SubPoint(point *Point) *Point {
	return &Point{X: this.X - point.X, Y: this.Y - point.Y}
}

func (this *Point) Equals(point *Point) bool {
	return (this.X == point.X) && (this.Y == point.Y)
}

//--------------------------------------------------------------------------
// Represents a rectangle in a 2D coordinate system
//  (floating-point coordinates)
//--------------------------------------------------------------------------

type RectF struct {
	X      REAL
	Y      REAL
	Width  REAL
	Height REAL
}

func NewRectF(x, y, width, height REAL) *RectF {
	return &RectF{X: x, Y: y, Width: width, Height: height}
}

func NewRectF2(location *PointF, size *SizeF) *RectF {
	return &RectF{X: location.X, Y: location.Y, Width: size.Width, Height: size.Height}
}

func (this *RectF) Clone() *RectF {
	return NewRectF(this.X, this.Y, this.Width, this.Height)
}

func (this *RectF) GetLocation(out_point *PointF) {
	out_point.X = this.X
	out_point.Y = this.Y
}

func (this *RectF) GetSize(out_size *SizeF) {
	out_size.Width = this.Width
	out_size.Height = this.Height
}

//VOID GetBounds(OUT RectF* rect) const
//{
//    rect->X = X;
//    rect->Y = Y;
//    rect->Width = Width;
//    rect->Height = Height;
//}

// Return the left, top, right, and bottom
// coordinates of the rectangle

func (this *RectF) GetLeft() REAL {
	return this.X
}

func (this *RectF) GetTop() REAL {
	return this.Y
}

func (this *RectF) GetRight() REAL {
	return this.X + this.Width
}

func (this *RectF) GetBottom() REAL {
	return this.Y + this.Height
}

// Determine if the rectangle is empty
func (this *RectF) IsEmptyArea() bool {
	return (this.Width <= REAL_EPSILON) || (this.Height <= REAL_EPSILON)
}

func (this *RectF) Equals(rect *RectF) bool {
	return this.X == rect.X &&
		this.Y == rect.Y &&
		this.Width == rect.Width &&
		this.Height == rect.Height
}

func (this *RectF) Contains(x, y REAL) bool {
	return x >= this.X && x < this.X+this.Width &&
		y >= this.Y && y < this.Y+this.Height
}

func (this *RectF) Contains2(pt *PointF) bool {
	return this.Contains(pt.X, pt.Y)
}

func (this *RectF) Contains3(rect *RectF) bool {
	return (this.X <= rect.X) && (rect.GetRight() <= this.GetRight()) &&
		(this.Y <= rect.Y) && (rect.GetBottom() <= this.GetBottom())
}

func (this *RectF) Inflate(dx, dy REAL) {
	this.X -= dx
	this.Y -= dy
	this.Width += 2 * dx
	this.Height += 2 * dy
}

func (this *RectF) Inflate2(point *PointF) {
	this.Inflate(point.X, point.Y)
}

// Intersect the current rect with the specified object

func (this *RectF) Intersect(rect *RectF) bool {
	r, ok := Intersect(this, rect)
	this.X, this.Y, this.Width, this.Height = r.X, r.Y, r.Width, r.Height
	return ok
}

// Intersect rect a and b and save the result into c
// Notice that c may be the same object as a or b.

func Intersect(in_a, in_b *RectF) (*RectF, bool) {

	right := Min((in_a.GetRight()), (in_b.GetRight()))
	bottom := Min((in_a.GetBottom()), (in_b.GetBottom()))
	left := Max((in_a.GetLeft()), (in_b.GetLeft()))
	top := Max((in_a.GetTop()), (in_b.GetTop()))

	c := &RectF{X: left, Y: top, Width: (right - left), Height: (bottom - top)}
	return c, !c.IsEmptyArea()
}

// Determine if the specified rect intersects with the
// current rect object.

func (this *RectF) IntersectsWith(rect *RectF) bool {
	return (this.GetLeft() < rect.GetRight() &&
		this.GetTop() < rect.GetTop() &&
		this.GetRight() > rect.GetLeft() &&
		this.GetBottom() > rect.GetTop())
}

func Union(a, b *RectF) (*RectF, bool) {
	right := Max((a.GetRight()), (b.GetRight()))
	bottom := Max((a.GetBottom()), (b.GetBottom()))
	left := Min((a.GetLeft()), (b.GetLeft()))
	top := Min((a.GetTop()), (b.GetTop()))

	c := &RectF{X: left, Y: top, Width: (right - left), Height: (bottom - top)}
	return c, !c.IsEmptyArea()
}

func (this *RectF) Offset(dx, dy REAL) {
	this.X += dx
	this.Y += dy
}

func (this *RectF) Offset2(point *PointF) {
	this.Offset(point.X, point.Y)
}

//--------------------------------------------------------------------------
// Represents a rectangle in a 2D coordinate system (integer coordinates)
//--------------------------------------------------------------------------

type Rect struct {
	X      INT
	Y      INT
	Width  INT
	Height INT
}

func NewRect(x, y, width, height INT) *Rect {
	return &Rect{X: x, Y: y, Width: width, Height: height}
}

func NewRect2(location *Point, size *Size) *Rect {
	return &Rect{X: location.X, Y: location.Y, Width: size.Width, Height: size.Height}
}

func (this *Rect) Clone() *Rect {
	return NewRect(this.X, this.Y, this.Width, this.Height)
}

func (this *Rect) GetLocation(out_point *Point) {
	out_point.X = this.X
	out_point.Y = this.Y
}

func (this *Rect) GetSize(out_size *Size) {
	out_size.Width = this.Width
	out_size.Height = this.Height
}

//VOID GetBounds(OUT Rect* rect) const
//{
//    rect->X = X;
//    rect->Y = Y;
//    rect->Width = Width;
//    rect->Height = Height;
//}

// Return the left, top, right, and bottom
// coordinates of the rectangle

func (this *Rect) GetLeft() INT {
	return this.X
}

func (this *Rect) GetTop() INT {
	return this.Y
}

func (this *Rect) GetRight() INT {
	return this.X + this.Width
}

func (this *Rect) GetBottom() INT {
	return this.Y + this.Height
}

// Determine if the rectangle is empty
func (this *Rect) IsEmptyArea() bool {
	return (this.Width <= 0) || (this.Height <= 0)
}

func (this *Rect) Equals(rect *Rect) bool {
	return this.X == rect.X &&
		this.Y == rect.Y &&
		this.Width == rect.Width &&
		this.Height == rect.Height
}

func (this *Rect) Contains(x, y INT) bool {
	return x >= this.X && x < this.X+this.Width &&
		y >= this.Y && y < this.Y+this.Height
}

func (this *Rect) Contains2(pt *Point) bool {
	return this.Contains(pt.X, pt.Y)
}

func (this *Rect) Contains3(rect *Rect) bool {
	return (this.X <= rect.X) && (rect.GetRight() <= this.GetRight()) &&
		(this.Y <= rect.Y) && (rect.GetBottom() <= this.GetBottom())
}

func (this *Rect) Inflate(dx, dy INT) {
	this.X -= dx
	this.Y -= dy
	this.Width += 2 * dx
	this.Height += 2 * dy
}

func (this *Rect) Inflate2(point *Point) {
	this.Inflate(point.X, point.Y)
}

// Intersect the current rect with the specified object

func (this *Rect) Intersect(rect *Rect) bool {
	r, ok := IntersectI(this, rect)
	this.X, this.Y, this.Width, this.Height = r.X, r.Y, r.Width, r.Height
	return ok
}

// Intersect rect a and b and save the result into c
// Notice that c may be the same object as a or b.

func IntersectI(in_a, in_b *Rect) (*Rect, bool) {

	right := MinI((in_a.GetRight()), (in_b.GetRight()))
	bottom := MinI((in_a.GetBottom()), (in_b.GetBottom()))
	left := MaxI((in_a.GetLeft()), (in_b.GetLeft()))
	top := MaxI((in_a.GetTop()), (in_b.GetTop()))

	c := &Rect{X: left, Y: top, Width: (right - left), Height: (bottom - top)}
	return c, !c.IsEmptyArea()
}

// Determine if the specified rect intersects with the
// current rect object.

func (this *Rect) IntersectsWith(rect *Rect) bool {
	return (this.GetLeft() < rect.GetRight() &&
		this.GetTop() < rect.GetTop() &&
		this.GetRight() > rect.GetLeft() &&
		this.GetBottom() > rect.GetTop())
}

func UnionI(a, b *Rect) (*Rect, bool) {
	right := MaxI((a.GetRight()), (b.GetRight()))
	bottom := MaxI((a.GetBottom()), (b.GetBottom()))
	left := MinI((a.GetLeft()), (b.GetLeft()))
	top := MinI((a.GetTop()), (b.GetTop()))
	c := &Rect{X: left, Y: top, Width: (right - left), Height: (bottom - top)}
	return c, !c.IsEmptyArea()
}

func (this *Rect) Offset(dx, dy INT) {
	this.X += dx
	this.Y += dy
}

func (this *Rect) Offset2(point *Point) {
	this.Offset(point.X, point.Y)
}

type PathData struct {
	Count   INT
	_points *PointF
	Types   []BYTE
	Points  []PointF
}

func NewPathData(count INT) *PathData {
	pathdata := &PathData{Count: count}
	if count > 0 {
		pathdata.Points = make([]PointF, count)
		pathdata._points = &pathdata.Points[0]
		pathdata.Types = make([]BYTE, count, count)
	}
	return pathdata
}

//var PathData_SIZE = UINT(unsafe.Sizeof(PathData{}))

type CharacterRange struct {
	First  INT
	Length INT
}
