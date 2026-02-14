//go:build netbsd

package system

import (
	"errors"
)

// getDiskStats returns disk statistics for a given path (NetBSD-specific)
// Note: NetBSD syscall API is limited, so we return an error for now
func getDiskStats(path string) (total, free, avail uint64, err error) {
	return 0, 0, 0, errors.New("disk stats not supported on NetBSD")
}
