package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://127.0.0.1:5500/" // La ruta a atacar
	totalRequests := 500                 // Cuántas flechas lanzamos
	concurrency := 50                    // Cuántas al mismo tiempo

	fmt.Printf("🚀 Iniciando ataque de estrés a %s...\n", url)
	fmt.Printf("Configuración: %d peticiones totales con %d en paralelo.\n", totalRequests, concurrency)

	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan struct{}, concurrency)

	success := 0
	failed := 0
	var mu sync.Mutex

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		ch <- struct{}{} // Límite de concurrencia
		go func() {
			defer wg.Done()
			defer func() { <-ch }()

			resp, err := http.Get(url)
			mu.Lock()
			if err != nil || resp.StatusCode != http.StatusOK {
				failed++
			} else {
				success++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Println("\n--- 📊 RESULTADOS DEL TORMENTO ---")
	fmt.Printf("Tiempo total: %v\n", duration)
	fmt.Printf("Peticiones por segundo (RPS): %.2f\n", float64(totalRequests)/duration.Seconds())
	fmt.Printf("✅ Exitosas: %d\n", success)
	fmt.Printf("❌ Fallidas: %d\n", failed)
}