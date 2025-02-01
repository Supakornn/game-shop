# 🎮 Game Shop API

A modern REST API for managing in-game item transactions, inventory, and virtual currency. Built with Go and clean architecture principles.

## ✨ Key Features

- 🛍️ **Item Shop System**
  - Purchase items with coin validation
  - Sell items back for partial refunds
  - Transaction history tracking
  - Automatic inventory management
- 🎒 **Inventory Management**
  - Track player item ownership
  - Automatic filling on purchase
  - Safe removal on selling
  - Item quantity validation
- 💰 **Virtual Currency**
  - Player coin balance tracking
  - Secure transaction handling
  - Automatic coin deduction/addition
  - Balance validation

## 🏗️ Clean Architecture

The project follows clean architecture with clear separation of concerns:

```
├── pkg/
│   ├── inventory/
│   │   ├── repository/     # Data access layer
│   │   ├── service/        # Business logic
│   │   ├── model/          # Data transfer objects
│   │   └── exception/      # Custom errors
│   ├── itemShop/
│   │   ├── controller/     # HTTP handlers
│   │   ├── service/        # Business logic
│   │   ├── repository/     # Data access
│   │   ├── model/          # DTOs
│   │   └── exception/      # Custom errors
│   └── playerCoin/
│       ├── repository/     # Data access layer
│       ├── service/        # Business logic
│       └── model/          # DTOs
└── test/                   # Integration tests
```

## 🧪 Testing

The service includes comprehensive test coverage using testify mocks:

```go
func TestItemShopService_Buying(t *testing.T) {
    mockItemShopRepo := new(_itemShopRepository.ItemShopRepositoryMock)
    mockPlayerCoinRepo := new(_playerCoinRepository.PlayerCoinRepositoryMock)
    mockInventoryRepo := new(_inventoryRepository.InventoryRepositoryMock)
    logger := echo.New().Logger

    service := _itemShopService.NewItemShopServiceImpl(
        mockItemShopRepo,
        mockPlayerCoinRepo,
        mockInventoryRepo,
        logger,
    )

    t.Run("should buy item successfully", func(t *testing.T) {
        // Test implementation
    })

    t.Run("should fail when not enough coins", func(t *testing.T) {
        // Test implementation
    })
}
```

## 📝 API Documentation

### Authentication

All endpoints require OAuth2 authentication with Google.

### Item Shop Endpoints

#### List Items

```http
GET /v1/item-shop
Query Parameters:
- name (optional): Filter by item name
- description (optional): Filter by description
- page (required): Page number (min: 1)
- size (required): Items per page (min: 1, max: 20)

Response:
{
    "items": [
        {
            "id": 1,
            "name": "Sword",
            "description": "A sword is a bladed melee weapon",
            "picture": "https://example.com/sword.jpg",
            "price": 100
        }
    ],
    "paginate": {
        "page": 1,
        "totalPage": 5
    }
}
```

#### Buy Item

```http
POST /v1/item-shop/buy
Content-Type: application/json

Request:
{
    "itemID": 1,
    "quantity": 2
}

Response:
{
    "id": 123,
    "playerID": "player1",
    "amount": 800,
    "createdAt": "2024-03-15T10:30:00Z"
}
```

#### Sell Item

```http
POST /v1/item-shop/sell
Content-Type: application/json

Request:
{
    "itemID": 1,
    "quantity": 1
}

Response:
{
    "id": 124,
    "playerID": "player1",
    "amount": 850,
    "createdAt": "2024-03-15T10:35:00Z"
}
```

### Inventory Endpoints

#### List Player Inventory

```http
GET /v1/inventory
Authorization: Required

Response:
{
    "items": [
        {
            "item": {
                "id": 1,
                "name": "Sword",
                "description": "A sword is a bladed melee weapon",
                "picture": "https://example.com/sword.jpg",
                "price": 100
            },
            "quantity": 5
        }
    ]
}
```

### Player Coin Endpoints

#### Get Coin Balance

```http
GET /v1/player-coin
Authorization: Required

Response:
{
    "playerID": "player1",
    "coin": 1000
}
```

#### Add Coins

```http
POST /v1/player-coin
Authorization: Required
Content-Type: application/json

Request:
{
    "amount": 100
}

Response:
{
    "id": 125,
    "playerID": "player1",
    "amount": 1100,
    "createdAt": "2024-03-15T10:40:00Z"
}
```

### Admin Endpoints

#### Create Item

```http
POST /v1/item-managing
Authorization: Admin Required
Content-Type: application/json

Request:
{
    "name": "New Sword",
    "description": "A powerful sword",
    "picture": "https://example.com/new-sword.jpg",
    "price": 200
}
```

#### Edit Item

```http
PATCH /v1/item-managing/{itemID}
Authorization: Admin Required
Content-Type: application/json

Request:
{
    "name": "Updated Sword",
    "description": "An updated sword",
    "picture": "https://example.com/updated-sword.jpg",
    "price": 250
}
```

#### Archive Item

```http
DELETE /v1/item-managing/{itemID}
Authorization: Admin Required
```

### Error Responses

```json
400 Bad Request
{
    "error": "Invalid request parameters"
}

401 Unauthorized
{
    "error": "Authentication required"
}

404 Not Found
{
    "error": "Item not found"
}

422 Unprocessable Entity
{
    "error": "coin not enough"
}

500 Internal Server Error
{
    "error": "Internal server error"
}
```

### Authentication Endpoints

#### Player Login

```http
GET /v1/oauth2/google/player/login
```

#### Admin Login

```http
GET /v1/oauth2/google/admin/login
```

#### Logout

```http
DELETE /v1/oauth2/logout
```

## 🛠️ Built With

- [Echo](https://echo.labstack.com/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Testify](https://github.com/stretchr/testify) - Testing toolkit
- [PostgreSQL](https://www.postgresql.org/) - Database

## 👥 Authors

- **Supakorn** - _Initial work_ - [supakornn](https://github.com/supakornn)

⭐️ Star us on GitHub — it helps!
