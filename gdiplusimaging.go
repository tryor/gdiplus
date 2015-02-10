package gdiplus

import (
	. "github.com/trygo/winapi"
	"unsafe"
)

//---------------------------------------------------------------------------
// Image file format identifiers
//---------------------------------------------------------------------------

var (
	ImageFormatUndefined *GUID = NewGUID(0xb96b3ca9, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatMemoryBMP *GUID = NewGUID(0xb96b3caa, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatBMP       *GUID = NewGUID(0xb96b3cab, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatEMF       *GUID = NewGUID(0xb96b3cac, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatWMF       *GUID = NewGUID(0xb96b3cad, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatJPEG      *GUID = NewGUID(0xb96b3cae, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatPNG       *GUID = NewGUID(0xb96b3caf, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatGIF       *GUID = NewGUID(0xb96b3cb0, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatTIFF      *GUID = NewGUID(0xb96b3cb1, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatEXIF      *GUID = NewGUID(0xb96b3cb2, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)
	ImageFormatIcon      *GUID = NewGUID(0xb96b3cb5, 0x0728, 0x11d3, 0x9d, 0x7b, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e)

	//---------------------------------------------------------------------------
	// Predefined multi-frame dimension IDs
	//---------------------------------------------------------------------------

	FrameDimensionTime       *GUID = NewGUID(0x6aedbd6d, 0x3fb5, 0x418a, 0x83, 0xa6, 0x7f, 0x45, 0x22, 0x9d, 0xc8, 0x72)
	FrameDimensionResolution *GUID = NewGUID(0x84236f7b, 0x3bd3, 0x428f, 0x8d, 0xab, 0x4e, 0xa1, 0x43, 0x9c, 0xa3, 0x15)
	FrameDimensionPage       *GUID = NewGUID(0x7462dc86, 0x6180, 0x4c7e, 0x8e, 0x3f, 0xee, 0x73, 0x33, 0xa7, 0xa4, 0x83)

	//---------------------------------------------------------------------------
	// Property sets
	//---------------------------------------------------------------------------

	FormatIDImageInformation *GUID = NewGUID(0xe5836cbe, 0x5eef, 0x4f1d, 0xac, 0xde, 0xae, 0x4c, 0x43, 0xb6, 0x08, 0xce)
	FormatIDJpegAppHeaders   *GUID = NewGUID(0x1c4afdcd, 0x6177, 0x43cf, 0xab, 0xc7, 0x5f, 0x51, 0xaf, 0x39, 0xee, 0x85)

	//---------------------------------------------------------------------------
	// Encoder parameter sets
	//---------------------------------------------------------------------------

	EncoderCompression      *GUID = NewGUID(0xe09d739d, 0xccd4, 0x44ee, 0x8e, 0xba, 0x3f, 0xbf, 0x8b, 0xe4, 0xfc, 0x58)
	EncoderColorDepth       *GUID = NewGUID(0x66087055, 0xad66, 0x4c7c, 0x9a, 0x18, 0x38, 0xa2, 0x31, 0x0b, 0x83, 0x37)
	EncoderScanMethod       *GUID = NewGUID(0x3a4e2661, 0x3109, 0x4e56, 0x85, 0x36, 0x42, 0xc1, 0x56, 0xe7, 0xdc, 0xfa)
	EncoderVersion          *GUID = NewGUID(0x24d18c76, 0x814a, 0x41a4, 0xbf, 0x53, 0x1c, 0x21, 0x9c, 0xcc, 0xf7, 0x97)
	EncoderRenderMethod     *GUID = NewGUID(0x6d42c53a, 0x229a, 0x4825, 0x8b, 0xb7, 0x5c, 0x99, 0xe2, 0xb9, 0xa8, 0xb8)
	EncoderQuality          *GUID = NewGUID(0x1d5be4b5, 0xfa4a, 0x452d, 0x9c, 0xdd, 0x5d, 0xb3, 0x51, 0x05, 0xe7, 0xeb)
	EncoderTransformation   *GUID = NewGUID(0x8d0eb2d1, 0xa58e, 0x4ea8, 0xaa, 0x14, 0x10, 0x80, 0x74, 0xb7, 0xb6, 0xf9)
	EncoderLuminanceTable   *GUID = NewGUID(0xedb33bce, 0x0266, 0x4a77, 0xb9, 0x04, 0x27, 0x21, 0x60, 0x99, 0xe7, 0x17)
	EncoderChrominanceTable *GUID = NewGUID(0xf2e455dc, 0x09b3, 0x4316, 0x82, 0x60, 0x67, 0x6a, 0xda, 0x32, 0x48, 0x1c)
	EncoderSaveFlag         *GUID = NewGUID(0x292266fc, 0xac40, 0x47bf, 0x8c, 0xfc, 0xa8, 0x5b, 0x89, 0xa6, 0x55, 0xde)

	//#if (GDIPVER >= 0x0110)
	EncoderColorSpace *GUID = NewGUID(0xae7a62a0, 0xee2c, 0x49d8, 0x9d, 0x7, 0x1b, 0xa8, 0xa9, 0x27, 0x59, 0x6e)
	EncoderImageItems *GUID = NewGUID(0x63875e13, 0x1f1d, 0x45ab, 0x91, 0x95, 0xa2, 0x9b, 0x60, 0x66, 0xa6, 0x50)
	EncoderSaveAsCMYK *GUID = NewGUID(0xa219bbc9, 0xa9d, 0x4005, 0xa3, 0xee, 0x3a, 0x42, 0x1b, 0x8b, 0xb0, 0x6c)
	//#endif //(GDIPVER >= 0x0110)

	CodecIImageBytes *GUID = NewGUID(0x025d1823, 0x6c7d, 0x447b, 0xbb, 0xdb, 0xa3, 0xcb, 0xc3, 0xdf, 0xa2, 0xfc)
)

//--------------------------------------------------------------------------
// ImageCodecInfo structure
//--------------------------------------------------------------------------

type ImageCodecInfo struct {
	Clsid             CLSID
	FormatID          GUID
	CodecName         *WCHAR
	DllName           *WCHAR
	FormatDescription *WCHAR
	FilenameExtension *WCHAR
	MimeType          *WCHAR
	Flags             DWORD
	Version           DWORD
	SigCount          DWORD
	SigSize           DWORD
	SigPattern        *BYTE
	SigMask           *BYTE
}

var ImageCodecInfo_SIZE = UINT(unsafe.Sizeof(ImageCodecInfo{}))

//--------------------------------------------------------------------------
// Information flags about image codecs
//--------------------------------------------------------------------------
type ImageCodecFlags INT

const (
	ImageCodecFlagsEncoder        ImageCodecFlags = 0x00000001
	ImageCodecFlagsDecoder        ImageCodecFlags = 0x00000002
	ImageCodecFlagsSupportBitmap  ImageCodecFlags = 0x00000004
	ImageCodecFlagsSupportVector  ImageCodecFlags = 0x00000008
	ImageCodecFlagsSeekableEncode ImageCodecFlags = 0x00000010
	ImageCodecFlagsBlockingDecode ImageCodecFlags = 0x00000020

	ImageCodecFlagsBuiltin ImageCodecFlags = 0x00010000
	ImageCodecFlagsSystem  ImageCodecFlags = 0x00020000
	ImageCodecFlagsUser    ImageCodecFlags = 0x00040000
)

//---------------------------------------------------------------------------
// Access modes used when calling Image::LockBits
//---------------------------------------------------------------------------
type ImageLockMode INT

const (
	ImageLockModeRead         ImageLockMode = 0x0001
	ImageLockModeWrite        ImageLockMode = 0x0002
	ImageLockModeUserInputBuf ImageLockMode = 0x0004
)

//---------------------------------------------------------------------------
// Information about image pixel data
//---------------------------------------------------------------------------

type BitmapData struct {
	Width       UINT
	Height      UINT
	Stride      INT
	PixelFormat PixelFormat
	Scan0       ULONG_PTR
	Reserved    UINT_PTR
}

//---------------------------------------------------------------------------
// Image flags
//---------------------------------------------------------------------------
type ImageFlags INT

const (
	ImageFlagsNone ImageFlags = 0

	// Low-word: shared with SINKFLAG_x

	ImageFlagsScalable          ImageFlags = 0x0001
	ImageFlagsHasAlpha          ImageFlags = 0x0002
	ImageFlagsHasTranslucent    ImageFlags = 0x0004
	ImageFlagsPartiallyScalable ImageFlags = 0x0008

	// Low-word: color space definition

	ImageFlagsColorSpaceRGB   ImageFlags = 0x0010
	ImageFlagsColorSpaceCMYK  ImageFlags = 0x0020
	ImageFlagsColorSpaceGRAY  ImageFlags = 0x0040
	ImageFlagsColorSpaceYCBCR ImageFlags = 0x0080
	ImageFlagsColorSpaceYCCK  ImageFlags = 0x0100

	// Low-word: image size info

	ImageFlagsHasRealDPI       ImageFlags = 0x1000
	ImageFlagsHasRealPixelSize ImageFlags = 0x2000

	// High-word

	ImageFlagsReadOnly ImageFlags = 0x00010000
	ImageFlagsCaching  ImageFlags = 0x00020000
)

type RotateFlipType INT

const (
	RotateNoneFlipNone RotateFlipType = 0
	Rotate90FlipNone   RotateFlipType = 1
	Rotate180FlipNone  RotateFlipType = 2
	Rotate270FlipNone  RotateFlipType = 3

	RotateNoneFlipX RotateFlipType = 4
	Rotate90FlipX   RotateFlipType = 5
	Rotate180FlipX  RotateFlipType = 6
	Rotate270FlipX  RotateFlipType = 7

	RotateNoneFlipY RotateFlipType = Rotate180FlipX
	Rotate90FlipY   RotateFlipType = Rotate270FlipX
	Rotate180FlipY  RotateFlipType = RotateNoneFlipX
	Rotate270FlipY  RotateFlipType = Rotate90FlipX

	RotateNoneFlipXY RotateFlipType = Rotate180FlipNone
	Rotate90FlipXY   RotateFlipType = Rotate270FlipNone
	Rotate180FlipXY  RotateFlipType = RotateNoneFlipNone
	Rotate270FlipXY  RotateFlipType = Rotate90FlipNone
)

//---------------------------------------------------------------------------
// Encoder Parameter structure
//---------------------------------------------------------------------------
type EncoderParameter struct {
	Guid           GUID     // GUID of the parameter
	NumberOfValues ULONG    // Number of the parameter values
	Type           ULONG    // Value type like ValueTypeLONG  etc.
	Value          DATA_PTR // A pointer to the parameter values
}

//---------------------------------------------------------------------------
// Encoder Parameters structure
//---------------------------------------------------------------------------
type EncoderParameters struct {
	Count     UINT                // Number of parameters in this structure
	Parameter [1]EncoderParameter // Parameter values
}

//#if (GDIPVER >= 0x0110)

type ItemDataPosition INT

const (
	ItemDataPositionAfterHeader  ItemDataPosition = 0x0
	ItemDataPositionAfterPalette ItemDataPosition = 0x1
	ItemDataPositionAfterBits    ItemDataPosition = 0x2
)

//---------------------------------------------------------------------------
// External Data Item
//---------------------------------------------------------------------------
type ImageItemData struct {
	Size     UINT     // size of the structure
	Position UINT     // flags describing how the data is to be used.
	Desc     DATA_PTR // description on how the data is to be saved. // it is different for every codec type.
	DescSize UINT     // size memory pointed by Desc
	Data     DATA_PTR // pointer to the data that is to be saved in the // file, could be anything saved directly.
	DataSize UINT     // size memory pointed by Data
	Cookie   UINT     // opaque for the apps data member used during // enumeration of image data items.
}

//#endif //(GDIPVER >= 0x0110)

//---------------------------------------------------------------------------
// Property Item
//---------------------------------------------------------------------------
type PropertyItem struct {
	Id     PROPID   // ID of this property
	Length ULONG    // Length of the property value, in bytes
	Type   WORD     // Type of the value, as one of TAG_TYPE_XXX // defined above
	Value  DATA_PTR // property value
}

var PropertyItem_SIZE = UINT(unsafe.Sizeof(PropertyItem{}))

const (
	//---------------------------------------------------------------------------
	// Image property types
	//---------------------------------------------------------------------------
	PropertyTagTypeByte      = 1
	PropertyTagTypeASCII     = 2
	PropertyTagTypeShort     = 3
	PropertyTagTypeLong      = 4
	PropertyTagTypeRational  = 5
	PropertyTagTypeUndefined = 7
	PropertyTagTypeSLONG     = 9
	PropertyTagTypeSRational = 10

	//---------------------------------------------------------------------------
	// Image property ID tags
	//---------------------------------------------------------------------------

	PropertyTagExifIFD = 0x8769
	PropertyTagGpsIFD  = 0x8825

	PropertyTagNewSubfileType        = 0x00FE
	PropertyTagSubfileType           = 0x00FF
	PropertyTagImageWidth            = 0x0100
	PropertyTagImageHeight           = 0x0101
	PropertyTagBitsPerSample         = 0x0102
	PropertyTagCompression           = 0x0103
	PropertyTagPhotometricInterp     = 0x0106
	PropertyTagThreshHolding         = 0x0107
	PropertyTagCellWidth             = 0x0108
	PropertyTagCellHeight            = 0x0109
	PropertyTagFillOrder             = 0x010A
	PropertyTagDocumentName          = 0x010D
	PropertyTagImageDescription      = 0x010E
	PropertyTagEquipMake             = 0x010F
	PropertyTagEquipModel            = 0x0110
	PropertyTagStripOffsets          = 0x0111
	PropertyTagOrientation           = 0x0112
	PropertyTagSamplesPerPixel       = 0x0115
	PropertyTagRowsPerStrip          = 0x0116
	PropertyTagStripBytesCount       = 0x0117
	PropertyTagMinSampleValue        = 0x0118
	PropertyTagMaxSampleValue        = 0x0119
	PropertyTagXResolution           = 0x011A // Image resolution in width direction
	PropertyTagYResolution           = 0x011B // Image resolution in height direction
	PropertyTagPlanarConfig          = 0x011C // Image data arrangement
	PropertyTagPageName              = 0x011D
	PropertyTagXPosition             = 0x011E
	PropertyTagYPosition             = 0x011F
	PropertyTagFreeOffset            = 0x0120
	PropertyTagFreeByteCounts        = 0x0121
	PropertyTagGrayResponseUnit      = 0x0122
	PropertyTagGrayResponseCurve     = 0x0123
	PropertyTagT4Option              = 0x0124
	PropertyTagT6Option              = 0x0125
	PropertyTagResolutionUnit        = 0x0128 // Unit of X and Y resolution
	PropertyTagPageNumber            = 0x0129
	PropertyTagTransferFuncition     = 0x012D
	PropertyTagSoftwareUsed          = 0x0131
	PropertyTagDateTime              = 0x0132
	PropertyTagArtist                = 0x013B
	PropertyTagHostComputer          = 0x013C
	PropertyTagPredictor             = 0x013D
	PropertyTagWhitePoint            = 0x013E
	PropertyTagPrimaryChromaticities = 0x013F
	PropertyTagColorMap              = 0x0140
	PropertyTagHalftoneHints         = 0x0141
	PropertyTagTileWidth             = 0x0142
	PropertyTagTileLength            = 0x0143
	PropertyTagTileOffset            = 0x0144
	PropertyTagTileByteCounts        = 0x0145
	PropertyTagInkSet                = 0x014C
	PropertyTagInkNames              = 0x014D
	PropertyTagNumberOfInks          = 0x014E
	PropertyTagDotRange              = 0x0150
	PropertyTagTargetPrinter         = 0x0151
	PropertyTagExtraSamples          = 0x0152
	PropertyTagSampleFormat          = 0x0153
	PropertyTagSMinSampleValue       = 0x0154
	PropertyTagSMaxSampleValue       = 0x0155
	PropertyTagTransferRange         = 0x0156

	PropertyTagJPEGProc               = 0x0200
	PropertyTagJPEGInterFormat        = 0x0201
	PropertyTagJPEGInterLength        = 0x0202
	PropertyTagJPEGRestartInterval    = 0x0203
	PropertyTagJPEGLosslessPredictors = 0x0205
	PropertyTagJPEGPointTransforms    = 0x0206
	PropertyTagJPEGQTables            = 0x0207
	PropertyTagJPEGDCTables           = 0x0208
	PropertyTagJPEGACTables           = 0x0209

	PropertyTagYCbCrCoefficients = 0x0211
	PropertyTagYCbCrSubsampling  = 0x0212
	PropertyTagYCbCrPositioning  = 0x0213
	PropertyTagREFBlackWhite     = 0x0214

	PropertyTagICCProfile = 0x8773 // This TAG is defined by ICC
	// for embedded ICC in TIFF
	PropertyTagGamma                = 0x0301
	PropertyTagICCProfileDescriptor = 0x0302
	PropertyTagSRGBRenderingIntent  = 0x0303

	PropertyTagImageTitle = 0x0320
	PropertyTagCopyright  = 0x8298

	// Extra TAGs (Like Adobe Image Information tags etc.)

	PropertyTagResolutionXUnit           = 0x5001
	PropertyTagResolutionYUnit           = 0x5002
	PropertyTagResolutionXLengthUnit     = 0x5003
	PropertyTagResolutionYLengthUnit     = 0x5004
	PropertyTagPrintFlags                = 0x5005
	PropertyTagPrintFlagsVersion         = 0x5006
	PropertyTagPrintFlagsCrop            = 0x5007
	PropertyTagPrintFlagsBleedWidth      = 0x5008
	PropertyTagPrintFlagsBleedWidthScale = 0x5009
	PropertyTagHalftoneLPI               = 0x500A
	PropertyTagHalftoneLPIUnit           = 0x500B
	PropertyTagHalftoneDegree            = 0x500C
	PropertyTagHalftoneShape             = 0x500D
	PropertyTagHalftoneMisc              = 0x500E
	PropertyTagHalftoneScreen            = 0x500F
	PropertyTagJPEGQuality               = 0x5010
	PropertyTagGridSize                  = 0x5011
	PropertyTagThumbnailFormat           = 0x5012 // 1 = JPEG, 0 = RAW RGB
	PropertyTagThumbnailWidth            = 0x5013
	PropertyTagThumbnailHeight           = 0x5014
	PropertyTagThumbnailColorDepth       = 0x5015
	PropertyTagThumbnailPlanes           = 0x5016
	PropertyTagThumbnailRawBytes         = 0x5017
	PropertyTagThumbnailSize             = 0x5018
	PropertyTagThumbnailCompressedSize   = 0x5019
	PropertyTagColorTransferFunction     = 0x501A
	PropertyTagThumbnailData             = 0x501B // RAW thumbnail bits in
	// JPEG format or RGB format
	// depends on
	// PropertyTagThumbnailFormat

	// Thumbnail related TAGs

	PropertyTagThumbnailImageWidth    = 0x5020 // Thumbnail width
	PropertyTagThumbnailImageHeight   = 0x5021 // Thumbnail height
	PropertyTagThumbnailBitsPerSample = 0x5022 // Number of bits per
	// component
	PropertyTagThumbnailCompression       = 0x5023 // Compression Scheme
	PropertyTagThumbnailPhotometricInterp = 0x5024 // Pixel composition
	PropertyTagThumbnailImageDescription  = 0x5025 // Image Tile
	PropertyTagThumbnailEquipMake         = 0x5026 // Manufacturer of Image
	// Input equipment
	PropertyTagThumbnailEquipModel = 0x5027 // Model of Image input
	// equipment
	PropertyTagThumbnailStripOffsets    = 0x5028 // Image data location
	PropertyTagThumbnailOrientation     = 0x5029 // Orientation of image
	PropertyTagThumbnailSamplesPerPixel = 0x502A // Number of components
	PropertyTagThumbnailRowsPerStrip    = 0x502B // Number of rows per strip
	PropertyTagThumbnailStripBytesCount = 0x502C // Bytes per compressed
	// strip
	PropertyTagThumbnailResolutionX = 0x502D // Resolution in width
	// direction
	PropertyTagThumbnailResolutionY = 0x502E // Resolution in height
	// direction
	PropertyTagThumbnailPlanarConfig   = 0x502F // Image data arrangement
	PropertyTagThumbnailResolutionUnit = 0x5030 // Unit of X and Y
	// Resolution
	PropertyTagThumbnailTransferFunction = 0x5031 // Transfer function
	PropertyTagThumbnailSoftwareUsed     = 0x5032 // Software used
	PropertyTagThumbnailDateTime         = 0x5033 // File change date and
	// time
	PropertyTagThumbnailArtist = 0x5034 // Person who created the
	// image
	PropertyTagThumbnailWhitePoint            = 0x5035 // White point chromaticity
	PropertyTagThumbnailPrimaryChromaticities = 0x5036
	// Chromaticities of
	// primaries
	PropertyTagThumbnailYCbCrCoefficients = 0x5037 // Color space transforma-
	// tion coefficients
	PropertyTagThumbnailYCbCrSubsampling = 0x5038 // Subsampling ratio of Y
	// to C
	PropertyTagThumbnailYCbCrPositioning = 0x5039 // Y and C position
	PropertyTagThumbnailRefBlackWhite    = 0x503A // Pair of black and white
	// reference values
	PropertyTagThumbnailCopyRight = 0x503B // CopyRight holder

	PropertyTagLuminanceTable   = 0x5090
	PropertyTagChrominanceTable = 0x5091

	PropertyTagFrameDelay = 0x5100
	PropertyTagLoopCount  = 0x5101

	//#if (GDIPVER >= =0x0110)
	PropertyTagGlobalPalette    = 0x5102
	PropertyTagIndexBackground  = 0x5103
	PropertyTagIndexTransparent = 0x5104
	//#endif //(GDIPVER >= =0x0110)

	PropertyTagPixelUnit        = 0x5110 // Unit specifier for pixel/unit
	PropertyTagPixelPerUnitX    = 0x5111 // Pixels per unit in X
	PropertyTagPixelPerUnitY    = 0x5112 // Pixels per unit in Y
	PropertyTagPaletteHistogram = 0x5113 // Palette histogram

	// EXIF specific tag

	PropertyTagExifExposureTime = 0x829A
	PropertyTagExifFNumber      = 0x829D

	PropertyTagExifExposureProg  = 0x8822
	PropertyTagExifSpectralSense = 0x8824
	PropertyTagExifISOSpeed      = 0x8827
	PropertyTagExifOECF          = 0x8828

	PropertyTagExifVer         = 0x9000
	PropertyTagExifDTOrig      = 0x9003 // Date & time of original
	PropertyTagExifDTDigitized = 0x9004 // Date & time of digital data generation

	PropertyTagExifCompConfig = 0x9101
	PropertyTagExifCompBPP    = 0x9102

	PropertyTagExifShutterSpeed = 0x9201
	PropertyTagExifAperture     = 0x9202
	PropertyTagExifBrightness   = 0x9203
	PropertyTagExifExposureBias = 0x9204
	PropertyTagExifMaxAperture  = 0x9205
	PropertyTagExifSubjectDist  = 0x9206
	PropertyTagExifMeteringMode = 0x9207
	PropertyTagExifLightSource  = 0x9208
	PropertyTagExifFlash        = 0x9209
	PropertyTagExifFocalLength  = 0x920A
	PropertyTagExifSubjectArea  = 0x9214 // exif 2.2 Subject Area
	PropertyTagExifMakerNote    = 0x927C
	PropertyTagExifUserComment  = 0x9286
	PropertyTagExifDTSubsec     = 0x9290 // Date & Time subseconds
	PropertyTagExifDTOrigSS     = 0x9291 // Date & Time original subseconds
	PropertyTagExifDTDigSS      = 0x9292 // Date & TIme digitized subseconds

	PropertyTagExifFPXVer        = 0xA000
	PropertyTagExifColorSpace    = 0xA001
	PropertyTagExifPixXDim       = 0xA002
	PropertyTagExifPixYDim       = 0xA003
	PropertyTagExifRelatedWav    = 0xA004 // related sound file
	PropertyTagExifInterop       = 0xA005
	PropertyTagExifFlashEnergy   = 0xA20B
	PropertyTagExifSpatialFR     = 0xA20C // Spatial Frequency Response
	PropertyTagExifFocalXRes     = 0xA20E // Focal Plane X Resolution
	PropertyTagExifFocalYRes     = 0xA20F // Focal Plane Y Resolution
	PropertyTagExifFocalResUnit  = 0xA210 // Focal Plane Resolution Unit
	PropertyTagExifSubjectLoc    = 0xA214
	PropertyTagExifExposureIndex = 0xA215
	PropertyTagExifSensingMethod = 0xA217
	PropertyTagExifFileSource    = 0xA300
	PropertyTagExifSceneType     = 0xA301
	PropertyTagExifCfaPattern    = 0xA302

	// New EXIF 2.2 properties

	PropertyTagExifCustomRendered        = 0xA401
	PropertyTagExifExposureMode          = 0xA402
	PropertyTagExifWhiteBalance          = 0xA403
	PropertyTagExifDigitalZoomRatio      = 0xA404
	PropertyTagExifFocalLengthIn35mmFilm = 0xA405
	PropertyTagExifSceneCaptureType      = 0xA406
	PropertyTagExifGainControl           = 0xA407
	PropertyTagExifContrast              = 0xA408
	PropertyTagExifSaturation            = 0xA409
	PropertyTagExifSharpness             = 0xA40A
	PropertyTagExifDeviceSettingDesc     = 0xA40B
	PropertyTagExifSubjectDistanceRange  = 0xA40C
	PropertyTagExifUniqueImageID         = 0xA420

	PropertyTagGpsVer              = 0x0000
	PropertyTagGpsLatitudeRef      = 0x0001
	PropertyTagGpsLatitude         = 0x0002
	PropertyTagGpsLongitudeRef     = 0x0003
	PropertyTagGpsLongitude        = 0x0004
	PropertyTagGpsAltitudeRef      = 0x0005
	PropertyTagGpsAltitude         = 0x0006
	PropertyTagGpsGpsTime          = 0x0007
	PropertyTagGpsGpsSatellites    = 0x0008
	PropertyTagGpsGpsStatus        = 0x0009
	PropertyTagGpsGpsMeasureMode   = 0x00A
	PropertyTagGpsGpsDop           = 0x000B // Measurement precision
	PropertyTagGpsSpeedRef         = 0x000C
	PropertyTagGpsSpeed            = 0x000D
	PropertyTagGpsTrackRef         = 0x000E
	PropertyTagGpsTrack            = 0x000F
	PropertyTagGpsImgDirRef        = 0x0010
	PropertyTagGpsImgDir           = 0x0011
	PropertyTagGpsMapDatum         = 0x0012
	PropertyTagGpsDestLatRef       = 0x0013
	PropertyTagGpsDestLat          = 0x0014
	PropertyTagGpsDestLongRef      = 0x0015
	PropertyTagGpsDestLong         = 0x0016
	PropertyTagGpsDestBearRef      = 0x0017
	PropertyTagGpsDestBear         = 0x0018
	PropertyTagGpsDestDistRef      = 0x0019
	PropertyTagGpsDestDist         = 0x001A
	PropertyTagGpsProcessingMethod = 0x001B
	PropertyTagGpsAreaInformation  = 0x001C
	PropertyTagGpsDate             = 0x001D
	PropertyTagGpsDifferential     = 0x001E
)
