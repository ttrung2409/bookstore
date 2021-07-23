package com.bookstore.ecommerce.app.operation.query;

public interface OrderQuery {
  Order getOrderToView(String orderId);
}
