package storage

import "time"

// ID represents the unique identifier of the project.
// Name represents the name of the project.
// Description represents the description of the project.
// CreatedAt represents the timestamp when the project was created.
type Project struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

// ID is the unique identifier for the classification task.
// ProjectID is the identifier for the project associated with the
// classification task.
// LLMInput is the input data for the classification task.
// LLMOutput is the output data generated by the classification task.
// CreatedAt is the timestamp when the classification task was created.
// Embedding is the binary representation of the task's embedding.
// LabelID is the identifier for the label associated with the classification task.
type ClassificationTask struct {
	ID        int64
	ProjectID int64
	LLMInput  string
	LLMOutput string
	CreatedAt time.Time
	Embedding []byte
	LabelID   *int64
}

// ID represents the unique identifier of the classification task label.
// ProjectID represents the identifier of the project associated with the
// classification task label.
// Label represents the name or description of the classification task label.
// CreatedAt represents the timestamp when the classification task label was
// created.
type ClassificationTaskLabel struct {
	ID        int64
	ProjectID int64
	Label     string
	CreatedAt time.Time
}
