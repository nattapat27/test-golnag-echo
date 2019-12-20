package model

import "encoding/json"

type Relation struct {
	tableName   struct{} `json:"-" pg:"relations" bson:"-"`
	id          int      `json:"-" pg:"id,default:DEFAULT"`
	Type        string   `json:"type" pg:"type"`
	RelatedName string   `json:"relatename" pg:"related_name"`
	UserId      int      `json:"userid" pg:"user_id"`
}

func (r Relation) String() string {
	bu, _ := json.Marshal(r)
	return string(bu)
}

func NewRelationWithParam(param map[string]interface{}, relation *Relation) *Relation{
	if(relation == nil){
		relation = new(Relation)
	}
	if v, ok := param["type"]; ok{
		relation.Type = v.(string)
	}
	if v, ok := param["relatedname"]; ok{
		relation.RelatedName = v.(string)
	}
	relation.UserId = 0
	return relation
}
