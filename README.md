# FxBoard

FxBoard 是一个前后端分离的汇率工具与金融观点社区项目。前端使用 Vue 3 + TypeScript + Element Plus 构建交互界面，后端使用 Go Gin + GORM 提供用户认证、汇率数据、文章内容和点赞统计接口，适合作为全栈开发求职作品进行本地演示。

## 技术栈

### 前端

- Vue 3
- TypeScript
- Vite
- Vue Router
- Pinia
- Axios
- Element Plus

### 后端

- Go 1.22+
- Gin
- GORM
- MySQL
- Redis
- Viper
- JWT
- bcrypt
- CORS

## 功能模块

- 用户注册、登录与退出
- JWT 登录态鉴权
- 汇率列表读取与汇率换算
- 市场观点列表、详情阅读与登录访问控制
- 文章点赞与 Redis 点赞数缓存
- MySQL 持久化用户、文章和汇率数据
- 前端加载态、空状态、错误提示和基础表单校验
- 后端启动时自动迁移数据表，降低首次运行成本

## 项目结构

```text
FxBoard/
├── backend/                   # Go + Gin 后端服务
│   ├── config/                # MySQL、Redis、应用配置
│   ├── controllers/           # 接口控制器
│   ├── global/                # 全局数据库与 Redis 客户端
│   ├── middlewares/           # JWT 鉴权中间件
│   ├── models/                # GORM 数据模型
│   ├── router/                # 路由配置
│   ├── utils/                 # 密码加密、JWT 工具
│   ├── go.mod
│   └── main.go
├── frontend/                  # Vue 3 前端应用
│   ├── src/
│   │   ├── components/        # 登录、注册组件
│   │   ├── router/            # 前端路由
│   │   ├── store/             # Pinia 状态管理
│   │   ├── views/             # 首页、汇率工具、市场观点页面
│   │   └── axios.ts           # Axios 请求封装
│   ├── .env.example
│   ├── package.json
│   └── vite.config.ts
└── README.md
```

## 本地运行

### 1. 准备环境

请先安装并启动：

- Node.js / npm
- Go 1.22+
- MySQL
- Redis

### 2. 准备数据库

默认配置使用 MySQL 的 `test` 数据库。首次运行前可以先创建数据库：

```sql
CREATE DATABASE IF NOT EXISTS test
  DEFAULT CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;
```

后端启动时会通过 GORM 自动创建或更新 `users`、`articles`、`exchange_rates` 等数据表。

### 3. 配置后端

按本地环境修改后端配置：

```yml
# backend/config/config.yml
app:
  name: FxBoard
  port: :3000

database:
  dsn: root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
  MaxIdleConns: 11
  MaxOpenConns: 114

redis:
  addr: localhost:6379
  DB: 0
  Password: ""
```

如果 Go 依赖下载较慢，可以在 PowerShell 中临时设置代理：

```powershell
$env:GOPROXY = "https://goproxy.cn,direct"
```

### 4. 启动后端

```bash
cd backend
go mod download
go run main.go
```

后端默认运行在 `http://localhost:3000`。

### 5. 启动前端

```bash
cd frontend
npm install
npm run dev
```

前端默认运行在 `http://localhost:5173`。接口地址默认是 `http://localhost:3000/api`，也可以复制 `.env.example` 为 `.env.local` 后修改：

```env
VITE_API_BASE_URL=http://localhost:3000/api
```

## 可选初始化数据

如果想让页面一打开就有汇率和市场观点数据，可以在后端首次启动并自动建表后，向 MySQL 写入示例数据：

```sql
INSERT INTO exchange_rates (from_currency, to_currency, rate, date)
VALUES
  ('USD', 'CNY', 7.20, NOW()),
  ('CNY', 'USD', 0.14, NOW()),
  ('EUR', 'CNY', 7.80, NOW()),
  ('CNY', 'EUR', 0.13, NOW());

INSERT INTO articles (created_at, updated_at, title, preview, content)
VALUES
  (NOW(), NOW(), '人民币汇率观察', '关注近期人民币兑美元走势。', '本文用于演示 FxBoard 的观点详情、鉴权访问和点赞功能。'),
  (NOW(), NOW(), '外汇交易基础', '整理常见外汇概念和风险提示。', '外汇交易需要关注汇率波动、流动性、手续费和风险控制。');
```

用户账号可直接通过前端注册页面创建。

## 常用命令

### 前端

```bash
npm run dev       # 启动开发服务
npm run build     # 类型检查并构建生产包
npm run preview   # 预览生产构建结果
```

### 后端

```bash
go run main.go    # 启动后端服务
go test ./...     # 编译并运行后端测试
go mod download   # 下载 Go 依赖
```

## 接口概览

| 方法 | 路径 | 说明 | 鉴权 |
| --- | --- | --- | --- |
| POST | `/api/auth/register` | 用户注册 | 否 |
| POST | `/api/auth/login` | 用户登录 | 否 |
| GET | `/api/exchangeRates` | 获取汇率列表 | 否 |
| POST | `/api/exchangeRates` | 新增汇率 | 是 |
| POST | `/api/articles` | 发布观点文章 | 是 |
| GET | `/api/articles` | 获取观点列表 | 是 |
| GET | `/api/articles/:id` | 获取观点详情 | 是 |
| POST | `/api/articles/:id/like` | 点赞观点 | 是 |
| GET | `/api/articles/:id/like` | 获取点赞数 | 是 |

需要鉴权的接口会读取请求头中的 `Authorization` 字段，前端登录后会自动把 token 写入请求头。

## 简历描述示例

FxBoard：基于 Vue 3、TypeScript、Element Plus、Go Gin、GORM、MySQL 和 Redis 实现的前后端分离汇率工具与金融观点社区，支持用户注册登录、JWT 鉴权、汇率查询换算、观点文章浏览和 Redis 点赞统计。项目采用 `frontend` / `backend` 的清晰目录结构，补充了自动建表、接口错误处理、前端加载/空状态、表单校验和本地运行文档，可稳定完成本地演示。

## 可讲技术亮点

- 使用 Pinia 管理登录态，并通过 Axios 拦截器统一携带 JWT。
- 后端使用 Gin 路由分组和鉴权中间件保护文章、点赞和新增汇率接口。
- 使用 GORM 自动迁移数据表，降低首次部署和本地演示成本。
- 使用 Redis 缓存文章列表和文章点赞数，展示缓存型业务场景。
- 通过清晰的 `frontend` / `backend` 目录和统一命名，减少教程项目痕迹，提升作品展示感。
