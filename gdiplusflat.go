package gdiplus

import (
	"errors"
	"os"
	"path"
	"strconv"
	"syscall"
	"unsafe"
)
import (
	. "github.com/tryor/winapi"
	. "github.com/tryor/winapi/gdi"
)

var (
	modgdiplus = func() *syscall.LazyDLL {
		name := "gdiplus.dll"
		appPath, _ := os.Getwd()
		filename := path.Join(appPath, "", name)
		_, err := os.Stat(filename)
		if err == nil {
			return syscall.NewLazyDLL(filename)
		}
		return syscall.NewLazyDLL(name)
	}()

	procGdipSetTextGammaValue                            = modgdiplus.NewProc("GdipSetTextGammaValue")
	procGdipGetTextGammaValue                            = modgdiplus.NewProc("GdipGetTextGammaValue")
	procGdipAddPathArc                                   = modgdiplus.NewProc("GdipAddPathArc")
	procGdipAddPathArcI                                  = modgdiplus.NewProc("GdipAddPathArcI")
	procGdipAddPathBezier                                = modgdiplus.NewProc("GdipAddPathBezier")
	procGdipAddPathBezierI                               = modgdiplus.NewProc("GdipAddPathBezierI")
	procGdipAddPathBeziers                               = modgdiplus.NewProc("GdipAddPathBeziers")
	procGdipAddPathBeziersI                              = modgdiplus.NewProc("GdipAddPathBeziersI")
	procGdipAddPathClosedCurve2                          = modgdiplus.NewProc("GdipAddPathClosedCurve2")
	procGdipAddPathClosedCurve2I                         = modgdiplus.NewProc("GdipAddPathClosedCurve2I")
	procGdipAddPathClosedCurve                           = modgdiplus.NewProc("GdipAddPathClosedCurve")
	procGdipAddPathClosedCurveI                          = modgdiplus.NewProc("GdipAddPathClosedCurveI")
	procGdipAddPathCurve2                                = modgdiplus.NewProc("GdipAddPathCurve2")
	procGdipAddPathCurve2I                               = modgdiplus.NewProc("GdipAddPathCurve2I")
	procGdipAddPathCurve3                                = modgdiplus.NewProc("GdipAddPathCurve3")
	procGdipAddPathCurve3I                               = modgdiplus.NewProc("GdipAddPathCurve3I")
	procGdipAddPathCurve                                 = modgdiplus.NewProc("GdipAddPathCurve")
	procGdipAddPathCurveI                                = modgdiplus.NewProc("GdipAddPathCurveI")
	procGdipAddPathEllipse                               = modgdiplus.NewProc("GdipAddPathEllipse")
	procGdipAddPathEllipseI                              = modgdiplus.NewProc("GdipAddPathEllipseI")
	procGdipAddPathLine2                                 = modgdiplus.NewProc("GdipAddPathLine2")
	procGdipAddPathLine2I                                = modgdiplus.NewProc("GdipAddPathLine2I")
	procGdipAddPathLine                                  = modgdiplus.NewProc("GdipAddPathLine")
	procGdipAddPathLineI                                 = modgdiplus.NewProc("GdipAddPathLineI")
	procGdipAddPathPath                                  = modgdiplus.NewProc("GdipAddPathPath")
	procGdipAddPathPie                                   = modgdiplus.NewProc("GdipAddPathPie")
	procGdipAddPathPieI                                  = modgdiplus.NewProc("GdipAddPathPieI")
	procGdipAddPathPolygon                               = modgdiplus.NewProc("GdipAddPathPolygon")
	procGdipAddPathPolygonI                              = modgdiplus.NewProc("GdipAddPathPolygonI")
	procGdipAddPathRectangle                             = modgdiplus.NewProc("GdipAddPathRectangle")
	procGdipAddPathRectangleI                            = modgdiplus.NewProc("GdipAddPathRectangleI")
	procGdipAddPathRectangles                            = modgdiplus.NewProc("GdipAddPathRectangles")
	procGdipAddPathRectanglesI                           = modgdiplus.NewProc("GdipAddPathRectanglesI")
	procGdipAddPathString                                = modgdiplus.NewProc("GdipAddPathString")
	procGdipAddPathStringI                               = modgdiplus.NewProc("GdipAddPathStringI")
	procGdipAlloc                                        = modgdiplus.NewProc("GdipAlloc")
	procGdipBeginContainer2                              = modgdiplus.NewProc("GdipBeginContainer2")
	procGdipBeginContainer                               = modgdiplus.NewProc("GdipBeginContainer")
	procGdipBeginContainerI                              = modgdiplus.NewProc("GdipBeginContainerI")
	procGdipBitmapGetPixel                               = modgdiplus.NewProc("GdipBitmapGetPixel")
	procGdipBitmapLockBits                               = modgdiplus.NewProc("GdipBitmapLockBits")
	procGdipBitmapSetPixel                               = modgdiplus.NewProc("GdipBitmapSetPixel")
	procGdipBitmapSetResolution                          = modgdiplus.NewProc("GdipBitmapSetResolution")
	procGdipBitmapUnlockBits                             = modgdiplus.NewProc("GdipBitmapUnlockBits")
	procGdipClearPathMarkers                             = modgdiplus.NewProc("GdipClearPathMarkers")
	procGdipCloneBitmapArea                              = modgdiplus.NewProc("GdipCloneBitmapArea")
	procGdipCloneBitmapAreaI                             = modgdiplus.NewProc("GdipCloneBitmapAreaI")
	procGdipCloneBrush                                   = modgdiplus.NewProc("GdipCloneBrush")
	procGdipCloneCustomLineCap                           = modgdiplus.NewProc("GdipCloneCustomLineCap")
	procGdipCloneFont                                    = modgdiplus.NewProc("GdipCloneFont")
	procGdipCloneFontFamily                              = modgdiplus.NewProc("GdipCloneFontFamily")
	procGdipCloneImage                                   = modgdiplus.NewProc("GdipCloneImage")
	procGdipCloneImageAttributes                         = modgdiplus.NewProc("GdipCloneImageAttributes")
	procGdipCloneMatrix                                  = modgdiplus.NewProc("GdipCloneMatrix")
	procGdipClonePath                                    = modgdiplus.NewProc("GdipClonePath")
	procGdipClonePen                                     = modgdiplus.NewProc("GdipClonePen")
	procGdipCloneRegion                                  = modgdiplus.NewProc("GdipCloneRegion")
	procGdipCloneStringFormat                            = modgdiplus.NewProc("GdipCloneStringFormat")
	procGdipClosePathFigure                              = modgdiplus.NewProc("GdipClosePathFigure")
	procGdipClosePathFigures                             = modgdiplus.NewProc("GdipClosePathFigures")
	procGdipCombineRegionPath                            = modgdiplus.NewProc("GdipCombineRegionPath")
	procGdipCombineRegionRect                            = modgdiplus.NewProc("GdipCombineRegionRect")
	procGdipCombineRegionRectI                           = modgdiplus.NewProc("GdipCombineRegionRectI")
	procGdipCombineRegionRegion                          = modgdiplus.NewProc("GdipCombineRegionRegion")
	procGdipComment                                      = modgdiplus.NewProc("GdipComment")
	procGdipCreateAdjustableArrowCap                     = modgdiplus.NewProc("GdipCreateAdjustableArrowCap")
	procGdipCreateBitmapFromDirectDrawSurface            = modgdiplus.NewProc("GdipCreateBitmapFromDirectDrawSurface")
	procGdipCreateBitmapFromFile                         = modgdiplus.NewProc("GdipCreateBitmapFromFile")
	procGdipCreateBitmapFromFileICM                      = modgdiplus.NewProc("GdipCreateBitmapFromFileICM")
	procGdipCreateBitmapFromGdiDib                       = modgdiplus.NewProc("GdipCreateBitmapFromGdiDib")
	procGdipCreateBitmapFromGraphics                     = modgdiplus.NewProc("GdipCreateBitmapFromGraphics")
	procGdipCreateBitmapFromHBITMAP                      = modgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	procGdipCreateBitmapFromHICON                        = modgdiplus.NewProc("GdipCreateBitmapFromHICON")
	procGdipCreateBitmapFromResource                     = modgdiplus.NewProc("GdipCreateBitmapFromResource")
	procGdipCreateBitmapFromScan0                        = modgdiplus.NewProc("GdipCreateBitmapFromScan0")
	procGdipCreateBitmapFromStream                       = modgdiplus.NewProc("GdipCreateBitmapFromStream")
	procGdipCreateBitmapFromStreamICM                    = modgdiplus.NewProc("GdipCreateBitmapFromStreamICM")
	procGdipCreateCachedBitmap                           = modgdiplus.NewProc("GdipCreateCachedBitmap")
	procGdipCreateCustomLineCap                          = modgdiplus.NewProc("GdipCreateCustomLineCap")
	procGdipCreateFont                                   = modgdiplus.NewProc("GdipCreateFont")
	procGdipCreateFontFamilyFromName                     = modgdiplus.NewProc("GdipCreateFontFamilyFromName")
	procGdipCreateFontFromDC                             = modgdiplus.NewProc("GdipCreateFontFromDC")
	procGdipCreateFontFromLogfontA                       = modgdiplus.NewProc("GdipCreateFontFromLogfontA")
	procGdipCreateFontFromLogfontW                       = modgdiplus.NewProc("GdipCreateFontFromLogfontW")
	procGdipCreateFromHDC2                               = modgdiplus.NewProc("GdipCreateFromHDC2")
	procGdipCreateFromHDC                                = modgdiplus.NewProc("GdipCreateFromHDC")
	procGdipCreateFromHWND                               = modgdiplus.NewProc("GdipCreateFromHWND")
	procGdipCreateFromHWNDICM                            = modgdiplus.NewProc("GdipCreateFromHWNDICM")
	procGdipCreateHBITMAPFromBitmap                      = modgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")
	procGdipCreateHICONFromBitmap                        = modgdiplus.NewProc("GdipCreateHICONFromBitmap")
	procGdipCreateHalftonePalette                        = modgdiplus.NewProc("GdipCreateHalftonePalette")
	procGdipCreateHatchBrush                             = modgdiplus.NewProc("GdipCreateHatchBrush")
	procGdipCreateImageAttributes                        = modgdiplus.NewProc("GdipCreateImageAttributes")
	procGdipCreateLineBrush                              = modgdiplus.NewProc("GdipCreateLineBrush")
	procGdipCreateLineBrushFromRect                      = modgdiplus.NewProc("GdipCreateLineBrushFromRect")
	procGdipCreateLineBrushFromRectI                     = modgdiplus.NewProc("GdipCreateLineBrushFromRectI")
	procGdipCreateLineBrushFromRectWithAngle             = modgdiplus.NewProc("GdipCreateLineBrushFromRectWithAngle")
	procGdipCreateLineBrushFromRectWithAngleI            = modgdiplus.NewProc("GdipCreateLineBrushFromRectWithAngleI")
	procGdipCreateLineBrushI                             = modgdiplus.NewProc("GdipCreateLineBrushI")
	procGdipCreateMatrix2                                = modgdiplus.NewProc("GdipCreateMatrix2")
	procGdipCreateMatrix3                                = modgdiplus.NewProc("GdipCreateMatrix3")
	procGdipCreateMatrix3I                               = modgdiplus.NewProc("GdipCreateMatrix3I")
	procGdipCreateMatrix                                 = modgdiplus.NewProc("GdipCreateMatrix")
	procGdipCreateMetafileFromEmf                        = modgdiplus.NewProc("GdipCreateMetafileFromEmf")
	procGdipCreateMetafileFromFile                       = modgdiplus.NewProc("GdipCreateMetafileFromFile")
	procGdipCreateMetafileFromStream                     = modgdiplus.NewProc("GdipCreateMetafileFromStream")
	procGdipCreateMetafileFromWmf                        = modgdiplus.NewProc("GdipCreateMetafileFromWmf")
	procGdipCreateMetafileFromWmfFile                    = modgdiplus.NewProc("GdipCreateMetafileFromWmfFile")
	procGdipCreatePath2                                  = modgdiplus.NewProc("GdipCreatePath2")
	procGdipCreatePath2I                                 = modgdiplus.NewProc("GdipCreatePath2I")
	procGdipCreatePath                                   = modgdiplus.NewProc("GdipCreatePath")
	procGdipCreatePathGradient                           = modgdiplus.NewProc("GdipCreatePathGradient")
	procGdipCreatePathGradientFromPath                   = modgdiplus.NewProc("GdipCreatePathGradientFromPath")
	procGdipCreatePathGradientI                          = modgdiplus.NewProc("GdipCreatePathGradientI")
	procGdipCreatePathIter                               = modgdiplus.NewProc("GdipCreatePathIter")
	procGdipCreatePen1                                   = modgdiplus.NewProc("GdipCreatePen1")
	procGdipCreatePen2                                   = modgdiplus.NewProc("GdipCreatePen2")
	procGdipCreateRegion                                 = modgdiplus.NewProc("GdipCreateRegion")
	procGdipCreateRegionHrgn                             = modgdiplus.NewProc("GdipCreateRegionHrgn")
	procGdipCreateRegionPath                             = modgdiplus.NewProc("GdipCreateRegionPath")
	procGdipCreateRegionRect                             = modgdiplus.NewProc("GdipCreateRegionRect")
	procGdipCreateRegionRectI                            = modgdiplus.NewProc("GdipCreateRegionRectI")
	procGdipCreateRegionRgnData                          = modgdiplus.NewProc("GdipCreateRegionRgnData")
	procGdipCreateSolidFill                              = modgdiplus.NewProc("GdipCreateSolidFill")
	procGdipCreateStreamOnFile                           = modgdiplus.NewProc("GdipCreateStreamOnFile")
	procGdipCreateStringFormat                           = modgdiplus.NewProc("GdipCreateStringFormat")
	procGdipCreateTexture2                               = modgdiplus.NewProc("GdipCreateTexture2")
	procGdipCreateTexture2I                              = modgdiplus.NewProc("GdipCreateTexture2I")
	procGdipCreateTexture                                = modgdiplus.NewProc("GdipCreateTexture")
	procGdipCreateTextureIA                              = modgdiplus.NewProc("GdipCreateTextureIA")
	procGdipCreateTextureIAI                             = modgdiplus.NewProc("GdipCreateTextureIAI")
	procGdipDeleteBrush                                  = modgdiplus.NewProc("GdipDeleteBrush")
	procGdipDeleteCachedBitmap                           = modgdiplus.NewProc("GdipDeleteCachedBitmap")
	procGdipDeleteCustomLineCap                          = modgdiplus.NewProc("GdipDeleteCustomLineCap")
	procGdipDeleteFont                                   = modgdiplus.NewProc("GdipDeleteFont")
	procGdipDeleteFontFamily                             = modgdiplus.NewProc("GdipDeleteFontFamily")
	procGdipDeleteGraphics                               = modgdiplus.NewProc("GdipDeleteGraphics")
	procGdipDeleteMatrix                                 = modgdiplus.NewProc("GdipDeleteMatrix")
	procGdipDeletePath                                   = modgdiplus.NewProc("GdipDeletePath")
	procGdipDeletePathIter                               = modgdiplus.NewProc("GdipDeletePathIter")
	procGdipDeletePen                                    = modgdiplus.NewProc("GdipDeletePen")
	procGdipDeletePrivateFontCollection                  = modgdiplus.NewProc("GdipDeletePrivateFontCollection")
	procGdipDeleteRegion                                 = modgdiplus.NewProc("GdipDeleteRegion")
	procGdipDeleteStringFormat                           = modgdiplus.NewProc("GdipDeleteStringFormat")
	procGdipDisposeImage                                 = modgdiplus.NewProc("GdipDisposeImage")
	procGdipDisposeImageAttributes                       = modgdiplus.NewProc("GdipDisposeImageAttributes")
	procGdipDrawArc                                      = modgdiplus.NewProc("GdipDrawArc")
	procGdipDrawArcI                                     = modgdiplus.NewProc("GdipDrawArcI")
	procGdipDrawBezier                                   = modgdiplus.NewProc("GdipDrawBezier")
	procGdipDrawBezierI                                  = modgdiplus.NewProc("GdipDrawBezierI")
	procGdipDrawBeziers                                  = modgdiplus.NewProc("GdipDrawBeziers")
	procGdipDrawBeziersI                                 = modgdiplus.NewProc("GdipDrawBeziersI")
	procGdipDrawCachedBitmap                             = modgdiplus.NewProc("GdipDrawCachedBitmap")
	procGdipDrawClosedCurve2                             = modgdiplus.NewProc("GdipDrawClosedCurve2")
	procGdipDrawClosedCurve2I                            = modgdiplus.NewProc("GdipDrawClosedCurve2I")
	procGdipDrawClosedCurve                              = modgdiplus.NewProc("GdipDrawClosedCurve")
	procGdipDrawClosedCurveI                             = modgdiplus.NewProc("GdipDrawClosedCurveI")
	procGdipDrawCurve2                                   = modgdiplus.NewProc("GdipDrawCurve2")
	procGdipDrawCurve2I                                  = modgdiplus.NewProc("GdipDrawCurve2I")
	procGdipDrawCurve3                                   = modgdiplus.NewProc("GdipDrawCurve3")
	procGdipDrawCurve3I                                  = modgdiplus.NewProc("GdipDrawCurve3I")
	procGdipDrawCurve                                    = modgdiplus.NewProc("GdipDrawCurve")
	procGdipDrawCurveI                                   = modgdiplus.NewProc("GdipDrawCurveI")
	procGdipDrawDriverString                             = modgdiplus.NewProc("GdipDrawDriverString")
	procGdipDrawEllipse                                  = modgdiplus.NewProc("GdipDrawEllipse")
	procGdipDrawEllipseI                                 = modgdiplus.NewProc("GdipDrawEllipseI")
	procGdipDrawImage                                    = modgdiplus.NewProc("GdipDrawImage")
	procGdipDrawImageI                                   = modgdiplus.NewProc("GdipDrawImageI")
	procGdipDrawImagePointRect                           = modgdiplus.NewProc("GdipDrawImagePointRect")
	procGdipDrawImagePointRectI                          = modgdiplus.NewProc("GdipDrawImagePointRectI")
	procGdipDrawImagePoints                              = modgdiplus.NewProc("GdipDrawImagePoints")
	procGdipDrawImagePointsI                             = modgdiplus.NewProc("GdipDrawImagePointsI")
	procGdipDrawImagePointsRect                          = modgdiplus.NewProc("GdipDrawImagePointsRect")
	procGdipDrawImagePointsRectI                         = modgdiplus.NewProc("GdipDrawImagePointsRectI")
	procGdipDrawImageRect                                = modgdiplus.NewProc("GdipDrawImageRect")
	procGdipDrawImageRectI                               = modgdiplus.NewProc("GdipDrawImageRectI")
	procGdipDrawImageRectRect                            = modgdiplus.NewProc("GdipDrawImageRectRect")
	procGdipDrawImageRectRectI                           = modgdiplus.NewProc("GdipDrawImageRectRectI")
	procGdipDrawLine                                     = modgdiplus.NewProc("GdipDrawLine")
	procGdipDrawLineI                                    = modgdiplus.NewProc("GdipDrawLineI")
	procGdipDrawLines                                    = modgdiplus.NewProc("GdipDrawLines")
	procGdipDrawLinesI                                   = modgdiplus.NewProc("GdipDrawLinesI")
	procGdipDrawPath                                     = modgdiplus.NewProc("GdipDrawPath")
	procGdipDrawPie                                      = modgdiplus.NewProc("GdipDrawPie")
	procGdipDrawPieI                                     = modgdiplus.NewProc("GdipDrawPieI")
	procGdipDrawPolygon                                  = modgdiplus.NewProc("GdipDrawPolygon")
	procGdipDrawPolygonI                                 = modgdiplus.NewProc("GdipDrawPolygonI")
	procGdipDrawRectangle                                = modgdiplus.NewProc("GdipDrawRectangle")
	procGdipDrawRectangleI                               = modgdiplus.NewProc("GdipDrawRectangleI")
	procGdipDrawRectangles                               = modgdiplus.NewProc("GdipDrawRectangles")
	procGdipDrawRectanglesI                              = modgdiplus.NewProc("GdipDrawRectanglesI")
	procGdipDrawString                                   = modgdiplus.NewProc("GdipDrawString")
	procGdipEmfToWmfBits                                 = modgdiplus.NewProc("GdipEmfToWmfBits")
	procGdipEndContainer                                 = modgdiplus.NewProc("GdipEndContainer")
	procGdipEnumerateMetafileDestPoint                   = modgdiplus.NewProc("GdipEnumerateMetafileDestPoint")
	procGdipEnumerateMetafileDestPointI                  = modgdiplus.NewProc("GdipEnumerateMetafileDestPointI")
	procGdipEnumerateMetafileDestPoints                  = modgdiplus.NewProc("GdipEnumerateMetafileDestPoints")
	procGdipEnumerateMetafileDestPointsI                 = modgdiplus.NewProc("GdipEnumerateMetafileDestPointsI")
	procGdipEnumerateMetafileDestRect                    = modgdiplus.NewProc("GdipEnumerateMetafileDestRect")
	procGdipEnumerateMetafileDestRectI                   = modgdiplus.NewProc("GdipEnumerateMetafileDestRectI")
	procGdipEnumerateMetafileSrcRectDestPoint            = modgdiplus.NewProc("GdipEnumerateMetafileSrcRectDestPoint")
	procGdipEnumerateMetafileSrcRectDestPointI           = modgdiplus.NewProc("GdipEnumerateMetafileSrcRectDestPointI")
	procGdipEnumerateMetafileSrcRectDestPoints           = modgdiplus.NewProc("GdipEnumerateMetafileSrcRectDestPoints")
	procGdipEnumerateMetafileSrcRectDestPointsI          = modgdiplus.NewProc("GdipEnumerateMetafileSrcRectDestPointsI")
	procGdipEnumerateMetafileSrcRectDestRect             = modgdiplus.NewProc("GdipEnumerateMetafileSrcRectDestRect")
	procGdipEnumerateMetafileSrcRectDestRectI            = modgdiplus.NewProc("GdipEnumerateMetafileSrcRectDestRectI")
	procGdipFillClosedCurve2                             = modgdiplus.NewProc("GdipFillClosedCurve2")
	procGdipFillClosedCurve2I                            = modgdiplus.NewProc("GdipFillClosedCurve2I")
	procGdipFillClosedCurve                              = modgdiplus.NewProc("GdipFillClosedCurve")
	procGdipFillClosedCurveI                             = modgdiplus.NewProc("GdipFillClosedCurveI")
	procGdipFillEllipse                                  = modgdiplus.NewProc("GdipFillEllipse")
	procGdipFillEllipseI                                 = modgdiplus.NewProc("GdipFillEllipseI")
	procGdipFillPath                                     = modgdiplus.NewProc("GdipFillPath")
	procGdipFillPie                                      = modgdiplus.NewProc("GdipFillPie")
	procGdipFillPieI                                     = modgdiplus.NewProc("GdipFillPieI")
	procGdipFillPolygon2                                 = modgdiplus.NewProc("GdipFillPolygon2")
	procGdipFillPolygon2I                                = modgdiplus.NewProc("GdipFillPolygon2I")
	procGdipFillPolygon                                  = modgdiplus.NewProc("GdipFillPolygon")
	procGdipFillPolygonI                                 = modgdiplus.NewProc("GdipFillPolygonI")
	procGdipFillRectangle                                = modgdiplus.NewProc("GdipFillRectangle")
	procGdipFillRectangleI                               = modgdiplus.NewProc("GdipFillRectangleI")
	procGdipFillRectangles                               = modgdiplus.NewProc("GdipFillRectangles")
	procGdipFillRectanglesI                              = modgdiplus.NewProc("GdipFillRectanglesI")
	procGdipFillRegion                                   = modgdiplus.NewProc("GdipFillRegion")
	procGdipFlattenPath                                  = modgdiplus.NewProc("GdipFlattenPath")
	procGdipFlush                                        = modgdiplus.NewProc("GdipFlush")
	procGdipFree                                         = modgdiplus.NewProc("GdipFree")
	procGdipGetAdjustableArrowCapFillState               = modgdiplus.NewProc("GdipGetAdjustableArrowCapFillState")
	procGdipGetAdjustableArrowCapHeight                  = modgdiplus.NewProc("GdipGetAdjustableArrowCapHeight")
	procGdipGetAdjustableArrowCapMiddleInset             = modgdiplus.NewProc("GdipGetAdjustableArrowCapMiddleInset")
	procGdipGetAdjustableArrowCapWidth                   = modgdiplus.NewProc("GdipGetAdjustableArrowCapWidth")
	procGdipGetAllPropertyItems                          = modgdiplus.NewProc("GdipGetAllPropertyItems")
	procGdipGetBrushType                                 = modgdiplus.NewProc("GdipGetBrushType")
	procGdipGetCellAscent                                = modgdiplus.NewProc("GdipGetCellAscent")
	procGdipGetCellDescent                               = modgdiplus.NewProc("GdipGetCellDescent")
	procGdipGetClip                                      = modgdiplus.NewProc("GdipGetClip")
	procGdipGetClipBounds                                = modgdiplus.NewProc("GdipGetClipBounds")
	procGdipGetClipBoundsI                               = modgdiplus.NewProc("GdipGetClipBoundsI")
	procGdipGetCompositingMode                           = modgdiplus.NewProc("GdipGetCompositingMode")
	procGdipGetCompositingQuality                        = modgdiplus.NewProc("GdipGetCompositingQuality")
	procGdipGetCustomLineCapBaseCap                      = modgdiplus.NewProc("GdipGetCustomLineCapBaseCap")
	procGdipGetCustomLineCapBaseInset                    = modgdiplus.NewProc("GdipGetCustomLineCapBaseInset")
	procGdipGetCustomLineCapStrokeCaps                   = modgdiplus.NewProc("GdipGetCustomLineCapStrokeCaps")
	procGdipGetCustomLineCapStrokeJoin                   = modgdiplus.NewProc("GdipGetCustomLineCapStrokeJoin")
	procGdipGetCustomLineCapType                         = modgdiplus.NewProc("GdipGetCustomLineCapType")
	procGdipGetCustomLineCapWidthScale                   = modgdiplus.NewProc("GdipGetCustomLineCapWidthScale")
	procGdipGetDC                                        = modgdiplus.NewProc("GdipGetDC")
	procGdipGetDpiX                                      = modgdiplus.NewProc("GdipGetDpiX")
	procGdipGetDpiY                                      = modgdiplus.NewProc("GdipGetDpiY")
	procGdipGetEmHeight                                  = modgdiplus.NewProc("GdipGetEmHeight")
	procGdipGetEncoderParameterList                      = modgdiplus.NewProc("GdipGetEncoderParameterList")
	procGdipGetEncoderParameterListSize                  = modgdiplus.NewProc("GdipGetEncoderParameterListSize")
	procGdipGetFamily                                    = modgdiplus.NewProc("GdipGetFamily")
	procGdipGetFamilyName                                = modgdiplus.NewProc("GdipGetFamilyName")
	procGdipGetFontCollectionFamilyCount                 = modgdiplus.NewProc("GdipGetFontCollectionFamilyCount")
	procGdipGetFontCollectionFamilyList                  = modgdiplus.NewProc("GdipGetFontCollectionFamilyList")
	procGdipGetFontHeight                                = modgdiplus.NewProc("GdipGetFontHeight")
	procGdipGetFontHeightGivenDPI                        = modgdiplus.NewProc("GdipGetFontHeightGivenDPI")
	procGdipGetFontSize                                  = modgdiplus.NewProc("GdipGetFontSize")
	procGdipGetFontStyle                                 = modgdiplus.NewProc("GdipGetFontStyle")
	procGdipGetFontUnit                                  = modgdiplus.NewProc("GdipGetFontUnit")
	procGdipGetGenericFontFamilyMonospace                = modgdiplus.NewProc("GdipGetGenericFontFamilyMonospace")
	procGdipGetGenericFontFamilySansSerif                = modgdiplus.NewProc("GdipGetGenericFontFamilySansSerif")
	procGdipGetGenericFontFamilySerif                    = modgdiplus.NewProc("GdipGetGenericFontFamilySerif")
	procGdipGetHatchBackgroundColor                      = modgdiplus.NewProc("GdipGetHatchBackgroundColor")
	procGdipGetHatchForegroundColor                      = modgdiplus.NewProc("GdipGetHatchForegroundColor")
	procGdipGetHatchStyle                                = modgdiplus.NewProc("GdipGetHatchStyle")
	procGdipGetHemfFromMetafile                          = modgdiplus.NewProc("GdipGetHemfFromMetafile")
	procGdipGetImageAttributesAdjustedPalette            = modgdiplus.NewProc("GdipGetImageAttributesAdjustedPalette")
	procGdipGetImageBounds                               = modgdiplus.NewProc("GdipGetImageBounds")
	procGdipGetImageDecoders                             = modgdiplus.NewProc("GdipGetImageDecoders")
	procGdipGetImageDecodersSize                         = modgdiplus.NewProc("GdipGetImageDecodersSize")
	procGdipGetImageDimension                            = modgdiplus.NewProc("GdipGetImageDimension")
	procGdipGetImageEncoders                             = modgdiplus.NewProc("GdipGetImageEncoders")
	procGdipGetImageEncodersSize                         = modgdiplus.NewProc("GdipGetImageEncodersSize")
	procGdipGetImageFlags                                = modgdiplus.NewProc("GdipGetImageFlags")
	procGdipGetImageGraphicsContext                      = modgdiplus.NewProc("GdipGetImageGraphicsContext")
	procGdipGetImageHeight                               = modgdiplus.NewProc("GdipGetImageHeight")
	procGdipGetImageHorizontalResolution                 = modgdiplus.NewProc("GdipGetImageHorizontalResolution")
	procGdipGetImagePalette                              = modgdiplus.NewProc("GdipGetImagePalette")
	procGdipGetImagePaletteSize                          = modgdiplus.NewProc("GdipGetImagePaletteSize")
	procGdipGetImagePixelFormat                          = modgdiplus.NewProc("GdipGetImagePixelFormat")
	procGdipGetImageRawFormat                            = modgdiplus.NewProc("GdipGetImageRawFormat")
	procGdipGetImageThumbnail                            = modgdiplus.NewProc("GdipGetImageThumbnail")
	procGdipGetImageType                                 = modgdiplus.NewProc("GdipGetImageType")
	procGdipGetImageVerticalResolution                   = modgdiplus.NewProc("GdipGetImageVerticalResolution")
	procGdipGetImageWidth                                = modgdiplus.NewProc("GdipGetImageWidth")
	procGdipGetInterpolationMode                         = modgdiplus.NewProc("GdipGetInterpolationMode")
	procGdipGetLineBlend                                 = modgdiplus.NewProc("GdipGetLineBlend")
	procGdipGetLineBlendCount                            = modgdiplus.NewProc("GdipGetLineBlendCount")
	procGdipGetLineColors                                = modgdiplus.NewProc("GdipGetLineColors")
	procGdipGetLineGammaCorrection                       = modgdiplus.NewProc("GdipGetLineGammaCorrection")
	procGdipGetLinePresetBlend                           = modgdiplus.NewProc("GdipGetLinePresetBlend")
	procGdipGetLinePresetBlendCount                      = modgdiplus.NewProc("GdipGetLinePresetBlendCount")
	procGdipGetLineRect                                  = modgdiplus.NewProc("GdipGetLineRect")
	procGdipGetLineRectI                                 = modgdiplus.NewProc("GdipGetLineRectI")
	procGdipGetLineSpacing                               = modgdiplus.NewProc("GdipGetLineSpacing")
	procGdipGetLineTransform                             = modgdiplus.NewProc("GdipGetLineTransform")
	procGdipGetLineWrapMode                              = modgdiplus.NewProc("GdipGetLineWrapMode")
	procGdipGetLogFontA                                  = modgdiplus.NewProc("GdipGetLogFontA")
	procGdipGetLogFontW                                  = modgdiplus.NewProc("GdipGetLogFontW")
	procGdipGetMatrixElements                            = modgdiplus.NewProc("GdipGetMatrixElements")
	procGdipGetMetafileDownLevelRasterizationLimit       = modgdiplus.NewProc("GdipGetMetafileDownLevelRasterizationLimit")
	procGdipGetMetafileHeaderFromEmf                     = modgdiplus.NewProc("GdipGetMetafileHeaderFromEmf")
	procGdipGetMetafileHeaderFromFile                    = modgdiplus.NewProc("GdipGetMetafileHeaderFromFile")
	procGdipGetMetafileHeaderFromMetafile                = modgdiplus.NewProc("GdipGetMetafileHeaderFromMetafile")
	procGdipGetMetafileHeaderFromStream                  = modgdiplus.NewProc("GdipGetMetafileHeaderFromStream")
	procGdipGetMetafileHeaderFromWmf                     = modgdiplus.NewProc("GdipGetMetafileHeaderFromWmf")
	procGdipGetNearestColor                              = modgdiplus.NewProc("GdipGetNearestColor")
	procGdipGetPageScale                                 = modgdiplus.NewProc("GdipGetPageScale")
	procGdipGetPageUnit                                  = modgdiplus.NewProc("GdipGetPageUnit")
	procGdipGetPathData                                  = modgdiplus.NewProc("GdipGetPathData")
	procGdipGetPathFillMode                              = modgdiplus.NewProc("GdipGetPathFillMode")
	procGdipGetPathGradientBlend                         = modgdiplus.NewProc("GdipGetPathGradientBlend")
	procGdipGetPathGradientBlendCount                    = modgdiplus.NewProc("GdipGetPathGradientBlendCount")
	procGdipGetPathGradientCenterColor                   = modgdiplus.NewProc("GdipGetPathGradientCenterColor")
	procGdipGetPathGradientCenterPoint                   = modgdiplus.NewProc("GdipGetPathGradientCenterPoint")
	procGdipGetPathGradientCenterPointI                  = modgdiplus.NewProc("GdipGetPathGradientCenterPointI")
	procGdipGetPathGradientFocusScales                   = modgdiplus.NewProc("GdipGetPathGradientFocusScales")
	procGdipGetPathGradientGammaCorrection               = modgdiplus.NewProc("GdipGetPathGradientGammaCorrection")
	procGdipGetPathGradientPath                          = modgdiplus.NewProc("GdipGetPathGradientPath")
	procGdipGetPathGradientPointCount                    = modgdiplus.NewProc("GdipGetPathGradientPointCount")
	procGdipGetPathGradientPresetBlend                   = modgdiplus.NewProc("GdipGetPathGradientPresetBlend")
	procGdipGetPathGradientPresetBlendCount              = modgdiplus.NewProc("GdipGetPathGradientPresetBlendCount")
	procGdipGetPathGradientRect                          = modgdiplus.NewProc("GdipGetPathGradientRect")
	procGdipGetPathGradientRectI                         = modgdiplus.NewProc("GdipGetPathGradientRectI")
	procGdipGetPathGradientSurroundColorCount            = modgdiplus.NewProc("GdipGetPathGradientSurroundColorCount")
	procGdipGetPathGradientSurroundColorsWithCount       = modgdiplus.NewProc("GdipGetPathGradientSurroundColorsWithCount")
	procGdipGetPathGradientTransform                     = modgdiplus.NewProc("GdipGetPathGradientTransform")
	procGdipGetPathGradientWrapMode                      = modgdiplus.NewProc("GdipGetPathGradientWrapMode")
	procGdipGetPathLastPoint                             = modgdiplus.NewProc("GdipGetPathLastPoint")
	procGdipGetPathPoints                                = modgdiplus.NewProc("GdipGetPathPoints")
	procGdipGetPathPointsI                               = modgdiplus.NewProc("GdipGetPathPointsI")
	procGdipGetPathTypes                                 = modgdiplus.NewProc("GdipGetPathTypes")
	procGdipGetPathWorldBounds                           = modgdiplus.NewProc("GdipGetPathWorldBounds")
	procGdipGetPathWorldBoundsI                          = modgdiplus.NewProc("GdipGetPathWorldBoundsI")
	procGdipGetPenBrushFill                              = modgdiplus.NewProc("GdipGetPenBrushFill")
	procGdipGetPenColor                                  = modgdiplus.NewProc("GdipGetPenColor")
	procGdipGetPenCompoundArray                          = modgdiplus.NewProc("GdipGetPenCompoundArray")
	procGdipGetPenCompoundCount                          = modgdiplus.NewProc("GdipGetPenCompoundCount")
	procGdipGetPenCustomEndCap                           = modgdiplus.NewProc("GdipGetPenCustomEndCap")
	procGdipGetPenCustomStartCap                         = modgdiplus.NewProc("GdipGetPenCustomStartCap")
	procGdipGetPenDashArray                              = modgdiplus.NewProc("GdipGetPenDashArray")
	procGdipGetPenDashCap197819                          = modgdiplus.NewProc("GdipGetPenDashCap197819")
	procGdipGetPenDashCount                              = modgdiplus.NewProc("GdipGetPenDashCount")
	procGdipGetPenDashOffset                             = modgdiplus.NewProc("GdipGetPenDashOffset")
	procGdipGetPenDashStyle                              = modgdiplus.NewProc("GdipGetPenDashStyle")
	procGdipGetPenEndCap                                 = modgdiplus.NewProc("GdipGetPenEndCap")
	procGdipGetPenFillType                               = modgdiplus.NewProc("GdipGetPenFillType")
	procGdipGetPenLineJoin                               = modgdiplus.NewProc("GdipGetPenLineJoin")
	procGdipGetPenMiterLimit                             = modgdiplus.NewProc("GdipGetPenMiterLimit")
	procGdipGetPenMode                                   = modgdiplus.NewProc("GdipGetPenMode")
	procGdipGetPenStartCap                               = modgdiplus.NewProc("GdipGetPenStartCap")
	procGdipGetPenTransform                              = modgdiplus.NewProc("GdipGetPenTransform")
	procGdipGetPenUnit                                   = modgdiplus.NewProc("GdipGetPenUnit")
	procGdipGetPenWidth                                  = modgdiplus.NewProc("GdipGetPenWidth")
	procGdipGetPixelOffsetMode                           = modgdiplus.NewProc("GdipGetPixelOffsetMode")
	procGdipGetPointCount                                = modgdiplus.NewProc("GdipGetPointCount")
	procGdipGetPropertyCount                             = modgdiplus.NewProc("GdipGetPropertyCount")
	procGdipGetPropertyIdList                            = modgdiplus.NewProc("GdipGetPropertyIdList")
	procGdipGetPropertyItem                              = modgdiplus.NewProc("GdipGetPropertyItem")
	procGdipGetPropertyItemSize                          = modgdiplus.NewProc("GdipGetPropertyItemSize")
	procGdipGetPropertySize                              = modgdiplus.NewProc("GdipGetPropertySize")
	procGdipGetRegionBounds                              = modgdiplus.NewProc("GdipGetRegionBounds")
	procGdipGetRegionBoundsI                             = modgdiplus.NewProc("GdipGetRegionBoundsI")
	procGdipGetRegionData                                = modgdiplus.NewProc("GdipGetRegionData")
	procGdipGetRegionDataSize                            = modgdiplus.NewProc("GdipGetRegionDataSize")
	procGdipGetRegionHRgn                                = modgdiplus.NewProc("GdipGetRegionHRgn")
	procGdipGetRegionScans                               = modgdiplus.NewProc("GdipGetRegionScans")
	procGdipGetRegionScansCount                          = modgdiplus.NewProc("GdipGetRegionScansCount")
	procGdipGetRegionScansI                              = modgdiplus.NewProc("GdipGetRegionScansI")
	procGdipGetRenderingOrigin                           = modgdiplus.NewProc("GdipGetRenderingOrigin")
	procGdipGetSmoothingMode                             = modgdiplus.NewProc("GdipGetSmoothingMode")
	procGdipGetSolidFillColor                            = modgdiplus.NewProc("GdipGetSolidFillColor")
	procGdipGetStringFormatAlign                         = modgdiplus.NewProc("GdipGetStringFormatAlign")
	procGdipGetStringFormatDigitSubstitution             = modgdiplus.NewProc("GdipGetStringFormatDigitSubstitution")
	procGdipGetStringFormatFlags                         = modgdiplus.NewProc("GdipGetStringFormatFlags")
	procGdipGetStringFormatHotkeyPrefix                  = modgdiplus.NewProc("GdipGetStringFormatHotkeyPrefix")
	procGdipGetStringFormatLineAlign                     = modgdiplus.NewProc("GdipGetStringFormatLineAlign")
	procGdipGetStringFormatMeasurableCharacterRangeCount = modgdiplus.NewProc("GdipGetStringFormatMeasurableCharacterRangeCount")
	procGdipGetStringFormatTabStopCount                  = modgdiplus.NewProc("GdipGetStringFormatTabStopCount")
	procGdipGetStringFormatTabStops                      = modgdiplus.NewProc("GdipGetStringFormatTabStops")
	procGdipGetStringFormatTrimming                      = modgdiplus.NewProc("GdipGetStringFormatTrimming")
	procGdipGetTextContrast                              = modgdiplus.NewProc("GdipGetTextContrast")
	procGdipGetTextRenderingHint                         = modgdiplus.NewProc("GdipGetTextRenderingHint")
	procGdipGetTextureImage                              = modgdiplus.NewProc("GdipGetTextureImage")
	procGdipGetTextureTransform                          = modgdiplus.NewProc("GdipGetTextureTransform")
	procGdipGetTextureWrapMode                           = modgdiplus.NewProc("GdipGetTextureWrapMode")
	procGdipGetVisibleClipBounds                         = modgdiplus.NewProc("GdipGetVisibleClipBounds")
	procGdipGetVisibleClipBoundsI                        = modgdiplus.NewProc("GdipGetVisibleClipBoundsI")
	procGdipGetWorldTransform                            = modgdiplus.NewProc("GdipGetWorldTransform")
	procGdipGraphicsClear                                = modgdiplus.NewProc("GdipGraphicsClear")
	procGdipImageForceValidation                         = modgdiplus.NewProc("GdipImageForceValidation")
	procGdipImageGetFrameCount                           = modgdiplus.NewProc("GdipImageGetFrameCount")
	procGdipImageGetFrameDimensionsCount                 = modgdiplus.NewProc("GdipImageGetFrameDimensionsCount")
	procGdipImageGetFrameDimensionsList                  = modgdiplus.NewProc("GdipImageGetFrameDimensionsList")
	procGdipImageRotateFlip                              = modgdiplus.NewProc("GdipImageRotateFlip")
	procGdipImageSelectActiveFrame                       = modgdiplus.NewProc("GdipImageSelectActiveFrame")
	procGdipInvertMatrix                                 = modgdiplus.NewProc("GdipInvertMatrix")
	procGdipIsClipEmpty                                  = modgdiplus.NewProc("GdipIsClipEmpty")
	procGdipIsEmptyRegion                                = modgdiplus.NewProc("GdipIsEmptyRegion")
	procGdipIsEqualRegion                                = modgdiplus.NewProc("GdipIsEqualRegion")
	procGdipIsInfiniteRegion                             = modgdiplus.NewProc("GdipIsInfiniteRegion")
	procGdipIsMatrixEqual                                = modgdiplus.NewProc("GdipIsMatrixEqual")
	procGdipIsMatrixIdentity                             = modgdiplus.NewProc("GdipIsMatrixIdentity")
	procGdipIsMatrixInvertible                           = modgdiplus.NewProc("GdipIsMatrixInvertible")
	procGdipIsOutlineVisiblePathPoint                    = modgdiplus.NewProc("GdipIsOutlineVisiblePathPoint")
	procGdipIsOutlineVisiblePathPointI                   = modgdiplus.NewProc("GdipIsOutlineVisiblePathPointI")
	procGdipIsStyleAvailable                             = modgdiplus.NewProc("GdipIsStyleAvailable")
	procGdipIsVisibleClipEmpty                           = modgdiplus.NewProc("GdipIsVisibleClipEmpty")
	procGdipIsVisiblePathPoint                           = modgdiplus.NewProc("GdipIsVisiblePathPoint")
	procGdipIsVisiblePathPointI                          = modgdiplus.NewProc("GdipIsVisiblePathPointI")
	procGdipIsVisiblePoint                               = modgdiplus.NewProc("GdipIsVisiblePoint")
	procGdipIsVisiblePointI                              = modgdiplus.NewProc("GdipIsVisiblePointI")
	procGdipIsVisibleRect                                = modgdiplus.NewProc("GdipIsVisibleRect")
	procGdipIsVisibleRectI                               = modgdiplus.NewProc("GdipIsVisibleRectI")
	procGdipIsVisibleRegionPoint                         = modgdiplus.NewProc("GdipIsVisibleRegionPoint")
	procGdipIsVisibleRegionPointI                        = modgdiplus.NewProc("GdipIsVisibleRegionPointI")
	procGdipIsVisibleRegionRect                          = modgdiplus.NewProc("GdipIsVisibleRegionRect")
	procGdipIsVisibleRegionRectI                         = modgdiplus.NewProc("GdipIsVisibleRegionRectI")
	procGdipLoadImageFromFile                            = modgdiplus.NewProc("GdipLoadImageFromFile")
	procGdipLoadImageFromFileICM                         = modgdiplus.NewProc("GdipLoadImageFromFileICM")
	procGdipLoadImageFromStream                          = modgdiplus.NewProc("GdipLoadImageFromStream")
	procGdipLoadImageFromStreamICM                       = modgdiplus.NewProc("GdipLoadImageFromStreamICM")
	procGdipMeasureCharacterRanges                       = modgdiplus.NewProc("GdipMeasureCharacterRanges")
	procGdipMeasureDriverString                          = modgdiplus.NewProc("GdipMeasureDriverString")
	procGdipMeasureString                                = modgdiplus.NewProc("GdipMeasureString")
	procGdipMultiplyLineTransform                        = modgdiplus.NewProc("GdipMultiplyLineTransform")
	procGdipMultiplyMatrix                               = modgdiplus.NewProc("GdipMultiplyMatrix")
	procGdipMultiplyPathGradientTransform                = modgdiplus.NewProc("GdipMultiplyPathGradientTransform")
	procGdipMultiplyPenTransform                         = modgdiplus.NewProc("GdipMultiplyPenTransform")
	procGdipMultiplyTextureTransform                     = modgdiplus.NewProc("GdipMultiplyTextureTransform")
	procGdipMultiplyWorldTransform                       = modgdiplus.NewProc("GdipMultiplyWorldTransform")
	procGdipNewInstalledFontCollection                   = modgdiplus.NewProc("GdipNewInstalledFontCollection")
	procGdipNewPrivateFontCollection                     = modgdiplus.NewProc("GdipNewPrivateFontCollection")
	procGdipPathIterCopyData                             = modgdiplus.NewProc("GdipPathIterCopyData")
	procGdipPathIterEnumerate                            = modgdiplus.NewProc("GdipPathIterEnumerate")
	procGdipPathIterGetCount                             = modgdiplus.NewProc("GdipPathIterGetCount")
	procGdipPathIterGetSubpathCount                      = modgdiplus.NewProc("GdipPathIterGetSubpathCount")
	procGdipPathIterHasCurve                             = modgdiplus.NewProc("GdipPathIterHasCurve")
	procGdipPathIterIsValid                              = modgdiplus.NewProc("GdipPathIterIsValid")
	procGdipPathIterNextMarker                           = modgdiplus.NewProc("GdipPathIterNextMarker")
	procGdipPathIterNextMarkerPath                       = modgdiplus.NewProc("GdipPathIterNextMarkerPath")
	procGdipPathIterNextPathType                         = modgdiplus.NewProc("GdipPathIterNextPathType")
	procGdipPathIterNextSubpath                          = modgdiplus.NewProc("GdipPathIterNextSubpath")
	procGdipPathIterNextSubpathPath                      = modgdiplus.NewProc("GdipPathIterNextSubpathPath")
	procGdipPathIterRewind                               = modgdiplus.NewProc("GdipPathIterRewind")
	procGdipPlayMetafileRecord                           = modgdiplus.NewProc("GdipPlayMetafileRecord")
	procGdipPrivateAddFontFile                           = modgdiplus.NewProc("GdipPrivateAddFontFile")
	procGdipPrivateAddMemoryFont                         = modgdiplus.NewProc("GdipPrivateAddMemoryFont")
	procGdipRecordMetafile                               = modgdiplus.NewProc("GdipRecordMetafile")
	procGdipRecordMetafileFileName                       = modgdiplus.NewProc("GdipRecordMetafileFileName")
	procGdipRecordMetafileFileNameI                      = modgdiplus.NewProc("GdipRecordMetafileFileNameI")
	procGdipRecordMetafileI                              = modgdiplus.NewProc("GdipRecordMetafileI")
	procGdipRecordMetafileStream                         = modgdiplus.NewProc("GdipRecordMetafileStream")
	procGdipRecordMetafileStreamI                        = modgdiplus.NewProc("GdipRecordMetafileStreamI")
	procGdipReleaseDC                                    = modgdiplus.NewProc("GdipReleaseDC")
	procGdipRemovePropertyItem                           = modgdiplus.NewProc("GdipRemovePropertyItem")
	procGdipResetClip                                    = modgdiplus.NewProc("GdipResetClip")
	procGdipResetImageAttributes                         = modgdiplus.NewProc("GdipResetImageAttributes")
	procGdipResetLineTransform                           = modgdiplus.NewProc("GdipResetLineTransform")
	procGdipResetPageTransform                           = modgdiplus.NewProc("GdipResetPageTransform")
	procGdipResetPath                                    = modgdiplus.NewProc("GdipResetPath")
	procGdipResetPathGradientTransform                   = modgdiplus.NewProc("GdipResetPathGradientTransform")
	procGdipResetPenTransform                            = modgdiplus.NewProc("GdipResetPenTransform")
	procGdipResetTextureTransform                        = modgdiplus.NewProc("GdipResetTextureTransform")
	procGdipResetWorldTransform                          = modgdiplus.NewProc("GdipResetWorldTransform")
	procGdipRestoreGraphics                              = modgdiplus.NewProc("GdipRestoreGraphics")
	procGdipReversePath                                  = modgdiplus.NewProc("GdipReversePath")
	procGdipRotateLineTransform                          = modgdiplus.NewProc("GdipRotateLineTransform")
	procGdipRotateMatrix                                 = modgdiplus.NewProc("GdipRotateMatrix")
	procGdipRotatePathGradientTransform                  = modgdiplus.NewProc("GdipRotatePathGradientTransform")
	procGdipRotatePenTransform                           = modgdiplus.NewProc("GdipRotatePenTransform")
	procGdipRotateTextureTransform                       = modgdiplus.NewProc("GdipRotateTextureTransform")
	procGdipRotateWorldTransform                         = modgdiplus.NewProc("GdipRotateWorldTransform")
	procGdipSaveAdd                                      = modgdiplus.NewProc("GdipSaveAdd")
	procGdipSaveAddImage                                 = modgdiplus.NewProc("GdipSaveAddImage")
	procGdipSaveGraphics                                 = modgdiplus.NewProc("GdipSaveGraphics")
	procGdipSaveImageToFile                              = modgdiplus.NewProc("GdipSaveImageToFile")
	procGdipSaveImageToStream                            = modgdiplus.NewProc("GdipSaveImageToStream")
	procGdipScaleLineTransform                           = modgdiplus.NewProc("GdipScaleLineTransform")
	procGdipScaleMatrix                                  = modgdiplus.NewProc("GdipScaleMatrix")
	procGdipScalePathGradientTransform                   = modgdiplus.NewProc("GdipScalePathGradientTransform")
	procGdipScalePenTransform                            = modgdiplus.NewProc("GdipScalePenTransform")
	procGdipScaleTextureTransform                        = modgdiplus.NewProc("GdipScaleTextureTransform")
	procGdipScaleWorldTransform                          = modgdiplus.NewProc("GdipScaleWorldTransform")
	procGdipSetAdjustableArrowCapFillState               = modgdiplus.NewProc("GdipSetAdjustableArrowCapFillState")
	procGdipSetAdjustableArrowCapHeight                  = modgdiplus.NewProc("GdipSetAdjustableArrowCapHeight")
	procGdipSetAdjustableArrowCapMiddleInset             = modgdiplus.NewProc("GdipSetAdjustableArrowCapMiddleInset")
	procGdipSetAdjustableArrowCapWidth                   = modgdiplus.NewProc("GdipSetAdjustableArrowCapWidth")
	procGdipSetClipGraphics                              = modgdiplus.NewProc("GdipSetClipGraphics")
	procGdipSetClipHrgn                                  = modgdiplus.NewProc("GdipSetClipHrgn")
	procGdipSetClipPath                                  = modgdiplus.NewProc("GdipSetClipPath")
	procGdipSetClipRect                                  = modgdiplus.NewProc("GdipSetClipRect")
	procGdipSetClipRectI                                 = modgdiplus.NewProc("GdipSetClipRectI")
	procGdipSetClipRegion                                = modgdiplus.NewProc("GdipSetClipRegion")
	procGdipSetCompositingMode                           = modgdiplus.NewProc("GdipSetCompositingMode")
	procGdipSetCompositingQuality                        = modgdiplus.NewProc("GdipSetCompositingQuality")
	procGdipSetCustomLineCapBaseCap                      = modgdiplus.NewProc("GdipSetCustomLineCapBaseCap")
	procGdipSetCustomLineCapBaseInset                    = modgdiplus.NewProc("GdipSetCustomLineCapBaseInset")
	procGdipSetCustomLineCapStrokeCaps                   = modgdiplus.NewProc("GdipSetCustomLineCapStrokeCaps")
	procGdipSetCustomLineCapStrokeJoin                   = modgdiplus.NewProc("GdipSetCustomLineCapStrokeJoin")
	procGdipSetCustomLineCapWidthScale                   = modgdiplus.NewProc("GdipSetCustomLineCapWidthScale")
	procGdipSetEmpty                                     = modgdiplus.NewProc("GdipSetEmpty")
	procGdipSetImageAttributesCachedBackground           = modgdiplus.NewProc("GdipSetImageAttributesCachedBackground")
	procGdipSetImageAttributesColorKeys                  = modgdiplus.NewProc("GdipSetImageAttributesColorKeys")
	procGdipSetImageAttributesColorMatrix                = modgdiplus.NewProc("GdipSetImageAttributesColorMatrix")
	procGdipSetImageAttributesGamma                      = modgdiplus.NewProc("GdipSetImageAttributesGamma")
	procGdipSetImageAttributesNoOp                       = modgdiplus.NewProc("GdipSetImageAttributesNoOp")
	procGdipSetImageAttributesOutputChannel              = modgdiplus.NewProc("GdipSetImageAttributesOutputChannel")
	procGdipSetImageAttributesOutputChannelColorProfile  = modgdiplus.NewProc("GdipSetImageAttributesOutputChannelColorProfile")
	procGdipSetImageAttributesRemapTable                 = modgdiplus.NewProc("GdipSetImageAttributesRemapTable")
	procGdipSetImageAttributesThreshold                  = modgdiplus.NewProc("GdipSetImageAttributesThreshold")
	procGdipSetImageAttributesToIdentity                 = modgdiplus.NewProc("GdipSetImageAttributesToIdentity")
	procGdipSetImageAttributesWrapMode                   = modgdiplus.NewProc("GdipSetImageAttributesWrapMode")
	procGdipSetImagePalette                              = modgdiplus.NewProc("GdipSetImagePalette")
	procGdipSetInfinite                                  = modgdiplus.NewProc("GdipSetInfinite")
	procGdipSetInterpolationMode                         = modgdiplus.NewProc("GdipSetInterpolationMode")
	procGdipSetLineBlend                                 = modgdiplus.NewProc("GdipSetLineBlend")
	procGdipSetLineColors                                = modgdiplus.NewProc("GdipSetLineColors")
	procGdipSetLineGammaCorrection                       = modgdiplus.NewProc("GdipSetLineGammaCorrection")
	procGdipSetLineLinearBlend                           = modgdiplus.NewProc("GdipSetLineLinearBlend")
	procGdipSetLinePresetBlend                           = modgdiplus.NewProc("GdipSetLinePresetBlend")
	procGdipSetLineSigmaBlend                            = modgdiplus.NewProc("GdipSetLineSigmaBlend")
	procGdipSetLineTransform                             = modgdiplus.NewProc("GdipSetLineTransform")
	procGdipSetLineWrapMode                              = modgdiplus.NewProc("GdipSetLineWrapMode")
	procGdipSetMatrixElements                            = modgdiplus.NewProc("GdipSetMatrixElements")
	procGdipSetMetafileDownLevelRasterizationLimit       = modgdiplus.NewProc("GdipSetMetafileDownLevelRasterizationLimit")
	procGdipSetPageScale                                 = modgdiplus.NewProc("GdipSetPageScale")
	procGdipSetPageUnit                                  = modgdiplus.NewProc("GdipSetPageUnit")
	procGdipSetPathFillMode                              = modgdiplus.NewProc("GdipSetPathFillMode")
	procGdipSetPathGradientBlend                         = modgdiplus.NewProc("GdipSetPathGradientBlend")
	procGdipSetPathGradientCenterColor                   = modgdiplus.NewProc("GdipSetPathGradientCenterColor")
	procGdipSetPathGradientCenterPoint                   = modgdiplus.NewProc("GdipSetPathGradientCenterPoint")
	procGdipSetPathGradientCenterPointI                  = modgdiplus.NewProc("GdipSetPathGradientCenterPointI")
	procGdipSetPathGradientFocusScales                   = modgdiplus.NewProc("GdipSetPathGradientFocusScales")
	procGdipSetPathGradientGammaCorrection               = modgdiplus.NewProc("GdipSetPathGradientGammaCorrection")
	procGdipSetPathGradientLinearBlend                   = modgdiplus.NewProc("GdipSetPathGradientLinearBlend")
	procGdipSetPathGradientPath                          = modgdiplus.NewProc("GdipSetPathGradientPath")
	procGdipSetPathGradientPresetBlend                   = modgdiplus.NewProc("GdipSetPathGradientPresetBlend")
	procGdipSetPathGradientSigmaBlend                    = modgdiplus.NewProc("GdipSetPathGradientSigmaBlend")
	procGdipSetPathGradientSurroundColorsWithCount       = modgdiplus.NewProc("GdipSetPathGradientSurroundColorsWithCount")
	procGdipSetPathGradientTransform                     = modgdiplus.NewProc("GdipSetPathGradientTransform")
	procGdipSetPathGradientWrapMode                      = modgdiplus.NewProc("GdipSetPathGradientWrapMode")
	procGdipSetPathMarker                                = modgdiplus.NewProc("GdipSetPathMarker")
	procGdipSetPenBrushFill                              = modgdiplus.NewProc("GdipSetPenBrushFill")
	procGdipSetPenColor                                  = modgdiplus.NewProc("GdipSetPenColor")
	procGdipSetPenCompoundArray                          = modgdiplus.NewProc("GdipSetPenCompoundArray")
	procGdipSetPenCustomEndCap                           = modgdiplus.NewProc("GdipSetPenCustomEndCap")
	procGdipSetPenCustomStartCap                         = modgdiplus.NewProc("GdipSetPenCustomStartCap")
	procGdipSetPenDashArray                              = modgdiplus.NewProc("GdipSetPenDashArray")
	procGdipSetPenDashCap197819                          = modgdiplus.NewProc("GdipSetPenDashCap197819")
	procGdipSetPenDashOffset                             = modgdiplus.NewProc("GdipSetPenDashOffset")
	procGdipSetPenDashStyle                              = modgdiplus.NewProc("GdipSetPenDashStyle")
	procGdipSetPenEndCap                                 = modgdiplus.NewProc("GdipSetPenEndCap")
	procGdipSetPenLineCap197819                          = modgdiplus.NewProc("GdipSetPenLineCap197819")
	procGdipSetPenLineJoin                               = modgdiplus.NewProc("GdipSetPenLineJoin")
	procGdipSetPenMiterLimit                             = modgdiplus.NewProc("GdipSetPenMiterLimit")
	procGdipSetPenMode                                   = modgdiplus.NewProc("GdipSetPenMode")
	procGdipSetPenStartCap                               = modgdiplus.NewProc("GdipSetPenStartCap")
	procGdipSetPenTransform                              = modgdiplus.NewProc("GdipSetPenTransform")
	procGdipSetPenUnit                                   = modgdiplus.NewProc("GdipSetPenUnit")
	procGdipSetPenWidth                                  = modgdiplus.NewProc("GdipSetPenWidth")
	procGdipSetPixelOffsetMode                           = modgdiplus.NewProc("GdipSetPixelOffsetMode")
	procGdipSetPropertyItem                              = modgdiplus.NewProc("GdipSetPropertyItem")
	procGdipSetRenderingOrigin                           = modgdiplus.NewProc("GdipSetRenderingOrigin")
	procGdipSetSmoothingMode                             = modgdiplus.NewProc("GdipSetSmoothingMode")
	procGdipSetSolidFillColor                            = modgdiplus.NewProc("GdipSetSolidFillColor")
	procGdipSetStringFormatAlign                         = modgdiplus.NewProc("GdipSetStringFormatAlign")
	procGdipSetStringFormatDigitSubstitution             = modgdiplus.NewProc("GdipSetStringFormatDigitSubstitution")
	procGdipSetStringFormatFlags                         = modgdiplus.NewProc("GdipSetStringFormatFlags")
	procGdipSetStringFormatHotkeyPrefix                  = modgdiplus.NewProc("GdipSetStringFormatHotkeyPrefix")
	procGdipSetStringFormatLineAlign                     = modgdiplus.NewProc("GdipSetStringFormatLineAlign")
	procGdipSetStringFormatMeasurableCharacterRanges     = modgdiplus.NewProc("GdipSetStringFormatMeasurableCharacterRanges")
	procGdipSetStringFormatTabStops                      = modgdiplus.NewProc("GdipSetStringFormatTabStops")
	procGdipSetStringFormatTrimming                      = modgdiplus.NewProc("GdipSetStringFormatTrimming")
	procGdipSetTextContrast                              = modgdiplus.NewProc("GdipSetTextContrast")
	procGdipSetTextRenderingHint                         = modgdiplus.NewProc("GdipSetTextRenderingHint")
	procGdipSetTextureTransform                          = modgdiplus.NewProc("GdipSetTextureTransform")
	procGdipSetTextureWrapMode                           = modgdiplus.NewProc("GdipSetTextureWrapMode")
	procGdipSetWorldTransform                            = modgdiplus.NewProc("GdipSetWorldTransform")
	procGdipShearMatrix                                  = modgdiplus.NewProc("GdipShearMatrix")
	procGdipStartPathFigure                              = modgdiplus.NewProc("GdipStartPathFigure")
	procGdipStringFormatGetGenericDefault                = modgdiplus.NewProc("GdipStringFormatGetGenericDefault")
	procGdipStringFormatGetGenericTypographic            = modgdiplus.NewProc("GdipStringFormatGetGenericTypographic")
	procGdipTestControl                                  = modgdiplus.NewProc("GdipTestControl")
	procGdipTransformMatrixPoints                        = modgdiplus.NewProc("GdipTransformMatrixPoints")
	procGdipTransformMatrixPointsI                       = modgdiplus.NewProc("GdipTransformMatrixPointsI")
	procGdipTransformPath                                = modgdiplus.NewProc("GdipTransformPath")
	procGdipTransformPoints                              = modgdiplus.NewProc("GdipTransformPoints")
	procGdipTransformPointsI                             = modgdiplus.NewProc("GdipTransformPointsI")
	procGdipTransformRegion                              = modgdiplus.NewProc("GdipTransformRegion")
	procGdipTranslateClip                                = modgdiplus.NewProc("GdipTranslateClip")
	procGdipTranslateClipI                               = modgdiplus.NewProc("GdipTranslateClipI")
	procGdipTranslateLineTransform                       = modgdiplus.NewProc("GdipTranslateLineTransform")
	procGdipTranslateMatrix                              = modgdiplus.NewProc("GdipTranslateMatrix")
	procGdipTranslatePathGradientTransform               = modgdiplus.NewProc("GdipTranslatePathGradientTransform")
	procGdipTranslatePenTransform                        = modgdiplus.NewProc("GdipTranslatePenTransform")
	procGdipTranslateRegion                              = modgdiplus.NewProc("GdipTranslateRegion")
	procGdipTranslateRegionI                             = modgdiplus.NewProc("GdipTranslateRegionI")
	procGdipTranslateTextureTransform                    = modgdiplus.NewProc("GdipTranslateTextureTransform")
	procGdipTranslateWorldTransform                      = modgdiplus.NewProc("GdipTranslateWorldTransform")
	procGdipVectorTransformMatrixPoints                  = modgdiplus.NewProc("GdipVectorTransformMatrixPoints")
	procGdipVectorTransformMatrixPointsI                 = modgdiplus.NewProc("GdipVectorTransformMatrixPointsI")
	procGdipWarpPath                                     = modgdiplus.NewProc("GdipWarpPath")
	procGdipWidenPath                                    = modgdiplus.NewProc("GdipWidenPath")
	procGdipWindingModeOutline                           = modgdiplus.NewProc("GdipWindingModeOutline")
	procGdiplusNotificationHook                          = modgdiplus.NewProc("GdiplusNotificationHook")
	procGdiplusNotificationUnhook                        = modgdiplus.NewProc("GdiplusNotificationUnhook")
	procGdiplusShutdown                                  = modgdiplus.NewProc("GdiplusShutdown")
	procGdiplusStartup                                   = modgdiplus.NewProc("GdiplusStartup")
	procGdipFindFirstImageItem                           = modgdiplus.NewProc("GdipFindFirstImageItem")
	procGdipFindNextImageItem                            = modgdiplus.NewProc("GdipFindNextImageItem")
	procGdipGetImageItemData                             = modgdiplus.NewProc("GdipGetImageItemData")
	procGdipCreateEffect                                 = modgdiplus.NewProc("GdipCreateEffect")
	procGdipDeleteEffect                                 = modgdiplus.NewProc("GdipDeleteEffect")
	procGdipGetEffectParameterSize                       = modgdiplus.NewProc("GdipGetEffectParameterSize")
	procGdipGetEffectParameters                          = modgdiplus.NewProc("GdipGetEffectParameters")
	procGdipSetEffectParameters                          = modgdiplus.NewProc("GdipSetEffectParameters")
	procGdipBitmapCreateApplyEffect                      = modgdiplus.NewProc("GdipBitmapCreateApplyEffect")
	procGdipBitmapApplyEffect                            = modgdiplus.NewProc("GdipBitmapApplyEffect")
	procGdipBitmapGetHistogram                           = modgdiplus.NewProc("GdipBitmapGetHistogram")
	procGdipBitmapGetHistogramSize                       = modgdiplus.NewProc("GdipBitmapGetHistogramSize")
	procGdipBitmapConvertFormat                          = modgdiplus.NewProc("GdipBitmapConvertFormat")
	procGdipInitializePalette                            = modgdiplus.NewProc("GdipInitializePalette")
	procGdipImageSetAbort                                = modgdiplus.NewProc("GdipImageSetAbort")
	procGdipGraphicsSetAbort                             = modgdiplus.NewProc("GdipGraphicsSetAbort")
	procGdipDrawImageFX                                  = modgdiplus.NewProc("GdipDrawImageFX")
	procGdipConvertToEmfPlus                             = modgdiplus.NewProc("GdipConvertToEmfPlus")
	procGdipConvertToEmfPlusToFile                       = modgdiplus.NewProc("GdipConvertToEmfPlusToFile")
	procGdipConvertToEmfPlusToStream                     = modgdiplus.NewProc("GdipConvertToEmfPlusToStream")
	procGdipPlayTSClientRecord                           = modgdiplus.NewProc("GdipPlayTSClientRecord")
)

