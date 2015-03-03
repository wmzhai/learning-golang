package models

import (
_	"errors"
	"fmt"
_	"reflect"

	"time"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/astaxie/beego"
)

type Hospitals struct {
	Name            string    `orm:"column(name);size(255);null"`
	SpellCode       string    `orm:"column(spell_code);size(255);null"`
	ShortName       string    `orm:"column(short_name);size(255);null"`
	Address         string    `orm:"column(address);size(255);null"`
	Phone           string    `orm:"column(phone);size(255);null"`
	Description     string    `orm:"column(description);null"`
	Rank            string    `orm:"column(rank);size(255);null"`
	Province        string    `orm:"column(province);size(255);null"`
	City            string    `orm:"column(city);size(255);null"`
	KeyDepartments  string    `orm:"column(key_departments);size(255);null"`
	OperationMode   string    `orm:"column(operation_mode);size(255);null"`
	Email           string    `orm:"column(email);size(255);null"`
	Website         string    `orm:"column(website);size(255);null"`
	Fax             string    `orm:"column(fax);size(255);null"`
	Ids             string    `orm:"column(ids);size(255);null"`
	ShortSpell      string    `orm:"column(short_spell);size(255);null"`
	Longitude       string    `orm:"column(longitude);size(255);null"`
	Latitude        string    `orm:"column(latitude);size(255);null"`
	Area            string    `orm:"column(area);size(255);null"`
	DepartmentCount int       `orm:"column(department_count);null"`
	DoctorCount     int       `orm:"column(doctor_count);null"`
	CreatedAt       time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(datetime);auto_now"`
}

// AddHospitals insert a new Hospitals into database and returns
// last inserted Id on success.
func AddHospitals(m *Hospitals) (v *Hospitals,err error) {
	session, err := mgo.Dial(beego.AppConfig.String("mongoDb_addr"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gMon").C("hospitals")
	err = c.Insert(&m)
	if err != nil {
		panic(err)
	}
	return m,err
}

// GetHospitalsById retrieves Hospitals by Id. Returns error if
// Id doesn't exist
func GetHospitalsById(id string) (v *Hospitals,err error) {
	session, err := mgo.Dial(beego.AppConfig.String("mongoDb_addr"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gMon").C("hospitals")
	v = new(Hospitals)
	objectId := bson.ObjectIdHex(id)
	c.FindId(objectId).One(&v)

	return v,err
}

// GetAllHospitals retrieves all Hospitals matches certain condition. Returns empty list if
// no records exist
func GetAllHospitals(query map[string]string, sortby string, limit int) (ml []interface{}, err error) {

	session, err := mgo.Dial(beego.AppConfig.String("mongoDb_addr"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gMon").C("hospitals")
	if sortby != "" {
	c.Find(query).Sort(sortby).Limit(limit).All(&ml)
	} else {
	c.Find(query).Limit(limit).All(&ml)
	}
	count,_ := c.Find(query).Count()
	fmt.Println(count)
	return ml, nil
}

// UpdateHospitals updates Hospitals by Id and returns error if
// the record to be updated doesn't exist
func UpdateHospitalsById(id string,m *Hospitals) (err error) {
	session, err := mgo.Dial(beego.AppConfig.String("mongoDb_addr"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gMon").C("hospitals")
	objectId := bson.ObjectIdHex(id)

	c.Update(bson.M{"_id":objectId}, m)
	return err
}

// DeleteHospitals deletes Hospitals by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHospitals(id string) (err error) {

	session, err := mgo.Dial(beego.AppConfig.String("mongoDb_addr"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gMon").C("hospitals")
	objectId := bson.ObjectIdHex(id)
	c.Remove(bson.M{"_id": objectId})

	return err
}
