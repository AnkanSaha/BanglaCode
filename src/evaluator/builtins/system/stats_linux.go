//go:build linux

package system

import (
	"syscall"
)

// getDiskStats returns disk statistics for a given path (Linux-specific)
func getDiskStats(path string) (total, free, avail uint64, err error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, 0, 0, err
	}

	// Total size = block size * total blocks
	total = stat.Blocks * uint64(stat.Bsize)
	// Free space = block size * free blocks
	free = stat.Bfree * uint64(stat.Bsize)
	// Available space = block size * available blocks
	avail = stat.Bavail * uint64(stat.Bsize)

	return total, free, avail, nil
}
