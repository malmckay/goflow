package main

import (
	"bytes"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/test"
	"github.com/nyaruka/goflow/utils"
)

var docSets = []struct {
	contextKey string
	searchDirs []string
	tag        string
	handler    handleFunc
}{
	{"contextDocs", []string{"flows"}, "@context", handleContextDoc},
	{"functionDocs", []string{"excellent/functions"}, "@function", handleFunctionDoc},
	{"testDocs", []string{"flows/routers/tests"}, "@test", handleFunctionDoc},
	{"actionDocs", []string{"flows/actions"}, "@action", handleActionDoc},
	{"eventDocs", []string{"flows/events"}, "@event", handleEventDoc},
}

var testServerPort = 49998

type documentedItem struct {
	typeName    string   // actual go type name
	tagName     string   // tag used to make this as a documented item
	tagValue    string   // value after @tag
	examples    []string // any indented line
	description []string // any other line
}

type handleFunc func(output *strings.Builder, item *documentedItem, session flows.Session) error

// builds all documentation from the given base directory
func buildDocs(baseDir string) (string, error) {
	fmt.Println("Generating docs...")

	server, err := test.NewTestHTTPServer(testServerPort)
	if err != nil {
		return "", fmt.Errorf("error starting mock HTTP server: %s", err)
	}
	defer server.Close()

	utils.SetRand(utils.NewSeededRand(123456))
	defer utils.SetRand(utils.DefaultRand)

	utils.SetUUIDGenerator(utils.NewSeededUUID4Generator(123456))
	defer utils.SetUUIDGenerator(utils.DefaultUUIDGenerator)

	session, err := test.CreateTestSession(testServerPort, nil)
	if err != nil {
		return "", fmt.Errorf("error creating example session: %s", err)
	}

	context := make(map[string]string)

	for _, ds := range docSets {
		if context[ds.contextKey], err = buildDocSet(baseDir, ds.searchDirs, ds.tag, ds.handler, session); err != nil {
			return "", err
		}
	}

	// generate our complete docs
	docTpl, err := template.ParseFiles(path.Join(baseDir, "cmd/docgen/templates/docs.md"))
	if err != nil {
		return "", fmt.Errorf("Error reading template file: %s", err)
	}

	output := bytes.Buffer{}
	err = docTpl.Execute(&output, context)
	if err != nil {
		return "", fmt.Errorf("Error executing template: %s", err)
	}

	return output.String(), nil
}

func buildDocSet(baseDir string, searchDirs []string, tag string, handler handleFunc, session flows.Session) (string, error) {
	items := make([]*documentedItem, 0)
	for _, searchDir := range searchDirs {
		fromDir, err := findDocumentedItems(baseDir, searchDir, tag)
		if err != nil {
			return "", err
		}
		items = append(items, fromDir...)
	}

	// sort documented items by their tag value
	sort.SliceStable(items, func(i, j int) bool { return items[i].tagValue < items[j].tagValue })

	fmt.Printf(" > found %d documented items with tag %s\n", len(items), tag)

	buffer := &strings.Builder{}

	for _, item := range items {
		if err := handler(buffer, item, session); err != nil {
			return "", fmt.Errorf("error parsing %s:%s: %s", item.tagName, item.tagValue, err)
		}
	}

	return buffer.String(), nil
}

// finds all documented items in go files in the given directory
func findDocumentedItems(baseDir string, searchDir string, tag string) ([]*documentedItem, error) {
	documentedItems := make([]*documentedItem, 0)

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path.Join(baseDir, searchDir), nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, f := range pkgs {
		p := doc.New(f, "./", 0)
		for _, t := range p.Types {
			if strings.Contains(t.Doc, tag) {
				documentedItems = append(documentedItems, parseDocString(tag, t.Doc, t.Name))
			}
		}
		for _, t := range p.Funcs {
			if strings.Contains(t.Doc, tag) {
				documentedItems = append(documentedItems, parseDocString(tag, t.Doc, t.Name))
			}
		}
	}

	return documentedItems, nil
}

func parseDocString(tag string, docString string, typeName string) *documentedItem {
	var tagValue string
	examples := make([]string, 0)
	description := make([]string, 0)

	docString = removeTypeNamePrefix(docString, typeName)

	for _, l := range strings.Split(docString, "\n") {
		trimmed := strings.TrimSpace(l)

		if strings.HasPrefix(l, tag) {
			tagValue = l[len(tag)+1:]
		} else if strings.HasPrefix(l, "  ") { // examples are indented by at least two spaces
			examples = append(examples, trimmed)
		} else {
			description = append(description, l)
		}
	}

	return &documentedItem{typeName: typeName, tagName: tag[1:], tagValue: tagValue, examples: examples, description: description}
}

// Golang convention is to start all docstrings with the type name, but the actual type name can differ from how the type is
// referenced in the flow spec, so we remove it.
func removeTypeNamePrefix(docString string, typeName string) string {
	// remove type name from start of description and capitalize the next word
	if strings.HasPrefix(docString, typeName) {
		docString = strings.Replace(docString, typeName, "", 1)
		docString = strings.TrimSpace(docString)
		docString = strings.ToUpper(docString[0:1]) + docString[1:]
	}
	return docString
}
