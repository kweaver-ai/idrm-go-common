package entity_change

import "gorm.io/gorm/schema"

type EntityGraphMap map[string][]string

type GormModel schema.Tabler

func NewEntityGraphMap() EntityGraphMap {
	return make(map[string][]string)
}

func (e EntityGraphMap) Record(graphName string, ts []GormModel) {
	for _, t := range ts {
		e[t.TableName()] = append(e[t.TableName()], graphName)
	}
}
func (e EntityGraphMap) DestGraph(modelName string) []string {
	graphNames, ok := e[modelName]
	if !ok {
		return []string{}
	}
	return graphNames
}
