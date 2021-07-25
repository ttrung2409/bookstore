package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.CompletableFuture;

import com.bookstore.ecommerce.app.domain.data.Order;

public interface OrderRepository {
  CompletableFuture<Order> get(String id);

  CompletableFuture<Void> create(Order order, Transaction tx) throws Exception;
}
