//go:build netbsd

package system

import (
	"os"
	"syscall"
)

// getFileOwnership extracts UID and GID from file info (NetBSD-specific)
func getFileOwnership(info os.FileInfo) (uid, gid uint32, ok bool) {
	if stat, statOk := info.Sys().(*syscall.Stat_t); statOk {
		return stat.Uid, stat.Gid, true
	}
	return 0, 0, false
}

// getAccessTime extracts access time from file info (NetBSD-specific)
func getAccessTime(info os.FileInfo) (int64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return int64(stat.Atimespec.Sec), true
	}
	return 0, false
}

// getCreationTime extracts creation time from file info (NetBSD-specific)
func getCreationTime(info os.FileInfo) (int64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return int64(stat.Ctimespec.Sec), true
	}
	return 0, false
}

// getFileInode extracts inode number from file info (NetBSD-specific)
func getFileInode(info os.FileInfo) (uint64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return stat.Ino, true
	}
	return 0, false
}

// getFileLinkCount extracts link count from file info (NetBSD-specific)
func getFileLinkCount(info os.FileInfo) (uint64, bool) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return uint64(stat.Nlink), true
	}
	return 0, false
}
