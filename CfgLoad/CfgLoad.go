package CfgLoad

import (
	"VisFusion/RegStruct"
	"encoding/json"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"strings"
)

var (
	CfgFile *ini.File
	err     error
)

func LoadCfg() *Cfg{

	//ShadowLoad And AllowShadows Load
	CfgFile, err = ini.LoadSources(
		ini.LoadOptions{
			AllowBooleanKeys: true,
			AllowShadows:     true,
		}, "./config/config.ini")
	checkErr(err)
	//make conf loader
	return cfgLoader(CfgFile)

}



func cfgLoader(cfgFile *ini.File) *Cfg {
	// MapTo Configurations
	sc := new(Cfg)
	err := cfgFile.MapTo(sc)
	sc.shellMapper()
	checkErr(err)
	return sc
}

func (c *Cfg)shellMapper()  {
	c.DevCfgMap=make(map[string]RegStruct.HttpsReg)
	for _,v:=range c.DevDefaultCfg {
		//DevDefaultCfg:=flow-./config/flow
		//s[0]->name s[1]->path
		var r RegStruct.HttpsReg
		s:=strings.Split(v,"-")
		b,err:=ioutil.ReadFile(s[1])
		checkErr(err)
		err=json.Unmarshal(b,&r)
		checkErr(err)
		c.DevCfgMap[s[0]]=r
	}
}


func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
