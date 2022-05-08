package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.Order;

public interface OrderRepository {
  CompletableFuture<Order> get(String id, Transaction tx) throws Exception;

  CompletableFuture<Void> create(Order order, Transaction tx) throws Exception;

  CompletableFuture<Void> update(Order order, Transaction tx) throws Exception;
}
