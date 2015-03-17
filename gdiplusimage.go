package gdiplus

import (
	. "github.com/trygo/winapi"
	"unsafe"
)

type IImage interface {
	GetNativeImage() *GpImage
}

type Image struct {
	GdiplusBase
	nativeImage *GpImage
}

func NewImage(filename string, useEmbeddedColorManagement ...BOOL) (*Image, error) {
	image := &Image{}
	if len(useEmbeddedColorManagement) > 0 && useEmbeddedColorManagement[0] {
		image.setStatus(GdipLoadImageFromFileICM(filename, &image.nativeImage))
	} else {
		image.setStatus(GdipLoadImageFromFile(filename, &image.nativeImage))
	}
	return image, image.LastError
}

func (this *Image) Release() {
	if this.nativeImage != nil {
		this.setStatus(GdipDisposeImage(this.nativeImage))
		this.nativeImage = nil
	}
}

func (this *Image) GetNativeImage() *GpImage {
	return this.nativeImage
}

func (this *Image) Clone() *Image {
	image := &Image{}
	this.setStatus(GdipCloneImage(this.nativeImage, &image.nativeImage))
	return image
}

func (this *Image) GetEncoderParameterListSize(clsidEncoder *CLSID) (size UINT) {
	this.setStatus(GdipGetEncoderParameterListSize(this.nativeImage,
		clsidEncoder,
		&size))
	return
}

func (this *Image) GetEncoderParameterList(
	clsidEncoder *CLSID, size UINT) (buffer []EncoderParameters, status Status) {

	buffer = make([]EncoderParameters, size)
	status = this.setStatus(GdipGetEncoderParameterList(this.nativeImage,
		clsidEncoder,
		buffer))
	return
}

func (this *Image) Save(
	filename string,
	clsidEncoder *CLSID,
	encoderParams *EncoderParameters) Status {
	return this.setStatus(GdipSaveImageToFile(this.nativeImage,
		filename,
		clsidEncoder,
		encoderParams))
}

func (this *Image) SaveAdd(encoderParams *EncoderParameters) Status {
	return this.setStatus(GdipSaveAdd(this.nativeImage, encoderParams))
}

func (this *Image) SaveAdd2(newImage IImage, encoderParams *EncoderParameters) Status {
	if newImage == nil {
		return this.setStatus(InvalidParameter, nil)
	}

	return this.setStatus(GdipSaveAddImage(this.nativeImage,
		newImage.GetNativeImage(), encoderParams))
}

func (this *Image) GetType() ImageType {
	var typ ImageType = ImageTypeUnknown
	this.setStatus(GdipGetImageType(this.nativeImage, &typ))
	return typ
}

func (this *Image) GetPhysicalDimension() (size *SizeF, status Status) {
	size = &SizeF{}
	status = this.setStatus(GdipGetImageDimension(this.nativeImage,
		&size.Width, &size.Height))
	return
}

func (this *Image) GetBounds() (srcRect *RectF, srcUnit Unit, status Status) {
	srcRect = &RectF{}
	status = this.setStatus(GdipGetImageBounds(this.nativeImage, srcRect, &srcUnit))
	return
}

func (this *Image) GetWidth() (width UINT) {
	this.setStatus(GdipGetImageWidth(this.nativeImage, &width))
	return
}

func (this *Image) GetHeight() (height UINT) {
	this.setStatus(GdipGetImageHeight(this.nativeImage, &height))
	return
}

func (this *Image) GetHorizontalResolution() (resolution REAL) {
	this.setStatus(GdipGetImageHorizontalResolution(this.nativeImage, &resolution))
	return
}

func (this *Image) GetVerticalResolution() (resolution REAL) {
	this.setStatus(GdipGetImageVerticalResolution(this.nativeImage, &resolution))
	return
}

func (this *Image) GetFlags() (flags UINT) {
	this.setStatus(GdipGetImageFlags(this.nativeImage, &flags))
	return
}

func (this *Image) GetRawFormat() (format GUID, status Status) {
	status = this.setStatus(GdipGetImageRawFormat(this.nativeImage, &format))
	return
}

func (this *Image) GetPixelFormat() (format PixelFormat) {
	this.setStatus(GdipGetImagePixelFormat(this.nativeImage, &format))
	return
}

func (this *Image) GetPaletteSize() (size INT) {
	this.setStatus(GdipGetImagePaletteSize(this.nativeImage, &size))
	return
}

func (this *Image) GetPalette(size INT) (palette *ColorPalette, status Status) {
	palette_ := make([]byte, size)
	palette = (*ColorPalette)(unsafe.Pointer(&palette_[0]))
	status = this.setStatus(GdipGetImagePalette(this.nativeImage, palette, size))
	return
}

