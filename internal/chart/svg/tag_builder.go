package svg

import (
	"fmt"
	"strings"
)

type TagBuilder struct {
	tagName string
	attrs   map[string]string
	content string
}

func NewTag(tagName string) *TagBuilder {
	return &TagBuilder{
		tagName: tagName,
		attrs:   make(map[string]string),
	}
}

func (t *TagBuilder) Attr(name, value string) *TagBuilder {
	t.attrs[name] = value
	return t
}

func (t *TagBuilder) Content(content string) *TagBuilder {
	t.content = content
	return t
}

func (t *TagBuilder) String() string {
	var attrs []string
	for k, v := range t.attrs {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
	}
	attrStr := strings.Join(attrs, " ")
	if t.content == "" {
		return fmt.Sprintf(`<%s %s />`, t.tagName, attrStr)
	}
	return fmt.Sprintf(`<%s %s>%s</%s>`, t.tagName, attrStr, t.content, t.tagName)
}