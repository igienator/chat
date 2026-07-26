package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/igienator/chat/e2e/cmd/p2p"
	"github.com/igienator/chat/e2e/cmd/user"
)

func main() {
	tox, err := user.LoadOrCreateProfile()
	if err != nil {
		log.Fatalf("failed to initialize tox profile: %v", err)
	}

	defer func() {
		if err := user.SaveProfile(tox); err != nil {
			log.Printf("warning: failed to save tox profile: %v", err)
		}
		tox.Kill()
	}()

	srv := p2p.NewServer()

	serverErr := make(chan error, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
		close(serverErr)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	log.Printf("server started on %s", srv.Addr)

	select {
	case sig := <-sigCh:
		log.Printf("received signal %s, shutting down", sig)
	case err := <-serverErr:
		if err != nil {
			log.Fatalf("server error: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	log.Println("shutdown complete")
}
