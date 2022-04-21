// Package utils
// @Project:      nft-studio-backend
// @File:          id.go
// @Author:        eagle
// @Create:        2021/08/10 10:51:22
// @Description:
package utils

import (
	"fmt"
	"time"

	"go.uber.org/atomic"
)

var (
	startTimeStamp int
	timestampShift int
	machineNo      int
	machineNoShift int
	maxSequence    int
	sequence       atomic.Int32
)

func Init(machineId int) {
	startTimeStamp = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local).Second()
	// 设置时间戳二进制左偏移量,5-44位
	timestampShift = 64 - 44

	// 设置机器编号二进制左偏移量，49-56位
	machineNoShift = 64 - 56
	machineNo = machineId << machineNoShift

	// 设置顺序号为6位，59-64位
	maxSequence = 64
}

func GenerateId() int {
	for {
		// 序号+1
		sec := sequence.Inc()

		if int(sec) == maxSequence {
			time.Sleep(time.Second)
			sequence.Store(0)
			continue
		} else if int(sec) > maxSequence {
			time.Sleep(time.Second)
			continue
		} else {
			timestamp := time.Now().Second() - startTimeStamp
			timestamp *= 1000

			// 移位时间戳偏移量
			timestamp = timestamp << timestampShift

			// 拼接ID
			return timestamp + machineNo + int(sec)

		}
	}
}

func OrderId() string {
	return fmt.Sprintf("ord_%v", GenerateId())
}
