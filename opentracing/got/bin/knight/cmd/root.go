package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/k8sland/code2/tracing/got/internal"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
)

const (
	port = ":4005"
	app  = "Knight"
)

var (
	// Version set via build tag
	Version = ""
	jaeger  string
	castle  string
	rootCmd = &cobra.Command{
		Use:   "north",
		Short: "G.O.T " + app + " Man",
		Long:  "G.O.T " + app + " Man",
		Run:   listen,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&jaeger, "jaeger", "j", "192.168.64.60:6831", "Specify Jaeger address")
	rootCmd.Flags().StringVarP(&castle, "castle", "c", "localhost:4000", "Specify a castle service address")
	rootCmd.Version = Version
}

// Execute a command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listen(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)

	tracer, closer := internal.InitJaeger(app, jaeger)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	r.Handle("/v1/melt", http.HandlerFunc(meltHandler)).Methods("POST")

	log.Printf("%s[%s] listening on port %s...\n", app, Version, port)
	log.Panic(http.ListenAndServe(port, m))
}

func meltHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	s := spanFromReq(r)
	defer s.Finish()

	var (
		ctx  = opentracing.ContextWithSpan(r.Context(), s)
		resp internal.Response
	)
	if err := internal.Call(ctx, "POST", urlFor("v1/melt"), r.Body, &resp); err != nil {
		internal.WriteErrOut(ctx, w, err)
		return
	}
	writeResponse(ctx, w, resp)
}

func spanFromReq(r *http.Request) opentracing.Span {
	s := opentracing.StartSpan(r.URL.String()).
		SetTag("component", app).
		SetTag("http.method", "POST").
		SetTag("http.url", r.URL)
	s.LogKV("message", internal.FuncName(1))
	return s
}

func writeResponse(ctx context.Context, w http.ResponseWriter, r internal.Response) {
	s := opentracing.StartSpan(
		internal.FuncName(0),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	defer s.Finish()

	sx := opentracing.ContextWithSpan(ctx, s)
	raw := bytes.NewBufferString("")
	if err := json.NewEncoder(raw).Encode(r); err != nil {
		internal.WriteErrOut(sx, w, err)
	}
	s.SetTag("http.status_code", http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, raw.String())
}

func urlFor(path string) string {
	return "http://" + castle + "/" + path
}
