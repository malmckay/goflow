package expressions

import (
	"fmt"
	"strings"

	"github.com/nyaruka/gocommon/urns"
)

type Resolvable interface {
	Resolve(key string) interface{}
	String() string
}

// ContextTopLevels are the allowed top-level identifiers in legacy expressions, i.e. @contact.bar is valid but @foo.bar isn't
var ContextTopLevels = []string{"channel", "child", "contact", "date", "extra", "flow", "parent", "step"}

// ExtraVarsMapping defines how @extra.* variables should be migrated
type ExtraVarsMapping int

// different ways of mapping @extra in legacy flows
const (
	ExtraAsWebhookJSON ExtraVarsMapping = iota
	ExtraAsTriggerParams
	ExtraAsFunction
)

type varMapper struct {
	// subitems that should be replaced completely with the given strings
	substitutions map[string]string

	// base for fixed subitems, e.g. "contact"
	base string

	// recognized fixed subitems, e.g. "name" or "uuid"
	baseVars map[string]interface{}

	// nesting for arbitrary subitems, e.g. contact fields or run results
	arbitraryNesting string

	// mapper for each arbitrary item
	arbitraryVars map[string]interface{}
}

// returns a copy of this mapper with a prefix applied to the previous base
func (v *varMapper) rebase(prefix string) *varMapper {
	var newBase string
	if prefix != "" {
		newBase = fmt.Sprintf("%s.%s", prefix, v.base)
	} else {
		newBase = v.base
	}
	return &varMapper{
		substitutions:    v.substitutions,
		base:             newBase,
		baseVars:         v.baseVars,
		arbitraryNesting: v.arbitraryNesting,
		arbitraryVars:    v.arbitraryVars,
	}
}

// Resolve resolves the given key to a mapped expression
func (v *varMapper) Resolve(key string) interface{} {
	key = strings.ToLower(key)

	// is this a complete substitution?
	if substitute, ok := v.substitutions[key]; ok {
		return substitute
	}

	newPath := make([]string, 0, 1)

	if v.base != "" {
		newPath = append(newPath, v.base)
	}

	// is it a fixed base item?
	value, ok := v.baseVars[key]
	if ok {
		// subitem may be a mapper itself
		asVarMapper, isVarMapper := value.(*varMapper)
		if isVarMapper {
			if len(newPath) > 0 {
				return asVarMapper.rebase(strings.Join(newPath, "."))
			}
			return asVarMapper
		}

		asExtraMapper, isExtraMapper := value.(*extraMapper)
		if isExtraMapper {
			return asExtraMapper
		}

		// or a simple string in which case we add to the end of the path and return that
		newPath = append(newPath, value.(string))
		return strings.Join(newPath, ".")
	}

	// then it must be an arbitrary item
	if v.arbitraryNesting != "" {
		newPath = append(newPath, v.arbitraryNesting)
	}

	newPath = append(newPath, key)

	if v.arbitraryVars != nil {
		return &varMapper{
			base:     strings.Join(newPath, "."),
			baseVars: v.arbitraryVars,
		}
	}

	return strings.Join(newPath, ".")
}

func (v *varMapper) String() string {
	sub, exists := v.substitutions["__default__"]
	if exists {
		return sub
	}
	return v.base
}

var _ Resolvable = (*varMapper)(nil)

// Migration of @extra requires its own mapper because it can map differently depending on the containing flow
type extraMapper struct {
	varMapper

	path    string
	extraAs ExtraVarsMapping
}

// Resolve resolves the given key to a new expression
func (m *extraMapper) Resolve(key string) interface{} {
	newPath := []string{}
	if m.path != "" {
		newPath = append(newPath, m.path)
	}
	newPath = append(newPath, key)

	if m.path == "" && key == "flow" {
		return &varMapper{
			base: "parent.results",
			arbitraryVars: map[string]interface{}{
				"category": "category_localized",
				"text":     "input",
				"time":     "created_on",
			},
		}
	}

	return &extraMapper{extraAs: m.extraAs, path: strings.Join(newPath, ".")}
}

