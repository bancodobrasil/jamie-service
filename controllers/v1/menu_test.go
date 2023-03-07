package v1

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func mockGin() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// test request, must instantiate a request first
	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header), // if you need to test headers
	}
	// finally set the request to the gin context
	c.Request = req

	return c, w
}

// type EvalServiceTestEvalHandlerWithoutMenuAndVersion struct {
// 	services.IEval
// 	t *testing.T
// }

// func (s EvalServiceTestEvalHandlerWithoutMenuAndVersion) LoadRemoteGRL(uuid string, version string) error {
// 	if uuid != services.DefaultMenuName || version != services.CurrentMenuVersion {
// 		s.t.Error("Did not load the default")
// 	}
// 	return nil
// }

// func (s EvalServiceTestEvalHandlerWithoutMenuAndVersion) GetKnowledgeLibrary() *ast.KnowledgeLibrary {
// 	return ast.NewKnowledgeLibrary()
// }

// func TestEvalHandlerWithoutMenuAndVersion(t *testing.T) {

// 	services.EvalService = EvalServiceTestEvalHandlerWithoutMenuAndVersion{
// 		t: t,
// 	}

// 	c, r := mockGin()
// 	EvalHandler()(c)
// 	gotStatus := r.Result().Status
// 	expectedStatus := "404 Not Found"

// 	if gotStatus != expectedStatus {
// 		t.Error("got error on request evalHandler func")
// 	}

// 	gotBody := r.Body.String()
// 	expectedBody := "Menu or version not founded!"

// 	if gotBody != expectedBody {
// 		t.Error("got error on request evalHandler func")
// 	}
// }

// type EvalServiceTestEvalHandlerLoadError struct {
// 	services.IEval
// 	t *testing.T
// }

// func (s EvalServiceTestEvalHandlerLoadError) LoadRemoteGRL(uuid string, version string) error {
// 	return fmt.Errorf("mock load error")
// }

// func (s EvalServiceTestEvalHandlerLoadError) GetKnowledgeLibrary() *ast.KnowledgeLibrary {
// 	return ast.NewKnowledgeLibrary()
// }

// func TestEvalHandlerLoadError(t *testing.T) {

// 	services.EvalService = EvalServiceTestEvalHandlerLoadError{
// 		t: t,
// 	}

// 	c, r := mockGin()
// 	EvalHandler()(c)
// 	gotStatus := r.Code
// 	expectedStatus := http.StatusInternalServerError

// 	if gotStatus != expectedStatus {
// 		t.Error("got error on request evalHandler func")
// 	}

// 	gotBody := r.Body.String()
// 	expectedBody := "Error on load menu and/or version"

// 	if gotBody != expectedBody {
// 		t.Error("got error on request evalHandler func")
// 	}
// }