func GdipGetStringFormatMeasurableCharacterRangeCount(format *GpStringFormat, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatMeasurableCharacterRangeCount,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(count)))
}

func GdipSetStringFormatMeasurableCharacterRanges(format *GpStringFormat, rangeCount INT, ranges *CharacterRange) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatMeasurableCharacterRanges,
		uintptr(unsafe.Pointer(format)),
		uintptr(rangeCount),
		uintptr(unsafe.Pointer(ranges)))
}

func GdipGetStringFormatTrimming(format *GpStringFormat, trimming *StringTrimming) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatTrimming,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(trimming)))
}

func GdipSetStringFormatTrimming(format *GpStringFormat, trimming StringTrimming) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatTrimming,
		uintptr(unsafe.Pointer(format)),
		uintptr(trimming))
}

func GdipGetStringFormatDigitSubstitution(format *GpStringFormat, language *LANGID, substitute *StringDigitSubstitute) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatDigitSubstitution,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(language)),
		uintptr(unsafe.Pointer(substitute)))
}

func GdipSetStringFormatDigitSubstitution(format *GpStringFormat, language LANGID, substitute StringDigitSubstitute) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatDigitSubstitution,
		uintptr(unsafe.Pointer(format)),
		uintptr(language),
		uintptr(substitute))
}

