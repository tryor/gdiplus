package gdiplus

import (
	. "github.com/trygo/winapi"
)

type StringFormat struct {
	GdiplusBase
	nativeFormat *GpStringFormat
}

func NewStringFormat() (*StringFormat, error) {
	var formatFlags INT = 0
	var language LANGID = LANG_NEUTRAL
	return NewStringFormat2(formatFlags, language)
}

func NewStringFormat2(formatFlags INT, language LANGID) (*StringFormat, error) {
	format := &StringFormat{}
	format.setStatus(GdipCreateStringFormat(formatFlags, language, &format.nativeFormat))
	return format, format.LastError

}

func NewStringFormat3(strfmt *StringFormat) (*StringFormat, error) {
	format := &StringFormat{}

	if strfmt == nil {
		format.setStatus(InvalidParameter, nil)
		return format, format.LastError
	}

	format.setStatus(GdipCloneStringFormat(
		strfmt.nativeFormat,
		&format.nativeFormat))

	return format, format.LastError
}

func (this *StringFormat) Release() {
	if this.nativeFormat != nil {
		this.setStatus(GdipDeleteStringFormat(this.nativeFormat))
	}
}

func (this *StringFormat) Clone() *StringFormat {
	format := &StringFormat{}
	status, _ := GdipCloneStringFormat(
		this.nativeFormat, &format.nativeFormat)
	if status == Ok {
		return format
	}
	return nil
}

func (this *StringFormat) SetFormatFlags(flags INT) Status {
	return this.setStatus(GdipSetStringFormatFlags(
		this.nativeFormat,
		flags))
}

func (this *StringFormat) GetFormatFlags() (flags INT) {
	this.setStatus(GdipGetStringFormatFlags(this.nativeFormat, &flags))
	return
}

func (this *StringFormat) SetAlignment(align StringAlignment) Status {
	return this.setStatus(GdipSetStringFormatAlign(
		this.nativeFormat,
		align))
}

func (this *StringFormat) GetAlignment() (align StringAlignment) {
	this.setStatus(GdipGetStringFormatAlign(
		this.nativeFormat, &align))
	return
}

func (this *StringFormat) SetLineAlignment(align StringAlignment) Status {
	return this.setStatus(GdipSetStringFormatLineAlign(
		this.nativeFormat,
		align))
}

func (this *StringFormat) GetLineAlignment() (align StringAlignment) {
	this.setStatus(GdipGetStringFormatLineAlign(
		this.nativeFormat,
		&align))
	return
}

func (this *StringFormat) SetHotkeyPrefix(hotkeyPrefix HotkeyPrefix) Status {
	return this.setStatus(GdipSetStringFormatHotkeyPrefix(
		this.nativeFormat,
		INT(hotkeyPrefix)))
}

func (this *StringFormat) GetHotkeyPrefix() (hotkeyPrefix HotkeyPrefix) {
	this.setStatus(GdipGetStringFormatHotkeyPrefix(
		this.nativeFormat, (*INT)(&hotkeyPrefix)))
	return
}

func (this *StringFormat) SetTabStops(
	firstTabOffset REAL,
	tabStops []REAL) Status {
	return this.setStatus(GdipSetStringFormatTabStops(
		this.nativeFormat,
		firstTabOffset,
		tabStops))
}

func (this *StringFormat) GetTabStopCount() (count INT) {
	this.setStatus(GdipGetStringFormatTabStopCount(this.nativeFormat, &count))
	return
}

func (this *StringFormat) GetTabStops(
	count INT) (firstTabOffset REAL, tabStops []REAL, status Status) {
	tabStops = make([]REAL, count)
	status = this.setStatus(GdipGetStringFormatTabStops(
		this.nativeFormat,
		count,
		&firstTabOffset,
		tabStops))
	return
}

func (this *StringFormat) SetDigitSubstitution(
	language LANGID,
	substitute StringDigitSubstitute) Status {
	return this.setStatus(GdipSetStringFormatDigitSubstitution(
		this.nativeFormat,
		language,
		substitute))
}

func (this *StringFormat) GetDigitSubstitutionLanguage() (language LANGID) {
	this.setStatus(GdipGetStringFormatDigitSubstitution(
		this.nativeFormat,
		&language,
		nil))
	return
}

func (this *StringFormat) GetDigitSubstitutionMethod() (substitute StringDigitSubstitute) {
	this.setStatus(GdipGetStringFormatDigitSubstitution(
		this.nativeFormat,
		nil,
		&substitute))
	return
}

func (this *StringFormat) SetTrimming(trimming StringTrimming) Status {
	return this.setStatus(GdipSetStringFormatTrimming(
		this.nativeFormat,
		trimming))
}

func (this *StringFormat) GetTrimming() (trimming StringTrimming) {
	this.setStatus(GdipGetStringFormatTrimming(
		this.nativeFormat,
		&trimming))
	return
}

func (this *StringFormat) SetMeasurableCharacterRanges(ranges []CharacterRange) Status {
	return this.setStatus(GdipSetStringFormatMeasurableCharacterRanges(
		this.nativeFormat,
		INT(len(ranges)),
		&ranges[0]))
}

func (this *StringFormat) GetMeasurableCharacterRangeCount() (count INT) {
	this.setStatus(GdipGetStringFormatMeasurableCharacterRangeCount(
		this.nativeFormat,
		&count))
	return
}

var genericTypographicStringFormat *StringFormat
var genericDefaultStringFormat *StringFormat

func GenericDefault() *StringFormat {

	if genericDefaultStringFormat != nil {
		return genericDefaultStringFormat
	}

	genericDefaultStringFormat = &StringFormat{}
	genericDefaultStringFormat.setStatus(GdipStringFormatGetGenericDefault(&genericDefaultStringFormat.nativeFormat))

	return genericDefaultStringFormat

}
func GenericTypographic() *StringFormat {

	if genericTypographicStringFormat != nil {
		return genericTypographicStringFormat
	}

	genericTypographicStringFormat = &StringFormat{}
	genericTypographicStringFormat.setStatus(GdipStringFormatGetGenericTypographic(&genericTypographicStringFormat.nativeFormat))

	return genericTypographicStringFormat
}
