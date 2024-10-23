//go:build !dev

package swagger

import (
	"embed"
)

//go:embed *.json
var assets embed.FS

func Assets() embed.FS {
	return assets
}
