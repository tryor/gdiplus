package gdiplus

import (
	. "github.com/trygo/winapi"
	"unsafe"
)

/*
 * 32-bit and 64-bit ARGB pixel value
 */

type ARGB uint32

const ARGB_SIZE = 4

type ARGB64 uint64

const ALPHA_SHIFT = 24
const RED_SHIFT = 16
const GREEN_SHIFT = 8
const BLUE_SHIFT = 0
const ALPHA_MASK = (ARGB(0xff) << ALPHA_SHIFT)

// In-memory pixel data formats:
// bits 0-7 = format index
// bits 8-15 = pixel size (in bits)
// bits 16-23 = flags
// bits 24-31 = reserved

type PixelFormat INT

const (
	PixelFormatIndexed   PixelFormat = 0x00010000 // Indexes into a palette
	PixelFormatGDI       PixelFormat = 0x00020000 // Is a GDI-supported format
	PixelFormatAlpha     PixelFormat = 0x00040000 // Has an alpha component
	PixelFormatPAlpha    PixelFormat = 0x00080000 // Pre-multiplied alpha
	PixelFormatExtended  PixelFormat = 0x00100000 // Extended color 16 bits/channel
	PixelFormatCanonical PixelFormat = 0x00200000

	PixelFormatUndefined PixelFormat = 0
	PixelFormatDontCare  PixelFormat = 0
)

var (
	PixelFormat1bppIndexed    PixelFormat = (1 | (1 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat4bppIndexed    PixelFormat = (2 | (4 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat8bppIndexed    PixelFormat = (3 | (8 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat16bppGrayScale PixelFormat = (4 | (16 << 8) | PixelFormatExtended)
	PixelFormat16bppRGB555    PixelFormat = (5 | (16 << 8) | PixelFormatGDI)
	PixelFormat16bppRGB565    PixelFormat = (6 | (16 << 8) | PixelFormatGDI)
	PixelFormat16bppARGB1555  PixelFormat = (7 | (16 << 8) | PixelFormatAlpha | PixelFormatGDI)
	PixelFormat24bppRGB       PixelFormat = (8 | (24 << 8) | PixelFormatGDI)
	PixelFormat32bppRGB       PixelFormat = (9 | (32 << 8) | PixelFormatGDI)
	PixelFormat32bppARGB      PixelFormat = (10 | (32 << 8) | PixelFormatAlpha | PixelFormatGDI | PixelFormatCanonical)
	PixelFormat32bppPARGB     PixelFormat = (11 | (32 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatGDI)
	PixelFormat48bppRGB       PixelFormat = (12 | (48 << 8) | PixelFormatExtended)
	PixelFormat64bppARGB      PixelFormat = (13 | (64 << 8) | PixelFormatAlpha | PixelFormatCanonical | PixelFormatExtended)
	PixelFormat64bppPARGB     PixelFormat = (14 | (64 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatExtended)
	PixelFormat32bppCMYK      PixelFormat = (15 | (32 << 8))
	PixelFormatMax            PixelFormat = 16
)

func GetPixelFormatSize(pixfmt PixelFormat) UINT {
	return UINT((pixfmt >> 8) & 0xff)
}

func IsIndexedPixelFormat(pixfmt PixelFormat) BOOL {
	return (pixfmt & PixelFormatIndexed) != 0
}

func IsAlphaPixelFormat(pixfmt PixelFormat) BOOL {
	return (pixfmt & PixelFormatAlpha) != 0
}

func IsExtendedPixelFormat(pixfmt PixelFormat) BOOL {
	return (pixfmt & PixelFormatExtended) != 0
}

//--------------------------------------------------------------------------
// Determine if the Pixel Format is Canonical format:
//   PixelFormat32bppARGB
//   PixelFormat32bppPARGB
//   PixelFormat64bppARGB
//   PixelFormat64bppPARGB
//--------------------------------------------------------------------------

func IsCanonicalPixelFormat(pixfmt PixelFormat) BOOL {
	return (pixfmt & PixelFormatCanonical) != 0
}

type PaletteFlags INT

const (
	PaletteFlagsHasAlpha  PaletteFlags = 0x0001
	PaletteFlagsGrayScale PaletteFlags = 0x0002
	PaletteFlagsHalftone  PaletteFlags = 0x0004
)

type ColorPalette struct {
	Flags   UINT    // Palette flags
	Count   UINT    // Number of color entries
	Entries [1]ARGB // Palette color entries
}

var ColorPalette_SIZE = UINT(unsafe.Sizeof(ColorPalette{}))

//#if (GDIPVER >= 0x0110)
//----------------------------------------------------------------------------
// Color format conversion parameters
//----------------------------------------------------------------------------
type PaletteType INT

const (
	// Arbitrary custom palette provided by caller.

	PaletteTypeCustom PaletteType = 0

	// Optimal palette generated using a median-cut algorithm.

	PaletteTypeOptimal PaletteType = 1

	// Black and white palette.

	PaletteTypeFixedBW PaletteType = 2

	// Symmetric halftone palettes.
	// Each of these halftone palettes will be a superset of the system palette.
	// E.g. Halftone8 will have it's 8-color on-off primaries and the 16 system
	// colors added. With duplicates removed, that leaves 16 colors.

	PaletteTypeFixedHalftone8   PaletteType = 3 // 8-color on-off primaries
	PaletteTypeFixedHalftone27  PaletteType = 4 // 3 intensity levels of each color
	PaletteTypeFixedHalftone64  PaletteType = 5 // 4 intensity levels of each color
	PaletteTypeFixedHalftone125 PaletteType = 6 // 5 intensity levels of each color
	PaletteTypeFixedHalftone216 PaletteType = 7 // 6 intensity levels of each color

	// Assymetric halftone palettes.
	// These are somewhat less useful than the symmetric ones, but are
	// included for completeness. These do not include all of the system
	// colors.

	PaletteTypeFixedHalftone252 PaletteType = 8 // 6-red, 7-green, 6-blue intensities
	PaletteTypeFixedHalftone256 PaletteType = 9 // 8-red, 8-green, 4-blue intensities
)

type DitherType INT

const (
	DitherTypeNone DitherType = 0

	// Solid color - picks the nearest matching color with no attempt to
	// halftone or dither. May be used on an arbitrary palette.

	DitherTypeSolid DitherType = 1

	// Ordered dithers and spiral dithers must be used with a fixed palette.

	// NOTE: DitherOrdered4x4 is unique in that it may apply to 16bpp
	// conversions also.

	DitherTypeOrdered4x4 DitherType = 2

	DitherTypeOrdered8x8    DitherType = 3
	DitherTypeOrdered16x16  DitherType = 4
	DitherTypeSpiral4x4     DitherType = 5
	DitherTypeSpiral8x8     DitherType = 6
	DitherTypeDualSpiral4x4 DitherType = 7
	DitherTypeDualSpiral8x8 DitherType = 8

	// Error diffusion. May be used with any palette.

	DitherTypeErrorDiffusion DitherType = 9

	DitherTypeMax DitherType = 10
)

//#endif //(GDIPVER >= 0x0110)
