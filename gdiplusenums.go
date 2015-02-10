package gdiplus

import (
	. "github.com/trygo/winapi"
)

//--------------------------------------------------------------------------
// Default bezier flattening tolerance in device pixels.
//--------------------------------------------------------------------------
const FlatnessDefault REAL = 1.0 / 4.0

//--------------------------------------------------------------------------
// Graphics and Container State cookies
//--------------------------------------------------------------------------

type GraphicsState UINT
type GraphicsContainer UINT

//--------------------------------------------------------------------------
// Fill mode constants
//--------------------------------------------------------------------------

type FillMode INT

const (
	FillModeAlternate FillMode = 0 // 0
	FillModeWinding   FillMode = 1 // 1
)

//--------------------------------------------------------------------------
// Quality mode constants
//--------------------------------------------------------------------------

type QualityMode int

const (
	QualityModeInvalid QualityMode = -1
	QualityModeDefault QualityMode = 0
	QualityModeLow     QualityMode = 1 // for apps that need the best performance
	QualityModeHigh    QualityMode = 2 // for apps that need the best rendering quality
)

//--------------------------------------------------------------------------
// Alpha compositing mode constants
//--------------------------------------------------------------------------

type CompositingMode INT

const (
	CompositingModeSourceOver CompositingMode = 0 // 0
	CompositingModeSourceCopy CompositingMode = 1 // 1
)

//--------------------------------------------------------------------------
// Alpha compositing quality constants
//--------------------------------------------------------------------------
type CompositingQuality int

const (
	CompositingQualityInvalid        CompositingQuality = CompositingQuality(QualityModeInvalid)
	CompositingQualityDefault        CompositingQuality = CompositingQuality(QualityModeDefault)
	CompositingQualityHighSpeed      CompositingQuality = CompositingQuality(QualityModeLow)
	CompositingQualityHighQuality    CompositingQuality = CompositingQuality(QualityModeHigh)
	CompositingQualityGammaCorrected                    = 3
	CompositingQualityAssumeLinear                      = 4
)

//--------------------------------------------------------------------------
// Unit constants
//--------------------------------------------------------------------------
type Unit INT

const (
	UnitWorld      Unit = iota // 0 -- World coordinate (non-physical unit)
	UnitDisplay                // 1 -- Variable -- for PageTransform only
	UnitPixel                  // 2 -- Each unit is one device pixel.
	UnitPoint                  // 3 -- Each unit is a printer's point, or 1/72 inch.
	UnitInch                   // 4 -- Each unit is 1 inch.
	UnitDocument               // 5 -- Each unit is 1/300 inch.
	UnitMillimeter             // 6 -- Each unit is 1 millimeter.
)

//--------------------------------------------------------------------------
// MetafileFrameUnit
//
// The frameRect for creating a metafile can be specified in any of these
// units.  There is an extra frame unit value (MetafileFrameUnitGdi) so
// that units can be supplied in the same units that GDI expects for
// frame rects -- these units are in .01 (1/100ths) millimeter units
// as defined by GDI.
//--------------------------------------------------------------------------

type MetafileFrameUnit INT

const (
	MetafileFrameUnitPixel      MetafileFrameUnit = MetafileFrameUnit(UnitPixel)
	MetafileFrameUnitPoint      MetafileFrameUnit = MetafileFrameUnit(UnitPoint)
	MetafileFrameUnitInch       MetafileFrameUnit = MetafileFrameUnit(UnitInch)
	MetafileFrameUnitDocument   MetafileFrameUnit = MetafileFrameUnit(UnitDocument)
	MetafileFrameUnitMillimeter MetafileFrameUnit = MetafileFrameUnit(UnitMillimeter)
	MetafileFrameUnitGdi        MetafileFrameUnit = 7 // GDI compatible .01 MM units
)

//--------------------------------------------------------------------------
// Coordinate space identifiers
//--------------------------------------------------------------------------

type CoordinateSpace INT

const (
	CoordinateSpaceWorld  CoordinateSpace = iota // 0
	CoordinateSpacePage                          // 1
	CoordinateSpaceDevice                        // 2
)

//--------------------------------------------------------------------------
// Various wrap modes for brushes
//--------------------------------------------------------------------------
type WrapMode INT

const (
	WrapModeTile       WrapMode = iota // 0
	WrapModeTileFlipX                  // 1
	WrapModeTileFlipY                  // 2
	WrapModeTileFlipXY                 // 3
	WrapModeClamp                      // 4
)

//--------------------------------------------------------------------------
// Various hatch styles
//--------------------------------------------------------------------------
type HatchStyle INT

const (
	HatchStyleHorizontal             HatchStyle        = iota // 0
	HatchStyleVertical                                        // 1
	HatchStyleForwardDiagonal                                 // 2
	HatchStyleBackwardDiagonal                                // 3
	HatchStyleCross                                           // 4
	HatchStyleDiagonalCross                                   // 5
	HatchStyle05Percent                                       // 6
	HatchStyle10Percent                                       // 7
	HatchStyle20Percent                                       // 8
	HatchStyle25Percent                                       // 9
	HatchStyle30Percent                                       // 10
	HatchStyle40Percent                                       // 11
	HatchStyle50Percent                                       // 12
	HatchStyle60Percent                                       // 13
	HatchStyle70Percent                                       // 14
	HatchStyle75Percent                                       // 15
	HatchStyle80Percent                                       // 16
	HatchStyle90Percent                                       // 17
	HatchStyleLightDownwardDiagonal                           // 18
	HatchStyleLightUpwardDiagonal                             // 19
	HatchStyleDarkDownwardDiagonal                            // 20
	HatchStyleDarkUpwardDiagonal                              // 21
	HatchStyleWideDownwardDiagonal                            // 22
	HatchStyleWideUpwardDiagonal                              // 23
	HatchStyleLightVertical                                   // 24
	HatchStyleLightHorizontal                                 // 25
	HatchStyleNarrowVertical                                  // 26
	HatchStyleNarrowHorizontal                                // 27
	HatchStyleDarkVertical                                    // 28
	HatchStyleDarkHorizontal                                  // 29
	HatchStyleDashedDownwardDiagonal                          // 30
	HatchStyleDashedUpwardDiagonal                            // 31
	HatchStyleDashedHorizontal                                // 32
	HatchStyleDashedVertical                                  // 33
	HatchStyleSmallConfetti                                   // 34
	HatchStyleLargeConfetti                                   // 35
	HatchStyleZigZag                                          // 36
	HatchStyleWave                                            // 37
	HatchStyleDiagonalBrick                                   // 38
	HatchStyleHorizontalBrick                                 // 39
	HatchStyleWeave                                           // 40
	HatchStylePlaid                                           // 41
	HatchStyleDivot                                           // 42
	HatchStyleDottedGrid                                      // 43
	HatchStyleDottedDiamond                                   // 44
	HatchStyleShingle                                         // 45
	HatchStyleTrellis                                         // 46
	HatchStyleSphere                                          // 47
	HatchStyleSmallGrid                                       // 48
	HatchStyleSmallCheckerBoard                               // 49
	HatchStyleLargeCheckerBoard                               // 50
	HatchStyleOutlinedDiamond                                 // 51
	HatchStyleSolidDiamond                                    // 52
	HatchStyleTotal                                           // must be after all unique hatch styles
	HatchStyleLargeGrid              = HatchStyleCross        // 4  an alias for the cross style
	HatchStyleMin                    = HatchStyleHorizontal
	HatchStyleMax                    = HatchStyleTotal - 1
)

//--------------------------------------------------------------------------
// Dash style constants
//--------------------------------------------------------------------------
type DashStyle INT

const (
	DashStyleSolid      DashStyle = iota // 0
	DashStyleDash                        // 1
	DashStyleDot                         // 2
	DashStyleDashDot                     // 3
	DashStyleDashDotDot                  // 4
	DashStyleCustom                      // 5
)

//--------------------------------------------------------------------------
// Dash cap constants
//--------------------------------------------------------------------------
type DashCap INT

const (
	DashCapFlat     DashCap = 0
	DashCapRound    DashCap = 2
	DashCapTriangle DashCap = 3
)

//--------------------------------------------------------------------------
// Line cap constants (only the lowest 8 bits are used).
//--------------------------------------------------------------------------
type LineCap INT

const (
	LineCapFlat     LineCap = 0
	LineCapSquare   LineCap = 1
	LineCapRound    LineCap = 2
	LineCapTriangle LineCap = 3

	LineCapNoAnchor      LineCap = 0x10 // corresponds to flat cap
	LineCapSquareAnchor  LineCap = 0x11 // corresponds to square cap
	LineCapRoundAnchor   LineCap = 0x12 // corresponds to round cap
	LineCapDiamondAnchor LineCap = 0x13 // corresponds to triangle cap
	LineCapArrowAnchor   LineCap = 0x14 // no correspondence

	LineCapCustom LineCap = 0xff // custom cap

	LineCapAnchorMask LineCap = 0xf0 // mask to check for anchor or not.
)

