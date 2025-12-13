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
