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

type GetTranslatedPokemonDescriptionTestSuite struct {
	suite.Suite
	sut              *handlers.GetTranslatedPokemonDescriptionHandler
	responseRecorder *httptest.ResponseRecorder
	requestBuilder   *http2.Request
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) TearDownTest() {
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) SetupTest() {
	suite.sut = handlers.NewGetTranslatedPokemonDescriptionHandler()
	suite.requestBuilder = httptest.NewRequest(http2.MethodGet, "/pokemon/translated/metapod", nil)
	suite.responseRecorder = httptest.NewRecorder()
}

func TestGetTranslatedPokemonDescriptionTestSuiteRunner(t *testing.T) {
	suite.Run(t, new(GetTranslatedPokemonDescriptionTestSuite))
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) TestReturns_StatusOk_On_Success_With_Correct_Response() {
	// if pokemon habitat is not cave, and it is not legendary, return shakespeare translation
	expectedName := "metapod"
	expectedTranslatedDescription := "Coequal though 't is encased in a sturdy shell,  the corse inside is tender. 't can’t withstand a harsh attack."
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
	suite.Assert().Equal(expectedTranslatedDescription, response.Description)
	suite.Assert().Equal(expectedHabitat, response.Habitat)
	suite.Assert().Equal(expectedLegendaryStatus, response.IsLegendary)
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) TestReturns_StatusOk_On_Success_With_Yoda_Translation_In_Response() {
	// if pokemon habitat is cave return yoda translation
	expectedTranslatedDescription := "Almost the same as mew’s,  its dna is.However,Vastly different,  its size and disposition are."

	request := httptest.NewRequest(http2.MethodGet, "/pokemon/translated/mewtwo", nil)

	suite.sut.Handle(suite.responseRecorder, request)
	suite.Assert().Equal(http2.StatusOK, suite.responseRecorder.Code)

	body, err := ioutil.ReadAll(suite.responseRecorder.Body)
	suite.Assert().Nil(err)

	var response handlers.GetPokemonInformationResponse
	err = json.Unmarshal(body, &response)
	suite.Assert().Nil(err)

	suite.Assert().Equal(expectedTranslatedDescription, response.Description)
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) TestReturns_StatusNotFound_If_Pokemon_Does_Not_Exist() {
	request := httptest.NewRequest(http2.MethodGet, "/pokemon/fakepokemon", nil)

	suite.sut.Handle(suite.responseRecorder, request)

	suite.Assert().Equal(http2.StatusNotFound, suite.responseRecorder.Code)
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) TestReturns_InternalServerErr_If_No_Name_Supplied_In_RequestUrl() {
	request := httptest.NewRequest(http2.MethodGet, "/pokemon/translated/", nil)

	suite.sut.Handle(suite.responseRecorder, request)

	suite.Assert().Equal(http2.StatusInternalServerError, suite.responseRecorder.Code)
}

func (suite *GetTranslatedPokemonDescriptionTestSuite) TestReturns_Error_If_Decoding_Fails() {
	request := suite.requestBuilder

	suite.sut.Handle(suite.responseRecorder, request)
	suite.Assert().Equal(http2.StatusOK, suite.responseRecorder.Code)

	body := []byte(`{"name":bad_value}`)

	var response handlers.GetPokemonInformationResponse
	err := json.Unmarshal(body, &response)
	suite.Assert().NotNil(err)
}