//--------------------------------------------------------------------------
// Custom Line cap type constants
//--------------------------------------------------------------------------
type CustomLineCapType INT

const (
	CustomLineCapTypeDefault         CustomLineCapType = 0
	CustomLineCapTypeAdjustableArrow CustomLineCapType = 1
)

//--------------------------------------------------------------------------
// Line join constants
//--------------------------------------------------------------------------
type LineJoin INT

const (
	LineJoinMiter        LineJoin = 0
	LineJoinBevel        LineJoin = 1
	LineJoinRound        LineJoin = 2
	LineJoinMiterClipped LineJoin = 3
)

//--------------------------------------------------------------------------
// Path point types (only the lowest 8 bits are used.)
//  The lowest 3 bits are interpreted as point type
//  The higher 5 bits are reserved for flags.
//--------------------------------------------------------------------------
type PathPointType INT

const (
	PathPointTypeStart        PathPointType = 0    // move
	PathPointTypeLine         PathPointType = 1    // line
	PathPointTypeBezier       PathPointType = 3    // default Beizer (= cubic Bezier)
	PathPointTypePathTypeMask PathPointType = 0x07 // type mask (lowest 3 bits).
	PathPointTypeDashMode     PathPointType = 0x10 // currently in dash mode.
	PathPointTypePathMarker   PathPointType = 0x20 // a marker for the path.
	PathPointTypeCloseSubpath PathPointType = 0x80 // closed flag

	// Path types used for advanced path.

	PathPointTypeBezier2 PathPointType = 2 // quadratic Beizer
	PathPointTypeBezier3 PathPointType = 3 // cubic Bezier
	PathPointTypeBezier4 PathPointType = 4 // quartic (4th order) Beizer
	PathPointTypeBezier5 PathPointType = 5 // quintic (5th order) Bezier
	PathPointTypeBezier6 PathPointType = 6 // hexaic (6th order) Bezier
)

//--------------------------------------------------------------------------
// WarpMode constants
//--------------------------------------------------------------------------
type WarpMode INT

const (
	WarpModePerspective WarpMode = iota // 0
	WarpModeBilinear                    // 1
)

//--------------------------------------------------------------------------
// LineGradient Mode
//--------------------------------------------------------------------------
type LinearGradientMode INT

const (
	LinearGradientModeHorizontal       LinearGradientMode = iota // 0
	LinearGradientModeVertical                                   // 1
	LinearGradientModeForwardDiagonal                            // 2
	LinearGradientModeBackwardDiagonal                           // 3
)

//--------------------------------------------------------------------------
// Region Comine Modes
//--------------------------------------------------------------------------
type CombineMode INT

const (
	CombineModeReplace    CombineMode = iota // 0
	CombineModeIntersect                     // 1
	CombineModeUnion                         // 2
	CombineModeXor                           // 3
	CombineModeExclude                       // 4
	CombineModeComplement                    // 5 (does exclude from)
)

//--------------------------------------------------------------------------
// Image types
//--------------------------------------------------------------------------
type ImageType INT

const (
	ImageTypeUnknown  = iota // 0
	ImageTypeBitmap          // 1
	ImageTypeMetafile        // 2
)

//--------------------------------------------------------------------------
// Interpolation modes
//--------------------------------------------------------------------------
type InterpolationMode int

const (
	InterpolationModeInvalid             InterpolationMode = InterpolationMode(QualityModeInvalid)
	InterpolationModeDefault             InterpolationMode = InterpolationMode(QualityModeDefault)
	InterpolationModeLowQuality          InterpolationMode = InterpolationMode(QualityModeLow)
	InterpolationModeHighQuality         InterpolationMode = InterpolationMode(QualityModeHigh)
	InterpolationModeBilinear            InterpolationMode = 3
	InterpolationModeBicubic             InterpolationMode = 4
	InterpolationModeNearestNeighbor     InterpolationMode = 5
	InterpolationModeHighQualityBilinear InterpolationMode = 6
	InterpolationModeHighQualityBicubic  InterpolationMode = 7
)

//--------------------------------------------------------------------------
// Pen types
//--------------------------------------------------------------------------
type PenAlignment INT

const (
	PenAlignmentCenter PenAlignment = 0
	PenAlignmentInset  PenAlignment = 1
	PenAlignmentOutset PenAlignment = 2
	PenAlignmentLeft   PenAlignment = 3
	PenAlignmentRight  PenAlignment = 4
)

//--------------------------------------------------------------------------
// Brush types
//--------------------------------------------------------------------------
type BrushType int

const (
	BrushTypeSolidColor     BrushType = 0
	BrushTypeHatchFill      BrushType = 1
	BrushTypeTextureFill    BrushType = 2
	BrushTypePathGradient   BrushType = 3
	BrushTypeLinearGradient BrushType = 4
)

//--------------------------------------------------------------------------
// Pen's Fill types
//--------------------------------------------------------------------------
type PenType int

const (
	PenTypeSolidColor     PenType = PenType(BrushTypeSolidColor)
	PenTypeHatchFill      PenType = PenType(BrushTypeHatchFill)
	PenTypeTextureFill    PenType = PenType(BrushTypeTextureFill)
	PenTypePathGradient   PenType = PenType(BrushTypePathGradient)
	PenTypeLinearGradient PenType = PenType(BrushTypeLinearGradient)
	PenTypeUnknown        PenType = -1
)

//--------------------------------------------------------------------------
// Matrix Order
//--------------------------------------------------------------------------

type MatrixOrder INT

const (
	MatrixOrderPrepend MatrixOrder = 0
	MatrixOrderAppend  MatrixOrder = 1
)

//--------------------------------------------------------------------------
// Generic font families
//--------------------------------------------------------------------------
type GenericFontFamily INT

const (
	GenericFontFamilySerif GenericFontFamily = iota
	GenericFontFamilySansSerif
	GenericFontFamilyMonospace
)

//--------------------------------------------------------------------------
// FontStyle: face types and common styles
//--------------------------------------------------------------------------

//  These should probably be flags

//  Must have:
//      Regular = 0
//      Bold = 1
//      Italic = 2
//      BoldItalic = 3
type FontStyle int

const (
	FontStyleRegular    FontStyle = 0
	FontStyleBold       FontStyle = 1
	FontStyleItalic     FontStyle = 2
	FontStyleBoldItalic FontStyle = 3
	FontStyleUnderline  FontStyle = 4
	FontStyleStrikeout  FontStyle = 8
)

//---------------------------------------------------------------------------
// Smoothing Mode
//---------------------------------------------------------------------------
type SmoothingMode int

const (
	SmoothingModeInvalid     SmoothingMode = SmoothingMode(QualityModeInvalid)
	SmoothingModeDefault     SmoothingMode = SmoothingMode(QualityModeDefault)
	SmoothingModeHighSpeed   SmoothingMode = SmoothingMode(QualityModeLow)
	SmoothingModeHighQuality SmoothingMode = SmoothingMode(QualityModeHigh)
	SmoothingModeNone        SmoothingMode = 3
	SmoothingModeAntiAlias   SmoothingMode = 4

	//#if (GDIPVER >= 0x0110)
	SmoothingModeAntiAlias8x4 = SmoothingModeAntiAlias
	SmoothingModeAntiAlias8x8 = 5

//#endif //(GDIPVER >= 0x0110)
)

//---------------------------------------------------------------------------
// Pixel Format Mode
//---------------------------------------------------------------------------
type PixelOffsetMode int

const (
	PixelOffsetModeInvalid     PixelOffsetMode = PixelOffsetMode(QualityModeInvalid)
	PixelOffsetModeDefault     PixelOffsetMode = PixelOffsetMode(QualityModeDefault)
	PixelOffsetModeHighSpeed   PixelOffsetMode = PixelOffsetMode(QualityModeLow)
	PixelOffsetModeHighQuality PixelOffsetMode = PixelOffsetMode(QualityModeHigh)
	PixelOffsetModeNone        PixelOffsetMode = 3 // no pixel offset
	PixelOffsetModeHalf        PixelOffsetMode = 4 // offset by -0.5, -0.5 for fast anti-alias perf
)

//---------------------------------------------------------------------------
// Text Rendering Hint
//---------------------------------------------------------------------------
type TextRenderingHint INT

