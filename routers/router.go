// @APIVersion 1.0
// @Title Cita Médica
// @Description Api para citas médicas dentro del módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co
// @TermsOfServiceUrl http://www.udistrital.edu.co/politicas-de-privacidad.pdf
// @License BSD-3-Clause
// @LicenseUrl http://opensource.org/licenses/BSD-3-Clause
// @BasePath /Cita/v1
package routers

import (
	"Cita/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/Cita",
		beego.NSNamespace("/Cita",
			beego.NSInclude(
				&controllers.CitaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