func GdipGetStringFormatTabStops(format *GpStringFormat, count INT, firstTabOffset *REAL, tabStops []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatTabStops,
		uintptr(unsafe.Pointer(format)),
		uintptr(count),
		uintptr(unsafe.Pointer(firstTabOffset)),
		uintptr(unsafe.Pointer(&tabStops[0])))
}

func GdipGetStringFormatTabStopCount(format *GpStringFormat, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatTabStopCount,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(count)))
}

func GdipSetStringFormatTabStops(format *GpStringFormat, firstTabOffset REAL, tabStops []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatTabStops,
		uintptr(unsafe.Pointer(format)),
		REALbits(firstTabOffset),
		uintptr(INT(len(tabStops))),
		uintptr(unsafe.Pointer(&tabStops[0])))
}

func GdipGetStringFormatHotkeyPrefix(format *GpStringFormat, hotkeyPrefix *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatHotkeyPrefix,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(hotkeyPrefix)))
}

func GdipSetStringFormatHotkeyPrefix(format *GpStringFormat, hotkeyPrefix INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatHotkeyPrefix,
		uintptr(unsafe.Pointer(format)),
		uintptr(hotkeyPrefix))
}

func GdipSetStringFormatLineAlign(format *GpStringFormat, align StringAlignment) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatLineAlign,
		uintptr(unsafe.Pointer(format)),
		uintptr(align))
}

func GdipGetStringFormatLineAlign(format *GpStringFormat, align *StringAlignment) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatLineAlign,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(align)))
}

func GdipSetStringFormatAlign(format *GpStringFormat, align StringAlignment) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatAlign,
		uintptr(unsafe.Pointer(format)),
		uintptr(align))
}

func GdipGetStringFormatAlign(format *GpStringFormat, align *StringAlignment) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatAlign,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(align)))
}

func GdipGetStringFormatFlags(format *GpStringFormat, flags *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetStringFormatFlags,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(flags)))
}

func GdipSetStringFormatFlags(format *GpStringFormat, flags INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetStringFormatFlags,
		uintptr(unsafe.Pointer(format)),
		uintptr(flags))
}

func GdipCloneStringFormat(format *GpStringFormat,
	newFormat **GpStringFormat) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneStringFormat,
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(newFormat)))
}

func GdipStringFormatGetGenericDefault(format **GpStringFormat) (status Status, err error) {
	return gdiplusSyscall(procGdipStringFormatGetGenericDefault,
		uintptr(unsafe.Pointer(format)))
}

func GdipStringFormatGetGenericTypographic(format **GpStringFormat) (status Status, err error) {
	return gdiplusSyscall(procGdipStringFormatGetGenericTypographic,
		uintptr(unsafe.Pointer(format)))
}

