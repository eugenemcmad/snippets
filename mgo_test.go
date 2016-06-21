package tests

import (
	"fmt"
	"testing"

	"bytes"
	"encoding/json"
	"strings"
	"xr/xutor/model/sites"
	"xr/xutor/model/sources"
	"xr/xutor/mongodb"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ()

var ()

func TestSafeUpdate(t *testing.T) {

	dialInfo := mgo.DialInfo{
		Addrs:          []string{mongodb.HOST_GVM6_X, mongodb.HOST_GVM7_X},
		Direct:         false,
		ReplicaSetName: mongodb.TEST_REPL_SET,
	}

	sess, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
		t.FailNow()
	}
	defer sess.Close()

	var obj sites.Site
	c := sess.DB(mongodb.DbName).C(mongodb.SITES_COLLECTION_NAME)

	err = c.Find(
		bson.M{"_id": 3, "senders._id": 8}).Select(bson.M{"senders": bson.M{"$elemMatch": bson.M{"_id": 8}}}).One(&obj)
	if err != nil {
		fmt.Printf("find err:'%v'\n", err)
		t.FailNow()
	}
	fmt.Printf("%v\n", obj.Senders)

	err = c.Update(
		bson.M{"_id": 3, "senders._id": 8},
		bson.M{"$set": bson.M{"senders.$.activationlimits.others": 11}})
	if err != nil {
		fmt.Printf("update err:'%v'\n", err)
	}

	err = c.Find(
		bson.M{"_id": 3, "senders._id": 8}).Select(bson.M{"senders": bson.M{"$elemMatch": bson.M{"_id": 8}}}).One(&obj)
	if err != nil {
		fmt.Printf("find err:'%v'\n", err)
		t.FailNow()
	}
	fmt.Printf("%v\n", obj.Senders)

	for _, snd := range obj.Senders {
		if snd.Id == 8 {
			snd.ActivationLimits.Others = 9
			err = c.Update(
				bson.M{"_id": 3, "senders._id": 8},
				bson.M{"$set": bson.M{"senders.$.activationlimits": snd.ActivationLimits}})
			if err != nil {
				fmt.Printf("update err:'%v'\n", err)
			}
			break
		}
	}

	err = c.Find(
		bson.M{"_id": 3, "senders._id": 8}).Select(bson.M{"senders": bson.M{"$elemMatch": bson.M{"_id": 8}}}).One(&obj)
	if err != nil {
		fmt.Printf("find err:'%v'\n", err)
		t.FailNow()
	}
	fmt.Printf("%v\n", obj.Senders)

}

func TestDialWithInfo(t *testing.T) {

	dialInfo := mgo.DialInfo{
		Addrs:          []string{mongodb.HOST_41_X, mongodb.HOST_42_X},
		Direct:         false,
		ReplicaSetName: mongodb.DEFAULT_REPL_SET,
	}

	sess, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	var partners []sources.Partner
	err = sess.DB(mongodb.DbName).C(mongodb.PARTNERS_COLLECTION_NAME).Find(nil).Limit(10).All(&partners)
	if err != nil {
		fmt.Printf("read err:'%v'\n", err)
	}

	for _, p := range partners {
		fmt.Printf("%#v\n", p)
	}
}

func TestMgoPartners(t *testing.T) {
	//host := mongodb.HOST_39_X
	host := mongodb.HOST_41_X
	sess, err := mongodb.GetSession(host)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	var partners []sources.Partner
	err = sess.DB(mongodb.DbName).C(mongodb.PARTNERS_COLLECTION_NAME).Find(nil).Limit(10).All(&partners)
	if err != nil {
		fmt.Printf("read err:'%v'\n", err)
	}

	for _, p := range partners {
		fmt.Printf("%#v\n", p)
	}
}

