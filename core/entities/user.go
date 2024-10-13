package entities

type User struct {
	UserName string `json:"user_name"` // Tên người dùng
	// Password   string `json:"password"`   // Mật khẩu người dùn
	Avatar     string `json:"avatar"`     // Đường dẫn tới avatar
	Age        int    `json:"age"`        // Tuổi người dùng
	Department string `json:"department"` // Phòng ban của người dùng
}
type GetUsers struct {
	ID         int64  `json:"id"`         // ID người dùng
	UserName   string `json:"user_name"`  // Tên người dùng
	Avatar     string `json:"avatar"`     // Đường dẫn tới avatar
	Age        int    `json:"age"`        // Tuổi người dùng
	Department string `json:"department"` // Phòng ban của người dùng
	Password   string `json:"password"`
	CreatedAt  int64  `json:"created_at"` // Thời gian tạo (timestamp)
	UpdatedAt  int64  `json:"updated_at"`
}

type UserRequestUpdate struct {
	ID         int64  `json:"id"`
	UserName   string `json:"user_name"`
	Avatar     string `json:"avatar"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}
