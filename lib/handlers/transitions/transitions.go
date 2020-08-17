package transitions

import (
	"jirago/lib/handlers"
	"jirago/lib/handlers/transitions/beginwork"
)

var Actions = handlers.Actions{
	"31":  beginwork.Handler, // beginwork
	"41":  beginwork.Handler, // toqa
	"51":  beginwork.Handler, // toclosed
	"161": beginwork.Handler, // toreview
	"171": beginwork.Handler, // toqa
	"191": beginwork.Handler, // totodo
}
