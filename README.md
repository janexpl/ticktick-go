# TickTick Go Client

Biblioteka Go do obsługi API aplikacji TickTick - narzędzia do zarządzania zadaniami i projektami.

## Instalacja

```bash
go get github.com/janexpl/ticktick-go
```

## Wymagania

- Go 1.21 lub nowszy
- Konto TickTick z dostępem do API
- Client ID i Client Secret z TickTick Developer Portal

## Konfiguracja OAuth2

Aby korzystać z tej biblioteki, musisz najpierw zarejestrować aplikację w TickTick Developer Portal:

1. Przejdź do https://developer.ticktick.com/
2. Utwórz nową aplikację
3. Zapisz Client ID i Client Secret
4. Ustaw Redirect URI

## Użycie

### Autentykacja OAuth2

```go
package main

import (
    "fmt"
    "github.com/janexpl/ticktick-go"
)

func main() {
    // Konfiguracja OAuth
    config := &ticktick.OAuthConfig{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        RedirectURI:  "http://localhost:8080/callback",
        Scope:        "tasks:read tasks:write",
    }

    // Generowanie URL autoryzacji
    authURL := config.GetAuthorizationURL("random-state-string")
    fmt.Println("Odwiedź:", authURL)

    // Po autoryzacji użytkownik otrzyma kod autoryzacji
    // Wymień kod na token dostępu
    code := "authorization-code-from-callback"
    token, err := config.ExchangeCode(code)
    if err != nil {
        panic(err)
    }

    fmt.Println("Access Token:", token.AccessToken)
}
```

### Tworzenie klienta

```go
package main

import (
    "github.com/janexpl/ticktick-go"
)

func main() {
    // Utwórz klienta z tokenem dostępu
    client := ticktick.NewClient("your-access-token")

    // Teraz możesz używać klienta do wykonywania operacji
}
```

### Operacje na projektach

```go
// Pobierz wszystkie projekty
projects, err := client.Projects.List()
if err != nil {
    panic(err)
}

// Utwórz nowy projekt
newProject, err := client.Projects.Create(&ticktick.CreateProjectRequest{
    Name:  "Mój nowy projekt",
    Color: "#FF5733",
})
if err != nil {
    panic(err)
}

// Pobierz konkretny projekt
project, err := client.Projects.Get("project-id")
if err != nil {
    panic(err)
}

// Zaktualizuj projekt
updatedProject, err := client.Projects.Update(&ticktick.UpdateProjectRequest{
    ID:    "project-id",
    Name:  "Zaktualizowana nazwa",
    Color: "#33FF57",
})
if err != nil {
    panic(err)
}

// Usuń projekt
err = client.Projects.Delete("project-id")
if err != nil {
    panic(err)
}
```

### Operacje na zadaniach

```go
// Pobierz wszystkie zadania z projektu
tasks, err := client.Tasks.List("project-id")
if err != nil {
    panic(err)
}

// Utwórz nowe zadanie
newTask, err := client.Tasks.Create(&ticktick.CreateTaskRequest{
    Title:     "Moje nowe zadanie",
    ProjectID: "project-id",
    Content:   "Opis zadania",
    Priority:  3,
})
if err != nil {
    panic(err)
}

// Pobierz konkretne zadanie
task, err := client.Tasks.Get("project-id", "task-id")
if err != nil {
    panic(err)
}

// Zaktualizuj zadanie
updatedTask, err := client.Tasks.Update("project-id", &ticktick.UpdateTaskRequest{
    ID:       "task-id",
    Title:    "Zaktualizowany tytuł",
    Priority: 5,
})
if err != nil {
    panic(err)
}

// Oznacz zadanie jako ukończone
completedTask, err := client.Tasks.Complete("project-id", "task-id")
if err != nil {
    panic(err)
}

// Usuń zadanie
err = client.Tasks.Delete("project-id", "task-id")
if err != nil {
    panic(err)
}
```

### Obsługa błędów

```go
task, err := client.Tasks.Get("project-id", "task-id")
if err != nil {
    if apiErr, ok := err.(*ticktick.APIError); ok {
        switch {
        case apiErr.IsNotFound():
            fmt.Println("Zadanie nie zostało znalezione")
        case apiErr.IsUnauthorized():
            fmt.Println("Nieautoryzowany - sprawdź token")
        case apiErr.IsRateLimited():
            fmt.Println("Przekroczono limit zapytań")
        case apiErr.IsServerError():
            fmt.Println("Błąd serwera TickTick")
        default:
            fmt.Printf("Błąd API: %v\n", apiErr)
        }
    } else {
        fmt.Printf("Błąd: %v\n", err)
    }
}
```

## Struktura projektu

```
.
├── client.go      # Główny klient API
├── auth.go        # Obsługa OAuth2
├── tasks.go       # Operacje na zadaniach
├── projects.go    # Operacje na projektach
├── types.go       # Definicje typów danych
├── errors.go      # Własne typy błędów
├── examples/      # Przykłady użycia
│   └── main.go
├── go.mod         # Zależności modułu
└── README.md      # Dokumentacja
```

## Funkcjonalności

- ✅ Autentykacja OAuth2
- ✅ Zarządzanie projektami (tworzenie, odczyt, aktualizacja, usuwanie)
- ✅ Zarządzanie zadaniami (tworzenie, odczyt, aktualizacja, usuwanie)
- ✅ Oznaczanie zadań jako ukończone
- ✅ Obsługa błędów API
- ✅ Wsparcie dla priorytetów zadań
- ✅ Wsparcie dla dat i przypomnień

## Licencja

MIT

## Autor

janexpl
