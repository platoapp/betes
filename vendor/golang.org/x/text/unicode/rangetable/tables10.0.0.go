// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

// +build go1.10

package rangetable

//go:generate go run gen.go --versions=4.1.0,5.1.0,5.2.0,5.0.0,6.1.0,6.2.0,6.3.0,6.0.0,7.0.0,8.0.0,9.0.0,10.0.0

import "unicode"

var assigned = map[string]*unicode.RangeTable{
	"4.1.0":  assigned4_1_0,
	"5.1.0":  assigned5_1_0,
	"5.2.0":  assigned5_2_0,
	"5.0.0":  assigned5_0_0,
	"6.1.0":  assigned6_1_0,
	"6.2.0":  assigned6_2_0,
	"6.3.0":  assigned6_3_0,
	"6.0.0":  assigned6_0_0,
	"7.0.0":  assigned7_0_0,
	"8.0.0":  assigned8_0_0,
	"9.0.0":  assigned9_0_0,
	"10.0.0": assigned10_0_0,
}

// size 2924 bytes (2 KiB)
var assigned4_1_0 = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0000, 0x0241, 1},
		{0x0250, 0x036f, 1},
		{0x0374, 0x0375, 1},
		{0x037a, 0x037e, 4},
		{0x0384, 0x038a, 1},
		{0x038c, 0x038e, 2},
		{0x038f, 0x03a1, 1},
		{0x03a3, 0x03ce, 1},
		{0x03d0, 0x0486, 1},
		{0x0488, 0x04ce, 1},
		{0x04d0, 0x04f9, 1},
		{0x0500, 0x050f, 1},
		{0x0531, 0x0556, 1},
		{0x0559, 0x055f, 1},
		{0x0561, 0x0587, 1},
		{0x0589, 0x058a, 1},
		{0x0591, 0x05b9, 1},
		{0x05bb, 0x05c7, 1},
		{0x05d0, 0x05ea, 1},
		{0x05f0, 0x05f4, 1},
		{0x0600, 0x0603, 1},
		{0x060b, 0x0615, 1},
		{0x061b, 0x061e, 3},
		{0x061f, 0x0621, 2},
		{0x0622, 0x063a, 1},
		{0x0640, 0x065e, 1},
		{0x0660, 0x070d, 1},
		{0x070f, 0x074a, 1},
		{0x074d, 0x076d, 1},
		{0x0780, 0x07b1, 1},
		{0x0901, 0x0939, 1},
		{0x093c, 0x094d, 1},
		{0x0950, 0x0954, 1},
		{0x0958, 0x0970, 1},
		{0x097d, 0x0981, 4},
		{0x0982, 0x0983, 1},
		{0x0985, 0x098c, 1},
		{0x098f, 0x0990, 1},
		{0x0993, 0x09a8, 1},
		{0x09aa, 0x09b0, 1},
		{0x09b2, 0x09b6, 4},
		{0x09b7, 0x09b9, 1},
		{0x09bc, 0x09c4, 1},
		{0x09c7, 0x09c8, 1},
		{0x09cb, 0x09ce, 1},
		{0x09d7, 0x09dc, 5},
		{0x09dd, 0x09df, 2},
		{0x09e0, 0x09e3, 1},
		{0x09e6, 0x09fa, 1},
		{0x0a01, 0x0a03, 1},
		{0x0a05, 0x0a0a, 1},
		{0x0a0f, 0x0a10, 1},
		{0x0a13, 0x0a28, 1},
		{0x0a2a, 0x0a30, 1},
		{0x0a32, 0x0a33, 1},
		{0x0a35, 0x0a36, 1},
		{0x0a38, 0x0a39, 1},
		{0x0a3c, 0x0a3e, 2},
		{0x0a3f, 0x0a42, 1},
		{0x0a47, 0x0a48, 1},
		{0x0a4b, 0x0a4d, 1},
		{0x0a59, 0x0a5c, 1},
		{0x0a5e, 0x0a66, 8},
		{0x0a67, 0x0a74, 1},
		{0x0a81, 0x0a83, 1},
		{0x0a85, 0x0a8d, 1},
		{0x0a8f, 0x0a91, 1},
		{0x0a93, 0x0aa8, 1},
		{0x0aaa, 0x0ab0, 1},
		{0x0ab2, 0x0ab3, 1},
		{0x0ab5, 0x0ab9, 1},
		{0x0abc, 0x0ac5, 1},
		{0x0ac7, 0x0ac9, 1},
		{0x0acb, 0x0acd, 1},
		{0x0ad0, 0x0ae0, 16},
		{0x0ae1, 0x0ae3, 1},
		{0x0ae6, 0x0aef, 1},
		{0x0af1, 0x0b01, 16},
		{0x0b02, 0x0b03, 1},
		{0x0b05, 0x0b0c, 1},
		{0x0b0f, 0x0b10, 1},
		{0x0b13, 0x0b28, 1},
		{0x0b2a, 0x0b30, 1},
		{0x0b32, 0x0b33, 1},
		{0x0b35, 0x0b39, 1},
		{0x0b3c, 0x0b43, 1},
		{0x0b47, 0x0b48, 1},
		{0x0b4b, 0x0b4d, 1},
		{0x0b56, 0x0b57, 1},
		{0x0b5c, 0x0b5d, 1},
		{0x0b5f, 0x0b61, 1},
		{0x0b66, 0x0b71, 1},
		{0x0b82, 0x0b83, 1},
		{0x0b85, 0x0b8a, 1},
		{0x0b8e, 0x0b90, 1},
		{0x0b92, 0x0b95, 1},
		{0x0b99, 0x0b9a, 1},
		{0x0b9c, 0x0b9e, 2},
		{0x0b9f, 0x0ba3, 4},
		{0x0ba4, 0x0ba8, 4},
		{0x0ba9, 0x0baa, 1},
		{0x0bae, 0x0bb9, 1},
		{0x0bbe, 0x0bc2, 1},
		{0x0bc6, 0x0bc8, 1},
		{0x0bca, 0x0bcd, 1},
		{0x0bd7, 0x0be6, 15},
		{0x0be7, 0x0bfa, 1},
		{0x0c01, 0x0c03, 1},
		{0x0c05, 0x0c0c, 1},
		{0x0c0e, 0x0c10, 1},
		{0x0c12, 0x0c28, 1},
		{0x0c2a, 0x0c33, 1},
		{0x0c35, 0x0c39, 1},
		{0x0c3e, 0x0c44, 1},
		{0x0c46, 0x0c48, 1},
		{0x0c4a, 0x0c4d, 1},
		{0x0c55, 0x0c56, 1},
		{0x0c60, 0x0c61, 1},
		{0x0c66, 0x0c6f, 1},
		{0x0c82, 0x0c83, 1},
		{0x0c85, 0x0c8c, 1},
		{0x0c8e, 0x0c90, 1},
		{0x0c92, 0x0ca8, 1},
		{0x0caa, 0x0cb3, 1},
		{0x0cb5, 0x0cb9, 1},
		{0x0cbc, 0x0cc4, 1},
		{0x0cc6, 0x0cc8, 1},
		{0x0cca, 0x0ccd, 1},
		{0x0cd5, 0x0cd6, 1},
		{0x0cde, 0x0ce0, 2},
		{0x0ce1, 0x0ce6, 5},
		{0x0ce7, 0x0cef, 1},
		{0x0d02, 0x0d03, 1},
		{0x0d05, 0x0d0c, 1},
		{0x0d0e, 0x0d10, 1},
		{0x0d12, 0x0d28, 1},
		{0x0d2a, 0x0d39, 1},
		{0x0d3e, 0x0d43, 1},
		{0x0d46, 0x0d48, 1},
		{0x0d4a, 0x0d4d, 1},
		{0x0d57, 0x0d60, 9},
		{0x0d61, 0x0d66, 5},
		{0x0d67, 0x0d6f, 1},
		{0x0d82, 0x0d83, 1},
		{0x0d85, 0x0d96, 1},
		{0x0d9a, 0x0db1, 1},
		{0x0db3, 0x0dbb, 1},
		{0x0dbd, 0x0dc0, 3},
		{0x0dc1, 0x0dc6, 1},
		{0x0dca, 0x0dcf, 5},
		{0x0dd0, 0x0dd4, 1},
		{0x0dd6, 0x0dd8, 2},
		{0x0dd9, 0x0ddf, 1},
		{0x0df2, 0x0df4, 1},
		{0x0e01, 0x0e3a, 1},
		{0x0e3f, 0x0e5b, 1},
		{0x0e81, 0x0e82, 1},
		{0x0e84, 0x0e87, 3},
		{0x0e88, 0x0e8a, 2},
		{0x0e8d, 0x0e94, 7},
		{0x0e95, 0x0e97, 1},
		{0x0e99, 0x0e9f, 1},
		{0x0ea1, 0x0ea3, 1},
		{0x0ea5, 0x0ea7, 2},
		{0x0eaa, 0x0eab, 1},
		{0x0ead, 0x0eb9, 1},
		{0x0ebb, 0x0ebd, 1},
		{0x0ec0, 0x0ec4, 1},
		{0x0ec6, 0x0ec8, 2},
		{0x0ec9, 0x0ecd, 1},
		{0x0ed0, 0x0ed9, 1},
		{0x0edc, 0x0edd, 1},
		{0x0f00, 0x0f47, 1},
		{0x0f49, 0x0f6a, 1},
		{0x0f71, 0x0f8b, 1},
		{0x0f90, 0x0f97, 1},
		{0x0f99, 0x0fbc, 1},
		{0x0fbe, 0x0fcc, 1},
		{0x0fcf, 0x0fd1, 1},
		{0x1000, 0x1021, 1},
		{0x1023, 0x1027, 1},
		{0x1029, 0x102a, 1},
		{0x102c, 0x1032, 1},
		{0x1036, 0x1039, 1},
		{0x1040, 0x1059, 1},
		{0x10a0, 0x10c5, 1},
		{0x10d0, 0x10fc, 1},
		{0x1100, 0x1159, 1},
		{0x115f, 0x11a2, 1},
		{0x11a8, 0x11f9, 1},
		{0x1200, 0x1248, 1},
		{0x124a, 0x124d, 1},
		{0x1250, 0x1256, 1},
		{0x1258, 0x125a, 2},
		{0x125b, 0x125d, 1},
		{0x1260, 0x1288, 1},
		{0x128a, 0x128d, 1},
		{0x1290, 0x12b0, 1},
		{0x12b2, 0x12b5, 1},
		{0x12b8, 0x12be, 1},
		{0x12c0, 0x12c2, 2},
		{0x12c3, 0x12c5, 1},
		{0x12c8, 0x12d6, 1},
		{0x12d8, 0x1310, 1},
		{0x1312, 0x1315, 1},
		{0x1318, 0x135a, 1},
		{0x135f, 0x137c, 1},
		{0x1380, 0x1399, 1},
		{0x13a0, 0x13f4, 1},
		{0x1401, 0x1676, 1},
		{0x1680, 0x169c, 1},
		{0x16a0, 0x16f0, 1},
		{0x1700, 0x170c, 1},
		{0x170e, 0x1714, 1},
		{0x1720, 0x1736, 1},
		{0x1740, 0x1753, 1},
		{0x1760, 0x176c, 1},
		{0x176e, 0x1770, 1},
		{0x1772, 0x1773, 1},
		{0x1780, 0x17dd, 1},
		{0x17e0, 0x17e9, 1},
		{0x17f0, 0x17f9, 1},
		{0x1800, 0x180e, 1},
		{0x1810, 0x1819, 1},
		{0x1820, 0x1877, 1},
		{0x1880, 0x18a9, 1},
		{0x1900, 0x191c, 1},
		{0x1920, 0x192b, 1},
		{0x1930, 0x193b, 1},
		{0x1940, 0x1944, 4},
		{0x1945, 0x196d, 1},
		{0x1970, 0x1974, 1},
		{0x1980, 0x19a9, 1},
		{0x19b0, 0x19c9, 1},
		{0x19d0, 0x19d9, 1},
		{0x19de, 0x1a1b, 1},
		{0x1a1e, 0x1a1f, 1},
		{0x1d00, 0x1dc3, 1},
		{0x1e00, 0x1e9b, 1},
		{0x1ea0, 0x1ef9, 1},
		{0x1f00, 0x1f15, 1},
		{0x1f18, 0x1f1d, 1},
		{0x1f20, 0x1f45, 1},
		{0x1f48, 0x1f4d, 1},
		{0x1f50, 0x1f57, 1},
		{0x1f59, 0x1f5f, 2},
		{0x1f60, 0x1f7d, 1},
		{0x1f80, 0x1fb4, 1},
		{0x1fb6, 0x1fc4, 1},
		{0x1fc6, 0x1fd3, 1},
		{0x1fd6, 0x1fdb, 1},
		{0x1fdd, 0x1fef, 1},
		{0x1ff2, 0x1ff4, 1},
		{0x1ff6, 0x1ffe, 1},
		{0x2000, 0x2063, 1},
		{0x206a, 0x2071, 1},
		{0x2074, 0x208e, 1},
		{0x2090, 0x2094, 1},
		{0x20a0, 0x20b5, 1},
		{0x20d0, 0x20eb, 1},
		{0x2100, 0x214c, 1},
		{0x2153, 0x2183, 1},
		{0x2190, 0x23db, 1},
		{0x2400, 0x2426, 1},
		{0x2440, 0x244a, 1},
		{0x2460, 0x269c, 1},
		{0x26a0, 0x26b1, 1},
		{0x2701, 0x2704, 1},
		{0x2706, 0x2709, 1},
		{0x270c, 0x2727, 1},
		{0x2729, 0x274b, 1},
		{0x274d, 0x274f, 2},
		{0x2750, 0x2752, 1},
		{0x2756, 0x2758, 2},
		{0x2759, 0x275e, 1},
		{0x2761, 0x2794, 1},
		{0x2798, 0x27af, 1},
		{0x27b1, 0x27be, 1},
		{0x27c0, 0x27c6, 1},
		{0x27d0, 0x27eb, 1},
		{0x27f0, 0x2b13, 1},
		{0x2c00, 0x2c2e, 1},
		{0x2c30, 0x2c5e, 1},
		{0x2c80, 0x2cea, 1},
		{0x2cf9, 0x2d25, 1},
		{0x2d30, 0x2d65, 1},
		{0x2d6f, 0x2d80, 17},
		{0x2d81, 0x2d96, 1},
		{0x2da0, 0x2da6, 1},
		{0x2da8, 0x2dae, 1},
		{0x2db0, 0x2db6, 1},
		{0x2db8, 0x2dbe, 1},
		{0x2dc0, 0x2dc6, 1},
		{0x2dc8, 0x2dce, 1},
		{0x2dd0, 0x2dd6, 1},
		{0x2dd8, 0x2dde, 1},
		{0x2e00, 0x2e17, 1},
		{0x2e1c, 0x2e1d, 1},
		{0x2e80, 0x2e99, 1},
		{0x2e9b, 0x2ef3, 1},
		{0x2f00, 0x2fd5, 1},
		{0x2ff0, 0x2ffb, 1},
		{0x3000, 0x303f, 1},
		{0x3041, 0x3096, 1},
		{0x3099, 0x30ff, 1},
		{0x3105, 0x312c, 1},
		{0x3131, 0x318e, 1},
		{0x3190, 0x31b7, 1},
		{0x31c0, 0x31cf, 1},
		{0x31f0, 0x321e, 1},
		{0x3220, 0x3243, 1},
		{0x3250, 0x32fe, 1},
		{0x3300, 0x4db5, 1},
		{0x4dc0, 0x9fbb, 1},
		{0xa000, 0xa48c, 1},
		{0xa490, 0xa4c6, 1},
		{0xa700, 0xa716, 1},
		{0xa800, 0xa82b, 1},
		{0xac00, 0xd7a3, 1},
		{0xd800, 0xfa2d, 1},
		{0xfa30, 0xfa6a, 1},
		{0xfa70, 0xfad9, 1},
		{0xfb00, 0xfb06, 1},
		{0xfb13, 0xfb17, 1},
		{0xfb1d, 0xfb36, 1},
		{0xfb38, 0xfb3c, 1},
		{0xfb3e, 0xfb40, 2},
		{0xfb41, 0xfb43, 2},
		{0xfb44, 0xfb46, 2},
		{0xfb47, 0xfbb1, 1},
		{0xfbd3, 0xfd3f, 1},
		{0xfd50, 0xfd8f, 1},
		{0xfd92, 0xfdc7, 1},
		{0xfdf0, 0xfdfd, 1},
		{0xfe00, 0xfe19, 1},
		{0xfe20, 0xfe23, 1},
		{0xfe30, 0xfe52, 1},
		{0xfe54, 0xfe66, 1},
		{0xfe68, 0xfe6b, 1},
		{0xfe70, 0xfe74, 1},
		{0xfe76, 0xfefc, 1},
		{0xfeff, 0xff01, 2},
		{0xff02, 0xffbe, 1},
		{0xffc2, 0xffc7, 1},
		{0xffca, 0xffcf, 1},
		{0xffd2, 0xffd7, 1},
		{0xffda, 0xffdc, 1},
		{0xffe0, 0xffe6, 1},
		{0xffe8, 0xffee, 1},
		{0xfff9, 0xfffd, 1},
	},
	R32: []unicode.Range32{
		{0x00010000, 0x0001000b, 1},
		{0x0001000d, 0x00010026, 1},
		{0x00010028, 0x0001003a, 1},
		{0x0001003c, 0x0001003d, 1},
		{0x0001003f, 0x0001004d, 1},
		{0x00010050, 0x0001005d, 1},
		{0x00010080, 0x000100fa, 1},
		{0x00010100, 0x00010102, 1},
		{0x00010107, 0x00010133, 1},
		{0x00010137, 0x0001018a, 1},
		{0x00010300, 0x0001031e, 1},
		{0x00010320, 0x00010323, 1},
		{0x00010330, 0x0001034a, 1},
		{0x00010380, 0x0001039d, 1},
		{0x0001039f, 0x000103c3, 1},
		{0x000103c8, 0x000103d5, 1},
		{0x00010400, 0x0001049d, 1},
		{0x000104a0, 0x000104a9, 1},
		{0x00010800, 0x00010805, 1},
		{0x00010808, 0x0001080a, 2},
		{0x0001080b, 0x00010835, 1},
		{0x00010837, 0x00010838, 1},
		{0x0001083c, 0x0001083f, 3},
		{0x00010a00, 0x00010a03, 1},
		{0x00010a05, 0x00010a06, 1},
		{0x00010a0c, 0x00010a13, 1},
		{0x00010a15, 0x00010a17, 1},
		{0x00010a19, 0x00010a33, 1},
		{0x00010a38, 0x00010a3a, 1},
		{0x00010a3f, 0x00010a47, 1},
		{0x00010a50, 0x00010a58, 1},
		{0x0001d000, 0x0001d0f5, 1},
		{0x0001d100, 0x0001d126, 1},
		{0x0001d12a, 0x0001d1dd, 1},
		{0x0001d200, 0x0001d245, 1},
		{0x0001d300, 0x0001d356, 1},
		{0x0001d400, 0x0001d454, 1},
		{0x0001d456, 0x0001d49c, 1},
		{0x0001d49e, 0x0001d49f, 1},
		{0x0001d4a2, 0x0001d4a5, 3},
		{0x0001d4a6, 0x0001d4a9, 3},
		{0x0001d4aa, 0x0001d4ac, 1},
		{0x0001d4ae, 0x0001d4b9, 1},
		{0x0001d4bb, 0x0001d4bd, 2},
		{0x0001d4be, 0x0001d4c3, 1},
		{0x0001d4c5, 0x0001d505, 1},
		{0x0001d507, 0x0001d50a, 1},
		{0x0001d50d, 0x0001d514, 1},
		{0x0001d516, 0x0001d51c, 1},
		{0x0001d51e, 0x0001d539, 1},
		{0x0001d53b, 0x0001d53e, 1},
		{0x0001d540, 0x0001d544, 1},
		{0x0001d546, 0x0001d54a, 4},
		{0x0001d54b, 0x0001d550, 1},
		{0x0001d552, 0x0001d6a5, 1},
		{0x0001d6a8, 0x0001d7c9, 1},
		{0x0001d7ce, 0x0001d7ff, 1},
		{0x00020000, 0x0002a6d6, 1},
		{0x0002f800, 0x0002fa1d, 1},
		{0x000e0001, 0x000e0020, 31},
		{0x000e0021, 0x000e007f, 1},
		{0x000e0100, 0x000e01ef, 1},
		{0x000f0000, 0x000ffffd, 1},
		{0x00100000, 0x0010fffd, 1},
	},
	LatinOffset: 0,
}

