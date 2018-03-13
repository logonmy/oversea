package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "oversea/routers"
	_"github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"oversea/utils"
	"github.com/astaxie/beego/logs"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}


// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

func TestPassword(t *testing.T) {
	salt := utils.NewNoDashUUID()
	password := "admin"
	pwd := utils.MD5(password + salt)

	logs.Info("salt=%v , password=%v" , salt, pwd)
}