package CfgLoad

import (
	"VisFusion/RegStruct"
	"gopkg.in/ini.v1"
	"log"
)

var (
	CfgFile *ini.File
	err     error
)

func LoadCfg() *RegStruct.HttpsReg{

	//ShadowLoad And AllowShadows Load
	CfgFile, err = ini.LoadSources(
		ini.LoadOptions{
			AllowBooleanKeys: true,
			AllowShadows:     true,
		}, "./config.ini")
	checkErr(err)
	//make conf loader
	return cfgLoader(CfgFile)

}



func cfgLoader(cfgFile *ini.File) *RegStruct.HttpsReg {
	// MapTo Configurations
	sc := new(RegStruct.HttpsReg)
	err := cfgFile.MapTo(sc)
	checkErr(err)
	return sc
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
