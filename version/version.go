package version

import "fmt"

var (
	BRANCH = "Unknown"

	TAG = "Unknown"

	REVISION = "Unknown"

	BUILDTIME = "Unknown"

	GOVERSION = "Unknown"
)

func String() string {
	return fmt.Sprintf(`-----------------------------------------
-----------------------------------------
Branch:       %v
Tag:          %v
Revision:     %v
Go:           %v
BuildTime:    %v
-----------------------------------------
-----------------------------------------
`, BRANCH, TAG, REVISION, GOVERSION, BUILDTIME)
}
