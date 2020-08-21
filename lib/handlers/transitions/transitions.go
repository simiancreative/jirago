package transitions

import (
	"jirago/lib/handlers"
	"jirago/lib/handlers/transitions/beginwork"
	"jirago/lib/handlers/transitions/common"
	"jirago/lib/handlers/transitions/toreview"
)

var Actions = handlers.Actions{
	"31":  beginwork.Handler, // begin_work
	"41":  common.Handler,    // to_qa
	"51":  common.Handler,    // to_closed
	"161": toreview.Handler,  // to_review
	"171": common.Handler,    // to_qa
	"191": common.Handler,    // to_todo
	"141": common.Handler,    // to_todo
}
