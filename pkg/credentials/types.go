package credentials

import (
	"github.com/gofrs/uuid"
	"log"
)

type KeyValue struct {
	Key   string
	Value string // any value is transformed to a string for readability and handling
}
type InfoEntry struct {
	uid    uuid.UUID
	values []KeyValue
}

func NewInfoEntry(values []KeyValue) *InfoEntry {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return &InfoEntry{uid, values}
}

//// Create a Version 4 UUID, panicking on error.
//// Use this form to initialize package-level variables.
//var u1 = uuid.Must(uuid.NewV4())
//
//func main() {
//	// Create a Version 4 UUID.
//	u2, err := uuid.NewV4()
//	if err != nil {
//		log.Fatalf("failed to generate UUID: %v", err)
//	}
//	log.Printf("generated Version 4 UUID %v", u2)
//
//	// Parse a UUID from a string.
//	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
//	u3, err := uuid.FromString(s)
//	if err != nil {
//		log.Fatalf("failed to parse UUID %q: %v", s, err)
//	}
//	log.Printf("successfully parsed UUID %v", u3)
//}
