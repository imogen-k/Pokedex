package handlers_test

import (
	"Pokedex/handlers"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	http2 "net/http"
	"net/http/httptest"
	"testing"
)

type GetPokemonInformationTestSuite struct {
	suite.Suite
	sut              *handlers.GetPokemonInformationHandler
	responseRecorder *httptest.ResponseRecorder
	requestBuilder   *http2.Request
}

func (suite *GetPokemonInformationTestSuite) SetupTest() {
	suite.sut = handlers.NewGetPokemonInformationHandler()
	suite.requestBuilder = httptest.NewRequest(http2.MethodGet, "/pokemon/metapod", nil)
	suite.responseRecorder = httptest.NewRecorder()
}

func TestGetPokemonInformationTestSuiteRunner(t *testing.T) {
	suite.Run(t, new(GetPokemonInformationTestSuite))
}

func (suite *GetPokemonInformationTestSuite) TestReturns_StatusOk_On_Success_With_Correct_Response_Values() {
	expectedName := "metapod"
	expectedDescription := "Even though it is encased in a sturdy shell, the body inside is tender. It can’t withstand a harsh attack."
	expectedHabitat := "forest"
	expectedLegendaryStatus := false

	request := suite.requestBuilder

	suite.sut.Handle(suite.responseRecorder, request)
	suite.Assert().Equal(http2.StatusOK, suite.responseRecorder.Code)

	body, err := ioutil.ReadAll(suite.responseRecorder.Body)
	suite.Assert().Nil(err)

	var response handlers.GetPokemonInformationResponse
	err = json.Unmarshal(body, &response)
	suite.Assert().Nil(err)

	suite.Assert().Equal(expectedName, response.Name)
	suite.Assert().Equal(expectedDescription, response.Description)
	suite.Assert().Equal(expectedHabitat, response.Habitat)
	suite.Assert().Equal(expectedLegendaryStatus, response.IsLegendary)
}

func (suite *GetPokemonInformationTestSuite) TestReturns_InternalServerErr_If_No_Name_Supplied_In_RequestUrl() {
	request := httptest.NewRequest(http2.MethodGet, "/pokemon/", nil)

	suite.sut.Handle(suite.responseRecorder, request)

	suite.Assert().Equal(http2.StatusInternalServerError, suite.responseRecorder.Code)
}

func (suite *GetPokemonInformationTestSuite) TestReturns_StatusNotFound_If_Pokemon_Does_Not_Exist() {
	request := httptest.NewRequest(http2.MethodGet, "/pokemon/fakepokemon", nil)

	suite.sut.Handle(suite.responseRecorder, request)

	suite.Assert().Equal(http2.StatusNotFound, suite.responseRecorder.Code)
}

func (suite *GetPokemonInformationTestSuite) TestReturns_Error_If_Decoding_Fails() {
	request := suite.requestBuilder

	suite.sut.Handle(suite.responseRecorder, request)

	suite.Assert().Equal(http2.StatusOK, suite.responseRecorder.Code)

	body := []byte(`{"name":bad_value}`)

	var response handlers.GetPokemonInformationResponse
	err := json.Unmarshal(body, &response)
	suite.Assert().NotNil(err)
}
