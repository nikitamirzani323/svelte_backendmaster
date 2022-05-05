package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/go_mastertoto/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controller.HealthCheck)
	api := app.Group("/api/")

	api.Post("init", controller.Init)
	api.Post("login", controller.CheckLogin)
	api.Post("dashboardwinlose", controller.DashboardComp_winlose)
	api.Post("dashboardwinlosecompany", controller.DashboardComp_winlosecompany)
	api.Post("dashboardlistcompany", controller.DashboardComp_listcompany)

	api.Post("log", controller.Log)
	api.Post("company", controller.Company)
	api.Post("companydetail", controller.Companydetail)
	api.Post("companylistadmin", controller.Companylistadmin)
	api.Post("companylistpasaran", controller.Companylistpasaran)
	api.Post("companylistkeluaran", controller.Companylistkeluaran)
	api.Post("companyinvoicemember", controller.Companyinvoicemember)
	api.Post("companyinvoicemembertemp", controller.Companyinvoicemembertemp)
	api.Post("companyinvoicemembersync", controller.Companyinvoicemembersync)
	api.Post("companyinvoicegrouppermainan", controller.Companyinvoicegrouppermainan)
	api.Post("companyinvoicelistpermainan", controller.Companyinvoicelistpermainan)
	api.Post("companyinvoicelistpermainanstatus", controller.Companyinvoicelistpermainanstatus)
	api.Post("companyinvoicelistpermainanmember", controller.Companyinvoicelistpermainanmember)
	api.Post("companypasaranconf", controller.Companypasaranconf)
	api.Post("companypasaranonline", controller.Companypasaranonline)
	api.Post("savecompany", controller.Companysave)
	api.Post("savecompanyadmin", controller.Companysaveadmin)
	api.Post("savecompanyupdatepasaran", controller.Companysaveupdatepasaran)
	api.Post("savecompanyupdatepasaranroyaltyfee", controller.Companysaveupdatepasaranroyaltyfee)
	api.Post("savecompanyupdatepasaranline", controller.Companysaveupdatepasaranline)
	api.Post("savecompanypasaranonline", controller.Companysavepasaranonline)
	api.Post("deletecompanypasaranonline", controller.Companydeletepasaranonline)
	api.Post("savecompanypasaran", controller.Companysavecompanypasaran)
	api.Post("updatecompanypasaran432", controller.Companyupdatepasaran432)
	api.Post("updatecompanypasarancolokbebas", controller.Companyupdatepasarancolokbebas)
	api.Post("updatecompanypasarancolokmacau", controller.Companyupdatepasarancolokmacau)
	api.Post("updatecompanypasarancoloknaga", controller.Companyupdatepasarancoloknaga)
	api.Post("updatecompanypasarancolokjitu", controller.Companyupdatepasarancolokjitu)
	api.Post("updatecompanypasaran5050umum", controller.Companyupdatepasaran5050umum)
	api.Post("updatecompanypasaran5050special", controller.Companyupdatepasaran5050special)
	api.Post("updatecompanypasaran5050kombinasi", controller.Companyupdatepasaran5050kombinasi)
	api.Post("updatecompanypasarankombinasi", controller.Companyupdatepasarankombinasi)
	api.Post("updatecompanypasarandasar", controller.Companyupdatepasarandasar)
	api.Post("updatecompanypasaranshio", controller.Companyupdatepasaranshio)
	api.Post("fetchpasaranlimitline", controller.Companyfetchpasaranlimitline)
	api.Post("fetchpasaran432", controller.Companyfetchpasaran432)
	api.Post("fetchpasarancbebas", controller.Companyfetchpasarancbebas)
	api.Post("fetchpasarancmacau", controller.Companyfetchpasarancmacau)
	api.Post("fetchpasarancnaga", controller.Companyfetchpasarancnaga)
	api.Post("fetchpasarancjitu", controller.Companyfetchpasarancjitu)
	api.Post("fetchpasaran5050umum", controller.Companyfetchpasaran5050umum)
	api.Post("fetchpasaran5050special", controller.Companyfetchpasaran5050special)
	api.Post("fetchpasaran5050kombinasi", controller.Companyfetchpasaran5050kombinasi)
	api.Post("fetchpasaranmacau", controller.Companyfetchpasaranmacau)
	api.Post("fetchpasarandasar", controller.Companyfetchpasarandasar)
	api.Post("fetchpasaranshio", controller.Companyfetchpasaranshio)

	api.Post("allpasaran", controller.Pasaran)
	api.Post("pasarandetail", controller.Pasarandetail)
	api.Post("pasarandetailconf", controller.Pasarandetailconf)
	api.Post("savepasaran", controller.Savepasaran)
	api.Post("savepasaranlimitline", controller.Savepasaranlimitline)
	api.Post("savepasaran432", controller.Savepasaran432)
	api.Post("savepasarancbebas", controller.Savepasarancbebas)
	api.Post("savepasarancmacau", controller.Savepasarancmacau)
	api.Post("savepasarancnaga", controller.Savepasarancnaga)
	api.Post("savepasarancjitu", controller.Savepasarancjitu)
	api.Post("savepasaran5050umum", controller.Savepasaran5050umum)
	api.Post("savepasaran5050special", controller.Savepasaran5050special)
	api.Post("savepasaran5050kombinasi", controller.Savepasaran5050kombinasi)
	api.Post("savepasaranmacau", controller.Savepasaranmacau)
	api.Post("savepasarandasar", controller.Savepasarandasar)
	api.Post("savepasaranshio", controller.Savepasaranshio)

	api.Post("invoice", controller.Invoice)
	api.Post("invoicedetail", controller.Invoicedetail)
	api.Post("saveinvoice", controller.Saveinvoice)
	api.Post("saveinvoicewinlose", controller.Saveinvoicewinlose)
	api.Post("saveinvoicewinlosepasaran", controller.Saveinvoicewinlosepasaran)
	api.Post("deleteinvoicewinlosepasaran", controller.Deleteinvoicewinlosepasaran)

	api.Post("setting", controller.Setting)
	api.Post("savesetting", controller.Settingsave)
	api.Post("domain", controller.Domain)
	api.Post("savedomain", controller.Savedomain)
	return app
}
