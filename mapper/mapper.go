package mapper

import (
	"strings"

	"poc-mapper/mapper/kind"
	"poc-mapper/mapper/model"
)

type Mapper struct {
	Mapper    map[string]map[string]Result
	RulesKeys []string
	Rules     map[string][]kind.Kind
}

func NewMapper() *Mapper {
	return &Mapper{
		Mapper: make(map[string]map[string]Result),
		Rules:  make(map[string][]kind.Kind),
	}
}

func (m *Mapper) Add(mapper map[string]Result, rules ...KeyRule) {
	mapperKey := new(strings.Builder)
	gettersKey := new(strings.Builder)
	getters := make([]kind.Kind, 0)

	for _, rule := range rules {
		mapperKey.WriteString(rule.Key)
		gettersKey.WriteString(rule.Kind.Name())
		getters = append(getters, rule.Kind)
	}
	m.Mapper[mapperKey.String()] = mapper
	m.RulesKeys = append(m.RulesKeys, gettersKey.String())
	m.Rules[gettersKey.String()] = getters
}

func (m *Mapper) Get(trx model.Transaction, status string) Result {
	keys := make([]string, 0)

	for _, ruleKey := range m.RulesKeys {
		key := new(strings.Builder)
		for _, r := range m.Rules[ruleKey] {
			key.WriteString(r.Key(trx))
		}
		keys = append(keys, key.String())
	}

	for _, key := range keys {
		if mapper, ok := m.Mapper[key]; ok {
			if result, ok := mapper[status]; ok {
				return result
			}
		}
	}

	return NotMapped
}
