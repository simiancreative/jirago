package handlers

type Params map[string]interface{}

type Run func(
	params Params,
) error

type Handler struct {
	Run
}

type Actions map[string]Handler
