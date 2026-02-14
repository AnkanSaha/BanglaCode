//go:build freebsd

package system

import (
	"syscall"
)

// getDiskStats returns disk statistics for a given path (FreeBSD-specific)
func getDiskStats(path string) (total, free, avail uint64, err error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, 0, 0, err
	}

	// FreeBSD uses slightly different field types
	blockSize := uint64(stat.Bsize)

	// Total size = block size * total blocks
	total = uint64(stat.Blocks) * blockSize
	// Free space = block size * free blocks
	free = uint64(stat.Bfree) * blockSize
	// Available space = block size * available blocks
	avail = uint64(stat.Bavail) * blockSize

	return total, free, avail, nil
}
