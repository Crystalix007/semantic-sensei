package storage

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
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

// GetProject returns the project with the provided ID.
func (d Database) GetProject(ctx context.Context, id int64) (*Project, error) {
	var p Project

	err := d.db.QueryRowContext(ctx, `
		SELECT id, name, description, created_at
		FROM projects
		WHERE id = $1
	`, id).Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error getting project: %w",
			err,
		)
	}

	return &p, nil
}

// FindProjects gets a list of all projects.
func (d Database) FindProjects(ctx context.Context) ([]Project, error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT id, name, description, created_at
		FROM projects
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error getting all projects: %w",
			err,
		)
	}

	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project

		if err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf(
				"storage: error scanning project row: %w",
				err,
			)
		}

		projects = append(projects, project)
	}

	return projects, nil
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

/// PendingClassificationTask methods.

// CreatePendingClassificationTask creates a new pending classification task in
// the database and returns the ID of the newly created pending classification
// task.
func (d Database) CreatePendingClassificationTask(ctx context.Context, pct PendingClassificationTask) (int64, error) {
	var id int64

	err := d.db.QueryRowContext(ctx, `
		INSERT INTO pending_classification_tasks (
			project_id,
			llm_input,
			llm_output,
			embedding
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, pct.ProjectID, pct.LLMInput, pct.LLMOutput, pct.Embedding).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf(
			"storage: error creating pending classification task: %w",
			err,
		)
	}

	return id, nil
}

// GetPendingClassificationTask returns the pending classification task with the
// provided ID.
func (d Database) GetPendingClassificationTask(ctx context.Context, id int64) (*PendingClassificationTask, error) {
	var pct PendingClassificationTask

	err := d.db.QueryRowContext(ctx, `
		SELECT id, project_id, llm_input, llm_output, created_at, embedding
		FROM pending_classification_tasks
		WHERE id = $1
	`, id).Scan(
		&pct.ID,
		&pct.ProjectID,
		&pct.LLMInput,
		&pct.LLMOutput,
		&pct.CreatedAt,
		&pct.Embedding,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error getting pending classification task: %w",
			err,
		)
	}

	return &pct, nil
}

// FindPendingClassificationTasksForProject returns a list of pending
// classification tasks for the specified project ID.
func (d Database) FindPendingClassificationTasksForProject(
	ctx context.Context,
	projectID int64,
	parameters Parameters,
) ([]PendingClassificationTask, error) {
	selectBuilder := sqlbuilder.NewSelectBuilder()

	selectBuilder.Select(
		"id",
		"project_id",
		"llm_input",
		"llm_output",
		"embedding",
		"created_at",
	).From(
		"pending_classification_tasks",
	).Where(
		selectBuilder.Equal("project_id", projectID),
	).Where(
		parameters.Where...,
	)

	if parameters.PageSize != 0 {
		selectBuilder.Offset(int(parameters.Page) * int(parameters.PageSize))
		selectBuilder.Limit(int(parameters.PageSize))
	}

	sql, binds := selectBuilder.Build()

	rows, err := d.db.QueryContext(ctx, sql, binds...)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error finding pending classification tasks for project: %w",
			err,
		)
	}

	defer rows.Close()

	var pendingClassificationTasks []PendingClassificationTask

	for rows.Next() {
		var pendingClassificationTask PendingClassificationTask

		if err := rows.Scan(
			&pendingClassificationTask.ID,
			&pendingClassificationTask.ProjectID,
			&pendingClassificationTask.LLMInput,
			&pendingClassificationTask.LLMOutput,
			&pendingClassificationTask.Embedding,
			&pendingClassificationTask.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf(
				"storage: error scanning pending classification task: %w",
				err,
			)
		}

		pendingClassificationTasks = append(
			pendingClassificationTasks,
			pendingClassificationTask,
		)
	}

	return pendingClassificationTasks, nil
}

// FindPendingClassificationTaskCountForProject returns the count of pending
// classification tasks for the specified project ID.
func (d Database) FindPendingClassificationTaskCountForProject(
	ctx context.Context,
	projectID int64,
	where ...string,
) (uint64, error) {
	selectBuilder := sqlbuilder.NewSelectBuilder()

	selectBuilder.Select(
		"count(*)",
	).From(
		"pending_classification_tasks",
	).Where(
		selectBuilder.Equal("project_id", projectID),
	).Where(
		where...,
	)

	var count uint64

	sql, binds := selectBuilder.Build()

	err := d.db.QueryRowContext(ctx, sql, binds...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf(
			"storage: error getting pending classification task count for project: %w",
			err,
		)
	}

	return count, nil
}

// UpdatePendingClassificationTask updates the pending classification task with
// the provided ID with the provided values.
func (d Database) UpdatePendingClassificationTask(ctx context.Context, pct PendingClassificationTask) error {
	_, err := d.db.ExecContext(ctx, `
		UPDATE pending_classification_tasks
		SET project_id = $1, llm_input = $2, llm_output = $3, embedding = $4
		WHERE id = $5
	`, pct.ProjectID, pct.LLMInput, pct.LLMOutput, pct.Embedding, pct.ID)
	if err != nil {
		return fmt.Errorf(
			"storage: error updating pending classification task: %w",
			err,
		)
	}

	return nil
}

// DeletePendingClassificationTask deletes a pending classification task from
// the database based on the given ID.
func (d Database) DeletePendingClassificationTask(ctx context.Context, id int64) error {
	_, err := d.db.ExecContext(ctx, `
		DELETE FROM pending_classification_tasks
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf(
			"storage: error deleting pending classification task: %w",
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

// GetClassificationTask returns the classification task with the provided ID.
func (d Database) GetClassificationTask(ctx context.Context, id int64) (*ClassificationTask, error) {
	var ct ClassificationTask

	err := d.db.QueryRowContext(ctx, `
		SELECT id, project_id, llm_input, llm_output, created_at, embedding, label_id
		FROM classification_tasks
		WHERE id = $1
	`, id).Scan(
		&ct.ID,
		&ct.ProjectID,
		&ct.LLMInput,
		&ct.LLMOutput,
		&ct.CreatedAt,
		&ct.Embedding,
		&ct.LabelID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error getting classification task: %w",
			err,
		)
	}

	return &ct, nil
}

// FindClassificationTasksForProject returns a list of classification tasks for
// the specified project ID.
func (d Database) FindClassificationTasksForProject(
	ctx context.Context,
	projectID int64,
	parameters Parameters,
) ([]ClassificationTask, error) {
	selectBuilder := sqlbuilder.NewSelectBuilder()

	selectBuilder.Select(
		"id",
		"project_id",
		"llm_input",
		"llm_output",
		"embedding",
		"label_id",
		"created_at",
	).From(
		"classification_tasks",
	).Where(
		selectBuilder.Equal("project_id", projectID),
	).Where(
		parameters.Where...,
	)

	if parameters.PageSize != 0 {
		selectBuilder.Offset(int(parameters.Page) * int(parameters.PageSize))
		selectBuilder.Limit(int(parameters.PageSize))
	}

	sql, binds := selectBuilder.Build()

	rows, err := d.db.QueryContext(ctx, sql, binds...)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error finding classification tasks for project: %w",
			err,
		)
	}

	defer rows.Close()

	var classificationTasks []ClassificationTask

	for rows.Next() {
		var classificationTask ClassificationTask

		if err := rows.Scan(
			&classificationTask.ID,
			&classificationTask.ProjectID,
			&classificationTask.LLMInput,
			&classificationTask.LLMOutput,
			&classificationTask.Embedding,
			&classificationTask.LabelID,
			&classificationTask.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf(
				"storage: error scanning classification task: %w",
				err,
			)
		}

		classificationTasks = append(classificationTasks, classificationTask)
	}

	return classificationTasks, nil
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

// GetClassificationTaskLabel returns the classification task label with the
// provided ID.
func (d Database) GetClassificationTaskLabel(ctx context.Context, id int64) (*ClassificationTaskLabel, error) {
	var ctl ClassificationTaskLabel

	err := d.db.QueryRowContext(ctx, `
		SELECT id, project_id, label, created_at
		FROM classification_task_labels
		WHERE id = $1
	`, id).Scan(&ctl.ID, &ctl.ProjectID, &ctl.Label, &ctl.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error getting classification task label: %w",
			err,
		)
	}

	return &ctl, nil
}

// FindClassificationTaskLabelsForProject returns the classification task label
// with the given project ID.
func (d Database) FindClassificationTaskLabelsForProject(
	ctx context.Context,
	projectID int64,
) ([]ClassificationTaskLabel, error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT id, project_id, label, created_at
		FROM classification_task_labels
		WHERE project_id = $1
		ORDER BY id ASC
	`, projectID)
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error getting classification task labels for project: %w",
			err,
		)
	}

	defer rows.Close()

	var ctls []ClassificationTaskLabel

	for rows.Next() {
		var ctl ClassificationTaskLabel

		if err := rows.Scan(
			&ctl.ID,
			&ctl.ProjectID,
			&ctl.Label,
			&ctl.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf(
				"storage: error scanning project classification task label row: %w",
				err,
			)
		}

		ctls = append(ctls, ctl)
	}

	return ctls, nil
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