func TestMgoEmails(t *testing.T) {
	sess, err := mongodb.GetSession(mongodb.HOST_39_X)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	var emails []sites.Email
	err = sess.DB(mongodb.DbName).C(mongodb.EmailsCollectionName).Find(nil).Limit(10).All(&emails)
	if err != nil {
		fmt.Printf("read err:'%v'\n", err)
	}
}

func TestMgoEmails3(t *testing.T) {

	sess, err := mongodb.GetSession(mongodb.HOST_39_X)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	var m bson.M
	err = sess.DB(mongodb.DbName).C(mongodb.EmailsCollectionName).Find(bson.M{"_id": 3}).One(&m)
	if err != nil {
		fmt.Printf("read err:'%v'\n", err)
	} else {
		fmt.Printf("'%+v'\n", m)
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&m); err != nil {
		fmt.Printf("buf err:'%v'\n", err)
	} else {
		fmt.Printf("buf.String()=%s\n", buf.String())

		decoded := sites.SentRule{}
		rd := strings.NewReader(buf.String())

		if err := json.NewDecoder(rd).Decode(&decoded); err != nil { // NOTE: ERROR т.к. по дефолту не сохраняются фейковые поля
			fmt.Printf("decode err:'%v'\n", err)
		}

		fmt.Printf("decoded=%#v\n", decoded)
	}

}

func TestMgoInterfacesOld(t *testing.T) {

	colName := "tgs"
	// db.tgs.find()
	// db.tgs.remove({})

	sess, err := mgo.Dial(mongodb.LOCALHOST)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	for i := 0; i < 4; i++ {
		var tt T12
		if i%2 != 0 {
			gg := G1{Id: i, Value: i, Value1: i}
			fmt.Printf("gg=g1=%v\n", gg)

			tt = T12{Id: i, GG: &gg}
			fmt.Printf("T1=%v\n", tt.String())
		} else {
			gg := G2{Id: i, Value: i, Value2: i}
			fmt.Printf("gg=g2=%v\n", gg)

			tt = T12{Id: i, GG: &gg}
			fmt.Printf("T1=%v\n", tt.String())
		}

		info, err := sess.DB(mongodb.DbName).C(colName).UpsertId(i, tt)
		fmt.Printf("update info:'%v'\n", info)
		if err != nil {
			fmt.Printf("update err:'%v'\n", err)
		} else {
			var ttm T12
			err = sess.DB(mongodb.DbName).C(colName).Find(bson.M{"_id": i}).One(&ttm)
			fmt.Printf("find ttm:%v [err:'%v']\n", ttm, err)
		}
	}
}

func TestMgoInterfaces(t *testing.T) {

	colName := "tgs"
	// db.tgs.find()
	// db.tgs.remove({})

	sess, err := mongodb.GetSession(mongodb.HOST_39_X)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	for i := 0; i < 4; i++ {
		var tt T12
		if i%2 != 0 {
			gg := G1{Id: i, Value: i, Value1: i}
			fmt.Printf("gg=g1=%v\n", gg)

			tt = T12{Id: i, GG: &gg}
			fmt.Printf("T1=%v\n", tt.String())
		} else {
			gg := G2{Id: i, Value: i, Value2: i}
			fmt.Printf("gg=g2=%v\n", gg)

			tt = T12{Id: i, GG: &gg}
			fmt.Printf("T1=%v\n", tt.String())
		}

		info, err := sess.DB(mongodb.TEST_DB_NAME).C(colName).UpsertId(i, tt)
		fmt.Printf("update info:'%v'\n", info)
		if err != nil {
			fmt.Printf("update err:'%v'\n", err)
		} else {
			var ttm T12
			err = sess.DB(mongodb.TEST_DB_NAME).C(colName).Find(bson.M{"_id": i}).One(&ttm)
			fmt.Printf("find ttm:%v [err:'%v']\n", ttm, err)
		}
	}
}

