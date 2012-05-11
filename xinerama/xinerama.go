// Package xinerama is the X client API for the XINERAMA extension.
package xinerama

/*
	This file was generated by xinerama.xml on May 10 2012 8:04:32pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the XINERAMA extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 8, "XINERAMA").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named XINERAMA could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["XINERAMA"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["XINERAMA"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["XINERAMA"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["XINERAMA"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["XINERAMA"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// 'ScreenInfo' struct definition
// Size: 8
type ScreenInfo struct {
	XOrg   int16
	YOrg   int16
	Width  uint16
	Height uint16
}

// Struct read ScreenInfo
func ScreenInfoRead(buf []byte, v *ScreenInfo) int {
	b := 0

	v.XOrg = int16(xgb.Get16(buf[b:]))
	b += 2

	v.YOrg = int16(xgb.Get16(buf[b:]))
	b += 2

	v.Width = xgb.Get16(buf[b:])
	b += 2

	v.Height = xgb.Get16(buf[b:])
	b += 2

	return b
}

// Struct list read ScreenInfo
func ScreenInfoReadList(buf []byte, dest []ScreenInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ScreenInfo{}
		b += ScreenInfoRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Struct write ScreenInfo
func (v ScreenInfo) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put16(buf[b:], uint16(v.XOrg))
	b += 2

	xgb.Put16(buf[b:], uint16(v.YOrg))
	b += 2

	xgb.Put16(buf[b:], v.Width)
	b += 2

	xgb.Put16(buf[b:], v.Height)
	b += 2

	return buf
}

// Write struct list ScreenInfo
func ScreenInfoListBytes(buf []byte, list []ScreenInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// Request QueryVersion
// size: 8
type QueryVersionCookie struct {
	*xgb.Cookie
}

func QueryVersion(c *xgb.Conn, Major byte, Minor byte) QueryVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, Major, Minor), cookie)
	return QueryVersionCookie{cookie}
}

func QueryVersionUnchecked(c *xgb.Conn, Major byte, Minor byte) QueryVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, Major, Minor), cookie)
	return QueryVersionCookie{cookie}
}

// Request reply for QueryVersion
// size: 12
type QueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Major uint16
	Minor uint16
}

// Waits and reads reply data from request QueryVersion
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// Read reply into structure from buffer for QueryVersion
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Major = xgb.Get16(buf[b:])
	b += 2

	v.Minor = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryVersion
func queryVersionRequest(c *xgb.Conn, Major byte, Minor byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XINERAMA"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Major
	b += 1

	buf[b] = Minor
	b += 1

	return buf
}

// Request GetState
// size: 8
type GetStateCookie struct {
	*xgb.Cookie
}

func GetState(c *xgb.Conn, Window xproto.Window) GetStateCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getStateRequest(c, Window), cookie)
	return GetStateCookie{cookie}
}

func GetStateUnchecked(c *xgb.Conn, Window xproto.Window) GetStateCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getStateRequest(c, Window), cookie)
	return GetStateCookie{cookie}
}

// Request reply for GetState
// size: 12
type GetStateReply struct {
	Sequence uint16
	Length   uint32
	State    byte
	Window   xproto.Window
}

// Waits and reads reply data from request GetState
func (cook GetStateCookie) Reply() (*GetStateReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getStateReply(buf), nil
}

// Read reply into structure from buffer for GetState
func getStateReply(buf []byte) *GetStateReply {
	v := new(GetStateReply)
	b := 1 // skip reply determinant

	v.State = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Window = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	return v
}

// Write request to wire for GetState
func getStateRequest(c *xgb.Conn, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XINERAMA"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// Request GetScreenCount
// size: 8
type GetScreenCountCookie struct {
	*xgb.Cookie
}

func GetScreenCount(c *xgb.Conn, Window xproto.Window) GetScreenCountCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getScreenCountRequest(c, Window), cookie)
	return GetScreenCountCookie{cookie}
}

func GetScreenCountUnchecked(c *xgb.Conn, Window xproto.Window) GetScreenCountCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getScreenCountRequest(c, Window), cookie)
	return GetScreenCountCookie{cookie}
}

// Request reply for GetScreenCount
// size: 12
type GetScreenCountReply struct {
	Sequence    uint16
	Length      uint32
	ScreenCount byte
	Window      xproto.Window
}

// Waits and reads reply data from request GetScreenCount
func (cook GetScreenCountCookie) Reply() (*GetScreenCountReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getScreenCountReply(buf), nil
}

// Read reply into structure from buffer for GetScreenCount
func getScreenCountReply(buf []byte) *GetScreenCountReply {
	v := new(GetScreenCountReply)
	b := 1 // skip reply determinant

	v.ScreenCount = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Window = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	return v
}

// Write request to wire for GetScreenCount
func getScreenCountRequest(c *xgb.Conn, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XINERAMA"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// Request GetScreenSize
// size: 12
type GetScreenSizeCookie struct {
	*xgb.Cookie
}

func GetScreenSize(c *xgb.Conn, Window xproto.Window, Screen uint32) GetScreenSizeCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getScreenSizeRequest(c, Window, Screen), cookie)
	return GetScreenSizeCookie{cookie}
}

func GetScreenSizeUnchecked(c *xgb.Conn, Window xproto.Window, Screen uint32) GetScreenSizeCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getScreenSizeRequest(c, Window, Screen), cookie)
	return GetScreenSizeCookie{cookie}
}

// Request reply for GetScreenSize
// size: 24
type GetScreenSizeReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Width  uint32
	Height uint32
	Window xproto.Window
	Screen uint32
}

// Waits and reads reply data from request GetScreenSize
func (cook GetScreenSizeCookie) Reply() (*GetScreenSizeReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getScreenSizeReply(buf), nil
}

// Read reply into structure from buffer for GetScreenSize
func getScreenSizeReply(buf []byte) *GetScreenSizeReply {
	v := new(GetScreenSizeReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Width = xgb.Get32(buf[b:])
	b += 4

	v.Height = xgb.Get32(buf[b:])
	b += 4

	v.Window = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	v.Screen = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for GetScreenSize
func getScreenSizeRequest(c *xgb.Conn, Window xproto.Window, Screen uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XINERAMA"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}

// Request IsActive
// size: 4
type IsActiveCookie struct {
	*xgb.Cookie
}

func IsActive(c *xgb.Conn) IsActiveCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(isActiveRequest(c), cookie)
	return IsActiveCookie{cookie}
}

func IsActiveUnchecked(c *xgb.Conn) IsActiveCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(isActiveRequest(c), cookie)
	return IsActiveCookie{cookie}
}

// Request reply for IsActive
// size: 12
type IsActiveReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	State uint32
}

// Waits and reads reply data from request IsActive
func (cook IsActiveCookie) Reply() (*IsActiveReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return isActiveReply(buf), nil
}

// Read reply into structure from buffer for IsActive
func isActiveReply(buf []byte) *IsActiveReply {
	v := new(IsActiveReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.State = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for IsActive
func isActiveRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XINERAMA"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request QueryScreens
// size: 4
type QueryScreensCookie struct {
	*xgb.Cookie
}

func QueryScreens(c *xgb.Conn) QueryScreensCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryScreensRequest(c), cookie)
	return QueryScreensCookie{cookie}
}

func QueryScreensUnchecked(c *xgb.Conn) QueryScreensCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryScreensRequest(c), cookie)
	return QueryScreensCookie{cookie}
}

// Request reply for QueryScreens
// size: (32 + xgb.Pad((int(Number) * 8)))
type QueryScreensReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Number uint32
	// padding: 20 bytes
	ScreenInfo []ScreenInfo // size: xgb.Pad((int(Number) * 8))
}

// Waits and reads reply data from request QueryScreens
func (cook QueryScreensCookie) Reply() (*QueryScreensReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryScreensReply(buf), nil
}

// Read reply into structure from buffer for QueryScreens
func queryScreensReply(buf []byte) *QueryScreensReply {
	v := new(QueryScreensReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Number = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.ScreenInfo = make([]ScreenInfo, v.Number)
	b += ScreenInfoReadList(buf[b:], v.ScreenInfo)

	return v
}

// Write request to wire for QueryScreens
func queryScreensRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XINERAMA"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
