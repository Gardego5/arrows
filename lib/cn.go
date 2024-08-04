package lib

import (
	html "github.com/Gardego5/htmdsl"
	"github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
)

func CN(args ...string) html.Attr {
	return html.Attr{"class", twmerge.Merge(args...)}
}
