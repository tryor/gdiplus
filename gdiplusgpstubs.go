package gdiplus

type IDirectDrawSurface7 interface{}

//---------------------------------------------------------------------------
// Internal GDI+ classes for internal type checking
//---------------------------------------------------------------------------
type GpGraphics interface{}

type GpBrush interface{}
type GpTexture interface {
	GpBrush
}
type GpSolidFill interface {
	GpBrush
}
type GpLineGradient interface {
	GpBrush
}
type GpPathGradient interface {
	GpBrush
}
type GpHatch interface {
	GpBrush
}

type GpPen interface{}
type GpCustomLineCap interface{}
type GpAdjustableArrowCap interface {
	GpCustomLineCap
}

type GpImage interface{}
type GpBitmap interface {
	GpImage
}
type GpMetafile interface {
	GpImage
}
type GpImageAttributes interface{}

type GpPath interface{}
type GpRegion interface{}
type GpPathIterator interface{}

type GpFontFamily interface{}
type GpFont interface{}
type GpStringFormat interface{}
type GpFontCollection interface{}
type GpInstalledFontCollection interface {
	GpFontCollection
}
type GpPrivateFontCollection interface {
	GpFontCollection
}

type GpCachedBitmap interface{}

//type GpStatus Status

type GpFillMode FillMode
type GpWrapMode WrapMode
type GpUnit Unit

type GpCoordinateSpace CoordinateSpace

type GpPointF PointF
type GpPoint Point
type GpRectF RectF
type GpRect Rect

//type SizeF GpSizeF;
//typedef HatchStyle GpHatchStyle;
//typedef DashStyle GpDashStyle;
//typedef LineCap GpLineCap;
//typedef DashCap GpDashCap;

//typedef PenAlignment GpPenAlignment;

//typedef LineJoin GpLineJoin;
//typedef PenType GpPenType;
type GpMatrix Matrix
type GpBrushType BrushType

//typedef MatrixOrder GpMatrixOrder;
type GpFlushIntention FlushIntention

type GpPathData PathData
