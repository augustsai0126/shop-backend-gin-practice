# Shopping Website Practice - Go Gin

這是一個購物網站的後端API專案，目標是練習使用 Golang 的 [Gin](https://github.com/gin-gonic/gin) 框架，並搭配 PostgreSQL 資料庫，實作一個具備完整電商功能的 RESTful API。

## 專案目標

- 學習 Golang 與 Gin 開發API
- 熟悉 PostgreSQL 資料庫設計與操作
- 體驗購物網站的後端架構
- 預計未來搭配 React 實作前端頁面
- 最終將專案部署到 AWS，並使用 Cloudflare 作為 CDN

## 主要功能

- 會員註冊 / 登入（JWT驗證）
- 商品查詢與管理
- 商品分類
- 購物車功能
- 訂單管理
- 優惠活動管理
- 基本的 RBAC 權限控管（進階目標）

## 技術棧

- Golang (Gin)
- PostgreSQL
- JWT
- Docker（後續將加入）
- AWS EC2 / RDS（後續部署）
- Cloudflare（後續 CDN）

## 使用說明

1. 安裝必要套件
    ```bash
    go mod tidy
    ```
2. 建立 `.env` 設定資料庫連線資訊

3. 執行主程式
    ```bash
    go run main.go
    ```

## 備註

此專案僅作為學習用途，歡迎任何意見與建議！

---