func GdipGetRegionScansI(region *GpRegion, rects []Rect, count *INT,
	matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionScansI,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(&rects[0])),
		uintptr(unsafe.Pointer(count)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipGetRegionScans(region *GpRegion, rects []RectF, count *INT,
	matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionScans,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(&rects[0])),
		uintptr(unsafe.Pointer(count)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipGetRegionScansCount(region *GpRegion, count *UINT, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionScansCount,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(count)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipIsVisibleRegionRectI(region *GpRegion, x, y, width, height INT,
	graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleRegionRectI,
		uintptr(unsafe.Pointer(region)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsVisibleRegionRect(region *GpRegion, x, y, width, height REAL,
	graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleRegionRect,
		uintptr(unsafe.Pointer(region)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsVisibleRegionPointI(region *GpRegion, x, y INT,
	graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleRegionPointI,
		uintptr(unsafe.Pointer(region)),
		uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsVisibleRegionPoint(region *GpRegion, x, y REAL,
	graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleRegionPoint,
		uintptr(unsafe.Pointer(region)),
		REALbits(x), REALbits(y),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipGetRegionDataSize(region *GpRegion, bufferSize *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionDataSize,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(bufferSize)))
}

func GdipIsEqualRegion(region *GpRegion, region2 *GpRegion, graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsEqualRegion,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(region2)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsInfiniteRegion(region *GpRegion, graphics *GpGraphics, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsInfiniteRegion,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(booln)))
}

func GdipIsEmptyRegion(region *GpRegion, graphics *GpGraphics, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsEmptyRegion,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(booln)))
}

func GdipGetRegionBoundsI(region *GpRegion, graphics *GpGraphics, rect *Rect) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionBoundsI,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetRegionBounds(region *GpRegion, graphics *GpGraphics, rect *RectF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionBounds,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipTransformRegion(region *GpRegion, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipTransformRegion,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipTranslateRegionI(region *GpRegion, dx, dy INT) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateRegionI,
		uintptr(unsafe.Pointer(region)),
		uintptr(dx), uintptr(dy))
}

func GdipTranslateRegion(region *GpRegion, dx, dy REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateRegion,
		uintptr(unsafe.Pointer(region)),
		REALbits(dx), REALbits(dy))
}

func GdipGetPenCompoundArray(pen *GpPen, compoundArray []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenCompoundArray,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&compoundArray[0])),
		uintptr(INT(len(compoundArray))))
}

func GdipGetPenCompoundCount(pen *GpPen, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenCompoundCount,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(count)))
}

func GdipSetPenCompoundArray(pen *GpPen, compoundArray []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenCompoundArray,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&compoundArray[0])),
		uintptr(INT(len(compoundArray))))
}

func GdipGetPenDashArray(pen *GpPen, dashArray []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenDashArray,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&dashArray[0])),
		uintptr(INT(len(dashArray))))
}

func GdipGetPenDashCount(pen *GpPen, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenDashCount,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(count)))
}

func GdipSetPenDashArray(pen *GpPen, dashArray []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenDashArray,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&dashArray[0])),
		uintptr(INT(len(dashArray))))
}

func GdipSetPenDashOffset(pen *GpPen, dashOffset REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenDashOffset,
		uintptr(unsafe.Pointer(pen)),
		REALbits(dashOffset))
}

func GdipGetPenDashOffset(pen *GpPen, dashOffset *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenDashOffset,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashOffset)))
}

func GdipSetPenDashStyle(pen *GpPen, dashStyle DashStyle) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenDashStyle,
		uintptr(unsafe.Pointer(pen)),
		uintptr(dashStyle))
}

func GdipGetPenDashStyle(pen *GpPen, dashStyle *DashStyle) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenDashStyle,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashStyle)))
}

func GdipGetHatchBackgroundColor(brush *GpBrush, argb *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetHatchBackgroundColor,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(argb)))
}

func GdipGetHatchForegroundColor(brush *GpBrush, argb *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetHatchForegroundColor,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(argb)))
}

func GdipGetHatchStyle(brush *GpBrush, hatchStyle *HatchStyle) (status Status, err error) {
	return gdiplusSyscall(procGdipGetHatchStyle,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(hatchStyle)))
}

func GdipGetPenBrushFill(pen *GpPen, brush **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenBrushFill,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(brush)))
}

func GdipCreateHatchBrush(hatchstyle HatchStyle, forecol,
	backcol ARGB, brush **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateHatchBrush,
		uintptr(hatchstyle),
		uintptr(forecol),
		uintptr(backcol),
		uintptr(unsafe.Pointer(brush)))
}

func GdipSetPenBrushFill(pen *GpPen, brush *GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenBrushFill,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(brush)))
}

func GdipGetPenColor(pen *GpPen, argb *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenColor,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(argb)))
}

func GdipSetPenColor(pen *GpPen, argb ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenColor,
		uintptr(unsafe.Pointer(pen)),
		uintptr(argb))
}

func GdipGetPenFillType(pen *GpPen, typ *PenType) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenFillType,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(typ)))
}

func GdipMultiplyPenTransform(pen *GpPen, matrix *GpMatrix,
	order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipMultiplyPenTransform,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
}

func GdipResetPenTransform(pen *GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipResetPenTransform,
		uintptr(unsafe.Pointer(pen)))
}

func GdipGetPenTransform(pen *GpPen, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenTransform,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipSetPenTransform(pen *GpPen, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenTransform,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipGetPenMode(pen *GpPen, penAlignment *PenAlignment) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenMode,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(penAlignment)))
}

func GdipSetPenMode(pen *GpPen, penAlignment PenAlignment) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenMode,
		uintptr(unsafe.Pointer(pen)),
		uintptr(penAlignment))
}

func GdipGetPenMiterLimit(pen *GpPen, miterLimit *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenMiterLimit,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(miterLimit)))
}

func GdipSetPenMiterLimit(pen *GpPen, miterLimit REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenMiterLimit,
		uintptr(unsafe.Pointer(pen)),
		REALbits(miterLimit))
}

func GdipGetPenCustomEndCap(pen *GpPen, customCap **GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenCustomEndCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
}

func GdipGetPenCustomStartCap(pen *GpPen, customCap **GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenCustomStartCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
}

func GdipGetPenLineJoin(pen *GpPen, lineJoin *LineJoin) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenLineJoin,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(lineJoin)))
}

func GdipSetPenLineJoin(pen *GpPen, lineJoin LineJoin) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenLineJoin,
		uintptr(unsafe.Pointer(pen)),
		uintptr(lineJoin))
}

func GdipGetPenDashCap197819(pen *GpPen, dashCap *DashCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenDashCap197819,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashCap)))
}

func GdipGetPenEndCap(pen *GpPen, endCap *LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenEndCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(endCap)))
}

func GdipGetPenStartCap(pen *GpPen, startCap *LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenStartCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(startCap)))
}

func GdipSetPenDashCap197819(pen *GpPen, dashCap DashCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenDashCap197819,
		uintptr(unsafe.Pointer(pen)),
		uintptr(dashCap))
}

func GdipSetPenEndCap(pen *GpPen, endCap LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenEndCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(endCap))
}

func GdipSetPenStartCap(pen *GpPen, startCap LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenStartCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(startCap))
}

func GdipSetPenLineCap197819(pen *GpPen, startCap, endCap LineCap,
	dashCap DashCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenLineCap197819,
		uintptr(unsafe.Pointer(pen)),
		uintptr(startCap), uintptr(endCap),
		uintptr(dashCap))
}

func GdipGetPenWidth(pen *GpPen, width *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPenWidth,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(width)))
}

func GdipSetPenWidth(pen *GpPen, width REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenWidth,
		uintptr(unsafe.Pointer(pen)),
		REALbits(width))
}

func GdipClonePen(pen *GpPen, clonepen **GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipClonePen,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(clonepen)))
}

func GdipCreatePen2(brush *GpBrush, width REAL, unit Unit,
	pen **GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePen2,
		uintptr(unsafe.Pointer(brush)),
		uintptr(REALbits(width)),
		uintptr(unit),
		uintptr(unsafe.Pointer(pen)))
}

func GdipConvertToEmfPlusToStream(
	refGraphics *GpGraphics,
	metafile *GpImage,
	conversionFailureFlag *INT,
	stream *IStream,
	emfType EmfType,
	description string,
	out_metafile **GpImage) (status Status, err error) {

	return gdiplusSyscall(procGdipConvertToEmfPlusToStream,
		uintptr(unsafe.Pointer(refGraphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(conversionFailureFlag)),
		uintptr(unsafe.Pointer(stream)),
		uintptr(emfType),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(out_metafile)))
}

func GdipConvertToEmfPlusToFile(
	refGraphics *GpGraphics,
	metafile *GpImage,
	conversionFailureFlag *INT,
	filename string,
	emfType EmfType,
	description string,
	out_metafile **GpImage) (status Status, err error) {

	return gdiplusSyscall(procGdipConvertToEmfPlusToFile,
		uintptr(unsafe.Pointer(refGraphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(conversionFailureFlag)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))),
		uintptr(emfType),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(out_metafile)))
}

func GdipConvertToEmfPlus(
	refGraphics *GpGraphics,
	metafile *GpImage,
	conversionFailureFlag *INT,
	emfType EmfType,
	description string,
	out_metafile **GpImage) (status Status, err error) {

	return gdiplusSyscall(procGdipConvertToEmfPlus,
		uintptr(unsafe.Pointer(refGraphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(conversionFailureFlag)),
		uintptr(emfType),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(out_metafile)))

}

func GdipEmfToWmfBits(
	hemf HENHMETAFILE,
	cbData16 UINT,
	pData16 LPBYTE,
	iMapMode INT,
	eFlags INT) (length UINT, err error) {

	rs, _, e := procGdipEmfToWmfBits.Call(uintptr(hemf), uintptr(cbData16),
		uintptr(unsafe.Pointer(pData16)),
		uintptr(iMapMode), uintptr(eFlags))

	return UINT(rs), e
}

func GdipGetMetafileDownLevelRasterizationLimit(metafile *GpImage,
	metafileRasterizationLimitDpi *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMetafileDownLevelRasterizationLimit,
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(metafileRasterizationLimitDpi)))
}

func GdipSetMetafileDownLevelRasterizationLimit(metafile *GpImage,
	metafileRasterizationLimitDpi UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetMetafileDownLevelRasterizationLimit,
		uintptr(unsafe.Pointer(metafile)),
		uintptr(metafileRasterizationLimitDpi))
}

func GdipPlayMetafileRecord(
	metafile *GpImage,
	recordType EmfPlusRecordType,
	flags UINT,
	data []BYTE) (status Status, err error) {
	return gdiplusSyscall(procGdipPlayMetafileRecord,
		uintptr(unsafe.Pointer(metafile)),
		uintptr(recordType), uintptr(flags),
		uintptr(INT(len(data))),
		uintptr(unsafe.Pointer(&data[0])))
}

func GdipGetHemfFromMetafile(metafile *GpImage,
	hEmf *HENHMETAFILE) (status Status, err error) {
	return gdiplusSyscall(procGdipGetHemfFromMetafile,
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(hEmf)))
}

func GdipGetMetafileHeaderFromMetafile(metafile *GpImage,
	header *MetafileHeader) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMetafileHeaderFromMetafile,
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(header)))
}

func GdipGetMetafileHeaderFromStream(
	stream *IStream,
	header *MetafileHeader) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMetafileHeaderFromStream,
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(header)))
}

func GdipGetMetafileHeaderFromEmf(
	hEmf HENHMETAFILE,
	header *MetafileHeader) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMetafileHeaderFromEmf,
		uintptr(hEmf),
		uintptr(unsafe.Pointer(header)))
}

func GdipGetMetafileHeaderFromWmf(
	hWmf HMETAFILE,
	wmfPlaceableFileHeader *WmfPlaceableFileHeader,
	header *MetafileHeader) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMetafileHeaderFromWmf,
		uintptr(hWmf),
		uintptr(unsafe.Pointer(wmfPlaceableFileHeader)),
		uintptr(unsafe.Pointer(header)))
}

func GdipRecordMetafileStreamI(
	stream *IStream,
	referenceHdc HDC,
	typ EmfType,
	frameRect *Rect,
	frameUnit MetafileFrameUnit,
	description string,
	metafile **GpImage) (status Status, err error) {

	return gdiplusSyscall(procGdipRecordMetafileStreamI,
		uintptr(unsafe.Pointer(stream)),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipRecordMetafileStream(
	stream *IStream,
	referenceHdc HDC,
	typ EmfType,
	frameRect *RectF,
	frameUnit MetafileFrameUnit,
	description string,
	metafile **GpImage) (status Status, err error) {

	return gdiplusSyscall(procGdipRecordMetafileStream,
		uintptr(unsafe.Pointer(stream)),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipRecordMetafileFileNameI(
	fileName string,
	referenceHdc HDC,
	typ EmfType,
	frameRect *Rect,
	frameUnit MetafileFrameUnit,
	description string,
	metafile **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipRecordMetafileFileNameI,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fileName))),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipRecordMetafileFileName(
	fileName string,
	referenceHdc HDC,
	typ EmfType,
	frameRect *RectF,
	frameUnit MetafileFrameUnit,
	description string,
	metafile **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipRecordMetafileFileName,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fileName))),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipRecordMetafileI(
	referenceHdc HDC,
	typ EmfType,
	frameRect *Rect,
	frameUnit MetafileFrameUnit,
	description string,
	metafile **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipRecordMetafileI,
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipRecordMetafile(
	referenceHdc HDC,
	typ EmfType,
	frameRect *RectF,
	frameUnit MetafileFrameUnit,
	description string,
	metafile **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipRecordMetafile,
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(description))),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipCreateMetafileFromStream(stream *IStream, metafile **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMetafileFromStream,
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipCreateMetafileFromEmf(hEmf HENHMETAFILE, deleteEmf BOOL,
	metafile **GpImage) (status Status, err error) {
	var deleteEmf_ int8
	if deleteEmf {
		deleteEmf_ = 1
	}
	return gdiplusSyscall(procGdipCreateMetafileFromEmf,
		uintptr(hEmf),
		uintptr(deleteEmf_),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipCreateMetafileFromWmf(hWmf HMETAFILE, deleteWmf BOOL,
	wmfPlaceableFileHeader *WmfPlaceableFileHeader,
	metafile **GpImage) (status Status, err error) {
	var deleteWmf_ int8
	if deleteWmf {
		deleteWmf_ = 1
	}
	return gdiplusSyscall(procGdipCreateMetafileFromWmf,
		uintptr(hWmf),
		uintptr(deleteWmf_),
		uintptr(unsafe.Pointer(wmfPlaceableFileHeader)),
		uintptr(unsafe.Pointer(metafile)))
}

func GdipIsMatrixEqual(matrix *GpMatrix, matrix2 *GpMatrix, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsMatrixEqual,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(matrix2)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsMatrixIdentity(matrix *GpMatrix, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsMatrixIdentity,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsMatrixInvertible(matrix *GpMatrix, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsMatrixInvertible,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(result)))
}

func GdipVectorTransformMatrixPointsI(matrix *GpMatrix, pts []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipVectorTransformMatrixPointsI,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(&pts[0])), uintptr(INT(len(pts))))
}

func GdipVectorTransformMatrixPoints(matrix *GpMatrix, pts []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipVectorTransformMatrixPoints,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(&pts[0])), uintptr(INT(len(pts))))
}

func GdipTransformMatrixPointsI(matrix *GpMatrix, pts []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipTransformMatrixPointsI,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(&pts[0])), uintptr(INT(len(pts))))
}

func GdipTransformMatrixPoints(matrix *GpMatrix, pts []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipTransformMatrixPoints,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(&pts[0])), uintptr(INT(len(pts))))
}

func GdipInvertMatrix(matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipInvertMatrix,
		uintptr(unsafe.Pointer(matrix)))
}

func GdipShearMatrix(matrix *GpMatrix, shearX, shearY REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipShearMatrix,
		uintptr(unsafe.Pointer(matrix)),
		REALbits(shearX), REALbits(shearY),
		uintptr(order))
}

func GdipRotateMatrix(matrix *GpMatrix, angle REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipRotateMatrix,
		uintptr(unsafe.Pointer(matrix)),
		REALbits(angle),
		uintptr(order))
}

func GdipScaleMatrix(matrix *GpMatrix, scaleX, scaleY REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipScaleMatrix,
		uintptr(unsafe.Pointer(matrix)),
		REALbits(scaleX), REALbits(scaleY),
		uintptr(order))
}

func GdipTranslateMatrix(matrix *GpMatrix, offsetX, offsetY REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateMatrix,
		uintptr(unsafe.Pointer(matrix)),
		REALbits(offsetX), REALbits(offsetY),
		uintptr(order))
}

func GdipMultiplyMatrix(matrix *GpMatrix, matrix2 *GpMatrix, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipMultiplyMatrix,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(matrix2)),
		uintptr(order))
}

func GdipSetMatrixElements(matrix *GpMatrix, m11, m12, m21, m22, dx,
	dy REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetMatrixElements,
		uintptr(unsafe.Pointer(matrix)),
		REALbits(m11), REALbits(m12), REALbits(m21), REALbits(m22),
		REALbits(dx), REALbits(dy))
}

func GdipGetMatrixElements(matrix *GpMatrix, ms []REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMatrixElements,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(&ms[0])))
}

func GdipCloneMatrix(matrix *GpMatrix, cloneMatrix **GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneMatrix,
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(cloneMatrix)))
}

func GdipCreateMatrix3I(rect *Rect, dstplg *Point, matrix **GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMatrix3I,
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(dstplg)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipCreateMatrix3(rect *RectF, dstplg *PointF, matrix **GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMatrix3,
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(dstplg)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipCreateMatrix2(m11, m12, m21, m22, dx,
	dy REAL, matrix **GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMatrix2,
		REALbits(m11), REALbits(m12), REALbits(m21), REALbits(m22),
		REALbits(dx), REALbits(dy),
		uintptr(unsafe.Pointer(matrix)))
}

//----------------------------------------------------------------------------
// AdjustableArrowCap APIs
//----------------------------------------------------------------------------

func GdipGetAdjustableArrowCapFillState(arrowCap *GpCustomLineCap, isFilled *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetAdjustableArrowCapFillState,
		uintptr(unsafe.Pointer(arrowCap)), uintptr(unsafe.Pointer(isFilled)))
}

func GdipSetAdjustableArrowCapFillState(arrowCap *GpCustomLineCap, isFilled BOOL) (status Status, err error) {
	var isFilled_ int8
	if isFilled {
		isFilled_ = 1
	}
	return gdiplusSyscall(procGdipSetAdjustableArrowCapFillState,
		uintptr(unsafe.Pointer(arrowCap)), uintptr(isFilled_))
}

func GdipGetAdjustableArrowCapMiddleInset(arrowCap *GpCustomLineCap, middleInset *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetAdjustableArrowCapMiddleInset,
		uintptr(unsafe.Pointer(arrowCap)), uintptr(unsafe.Pointer(middleInset)))
}

func GdipSetAdjustableArrowCapMiddleInset(arrowCap *GpCustomLineCap, middleInset REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetAdjustableArrowCapMiddleInset,
		uintptr(unsafe.Pointer(arrowCap)), REALbits(middleInset))
}

func GdipGetAdjustableArrowCapWidth(arrowCap *GpCustomLineCap, width *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetAdjustableArrowCapWidth,
		uintptr(unsafe.Pointer(arrowCap)), uintptr(unsafe.Pointer(width)))
}

func GdipSetAdjustableArrowCapWidth(arrowCap *GpCustomLineCap, width REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetAdjustableArrowCapWidth,
		uintptr(unsafe.Pointer(arrowCap)), REALbits(width))
}

func GdipGetAdjustableArrowCapHeight(arrowCap *GpCustomLineCap, height *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetAdjustableArrowCapHeight,
		uintptr(unsafe.Pointer(arrowCap)), uintptr(unsafe.Pointer(height)))
}

func GdipSetAdjustableArrowCapHeight(arrowCap *GpCustomLineCap, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetAdjustableArrowCapHeight,
		uintptr(unsafe.Pointer(arrowCap)), REALbits(height))
}

func GdipCreateAdjustableArrowCap(height, width REAL, isFilled BOOL,
	arrowCap **GpCustomLineCap) (status Status, err error) {

	var isFilled_ int8
	if isFilled {
		isFilled_ = 1
	}

	return gdiplusSyscall(procGdipCreateAdjustableArrowCap,
		REALbits(height), REALbits(width), uintptr(isFilled_),
		uintptr(unsafe.Pointer(arrowCap)))
}

func GdipCloneCustomLineCap(customCap *GpCustomLineCap,
	clonedCap **GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneCustomLineCap,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(clonedCap)))
}

func GdipGetCustomLineCapWidthScale(customCap *GpCustomLineCap,
	widthScale *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCustomLineCapWidthScale,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(widthScale)))
}

func GdipSetCustomLineCapWidthScale(customCap *GpCustomLineCap,
	widthScale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCustomLineCapWidthScale,
		uintptr(unsafe.Pointer(customCap)),
		REALbits(widthScale))
}

func GdipGetCustomLineCapBaseInset(customCap *GpCustomLineCap,
	inset *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCustomLineCapBaseInset,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(inset)))
}

func GdipSetCustomLineCapBaseInset(customCap *GpCustomLineCap,
	inset REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCustomLineCapBaseInset,
		uintptr(unsafe.Pointer(customCap)),
		REALbits(inset))
}

func GdipGetCustomLineCapBaseCap(customCap *GpCustomLineCap,
	baseCap *LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCustomLineCapBaseCap,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(baseCap)))
}

func GdipSetCustomLineCapBaseCap(customCap *GpCustomLineCap,
	baseCap LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCustomLineCapBaseCap,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(baseCap))
}

func GdipGetCustomLineCapStrokeJoin(customCap *GpCustomLineCap,
	lineJoin *LineJoin) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCustomLineCapStrokeJoin,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(lineJoin)))
}

func GdipSetCustomLineCapStrokeJoin(customCap *GpCustomLineCap,
	lineJoin LineJoin) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCustomLineCapStrokeJoin,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(lineJoin))
}

func GdipGetCustomLineCapStrokeCaps(customCap *GpCustomLineCap,
	startCap, endCap *LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCustomLineCapStrokeCaps,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(startCap)),
		uintptr(unsafe.Pointer(endCap)))
}

func GdipSetPenCustomEndCap(pen *GpPen, customCap *GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenCustomEndCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
}

func GdipSetPenCustomStartCap(pen *GpPen, customCap *GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPenCustomStartCap,
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
}

func GdipSetCustomLineCapStrokeCaps(customCap *GpCustomLineCap,
	startCap, endCap LineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCustomLineCapStrokeCaps,
		uintptr(unsafe.Pointer(customCap)),
		uintptr(startCap),
		uintptr(endCap))
}

func GdipDeleteCustomLineCap(customCap *GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteCustomLineCap,
		uintptr(unsafe.Pointer(customCap)))
}

func GdipCreateCustomLineCap(fillPath *GpPath, strokePath *GpPath,
	baseCap LineCap, baseInset REAL, customCap **GpCustomLineCap) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateCustomLineCap,
		uintptr(unsafe.Pointer(fillPath)),
		uintptr(unsafe.Pointer(strokePath)),
		uintptr(baseCap),
		REALbits(baseInset),
		uintptr(unsafe.Pointer(customCap)))
}

func GdipGetImageEncoders(numEncoders UINT, size UINT, decoders *ImageCodecInfo) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageEncoders,
		uintptr(numEncoders),
		uintptr(size),
		uintptr(unsafe.Pointer(decoders)))
}

