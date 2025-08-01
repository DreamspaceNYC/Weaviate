//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

//go:build !noasm && amd64

// AUTO-GENERATED BY GOAT -- DO NOT EDIT

package asm

import "unsafe"

//go:noescape
func popcnt_AVX2_lookup(vec, low_mask_vec, lookup_vec unsafe.Pointer)

//go:noescape
func popcnt_64bit(src, popcnt_constants unsafe.Pointer)

//go:noescape
func hamming_bitwise_256(a, b, res, len, lookup_avx, popcnt_constants unsafe.Pointer)
