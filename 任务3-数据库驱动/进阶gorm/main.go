package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;size:100;not null"`
	PostCount int    `gorm:"default:0"`

	Posts []Post
}

type Post struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"size:200;not null"`
	Content string `gorm:"type:text"`
	State   string `gorm:"size:20;default:'有评论'"`

	UserID uint
	User   User

	Comments []Comment
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`

	PostID uint
	Post   Post
}

func addTestData(db *gorm.DB) {
	user := User{
		Name:  "Alice",
		Email: "111@qq.com",
		Posts: []Post{
			{
				Title:   "First Post",
				Content: "This is the content of the first post.",
				Comments: []Comment{
					{Content: "Great post!"},
					{Content: "Thanks for sharing."},
				},
			},
			{
				Title:   "Second Post",
				Content: "This is the content of the second post.",
				Comments: []Comment{
					{Content: "Interesting read."},
				},
			},
		},
	}
	db.Create(&user)
	fmt.Println("测试数据添加完成")
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Printf("新帖子已创建: %s\n", p.Title)
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		fmt.Printf("帖子ID %d 的评论已全部删除，更新状态为 '无评论'\n", c.PostID)

		err = tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("state", "无评论").Error
		fmt.Println("更新帖子状态完成:", err)
		return err
	}
	return nil
}

func findUserAllPostsAndComments(db *gorm.DB, userID uint) {
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		fmt.Println("查询用户失败:", err)
		return
	}
	fmt.Println("查询用户成功:")
	for _, post := range user.Posts {
		fmt.Printf("Post: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf(" - Comment: %s\n", comment.Content)
		}
	}
}

func findMaxPostComments(db *gorm.DB) {
	type Result struct {
		PostID       uint
		CommentCount int
	}
	var results []Result
	db.Model(&Comment{}).Select("post_id, COUNT(*) as comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Scan(&results)

	if len(results) == 0 {
		fmt.Println("没有评论数据")
		return
	}
	maxCount := results[0].CommentCount

	var posts []Post
	db.Debug().Where("id IN (?)", db.Model(&Comment{}).
		Select("post_id").
		Group("post_id").
		Having("COUNT(*) = ?", maxCount)).
		Find(&posts)
	fmt.Println("评论数最多的帖子:", posts)
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/xxg?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("数据库连接成功")

	// if db.Migrator().HasTable(&User{}) {
	// 	fmt.Println("User 表已存在，跳过创建")
	// } else {
	// 	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// 	fmt.Println("自动迁移完成")
	// }

	// // addTestData(db)

	// findUserAllPostsAndComments(db, 1)
	// findMaxPostComments(db)

	// db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// type UserPostCount struct {
	// 	UserID uint
	// 	Count  int64
	// }
	// var results []UserPostCount
	// db.Model(&Post{}).
	// 	Select(("user_id, COUNT(*) as count")).
	// 	Group("user_id").
	// 	Scan(&results)

	// for _, r := range results {
	// 	db.Model(&User{}).
	// 		Where("id = ?", r.UserID).
	// 		Update("post_count", r.Count)
	// }

	// db.Migrator().DropColumn(&Comment{}, "state")

	// err = db.Model(&Post{}).
	// 	Where("id = ?", 2).
	// 	Update("state", "无评论").Error

	// fmt.Println(err)

	// stmt := db.Session(&gorm.Session{DryRun: true}).
	// 	Where("post_id = ?", 2).
	// 	Delete(&Comment{}).
	// 	Statement
	// fmt.Println(stmt.SQL.String())

	delC := Comment{PostID: 2}
	db.Where("post_id = ?", delC.PostID).Delete(&delC)
}
