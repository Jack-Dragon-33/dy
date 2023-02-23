package util

import (
	douyin_core "dy/biz/model/douyin_core"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	workerIdBitsMoveLen = uint(8)
	maxWorkerId         = int64(-1 ^ (-1 << workerIdBitsMoveLen))
	timerIdBitsMoveLen  = uint(17)
	maxNumId            = int64(-1 ^ (-1 << 9))
	TIME_FORMAT         = "2006-01-02 03:04:05"
	BASE_URL            = "http://f33qqu.natappfree.cc"
	URL_FORMAT          = "%s/static/%s"
)

var (
	TokenExpireDuration = time.Hour * 2
	MySecret            = []byte("sxl")
)

type Worker1 struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	workerId  int64      // 机器编码
	timestamp int64      // 记录时间戳
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

// 初始化ID生成结构体
// workerId 机器的编号
func NewWorker1(workerId int64) *Worker1 {
	if workerId > maxWorkerId {
		panic("workerId 不能大于最大值")
	}
	return &Worker1{workerId: workerId, timestamp: 0, number: 0}
}
func (w *Worker1) GetId() int64 {
	epoch := int64(1613811738) // 设置为去年今天的时间戳...因为位数变了后,几百年都用不完,,实际可以设置上线日期的
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixMilli() // 获得现在对应的时间戳
	if now < w.timestamp {
		// 当机器出现时钟回拨时报错
		panic("Clock moved backwards.  Refusing to generate id for %d milliseconds")
	}
	if w.timestamp == now {
		w.number++
		if w.number > maxNumId { //此处为最大节点ID,大概是2^9-1 511条,
			for now <= w.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		w.number = 0
		w.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	}
	ID := int64((now-epoch)<<timerIdBitsMoveLen | (w.workerId << workerIdBitsMoveLen) | (w.number))
	return ID
}
func GenerateID() int64 {
	worker := NewWorker1(55)
	return worker.GetId()
}
func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil // 这是我的secret
	}
}

// func GenToken(userId int,username string)(string,error){
// 	return nil,nil
// }

func GeneratorToken(UserInfo *douyin_core.User) (string, error) {
	claim := douyin_core.MyClaim{
		UserId: UserInfo.Id,
		Name:   UserInfo.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
		}}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString(MySecret)
}
func ParseToken(tokenss string) (*douyin_core.MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenss, &douyin_core.MyClaim{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*douyin_core.MyClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
func ParseTime(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format(TIME_FORMAT)
}
func GetURL(filename string) string {
	return fmt.Sprintf(URL_FORMAT, BASE_URL, filename)
}
