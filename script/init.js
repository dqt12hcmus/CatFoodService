db = db.getSiblingDB("orders")

db.createUser({
  user: 'orders',
  pwd: 'orders',
  roles: [
    {
      role: 'readWrite',
      db: 'orders',
    },
  ],
});