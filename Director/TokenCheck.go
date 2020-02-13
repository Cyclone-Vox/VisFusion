package Director

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func TokenCheck(tk string, r redis.Conn) error {
	if mc, err := redis.String(r.Do("GET", tk)); mc != "" {
		if err != nil {
			return err
		}
		fmt.Println(mc)

		if str1, err := redis.String(r.Do("GET", mc)); str1 != "" {
			//token合法后，刷新mc和tk超时时间
			if err != nil {
				return err
			}
			fmt.Println(str1)

		}
	}
	_, err := r.Do("EXPIRE", tk, 20)
	CheckError(err)
	return nil
}