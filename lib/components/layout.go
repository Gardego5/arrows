package components

import (
	"github.com/Gardego5/arrows/lib"
	. "github.com/Gardego5/htmdsl"
)

func Layout(title string, head []HTMLElement, children ...HTMLElement) HTMLElement {
	return Fragment{DOCTYPE, Html{Attrs{{"lang", "en"}},
		Head{
			Title{title},
			Script{Attrs{{"src", "/static/htmx.min.js"}}},
			head,
			Link{{"rel", "stylesheet"}, {"href", "/static/style.css"}},
		},
		Body{lib.CN("bg-pink-300 dark:bg-pink-800 h-screen"),
			children,
		},
	}}
}
