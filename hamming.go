package hamming

import (
	"errors"
	"strings"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	aReader := strings.NewReader(a)
	bReader := strings.NewReader(b)
	var distance int
	err := ValidateTwoString(aReader, bReader)
	if err == nil {
		distance = CaculateDistance(aReader, bReader)
	}
	return distance, err
}
func ValidateTwoString(aReader, bReader *strings.Reader) error {
	var err error
	if aReader.Size() != bReader.Size() {
		err = errors.New("Unequal length!")
	}
	return err
}
func CaculateDistance(aReader, bReader *strings.Reader) int {
	distance := 0
	for i := 0; i < int(aReader.Size()); i++ {
		elementA, _ := aReader.ReadByte()
		elementB, _ := bReader.ReadByte()
		if elementA != elementB {
			distance += 1
		}
	}
	return distance
}
