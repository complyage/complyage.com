package identity

import (
	"fmt"
	"time"

	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| VerifyDOB
//||------------------------------------------------------------------------------------------------||

func (i *Identity) UpdateAge(section string, age int, verification string) {
	//||------------------------------------------------------------------------------------------------||
	//|| Generate DOB
	//||------------------------------------------------------------------------------------------------||
	dobTime := time.Now().AddDate(-age, 0, 0)
	dob := types.DOB{
		Day:   dobTime.Day(),
		Month: int(dobTime.Month()),
		Year:  dobTime.Year(),
	}
	displayStr := fmt.Sprintf("%d", age)
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	idRecord := IdentityRecord{
		Verified:     true,
		Age:          age,
		DOB:          dob,
		Display:      displayStr,
		Verification: verification,
	}
	fmt.Println("Updating Age to", age, "for", section, "->", displayStr)
	fmt.Println("DOB:", dob.String())
	fmt.Println("Verification:", verification)
	fmt.Println(displayStr)
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	switch section {
	case "IDEN":
		i.IDCard = idRecord
	case "FACE":
		i.Face = idRecord
	case "CRCD":
		i.CreditCard = idRecord
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Save
	//||------------------------------------------------------------------------------------------------||
	i.Save()
}
