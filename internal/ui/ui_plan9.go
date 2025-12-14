// Copyright 2022 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	"sync/atomic"

	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver"
)

type graphicsDriverCreatorImpl struct {
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

func (g *graphicsDriverCreatorImpl) newMetal() (graphicsdriver.Graphics, error) {
	panic("newMetal not implemented")
	return nil, errors.New("no")
}

func (*graphicsDriverCreatorImpl) newPlayStation5() (graphicsdriver.Graphics, error) {
	panic("newPlayStation5 not implemented")
	return nil, errors.New("no")
}

func (u *UserInterface) SetUIView(uiview uintptr) error {
	panic("SetUIView not implemented")
	return errors.New("no")
}

func (u *UserInterface) IsGL() (bool, error) {
	panic("IsGL not implemented")
	return false, errors.New("no")
}

func dipToNativePixels(x float64, scale float64) float64 {
	panic("dipToNativePixels not implemented")
	return 0.0
}

func dipFromNativePixels(x float64, scale float64) float64 {
	panic("dipFromNativePixels not implemented")
	return 0.0
}

func (u *UserInterface) displayInfo() (int, int, float64, bool) {
	panic("displayInfo not implemented")
	return 0, 0, 0.0, false
}

// Copyright 2016 Hajime Hoshi
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

var (
	// renderCh receives when updating starts.
	renderCh = make(chan struct{})

	// renderEndCh receives when updating finishes.
	renderEndCh = make(chan struct{})
)

func (u *UserInterface) init() error {
	panic("init not implemented")
	return errors.New("no")
}

// Update is called from mobile/ebitenmobileview.
//
// Update must be called on the rendering thread.
func (u *UserInterface) Update() error {
	panic("Update not implemented")
	return errors.New("no")
}

type userInterfaceImpl struct {
	graphicsDriver        graphicsdriver.Graphics
	graphicsLibraryInitCh chan struct{}

	outsideWidth  float64
	outsideHeight float64

	foreground atomic.Bool
	errCh      chan error

	context *context

	inputState InputState
	touches    []TouchForInput

	fpsMode  atomic.Int32
	renderer Renderer

	// uiView is used only on iOS.
	uiView atomic.Uintptr

	m sync.RWMutex
}

func (u *UserInterface) SetForeground(foreground bool) error {
	panic("SetForeground not implemented")
	return errors.New("no")
}

/*
	func (u *UserInterface) Run(game Game, options *RunOptions) error {
		panic("Run not implemented")
		return errors.New("no")
	}
*/
func (u *UserInterface) RunWithoutMainLoop(game Game, options *RunOptions) {
	panic("RunWithoutMainLoop not implemented")
}

func (u *UserInterface) runMobile(game Game, options *RunOptions) (err error) {
	panic("runMobile not implemented")
	return errors.New("no")
}

// outsideSize must be called on the same goroutine as update().
func (u *UserInterface) outsideSize() (float64, float64) {
	panic("outsideSize not implemented")
	return 0.0, 0.0
}

func (u *UserInterface) update() error {
	panic("update not implemented")
	return errors.New("no")
}

// SetOutsideSize is called from mobile/ebitenmobileview.
//
// SetOutsideSize is concurrent safe.
func (u *UserInterface) SetOutsideSize(outsideWidth, outsideHeight float64) {
	panic("SetOutsideSize not implemented")
}

func (u *UserInterface) CursorMode() CursorMode {
	panic("CursorMode not implemented")
	return CursorModeHidden
}

func (u *UserInterface) SetCursorMode(mode CursorMode) {
	panic("SetCursorMode not implemented")
}

func (u *UserInterface) CursorShape() CursorShape {
	panic("CursorShape not implemented")
	return CursorShapeDefault
}

func (u *UserInterface) SetCursorShape(shape CursorShape) {
	panic("SetCursorShape not implemented")
}

func (u *UserInterface) IsFullscreen() bool {
	panic("IsFullscreen not implemented")
	return false
}

func (u *UserInterface) SetFullscreen(fullscreen bool) {
	panic("SetFullscreen not implemented")
}

func (u *UserInterface) IsFocused() bool {
	panic("IsFocused not implemented")
	return false
}

func (u *UserInterface) IsRunnableOnUnfocused() bool {
	panic("IsRunnableOnUnfocused not implemented")
	return false
}

func (u *UserInterface) SetRunnableOnUnfocused(runnableOnUnfocused bool) {
	panic("SetRunnableOnUnfocused not implemented")
}

func (u *UserInterface) FPSMode() FPSModeType {
	panic("FPSMode not implemented")
	return FPSModeType(0)
}

func (u *UserInterface) SetFPSMode(mode FPSModeType) {
	panic("SetFPSMode not implemented")
}

func (u *UserInterface) updateExplicitRenderingModeIfNeeded(fpsMode FPSModeType) {
	panic("updateExplicitRenderingModeIfNeeded not implemented")
}

func (u *UserInterface) readInputState(inputState *InputState) {
	panic("readInputState not implemented")
}

func (u *UserInterface) Window() Window {
	panic("Window not implemented")
	return nil
}

type Monitor struct {
	monitor monitor

	// expireAt is a tick when the cached values expire.
	// As there is no commmon way to detect monitor changes in Android and iOS,
	// the values are invalidated regularly.
	expireAt atomic.Int64

	m sync.Mutex
}

type monitor struct {
	width             int
	height            int
	deviceScaleFactor float64
}

var theMonitor = &Monitor{}

func (m *Monitor) Name() string {
	panic("Name not implemented")
	return ""
}

func (m *Monitor) ensureValues() monitor {
	panic("ensureValues not implemented")
	return monitor{}
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

func (u *UserInterface) UpdateInput(keyPressedTimes, keyReleasedTimes [KeyMax + 1]InputTime, runes []rune, touches []TouchForInput) {
	panic("UpdateInput not implemented")
}

type Renderer interface {
	SetExplicitRenderingMode(explicitRendering bool)
	RequestRenderIfNeeded()
}

func (u *UserInterface) SetRenderer(renderer Renderer) {
	panic("SetRenderer not implemented")
}

func (u *UserInterface) ScheduleFrame() {
	panic("ScheduleFrame not implemented")
}

func (u *UserInterface) updateIconIfNeeded() error {
	panic("updateIconIfNeeded not implemented")
	return errors.New("no")
}

func IsScreenTransparentAvailable() bool {
	panic("IsScreenTransparentAvailable not implemented")
	return false
}
