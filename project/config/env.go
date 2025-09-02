package config

type deployEnv uint8

const (
	DEV      deployEnv = 1 << iota // 开发环境
	SIT                            // SIT
	UAT                            // UAT
	PRE_PROD                       // 准生产
	PROD                           // 生产
)

func (e deployEnv) IsValid() bool {
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

func (e deployEnv) String() string {
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
