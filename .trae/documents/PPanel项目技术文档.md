# PPanel 项目技术文档

## 1. 项目概述

PPanel 是一个现代化的多用户代理面板工具，专为学习和实际使用而设计的纯净、专业、完美的开源代理面板解决方案。该项目采用 Go 语言开发，基于 Gin 框架构建 RESTful API，支持多种代理协议，提供完整的用户管理、订阅系统、支付集成、节点管理等功能。

### 1.1 主要特性
- **多协议支持**：支持多种代理协议和客户端配置
- **隐私保护**：注重用户隐私和数据安全
- **用户管理**：完整的用户注册、认证、权限管理系统
- **订阅系统**：灵活的订阅产品和套餐管理
- **支付集成**：支持多种支付方式（支付宝、Stripe等）
- **订单管理**：完整的订单生命周期管理
- **节点管理**：服务器节点的统一管理和监控
- **API框架**：基于 go-zero 的高性能 API 框架

### 1.2 技术栈
- **后端语言**：Go 1.21+
- **Web框架**：Gin
- **API框架**：go-zero
- **数据库**：MySQL 8.0+
- **缓存**：Redis 6.0+
- **任务队列**：Asynq
- **认证**：JWT
- **配置管理**：YAML
- **日志**：结构化日志系统

## 2. 项目架构

### 2.1 目录结构
```
PPanel/
├── .github/           # GitHub Actions 配置
├── apis/              # API 定义文件
│   ├── admin/         # 管理员 API
│   ├── app/           # 应用层 API
│   ├── auth/          # 认证相关 API
│   ├── common/        # 公共 API
│   └── public/        # 公开 API
├── cmd/               # 命令行入口
├── doc/               # 文档
├── etc/               # 配置文件
├── initialize/        # 系统初始化
├── internal/          # 内部业务逻辑
│   ├── config/        # 配置管理
│   ├── handler/       # HTTP 处理器
│   ├── logic/         # 业务逻辑
│   ├── middleware/    # 中间件
│   ├── model/         # 数据模型
│   ├── svc/           # 服务上下文
│   └── types/         # 类型定义
├── pkg/               # 公共包
├── queue/             # 异步任务队列
├── scheduler/         # 定时任务调度
└── script/            # 脚本文件
```

### 2.2 系统架构

```mermaid
graph TB
    A[客户端] --> B[Gin Web Server]
    B --> C[认证中间件]
    C --> D[业务处理器]
    D --> E[业务逻辑层]
    E --> F[数据访问层]
    F --> G[MySQL数据库]
    E --> H[Redis缓存]
    E --> I[异步任务队列]
    I --> J[任务处理器]
    K[定时调度器] --> I
```

### 2.3 核心组件

#### 2.3.1 Web服务器
- 基于 Gin 框架的 HTTP 服务器
- 支持 TLS/HTTPS
- 集成 CORS 中间件
- 会话管理（Redis存储）

#### 2.3.2 认证系统
- JWT Token 认证
- 多种登录方式：邮箱、手机、Apple、Google、Telegram
- 会话管理和权限控制
- 管理员权限验证

#### 2.3.3 异步任务系统
- 基于 Asynq 的任务队列
- 支持邮件发送、短信发送、订单处理等异步任务
- 任务重试和错误处理机制

#### 2.3.4 定时调度系统
- 订阅状态检查（每60秒）
- 服务器数据统计（每3分钟）
- 流量重置（每24小时）

## 3. 核心功能模块

### 3.1 用户管理模块

#### 功能特性
- 用户注册、登录、注销
- 多种认证方式支持
- 用户信息管理
- 邀请码系统
- 登录日志记录

#### API接口
- `GET /v1/admin/user/list` - 用户列表
- `GET /v1/admin/user/detail` - 用户详情
- `PUT /v1/admin/user/update` - 更新用户
- `POST /v1/admin/user/create` - 创建用户
- `DELETE /v1/admin/user/delete` - 删除用户

### 3.2 订阅管理模块

