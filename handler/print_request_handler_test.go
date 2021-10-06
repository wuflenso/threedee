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

	"github.com/stretchr/testify/suite"
	"github.com/subosito/gotenv"
)

var _ = gotenv.Load("../.env")

type PrintRequestHandlerTestSuite struct {
	suite.Suite
	mockPanelRepo   *mock.MockPrintRequestRepository
	handlerInstance handler.RequestHandler
}

func (suite *PrintRequestHandlerTestSuite) SetupTest() {
	suite.mockPanelRepo = &mock.MockPrintRequestRepository{}
	suite.handlerInstance = handler.RequestHandler{Repo: suite.mockPanelRepo, Norm: &normalizer.PrintRequestNormalizer{}}
}

//===============================================INDEX========================================================

func (suite *PrintRequestHandlerTestSuite) TestIndex() {
	var testCase = []struct {
		testcase     string
		isTimeout    bool
		isError      bool
		GetAllResult []*entity.PrintRequest
		GetAllError  error
	}{
		{
			testcase:     "success",
			isTimeout:    false,
			isError:      false,
			GetAllResult: []*entity.PrintRequest{entity.NewPrintRequest(), entity.NewPrintRequest()},
			GetAllError:  nil,
		},
		{
			testcase:     "returns error",
			isTimeout:    false,
			isError:      true,
			GetAllResult: nil,
			GetAllError:  errors.New("[TEST] Failed to Retrieve Data"),
		},
		{
			testcase:     "timeout",
			isTimeout:    true,
			isError:      true,
			GetAllResult: nil,
			GetAllError:  nil,
		},
	}
	for _, tc := range testCase {
		req, _ := http.NewRequest("GET", "/print-requests", nil)
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		suite.mockPanelRepo.On("GetAll").Return(tc.GetAllResult, tc.GetAllError).Times(1)

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

//===============================================TESTING========================================================

func TestPrintRequestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PrintRequestHandlerTestSuite))
}
