package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:DocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:SoporteDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documento_programa_crud/controllers:TipoDocumentoProgramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
