package helpers

import (
	"bytes"
	"go-boilerplate/internal/pkg/responses"
	"io"
	"net/http"
)

type httpHelper struct{}

func (h *httpHelper) CloneRequestBody(req *http.Request) (clonedBody []byte, err error) {
	isHasNoBody := (req.Body == http.NoBody || req.Body == nil)
	if isHasNoBody {
		return
	}
	defer req.Body.Close()

	clonedBody, err = io.ReadAll(req.Body)
	if err != nil {
		err = responses.NewError().
			WithCode(http.StatusInternalServerError).
			WithError(err).
			WithMessage("Failed to read request body.")
		return
	}

	bodyReader := bytes.NewReader(clonedBody)
	req.Body = io.NopCloser(bodyReader)
	return
}
