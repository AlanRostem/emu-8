package mu8

// Byte is an alias for uint8
type Byte uint8

// DByte is short for "double-byte"
type DByte uint16

func (db *DByte) AppendNibble(nib uint8) {
	if nib > 0x10 {
		panic("uint8 value is larger than a nibble")
	}
	*db = *db << 4
	*db |= DByte(nib)
}

// Nibble returns the nth nibble where 0 is the left-most nibble.
func (db DByte) Nibble(n uint8) uint8 {
	if n > 3 {
		panic("nibble index too high")
	}
	return uint8((db >> ((3 - n) * 4)) & 0xF)
}

func (db DByte) GetNibbles() [4]uint8 {
	nibbles := [4]uint8{}
	for i := range uint8(4) {
		nibbles[i] = db.Nibble(i)
	}
	return nibbles
}
