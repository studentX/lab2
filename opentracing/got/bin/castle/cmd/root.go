package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/k8sland/code2/tracing/got/internal"
	opentracing "github.com/opentracing/opentracing-go"
	ext "github.com/opentracing/opentracing-go/ext"
	"github.com/spf13/cobra"
)

const (
	app  = "Castle"
	port = ":4000"
)

var (
	// Version set via build tags
	Version = ""
	jaeger  string
	rootCmd = &cobra.Command{
		Use:   strings.ToLower(app),
		Short: "G.O.T " + app,
		Long:  "G.O.T " + app,
		Run:   web,
	}
)

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&jaeger, "jaeger", "j", "192.168.64.60:6831", "Specify Jaeger address")
	rootCmd.Version = Version
}

func web(cmd *cobra.Command, args []string) {
	tracer, closer := internal.InitJaeger(app, jaeger)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)

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

	s, err := startSpan(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	defer s.Finish()
	sx := opentracing.ContextWithSpan(r.Context(), s)

	var q internal.Quest
	if q, err = readQuest(sx, w, r); err != nil {
		return
	}
	s.SetTag("knight", q.Knight)

	if meltAuth(q.Knight) {
		if err := writeResponse(sx, w); err == nil {
			s.LogKV("message", fmt.Sprintf("castle successfully melted"))
		}
	} else {
		internal.WriteErrOut(sx, w, fmt.Errorf("only the nightking can melt"))
		return
	}
}

func readQuest(ctx context.Context, w http.ResponseWriter, r *http.Request) (internal.Quest, error) {
	s := opentracing.StartSpan(
		internal.FuncName(0),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	defer s.Finish()

	var q internal.Quest
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		internal.WriteErrOut(ctx, w, err)
		return q, err
	}
	s.SetTag("action", "castle.quest")
	s.LogKV("message", fmt.Sprintf("%s requested a melt", q.Knight))
	return q, nil
}

func writeResponse(ctx context.Context, w http.ResponseWriter) error {
	s := opentracing.StartSpan(
		internal.FuncName(0),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	defer s.Finish()

	resp := internal.Response{Status: "💀  Castle Melted!!"}
	if raw, err := json.Marshal(resp); err != nil {
		sx := opentracing.ContextWithSpan(ctx, s)
		internal.WriteErrOut(sx, w, err)
		return err
	}
	s.SetTag("action", "castle.melted")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(raw))
	return nil
}

func meltAuth(k string) bool {
	return strings.ToLower(k) == "nightking"
}

func startSpan(r *http.Request) (opentracing.Span, error) {
	var s opentracing.Span
	ctx, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		return s, err
	}

	s, err = opentracing.StartSpan(r.URL.String(), ext.RPCServerOption(ctx)), nil
	if err != nil {
		return s, err
	}
	// Tag the span the following tags component, http.method, http.url
	// YOUR_CODE!!
	return s, nil
}
