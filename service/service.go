package service

import (
	"Geeklanet/datasource"
	"Geeklanet/repository"
)

type Service struct {
	User		userService
	Post		postService
	Notice		noticeService
	Recommend	recommendService
	Tag			tagService
	Achievement	achievementService
}

func NewService() *Service {
	database := datasource.GetDataBase("mongoDB","Geeklanet")

	user := userService{
		r: repository.UserRepository{
			T: database.GetTable("user"),
		},
	}

	post := postService{
		postR: repository.PostRepository{
			T: database.GetTable("post"),
		},
		commentR: repository.CommentRepository{
			T: database.GetTable("comment"),
		},
		subCommentR: repository.SubCommentRepository{
			T: database.GetTable("subComment"),
		},
	}

	notice := noticeService{
		r: repository.NoticeRepository{
			T: database.GetTable("notice"),
		},
	}

	tag := tagService{
		r: repository.TagRepository{
			T: database.GetTable("tag"),
		},
	}

	achievement := achievementService{
		r: repository.AchievementRepository{
			T: database.GetTable("user"),
		},
	}

	recommend := recommendService{
		r: repository.RecommendRepository{
			T: database.GetTable("user"),
		},
	}

	return &Service{
		User:        user,
		Post:        post,
		Notice:      notice,
		Recommend:   recommend,
		Tag:         tag,
		Achievement: achievement,
	}
}