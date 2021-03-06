package actions

import (
	"fmt"

	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/flows/events"
)

// TypeSetContactChannel is the type for the set contact channel action
const TypeSetContactChannel string = "set_contact_channel"

// SetContactChannelAction can be used to update the preferred channel of the current contact.
//
// A `contact_channel_changed` event will be created with the set channel.
//
//   {
//     "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
//     "type": "set_contact_channel",
//     "channel": {"uuid": "4bb288a0-7fca-4da1-abe8-59a593aff648", "name": "FAcebook Channel"}
//   }
//
// @action set_contact_channel
type SetContactChannelAction struct {
	BaseAction
	Channel *flows.ChannelReference `json:"channel"`
}

// Type returns the type of this action
func (a *SetContactChannelAction) Type() string { return TypeSetContactChannel }

// Validate validates our action is valid and has all the assets it needs
func (a *SetContactChannelAction) Validate(assets flows.SessionAssets) error {
	_, err := assets.GetChannel(a.Channel.UUID)
	return err
}

func (a *SetContactChannelAction) Execute(run flows.FlowRun, step flows.Step, log flows.EventLog) error {
	if run.Contact() == nil {
		log.Add(events.NewFatalErrorEvent(fmt.Errorf("can't execute action in session without a contact")))
		return nil
	}

	log.Add(events.NewContactChannelChangedEvent(a.Channel))
	return nil
}
