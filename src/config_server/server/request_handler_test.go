package server_test

import (
	. "config_server/server"
	. "config_server/store/fakes"
    . "config_server/server/fakes"
	. "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    . "config_server/types/fakes"
    "errors"
    "net/http"
    "net/http/httptest"
    "strings"

    "config_server/types"
)

type BadMockStore struct{}

func (store BadMockStore) Get(key string) (string, error) {
	return "", errors.New("")
}

func (store BadMockStore) Put(key string, value string) error {
	return errors.New("")
}

var _ = Describe("RequestHandlerConcrete", func() {

	Describe("Given a nil store", func() {

		Context("creating the requestHandler", func() {
			It("should return an error", func() {
				putReq, _ := http.NewRequest("PUT", "/v1/config/bla", strings.NewReader("{\"value\":\"blabla\"}"))
				putRecorder := httptest.NewRecorder()

				requestHandler := NewRequestHandler(nil, types.NewValueGeneratorConcrete())
				requestHandler.ServeHTTP(putRecorder, putReq)

				Expect(putRecorder.Code).To(Equal(http.StatusInternalServerError))
                Expect(putRecorder.Body.String()).To(Equal("DB Store is nil\n"))
			})
		})
	})

	Describe("Given a server with store", func() {

		var requestHandler http.Handler
		var mockTokenValidator *FakeTokenValidator
        var mockStore *FakeStore
        var mockValueGeneratorFactory *FakeValueGeneratorFactory

		BeforeEach(func() {
            mockTokenValidator = &FakeTokenValidator{}
			mockStore = &FakeStore{}
            mockValueGeneratorFactory = &FakeValueGeneratorFactory{}
			requestHandler = NewRequestHandler(mockStore, mockValueGeneratorFactory)
		})

		Context("when URL path is invalid", func() {

			It("should return 404 Not Found for invalid paths", func() {
				invalidPaths := []string{"/v1/config/test/case", "/v1"}

				for _, path := range invalidPaths {
					req, _ := http.NewRequest("GET", path, nil)
					recorder := httptest.NewRecorder()
					requestHandler.ServeHTTP(recorder, req)

					Expect(recorder.Code).To(Equal(http.StatusNotFound))
				}
			})

			It("should return 404 Not Found for other methods", func() {
				invalidMethods := [...]string{"DELETE", "PATCH"}
				http.NewRequest("PUT", "/v1/config/bla", strings.NewReader("value=blabla"))

				for _, method := range invalidMethods {
					req, _ := http.NewRequest(method, "/v1/config/bla", nil)
                    req.Header.Set("Authorization", "bearer fake-auth-header")

					recorder := httptest.NewRecorder()
					requestHandler.ServeHTTP(recorder, req)

					Expect(recorder.Code).To(Equal(http.StatusNotFound))
				}
			})

			It("should return 404 Not Found when key is not provided for fetch", func() {
				req, _ := http.NewRequest("GET", "/v1/config/", nil)
                req.Header.Set("Authorization", "bearer fake-auth-header")

				getRecorder := httptest.NewRecorder()
				requestHandler.ServeHTTP(getRecorder, req)

				Expect(getRecorder.Code).To(Equal(http.StatusNotFound))
			})

			It("should return 404 Not Found when key is not provided for update", func() {
				req, _ := http.NewRequest("PUT", "/v1/config/", nil)
                req.Header.Set("Authorization", "bearer fake-auth-header")

				getRecorder := httptest.NewRecorder()
				requestHandler.ServeHTTP(getRecorder, req)

				Expect(getRecorder.Code).To(Equal(http.StatusNotFound))
			})
		})

		Context("when URL path is valid", func() {

            It("should store values as JSON", func() {
                req, _ := http.NewRequest("PUT", "/v1/config/bla", strings.NewReader("{\"value\":\"str\"}"))
                req.Header.Set("Authorization", "bearer fake-auth-header")

                putRecorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(putRecorder, req)

                key, value := mockStore.PutArgsForCall(0)

                Expect(key).To(Equal("bla"))
                Expect(value).To(Equal("{\"path\":\"bla\",\"value\":\"str\"}"))
            })

            It("should return 200 Status OK when an integer value is added", func() {
                req, _ := http.NewRequest("PUT", "/v1/config/bla", strings.NewReader("{\"value\":1}"))
                req.Header.Set("Authorization", "bearer fake-auth-header")

                putRecorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(putRecorder, req)

                Expect(putRecorder.Code).To(Equal(http.StatusOK))
            })

            It("should return 200 Status OK when a string value is added", func() {
                req, _ := http.NewRequest("PUT", "/v1/config/bla", strings.NewReader("{\"value\":\"str\"}"))
                req.Header.Set("Authorization", "bearer fake-auth-header")

                putRecorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(putRecorder, req)

                Expect(putRecorder.Code).To(Equal(http.StatusOK))
            })

            It("should return 200 OK when config value is updated", func() {
                req, _ := http.NewRequest("PUT", "/v1/config/bla", strings.NewReader("{\"value\":\"blabla\"}"))
                req.Header.Set("Authorization", "bearer fake-auth-header")

                recorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(recorder, req)

                Expect(recorder.Code).To(Equal(http.StatusOK))
            })

            It("should return 200 OK when valid key is retrieved", func() {
                mockStore.GetReturns("{\"path\":\"bla\",\"value\":\"blabla\"}", nil)

                getReq, _ := http.NewRequest("GET", "/v1/config/bla/", nil)
                getReq.Header.Set("Authorization", "bearer fake-auth-header")

                getRecorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(getRecorder, getReq)

                Expect(getRecorder.Code).To(Equal(http.StatusOK))
                Expect(getRecorder.Body.String()).To(Equal("{\"path\":\"bla\",\"value\":\"blabla\"}"))
            })

            It("should return 404 Not Found when key is not found", func() {
                req, _ := http.NewRequest("GET", "/v1/config/test", nil)
                req.Header.Set("Authorization", "bearer fake-auth-header")

                getRecorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(getRecorder, req)

                Expect(getRecorder.Code).To(Equal(http.StatusNotFound))
            })

            It("should return 400 Bad Request when value is not provided for update", func() {
                req, _ := http.NewRequest("PUT", "/v1/config/key", nil)
                req.Header.Set("Authorization", "bearer fake-auth-header")

                getRecorder := httptest.NewRecorder()
                requestHandler.ServeHTTP(getRecorder, req)

                Expect(getRecorder.Code).To(Equal(http.StatusBadRequest))
            })

            Context("Password generation", func() {
                It("should return generated password", func() {

                    generator := &FakeValueGenerator{}
                    generator.GenerateReturns("bXgsZD!aNukh$#sSRdBh", nil)

                    mockValueGeneratorFactory.GetGeneratorReturns(generator, nil)

                    postReq, _ := http.NewRequest("POST", "/v1/config/bla/", strings.NewReader("{\"type\":\"password\",\"parameters\":{}}"))
                    postReq.Header.Set("Authorization", "bearer fake-auth-header")

                    getRecorder := httptest.NewRecorder()
                    requestHandler.ServeHTTP(getRecorder, postReq)

                    Expect(getRecorder.Code).To(Equal(http.StatusOK))

                    key, value := mockStore.PutArgsForCall(0)
                    Expect(key).To(Equal("bla"))
                    Expect(value).To(Equal("{\"path\":\"bla\",\"value\":\"bXgsZD!aNukh$#sSRdBh\"}"))
                })

                It("should not generate a password if one already exists", func() {
                    mockStore.GetStub = func(key string) (string, error) {
                        if key == "bla" {
                            return "{\"path\":\"bla\",\"value\":\"value\"}", nil
                        }
                        return "", nil
                    }

                    postReq, _ := http.NewRequest("POST", "/v1/config/bla/", strings.NewReader("{\"type\":\"password\",\"parameters\":{}}"))
                    postReq.Header.Set("Authorization", "bearer fake-auth-header")

                    getRecorder := httptest.NewRecorder()
                    requestHandler.ServeHTTP(getRecorder, postReq)

                    Expect(getRecorder.Code).To(Equal(http.StatusOK))
                    Expect(getRecorder.Body.String()).To(Equal("{\"path\":\"bla\",\"value\":\"value\"}"))
                    Expect(mockValueGeneratorFactory.GetGeneratorCallCount()).To(Equal(0))
                })
            })
		})
	})
})
