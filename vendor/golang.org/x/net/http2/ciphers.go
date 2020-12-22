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
	cipher_TLS_DH_RSA_WITH_DES_CBC_SHA           uint16 = 0x000F
	cipher_TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA      uint16 = 0x0010
	cipher_TLS_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA uint16 = 0x0011
	cipher_TLS_DHE_DSS_WITH_DES_CBC_SHA          uint16 = 0x0012
	cipher_TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA     uint16 = 0x0013
	cipher_TLS_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA uint16 = 0x0014
	cipher_TLS_DHE_RSA_WITH_DES_CBC_SHA          uint16 = 0x0015
	cipher_TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA     uint16 = 0x0016
	cipher_TLS_DH_anon_EXPORT_WITH_RC4_40_MD5    uint16 = 0x0017
	cipher_TLS_DH_anon_WITH_RC4_128_MD5          uint16 = 0x0018
	cipher_TLS_DH_anon_EXPORT_WITH_DES40_CBC_SHA uint16 = 0x0019
	cipher_TLS_DH_anon_WITH_DES_CBC_SHA          uint16 = 0x001A
	cipher_TLS_DH_anon_WITH_3DES_EDE_CBC_SHA     uint16 = 0x001B
	// Reserved uint16 =  0x001C-1D
	cipher_TLS_KRB5_WITH_DES_CBC_SHA             uint16 = 0x001E
	cipher_TLS_KRB5_WITH_3DES_EDE_CBC_SHA        uint16 = 0x001F
	cipher_TLS_KRB5_WITH_RC4_128_SHA             uint16 = 0x0020
	cipher_TLS_KRB5_WITH_IDEA_CBC_SHA            uint16 = 0x0021
	cipher_TLS_KRB5_WITH_DES_CBC_MD5             uint16 = 0x0022
	cipher_TLS_KRB5_WITH_3DES_EDE_CBC_MD5        uint16 = 0x0023
	cipher_TLS_KRB5_WITH_RC4_128_MD5             uint16 = 0x0024
	cipher_TLS_KRB5_WITH_IDEA_CBC_MD5            uint16 = 0x0025
	cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_SHA   uint16 = 0x0026
	cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_SHA   uint16 = 0x0027
	cipher_TLS_KRB5_EXPORT_WITH_RC4_40_SHA       uint16 = 0x0028
	cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_MD5   uint16 = 0x0029
	cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_MD5   uint16 = 0x002A
	cipher_TLS_KRB5_EXPORT_WITH_RC4_40_MD5       uint16 = 0x002B
	cipher_TLS_PSK_WITH_NULL_SHA                 uint16 = 0x002C
	cipher_TLS_DHE_PSK_WITH_NULL_SHA             uint16 = 0x002D
	cipher_TLS_RSA_PSK_WITH_NULL_SHA             uint16 = 0x002E
	cipher_TLS_RSA_WITH_AES_128_CBC_SHA          uint16 = 0x002F
	cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA       uint16 = 0x0030
	cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA       uint16 = 0x0031
	cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA      uint16 = 0x0032
	cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA      uint16 = 0x0033
	cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA      uint16 = 0x0034
	cipher_TLS_RSA_WITH_AES_256_CBC_SHA          uint16 = 0x0035
	cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA       uint16 = 0x0036
	cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA       uint16 = 0x0037
	cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA      uint16 = 0x0038
	cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA      uint16 = 0x0039
	cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA      uint16 = 0x003A
	cipher_TLS_RSA_WITH_NULL_SHA256              uint16 = 0x003B
	cipher_TLS_RSA_WITH_AES_128_CBC_SHA256       uint16 = 0x003C
	cipher_TLS_RSA_WITH_AES_256_CBC_SHA256       uint16 = 0x003D
	cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA256    uint16 = 0x003E
	cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA256    uint16 = 0x003F
	cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA256   uint16 = 0x0040
	cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA     uint16 = 0x0041
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA  uint16 = 0x0042
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA  uint16 = 0x0043
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA uint16 = 0x0044
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA uint16 = 0x0045
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA uint16 = 0x0046
	// Reserved uint16 =  0x0047-4F
	// Reserved uint16 =  0x0050-58
	// Reserved uint16 =  0x0059-5C
	// Unassigned uint16 =  0x005D-5F
	// Reserved uint16 =  0x0060-66
	cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 uint16 =