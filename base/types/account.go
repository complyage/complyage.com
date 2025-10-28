package types

//||------------------------------------------------------------------------------------------------||
//|| Account Session
//||------------------------------------------------------------------------------------------------||

type AccountSessionChecked struct {
	ID         int64
	Salt       string
	Username   string
	Identifier string
	Status     string
	Level      int
	KeysLoaded bool
	Private    string
	Public     string
	CheckKey   string
}
