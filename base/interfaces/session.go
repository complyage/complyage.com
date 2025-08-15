package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Session Record
//||------------------------------------------------------------------------------------------------||

type SessionRecord struct {
	ID          int64
	Email       string
	Username    string
	Status      string
	Type        string
	Level       int8
	Security    int
	Private     string
	PrivateHash string
	Public      string
	Created     int64
	Expires     int64
	Identity    Identity
}
