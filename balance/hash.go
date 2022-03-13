package balance

import (
	"Amanisis/model/ServerModel"
	"hash/crc32"
)

func hashKey(key string) uint32 {
	scratch := []byte(key)
	return crc32.ChecksumIEEE(scratch)
}
func GetServerByHash(list []ServerModel.Server, ip string) ServerModel.Server {
	ServerList := GetUsedServer(list)
	sourcehash := int(hashKey(ip))
	key := sourcehash % len(list)
	return ServerList[key]
}
