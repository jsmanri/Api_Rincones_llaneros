package main

import (
	_ "github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	// Registra la base de datos
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))

	// Configuración de CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// Permitimos solicitudes de todos los orígenes
		AllowAllOrigins: true,
		// Permitimos los métodos que serán necesarios
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Permitimos estos encabezados
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		// Permitir credenciales si es necesario
		AllowCredentials: true,
	}))

	// Configuración para el modo de desarrollo
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Inicia el servidor
	beego.Run()
}