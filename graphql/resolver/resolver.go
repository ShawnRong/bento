package resolver

import (
	"context"

	"github.com/ShawnRong/bento/db"

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
func (r *articleResolver) Createdat(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}
func (r *articleResolver) Updatedat(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}
func (r *articleResolver) Deletedat(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (int, error) {
	panic("not implemented")
}
func (r *commentResolver) Createdat(ctx context.Context, obj *models.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) Updatedat(ctx context.Context, obj *models.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) Deletedat(ctx context.Context, obj *models.Comment) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, limit *int, offset *int) ([]models.User, error) {
	var users []models.User
	db.GetDB().Find(&users)
	return users, nil
}

type tagResolver struct{ *Resolver }

func (r *tagResolver) ID(ctx context.Context, obj *models.Tag) (int, error) {
	panic("not implemented")
}
func (r *tagResolver) Createdat(ctx context.Context, obj *models.Tag) (string, error) {
	panic("not implemented")
}
func (r *tagResolver) Updatedat(ctx context.Context, obj *models.Tag) (string, error) {
	panic("not implemented")
}
func (r *tagResolver) Deletedat(ctx context.Context, obj *models.Tag) (string, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (int, error) {
	return 1, nil
}
func (r *userResolver) Active(ctx context.Context, obj *models.User) (string, error) {
	return "1", nil
}
func (r *userResolver) Createdat(ctx context.Context, obj *models.User) (string, error) {
	return "1", nil
}
func (r *userResolver) Updatedat(ctx context.Context, obj *models.User) (string, error) {
	return "1", nil
}
func (r *userResolver) Deletedat(ctx context.Context, obj *models.User) (string, error) {
	return "1", nil
}
