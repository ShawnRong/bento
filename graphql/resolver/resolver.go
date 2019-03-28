package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ShawnRong/bento/db"

	"github.com/vektah/gqlparser/gqlerror"

	"github.com/ShawnRong/bento/graphql/generated"
	"github.com/ShawnRong/bento/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Article() generated.ArticleResolver {
	return &articleResolver{r}
}
func (r *Resolver) Comment() generated.CommentResolver {
	return &commentResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Tag() generated.TagResolver {
	return &tagResolver{r}
}
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type articleResolver struct{ *Resolver }

func (r *articleResolver) ID(ctx context.Context, obj *models.Article) (int, error) {
	panic("not implemented")
}
func (r *articleResolver) User(ctx context.Context, obj *models.Article) (*models.User, error) {
	panic("not implemented")
}
func (r *articleResolver) CreatedAt(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}
func (r *articleResolver) UpdatedAt(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}
func (r *articleResolver) DeletedAt(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (int, error) {
	panic("not implemented")
}
func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error) {
	panic("not implemented")
}
func (r *commentResolver) CreatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) UpdatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) DeletedAt(ctx context.Context, obj *models.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) Article(ctx context.Context, obj *models.Comment) (*models.Article, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Register(ctx context.Context, input models.NewUser) (*models.User, error) {
	newUser := models.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}
	db.GetDB().FirstOrCreate(&newUser, &newUser)
	return &newUser, nil
}
func (r *mutationResolver) UpdateProfile(ctx context.Context, input models.NewUser) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateArticle(ctx context.Context, input models.NewArticle) (*models.Article, error) {
	var tags []models.Tag
	db.GetDB().Where("id in (?)", input.Tags).Find(&tags)
	newArticle := models.Article{
		Content: input.Content,
		UserID:  uint(input.UserID),
		Tags:    tags,
	}
	db.GetDB().Create(&newArticle)
	return &newArticle, nil
}
func (r *mutationResolver) DeleteArticle(ctx context.Context, id int) (*models.Article, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateArticle(ctx context.Context, input models.NewArticle) (*models.Article, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTag(ctx context.Context, input models.NewTag) (*models.Tag, error) {
	newTag := models.Tag{
		Name: input.Name,
	}
	db.GetDB().FirstOrCreate(&newTag, &newTag)
	return &newTag, nil
}
func (r *mutationResolver) DeleteTag(ctx context.Context, id int) (*models.Tag, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateComment(ctx context.Context, input models.NewComment) (*models.Comment, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*models.Comment, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]models.User, error) {
	var users []models.User
	if err := db.GetDB().Find(&users).Error; err != nil {
		return []models.User{}, gqlerror.Errorf("err: %v", err)
	}
	return users, nil
}
func (r *queryResolver) Me(ctx context.Context, id int) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Article(ctx context.Context, id *int, tag *string) (*models.Article, error) {
	panic("not implemented")
}
func (r *queryResolver) Articles(ctx context.Context, limit *int, offset *int) ([]models.Article, error) {
	panic("not implemented")
}
func (r *queryResolver) Tags(ctx context.Context) ([]models.Tag, error) {
	panic("not implemented")
}

type tagResolver struct{ *Resolver }

func (r *tagResolver) ID(ctx context.Context, obj *models.Tag) (int, error) {
	panic("not implemented")
}
func (r *tagResolver) CreatedAt(ctx context.Context, obj *models.Tag) (string, error) {
	panic("not implemented")
}
func (r *tagResolver) UpdatedAt(ctx context.Context, obj *models.Tag) (string, error) {
	panic("not implemented")
}
func (r *tagResolver) DeletedAt(ctx context.Context, obj *models.Tag) (string, error) {
	panic("not implemented")
}
func (r *tagResolver) Articles(ctx context.Context, obj *models.Tag) ([]models.Article, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (int, error) {
	return int(obj.ID), nil
}
func (r *userResolver) Active(ctx context.Context, obj *models.User) (string, error) {
	return strconv.FormatBool(obj.Active), nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	t := obj.CreatedAt
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()), nil
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	t := obj.UpdatedAt
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()), nil
}
func (r *userResolver) DeletedAt(ctx context.Context, obj *models.User) (string, error) {
	t := obj.DeletedAt
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()), nil
}
func (r *userResolver) Comments(ctx context.Context, obj *models.User) ([]models.Comment, error) {
	panic("not implemented")
}
