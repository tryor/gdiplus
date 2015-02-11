package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

import (
	. "github.com/trygo/gdiplus"
	. "github.com/trygo/winapi"
	. "github.com/trygo/winapi/gdi"
)

func abortf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format, a...)
	os.Exit(1)
}

func abortErrNo(funcname string, err error) {
	abortf("%s failed: %d %s\n", funcname, err, err)
}

var (
	mh HANDLE
)

//var appn *Application

const Width, Height = 600, 450

var (
	hostDC   HDC
	bufferDC HDC
	hbitmap  HANDLE
	graphics *Graphics
	gpToken  ULONG_PTR
)

func startupGdiplus() {
	status, err := Startup(&gpToken, nil, nil)
	fmt.Println("Startup.status:", status)
	fmt.Println("Startup.err:", err)
}

////appn = NewApplication(hwnd, backBuffer)
func graphics_example(hwnd HANDLE) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v\n", r)
			for i := 1; ; i += 1 {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				log.Print(file, line)
			}
		}
	}()

	startupGdiplus()
	hostDC = GetDC(hwnd)

	// 创建双缓冲
	bufferDC = CreateCompatibleDC(HANDLE(hostDC))
	hbitmap = CreateCompatibleBitmap(hostDC, Width, Height)
	SelectObject(bufferDC, hbitmap)
	DeleteObject(hbitmap)

	var err error
	graphics, err = FromHDC(bufferDC)
	defer graphics.Close()
	fmt.Println("FromHDC.graphics:", graphics)
	fmt.Println("FromHDC.status:", graphics.LastResult)
	fmt.Println("FromHDC.err:", err)

	graphics.SetPageUnit(UnitPixel)
	graphics.SetSmoothingMode(SmoothingModeHighQuality)
	graphics.SetTextRenderingHint(TextRenderingHintClearTypeGridFit)
	//graphics.SetTextRenderingHint(TextRenderingHintAntiAlias)
	graphics.Clear(Color{Olive})

	pen, err := NewPen(Color{Argb: Red}, 3)
	defer pen.Close()
	fmt.Println("NewPen.err:", err)
	if err != nil {
		return
	}

	brush, err := NewSolidBrush(NewColor3(255, 200, 200, 100))
	defer brush.Close()
	fmt.Println("NewSolidBrush.brush:", brush)
	if err != nil {
		return
	}

	gpath, _ := NewGraphicsPath()
	defer gpath.Close()
	fmt.Println("gpath.status:", gpath.LastResult)
	fmt.Println("gpath.err:", gpath.LastError)

	gpath.AddRectangle(NewRectF(10, 10, 250, 100))
	fmt.Println("gpath.AddRectangle.status:", gpath.LastResult)
	fmt.Println("gpath.AddRectangle.err:", gpath.LastError)

	family, _ := NewFontFamily("宋体", nil) //"Courier New"
	defer family.Close()
	gpath.AddString("测试", family, FontStyleBold|FontStyleItalic, 80, &RectF{25, 25, 0.0, 0.0}, nil)
	fmt.Println("gpath.AddString.status:", gpath.LastResult)
	fmt.Println("gpath.AddString.err:", gpath.LastError)

	graphics.FillPath(brush, gpath)
	fmt.Println("FillPath.status:", graphics.LastResult)
	fmt.Println("FillPath.err:", graphics.LastError)

	graphics.DrawPath(pen, gpath)
	fmt.Println("DrawPath.status:", graphics.LastResult)
	fmt.Println("DrawPath.err:", graphics.LastError)

	font, _ := NewFont(family, 50, FontStyleBold|FontStyleItalic, UnitPixel)
	fmt.Println("NewFont.status:", font.LastResult)
	fmt.Println("NewFont.err:", font.LastError)
	defer font.Close()

	rect := &RectF{10, 150, 0, 0}
	rect, codepointsFitted, linesFilled, _ := graphics.MeasureString("测试", font, rect, nil)
	fmt.Println("MeasureString.status:", graphics.LastResult)
	fmt.Println("MeasureString.err:", graphics.LastError)
	fmt.Println("MeasureString.linesFilled:", linesFilled)
	fmt.Println("MeasureString.codepointsFitted:", codepointsFitted)
	fmt.Println("MeasureString.rect:", rect)

	graphics.DrawRectangle2(pen, rect)
	graphics.DrawString("测试", font, rect, nil, brush)
	fmt.Println("DrawString.status:", graphics.LastResult)
	fmt.Println("DrawString.err:", graphics.LastError)

	appPath, _ := os.Getwd()
	bitmap, _ := NewBitmap(appPath + "/penguins.jpg")
	defer bitmap.Close()
	fmt.Println("NewBitmap.status:", bitmap.LastResult)
	fmt.Println("NewBitmap.err:", bitmap.LastError)
	fmt.Println("NewBitmap.bitmap:", bitmap)

	graphics.DrawImageI4(bitmap, &Rect{200, 150, 200, 200})
	fmt.Println("DrawImageI4.status:", graphics.LastResult)
	fmt.Println("DrawImageI4.err:", graphics.LastError)

	graphics.DrawDriverString("测试TEST223", font, brush, &PointF{150, 150}, DriverStringOptionsRealizedAdvance|DriverStringOptionsVertical|DriverStringOptionsCmapLookup|DriverStringOptionsCompensateResolution, nil)
	fmt.Println("DrawDriverString.status:", graphics.LastResult)
	fmt.Println("DrawDriverString.err:", graphics.LastError)

	BitBlt(hostDC, 0, 0, Width, Height, bufferDC, 0, 0, SRCCOPY)
}