// size 3152 bytes (3 KiB)
var assigned5_1_0 = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0000, 0x0377, 1},
		{0x037a, 0x037e, 1},
		{0x0384, 0x038a, 1},
		{0x038c, 0x038e, 2},
		{0x038f, 0x03a1, 1},
		{0x03a3, 0x0523, 1},
		{0x0531, 0x0556, 1},
		{0x0559, 0x055f, 1},
		{0x0561, 0x0587, 1},
		{0x0589, 0x058a, 1},
		{0x0591, 0x05c7, 1},
		{0x05d0, 0x05ea, 1},
		{0x05f0, 0x05f4, 1},
		{0x0600, 0x0603, 1},
		{0x0606, 0x061b, 1},
		{0x061e, 0x061f, 1},
		{0x0621, 0x065e, 1},
		{0x0660, 0x070d, 1},
		{0x070f, 0x074a, 1},
		{0x074d, 0x07b1, 1},
		{0x07c0, 0x07fa, 1},
		{0x0901, 0x0939, 1},
		{0x093c, 0x094d, 1},
		{0x0950, 0x0954, 1},
		{0x0958, 0x0972, 1},
		{0x097b, 0x097f, 1},
		{0x0981, 0x0983, 1},
		{0x0985, 0x098c, 1},
		{0x098f, 0x0990, 1},
		{0x0993, 0x09a8, 1},
		{0x09aa, 0x09b0, 1},
		{0x09b2, 0x09b6, 4},
		{0x09b7, 0x09b9, 1},
		{0x09bc, 0x09c4, 1},
		{0x09c7, 0x09c8, 1},
		{0x09cb, 0x09ce, 1},
		{0x09d7, 0x09dc, 5},
		{0x09dd, 0x09df, 2},
		{0x09e0, 0x09e3, 1},
		{0x09e6, 0x09fa, 1},
		{0x0a01, 0x0a03, 1},
		{0x0a05, 0x0a0a, 1},
		{0x0a0f, 0x0a10, 1},
		{0x0a13, 0x0a28, 1},
		{0x0a2a, 0x0a30, 1},
		{0x0a32, 0x0a33, 1},
		{0x0a35, 0x0a36, 1},
		{0x0a38, 0x0a39, 1},
		{0x0a3c, 0x0a3e, 2},
		{0x0a3f, 0x0a42, 1},
		{0x0a47, 0x0a48, 1},
		{0x0a4b, 0x0a4d, 1},
		{0x0a51, 0x0a59, 8},
		{0x0a5a, 0x0a5c, 1},
		{0x0a5e, 0x0a66, 8},
		{0x0a67, 0x0a75, 1},
		{0x0a81, 0x0a83, 1},
		{0x0a85, 0x0a8d, 1},
		{0x0a8f, 0x0a91, 1},
		{0x0a93, 0x0aa8, 1},
		{0x0aaa, 0x0ab0, 1},
		{0x0ab2, 0x0ab3, 1},
		{0x0ab5, 0x0ab9, 1},
		{0x0abc, 0x0ac5, 1},
		{0x0ac7, 0x0ac9, 1},
		{0x0acb, 0x0acd, 1},
		{0x0ad0, 0x0ae0, 16},
		{0x0ae1, 0x0ae3, 1},
		{0x0ae6, 0x0aef, 1},
		{0x0af1, 0x0b01, 16},
		{0x0b02, 0x0b03, 1},
		{0x0b05, 0x0b0c, 1},
		{0x0b0f, 0x0b10, 1},
		{0x0b13, 0x0b28, 1},
		{0x0b2a, 0x0b30, 1},
		{0x0b32, 0x0b33, 1},
		{0x0b35, 0x0b39, 1},
		{0x0b3c, 0x0b44, 1},
		{0x0b47, 0x0b48, 1},
		{0x0b4b, 0x0b4d, 1},
		{0x0b56, 0x0b57, 1},
		{0x0b5c, 0x0b5d, 1},
		{0x0b5f, 0x0b63, 1},
		{0x0b66, 0x0b71, 1},
		{0x0b82, 0x0b83, 1},
		{0x0b85, 0x0b8a, 1},
		{0x0b8e, 0x0b90, 1},
		{0x0b92, 0x0b95, 1},
		{0x0b99, 0x0b9a, 1},
		{0x0b9c, 0x0b9e, 2},
		{0x0b9f, 0x0ba3, 4},
		{0x0ba4, 0x0ba8, 4},
		{0x0ba9, 0x0baa, 1},
		{0x0bae, 0x0bb9, 1},
		{0x0bbe, 0x0bc2, 1},
		{0x0bc6, 0x0bc8, 1},
		{0x0bca, 0x0bcd, 1},
		{0x0bd0, 0x0bd7, 7},
		{0x0be6, 0x0bfa, 1},
		{0x0c01, 0x0c03, 1},
		{0x0c05, 0x0c0c, 1},
		{0x0c0e, 0x0c10, 1},
		{0x0c12, 0x0c28, 1},
		{0x0c2a, 0x0c33, 1},
		{0x0c35, 0x0c39, 1},
		{0x0c3d, 0x0c44, 1},
		{0x0c46, 0x0c48, 1},
		{0x0c4a, 0x0c4d, 1},
		{0x0c55, 0x0c56, 1},
		{0x0c58, 0x0c59, 1},
		{0x0c60, 0x0c63, 1},
		{0x0c66, 0x0c6f, 1},
		{0x0c78, 0x0c7f, 1},
		{0x0c82, 0x0c83, 1},
		{0x0c85, 0x0c8c, 1},
		{0x0c8e, 0x0c90, 1},
		{0x0c92, 0x0ca8, 1},
		{0x0caa, 0x0cb3, 1},
		{0x0cb5, 0x0cb9, 1},
		{0x0cbc, 0x0cc4, 1},
		{0x0cc6, 0x0cc8, 1},
		{0x0cca, 0x0ccd, 1},
		{0x0cd5, 0x0cd6, 1},
		{0x0cde, 0x0ce0, 2},
		{0x0ce1, 0x0ce3, 1},
		{0x0ce6, 0x0cef, 1},
		{0x0cf1, 0x0cf2, 1},
		{0x0d02, 0x0d03, 1},
		{0x0d05, 0x0d0c, 1},
		{0x0d0e, 0x0d10, 1},
		{0x0d12, 0x0d28, 1},
		{0x0d2a, 0x0d39, 1},
		{0x0d3d, 0x0d44, 1},
		{0x0d46, 0x0d48, 1},
		{0x0d4a, 0x0d4d, 1},
		{0x0d57, 0x0d60, 9},
		{0x0d61, 0x0d63, 1},
		{0x0d66, 0x0d75, 1},
		{0x0d79, 0x0d7f, 1},
		{0x0d82, 0x0d83, 1},
		{0x0d85, 0x0d96, 1},
		{0x0d9a, 0x0db1, 1},
		{0x0db3, 0x0dbb, 1},
		{0x0dbd, 0x0dc0, 3},
		{0x0dc1, 0x0dc6, 1},
		{0x0dca, 0x0dcf, 5},
		{0x0dd0, 0x0dd4, 1},
		{0x0dd6, 0x0dd8, 2},
		{0x0dd9, 0x0ddf, 1},
		{0x0df2, 0x0df4, 1},
		{0x0e01, 0x0e3a, 1},
		{0x0e3f, 0x0e5b, 1},
		{0x0e81, 0x0e82, 1},
		{0x0e84, 0x0e87, 3},
		{0x0e88, 0x0e8a, 2},
		{0x0e8d, 0x0e94, 7},
		{0x0e95, 0x0e97, 1},
		{0x0e99, 0x0e9f, 1},
		{0x0ea1, 0x0ea3, 1},
		{0x0ea5, 0x0ea7, 2},
		{0x0eaa, 0x0eab, 1},
		{0x0ead, 0x0eb9, 1},
		{0x0ebb, 0x0ebd, 1},
		{0x0ec0, 0x0ec4, 1},
		{0x0ec6, 0x0ec8, 2},
		{0x0ec9, 0x0ecd, 1},
		{0x0ed0, 0x0ed9, 1},
		{0x0edc, 0x0edd, 1},
		{0x0f00, 0x0f47, 1},
		{0x0f49, 0x0f6c, 1},
		{0x0f71, 0x0f8b, 1},
		{0x0f90, 0x0f97, 1},
		{0x0f99, 0x0fbc, 1},
		{0x0fbe, 0x0fcc, 1},
		{0x0fce, 0x0fd4, 1},
		{0x1000, 0x1099, 1},
		{0x109e, 0x10c5, 1},
		{0x10d0, 0x10fc, 1},
		{0x1100, 0x1159, 1},
		{0x115f, 0x11a2, 1},
		{0x11a8, 0x11f9, 1},
		{0x1200, 0x1248, 1},
		{0x124a, 0x124d, 1},
		{0x1250, 0x1256, 1},
		{0x1258, 0x125a, 2},
		{0x125b, 0x125d, 1},
		{0x1260, 0x1288, 1},
		{0x128a, 0x128d, 1},
		{0x1290, 0x12b0, 1},
		{0x12b2, 0x12b5, 1},
		{0x12b8, 0x12be, 1},
		{0x12c0, 0x12c2, 2},
		{0x12c3, 0x12c5, 1},
		{0x12c8, 0x12d6, 1},
		{0x12d8, 0x1310, 1},
		{0x1312, 0x1315, 1},
		{0x1318, 0x135a, 1},
		{0x135f, 0x137c, 1},
		{0x1380, 0x1399, 1},
		{0x13a0, 0x13f4, 1},
		{0x1401, 0x1676, 1},
		{0x1680, 0x169c, 1},
		{0x16a0, 0x16f0, 1},
		{0x1700, 0x170c, 1},
		{0x170e, 0x1714, 1},
		{0x1720, 0x1736, 1},
		{0x1740, 0x1753, 1},
		{0x1760, 0x176c, 1},
		{0x176e, 0x1770, 1},
		{0x1772, 0x1773, 1},
		{0x1780, 0x17dd, 1},
		{0x17e0, 0x17e9, 1},
		{0x17f0, 0x17f9, 1},
		{0x1800, 0x180e, 1},
		{0x1810, 0x1819, 1},
		{0x1820, 0x1877, 1},
		{0x1880, 0x18aa, 1},
		{0x1900, 0x191c, 1},
		{0x1920, 0x192b, 1},
		{0x1930, 0x193b, 1},
		{0x1940, 0x1944, 4},
		{0x1945, 0x196d, 1},
		{0x1970, 0x1974, 1},
		{0x1980, 0x19a9, 1},
		{0x19b0, 0x19c9, 1},
		{0x19d0, 0x19d9, 1},
		{0x19de, 0x1a1b, 1},
		{0x1a1e, 0x1a1f, 1},
		{0x1b00, 0x1b4b, 1},
		{0x1b50, 0x1b7c, 1},
		{0x1b80, 0x1baa, 1},
		{0x1bae, 0x1bb9, 1},
		{0x1c00, 0x1c37, 1},
		{0x1c3b, 0x1c49, 1},
		{0x1c4d, 0x1c7f, 1},
		{0x1d00, 0x1de6, 1},
		{0x1dfe, 0x1f15, 1},
		{0x1f18, 0x1f1d, 1},
		{0x1f20, 0x1f45, 1},
		{0x1f48, 0x1f4d, 1},
		{0x1f50, 0x1f57, 1},
		{0x1f59, 0x1f5f, 2},
		{0x1f60, 0x1f7d, 1},
		{0x1f80, 0x1fb4, 1},
		{0x1fb6, 0x1fc4, 1},
		{0x1fc6, 0x1fd3, 1},
		{0x1fd6, 0x1fdb, 1},
		{0x1fdd, 0x1fef, 1},
		{0x1ff2, 0x1ff4, 1},
		{0x1ff6, 0x1ffe, 1},
		{0x2000, 0x2064, 1},
		{0x206a, 0x2071, 1},
		{0x2074, 0x208e, 1},
		{0x2090, 0x2094, 1},
		{0x20a0, 0x20b5, 1},
		{0x20d0, 0x20f0, 1},
		{0x2100, 0x214f, 1},
		{0x2153, 0x2188, 1},
		{0x2190, 0x23e7, 1},
		{0x2400, 0x2426, 1},
		{0x2440, 0x244a, 1},
		{0x2460, 0x269d, 1},
		{0x26a0, 0x26bc, 1},
		{0x26c0, 0x26c3, 1},
		{0x2701, 0x2704, 1},
		{0x2706, 0x2709, 1},
		{0x270c, 0x2727, 1},
		{0x2729, 0x274b, 1},
		{0x274d, 0x274f, 2},
		{0x2750, 0x2752, 1},
		{0x2756, 0x2758, 2},
		{0x2759, 0x275e, 1},
		{0x2761, 0x2794, 1},
		{0x2798, 0x27af, 1},
		{0x27b1, 0x27be, 1},
		{0x27c0, 0x27ca, 1},
		{0x27cc, 0x27d0, 4},
		{0x27d1, 0x2b4c, 1},
		{0x2b50, 0x2b54, 1},
		{0x2c00, 0x2c2e, 1},
		{0x2c30, 0x2c5e, 1},
		{0x2c60, 0x2c6f, 1},
		{0x2c71, 0x2c7d, 1},
		{0x2c80, 0x2cea, 1},
		{0x2cf9, 0x2d25, 1},
		{0x2d30, 0x2d65, 1},
		{0x2d6f, 0x2d80, 17},
		{0x2d81, 0x2d96, 1},
		{0x2da0, 0x2da6, 1},
		{0x2da8, 0x2dae, 1},
		{0x2db0, 0x2db6, 1},
		{0x2db8, 0x2dbe, 1},
		{0x2dc0, 0x2dc6, 1},
		{0x2dc8, 0x2dce, 1},
		{0x2dd0, 0x2dd6, 1},
		{0x2dd8, 0x2dde, 1},
		{0x2de0, 0x2e30, 1},
		{0x2e80, 0x2e99, 1},
		{0x2e9b, 0x2ef3, 1},
		{0x2f00, 0x2fd5, 1},
		{0x2ff0, 0x2ffb, 1},
		{0x3000, 0x303f, 1},
		{0x3041, 0x3096, 1},
		{0x3099, 0x30ff, 1},
		{0x3105, 0x312d, 1},
		{0x3131, 0x318e, 1},
		{0x3190, 0x31b7, 1},
		{0x31c0, 0x31e3, 1},
		{0x31f0, 0x321e, 1},
		{0x3220, 0x3243, 1},
		{0x3250, 0x32fe, 1},
		{0x3300, 0x4db5, 1},
		{0x4dc0, 0x9fc3, 1},
		{0xa000, 0xa48c, 1},
		{0xa490, 0xa4c6, 1},
		{0xa500, 0xa62b, 1},
		{0xa640, 0xa65f, 1},
		{0xa662, 0xa673, 1},
		{0xa67c, 0xa697, 1},
		{0xa700, 0xa78c, 1},
		{0xa7fb, 0xa82b, 1},
		{0xa840, 0xa877, 1},
		{0xa880, 0xa8c4, 1},
		{0xa8ce, 0xa8d9, 1},
		{0xa900, 0xa953, 1},
		{0xa95f, 0xaa00, 161},
		{0xaa01, 0xaa36, 1},
		{0xaa40, 0xaa4d, 1},
		{0xaa50, 0xaa59, 1},
		{0xaa5c, 0xaa5f, 1},
		{0xac00, 0xd7a3, 1},
		{0xd800, 0xfa2d, 1},
		{0xfa30, 0xfa6a, 1},
		{0xfa70, 0xfad9, 1},
		{0xfb00, 0xfb06, 1},
		{0xfb13, 0xfb17, 1},
		{0xfb1d, 0xfb36, 1},
		{0xfb38, 0xfb3c, 1},
		{0xfb3e, 0xfb40, 2},
		{0xfb41, 0xfb43, 2},
		{0xfb44, 0xfb46, 2},
		{0xfb47, 0xfbb1, 1},
		{0xfbd3, 0xfd3f, 1},
		{0xfd50, 0xfd8f, 1},
		{0xfd92, 0xfdc7, 1},
		{0xfdf0, 0xfdfd, 1},
		{0xfe00, 0xfe19, 1},
		{0xfe20, 0xfe26, 1},
		{0xfe30, 0xfe52, 1},
		{0xfe54, 0xfe66, 1},
		{0xfe68, 0xfe6b, 1},
		{0xfe70, 0xfe74, 1},
		{0xfe76, 0xfefc, 1},
		{0xfeff, 0xff01, 2},
		{0xff02, 0xffbe, 1},
		{0xffc2, 0xffc7, 1},
		{0xffca, 0xffcf, 1},
		{0xffd2, 0xffd7, 1},
		{0xffda, 0xffdc, 1},
		{0xffe0, 0xffe6, 1},
		{0xffe8, 0xffee, 1},
		{0xfff9, 0xfffd, 1},
	},
	R32: []unicode.Range32{
		{0x00010000, 0x0001000b, 1},
		{0x0001000d, 0x00010026, 1},
		{0x00010028, 0x0001003a, 1},
		{0x0001003c, 0x0001003d, 1},
		{0x0001003f, 0x0001004d, 1},
		{0x00010050, 0x0001005d, 1},
		{0x00010080, 0x000100fa, 1},
		{0x00010100, 0x00010102, 1},
		{0x00010107, 0x00010133, 1},
		{0x00010137, 0x0001018a, 1},
		{0x00010190, 0x0001019b, 1},
		{0x000101d0, 0x000101fd, 1},
		{0x00010280, 0x0001029c, 1},
		{0x000102a0, 0x000102d0, 1},
		{0x00010300, 0x0001031e, 1},
		{0x00010320, 0x00010323, 1},
		{0x00010330, 0x0001034a, 1},
		{0x00010380, 0x0001039d, 1},
		{0x0001039f, 0x000103c3, 1},
		{0x000103c8, 0x000103d5, 1},
		{0x00010400, 0x0001049d, 1},
		{0x000104a0, 0x000104a9, 1},
		{0x00010800, 0x00010805, 1},
		{0x00010808, 0x0001080a, 2},
		{0x0001080b, 0x00010835, 1},
		{0x00010837, 0x00010838, 1},
		{0x0001083c, 0x0001083f, 3},
		{0x00010900, 0x00010919, 1},
		{0x0001091f, 0x00010939, 1},
		{0x0001093f, 0x00010a00, 193},
		{0x00010a01, 0x00010a03, 1},
		{0x00010a05, 0x00010a06, 1},
		{0x00010a0c, 0x00010a13, 1},
		{0x00010a15, 0x00010a17, 1},
		{0x00010a19, 0x00010a33, 1},
		{0x00010a38, 0x00010a3a, 1},
		{0x00010a3f, 0x00010a47, 1},
		{0x00010a50, 0x00010a58, 1},
		{0x00012000, 0x0001236e, 1},
		{0x00012400, 0x00012462, 1},
		{0x00012470, 0x00012473, 1},
		{0x0001d000, 0x0001d0f5, 1},
		{0x0001d100, 0x0001d126, 1},
		{0x0001d129, 0x0001d1dd, 1},
		{0x0001d200, 0x0001d245, 1},
		{0x0001d300, 0x0001d356, 1},
		{0x0001d360, 0x0001d371, 1},
		{0x0001d400, 0x0001d454, 1},
		{0x0001d456, 0x0001d49c, 1},
		{0x0001d49e, 0x0001d49f, 1},
		{0x0001d4a2, 0x0001d4a5, 3},
		{0x0001d4a6, 0x0001d4a9, 3},
		{0x0001d4aa, 0x0001d4ac, 1},
		{0x0001d4ae, 0x0001d4b9, 1},
		{0x0001d4bb, 0x0001d4bd, 2},
		{0x0001d4be, 0x0001d4c3, 1},
		{0x0001d4c5, 0x0001d505, 1},
		{0x0001d507, 0x0001d50a, 1},
		{0x0001d50d, 0x0001d514, 1},
		{0x0001d516, 0x0001d51c, 1},
		{0x0001d51e, 0x0001d539, 1},
		{0x0001d53b, 0x0001d53e, 1},
		{0x0001d540, 0x0001d544, 1},
		{0x0001d546, 0x0001d54a, 4},
		{0x0001d54b, 0x0001d550, 1},
		{0x0001d552, 0x0001d6a5, 1},
		{0x0001d6a8, 0x0001d7cb, 1},
		{0x0001d7ce, 0x0001d7ff, 1},
		{0x0001f000, 0x0001f02b, 1},
		{0x0001f030, 0x0001f093, 1},
		{0x00020000, 0x0002a6d6, 1},
		{0x0002f800, 0x0002fa1d, 1},
		{0x000e0001, 0x000e0020, 31},
		{0x000e0021, 0x000e007f, 1},
		{0x000e0100, 0x000e01ef, 1},
		{0x000f0000, 0x000ffffd, 1},
		{0x00100000, 0x0010fffd, 1},
	},
	LatinOffset: 0,
}

