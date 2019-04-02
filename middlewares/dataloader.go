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
	UserLoader *dataloader.UserLoader
}

func DataLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := dataloader.UserLoaderConfig{
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
			Fetch: func(ids []int) ([]*models.User, []error) {
				var users []*models.User
				if len(ids) == 1 {
					db.GetDB().Where("id = ?", ids).Find(users)
				} else {
					db.GetDB().Where("id in (?)", ids).Find(users)
				}
				return users, nil
			},
		}
		userLoader := dataloader.NewUserLoader(config)
		ctx := context.WithValue(c, CtxKey, &userLoader)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