func (m *extraMapper) String() string {
	switch m.extraAs {
	case ExtraAsWebhookJSON:
		return fmt.Sprintf("run.webhook.json.%s", m.path)
	case ExtraAsTriggerParams:
		return fmt.Sprintf("trigger.params.%s", m.path)
	case ExtraAsFunction:
		return fmt.Sprintf("if(is_error(run.webhook.json.%s), trigger.params.%s, run.webhook.json.%s)", m.path, m.path, m.path)
	}
	return ""
}

var _ Resolvable = (*extraMapper)(nil)

func newMigrationBaseVars() map[string]interface{} {
	contact := &varMapper{
		base: "contact",
		baseVars: map[string]interface{}{
			"uuid":       "uuid",
			"id":         "id",
			"name":       "name",
			"first_name": "first_name",
			"language":   "language",
			"tel_e164":   "urns.tel.0.path",
		},
		substitutions: map[string]string{
			"groups": "join(contact.groups, \",\")",
		},
		arbitraryNesting: "fields",
	}

	for scheme := range urns.ValidSchemes {
		contact.baseVars[scheme] = &varMapper{
			substitutions: map[string]string{
				"__default__": fmt.Sprintf("format_urn(contact.urns.%s)", scheme),
				"display":     fmt.Sprintf("format_urn(contact.urns.%s)", scheme),
				"scheme":      fmt.Sprintf("contact.urns.%s.0.scheme", scheme),
				"path":        fmt.Sprintf("contact.urns.%s.0.path", scheme),
				"urn":         fmt.Sprintf("contact.urns.%s.0", scheme),
			},
			base: fmt.Sprintf("urns.%s", scheme),
		}
	}

	return map[string]interface{}{
		"contact": contact,
		"flow": &varMapper{
			baseVars: map[string]interface{}{
				"contact": contact,
			},
			arbitraryNesting: "run.results",
			arbitraryVars: map[string]interface{}{
				"category": "category_localized",
				"text":     "input",
				"time":     "created_on",
			},
		},
		"parent": &varMapper{
			base: "parent",
			baseVars: map[string]interface{}{
				"contact": contact,
			},
			arbitraryNesting: "results",
			arbitraryVars: map[string]interface{}{
				"category": "category_localized",
			},
		},
		"child": &varMapper{
			base: "child",
			baseVars: map[string]interface{}{
				"contact": contact,
			},
			arbitraryNesting: "results",
			arbitraryVars: map[string]interface{}{
				"category": "category_localized",
			},
		},
		"step": &varMapper{
			substitutions: map[string]string{
				"__default__": "run.input",
				"value":       "run.input",
				"text":        "run.input.text",
				"attachments": "run.input.attachments",
				"time":        "run.input.created_on",
			},
			baseVars: map[string]interface{}{
				"contact": contact,
			},
		},
		"channel": &varMapper{
			substitutions: map[string]string{
				"__default__": "contact.channel.address",
				"name":        "contact.channel.name",
				"tel":         "contact.channel.address",
				"tel_e164":    "contact.channel.address",
			},
		},
		"date": &varMapper{
			substitutions: map[string]string{
				"__default__": `now()`,
				"now":         `now()`,
				"today":       `today()`,
				"tomorrow":    `datetime_add(today(), 1, "D")`,
				"yesterday":   `datetime_add(today(), -1, "D")`,
			},
		},
	}
}

var migrationBaseVars = newMigrationBaseVars()

// creates a new var mapper for migrating expressions
func newMigrationVarMapper(extraAs ExtraVarsMapping) *varMapper {
	// copy the base migration vars
	baseVars := make(map[string]interface{})
	for k, v := range migrationBaseVars {
		baseVars[k] = v
	}

	// add a mapper for extra
	baseVars["extra"] = &extraMapper{extraAs: extraAs}

	return &varMapper{baseVars: baseVars}
}
