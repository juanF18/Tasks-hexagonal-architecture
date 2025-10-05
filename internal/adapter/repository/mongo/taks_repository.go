package mongorepo

import (
	"context"
	"test-hex-architecture/internal/core/domain/task"
	"test-hex-architecture/internal/core/port"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TaskRepository struct{ col *mongo.Collection }

func NewTaskRepository(db *mongo.Database) port.TaskRepository {
	return &TaskRepository{col: db.Collection("tasks")}
}

type taskDOC struct {
	ID          string `bson:"_id"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Done        bool   `bson:"done"`
	CreatedAt   int64  `bson:"created_at"`
	UpdatedAt   int64  `bson:"updated_at"`
}

func toDoc(t *task.Task) *taskDOC {
	return &taskDOC{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
}

func fromDoc(d *taskDOC) *task.Task {
	return &task.Task{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Done:        d.Done,
		CreatedAt:   time.Unix(d.CreatedAt, 0).UTC(),
		UpdatedAt:   time.Unix(d.UpdatedAt, 0).UTC(),
	}
}

func (r *TaskRepository) CountAll(ctx context.Context) (int64, error) {
	return r.col.CountDocuments(ctx, bson.M{})
}

func (r *TaskRepository) Save(ctx context.Context, t *task.Task) (string, error) {
	_, err := r.col.InsertOne(ctx, toDoc(t))
	if err != nil {
		return "", err
	}
	return t.ID, nil
}

func (r *TaskRepository) FindByID(ctx context.Context, id string) (*task.Task, error) {
	var d taskDOC
	if err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return fromDoc(&d), nil
}

func (r *TaskRepository) FindAll(ctx context.Context, offset, limit int) ([]*task.Task, error) {
	opts := options.Find()

	if offset >= 0 && limit > 0 {
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(limit))
	}

	opts.SetSort(bson.M{"created_at": -1})

	cursor, err := r.col.Find(ctx, bson.M{}, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*task.Task
	for cursor.Next(ctx) {
		var d taskDOC
		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}
		tasks = append(tasks, fromDoc(&d))
	}
	return tasks, cursor.Err()
}

func (r *TaskRepository) Update(ctx context.Context, t *task.Task) error {
	update := bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"done":        t.Done,
			"updatedAt":   t.UpdatedAt.Unix(),
		},
	}
	_, err := r.col.UpdateByID(ctx, t.ID, update)
	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
