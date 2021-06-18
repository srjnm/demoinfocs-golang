package mocking

import (
	demoinfocs "github.com/srjnm/demoinfocs-golang/pkg/demoinfocs"
	events "github.com/srjnm/demoinfocs-golang/pkg/demoinfocs/events"
)

func collectKills(parser demoinfocs.Parser) (kills []events.Kill, err error) {
	parser.RegisterEventHandler(func(kill events.Kill) {
		kills = append(kills, kill)
	})
	err = parser.ParseToEnd()
	return
}
