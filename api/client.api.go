package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"meteo/models"
)

var client = &http.Client{
	Timeout: 15 * time.Second,
}

func get(endpoint string, dest any) error {
	return getWithContext(context.Background(), endpoint, dest)
}

func getWithContext(ctx context.Context, endpoint string, dest any) error {
	return getURLWithContext(ctx, baseURL+endpoint, dest)
}

func getURL(requestURL string, dest any) error {
	return getURLWithContext(context.Background(), requestURL, dest)
}

func getURLWithContext(ctx context.Context, requestURL string, dest any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return fmt.Errorf("création de la requête : %w", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)

	response, err := client.Do(request)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return err
		}
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return models.ErrTimeout
		}
		return fmt.Errorf("%w : %v", models.ErrUnavailable, err)
	}
	defer response.Body.Close()

	switch {
	case response.StatusCode == http.StatusOK:
	case response.StatusCode == http.StatusNotFound:
		return models.ErrNotFound
	case response.StatusCode == http.StatusBadRequest:
		return models.ErrBadRequest
	case response.StatusCode >= 500:
		return fmt.Errorf("%w : %s", models.ErrUnavailable, response.Status)
	default:
		return fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return fmt.Errorf("décodage de la réponse : %w", err)
	}
	return nil
}