func GdipGetImageEncodersSize(numEncoders *UINT, size *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageEncodersSize,
		uintptr(unsafe.Pointer(numEncoders)),
		uintptr(unsafe.Pointer(size)))
}

func GdipGetImageDecoders(numDecoders UINT, size UINT, decoders *ImageCodecInfo) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageDecoders,
		uintptr(numDecoders),
		uintptr(size),
		uintptr(unsafe.Pointer(decoders)))
}

func GdipGetImageDecodersSize(numDecoders *UINT, size *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageDecodersSize,
		uintptr(unsafe.Pointer(numDecoders)),
		uintptr(unsafe.Pointer(size)))
}

func GdipGetImageAttributesAdjustedPalette(imageattr *GpImageAttributes,
	colorPalette *ColorPalette, colorAdjustType ColorAdjustType) (status Status, err error) {

	return gdiplusSyscall(procGdipGetImageAttributesAdjustedPalette,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(unsafe.Pointer(colorPalette)),
		uintptr(colorAdjustType))
}

func GdipSetImageAttributesWrapMode(imageattr *GpImageAttributes,
	wrap WrapMode, argb ARGB, clamp BOOL) (status Status, err error) {
	var clamp_ int8
	if clamp {
		clamp_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesWrapMode,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(wrap),
		uintptr(argb),
		uintptr(clamp_))
}

func GdipSetImageAttributesRemapTable(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL, maps []ColorMap) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	var mapsize uintptr
	var mapsptr uintptr
	if maps != nil {
		mapsize = uintptr(INT(len(maps)))
		mapsptr = uintptr(unsafe.Pointer(&maps[0]))
	}
	return gdiplusSyscall(procGdipSetImageAttributesRemapTable,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		mapsize,
		mapsptr)
}

func GdipSetImageAttributesOutputChannelColorProfile(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL, colorProfileFilename *string) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}

	var colorProfileFilename_ uintptr
	if colorProfileFilename != nil {
		colorProfileFilename_ = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(*colorProfileFilename)))
	}

	return gdiplusSyscall(procGdipSetImageAttributesOutputChannelColorProfile,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		colorProfileFilename_)
}

func GdipSetImageAttributesOutputChannel(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL, channelFlags ColorChannelFlags) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesOutputChannel,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		uintptr(channelFlags))
}

func GdipSetImageAttributesColorKeys(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL, colorLow, colorHigh ARGB) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesColorKeys,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		uintptr(colorLow), uintptr(colorHigh))
}

func GdipSetImageAttributesNoOp(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesNoOp,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_))
}

func GdipSetImageAttributesGamma(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL, gamma REAL) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesGamma,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		REALbits(gamma))
}

func GdipSetImageAttributesThreshold(imageattr *GpImageAttributes,
	typ ColorAdjustType, enableFlag BOOL, threshold REAL) (status Status, err error) {
	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesThreshold,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		REALbits(threshold))
}

func GdipSetImageAttributesColorMatrix(imageattr *GpImageAttributes,
	typ ColorAdjustType,
	enableFlag BOOL,
	colorMatrix *ColorMatrix,
	grayMatrix *ColorMatrix,
	flags ColorMatrixFlags) (status Status, err error) {

	var enableFlag_ int8
	if enableFlag {
		enableFlag_ = 1
	}
	return gdiplusSyscall(procGdipSetImageAttributesColorMatrix,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag_),
		uintptr(unsafe.Pointer(colorMatrix)),
		uintptr(unsafe.Pointer(grayMatrix)),
		uintptr(flags))
}

func GdipResetImageAttributes(imageattr *GpImageAttributes, typ ColorAdjustType) (status Status, err error) {
	return gdiplusSyscall(procGdipResetImageAttributes,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ))
}

func GdipSetImageAttributesToIdentity(imageattr *GpImageAttributes, typ ColorAdjustType) (status Status, err error) {
	return gdiplusSyscall(procGdipSetImageAttributesToIdentity,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ))
}

func GdipCloneImageAttributes(imageattr *GpImageAttributes, cloneImageattr **GpImageAttributes) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneImageAttributes,
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(unsafe.Pointer(cloneImageattr)))
}

func GdipGetLineSpacing(family *GpFontFamily, style INT, lineSpacing *UINT16) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineSpacing,
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(lineSpacing)))
}

func GdipGetCellDescent(family *GpFontFamily, style INT, cellDescent *UINT16) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCellDescent,
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(cellDescent)))
}

func GdipGetCellAscent(family *GpFontFamily, style INT, cellAscent *UINT16) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCellAscent,
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(cellAscent)))
}

func GdipGetEmHeight(family *GpFontFamily, style INT, emHeight *UINT16) (status Status, err error) {
	return gdiplusSyscall(procGdipGetEmHeight,
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(emHeight)))
}

func GdipIsStyleAvailable(family *GpFontFamily, style INT, isStyleAvailable *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFamilyName,
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(isStyleAvailable)))
}

func GdipGetFamilyName(family *GpFontFamily, name LPWSTR, language LANGID) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFamilyName,
		uintptr(unsafe.Pointer(family)),
		uintptr(unsafe.Pointer(name)),
		uintptr(language))
}

func GdipGetGenericFontFamilyMonospace(nativeFamily **GpFontFamily) (status Status, err error) {
	return gdiplusSyscall(procGdipGetGenericFontFamilyMonospace,
		uintptr(unsafe.Pointer(nativeFamily)))
}

func GdipGetGenericFontFamilySerif(nativeFamily **GpFontFamily) (status Status, err error) {
	return gdiplusSyscall(procGdipGetGenericFontFamilySerif,
		uintptr(unsafe.Pointer(nativeFamily)))
}

func GdipGetFontHeightGivenDPI(font *GpFont, dpi REAL, height *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontHeightGivenDPI,
		uintptr(unsafe.Pointer(font)),
		REALbits(dpi),
		uintptr(unsafe.Pointer(height)))
}

func GdipGetFontHeight(font *GpFont, graphics *GpGraphics, height *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontHeight,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(height)))
}

func GdipGetFontUnit(font *GpFont, unit *Unit) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontUnit,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(unit)))
}

func GdipGetFontSize(font *GpFont, size *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontSize,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(size)))
}

func GdipGetFontStyle(font *GpFont, style *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontStyle,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(style)))
}

func GdipGetFamily(font *GpFont, family **GpFontFamily) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFamily,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(family)))
}

func GdipCloneFont(font *GpFont, cloneFont **GpFont) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneFont,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(cloneFont)))
}

func GdipGetLogFontW(font *GpFont, graphics *GpGraphics, logfontW *LOGFONTW) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLogFontW,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(logfontW)))
}

func GdipGetLogFontA(font *GpFont, graphics *GpGraphics, logfontA *LOGFONTA) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLogFontA,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(logfontA)))
}

func GdipCreateFontFromLogfontW(hdc HDC, logfont *LOGFONTW, font **GpFont) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFontFromLogfontW,
		uintptr(hdc),
		uintptr(unsafe.Pointer(logfont)),
		uintptr(unsafe.Pointer(font)))
}

func GdipCreateFontFromLogfontA(hdc HDC, logfont *LOGFONTA, font **GpFont) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFontFromLogfontA,
		uintptr(hdc),
		uintptr(unsafe.Pointer(logfont)),
		uintptr(unsafe.Pointer(font)))
}

func GdipGetLineRectI(brush *GpBrush, rect *Rect) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineRectI,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetLineRect(brush *GpBrush, rect *RectF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineRect,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetLineColors(brush *GpBrush, colors []ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineColors,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&colors[0])))
}

func GdipSetLineColors(texture *GpBrush, color1, color2 ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLineColors,
		uintptr(unsafe.Pointer(texture)),
		uintptr(color1),
		uintptr(color2))
}

func GdipCreateLineBrushFromRectWithAngleI(rect *Rect,
	color1, color2 ARGB, angle REAL, isAngleScalable BOOL, wrapMode WrapMode,
	lineGradient **GpBrush) (status Status, err error) {

	var isAngleScalable_ int8
	if isAngleScalable {
		isAngleScalable_ = 1
	}

	return gdiplusSyscall(procGdipCreateLineBrushFromRectWithAngleI,
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1), uintptr(color2),
		REALbits(angle), uintptr(isAngleScalable_),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
}

func GdipCreateLineBrushFromRectWithAngle(rect *RectF,
	color1, color2 ARGB, angle REAL, isAngleScalable BOOL, wrapMode WrapMode,
	lineGradient **GpBrush) (status Status, err error) {

	var isAngleScalable_ int8
	if isAngleScalable {
		isAngleScalable_ = 1
	}

	return gdiplusSyscall(procGdipCreateLineBrushFromRectWithAngle,
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1), uintptr(color2),
		REALbits(angle), uintptr(isAngleScalable_),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
}

func GdipCreateLineBrushFromRectI(rect *Rect,
	color1, color2 ARGB, mode LinearGradientMode, wrapMode WrapMode,
	lineGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateLineBrushFromRectI,
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1), uintptr(color2),
		uintptr(mode),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
}

func GdipCreateLineBrushFromRect(rect *RectF,
	color1, color2 ARGB, mode LinearGradientMode, wrapMode WrapMode,
	lineGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateLineBrushFromRect,
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1), uintptr(color2),
		uintptr(mode),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
}

func GdipCreateLineBrushI(point1, point2 *Point,
	color1, color2 ARGB, wrapMode WrapMode,
	lineGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateLineBrushI,
		uintptr(unsafe.Pointer(point1)),
		uintptr(unsafe.Pointer(point2)),
		uintptr(color1), uintptr(color2), uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
}

func GdipCreateLineBrush(point1, point2 *PointF,
	color1, color2 ARGB, wrapMode WrapMode,
	lineGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateLineBrush,
		uintptr(unsafe.Pointer(point1)),
		uintptr(unsafe.Pointer(point2)),
		uintptr(color1), uintptr(color2), uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
}

func GdipGetTextureImage(texture *GpBrush, nativeImage **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipGetTextureImage,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(nativeImage)))
}

func GdipGetTextureWrapMode(texture *GpBrush, wrapMode *WrapMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetTextureWrapMode,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(wrapMode)))
}

func GdipSetTextureWrapMode(texture *GpBrush, wrapMode WrapMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetTextureWrapMode,
		uintptr(unsafe.Pointer(texture)),
		uintptr(wrapMode))
}

func GdipRotateTextureTransform(texture *GpBrush, angle REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipRotateTextureTransform,
		uintptr(unsafe.Pointer(texture)),
		REALbits(angle),
		uintptr(order))
}

func GdipScaleTextureTransform(texture *GpBrush, sx, sy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipScaleTextureTransform,
		uintptr(unsafe.Pointer(texture)),
		REALbits(sx), REALbits(sy),
		uintptr(order))
}

func GdipTranslateTextureTransform(texture *GpBrush, dx, dy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateTextureTransform,
		uintptr(unsafe.Pointer(texture)),
		REALbits(dx), REALbits(dy),
		uintptr(order))
}

func GdipMultiplyTextureTransform(texture *GpBrush, nativeMatrix *GpMatrix, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipMultiplyTextureTransform,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(nativeMatrix)),
		uintptr(order))
}

func GdipResetTextureTransform(texture *GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipResetTextureTransform,
		uintptr(unsafe.Pointer(texture)))
}

func GdipGetTextureTransform(texture *GpBrush, nativeMatrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetTextureTransform,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(nativeMatrix)))
}

func GdipSetTextureTransform(texture *GpBrush, nativeMatrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipSetTextureTransform,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(nativeMatrix)))
}

func GdipCreateTextureIAI(image *GpImage,
	imageAttributes *GpImageAttributes,
	x, y, width, height INT, texture **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateTextureIAI,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(unsafe.Pointer(texture)))
}

func GdipCreateTexture2I(image *GpImage, wrapmode WrapMode,
	x, y, width, height INT,
	texture **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateTexture2I,
		uintptr(unsafe.Pointer(image)),
		uintptr(wrapmode),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(unsafe.Pointer(texture)))
}

func GdipCreateTextureIA(image *GpImage,
	imageAttributes *GpImageAttributes,
	x, y, width, height REAL, texture **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateTextureIA,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(imageAttributes)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		uintptr(unsafe.Pointer(texture)))
}

func GdipCreateTexture2(image *GpImage, wrapmode WrapMode,
	x, y, width, height REAL,
	texture **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateTexture2,
		uintptr(unsafe.Pointer(image)),
		uintptr(wrapmode),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		uintptr(unsafe.Pointer(texture)))
}

func GdipCreateTexture(image *GpImage, wrapmode WrapMode, texture **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateTexture,
		uintptr(unsafe.Pointer(image)),
		uintptr(wrapmode),
		uintptr(unsafe.Pointer(texture)))
}

func GdipBitmapSetResolution(bitmap *GpImage, xdpi, ydpi REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapSetResolution,
		uintptr(unsafe.Pointer(bitmap)),
		REALbits(xdpi), REALbits(ydpi))
}

func GdipBitmapGetHistogramSize(format HistogramFormat, numberOfEntries *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapGetHistogramSize,
		uintptr(format),
		uintptr(unsafe.Pointer(numberOfEntries)))
}

func GdipBitmapGetHistogram(
	bitmap *GpImage,
	format HistogramFormat,
	numberOfEntries UINT,
	channel0 []UINT,
	channel1 []UINT,
	channel2 []UINT,
	channel3 []UINT) (status Status, err error) {

	return gdiplusSyscall(procGdipBitmapGetHistogram,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(format),
		uintptr(numberOfEntries),
		uintptr(unsafe.Pointer(&channel0[0])),
		uintptr(unsafe.Pointer(&channel1[0])),
		uintptr(unsafe.Pointer(&channel2[0])),
		uintptr(unsafe.Pointer(&channel3[0])))
}

func GdipBitmapApplyEffect(
	bitmap *GpImage,
	effect *CGpEffect,
	roi *RECT,
	useAuxData BOOL,
	auxData *DATA_PTR,
	auxDataSize *INT) (status Status, err error) {
	var useAuxData_ int8
	if useAuxData {
		useAuxData_ = 1
	}
	return gdiplusSyscall(procGdipBitmapApplyEffect,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(roi)),
		uintptr(useAuxData_),
		uintptr(unsafe.Pointer(auxData)),
		uintptr(unsafe.Pointer(auxDataSize)))
}

func GdipBitmapCreateApplyEffect(
	inputBitmaps **GpImage,
	numInputs INT,
	effect *CGpEffect,
	roi *RECT,
	outputRect *RECT,
	outputBitmap **GpImage,
	useAuxData BOOL,
	auxData *DATA_PTR,
	auxDataSize *INT) (status Status, err error) {
	var useAuxData_ int8
	if useAuxData {
		useAuxData_ = 1
	}
	return gdiplusSyscall(procGdipBitmapCreateApplyEffect,
		uintptr(unsafe.Pointer(inputBitmaps)),
		uintptr(numInputs),
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(roi)),
		uintptr(unsafe.Pointer(outputRect)),
		uintptr(unsafe.Pointer(outputBitmap)),
		uintptr(useAuxData_),
		uintptr(unsafe.Pointer(auxData)),
		uintptr(unsafe.Pointer(auxDataSize)))

}

func GdipInitializePalette(palette *ColorPalette, palettetype PaletteType,
	optimalColors INT, useTransparentColor BOOL, bitmap *GpImage) (status Status, err error) {
	var useTransparentColor_ int8
	if useTransparentColor {
		useTransparentColor_ = 1
	}
	return gdiplusSyscall(procGdipInitializePalette,
		uintptr(unsafe.Pointer(palette)),
		uintptr(palettetype),
		uintptr(optimalColors),
		uintptr(useTransparentColor_),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipBitmapConvertFormat(pInputBitmap *GpImage, format PixelFormat,
	dithertype DitherType, palettetype PaletteType, palette *ColorPalette,
	alphaThresholdPercent REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapConvertFormat,
		uintptr(unsafe.Pointer(pInputBitmap)),
		uintptr(format),
		uintptr(dithertype),
		uintptr(palettetype),
		uintptr(unsafe.Pointer(palette)),
		REALbits(alphaThresholdPercent))
}

func GdipImageSetAbort(pImage *GpImage, pIAbort *GdiplusAbort) (status Status, err error) {
	return gdiplusSyscall(procGdipImageSetAbort,
		uintptr(unsafe.Pointer(pImage)),
		uintptr(unsafe.Pointer(pIAbort)))
}

func GdipBitmapSetPixel(bitmap *GpImage, x, y INT, argb ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapSetPixel,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(x), uintptr(y),
		uintptr(argb))
}

func GdipBitmapGetPixel(bitmap *GpImage, x, y INT, argb *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapGetPixel,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(argb)))
}

func GdipBitmapUnlockBits(bitmap *GpImage, lockedBitmapData *BitmapData) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapUnlockBits,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(lockedBitmapData)))
}

func GdipBitmapLockBits(bitmap *GpImage, rect *Rect,
	flags UINT, format PixelFormat, lockedBitmapData *BitmapData) (status Status, err error) {
	return gdiplusSyscall(procGdipBitmapLockBits,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(flags),
		uintptr(format),
		uintptr(unsafe.Pointer(lockedBitmapData)))
}

func GdipCloneBitmapArea(x, y, width, height REAL, format PixelFormat,
	srcBitmap *GpImage, dstBitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneBitmapArea,
		REALbits(x), REALbits(y), REALbits(width), REALbits(height), uintptr(format),
		uintptr(unsafe.Pointer(srcBitmap)),
		uintptr(unsafe.Pointer(dstBitmap)))
}

func GdipCloneBitmapAreaI(x, y, width, height INT, format PixelFormat,
	srcBitmap *GpImage, dstBitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneBitmapAreaI,
		uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(format),
		uintptr(unsafe.Pointer(srcBitmap)),
		uintptr(unsafe.Pointer(dstBitmap)))
}

func GdipCreateHICONFromBitmap(bitmap *GpImage, hiconReturn *HICON) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateHICONFromBitmap,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hiconReturn)))
}

func GdipCreateHBITMAPFromBitmap(bitmap *GpImage, hbmReturn *HBITMAP,
	background ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateHBITMAPFromBitmap,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hbmReturn)),
		uintptr(background))
}

func GdipCreateBitmapFromResource(hInstance HINSTANCE, bitmapName string,
	bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromResource,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(bitmapName))),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromHICON(hicon HICON,
	bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromHICON,
		uintptr(hicon),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE,
	bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromHBITMAP,
		uintptr(hbm),
		uintptr(unsafe.Pointer(hpal)),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromGdiDib(gdiBitmapInfo *BITMAPINFO, gdiBitmapData DATA_PTR,
	bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromGdiDib,
		uintptr(unsafe.Pointer(gdiBitmapInfo)),
		uintptr(gdiBitmapData),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromDirectDrawSurface(surface *IDirectDrawSurface7, bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromDirectDrawSurface,
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromGraphics(width, height INT, target *GpGraphics, bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromGraphics,
		uintptr(width), uintptr(height),
		uintptr(unsafe.Pointer(target)),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromScan0(width, height, stride INT,
	format PixelFormat, scan0 *BYTE, bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromScan0,
		uintptr(width), uintptr(height), uintptr(stride), uintptr(format),
		uintptr(unsafe.Pointer(scan0)),
		uintptr(unsafe.Pointer(bitmap)))
}

func GdipGetImageItemData(image *GpImage, item *ImageItemData) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageItemData,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
}

func GdipFindNextImageItem(image *GpImage, item *ImageItemData) (status Status, err error) {
	return gdiplusSyscall(procGdipFindNextImageItem,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
}

func GdipFindFirstImageItem(image *GpImage, item *ImageItemData) (status Status, err error) {
	return gdiplusSyscall(procGdipFindFirstImageItem,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
}

func GdipSetPropertyItem(image *GpImage, item *PropertyItem) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPropertyItem,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
}

func GdipRemovePropertyItem(image *GpImage, propId PROPID) (status Status, err error) {
	return gdiplusSyscall(procGdipRemovePropertyItem,
		uintptr(unsafe.Pointer(image)),
		uintptr(propId))
}

func GdipGetAllPropertyItems(image *GpImage, totalBufferSize, numProperties UINT, allItems *PropertyItem) (status Status, err error) {
	return gdiplusSyscall(procGdipGetAllPropertyItems,
		uintptr(unsafe.Pointer(image)),
		uintptr(totalBufferSize),
		uintptr(numProperties),
		uintptr(unsafe.Pointer(allItems)))
}

func GdipGetPropertySize(image *GpImage, totalBufferSize, numProperties *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPropertySize,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(totalBufferSize)),
		uintptr(unsafe.Pointer(numProperties)))
}

//GpImage *image, PROPID propId,UINT propSize,
//                    PropertyItem* buffer
func GdipGetPropertyItem(image *GpImage, propId PROPID, propSize UINT, buffer *PropertyItem) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPropertyItem,
		uintptr(unsafe.Pointer(image)),
		uintptr(propId), uintptr(propSize),
		uintptr(unsafe.Pointer(buffer)))
}

func GdipGetPropertyItemSize(image *GpImage, propId PROPID, size *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPropertyItemSize,
		uintptr(unsafe.Pointer(image)),
		uintptr(propId),
		uintptr(unsafe.Pointer(size)))
}

//UINT numOfProperty, PROPID* list
func GdipGetPropertyIdList(image *GpImage, list []PROPID) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPropertyIdList,
		uintptr(unsafe.Pointer(image)),
		uintptr(INT(len(list))),
		uintptr(unsafe.Pointer(&list[0])))
}

func GdipGetPropertyCount(image *GpImage, numProperty *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPropertyCount,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(numProperty)))
}

func GdipImageRotateFlip(image *GpImage, rotateFlipType RotateFlipType) (status Status, err error) {
	return gdiplusSyscall(procGdipImageRotateFlip,
		uintptr(unsafe.Pointer(image)),
		uintptr(rotateFlipType))
}

func GdipImageSelectActiveFrame(image *GpImage, dimensionID *GUID, frameIndex UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipImageSelectActiveFrame,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dimensionID)),
		uintptr(frameIndex))
}

func GdipImageGetFrameCount(image *GpImage, dimensionID *GUID, count *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipImageGetFrameCount,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dimensionID)),
		uintptr(unsafe.Pointer(count)))
}

func GdipImageGetFrameDimensionsList(image *GpImage, dimensionIDs []GUID) (status Status, err error) {
	return gdiplusSyscall(procGdipImageGetFrameDimensionsList,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(&dimensionIDs[0])),
		uintptr(INT(len(dimensionIDs))))
}

func GdipImageGetFrameDimensionsCount(image *GpImage, count *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipImageGetFrameDimensionsCount,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetImageThumbnail(image *GpImage, thumbWidth, thumbHeight UINT,
	thumbImage **GpImage, callback GetThumbnailImageAbort, callbackData DATA_PTR) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipGetImageThumbnail,
		uintptr(unsafe.Pointer(image)),
		uintptr(thumbWidth), uintptr(thumbHeight),
		uintptr(unsafe.Pointer(thumbImage)),
		callbackptr, uintptr(callbackData))
}

func GdipSetImagePalette(image *GpImage, palette *ColorPalette) (status Status, err error) {
	return gdiplusSyscall(procGdipSetImagePalette,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(palette)))
}

func GdipGetImagePalette(image *GpImage, palette *ColorPalette, size INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImagePalette,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(palette)),
		uintptr(size))
}

func GdipGetImagePaletteSize(image *GpImage, size *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImagePaletteSize,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(size)))
}

func GdipGetImagePixelFormat(image *GpImage, format *PixelFormat) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImagePixelFormat,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(format)))
}

func GdipGetImageRawFormat(image *GpImage, format *GUID) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageRawFormat,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(format)))
}

func GdipGetImageFlags(image *GpImage, flags *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageFlags,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(flags)))
}

func GdipGetImageVerticalResolution(image *GpImage, resolution *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageVerticalResolution,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(resolution)))
}

func GdipGetImageHorizontalResolution(image *GpImage, resolution *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageHorizontalResolution,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(resolution)))
}

