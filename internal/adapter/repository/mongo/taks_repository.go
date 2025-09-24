package mongorepo

import (
	"context"
	"test-hex-architecture/internal/core/domain/task"
	"test-hex-architecture/internal/core/port"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func (r *TaskRepository) FindAll(ctx context.Context) ([]*task.Task, error) {
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []*task.Task
	for cur.Next(ctx) {
		var d taskDOC
		if err := cur.Decode(&d); err != nil {
			return nil, err
		}
		out = append(out, fromDoc(&d))
	}
	return out, cur.Err()
}

func (r *TaskRepository) Update(ctx context.Context, t *task.Task) error {
	_, err := r.col.UpdateOne(ctx, t.ID, bson.M{"$set": bson.M{
		"title":       t.Title,
		"description": t.Description,
		"done":        t.Done,
		"updated_at":  t.UpdatedAt.Unix(),
	}})
	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
