package gdiplus

import (
	. "github.com/trygo/winapi"
)

//#ifndef DCR_REMOVE_INTERNAL
var DCR_REMOVE_INTERNAL bool

func DCR_USE_NEW_105760() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_127084() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_135429() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_140782() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_175866() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_152154() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_146933() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_145804() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_145139() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_145138() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_145135() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_140861() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_140857() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_140855() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_235072() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_186151() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_174340() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_186764() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_168772() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_125467() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_186091() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_197819() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_202903() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_137252() bool { return !DCR_REMOVE_INTERNAL }
func DCR_USE_NEW_188922() bool { return !DCR_REMOVE_INTERNAL }

func MinI(x, y INT) INT {
	if x < y {
		return x
	}
	return y
}

func MaxI(x, y INT) INT {
	if x > y {
		return x
	}
	return y
}

func Min(x, y REAL) REAL {
	if x < y {
		return x
	}
	return y
}

func Max(x, y REAL) REAL {
	if x > y {
		return x
	}
	return y
}