func GdipGetImageHeight(image *GpImage, height *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageHeight,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(height)))
}

func GdipGetImageWidth(image *GpImage, width *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageWidth,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(width)))
}

func GdipGetImageBounds(image *GpImage, srcRect *RectF, srcUnit *Unit) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageBounds,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(unsafe.Pointer(srcUnit)))
}

func GdipGetImageDimension(image *GpImage, width, height *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageDimension,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(width)),
		uintptr(unsafe.Pointer(height)))
}

func GdipGetImageType(image *GpImage, typ *ImageType) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageType,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(typ)))
}

func GdipSaveAddImage(image *GpImage, newImage *GpImage, encoderParams *EncoderParameters) (status Status, err error) {
	return gdiplusSyscall(procGdipSaveAddImage,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(newImage)),
		uintptr(unsafe.Pointer(encoderParams)))
}

func GdipSaveAdd(image *GpImage, encoderParams *EncoderParameters) (status Status, err error) {
	return gdiplusSyscall(procGdipSaveAdd,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(encoderParams)))
}

func GdipSaveImageToFile(image *GpImage, filename string, clsidEncoder *CLSID, encoderParams *EncoderParameters) (status Status, err error) {
	return gdiplusSyscall(procGdipSaveImageToFile,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(encoderParams)))
}

func GdipGetEncoderParameterList(image *GpImage, clsidEncoder *CLSID, buffer []EncoderParameters) (status Status, err error) {
	return gdiplusSyscall(procGdipGetEncoderParameterList,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(INT(len(buffer))))
}

func GdipGetEncoderParameterListSize(image *GpImage, clsidEncoder *CLSID, size *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetEncoderParameterListSize,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(size)))
}

func GdipCloneImage(image *GpImage, cloneImage **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneImage,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(cloneImage)))
}

func GdipSetPathGradientWrapMode(brush *GpBrush, wrapMode WrapMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientWrapMode,
		uintptr(unsafe.Pointer(brush)),
		uintptr(wrapMode))
}

func GdipSetLineWrapMode(brush *GpBrush, wrapMode WrapMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLineWrapMode,
		uintptr(unsafe.Pointer(brush)),
		uintptr(wrapMode))
}

func GdipGetPathGradientWrapMode(brush *GpBrush, wrapMode *WrapMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientWrapMode,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(wrapMode)))
}

func GdipGetLineWrapMode(brush *GpBrush, wrapMode *WrapMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineWrapMode,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(wrapMode)))
}

func GdipSetPathGradientFocusScales(brush *GpBrush, xScale, yScale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientFocusScales,
		uintptr(unsafe.Pointer(brush)),
		REALbits(xScale),
		REALbits(yScale))
}

func GdipGetPathGradientFocusScales(brush *GpBrush, xScale, yScale *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientFocusScales,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(xScale)),
		uintptr(unsafe.Pointer(yScale)))
}

func GdipRotatePathGradientTransform(brush *GpBrush, angle REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipRotatePathGradientTransform,
		uintptr(unsafe.Pointer(brush)),
		REALbits(angle),
		uintptr(order))
}

func GdipRotatePenTransform(pen *GpPen, angle REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipRotatePenTransform,
		uintptr(unsafe.Pointer(pen)),
		REALbits(angle),
		uintptr(order))
}

func GdipRotateLineTransform(brush *GpBrush, angle REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipRotateLineTransform,
		uintptr(unsafe.Pointer(brush)),
		REALbits(angle),
		uintptr(order))
}

func GdipScalePathGradientTransform(brush *GpBrush, sx, sy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipScalePathGradientTransform,
		uintptr(unsafe.Pointer(brush)),
		REALbits(sx), REALbits(sy),
		uintptr(order))
}

func GdipScalePenTransform(pen *GpPen, sx, sy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipScalePenTransform,
		uintptr(unsafe.Pointer(pen)),
		REALbits(sx), REALbits(sy),
		uintptr(order))
}

func GdipScaleLineTransform(brush *GpBrush, sx, sy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipScaleLineTransform,
		uintptr(unsafe.Pointer(brush)),
		REALbits(sx), REALbits(sy),
		uintptr(order))
}

func GdipTranslatePathGradientTransform(brush *GpBrush, dx, dy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslatePathGradientTransform,
		uintptr(unsafe.Pointer(brush)),
		REALbits(dx), REALbits(dy),
		uintptr(order))
}

func GdipTranslatePenTransform(pen *GpPen, dx, dy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslatePenTransform,
		uintptr(unsafe.Pointer(pen)),
		REALbits(dx), REALbits(dy),
		uintptr(order))
}

func GdipTranslateLineTransform(brush *GpBrush, dx, dy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateLineTransform,
		uintptr(unsafe.Pointer(brush)),
		REALbits(dx), REALbits(dy),
		uintptr(order))
}

func GdipMultiplyPathGradientTransform(brush *GpBrush, matrix *GpMatrix, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipMultiplyPathGradientTransform,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
}

func GdipMultiplyLineTransform(brush *GpBrush, matrix *GpMatrix, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipMultiplyLineTransform,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
}

func GdipResetPathGradientTransform(brush *GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipResetPathGradientTransform,
		uintptr(unsafe.Pointer(brush)))
}
func GdipResetLineTransform(brush *GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipResetLineTransform,
		uintptr(unsafe.Pointer(brush)))
}
func GdipSetPathGradientTransform(brush *GpBrush, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientTransform,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipSetLineTransform(brush *GpBrush, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLineTransform,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipGetPathGradientTransform(brush *GpBrush, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientTransform,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipGetLineTransform(brush *GpBrush, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineTransform,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipSetPathGradientLinearBlend(brush *GpBrush, focus, scale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientLinearBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(REALbits(focus)),
		uintptr(REALbits(scale)))
}

func GdipSetLineLinearBlend(brush *GpBrush, focus, scale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLineLinearBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(REALbits(focus)),
		uintptr(REALbits(scale)))
}

func GdipSetPathGradientSigmaBlend(brush *GpBrush, focus, scale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientSigmaBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(REALbits(focus)),
		uintptr(REALbits(scale)))
}

func GdipSetLineSigmaBlend(brush *GpBrush, focus, scale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLineSigmaBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(REALbits(focus)),
		uintptr(REALbits(scale)))
}

func GdipGetPathGradientPresetBlend(brush *GpBrush, blend []ARGB, positions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientPresetBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blend[0])),
		uintptr(unsafe.Pointer(&positions[0])),
		uintptr(count))
}

func GdipGetLinePresetBlend(brush *GpBrush, blend []ARGB, positions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLinePresetBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blend[0])),
		uintptr(unsafe.Pointer(&positions[0])),
		uintptr(count))
}

func GdipSetPathGradientPresetBlend(brush *GpBrush, blend []ARGB, positions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientPresetBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blend[0])),
		uintptr(unsafe.Pointer(&positions[0])),
		uintptr(count))
}

func GdipSetLinePresetBlend(brush *GpBrush, blend []ARGB, positions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLinePresetBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blend[0])),
		uintptr(unsafe.Pointer(&positions[0])),
		uintptr(count))
}

func GdipGetPathGradientPresetBlendCount(brush *GpBrush, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientPresetBlendCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetLinePresetBlendCount(brush *GpBrush, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLinePresetBlendCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
}

func GdipSetPathGradientBlend(brush *GpBrush, blendFactors, blendPositions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blendFactors[0])),
		uintptr(unsafe.Pointer(&blendPositions[0])),
		uintptr(count))
}

func GdipSetLineBlend(brush *GpBrush, blendFactors, blendPositions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetLineBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blendFactors[0])),
		uintptr(unsafe.Pointer(&blendPositions[0])),
		uintptr(count))
}

func GdipGetPathGradientBlend(brush *GpBrush, blendFactors, blendPositions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blendFactors[0])),
		uintptr(unsafe.Pointer(&blendPositions[0])),
		uintptr(count))
}

func GdipGetLineBlend(brush *GpBrush, blendFactors, blendPositions []REAL, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineBlend,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&blendFactors[0])),
		uintptr(unsafe.Pointer(&blendPositions[0])),
		uintptr(count))
}

func GdipGetPathGradientBlendCount(brush *GpBrush, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientBlendCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetLineBlendCount(brush *GpBrush, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineBlendCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetPathGradientGammaCorrection(brush *GpBrush, useGammaCorrection *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientGammaCorrection,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(useGammaCorrection)))
}

func GdipGetLineGammaCorrection(brush *GpBrush, useGammaCorrection *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetLineGammaCorrection,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(useGammaCorrection)))
}

func GdipSetPathGradientGammaCorrection(brush *GpBrush, useGammaCorrection BOOL) (status Status, err error) {
	var useGammaCorrection_ int8
	if useGammaCorrection {
		useGammaCorrection_ = 1
	}
	return gdiplusSyscall(procGdipSetPathGradientGammaCorrection,
		uintptr(unsafe.Pointer(brush)),
		uintptr(useGammaCorrection_))
}

func GdipSetLineGammaCorrection(brush *GpBrush, useGammaCorrection BOOL) (status Status, err error) {
	var useGammaCorrection_ int8
	if useGammaCorrection {
		useGammaCorrection_ = 1
	}
	return gdiplusSyscall(procGdipSetLineGammaCorrection,
		uintptr(unsafe.Pointer(brush)),
		uintptr(useGammaCorrection_))
}

func GdipGetPathGradientRectI(brush *GpBrush, rect *Rect) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientRectI,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetPathGradientRect(brush *GpBrush, rect *RectF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientRect,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipSetPathGradientCenterPointI(brush *GpBrush, point *Point) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientCenterPointI,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(point)))
}

func GdipSetPathGradientCenterPoint(brush *GpBrush, point *PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientCenterPoint,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(point)))
}

func GdipGetPathGradientCenterPointI(brush *GpBrush, point *Point) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientCenterPointI,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(point)))
}

func GdipGetPathGradientCenterPoint(brush *GpBrush, point *PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientCenterPoint,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(point)))
}

func GdipSetPathGradientPath(brush *GpBrush, path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientPath,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(path)))
}

func GdipGetPathGradientPath(brush *GpBrush, path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientPath,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(path)))
}

func GdipSetPathGradientSurroundColorsWithCount(brush *GpBrush, colors []ARGB, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientSurroundColorsWithCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&colors[0])),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetPathGradientSurroundColorsWithCount(brush *GpBrush, colors []ARGB, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientSurroundColorsWithCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(&colors[0])),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetPathGradientSurroundColorCount(brush *GpBrush, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientSurroundColorCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
}

func GdipGetPathGradientPointCount(brush *GpBrush, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientPointCount,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
}

func GdipSetPathGradientCenterColor(brush *GpBrush, argb ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathGradientCenterColor,
		uintptr(unsafe.Pointer(brush)),
		uintptr(argb))
}

func GdipGetPathGradientCenterColor(brush *GpBrush, argb *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathGradientCenterColor,
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(argb)))
}

func GdipCreatePathGradientFromPath(path *GpPath,
	polyGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePathGradientFromPath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(polyGradient)))
}

func GdipCreatePathGradientI(points []Point, wrapMode GpWrapMode,
	polyGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePathGradientI,
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(polyGradient)))
}

func GdipCreatePathGradient(points []PointF, wrapMode GpWrapMode,
	polyGradient **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePathGradient,
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(polyGradient)))
}

func GdipPathIterCopyData(iterator *GpPathIterator, resultCount *INT,
	points []PointF, types []BYTE, startIndex, endIndex INT) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterCopyData,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(unsafe.Pointer(&types[0])),
		uintptr(startIndex), uintptr(endIndex))
}

func GdipPathIterEnumerate(iterator *GpPathIterator, resultCount *INT,
	points []PointF, types []BYTE, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterEnumerate,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(unsafe.Pointer(&types[0])),
		uintptr(count))
}

func GdipPathIterRewind(iterator *GpPathIterator) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterRewind,
		uintptr(unsafe.Pointer(iterator)))
}

func GdipPathIterHasCurve(iterator *GpPathIterator, hasCurve *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterHasCurve,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(hasCurve)))
}

func GdipPathIterGetSubpathCount(iterator *GpPathIterator, resultCount *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterGetSubpathCount,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)))
}

func GdipPathIterGetCount(iterator *GpPathIterator, resultCount *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterGetCount,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)))
}

func GdipPathIterNextMarkerPath(iterator *GpPathIterator, resultCount *INT, path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterNextMarkerPath,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(path)))
}

func GdipPathIterNextMarker(iterator *GpPathIterator, resultCount *INT, startIndex, endIndex *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterNextMarker,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(startIndex)),
		uintptr(unsafe.Pointer(endIndex)))
}

func GdipPathIterNextPathType(iterator *GpPathIterator, resultCount *INT, pathType *BYTE, startIndex, endIndex *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterNextPathType,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(pathType)),
		uintptr(unsafe.Pointer(startIndex)),
		uintptr(unsafe.Pointer(endIndex)))
}

func GdipPathIterNextSubpathPath(iterator *GpPathIterator, resultCount *INT, path *GpPath, isClosed *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterNextSubpathPath,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(isClosed)))
}

func GdipPathIterNextSubpath(iterator *GpPathIterator, resultCount, startIndex, endIndex *INT, isClosed *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipPathIterNextSubpath,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(startIndex)),
		uintptr(unsafe.Pointer(endIndex)),
		uintptr(unsafe.Pointer(isClosed)))
}

func GdipDeletePathIter(iterator *GpPathIterator) (status Status, err error) {
	return gdiplusSyscall(procGdipDeletePathIter,
		uintptr(unsafe.Pointer(iterator)))
}

//GdipCreatePathIter(GpPathIterator **iterator, GpPath* path);
func GdipCreatePathIter(iterator **GpPathIterator, path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePathIter,
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(path)))
}

func GdipWarpPath(path *GpPath, matrix *GpMatrix, points []PointF,
	srcx, srcy, srcwidth, srcheight REAL,
	warpMode WarpMode, flatness REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipWarpPath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(&points[0])), uintptr(INT(len(points))),
		REALbits(srcx), REALbits(srcy), REALbits(srcwidth), REALbits(srcheight),
		uintptr(warpMode), REALbits(flatness))
}

func GdipWindingModeOutline(path *GpPath, matrix *GpMatrix, flatness REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipWindingModeOutline,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)),
		REALbits(flatness))
}

func GdipWidenPath(path *GpPath, pen *GpPen, matrix *GpMatrix, flatness REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipWidenPath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)),
		REALbits(flatness))
}

func GdipFlattenPath(path *GpPath, matrix *GpMatrix, flatness REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipFlattenPath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)),
		REALbits(flatness))
}

func GdipTransformPath(path *GpPath, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipTransformPath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)))
}

func GdipAddPathPath(path *GpPath, addingPath *GpPath, connect BOOL) (status Status, err error) {
	var connect_ int8
	if connect {
		connect_ = 1
	}
	return gdiplusSyscall(procGdipAddPathPath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(addingPath)),
		uintptr(connect_))
}

func GdipAddPathPolygonI(nativePath *GpPath, points []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathPolygonI,
		uintptr(unsafe.Pointer(nativePath)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathPolygon(nativePath *GpPath, points []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathPolygon,
		uintptr(unsafe.Pointer(nativePath)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathPieI(nativePath *GpPath, x, y, width, height INT,
	startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathPieI, uintptr(unsafe.Pointer(nativePath)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		REALbits(startAngle), REALbits(sweepAngle))
}

func GdipAddPathPie(nativePath *GpPath, x, y, width, height,
	startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathPie, uintptr(unsafe.Pointer(nativePath)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		REALbits(startAngle), REALbits(sweepAngle))
}

func GdipAddPathEllipseI(nativePath *GpPath, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathEllipseI, uintptr(unsafe.Pointer(nativePath)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func GdipAddPathEllipse(nativePath *GpPath, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathEllipse, uintptr(unsafe.Pointer(nativePath)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height))
}

func GdipAddPathRectangleI(path *GpPath, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathRectangleI, uintptr(unsafe.Pointer(path)), uintptr((x)), uintptr((y)), uintptr((width)), uintptr((height)))
}

func GdipAddPathRectanglesI(path *GpPath, rects []Rect) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathRectanglesI,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&rects[0])),
		uintptr(INT(len(rects))))
}

func GdipAddPathRectangles(path *GpPath, rects []RectF) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathRectangles,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&rects[0])),
		uintptr(INT(len(rects))))
}

func GdipAddPathClosedCurve2I(path *GpPath, points []Point, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathClosedCurve2I,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		REALbits(tension))
}

func GdipAddPathClosedCurveI(path *GpPath, points []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathClosedCurveI,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathClosedCurve2(path *GpPath, points []PointF, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathClosedCurve2,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		REALbits(tension))
}

func GdipAddPathClosedCurve(path *GpPath, points []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathClosedCurve,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathCurve3I(path *GpPath, points []Point, offset INT, numberOfSegments INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathCurve3I,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		uintptr(offset), uintptr(numberOfSegments),
		REALbits(tension))
}

func GdipAddPathCurve2I(path *GpPath, points []Point, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathCurve2I,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		REALbits(tension))
}

func GdipAddPathCurveI(path *GpPath, points []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathCurveI,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathCurve3(path *GpPath, points []PointF, offset INT, numberOfSegments INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathCurve3,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		uintptr(offset), uintptr(numberOfSegments),
		REALbits(tension))
}

func GdipAddPathCurve2(path *GpPath, points []PointF, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathCurve2,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))),
		REALbits(tension))
}

func GdipAddPathCurve(path *GpPath, points []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathCurve,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathBeziersI(path *GpPath, points []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathBeziersI,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathBezierI(nativePath *GpPath, x1, y1, x2, y2, x3, y3, x4, y4 INT) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathBezierI, uintptr(unsafe.Pointer(nativePath)),
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		uintptr(x3), uintptr(y3), uintptr(x4), uintptr(y4))
}

func GdipAddPathBeziers(path *GpPath, points []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathBeziers,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathBezier(nativePath *GpPath, x1, y1, x2, y2, x3, y3, x4, y4 REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathBezier, uintptr(unsafe.Pointer(nativePath)),
		REALbits(x1), REALbits(y1), REALbits(x2), REALbits(y2),
		REALbits(x3), REALbits(y3), REALbits(x4), REALbits(y4))
}

func GdipAddPathArcI(nativePath *GpPath, x, y, width, height INT, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathArcI, uintptr(unsafe.Pointer(nativePath)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		REALbits(startAngle), REALbits(sweepAngle))
}

func GdipAddPathArc(nativePath *GpPath, x, y, width, height, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathArc, uintptr(unsafe.Pointer(nativePath)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		REALbits(startAngle), REALbits(sweepAngle))
}

func GdipAddPathLineI(nativePath *GpPath, x1, y1, x2, y2 INT) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathLineI, uintptr(unsafe.Pointer(nativePath)), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func GdipAddPathLine2I(path *GpPath, points []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathLine2I,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipAddPathLine2(path *GpPath, points []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathLine2,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipGetPathLastPoint(path *GpPath, point *PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathLastPoint,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(point)))
}

func GdipSetPathMarker(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathMarker,
		uintptr(unsafe.Pointer(path)))
}
func GdipClearPathMarkers(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipClearPathMarkers,
		uintptr(unsafe.Pointer(path)))
}
func GdipReversePath(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipReversePath,
		uintptr(unsafe.Pointer(path)))
}

func GdipClosePathFigures(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipClosePathFigures,
		uintptr(unsafe.Pointer(path)))
}

func GdipClosePathFigure(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipClosePathFigure,
		uintptr(unsafe.Pointer(path)))
}

func GdipStartPathFigure(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipStartPathFigure,
		uintptr(unsafe.Pointer(path)))
}

//func GdipAddPathCurve(path *GpPath, points []PointF) (status Status, err error) {
//	return gdiplusSyscall(procGdipAddPathCurve,
//		uintptr(unsafe.Pointer(path)),
//		uintptr(unsafe.Pointer(&points[0])),
//		uintptr(INT(len(points))))
//}

//Status WINGDIPAPI
//GdipGetPathData(GpPath *path, GpPathData* pathData);
func GdipGetPathData(path *GpPath, pathData *PathData) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathData,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(pathData)))
}

func GdipSetPathFillMode(path *GpPath, fillmode GpFillMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPathFillMode,
		uintptr(unsafe.Pointer(path)),
		uintptr(fillmode))
}

//GpPath *path, GpFillMode *fillmode
func GdipGetPathFillMode(path *GpPath, fillmode *GpFillMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathFillMode,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(fillmode)))
}

//Status WINGDIPAPI
//GdipGetPathPoints(GpPath*, GpPointF* points, INT count);
func GdipGetPathPoints(path *GpPath, points []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathPoints,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

func GdipGetPathPointsI(path *GpPath, points []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathPointsI,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(INT(len(points))))
}

//Status WINGDIPAPI
//GdipGetPathTypes(GpPath* path, BYTE* types, INT count);
func GdipGetPathTypes(path *GpPath, types []BYTE) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathTypes,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&types[0])),
		uintptr(INT(len(types))))
}

//GpPath* path, INT* count
func GdipGetPointCount(path *GpPath, count *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPointCount,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(count)))
}

func GdipResetPath(path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipResetPath,
		uintptr(unsafe.Pointer(path)))
}

func GdipClonePath(path *GpPath, clonePath **GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipClonePath,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(clonePath)))
}

func GdipCreatePath2I(points []Point, types []byte, fillMode FillMode, nativePath **GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePath2I,
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(unsafe.Pointer(&types[0])),
		uintptr(INT(len(types))),
		uintptr(fillMode),
		uintptr(unsafe.Pointer(nativePath)))
}

func GdipCreatePath2(points []PointF, types []byte, fillMode FillMode, nativePath **GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePath2,
		uintptr(unsafe.Pointer(&points[0])),
		uintptr(unsafe.Pointer(&types[0])),
		uintptr(INT(len(types))),
		uintptr(fillMode),
		uintptr(unsafe.Pointer(nativePath)))
}

func GdipIsOutlineVisiblePathPointI(path *GpPath, x, y INT, pen *GpPen, graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsOutlineVisiblePathPointI,
		uintptr(unsafe.Pointer(path)),
		uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsOutlineVisiblePathPoint(path *GpPath, x, y REAL, pen *GpPen, graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsOutlineVisiblePathPoint,
		uintptr(unsafe.Pointer(path)),
		REALbits(x), REALbits(y),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsVisiblePathPointI(path *GpPath, x, y INT, graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisiblePathPointI,
		uintptr(unsafe.Pointer(path)),
		uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipIsVisiblePathPoint(path *GpPath, x, y REAL, graphics *GpGraphics, result *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisiblePathPoint,
		uintptr(unsafe.Pointer(path)),
		REALbits(x), REALbits(y),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
}

func GdipGetPathWorldBoundsI(path *GpPath, bounds *Rect, matrix *GpMatrix, pen *GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathWorldBoundsI,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(bounds)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pen)))
}

func GdipGetPathWorldBounds(path *GpPath, bounds *RectF, matrix *GpMatrix, pen *GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPathWorldBounds,
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(bounds)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pen)))
}

func GdipCreateHalftonePalette() (hpalette HPALETTE, err error) {
	rs, _, e := syscall.Syscall(procGdipCreateHalftonePalette.Addr(), 0, 0, 0, 0)
	if rs == 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = syscall.EINVAL
		}
	} else {
		hpalette = HPALETTE(unsafe.Pointer(rs))
	}
	return
}

func GdipComment(nativeGraphics *GpGraphics, data []byte) (status Status, err error) {
	//UINT sizeData, GDIPCONST BYTE * data
	return gdiplusSyscall(procGdipComment,
		uintptr(UINT(len(data))),
		uintptr(unsafe.Pointer(&data[0])))
}

