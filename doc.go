// Package ccmaxclaudedist is a distribution shim: it carries prebuilt ccmax-claude
// wrapper binaries so they can be pulled through the Go module proxy (goproxy.cn),
// which is China-fast (~34 MB/s to Shenzhen) unlike GitHub release assets (GFW-throttled).
// There is no importable code here.
package ccmaxclaudedist