const (
	//#ifdef DCR_USE_NEW_186764
	TextRenderingHintSystemDefault            TextRenderingHint = 0 // Glyph with system default rendering hint
	TextRenderingHintSingleBitPerPixelGridFit TextRenderingHint = 1 // Glyph bitmap with hinting
	//#else
	//TextRenderingHintSingleBitPerPixelGridFit TextRenderingHint = 0 // Glyph bitmap with hinting
	//#endif // DCR_USE_NEW_186764
	TextRenderingHintSingleBitPerPixel TextRenderingHint = 2 // Glyph bitmap without hinting
	TextRenderingHintAntiAliasGridFit  TextRenderingHint = 3 // Glyph anti-alias bitmap with hinting
	TextRenderingHintAntiAlias         TextRenderingHint = 4 // Glyph anti-alias bitmap without hinting
	TextRenderingHintClearTypeGridFit  TextRenderingHint = 5 // Glyph CT bitmap with hinting
)

//---------------------------------------------------------------------------
// Metafile Types
//---------------------------------------------------------------------------
type MetafileType INT

//const (
//	MetafileTypeInvalid     MetafileType = iota // Invalid metafile
//	MetafileTypeWmf                             // Standard WMF
//	MetafileTypeWmfAldus                        // Aldus Placeable Metafile format
//	MetafileTypeEmf                             // EMF (not EMF+)
//	MetafileTypeEmfPlusOnly                     // EMF+ without dual, down-level records
//	MetafileTypeEmfPlusDual                     // EMF+ with dual, down-level records
//)

//---------------------------------------------------------------------------
// Metafile Types
//---------------------------------------------------------------------------

const (
	MetafileTypeInvalid      MetafileType = iota // Invalid metafile
	MetafileTypeWmf                              // Standard WMF
	MetafileTypeWmfPlaceable                     // Placeable WMF
	MetafileTypeEmf                              // EMF (not EMF+)
	MetafileTypeEmfPlusOnly                      // EMF+ without dual, down-level records
	MetafileTypeEmfPlusDual                      // EMF+ with dual, down-level records
)

// Specifies the type of EMF to record
type EmfType INT

const (
	EmfTypeEmfOnly     EmfType = EmfType(MetafileTypeEmf)         // no EMF+, only EMF
	EmfTypeEmfPlusOnly EmfType = EmfType(MetafileTypeEmfPlusOnly) // no EMF, only EMF+
	EmfTypeEmfPlusDual EmfType = EmfType(MetafileTypeEmfPlusDual) // both EMF+ and EMF
)

// All persistent objects must have a type listed here
type ObjectType INT

const (
	ObjectTypeInvalid ObjectType = iota
	ObjectTypeBrush
	ObjectTypePen
	ObjectTypePath
	ObjectTypeRegion
	ObjectTypeImage
	ObjectTypeFont
	ObjectTypeStringFormat
	ObjectTypeImageAttributes
	ObjectTypeCustomLineCap

	//#if (GDIPVER >= 0x0110)
	ObjectTypeGraphics

	ObjectTypeMax = ObjectTypeGraphics
	//#else
	//ObjectTypeMax = ObjectTypeCustomLineCap
	//#endif //(GDIPVER >= 0x0110)
	ObjectTypeMin = ObjectTypeBrush
)

func ObjectTypeIsValid(typ ObjectType) bool {
	return ((typ >= ObjectTypeMin) && (typ <= ObjectTypeMax))
}

//type RecordType INT
const (
	META_EOF                   = 0x0000
	META_REALIZEPALETTE        = 0x0035
	META_SETPALENTRIES         = 0x0037
	META_SETBKMODE             = 0x0102
	META_SETMAPMODE            = 0x0103
	META_SETROP2               = 0x0104
	META_SETRELABS             = 0x0105
	META_SETPOLYFILLMODE       = 0x0106
	META_SETSTRETCHBLTMODE     = 0x0107
	META_SETTEXTCHAREXTRA      = 0x0108
	META_RESTOREDC             = 0x0127
	META_RESIZEPALETTE         = 0x0139
	META_DIBCREATEPATTERNBRUSH = 0x0142
	META_SETLAYOUT             = 0x0149
	META_SETBKCOLOR            = 0x0201
	META_SETTEXTCOLOR          = 0x0209
	META_OFFSETVIEWPORTORG     = 0x0211
	META_LINETO                = 0x0213
	META_MOVETO                = 0x0214
	META_OFFSETCLIPRGN         = 0x0220
	META_FILLREGION            = 0x0228
	META_SETMAPPERFLAGS        = 0x0231
	META_SELECTPALETTE         = 0x0234
	META_POLYGON               = 0x0324
	META_POLYLINE              = 0x0325
	META_SETTEXTJUSTIFICATION  = 0x020A
	META_SETWINDOWORG          = 0x020B
	META_SETWINDOWEXT          = 0x020C
	META_SETVIEWPORTORG        = 0x020D
	META_SETVIEWPORTEXT        = 0x020E
	META_OFFSETWINDOWORG       = 0x020F
	META_SCALEWINDOWEXT        = 0x0410
	META_SCALEVIEWPORTEXT      = 0x0412
	META_EXCLUDECLIPRECT       = 0x0415
	META_INTERSECTCLIPRECT     = 0x0416
	META_ELLIPSE               = 0x0418
	META_FLOODFILL             = 0x0419
	META_FRAMEREGION           = 0x0429
	META_ANIMATEPALETTE        = 0x0436
	META_TEXTOUT               = 0x0521
	META_POLYPOLYGON           = 0x0538
	META_EXTFLOODFILL          = 0x0548
	META_RECTANGLE             = 0x041B
	META_SETPIXEL              = 0x041F
	META_ROUNDRECT             = 0x061C
	META_PATBLT                = 0x061D
	META_SAVEDC                = 0x001E
	META_PIE                   = 0x081A
	META_STRETCHBLT            = 0x0B23
	META_ESCAPE                = 0x0626
	META_INVERTREGION          = 0x012A
	META_PAINTREGION           = 0x012B
	META_SELECTCLIPREGION      = 0x012C
	META_SELECTOBJECT          = 0x012D
	META_SETTEXTALIGN          = 0x012E
	META_ARC                   = 0x0817
	META_CHORD                 = 0x0830
	META_BITBLT                = 0x0922
	META_EXTTEXTOUT            = 0x0a32
	META_SETDIBTODEV           = 0x0d33
	META_DIBBITBLT             = 0x0940
	META_DIBSTRETCHBLT         = 0x0b41
	META_STRETCHDIB            = 0x0f43
	META_DELETEOBJECT          = 0x01f0
	META_CREATEPALETTE         = 0x00f7
	META_CREATEPATTERNBRUSH    = 0x01F9
	META_CREATEPENINDIRECT     = 0x02FA
	META_CREATEFONTINDIRECT    = 0x02FB
	META_CREATEBRUSHINDIRECT   = 0x02FC
	META_CREATEREGION          = 0x06FF
)

