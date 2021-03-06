package events

import "github.com/nyaruka/goflow/flows"

// TypeInputLabelsAdded is the type of our add label action
const TypeInputLabelsAdded string = "input_labels_added"

// InputLabelsAddedEvent events will be created with the labels that were applied to the given input.
//
//   {
//     "type": "input_labels_added",
//     "created_on": "2006-01-02T15:04:05Z",
//     "input_uuid": "4aef4050-1895-4c80-999a-70368317a4f5",
//     "labels": [{"uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d", "name": "Spam"}]
//   }
//
// @event input_labels_added
type InputLabelsAddedEvent struct {
	baseEvent
	callerOrEngineEvent

	InputUUID flows.InputUUID         `json:"input_uuid" validate:"required,uuid4"`
	Labels    []*flows.LabelReference `json:"labels" validate:"required,min=1,dive"`
}

// NewInputLabelsAddedEvent returns a new add to group event
func NewInputLabelsAddedEvent(inputUUID flows.InputUUID, labels []*flows.LabelReference) *InputLabelsAddedEvent {
	return &InputLabelsAddedEvent{
		baseEvent: newBaseEvent(),
		InputUUID: inputUUID,
		Labels:    labels,
	}
}

// Type returns the type of this event
func (e *InputLabelsAddedEvent) Type() string { return TypeInputLabelsAdded }

// Validate validates our event is valid and has all the assets it needs
func (e *InputLabelsAddedEvent) Validate(assets flows.SessionAssets) error {
	for _, label := range e.Labels {
		_, err := assets.GetLabel(label.UUID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Apply applies this event to the given run
func (e *InputLabelsAddedEvent) Apply(run flows.FlowRun) error {
	return nil
}
