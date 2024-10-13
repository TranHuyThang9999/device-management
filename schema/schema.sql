-- Active: 1726675047873@@127.0.0.1@5432@medical_equipment
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    user_name VARCHAR(1024) UNIQUE,
    password VARCHAR(1024),
    avatar VARCHAR(1024),
    age INTEGER,
    address INTEGER,
    role INT,
    department VARCHAR(255),
    created_at BIGINT,
    updated_at BIGINT
);

-- Bảng devices (thiết bị)
CREATE TABLE IF NOT EXISTS devices (
    id BIGINT PRIMARY KEY, -- ID thiết bị
    device_name VARCHAR(255) NOT NULL, -- Tên thiết bị
    description TEXT, -- Mô tả thiết bị
    status INT, -- Tình trạng thiết bị (hoạt động, hỏng, bảo trì,...)
    created_at BIGINT, -- Thời gian tạo thiết bị
    updated_at BIGINT -- Thời gian cập nhật thiết bị
);

-- Bảng linh kiện (components)
CREATE TABLE IF NOT EXISTS components (
    id BIGINT PRIMARY KEY,
    device_id BIGINT REFERENCES devices (id) ON DELETE CASCADE, -- Liên kết với thiết bị
    component_name VARCHAR(255) NOT NULL, -- Tên linh kiện
    serial_number VARCHAR(255) UNIQUE, -- Số seri của linh kiện (nếu có)
    status INT, -- Tình trạng linh kiện
    installed_date BIGINT, -- Ngày lắp đặt
    warranty_expire_date BIGINT, -- Ngày hết hạn bảo hành
    created_at BIGINT,
    updated_at BIGINT
);

-- Bảng maintenance_history (lịch sử bảo trì thiết bị và linh kiện)
CREATE TABLE IF NOT EXISTS maintenance_history (
    id BIGINT PRIMARY KEY, -- ID lịch sử bảo trì
    device_id BIGINT REFERENCES devices (id) ON DELETE CASCADE, -- Liên kết với thiết bị
    user_id BIGINT REFERENCES users (id) ON DELETE SET NULL, -- Người thực hiện bảo trì
    maintenance_date BIGINT NOT NULL, -- Ngày bảo trì
    description TEXT, -- Mô tả quá trình bảo trì
    cost NUMERIC(10, 2), -- Chi phí bảo trì (nếu có)
    next_maintenance_date BIGINT, -- Ngày bảo trì tiếp theo dự kiến (nếu có)
    created_at BIGINT, -- Thời gian tạo bản ghi
    updated_at BIGINT -- Thời gian cập nhật bản ghi
);

-- Bảng maintenance_components (linh kiện được bảo trì)
CREATE TABLE IF NOT EXISTS maintenance_components (
    id BIGINT PRIMARY KEY, -- ID của bản ghi bảo trì linh kiện
    maintenance_id BIGINT REFERENCES maintenance_history (id) ON DELETE CASCADE, -- Liên kết với lịch sử bảo trì
    component_id BIGINT REFERENCES components (id) ON DELETE CASCADE, -- Liên kết với linh kiện
    description TEXT, -- Mô tả quá trình bảo trì của linh kiện
    created_at BIGINT, -- Thời gian tạo bản ghi
    updated_at BIGINT -- Thời gian cập nhật bản ghi
);


CREATE TABLE IF NOT EXISTS file_stores (
    id BIGINT PRIMARY KEY,
    any_id BIGINT,
    url VARCHAR(2048)
)