//TEnhMetaRecord.iType 的可能值:
const (
	EMR_HEADER                  = 1
	EMR_POLYBEZIER              = 2
	EMR_POLYGON                 = 3
	EMR_POLYLINE                = 4
	EMR_POLYBEZIERTO            = 5
	EMR_POLYLINETO              = 6
	EMR_POLYPOLYLINE            = 7
	EMR_POLYPOLYGON             = 8
	EMR_SETWINDOWEXTEX          = 9
	EMR_SETWINDOWORGEX          = 10
	EMR_SETVIEWPORTEXTEX        = 11
	EMR_SETVIEWPORTORGEX        = 12
	EMR_SETBRUSHORGEX           = 13
	EMR_EOF                     = 14
	EMR_SETPIXELV               = 15
	EMR_SETMAPPERFLAGS          = 0x10
	EMR_SETMAPMODE              = 17
	EMR_SETBKMODE               = 18
	EMR_SETPOLYFILLMODE         = 19
	EMR_SETROP2                 = 20
	EMR_SETSTRETCHBLTMODE       = 21
	EMR_SETTEXTALIGN            = 22
	EMR_SETCOLORADJUSTMENT      = 23
	EMR_SETTEXTCOLOR            = 24
	EMR_SETBKCOLOR              = 25
	EMR_OFFSETCLIPRGN           = 26
	EMR_MOVETOEX                = 27
	EMR_SETMETARGN              = 28
	EMR_EXCLUDECLIPRECT         = 29
	EMR_INTERSECTCLIPRECT       = 30
	EMR_SCALEVIEWPORTEXTEX      = 31
	EMR_SCALEWINDOWEXTEX        = 32
	EMR_SAVEDC                  = 33
	EMR_RESTOREDC               = 34
	EMR_SETWORLDTRANSFORM       = 35
	EMR_MODIFYWORLDTRANSFORM    = 36
	EMR_SELECTOBJECT            = 37
	EMR_CREATEPEN               = 38
	EMR_CREATEBRUSHINDIRECT     = 39
	EMR_DELETEOBJECT            = 40
	EMR_ANGLEARC                = 41
	EMR_ELLIPSE                 = 42
	EMR_RECTANGLE               = 43
	EMR_ROUNDRECT               = 44
	EMR_ARC                     = 45
	EMR_CHORD                   = 46
	EMR_PIE                     = 47
	EMR_SELECTPALETTE           = 48
	EMR_CREATEPALETTE           = 49
	EMR_SETPALETTEENTRIES       = 50
	EMR_RESIZEPALETTE           = 51
	EMR_REALIZEPALETTE          = 52
	EMR_EXTFLOODFILL            = 53
	EMR_LINETO                  = 54
	EMR_ARCTO                   = 55
	EMR_POLYDRAW                = 56
	EMR_SETARCDIRECTION         = 57
	EMR_SETMITERLIMIT           = 58
	EMR_BEGINPATH               = 59
	EMR_ENDPATH                 = 60
	EMR_CLOSEFIGURE             = 61
	EMR_FILLPATH                = 62
	EMR_STROKEANDFILLPATH       = 63
	EMR_STROKEPATH              = 0x40
	EMR_FLATTENPATH             = 65
	EMR_WIDENPATH               = 66
	EMR_SELECTCLIPPATH          = 67
	EMR_ABORTPATH               = 68
	EMR_GDICOMMENT              = 70
	EMR_FILLRGN                 = 71
	EMR_FRAMERGN                = 72
	EMR_INVERTRGN               = 73
	EMR_PAINTRGN                = 74
	EMR_EXTSELECTCLIPRGN        = 75
	EMR_BITBLT                  = 76
	EMR_STRETCHBLT              = 77
	EMR_MASKBLT                 = 78
	EMR_PLGBLT                  = 79
	EMR_SETDIBITSTODEVICE       = 80
	EMR_STRETCHDIBITS           = 81
	EMR_EXTCREATEFONTINDIRECTW  = 82
	EMR_EXTTEXTOUTA             = 83
	EMR_EXTTEXTOUTW             = 84
	EMR_POLYBEZIER16            = 85
	EMR_POLYGON16               = 86
	EMR_POLYLINE16              = 87
	EMR_POLYBEZIERTO16          = 88
	EMR_POLYLINETO16            = 89
	EMR_POLYPOLYLINE16          = 90
	EMR_POLYPOLYGON16           = 91
	EMR_POLYDRAW16              = 92
	EMR_CREATEMONOBRUSH         = 93
	EMR_CREATEDIBPATTERNBRUSHPT = 94
	EMR_EXTCREATEPEN            = 95
	EMR_POLYTEXTOUTA            = 96
	EMR_POLYTEXTOUTW            = 97
	EMR_SETICMMODE              = 98
	EMR_CREATECOLORSPACE        = 99
	EMR_SETCOLORSPACE           = 100
	EMR_DELETECOLORSPACE        = 101
	EMR_GLSRECORD               = 102
	EMR_GLSBOUNDEDRECORD        = 103
	EMR_PIXELFORMAT             = 104
	EMR_DRAWESCAPE              = 105
	EMR_EXTESCAPE               = 106
	EMR_STARTDOC                = 107
	EMR_SMALLTEXTOUT            = 108
	EMR_FORCEUFIMAPPING         = 109
	EMR_NAMEDESCAPE             = 110
	EMR_COLORCORRECTPALETTE     = 111
	EMR_SETICMPROFILEA          = 112
	EMR_SETICMPROFILEW          = 113
	EMR_ALPHABLEND              = 114
	EMR_ALPHADIBBLEND           = 115
	EMR_TRANSPARENTBLT          = 116
	EMR_TRANSPARENTDIB          = 117
	EMR_GRADIENTFILL            = 118
	EMR_SETLINKEDUFIS           = 119
	EMR_SETTEXTJUSTIFICATION    = 120
)

const GDIP_EMFPLUS_RECORD_BASE INT = 0x00004000
const GDIP_WMF_RECORD_BASE INT = 0x00010000

func GDIP_WMF_RECORD_TO_EMFPLUS(n INT) INT { return n | GDIP_WMF_RECORD_BASE }
func GDIP_EMFPLUS_RECORD_TO_WMF(n INT) INT { return n & (^GDIP_WMF_RECORD_BASE) }
func GDIP_IS_WMF_RECORDTYPE(n INT) bool    { return (n & GDIP_WMF_RECORD_BASE) != 0 }

type EmfPlusRecordType INT

