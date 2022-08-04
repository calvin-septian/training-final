package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"training-final/model"
)

func (s *SQLServer) AddSocialMedia(data model.SocialMedia, userId int, ctx context.Context) int {
	var sosmedId int
	err := s.LocalDB.QueryRow("insert into tb_socialmedia (Name, Social_media_url, User_Id, Created_at, Updated_at) values (@Name, @Social_media_url, @User_Id, @Created_at, @Updated_at); select Id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("Name", data.Name),
		sql.Named("Social_media_url", data.Social_media_url),
		sql.Named("User_Id", userId),
		sql.Named("Created_at", time.Now()),
		sql.Named("Updated_at", time.Now())).Scan(&sosmedId)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return sosmedId
}

func (s *SQLServer) GetSocialMediaById(sosmedId int, ctx context.Context) model.SocialMedia {
	var sosmed model.SocialMedia
	err := s.LocalDB.QueryRow("select Id, Name, Social_media_url, User_Id, Created_at, Updated_at from tb_socialmedia where Id = @Id",
		sql.Named("Id", sosmedId)).Scan(&sosmed.Id, &sosmed.Name, &sosmed.Social_media_url, &sosmed.User_Id, &sosmed.Created_at, &sosmed.Updated_at)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return sosmed
}

func (s *SQLServer) GetAllSocialMedia(ctx context.Context) []model.SocialMediaGetResponse {
	rows, err := s.LocalDB.Query("select Id, Name, Social_media_url, User_Id, Created_at, Updated_at from tb_socialmedia")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows.Close()

	var result []model.SocialMediaGetResponse
	for rows.Next() {
		sosmed := model.SocialMediaGetResponse{}

		err := rows.Scan(&sosmed.Id, &sosmed.Name, &sosmed.Social_media_url, &sosmed.User_Id, &sosmed.Created_at, &sosmed.Updated_at)
		if err != nil {
			fmt.Println("error: ", err)
			return nil
		}

		err = s.LocalDB.QueryRow("select Id, Username from tb_user where User_Id = @Id",
			sql.Named("Id", sosmed.User_Id)).Scan(&sosmed.User.Id, &sosmed.User.Username)
		if err != nil {
			fmt.Println("error: ", err)
		}

		result = append(result, sosmed)
	}

	return result
}

func (s *SQLServer) UpdateSocialMedia(data model.SocialMedia, sosmedId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("update tb_socialmedia set Name = @Name, Social_media_url = @Social_media_url, Updated_at = @Updated_at where Id = @Id",
		sql.Named("Name", data.Name),
		sql.Named("Social_media_url", data.Social_media_url),
		sql.Named("Updated_at", time.Now()),
		sql.Named("Id", sosmedId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (s *SQLServer) DeleteSocialMedia(sosmedId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("delete tb_socialmedia where Id = @Id",
		sql.Named("Id", sosmedId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
