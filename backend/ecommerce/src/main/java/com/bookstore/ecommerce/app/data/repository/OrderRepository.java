package com.bookstore.ecommerce.app.data.repository;

import com.bookstore.ecommerce.app.data.Order;

public interface OrderRepository {
  Order get(String id);

  void create(Order order);
}
