package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
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

//===============================================TESTING========================================================

// 4
func TestPrintRequestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PrintRequestHandlerTestSuite))
}
