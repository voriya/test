package main

import (
	"flag"
	"fmt"
	
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
)

type Group struct {
	ID          int64  `json:"id" orm:"column(id)"`
	GroupName   string `json:"group_name" orm:"column(group_name)"`
	Description string `json:"description" orm:"column(description)"`
	Extra       string `json:"extra" orm:"column(extra)"`
	Owner       string `json:"owner" orm:"column(owner)"`
	Namespace   string `json:"namespace" orm:"column(namespace)"`
	CreateTime  int64  `json:"create_time" orm:"column(create_time)"`
	UpdateTime  int64  `json:"update_time" orm:"column(update_time)"`
}

func init() {
	var configPath string
	flag.StringVar(&configPath, "config", "conf/service.conf", "server config.")
	flag.Parse()
	if _, err := toml.DecodeFile(configPath, &config.Config); err != nil {
		fmt.Printf("fail to read config.||err=%v||config=%v", err, configPath)
		os.Exit(1)
		return
	}

	dbInfo := config.Config.MysqlConfig
	loginUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbInfo.UserName, dbInfo.Password, dbInfo.Ip, dbInfo.Port, dbInfo.DbName)
	fmt.Printf("loginUrl:%s\n", loginUrl)
	orm.RegisterDataBase("default", "mysql", loginUrl, 30)
	orm.RegisterModel(new(Group))
	//orm.RunSyncdb("default", false, true)
}

func main() {

	//o := orm.NewOrm()
	//o.Using("default")

	// insert
	//group := Group{ID:5,GroupName: "aaaa", Description: "aaaaaaa", Owner: "lyx", Namespace: "sim"}
	//id, err := o.Insert(&group)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//update
	//g := Group{
	//	ID:          1,
	//	Namespace:   "sim",
	//	Description: "updateupdateupdateupdate",
	//}
	//group := Group{ID: 1}
	//if o.Read(&group) == nil {
	//	if num, err := o.Update(&g); err == nil {
	//		fmt.Println(num)
	//	}
	//}

	//search
	//var maps []orm.Params
	//_,err := o.QueryTable("group").OrderBy("ID").Limit(10,0).Values(&maps)
	//if err != nil{
	//	fmt.Printf("%s\n",err)
	//}
	//fmt.Printf("len = %v\n,", len(maps))
	//for _,row := range maps{
	//	fmt.Println(row["ID"])
	//}
	//o.Begin()

	//start := 0
	//leng := 10
	//var groups []*Group
	//_,er := o.Raw("SELECT * FROM `group` WHERE namespace = ? ORDER BY id ASC LIMIT ?,?","sim",start,leng).QueryRows(&groups)
	//
	////sql:= "SELECT * FROM `group` ORDER BY id ASC LIMIT %s,%s"
	////_, er := o.Raw(sql).QueryRows(&groups)
	//if er != nil {
	//	fmt.Println(er)
	//} else {
	//	fmt.Printf("students2 is : %v\n", groups)
	//	for index, _ := range groups {
	//		fmt.Printf("第%d个学生个人信息：", index + 1)
	//	}
	//}

	//search one

	//var groups []*Group
	//_,er := o.Raw("SELECT * FROM `group` WHERE id = ?  ",3).QueryRows(&groups)
	//
	////sql:= "SELECT * FROM `group` ORDER BY id ASC LIMIT %s,%s"
	////_, er := o.Raw(sql).QueryRows(&groups)
	//if er != nil {
	//	fmt.Println(er)
	//} else {
	//	fmt.Printf("group=%v", groups[0])
	//}

	// update
	//user.Name = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//// read one
	//u := User{Id: user.Id}
	//err = o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//
	//// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//if _, err := o.Delete(&Group{ID: 1}); err != nil {
	//	fmt.Printf("err:%v", err)
	//}

	//serarch one
	//var group Group
	//err := o.QueryTable("group").Filter("id", 5).One(&group, "id","group_name","owner")
	//if err != nil {
	//	fmt.Printf("Err:%v\n",err)
	//}
	//fmt.Printf("group:%n",group)

	group := Group{ID:1,GroupName:"testreflect"}

	tg := reflect.TypeOf(group)
	vg := reflect.ValueOf(group)
	kd := vg.Kind()
	fmt.Println(vg,kd)
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	fmt.Println(tg.Name())
	fmt.Println(tg.Kind())

	fmt.Println(tg.FieldByName("id"))
	fmt.Println(tg.FieldByName("ID"))

	for i := 0; i<vg.NumField(); i++ {
		//field := tg.Field(i)
		//tag := field.Tag.Get("json")



		//fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		fmt.Printf("%d %v\n",i,vg.Field(i))
	}

}
