// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package nfsd implements parsing of /proc/net/rpc/nfsd.
// Fields are documented in https://www.svennd.be/nfsd-stats-explained-procnetrpcnfsd/
package nfs

// ReplyCache models the "rc" line.
type ReplyCache struct {
	Hits    uint64
	Misses  uint64
	NoCache uint64
}

// FileHandles models the "fh" line.
type FileHandles struct {
	Stale        uint64
	TotalLookups uint64
	AnonLookups  uint64
	DirNoCache   uint64
	NoDirNoCache uint64
}

// InputOutput models the "io" line.
type InputOutput struct {
	Read  uint64
	Write uint64
}

// Threads models the "th" line.
type Threads struct {
	Threads uint64
	FullCnt uint64
}

// ReadAheadCache models the "ra" line.
type ReadAheadCache struct {
	CacheSize      uint64
	CacheHistogram []uint64
	NotFound       uint64
}

// Network models the "net" line.
type Network struct {
	NetCount   uint64
	UDPCount   uint64
	TCPCount   uint64
	TCPConnect uint64
}

// ClientRPC models the nfs "rpc" line.
type ClientRPC struct {
	RPCCount        uint64
	Retransmissions uint64
	AuthRefreshe