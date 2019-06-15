package kubeless

import (
	"github.com/kubeless/kubeless/pkg/functions"
)

func Handler(event functions.Event, context functions.Context) (string, error) {
	return "Hello go!", nil
}
