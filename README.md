# 🌻 Gin-Blog

Gin-Blog 是一款基于 Go 语言和 Vue3 框架的全栈开发项目, 旨在熟练掌握这两种技术的开发流程。

## 💻 项目进度

- 后端: 已完成基于 Go 和 Gin 框架的 API 开发
- 前端: 准备开始 Vue3 项目的开发

### 环境依赖

| 技术  | 版本   |
| ----- | ------ |
| Go    | 1.19   |
| MySQL | >= 8.x |
| Redis | 7.x    |
| Vue   | 3      |

### 待办事项

#### 后端待办

- [ ] Docker 部署
- [ ] 评论 API
- [ ] 用户权限 API
- [ ] 引入 Redis 作为缓存
- [ ] 规范代码风格
- [ ] 补充文档
- [ ] 引入校验器

#### 前端待办

- 开发中...欢迎贡献代码!

## 🗃 项目结构

### 后端结构

```
blog-server
├── cmd
│   └── main.go # 后端入口
├── config.yml
├── go.mod
├── go.sum
└── internal
    ├── global # 全局配置
    │   ├── config.go
    │   ├── keys.go
    │   └── result.go
    ├── handle # API 处理器
    │   ├── article.go
    │   ├── base.go
    │   ├── category.go
    │   ├── login.go
    │   ├── tag.go
    │   ├── user.go
    │   └── view.go
    ├── middleware # 中间件
    │   ├── base.go
    │   ├── cors.go
    │   └── logger.go
    ├── model # ORM 映射
    │   ├── article.go
    │   ├── base.go
    │   ├── category.go
    │   ├── comment.go
    │   ├── config.go
    │   ├── tag.go
    │   ├── user.go
    │   └── view.go
    ├── router # 路由聚合
    │   └── manager.go
    └── utils
        ├── email.go
        ├── encrypt.go
        ├── jwt
        │   └── jwt.go
        └── validator.go
```

### 前端结构

```
开发中...
欢迎贡献代码
```

## 🛠 技术栈

### 后端技术

- Go
- Gin 框架
- GORM + MySQL (后续扩展 PostgreSQL, SQLite)
- Redis (待接入)

### 前端技术

- Vue 3

# License

MIT License

Copyright (c) 2024 Peter Lee
