package DataConn

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var(
	RedisPool *redis.Pool
	SubRedisPool *redis.Pool
)
//RedisSetup
func RedisSetup(RedisIP, RedisPW string, sc *Cfg) {
	//Config, err := dbCfg.GetSection("Redis")
	//checkErr(err)
	fmt.Println("Redis Config init....")

	//Set up Sc.RedisPool for Normal request
	RedisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   30,
		IdleTimeout: 3 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", RedisIP,
				redis.DialPassword(RedisPW),
				redis.DialConnectTimeout(2*time.Second),
				redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(2*time.Second))

			if err != nil {
				return nil, err
			}
			_, err = con.Do("select", 2)
			CheckError(err)
			/* 挑战字方式登录暂时不考虑 故先注释 日后又加*/
			//if RedisConnFlag == "1" {
			//	challengeCode, err := redis.String(con.Do("challenge"))
			//	CheckError(err)
			//	origData := []byte("wwwitestcom" + challengeCode)
			//
			//	authbyte, err := RsaEncrypt(origData, PubKeyPath)
			//	encodeString := base64.StdEncoding.EncodeToString(authbyte)
			//	_, err = con.Do("auth", encodeString)
			//	CheckError(err)
			//}
			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	//Set up sc.SubRedisPool
	SubRedisPool = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   8,
		IdleTimeout: 3 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", RedisIP,
				redis.DialPassword(RedisPW),
				redis.DialConnectTimeout(2*time.Second),
				//redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(2*time.Second))

			if err != nil {
				return nil, err

			}
			_, err = con.Do("select", 2)
			CheckError(err)

			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	//初始化写入shell名称列表
	Con := RedisPool.Get()
	_, err := Con.Do("SET", "ShellCfg", sc.ShellList)
	CheckError(err)
	//////处理服务器之前的残余token
	//s, err := redis.Strings(Con.Do("KEYS", "token#*#*"))
	//checkErr(err)
	//for _, v := range s {
	//	_, err := Con.Do("DEL", v)
	//	checkErr(err)
	//}
	Con.Close()
}

