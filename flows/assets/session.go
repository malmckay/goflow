package assets

import (
	"fmt"

	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/utils"
)

// our implementation of SessionAssets - the high-level API for asset access from the engine
type sessionAssets struct {
	cache  *AssetCache
	server AssetServer
}

var _ flows.SessionAssets = (*sessionAssets)(nil)

// NewSessionAssets creates a new session assets instance with the provided base URLs
func NewSessionAssets(cache *AssetCache, server AssetServer) flows.SessionAssets {
	return &sessionAssets{cache: cache, server: server}
}

// HasLocations returns whether locations are supported as an asset item type
func (s *sessionAssets) HasLocations() bool {
	return s.server.isTypeSupported(assetTypeLocationHierarchy)
}

// GetLocationHierarchy gets the location hierarchy asset for the session
func (s *sessionAssets) GetLocationHierarchy() (*utils.LocationHierarchy, error) {
	asset, err := s.cache.GetAsset(s.server, assetTypeLocationHierarchy, "")
	if err != nil {
		return nil, err
	}
	hierarchy, isType := asset.(*utils.LocationHierarchy)
	if !isType {
		return nil, fmt.Errorf("asset cache contains asset with wrong type")
	}
	return hierarchy, nil
}

// GetChannel gets a channel asset for the session
func (s *sessionAssets) GetChannel(uuid flows.ChannelUUID) (flows.Channel, error) {
	channels, err := s.GetChannelSet()
	if err != nil {
		return nil, err
	}
	channel := channels.FindByUUID(uuid)
	if channel == nil {
		return nil, fmt.Errorf("no such channel with uuid '%s'", uuid)
	}
	return channel, nil
}

// GetChannelSet gets the set of all channels asset for the session
func (s *sessionAssets) GetChannelSet() (*flows.ChannelSet, error) {
	asset, err := s.cache.GetAsset(s.server, assetTypeChannelSet, "")
	if err != nil {
		return nil, err
	}
	channels, isType := asset.(*flows.ChannelSet)
	if !isType {
		return nil, fmt.Errorf("asset cache contains asset with wrong type")
	}
	return channels, nil
}

// GetField gets a contact field asset for the session
func (s *sessionAssets) GetField(key string) (*flows.Field, error) {
	fields, err := s.GetFieldSet()
	if err != nil {
		return nil, err
	}
	field := fields.FindByKey(key)
	if field == nil {
		return nil, fmt.Errorf("no such field with key '%s'", key)
	}
	return field, nil
}

// GetFieldSet gets the set of all fields asset for the session
func (s *sessionAssets) GetFieldSet() (*flows.FieldSet, error) {
	asset, err := s.cache.GetAsset(s.server, assetTypeFieldSet, "")
	if err != nil {
		return nil, err
	}
	fields, isType := asset.(*flows.FieldSet)
	if !isType {
		return nil, fmt.Errorf("asset cache contains asset with wrong type")
	}
	return fields, nil
}

// GetFlow gets a flow asset for the session
func (s *sessionAssets) GetFlow(uuid flows.FlowUUID) (flows.Flow, error) {
	asset, err := s.cache.GetAsset(s.server, assetTypeFlow, string(uuid))
	if err != nil {
		return nil, err
	}
	flow, isType := asset.(flows.Flow)
	if !isType {
		return nil, fmt.Errorf("asset cache contains asset with wrong type for UUID '%s'", uuid)
	}
	return flow, nil
}

// GetGroup gets a contact group asset for the session
func (s *sessionAssets) GetGroup(uuid flows.GroupUUID) (*flows.Group, error) {
	groups, err := s.GetGroupSet()
	if err != nil {
		return nil, err
	}
	group := groups.FindByUUID(uuid)
	if group == nil {
		return nil, fmt.Errorf("no such group with uuid '%s'", uuid)
	}
	return group, nil
}

// GetGroupSet gets the set of all groups asset for the session
func (s *sessionAssets) GetGroupSet() (*flows.GroupSet, error) {
	asset, err := s.cache.GetAsset(s.server, assetTypeGroupSet, "")
	if err != nil {
		return nil, err
	}
	groups, isType := asset.(*flows.GroupSet)
	if !isType {
		return nil, fmt.Errorf("asset cache contains asset with wrong type")
	}
	return groups, nil
}

// GetLabel gets a message label asset for the session
func (s *sessionAssets) GetLabel(uuid flows.LabelUUID) (*flows.Label, error) {
	labels, err := s.GetLabelSet()
	if err != nil {
		return nil, err
	}
	label := labels.FindByUUID(uuid)
	if label == nil {
		return nil, fmt.Errorf("no such label with uuid '%s'", uuid)
	}
	return label, nil
}

func (s *sessionAssets) GetLabelSet() (*flows.LabelSet, error) {
	asset, err := s.cache.GetAsset(s.server, assetTypeLabelSet, "")
	if err != nil {
		return nil, err
	}
	labels, isType := asset.(*flows.LabelSet)
	if !isType {
		return nil, fmt.Errorf("asset cache contains asset with wrong type")
	}
	return labels, nil
}
