package pkg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// AllTasks возвращает список всех задач из БД.
func (s *Storage) AllTasks() ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

// TasksByAuthor возвращает список задач по автору.
func (s *Storage) TasksByAuthor(authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			author_id = $1
		ORDER BY id;
	`,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

// TasksByLabel возвращает список задач по метке.
func (s *Storage) TasksByLabel(labelID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
SELECT
    tasks.id,
    tasks.opened,
    tasks.closed,
    tasks.author_id,
    tasks.assigned_id,
    tasks.title,
    tasks.content
FROM tasks JOIN tasks_labels
                ON tasks.id = tasks_labels.task_id
WHERE
    label_id = 2
ORDER BY id;
	`,
		labelID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

// UpdateTask обновляет задачу по id.
func (s *Storage) UpdateTask(taskID int) error {
	_, err := s.db.Query(context.Background(), `
update tasks
set title  = 'первая в очереди',
    closed = extract(epoch from now())

WHERE id = 2;
	`,
		taskID,
	)

	return err
}

// DeleteTask Удаляет задачу по id.
func (s *Storage) DeleteTask(taskID int) error {
	_, err := s.db.Query(context.Background(), `
DELETE FROM tasks_labels
WHERE
        task_id = $1;
DELETE FROM tasks
WHERE
        id = $1;
	`,
		taskID,
	)
	return err
}

// NewTask создаёт новую задачу
func (s *Storage) NewTask(t Task) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
INSERT INTO tasks (title, content, assigned_id, author_id)
VALUES ($1, $2, $3, $4);
		`,
		t.AssignedID,
		t.AuthorID,
		t.Title,
		t.Content,
	).Scan(&id)
	return err
}
