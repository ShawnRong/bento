package resolver

import (
	"context"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/ShawnRong/bento/db"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/ShawnRong/bento/graphql/generated"
	"github.com/ShawnRong/bento/models"
)

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

func (r *articleResolver) User(ctx context.Context, obj *models.Article) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where("id = ?", obj.UserID).Find(&user).Error
	return &user, err
}
func (r *articleResolver) Tags(ctx context.Context, obj *models.Article) ([]models.Tag, error) {
	var tags []models.Tag
	err := db.GetDB().Joins("join `article_tags` on `article_tags`.`tag_id` = `tags`.`id`").
		Where("`article_tags`.`article_id` = ?", obj.ID).
		Find(&tags).Error
	return tags, err
}
func (r *articleResolver) Comments(ctx context.Context, obj *models.Article) ([]models.Comment, error) {
	var comments []models.Comment
	err := db.GetDB().Where("article_id = ?", obj.ID).Find(&comments).Error
	return comments, err
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where("id = ?", obj.UserID).Find(&user).Error
	return &user, err
}
func (r *commentResolver) Article(ctx context.Context, obj *models.Comment) (*models.Article, error) {
	var article models.Article
	err := db.GetDB().Where("id = ?", obj.ArticleID).Find(&article).Error
	return &article, err
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Register(ctx context.Context, input models.NewUser) (*models.User, error) {
	//Check email if unique
	var user models.User
	if err := db.GetDB().Model(&user).Where("email = ?", input.Email).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	if user.Email != "" {
		return nil, gqlerror.Errorf("User email already exist!")
	}

	newUser := models.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}
	if err := db.GetDB().FirstOrCreate(&newUser, &newUser).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &newUser, nil
}
func (r *mutationResolver) UpdateProfile(ctx context.Context, input models.UpdateUser) (*models.User, error) {
	var user models.User
	newUser := models.User{}
	if name := input.Name; name != nil {
		newUser.Name = *name
	}
	if email := input.Email; email != nil {
		newUser.Email = *email
	}
	if password := input.Password; password != nil {
		newUser.Password = *password
	}
	if err := db.GetDB().Model(&user).Where("id = ?", input.ID).Update(&newUser).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &newUser, nil
}
func (r *mutationResolver) CreateArticle(ctx context.Context, input models.NewArticle) (*models.Article, error) {
	var tags []models.Tag
	if err := db.GetDB().Where("id in (?)", input.Tags).Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %s", err)
	}
	newArticle := models.Article{
		Content: input.Content,
		UserID:  uint(input.UserID),
		//Tags:    tags,
	}
	if err := db.GetDB().Create(&newArticle).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &newArticle, nil
}
func (r *mutationResolver) DeleteArticle(ctx context.Context, id int) (*models.Article, error) {
	var article models.Article
	if err := db.GetDB().Where("id = ?", id).Delete(&article).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &article, nil
}
func (r *mutationResolver) UpdateArticle(ctx context.Context, input models.UpdateArticle) (*models.Article, error) {
	var tags []models.Tag
	var article models.Article
	if err := db.GetDB().Where("id in (?)", input.Tags).Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	newArticle := models.Article{}
	if content := input.Content; content != nil {
		newArticle.Content = *content
	}
	if userID := input.UserID; userID != nil {
		newArticle.UserID = uint(*userID)
	}
	//@TODO
	//newArticle.Tags = tags

	if err := db.GetDB().Model(&article).Where("id = ?", input.ID).Update(&newArticle).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &newArticle, nil
}
func (r *mutationResolver) CreateTag(ctx context.Context, input models.NewTag) (*models.Tag, error) {
	newTag := models.Tag{
		Name: input.Name,
	}
	if err := db.GetDB().FirstOrCreate(&newTag, &newTag).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &newTag, nil
}
func (r *mutationResolver) DeleteTag(ctx context.Context, id int) (*models.Tag, error) {
	var tag models.Tag
	if err := db.GetDB().Where("id = ?", id).Delete(&tag).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &tag, nil
}
func (r *mutationResolver) UpdateTag(ctx context.Context, input models.UpdateTag) (*models.Tag, error) {
	var tag models.Tag
	updateTag := models.Tag{}
	if name := input.Name; name != nil {
		updateTag.Name = *name
	}
	if err := db.GetDB().Model(&tag).Where("id = ?", input.ID).Update(&updateTag).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &tag, nil
}
func (r *mutationResolver) CreateComment(ctx context.Context, input models.NewComment) (*models.Comment, error) {
	newComment := models.Comment{
		Content:   input.Content,
		UserID:    uint(input.UserID),
		ArticleID: uint(input.ArticleID),
	}
	if err := db.GetDB().Create(&newComment).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &newComment, nil
}
func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*models.Comment, error) {
	var comment models.Comment
	if err := db.GetDB().Where("id = ?", id).Delete(&comment).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return &comment, nil
}
func (r *mutationResolver) UpdateComment(ctx context.Context, input models.UpdateComment) (*models.Comment, error) {
	var comment models.Comment
	newComment := models.Comment{}
	if content := input.Content; content != nil {
		newComment.Content = *content
	}
	if userID := input.UserID; userID != nil {
		newComment.UserID = uint(*userID)
	}
	if articleID := input.ArticleID; articleID != nil {
		newComment.ArticleID = uint(*articleID)
	}
	if err := db.GetDB().Model(&comment).Where("id = ?", input.ID).Update(newComment).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}

	return &newComment, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]models.User, error) {
	var users []models.User
	if err := db.GetDB().Limit(*limit).Offset(*offset).Find(&users).Error; err != nil {
		return []models.User{}, gqlerror.Errorf("db error: %v", err)
	}
	return users, nil
}
func (r *queryResolver) Me(ctx context.Context, id int) (*models.User, error) {
	var user *models.User
	if err := db.GetDB().Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return user, nil
}
func (r *queryResolver) Article(ctx context.Context, id *int, tag *string) (*models.Article, error) {
	var article *models.Article
	if id != nil {
		if err := db.GetDB().Where("id = ?", id).Find(&article).Error; err != nil {
			return nil, gqlerror.Errorf("db error: %v", err)
		}
	}
	if tag != nil {
		if err := db.GetDB().Where("name LIKE ?", tag).Find(&article).Error; err != nil {
			return nil, gqlerror.Errorf("db error: %v", err)
		}
	}
	return article, nil
}
func (r *queryResolver) Articles(ctx context.Context, limit *int, offset *int) ([]models.Article, error) {
	var articles []models.Article
	if err := db.GetDB().Limit(*limit).Offset(*offset).Find(&articles).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return articles, nil
}
func (r *queryResolver) Tags(ctx context.Context) ([]models.Tag, error) {
	var tags []models.Tag
	if err := db.GetDB().Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("db error: %v", err)
	}
	return tags, nil
}

type tagResolver struct{ *Resolver }

func (r *tagResolver) Articles(ctx context.Context, obj *models.Tag) ([]models.Article, error) {
	var articles []models.Article
	err := db.GetDB().Joins("join `article_tags` on `article_tags`.`article_id` = `articles`.`id`").
		Where("`article_tags`.`tag_id` = ?", obj.ID).
		Find(&articles).Error
	return articles, err
}

type userResolver struct{ *Resolver }

func (r *userResolver) Active(ctx context.Context, obj *models.User) (string, error) {
	return strconv.FormatBool(obj.Active), nil
}
func (r *userResolver) Articles(ctx context.Context, obj *models.User) ([]models.Article, error) {
	var articles []models.Article
	err := db.GetDB().Where("user_id = ?", obj.ID).Find(&articles).Error
	return articles, err
}
func (r *userResolver) Comments(ctx context.Context, obj *models.User) ([]models.Comment, error) {
	var comments []models.Comment
	err := db.GetDB().Where("user_id = ?", obj.ID).Find(&comments).Error
	return comments, err
}