// size 3518 bytes (3 KiB)
var assigned5_2_0 = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0000, 0x0377, 1},
		{0x037a, 0x037e, 1},
		{0x0384, 0x038a, 1},
		{0x038c, 0x038e, 2},
		{0x038f, 0x03a1, 1},
		{0x03a3, 0x0525, 1},
		{0x0531, 0x0556, 1},
		{0x0559, 0x055f, 1},
		{0x0561, 0x0587, 1},
		{0x0589, 0x058a, 1},
		{0x0591, 0x05c7, 1},
		{0x05d0, 0x05ea, 1},
		{0x05f0, 0x05f4, 1},
		{0x0600, 0x0603, 1},
		{0x0606, 0x061b, 1},
		{0x061e, 0x061f, 1},
		{0x0621, 0x065e, 1},
		{0x0660, 0x070d, 1},
		{0x070f, 0x074a, 1},
		{0x074d, 0x07b1, 1},
		{0x07c0, 0x07fa, 1},
		{0x0800, 0x082d, 1},
		{0x0830, 0x083e, 1},
		{0x0900, 0x0939, 1},
		{0x093c, 0x094e, 1},
		{0x0950, 0x0955, 1},
		{0x0958, 0x0972, 1},
		{0x0979, 0x097f, 1},
		{0x0981, 0x0983, 1},
		{0x0985, 0x098c, 1},
		{0x098f, 0x0990, 1},
		{0x0993, 0x09a8, 1},
		{0x09aa, 0x09b0, 1},
		{0x09b2, 0x09b6, 4},
		{0x09b7, 0x09b9, 1},
		{0x09bc, 0x09c4, 1},
		{0x09c7, 0x09c8, 1},
		{0x09cb, 0x09ce, 1},
		{0x09d7, 0x09dc, 5},
		{0x09dd, 0x09df, 2},
		{0x09e0, 0x09e3, 1},
		{0x09e6, 0x09fb, 1},
		{0x0a01, 0x0a03, 1},
		{0x0a05, 0x0a0a, 1},
		{0x0a0f, 0x0a10, 1},
		{0x0a13, 0x0a28, 1},
		{0x0a2a, 0x0a30, 1},
		{0x0a32, 0x0a33, 1},
		{0x0a35, 0x0a36, 1},
		{0x0a38, 0x0a39, 1},
		{0x0a3c, 0x0a3e, 2},
		{0x0a3f, 0x0a42, 1},
		{0x0a47, 0x0a48, 1},
		{0x0a4b, 0x0a4d, 1},
		{0x0a51, 0x0a59, 8},
		{0x0a5a, 0x0a5c, 1},
		{0x0a5e, 0x0a66, 8},
		{0x0a67, 0x0a75, 1},
		{0x0a81, 0x0a83, 1},
		{0x0a85, 0x0a8d, 1},
		{0x0a8f, 0x0a91, 1},
		{0x0a93, 0x0aa8, 1},
		{0x0aaa, 0x0ab0, 1},
		{0x0ab2, 0x0ab3, 1},
		{0x0ab5, 0x0ab9, 1},
		{0x0abc, 0x0ac5, 1},
		{0x0ac7, 0x0ac9, 1},
		{0x0acb, 0x0acd, 1},
		{0x0ad0, 0x0ae0, 16},
		{0x0ae1, 0x0ae3, 1},
		{0x0ae6, 0x0aef, 1},
		{0x0af1, 0x0b01, 16},
		{0x0b02, 0x0b03, 1},
		{0x0b05, 0x0b0c, 1},
		{0x0b0f, 0x0b10, 1},
		{0x0b13, 0x0b28, 1},
		{0x0b2a, 0x0b30, 1},
		{0x0b32, 0x0b33, 1},
		{0x0b35, 0x0b39, 1},
		{0x0b3c, 0x0b44, 1},
		{0x0b47, 0x0b48, 1},
		{0x0b4b, 0x0b4d, 1},
		{0x0b56, 0x0b57, 1},
		{0x0b5c, 0x0b5d, 1},
		{0x0b5f, 0x0b63, 1},
		{0x0b66, 0x0b71, 1},
		{0x0b82, 0x0b83, 1},
		{0x0b85, 0x0b8a, 1},
		{0x0b8e, 0x0b90, 1},
		{0x0b92, 0x0b95, 1},
		{0x0b99, 0x0b9a, 1},
		{0x0b9c, 0x0b9e, 2},
		{0x0b9f, 0x0ba3, 4},
		{0x0ba4, 0x0ba8, 4},
		{0x0ba9, 0x0baa, 1},
		{0x0bae, 0x0bb9, 1},
		{0x0bbe, 0x0bc2, 1},
		{0x0bc6, 0x0bc8, 1},
		{0x0bca, 0x0bcd, 1},
		{0x0bd0, 0x0bd7, 7},
		{0x0be6, 0x0bfa, 1},
		{0x0c01, 0x0c03, 1},
		{0x0c05, 0x0c0c, 1},
		{0x0c0e, 0x0c10, 1},
		{0x0c12, 0x0c28, 1},
		{0x0c2a, 0x0c33, 1},
		{0x0c35, 0x0c39, 1},
		{0x0c3d, 0x0c44, 1},
		{0x0c46, 0x0c48, 1},
		{0x0c4a, 0x0c4d, 1},
		{0x0c55, 0x0c56, 1},
		{0x0c58, 0x0c59, 1},
		{0x0c60, 0x0c63, 1},
		{0x0c66, 0x0c6f, 1},
		{0x0c78, 0x0c7f, 1},
		{0x0c82, 0x0c83, 1},
		{0x0c85, 0x0c8c, 1},
		{0x0c8e, 0x0c90, 1},
		{0x0c92, 0x0ca8, 1},
		{0x0caa, 0x0cb3, 1},
		{0x0cb5, 0x0cb9, 1},
		{0x0cbc, 0x0cc4, 1},
		{0x0cc6, 0x0cc8, 1},
		{0x0cca, 0x0ccd, 1},
		{0x0cd5, 0x0cd6, 1},
		{0x0cde, 0x0ce0, 2},
		{0x0ce1, 0x0ce3, 1},
		{0x0ce6, 0x0cef, 1},
		{0x0cf1, 0x0cf2, 1},
		{0x0d02, 0x0d03, 1},
		{0x0d05, 0x0d0c, 1},
		{0x0d0e, 0x0d10, 1},
		{0x0d12, 0x0d28, 1},
		{0x0d2a, 0x0d39, 1},
		{0x0d3d, 0x0d44, 1},
		{0x0d46, 0x0d48, 1},
		{0x0d4a, 0x0d4d, 1},
		{0x0d57, 0x0d60, 9},
		{0x0d61, 0x0d63, 1},
		{0x0d66, 0x0d75, 1},
		{0x0d79, 0x0d7f, 1},
		{0x0d82, 0x0d83, 1},
		{0x0d85, 0x0d96, 1},
		{0x0d9a, 0x0db1, 1},
		{0x0db3, 0x0dbb, 1},
		{0x0dbd, 0x0dc0, 3},
		{0x0dc1, 0x0dc6, 1},
		{0x0dca, 0x0dcf, 5},
		{0x0dd0, 0x0dd4, 1},
		{0x0dd6, 0x0dd8, 2},
		{0x0dd9, 0x0ddf, 1},
		{0x0df2, 0x0df4, 1},
		{0x0e01, 0x0e3a, 1},
		{0x0e3f, 0x0e5b, 1},
		{0x0e81, 0x0e82, 1},
		{0x0e84, 0x0e87, 3},
		{0x0e88, 0x0e8a, 2},
		{0x0e8d, 0x0e94, 7},
		{0x0e95, 0x0e97, 1},
		{0x0e99, 0x0e9f, 1},
		{0x0ea1, 0x0ea3, 1},
		{0x0ea5, 0x0ea7, 2},
		{0x0eaa, 0x0eab, 1},
		{0x0ead, 0x0eb9, 1},
		{0x0ebb, 0x0ebd, 1},
		{0x0ec0, 0x0ec4, 1},
		{0x0ec6, 0x0ec8, 2},
		{0x0ec9, 0x0ecd, 1},
		{0x0ed0, 0x0ed9, 1},
		{0x0edc, 0x0edd, 1},
		{0x0f00, 0x0f47, 1},
		{0x0f49, 0x0f6c, 1},
		{0x0f71, 0x0f8b, 1},
		{0x0f90, 0x0f97, 1},
		{0x0f99, 0x0fbc, 1},
		{0x0fbe, 0x0fcc, 1},
		{0x0fce, 0x0fd8, 1},
		{0x1000, 0x10c5, 1},
		{0x10d0, 0x10fc, 1},
		{0x1100, 0x1248, 1},
		{0x124a, 0x124d, 1},
		{0x1250, 0x1256, 1},
		{0x1258, 0x125a, 2},
		{0x125b, 0x125d, 1},
		{0x1260, 0x1288, 1},
		{0x128a, 0x128d, 1},
		{0x1290, 0x12b0, 1},
		{0x12b2, 0x12b5, 1},
		{0x12b8, 0x12be, 1},
		{0x12c0, 0x12c2, 2},
		{0x12c3, 0x12c5, 1},
		{0x12c8, 0x12d6, 1},
		{0x12d8, 0x1310, 1},
		{0x1312, 0x1315, 1},
		{0x1318, 0x135a, 1},
		{0x135f, 0x137c, 1},
		{0x1380, 0x1399, 1},
		{0x13a0, 0x13f4, 1},
		{0x1400, 0x169c, 1},
		{0x16a0, 0x16f0, 1},
		{0x1700, 0x170c, 1},
		{0x170e, 0x1714, 1},
		{0x1720, 0x1736, 1},
		{0x1740, 0x1753, 1},
		{0x1760, 0x176c, 1},
		{0x176e, 0x1770, 1},
		{0x1772, 0x1773, 1},
		{0x1780, 0x17dd, 1},
		{0x17e0, 0x17e9, 1},
		{0x17f0, 0x17f9, 1},
		{0x1800, 0x180e, 1},
		{0x1810, 0x1819, 1},
		{0x1820, 0x1877, 1},
		{0x1880, 0x18aa, 1},
		{0x18b0, 0x18f5, 1},
		{0x1900, 0x191c, 1},
		{0x1920, 0x192b, 1},
		{0x1930, 0x193b, 1},
		{0x1940, 0x1944, 4},
		{0x1945, 0x196d, 1},
		{0x1970, 0x1974, 1},
		{0x1980, 0x19ab, 1},
		{0x19b0, 0x19c9, 1},
		{0x19d0, 0x19da, 1},
		{0x19de, 0x1a1b, 1},
		{0x1a1e, 0x1a5e, 1},
		{0x1a60, 0x1a7c, 1},
		{0x1a7f, 0x1a89, 1},
		{0x1a90, 0x1a99, 1},
		{0x1aa0, 0x1aad, 1},
		{0x1b00, 0x1b4b, 1},
		{0x1b50, 0x1b7c, 1},
		{0x1b80, 0x1baa, 1},
		{0x1bae, 0x1bb9, 1},
		{0x1c00, 0x1c37, 1},
		{0x1c3b, 0x1c49, 1},
		{0x1c4d, 0x1c7f, 1},
		{0x1cd0, 0x1cf2, 1},
		{0x1d00, 0x1de6, 1},
		{0x1dfd, 0x1f15, 1},
		{0x1f18, 0x1f1d, 1},
		{0x1f20, 0x1f45, 1},
		{0x1f48, 0x1f4d, 1},
		{0x1f50, 0x1f57, 1},
		{0x1f59, 0x1f5f, 2},
		{0x1f60, 0x1f7d, 1},
		{0x1f80, 0x1fb4, 1},
		{0x1fb6, 0x1fc4, 1},
		{0x1fc6, 0x1fd3, 1},
		{0x1fd6, 0x1fdb, 1},
		{0x1fdd, 0x1fef, 1},
		{0x1ff2, 0x1ff4, 1},
		{0x1ff6, 0x1ffe, 1},
		{0x2000, 0x2064, 1},
		{0x206a, 0x2071, 1},
		{0x2074, 0x208e, 1},
		{0x2090, 0x2094, 1},
		{0x20a0, 0x20b8, 1},
		{0x20d0, 0x20f0, 1},
		{0x2100, 0x2189, 1},
		{0x2190, 0x23e8, 1},
		{0x2400, 0x2426, 1},
		{0x2440, 0x244a, 1},
		{0x2460, 0x26cd, 1},
		{0x26cf, 0x26e1, 1},
		{0x26e3, 0x26e8, 5},
		{0x26e9, 0x26ff, 1},
		{0x2701, 0x2704, 1},
		{0x2706, 0x2709, 1},
		{0x270c, 0x2727, 1},
		{0x2729, 0x274b, 1},
		{0x274d, 0x274f, 2},
		{0x2750, 0x2752, 1},
		{0x2756, 0x275e, 1},
		{0x2761, 0x2794, 1},
		{0x2798, 0x27af, 1},
		{0x27b1, 0x27be, 1},
		{0x27c0, 0x27ca, 1},
		{0x27cc, 0x27d0, 4},
		{0x27d1, 0x2b4c, 1},
		{0x2b50, 0x2b59, 1},
		{0x2c00, 0x2c2e, 1},
		{0x2c30, 0x2c5e, 1},
		{0x2c60, 0x2cf1, 1},
		{0x2cf9, 0x2d25, 1},
		{0x2d30, 0x2d65, 1},
		{0x2d6f, 0x2d80, 17},
		{0x2d81, 0x2d96, 1},
		{0x2da0, 0x2da6, 1},
		{0x2da8, 0x2dae, 1},
		{0x2db0, 0x2db6, 1},
		{0x2db8, 0x2dbe, 1},
		{0x2dc0, 0x2dc6, 1},
		{0x2dc8, 0x2dce, 1},
		{0x2dd0, 0x2dd6, 1},
		{0x2dd8, 0x2dde, 1},
		{0x2de0, 0x2e31, 1},
		{0x2e80, 0x2e99, 1},
		{0x2e9b, 0x2ef3, 1},
		{0x2f00, 0x2fd5, 1},
		{0x2ff0, 0x2ffb, 1},
		{0x3000, 0x303f, 1},
		{0x3041, 0x3096, 1},
		{0x3099, 0x30ff, 1},
		{0x3105, 0x312d, 1},
		{0x3131, 0x318e, 1},
		{0x3190, 0x31b7, 1},
		{0x31c0, 0x31e3, 1},
		{0x31f0, 0x321e, 1},
		{0x3220, 0x32fe, 1},
		{0x3300, 0x4db5, 1},
		{0x4dc0, 0x9fcb, 1},
		{0xa000, 0xa48c, 1},
		{0xa490, 0xa4c6, 1},
		{0xa4d0, 0xa62b, 1},
		{0xa640, 0xa65f, 1},
		{0xa662, 0xa673, 1},
		{0xa67c, 0xa697, 1},
		{0xa6a0, 0xa6f7, 1},
		{0xa700, 0xa78c, 1},
		{0xa7fb, 0xa82b, 1},
		{0xa830, 0xa839, 1},
		{0xa840, 0xa877, 1},
		{0xa880, 0xa8c4, 1},
		{0xa8ce, 0xa8d9, 1},
		{0xa8e0, 0xa8fb, 1},
		{0xa900, 0xa953, 1},
		{0xa95f, 0xa97c, 1},
		{0xa980, 0xa9cd, 1},
		{0xa9cf, 0xa9d9, 1},
		{0xa9de, 0xa9df, 1},
		{0xaa00, 0xaa36, 1},
		{0xaa40, 0xaa4d, 1},
		{0xaa50, 0xaa59, 1},
		{0xaa5c, 0xaa7b, 1},
		{0xaa80, 0xaac2, 1},
		{0xaadb, 0xaadf, 1},
		{0xabc0, 0xabed, 1},
		{0xabf0, 0xabf9, 1},
		{0xac00, 0xd7a3, 1},
		{0xd7b0, 0xd7c6, 1},
		{0xd7cb, 0xd7fb, 1},
		{0xd800, 0xfa2d, 1},
		{0xfa30, 0xfa6d, 1},
		{0xfa70, 0xfad9, 1},
		{0xfb00, 0xfb06, 1},
		{0xfb13, 0xfb17, 1},
		{0xfb1d, 0xfb36, 1},
		{0xfb38, 0xfb3c, 1},
		{0xfb3e, 0xfb40, 2},
		{0xfb41, 0xfb43, 2},
		{0xfb44, 0xfb46, 2},
		{0xfb47, 0xfbb1, 1},
		{0xfbd3, 0xfd3f, 1},
		{0xfd50, 0xfd8f, 1},
		{0xfd92, 0xfdc7, 1},
		{0xfdf0, 0xfdfd, 1},
		{0xfe00, 0xfe19, 1},
		{0xfe20, 0xfe26, 1},
		{0xfe30, 0xfe52, 1},
		{0xfe54, 0xfe66, 1},
		{0xfe68, 0xfe6b, 1},
		{0xfe70, 0xfe74, 1},
		{0xfe76, 0xfefc, 1},
		{0xfeff, 0xff01, 2},
		{0xff02, 0xffbe, 1},
		{0xffc2, 0xffc7, 1},
		{0xffca, 0xffcf, 1},
		{0xffd2, 0xffd7, 1},
		{0xffda, 0xffdc, 1},
		{0xffe0, 0xffe6, 1},
		{0xffe8, 0xffee, 1},
		{0xfff9, 0xfffd, 1},
	},
	R32: []unicode.Range32{
		{0x00010000, 0x0001000b, 1},
		{0x0001000d, 0x00010026, 1},
		{0x00010028, 0x0001003a, 1},
		{0x0001003c, 0x0001003d, 1},
		{0x0001003f, 0x0001004d, 1},
		{0x00010050, 0x0001005d, 1},
		{0x00010080, 0x000100fa, 1},
		{0x00010100, 0x00010102, 1},
		{0x00010107, 0x00010133, 1},
		{0x00010137, 0x0001018a, 1},
		{0x00010190, 0x0001019b, 1},
		{0x000101d0, 0x000101fd, 1},
		{0x00010280, 0x0001029c, 1},
		{0x000102a0, 0x000102d0, 1},
		{0x00010300, 0x0001031e, 1},
		{0x00010320, 0x00010323, 1},
		{0x00010330, 0x0001034a, 1},
		{0x00010380, 0x0001039d, 1},
		{0x0001039f, 0x000103c3, 1},
		{0x000103c8, 0x000103d5, 1},
		{0x00010400, 0x0001049d, 1},
		{0x000104a0, 0x000104a9, 1},
		{0x00010800, 0x00010805, 1},
		{0x00010808, 0x0001080a, 2},
		{0x0001080b, 0x00010835, 1},
		{0x00010837, 0x00010838, 1},
		{0x0001083c, 0x0001083f, 3},
		{0x00010840, 0x00010855, 1},
		{0x00010857, 0x0001085f, 1},
		{0x00010900, 0x0001091b, 1},
		{0x0001091f, 0x00010939, 1},
		{0x0001093f, 0x00010a00, 193},
		{0x00010a01, 0x00010a03, 1},
		{0x00010a05, 0x00010a06, 1},
		{0x00010a0c, 0x00010a13, 1},
		{0x00010a15, 0x00010a17, 1},
		{0x00010a19, 0x00010a33, 1},
		{0x00010a38, 0x00010a3a, 1},
		{0x00010a3f, 0x00010a47, 1},
		{0x00010a50, 0x00010a58, 1},
		{0x00010a60, 0x00010a7f, 1},
		{0x00010b00, 0x00010b35, 1},
		{0x00010b39, 0x00010b55, 1},
		{0x00010b58, 0x00010b72, 1},
		{0x00010b78, 0x00010b7f, 1},
		{0x00010c00, 0x00010c48, 1},
		{0x00010e60, 0x00010e7e, 1},
		{0x00011080, 0x000110c1, 1},
		{0x00012000, 0x0001236e, 1},
		{0x00012400, 0x00012462, 1},
		{0x00012470, 0x00012473, 1},
		{0x00013000, 0x0001342e, 1},
		{0x0001d000, 0x0001d0f5, 1},
		{0x0001d100, 0x0001d126, 1},
		{0x0001d129, 0x0001d1dd, 1},
		{0x0001d200, 0x0001d245, 1},
		{0x0001d300, 0x0001d356, 1},
		{0x0001d360, 0x0001d371, 1},
		{0x0001d400, 0x0001d454, 1},
		{0x0001d456, 0x0001d49c, 1},
		{0x0001d49e, 0x0001d49f