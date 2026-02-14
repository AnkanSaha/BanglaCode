//go:build windows

package system

import (
	"os"
)

// getFileOwnership extracts UID and GID from file info (Windows fallback)
// Windows doesn't have UID/GID, so we return default values
func getFileOwnership(info os.FileInfo) (uid, gid uint32, ok bool) {
	return 0, 0, false
}

// getAccessTime extracts access time from file info (Windows fallback)
// Windows doesn't expose access time via syscall.Stat_t
func getAccessTime(info os.FileInfo) (int64, bool) {
	return 0, false
}

// getCreationTime extracts creation time from file info (Windows fallback)
// Windows doesn't expose creation time via syscall.Stat_t
func getCreationTime(info os.FileInfo) (int64, bool) {
	return 0, false
}

// getFileInode extracts inode number from file info (Windows fallback)
// Windows doesn't have inodes
func getFileInode(info os.FileInfo) (uint64, bool) {
	return 0, false
}

// getFileLinkCount extracts link count from file info (Windows fallback)
// Windows doesn't expose link count in the same way
func getFileLinkCount(info os.FileInfo) (uint64, bool) {
	return 0, false
}
