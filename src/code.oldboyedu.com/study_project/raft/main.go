package main

import (
	"sync"
)

//1.实现3节点选举
//2.改造代码成分布式选举代码，加入RPC调用
//3.演示完整代码 自动选主 日志复制

// 定义3节点常量
const raftCount = 3

// 声明leader对象
type Leader struct {
	//任期
	term int
	//LeaderId编号
	leaderId int
}

// 声明raft
type Raft struct {
	//锁
	mu sync.Mutex
	//节点编号
	me int
	//当前任期
	currentTerm int
	//为哪个节点投票
	votedFor int
	//3个状态
	//0 follower 1 candidate 2 leader
	state int
	//发送最后一条数据的时间
	lastMessageTime int64
	//设置当前节点的领导
	currentLeader int
	//节点间发信息的通道
	message chan bool
	//选举通道
	electCh chan bool
	//心跳信号的通道
	heartBeat chan bool
	//返回心跳信号的通道
	heartBeatRe chan bool
	//超时时间
	timeout int
}

// 0 还没上任 -1 没有编号
var leader = Leader{0, -1}

func main() {
	//过程:有三个节点,最初都是follower
	//若有candidate状态，进行投票拉票
	//会产生leader

	//创建3个节点
	for i := 0; i < raftCount; i++ {
		//创建3个raft节点

	}
}

func Make(me int) *Raft {

}
