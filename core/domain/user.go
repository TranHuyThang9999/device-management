package domain

import "context"

type User struct {
	ID         int64  `json:"id"` // ID người dùng
	UserName   string `json:"user_name" gorm:"uniqueIndex"`
	Password   string `json:"password"`   // Mật khẩu người dùng (nên mã hóa)
	Avatar     string `json:"avatar"`     // Đường dẫn tới avatar
	Age        int    `json:"age"`        // Tuổi người dùng
	Role       int    `json:"role"`       // Vai trò của người dùng
	Department string `json:"department"` // Phòng ban của người dùng
	CreatedAt  int64  `json:"created_at"` // Thời gian tạo (timestamp)
	UpdatedAt  int64  `json:"updated_at"` // Thời gian cập nhật (timestamp)
}

type RepositoryUser interface {
	AddUser(ctx context.Context, req *User) error
	DeleteUserById(ctx context.Context, id int64) error
	UpdateUserById(ctx context.Context, req *User) error
	GetListUser(ctx context.Context, limit, offset int) ([]*User, error)
	GetUserByUserName(ctx context.Context, username string) (*User, error)
}
