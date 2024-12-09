package sockets

var GloabConfig WxMiniApp

type WxMiniApp map[string]string

func (w WxMiniApp) GetKeys() []string {
	return []string{"#601919"}
}
func init() {
	GloabConfig = WxMiniApp{
		"#601919": "#小程序://腾讯微证券/T5VHtrDV0EMTa4k",
	}
}
