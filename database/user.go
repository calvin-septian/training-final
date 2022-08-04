package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"training-final/model"
)

func (s *SQLServer) AddUser(data model.User, ctx context.Context) int {
	var userId int
	err := s.LocalDB.QueryRow("insert into tb_user (Age, Email, Username, Password, Created_at, Updated_at) values (@Age, @Email, @Username, @Password, @Created_at, @Updated_at); select Id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("Age", data.Age),
		sql.Named("Email", data.Email),
		sql.Named("Username", data.Username),
		sql.Named("Password", data.Password),
		sql.Named("Created_at", time.Now()),
		sql.Named("Updated_at", time.Now())).Scan(&userId)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return userId
}

func (s *SQLServer) GetUserById(userId int, ctx context.Context) model.User {
	var user model.User
	err := s.LocalDB.QueryRow("select Id, Username, Email, Age, Created_at, Updated_at from tb_user where Id = @Id",
		sql.Named("Id", userId)).Scan(&user.Id, &user.Username, &user.Email, &user.Age, &user.Created_at, &user.Updated_at)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return user
}

func (s *SQLServer) GetPasswordByEmail(email string, ctx context.Context) model.User {
	var user model.User
	err := s.LocalDB.QueryRow("select Id, Password from tb_user where Email = @Email",
		sql.Named("Email", email)).Scan(&user.Id, &user.Password)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return user
}

func (s *SQLServer) UpdateUser(data model.User, userId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("update tb_user set Email = @Email, Username = @Username, Updated_at = @Updated_at where Id = @Id",
		sql.Named("Email", data.Email),
		sql.Named("Username", data.Username),
		sql.Named("Updated_at", time.Now()),
		sql.Named("Id", userId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (s *SQLServer) DeleteUser(userId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("delete from tb_user where Id = @Id",
		sql.Named("Id", userId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
