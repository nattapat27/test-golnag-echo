package model

import (
	"encoding/json"
	"time"
)

type User struct {
	tableName struct{} 		`json:"-" pg:"users" bson:"-"`

	Id int					`json:"id" pg:"id,default:DEFAULT"`
	FirstName string		`json:"firstname" pg:"firstname"`
	LastName string			`json:"lastname" pg:"lastname"`
	Age int					`json:"age,omitempty" pg:"age"`
	Relation []*Relation	`json:"relation" pg:"-"`
	CreateTime time.Time	`json:"createtime" pg:"create_time"`
}

type Users []*User

func (u User) String() string {
	bu, _ := json.Marshal(u)
	return  string(bu)
}

func NewUserWithParam(param map[string]interface{}, user *User) *User {
	if user == nil {
		user = new(User)
	}
	if v, ok := param["id"]; ok {
		user.Id = int(v.(float64))
	}
	if v, ok := param["firstname"]; ok {
		user.FirstName = v.(string)
	}
	if v, ok := param["lastname"]; ok {
		user.LastName = v.(string)
	}
	if v, ok := param["age"]; ok {
		user.Age = int(v.(float64))
	}
	user.CreateTime = time.Now()
	for i := range param["relation"].([]interface{}){
		relation := NewRelationWithParam(param["relation"].([]interface{})[i].(map[string]interface{}), nil)
		user.Relation = append(user.Relation, relation)
	}

	return user
}