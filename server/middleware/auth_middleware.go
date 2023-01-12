package middleware

import (
    "time"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    jwt "github.com/appleboy/gin-jwt/v2"

    "app/models"
)

func Auth(identityKey string, userRepo models.UserRepository) *jwt.GinJWTMiddleware {
    m, err := jwt.New(
        &jwt.GinJWTMiddleware{
            Realm:       "test zone",
            Key:         []byte(os.Getenv("JWT_SECRET_KEY")),
            Timeout:     time.Hour,
            MaxRefresh:  time.Hour,
            IdentityKey: identityKey,
            PayloadFunc: func(data interface{}) jwt.MapClaims {
                if v, ok := data.(*models.UserModel); ok {
                    return jwt.MapClaims{
                        identityKey: v.UserName,
                    }
                }
                return jwt.MapClaims{}
            },
            IdentityHandler: func(c *gin.Context) interface{} {
                claims := jwt.ExtractClaims(c)
                return &models.UserModel{
                    UserName: claims[identityKey].(string),
                }
            },
            Authenticator: func(c *gin.Context) (interface{}, error) {
                var loginVals models.UserModel
                if err := c.ShouldBind(&loginVals); err != nil {
                    return "", jwt.ErrMissingLoginValues
                }
                userName := loginVals.UserName
                password := loginVals.Password

                if ( userRepo.Exists(userName, password) ) {
                    return &models.UserModel{
                        UserName:  userName,
                    }, nil
                }

                return nil, jwt.ErrFailedAuthentication
            },
            Authorizator: func(data interface{}, c *gin.Context) bool {
                if v, ok := data.(*models.UserModel); ok && v.UserName == "admin" {
                    return true
                }

                return false
            },
            Unauthorized: func(c *gin.Context, code int, message string) {
                c.JSON(code, gin.H{
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