var (
	// Since we have to enumerate GDI records right along with GDI+ records
	// we list all the GDI records here so that they can be part of the
	// same enumeration type which is used in the enumeration callback.

	WmfRecordTypeSetBkColor            EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETBKCOLOR))
	WmfRecordTypeSetBkMode             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETBKMODE))
	WmfRecordTypeSetMapMode            EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETMAPMODE))
	WmfRecordTypeSetROP2               EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETROP2))
	WmfRecordTypeSetRelAbs             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETRELABS))
	WmfRecordTypeSetPolyFillMode       EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETPOLYFILLMODE))
	WmfRecordTypeSetStretchBltMode     EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETSTRETCHBLTMODE))
	WmfRecordTypeSetTextCharExtra      EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETTEXTCHAREXTRA))
	WmfRecordTypeSetTextColor          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETTEXTCOLOR))
	WmfRecordTypeSetTextJustification  EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETTEXTJUSTIFICATION))
	WmfRecordTypeSetWindowOrg          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETWINDOWORG))
	WmfRecordTypeSetWindowExt          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETWINDOWEXT))
	WmfRecordTypeSetViewportOrg        EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETVIEWPORTORG))
	WmfRecordTypeSetViewportExt        EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETVIEWPORTEXT))
	WmfRecordTypeOffsetWindowOrg       EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_OFFSETWINDOWORG))
	WmfRecordTypeScaleWindowExt        EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SCALEWINDOWEXT))
	WmfRecordTypeOffsetViewportOrg     EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_OFFSETVIEWPORTORG))
	WmfRecordTypeScaleViewportExt      EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SCALEVIEWPORTEXT))
	WmfRecordTypeLineTo                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_LINETO))
	WmfRecordTypeMoveTo                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_MOVETO))
	WmfRecordTypeExcludeClipRect       EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_EXCLUDECLIPRECT))
	WmfRecordTypeIntersectClipRect     EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_INTERSECTCLIPRECT))
	WmfRecordTypeArc                   EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_ARC))
	WmfRecordTypeEllipse               EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_ELLIPSE))
	WmfRecordTypeFloodFill             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_FLOODFILL))
	WmfRecordTypePie                   EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_PIE))
	WmfRecordTypeRectangle             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_RECTANGLE))
	WmfRecordTypeRoundRect             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_ROUNDRECT))
	WmfRecordTypePatBlt                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_PATBLT))
	WmfRecordTypeSaveDC                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SAVEDC))
	WmfRecordTypeSetPixel              EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETPIXEL))
	WmfRecordTypeOffsetClipRgn         EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_OFFSETCLIPRGN))
	WmfRecordTypeTextOut               EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_TEXTOUT))
	WmfRecordTypeBitBlt                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_BITBLT))
	WmfRecordTypeStretchBlt            EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_STRETCHBLT))
	WmfRecordTypePolygon               EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_POLYGON))
	WmfRecordTypePolyline              EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_POLYLINE))
	WmfRecordTypeEscape                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_ESCAPE))
	WmfRecordTypeRestoreDC             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_RESTOREDC))
	WmfRecordTypeFillRegion            EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_FILLREGION))
	WmfRecordTypeFrameRegion           EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_FRAMEREGION))
	WmfRecordTypeInvertRegion          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_INVERTREGION))
	WmfRecordTypePaintRegion           EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_PAINTREGION))
	WmfRecordTypeSelectClipRegion      EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SELECTCLIPREGION))
	WmfRecordTypeSelectObject          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SELECTOBJECT))
	WmfRecordTypeSetTextAlign          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETTEXTALIGN))
	WmfRecordTypeDrawText              EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x062F)) // META_DRAWTEXT
	WmfRecordTypeChord                 EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CHORD))
	WmfRecordTypeSetMapperFlags        EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETMAPPERFLAGS))
	WmfRecordTypeExtTextOut            EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_EXTTEXTOUT))
	WmfRecordTypeSetDIBToDev           EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETDIBTODEV))
	WmfRecordTypeSelectPalette         EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SELECTPALETTE))
	WmfRecordTypeRealizePalette        EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_REALIZEPALETTE))
	WmfRecordTypeAnimatePalette        EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_ANIMATEPALETTE))
	WmfRecordTypeSetPalEntries         EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_SETPALENTRIES))
	WmfRecordTypePolyPolygon           EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_POLYPOLYGON))
	WmfRecordTypeResizePalette         EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_RESIZEPALETTE))
	WmfRecordTypeDIBBitBlt             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_DIBBITBLT))
	WmfRecordTypeDIBStretchBlt         EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_DIBSTRETCHBLT))
	WmfRecordTypeDIBCreatePatternBrush EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_DIBCREATEPATTERNBRUSH))
	WmfRecordTypeStretchDIB            EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_STRETCHDIB))
	WmfRecordTypeExtFloodFill          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_EXTFLOODFILL))
	WmfRecordTypeSetLayout             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x0149)) // META_SETLAYOUT
	WmfRecordTypeResetDC               EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x014C)) // META_RESETDC
	WmfRecordTypeStartDoc              EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x014D)) // META_STARTDOC
	WmfRecordTypeStartPage             EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x004F)) // META_STARTPAGE
	WmfRecordTypeEndPage               EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x0050)) // META_ENDPAGE
	WmfRecordTypeAbortDoc              EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x0052)) // META_ABORTDOC
	WmfRecordTypeEndDoc                EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x005E)) // META_ENDDOC
	WmfRecordTypeDeleteObject          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_DELETEOBJECT))
	WmfRecordTypeCreatePalette         EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CREATEPALETTE))
	WmfRecordTypeCreateBrush           EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x00F8)) // META_CREATEBRUSH
	WmfRecordTypeCreatePatternBrush    EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CREATEPATTERNBRUSH))
	WmfRecordTypeCreatePenIndirect     EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CREATEPENINDIRECT))
	WmfRecordTypeCreateFontIndirect    EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CREATEFONTINDIRECT))
	WmfRecordTypeCreateBrushIndirect   EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CREATEBRUSHINDIRECT))
	WmfRecordTypeCreateBitmapIndirect  EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x02FD)) // META_CREATEBITMAPINDIRECT
	WmfRecordTypeCreateBitmap          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(0x06FE)) // META_CREATEBITMAP
	WmfRecordTypeCreateRegion          EmfPlusRecordType = EmfPlusRecordType(GDIP_WMF_RECORD_TO_EMFPLUS(META_CREATEREGION))

	EmfRecordTypeHeader EmfPlusRecordType = EMR_HEADER

	EmfRecordTypePolyBezier              EmfPlusRecordType = EMR_POLYBEZIER
	EmfRecordTypePolygon                 EmfPlusRecordType = EMR_POLYGON
	EmfRecordTypePolyline                EmfPlusRecordType = EMR_POLYLINE
	EmfRecordTypePolyBezierTo            EmfPlusRecordType = EMR_POLYBEZIERTO
	EmfRecordTypePolyLineTo              EmfPlusRecordType = EMR_POLYLINETO
	EmfRecordTypePolyPolyline            EmfPlusRecordType = EMR_POLYPOLYLINE
	EmfRecordTypePolyPolygon             EmfPlusRecordType = EMR_POLYPOLYGON
	EmfRecordTypeSetWindowExtEx          EmfPlusRecordType = EMR_SETWINDOWEXTEX
	EmfRecordTypeSetWindowOrgEx          EmfPlusRecordType = EMR_SETWINDOWORGEX
	EmfRecordTypeSetViewportExtEx        EmfPlusRecordType = EMR_SETVIEWPORTEXTEX
	EmfRecordTypeSetViewportOrgEx        EmfPlusRecordType = EMR_SETVIEWPORTORGEX
	EmfRecordTypeSetBrushOrgEx           EmfPlusRecordType = EMR_SETBRUSHORGEX
	EmfRecordTypeEOF                     EmfPlusRecordType = EMR_EOF
	EmfRecordTypeSetPixelV               EmfPlusRecordType = EMR_SETPIXELV
	EmfRecordTypeSetMapperFlags          EmfPlusRecordType = EMR_SETMAPPERFLAGS
	EmfRecordTypeSetMapMode              EmfPlusRecordType = EMR_SETMAPMODE
	EmfRecordTypeSetBkMode               EmfPlusRecordType = EMR_SETBKMODE
	EmfRecordTypeSetPolyFillMode         EmfPlusRecordType = EMR_SETPOLYFILLMODE
	EmfRecordTypeSetROP2                 EmfPlusRecordType = EMR_SETROP2
	EmfRecordTypeSetStretchBltMode       EmfPlusRecordType = EMR_SETSTRETCHBLTMODE
	EmfRecordTypeSetTextAlign            EmfPlusRecordType = EMR_SETTEXTALIGN
	EmfRecordTypeSetColorAdjustment      EmfPlusRecordType = EMR_SETCOLORADJUSTMENT
	EmfRecordTypeSetTextColor            EmfPlusRecordType = EMR_SETTEXTCOLOR
	EmfRecordTypeSetBkColor              EmfPlusRecordType = EMR_SETBKCOLOR
	EmfRecordTypeOffsetClipRgn           EmfPlusRecordType = EMR_OFFSETCLIPRGN
	EmfRecordTypeMoveToEx                EmfPlusRecordType = EMR_MOVETOEX
	EmfRecordTypeSetMetaRgn              EmfPlusRecordType = EMR_SETMETARGN
	EmfRecordTypeExcludeClipRect         EmfPlusRecordType = EMR_EXCLUDECLIPRECT
	EmfRecordTypeIntersectClipRect       EmfPlusRecordType = EMR_INTERSECTCLIPRECT
	EmfRecordTypeScaleViewportExtEx      EmfPlusRecordType = EMR_SCALEVIEWPORTEXTEX
	EmfRecordTypeScaleWindowExtEx        EmfPlusRecordType = EMR_SCALEWINDOWEXTEX
	EmfRecordTypeSaveDC                  EmfPlusRecordType = EMR_SAVEDC
	EmfRecordTypeRestoreDC               EmfPlusRecordType = EMR_RESTOREDC
	EmfRecordTypeSetWorldTransform       EmfPlusRecordType = EMR_SETWORLDTRANSFORM
	EmfRecordTypeModifyWorldTransform    EmfPlusRecordType = EMR_MODIFYWORLDTRANSFORM
	EmfRecordTypeSelectObject            EmfPlusRecordType = EMR_SELECTOBJECT
	EmfRecordTypeCreatePen               EmfPlusRecordType = EMR_CREATEPEN
	EmfRecordTypeCreateBrushIndirect     EmfPlusRecordType = EMR_CREATEBRUSHINDIRECT
	EmfRecordTypeDeleteObject            EmfPlusRecordType = EMR_DELETEOBJECT
	EmfRecordTypeAngleArc                EmfPlusRecordType = EMR_ANGLEARC
	EmfRecordTypeEllipse                 EmfPlusRecordType = EMR_ELLIPSE
	EmfRecordTypeRectangle               EmfPlusRecordType = EMR_RECTANGLE
	EmfRecordTypeRoundRect               EmfPlusRecordType = EMR_ROUNDRECT
	EmfRecordTypeArc                     EmfPlusRecordType = EMR_ARC
	EmfRecordTypeChord                   EmfPlusRecordType = EMR_CHORD
	EmfRecordTypePie                     EmfPlusRecordType = EMR_PIE
	EmfRecordTypeSelectPalette           EmfPlusRecordType = EMR_SELECTPALETTE
	EmfRecordTypeCreatePalette           EmfPlusRecordType = EMR_CREATEPALETTE
	EmfRecordTypeSetPaletteEntries       EmfPlusRecordType = EMR_SETPALETTEENTRIES
	EmfRecordTypeResizePalette           EmfPlusRecordType = EMR_RESIZEPALETTE
	EmfRecordTypeRealizePalette          EmfPlusRecordType = EMR_REALIZEPALETTE
	EmfRecordTypeExtFloodFill            EmfPlusRecordType = EMR_EXTFLOODFILL
	EmfRecordTypeLineTo                  EmfPlusRecordType = EMR_LINETO
	EmfRecordTypeArcTo                   EmfPlusRecordType = EMR_ARCTO
	EmfRecordTypePolyDraw                EmfPlusRecordType = EMR_POLYDRAW
	EmfRecordTypeSetArcDirection         EmfPlusRecordType = EMR_SETARCDIRECTION
	EmfRecordTypeSetMiterLimit           EmfPlusRecordType = EMR_SETMITERLIMIT
	EmfRecordTypeBeginPath               EmfPlusRecordType = EMR_BEGINPATH
	EmfRecordTypeEndPath                 EmfPlusRecordType = EMR_ENDPATH
	EmfRecordTypeCloseFigure             EmfPlusRecordType = EMR_CLOSEFIGURE
	EmfRecordTypeFillPath                EmfPlusRecordType = EMR_FILLPATH
	EmfRecordTypeStrokeAndFillPath       EmfPlusRecordType = EMR_STROKEANDFILLPATH
	EmfRecordTypeStrokePath              EmfPlusRecordType = EMR_STROKEPATH
	EmfRecordTypeFlattenPath             EmfPlusRecordType = EMR_FLATTENPATH
	EmfRecordTypeWidenPath               EmfPlusRecordType = EMR_WIDENPATH
	EmfRecordTypeSelectClipPath          EmfPlusRecordType = EMR_SELECTCLIPPATH
	EmfRecordTypeAbortPath               EmfPlusRecordType = EMR_ABORTPATH
	EmfRecordTypeReserved_069            EmfPlusRecordType = 69 // Not Used
	EmfRecordTypeGdiComment              EmfPlusRecordType = EMR_GDICOMMENT
	EmfRecordTypeFillRgn                 EmfPlusRecordType = EMR_FILLRGN
	EmfRecordTypeFrameRgn                EmfPlusRecordType = EMR_FRAMERGN
	EmfRecordTypeInvertRgn               EmfPlusRecordType = EMR_INVERTRGN
	EmfRecordTypePaintRgn                EmfPlusRecordType = EMR_PAINTRGN
	EmfRecordTypeExtSelectClipRgn        EmfPlusRecordType = EMR_EXTSELECTCLIPRGN
	EmfRecordTypeBitBlt                  EmfPlusRecordType = EMR_BITBLT
	EmfRecordTypeStretchBlt              EmfPlusRecordType = EMR_STRETCHBLT
	EmfRecordTypeMaskBlt                 EmfPlusRecordType = EMR_MASKBLT
	EmfRecordTypePlgBlt                  EmfPlusRecordType = EMR_PLGBLT
	EmfRecordTypeSetDIBitsToDevice       EmfPlusRecordType = EMR_SETDIBITSTODEVICE
	EmfRecordTypeStretchDIBits           EmfPlusRecordType = EMR_STRETCHDIBITS
	EmfRecordTypeExtCreateFontIndirect   EmfPlusRecordType = EMR_EXTCREATEFONTINDIRECTW
	EmfRecordTypeExtTextOutA             EmfPlusRecordType = EMR_EXTTEXTOUTA
	EmfRecordTypeExtTextOutW             EmfPlusRecordType = EMR_EXTTEXTOUTW
	EmfRecordTypePolyBezier16            EmfPlusRecordType = EMR_POLYBEZIER16
	EmfRecordTypePolygon16               EmfPlusRecordType = EMR_POLYGON16
	EmfRecordTypePolyline16              EmfPlusRecordType = EMR_POLYLINE16
	EmfRecordTypePolyBezierTo16          EmfPlusRecordType = EMR_POLYBEZIERTO16
	EmfRecordTypePolylineTo16            EmfPlusRecordType = EMR_POLYLINETO16
	EmfRecordTypePolyPolyline16          EmfPlusRecordType = EMR_POLYPOLYLINE16
	EmfRecordTypePolyPolygon16           EmfPlusRecordType = EMR_POLYPOLYGON16
	EmfRecordTypePolyDraw16              EmfPlusRecordType = EMR_POLYDRAW16
	EmfRecordTypeCreateMonoBrush         EmfPlusRecordType = EMR_CREATEMONOBRUSH
	EmfRecordTypeCreateDIBPatternBrushPt EmfPlusRecordType = EMR_CREATEDIBPATTERNBRUSHPT
	EmfRecordTypeExtCreatePen            EmfPlusRecordType = EMR_EXTCREATEPEN
	EmfRecordTypePolyTextOutA            EmfPlusRecordType = EMR_POLYTEXTOUTA
	EmfRecordTypePolyTextOutW            EmfPlusRecordType = EMR_POLYTEXTOUTW
	EmfRecordTypeSetICMMode              EmfPlusRecordType = 98  // EMR_SETICMMODE
	EmfRecordTypeCreateColorSpace        EmfPlusRecordType = 99  // EMR_CREATECOLORSPACE
	EmfRecordTypeSetColorSpace           EmfPlusRecordType = 100 // EMR_SETCOLORSPACE
	EmfRecordTypeDeleteColorSpace        EmfPlusRecordType = 101 // EMR_DELETECOLORSPACE
	EmfRecordTypeGLSRecord               EmfPlusRecordType = 102 // EMR_GLSRECORD
	EmfRecordTypeGLSBoundedRecord        EmfPlusRecordType = 103 // EMR_GLSBOUNDEDRECORD
	EmfRecordTypePixelFormat             EmfPlusRecordType = 104 // EMR_PIXELFORMAT
	EmfRecordTypeDrawEscape              EmfPlusRecordType = 105 // EMR_RESERVED_105
	EmfRecordTypeExtEscape               EmfPlusRecordType = 106 // EMR_RESERVED_106
	EmfRecordTypeStartDoc                EmfPlusRecordType = 107 // EMR_RESERVED_107
	EmfRecordTypeSmallTextOut            EmfPlusRecordType = 108 // EMR_RESERVED_108
	EmfRecordTypeForceUFIMapping         EmfPlusRecordType = 109 // EMR_RESERVED_109
	EmfRecordTypeNamedEscape             EmfPlusRecordType = 110 // EMR_RESERVED_110
	EmfRecordTypeColorCorrectPalette     EmfPlusRecordType = 111 // EMR_COLORCORRECTPALETTE
	EmfRecordTypeSetICMProfileA          EmfPlusRecordType = 112 // EMR_SETICMPROFILEA
	EmfRecordTypeSetICMProfileW          EmfPlusRecordType = 113 // EMR_SETICMPROFILEW
	EmfRecordTypeAlphaBlend              EmfPlusRecordType = 114 // EMR_ALPHABLEND
	EmfRecordTypeSetLayout               EmfPlusRecordType = 115 // EMR_SETLAYOUT
	EmfRecordTypeTransparentBlt          EmfPlusRecordType = 116 // EMR_TRANSPARENTBLT
	EmfRecordTypeReserved_117            EmfPlusRecordType = 117 // Not Used
	EmfRecordTypeGradientFill            EmfPlusRecordType = 118 // EMR_GRADIENTFILL
	EmfRecordTypeSetLinkedUFIs           EmfPlusRecordType = 119 // EMR_RESERVED_119
	EmfRecordTypeSetTextJustification    EmfPlusRecordType = 120 // EMR_RESERVED_120
	EmfRecordTypeColorMatchToTargetW     EmfPlusRecordType = 121 // EMR_COLORMATCHTOTARGETW
	EmfRecordTypeCreateColorSpaceW       EmfPlusRecordType = 122 // EMR_CREATECOLORSPACEW
	EmfRecordTypeMax                     EmfPlusRecordType = 122
	EmfRecordTypeMin                     EmfPlusRecordType = 1

	// That is the END of the GDI EMF records.

	// Now we start the list of EMF+ records.  We leave quite
	// a bit of room here for the addition of any new GDI
	// records that may be added later.
)

