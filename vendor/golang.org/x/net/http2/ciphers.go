// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

// A list of the possible cipher suite ids. Taken from
// https://www.iana.org/assignments/tls-parameters/tls-parameters.txt

const (
	cipher_TLS_NULL_WITH_NULL_NULL               uint16 = 0x0000
	cipher_TLS_RSA_WITH_NULL_MD5                 uint16 = 0x0001
	cipher_TLS_RSA_WITH_NULL_SHA                 uint16 = 0x0002
	cipher_TLS_RSA_EXPORT_WITH_RC4_40_MD5        uint16 = 0x0003
	cipher_TLS_RSA_WITH_RC4_128_MD5              uint16 = 0x0004
	cipher_TLS_RSA_WITH_RC4_128_SHA              uint16 = 0x0005
	cipher_TLS_RSA_EXPORT_WITH_RC2_CBC_40_MD5    uint16 = 0x0006
	cipher_TLS_RSA_WITH_IDEA_CBC_SHA             uint16 = 0x0007
	cipher_TLS_RSA_EXPORT_WITH_DES40_CBC_SHA     uint16 = 0x0008
	cipher_TLS_RSA_WITH_DES_CBC_SHA              uint16 = 0x0009
	cipher_TLS_RSA_WITH_3DES_EDE_CBC_SHA         uint16 = 0x000A
	cipher_TLS_DH_DSS_EXPORT_WITH_DES40_CBC_SHA  uint16 = 0x000B
	cipher_TLS_DH_DSS_WITH_DES_CBC_SHA           uint16 = 0x000C
	cipher_TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA      uint16 = 0x000D
	cipher_TLS_DH_RSA_EXPORT_WITH_DES40_CBC_SHA  uint16 = 0x000E
	cipher_TLS_DH_RSA_WITH_DES_CBC_SHA           uint