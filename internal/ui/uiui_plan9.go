// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"errors"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2/internal/color"
	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver"
)

type graphicsDriverCreatorImpl struct {
	canvas     js.Value
	colorSpace color.ColorSpace
}

func (g *graphicsDriverCreatorImpl) newAuto() (graphicsdriver.Graphics, GraphicsLibrary, error) {
	panic("newAuto not implemented")
	return nil, 0, errors.New("no")
}

func (g *graphicsDriverCreatorImpl) newOpenGL() (graphicsdriver.Graphics, error) {
	panic("newOpenGL not implemented")
	return nil, errors.New("no")
}

func (*graphicsDriverCreatorImpl) newDirectX() (graphicsdriver.Graphics, error) {
	panic("newDirectX not implemented")
	return nil, errors.New("no")
}

func (*graphicsDriverCreatorImpl) newMetal() (graphicsdriver.Graphics, error) {
	panic("newMetal not implemented")
	return nil, errors.New("no")
}

func (*graphicsDriverCreatorImpl) newPlayStation5() (graphicsdriver.Graphics, error) {
	panic("newPlayStation5 not implemented")
	return nil, errors.New("no")
}

var (
	stringNone        = js.ValueOf("none")
	stringTransparent = js.ValueOf("transparent")
)

func driverCursorShapeToCSSCursor(cursor CursorShape) string {
	panic("driverCursorShapeToCSSCursor not implemented")
	return ""
}

type userInterfaceImpl struct {
	graphicsDriver graphicsdriver.Graphics

	runnableOnUnfocused bool
	fpsMode             FPSModeType
	renderingScheduled  bool
	cursorMode          CursorMode
	cursorPrevMode      CursorMode
	captureCursorLater  bool
	cursorShape         CursorShape
	onceUpdateCalled    bool
	lastCaptureExitTime time.Time
	hiDPIEnabled        bool

	context             *context
	inputState          InputState
	cursorXInClient     float64
	cursorYInClient     float64
	origCursorXInClient float64
	origCursorYInClient float64
	touchesInClient     []touchInClient

	savedCursorX              float64
	savedCursorY              float64
	savedOutsideWidth         float64
	savedOutsideHeight        float64
	outsideSizeUnchangedCount int

	keyboardLayoutMap js.Value

	m         sync.Mutex
	dropFileM sync.Mutex
}

var (
	window                = js.Global().Get("window")
	document              = js.Global().Get("document")
	screen                = js.Global().Get("screen")
	canvas                js.Value
	requestAnimationFrame = js.Global().Get("requestAnimationFrame")
	setTimeout            = js.Global().Get("setTimeout")
)

var (
	documentHasFocus = document.Get("hasFocus").Call("bind", document)
	documentHidden   = js.Global().Get("Object").Call("getOwnPropertyDescriptor", js.Global().Get("Document").Get("prototype"), "hidden").Get("get").Call("bind", document)
)

func (u *UserInterface) SetFullscreen(fullscreen bool) {
	panic("SetFullscreen not implemented")
}

func (u *UserInterface) IsFullscreen() bool {
	panic("IsFullscreen not implemented")
	return false
}

func (u *UserInterface) IsFocused() bool {
	panic("IsFocused not implemented")
	return false
}

func (u *UserInterface) SetRunnableOnUnfocused(runnableOnUnfocused bool) {
	panic("SetRunnableOnUnfocused not implemented")
}

func (u *UserInterface) IsRunnableOnUnfocused() bool {
	panic("IsRunnableOnUnfocused not implemented")
	return false
}

func (u *UserInterface) FPSMode() FPSModeType {
	panic("FPSMode not implemented")
	return FPSModeType(0)
}

func (u *UserInterface) SetFPSMode(mode FPSModeType) {
	panic("SetFPSMode not implemented")
}

func (u *UserInterface) ScheduleFrame() {
	panic("ScheduleFrame not implemented")
}

