package main

import (
	"fmt"
	"strings"
)
func ParseCookie(cookie string) (res map[string]string) {
	parts := strings.Split(strings.TrimSpace(cookie), ";")
	res = make(map[string]string, 0)
	for _, item := range parts {
		kv := strings.Split(item, "=")
		if len(kv) == 2 {
			res[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return
}

func testStringInit(){
	str := make([]string, 10)
	fmt.Println(str)
}

func main(){
	cookie := `_ntes_nuid=520ebc7b705970f69b2621a98ed63204; NTESSTUDYSI=2be41dd227aa4c5697d66e21a61b5a8e; EDUWEBDEVICE=29687d7ea7294ff4970873b529e0eb55; eds_utm=eyJjIjoiIiwiY3QiOiIiLCJpIjoiIiwibSI6IiIsInMiOiIiLCJ0IjoiIn0=|aHR0cHM6Ly93d3cuYmFpZHUuY29tL2xpbms/dXJsPTR2LXoySlpMcWdMcndOYVFHdjlPNmpWS2xhNDVaZHlSNHRoLUVuYlhpY1Mmd2Q9JmVxaWQ9YzQyMDA5MTAwMDEzMDFiMjAwMDAwMDAzNWQ0M2Y3ZGE=; hb_MA-BFF5-63705950A31C_source=www.baidu.com; __utmc=129633230; __utmz=129633230.1564735453.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; EDU-YKT-MODULE_GLOBAL_PRIVACY_DIALOG=true; WX_WEBVIEW_AUTH=1; STUDY_MIND_TELBIND=1; NETEASE_WDA_UID=1021363637#|#1477146771955; STUDY_PRIVACY_CONFIRMED=1; hasVolume=true; videoVolume=0.8; _ntes_nnid=96a339d1b867b0d4cd0b5a194fc26d15,1565250783229; __oc_uuid=8cbedf60-b9b1-11e9-a1a3-01f274ac21e9; utm=eyJjIjoiIiwiY3QiOiIiLCJpIjoiIiwibSI6IiIsInMiOiIiLCJ0IjoiIn0=|aHR0cHM6Ly9zdHVkeS4xNjMuY29tL2NvdXJzZS9jb3Vyc2VMZWFybi5odG0/Y291cnNlSWQ9MTIwOTQzMjgyOA==; __utma=129633230.1750060171.1564735453.1564974582.1565250980.6; STUDY_UUID=1658fd9e-70f6-4087-a7e3-ddb9293d4964; STUDY_SESS="r6fKC2yTipcZAutOODlFOLKeOy2pbcvm1mlZ1qv9zRx/qU/e0VAIplVcPbKAIBLHNcIx/vj873KA3Z7hDwu/pkHjJMSAaHCBjWuoWZxEykOlJ3Qn0hSMRqs+0sDW/ktboOPLzRSjMO1WaO3dAK4/7YxuTF1l6Yd9yQiOz0WdefgLhur2Nm2wEb9HcEikV+3FTI8+lZKyHhiycNQo+g+/oA=="; STUDY_INFO=oP4xHuIdpX4Ic1852l-5tDq02ZNM|6|1021363637|1565251098964; NTES_STUDY_YUNXIN_ACCID=s-1021363637; NTES_STUDY_YUNXIN_TOKEN=8e1c306625b6387bd5b78ee3281b58cf; __utmb=129633230.12.7.1565251099640`
	hash := ParseCookie(cookie)
	session, ok := hash["NTESSTUDYSI"]
	if ok == false{
		fmt.Print("why")
	}
	fmt.Println(session)
	testStringInit()
}
