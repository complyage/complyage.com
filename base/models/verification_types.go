package models

//||------------------------------------------------------------------------------------------------||
//|| VerificationType Model (maps to `verification_types` table)
//||------------------------------------------------------------------------------------------------||

type VerificationType struct {
	IDVerificationType      uint   `gorm:"column:id_verification_type;primaryKey;autoIncrement" json:"id"`
	VerificationCode        string `gorm:"column:verification_code;size:4" json:"code"`
	VerificationDescription string `gorm:"column:verification_description;size:60" json:"description"`
	VerificationLevel       uint8  `gorm:"column:verification_level;default:1" json:"level"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table
//||------------------------------------------------------------------------------------------------||

func (VerificationType) TableName() string {
	return "verification_types"
}
