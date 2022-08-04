package provider

import "strings"

// TL;DR: this is a help package, to hide the implementation details
// this package is intended to provide simple interface to the underlying
// map of maps, and to set them without excessive checks.

// NormalizedTagParam
// TODO: consider removing redundant fields (tagName, tagKey)
// TODO: consider adding modifiers list, to reuse later on population
type NormalizedTagParam struct {
	TagName string
	TagKey  string
	Values  []string
}

// TagKeyProvider
// tagKey: {name, key, values}
type TagKeyProvider map[string]NormalizedTagParam

// TagProvider
// tagName:
//   tagKey: {name, key, values}
type TagProvider map[string]TagKeyProvider

func NewTagKeyProvider() TagKeyProvider {
	return make(TagKeyProvider)
}

func (p TagKeyProvider) Get(tagKey string) []string {
	if _, ok := p[tagKey]; !ok {
		return nil
	}
	return p[tagKey].Values
}

func (p TagKeyProvider) Set(tagParam NormalizedTagParam) {
	p[tagParam.TagKey] = tagParam
}

func (p TagKeyProvider) Has(tagKey string) bool {
	_, has := p[tagKey]
	return has
}

func New() TagProvider {
	return make(TagProvider)
}

func (p TagProvider) Get(tagName, tagKey string) []string {
	if _, ok := p[tagName]; !ok {
		return nil
	}
	return p[tagName].Get(tagKey)
}

// Set TODO: make a note
// that supported only either filter[ids]=1,2,3 or filter[ids]=1&filter[ids]=2&filter[ids]=3
// filter[ids]=1&filter[ids]=2,3 NOT supported
func (p TagProvider) Set(tagName, tagKey string, values []string) {
	if _, ok := p[tagName]; !ok {
		p[tagName] = NewTagKeyProvider()
	}

	tagValues := values
	if len(values) == 1 {
		tagValues = strings.Split(values[0], ",")
	}
	p[tagName].Set(NormalizedTagParam{
		TagName: tagName,
		TagKey:  tagKey,
		Values:  tagValues,
	})
}

func (p TagProvider) Has(tagName, tagKey string) bool {
	_, has := p[tagName]
	return has && p[tagName].Has(tagKey)
}

func (p TagProvider) HasTag(tagName string) bool {
	_, has := p[tagName]
	return has
}

func (p TagProvider) Merge(other TagProvider) {
	for tagName, tagKeyProvider := range other {
		if !p.HasTag(tagName) {
			p[tagName] = other[tagName]
			continue
		}
		for tagKey, tagParam := range tagKeyProvider {
			if !p.Has(tagName, tagKey) {
				p[tagName][tagKey] = tagParam
			}
			if len(p.Get(tagName, tagKey)) == 0 {
				p.Set(tagName, tagKey, tagParam.Values)
			}
		}
	}
}
