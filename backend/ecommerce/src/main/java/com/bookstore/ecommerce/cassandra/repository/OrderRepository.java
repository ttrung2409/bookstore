package com.bookstore.ecommerce.cassandra.repository;

import com.bookstore.ecommerce.app.data.Order;

public class OrderRepository
    implements com.bookstore.ecommerce.app.data.repository.OrderRepository {
  @Override
  public Order get(String id) {}

  @Override
  public void create(Order order) {}
}
