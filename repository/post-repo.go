package repository

import (
	"log"
	"cloud.google.com/go/firestore"

	"../entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {}

//NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId string = "amit-thapa"
	collectionName string = "posts"
)

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": post.ID,
		"Title": post.Title,
		"Text": post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID: doc.Data()["ID"](int),
			Title: doc.Data()["title"](string),
			Text: doc.Data()["text"](string),
		}

		posts = append(posts, post)
	}
	return posts, nil
}