func WndProc_(hwnd HANDLE, msg uint32, wparam, lparam uintptr) (rc uintptr) {
	rc = DefWindowProc(hwnd, msg, wparam, lparam)
	return
}

// WinProc called by windows to notify us of all windows events we might be interested in.
func WndProc(hwnd HANDLE, msg uint32, wparam, lparam uintptr) (rc uintptr) {
	switch msg {
	case WM_CREATE:
		rc = DefWindowProc(hwnd, msg, wparam, lparam)
	case WM_CLOSE:
		Shutdown(gpToken)
		DestroyWindow(hwnd)
	case WM_COMMAND:
		switch HANDLE(lparam) {
		default:
			rc = DefWindowProc(hwnd, msg, wparam, lparam)
		}
	case WM_PAINT:
		BitBlt(hostDC, 0, 0, Width, Height, bufferDC, 0, 0, SRCCOPY)
		rc = DefWindowProc(hwnd, msg, wparam, lparam)
	case WM_DESTROY:
		PostQuitMessage(0)
	case WM_MOUSEMOVE:
		//appn.TrackMouseMoveEvent(int(LOWORD(lparam)), int(HIWORD(lparam)), easydraw.MButton(wparam))
	case WM_LBUTTONDOWN, WM_RBUTTONDOWN:
		//appn.TrackMousePressEvent(int(LOWORD(lparam)), int(HIWORD(lparam)), easydraw.MButton(wparam))
	case WM_LBUTTONUP, WM_RBUTTONUP:
		//appn.TrackMouseReleaseEvent(int(LOWORD(lparam)), int(HIWORD(lparam)), easydraw.MButton(wparam))
	case WM_LBUTTONDBLCLK:
		//onMouseLeftDoubleClick
	case WM_SETCURSOR:
		//if appn.SetCursor() {
		//	return
		//} else {
		//	return DefWindowProc(hwnd, msg, wparam, lparam)
		//}
		return DefWindowProc(hwnd, msg, wparam, lparam)
	case WM_KEYDOWN:
		//appn.TrackKeyPressEvent(int(wparam))
	case WM_KEYUP:
		//onKeyRelease(wParam)

	case WM_CHAR:
		//		s := string(int(wparam))
		//		enc := mahonia.NewDecoder("utf-8")
		//		v := enc.ConvertString(s)
		//		var c rune
		//		if v != "" {
		//			runes := bytes.Runes([]byte(v))
		//			c = runes[0]
		//		}
		//		fmt.Println("c:", c, string(c))
		//appn.TrackKeyCharEvent(c)

	default:
		rc = DefWindowProc(hwnd, msg, wparam, lparam)
	}
	return
}

func rungui() int {
	var e error

	// GetModuleHandle
	mh, e = GetModuleHandle(nil)
	if e != nil {
		abortErrNo("GetModuleHandle", e)
	}

	// Get icon we're going to use.
	myicon, e := LoadIcon(0, IDI_APPLICATION)
	if e != nil {
		abortErrNo("LoadIcon", e)
	}

	// Get cursor we're going to use.
	mycursor, e := LoadCursor(0, IDC_ARROW)
	if e != nil {
		abortErrNo("LoadCursor", e)
	}

	// Create callback
	wproc := syscall.NewCallback(WndProc)

	// RegisterClassEx
	wcname := syscall.StringToUTF16Ptr("my Window Class")
	var wc Wndclassex
	wc.Size = uint32(unsafe.Sizeof(wc))
	wc.WndProc = wproc
	wc.Instance = mh
	wc.Icon = myicon
	wc.Cursor = mycursor
	wc.Background = COLOR_BTNFACE + 1
	wc.MenuName = nil
	wc.ClassName = wcname
	wc.IconSm = myicon
	if _, e := RegisterClassEx(&wc); e != nil {
		abortErrNo("RegisterClassEx", e)
	}

	// CreateWindowEx
	wh, e := CreateWindowEx(
		WS_EX_CLIENTEDGE,
		wcname,
		syscall.StringToUTF16Ptr("My window"),
		WS_OVERLAPPEDWINDOW|WS_VISIBLE,
		CW_USEDEFAULT, CW_USEDEFAULT,
		Width+20, Height+40+25,
		0, 0, mh, 0)

	fmt.Printf("e %v\n", e)
	if e != nil {
		abortErrNo("CreateWindowEx", e)
	}
	fmt.Printf("main window handle is %x\n", wh)

	// ShowWindow
	ShowWindow(wh, SW_SHOWDEFAULT)

	if e := UpdateWindow(wh); e != nil {
		abortErrNo("UpdateWindow", e)
	}

	graphics_example(wh)

	// Process all windows messages until WM_QUIT.
	var m Msg
	for {
		r, e := GetMessage(&m, 0, 0, 0)
		if e != nil {
			abortErrNo("GetMessage", e)
		}
		if r == 0 {
			break
		}
		TranslateMessage(&m)
		DispatchMessage(&m)
	}
	return int(m.Wparam)
}

func main() {
	//	FreeConsole()
	rc := rungui()
	os.Exit(rc)
}
