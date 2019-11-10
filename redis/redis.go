package redis

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/redis.v5"
)

var (
	ClientRed *redis.Client
	address   string
	port      string
	dbtype    int
	queuename string
)

func init() {
	address = beego.AppConfig.DefaultString("redis.address", "54.254.222.187")
	port = beego.AppConfig.DefaultString("redis.port", "6077")
	dbtype = beego.AppConfig.DefaultInt("redis.dbtype", 10)
	queuename = beego.AppConfig.DefaultString("redis.queuename", "queOasis")

	ClientRed = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: "",     // no password set
		DB:       dbtype, // use default DB
	})

	pong, err := ClientRed.Ping().Result()
	fmt.Println("Redis Ping ", pong)
	fmt.Println("Redis Ping ", err)
	//RunSubscriber()
}

// GetRedisConnection ...
func GetRedisConnection() *redis.Client {
	return ClientRed
}

// GetRedisUri ...
func GetRedisUri() string {
	return "redis://" + address + ":" + port + "/"
}

// GetQueueName ...
func GetQueueName() string {
	return queuename
}

/*
 Redis Standard Set
*/
func SaveRedis(key string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		err = ClientRed.Set(key, val, 0).Err()
		if err == nil {
			break
		}
	}
	return err
}

/*
 Redis Standard Get
*/
func GetRedisKey(Key string) (string, error) {
	val2, err := ClientRed.Get(Key).Result()
	if err == redis.Nil {
		err = errors.New("Key Does Not Exists")
		fmt.Println("keystruct does not exists")
	} else if err != nil {
		fmt.Println("Error : ", err.Error())
	} //else {
	//fmt.Println("keystruct", val2)
	//}
	return val2, err
}

// GetRedisKeyPattern ...
func GetRedisKeyPattern(Key string) ([]string, error) {
	val2, err := ClientRed.Keys(Key).Result()
	if err == redis.Nil {
		err = errors.New("Key Does Not Exists")
		fmt.Println("keystruct does not exists")
	} else if err != nil {
		fmt.Println("Error : ", err.Error())
	} //else {
	//fmt.Println("keystruct", val2)
	//}
	return val2, err
}

func DelRedisKey(key string) error {
	return ClientRed.Del(key).Err()
}

/*
delayto * max = total timeout
*/
func GetDataRedis(key string, delayto, max int) (bool, string) {
	for i := 0; i < max; i++ {
		data, err := GetRedisKey(key)
		fmt.Println(" Err : ", err)
		fmt.Println(" data : ", data)
		if err == nil {
			return true, data
		}
		time.Sleep(time.Duration(delayto) * time.Second)
	}
	return false, ""
}

/*
 Redis Standard Set Expired
*/
func SaveRedisExp(key string, menit string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit)
		err = ClientRed.Set(key, val, duration).Err()
		if err == nil {
			break
		}
		fmt.Println("Error : ", err)
	}
	return err
}

func SaveRedisCounter(key string) (int64, error) {
	incr := ClientRed.Incr(key)
	return incr.Val(), incr.Err()
}

func SaveRedisCounterAuto(key string, autonom int64) (int64, error) {
	incr := ClientRed.IncrBy(key, autonom)
	return incr.Val(), incr.Err()
}

func GetRedisCounter(key string) (int64, error) {
	decr := ClientRed.Decr(key)
	return decr.Val(), decr.Err()

}

func GetRedisCounterIncr(key string) (int64, error) {
	decr := ClientRed.Incr(key)
	return decr.Val(), decr.Err()

}
