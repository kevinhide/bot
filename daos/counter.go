package daos

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionCounter = "counters"
)

type counter struct {
	Key   string `json:"key,omitempty" bson:"key,omitempty" form:"key"`
	Value int64  `json:"value,omitempty" bson:"value,omitempty" form:"value"`
}

// Next : this function returns next value
func (d *Daos) Next(key string) int64 {
	// Instance of DB
	// INFO: dependent on mongo package
	db := d.GetDB()
	defer db.Session.Close()
	result := counter{}
	if _, err := db.C(collectionCounter).Find(bson.M{"key": key}).Apply(mgo.Change{
		Update:    bson.M{"$set": bson.M{"key": key}, "$inc": bson.M{"value": 1}},
		Upsert:    true,
		ReturnNew: true,
	}, &result); err != nil {
		fmt.Println("Autoincrement error(1) : ", err.Error())
		return 0
	}
	return result.Value
}