// type EvalServiceTestEvalHandlerWithDefaultMenu struct {
// 	services.IEval
// 	kl *ast.KnowledgeLibrary
// 	t  *testing.T
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenu) LoadRemoteGRL(uuid string, version string) error {
// 	return nil
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenu) GetKnowledgeLibrary() *ast.KnowledgeLibrary {
// 	return s.kl
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenu) Eval(ctx *types.Context, menu *ast.Menu) (*types.Result, error) {
// 	return types.NewResult(), nil
// }

// func TestEvalHandlerWithDefaultMenu(t *testing.T) {

// 	services.EvalService = EvalServiceTestEvalHandlerWithDefaultMenu{
// 		t:  t,
// 		kl: ast.NewKnowledgeLibrary(),
// 	}

// 	drls := `
// 		rule DefaultValues salience 10 {
// 			when
// 				true
// 			then
// 				Retract("DefaultValues");
// 		}
// 	`

// 	ruleBuilder := builder.NewRuleBuilder(services.EvalService.GetKnowledgeLibrary())
// 	bs := pkg.NewBytesResource([]byte(drls))
// 	ruleBuilder.BuildRuleFromResource(services.DefaultMenuName, services.CurrentMenuVersion, bs)

// 	c, r := mockGin()

// 	stringReader := strings.NewReader("{}")
// 	c.Request.Body = io.NopCloser(stringReader)

// 	EvalHandler()(c)
// 	gotStatus := r.Code
// 	expectedStatus := http.StatusOK

// 	if gotStatus != expectedStatus {
// 		t.Error("got error on request evalHandler func")
// 	}

// 	gotBody := r.Body.String()
// 	expectedBody := "{}"

// 	if gotBody != expectedBody {
// 		t.Error("got error on request evalHandler func")
// 	}
// }

// type EvalServiceTestEvalHandlerWithDefaultMenuAndWrongJSON struct {
// 	services.IEval
// 	kl *ast.KnowledgeLibrary
// 	t  *testing.T
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenuAndWrongJSON) LoadRemoteGRL(uuid string, version string) error {
// 	return nil
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenuAndWrongJSON) GetKnowledgeLibrary() *ast.KnowledgeLibrary {
// 	return s.kl
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenuAndWrongJSON) Eval(ctx *types.Context, menu *ast.Menu) (*types.Result, error) {
// 	return types.NewResult(), nil
// }

// func TestEvalHandlerWithDefaultMenuAndWrongJSON(t *testing.T) {

// 	services.EvalService = EvalServiceTestEvalHandlerWithDefaultMenuAndWrongJSON{
// 		t:  t,
// 		kl: ast.NewKnowledgeLibrary(),
// 	}

// 	drls := `
// 		rule DefaultValues salience 10 {
// 			when
// 				true
// 			then
// 				Retract("DefaultValues");
// 		}
// 	`

// 	ruleBuilder := builder.NewRuleBuilder(services.EvalService.GetKnowledgeLibrary())
// 	bs := pkg.NewBytesResource([]byte(drls))
// 	ruleBuilder.BuildRuleFromResource(services.DefaultMenuName, services.CurrentMenuVersion, bs)

// 	c, r := mockGin()

// 	stringReader := strings.NewReader("")
// 	c.Request.Body = io.NopCloser(stringReader)

// 	EvalHandler()(c)
// 	gotStatus := r.Code
// 	expectedStatus := http.StatusInternalServerError

// 	if gotStatus != expectedStatus {
// 		t.Error("got error on request evalHandler func")
// 	}

// 	gotBody := r.Body.String()
// 	expectedBody := "Error on json decode"

// 	if gotBody != expectedBody {
// 		t.Error("we expect error and the didn't came out")
// 	}
// }

// type EvalServiceTestEvalHandlerWithDefaultMenuEvalError struct {
// 	services.IEval
// 	kl *ast.KnowledgeLibrary
// 	t  *testing.T
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenuEvalError) LoadRemoteGRL(uuid string, version string) error {
// 	return nil
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenuEvalError) GetKnowledgeLibrary() *ast.KnowledgeLibrary {
// 	return s.kl
// }

// func (s EvalServiceTestEvalHandlerWithDefaultMenuEvalError) Eval(ctx *types.Context, menu *ast.Menu) (*types.Result, error) {
// 	return nil, fmt.Errorf("mock error")
// }

// func TestEvalHandlerWithDefaultMenuEvalError(t *testing.T) {

// 	services.EvalService = EvalServiceTestEvalHandlerWithDefaultMenuEvalError{
// 		t:  t,
// 		kl: ast.NewKnowledgeLibrary(),
// 	}

// 	drls := `
// 		rule DefaultValues salience 10 {
// 			when
// 				true
// 			then
// 				Retract("DefaultValues");
// 		}
// 	`

// 	ruleBuilder := builder.NewRuleBuilder(services.EvalService.GetKnowledgeLibrary())
// 	bs := pkg.NewBytesResource([]byte(drls))
// 	ruleBuilder.BuildRuleFromResource(services.DefaultMenuName, services.CurrentMenuVersion, bs)

// 	c, r := mockGin()

// 	stringReader := strings.NewReader("{}")
// 	c.Request.Body = io.NopCloser(stringReader)

// 	EvalHandler()(c)
// 	gotStatus := r.Code
// 	expectedStatus := http.StatusInternalServerError

// 	if gotStatus != expectedStatus {
// 		t.Error("got error on request evalHandler func")
// 	}

// 	gotBody := r.Body.String()
// 	expectedBody := "Error on eval"

// 	if gotBody != expectedBody {
// 		t.Error("we expect error and the didn't came out")
// 	}

// }
