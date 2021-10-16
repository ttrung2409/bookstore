package com.bookstore.ecommerce.app.order.query;

public interface OrderQuery {
  Order getOrderToView(String orderId);
}
