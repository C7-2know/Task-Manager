package data

import (
	"context"
	"errors"

	model "task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
)

func Get_task(id string) (model.Task, error) {
	collection := Client.Database("task_manager").Collection("tasks")
	filter := bson.D{{"id", id}}
	result := collection.FindOne(context.TODO(), filter)
	var task model.Task
	err := result.Decode(&task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func Get_tasks() []model.Task {
	if Client == nil {
		return []model.Task{}
	}
	collection := Client.Database("task_manager").Collection("tasks")

	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []model.Task{}
	}
	var tasks []model.Task
	for cursor.Next(context.TODO()) {
		var task model.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func Create_task(task model.Task) error {
	collection := Client.Database("task_manager").Collection("tasks")

	_, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

func Update_task(id string, updated model.Task) error {
	collection := Client.Database("task_manager").Collection("tasks")
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"title", updated.Title},
			{"description", updated.Description},
			{"status", updated.Status}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result == nil || result.ModifiedCount == 0 {
		return errors.New("Server error")
	}
	return nil
}

func Delete_task(id string) error {
	collection := Client.Database("task_manager").Collection("tasks")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
