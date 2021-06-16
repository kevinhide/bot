package daos

import (
	"bot/constants"
	"bot/models"
	"bot/shared"
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//Category : ""
func (d *Daos) Category(isParent string) ([]models.Category, error) {
	db := d.GetDB()
	defer db.Session.Close()

	var query []bson.M

	data := db.C(constants.COLLECTIONBOT)

	val := new(models.Category)

	val.IsParent = isParent

	var category []models.Category

	if val.IsParent == constants.ISPARENT {
		if val != nil {
			if len(val.IsParent) > 0 {
				fmt.Println("False Values >>>>>>>>>>>>>>>>>", val.IsParent)
				query = append(query, bson.M{"$find": bson.M{"isParent": bson.M{"$eq": val.IsParent}}})

			} else {
				errors.New("not found")
			}
		}
	}
	if val.IsParent == constants.ISPARENTS {
		if val != nil {
			if len(val.IsParent) > 0 {
				fmt.Println("False Values >>>>>>>>>>>>>>>>>", val.IsParent)
				query = append(query, bson.M{"$find": bson.M{"isParent": bson.M{"$eq": val.IsParent}}})

			} else {
				errors.New("not found")
			}
		}
	}

	shared.BsonToJSONPrint(query)

	err := data.Find(bson.M{"query": query}).All(&category)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return category, nil
}
