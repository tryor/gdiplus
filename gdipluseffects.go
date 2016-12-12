package gdiplus

import (
	"errors"
	"unsafe"

	. "github.com/tryor/winapi"
)

//#if (GDIPVER >= 0x0110)

//-----------------------------------------------------------------------------
// GDI+ effect GUIDs
//-----------------------------------------------------------------------------

var (
	// {633C80A4-1843-482b-9EF2-BE2834C5FDD4}
	BlurEffectGuid = &GUID{0x633c80a4, 0x1843, 0x482b, [8]BYTE{0x9e, 0xf2, 0xbe, 0x28, 0x34, 0xc5, 0xfd, 0xd4}}
	// {63CBF3EE-C526-402c-8F71-62C540BF5142}
	SharpenEffectGuid = &GUID{0x63cbf3ee, 0xc526, 0x402c, [8]BYTE{0x8f, 0x71, 0x62, 0xc5, 0x40, 0xbf, 0x51, 0x42}}

	// {718F2615-7933-40e3-A511-5F68FE14DD74}
	ColorMatrixEffectGuid = &GUID{0x718f2615, 0x7933, 0x40e3, [8]BYTE{0xa5, 0x11, 0x5f, 0x68, 0xfe, 0x14, 0xdd, 0x74}}

	// {A7CE72A9-0F7F-40d7-B3CC-D0C02D5C3212}
	ColorLUTEffectGuid = &GUID{0xa7ce72a9, 0xf7f, 0x40d7, [8]BYTE{0xb3, 0xcc, 0xd0, 0xc0, 0x2d, 0x5c, 0x32, 0x12}}

	// {D3A1DBE1-8EC4-4c17-9F4C-EA97AD1C343D}
	BrightnessContrastEffectGuid = &GUID{0xd3a1dbe1, 0x8ec4, 0x4c17, [8]BYTE{0x9f, 0x4c, 0xea, 0x97, 0xad, 0x1c, 0x34, 0x3d}}

	// {8B2DD6C3-EB07-4d87-A5F0-7108E26A9C5F}
	HueSaturationLightnessEffectGuid = &GUID{0x8b2dd6c3, 0xeb07, 0x4d87, [8]BYTE{0xa5, 0xf0, 0x71, 0x8, 0xe2, 0x6a, 0x9c, 0x5f}}

	// {99C354EC-2A31-4f3a-8C34-17A803B33A25}
	LevelsEffectGuid = &GUID{0x99c354ec, 0x2a31, 0x4f3a, [8]BYTE{0x8c, 0x34, 0x17, 0xa8, 0x3, 0xb3, 0x3a, 0x25}}

	// {1077AF00-2848-4441-9489-44AD4C2D7A2C}
	TintEffectGuid = &GUID{0x1077af00, 0x2848, 0x4441, [8]BYTE{0x94, 0x89, 0x44, 0xad, 0x4c, 0x2d, 0x7a, 0x2c}}

	// {537E597D-251E-48da-9664-29CA496B70F8}
	ColorBalanceEffectGuid = &GUID{0x537e597d, 0x251e, 0x48da, [8]BYTE{0x96, 0x64, 0x29, 0xca, 0x49, 0x6b, 0x70, 0xf8}}

	// {74D29D05-69A4-4266-9549-3CC52836B632}
	RedEyeCorrectionEffectGuid = &GUID{0x74d29d05, 0x69a4, 0x4266, [8]BYTE{0x95, 0x49, 0x3c, 0xc5, 0x28, 0x36, 0xb6, 0x32}}

	// {DD6A0022-58E4-4a67-9D9B-D48EB881A53D}
	ColorCurveEffectGuid = &GUID{0xdd6a0022, 0x58e4, 0x4a67, [8]BYTE{0x9d, 0x9b, 0xd4, 0x8e, 0xb8, 0x81, 0xa5, 0x3d}}
)

//-----------------------------------------------------------------------------

type SharpenParams struct {
	Radius FLOAT
	Amount FLOAT
}

var SharpenParams_SIZE = UINT(unsafe.Sizeof(SharpenParams{}))

type BlurParams struct {
	Radius     FLOAT
	ExpandEdge BOOL
}

var BlurParams_SIZE = UINT(unsafe.Sizeof(BlurParams{}))

type BrightnessContrastParams struct {
	BrightnessLevel INT
	ContrastLevel   INT
}

var BrightnessContrastParams_SIZE = UINT(unsafe.Sizeof(BrightnessContrastParams{}))

type RedEyeCorrectionParams struct {
	NumberOfAreas UINT
	Areas         []RECT
}

var RedEyeCorrectionParams_SIZE = UINT(unsafe.Sizeof(RedEyeCorrectionParams{}))

type HueSaturationLightnessParams struct {
	HueLevel        INT
	SaturationLevel INT
	LightnessLevel  INT
}

var HueSaturationLightnessParams_SIZE = UINT(unsafe.Sizeof(HueSaturationLightnessParams{}))

type TintParams struct {
	Hue    INT
	Amount INT
}

var TintParams_SIZE = UINT(unsafe.Sizeof(TintParams{}))

