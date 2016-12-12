package gdiplus

import (
	. "github.com/tryor/winapi"
	. "github.com/tryor/winapi/gdi"
)

type Bitmap struct {
	Image
}

func NewBitmap(filename string, useEmbeddedColorManagement ...BOOL) (*Bitmap, error) {
	bitmap := &Bitmap{}
	if len(useEmbeddedColorManagement) > 0 && useEmbeddedColorManagement[0] {
		bitmap.setStatus(GdipCreateBitmapFromFileICM(filename, &bitmap.Image.nativeImage))
	} else {
		bitmap.setStatus(GdipCreateBitmapFromFile(filename, &bitmap.Image.nativeImage))
	}
	return bitmap, bitmap.LastError
}

func NewBitmap2(width, height, stride INT,
	format PixelFormat, scan0 *BYTE) (*Bitmap, error) {

	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromScan0(width,
		height,
		stride,
		format,
		scan0,
		&bitmap.Image.nativeImage))

	return bitmap, bitmap.LastError
}

func NewBitmap3(width, height INT, format PixelFormat) (*Bitmap, error) {
	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromScan0(width,
		height,
		0,
		format,
		nil,
		&bitmap.Image.nativeImage))

	return bitmap, bitmap.LastError
}

func NewBitmap4(width, height INT, target *Graphics) (*Bitmap, error) {
	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromGraphics(width,
		height,
		target.nativeGraphics,
		&bitmap.Image.nativeImage))

	return bitmap, bitmap.LastError
}

func NewBitmap5(surface *IDirectDrawSurface7) (*Bitmap, error) {
	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromDirectDrawSurface(surface,
		&bitmap.Image.nativeImage))
	return bitmap, bitmap.LastError
}

func NewBitmap6(gdiBitmapInfo *BITMAPINFO, gdiBitmapData DATA_PTR) (*Bitmap, error) {
	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromGdiDib(gdiBitmapInfo,
		gdiBitmapData,
		&bitmap.Image.nativeImage))

	return bitmap, bitmap.LastError
}

func NewBitmap7(hbm HBITMAP, hpal HPALETTE) (*Bitmap, error) {
	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromHBITMAP(hbm, hpal, &bitmap.Image.nativeImage))
	return bitmap, bitmap.LastError
}

func NewBitmap8(hicon HICON) (*Bitmap, error) {
	bitmap := &Bitmap{}
	bitmap.setStatus(GdipCreateBitmapFromHICON(hicon, &bitmap.Image.nativeImage))
	return bitmap, bitmap.LastError
}

func NewBitmap9(hInstance HINSTANCE, bitmapName string) (*Bitmap, error) {
	bitmap := &Bitmap{}

	bitmap.setStatus(GdipCreateBitmapFromResource(hInstance,
		bitmapName,
		&bitmap.Image.nativeImage))

	return bitmap, bitmap.LastError
}

func BitmapFromFile(filename string, useEmbeddedColorManagement ...BOOL) (*Bitmap, error) {
	return NewBitmap(filename, useEmbeddedColorManagement...)
}

func BitmapFromDirectDrawSurface7(surface *IDirectDrawSurface7) (*Bitmap, error) {
	return NewBitmap5(surface)
}

func BitmapFromBITMAPINFO(gdiBitmapInfo *BITMAPINFO, gdiBitmapData DATA_PTR) (*Bitmap, error) {
	return NewBitmap6(gdiBitmapInfo, gdiBitmapData)
}

func BitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE) (*Bitmap, error) {
	return NewBitmap7(hbm, hpal)
}

func BitmapFromHICON(hicon HICON) (*Bitmap, error) {
	return NewBitmap8(hicon)
}

func BitmapFromResource(hInstance HINSTANCE, bitmapName string) (*Bitmap, error) {
	return NewBitmap9(hInstance, bitmapName)
}

func (this *Bitmap) GetHBITMAP(colorBackground Color) (hbmReturn HBITMAP, status Status) {
	status = this.setStatus(GdipCreateHBITMAPFromBitmap(
		this.nativeImage,
		&hbmReturn,
		colorBackground.GetValue()))
	return
}

func (this *Bitmap) GetHICON() (hiconReturn HICON, status Status) {
	status = this.setStatus(GdipCreateHICONFromBitmap(
		this.nativeImage, &hiconReturn))
	return
}

func (this *Bitmap) CloneI(x, y, width, height INT, format PixelFormat) *Bitmap {
	var gpdstBitmap *GpImage
	if this.setStatus(GdipCloneBitmapAreaI(
		x, y, width, height, format,
		this.nativeImage, &gpdstBitmap)) == Ok {
		return &Bitmap{Image: Image{nativeImage: gpdstBitmap}}
	} else {
		return nil
	}
}

func (this *Bitmap) CloneI2(rect *Rect, format PixelFormat) *Bitmap {
	return this.CloneI(rect.X, rect.Y, rect.Width, rect.Height, format)
}

func (this *Bitmap) Clone(x, y, width, height REAL, format PixelFormat) *Bitmap {
	var gpdstBitmap *GpImage
	if this.setStatus(GdipCloneBitmapArea(
		x, y, width, height, format,
		this.nativeImage, &gpdstBitmap)) == Ok {
		return &Bitmap{Image: Image{nativeImage: gpdstBitmap}}
	} else {
		return nil
	}
}

func (this *Bitmap) Clone2(rect *RectF, format PixelFormat) *Bitmap {
	return this.Clone(rect.X, rect.Y, rect.Width, rect.Height, format)
}

