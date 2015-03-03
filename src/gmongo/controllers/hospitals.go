package controllers

import (
	"encoding/json"
	"errors"
	"gmongo/models"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Hospitals
type HospitalsController struct {
	beego.Controller
}

func (c *HospitalsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Hospitals
// @Param	body		body 	models.Hospitals	true		"body for Hospitals content"
// @Success 200 {int} models.Hospitals.Id
// @Failure 403 body is empty
// @router / [post]
func (c *HospitalsController) Post() {

	var v models.Hospitals
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if v,err := models.AddHospitals(&v); err == nil {
		var s = make(map[string]interface{})
		s["success"] = "true"
		s["hospitals"] = v
		c.Data["json"] = s
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Hospitals by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Hospitals
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HospitalsController) GetOne() {

	idStr := c.Ctx.Input.Params[":id"]
	v, err := models.GetHospitalsById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var s = make(map[string]interface{})
		s["success"] = "true"
		s["hospitals"] = v
		c.Data["json"] = s
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Hospitals
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Success 200 {object} models.Hospitals
// @Failure 403
// @router / [get]
func (c *HospitalsController) GetAll() {

	var sortby string
	var query map[string]string = make(map[string]string)
	var limit int = 10

	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}

	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = v
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllHospitals(query, sortby, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var s = make(map[string]interface{})
		s["success"] = "true"
		s["hospitals"] = l
		c.Data["json"] = s
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Hospitals
// @Param	id		path 	 string	true		"The id you want to update"
// @Param	body		body 	models.Hospitals	true		"body for Hospitals content"
// @Success 200 {object} models.Hospitals
// @Failure 403 :id is not int
// @router /:id [put]
func (c *HospitalsController) Put() {

	idStr := c.Ctx.Input.Params[":id"]

	v := models.Hospitals{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateHospitalsById(idStr,&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Hospitals
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *HospitalsController) Delete() {

	idStr := c.Ctx.Input.Params[":id"]
	if err := models.DeleteHospitals(idStr); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()

}
