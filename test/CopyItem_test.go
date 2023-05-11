package test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// 通过FindItem获取AAAWAGhlemhhb3d1MUBraW5nc29mdC5jb20ARgAAAAAAC9AYe7rimkOhvpHgvNsOzgcAOjhiyJgKLEO13EB8uUUY2gAAAAABDAAAgoNscah2/UO0tMiYJBHiWQAB0psqxAAA
const itemId = "AAAWAGhlemhhb3d1MUBraW5nc29mdC5jb20ARgAAAAAAC9AYe7rimkOhvpHgvNsOzgcAOjhiyJgKLEO13EB8uUUY2gAAAAABDAAAgoNscah2/UO0tMiYJBHiWQAB0psqxAAA"

func TestCopyItem(t *testing.T) {
	defer dumpFile.Sync()

}

func TestIdInfo(t *testing.T) {
	data, err := base64.RawStdEncoding.DecodeString(itemId)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
