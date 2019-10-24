package models

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"log"
)

/*BomberCharacter - identification model bomber for identication in multicast discovery*/
type BomberCharacter struct {
	ID string `json:"bomber_id"`
}

/*InitializeBomber - setup new uuid for current bomber*/
func InitializeBomber() *BomberCharacter {
	return &BomberCharacter{
		ID: generateUUID(),
	}
}

/*GetBytes - getting bytes for sending in multicast layer*/
func (bomb *BomberCharacter) GetBytes() []byte {
	var encode bytes.Buffer
	enc := gob.NewEncoder(&encode)
	if err := enc.Encode(bomb); err != nil {
		return nil
	}
	return encode.Bytes()
}

func GetStructureFromBytes(input []byte) *BomberCharacter {
	decode := bytes.NewBuffer(input)
	dec := gob.NewDecoder(decode)
	var bomber BomberCharacter
	if err := dec.Decode(&bomber); err != nil {
		return nil
	}
	return &bomber
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
