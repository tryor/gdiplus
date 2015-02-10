package gdiplus

import (
	. "github.com/trygo/winapi"
)

type Matrix struct {
	GdiplusBase
	nativeMatrix *GpMatrix
}

func NewMatrix() (*Matrix, error) {
	matrix := &Matrix{}
	matrix.setStatus(GdipCreateMatrix(&matrix.nativeMatrix))
	return matrix, matrix.LastError
}

func (this *Matrix) Close() {
	if this.nativeMatrix != nil {
		GdipDeleteMatrix(this.nativeMatrix)
	}
}

func NewMatrix2(m11, m12, m21, m22, dx, dy REAL) (*Matrix, error) {
	matrix := &Matrix{}

	matrix.setStatus(GdipCreateMatrix2(m11, m12, m21, m22,
		dx, dy, &matrix.nativeMatrix))
	return matrix, matrix.LastError
}

func NewMatrix3(rect *RectF,
	dstplg *PointF) (*Matrix, error) {
	matrix := &Matrix{}
	matrix.setStatus(GdipCreateMatrix3(rect,
		dstplg, &matrix.nativeMatrix))
	return matrix, matrix.LastError
}

func NewMatrix3I(rect *Rect,
	dstplg *Point) (*Matrix, error) {
	matrix := &Matrix{}
	matrix.setStatus(GdipCreateMatrix3I(rect,
		dstplg, &matrix.nativeMatrix))
	return matrix, matrix.LastError
}

func (this *Matrix) Clone() *Matrix {
	matrix := &Matrix{}
	this.setStatus(GdipCloneMatrix(this.nativeMatrix,
		&matrix.nativeMatrix))
	return matrix
}

func (this *Matrix) GetElements() (ms []REAL) {
	ms = make([]REAL, 6)
	this.setStatus(GdipGetMatrixElements(this.nativeMatrix, ms))
	return
}

func (this *Matrix) SetElements(m11, m12, m21, m22,
	dx, dy REAL) Status {
	return this.setStatus(GdipSetMatrixElements(this.nativeMatrix,
		m11, m12, m21, m22, dx, dy))
}

func (this *Matrix) OffsetX() REAL {
	ms := this.GetElements()
	if this.LastResult == Ok {
		return ms[4]
	} else {
		return 0.0
	}
}

func (this *Matrix) OffsetY() REAL {
	ms := this.GetElements()
	if this.LastResult == Ok {
		return ms[5]
	} else {
		return 0.0
	}
}

func (this *Matrix) Reset() Status {
	// set identity matrix elements
	return this.setStatus(GdipSetMatrixElements(this.nativeMatrix,
		1.0, 0.0, 0.0, 1.0, 0.0, 0.0))
}

//order = MatrixOrderPrepend
func (this *Matrix) Multiply(matrix *Matrix,
	order MatrixOrder) Status {
	return this.setStatus(GdipMultiplyMatrix(this.nativeMatrix,
		matrix.nativeMatrix, order))
}

//order = MatrixOrderPrepend
func (this *Matrix) Translate(offsetX, offsetY REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipTranslateMatrix(this.nativeMatrix, offsetX,
		offsetY, order))
}

func (this *Matrix) Scale(scaleX, scaleY REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipScaleMatrix(this.nativeMatrix, scaleX,
		scaleY, order))
}

func (this *Matrix) Rotate(angle REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipRotateMatrix(this.nativeMatrix, angle, order))
}

//order = MatrixOrderPrepend
func (this *Matrix) RotateAt(angle REAL, center *PointF,
	order MatrixOrder) Status {
	if order == MatrixOrderPrepend {
		this.setStatus(GdipTranslateMatrix(this.nativeMatrix, center.X, center.Y, order))
		this.setStatus(GdipRotateMatrix(this.nativeMatrix, angle, order))
		return this.setStatus(GdipTranslateMatrix(this.nativeMatrix,
			-center.X, -center.Y, order))
	} else {
		this.setStatus(GdipTranslateMatrix(this.nativeMatrix,
			-center.X, -center.Y, order))
		this.setStatus(GdipRotateMatrix(this.nativeMatrix, angle, order))
		return this.setStatus(GdipTranslateMatrix(this.nativeMatrix,
			center.X, center.Y, order))
	}
}

func (this *Matrix) Shear(shearX, shearY REAL,
	order MatrixOrder) Status {
	return this.setStatus(GdipShearMatrix(this.nativeMatrix, shearX,
		shearY, order))
}

func (this *Matrix) Invert() Status {
	return this.setStatus(GdipInvertMatrix(this.nativeMatrix))
}

//float version
//pts IN OUT
func (this *Matrix) TransformPoints(pts []PointF) Status {
	return this.setStatus(GdipTransformMatrixPoints(this.nativeMatrix, pts))
}

//int version
func (this *Matrix) TransformPointsI(pts []Point) Status {
	return this.setStatus(GdipTransformMatrixPointsI(this.nativeMatrix, pts))
}

func (this *Matrix) TransformVectors(pts []PointF) Status {
	return this.setStatus(GdipVectorTransformMatrixPoints(
		this.nativeMatrix, pts))
}

func (this *Matrix) TransformVectorsI(pts []Point) Status {
	return this.setStatus(GdipVectorTransformMatrixPointsI(
		this.nativeMatrix, pts))
}

func (this *Matrix) IsInvertible() (result BOOL) {
	this.setStatus(GdipIsMatrixInvertible(this.nativeMatrix, &result))
	return
}

func (this *Matrix) IsIdentity() (result BOOL) {
	this.setStatus(GdipIsMatrixIdentity(this.nativeMatrix, &result))
	return
}

func (this *Matrix) Equals(matrix *Matrix) (result BOOL) {
	this.setStatus(GdipIsMatrixEqual(this.nativeMatrix,
		matrix.nativeMatrix, &result))
	return
}