func (this *Bitmap) LockBits(rect *Rect, flags UINT, format PixelFormat) (lockedBitmapData *BitmapData, status Status) {
	lockedBitmapData = &BitmapData{}
	status = this.setStatus(GdipBitmapLockBits(
		this.nativeImage,
		rect,
		flags,
		format,
		lockedBitmapData))
	return
}

func (this *Bitmap) UnlockBits(lockedBitmapData *BitmapData) Status {
	return this.setStatus(GdipBitmapUnlockBits(
		this.nativeImage,
		lockedBitmapData))
}

func (this *Bitmap) GetPixel(x, y INT) (color Color, status Status) {
	status = this.setStatus(GdipBitmapGetPixel(
		this.nativeImage, x, y, &color.Argb))
	return
}

func (this *Bitmap) SetPixel(x, y INT, color Color) Status {
	return this.setStatus(GdipBitmapSetPixel(
		this.nativeImage,
		x, y, color.GetValue()))
}

//#if (GDIPVER >= 0x0110)
func (this *Bitmap) SetAbort(pIAbort *GdiplusAbort) Status {
	return this.setStatus(GdipImageSetAbort(
		this.nativeImage, pIAbort))
}

func (this *Bitmap) ConvertFormat(
	format PixelFormat,
	dithertype DitherType,
	palettetype PaletteType,
	palette *ColorPalette,
	alphaThresholdPercent REAL) Status {
	return this.setStatus(GdipBitmapConvertFormat(
		this.nativeImage,
		format,
		dithertype,
		palettetype,
		palette,
		alphaThresholdPercent))
}

func (this *Bitmap) InitializePalette(
	palette *ColorPalette, // output palette. must be allocated.
	palettetype PaletteType, // palette enumeration type.
	optimalColors INT, // how many optimal colors
	useTransparentColor BOOL, // add a transparent color to the palette.
	bitmap *Bitmap) Status { // optional bitmap for median cut.

	var bitmapNativeImage *GpImage
	if bitmap != nil {
		bitmapNativeImage = bitmap.nativeImage
	}

	return this.setStatus(GdipInitializePalette(
		palette,
		palettetype,
		optimalColors,
		useTransparentColor,
		bitmapNativeImage))
}

//TODO 未测试
func ApplyEffect(inputs []*Bitmap, // IN
	numInputs INT, //IN
	effect IEffect, //IN
	ROI *RECT, //IN optional parameter.
	outputRect *RECT, //OUT optional parameter.
	output **Bitmap) Status { //OUT

	if numInputs < 0 {
		return InvalidParameter
	}

	var outputNative *GpImage
	//GpBitmap **nativeInputs = new GpBitmap * [numInputs];
	nativeInputs := make([]*GpImage, numInputs)

	for i := 0; i < int(numInputs); i++ {
		//nativeInputs[i] = static_cast<GpBitmap*>(inputs[i]->nativeImage);
		nativeInputs[i] = inputs[i].nativeImage
	}

	if effect.GetAuxData() != 0 {
		GdipFree(uintptr(effect.GetAuxData()))
		effect.SetAuxData(0)
		effect.SetAuxDataSize(0)
	}

	status, _ := GdipBitmapCreateApplyEffect(
		&nativeInputs[0],
		numInputs,
		effect.GetNativeEffect(),
		ROI,
		outputRect,
		&outputNative,
		effect.IsUseAuxData(),
		effect.GetAuxDataPtr(),
		effect.GetAuxDataSizePtr())

	if (Ok == status) && outputNative != nil {
		*output = &Bitmap{Image: Image{nativeImage: outputNative}}
		//*output = new Bitmap(outputNative);
		//if (NULL == *output)
		//{
		//    status = OutOfMemory;
		//    DllExports::GdipDisposeImage(outputNative);
		//}
	} else {
		*output = nil
	}
	return Status(status)
}

//TODO 测试时失败
func (this *Bitmap) ApplyEffect(effect IEffect, ROI *RECT) Status {

	if effect.GetAuxData() != 0 {
		GdipFree(uintptr(effect.GetAuxData()))
		effect.SetAuxData(0)
		effect.SetAuxDataSize(0)
	}

	return this.setStatus(GdipBitmapApplyEffect(
		this.nativeImage,
		effect.GetNativeEffect(),
		ROI,
		effect.IsUseAuxData(),
		effect.GetAuxDataPtr(),
		effect.GetAuxDataSizePtr()))
}

func (this *Bitmap) SetResolution(xdpi, ydpi REAL) Status {
	return this.setStatus(GdipBitmapSetResolution(
		this.nativeImage, xdpi, ydpi))
}

//TODO 测试时失败
func (this *Bitmap) GetHistogram(
	format HistogramFormat, numberOfEntries UINT) (channel0 []UINT, channel1 []UINT,
	channel2 []UINT, channel3 []UINT, status Status) {
	//__out_bcount(sizeof(UINT) * 256)
	channel0 = make([]UINT, numberOfEntries)
	channel1 = make([]UINT, numberOfEntries)
	channel2 = make([]UINT, numberOfEntries)
	channel3 = make([]UINT, numberOfEntries)
	status = this.setStatus(GdipBitmapGetHistogram(
		this.nativeImage,
		format,
		numberOfEntries,
		channel0,
		channel1,
		channel2,
		channel3))

	return
}

func GetHistogramSize(format HistogramFormat) (numberOfEntries UINT, status Status) {
	gpstatus, _ := GdipBitmapGetHistogramSize(format, &numberOfEntries)
	status = Status(gpstatus)
	return
}
