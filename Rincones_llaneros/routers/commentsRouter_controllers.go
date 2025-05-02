package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:EventosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:SitiosTuristicosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Patch",
            Router: "/:id",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/Franckk24/Api_Rincones_llaneros/Rincones_llaneros/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
