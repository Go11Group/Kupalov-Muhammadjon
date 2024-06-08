package model

type TopicProblem struct {
	Id        string
	TopicId  string
	ProblemId string
	Time
}

type TopicProblemFilter struct{
	TopicsId  *string
	ProblemId *string
}