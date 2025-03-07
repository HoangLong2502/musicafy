package common

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
)

// UID is method to generate an virtual unique indentifier for whole system
// its structure contains 62 bits: LocalID - ObjectType - ShardID
// 32 bits for Local ID, max (2^32) - 1
// 10 bits for Object Type
// 18 bits for Shard ID
type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objType,
		shardID:    shardID,
	}
}

// Shard: 1, Object: 1, ID: 1 => 0001 0001 0001
// 1 << 8 = 0001 0000 0000
// 1 << 4 = 		1 0000
// 1 << 0 = 			 1
// => 0001 0001 0001
func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetShardID() uint32 {
	return uid.shardID
}

func (uid UID) GetObjType() int {
	return uid.objectType
}

func DecompseUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}

	// x = 1110 1110 0101 => x >> 4 = 1110 1110 & 0000 1111 = 1110
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecompseUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJson() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}
