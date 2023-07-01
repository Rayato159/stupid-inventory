<h1>Stupid Inventory App</h1>

<h2>👌 Architecture</h2>
<img src="./screenshots/arch.png">

<h2>🍃 Run MongoDB on Docker</h2>

```bash
# start user_db container
docker run --name user_db -e MONGO_INITDB_ROOT_USERNAME=rayato159 -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 2000:27017 -d mongo

# start item_db container
docker run --name item_db -e MONGO_INITDB_ROOT_USERNAME=rayato159 -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 2001:27017 -d mongo
```

<h2>🔌 Connect to a database</h2>

<h3>user_db</h3>

```bash
docker exec -it user_db bash
```

```bash
mongosh mongodb://rayato159:123456@0.0.0.0:27017
```

```bash
use user_db
```

<h3>item_db</h3>

```bash
docker exec -it item_db bash
```

```bash
mongosh mongodb://rayato159:123456@0.0.0.0:27017
```

```bash
use item_db
```

<h2>🪄 Migrate collection</h2>

<h3>items</h3>

```bash
use item_db
```

```bash
db.items.insertMany([
   {
    _id: ObjectId("649f133fce9ca44af3239ca6"),
    title: 'Sword of Fire',
    description: 'A powerful sword engulfed in flames',
    damage: 50
  },
  {
    _id: ObjectId("649f133fce9ca44af3239ca7"),
    title: 'Elven Bow',
    description: 'A beautifully crafted bow with enhanced accuracy',
    damage: 35
  },
  {
    _id: ObjectId("649f133fce9ca44af3239ca8"),
    title: 'Dwarven Axe',
    description: 'A sturdy axe forged by skilled dwarven smiths',
    damage: 45
  },
  {
    _id: ObjectId("649f133fce9ca44af3239ca9"),
    title: 'Wand of Thunder',
    description: 'A magical wand that conjures bolts of lightning',
    damage: 40
  },
  {
    _id: ObjectId("649f133fce9ca44af3239caa"),
    title: 'Staff of Healing',
    description: 'A staff imbued with restorative powers',
    damage: 10
  }
]);
```

<h3>users</h3>

```bash
use user_db
```

```bash
db.users.insertMany([
  {
    _id: ObjectId("649f14208624cc23d3a6e7c5"),
    username: 'john_doe' },
  },
  {
    _id: ObjectId("649f14208624cc23d3a6e7c6"),
    username: 'jane_smith' },
  },
  {
    _id: ObjectId("649f14208624cc23d3a6e7c7"),
    username: 'alex_miller'
  }
]);
```

```bash
db.users_items.insertMany([
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c5"),
    item_id: ObjectId("649f133fce9ca44af3239ca6"),
  },
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c5"),
    item_id: ObjectId("649f133fce9ca44af3239ca7"),
  },
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c5"),
    item_id: ObjectId("649f133fce9ca44af3239ca8"),
  },
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c6"),
    item_id: ObjectId("649f133fce9ca44af3239ca6"),
  },
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c6"),
    item_id: ObjectId("649f133fce9ca44af3239ca7"),
  },
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c7"),
    item_id: ObjectId("649f133fce9ca44af3239ca9"),
  },
  {
    user_id: ObjectId("649f14208624cc23d3a6e7c7"),
    item_id: ObjectId("649f133fce9ca44af3239caa"),
  }
]);
```

<h2>🌞 Generate proto file</h2>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/proto/item/item.proto
```