func TestMgoInterfacesUnmarsh(t *testing.T) { //TODO BSON TO JSON UNMARSH

	colName := "tgs"
	//> use test
	//> db.tgs.find().pretty()
	//{ "_id" : 0, "gg" : { "_id" : 0, "value" : 0, "value2" : 0 } }
	//{ "_id" : 1, "gg" : { "_id" : 1, "value" : 1, "value1" : 1 } }
	//{ "_id" : 2, "gg" : { "_id" : 2, "value" : 2, "value2" : 2 } }
	//{ "_id" : 3, "gg" : { "_id" : 3, "value" : 3, "value1" : 3 } }
	//> db.tgs.remove({})

	sess, err := mongodb.GetSession(mongodb.HOST_39_X)
	if err != nil {
		fmt.Printf("dial err:'%v'\n", err)
	}
	defer sess.Close()

	for i := 0; i < 4; i++ {

		var ttm T12

		var m bson.M
		err = sess.DB(mongodb.TEST_DB_NAME).C(colName).Find(bson.M{"_id": i}).One(&m)
		if err != nil {
			fmt.Printf("read err:%v\n", err)
		} else {
			fmt.Printf("find:%+v\n", m)
			_s := fmt.Sprintf("%#v", m)
			fmt.Printf("_s:%s\n", _s)
			str := strings.Replace(_s, `bson.M`, ``, 10)
			fmt.Printf("str:%s\n", str)
			err = json.Unmarshal([]byte(str), &ttm)
			fmt.Printf("unmarsh:%+v [err:%v]\n", ttm, err)
		}

		err = sess.DB(mongodb.TEST_DB_NAME).C(colName).Find(bson.M{"_id": i}).One(&ttm)
		fmt.Printf("\nfind ttm:%+v [err:%v]\n\n", ttm, err)
	}
}

type T12 struct {
	Id int  `bson:"_id,omitempty" json:"_id,omitempty"`
	GG IGet `bson:"gg" json:"gg,omitempty"`
	//FF func() string `bson:"-"`
}

func (t *T12) String() string {
	//return fmt.Sprintf("T1: Id=%v, GG=%v, FF=%v", t.Id, t.GG, t.FF)
	return fmt.Sprintf("T1: Id=%v, GG=%v", t.Id, t.GG)
}

type IGet interface {
	Get() int
}

type G1 struct {
	Id     int `bson:"_id,omitempty" json:"_id,omitempty"`
	Value  int
	Value1 int
}

func (g *G1) Get() int {
	return g.Value1
}

func (g *G1) MarshalJSON() ([]byte, error) { // not work
	var j interface{}
	b, _ := bson.Marshal(g)
	bson.Unmarshal(b, &j)
	return json.Marshal(&j)
}

func (g *G1) UnmarshalJSON(b []byte) error { // not work
	var j map[string]interface{}
	json.Unmarshal(b, &j)
	b, _ = bson.Marshal(&j)
	return bson.Unmarshal(b, g)
}

type G2 struct {
	Id     int `bson:"_id,omitempty" json:"_id,omitempty"`
	Value  int
	Value2 int
}

func (g *G2) Get() int {
	return g.Value2
}

func (g *G2) MarshalJSON() ([]byte, error) {
	var j interface{}
	b, _ := bson.Marshal(g)
	bson.Unmarshal(b, &j)
	return json.Marshal(&j)
}

func (g *G2) UnmarshalJSON(b []byte) error {
	var j map[string]interface{}
	json.Unmarshal(b, &j)
	b, _ = bson.Marshal(&j)
	return bson.Unmarshal(b, g)
}

type Forum struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name       string        `bson:",omitempty" json:"name"`
	Slug       string        `bson:",omitempty" json:"slug"`
	Text       string        `bson:",omitempty" json:"text"`
	Moderators []interface{} `bson:",omitempty" json:"moderators"` //userid
}

// /opt/mongodb/bin/mongo
// use xr
// db.partners.find()
// db.partners.remove({})
// db.fieldtypes.find()
// db.fieldtypes.remove({})
