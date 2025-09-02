package service

// HealthService defines health check logic
type HealthService struct{}

// NewHealthService returns a new instance
func NewHealthService() *HealthService {
	return &HealthService{}
}

// Check returns the current health status of the app
func (hs *HealthService) Check() map[string]string {
	return map[string]string{"status": "ok"}
}
