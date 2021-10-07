package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"threedee/entity"
	"threedee/handler"
	"threedee/testdata/mock"
	"threedee/utility/normalizer"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/suite"
	"github.com/subosito/gotenv"
)

/*
 * Unit Testing HTTP Requests in Go
 *
 * UT must-haves:
 * 1. Declare TestSuite struct including its parts (mock it if needed)
 * 2. SetupTest() to instantiate the testsuite upon go test launch
 * 3. Create TestMethod() that contains the testcases, input datas, mocking mechanisms, and
 *    endpoint call.
 * 4. Create the TestTestSuite function (is a must!)
 * 5. execute "go test ./..." to test all the testable items
 *
 * If you mock a repo/service, create them under "testdata/mock" directory under package mock
 */

var _ = gotenv.Load("../.env")

// 1
type PrintRequestHandlerTestSuite struct {
	suite.Suite
	mockPanelRepo   *mock.MockPrintRequestRepository
	handlerInstance handler.RequestHandler
}

// 2
func (suite *PrintRequestHandlerTestSuite) SetupTest() {
	suite.mockPanelRepo = &mock.MockPrintRequestRepository{}
	suite.handlerInstance = handler.RequestHandler{Repo: suite.mockPanelRepo, Norm: &normalizer.PrintRequestNormalizer{}}
}

//===============================================INDEX========================================================

// 3
func (suite *PrintRequestHandlerTestSuite) TestIndex() {
	var testCase = []struct {
		testcase     string
		isTimeout    bool
		isError      bool
		getAllResult []*entity.PrintRequest
		getAllError  error
	}{
		{
			testcase:     "success",
			isTimeout:    false,
			isError:      false,
			getAllResult: []*entity.PrintRequest{entity.NewPrintRequest(), entity.NewPrintRequest()},
			getAllError:  nil,
		},
		{
			testcase:     "returns error",
			isTimeout:    false,
			isError:      true,
			getAllResult: nil,
			getAllError:  errors.New("[TEST] Failed to Retrieve Data"),
		},
		{
			testcase:     "timeout",
			isTimeout:    true,
			isError:      true,
			getAllResult: nil,
			getAllError:  nil,
		},
	}
	for _, tc := range testCase {
		req, _ := http.NewRequest("GET", "/print-requests", nil)
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		suite.mockPanelRepo.On("GetAll").Return(tc.getAllResult, tc.getAllError).Times(1)

		var err error
		if tc.isTimeout {
			ctx, cancel := context.WithTimeout(req.Context(), -7*time.Hour)
			defer cancel()
			_, err = suite.handlerInstance.Index(responseRecorder, req.WithContext(ctx), nil)
		} else {
			_, err = suite.handlerInstance.Index(responseRecorder, req, nil)
		}

		if tc.isError {
			suite.NotNil(err)
		} else {
			suite.Nil(err)
		}
	}
}

//===============================================SHOW========================================================

func (suite *PrintRequestHandlerTestSuite) TestShow() {
	var testCase = []struct {
		testcase   string
		id         string
		isTimeout  bool
		isError    bool
		showResult *entity.PrintRequest
		showError  error
	}{
		{
			testcase:   "success",
			id:         "1",
			isTimeout:  false,
			isError:    false,
			showResult: &entity.PrintRequest{Id: 1},
			showError:  nil,
		},
		{
			testcase:   "returns error",
			id:         "1",
			isTimeout:  false,
			isError:    true,
			showResult: nil,
			showError:  errors.New("[TEST] Failed to Retrieve Data"),
		},
		{
			testcase:   "timeout",
			id:         "1",
			isTimeout:  true,
			isError:    true,
			showResult: nil,
			showError:  nil,
		},
		{
			testcase:   "id is not a number",
			id:         "a",
			isTimeout:  false,
			isError:    true,
			showResult: nil,
			showError:  nil,
		},
	}
	for _, tc := range testCase {
		req, _ := http.NewRequest("GET", "/print-requests/:id", nil)
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		suite.mockPanelRepo.On("GetById", 1).Return(tc.showResult, tc.showError).Times(1)

		var err error
		if tc.isTimeout {
			ctx, cancel := context.WithTimeout(req.Context(), -7*time.Hour)
			defer cancel()
			_, err = suite.handlerInstance.Show(responseRecorder, req.WithContext(ctx), []httprouter.Param{{Key: "id", Value: tc.id}})
		} else {
			_, err = suite.handlerInstance.Show(responseRecorder, req, []httprouter.Param{{Key: "id", Value: tc.id}})
		}

		if tc.isError {
			suite.NotNil(err)
		} else {
			suite.Nil(err)
		}
	}
}

//===============================================CREATE========================================================

