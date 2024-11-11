package centrifuge

import "fmt"

func getUid(userId, platform string) string {
	return fmt.Sprintf("%s:%s", userId, platform)
}
