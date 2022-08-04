package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"training-final/model"
)

func (s *SQLServer) AddComment(data model.Comment, userId int, ctx context.Context) int {
	var photoId int
	err := s.LocalDB.QueryRow("select Id from tb_photo where User_Id = @Id",
		sql.Named("Id", userId)).Scan(&photoId)
	if err != nil {
		fmt.Println("error: ", err)
	}

	var commentId int
	err = s.LocalDB.QueryRow("insert into tb_comment (Message, Photo_Id, User_Id, Created_at, Updated_at) values (@Title, @Caption, @Photo_Url, @User_Id, @Created_at, @Updated_at); select Id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("Message", data.Message),
		sql.Named("Photo_Id", photoId),
		sql.Named("User_Id", userId),
		sql.Named("Created_at", time.Now()),
		sql.Named("Updated_at", time.Now())).Scan(&commentId)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return commentId
}

func (s *SQLServer) GetCommentById(commentId int, ctx context.Context) model.Comment {
	var comment model.Comment
	err := s.LocalDB.QueryRow("select Id, Message, Photo_Id, User_Id, Created_at, Updated_at from tb_comment where Id = @Id",
		sql.Named("Id", commentId)).Scan(&comment.Id, &comment.Message, &comment.Photo_Id, &comment.User_Id, &comment.Created_at, &comment.Updated_at)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return comment
}

func (s *SQLServer) GetAllComments(ctx context.Context) []model.CommentGetResponse {
	rows, err := s.LocalDB.Query("select Id, Message, Photo_Id, User_Id, Created_at, Updated_at from tb_comment")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows.Close()

	var result []model.CommentGetResponse
	for rows.Next() {
		comment := model.CommentGetResponse{}

		err := rows.Scan(&comment.Id, &comment.Message, &comment.Photo_Id, &comment.User_Id, &comment.Created_at, &comment.Updated_at)
		if err != nil {
			fmt.Println("error: ", err)
			return nil
		}

		err = s.LocalDB.QueryRow("select Id, Email, Username from tb_user where User_Id = @Id",
			sql.Named("Id", comment.User_Id)).Scan(&comment.User.Id, &comment.User.Email, &comment.User.Username)
		if err != nil {
			fmt.Println("error: ", err)
		}

		err = s.LocalDB.QueryRow("select Id, Title, Caption, Photo_Url, User_Id from tb_photo where User_Id = @Id",
			sql.Named("Id", comment.User_Id)).Scan(&comment.Photo.Id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.Photo_Url, &comment.Photo.User_Id)
		if err != nil {
			fmt.Println("error: ", err)
		}

		result = append(result, comment)
	}

	return result
}

func (s *SQLServer) UpdateComment(data model.Comment, commentId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("update tb_comment set Message = @Message, Updated_at = @Updated_at where Id = @Id",
		sql.Named("Message", data.Message),
		sql.Named("Updated_at", time.Now()),
		sql.Named("Id", commentId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (s *SQLServer) DeleteComment(commentId int, ctx context.Context) {
	_, err := s.LocalDB.Exec("delete tb_comment where Id = @Id",
		sql.Named("Id", commentId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
