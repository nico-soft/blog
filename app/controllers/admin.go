package controllers

import (
	"strconv"
	"time"

	"blog/app/models"

	"github.com/revel/revel"
)

//Admin controller.
type Admin struct {
	*revel.Controller
}

//Main page.
func (a *Admin) Main() revel.Result {
	return a.Render()
}

type PostData struct {
	Title    string
	Context  string
	Date     time.Time
	Label    string
	Tag      string
	Keywords string
	passwd   string
}

// Add new post.
func (a *Admin) NewPostHandler() revel.Result {

	data := new(PostData)
	a.Params.Bind(&data, "data")

	a.Validation.Required(data.Title).Message("title can't be null.")
	a.Validation.Required(data.Context).Message("context can't be null.")
	a.Validation.Required(data.Date).Message("date can't be null.")

	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		// TODO Redirect new post page.
	}

	blog := new(models.Blogger)
	blog.Title = data.Title
	blog.Context = data.Context
	blog.CreateTime = data.Date

	uid := a.Session["UID"]
	id, _ := strconv.Atoi(uid)

	blog.CreateBy = id

	if data.passwd != "" {
		blog.Passwd = data.passwd
	}

	has, err := blog.New()

	if err != nil || !has {
		a.Flash.Error("msg", "create new blogger post error.")
		// TODO Redirect new post page.
	}

	return a.RenderHtml("ok")
}