type LevelsParams struct {
	Highlight INT
	Midtone   INT
	Shadow    INT
}

var LevelsParams_SIZE = UINT(unsafe.Sizeof(LevelsParams{}))

type ColorBalanceParams struct {
	CyanRed      INT
	MagentaGreen INT
	YellowBlue   INT
}

var ColorBalanceParams_SIZE = UINT(unsafe.Sizeof(ColorBalanceParams{}))

type ColorLUTParams struct {
	// look up tables for each color channel.
	LutB ColorChannelLUT
	LutG ColorChannelLUT
	LutR ColorChannelLUT
	LutA ColorChannelLUT
}

var ColorLUTParams_SIZE = UINT(unsafe.Sizeof(ColorLUTParams{}))

type CurveAdjustments INT

const (
	AdjustExposure CurveAdjustments = iota
	AdjustDensity
	AdjustContrast
	AdjustHighlight
	AdjustShadow
	AdjustMidtone
	AdjustWhiteSaturation
	AdjustBlackSaturation
)

type CurveChannel INT

const (
	CurveChannelAll CurveChannel = iota
	CurveChannelRed
	CurveChannelGreen
	CurveChannelBlue
)

type ColorCurveParams struct {
	adjustment  CurveAdjustments
	channel     CurveChannel
	adjustValue INT
}

var ColorCurveParams_SIZE = UINT(unsafe.Sizeof(ColorCurveParams{}))

type CGpEffect interface{}

type IEffect interface {
	GetNativeEffect() *CGpEffect
	GetAuxDataSize() INT
	SetAuxDataSize(v INT)
	GetAuxData() DATA_PTR
	SetAuxData(ptr DATA_PTR)
	IsUseAuxData() BOOL
	SetUseAuxData(useAuxDataFlag BOOL)

	GetAuxDataPtr() *DATA_PTR
	GetAuxDataSizePtr() *INT
}

type Effect struct {
	GdiplusBase
	nativeEffect *CGpEffect
	auxDataSize  INT
	auxData      DATA_PTR
	useAuxData   BOOL
}

func NewEffect(guid *GUID) (*Effect, error) {
	effect := &Effect{}
	effect.setStatus(GdipCreateEffect(guid, &effect.nativeEffect))
	return effect, effect.LastError
}

func (this *Effect) Release() {
	if this.nativeEffect != nil {
		this.setStatus(GdipDeleteEffect(this.nativeEffect))
		this.nativeEffect = nil
	}
}

func (this *Effect) GetAuxDataPtr() *DATA_PTR {
	return &this.auxData
}
func (this *Effect) GetAuxDataSizePtr() *INT {
	return &this.auxDataSize
}

func (this *Effect) GetAuxDataSize() INT {
	return this.auxDataSize
}

func (this *Effect) SetAuxDataSize(v INT) {
	this.auxDataSize = v
}

func (this *Effect) GetAuxData() DATA_PTR {
	return this.auxData
}

func (this *Effect) SetAuxData(ptr DATA_PTR) {
	this.auxData = ptr
}

func (this *Effect) SetUseAuxData(useAuxDataFlag BOOL) {
	this.useAuxData = useAuxDataFlag
}

func (this *Effect) IsUseAuxData() BOOL {
	return this.useAuxData
}

func (this *Effect) GetParameterSize() (size UINT, status Status) {
	status = this.setStatus(GdipGetEffectParameterSize(this.nativeEffect, &size))
	return
}

func (this *Effect) GetNativeEffect() *CGpEffect {
	return this.nativeEffect
}

func (this *Effect) setParameters(params ULONG_PTR, size UINT) Status {
	return this.setStatus(GdipSetEffectParameters(this.nativeEffect, params, size))
}

func (this *Effect) getParameters(size *UINT, params ULONG_PTR) Status {
	return this.setStatus(GdipGetEffectParameters(this.nativeEffect, size, params))
}

type Blur struct {
	*Effect
}

func NewBlur() (*Blur, error) {
	effect, err := NewEffect(BlurEffectGuid)
	blur := &Blur{Effect: effect}
	return blur, err
}

func (this *Blur) SetParameters(parameters *BlurParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), BlurParams_SIZE)
}

/*
Usage:
	blur, _ := gdiplus.NewBlur()
	defer blur.Release()
	var bpsize UINT
	blur.GetParameterSize(&bpsize)
	blurbuffer := make([]byte, bpsize)
	blurParamsBuffer := (*gdiplus.BlurParams)(unsafe.Pointer(&blurbuffer[0]))
	blur.GetParameters(&bpsize, blurParamsBuffer)
*/
func (this *Blur) GetParameters(size *UINT, parameters *BlurParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type Sharpen struct {
	*Effect
}

func NewSharpen() (*Sharpen, error) {
	effect, err := NewEffect(SharpenEffectGuid)
	sharpen := &Sharpen{Effect: effect}
	return sharpen, err
}

func (this *Sharpen) SetParameters(parameters *SharpenParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), SharpenParams_SIZE)
}

