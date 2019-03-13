package gls

import (
	"reflect"
	"runtime"
)


type flagFunc func(rem uint32, cb func())

var fs []flagFunc
func initFlagFuncs() {
	fs = []flagFunc{
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 00
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 01
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 02
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 03
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 04
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 05
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 06
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 07
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 08
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 09
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 0a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 0b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 0c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 0d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 0e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 0f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 10
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 11
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 12
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 13
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 14
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 15
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 16
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 17
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 18
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 19
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 1a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 1b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 1c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 1d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 1e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 1f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 20
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 21
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 22
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 23
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 24
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 25
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 26
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 27
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 28
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 29
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 2a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 2b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 2c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 2d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 2e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 2f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 30
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 31
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 32
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 33
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 34
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 35
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 36
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 37
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 38
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 39
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 3a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 3b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 3c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 3d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 3e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 3f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 40
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 41
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 42
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 43
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 44
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 45
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 46
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 47
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 48
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 49
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 4a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 4b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 4c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 4d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 4e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 4f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 50
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 51
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 52
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 53
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 54
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 55
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 56
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 57
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 58
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 59
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 5a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 5b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 5c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 5d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 5e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 5f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 60
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 61
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 62
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 63
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 64
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 65
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 66
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 67
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 68
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 69
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 6a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 6b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 6c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 6d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 6e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 6f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 70
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 71
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 72
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 73
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 74
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 75
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 76
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 77
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 78
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 79
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 7a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 7b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 7c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 7d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 7e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 7f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 80
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 81
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 82
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 83
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 84
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 85
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 86
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 87
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 88
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 89
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 8a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 8b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 8c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 8d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 8e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 8f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 90
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 91
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 92
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 93
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 94
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 95
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 96
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 97
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 98
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 99
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 9a
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 9b
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 9c
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 9d
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 9e
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // 9f
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a0
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a1
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a2
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a3
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a4
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a5
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a6
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a7
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a8
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // a9
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // aa
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ab
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ac
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ad
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ae
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // af
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b0
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b1
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b2
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b3
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b4
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b5
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b6
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b7
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b8
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // b9
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ba
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // bb
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // bc
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // bd
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // be
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // bf
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c0
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c1
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c2
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c3
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c4
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c5
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c6
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c7
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c8
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // c9
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ca
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // cb
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // cc
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // cd
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ce
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // cf
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d0
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d1
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d2
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d3
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d4
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d5
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d6
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d7
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d8
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // d9
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // da
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // db
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // dc
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // dd
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // de
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // df
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e0
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e1
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e2
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e3
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e4
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e5
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e6
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e7
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e8
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // e9
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ea
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // eb
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ec
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ed
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ee
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ef
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f0
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f1
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f2
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f3
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f4
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f5
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f6
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f7
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f8
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // f9
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // fa
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // fb
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // fc
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // fd
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // fe
		func(rem uint32, cb func()) { if rem == 0 { cb() } else { fs[rem & 0xff](rem >> 8, cb) } }, // ff
	}
}
var startPc uintptr
var pcToN = make(map[uintptr]uint32, 256)

func start(rem uint32, cb func()){
	if rem == 0 {
		cb()
	} else {
		fs[rem & 0xff](rem >> 8, cb)
	}
}

func init() {
	initFlagFuncs()
	for i := uint32(0); i < 256; i++ {
		pc := reflect.ValueOf(fs[i]).Pointer()
		pcToN[pc] = i
	}
	startPc = reflect.ValueOf(start).Pointer()
}

func getContextId() uint32 {
	var ret uint32 = 0
	for i := 1;; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if ! ok {
			break
		}
		pcFunc := runtime.FuncForPC(pc)
		if pcFunc == nil {
			continue
		}
		fpc := pcFunc.Entry()
		n, ok := pcToN[fpc]
		if ok {
			ret <<= 8
			ret+= n
		}
		if fpc == startPc {
			break
		}
	}
	return ret
}


