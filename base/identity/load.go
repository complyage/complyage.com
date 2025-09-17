package identity

import (
	"encoding/json"
	"fmt"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| fromJSON - Populates Identity struct from JSON string
//||------------------------------------------------------------------------------------------------||

func Load(id int64) (Identity, error) {
	tmp := Identity{ID: id}
	file := tmp.Filename()
	data, err := app.Storages["identity"].Get(file)
	if err != nil {
		return Identity{}, fmt.Errorf("failed to load identity %s: %w", file, err)
	}
	var identity Identity
	if err := json.Unmarshal(data, &identity); err != nil {
		return Identity{}, fmt.Errorf("failed to parse identity %s: %w", file, err)
	}
	return identity, nil
}
