package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/cyanrad/university/db/mock"
	db "github.com/cyanrad/university/db/sqlc"
	"github.com/cyanrad/university/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateUserRequest(t *testing.T) {
	user, pass := generateUserData(t)

	// >> different test cases
	// TODO: implement db mocking
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{ // >> everything should be fine and dandy
			name: "Ok",
			body: gin.H{
				"id":        user.ID,
				"password":  pass,
				"full_name": user.FullName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// >> we expect CreateUser to run once, and return the user
				arg := db.CreateUserParams{
					ID:       user.ID,
					FullName: user.FullName,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUser(arg, pass)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				// >> Since data is valid, response should be <200 OK>
				require.Equal(t, http.StatusOK, recorder.Code)

				// >> comparing resp and generated user data
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalServererror",
			body: gin.H{
				"id":        user.ID,
				"password":  pass,
				"full_name": user.FullName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateName",
			body: gin.H{
				"id":        user.ID,
				"password":  pass,
				"full_name": user.FullName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"id":        "ajhlf34324s",
				"password":  pass,
				"full_name": user.FullName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ShortPassword",
			body: gin.H{
				"id":        user.ID,
				"password":  "123",
				"full_name": user.FullName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	// >> running tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// >> creating the mock ecosystem controller
			ctrl := gomock.NewController(t)

			// >> creating the mock store
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// >> creating JSON from user data
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// >> creating test server
			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			// >> Creating the request
			url := pathsURI["createUser"]
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			// >> sending the request
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

// >> cheacking if the user from the resp is the same as the one that was generated
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body) // reading the resp byts
	require.NoError(t, err)

	// >> creating the user obj & filling with resp data
	var gotUser createUserResponse
	err = json.Unmarshal(data, &gotUser)

	// >> testing created and resp user data validity
	require.NoError(t, err)
	require.Equal(t, user.ID, gotUser.ID)
	require.Equal(t, user.FullName, gotUser.FullName)
	//require.Empty(t, gotUser.HashedPassword)
}

// >> generates random data for a random user
func generateUserData(t *testing.T) (user db.User, password string) {
	password = util.RandomString(8)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		ID:             util.RandomID(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomString(10),
	}
	return
}
