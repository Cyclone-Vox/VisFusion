package Director

import (
	"bufio"
	"fmt"
	//"github.com/spaolacci/murmur3"
	"github.com/valyala/fasthttp"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"os/exec"
)

func licgen(ctx *fasthttp.RequestCtx) {
	//println("here!!!")
	key := ctx.Request.Body()
	fmt.Println(string(ctx.Path()), string(ctx.Request.Header.Peek("token")), string(key))
	//建立自定义license文件用于外部程序的执行参数
	randNum := rand.Int()
	f, err := os.Create("./data/" + strconv.Itoa(randNum) + ".lic")
	CheckError(err)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(key))
	CheckError(err)
	err = w.Flush()
	CheckError(err)
	f.Close()

	//载入license外部程序
	cmd := exec.Command("./licgen/licgen", "-i", "./data/"+strconv.Itoa(randNum)+".lic")
	//fmt.Printf("%+v\n",cmd)
	output, err := cmd.CombinedOutput()
	CheckError(err)
	MsgSplit := strings.Split(string(output), "\n")
	fmt.Println("license:", MsgSplit[0])

	//运行结束后删除外部程序
	err = os.Remove("./data/" + strconv.Itoa(randNum) + ".lic")
	CheckError(err)
	//ps, err := os.StartProcess("./data/licgen", []string{string(key)}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	//sp,err:=os.StartProcess("./data/licgen",keystr,nil)

	ctx.Response.SetBody([]byte(MsgSplit[0]))
	ctx.Request.ConnectionClose()
	ctx.SetConnectionClose()
	return
}
