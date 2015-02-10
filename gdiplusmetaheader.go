package gdiplus

import (
	. "github.com/trygo/winapi"
	. "github.com/trygo/winapi/gdi"
	"unsafe"
)

type ENHMETAHEADER3 struct {
	IType DWORD // Record type EMR_HEADER
	NSize DWORD // Record size in bytes.  This may be greater
	// than the sizeof(ENHMETAHEADER).
	RclBounds  RECTL // Inclusive-inclusive bounds in device units
	RclFrame   RECTL // Inclusive-inclusive Picture Frame .01mm unit
	DSignature DWORD // Signature.  Must be ENHMETA_SIGNATURE.
	NVersion   DWORD // Version number
	NBytes     DWORD // Size of the metafile in bytes
	NRecords   DWORD // Number of records in the metafile
	NHandles   WORD  // Number of handles in the handle table
	// Handle index zero is reserved.
	SReserved    WORD  // Reserved.  Must be zero.
	NDescription DWORD // Number of chars in the unicode desc string
	// This is 0 if there is no description string
	OffDescription DWORD // Offset to the metafile description record.
	// This is 0 if there is no description string
	NPalEntries    DWORD // Number of entries in the metafile palette.
	SzlDevice      SIZEL // Size of the reference device in pels
	SzlMillimeters SIZEL // Size of the reference device in millimeters
}

var ENHMETAHEADER3_SIZE = UINT(unsafe.Sizeof(ENHMETAHEADER3{}))

type PWMFRect16 struct {
	Left   INT16
	Top    INT16
	Right  INT16
	Bottom INT16
}

type WmfPlaceableFileHeader struct {
	Key         UINT32     // GDIP_WMF_PLACEABLEKEY
	Hmf         INT16      // Metafile HANDLE number (always 0)
	BoundingBox PWMFRect16 // Coordinates in metafile units
	Inch        INT16      // Number of metafile units per inch
	Reserved    UINT32     // Reserved (always 0)
	Checksum    INT16      // Checksum value for previous 10 WORDs
}

// Key contains a special identification value that indicates the presence
// of a placeable metafile header and is always 0x9AC6CDD7.

// Handle is used to stored the handle of the metafile in memory. When written
// to disk, this field is not used and will always contains the value 0.

// Left, Top, Right, and Bottom contain the coordinates of the upper-left
// and lower-right corners of the image on the output device. These are
// measured in twips.

// A twip (meaning "twentieth of a point") is the logical unit of measurement
// used in Windows Metafiles. A twip is equal to 1/1440 of an inch. Thus 720
// twips equal 1/2 inch, while 32,768 twips is 22.75 inches.

// Inch contains the number of twips per inch used to represent the image.
// Normally, there are 1440 twips per inch; however, this number may be
// changed to scale the image. A value of 720 indicates that the image is
// double its normal size, or scaled to a factor of 2:1. A value of 360
// indicates a scale of 4:1, while a value of 2880 indicates that the image
// is scaled down in size by a factor of two. A value of 1440 indicates
// a 1:1 scale ratio.

// Reserved is not used and is always set to 0.

// Checksum contains a checksum value for the previous 10 WORDs in the header.
// This value can be used in an attempt to detect if the metafile has become
// corrupted. The checksum is calculated by XORing each WORD value to an
// initial value of 0.

// If the metafile was recorded with a reference Hdc that was a display.

const GDIP_EMFPLUSFLAGS_DISPLAY = 0x00000001

type HeaderUnion struct {
	EmfHeader ENHMETAHEADER3 //sizeof is 88
	//WmfHeader METAHEADER   //sizeof is 24
}

type MetafileHeader struct {
	Type         MetafileType
	Size         UINT // Size of the metafile (in bytes)
	Version      UINT // EMF+, EMF, or WMF version
	EmfPlusFlags UINT
	DpiX         REAL
	DpiY         REAL
	X            INT // Bounds in device units
	Y            INT
	Width        INT
	Height       INT
	//union
	//{
	//    METAHEADER      WmfHeader;
	//    ENHMETAHEADER3  EmfHeader;
	//};
	Header            HeaderUnion
	EmfPlusHeaderSize INT // size of the EMF+ header in file
	LogicalDpiX       INT // Logical Dpi of reference Hdc
	LogicalDpiY       INT // usually valid only for EMF+

}

func (this *MetafileHeader) GetBounds() *Rect {
	return &Rect{this.X, this.Y, this.Width, this.Height}
}

// Is it any type of WMF (standard or Placeable Metafile)?
func (this *MetafileHeader) IsWmf() bool {
	return ((this.Type == MetafileTypeWmf) || (this.Type == MetafileTypeWmfPlaceable))
}

// Is this an Placeable Metafile?
func (this *MetafileHeader) IsWmfPlaceable() bool { return (this.Type == MetafileTypeWmfPlaceable) }

// Is this an EMF (not an EMF+)?
func (this *MetafileHeader) IsEmf() bool { return (this.Type == MetafileTypeEmf) }

// Is this an EMF or EMF+ file?
func (this *MetafileHeader) IsEmfOrEmfPlus() bool { return (this.Type >= MetafileTypeEmf) }

// Is this an EMF+ file?
func (this *MetafileHeader) IsEmfPlus() bool { return (this.Type >= MetafileTypeEmfPlusOnly) }

// Is this an EMF+ dual (has dual, down-level records) file?
func (this *MetafileHeader) IsEmfPlusDual() bool { return (this.Type == MetafileTypeEmfPlusDual) }

// Is this an EMF+ only (no dual records) file?
func (this *MetafileHeader) IsEmfPlusOnly() bool { return (this.Type == MetafileTypeEmfPlusOnly) }

// If it's an EMF+ file, was it recorded against a display Hdc?
func (this *MetafileHeader) IsDisplay() bool {
	return (this.IsEmfPlus() &&
		((this.EmfPlusFlags & GDIP_EMFPLUSFLAGS_DISPLAY) != 0))
}

// Get the WMF header of the metafile (if it is a WMF)
func (this *MetafileHeader) GetWmfHeader() *METAHEADER {
	if this.IsWmf() {
		return (*METAHEADER)(unsafe.Pointer(&this.Header))
	}
	return nil
}

// Get the EMF header of the metafile (if it is an EMF)
func (this *MetafileHeader) GetEmfHeader() *ENHMETAHEADER3 {
	if this.IsEmfOrEmfPlus() {
		return (*ENHMETAHEADER3)(unsafe.Pointer(&this.Header))
	}
	return nil
}
