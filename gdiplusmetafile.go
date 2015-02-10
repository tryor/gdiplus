package gdiplus

import (
	. "github.com/trygo/winapi"
)

type Metafile struct {
	Image
}

func NewMetafile(filename string) (*Metafile, error) {
	metafile := &Metafile{}
	metafile.setStatus(GdipCreateMetafileFromFile(filename, &metafile.Image.nativeImage))
	return metafile, metafile.LastError
}

func NewMetafile2(filename string, wmfPlaceableFileHeader *WmfPlaceableFileHeader) (*Metafile, error) {
	metafile := &Metafile{}
	metafile.setStatus(GdipCreateMetafileFromWmfFile(filename, wmfPlaceableFileHeader, &metafile.Image.nativeImage))
	return metafile, metafile.LastError
}

func NewMetafile3(
	hWmf HMETAFILE,
	wmfPlaceableFileHeader *WmfPlaceableFileHeader,
	deleteWmf BOOL) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipCreateMetafileFromWmf(hWmf, deleteWmf,
		wmfPlaceableFileHeader, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile4(
	hEmf HENHMETAFILE,
	deleteEmf BOOL) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipCreateMetafileFromEmf(hEmf, deleteEmf,
		&metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile5(stream *IStream) (*Metafile, error) {
	metafile := &Metafile{}
	metafile.setStatus(GdipCreateMetafileFromStream(stream,
		&metafile.nativeImage))
	return metafile, metafile.LastError
}

func NewMetafile6(
	referenceHdc HDC,
	frameRect *RectF,
	frameUnit MetafileFrameUnit,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafile(
		referenceHdc, typ, frameRect, frameUnit,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile7(referenceHdc HDC, typ EmfType, description string) (*Metafile, error) {
	metafile := &Metafile{}
	metafile.setStatus(GdipRecordMetafile(referenceHdc, typ, nil, MetafileFrameUnitGdi,
		description, &metafile.nativeImage))
	return metafile, metafile.LastError
}

func NewMetafile6I(
	referenceHdc HDC,
	frameRect *Rect,
	frameUnit MetafileFrameUnit,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileI(
		referenceHdc, typ, frameRect, frameUnit,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile7I(referenceHdc HDC, typ EmfType, description string) (*Metafile, error) {
	metafile := &Metafile{}
	metafile.setStatus(GdipRecordMetafileI(referenceHdc, typ, nil, MetafileFrameUnitGdi,
		description, &metafile.nativeImage))
	return metafile, metafile.LastError
}

func NewMetafile8(
	fileName string,
	referenceHdc HDC,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileFileName(fileName,
		referenceHdc, typ, nil, MetafileFrameUnitGdi,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile9(
	fileName string,
	referenceHdc HDC,
	frameRect *RectF,
	frameUnit MetafileFrameUnit,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileFileName(fileName,
		referenceHdc, typ, frameRect, frameUnit,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile9I(
	fileName string,
	referenceHdc HDC,
	frameRect *Rect,
	frameUnit MetafileFrameUnit,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileFileNameI(fileName,
		referenceHdc, typ, frameRect, frameUnit,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile10(
	stream *IStream,
	referenceHdc HDC,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileStream(stream,
		referenceHdc, typ, nil, MetafileFrameUnitGdi,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile11(
	stream *IStream,
	referenceHdc HDC,
	frameRect *RectF,
	frameUnit MetafileFrameUnit,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileStream(stream,
		referenceHdc, typ, frameRect, frameUnit,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func NewMetafile11I(
	stream *IStream,
	referenceHdc HDC,
	frameRect *Rect,
	frameUnit MetafileFrameUnit,
	typ EmfType,
	description string) (*Metafile, error) {
	metafile := &Metafile{}

	metafile.setStatus(GdipRecordMetafileStreamI(stream,
		referenceHdc, typ, frameRect, frameUnit,
		description, &metafile.nativeImage))

	return metafile, metafile.LastError
}

func GetMetafileHeader(filename string) (header *MetafileHeader, status Status) {
	header = &MetafileHeader{}
	status, _ = GdipGetMetafileHeaderFromFile(filename, header)
	return
}

func GetMetafileHeader2(
	hWmf HMETAFILE,
	wmfPlaceableFileHeader *WmfPlaceableFileHeader) (header *MetafileHeader, status Status) {
	header = &MetafileHeader{}
	status, _ = GdipGetMetafileHeaderFromWmf(hWmf,
		wmfPlaceableFileHeader, header)
	return
}

func GetMetafileHeader3(
	hEmf HENHMETAFILE) (header *MetafileHeader, status Status) {
	header = &MetafileHeader{}
	status, _ = GdipGetMetafileHeaderFromEmf(hEmf, header)
	return
}

func GetMetafileHeader4(
	stream *IStream) (header *MetafileHeader, status Status) {
	header = &MetafileHeader{}
	status, _ = GdipGetMetafileHeaderFromStream(stream, header)
	return
}

func (this *Metafile) GetMetafileHeader() (header *MetafileHeader, status Status) {
	header = &MetafileHeader{}
	status = this.setStatus(GdipGetMetafileHeaderFromMetafile(
		this.nativeImage, header))
	return
}

// Once this method is called, the Metafile object is in an invalid state
// and can no longer be used.  It is the responsiblity of the caller to
// invoke DeleteEnhMetaFile to delete this hEmf.
func (this *Metafile) GetHENHMETAFILE() (hEmf HENHMETAFILE) {
	this.setStatus(GdipGetHemfFromMetafile(
		this.nativeImage,
		&hEmf))
	return
}

// Used in conjuction with Graphics::EnumerateMetafile to play an EMF+
// The data must be DWORD aligned if it's an EMF or EMF+.  It must be
// WORD aligned if it's a WMF.
func (this *Metafile) PlayRecord(
	recordType EmfPlusRecordType,
	flags UINT,
	data []BYTE) Status {
	return this.setStatus(GdipPlayMetafileRecord(
		this.nativeImage,
		recordType,
		flags,
		data))
}

// If you're using a printer HDC for the metafile, but you want the
// metafile rasterized at screen resolution, then use this API to set
// the rasterization dpi of the metafile to the screen resolution,
// e.g. 96 dpi or 120 dpi.
func (this *Metafile) SetDownLevelRasterizationLimit(
	metafileRasterizationLimitDpi UINT) Status {
	return this.setStatus(GdipSetMetafileDownLevelRasterizationLimit(
		this.nativeImage,
		metafileRasterizationLimitDpi))
}

func (this *Metafile) GetDownLevelRasterizationLimit() UINT {
	var metafileRasterizationLimitDpi UINT

	this.setStatus(GdipGetMetafileDownLevelRasterizationLimit(
		this.nativeImage,
		&metafileRasterizationLimitDpi))

	return metafileRasterizationLimitDpi
}

//typedef BYTE far            *LPBYTE;
//OUT             pData16
func EmfToWmfBits(
	hemf HENHMETAFILE,
	cbData16 UINT,
	pData16 LPBYTE,
	iMapMode INT,
	eFlags INT) (UINT, error) {
	return GdipEmfToWmfBits(
		hemf,
		cbData16,
		pData16,
		iMapMode,
		eFlags)
}

//#if (GDIPVER >= 0x0110)
func (this *Metafile) ConvertToEmfPlus(
	refGraphics *Graphics,
	inConversionFailureFlag INT,
	emfType EmfType,
	description string) (outConversionFailureFlag INT, status Status) {
	var metafile *GpImage
	outConversionFailureFlag = inConversionFailureFlag
	status, _ = GdipConvertToEmfPlus(
		refGraphics.nativeGraphics,
		this.nativeImage,
		&outConversionFailureFlag,
		emfType, description, &metafile)

	if metafile != nil {
		if status == Ok {
			GdipDisposeImage(this.nativeImage)
			//SetNativeImage(metafile);
			this.nativeImage = metafile
		} else {
			GdipDisposeImage(metafile)
		}
	}
	return
}

func (this *Metafile) ConvertToEmfPlusToFile(
	refGraphics *Graphics,
	filename string,
	inConversionFailureFlag INT,
	emfType EmfType,
	description string) (outConversionFailureFlag INT, status Status) {

	var metafile *GpImage
	outConversionFailureFlag = inConversionFailureFlag

	status, _ = GdipConvertToEmfPlusToFile(
		refGraphics.nativeGraphics,
		this.nativeImage,
		&outConversionFailureFlag,
		filename, emfType, description, &metafile)

	if metafile != nil {
		if status == Ok {
			GdipDisposeImage(this.nativeImage)
			this.nativeImage = metafile
		} else {
			GdipDisposeImage(metafile)
		}
	}
	return
}

func (this *Metafile) ConvertToEmfPlusToStream(
	refGraphics *Graphics,
	stream *IStream,
	inConversionFailureFlag INT,
	emfType EmfType,
	description string) (outConversionFailureFlag INT, status Status) {

	var metafile *GpImage
	outConversionFailureFlag = inConversionFailureFlag

	status, _ = GdipConvertToEmfPlusToStream(
		refGraphics.nativeGraphics,
		this.nativeImage,
		&outConversionFailureFlag,
		stream, emfType, description, &metafile)

	if metafile != nil {
		if status == Ok {
			GdipDisposeImage(this.nativeImage)
			this.nativeImage = metafile
		} else {
			GdipDisposeImage(metafile)
		}
	}
	return
}