#### 功能特性
- 订阅产品管理
- 订阅组管理
- 价格和周期配置
- 订阅状态跟踪

#### API接口
- `GET /v1/admin/subscribe/group/list` - 订阅组列表
- `POST /v1/admin/subscribe/group/create` - 创建订阅组
- `GET /v1/admin/subscribe/list` - 订阅产品列表
- `POST /v1/admin/subscribe/create` - 创建订阅产品

### 3.3 节点管理模块

#### 功能特性
- 服务器节点管理
- 节点组管理
- 规则组配置
- 节点状态监控

#### API接口
- `GET /v1/admin/server/list` - 节点列表
- `POST /v1/admin/server/create` - 创建节点
- `PUT /v1/admin/server/update` - 更新节点
- `DELETE /v1/admin/server/delete` - 删除节点

### 3.4 支付系统模块

#### 功能特性
- 多种支付方式集成
- 订单生成和管理
- 支付回调处理
- 退款处理

#### 支持的支付方式
- 支付宝（Alipay）
- Stripe
- EPay

### 3.5 游戏管理模块

#### 功能特性
- 游戏列表管理
- 游戏详情配置
- 游戏状态控制

#### API接口
- `GET /v1/admin/game/list` - 游戏列表
- `POST /v1/admin/game/create` - 创建游戏
- `PUT /v1/admin/game/update` - 更新游戏
- `DELETE /v1/admin/game/delete` - 删除游戏

## 4. 数据库设计

### 4.1 核心数据表

#### 用户表 (users)
- id: 主键
- password: 密码（加密）
- is_admin: 管理员标识
- refer_code: 邀请码
- created_at: 创建时间
- updated_at: 更新时间

#### 认证方法表 (auth_methods)
- id: 主键
- user_id: 用户ID
- auth_type: 认证类型（email/phone/apple/google/telegram）
- auth_identifier: 认证标识符
- verified: 验证状态

#### 订阅表 (subscriptions)
- id: 主键
- name: 订阅名称
- price: 价格
- duration: 周期
- traffic_limit: 流量限制
- device_limit: 设备限制

#### 服务器表 (servers)
- id: 主键
- name: 服务器名称
- host: 主机地址
- port: 端口
- protocol: 协议类型
- status: 状态

### 4.2 数据库迁移

系统启动时会自动执行数据库迁移，创建必要的表结构和初始数据。

## 5. API设计

### 5.1 API规范

#### 请求格式
- 基础URL: `http://localhost:8080`
- API版本: `/v1`
- 内容类型: `application/json`
- 认证方式: `Authorization: Bearer <token>`

#### 响应格式
```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 5.2 认证流程

1. 用户登录获取 JWT Token
2. 请求头携带 `Authorization: Bearer <token>`
3. 服务器验证 Token 有效性
4. 检查用户权限
5. 返回相应数据

### 5.3 错误处理

#### 常见错误码
- 200: 成功
- 400: 请求参数错误
- 401: 未授权
- 403: 权限不足
- 404: 资源不存在
- 500: 服务器内部错误

## 6. 配置管理

### 6.1 配置文件结构

```yaml
# ppanel.yaml
Host: 0.0.0.0
Port: 8080
Debug: true
TLS:
  Enable: false
  CertFile: ""
  KeyFile: ""
JwtAuth:
  AccessSecret: "your-secret-key"
  AccessExpire: 86400
Logger:
  Level: "info"
  Filename: "logs/ppanel.log"
  MaxSize: 100
  MaxAge: 30
  MaxBackups: 10
MySQL:
  Addr: "localhost:3306"
  Username: "root"
  Password: "password"
  Database: "ppanel"
Redis:
  Host: "localhost:6379"
  Pass: ""
  DB: 0