const (
	EmfPlusRecordTypeInvalid   EmfPlusRecordType = EmfPlusRecordType(GDIP_EMFPLUS_RECORD_BASE)
	EmfPlusRecordTypeHeader    EmfPlusRecordType = EmfPlusRecordTypeInvalid + 1
	EmfPlusRecordTypeEndOfFile EmfPlusRecordType = EmfPlusRecordTypeInvalid + 2

	EmfPlusRecordTypeComment EmfPlusRecordType = EmfPlusRecordTypeInvalid + 3

	EmfPlusRecordTypeGetDC EmfPlusRecordType = EmfPlusRecordTypeInvalid + 4 // the application grabbed the metafile dc

	EmfPlusRecordTypeMultiFormatStart   EmfPlusRecordType = EmfPlusRecordTypeInvalid + 5
	EmfPlusRecordTypeMultiFormatSection EmfPlusRecordType = EmfPlusRecordTypeInvalid + 6
	EmfPlusRecordTypeMultiFormatEnd     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 7

	// For all persistent objects
	EmfPlusRecordTypeObject EmfPlusRecordType = EmfPlusRecordTypeInvalid + 8 // brushpenpathregionimagefontstring-format

	// Drawing Records
	EmfPlusRecordTypeClear           EmfPlusRecordType = EmfPlusRecordTypeInvalid + 9
	EmfPlusRecordTypeFillRects       EmfPlusRecordType = EmfPlusRecordTypeInvalid + 10
	EmfPlusRecordTypeDrawRects       EmfPlusRecordType = EmfPlusRecordTypeInvalid + 11
	EmfPlusRecordTypeFillPolygon     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 12
	EmfPlusRecordTypeDrawLines       EmfPlusRecordType = EmfPlusRecordTypeInvalid + 13
	EmfPlusRecordTypeFillEllipse     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 14
	EmfPlusRecordTypeDrawEllipse     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 15
	EmfPlusRecordTypeFillPie         EmfPlusRecordType = EmfPlusRecordTypeInvalid + 16
	EmfPlusRecordTypeDrawPie         EmfPlusRecordType = EmfPlusRecordTypeInvalid + 17
	EmfPlusRecordTypeDrawArc         EmfPlusRecordType = EmfPlusRecordTypeInvalid + 18
	EmfPlusRecordTypeFillRegion      EmfPlusRecordType = EmfPlusRecordTypeInvalid + 19
	EmfPlusRecordTypeFillPath        EmfPlusRecordType = EmfPlusRecordTypeInvalid + 20
	EmfPlusRecordTypeDrawPath        EmfPlusRecordType = EmfPlusRecordTypeInvalid + 21
	EmfPlusRecordTypeFillClosedCurve EmfPlusRecordType = EmfPlusRecordTypeInvalid + 22
	EmfPlusRecordTypeDrawClosedCurve EmfPlusRecordType = EmfPlusRecordTypeInvalid + 23
	EmfPlusRecordTypeDrawCurve       EmfPlusRecordType = EmfPlusRecordTypeInvalid + 24
	EmfPlusRecordTypeDrawBeziers     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 25
	EmfPlusRecordTypeDrawImage       EmfPlusRecordType = EmfPlusRecordTypeInvalid + 26
	EmfPlusRecordTypeDrawImagePoints EmfPlusRecordType = EmfPlusRecordTypeInvalid + 27
	EmfPlusRecordTypeDrawString      EmfPlusRecordType = EmfPlusRecordTypeInvalid + 28

	// Graphics State Records
	EmfPlusRecordTypeSetRenderingOrigin   EmfPlusRecordType = EmfPlusRecordTypeInvalid + 29
	EmfPlusRecordTypeSetAntiAliasMode     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 30
	EmfPlusRecordTypeSetTextRenderingHint EmfPlusRecordType = EmfPlusRecordTypeInvalid + 31

	//#ifdef DCR_USE_NEW_188922
	EmfPlusRecordTypeSetTextContrast EmfPlusRecordType = EmfPlusRecordTypeInvalid + 32

	//#else
	//   EmfPlusRecordTypeSetGammaValue EmfPlusRecordType = EmfPlusRecordTypeInvalid + 32
	//#endif // DCR_USE_NEW_188922

	EmfPlusRecordTypeSetInterpolationMode   EmfPlusRecordType = EmfPlusRecordTypeInvalid + 33
	EmfPlusRecordTypeSetPixelOffsetMode     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 34
	EmfPlusRecordTypeSetCompositingMode     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 35
	EmfPlusRecordTypeSetCompositingQuality  EmfPlusRecordType = EmfPlusRecordTypeInvalid + 36
	EmfPlusRecordTypeSave                   EmfPlusRecordType = EmfPlusRecordTypeInvalid + 37
	EmfPlusRecordTypeRestore                EmfPlusRecordType = EmfPlusRecordTypeInvalid + 38
	EmfPlusRecordTypeBeginContainer         EmfPlusRecordType = EmfPlusRecordTypeInvalid + 39
	EmfPlusRecordTypeBeginContainerNoParams EmfPlusRecordType = EmfPlusRecordTypeInvalid + 40

	EmfPlusRecordTypeEndContainer            EmfPlusRecordType = EmfPlusRecordTypeInvalid + 41
	EmfPlusRecordTypeSetWorldTransform       EmfPlusRecordType = EmfPlusRecordTypeInvalid + 42
	EmfPlusRecordTypeResetWorldTransform     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 43
	EmfPlusRecordTypeMultiplyWorldTransform  EmfPlusRecordType = EmfPlusRecordTypeInvalid + 44
	EmfPlusRecordTypeTranslateWorldTransform EmfPlusRecordType = EmfPlusRecordTypeInvalid + 45
	EmfPlusRecordTypeScaleWorldTransform     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 46
	EmfPlusRecordTypeRotateWorldTransform    EmfPlusRecordType = EmfPlusRecordTypeInvalid + 47
	EmfPlusRecordTypeSetPageTransform        EmfPlusRecordType = EmfPlusRecordTypeInvalid + 48
	EmfPlusRecordTypeResetClip               EmfPlusRecordType = EmfPlusRecordTypeInvalid + 49
	EmfPlusRecordTypeSetClipRect             EmfPlusRecordType = EmfPlusRecordTypeInvalid + 50
	EmfPlusRecordTypeSetClipPath             EmfPlusRecordType = EmfPlusRecordTypeInvalid + 51
	EmfPlusRecordTypeSetClipRegion           EmfPlusRecordType = EmfPlusRecordTypeInvalid + 52
	EmfPlusRecordTypeOffsetClip              EmfPlusRecordType = EmfPlusRecordTypeInvalid + 53

	// New record types must be added here (at the end) -- do not add above
	// since that will invalidate previous metafiles!
	EmfPlusRecordTypeDrawDriverString EmfPlusRecordType = EmfPlusRecordTypeInvalid + 54

	//#if (GDIPVER >= 0x0110)
	EmfPlusRecordTypeStrokeFillPath     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 55
	EmfPlusRecordTypeSerializableObject EmfPlusRecordType = EmfPlusRecordTypeInvalid + 56

	EmfPlusRecordTypeSetTSGraphics EmfPlusRecordType = EmfPlusRecordTypeInvalid + 57
	EmfPlusRecordTypeSetTSClip     EmfPlusRecordType = EmfPlusRecordTypeInvalid + 58
	//#endif

	// Have this here so you don't need to keep changing the value of
	// EmfPlusRecordTypeMax every time you add a new record.

	EmfPlusRecordTotal EmfPlusRecordType = EmfPlusRecordTypeInvalid + 59

	EmfPlusRecordTypeMax = EmfPlusRecordTotal - 1
	EmfPlusRecordTypeMin = EmfPlusRecordTypeHeader
)

