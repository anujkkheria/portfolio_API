package blogs

import (
	"github.com/anujkkheria/portfolio_API/common"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

	type Newblog struct{
		Title string `json:"title" bson:"title"`
		Type string `json:"type" bson:"type"`
		Link string `json:"link" bson:"link"`
	}

	type Blog struct{
		ID primitive.ObjectID `json:"id" bson:"_id"`
		Title string `json:"title" bson:"title"`
		Type string `json:"type" bson:"type"`
		Link string `json:"link" bson:"link"`
	}

	func CreateBlogs(c *fiber.Ctx) error{

	body := new(Newblog)
	if err := c.BodyParser(body); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"error":"invalid Body",
		})
	}

	coll := common.GetDBCollection("blogs")

	result,err := coll.InsertOne(c.Context(), body)

	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message" :"failed to create book",
			"error": err.Error(),
		})
	}

return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	"result":result,
})
}

func GetBlogs(c *fiber.Ctx) error{
	coll := common.GetDBCollection("blogs")
	blogs := make([]Blog, 0)
	cursor,err := coll.Find(c.Context(),bson.M{})
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"failed to get blogs",
			"error":err.Error(),
		})
	}

	for	cursor.Next(c.Context()){
		blog:= Blog{}
		err := cursor.Decode(&blog)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message":"failed to decode blog",
				"error":err.Error(),
			})
		}
		blogs = append(blogs, blog)
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blogs":blogs,
	})
}

func UpdateBlog(c *fiber.Ctx) error{
	id := c.Params("id")
	
	if id =="" {
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	"message":"id is required",
})
	}
	
	objectId,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"invalid id",
			"error":err.Error(),
		})
	}

	body := new(Blog)
	if err := c.BodyParser(body); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"invalid body",
			"error":err.Error(),
		})
}

coll := common.GetDBCollection("blogs")

result,err := coll.UpdateOne(c.Context(), bson.M{"_id":objectId},bson.M{"$set":body}) 

if err !=nil{
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message":"failed to update blog",
		"error":err.Error(),
	})
}

return c.Status(fiber.StatusOK).JSON(fiber.Map{
	"message":"blog updated successfully",
	"result":result,
})
}
