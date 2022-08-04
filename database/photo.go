package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"training-final/model"
)

func (s *SQLServer) AddPhoto(data model.Photo, userId int, ctx context.Context) int {
	var photoId int
	err := s.LocalDB.QueryRow("insert into tb_photo (Title, Caption, Photo_Url, User_Id, Created_at, Updated_at) values (@Title, @Caption, @Photo_Url, @User_Id, @Created_at, @Updated_at); select Id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("Title", data.Title),
		sql.Named("Caption", data.Caption),
		sql.Named("Photo_Url", data.Photo_Url),
		sql.Named("User_Id", userId),
		sql.Named("Created_at", time.Now()),
		sql.Named("Updated_at", time.Now())).Scan(&photoId)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return photoId
}

func (s *SQLServer) GetPhotoById(photoId int, ctx context.Context) model.Photo {
	var photo model.Photo
	err := s.LocalDB.QueryRow("select Id, Title, Caption, Photo_Url, User_Id, Created_at, Updated_at from tb_photo where Id = @Id",
		sql.Named("Id", photoId)).Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Photo_Url, &photo.User_Id, &photo.Created_at, &photo.Updated_at)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return photo
}

func (s *SQLServer) GetAllPhoto(ctx context.Context) []model.PhotoGetResponse {
	rows, err := s.LocalDB.Query("select Id, Title, Caption, Photo_Url, User_Id, Created_at, Updated_at from tb_photo")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows.Close()

	var result []model.PhotoGetResponse
	for rows.Next() {
		photo := model.PhotoGetResponse{}

		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Photo_Url, &photo.User_Id, &photo.Created_at, &photo.Updated_at)
		if err != nil {
			fmt.Println("error: ", err)
			return nil
		}

		err = s.LocalDB.QueryRow("select Email, Username from tb_user where Id = @Id",
			sql.Named("Id", photo.User_Id)).Scan(&photo.User.Email, &photo.User.Username)
		if err != nil {
			fmt.Println("error: ", err)
		}

		result = append(result, photo)
	}

	return result
}

func (s *SQLServer) UpdatePhoto(data model.Photo, photoId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("update tb_photo set Title = @Title, Caption = @Caption, Photo_Url = @Photo_Url, Updated_at = @Updated_at where Id = @Id",
		sql.Named("Title", data.Title),
		sql.Named("Caption", data.Caption),
		sql.Named("Photo_Url", data.Photo_Url),
		sql.Named("Updated_at", time.Now()),
		sql.Named("Id", photoId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (s *SQLServer) DeletePhoto(photoId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("delete from tb_photo where Id = @Id",
		sql.Named("Id", photoId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
