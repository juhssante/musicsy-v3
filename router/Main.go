package маршрутизатор

import ( 
	"Приложение/хттн/контроллеры"
	"Приложение/хттн/промежуточное"

	"github.com/gin-gonic/gin"
)

func ОбъявлятьМаршрутизатор(двигатель *gin.Engine) {
	двигатель.Use(промежуточное.Журнал)
	двигатель.GET("/пинг", контроллеры.ГлавныйКонтроллер{}.Пинг)

	апи := двигатель.Group("/апи")
	{
		апи.GET("/", контроллеры.ГлавныйКонтроллер{}.Домашняястраница)

		пользователи := апи.Group("/пользователи")
		пользователи.
			Use(middlewares.LoadUser).
			Use(middlewares.LogLatency)
		{
			userController := controllers.UserController{}
			users.GET("", userController.List)
			users.POST("", userController.Post)
			users.PATCH("/:userId", userController.Patch)
			users.GET("/:userId", userController.Get)
			users.DELETE("/:userId", userController.Delete)
		}
	}
	engine.NoRoute(controllers.MainController{}.Client)
	engine.Static("/static", "./dist")
}
