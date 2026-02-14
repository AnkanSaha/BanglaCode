//go:build openbsd

package system

import (
	"syscall"
)

// getDiskStats returns disk statistics for a given path (OpenBSD-specific)
func getDiskStats(path string) (total, free, avail uint64, err error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, 0, 0, err
	}

	// OpenBSD Statfs_t fields
	blockSize := uint64(stat.F_bsize)

	// Total size = block size * total blocks
	total = stat.F_blocks * blockSize
	// Free space = block size * free blocks
	free = stat.F_bfree * blockSize
	// Available space = block size * available blocks
	avail = uint64(stat.F_bavail) * blockSize

	return total, free, avail, nil
}
