// It takes a string and returns a random string of the same length
package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	adminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjY5NTkyMTExODUsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiUGhvbmUiOiIwMDAwMDAwMDAwIiwiUm9sZSI6ImFkbWluIn0.QqrBgRE6GLAtB7GmuKHkLJ7rKUWtJ1ebB9S6em3lMXI"

	driverPhone    = "2222222222"
	driverPassword = "password"
	driverName     = randomStringNumber(10)
	driverRole     = "driver"
	driverDevice   = "dm38o087QRm6Bu8jTXB5Qq:APA91bGjz940Ynf-49xwX0wl0vn67QL3pNKpiKlUJ_06NV5KtOYxSMSY5RRMS3vVgIKU1x22sPcE63OLf5ZQK_ttkD4xt1Ved4yaNcSitxPN9GRf2DKWTD6pA8R3vz1o1ju8RAsyDee3"
	driverToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI2OTg5NTEsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6OCwiUGhvbmUiOiIyMjIyMjIyMjIyIiwiUm9sZSI6ImRyaXZlciJ9.kKQhJe9JJcpJiy_WhMHwwsEeBqCoRhBdtgLEVRG007s"

	passengerPhone    = "1111111111"
	passengerPassword = "password"
	passengerName     = randomStringNumber(10)
	passengerRole     = "passenger"
	passengerDevice   = "dm38o087QRm6Bu8jTXB5Qq:APA91bGjz940Ynf-49xwX0wl0vn67QL3pNKpiKlUJ_06NV5KtOYxSMSY5RRMS3vVgIKU1x22sPcE63OLf5ZQK_ttkD4xt1Ved4yaNcSitxPN9GRf2DKWTD6pA8R3vz1o1ju8RAsyDee3"
	passengerToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI2OTg5MDMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MiwiUGhvbmUiOiIxMTExMTExMTExIiwiUm9sZSI6InBhc3NlbmdlciJ9.1QfoaCb1QSWcmCIrsEOlkh5nnGe9oYtulfoyxV0-rEc"
)

func TestRegister(t *testing.T) {

	testCases := []struct {
		name string
		data map[string]string
	}{
		{
			name: "register driver",
			data: map[string]string{
				"phone":    randomStringNumber(10) + "1",
				"password": driverPassword,
				"name":     driverName,
				"role":     driverRole,
			},
		},
		{
			name: "register passenger",
			data: map[string]string{
				"phone":    randomStringNumber(10) + "1",
				"password": passengerPassword,
				"name":     passengerName,
				"role":     passengerRole,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			url := "/auth/register"
			body := tc.data
			data, err := json.Marshal(body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			r.ServeHTTP(recorder, request)

			require.Equal(t, http.StatusOK, recorder.Code)

			rsp, err := ioutil.ReadAll(recorder.Body)
			require.NoError(t, err)

			var res response
			err = json.Unmarshal(rsp, &res)
			require.NoError(t, err)
			require.Equal(t, http.StatusCreated, res.Status)
		})
	}
}

func TestLogin(t *testing.T) {

	testCases := []struct {
		name string
		data map[string]string
		role string
	}{
		{
			name: "login driver",
			data: map[string]string{
				"phone":        driverPhone,
				"password":     driverPassword,
				"device_token": driverDevice,
			},
			role: driverRole,
		},
		{
			name: "login passenger",
			data: map[string]string{
				"phone":        passengerPhone,
				"password":     passengerPassword,
				"device_token": passengerDevice,
			},
			role: passengerRole,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/auth/login/%s", tc.role)
			body := tc.data
			data, err := json.Marshal(body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			r.ServeHTTP(recorder, request)

			require.Equal(t, http.StatusOK, recorder.Code)

			rsp, err := ioutil.ReadAll(recorder.Body)
			require.NoError(t, err)

			var res response
			err = json.Unmarshal(rsp, &res)
			require.NoError(t, err)
			require.Empty(t, res.Error)
			require.Equal(t, http.StatusOK, res.Status)
		})
	}
}

func TestVerify(t *testing.T) {

	testCases := []struct {
		name string
		data map[string]string
	}{
		{
			name: "verify",
			data: map[string]string{
				"phone": driverPhone,
				"token": adminToken,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			url := "/auth/verify"
			body := tc.data
			data, err := json.Marshal(body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			r.ServeHTTP(recorder, request)

			require.Equal(t, http.StatusOK, recorder.Code)

			rsp, err := ioutil.ReadAll(recorder.Body)
			require.NoError(t, err)

			var res response
			err = json.Unmarshal(rsp, &res)
			require.NoError(t, err)
		})
	}
}
