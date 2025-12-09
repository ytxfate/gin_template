package deployenv

import "errors"

type DeployEnv uint8

const (
	DEV      DeployEnv = 1 << iota // 开发环境
	SIT                            // SIT
	UAT                            // UAT
	PRE_PROD                       // 准生产
	PROD                           // 生产
)

func (e DeployEnv) IsValid() bool {
	switch {
	case e&DEV > 0:
		return true
	case e&SIT > 0:
		return true
	case e&UAT > 0:
		return true
	case e&PRE_PROD > 0:
		return true
	case e&PROD > 0:
		return true
	}
	return false
}

func (e DeployEnv) String() string {
	switch {
	case e&DEV > 0:
		return "DEV"
	case e&SIT > 0:
		return "SIT"
	case e&UAT > 0:
		return "UAT"
	case e&PRE_PROD > 0:
		return "PRE_PROD"
	case e&PROD > 0:
		return "PROD"
	}
	return "UNKNOW"
}

func IsDeployEnv(s string) (DeployEnv, error) {
	switch {
	case s == "DEV":
		return DEV, nil
	case s == "SIT":
		return SIT, nil
	case s == "UAT":
		return UAT, nil
	case s == "PRE_PROD":
		return PRE_PROD, nil
	case s == "PROD":
		return PROD, nil
	}
	return 0, errors.New("unknow enum")
}
