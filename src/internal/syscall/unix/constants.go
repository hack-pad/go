// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm)

package unix

const (
	R_OK = 0x4
	W_OK = 0x2
	X_OK = 0x1
)