func (this *Sharpen) GetParameters(size *UINT, parameters *SharpenParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type RedEyeCorrection struct {
	*Effect
}

func NewRedEyeCorrection() (*RedEyeCorrection, error) {
	effect, err := NewEffect(RedEyeCorrectionEffectGuid)
	obj := &RedEyeCorrection{Effect: effect}
	return obj, err
}

func (this *RedEyeCorrection) SetParameters(parameters *RedEyeCorrectionParams) Status {
	if parameters != nil {
		size := RedEyeCorrectionParams_SIZE + parameters.NumberOfAreas*RECT_SIZE
		return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), size)
	} else {
		return this.setStatus(InvalidParameter, errors.New(StatusText[InvalidParameter]))
	}
}

/*
USE:
	var psize UINT
	rec.GetParameterSize(&psize)
	buffer := make([]byte, psize)
	recParamsBuffer := (*gdiplus.RedEyeCorrectionParams)(unsafe.Pointer(&buffer[0]))
	rec.GetParameters(&psize, recParamsBuffer)
*/
func (this *RedEyeCorrection) GetParameters(size *UINT, parameters *RedEyeCorrectionParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

// Brightness/Contrast
type BrightnessContrast struct {
	*Effect
}

func NewBrightnessContrast() (*BrightnessContrast, error) {
	effect, err := NewEffect(BrightnessContrastEffectGuid)
	obj := &BrightnessContrast{Effect: effect}
	return obj, err
}

func (this *BrightnessContrast) SetParameters(parameters *BrightnessContrastParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), BrightnessContrastParams_SIZE)
}

func (this *BrightnessContrast) GetParameters(size *UINT, parameters *BrightnessContrastParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type HueSaturationLightness struct {
	*Effect
}

func NewHueSaturationLightness() (*HueSaturationLightness, error) {
	effect, err := NewEffect(HueSaturationLightnessEffectGuid)
	obj := &HueSaturationLightness{Effect: effect}
	return obj, err
}

func (this *HueSaturationLightness) SetParameters(parameters *HueSaturationLightnessParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), HueSaturationLightnessParams_SIZE)
}

func (this *HueSaturationLightness) GetParameters(size *UINT, parameters *HueSaturationLightnessParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type Levels struct {
	*Effect
}

func NewLevels() (*Levels, error) {
	effect, err := NewEffect(LevelsEffectGuid)
	obj := &Levels{Effect: effect}
	return obj, err
}

func (this *Levels) SetParameters(parameters *LevelsParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), LevelsParams_SIZE)
}

func (this *Levels) GetParameters(size *UINT, parameters *LevelsParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type Tint struct {
	*Effect
}

func NewTint() (*Tint, error) {
	effect, err := NewEffect(TintEffectGuid)
	obj := &Tint{Effect: effect}
	return obj, err
}

func (this *Tint) SetParameters(parameters *TintParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), TintParams_SIZE)
}

func (this *Tint) GetParameters(size *UINT, parameters *TintParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type ColorBalance struct {
	*Effect
}

func NewColorBalance() (*ColorBalance, error) {
	effect, err := NewEffect(ColorBalanceEffectGuid)
	obj := &ColorBalance{Effect: effect}
	return obj, err
}

func (this *ColorBalance) SetParameters(parameters *ColorBalanceParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), ColorBalanceParams_SIZE)
}

func (this *ColorBalance) GetParameters(size *UINT, parameters *ColorBalanceParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type ColorMatrixEffect struct {
	*Effect
}

func NewColorMatrixEffect() (*ColorMatrixEffect, error) {
	effect, err := NewEffect(ColorMatrixEffectGuid)
	obj := &ColorMatrixEffect{Effect: effect}
	return obj, err
}

func (this *ColorMatrixEffect) SetParameters(parameters *ColorMatrix) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), ColorMatrix_SIZE)
}

func (this *ColorMatrixEffect) GetParameters(size *UINT, parameters *ColorMatrix) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type ColorLUT struct {
	*Effect
}

func NewColorLUT() (*ColorLUT, error) {
	effect, err := NewEffect(ColorLUTEffectGuid)
	obj := &ColorLUT{Effect: effect}
	return obj, err
}

func (this *ColorLUT) SetParameters(parameters *ColorLUTParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), ColorLUTParams_SIZE)
}

func (this *ColorLUT) GetParameters(size *UINT, parameters *ColorLUTParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

type ColorCurve struct {
	*Effect
}

func NewColorCurve() (*ColorCurve, error) {
	effect, err := NewEffect(ColorCurveEffectGuid)
	obj := &ColorCurve{Effect: effect}
	return obj, err
}

func (this *ColorCurve) SetParameters(parameters *ColorCurveParams) Status {
	return this.Effect.setParameters(ULONG_PTR(unsafe.Pointer(parameters)), ColorCurveParams_SIZE)
}

func (this *ColorCurve) GetParameters(size *UINT, parameters *ColorCurveParams) Status {
	return this.Effect.getParameters(size, ULONG_PTR(unsafe.Pointer(parameters)))
}

//#endif //(GDIPVER >= 0x0110)
