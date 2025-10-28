package encrypted

func (e Encrypted) Filename() string {
	if e.UUID == "" {
		panic("UUID is required to generate filename")
	}
	return e.UUID + ".bin"
}