func GdipEndContainer(nativeGraphics *GpGraphics, state GraphicsContainer) (status Status, err error) {
	return gdiplusSyscall(procGdipEndContainer,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(state))
}

func GdipBeginContainer2(nativeGraphics *GpGraphics, state *GraphicsContainer) (status Status, err error) {
	return gdiplusSyscall(procGdipBeginContainer2,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(state)))
}

func GdipBeginContainerI(nativeGraphics *GpGraphics, dstrect,
	srcrect *Rect, unit Unit, state *GraphicsContainer) (status Status, err error) {
	return gdiplusSyscall(procGdipBeginContainerI,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(dstrect)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unit),
		uintptr(unsafe.Pointer(state)))
}

func GdipBeginContainer(nativeGraphics *GpGraphics, dstrect,
	srcrect *RectF, unit Unit, state *GraphicsContainer) (status Status, err error) {
	return gdiplusSyscall(procGdipBeginContainer,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(dstrect)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unit),
		uintptr(unsafe.Pointer(state)))
}

func GdipRestoreGraphics(nativeGraphics *GpGraphics, gstate GraphicsState) (status Status, err error) {
	return gdiplusSyscall(procGdipRestoreGraphics,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(gstate))
}

func GdipSaveGraphics(nativeGraphics *GpGraphics, state *GraphicsState) (status Status, err error) {
	return gdiplusSyscall(procGdipSaveGraphics,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(state)))
}

func GdipIsVisibleRect(nativeGraphics *GpGraphics, x, y, width, height REAL, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleRect,
		uintptr(unsafe.Pointer(nativeGraphics)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		uintptr(unsafe.Pointer(booln)))
}

func GdipIsVisiblePoint(nativeGraphics *GpGraphics, x, y REAL, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisiblePoint,
		uintptr(unsafe.Pointer(nativeGraphics)),
		REALbits(x), REALbits(y),
		uintptr(unsafe.Pointer(booln)))
}

func GdipIsVisibleRectI(nativeGraphics *GpGraphics, x, y, width, height INT, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleRectI,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(unsafe.Pointer(booln)))
}

func GdipIsVisiblePointI(nativeGraphics *GpGraphics, x, y INT, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisiblePointI,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(booln)))
}

func GdipIsVisibleClipEmpty(nativeGraphics *GpGraphics, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsVisibleClipEmpty,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(booln)))
}

func GdipGetVisibleClipBoundsI(nativeGraphics *GpGraphics, rect *Rect) (status Status, err error) {
	return gdiplusSyscall(procGdipGetVisibleClipBoundsI,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetVisibleClipBounds(nativeGraphics *GpGraphics, rect *RectF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetVisibleClipBounds,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipIsClipEmpty(nativeGraphics *GpGraphics, booln *BOOL) (status Status, err error) {
	return gdiplusSyscall(procGdipIsClipEmpty,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(booln)))
}

func GdipGetClipBoundsI(nativeGraphics *GpGraphics, rect *Rect) (status Status, err error) {
	return gdiplusSyscall(procGdipGetClipBoundsI,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetClipBounds(nativeGraphics *GpGraphics, rect *RectF) (status Status, err error) {
	return gdiplusSyscall(procGdipGetClipBounds,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(rect)))
}

func GdipGetClip(nativeGraphics *GpGraphics, region *GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipGetClip,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(unsafe.Pointer(region)))
}

func GdipTranslateClipI(nativeGraphics *GpGraphics, dx, dy INT) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateClipI,
		uintptr(unsafe.Pointer(nativeGraphics)),
		uintptr(dx), uintptr(dy))
}

func GdipTranslateClip(nativeGraphics *GpGraphics, dx, dy REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateClip,
		uintptr(unsafe.Pointer(nativeGraphics)),
		REALbits(dx), REALbits(dy))
}

func GdipPrivateAddMemoryFont(fontCollection *GpFontCollection, memory []byte) (status Status, err error) {
	return gdiplusSyscall(procGdipPrivateAddMemoryFont,
		uintptr(unsafe.Pointer(&memory[0])),
		uintptr(INT(len(memory))))
}

func GdipPrivateAddFontFile(fontCollection *GpFontCollection, filename string) (status Status, err error) {
	return gdiplusSyscall(procGdipPrivateAddFontFile,
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))))
}

func GdipDeletePrivateFontCollection(fontCollection **GpFontCollection) (status Status, err error) {
	return gdiplusSyscall(procGdipDeletePrivateFontCollection,
		uintptr(unsafe.Pointer(fontCollection)))
}

func GdipNewPrivateFontCollection(fontCollection **GpFontCollection) (status Status, err error) {
	return gdiplusSyscall(procGdipNewPrivateFontCollection,
		uintptr(unsafe.Pointer(fontCollection)))
}

func GdipNewInstalledFontCollection(fontCollection **GpFontCollection) (status Status, err error) {
	return gdiplusSyscall(procGdipNewInstalledFontCollection,
		uintptr(unsafe.Pointer(fontCollection)))
}

func GdipCloneFontFamily(fontFamily *GpFontFamily, clonedFontFamily **GpFontFamily) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneFontFamily,
		uintptr(unsafe.Pointer(fontFamily)),
		uintptr(unsafe.Pointer(clonedFontFamily)))
}

func GdipGetFontCollectionFamilyList(
	fontCollection *GpFontCollection,
	numSought INT,
	gpfamilies []*GpFontFamily,
	numFound *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontCollectionFamilyList,
		uintptr(unsafe.Pointer(fontCollection)), uintptr(numSought),
		uintptr(unsafe.Pointer(&gpfamilies[0])), uintptr(unsafe.Pointer(numFound)))
}

func GdipGetFontCollectionFamilyCount(
	fontCollection *GpFontCollection,
	numFound *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetFontCollectionFamilyCount,
		uintptr(unsafe.Pointer(fontCollection)), uintptr(unsafe.Pointer(numFound)))
}

func GdipGetGenericFontFamilySansSerif(nativeFamily **GpFontFamily) (status Status, err error) {
	return gdiplusSyscall(procGdipGetGenericFontFamilySansSerif,
		uintptr(unsafe.Pointer(nativeFamily)))
}

func GdipResetClip(graphics *GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipResetClip,
		uintptr(unsafe.Pointer(graphics)))
}

func GdipSetClipHrgn(graphics *GpGraphics, hRgn HRGN, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetClipHrgn,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(hRgn),
		uintptr(combineMode))
}

func GdipSetClipRegion(graphics *GpGraphics, region *GpRegion, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetClipRegion,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(region)),
		uintptr(combineMode))
}

func GdipSetClipPath(graphics *GpGraphics, path *GpPath, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetClipPath,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(path)),
		uintptr(combineMode))
}

func GdipSetClipRectI(graphics *GpGraphics, x, y, width, height INT, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetClipRectI,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(combineMode))
}

func GdipSetClipRect(graphics *GpGraphics, x, y, width, height REAL, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetClipRect,
		uintptr(unsafe.Pointer(graphics)),
		REALbits(x), REALbits(y), REALbits(width), REALbits(height),
		uintptr(combineMode))
}

func GdipSetClipGraphics(graphics *GpGraphics, srcgraphics *GpGraphics, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetClipGraphics,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(srcgraphics)),
		uintptr(combineMode))
}

func GdipEnumerateMetafileSrcRectDestPointsI(graphics *GpGraphics, metafile *GpMetafile,
	destPoints []Point, srcRect *Rect, srcUnit Unit,
	callback EnumerateMetafileProc, callbackData DATA_PTR,
	imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileSrcRectDestPointsI, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((&destPoints[0]))), uintptr(INT(len(destPoints))),
		uintptr(unsafe.Pointer((srcRect))), uintptr(srcUnit),
		callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileSrcRectDestPoints(graphics *GpGraphics, metafile *GpMetafile,
	destPoints []PointF, srcRect *RectF, srcUnit Unit,
	callback EnumerateMetafileProc, callbackData DATA_PTR,
	imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileSrcRectDestPoints, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((&destPoints[0]))), uintptr(INT(len(destPoints))),
		uintptr(unsafe.Pointer((srcRect))), uintptr(srcUnit),
		callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileSrcRectDestRectI(graphics *GpGraphics, metafile *GpMetafile,
	destRect *Rect, srcRect *Rect, srcUnit Unit,
	callback EnumerateMetafileProc, callbackData DATA_PTR,
	imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileSrcRectDestRectI, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destRect))),
		uintptr(unsafe.Pointer((srcRect))), uintptr(srcUnit),
		callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileSrcRectDestRect(graphics *GpGraphics, metafile *GpMetafile,
	destRect *RectF, srcRect *RectF, srcUnit Unit,
	callback EnumerateMetafileProc, callbackData DATA_PTR,
	imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileSrcRectDestRect, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destRect))),
		uintptr(unsafe.Pointer((srcRect))), uintptr(srcUnit),
		callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileSrcRectDestPointI(graphics *GpGraphics, metafile *GpMetafile,
	destPoint *Point, srcRect *Rect, srcUnit Unit,
	callback EnumerateMetafileProc, callbackData DATA_PTR,
	imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileSrcRectDestPointI, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destPoint))),
		uintptr(unsafe.Pointer((srcRect))), uintptr(srcUnit),
		callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileSrcRectDestPoint(graphics *GpGraphics, metafile *GpMetafile,
	destPoint *PointF, srcRect *RectF, srcUnit Unit,
	callback EnumerateMetafileProc, callbackData DATA_PTR,
	imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileSrcRectDestPoint, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destPoint))),
		uintptr(unsafe.Pointer((srcRect))), uintptr(srcUnit),
		callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipGetMetafileHeaderFromFile(filename string, header *MetafileHeader) (status Status, err error) {
	return gdiplusSyscall(procGdipGetMetafileHeaderFromFile, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(header)))
}

func GdipEnumerateMetafileDestRectI(graphics *GpGraphics, metafile *GpMetafile, destRect *Rect, callback EnumerateMetafileProc, callbackData DATA_PTR, imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileDestRectI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destRect))), callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileDestRect(graphics *GpGraphics, metafile *GpMetafile, destRect *RectF, callback EnumerateMetafileProc, callbackData DATA_PTR, imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileDestRect, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destRect))), callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileDestPointI(graphics *GpGraphics, metafile *GpMetafile, destPoint *Point, callback EnumerateMetafileProc, callbackData DATA_PTR, imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileDestPointI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destPoint))), callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileDestPoint(graphics *GpGraphics, metafile *GpMetafile, destPoint *PointF, callback EnumerateMetafileProc, callbackData DATA_PTR, imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileDestPoint, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer((destPoint))), callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileDestPoints(graphics *GpGraphics, metafile *GpMetafile, destPoints []PointF, callback EnumerateMetafileProc, callbackData DATA_PTR, imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileDestPoints, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer(&destPoints[0])), uintptr(INT(len(destPoints))), callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipEnumerateMetafileDestPointsI(graphics *GpGraphics, metafile *GpMetafile, destPoints []Point, callback EnumerateMetafileProc, callbackData DATA_PTR, imageAttributes *GpImageAttributes) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipEnumerateMetafileDestPointsI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(metafile)), uintptr(unsafe.Pointer(&destPoints[0])), uintptr(INT(len(destPoints))), callbackptr, uintptr((callbackData)), uintptr(unsafe.Pointer(imageAttributes)))
}

func GdipCreateMetafileFromWmfFile(filename string, wmfPlaceableFileHeader *WmfPlaceableFileHeader, image **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMetafileFromWmfFile, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(wmfPlaceableFileHeader)), uintptr(unsafe.Pointer(image)))
}

func GdipCreateMetafileFromFile(filename string, image **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMetafileFromFile, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(image)))
}

func GdipGetEffectParameters(nativeEffect *CGpEffect, size *UINT, params ULONG_PTR) (status Status, err error) {
	return gdiplusSyscall(procGdipGetEffectParameters, uintptr(unsafe.Pointer(nativeEffect)), uintptr(unsafe.Pointer(size)), uintptr(params))
}

func GdipSetEffectParameters(nativeEffect *CGpEffect, params ULONG_PTR, size UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetEffectParameters, uintptr(unsafe.Pointer(nativeEffect)), uintptr(params), uintptr(size))
}

func GdipGetEffectParameterSize(nativeEffect *CGpEffect, size *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetEffectParameterSize, uintptr(unsafe.Pointer(nativeEffect)), uintptr(unsafe.Pointer(size)))
}

func GdipDeleteEffect(nativeEffect *CGpEffect) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteEffect, uintptr(unsafe.Pointer(nativeEffect)))
}

//Private Declare Function GdipCreateEffect Lib "gdiplus.dll" (ByVal cid1 As Long, ByVal cid2 As Long, ByVal cid3 As Long, ByVal cid4 As Long, Effect As Long) As Long
func GdipCreateEffect(guid *GUID, nativeEffect **CGpEffect) (status Status, err error) {
	//GdipCreateEffect(D3A1DBE1, 4C178EC4, 97EA4C9F, 3D341CAD, lEffect)
	//{D3A1DBE1-8EC4-4c17-9F4C-EA97AD1C343D}
	//cid1 := guid.Data1
	//cid2 := (DWORD(guid.Data3) << 16) | DWORD(guid.Data2)
	//cid3 := *(*DWORD)(unsafe.Pointer(&guid.Data4[0]))
	//cid4 := *(*DWORD)(unsafe.Pointer(&guid.Data4[4]))
	return gdiplusSyscall(procGdipCreateEffect, uintptr(guid.Data1), uintptr((DWORD(guid.Data3)<<16)|DWORD(guid.Data2)), uintptr(*(*DWORD)(unsafe.Pointer(&guid.Data4[0]))), uintptr(*(*DWORD)(unsafe.Pointer(&guid.Data4[4]))), uintptr(unsafe.Pointer(nativeEffect)))
}

func GdipDrawImageFX(graphics *GpGraphics, nativeImage *GpImage, source *GpRectF,
	xForm *GpMatrix, effect *CGpEffect, imageAttributes *GpImageAttributes, srcUnit GpUnit) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImageFX, uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(nativeImage)), uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(xForm)), uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(imageAttributes)), uintptr(srcUnit))
}

func GdipDrawImagePointsRectI(graphics *GpGraphics, nativeImage *GpImage,
	destPoints []Point,
	srcx, srcy, srcwidth, srcheight INT,
	srcUnit GpUnit, imageAttributes *GpImageAttributes,
	callback DrawImageAbort, callbackData DATA_PTR) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipDrawImagePointsRectI,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(nativeImage)),
		uintptr(unsafe.Pointer(&destPoints[0])), uintptr(INT(len(destPoints))),
		uintptr(srcx), uintptr(srcy), uintptr(srcwidth), uintptr(srcheight),
		uintptr(srcUnit), uintptr(unsafe.Pointer(imageAttributes)),
		callbackptr, uintptr(callbackData))
}

func GdipDrawImagePointsRect(graphics *GpGraphics, nativeImage *GpImage,
	destPoints []PointF,
	srcx, srcy, srcwidth, srcheight REAL,
	srcUnit GpUnit, imageAttributes *GpImageAttributes,
	callback DrawImageAbort, callbackData DATA_PTR) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipDrawImagePointsRect,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(nativeImage)),
		uintptr(unsafe.Pointer(&destPoints[0])), uintptr(INT(len(destPoints))),
		REALbits(srcx), REALbits(srcy), REALbits(srcwidth), REALbits(srcheight),
		uintptr(srcUnit), uintptr(unsafe.Pointer(imageAttributes)),
		callbackptr, uintptr(callbackData))
}

func GdipDrawImageRectRectI(graphics *GpGraphics, nativeImage *GpImage,
	dstx, dsty, dstwidth, dstheight,
	srcx, srcy, srcwidth, srcheight INT,
	srcUnit GpUnit, imageAttributes *GpImageAttributes,
	callback DrawImageAbort, callbackData DATA_PTR) (status Status, err error) {
	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}
	return gdiplusSyscall(procGdipDrawImageRectRectI,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(nativeImage)),
		uintptr(dstx), uintptr(dsty), uintptr(dstwidth), uintptr(dstheight),
		uintptr(srcx), uintptr(srcy), uintptr(srcwidth), uintptr(srcheight),
		uintptr(srcUnit), uintptr(unsafe.Pointer(imageAttributes)),
		callbackptr, uintptr(callbackData))
}

func GdipDrawImageRectRect(graphics *GpGraphics, nativeImage *GpImage,
	dstx, dsty, dstwidth, dstheight,
	srcx, srcy, srcwidth, srcheight REAL,
	srcUnit GpUnit, imageAttributes *GpImageAttributes,
	callback DrawImageAbort, callbackData DATA_PTR) (status Status, err error) {

	var callbackptr uintptr
	if callback != nil {
		callbackptr = syscall.NewCallback(callback)
	}

	return gdiplusSyscall(procGdipDrawImageRectRect,
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(nativeImage)),
		REALbits(dstx), REALbits(dsty), REALbits(dstwidth), REALbits(dstheight),
		REALbits(srcx), REALbits(srcy), REALbits(srcwidth), REALbits(srcheight),
		uintptr(srcUnit), uintptr(unsafe.Pointer(imageAttributes)),
		callbackptr, uintptr(callbackData))
}

func GdipDisposeImageAttributes(nativeImageAttr *GpImageAttributes) (status Status, err error) {
	return gdiplusSyscall(procGdipDisposeImageAttributes, uintptr(unsafe.Pointer(nativeImageAttr)))
}

func GdipCreateImageAttributes(nativeImageAttr **GpImageAttributes) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateImageAttributes, uintptr(unsafe.Pointer(nativeImageAttr)))
}

func GdipDrawImagePointRectI(graphics *GpGraphics, nativeImage *GpImage, x, y, srcx, srcy, srcwidth, srcheight INT, srcUnit GpUnit) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImagePointRectI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), uintptr(x), uintptr(y), uintptr(srcx), uintptr(srcy), uintptr(srcwidth), uintptr(srcheight), uintptr(srcUnit))
}

func GdipDrawImagePointRect(graphics *GpGraphics, nativeImage *GpImage, x, y, srcx, srcy, srcwidth, srcheight REAL, srcUnit GpUnit) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImagePointRect, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), REALbits(x), REALbits(y), REALbits(srcx), REALbits(srcy), REALbits(srcwidth), REALbits(srcheight), uintptr(srcUnit))
}

func GdipDrawImagePointsI(graphics *GpGraphics, nativeImage *GpImage, destPoints []Point) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImagePointsI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), uintptr(unsafe.Pointer(&destPoints[0])), uintptr(INT(len(destPoints))))
}

func GdipDrawImagePoints(graphics *GpGraphics, nativeImage *GpImage, destPoints []PointF) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImagePoints, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), uintptr(unsafe.Pointer(&destPoints[0])), uintptr(INT(len(destPoints))))
}

func GdipDrawImageRectI(graphics *GpGraphics, nativeImage *GpImage, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImageRectI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func GdipDrawImageRect(graphics *GpGraphics, nativeImage *GpImage, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImageRect, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), REALbits(x), REALbits(y), REALbits(width), REALbits(height))
}

func GdipDrawImageI(graphics *GpGraphics, nativeImage *GpImage, x, y INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImageI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), uintptr(x), uintptr(y))
}

func GdipDrawImage(graphics *GpGraphics, nativeImage *GpImage, x, y REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawImage, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(nativeImage)), REALbits(x), REALbits(y))
}

func GdipDrawCachedBitmap(graphics *GpGraphics, cachedBitmap *GpCachedBitmap, x, y INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCachedBitmap, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(cachedBitmap)), uintptr(x), uintptr(y))
}

func GdipCreateBitmapFromFileICM(filename string, bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromFileICM, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(bitmap)))
}

func GdipCreateBitmapFromFile(filename string, bitmap **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateBitmapFromFile, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(bitmap)))
}

func GdipDisposeImage(nativeImage *GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipDisposeImage, uintptr(unsafe.Pointer(nativeImage)))
}

func GdipLoadImageFromFile(filename string, image **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipLoadImageFromFile, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(image)))
}

func GdipLoadImageFromFileICM(filename string, image **GpImage) (status Status, err error) {
	return gdiplusSyscall(procGdipLoadImageFromFileICM, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))), uintptr(unsafe.Pointer(image)))
}

func GdipDeleteCachedBitmap(nativeCachedBitmap *GpCachedBitmap) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteCachedBitmap, uintptr(unsafe.Pointer(nativeCachedBitmap)))
}
func GdipCreateCachedBitmap(bitmap *GpBitmap, graphics *GpGraphics, cachedBitmap **GpCachedBitmap) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateCachedBitmap, uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(cachedBitmap)))
}

func GdipDeleteMatrix(matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteMatrix, uintptr(unsafe.Pointer(matrix)))
}

func GdipCreateMatrix(matrix **GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateMatrix, uintptr(unsafe.Pointer(matrix)))
}

func GdipMeasureDriverString(graphics *GpGraphics, text string, font *GpFont, positions *PointF, flags INT, matrix *GpMatrix, boundingBox *RectF) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	strlen := INT(-1)
	return gdiplusSyscall(procGdipMeasureDriverString, uintptr(unsafe.Pointer(graphics)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(positions)), uintptr(flags), uintptr(unsafe.Pointer(matrix)), uintptr(unsafe.Pointer(boundingBox)))
}

func GdipMeasureCharacterRanges(graphics *GpGraphics, text string, font *GpFont, layoutRect *RectF, stringFormat *GpStringFormat, nativeRegions []*GpRegion) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	strlen := INT(-1)
	return gdiplusSyscall(procGdipMeasureCharacterRanges, uintptr(unsafe.Pointer(graphics)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(layoutRect)), uintptr(unsafe.Pointer(stringFormat)), uintptr(INT(len(nativeRegions))), uintptr(unsafe.Pointer(&nativeRegions[0])))

}

func GdipAddPathString(path *GpPath, text string, family *GpFontFamily, style FontStyle, emSize REAL, layoutRect *RectF, format *GpStringFormat) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	strlen := INT(-1)
	return gdiplusSyscall(procGdipAddPathString, uintptr(unsafe.Pointer(path)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(family)), uintptr(style), REALbits(emSize), uintptr(unsafe.Pointer(layoutRect)), uintptr(unsafe.Pointer(format)))

}

func GdipAddPathStringI(path *GpPath, text string, family *GpFontFamily, style FontStyle, emSize REAL, layoutRect *Rect, format *GpStringFormat) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	strlen := INT(-1)
	return gdiplusSyscall(procGdipAddPathStringI, uintptr(unsafe.Pointer(path)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(family)), uintptr(style), REALbits(emSize), uintptr(unsafe.Pointer(layoutRect)), uintptr(unsafe.Pointer(format)))

}

func GdipMeasureString(graphics *GpGraphics, text string, font *GpFont, layoutRect *RectF, format *GpStringFormat, boundingBox *RectF, codepointsFitted *INT, linesFilled *INT) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	strlen := INT(-1)
	return gdiplusSyscall(procGdipMeasureString, uintptr(unsafe.Pointer(graphics)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(layoutRect)), uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(boundingBox)), uintptr(unsafe.Pointer(codepointsFitted)), uintptr(unsafe.Pointer(linesFilled)))

}

func GdipCreateFontFamilyFromName(name string, fontCollection *GpFontCollection, fontFamily **GpFontFamily) (status Status, err error) {
	nameaddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))
	return gdiplusSyscall(procGdipCreateFontFamilyFromName, nameaddr, uintptr(unsafe.Pointer(fontCollection)), uintptr(unsafe.Pointer(fontFamily)))
}

func GdipDeleteStringFormat(format *GpStringFormat) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteStringFormat, uintptr(unsafe.Pointer(format)))
}

func GdipDeleteFontFamily(nativeFamily *GpFontFamily) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteFontFamily, uintptr(unsafe.Pointer(nativeFamily)))
}

func GdipDeleteFont(font *GpFont) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteFont, uintptr(unsafe.Pointer(font)))
}

