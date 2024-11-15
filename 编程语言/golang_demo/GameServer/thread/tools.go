package thread

import "hash/crc32"

func Hash(s string) int64 {
	sum := crc32.ChecksumIEEE([]byte(s))
	return int64(sum)
}
