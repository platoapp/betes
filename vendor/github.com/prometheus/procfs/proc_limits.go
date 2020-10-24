package procfs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// ProcLimits represents the soft limits for each of the process's resource
// limits. For more information see getrlimit(2):
// http://man7.org/linux/man-pages/man2/getrlimit.2.html.
type ProcLimits struct {
	// CPU time limit in seconds.
	CPUTime int64
	// Maximum size of files that the process may create.
	FileSize int64
	// Maximum size of the process's data segment (initialized data,
	// uninitialized data, and heap).
	DataSize int64
	// Maximum size of the process stack in bytes.
	StackSize int64
	// Maximum size of a core file.
	CoreFileSize int64
	// Limit of the process's resident set in pages.
	ResidentSet int64
	// Maximum number of processes that can be created for the real user ID of
	// the calling process.
	Processes int64
	// Value one greater than the maximum file descriptor number that can be
	// opened by this process.
	OpenFiles int64
	// Maximum number of bytes of memory that may be locked into RAM.
	LockedMemory int64
	// Maximum size of the process's virtual memory address space in bytes.
	AddressSpace int64
	// Limit on the combined number of flock(2) locks and fcntl(2) leases that
	// this process may establish.
	FileLocks int64
	// Limit of signals that may be queued for the real user ID of the calling
	// process.
	PendingSignals int64
	// Limit on the number of bytes that can be allocated for POSIX message
	// queues for the real user ID of the calling process.
	MsqqueueSize int64
	// Limit of the nice priority set using setpriority(2) or nice(2).
	NicePriority 