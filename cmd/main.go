package main

import (
	"log"

	"github.com/freitasmatheusrn/agent-calendar/internal/infra/server"
)

func main() {

	httpServer, err := server.InitializeServer()
	if err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}

	log.Printf("Servidor HTTP rodando em %s", httpServer.Addr)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("erro no servidor: %v", err)
	}
}
