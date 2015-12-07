package mygo01
import (
	"gopkg.in/redis.v3"
	"strconv"
)

type Channel struct {
	id int
	chaters []Chater
}

type Chater struct {
	userid string
	readMsgId int

}

func onMsg(userid, channel, msg string) {

}

type MsgCenter struct{
	conn	*redis.Client

}

func (m *MsgCenter) reciveMsg(userid, channel, msg string){
	id := m.conn.Incr("msg:"+channel+":id:")
	m.conn.HSet("msg"+strconv.Itoa(id), msg)


}