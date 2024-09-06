package projects

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/anujkkheria/portfolio_API/common"
	"github.com/anujkkheria/portfolio_API/common/utils"
	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
)

type ProjectBody struct{
	Link string `json:"link"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Category string `json:"category"`
}

type NewProject struct{
	Link string `json:"link" bson:"link"`
	Title string `json:"title" bson:"title"`
	Image string `json:"image" bson:"image"`
	Desc string `json:"desc" bson:"desc"`
	Category string `json:"category" bson:"category"`
}

type Project struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Link string `json:"link" bson:"link"`
	Title string `json:"title" bson:"title"`
	Image string `json:"image" bson:"image"`
	Desc string `json:"desc" bson:"desc"`
	Category string `json:"category" bson:"category"`
}



func CreateNewProject(c *fiber.Ctx)error{

	 project := new(ProjectBody)
		if err := c.BodyParser(project); err != nil{
			return c.Status(400).JSON(fiber.Map{
				"message":"invalid Body please check the body structure",
			})
		}
imageMetadata,err := getScreenShotfromLink(project.Link, project.Title)
	if err !=nil{
		return c.Status(500).JSON(fiber.Map{
			"message":err.Error(),
		})
	}

	projectMetaData:= NewProject{
		Link:project.Link,
		Title: project.Title,
		Category: project.Category,
		Desc:project.Desc,
		Image: imageMetadata.Location,
		}

		coll := common.GetDBCollection("projects")
		
		result, err := coll.InsertOne(c.Context(),projectMetaData)

		if err!=nil{
			return c.Status(500).JSON(fiber.Map{"message":"DB Storage Failed"})
		}
fmt.Println(result)
		return c.Status(201).JSON(fiber.Map{
			"message" :"Tesing the API",
			"result":projectMetaData,
		})
}

func GetProjects(c *fiber.Ctx) error {
	coll := common.GetDBCollection("projects")
	projects := make([]Project,0)
	cursor,err := coll.Find(c.Context(),bson.M{})
	if err != nil {
fmt.Println(err.Error())
c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	"message":"failed to get data",
	"error": err.Error(),
})
	}
for cursor.Next(c.Context()){
project := Project{}
err:=cursor.Decode(&project)
if err!=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message":"failed to decode blog",
				"error":err.Error(),
			})
}
projects = append(projects, project)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Successful",
		"result":projects,
	})
}

func DeleteProject(c *fiber.Ctx)error{
	id := c.Params("id")
	if id == ""{
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	"message":"no id found",
	})
	}
	objectId,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid Id",
		})
	} 
	coll := common.GetDBCollection("projects")
	result,err := coll.DeleteOne(c.Context(), bson.M{"_id":objectId})
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"The error occured due to error",
			"error":err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"successfully deleted",
		"result":result,
	}) 
}

	func getScreenShotfromLink(url string, title string) (*manager.UploadOutput,error){
		ctx,cancel:=	chromedp.NewContext(context.Background())
		defer cancel()
		
		fmt.Print(ctx,url)
var buf []byte
 
err:= chromedp.Run(ctx, chromedp.Tasks{
	chromedp.Navigate(url),
	chromedp.FullScreenshot(&buf,90),
})

if err !=nil{
	fmt.Println(err.Error())
	return nil,err
}

// err = SaveImage(title, buf)

result,err := utils.UploadS3(title, buf)

if err != nil{
	fmt.Println(err.Error())
	return nil,err
}



return result, nil

}


func SaveImage(title string, buf []byte)error{
	filename := fmt.Sprintf("%s.%s",title,".png")
	err := os.WriteFile(filename, buf, 0644)
if err != nil{ 
	return err
}
	return nil
} 