type StringFormatFlags INT

const (
	StringFormatFlagsDirectionRightToLeft StringFormatFlags = 0x00000001
	StringFormatFlagsDirectionVertical    StringFormatFlags = 0x00000002
	StringFormatFlagsNoFitBlackBox        StringFormatFlags = 0x00000004
	//#ifndef DCR_USE_NEW_137252
	StringFormatFlagsNumberContextArabic         StringFormatFlags = 0x00000008
	StringFormatFlagsDisableKashidaJustification StringFormatFlags = 0x00000010
	//#endif
	StringFormatFlagsDisplayFormatControl StringFormatFlags = 0x00000020
	//#ifndef DCR_USE_NEW_137252
	StringFormatFlagsDisableKerning   StringFormatFlags = 0x00000040
	StringFormatFlagsDisableLigatures StringFormatFlags = 0x00000080
	StringFormatFlagsLayoutLegacyBidi StringFormatFlags = 0x00000100
	StringFormatFlagsNoChanges        StringFormatFlags = 0x00000200
	//#endif
	StringFormatFlagsNoFontFallback        StringFormatFlags = 0x00000400
	StringFormatFlagsMeasureTrailingSpaces StringFormatFlags = 0x00000800
	StringFormatFlagsNoWrap                StringFormatFlags = 0x00001000
	StringFormatFlagsLineLimit             StringFormatFlags = 0x00002000

	StringFormatFlagsNoClip StringFormatFlags = 0x00004000
)

//---------------------------------------------------------------------------
// StringTrimming
//---------------------------------------------------------------------------
type StringTrimming INT

const (
	StringTrimmingNone              StringTrimming = 0
	StringTrimmingCharacter         StringTrimming = 1
	StringTrimmingWord              StringTrimming = 2
	StringTrimmingEllipsisCharacter StringTrimming = 3
	StringTrimmingEllipsisWord      StringTrimming = 4
	StringTrimmingEllipsisPath      StringTrimming = 5
)

//---------------------------------------------------------------------------
// String units
//
// String units are like length units in CSS, they may be absolute, or
// they may be relative to a font size.
//
//---------------------------------------------------------------------------

type StringUnit INT

const (
	StringUnitWorld      StringUnit = StringUnit(UnitWorld)
	StringUnitDisplay    StringUnit = StringUnit(UnitDisplay)
	StringUnitPixel      StringUnit = StringUnit(UnitPixel)
	StringUnitPoint      StringUnit = StringUnit(UnitPoint)
	StringUnitInch       StringUnit = StringUnit(UnitInch)
	StringUnitDocument   StringUnit = StringUnit(UnitDocument)
	StringUnitMillimeter StringUnit = StringUnit(UnitMillimeter)
	StringUnitEm         StringUnit = 32
)

//---------------------------------------------------------------------------
// Line spacing flags
//---------------------------------------------------------------------------
type LineSpacing INT

const (
	LineSpacingWorld      LineSpacing = LineSpacing(UnitWorld)
	LineSpacingDisplay    LineSpacing = LineSpacing(UnitDisplay)
	LineSpacingPixel      LineSpacing = LineSpacing(UnitPixel)
	LineSpacingPoint      LineSpacing = LineSpacing(UnitPoint)
	LineSpacingInch       LineSpacing = LineSpacing(UnitInch)
	LineSpacingDocument   LineSpacing = LineSpacing(UnitDocument)
	LineSpacingMillimeter LineSpacing = LineSpacing(UnitMillimeter)

	LineSpacingRecommended         = 32
	LineSpacingAtLeast             = 33
	LineSpacingAtLeastMultiple     = 34
	LineSpacingCell                = 35
	LineSpacingCellAtLeast         = 36
	LineSpacingCellAtLeastMultiple = 37
)

