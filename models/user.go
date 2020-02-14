package models

import "log"

type User struct {
	Name 	 string
	Password string
}

func(db *DB) AddUser(user User) error{
	_, err := db.Exec("insert into users values(?,?)", user.Name, user.Password)
	if err!=nil{
		log.Println(err.Error())
	}
	return err
}
func (db *DB) GetUser(name string)(User, error){
	user := User{}
	row := db.QueryRow("select * from users where name=?", name)
	err:= row.Scan(&user.Name, &user.Password)
	if err!=nil{
		log.Println(err.Error())
	}
	return user, err
}

func(db *DB)CompareUser(name, password string)(bool, error){
	user, err:= db.GetUser(name)
	if err!=nil{
		return false, err
	}
	if user.Password==password{
		return true, nil
	}
	return false, nil
}