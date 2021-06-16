package daos

import (
	"bot/config"
	"bot/constants"
	"bot/redis"
	"bot/shared"
	"fmt"
	"log"
	"os"

	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongo_Url string = ""
var mongo_DB string = ""

//Daos : ""
type Daos struct {
	mongoURL string
	mongoDB  string
	Shared   *shared.Shared
	Redis    *redis.RedisCli
}

//Init : ""
func init() {
	conf := config.ConfigReader
	fmt.Println("initilising mongodb config with ", conf.GetString(os.Getenv(constants.SOILPROTECTIONENV)))
	log.Println("Mongo URL - " + conf.GetString(os.Getenv(constants.SOILPROTECTIONENV)+".mongo_Url"))
	log.Println("Mongo DB - " + conf.GetString(os.Getenv(constants.SOILPROTECTIONENV)+".mongo_DB"))
	mongo_Url = conf.GetString(os.Getenv(constants.SOILPROTECTIONENV) + ".mongo_Url")
	mongo_DB = conf.GetString(os.Getenv(constants.SOILPROTECTIONENV) + ".mongo_DB")
}

// GetDB :
func GetDB() *mgo.Database {
	fmt.Println("enter main - connecting to mongo")
	maxWait := time.Duration(5 * time.Second)

	s, err := mgo.DialWithTimeout(mongo_Url, maxWait)
	// tried doing this - doesn't work as intended
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Detected panic=>")
			var ok bool
			err, ok := r.(error)
			if !ok {
				fmt.Printf("pkg:  %v,  error: %s", r, err)
			}
		}
	}()
	// s, err := mgo.Dial(mongo_Url)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	} else {
		s.SetMode(mgo.Monotonic, true)
	}

	// fmt.Println("Connected to", mongo_Url)
	defer s.Close()
	databaseName := mongo_DB
	if len(databaseName) == 0 {
		databaseName = mongo_DB
	}
	return s.Copy().DB(databaseName)
}

// GetDBV2 : ""
func GetDBV2() *mgo.Database {
	s, err := mgo.Dial(mongo_Url)
	// s, err := mgo.DialWithTimeout(mongo_Url, maxWait)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	// fmt.Println("Connected to", mongo_Url)
	defer s.Close()
	databaseName := mongo_DB
	if len(databaseName) == 0 {
		databaseName = mongo_DB
	}
	return s.Copy().DB(databaseName)
}

//GetDaos : ""
func GetDaos(s *shared.Shared, Redis *redis.RedisCli) *Daos {
	conf := config.ConfigReader
	fmt.Println(s.GetCmdArg(constants.ENV) + ".mongo_Url")
	return &Daos{conf.GetString(s.GetCmdArg(constants.ENV) + ".mongo_Url"),
		conf.GetString(s.GetCmdArg(constants.ENV) + ".mongo_DB"),
		s,
		Redis,
	}
}

// GetDB :
func (d *Daos) GetDB() *mgo.Database {
	s, err := mgo.Dial(d.GetDBURL())
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	fmt.Println("Connected to", d.GetDBURL())
	defer s.Close()
	databaseName := d.GetDBName()
	if len(databaseName) == 0 {
		databaseName = d.GetDBName()
	}
	return s.Copy().DB(databaseName)
}

//GetDBURL : ""
func (d *Daos) GetDBURL() string {
	return d.mongoURL
}

//GetDBName : ""
func (d *Daos) GetDBName() string {
	return d.mongoDB
}

// GenerateUniqueID : this function returns next value
func (d *Daos) GenerateUniqueID(key string) string {
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
		return ""
	}
	var collReg CollectionRegistory
	if err := db.C("collectionregister").Find(bson.M{"code": key}).One(&collReg); err != nil {
		fmt.Println("reg error : ", err.Error())
		return ""
	}
	dig := fmt.Sprintf("%dd", collReg.Digits)
	str := "%v%0" + dig + "%v"
	fmt.Println(dig, str)
	return fmt.Sprintf(str, collReg.Prefix, result.Value, collReg.SUffix)
}

//CollectionRegistory : ""
type CollectionRegistory struct {
	Code   string `json:"code" bson:"code,omitempty"`
	SUffix string `json:"suffix" bson:"suffix,omitempty"`
	Prefix string `json:"prefix" bson:"prefix,omitempty"`
	Digits int    `json:"digits" bson:"digits,omitempty"`
}
