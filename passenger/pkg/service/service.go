package service

// PassengerService describes the service.
type PassengerService interface {
	HealthCheck() bool
	PassengerInfo(passengerID int64) (map[string]interface{}, error)
}
