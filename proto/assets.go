//go:build !dev
// +build !dev

package proto

import "embed"

//go:embed *.proto google/* google/api/*.proto google/rpc/*.proto
var assets embed.FS

func Assets() embed.FS {
	return assets
}