//---------------------------------------------------------------------------
// National language digit substitution
//---------------------------------------------------------------------------
type StringDigitSubstitute INT

const (
	StringDigitSubstituteUser        StringDigitSubstitute = 0 // As NLS setting
	StringDigitSubstituteNone        StringDigitSubstitute = 1
	StringDigitSubstituteNational    StringDigitSubstitute = 2
	StringDigitSubstituteTraditional StringDigitSubstitute = 3
)

//---------------------------------------------------------------------------
// Hotkey prefix interpretation
//---------------------------------------------------------------------------
type HotkeyPrefix INT

const (
	HotkeyPrefixNone HotkeyPrefix = 0
	HotkeyPrefixShow HotkeyPrefix = 1
	HotkeyPrefixHide HotkeyPrefix = 2
)

//---------------------------------------------------------------------------
// Text alignment flags
//---------------------------------------------------------------------------
type StringAlignment INT

const (
	// Left edge for left-to-right text,
	// right for right-to-left text,
	// and top for vertical
	StringAlignmentNear   StringAlignment = 0
	StringAlignmentCenter StringAlignment = 1
	StringAlignmentFar    StringAlignment = 2
)

//---------------------------------------------------------------------------
// DriverStringOptions
//---------------------------------------------------------------------------
type DriverStringOptions INT

const (
	DriverStringOptionsCmapLookup      DriverStringOptions = 1
	DriverStringOptionsVertical        DriverStringOptions = 2
	DriverStringOptionsRealizedAdvance DriverStringOptions = 4
	//#ifndef DCR_USE_NEW_137252
	DriverStringOptionsCompensateResolution DriverStringOptions = 8

//#endif
)

//---------------------------------------------------------------------------
// Flush Intention flags
//---------------------------------------------------------------------------
type FlushIntention INT

const (
	FlushIntentionFlush FlushIntention = 0 // Flush all batched rendering operations
	FlushIntentionSync  FlushIntention = 1 // Flush all batched rendering operations
	// and wait for them to complete
)

//---------------------------------------------------------------------------
// Window Change Notification types
//---------------------------------------------------------------------------
type WindowNotifyEnum INT

const (
	WindowNotifyEnumEnable WindowNotifyEnum = iota
	WindowNotifyEnumDisable
	WindowNotifyEnumPalette
	WindowNotifyEnumDisplay
	WindowNotifyEnumSysColor
)

//---------------------------------------------------------------------------
// Image encoder parameter related types
//---------------------------------------------------------------------------
type EncoderParameterValueType INT

const (
	EncoderParameterValueTypeByte          EncoderParameterValueType = 1 // 8-bit unsigned int
	EncoderParameterValueTypeASCII         EncoderParameterValueType = 2 // 8-bit byte containing one 7-bit ASCII // code. NULL terminated.
	EncoderParameterValueTypeShort         EncoderParameterValueType = 3 // 16-bit unsigned int
	EncoderParameterValueTypeLong          EncoderParameterValueType = 4 // 32-bit unsigned int
	EncoderParameterValueTypeRational      EncoderParameterValueType = 5 // Two Longs. The first Long is the	// numerator, the second Long expresses the	// denomintor.
	EncoderParameterValueTypeLongRange     EncoderParameterValueType = 6 // Two longs which specify a range of	// integer values. The first Long specifies	// the lower end and the second one	// specifies the higher end. All values	// are inclusive at both ends
	EncoderParameterValueTypeUndefined     EncoderParameterValueType = 7 // 8-bit byte that can take any value	// depending on field definition
	EncoderParameterValueTypeRationalRange EncoderParameterValueType = 8 // Two Rationals. The first Rational	// specifies the lower end and the second	// specifies the higher end. All values	// are inclusive at both ends

	//#if (GDIPVER >= 0x0110)
	EncoderParameterValueTypePointer EncoderParameterValueType = 9 // a pointer to a parameter defined data.
//#endif //(GDIPVER >= 0x0110)
)

type ValueType INT

const (
	ValueTypeByte          ValueType = 1 // 8-bit unsigned int
	ValueTypeASCII         ValueType = 2 // 8-bit byte containing one 7-bit ASCII // code. NULL terminated.
	ValueTypeShort         ValueType = 3 // 16-bit unsigned int
	ValueTypeLong          ValueType = 4 // 32-bit unsigned int
	ValueTypeRational      ValueType = 5 // Two Longs. The first Long is the	// numerator, the second Long expresses the	// denomintor.
	ValueTypeLongRange     ValueType = 6 // Two longs which specify a range of	// integer values. The first Long specifies	// the lower end and the second one	// specifies the higher end. All values	// are inclusive at both ends
	ValueTypeUndefined     ValueType = 7 // 8-bit byte that can take any value	// depending on field definition
	ValueTypeRationalRange ValueType = 8 // Two Rationals. The first Rational	// specifies the lower end and the second	// specifies the higher end. All values	// are inclusive at both ends
)

//---------------------------------------------------------------------------
// Image encoder value types
//---------------------------------------------------------------------------
type EncoderValue INT

const (
	EncoderValueColorTypeCMYK EncoderValue = iota
	EncoderValueColorTypeYCCK
	EncoderValueCompressionLZW
	EncoderValueCompressionCCITT3
	EncoderValueCompressionCCITT4
	EncoderValueCompressionRle
	EncoderValueCompressionNone
	EncoderValueScanMethodInterlaced
	EncoderValueScanMethodNonInterlaced
	EncoderValueVersionGif87
	EncoderValueVersionGif89
	EncoderValueRenderProgressive
	EncoderValueRenderNonProgressive
	EncoderValueTransformRotate90
	EncoderValueTransformRotate180
	EncoderValueTransformRotate270
	EncoderValueTransformFlipHorizontal
	EncoderValueTransformFlipVertical
	//#ifdef DCR_USE_NEW_140861
	EncoderValueMultiFrame
	//#else
	EncodeValueMultiFrame
	//#endif
	EncoderValueLastFrame
	EncoderValueFlush
	//#ifdef DCR_USE_NEW_140861
	EncoderValueFrameDimensionTime
	EncoderValueFrameDimensionResolution
	EncoderValueFrameDimensionPage
	//#else
	EncodeValueFrameDimensionTime
	EncodeValueFrameDimensionResolution
	EncodeValueFrameDimensionPage
	//#endif

	//#if (GDIPVER >= 0x0110)
	EncoderValueColorTypeGray
	EncoderValueColorTypeRGB

//#endif
)

//---------------------------------------------------------------------------
// Graphics layout values (support for Middle East localization)
//---------------------------------------------------------------------------
type GraphicsLayout INT

const (
	GraphicsLayoutNormal GraphicsLayout = iota
	GraphicsLayoutMirrored
	GraphicsLayoutMirroredIgnoreImages
	GraphicsLayoutMirroredForceImages
)

//---------------------------------------------------------------------------
// Image layout values (support for Middle East localization)
//---------------------------------------------------------------------------
type ImageLayout INT

const (
	ImageLayoutNormal ImageLayout = iota
	ImageLayoutIgnoreMirrored
)

type EmfToWmfBitsFlags INT

const (
	EmfToWmfBitsFlagsDefault    EmfToWmfBitsFlags = 0x00000000
	EmfToWmfBitsFlagsEmbedEmf   EmfToWmfBitsFlags = 0x00000001
	EmfToWmfBitsFlagsIncludeAPM EmfToWmfBitsFlags = 0x00000002
	EmfToWmfBitsFlagsNoXORClip  EmfToWmfBitsFlags = 0x00000004
)

//---------------------------------------------------------------------------
// Conversion of Emf To Emf+ Bits flags
//---------------------------------------------------------------------------
//#if (GDIPVER >= 0x0110)
type ConvertToEmfPlusFlags INT

const (
	ConvertToEmfPlusFlagsDefault       ConvertToEmfPlusFlags = 0x00000000
	ConvertToEmfPlusFlagsRopUsed       ConvertToEmfPlusFlags = 0x00000001
	ConvertToEmfPlusFlagsText          ConvertToEmfPlusFlags = 0x00000002
	ConvertToEmfPlusFlagsInvalidRecord ConvertToEmfPlusFlags = 0x00000004
)

//#endif

//---------------------------------------------------------------------------
// Test Control flags
//---------------------------------------------------------------------------
type GpTestControlEnum INT

const (
	TestControlForceBilinear  GpTestControlEnum = 0
	TestControlNoICM          GpTestControlEnum = 1
	TestControlGetBuildNumber GpTestControlEnum = 2
)