```

### 6.2 环境变量支持

- `PPANEL_DB`: 数据库连接字符串
- `PPANEL_REDIS`: Redis连接字符串

## 7. 部署指南

### 7.1 源码部署

#### 环境要求
- Go 1.21+
- MySQL 8.0+
- Redis 6.0+

#### 部署步骤

1. **克隆代码**
```bash
git clone https://github.com/your-repo/ppanel.git
cd ppanel
```

2. **安装依赖**
```bash
go mod download
```

3. **配置文件**
```bash
cp etc/ppanel.yaml.example etc/ppanel.yaml
# 编辑配置文件
```

4. **数据库初始化**
```bash
# 创建数据库
mysql -u root -p -e "CREATE DATABASE ppanel;"
```

5. **启动服务**
```bash
go run ppanel.go run
```

### 7.2 Docker部署

#### 使用Docker Compose

```yaml
# docker-compose.yml
version: '3.8'
services:
  ppanel:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PPANEL_DB=mysql://root:password@mysql:3306/ppanel
      - PPANEL_REDIS=redis://redis:6379/0
    depends_on:
      - mysql
      - redis
  
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: ppanel
    volumes:
      - mysql_data:/var/lib/mysql
  
  redis:
    image: redis:6.0-alpine
    volumes:
      - redis_data:/data

volumes:
  mysql_data:
  redis_data:
```

#### 启动命令
```bash
docker-compose up -d
```

### 7.3 生产环境配置

#### 性能优化
- 启用 Gin 生产模式
- 配置适当的日志级别
- 设置合理的数据库连接池
- 配置 Redis 持久化

#### 安全配置
- 使用强密码
- 启用 HTTPS
- 配置防火墙
- 定期备份数据

## 8. 开发指南

### 8.1 开发环境搭建

1. **安装Go环境**
2. **安装开发工具**
   - IDE: GoLand 或 VS Code
   - 数据库工具: MySQL Workbench
   - API测试: Postman

3. **代码规范**
   - 遵循 Go 官方代码规范
   - 使用 gofmt 格式化代码
   - 编写单元测试

### 8.2 添加新功能

1. **定义API**
   - 在 `apis/` 目录下定义API接口
   - 更新类型定义

2. **实现处理器**
   - 在 `internal/handler/` 下实现HTTP处理器
   - 在 `internal/logic/` 下实现业务逻辑

3. **数据模型**
   - 在 `internal/model/` 下定义数据模型
   - 实现数据访问方法

4. **测试**
   - 编写单元测试
   - 进行集成测试

### 8.3 调试技巧

- 使用结构化日志记录关键信息
- 利用 Redis 监控工具查看缓存状态
- 使用数据库慢查询日志优化性能
- 通过 Asynq 监控界面查看任务状态

## 9. 监控与运维

### 9.1 日志管理

- 日志文件位置: `logs/ppanel.log`
- 日志轮转: 按大小和时间自动轮转
- 日志级别: debug, info, warn, error

### 9.2 性能监控

- 数据库连接池监控
- Redis 内存使用监控
- API 响应时间监控
- 任务队列积压监控

### 9.3 备份策略

- 数据库定期备份
- 配置文件备份
- 日志文件归档

## 10. 常见问题

### 10.1 启动问题

**Q: 服务启动失败**
A: 检查配置文件格式、数据库连接、端口占用等

**Q: 数据库连接失败**
A: 验证数据库服务状态、连接参数、网络连通性

### 10.2 认证问题

**Q: JWT Token 验证失败**
A: 检查 Token 格式、过期时间、密钥配置

**Q: 权限不足**
A: 确认用户角色、管理员权限设置

### 10.3 性能问题

**Q: API 响应慢**
A: 检查数据库查询、缓存命中率、网络延迟

**Q: 内存占用高**
A: 检查缓存配置、连接池设置、内存泄漏

## 11. 版本历史

- v1.0.0: 初始版本发布
- 包含基础的用户管理、订阅系统、节点管理功能
- 支持多种认证方式和支付集成

## 12. 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交代码变更
4. 创建 Pull Request
5. 代码审查和合并

## 13. 许可证

本项目采用开源许可证，具体条款请参考 LICENSE 文件。

---

**文档版本**: v1.0  
**最后更新**: 2024年12月  
**维护者**: PPanel 开发团队