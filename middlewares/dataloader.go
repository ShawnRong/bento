package middlewares

import (
	"context"
	"time"

	"github.com/ShawnRong/bento/db"

	"github.com/ShawnRong/bento/models"

	"github.com/ShawnRong/bento/graphql/dataloader"
	"github.com/gin-gonic/gin"
)

type ctxKeyType struct{ name string }

var CtxKey = ctxKeyType{"dataloaderctx"}

type Loaders struct {
	ArticleLoader *dataloader.ArticleLoader
}

func DataLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := dataloader.ArticleLoaderConfig{
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
			Fetch: func(ids []int) ([]*models.Article, []error) {
				var articles []*models.Article
				db.GetDB().Find(&articles)
				return articles, nil
			},
		}
		articleLoader := dataloader.NewArticleLoader(config)
		ctx := context.WithValue(c, CtxKey, &articleLoader)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
