/*
 * @Author: dowell87
 * @Date: 2021-11-14 20:40:57
 * @Descripttion:
 * @LastEditTime: 2021-11-15 11:38:38
 */
package algs15

import "strings"

// // 代码超过了运行速度

// func Robot(command string, obstacles [][]int, x int, y int) bool {
// 	newX, newY := 0, 0
// 	j := len(command)
// 	i := 0
// 	for {
// 		z := i % j
// 		if command[z:z+1] == "R" {
// 			newX++
// 		} else {
// 			newY++
// 		}
// 		// 到达目的地返回成功
// 		if newX == x && newY == y {
// 			return true
// 		}
// 		for w := 0; w < len(obstacles); w++ {
// 			// 匹配是否碰到障碍物
// 			if obstacles[w][0] == newX && obstacles[w][1] == newY {
// 				return false
// 			}
// 			// 已经超越障碍物从障碍物中删除
// 			if obstacles[w][0] < newX || obstacles[w][1] < newY {
// 				obstacles = append(obstacles[:w], obstacles[w+1:]...)
// 				w--
// 			}
// 		}
// 		// 超过终点终止运行
// 		if x+1 < newX || y+1 < newY {
// 			return false
// 		}
// 		i++
// 	}
// }

func Robot(command string, obstacles [][]int, x int, y int) bool {
	// 如果目标点不在路径上，返回失败
	if !isOnThePath(command, x, y) {
		return false
	}
	for _, V := range obstacles {
		// 判断有效的故障点是否在路径上（故障的步数大于等于目标的点，视为无效故障）
		if (x+y > V[0]+V[1]) && isOnThePath(command, V[0], V[1]) {
			return false
		}
	}
	return true
}

func isOnThePath(command string, x int, y int) bool {
	uNum := strings.Count(command, "U")*((x+y)/len(command)) + strings.Count(command[0:(x+y)%len(command)], "U")
	rNum := strings.Count(command, "R")*((x+y)/len(command)) + strings.Count(command[0:(x+y)%len(command)], "R")
	if uNum == y && rNum == x {
		return true
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// moveX, moveY := 0, 0 //向上的步长
// //计算一轮指令的向上和向右的步长
// for i := 0; i < len(command); i++ {
// 	if command[i:i+1] == "U" {
// 		moveY++
// 	} else {
// 		moveX++
// 	}
// }
// //计算每一个在范围的障碍物是否可达
// for i := 0; i < len(obstacles); i++ {
// 	//如果障碍物超出范围，则跳过
// 	if obstacles[i][0] > x || obstacles[i][1] > y {
// 		continue
// 	}
// 	//cycle为判断障碍物走过了多少次循环，取横坐标和纵坐标除以相应步长的最小值。
// 	cycle := Max(obstacles[i][0]/moveX, obstacles[i][1]/moveY)
// 	//减去走过的步长，剩下的则为此次循环应走的步长。
// 	targetX := obstacles[i][0] - cycle*moveX
// 	targetY := obstacles[i][1] - cycle*moveY
// 	//执行指令，判断障碍物是否可达
// 	for j := 0; j < len(command) && (targetX > 0 || targetY > 0); j++ {
// 		if command[j:j+1] == "U" {
// 			targetY--
// 		} else {
// 			targetX--
// 		}
// 	}
// 	//如果障碍物可达，则为终点不可达，直接返回
// 	if targetX == 0 && targetY == 0 {
// 		return false
// 	}
// }
// //跟障碍物同理，判断终点是否可达
// cycle := Max(x/moveX, y/moveY)
// x = x - cycle*moveX
// y = y - cycle*moveY
// for i := 0; i < len(command) && (y > 0 || x > 0); i++ {
// 	if command[i:i+1] == "U" {
// 		y--
// 	} else {
// 		x--
// 	}
// }
// if x == 0 && y == 0 {
// 	return true
// }
// return false
