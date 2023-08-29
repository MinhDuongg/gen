package ulti

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"os"
)

func PathConcat(destination string, name string) (string, error) {
	return fmt.Sprintf("%s%s%s", destination, string(os.PathSeparator), name), nil
}

func InArray[T any](array []T, search T) bool {
	for _, v := range array {
		if cmp.Equal(v, search) {
			return true
		}
	}

	return false
}

// THIS SHOULD BE A BACKUP MECHANISM ONLY
func GenerateFileName() (string, error) {
	return fmt.Sprintf("%s_GENERATED", uuid.NewString()), nil
}
