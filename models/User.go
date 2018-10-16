package models

import (
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id int 
    Name string 
    Client string 
    Url string 
    Notes string 
}

func (a *User) TableName() string {
	return TableName("user")
}

func GetUserById(id int) (*User, error){
    user := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
