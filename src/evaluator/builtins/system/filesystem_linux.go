//go:build linux

package system

import (
	"os"
	"syscall"
)

// getFileOwnership extracts UID and GID from file info (Linux-specific)
func getFileOwnership(info os.FileInfo) (uid, gid uint32, ok bool) {
	if stat, statOk := info.Sys().(*syscall.Stat_t); statOk {
		return stat.Uid, stat.Gid, true
	}
	return 0, 0, false
}

// getAccessTime extracts access time from file info (Linux-specific)
func getAccessTime(info os.FileInfo) (int64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return int64(stat.Atim.Sec), true
	}
	return 0, false
}

// getCreationTime extracts creation time from file info (Linux-specific)
// Note: Linux uses change time (Ctim), not birth time
func getCreationTime(info os.FileInfo) (int64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return int64(stat.Ctim.Sec), true
	}
	return 0, false
}

// getFileInode extracts inode number from file info (Linux-specific)
func getFileInode(info os.FileInfo) (uint64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return stat.Ino, true
	}
	return 0, false
}

// getFileLinkCount extracts link count from file info (Linux-specific)
func getFileLinkCount(info os.FileInfo) (uint64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return uint64(stat.Nlink), true
	}
	return 0, false
}