func (u *UserInterface) CursorMode() CursorMode {
	panic("CursorMode not implemented")
	return CursorModeHidden
}

func (u *UserInterface) SetCursorMode(mode CursorMode) {
	panic("SetCursorMode not implemented")
}

func (u *UserInterface) setCursorMode(mode CursorMode) {
	panic("setCursorMode not implemented")
}

func (u *UserInterface) recoverCursorMode() {
	panic("recoverCursorMode not implemented")
}

func (u *UserInterface) CursorShape() CursorShape {
	panic("CursorShape not implemented")
	return CursorShapeDefault
}

func (u *UserInterface) SetCursorShape(shape CursorShape) {
	panic("SetCursorShape not implemented")
}

func (u *UserInterface) outsideSize() (float64, float64) {
	panic("outsideSize not implemented")
	return 0.0, 0.0
}

func (u *UserInterface) suspended() bool {
	panic("suspended not implemented")
	return false
}

func (u *UserInterface) isFocused() bool {
	panic("isFocused not implemented")
	return false
}

func (u *UserInterface) canCaptureCursor() bool {
	panic("canCaptureCursor not implemented")
	return false
}

func (u *UserInterface) update() error {
	panic("update not implemented")
	return errors.New("no")
}

func (u *UserInterface) updateImpl(force bool) error {
	panic("updateImpl not implemented")
	return errors.New("no")
}

func (u *UserInterface) needsUpdate() bool {
	panic("needsUpdate not implemented")
	return false
}

func (u *UserInterface) loopGame() error {
	panic("loopGame not implemented")
	return errors.New("no")
}

func (u *UserInterface) init() error {
	panic("init not implemented")
	return errors.New("no")
}

func (u *UserInterface) setWindowEventHandlers(v js.Value) {
	panic("setWindowEventHandlers not implemented")
}

func (u *UserInterface) setCanvasEventHandlers(v js.Value) {
	panic("setCanvasEventHandlers not implemented")
}

func (u *UserInterface) appendDroppedFiles(data js.Value) {
	panic("appendDroppedFiles not implemented")
}

func (u *UserInterface) forceUpdateOnMinimumFPSMode() {
	panic("forceUpdateOnMinimumFPSMode not implemented")
}

func (u *UserInterface) shouldFocusFirst(options *RunOptions) bool {
	panic("shouldFocusFirst not implemented")
	return false
}

func (u *UserInterface) initOnMainThread(options *RunOptions) error {
	panic("initOnMainThread not implemented")
	return errors.New("no")
}

func (u *UserInterface) updateScreenSize() {
	panic("updateScreenSize not implemented")
}

func (u *UserInterface) readInputState(inputState *InputState) {
	panic("readInputState not implemented")
}

func (u *UserInterface) Window() Window {
	panic("Window not implemented")
	return nil
}

type Monitor struct {
	deviceScaleFactor float64
}

var theMonitor = &Monitor{}

func (m *Monitor) Name() string {
	panic("Name not implemented")
	return ""
}

func (m *Monitor) DeviceScaleFactor() float64 {
	panic("DeviceScaleFactor not implemented")
	return 0.0
}

func (m *Monitor) Size() (int, int) {
	panic("Size not implemented")
	return 0, 0
}

func (u *UserInterface) AppendMonitors(mons []*Monitor) []*Monitor {
	panic("AppendMonitors not implemented")
	return nil
}

func (u *UserInterface) Monitor() *Monitor {
	panic("Monitor not implemented")
	return nil
}

func (u *UserInterface) updateIconIfNeeded() error {
	panic("updateIconIfNeeded not implemented")
	return errors.New("no")
}

func IsScreenTransparentAvailable() bool {
	panic("IsScreenTransparentAvailable not implemented")
	return false
}

func dipToNativePixels(x float64, scale float64) float64 {
	panic("dipToNativePixels not implemented")
	return 0.0
}
