package middleware

import (
    "time"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    jwt "github.com/appleboy/gin-jwt/v2"

    "app/models"
    "app/services"
)

var IdentityKey = "id"

func Auth(userRepo models.UserRepository) *jwt.GinJWTMiddleware {
    m, err := jwt.New(
        &jwt.GinJWTMiddleware{
            Realm:       "test zone",
            Key:         []byte(os.Getenv("JWT_SECRET_KEY")),
            Timeout:     time.Hour,
            MaxRefresh:  time.Hour*24*31,
            IdentityKey: IdentityKey,
            PayloadFunc: func(data interface{}) jwt.MapClaims {
                if v, ok := data.(*models.UserModel); ok {
                    return jwt.MapClaims{
                        IdentityKey: v.Email,
                    }
                }
                return jwt.MapClaims{}
            },
            IdentityHandler: func(ctx *gin.Context) interface{} {
                claims := jwt.ExtractClaims(ctx)
                return &models.UserModel{
                    Email: claims[IdentityKey].(string),
                }
            },
            Authenticator: func(ctx *gin.Context) (interface{}, error) {
                var user models.UserModel
                if err := ctx.ShouldBind(&user); err != nil {
                    return "", jwt.ErrMissingLoginValues
                }
                user.Password = services.Hash().Make(user.Password)

                if ( userRepo.Exists(user) ) {
                    return &models.UserModel{
                        Email:  user.Email,
                    }, nil
                }

                return nil, jwt.ErrFailedAuthentication
            },
            Authorizator: func(data interface{}, ctx *gin.Context) bool {
                if _, ok := data.(*models.UserModel); ok {
                    return true
                }

                return false
            },
            Unauthorized: func(ctx *gin.Context, code int, message string) {
                ctx.JSON(code, gin.H{
                    "code":    code,
                    "message": message,
                })
            },
            TokenLookup: "header: Authorization, query: token, cookie: jwt",
            TokenHeadName: "Bearer",
            TimeFunc: time.Now,
        },
    )

    if err != nil {
      log.Fatal("JWT Error:" + err.Error())
    }

    return m
}