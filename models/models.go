package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
type UserCreate struct {
	Username string `json:"username" binding:"required,min=3,max=15,alphanum"`
	FullName string `json:"full_name" binding:"required,min=3,max=30,alphanum"`
}

type taskStatus int

const (
	ToDo taskStatus = iota
	InProgress
	Done
)

type TaskCreate struct {
	Title       string     `json:"title" binding:"required,min=3,max=30,alphanum"`
	Description string     `json:"description" binding:"required,min=3,max=255,alphanum"`
	Status      taskStatus `json:"status"`
	AssigneeId  int        `json:"assignee_id"`
}
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      taskStatus `json:"status"`
	AssigneeId  int        `json:"assignee_id"`
}

var Users = []User{
	{ID: 1, Username: "some_guy", FullName: "Guy Bernfeld"},
	{ID: 6, Username: "other_guy", FullName: "Yogev Gabay"},
	{ID: 2, Username: "jane_doe", FullName: "Jane Doe"},
	{ID: 3, Username: "john_smith", FullName: "John Smith"},
	{ID: 4, Username: "alice_wonder", FullName: "Alice Wonderland"},
	{ID: 5, Username: "bob_builder", FullName: "Bob Builder"},
}

var Tasks = []Task{
	{
		ID:          1,
		Title:       "Setup CI/CD Pipeline",
		Description: "Configure automated build and deployment pipeline using GitHub Actions.",
		Status:      ToDo,
		AssigneeId:  2,
	},
	{
		ID:          2,
		Title:       "Provision Kubernetes Cluster",
		Description: "Create and configure a production-ready Kubernetes cluster on cloud.",
		Status:      InProgress,
		AssigneeId:  3,
	},
	{
		ID:          3,
		Title:       "Implement Monitoring",
		Description: "Integrate Prometheus and Grafana for system monitoring and alerting.",
		Status:      ToDo,
		AssigneeId:  4,
	},
	{
		ID:          4,
		Title:       "Containerize Application",
		Description: "Dockerize all microservices and update deployment manifests.",
		Status:      Done,
		AssigneeId:  1,
	},
	{
		ID:          5,
		Title:       "Configure Secrets Management",
		Description: "Set up Vault for secure storage and retrieval of secrets.",
		Status:      InProgress,
		AssigneeId:  2,
	},
}
