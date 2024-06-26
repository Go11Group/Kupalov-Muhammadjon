// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__skip_one = 144
)

const (
    _stack__skip_one = 160
)

const (
    _size__skip_one = 10040
)

var (
    _pcsp__skip_one = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {9936, 160},
        {9940, 48},
        {9941, 40},
        {9943, 32},
        {9945, 24},
        {9947, 16},
        {9949, 8},
        {9950, 0},
        {10040, 160},
    }
)

var _cfunc_skip_one = []loader.CFunc{
    {"_skip_one_entry", 0,  _entry__skip_one, 0, nil},
    {"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
}
