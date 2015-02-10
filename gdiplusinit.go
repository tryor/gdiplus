package gdiplus

import (
	. "github.com/trygo/winapi"
)

// Used for debug event notification (debug builds only)

type DebugEventLevel uint

const (
	DebugEventLevelFatal DebugEventLevel = iota
	DebugEventLevelWarning
)

// Callback function that GDI+ can call, on debug builds, for assertions
// and warnings.

//typedef VOID (WINAPI *DebugEventProc)(DebugEventLevel level, CHAR *message);
type DebugEventProc func(level DebugEventLevel, message *CHAR)

// Notification functions which the user must call appropriately if
// "SuppressBackgroundThread" (below) is set.

//typedef Status (WINAPI *NotificationHookProc)(OUT ULONG_PTR *token);
//typedef VOID (WINAPI *NotificationUnhookProc)(ULONG_PTR token);
type NotificationHookProc func(token *ULONG_PTR) Status
type NotificationUnhookProc func(token ULONG_PTR)

// Input structure for GdiplusStartup()

type GdiplusStartupInput struct {
	GdiplusVersion           UINT32         // Must be 1
	DebugEventCallback       DebugEventProc // Ignored on free builds
	SuppressBackgroundThread BOOL           // FALSE unless you're prepared to call
	// the hook/unhook functions properly
	SuppressExternalCodecs BOOL // FALSE unless you want GDI+ only to use
	// its internal image codecs.

}

func NewGdiplusStartupInput(
	debugEventCallback DebugEventProc,
	suppressBackgroundThread BOOL,
	suppressExternalCodecs BOOL) *GdiplusStartupInput {
	return &GdiplusStartupInput{GdiplusVersion: 1, DebugEventCallback: debugEventCallback,
		SuppressBackgroundThread: suppressBackgroundThread,
		SuppressExternalCodecs:   suppressExternalCodecs}
}

// Output structure for GdiplusStartup()

type GdiplusStartupOutput struct {
	// The following 2 fields are NULL if SuppressBackgroundThread is FALSE.
	// Otherwise, they are functions which must be called appropriately to
	// replace the background thread.
	//
	// These should be called on the application's main message loop - i.e.
	// a message loop which is active for the lifetime of GDI+.
	// "NotificationHook" should be called before starting the loop,
	// and "NotificationUnhook" should be called after the loop ends.

	NotificationHook   NotificationHookProc
	NotificationUnhook NotificationUnhookProc
}

// GDI+ initialization. Must be called before GDI+ API's are used.
//
// token  - may not be NULL - accepts a token to be passed in the corresponding
//          GdiplusShutdown call.
// input  - may not be NULL
// output - may be NULL only if input->SuppressBackgroundThread is FALSE.

//extern "C" Status WINAPI GdiplusStartup(
//    OUT ULONG_PTR *token,
//    const GdiplusStartupInput *input,
//    OUT GdiplusStartupOutput *output);

func Startup(token *ULONG_PTR, input *GdiplusStartupInput, output *GdiplusStartupOutput) (Status, error) {
	if input == nil {
		input = NewGdiplusStartupInput(nil, false, false)
	}
	status, err := GdiplusStartup(token, input, output)
	return Status(status), err
}

// GDI+ termination. Must be called before GDI+ is unloaded. GDI+ API's may not
// be called after this.

//extern "C" VOID WINAPI GdiplusShutdown(ULONG_PTR token);

func Shutdown(token ULONG_PTR) (Status, error) {
	status, err := GdiplusShutdown(token)
	return Status(status), err
}
