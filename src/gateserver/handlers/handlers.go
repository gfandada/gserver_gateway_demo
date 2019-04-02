package handlers

import (
	Services "github.com/gfandada/gserver/services"
)

func NewHandlers() {
	Services.Register(uint16(500), TouristReqHandler)
	Services.Register(uint16(502), TouristLoginReqHandler)
}
