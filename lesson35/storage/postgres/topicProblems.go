package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type TopicProblemRepo struct {
	Db *sql.DB
}

func NewTopicProblemRepo(db *sql.DB) *TopicProblemRepo {
	return &TopicProblemRepo{db}
}

// Create
func (l *TopicProblemRepo) CreateTopicProblem(tp model.TopicProblem) error {

	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := "insert into topics_problems(topic_id, problem_id) values($1, $2)"
	_, err = tx.Exec(query, tp.TopicId, tp.ProblemId)

	return err
}

// Read
func (l *TopicProblemRepo) GetTopicProblemById(id string) model.TopicProblem {
	topicProblem := model.TopicProblem{}
	query := `
	select * from topics_problems
	where
		id = $1 and deleted_at is null
	`
	row := l.Db.QueryRow(query, id)
	row.Scan(&topicProblem.Id, &topicProblem.TopicId, &topicProblem.ProblemId, &topicProblem.Created_at, &topicProblem.Updated_at, &topicProblem.Deleted_at)
	return topicProblem
}
func (l *TopicProblemRepo) GetTopicProblems(filter model.TopicProblemFilter) (*[]model.TopicProblem, error) {
	params := []interface{}{}
	paramcount := 0
	query := `
	select * from topics_problems where deleted_at is null`
	if filter.TopicsId != nil {
		query += fmt.Sprintf(" and topic_id=$%d", paramcount)
		params = append(params, *filter.ProblemId)
		paramcount++
	}
	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and problem_id=$%d", paramcount)
		params = append(params, *filter.ProblemId)
		paramcount++
	}

	rows, err := l.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	topicProblems := []model.TopicProblem{}
	for rows.Next() {
		topicProblem := model.TopicProblem{}
		err = rows.Scan(&topicProblem.Id, &topicProblem.TopicId, &topicProblem.ProblemId, &topicProblem.Created_at,
			&topicProblem.Updated_at, &topicProblem.Deleted_at)
		if err != nil {
			return nil, err
		}
		topicProblems = append(topicProblems, topicProblem)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &topicProblems, nil
}

// Update
func (t *TopicProblemRepo) UpdateTopicProblem(tp model.TopicProblem) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics_problems 
	set 
		topic_id = $1,
		problem_id = $2
	where 
		deleted_at is null and id = $3`
	_, err = tx.Exec(query, tp.TopicId, tp.ProblemId, time.Now(), tp.Id)

	return err
}

// Delete
func (t *TopicProblemRepo) DeleteTopicProblem(id string) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics_problems 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
