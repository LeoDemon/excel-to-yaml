package main

import (
	"github.com/leodemon/excel-to-yaml/pkg1"
	"log"
)

func main() {
	// reading data from excel file
	alertRows, err1 := pkg1.ReadRules()
	if nil != err1 {
		log.Printf("read file error: [%v]\n", err1)
		return
	}
	// convert data to customized entity
	alerts, err2 := pkg1.ConvertAlertRow(alertRows)
	if nil != err2 {
		log.Printf("convert entity error: [%v]\n", err2)
		return
	}
	// replace file template by customized entity
	for i := 0; i < len(alerts); i++ {
		content, err3 := pkg1.FillTemplate("./templ/alert_rule_tmpl.yml", alerts[i])
		if nil != err3 {
			log.Printf("fill template error: [%v]\n", err3)
			continue
		}
		if len(content) == 0 {
			log.Printf("content data is empty\n")
			continue
		}

		// save data to file
		err4 := pkg1.SaveDataToFile("./output/"+alerts[i].RuleName+".yml", content)
		if err4 != nil {
			log.Printf("save template content error: [%v]\n", err4)
			continue
		}
	}
}
