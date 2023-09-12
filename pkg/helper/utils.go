package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/dto/request"
	"log"
	"os"
	"regexp"
	"strings"
)

func HeaderGetter(ctx *gin.Context) (header request.HeaderDto) {
	ip := ctx.ClientIP()
	location := ctx.GetHeader("Origin")
	userAgent := ctx.GetHeader("User-Agent")

	header = request.HeaderDto{
		Ip:        ip,
		Location:  location,
		UserAgent: userAgent,
	}
	return
}

// SetEnv is function for check and return env variable
func SetEnv(key string, defaultVal string) string {

	rootPath := SetRootPath()
	err := godotenv.Load(rootPath + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// set and check env is exists
	strVal, isFound := os.LookupEnv(key)
	if isFound {
		return strVal
	} else {
		return defaultVal
	}
}

func SetRootPath() string {
	_, ok := os.LookupEnv("ENVIRONMENT")
	mainPath := "build"
	if ok {
		mainPath = "build"
	} else {
		mainPath = "privyid-golang-test"
	}

	projectName := regexp.MustCompile(`^(.*` + mainPath + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func IsExistsInSlice[T comparable](sl []T, val T) bool {
	for _, v := range sl {
		if v == val {
			return true
		}
	}
	return false
}

func RemoveFromSlice[T comparable](sl []T, val T) []T {
	for i, v := range sl {
		if v == val {
			return append(sl[:i], sl[i+1:]...)
		}
	}
	return sl
}
