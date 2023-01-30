package resources

type Path struct {
	Url           string
	Login         string
	Menu          string
	Lottery       string
	GroundMenu    string
	GroundNumMenu string
	Apply         string
	Complete      string
}

func GetPath(url string) (path Path) {
	path = Path{
		Url:           url,
		Login:         url + "/rsvWUserAttestationAction.do",
		Menu:          url + "/lotWTransLotAcceptListAction.do",
		Lottery:       url + "/lotWTransLotBldGrpAction.do",
		GroundMenu:    url + "/lotWTransLotInstGrpAction.do",
		GroundNumMenu: url + "/lotWTransLotInstSrchVacantAction.do",
		Apply:         url + "/lotWInstTempLotApplyAction.do",
		Complete:      url + "/lotWInstLotApplyAction.do",
	}
	return path
}
