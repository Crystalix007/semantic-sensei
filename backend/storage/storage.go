package storage

import (
	"context"
	"fmt"
)

/// Project methods.

// CreateProject creates a new project in the database and returns the ID of the
// newly created project.
func (d Database) CreateProject(ctx context.Context, p Project) (int64, error) {
	var id int64

	err := d.db.QueryRowContext(ctx, `
		INSERT INTO projects (name, description)
		VALUES ($1, $2)
		RETURNING id
	`, p.Name, p.Description).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf(
			"storage: error creating project: %w",
			err,
		)
	}

	return id, nil
}

// UpdateProject updates the project with the provided ID with the provided
// values.
func (d Database) UpdateProject(ctx context.Context, p Project) error {
	_, err := d.db.ExecContext(ctx, `
		UPDATE projects
		SET name = $1, description = $2
		WHERE id = $3
	`, p.Name, p.Description, p.ID)
	if err != nil {
		return fmt.Errorf(
			"storage: error updating project: %w",
			err,
		)
	}

	return nil
}

// DeleteProject deletes a project from the database based on the given ID.
// It takes a context.Context and an int64 ID as parameters.
// It returns an error if there was a problem deleting the project.
func (d Database) DeleteProject(ctx context.Context, id int64) error {
	_, err := d.db.ExecContext(ctx, `
		DELETE FROM projects
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf(
			"storage: error deleting project: %w",
			err,
		)
	}

	return nil
}

/// ClassificationTask methods.

// CreateClassificationTask creates a new classification task in the database
// and returns the ID of the newly created classification task.
func (d Database) CreateClassificationTask(ctx context.Context, ct ClassificationTask) (int64, error) {
	var id int64

	err := d.db.QueryRowContext(ctx, `
		INSERT INTO classification_tasks (
			project_id,
			llm_input,
			llm_output,
			embedding,
			label_id
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, ct.ProjectID, ct.LLMInput, ct.LLMOutput, ct.Embedding, ct.LabelID).Scan(
		&id,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"storage: error creating classification task: %w",
			err,
		)
	}

	return id, nil
}

// UpdateClassificationTask updates the classification task with the provided ID
// with the provided values.
func (d Database) UpdateClassificationTask(ctx context.Context, ct ClassificationTask) error {
	_, err := d.db.ExecContext(ctx, `
		UPDATE classification_tasks
		SET project_id = $1, llm_input = $2, llm_output = $3, embedding = $4, label_id = $5
		WHERE id = $6
	`, ct.ProjectID, ct.LLMInput, ct.LLMOutput, ct.Embedding, ct.LabelID, ct.ID)
	if err != nil {
		return fmt.Errorf(
			"storage: error updating classification task: %w",
			err,
		)
	}

	return nil
}

// DeleteClassificationTask deletes a classification task from the database
// based on the given ID.
func (d Database) DeleteClassificationTask(ctx context.Context, id int64) error {
	_, err := d.db.ExecContext(ctx, `
		DELETE FROM classification_tasks
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf(
			"storage: error deleting classification task: %w",
			err,
		)
	}

	return nil
}

/// ClassificationTaskLabels methods.

// CreateClassificationTaskLabel creates a new classification task label in the
// database and returns the ID of the newly created classification task label.
func (d Database) CreateClassificationTaskLabel(ctx context.Context, ctl ClassificationTaskLabel) (int64, error) {
	var id int64

	err := d.db.QueryRowContext(ctx, `
		INSERT INTO classification_task_labels (project_id, label)
		VALUES ($1, $2)
		RETURNING id
	`, ctl.ProjectID, ctl.Label).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf(
			"storage: error creating classification task label: %w",
			err,
		)
	}

	return id, nil
}

// UpdateClassificationTaskLabel updates the classification task label with the
// provided ID with the provided values.
func (d Database) UpdateClassificationTaskLabel(ctx context.Context, ctl ClassificationTaskLabel) error {
	_, err := d.db.ExecContext(ctx, `
		UPDATE classification_task_labels
		SET project_id = $1, label = $2
		WHERE id = $3
	`, ctl.ProjectID, ctl.Label, ctl.ID)
	if err != nil {
		return fmt.Errorf(
			"storage: error updating classification task label: %w",
			err,
		)
	}

	return nil
}

// DeleteClassificationTaskLabel deletes a classification task label from the
// database based on the given ID.
func (d Database) DeleteClassificationTaskLabel(ctx context.Context, id int64) error {
	_, err := d.db.ExecContext(ctx, `
		DELETE FROM classification_task_labels
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf(
			"storage: error deleting classification task label: %w",
			err,
		)
	}

	return nil
}
