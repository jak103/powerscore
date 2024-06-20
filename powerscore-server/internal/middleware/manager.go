package middleware

import (
	"powerscore-server/internal/utils/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func Setup(app *fiber.App) {
	app.Use(helmet.New())

	app.Use(requestid.New(requestid.Config{
		Generator:  utils.UUIDv4,
		ContextKey: constants.RequestIdLocal,
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		// StackTraceHandler: appUtils.StackTraceHandler,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))

	// if config.Vars.Env == constants.Local || config.Vars.Env == constants.Test {
	// 	log.Alert("Setting permissive CORS")
	// 	// CORS https://docs.gofiber.io/api/middleware/cors /
	// 	app.Use(cors.New(cors.Config{
	// 		AllowOrigins:     "http://localhost:9002",
	// 		AllowCredentials: true,
	// 		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
	// 		AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cookie",
	// 		ExposeHeaders:    "Set-Cookie",
	// 	}))
	// }

}
