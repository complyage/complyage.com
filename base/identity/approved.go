package identity

import "strings"

//||------------------------------------------------------------------------------------------------||
//|| Add Approved to the i.Approved
//||------------------------------------------------------------------------------------------------||

func (i *Identity) ApprovedAdd(section string) {
	section = strings.ToUpper(strings.TrimSpace(section))
	for _, v := range i.Approved {
		if v == section {
			return
		}
	}
	i.Approved = append(i.Approved, section)
}

//||------------------------------------------------------------------------------------------------||
//|| ApprovedRemove :: Remove Item from the  i.Approved
//||------------------------------------------------------------------------------------------------||

func (i *Identity) ApprovedRemove(section string) {
	section = strings.ToUpper(strings.TrimSpace(section))
	newList := make([]string, 0, len(i.Approved))
	for _, v := range i.Approved {
		if v != section {
			newList = append(newList, v)
		}
	}
	i.Approved = newList
}
