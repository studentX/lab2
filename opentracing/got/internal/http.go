package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	ot "github.com/opentracing/opentracing-go"
)

type (
	// Response a generic Quest response
	Response struct {
		Status string `json:"status"`
	}

	// Quest represents a knight's quest
	Quest struct {
		Knight string `json:"knight"`
	}
)

// Call a remote http service
func Call(ctx context.Context, method, url string, payload io.Reader, res interface{}) error {
	p := ot.SpanFromContext(ctx)
	s := ot.StartSpan(FuncName(0), ot.ChildOf(p.Context()))
	defer s.Finish()
	s.SetTag("http.method", method)
	s.SetTag("http.url", url)

	sx := ot.ContextWithSpan(ctx, s)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		SpanError(sx, err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	if err = ot.GlobalTracer().Inject(p.Context(), ot.HTTPHeaders, ot.HTTPHeadersCarrier(req.Header)); err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		SpanError(sx, fmt.Errorf("remote call failed %s", err))
		return err
	}

	if resp.StatusCode != http.StatusOK {
		e, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			SpanError(sx, err)
			return err
		}
		SpanError(sx, fmt.Errorf("call `%s failed (%d)-%s", url, resp.StatusCode, string(e)))
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		SpanError(sx, err)
		return err
	}
	return nil
}

// SpanError decorates span with error
func SpanError(ctx context.Context, err error) {
	s := ot.SpanFromContext(ctx)
	s.SetTag("error", true)
	s.LogKV("event", "error", "message", err)
}

// WriteErrOut formulate err response and decorte span
func WriteErrOut(ctx context.Context, w http.ResponseWriter, err error) {
	code := http.StatusExpectationFailed
	SpanError(ctx, err)
	s := ot.SpanFromContext(ctx)
	s.SetTag("http.status_code", code)
	http.Error(w, err.Error(), code)
}
