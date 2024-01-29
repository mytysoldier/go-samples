package funcs

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-rest/model"
	"log"
)

func InsertUser(db *sql.DB, user model.User) (model.User, error) {
	var insertedUser model.User

	err := db.QueryRow(
		"INSERT INTO public.user (name, age)"+
			"VALUES ($1, $2) RETURNING id, name, age",
		user.Name, user.Age).Scan(
		&insertedUser.Id, &insertedUser.Name, &insertedUser.Age,
	)

	if err != nil {
		log.Fatal("error inserting user")
		fmt.Println(err)
		return model.User{}, err
	}

	return insertedUser, nil
}

func GetUserByID(db *sql.DB, id int) (model.User, error) {
	row := db.QueryRow("SELECT id, name, age FROM public.user WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		log.Fatal("error scanning user")
		return model.User{}, err
	}

	return user, nil
}

func UpdateUser(db *sql.DB, user model.User) (model.User, error) {
	var updatedUser model.User

	err := db.QueryRow(
		"UPDATE public.user SET name=$1, age=$2 WHERE id=$3"+
			" RETURNING id, name, age",
		user.Name, user.Age, user.Id,
	).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Age)

	if err != nil {
		log.Fatal("error updating user")
		return model.User{}, err
	}

	return updatedUser, nil
}

func DeleteUser(db *sql.DB, id int) error {
	existingUser, err := GetUserByID(db, id)
	if err != nil {
		log.Fatal("error deleting user")
		return err
	}

	// ユーザーが存在しない場合はエラーを返す
	if existingUser.Id == 0 {
		return errors.New("user not found")
	}

	// ユーザーが存在する場合は削除を実行
	_, err = db.Exec("DELETE FROM public.user WHERE id = $1", id)
	if err != nil {
		log.Fatal("error deleting user")
		return err
	}

	return nil
}
