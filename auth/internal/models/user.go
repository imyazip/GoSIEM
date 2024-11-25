package models

type User struct {
	ID        int64
	Username  string
	Password  string
	RoleID    int64
	CreatedAt string
	UpdatedAt string
}

type Sensor struct {
	ID            int64
	Sensor_id     string
	Name          string
	Hostname      string
	Os_version    string
	Sensor_type   string
	Agent_version string
	Created_at    string
}