func (this *Image) SetPalette(palette *ColorPalette) Status {
	return this.setStatus(GdipSetImagePalette(this.nativeImage, palette))
}

func (this *Image) GetThumbnailImage(thumbWidth UINT, thumbHeight UINT,
	callback GetThumbnailImageAbort, callbackData DATA_PTR) *Image {
	newImage := &Image{}
	this.setStatus(GdipGetImageThumbnail(this.nativeImage,
		thumbWidth, thumbHeight,
		&newImage.nativeImage,
		callback, callbackData))

	return newImage
}

func (this *Image) GetFrameDimensionsCount() (count UINT) {
	this.setStatus(GdipImageGetFrameDimensionsCount(this.nativeImage,
		&count))
	return
}

func (this *Image) GetFrameDimensionsList(count UINT) (dimensionIDs []GUID, status Status) {
	dimensionIDs = make([]GUID, count)
	status = this.setStatus(GdipImageGetFrameDimensionsList(this.nativeImage,
		dimensionIDs))
	return
}

func (this *Image) GetFrameCount(dimensionID *GUID) (count UINT) {
	this.setStatus(GdipImageGetFrameCount(this.nativeImage,
		dimensionID, &count))
	return
}

func (this *Image) SelectActiveFrame(
	dimensionID *GUID, frameIndex UINT) Status {
	return this.setStatus(GdipImageSelectActiveFrame(this.nativeImage,
		dimensionID, frameIndex))
}

func (this *Image) RotateFlip(rotateFlipType RotateFlipType) Status {
	return this.setStatus(GdipImageRotateFlip(this.nativeImage,
		rotateFlipType))
}

func (this *Image) GetPropertyCount() (numProperty UINT) {
	this.setStatus(GdipGetPropertyCount(this.nativeImage,
		&numProperty))
	return
}

func (this *Image) GetPropertyIdList(numOfProperty UINT) (list []PROPID, status Status) {
	list = make([]PROPID, numOfProperty)
	status = this.setStatus(GdipGetPropertyIdList(this.nativeImage, list))
	return
}

func (this *Image) GetPropertyItemSize(
	propId PROPID) (size UINT) {
	this.setStatus(GdipGetPropertyItemSize(this.nativeImage,
		propId, &size))
	return
}

func (this *Image) GetPropertyItem(propId PROPID,
	propSize UINT) (buffer *PropertyItem, status Status) {

	if propSize == 0 {
		status = this.setStatus(InvalidParameter, nil)
		return
	}

	buffer_ := make([]byte, propSize)
	buffer = (*PropertyItem)(unsafe.Pointer(&buffer_[0]))

	status = this.setStatus(GdipGetPropertyItem(this.nativeImage,
		propId, propSize, buffer))
	return
}

func (this *Image) GetPropertySize() (totalBufferSize, numProperties UINT, status Status) {
	status = this.setStatus(GdipGetPropertySize(this.nativeImage,
		&totalBufferSize,
		&numProperties))
	return
}

func (this *Image) GetAllPropertyItems(
	totalBufferSize, numProperties UINT) (allItems []PropertyItem, status Status) {

	allItems_ := make([]byte, totalBufferSize)
	allItemsPtr := (*PropertyItem)(unsafe.Pointer(&allItems_[0]))

	status = this.setStatus(GdipGetAllPropertyItems(this.nativeImage,
		totalBufferSize,
		numProperties,
		allItemsPtr))

	if status == Ok {
		allItems = make([]PropertyItem, numProperties)
		for j := UINT(0); j < numProperties; j++ {
			item := (*PropertyItem)(unsafe.Pointer(uintptr(unsafe.Pointer(allItemsPtr)) + uintptr(PropertyItem_SIZE*j)))
			allItems[j] = *item
		}
	}
	return
}

func (this *Image) RemovePropertyItem(propId PROPID) Status {
	return this.setStatus(GdipRemovePropertyItem(this.nativeImage, propId))
}

func (this *Image) SetPropertyItem(item *PropertyItem) Status {
	return this.setStatus(GdipSetPropertyItem(this.nativeImage, item))
}

//#if (GDIPVER >= 0x0110)
func (this *Image) FindFirstItem(item *ImageItemData) Status {
	return this.setStatus(GdipFindFirstImageItem(this.nativeImage, item))
}

func (this *Image) FindNextItem(item *ImageItemData) Status {
	return this.setStatus(GdipFindNextImageItem(this.nativeImage, item))
}

func (this *Image) GetItemData(item *ImageItemData) Status {
	return this.setStatus(GdipGetImageItemData(this.nativeImage, item))
}

//#endif //(GDIPVER >= 0x0110)
