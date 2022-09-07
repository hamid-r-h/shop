package comment

import (
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func AddComment(c *fiber.Ctx) error {

	var commnets models.Comment
	if err := c.BodyParser(&commnets); err != nil {
		c.Status(fiber.StatusBadRequest).JSON("wrong input fromat")
	}
	_, err := govalidator.ValidateStruct(commnets)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("please input in correct format")
	}
	product_id, err := c.ParamsInt("productid")
	var product models.Product
	if err := db.Database.Db.Find(&product, "id = ?", product_id); err != nil {
		c.Status(fiber.StatusBadRequest).JSON("product does not exist")
	}

	user := c.Locals("user").(models.User)
	commnets.UserID = user.ID
	db.Database.Db.Create(&commnets)

	product_comments := product.Comments
	product_comments = append(product_comments, commnets)
	product.Comments = product_comments
	db.Database.Db.Save(&product)
	return c.Status(fiber.StatusOK).JSON(product)
}

func ReplyOn(c *fiber.Ctx) error {

	var reply models.Comment
	if err := c.BodyParser(&reply); err != nil {
		c.Status(fiber.StatusBadRequest).JSON("wrong input fromat")
	}
	_, err := govalidator.ValidateStruct(reply)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("please input in correct format")
	}
	comment_id, err := c.ParamsInt("commentid")
	var comment models.Comment
	if err := db.Database.Db.Find(&comment, "id = ?", comment_id); err != nil {
		c.Status(fiber.StatusBadRequest).JSON("comment does not exist")
	}

	user := c.Locals("user").(models.User)
	reply.UserID = user.ID
	db.Database.Db.Create(&reply)

	reply_comment:=comment.Reply
	reply_comment = append(reply_comment, reply)
	comment.Reply = reply_comment
	db.Database.Db.Save(&comment)
	return c.Status(fiber.StatusOK).JSON(comment)
}

