package gdiplus

import (
	. "github.com/trygo/winapi"
)

type IFontCollection interface {
	GetNativeFontCollection() *GpFontCollection
	GetFamilyCount() (numFound INT)
	GetFamilies(numSought INT, gpfamilies []*FontFamily) Status
	Release()
}

type FontCollection struct {
	GdiplusBase
	nativeFontCollection *GpFontCollection
}

func (this *FontCollection) GetNativeFontCollection() *GpFontCollection {
	return this.nativeFontCollection
}

func (this *FontCollection) Release() {
}

func (this *FontCollection) GetFamilyCount() (numFound INT) {
	this.setStatus(GdipGetFontCollectionFamilyCount(this.nativeFontCollection, &numFound))
	return
}

func (this *FontCollection) GetFamilies(numSought INT, gpfamilies []*FontFamily) Status {
	if numSought <= 0 || gpfamilies == nil || len(gpfamilies) == 0 {
		return this.setStatus(InvalidParameter, nil)
	}

	numFound := INT(0)
	//GpFontFamily **nativeFamilyList = new GpFontFamily*[numSought];
	nativeFamilyList := make([]*GpFontFamily, numSought)

	this.setStatus(GdipGetFontCollectionFamilyList(
		this.nativeFontCollection,
		numSought,
		nativeFamilyList,
		&numFound))

	if this.LastResult == Ok {
		for i := INT(0); i < numFound; i++ {
			GdipCloneFontFamily(nativeFamilyList[i], &gpfamilies[i].nativeFamily)
		}
	}

	return this.LastResult
}

type InstalledFontCollection struct {
	FontCollection
}

func NewInstalledFontCollection() (*InstalledFontCollection, error) {
	fc := &InstalledFontCollection{}
	fc.setStatus(GdipNewInstalledFontCollection(&fc.nativeFontCollection))
	return fc, fc.LastError
}

func (this *InstalledFontCollection) Release() {
}

type PrivateFontCollection struct {
	FontCollection
}

func NewPrivateFontCollection() (*PrivateFontCollection, error) {
	fc := &PrivateFontCollection{}
	fc.setStatus(GdipNewPrivateFontCollection(&fc.nativeFontCollection))
	return fc, fc.LastError
}

func (this *PrivateFontCollection) Release() {
	if this.nativeFontCollection != nil {
		this.setStatus(GdipDeletePrivateFontCollection(&this.nativeFontCollection))
		this.nativeFontCollection = nil
	}
}

func (this *PrivateFontCollection) AddFontFile(filename string) Status {
	return this.setStatus(GdipPrivateAddFontFile(this.nativeFontCollection, filename))
}

func (this *PrivateFontCollection) AddMemoryFont(memory []byte) Status {
	return this.setStatus(GdipPrivateAddMemoryFont(this.nativeFontCollection, memory))
}
