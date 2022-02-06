package field

import (
	"golang-echo-graphql/pkg/model"

	"github.com/graphql-go/graphql"

	"github.com/jinzhu/gorm"
)

var user = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"name":      &graphql.Field{Type: graphql.String},
			"age":       &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.String},
			"updatedAt": &graphql.Field{Type: graphql.String},
			"deletedAt": &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)

func NewUsers(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(user),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id, ok := p.Args["id"].(int)
			var u []*model.User
			if ok {
				if err := db.Where("id = ?", id).Find(&u).Error; err != nil {
					// do something
				}
			} else {
				if err := db.Find(&u).Error; err != nil {
					// do something
				}
			}
			return u, nil
		},
		Description: "user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
	}
}