func GdipCreateFont(fontFamily *GpFontFamily, emSize REAL, style FontStyle, unit Unit, font **GpFont) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFont, uintptr(unsafe.Pointer(fontFamily)), REALbits(emSize), uintptr(style), uintptr(unit), uintptr(unsafe.Pointer(font)))
}

func GdipDrawString(graphics *GpGraphics, text string, font *GpFont, layoutRect *RectF, format *GpStringFormat, brush *GpBrush) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))) //uintptr(unsafe.Pointer(&wchar[0]))
	strlen := INT(-1)
	//procGdipDrawString.Call(uintptr(unsafe.Pointer(graphics)), StringToUintptr(str), uintptr(len(str)), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(layoutRect)), uintptr(unsafe.Pointer(stringFormat)), uintptr(unsafe.Pointer(brush)))
	return gdiplusSyscall(procGdipDrawString, uintptr(unsafe.Pointer(graphics)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(layoutRect)), uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(brush)))
}

func GdipDrawDriverString(graphics *GpGraphics, text string, font *GpFont, brush *GpBrush, positions *PointF, flags DriverStringOptions, matrix *GpMatrix) (status Status, err error) {
	straddr := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	strlen := INT(-1)
	return gdiplusSyscall(procGdipDrawDriverString, uintptr(unsafe.Pointer(graphics)), straddr, uintptr(strlen), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(positions)), uintptr(flags), uintptr(unsafe.Pointer(matrix)))
}

func GdipCreateStringFormat(formatAttributes INT, language LANGID, format **GpStringFormat) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateStringFormat, uintptr(formatAttributes), uintptr(language), uintptr(unsafe.Pointer(format)))
}

func GdipCreateFontFromDC(hdc HDC, font **GpFont) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFontFromDC, uintptr(hdc), uintptr(unsafe.Pointer(font)))
}

func GdipCombineRegionRect(region *GpRegion, rect *RectF, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipCombineRegionRect, uintptr(unsafe.Pointer(region)), uintptr(unsafe.Pointer(rect)), uintptr(combineMode))
}

func GdipCombineRegionRectI(region *GpRegion, rect *Rect, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipCombineRegionRectI, uintptr(unsafe.Pointer(region)), uintptr(unsafe.Pointer(rect)), uintptr(combineMode))
}

func GdipCombineRegionRegion(region *GpRegion, region2 *GpRegion, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipCombineRegionRegion,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(region2)),
		uintptr(combineMode))
}

func GdipCombineRegionPath(region *GpRegion, path *GpPath, combineMode CombineMode) (status Status, err error) {
	return gdiplusSyscall(procGdipCombineRegionPath,
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(path)),
		uintptr(combineMode))
}

func GdipSetEmpty(region *GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipSetEmpty, uintptr(unsafe.Pointer(region)))
}
func GdipSetInfinite(region *GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipSetInfinite, uintptr(unsafe.Pointer(region)))
}

func GdipCloneRegion(region *GpRegion, cloneRegion **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneRegion, uintptr(unsafe.Pointer(region)), uintptr(unsafe.Pointer(cloneRegion)))
}

func GdipGetRegionHRgn(region *GpRegion, graphics *GpGraphics, hRgn *HRGN) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionHRgn, uintptr(unsafe.Pointer(region)), uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(hRgn)))
}

func GdipCreateRegionHrgn(hRgn HRGN, region **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateRegionHrgn, uintptr(hRgn), uintptr(unsafe.Pointer(region)))
}

func GdipGetRegionData(region *GpRegion, buffer *BYTE, bufferSize UINT, sizeFilled *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRegionData, uintptr(unsafe.Pointer(region)), uintptr(unsafe.Pointer(buffer)), uintptr(bufferSize), uintptr(unsafe.Pointer(sizeFilled)))
}

func GdipCreateRegionRgnData(regionData *BYTE, size INT, region **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateRegionRgnData, uintptr(unsafe.Pointer(regionData)), uintptr(size), uintptr(unsafe.Pointer(region)))
}

func GdipCreateRegionPath(path *GpPath, region **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateRegionPath, uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(region)))
}

func GdipCreateRegionRectI(rect *GpRect, region **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateRegionRectI, uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(region)))
}

func GdipCreateRegionRect(rect *GpRectF, region **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateRegionRect, uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(region)))
}

func GdipDeleteRegion(region *GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteRegion, uintptr(unsafe.Pointer(region)))
}

func GdipCreateRegion(region **GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateRegion, uintptr(unsafe.Pointer(region)))
}

func GdipFillRegion(graphics *GpGraphics, brush *GpBrush, region *GpRegion) (status Status, err error) {
	return gdiplusSyscall(procGdipFillRegion, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(region)))
}

func GdipFillClosedCurve2I(graphics *GpGraphics, brush *GpBrush, points *GpPoint, count INT, tension REAL, fillMode FillMode) (status Status, err error) {
	return gdiplusSyscall(procGdipFillClosedCurve2I, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(points)), uintptr(count), REALbits(tension), uintptr(fillMode))
}

func GdipFillClosedCurveI(graphics *GpGraphics, brush *GpBrush, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipFillClosedCurveI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipFillClosedCurve2(graphics *GpGraphics, brush *GpBrush, points *GpPointF, count INT, tension REAL, fillMode FillMode) (status Status, err error) {
	return gdiplusSyscall(procGdipFillClosedCurve2, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(points)), uintptr(count), REALbits(tension), uintptr(fillMode))
}

func GdipFillClosedCurve(graphics *GpGraphics, brush *GpBrush, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipFillClosedCurve, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipFillPieI(graphics *GpGraphics, brush *GpBrush, x, y, width, height INT, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipFillPieI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), REALbits(startAngle), REALbits(sweepAngle))
}

func GdipFillPie(graphics *GpGraphics, brush *GpBrush, x, y, width, height, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipFillPie, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), REALbits(x), REALbits(y), REALbits(width), REALbits(height), REALbits(startAngle), REALbits(sweepAngle))
}

func GdipFillEllipseI(graphics *GpGraphics, brush *GpBrush, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipFillEllipseI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func GdipFillEllipse(graphics *GpGraphics, brush *GpBrush, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipFillEllipse, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), REALbits(x), REALbits(y), REALbits(width), REALbits(height))
}

func GdipFillPolygonI(graphics *GpGraphics, brush *GpBrush, points *GpPoint, count INT, fillMode FillMode) (status Status, err error) {
	return gdiplusSyscall(procGdipFillPolygonI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(points)), uintptr(count), uintptr(fillMode))
}

func GdipFillPolygon(graphics *GpGraphics, brush *GpBrush, points *GpPointF, count INT, fillMode FillMode) (status Status, err error) {
	return gdiplusSyscall(procGdipFillPolygon, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(points)), uintptr(count), uintptr(fillMode))
}

func GdipFillRectanglesI(graphics *GpGraphics, brush *GpBrush, rects *GpRect, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipFillRectanglesI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(rects)), uintptr(count))
}

func GdipFillRectangleI(graphics *GpGraphics, brush *GpBrush, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipFillRectangleI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func GdipFillRectangles(graphics *GpGraphics, brush *GpBrush, rects *GpRectF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipFillRectangles, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(rects)), uintptr(count))
}

func GdipFillRectangle(graphics *GpGraphics, brush *GpBrush, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipFillRectangle, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), REALbits(x), REALbits(y), REALbits(width), REALbits(height))
}

func GdipDrawClosedCurve2I(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawClosedCurve2I, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count), REALbits(tension))
}

func GdipDrawClosedCurveI(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawClosedCurveI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawClosedCurve2(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawClosedCurve2, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count), REALbits(tension))
}

func GdipDrawClosedCurve(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawClosedCurve, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawCurve3I(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT, offset, numberOfSegments INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCurve3I, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count), uintptr(offset), uintptr(numberOfSegments), REALbits(tension))
}

func GdipDrawCurve2I(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCurve2I, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count), REALbits(tension))
}

func GdipDrawCurveI(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCurveI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawCurve3(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT, offset, numberOfSegments INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCurve3, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count), uintptr(offset), uintptr(numberOfSegments), REALbits(tension))
}

func GdipDrawCurve2(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT, tension REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCurve2, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count), REALbits(tension))
}

func GdipDrawCurve(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawCurve, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawPolygonI(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawPolygonI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawPolygon(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawPolygon, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawPieI(graphics *GpGraphics, pen *GpPen, x, y, width, height INT, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawPieI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), REALbits(startAngle), REALbits(sweepAngle))
}

func GdipDrawPie(graphics *GpGraphics, pen *GpPen, x, y, width, height, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawPie, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), REALbits(x), REALbits(y), REALbits(width), REALbits(height), REALbits(startAngle), REALbits(sweepAngle))
}

func GdipDrawEllipseI(graphics *GpGraphics, pen *GpPen, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawEllipseI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func GdipDrawEllipse(graphics *GpGraphics, pen *GpPen, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawEllipse, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), REALbits(x), REALbits(y), REALbits(width), REALbits(height))
}

func GdipDrawRectangleI(graphics *GpGraphics, pen *GpPen, x, y, width, height INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawRectangleI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func GdipDrawRectanglesI(graphics *GpGraphics, pen *GpPen, rects *GpRect, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawRectanglesI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(rects)), uintptr(count))
}

func GdipDrawRectangles(graphics *GpGraphics, pen *GpPen, rects *GpRectF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawRectangles, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(rects)), uintptr(count))
}

func GdipDrawBeziersI(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawBeziersI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawBezierI(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2, x3, y3, x4, y4 INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawBezierI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(x3), uintptr(y3), uintptr(x4), uintptr(y4))
}

func GdipDrawBeziers(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawBeziers, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawBezier(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2, x3, y3, x4, y4 REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawBezier, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), REALbits(x1), REALbits(y1), REALbits(x2), REALbits(y2), REALbits(x3), REALbits(y3), REALbits(x4), REALbits(y4))
}

func GdipDrawLinesI(graphics *GpGraphics, pen *GpPen, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawLinesI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipDrawLines(graphics *GpGraphics, pen *GpPen, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawLines, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipGetNearestColor(graphics *GpGraphics, argb *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetNearestColor, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(argb)))
}

func GdipTransformPointsI(graphics *GpGraphics, destSpace GpCoordinateSpace,
	srcSpace GpCoordinateSpace, points *GpPoint, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipTransformPointsI, uintptr(unsafe.Pointer(graphics)), uintptr(destSpace), uintptr(srcSpace), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipTransformPoints(graphics *GpGraphics, destSpace GpCoordinateSpace,
	srcSpace GpCoordinateSpace, points *GpPointF, count INT) (status Status, err error) {
	return gdiplusSyscall(procGdipTransformPoints, uintptr(unsafe.Pointer(graphics)), uintptr(destSpace), uintptr(srcSpace), uintptr(unsafe.Pointer(points)), uintptr(count))
}

func GdipGetDpiY(graphics *GpGraphics, dpi *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetDpiY, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(dpi)))
}
func GdipGetDpiX(graphics *GpGraphics, dpi *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetDpiX, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(dpi)))
}
func GdipGetPageScale(graphics *GpGraphics, scale *REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPageScale, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(scale)))
}
func GdipGetPageUnit(graphics *GpGraphics, unit *GpUnit) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPageUnit, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(unit)))
}

func GdipSetPageScale(graphics *GpGraphics, scale REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPageScale, uintptr(unsafe.Pointer(graphics)), REALbits(scale))
}
func GdipGetWorldTransform(graphics *GpGraphics, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipGetWorldTransform, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(matrix)))
}
func GdipRotateWorldTransform(graphics *GpGraphics, angle REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipRotateWorldTransform, uintptr(unsafe.Pointer(graphics)), REALbits(angle), uintptr(order))
}

func GdipScaleWorldTransform(graphics *GpGraphics, sx, sy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipScaleWorldTransform, uintptr(unsafe.Pointer(graphics)), REALbits((sx)), REALbits((sy)), uintptr(order))
}

func GdipTranslateWorldTransform(graphics *GpGraphics, dx, dy REAL, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipTranslateWorldTransform, uintptr(unsafe.Pointer(graphics)), REALbits((dx)), REALbits((dy)), uintptr(order))
}

func GdipFillPath(graphics *GpGraphics, brush *GpBrush, path *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipFillPath, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(path)))
}

//GpPath *path, REAL x, REAL y, REAL width, REAL height
func GdipAddPathRectangle(path *GpPath, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathRectangle, uintptr(unsafe.Pointer(path)), REALbits((x)), REALbits((y)), REALbits((width)), REALbits((height)))
}

func GdipSetSolidFillColor(brush *GpSolidFill, color ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipSetSolidFillColor, uintptr(unsafe.Pointer(brush)), uintptr(color))
}

func GdipGetSolidFillColor(brush *GpSolidFill, color *ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGetSolidFillColor, uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(color)))
}

func GdipCreateSolidFill(color ARGB, brush **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateSolidFill, uintptr(color), uintptr(unsafe.Pointer(brush)))
}

func GdipDeleteBrush(brush *GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteBrush, uintptr(unsafe.Pointer(brush)))
}

func GdipCloneBrush(brush *GpBrush, cloneBrush **GpBrush) (status Status, err error) {
	return gdiplusSyscall(procGdipCloneBrush, uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(cloneBrush)))
}

func GdipGetBrushType(brush *GpBrush, typ *GpBrushType) (status Status, err error) {
	return gdiplusSyscall(procGdipGetBrushType, uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(typ)))
}

func GdipDrawRectangle(graphics *GpGraphics, pen *GpPen, x, y, width, height REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawRectangle, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), REALbits(x), REALbits(y), REALbits(width), REALbits(height))
}

func GdiplusStartup(token *ULONG_PTR, input *GdiplusStartupInput, output *GdiplusStartupOutput) (status Status, err error) {
	return gdiplusSyscall(procGdiplusStartup, uintptr(unsafe.Pointer(token)), uintptr(unsafe.Pointer(input)), uintptr(unsafe.Pointer(output)))
}

func GdiplusShutdown(token ULONG_PTR) (status Status, err error) {
	return gdiplusSyscall(procGdiplusShutdown, uintptr(token))
}

func GdipDrawPath(graphics *GpGraphics, npen *GpPen, npath *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawPath, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(npen)), uintptr(unsafe.Pointer(npath)))
}

func GdipAddPathLine(nativePath *GpPath, x1, y1, x2, y2 REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipAddPathLine, uintptr(unsafe.Pointer(nativePath)), REALbits(x1), REALbits(y1), REALbits(x2), REALbits(y2))
}

func GdipDeletePath(nativePath *GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipDeletePath, uintptr(unsafe.Pointer(nativePath)))
}

func GdipCreatePath(fillMode FillMode, nativePath **GpPath) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePath, uintptr(fillMode), uintptr(unsafe.Pointer(nativePath)))
}

func GdipGraphicsClear(graphics *GpGraphics, color ARGB) (status Status, err error) {
	return gdiplusSyscall(procGdipGraphicsClear, uintptr(unsafe.Pointer(graphics)), uintptr(color))
}

func GdipSetPageUnit(graphics *GpGraphics, unit Unit) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPageUnit, uintptr(unsafe.Pointer(graphics)), uintptr(unit))
}

func GdipDeletePen(pen *GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipDeletePen, uintptr(unsafe.Pointer(pen)))
}

func GdipCreatePen1(color ARGB, width REAL, unit GpUnit, pen **GpPen) (status Status, err error) {
	return gdiplusSyscall(procGdipCreatePen1, uintptr(color), REALbits(width), uintptr(unit), uintptr(unsafe.Pointer(pen)))
}

func GdipDrawLine(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2 REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawLine, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), REALbits((x1)), REALbits((y1)), REALbits((x2)), REALbits((y2)))
}

func GdipDrawLineI(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2 INT) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawLineI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func GdipDrawArc(graphics *GpGraphics, pen *GpPen, x, y, width, height, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawArc, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), REALbits(x), REALbits(y), REALbits(width), REALbits(height), REALbits(startAngle), REALbits(sweepAngle))
}

func GdipDrawArcI(graphics *GpGraphics, pen *GpPen, x, y, width, height INT, startAngle, sweepAngle REAL) (status Status, err error) {
	return gdiplusSyscall(procGdipDrawArcI, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(pen)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), REALbits(startAngle), REALbits(sweepAngle))
}

func GdipMultiplyWorldTransform(graphics *GpGraphics, matrix *GpMatrix, order MatrixOrder) (status Status, err error) {
	return gdiplusSyscall(procGdipMultiplyWorldTransform, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(matrix)), uintptr(order))
}

func GdipResetWorldTransform(graphics *GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipResetWorldTransform, uintptr(unsafe.Pointer(graphics)))
}

func GdipSetWorldTransform(graphics *GpGraphics, matrix *GpMatrix) (status Status, err error) {
	return gdiplusSyscall(procGdipSetWorldTransform, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(matrix)))
}

func GdipGetPixelOffsetMode(graphics *GpGraphics, mode *PixelOffsetMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetPixelOffsetMode, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(mode)))
}

func GdipSetPixelOffsetMode(graphics *GpGraphics, mode PixelOffsetMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetPixelOffsetMode, uintptr(unsafe.Pointer(graphics)), uintptr(mode))
}

func GdipGetSmoothingMode(graphics *GpGraphics, mode *SmoothingMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetSmoothingMode, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(mode)))
}

func GdipSetSmoothingMode(graphics *GpGraphics, mode SmoothingMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetSmoothingMode, uintptr(unsafe.Pointer(graphics)), uintptr(mode))
}

func GdipGetInterpolationMode(graphics *GpGraphics, mode *InterpolationMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetInterpolationMode, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(mode)))
}

func GdipSetInterpolationMode(graphics *GpGraphics, mode InterpolationMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetInterpolationMode, uintptr(unsafe.Pointer(graphics)), uintptr(mode))
}

func GdipGetTextGammaValue(graphics *GpGraphics, gammaValue *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetTextGammaValue, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(gammaValue)))
}

func GdipSetTextGammaValue(graphics *GpGraphics, gammaValue UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetTextGammaValue, uintptr(unsafe.Pointer(graphics)), uintptr(gammaValue))
}

func GdipGetTextContrast(graphics *GpGraphics, contrast *UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetTextContrast, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(contrast)))
}

func GdipSetTextContrast(graphics *GpGraphics, contrast UINT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetTextContrast, uintptr(unsafe.Pointer(graphics)), uintptr(contrast))
}

func GdipGetTextRenderingHint(graphics *GpGraphics, newMode *TextRenderingHint) (status Status, err error) {
	return gdiplusSyscall(procGdipGetTextRenderingHint, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(newMode)))
}

func GdipSetTextRenderingHint(graphics *GpGraphics, newMode TextRenderingHint) (status Status, err error) {
	return gdiplusSyscall(procGdipSetTextRenderingHint, uintptr(unsafe.Pointer(graphics)), uintptr(newMode))
}

func GdipGetCompositingQuality(graphics *GpGraphics, compositingQuality *CompositingQuality) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCompositingQuality, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(compositingQuality)))
}

func GdipSetCompositingQuality(graphics *GpGraphics, compositingQuality CompositingQuality) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCompositingQuality, uintptr(unsafe.Pointer(graphics)), uintptr(compositingQuality))
}

func GdipGetCompositingMode(graphics *GpGraphics, compositingMode *CompositingMode) (status Status, err error) {
	return gdiplusSyscall(procGdipGetCompositingMode, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(compositingMode)))
}

func GdipSetCompositingMode(graphics *GpGraphics, compositingMode CompositingMode) (status Status, err error) {
	return gdiplusSyscall(procGdipSetCompositingMode, uintptr(unsafe.Pointer(graphics)), uintptr(compositingMode))
}

func GdipGetRenderingOrigin(graphics *GpGraphics, x, y *INT) (status Status, err error) {
	return gdiplusSyscall(procGdipGetRenderingOrigin, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y)))
}

func GdipSetRenderingOrigin(graphics *GpGraphics, x, y INT) (status Status, err error) {
	return gdiplusSyscall(procGdipSetRenderingOrigin, uintptr(unsafe.Pointer(graphics)), uintptr(x), uintptr(y))
}

func GdipReleaseDC(graphics *GpGraphics, hdc HDC) (status Status, err error) {
	return gdiplusSyscall(procGdipReleaseDC, uintptr(unsafe.Pointer(graphics)), uintptr(hdc))
}

func GdipFree(in_pvoid uintptr) (err error) {
	_, _, e1 := syscall.Syscall(procGdipFree.Addr(), 1, in_pvoid, 0, 0)
	if e1 != 0 {
		err = error(e1)
	}
	return
}

func GdipAlloc(in_size INT) (ptr uintptr, err error) {
	ptr, _, e := syscall.Syscall(procGdipAlloc.Addr(), 1, uintptr(in_size), 0, 0)
	if int(ptr) == 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GdipCreateFromHDC(hdc HDC, graphics **GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFromHDC, uintptr(hdc), uintptr(unsafe.Pointer(graphics)))
}

func GdipCreateFromHDC2(hdc HDC, hdevice HANDLE, graphics **GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFromHDC2, uintptr(hdc), uintptr(hdevice), uintptr(unsafe.Pointer(graphics)))
}

func GdipCreateFromHWNDICM(hwnd HANDLE, graphics **GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFromHWNDICM, uintptr(hwnd), uintptr(unsafe.Pointer(graphics)))
}

func GdipCreateFromHWND(hwnd HANDLE, graphics **GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipCreateFromHWND, uintptr(hwnd), uintptr(unsafe.Pointer(graphics)))
}

func GdipGetImageGraphicsContext(image *GpImage, graphics **GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipGetImageGraphicsContext, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(graphics)))
}

func GdipDeleteGraphics(graphics *GpGraphics) (status Status, err error) {
	return gdiplusSyscall(procGdipDeleteGraphics, uintptr(unsafe.Pointer(graphics)))
}

func GdipFlush(graphics *GpGraphics, intention GpFlushIntention) (status Status, err error) {
	return gdiplusSyscall(procGdipFlush, uintptr(unsafe.Pointer(graphics)), uintptr(intention))
}

func GdipGetDC(graphics *GpGraphics, hdc *HDC) (status Status, err error) {
	return gdiplusSyscall(procGdipGetDC, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(hdc)))
}

func gdiplusSyscall(proc *syscall.LazyProc, a ...uintptr) (status Status, err error) {
	var rs uintptr
	var e syscall.Errno
	switch len(a) {
	case 0:
		rs, _, e = syscall.Syscall(proc.Addr(), uintptr(len(a)), 0, 0, 0)
	case 1:
		rs, _, e = syscall.Syscall(proc.Addr(), uintptr(len(a)), a[0], 0, 0)
	case 2:
		rs, _, e = syscall.Syscall(proc.Addr(), uintptr(len(a)), a[0], a[1], 0)
	case 3:
		rs, _, e = syscall.Syscall(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2])
	case 4:
		rs, _, e = syscall.Syscall6(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], 0, 0)
	case 5:
		rs, _, e = syscall.Syscall6(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], 0)
	case 6:
		rs, _, e = syscall.Syscall6(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5])
	case 7:
		rs, _, e = syscall.Syscall9(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0)
	case 8:
		rs, _, e = syscall.Syscall9(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0)
	case 9:
		rs, _, e = syscall.Syscall9(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
	case 10:
		rs, _, e = syscall.Syscall12(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0)
	case 11:
		rs, _, e = syscall.Syscall12(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0)
	case 12:
		rs, _, e = syscall.Syscall12(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
	case 13:
		rs, _, e = syscall.Syscall15(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0)
	case 14:
		rs, _, e = syscall.Syscall15(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0)
	case 15:
		rs, _, e = syscall.Syscall15(proc.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14])
	default:
		return Status(InvalidParameter), errors.New("Call " + proc.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
	}

	status = Status(rs)
	if status != Ok {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(StatusText[Status(status)])
		}
	}
	return
}