func (suite *PrintRequestHandlerTestSuite) TestCreate() {
	model := entity.PrintRequest{
		ItemName:                "Bertaburan Bunga v2",
		EstimatedWeight:         37.5,
		EstimatedFilamentLength: 5000,
		EstimatedDuration:       9000,
		FileUrl:                 "http://drive.google.com/filez/100",
		Requestor:               "Karim Hartono",
	}
	reqBodyBytes, _ := json.Marshal(model)

	var testCase = []struct {
		testcase     string
		reqBody      []byte
		isTimeout    bool
		isError      bool
		createResult int
		createError  error
		showResult   *entity.PrintRequest
	}{
		{
			testcase:     "success",
			reqBody:      reqBodyBytes,
			isTimeout:    false,
			isError:      false,
			createResult: 1,
			createError:  nil,
			showResult:   &entity.PrintRequest{Id: 1},
		},
		{
			testcase:     "returns error",
			reqBody:      reqBodyBytes,
			isTimeout:    false,
			isError:      true,
			createResult: 0,
			createError:  errors.New("[TEST] Failed to Insert Data"),
			showResult:   nil,
		},
		{
			testcase:     "timeout",
			reqBody:      reqBodyBytes,
			isTimeout:    true,
			isError:      true,
			createResult: 0,
			createError:  nil,
			showResult:   nil,
		},
	}
	for _, tc := range testCase {
		req, _ := http.NewRequest("POST", "/print-requests", strings.NewReader(string(tc.reqBody)))
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		suite.mockPanelRepo.On("Insert", &model).Return(tc.createResult, tc.createError).Times(1)
		suite.mockPanelRepo.On("GetById", tc.createResult).Return(tc.showResult, nil).Times(1)

		var err error
		if tc.isTimeout {
			ctx, cancel := context.WithTimeout(req.Context(), -7*time.Hour)
			defer cancel()
			_, err = suite.handlerInstance.Create(responseRecorder, req.WithContext(ctx), nil)
		} else {
			_, err = suite.handlerInstance.Create(responseRecorder, req, nil)
		}

		if tc.isError {
			suite.NotNil(err)
		} else {
			suite.Nil(err)
		}
	}
}

//===============================================UPDATE========================================================

func (suite *PrintRequestHandlerTestSuite) TestUpdate() {
	model := entity.PrintRequest{
		ItemName:                "Bertaburan Bunga v2",
		EstimatedWeight:         37.5,
		EstimatedFilamentLength: 5000,
		EstimatedDuration:       9000,
		FileUrl:                 "http://drive.google.com/filez/100",
		Requestor:               "Karim Hartono",
	}
	reqBodyBytes, _ := json.Marshal(model)

	showModel := entity.PrintRequest{
		Id:                      1,
		ItemName:                "Bertaburan Bunga v2",
		EstimatedWeight:         37.5,
		EstimatedFilamentLength: 5000,
		EstimatedDuration:       9000,
		FileUrl:                 "http://drive.google.com/filez/100",
		Requestor:               "Karim Hartono",
	}

	var testCase = []struct {
		testcase     string
		id           string
		reqBody      []byte
		isTimeout    bool
		isError      bool
		updateResult bool
		updateError  error
		showResult   *entity.PrintRequest
	}{
		{
			testcase:     "success",
			id:           "1",
			reqBody:      reqBodyBytes,
			isTimeout:    false,
			isError:      false,
			updateResult: true,
			updateError:  nil,
			showResult:   &showModel,
		},
		{
			testcase:     "returns error",
			id:           "1",
			reqBody:      reqBodyBytes,
			isTimeout:    false,
			isError:      true,
			updateResult: false,
			updateError:  errors.New("[TEST] Failed to Insert Data"),
			showResult:   &showModel,
		},
		{
			testcase:     "timeout",
			id:           "1",
			reqBody:      reqBodyBytes,
			isTimeout:    true,
			isError:      true,
			updateResult: false,
			updateError:  nil,
			showResult:   nil,
		},
	}
	for _, tc := range testCase {
		req, _ := http.NewRequest("PUT", "/print-requests/:id", strings.NewReader(string(tc.reqBody)))
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		suite.mockPanelRepo.On("GetById", 1).Return(tc.showResult, nil).Times(1)
		suite.mockPanelRepo.On("Update", &showModel).Return(tc.updateResult, tc.updateError).Times(1)

		var err error
		if tc.isTimeout {
			ctx, cancel := context.WithTimeout(req.Context(), -7*time.Hour)
			defer cancel()
			_, err = suite.handlerInstance.Update(responseRecorder, req.WithContext(ctx), []httprouter.Param{{Key: "id", Value: tc.id}})
		} else {
			_, err = suite.handlerInstance.Update(responseRecorder, req, []httprouter.Param{{Key: "id", Value: tc.id}})
		}

		if tc.isError {
			suite.NotNil(err)
		} else {
			suite.Nil(err)
		}
	}
}

//===============================================TESTING========================================================

// 4
func TestPrintRequestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PrintRequestHandlerTestSuite))
}
