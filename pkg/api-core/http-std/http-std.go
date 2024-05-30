package httpstd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
)

type httpStd struct{}

var O httpStd = httpStd{}
var wg sync.WaitGroup

func (o *httpStd) Init(c *t.HttpConf, h *http.Handler) {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", c.Host, c.Port),
		Handler:      requestLog(*h),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Gracefull shutdown
	shutdownError := make(chan error)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		s := <-sig

		log.Println("Stopping signal " + s.String() + " is received")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		wg.Wait()
		shutdownError <- nil
	}()

	log.Printf("Instantiation for http server using net/http at \"%v\", status: LISTENING...", srv.Addr)
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err.Error())
	}

	err = <-shutdownError
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Closing http server is done")
}
