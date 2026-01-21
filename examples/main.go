package main

import (
	"fmt"
	"log"
	"time"

	"github.com/janexpl/ticktick-go"
)

func main() {
	// Przykład użycia biblioteki TickTick Go Client

	// 1. Konfiguracja OAuth (jednorazowa)
	exampleOAuth()

	// 2. Użycie klienta z tokenem dostępu
	exampleClient()
}

func exampleOAuth() {
	fmt.Println("=== Przykład OAuth2 ===")

	// Konfiguracja OAuth
	config := &ticktick.OAuthConfig{
		ClientID:     "9QU3dxoDFhATR9mo11",
		ClientSecret: "GmesN4P8haha08mP99rzvCdA43cQKcKn",
		RedirectURI:  "http://localhost:8080/callback",
		Scope:        "tasks:read tasks:write",
	}

	// Generowanie URL autoryzacji
	authURL := config.GetAuthorizationURL("random-state-string")
	fmt.Printf("1. Odwiedź URL autoryzacji:\n%s\n\n", authURL)

	// Po autoryzacji użytkownik otrzyma kod
	// W prawdziwej aplikacji kod będzie pobrany z callback URL
	code := "authorization-code-from-callback"

	// Wymiana kodu na token
	token, err := config.ExchangeCode(code)
	if err != nil {
		log.Printf("Błąd wymiany kodu: %v\n", err)
		return
	}

	fmt.Printf("2. Token dostępu: %s\n", token.AccessToken)
	fmt.Printf("3. Token wygasa za: %d sekund\n\n", token.ExpiresIn)
}

func exampleClient() {
	fmt.Println("=== Przykład użycia klienta ===")

	// Utwórz klienta z tokenem dostępu
	client := ticktick.NewClient("your-access-token-here")

	// Przykład 1: Lista projektów
	fmt.Println("\n--- Lista projektów ---")
	projects, err := client.Projects.List()
	if err != nil {
		log.Printf("Błąd pobierania projektów: %v\n", err)
	} else {
		for _, project := range projects {
			fmt.Printf("- %s (ID: %s, Kolor: %s)\n", project.Name, project.ID, project.Color)
		}
	}

	// Przykład 2: Tworzenie nowego projektu
	fmt.Println("\n--- Tworzenie nowego projektu ---")
	newProject, err := client.Projects.Create(&ticktick.CreateProjectRequest{
		Name:  "Projekt testowy",
		Color: "#FF5733",
	})
	if err != nil {
		log.Printf("Błąd tworzenia projektu: %v\n", err)
	} else {
		fmt.Printf("Utworzono projekt: %s (ID: %s)\n", newProject.Name, newProject.ID)
	}

	// Przykład 3: Tworzenie zadania
	fmt.Println("\n--- Tworzenie zadania ---")
	dueDate := time.Now().Add(24 * time.Hour)
	newTask, err := client.Tasks.Create(&ticktick.CreateTaskRequest{
		Title:     "Przykładowe zadanie",
		ProjectID: newProject.ID, // Użyj ID utworzonego projektu
		Content:   "To jest szczegółowy opis zadania",
		Priority:  3,
		DueDate:   &dueDate,
		Tags:      []string{"przykład", "test"},
	})
	if err != nil {
		log.Printf("Błąd tworzenia zadania: %v\n", err)
	} else {
		fmt.Printf("Utworzono zadanie: %s (ID: %s)\n", newTask.Title, newTask.ID)
	}

	// Przykład 4: Lista zadań z projektu
	fmt.Println("\n--- Lista zadań z projektu ---")
	tasks, err := client.Tasks.List(newProject.ID)
	if err != nil {
		log.Printf("Błąd pobierania zadań: %v\n", err)
	} else {
		for _, task := range tasks {
			fmt.Printf("- %s (Priorytet: %d, Status: %d)\n", task.Title, task.Priority, task.Status)
		}
	}

	// Przykład 5: Aktualizacja zadania
	fmt.Println("\n--- Aktualizacja zadania ---")
	updatedTask, err := client.Tasks.Update(newProject.ID, &ticktick.UpdateTaskRequest{
		ID:       newTask.ID,
		Title:    "Zaktualizowane zadanie",
		Priority: 5,
	})
	if err != nil {
		log.Printf("Błąd aktualizacji zadania: %v\n", err)
	} else {
		fmt.Printf("Zaktualizowano zadanie: %s (Priorytet: %d)\n", updatedTask.Title, updatedTask.Priority)
	}

	// Przykład 6: Oznaczanie zadania jako ukończone
	fmt.Println("\n--- Oznaczanie zadania jako ukończone ---")
	completedTask, err := client.Tasks.Complete(newProject.ID, newTask.ID)
	if err != nil {
		log.Printf("Błąd oznaczania zadania: %v\n", err)
	} else {
		fmt.Printf("Ukończono zadanie: %s (Status: %d)\n", completedTask.Title, completedTask.Status)
	}

	// Przykład 7: Obsługa błędów
	fmt.Println("\n--- Obsługa błędów ---")
	_, err = client.Tasks.Get(newProject.ID, "nieistniejace-id")
	if err != nil {
		if apiErr, ok := err.(*ticktick.APIError); ok {
			switch {
			case apiErr.IsNotFound():
				fmt.Println("✓ Prawidłowo wykryto błąd: Zadanie nie znalezione")
			case apiErr.IsUnauthorized():
				fmt.Println("Błąd autoryzacji")
			case apiErr.IsRateLimited():
				fmt.Println("Przekroczono limit zapytań")
			default:
				fmt.Printf("Błąd API: %v\n", apiErr)
			}
		}
	}

	// Przykład 8: Czyszczenie - usuwanie zasobów
	fmt.Println("\n--- Czyszczenie zasobów ---")
	err = client.Tasks.Delete(newProject.ID, newTask.ID)
	if err != nil {
		log.Printf("Błąd usuwania zadania: %v\n", err)
	} else {
		fmt.Println("✓ Usunięto zadanie")
	}

	err = client.Projects.Delete(newProject.ID)
	if err != nil {
		log.Printf("Błąd usuwania projektu: %v\n", err)
	} else {
		fmt.Println("✓ Usunięto projekt")
	}

	fmt.Println("\n=== Koniec przykładów ===")
}
